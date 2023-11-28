package e2e

// Copyright (c) Microsoft Corporation.
// Licensed under the Apache License 2.0.

import (
	"context"
	"math/rand"
	"strconv"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	mgmtnetwork "github.com/Azure/azure-sdk-for-go/services/network/mgmt/2020-08-01/network"
	"github.com/Azure/go-autorest/autorest/to"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	mgmtredhatopenshift20220904 "github.com/Azure/ARO-RP/pkg/client/services/redhatopenshift/mgmt/2022-09-04/redhatopenshift"
	mgmtredhatopenshift20230701preview "github.com/Azure/ARO-RP/pkg/client/services/redhatopenshift/mgmt/2023-07-01-preview/redhatopenshift"
	"github.com/Azure/ARO-RP/pkg/util/stringutils"
)

var _ = Describe("Update clusters", func() {
	It("must restart the aro-operator-master Deployment", func(ctx context.Context) {
		By("saving the current revision of the aro-operator-master Deployment")
		d, err := clients.Kubernetes.AppsV1().Deployments("openshift-azure-operator").Get(ctx, "aro-operator-master", metav1.GetOptions{})
		Expect(err).NotTo(HaveOccurred())

		Expect(d.ObjectMeta.Annotations).To(HaveKey("deployment.kubernetes.io/revision"))

		oldRevision, err := strconv.Atoi(d.ObjectMeta.Annotations["deployment.kubernetes.io/revision"])
		Expect(err).NotTo(HaveOccurred())

		By("sending the PATCH request to update the cluster")
		err = clients.OpenshiftClusters.UpdateAndWait(ctx, vnetResourceGroup, clusterName, mgmtredhatopenshift20220904.OpenShiftClusterUpdate{})
		Expect(err).NotTo(HaveOccurred())

		By("checking that the aro-operator-master Deployment was restarted")
		d, err = clients.Kubernetes.AppsV1().Deployments("openshift-azure-operator").Get(ctx, "aro-operator-master", metav1.GetOptions{})
		Expect(err).NotTo(HaveOccurred())

		Expect(d.Spec.Template.Annotations).To(HaveKey("kubectl.kubernetes.io/restartedAt"))

		Expect(d.ObjectMeta.Annotations).To(HaveKey("deployment.kubernetes.io/revision"))

		newRevision, err := strconv.Atoi(d.ObjectMeta.Annotations["deployment.kubernetes.io/revision"])
		Expect(err).NotTo(HaveOccurred())
		Expect(newRevision).To(Equal(oldRevision + 1))
	})

	// This tests the API which is most commonly generated by
	// az resource tag --tags key=value --ids /subscriptions/xxx/resourceGroups/xxx/providers/Microsoft.RedHatOpenShift/openShiftClusters/xxx
	It("must be possible to set tags on a cluster resource via PUT", func(ctx context.Context) {
		value := strconv.Itoa(rand.Int())

		By("getting cluster resource")
		oc, err := clients.OpenshiftClusters.Get(ctx, vnetResourceGroup, clusterName)
		Expect(err).NotTo(HaveOccurred())
		Expect(oc.Tags).NotTo(HaveKeyWithValue("key", &value))

		By("adding a new test tag")
		if oc.Tags == nil {
			oc.Tags = map[string]*string{}
		}
		oc.Tags["key"] = &value

		By("sending the PUT request to update the resource")
		err = clients.OpenshiftClusters.CreateOrUpdateAndWait(ctx, vnetResourceGroup, clusterName, oc)
		Expect(err).NotTo(HaveOccurred())

		By("getting the cluster resource")
		oc, err = clients.OpenshiftClusters.Get(ctx, vnetResourceGroup, clusterName)
		Expect(err).NotTo(HaveOccurred())

		By("checking that the tag has expected value")
		Expect(oc.Tags).To(HaveKeyWithValue("key", &value))
	})
})

var _ = Describe("Update cluster Managed Outbound IPs", func() {
	var lbName string
	var rgName string

	var _ = BeforeEach(func(ctx context.Context) {
		By("ensuring the public loadbalancer starts with one outbound IP")
		oc, err := clients.OpenshiftClustersPreview.Get(ctx, vnetResourceGroup, clusterName)
		Expect(err).NotTo(HaveOccurred())

		lbName, err = getPublicLoadBalancerName(ctx)
		Expect(err).NotTo(HaveOccurred())

		rgName = stringutils.LastTokenByte(*oc.ClusterProfile.ResourceGroupID, '/')
		lb, err := clients.LoadBalancers.Get(ctx, rgName, lbName, "")
		Expect(err).NotTo(HaveOccurred())

		if getOutboundIPsCount(lb) != 1 {
			By("sending the PATCH request to set ManagedOutboundIPs.Count to 1")
			err = clients.OpenshiftClustersPreview.UpdateAndWait(ctx, vnetResourceGroup, clusterName, newManagedOutboundIPUpdateBody(1))
			Expect(err).NotTo(HaveOccurred())
		}
	})

	It("must be possible to increase and decrease IP Addresses on the public loadbalancer", func(ctx context.Context) {
		By("sending the PATCH request to increase Managed Outbound IPs")
		err := clients.OpenshiftClustersPreview.UpdateAndWait(ctx, vnetResourceGroup, clusterName, newManagedOutboundIPUpdateBody(5))
		Expect(err).NotTo(HaveOccurred())

		By("getting the cluster resource")
		oc, err := clients.OpenshiftClustersPreview.Get(ctx, vnetResourceGroup, clusterName)
		Expect(err).NotTo(HaveOccurred())

		By("checking effectiveOutboundIPs has been updated")
		Expect(*oc.NetworkProfile.LoadBalancerProfile.EffectiveOutboundIps).To(HaveLen(5))

		By("checking outbound-rule-4 has required number IPs")
		lb, err := clients.LoadBalancers.Get(ctx, rgName, lbName, "")
		Expect(err).NotTo(HaveOccurred())
		Expect(getOutboundIPsCount(lb)).To(Equal(5))

		By("sending the PUT request to decrease Managed Outbound IPs")
		oc.OpenShiftClusterProperties.NetworkProfile.LoadBalancerProfile.ManagedOutboundIps.Count = to.Int32Ptr(1)
		err = clients.OpenshiftClustersPreview.CreateOrUpdateAndWait(ctx, vnetResourceGroup, clusterName, oc)
		Expect(err).NotTo(HaveOccurred())

		By("getting the cluster resource")
		oc, err = clients.OpenshiftClustersPreview.Get(ctx, vnetResourceGroup, clusterName)
		Expect(err).NotTo(HaveOccurred())

		By("checking effectiveOutboundIPs has been updated")
		Expect(*oc.NetworkProfile.LoadBalancerProfile.EffectiveOutboundIps).To(HaveLen(1))

		By("checking outbound-rule-4 has required number of IPs")
		lb, err = clients.LoadBalancers.Get(ctx, rgName, lbName, "")
		Expect(err).NotTo(HaveOccurred())
		Expect(getOutboundIPsCount(lb)).To(Equal(1))
	})
})

func getPublicLoadBalancerName(ctx context.Context) (string, error) {
	co, err := clients.AROClusters.AroV1alpha1().Clusters().Get(ctx, "cluster", metav1.GetOptions{})
	if err != nil {
		return "", err
	}
	return co.Spec.InfraID, err
}

func newManagedOutboundIPUpdateBody(managedOutboundIPCount int32) mgmtredhatopenshift20230701preview.OpenShiftClusterUpdate {
	return mgmtredhatopenshift20230701preview.OpenShiftClusterUpdate{
		OpenShiftClusterProperties: &mgmtredhatopenshift20230701preview.OpenShiftClusterProperties{
			NetworkProfile: &mgmtredhatopenshift20230701preview.NetworkProfile{
				LoadBalancerProfile: &mgmtredhatopenshift20230701preview.LoadBalancerProfile{
					ManagedOutboundIps: &mgmtredhatopenshift20230701preview.ManagedOutboundIPs{
						Count: to.Int32Ptr(managedOutboundIPCount),
					},
				},
			},
		},
	}
}

func getOutboundIPsCount(lb mgmtnetwork.LoadBalancer) int {
	numOfIPs := 0
	for _, obRule := range *lb.LoadBalancerPropertiesFormat.OutboundRules {
		if *obRule.Name == "outbound-rule-v4" {
			numOfIPs = len(*obRule.FrontendIPConfigurations)
		}
	}
	return numOfIPs
}
