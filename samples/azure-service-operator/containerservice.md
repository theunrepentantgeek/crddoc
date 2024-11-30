v1api20210501
=============

| Metadata             | Value                                                                         |
|----------------------|-------------------------------------------------------------------------------|
| Group                | containerservice.azure.com                                                    |
| Version              | v1api20210501                                                                 |
| Module               | github.com/Azure/azure-service-operator/v2/api/containerservice/v1api20210501 |
| Property Optionality | Optional                                                                      |

<a id="APIVersion"></a>APIVersion
---------------------------------

| Value        | Description |
|--------------|-------------|
| "2021-05-01" |             |

<a id="ManagedCluster"></a>ManagedCluster
-----------------------------------------

Generator information: - Generated from: /containerservice/resource-manager/Microsoft.ContainerService/aks/stable/2021-05-01/managedClusters.json - ARM URI: /&ZeroWidthSpace;subscriptions/&ZeroWidthSpace;{subscriptionId}/&ZeroWidthSpace;resourceGroups/&ZeroWidthSpace;{resourceGroupName}/&ZeroWidthSpace;providers/&ZeroWidthSpace;Microsoft.ContainerService/managedClusters/{resourceName}

Used by: [ManagedClusterList](#ManagedClusterList).

| Property          | Description | Type                                                                        |
|-------------------|-------------|-----------------------------------------------------------------------------|
| metav1.TypeMeta   |             |                                                                             |
| metav1.ObjectMeta |             |                                                                             |
| spec              |             | [ManagedCluster_Spec](#ManagedCluster_Spec)<br/><small>Optional</small>     |
| status            |             | [ManagedCluster_STATUS](#ManagedCluster_STATUS)<br/><small>Optional</small> |

### <a id="ManagedCluster_Spec"></a>ManagedCluster_Spec

| Property                     | Description                                                                                                                                                                                                                                                                                                                                                                                 | Type                                                                                                                  |
|------------------------------|---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|-----------------------------------------------------------------------------------------------------------------------|
| aadProfile                   | The Azure Active Directory configuration.                                                                                                                                                                                                                                                                                                                                                   | [ManagedClusterAADProfile](#ManagedClusterAADProfile)<br/><small>Optional</small>                                     |
| addonProfiles                | The profile of managed cluster add-on.                                                                                                                                                                                                                                                                                                                                                      | [map[string]ManagedClusterAddonProfile](#ManagedClusterAddonProfile)<br/><small>Optional</small>                      |
| agentPoolProfiles            | The agent pool properties.                                                                                                                                                                                                                                                                                                                                                                  | [ManagedClusterAgentPoolProfile[]](#ManagedClusterAgentPoolProfile)<br/><small>Optional</small>                       |
| apiServerAccessProfile       | The access profile for managed cluster API server.                                                                                                                                                                                                                                                                                                                                          | [ManagedClusterAPIServerAccessProfile](#ManagedClusterAPIServerAccessProfile)<br/><small>Optional</small>             |
| autoScalerProfile            | Parameters to be applied to the cluster-autoscaler when enabled                                                                                                                                                                                                                                                                                                                             | [ManagedClusterProperties_AutoScalerProfile](#ManagedClusterProperties_AutoScalerProfile)<br/><small>Optional</small> |
| autoUpgradeProfile           | The auto upgrade configuration.                                                                                                                                                                                                                                                                                                                                                             | [ManagedClusterAutoUpgradeProfile](#ManagedClusterAutoUpgradeProfile)<br/><small>Optional</small>                     |
| azureName                    | The name of the resource in Azure. This is often the same as the name of the resource in Kubernetes but it doesn't have to be.                                                                                                                                                                                                                                                              | string<br/><small>Optional</small>                                                                                    |
| disableLocalAccounts         | If set to true, getting static credentials will be disabled for this cluster. This must only be used on Managed Clusters that are AAD enabled. For more details see [disable local accounts](https://docs.microsoft.com/azure/aks/managed-aad#disable-local-accounts-preview).                                                                                                              | bool<br/><small>Optional</small>                                                                                      |
| diskEncryptionSetIDReference | This is of the form: '/&ZeroWidthSpace;subscriptions/&ZeroWidthSpace;{subscriptionId}/&ZeroWidthSpace;resourceGroups/&ZeroWidthSpace;{resourceGroupName}/&ZeroWidthSpace;providers/&ZeroWidthSpace;Microsoft.Compute/diskEncryptionSets/{encryptionSetName}'                                                                                                                                | genruntime.ResourceReference<br/><small>Optional</small>                                                              |
| dnsPrefix                    | This cannot be updated once the Managed Cluster has been created.                                                                                                                                                                                                                                                                                                                           | string<br/><small>Optional</small>                                                                                    |
| enablePodSecurityPolicy      | (DEPRECATING) Whether to enable Kubernetes pod security policy (preview). This feature is set for removal on October 15th, 2020. Learn more at aka.ms/aks/azpodpolicy.                                                                                                                                                                                                                      | bool<br/><small>Optional</small>                                                                                      |
| enableRBAC                   | Whether to enable Kubernetes Role-Based Access Control.                                                                                                                                                                                                                                                                                                                                     | bool<br/><small>Optional</small>                                                                                      |
| extendedLocation             | The extended location of the Virtual Machine.                                                                                                                                                                                                                                                                                                                                               | [ExtendedLocation](#ExtendedLocation)<br/><small>Optional</small>                                                     |
| fqdnSubdomain                | This cannot be updated once the Managed Cluster has been created.                                                                                                                                                                                                                                                                                                                           | string<br/><small>Optional</small>                                                                                    |
| httpProxyConfig              | Configurations for provisioning the cluster with HTTP proxy servers.                                                                                                                                                                                                                                                                                                                        | [ManagedClusterHTTPProxyConfig](#ManagedClusterHTTPProxyConfig)<br/><small>Optional</small>                           |
| identity                     | The identity of the managed cluster, if configured.                                                                                                                                                                                                                                                                                                                                         | [ManagedClusterIdentity](#ManagedClusterIdentity)<br/><small>Optional</small>                                         |
| identityProfile              | Identities associated with the cluster.                                                                                                                                                                                                                                                                                                                                                     | [map[string]UserAssignedIdentity](#UserAssignedIdentity)<br/><small>Optional</small>                                  |
| kubernetesVersion            | When you upgrade a supported AKS cluster, Kubernetes minor versions cannot be skipped. All upgrades must be performed sequentially by major version number. For example, upgrades between 1.14.x -> 1.15.x or 1.15.x -> 1.16.x are allowed, however 1.14.x -> 1.16.x is not allowed. See [upgrading an AKS cluster](https://docs.microsoft.com/azure/aks/upgrade-cluster) for more details. | string<br/><small>Optional</small>                                                                                    |
| linuxProfile                 | The profile for Linux VMs in the Managed Cluster.                                                                                                                                                                                                                                                                                                                                           | [ContainerServiceLinuxProfile](#ContainerServiceLinuxProfile)<br/><small>Optional</small>                             |
| location                     | Resource location                                                                                                                                                                                                                                                                                                                                                                           | string<br/><small>Required</small>                                                                                    |
| networkProfile               | The network configuration profile.                                                                                                                                                                                                                                                                                                                                                          | [ContainerServiceNetworkProfile](#ContainerServiceNetworkProfile)<br/><small>Optional</small>                         |
| nodeResourceGroup            | The name of the resource group containing agent pool nodes.                                                                                                                                                                                                                                                                                                                                 | string<br/><small>Optional</small>                                                                                    |
| operatorSpec                 | The specification for configuring operator behavior. This field is interpreted by the operator and not passed directly to Azure                                                                                                                                                                                                                                                             | [ManagedClusterOperatorSpec](#ManagedClusterOperatorSpec)<br/><small>Optional</small>                                 |
| owner                        | The owner of the resource. The owner controls where the resource goes when it is deployed. The owner also controls the resources lifecycle. When the owner is deleted the resource will also be deleted. Owner is expected to be a reference to a resources.azure.com/ResourceGroup resource                                                                                                | genruntime.KnownResourceReference<br/><small>Required</small>                                                         |
| podIdentityProfile           | See [use AAD pod identity](https://docs.microsoft.com/azure/aks/use-azure-ad-pod-identity) for more details on AAD pod identity integration.                                                                                                                                                                                                                                                | [ManagedClusterPodIdentityProfile](#ManagedClusterPodIdentityProfile)<br/><small>Optional</small>                     |
| privateLinkResources         | Private link resources associated with the cluster.                                                                                                                                                                                                                                                                                                                                         | [PrivateLinkResource[]](#PrivateLinkResource)<br/><small>Optional</small>                                             |
| servicePrincipalProfile      | Information about a service principal identity for the cluster to use for manipulating Azure APIs.                                                                                                                                                                                                                                                                                          | [ManagedClusterServicePrincipalProfile](#ManagedClusterServicePrincipalProfile)<br/><small>Optional</small>           |
| sku                          | The managed cluster SKU.                                                                                                                                                                                                                                                                                                                                                                    | [ManagedClusterSKU](#ManagedClusterSKU)<br/><small>Optional</small>                                                   |
| tags                         | Resource tags                                                                                                                                                                                                                                                                                                                                                                               | map[string]string<br/><small>Optional</small>                                                                         |
| windowsProfile               | The profile for Windows VMs in the Managed Cluster.                                                                                                                                                                                                                                                                                                                                         | [ManagedClusterWindowsProfile](#ManagedClusterWindowsProfile)<br/><small>Optional</small>                             |

### <a id="ManagedCluster_STATUS"></a>ManagedCluster_STATUS

| Property                | Description                                                                                                                                                                                                                                                                                                                                                                                 | Type                                                                                                                                |
|-------------------------|---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|-------------------------------------------------------------------------------------------------------------------------------------|
| aadProfile              | The Azure Active Directory configuration.                                                                                                                                                                                                                                                                                                                                                   | [ManagedClusterAADProfile_STATUS](#ManagedClusterAADProfile_STATUS)<br/><small>Optional</small>                                     |
| addonProfiles           | The profile of managed cluster add-on.                                                                                                                                                                                                                                                                                                                                                      | [map[string]ManagedClusterAddonProfile_STATUS](#ManagedClusterAddonProfile_STATUS)<br/><small>Optional</small>                      |
| agentPoolProfiles       | The agent pool properties.                                                                                                                                                                                                                                                                                                                                                                  | [ManagedClusterAgentPoolProfile_STATUS[]](#ManagedClusterAgentPoolProfile_STATUS)<br/><small>Optional</small>                       |
| apiServerAccessProfile  | The access profile for managed cluster API server.                                                                                                                                                                                                                                                                                                                                          | [ManagedClusterAPIServerAccessProfile_STATUS](#ManagedClusterAPIServerAccessProfile_STATUS)<br/><small>Optional</small>             |
| autoScalerProfile       | Parameters to be applied to the cluster-autoscaler when enabled                                                                                                                                                                                                                                                                                                                             | [ManagedClusterProperties_AutoScalerProfile_STATUS](#ManagedClusterProperties_AutoScalerProfile_STATUS)<br/><small>Optional</small> |
| autoUpgradeProfile      | The auto upgrade configuration.                                                                                                                                                                                                                                                                                                                                                             | [ManagedClusterAutoUpgradeProfile_STATUS](#ManagedClusterAutoUpgradeProfile_STATUS)<br/><small>Optional</small>                     |
| azurePortalFQDN         | The Azure Portal requires certain Cross-Origin Resource Sharing (CORS) headers to be sent in some responses, which Kubernetes APIServer doesn't handle by default. This special FQDN supports CORS, allowing the Azure Portal to function properly.                                                                                                                                         | string<br/><small>Optional</small>                                                                                                  |
| conditions              | The observed state of the resource                                                                                                                                                                                                                                                                                                                                                          | conditions.Condition[]<br/><small>Optional</small>                                                                                  |
| disableLocalAccounts    | If set to true, getting static credentials will be disabled for this cluster. This must only be used on Managed Clusters that are AAD enabled. For more details see [disable local accounts](https://docs.microsoft.com/azure/aks/managed-aad#disable-local-accounts-preview).                                                                                                              | bool<br/><small>Optional</small>                                                                                                    |
| diskEncryptionSetID     | This is of the form: '/&ZeroWidthSpace;subscriptions/&ZeroWidthSpace;{subscriptionId}/&ZeroWidthSpace;resourceGroups/&ZeroWidthSpace;{resourceGroupName}/&ZeroWidthSpace;providers/&ZeroWidthSpace;Microsoft.Compute/diskEncryptionSets/{encryptionSetName}'                                                                                                                                | string<br/><small>Optional</small>                                                                                                  |
| dnsPrefix               | This cannot be updated once the Managed Cluster has been created.                                                                                                                                                                                                                                                                                                                           | string<br/><small>Optional</small>                                                                                                  |
| enablePodSecurityPolicy | (DEPRECATING) Whether to enable Kubernetes pod security policy (preview). This feature is set for removal on October 15th, 2020. Learn more at aka.ms/aks/azpodpolicy.                                                                                                                                                                                                                      | bool<br/><small>Optional</small>                                                                                                    |
| enableRBAC              | Whether to enable Kubernetes Role-Based Access Control.                                                                                                                                                                                                                                                                                                                                     | bool<br/><small>Optional</small>                                                                                                    |
| extendedLocation        | The extended location of the Virtual Machine.                                                                                                                                                                                                                                                                                                                                               | [ExtendedLocation_STATUS](#ExtendedLocation_STATUS)<br/><small>Optional</small>                                                     |
| fqdn                    | The FQDN of the master pool.                                                                                                                                                                                                                                                                                                                                                                | string<br/><small>Optional</small>                                                                                                  |
| fqdnSubdomain           | This cannot be updated once the Managed Cluster has been created.                                                                                                                                                                                                                                                                                                                           | string<br/><small>Optional</small>                                                                                                  |
| httpProxyConfig         | Configurations for provisioning the cluster with HTTP proxy servers.                                                                                                                                                                                                                                                                                                                        | [ManagedClusterHTTPProxyConfig_STATUS](#ManagedClusterHTTPProxyConfig_STATUS)<br/><small>Optional</small>                           |
| id                      | Resource Id                                                                                                                                                                                                                                                                                                                                                                                 | string<br/><small>Optional</small>                                                                                                  |
| identity                | The identity of the managed cluster, if configured.                                                                                                                                                                                                                                                                                                                                         | [ManagedClusterIdentity_STATUS](#ManagedClusterIdentity_STATUS)<br/><small>Optional</small>                                         |
| identityProfile         | Identities associated with the cluster.                                                                                                                                                                                                                                                                                                                                                     | [map[string]UserAssignedIdentity_STATUS](#UserAssignedIdentity_STATUS)<br/><small>Optional</small>                                  |
| kubernetesVersion       | When you upgrade a supported AKS cluster, Kubernetes minor versions cannot be skipped. All upgrades must be performed sequentially by major version number. For example, upgrades between 1.14.x -> 1.15.x or 1.15.x -> 1.16.x are allowed, however 1.14.x -> 1.16.x is not allowed. See [upgrading an AKS cluster](https://docs.microsoft.com/azure/aks/upgrade-cluster) for more details. | string<br/><small>Optional</small>                                                                                                  |
| linuxProfile            | The profile for Linux VMs in the Managed Cluster.                                                                                                                                                                                                                                                                                                                                           | [ContainerServiceLinuxProfile_STATUS](#ContainerServiceLinuxProfile_STATUS)<br/><small>Optional</small>                             |
| location                | Resource location                                                                                                                                                                                                                                                                                                                                                                           | string<br/><small>Optional</small>                                                                                                  |
| maxAgentPools           | The max number of agent pools for the managed cluster.                                                                                                                                                                                                                                                                                                                                      | int<br/><small>Optional</small>                                                                                                     |
| name                    | Resource name                                                                                                                                                                                                                                                                                                                                                                               | string<br/><small>Optional</small>                                                                                                  |
| networkProfile          | The network configuration profile.                                                                                                                                                                                                                                                                                                                                                          | [ContainerServiceNetworkProfile_STATUS](#ContainerServiceNetworkProfile_STATUS)<br/><small>Optional</small>                         |
| nodeResourceGroup       | The name of the resource group containing agent pool nodes.                                                                                                                                                                                                                                                                                                                                 | string<br/><small>Optional</small>                                                                                                  |
| podIdentityProfile      | See [use AAD pod identity](https://docs.microsoft.com/azure/aks/use-azure-ad-pod-identity) for more details on AAD pod identity integration.                                                                                                                                                                                                                                                | [ManagedClusterPodIdentityProfile_STATUS](#ManagedClusterPodIdentityProfile_STATUS)<br/><small>Optional</small>                     |
| powerState              | The Power State of the cluster.                                                                                                                                                                                                                                                                                                                                                             | [PowerState_STATUS](#PowerState_STATUS)<br/><small>Optional</small>                                                                 |
| privateFQDN             | The FQDN of private cluster.                                                                                                                                                                                                                                                                                                                                                                | string<br/><small>Optional</small>                                                                                                  |
| privateLinkResources    | Private link resources associated with the cluster.                                                                                                                                                                                                                                                                                                                                         | [PrivateLinkResource_STATUS[]](#PrivateLinkResource_STATUS)<br/><small>Optional</small>                                             |
| provisioningState       | The current provisioning state.                                                                                                                                                                                                                                                                                                                                                             | string<br/><small>Optional</small>                                                                                                  |
| servicePrincipalProfile | Information about a service principal identity for the cluster to use for manipulating Azure APIs.                                                                                                                                                                                                                                                                                          | [ManagedClusterServicePrincipalProfile_STATUS](#ManagedClusterServicePrincipalProfile_STATUS)<br/><small>Optional</small>           |
| sku                     | The managed cluster SKU.                                                                                                                                                                                                                                                                                                                                                                    | [ManagedClusterSKU_STATUS](#ManagedClusterSKU_STATUS)<br/><small>Optional</small>                                                   |
| tags                    | Resource tags                                                                                                                                                                                                                                                                                                                                                                               | map[string]string<br/><small>Optional</small>                                                                                       |
| type                    | Resource type                                                                                                                                                                                                                                                                                                                                                                               | string<br/><small>Optional</small>                                                                                                  |
| windowsProfile          | The profile for Windows VMs in the Managed Cluster.                                                                                                                                                                                                                                                                                                                                         | [ManagedClusterWindowsProfile_STATUS](#ManagedClusterWindowsProfile_STATUS)<br/><small>Optional</small>                             |

<a id="ManagedClusterList"></a>ManagedClusterList
-------------------------------------------------

Generator information: - Generated from: /containerservice/resource-manager/Microsoft.ContainerService/aks/stable/2021-05-01/managedClusters.json - ARM URI: /&ZeroWidthSpace;subscriptions/&ZeroWidthSpace;{subscriptionId}/&ZeroWidthSpace;resourceGroups/&ZeroWidthSpace;{resourceGroupName}/&ZeroWidthSpace;providers/&ZeroWidthSpace;Microsoft.ContainerService/managedClusters/{resourceName}

| Property        | Description | Type                                                            |
|-----------------|-------------|-----------------------------------------------------------------|
| metav1.TypeMeta |             |                                                                 |
| metav1.ListMeta |             |                                                                 |
| items           |             | [ManagedCluster[]](#ManagedCluster)<br/><small>Optional</small> |

<a id="ManagedClustersAgentPool"></a>ManagedClustersAgentPool
-------------------------------------------------------------

Generator information: - Generated from: /containerservice/resource-manager/Microsoft.ContainerService/aks/stable/2021-05-01/managedClusters.json - ARM URI: /&ZeroWidthSpace;subscriptions/&ZeroWidthSpace;{subscriptionId}/&ZeroWidthSpace;resourceGroups/&ZeroWidthSpace;{resourceGroupName}/&ZeroWidthSpace;providers/&ZeroWidthSpace;Microsoft.ContainerService/managedClusters/{resourceName}/agentPools/{agentPoolName}

Used by: [ManagedClustersAgentPoolList](#ManagedClustersAgentPoolList).

| Property          | Description | Type                                                                                              |
|-------------------|-------------|---------------------------------------------------------------------------------------------------|
| metav1.TypeMeta   |             |                                                                                                   |
| metav1.ObjectMeta |             |                                                                                                   |
| spec              |             | [ManagedClusters_AgentPool_Spec](#ManagedClusters_AgentPool_Spec)<br/><small>Optional</small>     |
| status            |             | [ManagedClusters_AgentPool_STATUS](#ManagedClusters_AgentPool_STATUS)<br/><small>Optional</small> |

### <a id="ManagedClusters_AgentPool_Spec"></a>ManagedClusters_AgentPool_Spec

| Property                      | Description                                                                                                                                                                                                                                                                                                                                                                                                                                                                                  | Type                                                                              |
|-------------------------------|----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|-----------------------------------------------------------------------------------|
| availabilityZones             | The list of Availability zones to use for nodes. This can only be specified if the AgentPoolType property is 'VirtualMachineScaleSets'.                                                                                                                                                                                                                                                                                                                                                      | string[]<br/><small>Optional</small>                                              |
| azureName                     | The name of the resource in Azure. This is often the same as the name of the resource in Kubernetes but it doesn't have to be.                                                                                                                                                                                                                                                                                                                                                               | string<br/><small>Optional</small>                                                |
| count                         | Number of agents (VMs) to host docker containers. Allowed values must be in the range of 0 to 1000 (inclusive) for user pools and in the range of 1 to 1000 (inclusive) for system pools. The default value is 1.                                                                                                                                                                                                                                                                            | int<br/><small>Optional</small>                                                   |
| enableAutoScaling             | Whether to enable auto-scaler                                                                                                                                                                                                                                                                                                                                                                                                                                                                | bool<br/><small>Optional</small>                                                  |
| enableEncryptionAtHost        | This is only supported on certain VM sizes and in certain Azure regions. For more information, see: https://docs.microsoft.com/azure/aks/enable-host-encryption                                                                                                                                                                                                                                                                                                                              | bool<br/><small>Optional</small>                                                  |
| enableFIPS                    | See [Add a FIPS-enabled node pool](https://docs.microsoft.com/azure/aks/use-multiple-node-pools#add-a-fips-enabled-node-pool-preview) for more details.                                                                                                                                                                                                                                                                                                                                      | bool<br/><small>Optional</small>                                                  |
| enableNodePublicIP            | Some scenarios may require nodes in a node pool to receive their own dedicated public IP addresses. A common scenario is for gaming workloads, where a console needs to make a direct connection to a cloud virtual machine to minimize hops. For more information see [assigning a public IP per node](https://docs.microsoft.com/azure/aks/use-multiple-node-pools#assign-a-public-ip-per-node-for-your-node-pools). The default is false.                                                 | bool<br/><small>Optional</small>                                                  |
| enableUltraSSD                | Whether to enable UltraSSD                                                                                                                                                                                                                                                                                                                                                                                                                                                                   | bool<br/><small>Optional</small>                                                  |
| gpuInstanceProfile            | GPUInstanceProfile to be used to specify GPU MIG instance profile for supported GPU VM SKU.                                                                                                                                                                                                                                                                                                                                                                                                  | [GPUInstanceProfile](#GPUInstanceProfile)<br/><small>Optional</small>             |
| kubeletConfig                 | The Kubelet configuration on the agent pool nodes.                                                                                                                                                                                                                                                                                                                                                                                                                                           | [KubeletConfig](#KubeletConfig)<br/><small>Optional</small>                       |
| kubeletDiskType               | Determines the placement of emptyDir volumes, container runtime data root, and Kubelet ephemeral storage.                                                                                                                                                                                                                                                                                                                                                                                    | [KubeletDiskType](#KubeletDiskType)<br/><small>Optional</small>                   |
| linuxOSConfig                 | The OS configuration of Linux agent nodes.                                                                                                                                                                                                                                                                                                                                                                                                                                                   | [LinuxOSConfig](#LinuxOSConfig)<br/><small>Optional</small>                       |
| maxCount                      | The maximum number of nodes for auto-scaling                                                                                                                                                                                                                                                                                                                                                                                                                                                 | int<br/><small>Optional</small>                                                   |
| maxPods                       | The maximum number of pods that can run on a node.                                                                                                                                                                                                                                                                                                                                                                                                                                           | int<br/><small>Optional</small>                                                   |
| minCount                      | The minimum number of nodes for auto-scaling                                                                                                                                                                                                                                                                                                                                                                                                                                                 | int<br/><small>Optional</small>                                                   |
| mode                          | A cluster must have at least one 'System' Agent Pool at all times. For additional information on agent pool restrictions and best practices, see: https://docs.microsoft.com/azure/aks/use-system-pools                                                                                                                                                                                                                                                                                      | [AgentPoolMode](#AgentPoolMode)<br/><small>Optional</small>                       |
| nodeLabels                    | The node labels to be persisted across all nodes in agent pool.                                                                                                                                                                                                                                                                                                                                                                                                                              | map[string]string<br/><small>Optional</small>                                     |
| nodePublicIPPrefixIDReference | This is of the form: /&ZeroWidthSpace;subscriptions/&ZeroWidthSpace;{subscriptionId}/&ZeroWidthSpace;resourceGroups/&ZeroWidthSpace;{resourceGroupName}/&ZeroWidthSpace;providers/&ZeroWidthSpace;Microsoft.Network/publicIPPrefixes/{publicIPPrefixName}                                                                                                                                                                                                                                    | genruntime.ResourceReference<br/><small>Optional</small>                          |
| nodeTaints                    | The taints added to new nodes during node pool create and scale. For example, key=value:NoSchedule.                                                                                                                                                                                                                                                                                                                                                                                          | string[]<br/><small>Optional</small>                                              |
| orchestratorVersion           | As a best practice, you should upgrade all node pools in an AKS cluster to the same Kubernetes version. The node pool version must have the same major version as the control plane. The node pool minor version must be within two minor versions of the control plane version. The node pool version cannot be greater than the control plane version. For more information see [upgrading a node pool](https://docs.microsoft.com/azure/aks/use-multiple-node-pools#upgrade-a-node-pool). | string<br/><small>Optional</small>                                                |
| osDiskSizeGB                  |                                                                                                                                                                                                                                                                                                                                                                                                                                                                                              | [ContainerServiceOSDisk](#ContainerServiceOSDisk)<br/><small>Optional</small>     |
| osDiskType                    | The default is 'Ephemeral' if the VM supports it and has a cache disk larger than the requested OSDiskSizeGB. Otherwise, defaults to 'Managed'. May not be changed after creation. For more information see [Ephemeral OS](https://docs.microsoft.com/azure/aks/cluster-configuration#ephemeral-os).                                                                                                                                                                                         | [OSDiskType](#OSDiskType)<br/><small>Optional</small>                             |
| osSKU                         | Specifies an OS SKU. This value must not be specified if OSType is Windows.                                                                                                                                                                                                                                                                                                                                                                                                                  | [OSSKU](#OSSKU)<br/><small>Optional</small>                                       |
| osType                        | The operating system type. The default is Linux.                                                                                                                                                                                                                                                                                                                                                                                                                                             | [OSType](#OSType)<br/><small>Optional</small>                                     |
| owner                         | The owner of the resource. The owner controls where the resource goes when it is deployed. The owner also controls the resources lifecycle. When the owner is deleted the resource will also be deleted. Owner is expected to be a reference to a containerservice.azure.com/ManagedCluster resource                                                                                                                                                                                         | genruntime.KnownResourceReference<br/><small>Required</small>                     |
| podSubnetIDReference          | If omitted, pod IPs are statically assigned on the node subnet (see vnetSubnetID for more details). This is of the form: /&ZeroWidthSpace;subscriptions/&ZeroWidthSpace;{subscriptionId}/&ZeroWidthSpace;resourceGroups/&ZeroWidthSpace;{resourceGroupName}/&ZeroWidthSpace;providers/&ZeroWidthSpace;Microsoft.Network/virtualNetworks/{virtualNetworkName}/subnets/{subnetName}                                                                                                            | genruntime.ResourceReference<br/><small>Optional</small>                          |
| proximityPlacementGroupID     | The ID for Proximity Placement Group.                                                                                                                                                                                                                                                                                                                                                                                                                                                        | string<br/><small>Optional</small>                                                |
| scaleSetEvictionPolicy        | This cannot be specified unless the scaleSetPriority is 'Spot'. If not specified, the default is 'Delete'.                                                                                                                                                                                                                                                                                                                                                                                   | [ScaleSetEvictionPolicy](#ScaleSetEvictionPolicy)<br/><small>Optional</small>     |
| scaleSetPriority              | The Virtual Machine Scale Set priority. If not specified, the default is 'Regular'.                                                                                                                                                                                                                                                                                                                                                                                                          | [ScaleSetPriority](#ScaleSetPriority)<br/><small>Optional</small>                 |
| spotMaxPrice                  | Possible values are any decimal value greater than zero or -1 which indicates the willingness to pay any on-demand price. For more details on spot pricing, see [spot VMs pricing](https://docs.microsoft.com/azure/virtual-machines/spot-vms#pricing)                                                                                                                                                                                                                                       | float64<br/><small>Optional</small>                                               |
| tags                          | The tags to be persisted on the agent pool virtual machine scale set.                                                                                                                                                                                                                                                                                                                                                                                                                        | map[string]string<br/><small>Optional</small>                                     |
| type                          | The type of Agent Pool.                                                                                                                                                                                                                                                                                                                                                                                                                                                                      | [AgentPoolType](#AgentPoolType)<br/><small>Optional</small>                       |
| upgradeSettings               | Settings for upgrading the agentpool                                                                                                                                                                                                                                                                                                                                                                                                                                                         | [AgentPoolUpgradeSettings](#AgentPoolUpgradeSettings)<br/><small>Optional</small> |
| vmSize                        | VM size availability varies by region. If a node contains insufficient compute resources (memory, cpu, etc) pods might fail to run correctly. For more details on restricted VM sizes, see: https://docs.microsoft.com/azure/aks/quotas-skus-regions                                                                                                                                                                                                                                         | string<br/><small>Optional</small>                                                |
| vnetSubnetIDReference         | If this is not specified, a VNET and subnet will be generated and used. If no podSubnetID is specified, this applies to nodes and pods, otherwise it applies to just nodes. This is of the form: /&ZeroWidthSpace;subscriptions/&ZeroWidthSpace;{subscriptionId}/&ZeroWidthSpace;resourceGroups/&ZeroWidthSpace;{resourceGroupName}/&ZeroWidthSpace;providers/&ZeroWidthSpace;Microsoft.Network/virtualNetworks/{virtualNetworkName}/subnets/{subnetName}                                    | genruntime.ResourceReference<br/><small>Optional</small>                          |

### <a id="ManagedClusters_AgentPool_STATUS"></a>ManagedClusters_AgentPool_STATUS

| Property                  | Description                                                                                                                                                                                                                                                                                                                                                                                                                                                                                  | Type                                                                                            |
|---------------------------|----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|-------------------------------------------------------------------------------------------------|
| availabilityZones         | The list of Availability zones to use for nodes. This can only be specified if the AgentPoolType property is 'VirtualMachineScaleSets'.                                                                                                                                                                                                                                                                                                                                                      | string[]<br/><small>Optional</small>                                                            |
| conditions                | The observed state of the resource                                                                                                                                                                                                                                                                                                                                                                                                                                                           | conditions.Condition[]<br/><small>Optional</small>                                              |
| count                     | Number of agents (VMs) to host docker containers. Allowed values must be in the range of 0 to 1000 (inclusive) for user pools and in the range of 1 to 1000 (inclusive) for system pools. The default value is 1.                                                                                                                                                                                                                                                                            | int<br/><small>Optional</small>                                                                 |
| enableAutoScaling         | Whether to enable auto-scaler                                                                                                                                                                                                                                                                                                                                                                                                                                                                | bool<br/><small>Optional</small>                                                                |
| enableEncryptionAtHost    | This is only supported on certain VM sizes and in certain Azure regions. For more information, see: https://docs.microsoft.com/azure/aks/enable-host-encryption                                                                                                                                                                                                                                                                                                                              | bool<br/><small>Optional</small>                                                                |
| enableFIPS                | See [Add a FIPS-enabled node pool](https://docs.microsoft.com/azure/aks/use-multiple-node-pools#add-a-fips-enabled-node-pool-preview) for more details.                                                                                                                                                                                                                                                                                                                                      | bool<br/><small>Optional</small>                                                                |
| enableNodePublicIP        | Some scenarios may require nodes in a node pool to receive their own dedicated public IP addresses. A common scenario is for gaming workloads, where a console needs to make a direct connection to a cloud virtual machine to minimize hops. For more information see [assigning a public IP per node](https://docs.microsoft.com/azure/aks/use-multiple-node-pools#assign-a-public-ip-per-node-for-your-node-pools). The default is false.                                                 | bool<br/><small>Optional</small>                                                                |
| enableUltraSSD            | Whether to enable UltraSSD                                                                                                                                                                                                                                                                                                                                                                                                                                                                   | bool<br/><small>Optional</small>                                                                |
| gpuInstanceProfile        | GPUInstanceProfile to be used to specify GPU MIG instance profile for supported GPU VM SKU.                                                                                                                                                                                                                                                                                                                                                                                                  | [GPUInstanceProfile_STATUS](#GPUInstanceProfile_STATUS)<br/><small>Optional</small>             |
| id                        | Resource ID.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                 | string<br/><small>Optional</small>                                                              |
| kubeletConfig             | The Kubelet configuration on the agent pool nodes.                                                                                                                                                                                                                                                                                                                                                                                                                                           | [KubeletConfig_STATUS](#KubeletConfig_STATUS)<br/><small>Optional</small>                       |
| kubeletDiskType           | Determines the placement of emptyDir volumes, container runtime data root, and Kubelet ephemeral storage.                                                                                                                                                                                                                                                                                                                                                                                    | [KubeletDiskType_STATUS](#KubeletDiskType_STATUS)<br/><small>Optional</small>                   |
| linuxOSConfig             | The OS configuration of Linux agent nodes.                                                                                                                                                                                                                                                                                                                                                                                                                                                   | [LinuxOSConfig_STATUS](#LinuxOSConfig_STATUS)<br/><small>Optional</small>                       |
| maxCount                  | The maximum number of nodes for auto-scaling                                                                                                                                                                                                                                                                                                                                                                                                                                                 | int<br/><small>Optional</small>                                                                 |
| maxPods                   | The maximum number of pods that can run on a node.                                                                                                                                                                                                                                                                                                                                                                                                                                           | int<br/><small>Optional</small>                                                                 |
| minCount                  | The minimum number of nodes for auto-scaling                                                                                                                                                                                                                                                                                                                                                                                                                                                 | int<br/><small>Optional</small>                                                                 |
| mode                      | A cluster must have at least one 'System' Agent Pool at all times. For additional information on agent pool restrictions and best practices, see: https://docs.microsoft.com/azure/aks/use-system-pools                                                                                                                                                                                                                                                                                      | [AgentPoolMode_STATUS](#AgentPoolMode_STATUS)<br/><small>Optional</small>                       |
| name                      | The name of the resource that is unique within a resource group. This name can be used to access the resource.                                                                                                                                                                                                                                                                                                                                                                               | string<br/><small>Optional</small>                                                              |
| nodeImageVersion          | The version of node image                                                                                                                                                                                                                                                                                                                                                                                                                                                                    | string<br/><small>Optional</small>                                                              |
| nodeLabels                | The node labels to be persisted across all nodes in agent pool.                                                                                                                                                                                                                                                                                                                                                                                                                              | map[string]string<br/><small>Optional</small>                                                   |
| nodePublicIPPrefixID      | This is of the form: /&ZeroWidthSpace;subscriptions/&ZeroWidthSpace;{subscriptionId}/&ZeroWidthSpace;resourceGroups/&ZeroWidthSpace;{resourceGroupName}/&ZeroWidthSpace;providers/&ZeroWidthSpace;Microsoft.Network/publicIPPrefixes/{publicIPPrefixName}                                                                                                                                                                                                                                    | string<br/><small>Optional</small>                                                              |
| nodeTaints                | The taints added to new nodes during node pool create and scale. For example, key=value:NoSchedule.                                                                                                                                                                                                                                                                                                                                                                                          | string[]<br/><small>Optional</small>                                                            |
| orchestratorVersion       | As a best practice, you should upgrade all node pools in an AKS cluster to the same Kubernetes version. The node pool version must have the same major version as the control plane. The node pool minor version must be within two minor versions of the control plane version. The node pool version cannot be greater than the control plane version. For more information see [upgrading a node pool](https://docs.microsoft.com/azure/aks/use-multiple-node-pools#upgrade-a-node-pool). | string<br/><small>Optional</small>                                                              |
| osDiskSizeGB              |                                                                                                                                                                                                                                                                                                                                                                                                                                                                                              | int<br/><small>Optional</small>                                                                 |
| osDiskType                | The default is 'Ephemeral' if the VM supports it and has a cache disk larger than the requested OSDiskSizeGB. Otherwise, defaults to 'Managed'. May not be changed after creation. For more information see [Ephemeral OS](https://docs.microsoft.com/azure/aks/cluster-configuration#ephemeral-os).                                                                                                                                                                                         | [OSDiskType_STATUS](#OSDiskType_STATUS)<br/><small>Optional</small>                             |
| osSKU                     | Specifies an OS SKU. This value must not be specified if OSType is Windows.                                                                                                                                                                                                                                                                                                                                                                                                                  | [OSSKU_STATUS](#OSSKU_STATUS)<br/><small>Optional</small>                                       |
| osType                    | The operating system type. The default is Linux.                                                                                                                                                                                                                                                                                                                                                                                                                                             | [OSType_STATUS](#OSType_STATUS)<br/><small>Optional</small>                                     |
| podSubnetID               | If omitted, pod IPs are statically assigned on the node subnet (see vnetSubnetID for more details). This is of the form: /&ZeroWidthSpace;subscriptions/&ZeroWidthSpace;{subscriptionId}/&ZeroWidthSpace;resourceGroups/&ZeroWidthSpace;{resourceGroupName}/&ZeroWidthSpace;providers/&ZeroWidthSpace;Microsoft.Network/virtualNetworks/{virtualNetworkName}/subnets/{subnetName}                                                                                                            | string<br/><small>Optional</small>                                                              |
| powerState                | Describes whether the Agent Pool is Running or Stopped                                                                                                                                                                                                                                                                                                                                                                                                                                       | [PowerState_STATUS](#PowerState_STATUS)<br/><small>Optional</small>                             |
| properties_type           | The type of Agent Pool.                                                                                                                                                                                                                                                                                                                                                                                                                                                                      | [AgentPoolType_STATUS](#AgentPoolType_STATUS)<br/><small>Optional</small>                       |
| provisioningState         | The current deployment or provisioning state.                                                                                                                                                                                                                                                                                                                                                                                                                                                | string<br/><small>Optional</small>                                                              |
| proximityPlacementGroupID | The ID for Proximity Placement Group.                                                                                                                                                                                                                                                                                                                                                                                                                                                        | string<br/><small>Optional</small>                                                              |
| scaleSetEvictionPolicy    | This cannot be specified unless the scaleSetPriority is 'Spot'. If not specified, the default is 'Delete'.                                                                                                                                                                                                                                                                                                                                                                                   | [ScaleSetEvictionPolicy_STATUS](#ScaleSetEvictionPolicy_STATUS)<br/><small>Optional</small>     |
| scaleSetPriority          | The Virtual Machine Scale Set priority. If not specified, the default is 'Regular'.                                                                                                                                                                                                                                                                                                                                                                                                          | [ScaleSetPriority_STATUS](#ScaleSetPriority_STATUS)<br/><small>Optional</small>                 |
| spotMaxPrice              | Possible values are any decimal value greater than zero or -1 which indicates the willingness to pay any on-demand price. For more details on spot pricing, see [spot VMs pricing](https://docs.microsoft.com/azure/virtual-machines/spot-vms#pricing)                                                                                                                                                                                                                                       | float64<br/><small>Optional</small>                                                             |
| tags                      | The tags to be persisted on the agent pool virtual machine scale set.                                                                                                                                                                                                                                                                                                                                                                                                                        | map[string]string<br/><small>Optional</small>                                                   |
| type                      | Resource type                                                                                                                                                                                                                                                                                                                                                                                                                                                                                | string<br/><small>Optional</small>                                                              |
| upgradeSettings           | Settings for upgrading the agentpool                                                                                                                                                                                                                                                                                                                                                                                                                                                         | [AgentPoolUpgradeSettings_STATUS](#AgentPoolUpgradeSettings_STATUS)<br/><small>Optional</small> |
| vmSize                    | VM size availability varies by region. If a node contains insufficient compute resources (memory, cpu, etc) pods might fail to run correctly. For more details on restricted VM sizes, see: https://docs.microsoft.com/azure/aks/quotas-skus-regions                                                                                                                                                                                                                                         | string<br/><small>Optional</small>                                                              |
| vnetSubnetID              | If this is not specified, a VNET and subnet will be generated and used. If no podSubnetID is specified, this applies to nodes and pods, otherwise it applies to just nodes. This is of the form: /&ZeroWidthSpace;subscriptions/&ZeroWidthSpace;{subscriptionId}/&ZeroWidthSpace;resourceGroups/&ZeroWidthSpace;{resourceGroupName}/&ZeroWidthSpace;providers/&ZeroWidthSpace;Microsoft.Network/virtualNetworks/{virtualNetworkName}/subnets/{subnetName}                                    | string<br/><small>Optional</small>                                                              |

<a id="ManagedClustersAgentPoolList"></a>ManagedClustersAgentPoolList
---------------------------------------------------------------------

Generator information: - Generated from: /containerservice/resource-manager/Microsoft.ContainerService/aks/stable/2021-05-01/managedClusters.json - ARM URI: /&ZeroWidthSpace;subscriptions/&ZeroWidthSpace;{subscriptionId}/&ZeroWidthSpace;resourceGroups/&ZeroWidthSpace;{resourceGroupName}/&ZeroWidthSpace;providers/&ZeroWidthSpace;Microsoft.ContainerService/managedClusters/{resourceName}/agentPools/{agentPoolName}

| Property        | Description | Type                                                                                |
|-----------------|-------------|-------------------------------------------------------------------------------------|
| metav1.TypeMeta |             |                                                                                     |
| metav1.ListMeta |             |                                                                                     |
| items           |             | [ManagedClustersAgentPool[]](#ManagedClustersAgentPool)<br/><small>Optional</small> |

<a id="ManagedCluster_Spec"></a>ManagedCluster_Spec
---------------------------------------------------

Used by: [ManagedCluster](#ManagedCluster).

| Property                     | Description                                                                                                                                                                                                                                                                                                                                                                                 | Type                                                                                                                  |
|------------------------------|---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|-----------------------------------------------------------------------------------------------------------------------|
| aadProfile                   | The Azure Active Directory configuration.                                                                                                                                                                                                                                                                                                                                                   | [ManagedClusterAADProfile](#ManagedClusterAADProfile)<br/><small>Optional</small>                                     |
| addonProfiles                | The profile of managed cluster add-on.                                                                                                                                                                                                                                                                                                                                                      | [map[string]ManagedClusterAddonProfile](#ManagedClusterAddonProfile)<br/><small>Optional</small>                      |
| agentPoolProfiles            | The agent pool properties.                                                                                                                                                                                                                                                                                                                                                                  | [ManagedClusterAgentPoolProfile[]](#ManagedClusterAgentPoolProfile)<br/><small>Optional</small>                       |
| apiServerAccessProfile       | The access profile for managed cluster API server.                                                                                                                                                                                                                                                                                                                                          | [ManagedClusterAPIServerAccessProfile](#ManagedClusterAPIServerAccessProfile)<br/><small>Optional</small>             |
| autoScalerProfile            | Parameters to be applied to the cluster-autoscaler when enabled                                                                                                                                                                                                                                                                                                                             | [ManagedClusterProperties_AutoScalerProfile](#ManagedClusterProperties_AutoScalerProfile)<br/><small>Optional</small> |
| autoUpgradeProfile           | The auto upgrade configuration.                                                                                                                                                                                                                                                                                                                                                             | [ManagedClusterAutoUpgradeProfile](#ManagedClusterAutoUpgradeProfile)<br/><small>Optional</small>                     |
| azureName                    | The name of the resource in Azure. This is often the same as the name of the resource in Kubernetes but it doesn't have to be.                                                                                                                                                                                                                                                              | string<br/><small>Optional</small>                                                                                    |
| disableLocalAccounts         | If set to true, getting static credentials will be disabled for this cluster. This must only be used on Managed Clusters that are AAD enabled. For more details see [disable local accounts](https://docs.microsoft.com/azure/aks/managed-aad#disable-local-accounts-preview).                                                                                                              | bool<br/><small>Optional</small>                                                                                      |
| diskEncryptionSetIDReference | This is of the form: '/&ZeroWidthSpace;subscriptions/&ZeroWidthSpace;{subscriptionId}/&ZeroWidthSpace;resourceGroups/&ZeroWidthSpace;{resourceGroupName}/&ZeroWidthSpace;providers/&ZeroWidthSpace;Microsoft.Compute/diskEncryptionSets/{encryptionSetName}'                                                                                                                                | genruntime.ResourceReference<br/><small>Optional</small>                                                              |
| dnsPrefix                    | This cannot be updated once the Managed Cluster has been created.                                                                                                                                                                                                                                                                                                                           | string<br/><small>Optional</small>                                                                                    |
| enablePodSecurityPolicy      | (DEPRECATING) Whether to enable Kubernetes pod security policy (preview). This feature is set for removal on October 15th, 2020. Learn more at aka.ms/aks/azpodpolicy.                                                                                                                                                                                                                      | bool<br/><small>Optional</small>                                                                                      |
| enableRBAC                   | Whether to enable Kubernetes Role-Based Access Control.                                                                                                                                                                                                                                                                                                                                     | bool<br/><small>Optional</small>                                                                                      |
| extendedLocation             | The extended location of the Virtual Machine.                                                                                                                                                                                                                                                                                                                                               | [ExtendedLocation](#ExtendedLocation)<br/><small>Optional</small>                                                     |
| fqdnSubdomain                | This cannot be updated once the Managed Cluster has been created.                                                                                                                                                                                                                                                                                                                           | string<br/><small>Optional</small>                                                                                    |
| httpProxyConfig              | Configurations for provisioning the cluster with HTTP proxy servers.                                                                                                                                                                                                                                                                                                                        | [ManagedClusterHTTPProxyConfig](#ManagedClusterHTTPProxyConfig)<br/><small>Optional</small>                           |
| identity                     | The identity of the managed cluster, if configured.                                                                                                                                                                                                                                                                                                                                         | [ManagedClusterIdentity](#ManagedClusterIdentity)<br/><small>Optional</small>                                         |
| identityProfile              | Identities associated with the cluster.                                                                                                                                                                                                                                                                                                                                                     | [map[string]UserAssignedIdentity](#UserAssignedIdentity)<br/><small>Optional</small>                                  |
| kubernetesVersion            | When you upgrade a supported AKS cluster, Kubernetes minor versions cannot be skipped. All upgrades must be performed sequentially by major version number. For example, upgrades between 1.14.x -> 1.15.x or 1.15.x -> 1.16.x are allowed, however 1.14.x -> 1.16.x is not allowed. See [upgrading an AKS cluster](https://docs.microsoft.com/azure/aks/upgrade-cluster) for more details. | string<br/><small>Optional</small>                                                                                    |
| linuxProfile                 | The profile for Linux VMs in the Managed Cluster.                                                                                                                                                                                                                                                                                                                                           | [ContainerServiceLinuxProfile](#ContainerServiceLinuxProfile)<br/><small>Optional</small>                             |
| location                     | Resource location                                                                                                                                                                                                                                                                                                                                                                           | string<br/><small>Required</small>                                                                                    |
| networkProfile               | The network configuration profile.                                                                                                                                                                                                                                                                                                                                                          | [ContainerServiceNetworkProfile](#ContainerServiceNetworkProfile)<br/><small>Optional</small>                         |
| nodeResourceGroup            | The name of the resource group containing agent pool nodes.                                                                                                                                                                                                                                                                                                                                 | string<br/><small>Optional</small>                                                                                    |
| operatorSpec                 | The specification for configuring operator behavior. This field is interpreted by the operator and not passed directly to Azure                                                                                                                                                                                                                                                             | [ManagedClusterOperatorSpec](#ManagedClusterOperatorSpec)<br/><small>Optional</small>                                 |
| owner                        | The owner of the resource. The owner controls where the resource goes when it is deployed. The owner also controls the resources lifecycle. When the owner is deleted the resource will also be deleted. Owner is expected to be a reference to a resources.azure.com/ResourceGroup resource                                                                                                | genruntime.KnownResourceReference<br/><small>Required</small>                                                         |
| podIdentityProfile           | See [use AAD pod identity](https://docs.microsoft.com/azure/aks/use-azure-ad-pod-identity) for more details on AAD pod identity integration.                                                                                                                                                                                                                                                | [ManagedClusterPodIdentityProfile](#ManagedClusterPodIdentityProfile)<br/><small>Optional</small>                     |
| privateLinkResources         | Private link resources associated with the cluster.                                                                                                                                                                                                                                                                                                                                         | [PrivateLinkResource[]](#PrivateLinkResource)<br/><small>Optional</small>                                             |
| servicePrincipalProfile      | Information about a service principal identity for the cluster to use for manipulating Azure APIs.                                                                                                                                                                                                                                                                                          | [ManagedClusterServicePrincipalProfile](#ManagedClusterServicePrincipalProfile)<br/><small>Optional</small>           |
| sku                          | The managed cluster SKU.                                                                                                                                                                                                                                                                                                                                                                    | [ManagedClusterSKU](#ManagedClusterSKU)<br/><small>Optional</small>                                                   |
| tags                         | Resource tags                                                                                                                                                                                                                                                                                                                                                                               | map[string]string<br/><small>Optional</small>                                                                         |
| windowsProfile               | The profile for Windows VMs in the Managed Cluster.                                                                                                                                                                                                                                                                                                                                         | [ManagedClusterWindowsProfile](#ManagedClusterWindowsProfile)<br/><small>Optional</small>                             |

<a id="ManagedCluster_STATUS"></a>ManagedCluster_STATUS
-------------------------------------------------------

Managed cluster.

Used by: [ManagedCluster](#ManagedCluster).

| Property                | Description                                                                                                                                                                                                                                                                                                                                                                                 | Type                                                                                                                                |
|-------------------------|---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|-------------------------------------------------------------------------------------------------------------------------------------|
| aadProfile              | The Azure Active Directory configuration.                                                                                                                                                                                                                                                                                                                                                   | [ManagedClusterAADProfile_STATUS](#ManagedClusterAADProfile_STATUS)<br/><small>Optional</small>                                     |
| addonProfiles           | The profile of managed cluster add-on.                                                                                                                                                                                                                                                                                                                                                      | [map[string]ManagedClusterAddonProfile_STATUS](#ManagedClusterAddonProfile_STATUS)<br/><small>Optional</small>                      |
| agentPoolProfiles       | The agent pool properties.                                                                                                                                                                                                                                                                                                                                                                  | [ManagedClusterAgentPoolProfile_STATUS[]](#ManagedClusterAgentPoolProfile_STATUS)<br/><small>Optional</small>                       |
| apiServerAccessProfile  | The access profile for managed cluster API server.                                                                                                                                                                                                                                                                                                                                          | [ManagedClusterAPIServerAccessProfile_STATUS](#ManagedClusterAPIServerAccessProfile_STATUS)<br/><small>Optional</small>             |
| autoScalerProfile       | Parameters to be applied to the cluster-autoscaler when enabled                                                                                                                                                                                                                                                                                                                             | [ManagedClusterProperties_AutoScalerProfile_STATUS](#ManagedClusterProperties_AutoScalerProfile_STATUS)<br/><small>Optional</small> |
| autoUpgradeProfile      | The auto upgrade configuration.                                                                                                                                                                                                                                                                                                                                                             | [ManagedClusterAutoUpgradeProfile_STATUS](#ManagedClusterAutoUpgradeProfile_STATUS)<br/><small>Optional</small>                     |
| azurePortalFQDN         | The Azure Portal requires certain Cross-Origin Resource Sharing (CORS) headers to be sent in some responses, which Kubernetes APIServer doesn't handle by default. This special FQDN supports CORS, allowing the Azure Portal to function properly.                                                                                                                                         | string<br/><small>Optional</small>                                                                                                  |
| conditions              | The observed state of the resource                                                                                                                                                                                                                                                                                                                                                          | conditions.Condition[]<br/><small>Optional</small>                                                                                  |
| disableLocalAccounts    | If set to true, getting static credentials will be disabled for this cluster. This must only be used on Managed Clusters that are AAD enabled. For more details see [disable local accounts](https://docs.microsoft.com/azure/aks/managed-aad#disable-local-accounts-preview).                                                                                                              | bool<br/><small>Optional</small>                                                                                                    |
| diskEncryptionSetID     | This is of the form: '/&ZeroWidthSpace;subscriptions/&ZeroWidthSpace;{subscriptionId}/&ZeroWidthSpace;resourceGroups/&ZeroWidthSpace;{resourceGroupName}/&ZeroWidthSpace;providers/&ZeroWidthSpace;Microsoft.Compute/diskEncryptionSets/{encryptionSetName}'                                                                                                                                | string<br/><small>Optional</small>                                                                                                  |
| dnsPrefix               | This cannot be updated once the Managed Cluster has been created.                                                                                                                                                                                                                                                                                                                           | string<br/><small>Optional</small>                                                                                                  |
| enablePodSecurityPolicy | (DEPRECATING) Whether to enable Kubernetes pod security policy (preview). This feature is set for removal on October 15th, 2020. Learn more at aka.ms/aks/azpodpolicy.                                                                                                                                                                                                                      | bool<br/><small>Optional</small>                                                                                                    |
| enableRBAC              | Whether to enable Kubernetes Role-Based Access Control.                                                                                                                                                                                                                                                                                                                                     | bool<br/><small>Optional</small>                                                                                                    |
| extendedLocation        | The extended location of the Virtual Machine.                                                                                                                                                                                                                                                                                                                                               | [ExtendedLocation_STATUS](#ExtendedLocation_STATUS)<br/><small>Optional</small>                                                     |
| fqdn                    | The FQDN of the master pool.                                                                                                                                                                                                                                                                                                                                                                | string<br/><small>Optional</small>                                                                                                  |
| fqdnSubdomain           | This cannot be updated once the Managed Cluster has been created.                                                                                                                                                                                                                                                                                                                           | string<br/><small>Optional</small>                                                                                                  |
| httpProxyConfig         | Configurations for provisioning the cluster with HTTP proxy servers.                                                                                                                                                                                                                                                                                                                        | [ManagedClusterHTTPProxyConfig_STATUS](#ManagedClusterHTTPProxyConfig_STATUS)<br/><small>Optional</small>                           |
| id                      | Resource Id                                                                                                                                                                                                                                                                                                                                                                                 | string<br/><small>Optional</small>                                                                                                  |
| identity                | The identity of the managed cluster, if configured.                                                                                                                                                                                                                                                                                                                                         | [ManagedClusterIdentity_STATUS](#ManagedClusterIdentity_STATUS)<br/><small>Optional</small>                                         |
| identityProfile         | Identities associated with the cluster.                                                                                                                                                                                                                                                                                                                                                     | [map[string]UserAssignedIdentity_STATUS](#UserAssignedIdentity_STATUS)<br/><small>Optional</small>                                  |
| kubernetesVersion       | When you upgrade a supported AKS cluster, Kubernetes minor versions cannot be skipped. All upgrades must be performed sequentially by major version number. For example, upgrades between 1.14.x -> 1.15.x or 1.15.x -> 1.16.x are allowed, however 1.14.x -> 1.16.x is not allowed. See [upgrading an AKS cluster](https://docs.microsoft.com/azure/aks/upgrade-cluster) for more details. | string<br/><small>Optional</small>                                                                                                  |
| linuxProfile            | The profile for Linux VMs in the Managed Cluster.                                                                                                                                                                                                                                                                                                                                           | [ContainerServiceLinuxProfile_STATUS](#ContainerServiceLinuxProfile_STATUS)<br/><small>Optional</small>                             |
| location                | Resource location                                                                                                                                                                                                                                                                                                                                                                           | string<br/><small>Optional</small>                                                                                                  |
| maxAgentPools           | The max number of agent pools for the managed cluster.                                                                                                                                                                                                                                                                                                                                      | int<br/><small>Optional</small>                                                                                                     |
| name                    | Resource name                                                                                                                                                                                                                                                                                                                                                                               | string<br/><small>Optional</small>                                                                                                  |
| networkProfile          | The network configuration profile.                                                                                                                                                                                                                                                                                                                                                          | [ContainerServiceNetworkProfile_STATUS](#ContainerServiceNetworkProfile_STATUS)<br/><small>Optional</small>                         |
| nodeResourceGroup       | The name of the resource group containing agent pool nodes.                                                                                                                                                                                                                                                                                                                                 | string<br/><small>Optional</small>                                                                                                  |
| podIdentityProfile      | See [use AAD pod identity](https://docs.microsoft.com/azure/aks/use-azure-ad-pod-identity) for more details on AAD pod identity integration.                                                                                                                                                                                                                                                | [ManagedClusterPodIdentityProfile_STATUS](#ManagedClusterPodIdentityProfile_STATUS)<br/><small>Optional</small>                     |
| powerState              | The Power State of the cluster.                                                                                                                                                                                                                                                                                                                                                             | [PowerState_STATUS](#PowerState_STATUS)<br/><small>Optional</small>                                                                 |
| privateFQDN             | The FQDN of private cluster.                                                                                                                                                                                                                                                                                                                                                                | string<br/><small>Optional</small>                                                                                                  |
| privateLinkResources    | Private link resources associated with the cluster.                                                                                                                                                                                                                                                                                                                                         | [PrivateLinkResource_STATUS[]](#PrivateLinkResource_STATUS)<br/><small>Optional</small>                                             |
| provisioningState       | The current provisioning state.                                                                                                                                                                                                                                                                                                                                                             | string<br/><small>Optional</small>                                                                                                  |
| servicePrincipalProfile | Information about a service principal identity for the cluster to use for manipulating Azure APIs.                                                                                                                                                                                                                                                                                          | [ManagedClusterServicePrincipalProfile_STATUS](#ManagedClusterServicePrincipalProfile_STATUS)<br/><small>Optional</small>           |
| sku                     | The managed cluster SKU.                                                                                                                                                                                                                                                                                                                                                                    | [ManagedClusterSKU_STATUS](#ManagedClusterSKU_STATUS)<br/><small>Optional</small>                                                   |
| tags                    | Resource tags                                                                                                                                                                                                                                                                                                                                                                               | map[string]string<br/><small>Optional</small>                                                                                       |
| type                    | Resource type                                                                                                                                                                                                                                                                                                                                                                               | string<br/><small>Optional</small>                                                                                                  |
| windowsProfile          | The profile for Windows VMs in the Managed Cluster.                                                                                                                                                                                                                                                                                                                                         | [ManagedClusterWindowsProfile_STATUS](#ManagedClusterWindowsProfile_STATUS)<br/><small>Optional</small>                             |

<a id="ManagedClusters_AgentPool_Spec"></a>ManagedClusters_AgentPool_Spec
-------------------------------------------------------------------------

Used by: [ManagedClustersAgentPool](#ManagedClustersAgentPool).

| Property                      | Description                                                                                                                                                                                                                                                                                                                                                                                                                                                                                  | Type                                                                              |
|-------------------------------|----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|-----------------------------------------------------------------------------------|
| availabilityZones             | The list of Availability zones to use for nodes. This can only be specified if the AgentPoolType property is 'VirtualMachineScaleSets'.                                                                                                                                                                                                                                                                                                                                                      | string[]<br/><small>Optional</small>                                              |
| azureName                     | The name of the resource in Azure. This is often the same as the name of the resource in Kubernetes but it doesn't have to be.                                                                                                                                                                                                                                                                                                                                                               | string<br/><small>Optional</small>                                                |
| count                         | Number of agents (VMs) to host docker containers. Allowed values must be in the range of 0 to 1000 (inclusive) for user pools and in the range of 1 to 1000 (inclusive) for system pools. The default value is 1.                                                                                                                                                                                                                                                                            | int<br/><small>Optional</small>                                                   |
| enableAutoScaling             | Whether to enable auto-scaler                                                                                                                                                                                                                                                                                                                                                                                                                                                                | bool<br/><small>Optional</small>                                                  |
| enableEncryptionAtHost        | This is only supported on certain VM sizes and in certain Azure regions. For more information, see: https://docs.microsoft.com/azure/aks/enable-host-encryption                                                                                                                                                                                                                                                                                                                              | bool<br/><small>Optional</small>                                                  |
| enableFIPS                    | See [Add a FIPS-enabled node pool](https://docs.microsoft.com/azure/aks/use-multiple-node-pools#add-a-fips-enabled-node-pool-preview) for more details.                                                                                                                                                                                                                                                                                                                                      | bool<br/><small>Optional</small>                                                  |
| enableNodePublicIP            | Some scenarios may require nodes in a node pool to receive their own dedicated public IP addresses. A common scenario is for gaming workloads, where a console needs to make a direct connection to a cloud virtual machine to minimize hops. For more information see [assigning a public IP per node](https://docs.microsoft.com/azure/aks/use-multiple-node-pools#assign-a-public-ip-per-node-for-your-node-pools). The default is false.                                                 | bool<br/><small>Optional</small>                                                  |
| enableUltraSSD                | Whether to enable UltraSSD                                                                                                                                                                                                                                                                                                                                                                                                                                                                   | bool<br/><small>Optional</small>                                                  |
| gpuInstanceProfile            | GPUInstanceProfile to be used to specify GPU MIG instance profile for supported GPU VM SKU.                                                                                                                                                                                                                                                                                                                                                                                                  | [GPUInstanceProfile](#GPUInstanceProfile)<br/><small>Optional</small>             |
| kubeletConfig                 | The Kubelet configuration on the agent pool nodes.                                                                                                                                                                                                                                                                                                                                                                                                                                           | [KubeletConfig](#KubeletConfig)<br/><small>Optional</small>                       |
| kubeletDiskType               | Determines the placement of emptyDir volumes, container runtime data root, and Kubelet ephemeral storage.                                                                                                                                                                                                                                                                                                                                                                                    | [KubeletDiskType](#KubeletDiskType)<br/><small>Optional</small>                   |
| linuxOSConfig                 | The OS configuration of Linux agent nodes.                                                                                                                                                                                                                                                                                                                                                                                                                                                   | [LinuxOSConfig](#LinuxOSConfig)<br/><small>Optional</small>                       |
| maxCount                      | The maximum number of nodes for auto-scaling                                                                                                                                                                                                                                                                                                                                                                                                                                                 | int<br/><small>Optional</small>                                                   |
| maxPods                       | The maximum number of pods that can run on a node.                                                                                                                                                                                                                                                                                                                                                                                                                                           | int<br/><small>Optional</small>                                                   |
| minCount                      | The minimum number of nodes for auto-scaling                                                                                                                                                                                                                                                                                                                                                                                                                                                 | int<br/><small>Optional</small>                                                   |
| mode                          | A cluster must have at least one 'System' Agent Pool at all times. For additional information on agent pool restrictions and best practices, see: https://docs.microsoft.com/azure/aks/use-system-pools                                                                                                                                                                                                                                                                                      | [AgentPoolMode](#AgentPoolMode)<br/><small>Optional</small>                       |
| nodeLabels                    | The node labels to be persisted across all nodes in agent pool.                                                                                                                                                                                                                                                                                                                                                                                                                              | map[string]string<br/><small>Optional</small>                                     |
| nodePublicIPPrefixIDReference | This is of the form: /&ZeroWidthSpace;subscriptions/&ZeroWidthSpace;{subscriptionId}/&ZeroWidthSpace;resourceGroups/&ZeroWidthSpace;{resourceGroupName}/&ZeroWidthSpace;providers/&ZeroWidthSpace;Microsoft.Network/publicIPPrefixes/{publicIPPrefixName}                                                                                                                                                                                                                                    | genruntime.ResourceReference<br/><small>Optional</small>                          |
| nodeTaints                    | The taints added to new nodes during node pool create and scale. For example, key=value:NoSchedule.                                                                                                                                                                                                                                                                                                                                                                                          | string[]<br/><small>Optional</small>                                              |
| orchestratorVersion           | As a best practice, you should upgrade all node pools in an AKS cluster to the same Kubernetes version. The node pool version must have the same major version as the control plane. The node pool minor version must be within two minor versions of the control plane version. The node pool version cannot be greater than the control plane version. For more information see [upgrading a node pool](https://docs.microsoft.com/azure/aks/use-multiple-node-pools#upgrade-a-node-pool). | string<br/><small>Optional</small>                                                |
| osDiskSizeGB                  |                                                                                                                                                                                                                                                                                                                                                                                                                                                                                              | [ContainerServiceOSDisk](#ContainerServiceOSDisk)<br/><small>Optional</small>     |
| osDiskType                    | The default is 'Ephemeral' if the VM supports it and has a cache disk larger than the requested OSDiskSizeGB. Otherwise, defaults to 'Managed'. May not be changed after creation. For more information see [Ephemeral OS](https://docs.microsoft.com/azure/aks/cluster-configuration#ephemeral-os).                                                                                                                                                                                         | [OSDiskType](#OSDiskType)<br/><small>Optional</small>                             |
| osSKU                         | Specifies an OS SKU. This value must not be specified if OSType is Windows.                                                                                                                                                                                                                                                                                                                                                                                                                  | [OSSKU](#OSSKU)<br/><small>Optional</small>                                       |
| osType                        | The operating system type. The default is Linux.                                                                                                                                                                                                                                                                                                                                                                                                                                             | [OSType](#OSType)<br/><small>Optional</small>                                     |
| owner                         | The owner of the resource. The owner controls where the resource goes when it is deployed. The owner also controls the resources lifecycle. When the owner is deleted the resource will also be deleted. Owner is expected to be a reference to a containerservice.azure.com/ManagedCluster resource                                                                                                                                                                                         | genruntime.KnownResourceReference<br/><small>Required</small>                     |
| podSubnetIDReference          | If omitted, pod IPs are statically assigned on the node subnet (see vnetSubnetID for more details). This is of the form: /&ZeroWidthSpace;subscriptions/&ZeroWidthSpace;{subscriptionId}/&ZeroWidthSpace;resourceGroups/&ZeroWidthSpace;{resourceGroupName}/&ZeroWidthSpace;providers/&ZeroWidthSpace;Microsoft.Network/virtualNetworks/{virtualNetworkName}/subnets/{subnetName}                                                                                                            | genruntime.ResourceReference<br/><small>Optional</small>                          |
| proximityPlacementGroupID     | The ID for Proximity Placement Group.                                                                                                                                                                                                                                                                                                                                                                                                                                                        | string<br/><small>Optional</small>                                                |
| scaleSetEvictionPolicy        | This cannot be specified unless the scaleSetPriority is 'Spot'. If not specified, the default is 'Delete'.                                                                                                                                                                                                                                                                                                                                                                                   | [ScaleSetEvictionPolicy](#ScaleSetEvictionPolicy)<br/><small>Optional</small>     |
| scaleSetPriority              | The Virtual Machine Scale Set priority. If not specified, the default is 'Regular'.                                                                                                                                                                                                                                                                                                                                                                                                          | [ScaleSetPriority](#ScaleSetPriority)<br/><small>Optional</small>                 |
| spotMaxPrice                  | Possible values are any decimal value greater than zero or -1 which indicates the willingness to pay any on-demand price. For more details on spot pricing, see [spot VMs pricing](https://docs.microsoft.com/azure/virtual-machines/spot-vms#pricing)                                                                                                                                                                                                                                       | float64<br/><small>Optional</small>                                               |
| tags                          | The tags to be persisted on the agent pool virtual machine scale set.                                                                                                                                                                                                                                                                                                                                                                                                                        | map[string]string<br/><small>Optional</small>                                     |
| type                          | The type of Agent Pool.                                                                                                                                                                                                                                                                                                                                                                                                                                                                      | [AgentPoolType](#AgentPoolType)<br/><small>Optional</small>                       |
| upgradeSettings               | Settings for upgrading the agentpool                                                                                                                                                                                                                                                                                                                                                                                                                                                         | [AgentPoolUpgradeSettings](#AgentPoolUpgradeSettings)<br/><small>Optional</small> |
| vmSize                        | VM size availability varies by region. If a node contains insufficient compute resources (memory, cpu, etc) pods might fail to run correctly. For more details on restricted VM sizes, see: https://docs.microsoft.com/azure/aks/quotas-skus-regions                                                                                                                                                                                                                                         | string<br/><small>Optional</small>                                                |
| vnetSubnetIDReference         | If this is not specified, a VNET and subnet will be generated and used. If no podSubnetID is specified, this applies to nodes and pods, otherwise it applies to just nodes. This is of the form: /&ZeroWidthSpace;subscriptions/&ZeroWidthSpace;{subscriptionId}/&ZeroWidthSpace;resourceGroups/&ZeroWidthSpace;{resourceGroupName}/&ZeroWidthSpace;providers/&ZeroWidthSpace;Microsoft.Network/virtualNetworks/{virtualNetworkName}/subnets/{subnetName}                                    | genruntime.ResourceReference<br/><small>Optional</small>                          |

<a id="ManagedClusters_AgentPool_STATUS"></a>ManagedClusters_AgentPool_STATUS
-----------------------------------------------------------------------------

Used by: [ManagedClustersAgentPool](#ManagedClustersAgentPool).

| Property                  | Description                                                                                                                                                                                                                                                                                                                                                                                                                                                                                  | Type                                                                                            |
|---------------------------|----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|-------------------------------------------------------------------------------------------------|
| availabilityZones         | The list of Availability zones to use for nodes. This can only be specified if the AgentPoolType property is 'VirtualMachineScaleSets'.                                                                                                                                                                                                                                                                                                                                                      | string[]<br/><small>Optional</small>                                                            |
| conditions                | The observed state of the resource                                                                                                                                                                                                                                                                                                                                                                                                                                                           | conditions.Condition[]<br/><small>Optional</small>                                              |
| count                     | Number of agents (VMs) to host docker containers. Allowed values must be in the range of 0 to 1000 (inclusive) for user pools and in the range of 1 to 1000 (inclusive) for system pools. The default value is 1.                                                                                                                                                                                                                                                                            | int<br/><small>Optional</small>                                                                 |
| enableAutoScaling         | Whether to enable auto-scaler                                                                                                                                                                                                                                                                                                                                                                                                                                                                | bool<br/><small>Optional</small>                                                                |
| enableEncryptionAtHost    | This is only supported on certain VM sizes and in certain Azure regions. For more information, see: https://docs.microsoft.com/azure/aks/enable-host-encryption                                                                                                                                                                                                                                                                                                                              | bool<br/><small>Optional</small>                                                                |
| enableFIPS                | See [Add a FIPS-enabled node pool](https://docs.microsoft.com/azure/aks/use-multiple-node-pools#add-a-fips-enabled-node-pool-preview) for more details.                                                                                                                                                                                                                                                                                                                                      | bool<br/><small>Optional</small>                                                                |
| enableNodePublicIP        | Some scenarios may require nodes in a node pool to receive their own dedicated public IP addresses. A common scenario is for gaming workloads, where a console needs to make a direct connection to a cloud virtual machine to minimize hops. For more information see [assigning a public IP per node](https://docs.microsoft.com/azure/aks/use-multiple-node-pools#assign-a-public-ip-per-node-for-your-node-pools). The default is false.                                                 | bool<br/><small>Optional</small>                                                                |
| enableUltraSSD            | Whether to enable UltraSSD                                                                                                                                                                                                                                                                                                                                                                                                                                                                   | bool<br/><small>Optional</small>                                                                |
| gpuInstanceProfile        | GPUInstanceProfile to be used to specify GPU MIG instance profile for supported GPU VM SKU.                                                                                                                                                                                                                                                                                                                                                                                                  | [GPUInstanceProfile_STATUS](#GPUInstanceProfile_STATUS)<br/><small>Optional</small>             |
| id                        | Resource ID.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                 | string<br/><small>Optional</small>                                                              |
| kubeletConfig             | The Kubelet configuration on the agent pool nodes.                                                                                                                                                                                                                                                                                                                                                                                                                                           | [KubeletConfig_STATUS](#KubeletConfig_STATUS)<br/><small>Optional</small>                       |
| kubeletDiskType           | Determines the placement of emptyDir volumes, container runtime data root, and Kubelet ephemeral storage.                                                                                                                                                                                                                                                                                                                                                                                    | [KubeletDiskType_STATUS](#KubeletDiskType_STATUS)<br/><small>Optional</small>                   |
| linuxOSConfig             | The OS configuration of Linux agent nodes.                                                                                                                                                                                                                                                                                                                                                                                                                                                   | [LinuxOSConfig_STATUS](#LinuxOSConfig_STATUS)<br/><small>Optional</small>                       |
| maxCount                  | The maximum number of nodes for auto-scaling                                                                                                                                                                                                                                                                                                                                                                                                                                                 | int<br/><small>Optional</small>                                                                 |
| maxPods                   | The maximum number of pods that can run on a node.                                                                                                                                                                                                                                                                                                                                                                                                                                           | int<br/><small>Optional</small>                                                                 |
| minCount                  | The minimum number of nodes for auto-scaling                                                                                                                                                                                                                                                                                                                                                                                                                                                 | int<br/><small>Optional</small>                                                                 |
| mode                      | A cluster must have at least one 'System' Agent Pool at all times. For additional information on agent pool restrictions and best practices, see: https://docs.microsoft.com/azure/aks/use-system-pools                                                                                                                                                                                                                                                                                      | [AgentPoolMode_STATUS](#AgentPoolMode_STATUS)<br/><small>Optional</small>                       |
| name                      | The name of the resource that is unique within a resource group. This name can be used to access the resource.                                                                                                                                                                                                                                                                                                                                                                               | string<br/><small>Optional</small>                                                              |
| nodeImageVersion          | The version of node image                                                                                                                                                                                                                                                                                                                                                                                                                                                                    | string<br/><small>Optional</small>                                                              |
| nodeLabels                | The node labels to be persisted across all nodes in agent pool.                                                                                                                                                                                                                                                                                                                                                                                                                              | map[string]string<br/><small>Optional</small>                                                   |
| nodePublicIPPrefixID      | This is of the form: /&ZeroWidthSpace;subscriptions/&ZeroWidthSpace;{subscriptionId}/&ZeroWidthSpace;resourceGroups/&ZeroWidthSpace;{resourceGroupName}/&ZeroWidthSpace;providers/&ZeroWidthSpace;Microsoft.Network/publicIPPrefixes/{publicIPPrefixName}                                                                                                                                                                                                                                    | string<br/><small>Optional</small>                                                              |
| nodeTaints                | The taints added to new nodes during node pool create and scale. For example, key=value:NoSchedule.                                                                                                                                                                                                                                                                                                                                                                                          | string[]<br/><small>Optional</small>                                                            |
| orchestratorVersion       | As a best practice, you should upgrade all node pools in an AKS cluster to the same Kubernetes version. The node pool version must have the same major version as the control plane. The node pool minor version must be within two minor versions of the control plane version. The node pool version cannot be greater than the control plane version. For more information see [upgrading a node pool](https://docs.microsoft.com/azure/aks/use-multiple-node-pools#upgrade-a-node-pool). | string<br/><small>Optional</small>                                                              |
| osDiskSizeGB              |                                                                                                                                                                                                                                                                                                                                                                                                                                                                                              | int<br/><small>Optional</small>                                                                 |
| osDiskType                | The default is 'Ephemeral' if the VM supports it and has a cache disk larger than the requested OSDiskSizeGB. Otherwise, defaults to 'Managed'. May not be changed after creation. For more information see [Ephemeral OS](https://docs.microsoft.com/azure/aks/cluster-configuration#ephemeral-os).                                                                                                                                                                                         | [OSDiskType_STATUS](#OSDiskType_STATUS)<br/><small>Optional</small>                             |
| osSKU                     | Specifies an OS SKU. This value must not be specified if OSType is Windows.                                                                                                                                                                                                                                                                                                                                                                                                                  | [OSSKU_STATUS](#OSSKU_STATUS)<br/><small>Optional</small>                                       |
| osType                    | The operating system type. The default is Linux.                                                                                                                                                                                                                                                                                                                                                                                                                                             | [OSType_STATUS](#OSType_STATUS)<br/><small>Optional</small>                                     |
| podSubnetID               | If omitted, pod IPs are statically assigned on the node subnet (see vnetSubnetID for more details). This is of the form: /&ZeroWidthSpace;subscriptions/&ZeroWidthSpace;{subscriptionId}/&ZeroWidthSpace;resourceGroups/&ZeroWidthSpace;{resourceGroupName}/&ZeroWidthSpace;providers/&ZeroWidthSpace;Microsoft.Network/virtualNetworks/{virtualNetworkName}/subnets/{subnetName}                                                                                                            | string<br/><small>Optional</small>                                                              |
| powerState                | Describes whether the Agent Pool is Running or Stopped                                                                                                                                                                                                                                                                                                                                                                                                                                       | [PowerState_STATUS](#PowerState_STATUS)<br/><small>Optional</small>                             |
| properties_type           | The type of Agent Pool.                                                                                                                                                                                                                                                                                                                                                                                                                                                                      | [AgentPoolType_STATUS](#AgentPoolType_STATUS)<br/><small>Optional</small>                       |
| provisioningState         | The current deployment or provisioning state.                                                                                                                                                                                                                                                                                                                                                                                                                                                | string<br/><small>Optional</small>                                                              |
| proximityPlacementGroupID | The ID for Proximity Placement Group.                                                                                                                                                                                                                                                                                                                                                                                                                                                        | string<br/><small>Optional</small>                                                              |
| scaleSetEvictionPolicy    | This cannot be specified unless the scaleSetPriority is 'Spot'. If not specified, the default is 'Delete'.                                                                                                                                                                                                                                                                                                                                                                                   | [ScaleSetEvictionPolicy_STATUS](#ScaleSetEvictionPolicy_STATUS)<br/><small>Optional</small>     |
| scaleSetPriority          | The Virtual Machine Scale Set priority. If not specified, the default is 'Regular'.                                                                                                                                                                                                                                                                                                                                                                                                          | [ScaleSetPriority_STATUS](#ScaleSetPriority_STATUS)<br/><small>Optional</small>                 |
| spotMaxPrice              | Possible values are any decimal value greater than zero or -1 which indicates the willingness to pay any on-demand price. For more details on spot pricing, see [spot VMs pricing](https://docs.microsoft.com/azure/virtual-machines/spot-vms#pricing)                                                                                                                                                                                                                                       | float64<br/><small>Optional</small>                                                             |
| tags                      | The tags to be persisted on the agent pool virtual machine scale set.                                                                                                                                                                                                                                                                                                                                                                                                                        | map[string]string<br/><small>Optional</small>                                                   |
| type                      | Resource type                                                                                                                                                                                                                                                                                                                                                                                                                                                                                | string<br/><small>Optional</small>                                                              |
| upgradeSettings           | Settings for upgrading the agentpool                                                                                                                                                                                                                                                                                                                                                                                                                                                         | [AgentPoolUpgradeSettings_STATUS](#AgentPoolUpgradeSettings_STATUS)<br/><small>Optional</small> |
| vmSize                    | VM size availability varies by region. If a node contains insufficient compute resources (memory, cpu, etc) pods might fail to run correctly. For more details on restricted VM sizes, see: https://docs.microsoft.com/azure/aks/quotas-skus-regions                                                                                                                                                                                                                                         | string<br/><small>Optional</small>                                                              |
| vnetSubnetID              | If this is not specified, a VNET and subnet will be generated and used. If no podSubnetID is specified, this applies to nodes and pods, otherwise it applies to just nodes. This is of the form: /&ZeroWidthSpace;subscriptions/&ZeroWidthSpace;{subscriptionId}/&ZeroWidthSpace;resourceGroups/&ZeroWidthSpace;{resourceGroupName}/&ZeroWidthSpace;providers/&ZeroWidthSpace;Microsoft.Network/virtualNetworks/{virtualNetworkName}/subnets/{subnetName}                                    | string<br/><small>Optional</small>                                                              |

<a id="AgentPoolMode"></a>AgentPoolMode
---------------------------------------

A cluster must have at least one 'System' Agent Pool at all times. For additional information on agent pool restrictions and best practices, see: https://docs.microsoft.com/azure/aks/use-system-pools

Used by: [ManagedClusterAgentPoolProfile](#ManagedClusterAgentPoolProfile), and [ManagedClusters_AgentPool_Spec](#ManagedClusters_AgentPool_Spec).

| Value    | Description |
|----------|-------------|
| "System" |             |
| "User"   |             |

<a id="AgentPoolMode_STATUS"></a>AgentPoolMode_STATUS
-----------------------------------------------------

A cluster must have at least one 'System' Agent Pool at all times. For additional information on agent pool restrictions and best practices, see: https://docs.microsoft.com/azure/aks/use-system-pools

Used by: [ManagedClusterAgentPoolProfile_STATUS](#ManagedClusterAgentPoolProfile_STATUS), and [ManagedClusters_AgentPool_STATUS](#ManagedClusters_AgentPool_STATUS).

| Value    | Description |
|----------|-------------|
| "System" |             |
| "User"   |             |

<a id="AgentPoolType"></a>AgentPoolType
---------------------------------------

The type of Agent Pool.

Used by: [ManagedClusterAgentPoolProfile](#ManagedClusterAgentPoolProfile), and [ManagedClusters_AgentPool_Spec](#ManagedClusters_AgentPool_Spec).

| Value                     | Description |
|---------------------------|-------------|
| "AvailabilitySet"         |             |
| "VirtualMachineScaleSets" |             |

<a id="AgentPoolType_STATUS"></a>AgentPoolType_STATUS
-----------------------------------------------------

The type of Agent Pool.

Used by: [ManagedClusterAgentPoolProfile_STATUS](#ManagedClusterAgentPoolProfile_STATUS), and [ManagedClusters_AgentPool_STATUS](#ManagedClusters_AgentPool_STATUS).

| Value                     | Description |
|---------------------------|-------------|
| "AvailabilitySet"         |             |
| "VirtualMachineScaleSets" |             |

<a id="AgentPoolUpgradeSettings"></a>AgentPoolUpgradeSettings
-------------------------------------------------------------

Settings for upgrading an agentpool

Used by: [ManagedClusterAgentPoolProfile](#ManagedClusterAgentPoolProfile), and [ManagedClusters_AgentPool_Spec](#ManagedClusters_AgentPool_Spec).

| Property | Description                                                                                                                                                                                                                                                                                                                                                                                                            | Type                               |
|----------|------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|------------------------------------|
| maxSurge | This can either be set to an integer (e.g. '5') or a percentage (e.g. '50%'). If a percentage is specified, it is the percentage of the total agent pool size at the time of the upgrade. For percentages, fractional nodes are rounded up. If not specified, the default is 1. For more information, including best practices, see: https://docs.microsoft.com/azure/aks/upgrade-cluster#customize-node-surge-upgrade | string<br/><small>Optional</small> |

<a id="AgentPoolUpgradeSettings_STATUS"></a>AgentPoolUpgradeSettings_STATUS
---------------------------------------------------------------------------

Settings for upgrading an agentpool

Used by: [ManagedClusterAgentPoolProfile_STATUS](#ManagedClusterAgentPoolProfile_STATUS), and [ManagedClusters_AgentPool_STATUS](#ManagedClusters_AgentPool_STATUS).

| Property | Description                                                                                                                                                                                                                                                                                                                                                                                                            | Type                               |
|----------|------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|------------------------------------|
| maxSurge | This can either be set to an integer (e.g. '5') or a percentage (e.g. '50%'). If a percentage is specified, it is the percentage of the total agent pool size at the time of the upgrade. For percentages, fractional nodes are rounded up. If not specified, the default is 1. For more information, including best practices, see: https://docs.microsoft.com/azure/aks/upgrade-cluster#customize-node-surge-upgrade | string<br/><small>Optional</small> |

<a id="ContainerServiceLinuxProfile"></a>ContainerServiceLinuxProfile
---------------------------------------------------------------------

Profile for Linux VMs in the container service cluster.

Used by: [ManagedCluster_Spec](#ManagedCluster_Spec).

| Property      | Description                                                 | Type                                                                                              |
|---------------|-------------------------------------------------------------|---------------------------------------------------------------------------------------------------|
| adminUsername | The administrator username to use for Linux VMs.            | string<br/><small>Required</small>                                                                |
| ssh           | The SSH configuration for Linux-based VMs running on Azure. | [ContainerServiceSshConfiguration](#ContainerServiceSshConfiguration)<br/><small>Required</small> |

<a id="ContainerServiceLinuxProfile_STATUS"></a>ContainerServiceLinuxProfile_STATUS
-----------------------------------------------------------------------------------

Profile for Linux VMs in the container service cluster.

Used by: [ManagedCluster_STATUS](#ManagedCluster_STATUS).

| Property      | Description                                                 | Type                                                                                                            |
|---------------|-------------------------------------------------------------|-----------------------------------------------------------------------------------------------------------------|
| adminUsername | The administrator username to use for Linux VMs.            | string<br/><small>Optional</small>                                                                              |
| ssh           | The SSH configuration for Linux-based VMs running on Azure. | [ContainerServiceSshConfiguration_STATUS](#ContainerServiceSshConfiguration_STATUS)<br/><small>Optional</small> |

<a id="ContainerServiceNetworkProfile"></a>ContainerServiceNetworkProfile
-------------------------------------------------------------------------

Profile of network configuration.

Used by: [ManagedCluster_Spec](#ManagedCluster_Spec).

| Property            | Description                                                                                                                                                                           | Type                                                                                                                          |
|---------------------|---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|-------------------------------------------------------------------------------------------------------------------------------|
| dnsServiceIP        | An IP address assigned to the Kubernetes DNS service. It must be within the Kubernetes service address range specified in serviceCidr.                                                | string<br/><small>Optional</small>                                                                                            |
| dockerBridgeCidr    | A CIDR notation IP range assigned to the Docker bridge network. It must not overlap with any Subnet IP ranges or the Kubernetes service address range.                                | string<br/><small>Optional</small>                                                                                            |
| loadBalancerProfile | Profile of the cluster load balancer.                                                                                                                                                 | [ManagedClusterLoadBalancerProfile](#ManagedClusterLoadBalancerProfile)<br/><small>Optional</small>                           |
| loadBalancerSku     | The default is 'standard'. See [Azure Load Balancer SKUs](https://docs.microsoft.com/azure/load-balancer/skus) for more information about the differences between load balancer SKUs. | [ContainerServiceNetworkProfile_LoadBalancerSku](#ContainerServiceNetworkProfile_LoadBalancerSku)<br/><small>Optional</small> |
| networkMode         | This cannot be specified if networkPlugin is anything other than 'azure'.                                                                                                             | [ContainerServiceNetworkProfile_NetworkMode](#ContainerServiceNetworkProfile_NetworkMode)<br/><small>Optional</small>         |
| networkPlugin       | Network plugin used for building the Kubernetes network.                                                                                                                              | [ContainerServiceNetworkProfile_NetworkPlugin](#ContainerServiceNetworkProfile_NetworkPlugin)<br/><small>Optional</small>     |
| networkPolicy       | Network policy used for building the Kubernetes network.                                                                                                                              | [ContainerServiceNetworkProfile_NetworkPolicy](#ContainerServiceNetworkProfile_NetworkPolicy)<br/><small>Optional</small>     |
| outboundType        | This can only be set at cluster creation time and cannot be changed later. For more information see [egress outbound type](https://docs.microsoft.com/azure/aks/egress-outboundtype). | [ContainerServiceNetworkProfile_OutboundType](#ContainerServiceNetworkProfile_OutboundType)<br/><small>Optional</small>       |
| podCidr             | A CIDR notation IP range from which to assign pod IPs when kubenet is used.                                                                                                           | string<br/><small>Optional</small>                                                                                            |
| serviceCidr         | A CIDR notation IP range from which to assign service cluster IPs. It must not overlap with any Subnet IP ranges.                                                                     | string<br/><small>Optional</small>                                                                                            |

<a id="ContainerServiceNetworkProfile_STATUS"></a>ContainerServiceNetworkProfile_STATUS
---------------------------------------------------------------------------------------

Profile of network configuration.

Used by: [ManagedCluster_STATUS](#ManagedCluster_STATUS).

| Property            | Description                                                                                                                                                                           | Type                                                                                                                                        |
|---------------------|---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|---------------------------------------------------------------------------------------------------------------------------------------------|
| dnsServiceIP        | An IP address assigned to the Kubernetes DNS service. It must be within the Kubernetes service address range specified in serviceCidr.                                                | string<br/><small>Optional</small>                                                                                                          |
| dockerBridgeCidr    | A CIDR notation IP range assigned to the Docker bridge network. It must not overlap with any Subnet IP ranges or the Kubernetes service address range.                                | string<br/><small>Optional</small>                                                                                                          |
| loadBalancerProfile | Profile of the cluster load balancer.                                                                                                                                                 | [ManagedClusterLoadBalancerProfile_STATUS](#ManagedClusterLoadBalancerProfile_STATUS)<br/><small>Optional</small>                           |
| loadBalancerSku     | The default is 'standard'. See [Azure Load Balancer SKUs](https://docs.microsoft.com/azure/load-balancer/skus) for more information about the differences between load balancer SKUs. | [ContainerServiceNetworkProfile_LoadBalancerSku_STATUS](#ContainerServiceNetworkProfile_LoadBalancerSku_STATUS)<br/><small>Optional</small> |
| networkMode         | This cannot be specified if networkPlugin is anything other than 'azure'.                                                                                                             | [ContainerServiceNetworkProfile_NetworkMode_STATUS](#ContainerServiceNetworkProfile_NetworkMode_STATUS)<br/><small>Optional</small>         |
| networkPlugin       | Network plugin used for building the Kubernetes network.                                                                                                                              | [ContainerServiceNetworkProfile_NetworkPlugin_STATUS](#ContainerServiceNetworkProfile_NetworkPlugin_STATUS)<br/><small>Optional</small>     |
| networkPolicy       | Network policy used for building the Kubernetes network.                                                                                                                              | [ContainerServiceNetworkProfile_NetworkPolicy_STATUS](#ContainerServiceNetworkProfile_NetworkPolicy_STATUS)<br/><small>Optional</small>     |
| outboundType        | This can only be set at cluster creation time and cannot be changed later. For more information see [egress outbound type](https://docs.microsoft.com/azure/aks/egress-outboundtype). | [ContainerServiceNetworkProfile_OutboundType_STATUS](#ContainerServiceNetworkProfile_OutboundType_STATUS)<br/><small>Optional</small>       |
| podCidr             | A CIDR notation IP range from which to assign pod IPs when kubenet is used.                                                                                                           | string<br/><small>Optional</small>                                                                                                          |
| serviceCidr         | A CIDR notation IP range from which to assign service cluster IPs. It must not overlap with any Subnet IP ranges.                                                                     | string<br/><small>Optional</small>                                                                                                          |

<a id="ContainerServiceOSDisk"></a>ContainerServiceOSDisk
---------------------------------------------------------

Used by: [ManagedClusterAgentPoolProfile](#ManagedClusterAgentPoolProfile), and [ManagedClusters_AgentPool_Spec](#ManagedClusters_AgentPool_Spec).

<a id="ExtendedLocation"></a>ExtendedLocation
---------------------------------------------

The complex type of the extended location.

Used by: [ManagedCluster_Spec](#ManagedCluster_Spec).

| Property | Description                        | Type                                                                      |
|----------|------------------------------------|---------------------------------------------------------------------------|
| name     | The name of the extended location. | string<br/><small>Optional</small>                                        |
| type     | The type of the extended location. | [ExtendedLocationType](#ExtendedLocationType)<br/><small>Optional</small> |

<a id="ExtendedLocation_STATUS"></a>ExtendedLocation_STATUS
-----------------------------------------------------------

The complex type of the extended location.

Used by: [ManagedCluster_STATUS](#ManagedCluster_STATUS).

| Property | Description                        | Type                                                                                    |
|----------|------------------------------------|-----------------------------------------------------------------------------------------|
| name     | The name of the extended location. | string<br/><small>Optional</small>                                                      |
| type     | The type of the extended location. | [ExtendedLocationType_STATUS](#ExtendedLocationType_STATUS)<br/><small>Optional</small> |

<a id="GPUInstanceProfile"></a>GPUInstanceProfile
-------------------------------------------------

GPUInstanceProfile to be used to specify GPU MIG instance profile for supported GPU VM SKU.

Used by: [ManagedClusterAgentPoolProfile](#ManagedClusterAgentPoolProfile), and [ManagedClusters_AgentPool_Spec](#ManagedClusters_AgentPool_Spec).

| Value   | Description |
|---------|-------------|
| "MIG1g" |             |
| "MIG2g" |             |
| "MIG3g" |             |
| "MIG4g" |             |
| "MIG7g" |             |

<a id="GPUInstanceProfile_STATUS"></a>GPUInstanceProfile_STATUS
---------------------------------------------------------------

GPUInstanceProfile to be used to specify GPU MIG instance profile for supported GPU VM SKU.

Used by: [ManagedClusterAgentPoolProfile_STATUS](#ManagedClusterAgentPoolProfile_STATUS), and [ManagedClusters_AgentPool_STATUS](#ManagedClusters_AgentPool_STATUS).

| Value   | Description |
|---------|-------------|
| "MIG1g" |             |
| "MIG2g" |             |
| "MIG3g" |             |
| "MIG4g" |             |
| "MIG7g" |             |

<a id="KubeletConfig"></a>KubeletConfig
---------------------------------------

See [AKS custom node configuration](https://docs.microsoft.com/azure/aks/custom-node-configuration) for more details.

Used by: [ManagedClusterAgentPoolProfile](#ManagedClusterAgentPoolProfile), and [ManagedClusters_AgentPool_Spec](#ManagedClusters_AgentPool_Spec).

| Property              | Description                                                                                                                                                                                                                          | Type                                 |
|-----------------------|--------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|--------------------------------------|
| allowedUnsafeSysctls  | Allowed list of unsafe sysctls or unsafe sysctl patterns (ending in `*`).                                                                                                                                                            | string[]<br/><small>Optional</small> |
| containerLogMaxFiles  | The maximum number of container log files that can be present for a container. The number must be ≥ 2.                                                                                                                               | int<br/><small>Optional</small>      |
| containerLogMaxSizeMB | The maximum size (e.g. 10Mi) of container log file before it is rotated.                                                                                                                                                             | int<br/><small>Optional</small>      |
| cpuCfsQuota           | The default is true.                                                                                                                                                                                                                 | bool<br/><small>Optional</small>     |
| cpuCfsQuotaPeriod     | The default is '100ms.' Valid values are a sequence of decimal numbers with an optional fraction and a unit suffix. For example: '300ms', '2h45m'. Supported units are 'ns', 'us', 'ms', 's', 'm', and 'h'.                          | string<br/><small>Optional</small>   |
| cpuManagerPolicy      | The default is 'none'. See [Kubernetes CPU management policies](https://kubernetes.io/docs/tasks/administer-cluster/cpu-management-policies/#cpu-management-policies) for more information. Allowed values are 'none' and 'static'.  | string<br/><small>Optional</small>   |
| failSwapOn            | If set to true it will make the Kubelet fail to start if swap is enabled on the node.                                                                                                                                                | bool<br/><small>Optional</small>     |
| imageGcHighThreshold  | To disable image garbage collection, set to 100. The default is 85%                                                                                                                                                                  | int<br/><small>Optional</small>      |
| imageGcLowThreshold   | This cannot be set higher than imageGcHighThreshold. The default is 80%                                                                                                                                                              | int<br/><small>Optional</small>      |
| podMaxPids            | The maximum number of processes per pod.                                                                                                                                                                                             | int<br/><small>Optional</small>      |
| topologyManagerPolicy | For more information see [Kubernetes Topology Manager](https://kubernetes.io/docs/tasks/administer-cluster/topology-manager). The default is 'none'. Allowed values are 'none', 'best-effort', 'restricted', and 'single-numa-node'. | string<br/><small>Optional</small>   |

<a id="KubeletConfig_STATUS"></a>KubeletConfig_STATUS
-----------------------------------------------------

See [AKS custom node configuration](https://docs.microsoft.com/azure/aks/custom-node-configuration) for more details.

Used by: [ManagedClusterAgentPoolProfile_STATUS](#ManagedClusterAgentPoolProfile_STATUS), and [ManagedClusters_AgentPool_STATUS](#ManagedClusters_AgentPool_STATUS).

| Property              | Description                                                                                                                                                                                                                          | Type                                 |
|-----------------------|--------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|--------------------------------------|
| allowedUnsafeSysctls  | Allowed list of unsafe sysctls or unsafe sysctl patterns (ending in `*`).                                                                                                                                                            | string[]<br/><small>Optional</small> |
| containerLogMaxFiles  | The maximum number of container log files that can be present for a container. The number must be ≥ 2.                                                                                                                               | int<br/><small>Optional</small>      |
| containerLogMaxSizeMB | The maximum size (e.g. 10Mi) of container log file before it is rotated.                                                                                                                                                             | int<br/><small>Optional</small>      |
| cpuCfsQuota           | The default is true.                                                                                                                                                                                                                 | bool<br/><small>Optional</small>     |
| cpuCfsQuotaPeriod     | The default is '100ms.' Valid values are a sequence of decimal numbers with an optional fraction and a unit suffix. For example: '300ms', '2h45m'. Supported units are 'ns', 'us', 'ms', 's', 'm', and 'h'.                          | string<br/><small>Optional</small>   |
| cpuManagerPolicy      | The default is 'none'. See [Kubernetes CPU management policies](https://kubernetes.io/docs/tasks/administer-cluster/cpu-management-policies/#cpu-management-policies) for more information. Allowed values are 'none' and 'static'.  | string<br/><small>Optional</small>   |
| failSwapOn            | If set to true it will make the Kubelet fail to start if swap is enabled on the node.                                                                                                                                                | bool<br/><small>Optional</small>     |
| imageGcHighThreshold  | To disable image garbage collection, set to 100. The default is 85%                                                                                                                                                                  | int<br/><small>Optional</small>      |
| imageGcLowThreshold   | This cannot be set higher than imageGcHighThreshold. The default is 80%                                                                                                                                                              | int<br/><small>Optional</small>      |
| podMaxPids            | The maximum number of processes per pod.                                                                                                                                                                                             | int<br/><small>Optional</small>      |
| topologyManagerPolicy | For more information see [Kubernetes Topology Manager](https://kubernetes.io/docs/tasks/administer-cluster/topology-manager). The default is 'none'. Allowed values are 'none', 'best-effort', 'restricted', and 'single-numa-node'. | string<br/><small>Optional</small>   |

<a id="KubeletDiskType"></a>KubeletDiskType
-------------------------------------------

Determines the placement of emptyDir volumes, container runtime data root, and Kubelet ephemeral storage.

Used by: [ManagedClusterAgentPoolProfile](#ManagedClusterAgentPoolProfile), and [ManagedClusters_AgentPool_Spec](#ManagedClusters_AgentPool_Spec).

| Value       | Description |
|-------------|-------------|
| "OS"        |             |
| "Temporary" |             |

<a id="KubeletDiskType_STATUS"></a>KubeletDiskType_STATUS
---------------------------------------------------------

Determines the placement of emptyDir volumes, container runtime data root, and Kubelet ephemeral storage.

Used by: [ManagedClusterAgentPoolProfile_STATUS](#ManagedClusterAgentPoolProfile_STATUS), and [ManagedClusters_AgentPool_STATUS](#ManagedClusters_AgentPool_STATUS).

| Value       | Description |
|-------------|-------------|
| "OS"        |             |
| "Temporary" |             |

<a id="LinuxOSConfig"></a>LinuxOSConfig
---------------------------------------

See [AKS custom node configuration](https://docs.microsoft.com/azure/aks/custom-node-configuration) for more details.

Used by: [ManagedClusterAgentPoolProfile](#ManagedClusterAgentPoolProfile), and [ManagedClusters_AgentPool_Spec](#ManagedClusters_AgentPool_Spec).

| Property                   | Description                                                                                                                                                                                                                                         | Type                                                      |
|----------------------------|-----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|-----------------------------------------------------------|
| swapFileSizeMB             | The size in MB of a swap file that will be created on each node.                                                                                                                                                                                    | int<br/><small>Optional</small>                           |
| sysctls                    | Sysctl settings for Linux agent nodes.                                                                                                                                                                                                              | [SysctlConfig](#SysctlConfig)<br/><small>Optional</small> |
| transparentHugePageDefrag  | Valid values are 'always', 'defer', 'defer+madvise', 'madvise' and 'never'. The default is 'madvise'. For more information see [Transparent Hugepages](https://www.kernel.org/doc/html/latest/admin-guide/mm/transhuge.html#admin-guide-transhuge). | string<br/><small>Optional</small>                        |
| transparentHugePageEnabled | Valid values are 'always', 'madvise', and 'never'. The default is 'always'. For more information see [Transparent Hugepages](https://www.kernel.org/doc/html/latest/admin-guide/mm/transhuge.html#admin-guide-transhuge).                           | string<br/><small>Optional</small>                        |

<a id="LinuxOSConfig_STATUS"></a>LinuxOSConfig_STATUS
-----------------------------------------------------

See [AKS custom node configuration](https://docs.microsoft.com/azure/aks/custom-node-configuration) for more details.

Used by: [ManagedClusterAgentPoolProfile_STATUS](#ManagedClusterAgentPoolProfile_STATUS), and [ManagedClusters_AgentPool_STATUS](#ManagedClusters_AgentPool_STATUS).

| Property                   | Description                                                                                                                                                                                                                                         | Type                                                                    |
|----------------------------|-----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|-------------------------------------------------------------------------|
| swapFileSizeMB             | The size in MB of a swap file that will be created on each node.                                                                                                                                                                                    | int<br/><small>Optional</small>                                         |
| sysctls                    | Sysctl settings for Linux agent nodes.                                                                                                                                                                                                              | [SysctlConfig_STATUS](#SysctlConfig_STATUS)<br/><small>Optional</small> |
| transparentHugePageDefrag  | Valid values are 'always', 'defer', 'defer+madvise', 'madvise' and 'never'. The default is 'madvise'. For more information see [Transparent Hugepages](https://www.kernel.org/doc/html/latest/admin-guide/mm/transhuge.html#admin-guide-transhuge). | string<br/><small>Optional</small>                                      |
| transparentHugePageEnabled | Valid values are 'always', 'madvise', and 'never'. The default is 'always'. For more information see [Transparent Hugepages](https://www.kernel.org/doc/html/latest/admin-guide/mm/transhuge.html#admin-guide-transhuge).                           | string<br/><small>Optional</small>                                      |

<a id="ManagedClusterAADProfile"></a>ManagedClusterAADProfile
-------------------------------------------------------------

For more details see [managed AAD on AKS](https://docs.microsoft.com/azure/aks/managed-aad).

Used by: [ManagedCluster_Spec](#ManagedCluster_Spec).

| Property            | Description                                                                                                        | Type                                 |
|---------------------|--------------------------------------------------------------------------------------------------------------------|--------------------------------------|
| adminGroupObjectIDs | The list of AAD group object IDs that will have admin role of the cluster.                                         | string[]<br/><small>Optional</small> |
| clientAppID         | The client AAD application ID.                                                                                     | string<br/><small>Optional</small>   |
| enableAzureRBAC     | Whether to enable Azure RBAC for Kubernetes authorization.                                                         | bool<br/><small>Optional</small>     |
| managed             | Whether to enable managed AAD.                                                                                     | bool<br/><small>Optional</small>     |
| serverAppID         | The server AAD application ID.                                                                                     | string<br/><small>Optional</small>   |
| serverAppSecret     | The server AAD application secret.                                                                                 | string<br/><small>Optional</small>   |
| tenantID            | The AAD tenant ID to use for authentication. If not specified, will use the tenant of the deployment subscription. | string<br/><small>Optional</small>   |

<a id="ManagedClusterAADProfile_STATUS"></a>ManagedClusterAADProfile_STATUS
---------------------------------------------------------------------------

For more details see [managed AAD on AKS](https://docs.microsoft.com/azure/aks/managed-aad).

Used by: [ManagedCluster_STATUS](#ManagedCluster_STATUS).

| Property            | Description                                                                                                        | Type                                 |
|---------------------|--------------------------------------------------------------------------------------------------------------------|--------------------------------------|
| adminGroupObjectIDs | The list of AAD group object IDs that will have admin role of the cluster.                                         | string[]<br/><small>Optional</small> |
| clientAppID         | The client AAD application ID.                                                                                     | string<br/><small>Optional</small>   |
| enableAzureRBAC     | Whether to enable Azure RBAC for Kubernetes authorization.                                                         | bool<br/><small>Optional</small>     |
| managed             | Whether to enable managed AAD.                                                                                     | bool<br/><small>Optional</small>     |
| serverAppID         | The server AAD application ID.                                                                                     | string<br/><small>Optional</small>   |
| serverAppSecret     | The server AAD application secret.                                                                                 | string<br/><small>Optional</small>   |
| tenantID            | The AAD tenant ID to use for authentication. If not specified, will use the tenant of the deployment subscription. | string<br/><small>Optional</small>   |

<a id="ManagedClusterAddonProfile"></a>ManagedClusterAddonProfile
-----------------------------------------------------------------

A Kubernetes add-on profile for a managed cluster.

Used by: [ManagedCluster_Spec](#ManagedCluster_Spec).

| Property | Description                                | Type                                          |
|----------|--------------------------------------------|-----------------------------------------------|
| config   | Key-value pairs for configuring an add-on. | map[string]string<br/><small>Optional</small> |
| enabled  | Whether the add-on is enabled or not.      | bool<br/><small>Required</small>              |

<a id="ManagedClusterAddonProfile_STATUS"></a>ManagedClusterAddonProfile_STATUS
-------------------------------------------------------------------------------

A Kubernetes add-on profile for a managed cluster.

Used by: [ManagedCluster_STATUS](#ManagedCluster_STATUS).

| Property | Description                                                | Type                                                                                    |
|----------|------------------------------------------------------------|-----------------------------------------------------------------------------------------|
| config   | Key-value pairs for configuring an add-on.                 | map[string]string<br/><small>Optional</small>                                           |
| enabled  | Whether the add-on is enabled or not.                      | bool<br/><small>Optional</small>                                                        |
| identity | Information of user assigned identity used by this add-on. | [UserAssignedIdentity_STATUS](#UserAssignedIdentity_STATUS)<br/><small>Optional</small> |

<a id="ManagedClusterAgentPoolProfile"></a>ManagedClusterAgentPoolProfile
-------------------------------------------------------------------------

Profile for the container service agent pool.

Used by: [ManagedCluster_Spec](#ManagedCluster_Spec).

| Property                      | Description                                                                                                                                                                                                                                                                                                                                                                                                                                                                                  | Type                                                                              |
|-------------------------------|----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|-----------------------------------------------------------------------------------|
| availabilityZones             | The list of Availability zones to use for nodes. This can only be specified if the AgentPoolType property is 'VirtualMachineScaleSets'.                                                                                                                                                                                                                                                                                                                                                      | string[]<br/><small>Optional</small>                                              |
| count                         | Number of agents (VMs) to host docker containers. Allowed values must be in the range of 0 to 1000 (inclusive) for user pools and in the range of 1 to 1000 (inclusive) for system pools. The default value is 1.                                                                                                                                                                                                                                                                            | int<br/><small>Optional</small>                                                   |
| enableAutoScaling             | Whether to enable auto-scaler                                                                                                                                                                                                                                                                                                                                                                                                                                                                | bool<br/><small>Optional</small>                                                  |
| enableEncryptionAtHost        | This is only supported on certain VM sizes and in certain Azure regions. For more information, see: https://docs.microsoft.com/azure/aks/enable-host-encryption                                                                                                                                                                                                                                                                                                                              | bool<br/><small>Optional</small>                                                  |
| enableFIPS                    | See [Add a FIPS-enabled node pool](https://docs.microsoft.com/azure/aks/use-multiple-node-pools#add-a-fips-enabled-node-pool-preview) for more details.                                                                                                                                                                                                                                                                                                                                      | bool<br/><small>Optional</small>                                                  |
| enableNodePublicIP            | Some scenarios may require nodes in a node pool to receive their own dedicated public IP addresses. A common scenario is for gaming workloads, where a console needs to make a direct connection to a cloud virtual machine to minimize hops. For more information see [assigning a public IP per node](https://docs.microsoft.com/azure/aks/use-multiple-node-pools#assign-a-public-ip-per-node-for-your-node-pools). The default is false.                                                 | bool<br/><small>Optional</small>                                                  |
| enableUltraSSD                | Whether to enable UltraSSD                                                                                                                                                                                                                                                                                                                                                                                                                                                                   | bool<br/><small>Optional</small>                                                  |
| gpuInstanceProfile            | GPUInstanceProfile to be used to specify GPU MIG instance profile for supported GPU VM SKU.                                                                                                                                                                                                                                                                                                                                                                                                  | [GPUInstanceProfile](#GPUInstanceProfile)<br/><small>Optional</small>             |
| kubeletConfig                 | The Kubelet configuration on the agent pool nodes.                                                                                                                                                                                                                                                                                                                                                                                                                                           | [KubeletConfig](#KubeletConfig)<br/><small>Optional</small>                       |
| kubeletDiskType               | Determines the placement of emptyDir volumes, container runtime data root, and Kubelet ephemeral storage.                                                                                                                                                                                                                                                                                                                                                                                    | [KubeletDiskType](#KubeletDiskType)<br/><small>Optional</small>                   |
| linuxOSConfig                 | The OS configuration of Linux agent nodes.                                                                                                                                                                                                                                                                                                                                                                                                                                                   | [LinuxOSConfig](#LinuxOSConfig)<br/><small>Optional</small>                       |
| maxCount                      | The maximum number of nodes for auto-scaling                                                                                                                                                                                                                                                                                                                                                                                                                                                 | int<br/><small>Optional</small>                                                   |
| maxPods                       | The maximum number of pods that can run on a node.                                                                                                                                                                                                                                                                                                                                                                                                                                           | int<br/><small>Optional</small>                                                   |
| minCount                      | The minimum number of nodes for auto-scaling                                                                                                                                                                                                                                                                                                                                                                                                                                                 | int<br/><small>Optional</small>                                                   |
| mode                          | A cluster must have at least one 'System' Agent Pool at all times. For additional information on agent pool restrictions and best practices, see: https://docs.microsoft.com/azure/aks/use-system-pools                                                                                                                                                                                                                                                                                      | [AgentPoolMode](#AgentPoolMode)<br/><small>Optional</small>                       |
| name                          | Windows agent pool names must be 6 characters or less.                                                                                                                                                                                                                                                                                                                                                                                                                                       | string<br/><small>Optional</small>                                                |
| nodeLabels                    | The node labels to be persisted across all nodes in agent pool.                                                                                                                                                                                                                                                                                                                                                                                                                              | map[string]string<br/><small>Optional</small>                                     |
| nodePublicIPPrefixIDReference | This is of the form: /&ZeroWidthSpace;subscriptions/&ZeroWidthSpace;{subscriptionId}/&ZeroWidthSpace;resourceGroups/&ZeroWidthSpace;{resourceGroupName}/&ZeroWidthSpace;providers/&ZeroWidthSpace;Microsoft.Network/publicIPPrefixes/{publicIPPrefixName}                                                                                                                                                                                                                                    | genruntime.ResourceReference<br/><small>Optional</small>                          |
| nodeTaints                    | The taints added to new nodes during node pool create and scale. For example, key=value:NoSchedule.                                                                                                                                                                                                                                                                                                                                                                                          | string[]<br/><small>Optional</small>                                              |
| orchestratorVersion           | As a best practice, you should upgrade all node pools in an AKS cluster to the same Kubernetes version. The node pool version must have the same major version as the control plane. The node pool minor version must be within two minor versions of the control plane version. The node pool version cannot be greater than the control plane version. For more information see [upgrading a node pool](https://docs.microsoft.com/azure/aks/use-multiple-node-pools#upgrade-a-node-pool). | string<br/><small>Optional</small>                                                |
| osDiskSizeGB                  |                                                                                                                                                                                                                                                                                                                                                                                                                                                                                              | [ContainerServiceOSDisk](#ContainerServiceOSDisk)<br/><small>Optional</small>     |
| osDiskType                    | The default is 'Ephemeral' if the VM supports it and has a cache disk larger than the requested OSDiskSizeGB. Otherwise, defaults to 'Managed'. May not be changed after creation. For more information see [Ephemeral OS](https://docs.microsoft.com/azure/aks/cluster-configuration#ephemeral-os).                                                                                                                                                                                         | [OSDiskType](#OSDiskType)<br/><small>Optional</small>                             |
| osSKU                         | Specifies an OS SKU. This value must not be specified if OSType is Windows.                                                                                                                                                                                                                                                                                                                                                                                                                  | [OSSKU](#OSSKU)<br/><small>Optional</small>                                       |
| osType                        | The operating system type. The default is Linux.                                                                                                                                                                                                                                                                                                                                                                                                                                             | [OSType](#OSType)<br/><small>Optional</small>                                     |
| podSubnetIDReference          | If omitted, pod IPs are statically assigned on the node subnet (see vnetSubnetID for more details). This is of the form: /&ZeroWidthSpace;subscriptions/&ZeroWidthSpace;{subscriptionId}/&ZeroWidthSpace;resourceGroups/&ZeroWidthSpace;{resourceGroupName}/&ZeroWidthSpace;providers/&ZeroWidthSpace;Microsoft.Network/virtualNetworks/{virtualNetworkName}/subnets/{subnetName}                                                                                                            | genruntime.ResourceReference<br/><small>Optional</small>                          |
| proximityPlacementGroupID     | The ID for Proximity Placement Group.                                                                                                                                                                                                                                                                                                                                                                                                                                                        | string<br/><small>Optional</small>                                                |
| scaleSetEvictionPolicy        | This cannot be specified unless the scaleSetPriority is 'Spot'. If not specified, the default is 'Delete'.                                                                                                                                                                                                                                                                                                                                                                                   | [ScaleSetEvictionPolicy](#ScaleSetEvictionPolicy)<br/><small>Optional</small>     |
| scaleSetPriority              | The Virtual Machine Scale Set priority. If not specified, the default is 'Regular'.                                                                                                                                                                                                                                                                                                                                                                                                          | [ScaleSetPriority](#ScaleSetPriority)<br/><small>Optional</small>                 |
| spotMaxPrice                  | Possible values are any decimal value greater than zero or -1 which indicates the willingness to pay any on-demand price. For more details on spot pricing, see [spot VMs pricing](https://docs.microsoft.com/azure/virtual-machines/spot-vms#pricing)                                                                                                                                                                                                                                       | float64<br/><small>Optional</small>                                               |
| tags                          | The tags to be persisted on the agent pool virtual machine scale set.                                                                                                                                                                                                                                                                                                                                                                                                                        | map[string]string<br/><small>Optional</small>                                     |
| type                          | The type of Agent Pool.                                                                                                                                                                                                                                                                                                                                                                                                                                                                      | [AgentPoolType](#AgentPoolType)<br/><small>Optional</small>                       |
| upgradeSettings               | Settings for upgrading the agentpool                                                                                                                                                                                                                                                                                                                                                                                                                                                         | [AgentPoolUpgradeSettings](#AgentPoolUpgradeSettings)<br/><small>Optional</small> |
| vmSize                        | VM size availability varies by region. If a node contains insufficient compute resources (memory, cpu, etc) pods might fail to run correctly. For more details on restricted VM sizes, see: https://docs.microsoft.com/azure/aks/quotas-skus-regions                                                                                                                                                                                                                                         | string<br/><small>Optional</small>                                                |
| vnetSubnetIDReference         | If this is not specified, a VNET and subnet will be generated and used. If no podSubnetID is specified, this applies to nodes and pods, otherwise it applies to just nodes. This is of the form: /&ZeroWidthSpace;subscriptions/&ZeroWidthSpace;{subscriptionId}/&ZeroWidthSpace;resourceGroups/&ZeroWidthSpace;{resourceGroupName}/&ZeroWidthSpace;providers/&ZeroWidthSpace;Microsoft.Network/virtualNetworks/{virtualNetworkName}/subnets/{subnetName}                                    | genruntime.ResourceReference<br/><small>Optional</small>                          |

<a id="ManagedClusterAgentPoolProfile_STATUS"></a>ManagedClusterAgentPoolProfile_STATUS
---------------------------------------------------------------------------------------

Profile for the container service agent pool.

Used by: [ManagedCluster_STATUS](#ManagedCluster_STATUS).

| Property                  | Description                                                                                                                                                                                                                                                                                                                                                                                                                                                                                  | Type                                                                                            |
|---------------------------|----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|-------------------------------------------------------------------------------------------------|
| availabilityZones         | The list of Availability zones to use for nodes. This can only be specified if the AgentPoolType property is 'VirtualMachineScaleSets'.                                                                                                                                                                                                                                                                                                                                                      | string[]<br/><small>Optional</small>                                                            |
| count                     | Number of agents (VMs) to host docker containers. Allowed values must be in the range of 0 to 1000 (inclusive) for user pools and in the range of 1 to 1000 (inclusive) for system pools. The default value is 1.                                                                                                                                                                                                                                                                            | int<br/><small>Optional</small>                                                                 |
| enableAutoScaling         | Whether to enable auto-scaler                                                                                                                                                                                                                                                                                                                                                                                                                                                                | bool<br/><small>Optional</small>                                                                |
| enableEncryptionAtHost    | This is only supported on certain VM sizes and in certain Azure regions. For more information, see: https://docs.microsoft.com/azure/aks/enable-host-encryption                                                                                                                                                                                                                                                                                                                              | bool<br/><small>Optional</small>                                                                |
| enableFIPS                | See [Add a FIPS-enabled node pool](https://docs.microsoft.com/azure/aks/use-multiple-node-pools#add-a-fips-enabled-node-pool-preview) for more details.                                                                                                                                                                                                                                                                                                                                      | bool<br/><small>Optional</small>                                                                |
| enableNodePublicIP        | Some scenarios may require nodes in a node pool to receive their own dedicated public IP addresses. A common scenario is for gaming workloads, where a console needs to make a direct connection to a cloud virtual machine to minimize hops. For more information see [assigning a public IP per node](https://docs.microsoft.com/azure/aks/use-multiple-node-pools#assign-a-public-ip-per-node-for-your-node-pools). The default is false.                                                 | bool<br/><small>Optional</small>                                                                |
| enableUltraSSD            | Whether to enable UltraSSD                                                                                                                                                                                                                                                                                                                                                                                                                                                                   | bool<br/><small>Optional</small>                                                                |
| gpuInstanceProfile        | GPUInstanceProfile to be used to specify GPU MIG instance profile for supported GPU VM SKU.                                                                                                                                                                                                                                                                                                                                                                                                  | [GPUInstanceProfile_STATUS](#GPUInstanceProfile_STATUS)<br/><small>Optional</small>             |
| kubeletConfig             | The Kubelet configuration on the agent pool nodes.                                                                                                                                                                                                                                                                                                                                                                                                                                           | [KubeletConfig_STATUS](#KubeletConfig_STATUS)<br/><small>Optional</small>                       |
| kubeletDiskType           | Determines the placement of emptyDir volumes, container runtime data root, and Kubelet ephemeral storage.                                                                                                                                                                                                                                                                                                                                                                                    | [KubeletDiskType_STATUS](#KubeletDiskType_STATUS)<br/><small>Optional</small>                   |
| linuxOSConfig             | The OS configuration of Linux agent nodes.                                                                                                                                                                                                                                                                                                                                                                                                                                                   | [LinuxOSConfig_STATUS](#LinuxOSConfig_STATUS)<br/><small>Optional</small>                       |
| maxCount                  | The maximum number of nodes for auto-scaling                                                                                                                                                                                                                                                                                                                                                                                                                                                 | int<br/><small>Optional</small>                                                                 |
| maxPods                   | The maximum number of pods that can run on a node.                                                                                                                                                                                                                                                                                                                                                                                                                                           | int<br/><small>Optional</small>                                                                 |
| minCount                  | The minimum number of nodes for auto-scaling                                                                                                                                                                                                                                                                                                                                                                                                                                                 | int<br/><small>Optional</small>                                                                 |
| mode                      | A cluster must have at least one 'System' Agent Pool at all times. For additional information on agent pool restrictions and best practices, see: https://docs.microsoft.com/azure/aks/use-system-pools                                                                                                                                                                                                                                                                                      | [AgentPoolMode_STATUS](#AgentPoolMode_STATUS)<br/><small>Optional</small>                       |
| name                      | Windows agent pool names must be 6 characters or less.                                                                                                                                                                                                                                                                                                                                                                                                                                       | string<br/><small>Optional</small>                                                              |
| nodeImageVersion          | The version of node image                                                                                                                                                                                                                                                                                                                                                                                                                                                                    | string<br/><small>Optional</small>                                                              |
| nodeLabels                | The node labels to be persisted across all nodes in agent pool.                                                                                                                                                                                                                                                                                                                                                                                                                              | map[string]string<br/><small>Optional</small>                                                   |
| nodePublicIPPrefixID      | This is of the form: /&ZeroWidthSpace;subscriptions/&ZeroWidthSpace;{subscriptionId}/&ZeroWidthSpace;resourceGroups/&ZeroWidthSpace;{resourceGroupName}/&ZeroWidthSpace;providers/&ZeroWidthSpace;Microsoft.Network/publicIPPrefixes/{publicIPPrefixName}                                                                                                                                                                                                                                    | string<br/><small>Optional</small>                                                              |
| nodeTaints                | The taints added to new nodes during node pool create and scale. For example, key=value:NoSchedule.                                                                                                                                                                                                                                                                                                                                                                                          | string[]<br/><small>Optional</small>                                                            |
| orchestratorVersion       | As a best practice, you should upgrade all node pools in an AKS cluster to the same Kubernetes version. The node pool version must have the same major version as the control plane. The node pool minor version must be within two minor versions of the control plane version. The node pool version cannot be greater than the control plane version. For more information see [upgrading a node pool](https://docs.microsoft.com/azure/aks/use-multiple-node-pools#upgrade-a-node-pool). | string<br/><small>Optional</small>                                                              |
| osDiskSizeGB              |                                                                                                                                                                                                                                                                                                                                                                                                                                                                                              | int<br/><small>Optional</small>                                                                 |
| osDiskType                | The default is 'Ephemeral' if the VM supports it and has a cache disk larger than the requested OSDiskSizeGB. Otherwise, defaults to 'Managed'. May not be changed after creation. For more information see [Ephemeral OS](https://docs.microsoft.com/azure/aks/cluster-configuration#ephemeral-os).                                                                                                                                                                                         | [OSDiskType_STATUS](#OSDiskType_STATUS)<br/><small>Optional</small>                             |
| osSKU                     | Specifies an OS SKU. This value must not be specified if OSType is Windows.                                                                                                                                                                                                                                                                                                                                                                                                                  | [OSSKU_STATUS](#OSSKU_STATUS)<br/><small>Optional</small>                                       |
| osType                    | The operating system type. The default is Linux.                                                                                                                                                                                                                                                                                                                                                                                                                                             | [OSType_STATUS](#OSType_STATUS)<br/><small>Optional</small>                                     |
| podSubnetID               | If omitted, pod IPs are statically assigned on the node subnet (see vnetSubnetID for more details). This is of the form: /&ZeroWidthSpace;subscriptions/&ZeroWidthSpace;{subscriptionId}/&ZeroWidthSpace;resourceGroups/&ZeroWidthSpace;{resourceGroupName}/&ZeroWidthSpace;providers/&ZeroWidthSpace;Microsoft.Network/virtualNetworks/{virtualNetworkName}/subnets/{subnetName}                                                                                                            | string<br/><small>Optional</small>                                                              |
| powerState                | Describes whether the Agent Pool is Running or Stopped                                                                                                                                                                                                                                                                                                                                                                                                                                       | [PowerState_STATUS](#PowerState_STATUS)<br/><small>Optional</small>                             |
| provisioningState         | The current deployment or provisioning state.                                                                                                                                                                                                                                                                                                                                                                                                                                                | string<br/><small>Optional</small>                                                              |
| proximityPlacementGroupID | The ID for Proximity Placement Group.                                                                                                                                                                                                                                                                                                                                                                                                                                                        | string<br/><small>Optional</small>                                                              |
| scaleSetEvictionPolicy    | This cannot be specified unless the scaleSetPriority is 'Spot'. If not specified, the default is 'Delete'.                                                                                                                                                                                                                                                                                                                                                                                   | [ScaleSetEvictionPolicy_STATUS](#ScaleSetEvictionPolicy_STATUS)<br/><small>Optional</small>     |
| scaleSetPriority          | The Virtual Machine Scale Set priority. If not specified, the default is 'Regular'.                                                                                                                                                                                                                                                                                                                                                                                                          | [ScaleSetPriority_STATUS](#ScaleSetPriority_STATUS)<br/><small>Optional</small>                 |
| spotMaxPrice              | Possible values are any decimal value greater than zero or -1 which indicates the willingness to pay any on-demand price. For more details on spot pricing, see [spot VMs pricing](https://docs.microsoft.com/azure/virtual-machines/spot-vms#pricing)                                                                                                                                                                                                                                       | float64<br/><small>Optional</small>                                                             |
| tags                      | The tags to be persisted on the agent pool virtual machine scale set.                                                                                                                                                                                                                                                                                                                                                                                                                        | map[string]string<br/><small>Optional</small>                                                   |
| type                      | The type of Agent Pool.                                                                                                                                                                                                                                                                                                                                                                                                                                                                      | [AgentPoolType_STATUS](#AgentPoolType_STATUS)<br/><small>Optional</small>                       |
| upgradeSettings           | Settings for upgrading the agentpool                                                                                                                                                                                                                                                                                                                                                                                                                                                         | [AgentPoolUpgradeSettings_STATUS](#AgentPoolUpgradeSettings_STATUS)<br/><small>Optional</small> |
| vmSize                    | VM size availability varies by region. If a node contains insufficient compute resources (memory, cpu, etc) pods might fail to run correctly. For more details on restricted VM sizes, see: https://docs.microsoft.com/azure/aks/quotas-skus-regions                                                                                                                                                                                                                                         | string<br/><small>Optional</small>                                                              |
| vnetSubnetID              | If this is not specified, a VNET and subnet will be generated and used. If no podSubnetID is specified, this applies to nodes and pods, otherwise it applies to just nodes. This is of the form: /&ZeroWidthSpace;subscriptions/&ZeroWidthSpace;{subscriptionId}/&ZeroWidthSpace;resourceGroups/&ZeroWidthSpace;{resourceGroupName}/&ZeroWidthSpace;providers/&ZeroWidthSpace;Microsoft.Network/virtualNetworks/{virtualNetworkName}/subnets/{subnetName}                                    | string<br/><small>Optional</small>                                                              |

<a id="ManagedClusterAPIServerAccessProfile"></a>ManagedClusterAPIServerAccessProfile
-------------------------------------------------------------------------------------

Access profile for managed cluster API server.

Used by: [ManagedCluster_Spec](#ManagedCluster_Spec).

| Property                       | Description                                                                                                                                                                                                                                                                                                                   | Type                                 |
|--------------------------------|-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|--------------------------------------|
| authorizedIPRanges             | IP ranges are specified in CIDR format, e.g. 137.117.106.88/29. This feature is not compatible with clusters that use Public IP Per Node, or clusters that are using a Basic Load Balancer. For more information see [API server authorized IP ranges](https://docs.microsoft.com/azure/aks/api-server-authorized-ip-ranges). | string[]<br/><small>Optional</small> |
| enablePrivateCluster           | For more details, see [Creating a private AKS cluster](https://docs.microsoft.com/azure/aks/private-clusters).                                                                                                                                                                                                                | bool<br/><small>Optional</small>     |
| enablePrivateClusterPublicFQDN | Whether to create additional public FQDN for private cluster or not.                                                                                                                                                                                                                                                          | bool<br/><small>Optional</small>     |
| privateDNSZone                 | The default is System. For more details see [configure private DNS zone](https://docs.microsoft.com/azure/aks/private-clusters#configure-private-dns-zone). Allowed values are 'system' and 'none'.                                                                                                                           | string<br/><small>Optional</small>   |

<a id="ManagedClusterAPIServerAccessProfile_STATUS"></a>ManagedClusterAPIServerAccessProfile_STATUS
---------------------------------------------------------------------------------------------------

Access profile for managed cluster API server.

Used by: [ManagedCluster_STATUS](#ManagedCluster_STATUS).

| Property                       | Description                                                                                                                                                                                                                                                                                                                   | Type                                 |
|--------------------------------|-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|--------------------------------------|
| authorizedIPRanges             | IP ranges are specified in CIDR format, e.g. 137.117.106.88/29. This feature is not compatible with clusters that use Public IP Per Node, or clusters that are using a Basic Load Balancer. For more information see [API server authorized IP ranges](https://docs.microsoft.com/azure/aks/api-server-authorized-ip-ranges). | string[]<br/><small>Optional</small> |
| enablePrivateCluster           | For more details, see [Creating a private AKS cluster](https://docs.microsoft.com/azure/aks/private-clusters).                                                                                                                                                                                                                | bool<br/><small>Optional</small>     |
| enablePrivateClusterPublicFQDN | Whether to create additional public FQDN for private cluster or not.                                                                                                                                                                                                                                                          | bool<br/><small>Optional</small>     |
| privateDNSZone                 | The default is System. For more details see [configure private DNS zone](https://docs.microsoft.com/azure/aks/private-clusters#configure-private-dns-zone). Allowed values are 'system' and 'none'.                                                                                                                           | string<br/><small>Optional</small>   |

<a id="ManagedClusterAutoUpgradeProfile"></a>ManagedClusterAutoUpgradeProfile
-----------------------------------------------------------------------------

Auto upgrade profile for a managed cluster.

Used by: [ManagedCluster_Spec](#ManagedCluster_Spec).

| Property       | Description                                                                                                                                             | Type                                                                                                                            |
|----------------|---------------------------------------------------------------------------------------------------------------------------------------------------------|---------------------------------------------------------------------------------------------------------------------------------|
| upgradeChannel | For more information see [setting the AKS cluster auto-upgrade channel](https://docs.microsoft.com/azure/aks/upgrade-cluster#set-auto-upgrade-channel). | [ManagedClusterAutoUpgradeProfile_UpgradeChannel](#ManagedClusterAutoUpgradeProfile_UpgradeChannel)<br/><small>Optional</small> |

<a id="ManagedClusterAutoUpgradeProfile_STATUS"></a>ManagedClusterAutoUpgradeProfile_STATUS
-------------------------------------------------------------------------------------------

Auto upgrade profile for a managed cluster.

Used by: [ManagedCluster_STATUS](#ManagedCluster_STATUS).

| Property       | Description                                                                                                                                             | Type                                                                                                                                          |
|----------------|---------------------------------------------------------------------------------------------------------------------------------------------------------|-----------------------------------------------------------------------------------------------------------------------------------------------|
| upgradeChannel | For more information see [setting the AKS cluster auto-upgrade channel](https://docs.microsoft.com/azure/aks/upgrade-cluster#set-auto-upgrade-channel). | [ManagedClusterAutoUpgradeProfile_UpgradeChannel_STATUS](#ManagedClusterAutoUpgradeProfile_UpgradeChannel_STATUS)<br/><small>Optional</small> |

<a id="ManagedClusterHTTPProxyConfig"></a>ManagedClusterHTTPProxyConfig
-----------------------------------------------------------------------

Cluster HTTP proxy configuration.

Used by: [ManagedCluster_Spec](#ManagedCluster_Spec).

| Property   | Description                                                 | Type                                 |
|------------|-------------------------------------------------------------|--------------------------------------|
| httpProxy  | The HTTP proxy server endpoint to use.                      | string<br/><small>Optional</small>   |
| httpsProxy | The HTTPS proxy server endpoint to use.                     | string<br/><small>Optional</small>   |
| noProxy    | The endpoints that should not go through proxy.             | string[]<br/><small>Optional</small> |
| trustedCa  | Alternative CA cert to use for connecting to proxy servers. | string<br/><small>Optional</small>   |

<a id="ManagedClusterHTTPProxyConfig_STATUS"></a>ManagedClusterHTTPProxyConfig_STATUS
-------------------------------------------------------------------------------------

Cluster HTTP proxy configuration.

Used by: [ManagedCluster_STATUS](#ManagedCluster_STATUS).

| Property   | Description                                                 | Type                                 |
|------------|-------------------------------------------------------------|--------------------------------------|
| httpProxy  | The HTTP proxy server endpoint to use.                      | string<br/><small>Optional</small>   |
| httpsProxy | The HTTPS proxy server endpoint to use.                     | string<br/><small>Optional</small>   |
| noProxy    | The endpoints that should not go through proxy.             | string[]<br/><small>Optional</small> |
| trustedCa  | Alternative CA cert to use for connecting to proxy servers. | string<br/><small>Optional</small>   |

<a id="ManagedClusterIdentity"></a>ManagedClusterIdentity
---------------------------------------------------------

Identity for the managed cluster.

Used by: [ManagedCluster_Spec](#ManagedCluster_Spec).

| Property               | Description                                                                                                                                                                                                                                                                                    | Type                                                                                      |
|------------------------|------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|-------------------------------------------------------------------------------------------|
| type                   | For more information see [use managed identities in AKS](https://docs.microsoft.com/azure/aks/use-managed-identity).                                                                                                                                                                           | [ManagedClusterIdentity_Type](#ManagedClusterIdentity_Type)<br/><small>Optional</small>   |
| userAssignedIdentities | The keys must be ARM resource IDs in the form: '/&ZeroWidthSpace;subscriptions/&ZeroWidthSpace;{subscriptionId}/&ZeroWidthSpace;resourceGroups/&ZeroWidthSpace;{resourceGroupName}/&ZeroWidthSpace;providers/&ZeroWidthSpace;Microsoft.ManagedIdentity/userAssignedIdentities/{identityName}'. | [UserAssignedIdentityDetails[]](#UserAssignedIdentityDetails)<br/><small>Optional</small> |

<a id="ManagedClusterIdentity_STATUS"></a>ManagedClusterIdentity_STATUS
-----------------------------------------------------------------------

Identity for the managed cluster.

Used by: [ManagedCluster_STATUS](#ManagedCluster_STATUS).

| Property               | Description                                                                                                                                                                                                                                                                                    | Type                                                                                                                                                 |
|------------------------|------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|------------------------------------------------------------------------------------------------------------------------------------------------------|
| principalId            | The principal id of the system assigned identity which is used by master components.                                                                                                                                                                                                           | string<br/><small>Optional</small>                                                                                                                   |
| tenantId               | The tenant id of the system assigned identity which is used by master components.                                                                                                                                                                                                              | string<br/><small>Optional</small>                                                                                                                   |
| type                   | For more information see [use managed identities in AKS](https://docs.microsoft.com/azure/aks/use-managed-identity).                                                                                                                                                                           | [ManagedClusterIdentity_Type_STATUS](#ManagedClusterIdentity_Type_STATUS)<br/><small>Optional</small>                                                |
| userAssignedIdentities | The keys must be ARM resource IDs in the form: '/&ZeroWidthSpace;subscriptions/&ZeroWidthSpace;{subscriptionId}/&ZeroWidthSpace;resourceGroups/&ZeroWidthSpace;{resourceGroupName}/&ZeroWidthSpace;providers/&ZeroWidthSpace;Microsoft.ManagedIdentity/userAssignedIdentities/{identityName}'. | [map[string]ManagedClusterIdentity_UserAssignedIdentities_STATUS](#ManagedClusterIdentity_UserAssignedIdentities_STATUS)<br/><small>Optional</small> |

<a id="ManagedClusterOperatorSpec"></a>ManagedClusterOperatorSpec
-----------------------------------------------------------------

Details for configuring operator behavior. Fields in this struct are interpreted by the operator directly rather than being passed to Azure

Used by: [ManagedCluster_Spec](#ManagedCluster_Spec).

| Property | Description                                        | Type                                                                                        |
|----------|----------------------------------------------------|---------------------------------------------------------------------------------------------|
| secrets  | configures where to place Azure generated secrets. | [ManagedClusterOperatorSecrets](#ManagedClusterOperatorSecrets)<br/><small>Optional</small> |

<a id="ManagedClusterPodIdentityProfile"></a>ManagedClusterPodIdentityProfile
-----------------------------------------------------------------------------

See [use AAD pod identity](https://docs.microsoft.com/azure/aks/use-azure-ad-pod-identity) for more details on pod identity integration.

Used by: [ManagedCluster_Spec](#ManagedCluster_Spec).

| Property                       | Description                                                                                                                                                                                                                                                                                                                                                   | Type                                                                                                    |
|--------------------------------|---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|---------------------------------------------------------------------------------------------------------|
| allowNetworkPluginKubenet      | Running in Kubenet is disabled by default due to the security related nature of AAD Pod Identity and the risks of IP spoofing. See [using Kubenet network plugin with AAD Pod Identity](https://docs.microsoft.com/azure/aks/use-azure-ad-pod-identity#using-kubenet-network-plugin-with-azure-active-directory-pod-managed-identities) for more information. | bool<br/><small>Optional</small>                                                                        |
| enabled                        | Whether the pod identity addon is enabled.                                                                                                                                                                                                                                                                                                                    | bool<br/><small>Optional</small>                                                                        |
| userAssignedIdentities         | The pod identities to use in the cluster.                                                                                                                                                                                                                                                                                                                     | [ManagedClusterPodIdentity[]](#ManagedClusterPodIdentity)<br/><small>Optional</small>                   |
| userAssignedIdentityExceptions | The pod identity exceptions to allow.                                                                                                                                                                                                                                                                                                                         | [ManagedClusterPodIdentityException[]](#ManagedClusterPodIdentityException)<br/><small>Optional</small> |

<a id="ManagedClusterPodIdentityProfile_STATUS"></a>ManagedClusterPodIdentityProfile_STATUS
-------------------------------------------------------------------------------------------

See [use AAD pod identity](https://docs.microsoft.com/azure/aks/use-azure-ad-pod-identity) for more details on pod identity integration.

Used by: [ManagedCluster_STATUS](#ManagedCluster_STATUS).

| Property                       | Description                                                                                                                                                                                                                                                                                                                                                   | Type                                                                                                                  |
|--------------------------------|---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|-----------------------------------------------------------------------------------------------------------------------|
| allowNetworkPluginKubenet      | Running in Kubenet is disabled by default due to the security related nature of AAD Pod Identity and the risks of IP spoofing. See [using Kubenet network plugin with AAD Pod Identity](https://docs.microsoft.com/azure/aks/use-azure-ad-pod-identity#using-kubenet-network-plugin-with-azure-active-directory-pod-managed-identities) for more information. | bool<br/><small>Optional</small>                                                                                      |
| enabled                        | Whether the pod identity addon is enabled.                                                                                                                                                                                                                                                                                                                    | bool<br/><small>Optional</small>                                                                                      |
| userAssignedIdentities         | The pod identities to use in the cluster.                                                                                                                                                                                                                                                                                                                     | [ManagedClusterPodIdentity_STATUS[]](#ManagedClusterPodIdentity_STATUS)<br/><small>Optional</small>                   |
| userAssignedIdentityExceptions | The pod identity exceptions to allow.                                                                                                                                                                                                                                                                                                                         | [ManagedClusterPodIdentityException_STATUS[]](#ManagedClusterPodIdentityException_STATUS)<br/><small>Optional</small> |

<a id="ManagedClusterProperties_AutoScalerProfile"></a>ManagedClusterProperties_AutoScalerProfile
-------------------------------------------------------------------------------------------------

Used by: [ManagedCluster_Spec](#ManagedCluster_Spec).

| Property                         | Description                                                                                                                                                                                                                                                                                                                         | Type                                                                                                                                    |
|----------------------------------|-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|-----------------------------------------------------------------------------------------------------------------------------------------|
| balance-similar-node-groups      | Valid values are 'true' and 'false'                                                                                                                                                                                                                                                                                                 | string<br/><small>Optional</small>                                                                                                      |
| expander                         | If not specified, the default is 'random'. See [expanders](https://github.com/kubernetes/autoscaler/blob/master/cluster-autoscaler/FAQ.md#what-are-expanders) for more information.                                                                                                                                                 | [ManagedClusterProperties_AutoScalerProfile_Expander](#ManagedClusterProperties_AutoScalerProfile_Expander)<br/><small>Optional</small> |
| max-empty-bulk-delete            | The default is 10.                                                                                                                                                                                                                                                                                                                  | string<br/><small>Optional</small>                                                                                                      |
| max-graceful-termination-sec     | The default is 600.                                                                                                                                                                                                                                                                                                                 | string<br/><small>Optional</small>                                                                                                      |
| max-node-provision-time          | The default is '15m'. Values must be an integer followed by an 'm'. No unit of time other than minutes (m) is supported.                                                                                                                                                                                                            | string<br/><small>Optional</small>                                                                                                      |
| max-total-unready-percentage     | The default is 45. The maximum is 100 and the minimum is 0.                                                                                                                                                                                                                                                                         | string<br/><small>Optional</small>                                                                                                      |
| new-pod-scale-up-delay           | For scenarios like burst/batch scale where you don't want CA to act before the kubernetes scheduler could schedule all the pods, you can tell CA to ignore unscheduled pods before they're a certain age. The default is '0s'. Values must be an integer followed by a unit ('s' for seconds, 'm' for minutes, 'h' for hours, etc). | string<br/><small>Optional</small>                                                                                                      |
| ok-total-unready-count           | This must be an integer. The default is 3.                                                                                                                                                                                                                                                                                          | string<br/><small>Optional</small>                                                                                                      |
| scale-down-delay-after-add       | The default is '10m'. Values must be an integer followed by an 'm'. No unit of time other than minutes (m) is supported.                                                                                                                                                                                                            | string<br/><small>Optional</small>                                                                                                      |
| scale-down-delay-after-delete    | The default is the scan-interval. Values must be an integer followed by an 'm'. No unit of time other than minutes (m) is supported.                                                                                                                                                                                                | string<br/><small>Optional</small>                                                                                                      |
| scale-down-delay-after-failure   | The default is '3m'. Values must be an integer followed by an 'm'. No unit of time other than minutes (m) is supported.                                                                                                                                                                                                             | string<br/><small>Optional</small>                                                                                                      |
| scale-down-unneeded-time         | The default is '10m'. Values must be an integer followed by an 'm'. No unit of time other than minutes (m) is supported.                                                                                                                                                                                                            | string<br/><small>Optional</small>                                                                                                      |
| scale-down-unready-time          | The default is '20m'. Values must be an integer followed by an 'm'. No unit of time other than minutes (m) is supported.                                                                                                                                                                                                            | string<br/><small>Optional</small>                                                                                                      |
| scale-down-utilization-threshold | The default is '0.5'.                                                                                                                                                                                                                                                                                                               | string<br/><small>Optional</small>                                                                                                      |
| scan-interval                    | The default is '10'. Values must be an integer number of seconds.                                                                                                                                                                                                                                                                   | string<br/><small>Optional</small>                                                                                                      |
| skip-nodes-with-local-storage    | The default is true.                                                                                                                                                                                                                                                                                                                | string<br/><small>Optional</small>                                                                                                      |
| skip-nodes-with-system-pods      | The default is true.                                                                                                                                                                                                                                                                                                                | string<br/><small>Optional</small>                                                                                                      |

<a id="ManagedClusterProperties_AutoScalerProfile_STATUS"></a>ManagedClusterProperties_AutoScalerProfile_STATUS
---------------------------------------------------------------------------------------------------------------

Used by: [ManagedCluster_STATUS](#ManagedCluster_STATUS).

| Property                         | Description                                                                                                                                                                                                                                                                                                                         | Type                                                                                                                                                  |
|----------------------------------|-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|-------------------------------------------------------------------------------------------------------------------------------------------------------|
| balance-similar-node-groups      | Valid values are 'true' and 'false'                                                                                                                                                                                                                                                                                                 | string<br/><small>Optional</small>                                                                                                                    |
| expander                         | If not specified, the default is 'random'. See [expanders](https://github.com/kubernetes/autoscaler/blob/master/cluster-autoscaler/FAQ.md#what-are-expanders) for more information.                                                                                                                                                 | [ManagedClusterProperties_AutoScalerProfile_Expander_STATUS](#ManagedClusterProperties_AutoScalerProfile_Expander_STATUS)<br/><small>Optional</small> |
| max-empty-bulk-delete            | The default is 10.                                                                                                                                                                                                                                                                                                                  | string<br/><small>Optional</small>                                                                                                                    |
| max-graceful-termination-sec     | The default is 600.                                                                                                                                                                                                                                                                                                                 | string<br/><small>Optional</small>                                                                                                                    |
| max-node-provision-time          | The default is '15m'. Values must be an integer followed by an 'm'. No unit of time other than minutes (m) is supported.                                                                                                                                                                                                            | string<br/><small>Optional</small>                                                                                                                    |
| max-total-unready-percentage     | The default is 45. The maximum is 100 and the minimum is 0.                                                                                                                                                                                                                                                                         | string<br/><small>Optional</small>                                                                                                                    |
| new-pod-scale-up-delay           | For scenarios like burst/batch scale where you don't want CA to act before the kubernetes scheduler could schedule all the pods, you can tell CA to ignore unscheduled pods before they're a certain age. The default is '0s'. Values must be an integer followed by a unit ('s' for seconds, 'm' for minutes, 'h' for hours, etc). | string<br/><small>Optional</small>                                                                                                                    |
| ok-total-unready-count           | This must be an integer. The default is 3.                                                                                                                                                                                                                                                                                          | string<br/><small>Optional</small>                                                                                                                    |
| scale-down-delay-after-add       | The default is '10m'. Values must be an integer followed by an 'm'. No unit of time other than minutes (m) is supported.                                                                                                                                                                                                            | string<br/><small>Optional</small>                                                                                                                    |
| scale-down-delay-after-delete    | The default is the scan-interval. Values must be an integer followed by an 'm'. No unit of time other than minutes (m) is supported.                                                                                                                                                                                                | string<br/><small>Optional</small>                                                                                                                    |
| scale-down-delay-after-failure   | The default is '3m'. Values must be an integer followed by an 'm'. No unit of time other than minutes (m) is supported.                                                                                                                                                                                                             | string<br/><small>Optional</small>                                                                                                                    |
| scale-down-unneeded-time         | The default is '10m'. Values must be an integer followed by an 'm'. No unit of time other than minutes (m) is supported.                                                                                                                                                                                                            | string<br/><small>Optional</small>                                                                                                                    |
| scale-down-unready-time          | The default is '20m'. Values must be an integer followed by an 'm'. No unit of time other than minutes (m) is supported.                                                                                                                                                                                                            | string<br/><small>Optional</small>                                                                                                                    |
| scale-down-utilization-threshold | The default is '0.5'.                                                                                                                                                                                                                                                                                                               | string<br/><small>Optional</small>                                                                                                                    |
| scan-interval                    | The default is '10'. Values must be an integer number of seconds.                                                                                                                                                                                                                                                                   | string<br/><small>Optional</small>                                                                                                                    |
| skip-nodes-with-local-storage    | The default is true.                                                                                                                                                                                                                                                                                                                | string<br/><small>Optional</small>                                                                                                                    |
| skip-nodes-with-system-pods      | The default is true.                                                                                                                                                                                                                                                                                                                | string<br/><small>Optional</small>                                                                                                                    |

<a id="ManagedClusterServicePrincipalProfile"></a>ManagedClusterServicePrincipalProfile
---------------------------------------------------------------------------------------

Information about a service principal identity for the cluster to use for manipulating Azure APIs.

Used by: [ManagedCluster_Spec](#ManagedCluster_Spec).

| Property | Description                                                              | Type                                                   |
|----------|--------------------------------------------------------------------------|--------------------------------------------------------|
| clientId | The ID for the service principal.                                        | string<br/><small>Required</small>                     |
| secret   | The secret password associated with the service principal in plain text. | genruntime.SecretReference<br/><small>Optional</small> |

<a id="ManagedClusterServicePrincipalProfile_STATUS"></a>ManagedClusterServicePrincipalProfile_STATUS
-----------------------------------------------------------------------------------------------------

Information about a service principal identity for the cluster to use for manipulating Azure APIs.

Used by: [ManagedCluster_STATUS](#ManagedCluster_STATUS).

| Property | Description                       | Type                               |
|----------|-----------------------------------|------------------------------------|
| clientId | The ID for the service principal. | string<br/><small>Optional</small> |

<a id="ManagedClusterSKU"></a>ManagedClusterSKU
-----------------------------------------------

The SKU of a Managed Cluster.

Used by: [ManagedCluster_Spec](#ManagedCluster_Spec).

| Property | Description                                                                                                                  | Type                                                                          |
|----------|------------------------------------------------------------------------------------------------------------------------------|-------------------------------------------------------------------------------|
| name     | The name of a managed cluster SKU.                                                                                           | [ManagedClusterSKU_Name](#ManagedClusterSKU_Name)<br/><small>Optional</small> |
| tier     | If not specified, the default is 'Free'. See [uptime SLA](https://docs.microsoft.com/azure/aks/uptime-sla) for more details. | [ManagedClusterSKU_Tier](#ManagedClusterSKU_Tier)<br/><small>Optional</small> |

<a id="ManagedClusterSKU_STATUS"></a>ManagedClusterSKU_STATUS
-------------------------------------------------------------

The SKU of a Managed Cluster.

Used by: [ManagedCluster_STATUS](#ManagedCluster_STATUS).

| Property | Description                                                                                                                  | Type                                                                                        |
|----------|------------------------------------------------------------------------------------------------------------------------------|---------------------------------------------------------------------------------------------|
| name     | The name of a managed cluster SKU.                                                                                           | [ManagedClusterSKU_Name_STATUS](#ManagedClusterSKU_Name_STATUS)<br/><small>Optional</small> |
| tier     | If not specified, the default is 'Free'. See [uptime SLA](https://docs.microsoft.com/azure/aks/uptime-sla) for more details. | [ManagedClusterSKU_Tier_STATUS](#ManagedClusterSKU_Tier_STATUS)<br/><small>Optional</small> |

<a id="ManagedClusterWindowsProfile"></a>ManagedClusterWindowsProfile
---------------------------------------------------------------------

Profile for Windows VMs in the managed cluster.

Used by: [ManagedCluster_Spec](#ManagedCluster_Spec).

| Property       | Description                                                                                                                                                                                                                                                                                                                                                                                                                                                        | Type                                                                                                              |
|----------------|--------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|-------------------------------------------------------------------------------------------------------------------|
| adminPassword  | Specifies the password of the administrator account. Minimum-length: 8 characters Max-length: 123 characters Complexity requirements: 3 out of 4 conditions below need to be fulfilled Has lower characters Has upper characters Has a digit Has a special character (Regex match \[\W_]) Disallowed values: "abc@123", "P@$$w0rd", "P@ssw0rd", "P@ssword123", "Pa$$word", "pass@word1", "Password!", "Password1", "Password22", "iloveyou!"                       | genruntime.SecretReference<br/><small>Optional</small>                                                            |
| adminUsername  | Specifies the name of the administrator account. Restriction: Cannot end in "." Disallowed values: "administrator", "admin", "user", "user1", "test", "user2", "test1", "user3", "admin1", "1", "123", "a", "actuser", "adm", "admin2", "aspnet", "backup", "console", "david", "guest", "john", "owner", "root", "server", "sql", "support", "support_388945a0", "sys", "test2", "test3", "user4", "user5". Minimum-length: 1 character Max-length: 20 characters | string<br/><small>Required</small>                                                                                |
| enableCSIProxy | For more details on CSI proxy, see the [CSI proxy GitHub repo](https://github.com/kubernetes-csi/csi-proxy).                                                                                                                                                                                                                                                                                                                                                       | bool<br/><small>Optional</small>                                                                                  |
| licenseType    | The license type to use for Windows VMs. See [Azure Hybrid User Benefits](https://azure.microsoft.com/pricing/hybrid-benefit/faq/) for more details.                                                                                                                                                                                                                                                                                                               | [ManagedClusterWindowsProfile_LicenseType](#ManagedClusterWindowsProfile_LicenseType)<br/><small>Optional</small> |

<a id="ManagedClusterWindowsProfile_STATUS"></a>ManagedClusterWindowsProfile_STATUS
-----------------------------------------------------------------------------------

Profile for Windows VMs in the managed cluster.

Used by: [ManagedCluster_STATUS](#ManagedCluster_STATUS).

| Property       | Description                                                                                                                                                                                                                                                                                                                                                                                                                                                        | Type                                                                                                                            |
|----------------|--------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|---------------------------------------------------------------------------------------------------------------------------------|
| adminUsername  | Specifies the name of the administrator account. Restriction: Cannot end in "." Disallowed values: "administrator", "admin", "user", "user1", "test", "user2", "test1", "user3", "admin1", "1", "123", "a", "actuser", "adm", "admin2", "aspnet", "backup", "console", "david", "guest", "john", "owner", "root", "server", "sql", "support", "support_388945a0", "sys", "test2", "test3", "user4", "user5". Minimum-length: 1 character Max-length: 20 characters | string<br/><small>Optional</small>                                                                                              |
| enableCSIProxy | For more details on CSI proxy, see the [CSI proxy GitHub repo](https://github.com/kubernetes-csi/csi-proxy).                                                                                                                                                                                                                                                                                                                                                       | bool<br/><small>Optional</small>                                                                                                |
| licenseType    | The license type to use for Windows VMs. See [Azure Hybrid User Benefits](https://azure.microsoft.com/pricing/hybrid-benefit/faq/) for more details.                                                                                                                                                                                                                                                                                                               | [ManagedClusterWindowsProfile_LicenseType_STATUS](#ManagedClusterWindowsProfile_LicenseType_STATUS)<br/><small>Optional</small> |

<a id="OSDiskType"></a>OSDiskType
---------------------------------

The default is 'Ephemeral' if the VM supports it and has a cache disk larger than the requested OSDiskSizeGB. Otherwise, defaults to 'Managed'. May not be changed after creation. For more information see [Ephemeral OS](https://docs.microsoft.com/azure/aks/cluster-configuration#ephemeral-os).

Used by: [ManagedClusterAgentPoolProfile](#ManagedClusterAgentPoolProfile), and [ManagedClusters_AgentPool_Spec](#ManagedClusters_AgentPool_Spec).

| Value       | Description |
|-------------|-------------|
| "Ephemeral" |             |
| "Managed"   |             |

<a id="OSDiskType_STATUS"></a>OSDiskType_STATUS
-----------------------------------------------

The default is 'Ephemeral' if the VM supports it and has a cache disk larger than the requested OSDiskSizeGB. Otherwise, defaults to 'Managed'. May not be changed after creation. For more information see [Ephemeral OS](https://docs.microsoft.com/azure/aks/cluster-configuration#ephemeral-os).

Used by: [ManagedClusterAgentPoolProfile_STATUS](#ManagedClusterAgentPoolProfile_STATUS), and [ManagedClusters_AgentPool_STATUS](#ManagedClusters_AgentPool_STATUS).

| Value       | Description |
|-------------|-------------|
| "Ephemeral" |             |
| "Managed"   |             |

<a id="OSSKU"></a>OSSKU
-----------------------

Specifies an OS SKU. This value must not be specified if OSType is Windows.

Used by: [ManagedClusterAgentPoolProfile](#ManagedClusterAgentPoolProfile), and [ManagedClusters_AgentPool_Spec](#ManagedClusters_AgentPool_Spec).

| Value        | Description |
|--------------|-------------|
| "CBLMariner" |             |
| "Ubuntu"     |             |

<a id="OSSKU_STATUS"></a>OSSKU_STATUS
-------------------------------------

Specifies an OS SKU. This value must not be specified if OSType is Windows.

Used by: [ManagedClusterAgentPoolProfile_STATUS](#ManagedClusterAgentPoolProfile_STATUS), and [ManagedClusters_AgentPool_STATUS](#ManagedClusters_AgentPool_STATUS).

| Value        | Description |
|--------------|-------------|
| "CBLMariner" |             |
| "Ubuntu"     |             |

<a id="OSType"></a>OSType
-------------------------

The operating system type. The default is Linux.

Used by: [ManagedClusterAgentPoolProfile](#ManagedClusterAgentPoolProfile), and [ManagedClusters_AgentPool_Spec](#ManagedClusters_AgentPool_Spec).

| Value     | Description |
|-----------|-------------|
| "Linux"   |             |
| "Windows" |             |

<a id="OSType_STATUS"></a>OSType_STATUS
---------------------------------------

The operating system type. The default is Linux.

Used by: [ManagedClusterAgentPoolProfile_STATUS](#ManagedClusterAgentPoolProfile_STATUS), and [ManagedClusters_AgentPool_STATUS](#ManagedClusters_AgentPool_STATUS).

| Value     | Description |
|-----------|-------------|
| "Linux"   |             |
| "Windows" |             |

<a id="PowerState_STATUS"></a>PowerState_STATUS
-----------------------------------------------

Describes the Power State of the cluster

Used by: [ManagedCluster_STATUS](#ManagedCluster_STATUS), [ManagedClusterAgentPoolProfile_STATUS](#ManagedClusterAgentPoolProfile_STATUS), and [ManagedClusters_AgentPool_STATUS](#ManagedClusters_AgentPool_STATUS).

| Property | Description                                     | Type                                                                          |
|----------|-------------------------------------------------|-------------------------------------------------------------------------------|
| code     | Tells whether the cluster is Running or Stopped | [PowerState_Code_STATUS](#PowerState_Code_STATUS)<br/><small>Optional</small> |

<a id="PrivateLinkResource"></a>PrivateLinkResource
---------------------------------------------------

A private link resource

Used by: [ManagedCluster_Spec](#ManagedCluster_Spec).

| Property        | Description                            | Type                                                     |
|-----------------|----------------------------------------|----------------------------------------------------------|
| groupId         | The group ID of the resource.          | string<br/><small>Optional</small>                       |
| name            | The name of the private link resource. | string<br/><small>Optional</small>                       |
| reference       | The ID of the private link resource.   | genruntime.ResourceReference<br/><small>Optional</small> |
| requiredMembers | The RequiredMembers of the resource    | string[]<br/><small>Optional</small>                     |
| type            | The resource type.                     | string<br/><small>Optional</small>                       |

<a id="PrivateLinkResource_STATUS"></a>PrivateLinkResource_STATUS
-----------------------------------------------------------------

A private link resource

Used by: [ManagedCluster_STATUS](#ManagedCluster_STATUS).

| Property             | Description                                                                                | Type                                 |
|----------------------|--------------------------------------------------------------------------------------------|--------------------------------------|
| groupId              | The group ID of the resource.                                                              | string<br/><small>Optional</small>   |
| id                   | The ID of the private link resource.                                                       | string<br/><small>Optional</small>   |
| name                 | The name of the private link resource.                                                     | string<br/><small>Optional</small>   |
| privateLinkServiceID | The private link service ID of the resource, this field is exposed only to NRP internally. | string<br/><small>Optional</small>   |
| requiredMembers      | The RequiredMembers of the resource                                                        | string[]<br/><small>Optional</small> |
| type                 | The resource type.                                                                         | string<br/><small>Optional</small>   |

<a id="ScaleSetEvictionPolicy"></a>ScaleSetEvictionPolicy
---------------------------------------------------------

The eviction policy specifies what to do with the VM when it is evicted. The default is Delete. For more information about eviction see [spot VMs](https://docs.microsoft.com/azure/virtual-machines/spot-vms)

Used by: [ManagedClusterAgentPoolProfile](#ManagedClusterAgentPoolProfile), and [ManagedClusters_AgentPool_Spec](#ManagedClusters_AgentPool_Spec).

| Value        | Description |
|--------------|-------------|
| "Deallocate" |             |
| "Delete"     |             |

<a id="ScaleSetEvictionPolicy_STATUS"></a>ScaleSetEvictionPolicy_STATUS
-----------------------------------------------------------------------

The eviction policy specifies what to do with the VM when it is evicted. The default is Delete. For more information about eviction see [spot VMs](https://docs.microsoft.com/azure/virtual-machines/spot-vms)

Used by: [ManagedClusterAgentPoolProfile_STATUS](#ManagedClusterAgentPoolProfile_STATUS), and [ManagedClusters_AgentPool_STATUS](#ManagedClusters_AgentPool_STATUS).

| Value        | Description |
|--------------|-------------|
| "Deallocate" |             |
| "Delete"     |             |

<a id="ScaleSetPriority"></a>ScaleSetPriority
---------------------------------------------

The Virtual Machine Scale Set priority.

Used by: [ManagedClusterAgentPoolProfile](#ManagedClusterAgentPoolProfile), and [ManagedClusters_AgentPool_Spec](#ManagedClusters_AgentPool_Spec).

| Value     | Description |
|-----------|-------------|
| "Regular" |             |
| "Spot"    |             |

<a id="ScaleSetPriority_STATUS"></a>ScaleSetPriority_STATUS
-----------------------------------------------------------

The Virtual Machine Scale Set priority.

Used by: [ManagedClusterAgentPoolProfile_STATUS](#ManagedClusterAgentPoolProfile_STATUS), and [ManagedClusters_AgentPool_STATUS](#ManagedClusters_AgentPool_STATUS).

| Value     | Description |
|-----------|-------------|
| "Regular" |             |
| "Spot"    |             |

<a id="UserAssignedIdentity"></a>UserAssignedIdentity
-----------------------------------------------------

Details about a user assigned identity.

Used by: [ManagedCluster_Spec](#ManagedCluster_Spec), and [ManagedClusterPodIdentity](#ManagedClusterPodIdentity).

| Property          | Description                                    | Type                                                     |
|-------------------|------------------------------------------------|----------------------------------------------------------|
| clientId          | The client ID of the user assigned identity.   | string<br/><small>Optional</small>                       |
| objectId          | The object ID of the user assigned identity.   | string<br/><small>Optional</small>                       |
| resourceReference | The resource ID of the user assigned identity. | genruntime.ResourceReference<br/><small>Optional</small> |

<a id="UserAssignedIdentity_STATUS"></a>UserAssignedIdentity_STATUS
-------------------------------------------------------------------

Details about a user assigned identity.

Used by: [ManagedCluster_STATUS](#ManagedCluster_STATUS), [ManagedClusterAddonProfile_STATUS](#ManagedClusterAddonProfile_STATUS), and [ManagedClusterPodIdentity_STATUS](#ManagedClusterPodIdentity_STATUS).

| Property   | Description                                    | Type                               |
|------------|------------------------------------------------|------------------------------------|
| clientId   | The client ID of the user assigned identity.   | string<br/><small>Optional</small> |
| objectId   | The object ID of the user assigned identity.   | string<br/><small>Optional</small> |
| resourceId | The resource ID of the user assigned identity. | string<br/><small>Optional</small> |

<a id="ContainerServiceNetworkProfile_LoadBalancerSku"></a>ContainerServiceNetworkProfile_LoadBalancerSku
---------------------------------------------------------------------------------------------------------

Used by: [ContainerServiceNetworkProfile](#ContainerServiceNetworkProfile).

| Value      | Description |
|------------|-------------|
| "basic"    |             |
| "standard" |             |

<a id="ContainerServiceNetworkProfile_LoadBalancerSku_STATUS"></a>ContainerServiceNetworkProfile_LoadBalancerSku_STATUS
-----------------------------------------------------------------------------------------------------------------------

Used by: [ContainerServiceNetworkProfile_STATUS](#ContainerServiceNetworkProfile_STATUS).

| Value      | Description |
|------------|-------------|
| "basic"    |             |
| "standard" |             |

<a id="ContainerServiceNetworkProfile_NetworkMode"></a>ContainerServiceNetworkProfile_NetworkMode
-------------------------------------------------------------------------------------------------

Used by: [ContainerServiceNetworkProfile](#ContainerServiceNetworkProfile).

| Value         | Description |
|---------------|-------------|
| "bridge"      |             |
| "transparent" |             |

<a id="ContainerServiceNetworkProfile_NetworkMode_STATUS"></a>ContainerServiceNetworkProfile_NetworkMode_STATUS
---------------------------------------------------------------------------------------------------------------

Used by: [ContainerServiceNetworkProfile_STATUS](#ContainerServiceNetworkProfile_STATUS).

| Value         | Description |
|---------------|-------------|
| "bridge"      |             |
| "transparent" |             |

<a id="ContainerServiceNetworkProfile_NetworkPlugin"></a>ContainerServiceNetworkProfile_NetworkPlugin
-----------------------------------------------------------------------------------------------------

Used by: [ContainerServiceNetworkProfile](#ContainerServiceNetworkProfile).

| Value     | Description |
|-----------|-------------|
| "azure"   |             |
| "kubenet" |             |

<a id="ContainerServiceNetworkProfile_NetworkPlugin_STATUS"></a>ContainerServiceNetworkProfile_NetworkPlugin_STATUS
-------------------------------------------------------------------------------------------------------------------

Used by: [ContainerServiceNetworkProfile_STATUS](#ContainerServiceNetworkProfile_STATUS).

| Value     | Description |
|-----------|-------------|
| "azure"   |             |
| "kubenet" |             |

<a id="ContainerServiceNetworkProfile_NetworkPolicy"></a>ContainerServiceNetworkProfile_NetworkPolicy
-----------------------------------------------------------------------------------------------------

Used by: [ContainerServiceNetworkProfile](#ContainerServiceNetworkProfile).

| Value    | Description |
|----------|-------------|
| "azure"  |             |
| "calico" |             |

<a id="ContainerServiceNetworkProfile_NetworkPolicy_STATUS"></a>ContainerServiceNetworkProfile_NetworkPolicy_STATUS
-------------------------------------------------------------------------------------------------------------------

Used by: [ContainerServiceNetworkProfile_STATUS](#ContainerServiceNetworkProfile_STATUS).

| Value    | Description |
|----------|-------------|
| "azure"  |             |
| "calico" |             |

<a id="ContainerServiceNetworkProfile_OutboundType"></a>ContainerServiceNetworkProfile_OutboundType
---------------------------------------------------------------------------------------------------

Used by: [ContainerServiceNetworkProfile](#ContainerServiceNetworkProfile).

| Value                | Description |
|----------------------|-------------|
| "loadBalancer"       |             |
| "userDefinedRouting" |             |

<a id="ContainerServiceNetworkProfile_OutboundType_STATUS"></a>ContainerServiceNetworkProfile_OutboundType_STATUS
-----------------------------------------------------------------------------------------------------------------

Used by: [ContainerServiceNetworkProfile_STATUS](#ContainerServiceNetworkProfile_STATUS).

| Value                | Description |
|----------------------|-------------|
| "loadBalancer"       |             |
| "userDefinedRouting" |             |

<a id="ContainerServiceSshConfiguration"></a>ContainerServiceSshConfiguration
-----------------------------------------------------------------------------

SSH configuration for Linux-based VMs running on Azure.

Used by: [ContainerServiceLinuxProfile](#ContainerServiceLinuxProfile).

| Property   | Description                                                                                                 | Type                                                                                        |
|------------|-------------------------------------------------------------------------------------------------------------|---------------------------------------------------------------------------------------------|
| publicKeys | The list of SSH public keys used to authenticate with Linux-based VMs. A maximum of 1 key may be specified. | [ContainerServiceSshPublicKey[]](#ContainerServiceSshPublicKey)<br/><small>Required</small> |

<a id="ContainerServiceSshConfiguration_STATUS"></a>ContainerServiceSshConfiguration_STATUS
-------------------------------------------------------------------------------------------

SSH configuration for Linux-based VMs running on Azure.

Used by: [ContainerServiceLinuxProfile_STATUS](#ContainerServiceLinuxProfile_STATUS).

| Property   | Description                                                                                                 | Type                                                                                                      |
|------------|-------------------------------------------------------------------------------------------------------------|-----------------------------------------------------------------------------------------------------------|
| publicKeys | The list of SSH public keys used to authenticate with Linux-based VMs. A maximum of 1 key may be specified. | [ContainerServiceSshPublicKey_STATUS[]](#ContainerServiceSshPublicKey_STATUS)<br/><small>Optional</small> |

<a id="ExtendedLocationType"></a>ExtendedLocationType
-----------------------------------------------------

The type of extendedLocation.

Used by: [ExtendedLocation](#ExtendedLocation).

| Value      | Description |
|------------|-------------|
| "EdgeZone" |             |

<a id="ExtendedLocationType_STATUS"></a>ExtendedLocationType_STATUS
-------------------------------------------------------------------

The type of extendedLocation.

Used by: [ExtendedLocation_STATUS](#ExtendedLocation_STATUS).

| Value      | Description |
|------------|-------------|
| "EdgeZone" |             |

<a id="ManagedClusterAutoUpgradeProfile_UpgradeChannel"></a>ManagedClusterAutoUpgradeProfile_UpgradeChannel
-----------------------------------------------------------------------------------------------------------

Used by: [ManagedClusterAutoUpgradeProfile](#ManagedClusterAutoUpgradeProfile).

| Value        | Description |
|--------------|-------------|
| "node-image" |             |
| "none"       |             |
| "patch"      |             |
| "rapid"      |             |
| "stable"     |             |

<a id="ManagedClusterAutoUpgradeProfile_UpgradeChannel_STATUS"></a>ManagedClusterAutoUpgradeProfile_UpgradeChannel_STATUS
-------------------------------------------------------------------------------------------------------------------------

Used by: [ManagedClusterAutoUpgradeProfile_STATUS](#ManagedClusterAutoUpgradeProfile_STATUS).

| Value        | Description |
|--------------|-------------|
| "node-image" |             |
| "none"       |             |
| "patch"      |             |
| "rapid"      |             |
| "stable"     |             |

<a id="ManagedClusterIdentity_Type"></a>ManagedClusterIdentity_Type
-------------------------------------------------------------------

Used by: [ManagedClusterIdentity](#ManagedClusterIdentity).

| Value            | Description |
|------------------|-------------|
| "None"           |             |
| "SystemAssigned" |             |
| "UserAssigned"   |             |

<a id="ManagedClusterIdentity_Type_STATUS"></a>ManagedClusterIdentity_Type_STATUS
---------------------------------------------------------------------------------

Used by: [ManagedClusterIdentity_STATUS](#ManagedClusterIdentity_STATUS).

| Value            | Description |
|------------------|-------------|
| "None"           |             |
| "SystemAssigned" |             |
| "UserAssigned"   |             |

<a id="ManagedClusterIdentity_UserAssignedIdentities_STATUS"></a>ManagedClusterIdentity_UserAssignedIdentities_STATUS
---------------------------------------------------------------------------------------------------------------------

Used by: [ManagedClusterIdentity_STATUS](#ManagedClusterIdentity_STATUS).

| Property    | Description                                 | Type                               |
|-------------|---------------------------------------------|------------------------------------|
| clientId    | The client id of user assigned identity.    | string<br/><small>Optional</small> |
| principalId | The principal id of user assigned identity. | string<br/><small>Optional</small> |

<a id="ManagedClusterLoadBalancerProfile"></a>ManagedClusterLoadBalancerProfile
-------------------------------------------------------------------------------

Profile of the managed cluster load balancer.

Used by: [ContainerServiceNetworkProfile](#ContainerServiceNetworkProfile).

| Property               | Description                                                                                                                                                                               | Type                                                                                                                                      |
|------------------------|-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|-------------------------------------------------------------------------------------------------------------------------------------------|
| allocatedOutboundPorts | The desired number of allocated SNAT ports per VM. Allowed values are in the range of 0 to 64000 (inclusive). The default value is 0 which results in Azure dynamically allocating ports. | int<br/><small>Optional</small>                                                                                                           |
| effectiveOutboundIPs   | The effective outbound IP resources of the cluster load balancer.                                                                                                                         | [ResourceReference[]](#ResourceReference)<br/><small>Optional</small>                                                                     |
| idleTimeoutInMinutes   | Desired outbound flow idle timeout in minutes. Allowed values are in the range of 4 to 120 (inclusive). The default value is 30 minutes.                                                  | int<br/><small>Optional</small>                                                                                                           |
| managedOutboundIPs     | Desired managed outbound IPs for the cluster load balancer.                                                                                                                               | [ManagedClusterLoadBalancerProfile_ManagedOutboundIPs](#ManagedClusterLoadBalancerProfile_ManagedOutboundIPs)<br/><small>Optional</small> |
| outboundIPPrefixes     | Desired outbound IP Prefix resources for the cluster load balancer.                                                                                                                       | [ManagedClusterLoadBalancerProfile_OutboundIPPrefixes](#ManagedClusterLoadBalancerProfile_OutboundIPPrefixes)<br/><small>Optional</small> |
| outboundIPs            | Desired outbound IP resources for the cluster load balancer.                                                                                                                              | [ManagedClusterLoadBalancerProfile_OutboundIPs](#ManagedClusterLoadBalancerProfile_OutboundIPs)<br/><small>Optional</small>               |

<a id="ManagedClusterLoadBalancerProfile_STATUS"></a>ManagedClusterLoadBalancerProfile_STATUS
---------------------------------------------------------------------------------------------

Profile of the managed cluster load balancer.

Used by: [ContainerServiceNetworkProfile_STATUS](#ContainerServiceNetworkProfile_STATUS).

| Property               | Description                                                                                                                                                                               | Type                                                                                                                                                    |
|------------------------|-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|---------------------------------------------------------------------------------------------------------------------------------------------------------|
| allocatedOutboundPorts | The desired number of allocated SNAT ports per VM. Allowed values are in the range of 0 to 64000 (inclusive). The default value is 0 which results in Azure dynamically allocating ports. | int<br/><small>Optional</small>                                                                                                                         |
| effectiveOutboundIPs   | The effective outbound IP resources of the cluster load balancer.                                                                                                                         | [ResourceReference_STATUS[]](#ResourceReference_STATUS)<br/><small>Optional</small>                                                                     |
| idleTimeoutInMinutes   | Desired outbound flow idle timeout in minutes. Allowed values are in the range of 4 to 120 (inclusive). The default value is 30 minutes.                                                  | int<br/><small>Optional</small>                                                                                                                         |
| managedOutboundIPs     | Desired managed outbound IPs for the cluster load balancer.                                                                                                                               | [ManagedClusterLoadBalancerProfile_ManagedOutboundIPs_STATUS](#ManagedClusterLoadBalancerProfile_ManagedOutboundIPs_STATUS)<br/><small>Optional</small> |
| outboundIPPrefixes     | Desired outbound IP Prefix resources for the cluster load balancer.                                                                                                                       | [ManagedClusterLoadBalancerProfile_OutboundIPPrefixes_STATUS](#ManagedClusterLoadBalancerProfile_OutboundIPPrefixes_STATUS)<br/><small>Optional</small> |
| outboundIPs            | Desired outbound IP resources for the cluster load balancer.                                                                                                                              | [ManagedClusterLoadBalancerProfile_OutboundIPs_STATUS](#ManagedClusterLoadBalancerProfile_OutboundIPs_STATUS)<br/><small>Optional</small>               |

<a id="ManagedClusterOperatorSecrets"></a>ManagedClusterOperatorSecrets
-----------------------------------------------------------------------

Used by: [ManagedClusterOperatorSpec](#ManagedClusterOperatorSpec).

| Property         | Description                                                                                                            | Type                                                     |
|------------------|------------------------------------------------------------------------------------------------------------------------|----------------------------------------------------------|
| adminCredentials | indicates where the AdminCredentials secret should be placed. If omitted, the secret will not be retrieved from Azure. | genruntime.SecretDestination<br/><small>Optional</small> |
| userCredentials  | indicates where the UserCredentials secret should be placed. If omitted, the secret will not be retrieved from Azure.  | genruntime.SecretDestination<br/><small>Optional</small> |

<a id="ManagedClusterPodIdentity"></a>ManagedClusterPodIdentity
---------------------------------------------------------------

Details about the pod identity assigned to the Managed Cluster.

Used by: [ManagedClusterPodIdentityProfile](#ManagedClusterPodIdentityProfile).

| Property        | Description                                                        | Type                                                                      |
|-----------------|--------------------------------------------------------------------|---------------------------------------------------------------------------|
| bindingSelector | The binding selector to use for the AzureIdentityBinding resource. | string<br/><small>Optional</small>                                        |
| identity        | The user assigned identity details.                                | [UserAssignedIdentity](#UserAssignedIdentity)<br/><small>Required</small> |
| name            | The name of the pod identity.                                      | string<br/><small>Required</small>                                        |
| namespace       | The namespace of the pod identity.                                 | string<br/><small>Required</small>                                        |

<a id="ManagedClusterPodIdentity_STATUS"></a>ManagedClusterPodIdentity_STATUS
-----------------------------------------------------------------------------

Details about the pod identity assigned to the Managed Cluster.

Used by: [ManagedClusterPodIdentityProfile_STATUS](#ManagedClusterPodIdentityProfile_STATUS).

| Property          | Description                                                        | Type                                                                                                                                  |
|-------------------|--------------------------------------------------------------------|---------------------------------------------------------------------------------------------------------------------------------------|
| bindingSelector   | The binding selector to use for the AzureIdentityBinding resource. | string<br/><small>Optional</small>                                                                                                    |
| identity          | The user assigned identity details.                                | [UserAssignedIdentity_STATUS](#UserAssignedIdentity_STATUS)<br/><small>Optional</small>                                               |
| name              | The name of the pod identity.                                      | string<br/><small>Optional</small>                                                                                                    |
| namespace         | The namespace of the pod identity.                                 | string<br/><small>Optional</small>                                                                                                    |
| provisioningInfo  |                                                                    | [ManagedClusterPodIdentity_ProvisioningInfo_STATUS](#ManagedClusterPodIdentity_ProvisioningInfo_STATUS)<br/><small>Optional</small>   |
| provisioningState | The current provisioning state of the pod identity.                | [ManagedClusterPodIdentity_ProvisioningState_STATUS](#ManagedClusterPodIdentity_ProvisioningState_STATUS)<br/><small>Optional</small> |

<a id="ManagedClusterPodIdentityException"></a>ManagedClusterPodIdentityException
---------------------------------------------------------------------------------

See [disable AAD Pod Identity for a specific Pod/Application](https://azure.github.io/aad-pod-identity/docs/configure/application_exception/) for more details.

Used by: [ManagedClusterPodIdentityProfile](#ManagedClusterPodIdentityProfile).

| Property  | Description                                  | Type                                          |
|-----------|----------------------------------------------|-----------------------------------------------|
| name      | The name of the pod identity exception.      | string<br/><small>Required</small>            |
| namespace | The namespace of the pod identity exception. | string<br/><small>Required</small>            |
| podLabels | The pod labels to match.                     | map[string]string<br/><small>Required</small> |

<a id="ManagedClusterPodIdentityException_STATUS"></a>ManagedClusterPodIdentityException_STATUS
-----------------------------------------------------------------------------------------------

See [disable AAD Pod Identity for a specific Pod/Application](https://azure.github.io/aad-pod-identity/docs/configure/application_exception/) for more details.

Used by: [ManagedClusterPodIdentityProfile_STATUS](#ManagedClusterPodIdentityProfile_STATUS).

| Property  | Description                                  | Type                                          |
|-----------|----------------------------------------------|-----------------------------------------------|
| name      | The name of the pod identity exception.      | string<br/><small>Optional</small>            |
| namespace | The namespace of the pod identity exception. | string<br/><small>Optional</small>            |
| podLabels | The pod labels to match.                     | map[string]string<br/><small>Optional</small> |

<a id="ManagedClusterProperties_AutoScalerProfile_Expander"></a>ManagedClusterProperties_AutoScalerProfile_Expander
-------------------------------------------------------------------------------------------------------------------

Used by: [ManagedClusterProperties_AutoScalerProfile](#ManagedClusterProperties_AutoScalerProfile).

| Value         | Description |
|---------------|-------------|
| "least-waste" |             |
| "most-pods"   |             |
| "priority"    |             |
| "random"      |             |

<a id="ManagedClusterProperties_AutoScalerProfile_Expander_STATUS"></a>ManagedClusterProperties_AutoScalerProfile_Expander_STATUS
---------------------------------------------------------------------------------------------------------------------------------

Used by: [ManagedClusterProperties_AutoScalerProfile_STATUS](#ManagedClusterProperties_AutoScalerProfile_STATUS).

| Value         | Description |
|---------------|-------------|
| "least-waste" |             |
| "most-pods"   |             |
| "priority"    |             |
| "random"      |             |

<a id="ManagedClusterSKU_Name"></a>ManagedClusterSKU_Name
---------------------------------------------------------

Used by: [ManagedClusterSKU](#ManagedClusterSKU).

| Value   | Description |
|---------|-------------|
| "Basic" |             |

<a id="ManagedClusterSKU_Name_STATUS"></a>ManagedClusterSKU_Name_STATUS
-----------------------------------------------------------------------

Used by: [ManagedClusterSKU_STATUS](#ManagedClusterSKU_STATUS).

| Value   | Description |
|---------|-------------|
| "Basic" |             |

<a id="ManagedClusterSKU_Tier"></a>ManagedClusterSKU_Tier
---------------------------------------------------------

Used by: [ManagedClusterSKU](#ManagedClusterSKU).

| Value  | Description |
|--------|-------------|
| "Free" |             |
| "Paid" |             |

<a id="ManagedClusterSKU_Tier_STATUS"></a>ManagedClusterSKU_Tier_STATUS
-----------------------------------------------------------------------

Used by: [ManagedClusterSKU_STATUS](#ManagedClusterSKU_STATUS).

| Value  | Description |
|--------|-------------|
| "Free" |             |
| "Paid" |             |

<a id="ManagedClusterWindowsProfile_LicenseType"></a>ManagedClusterWindowsProfile_LicenseType
---------------------------------------------------------------------------------------------

Used by: [ManagedClusterWindowsProfile](#ManagedClusterWindowsProfile).

| Value            | Description |
|------------------|-------------|
| "None"           |             |
| "Windows_Server" |             |

<a id="ManagedClusterWindowsProfile_LicenseType_STATUS"></a>ManagedClusterWindowsProfile_LicenseType_STATUS
-----------------------------------------------------------------------------------------------------------

Used by: [ManagedClusterWindowsProfile_STATUS](#ManagedClusterWindowsProfile_STATUS).

| Value            | Description |
|------------------|-------------|
| "None"           |             |
| "Windows_Server" |             |

<a id="PowerState_Code_STATUS"></a>PowerState_Code_STATUS
---------------------------------------------------------

Used by: [PowerState_STATUS](#PowerState_STATUS).

| Value     | Description |
|-----------|-------------|
| "Running" |             |
| "Stopped" |             |

<a id="SysctlConfig"></a>SysctlConfig
-------------------------------------

Sysctl settings for Linux agent nodes.

Used by: [LinuxOSConfig](#LinuxOSConfig).

| Property                       | Description                                        | Type                               |
|--------------------------------|----------------------------------------------------|------------------------------------|
| fsAioMaxNr                     | Sysctl setting fs.aio-max-nr.                      | int<br/><small>Optional</small>    |
| fsFileMax                      | Sysctl setting fs.file-max.                        | int<br/><small>Optional</small>    |
| fsInotifyMaxUserWatches        | Sysctl setting fs.inotify.max_user_watches.        | int<br/><small>Optional</small>    |
| fsNrOpen                       | Sysctl setting fs.nr_open.                         | int<br/><small>Optional</small>    |
| kernelThreadsMax               | Sysctl setting kernel.threads-max.                 | int<br/><small>Optional</small>    |
| netCoreNetdevMaxBacklog        | Sysctl setting net.core.netdev_max_backlog.        | int<br/><small>Optional</small>    |
| netCoreOptmemMax               | Sysctl setting net.core.optmem_max.                | int<br/><small>Optional</small>    |
| netCoreRmemDefault             | Sysctl setting net.core.rmem_default.              | int<br/><small>Optional</small>    |
| netCoreRmemMax                 | Sysctl setting net.core.rmem_max.                  | int<br/><small>Optional</small>    |
| netCoreSomaxconn               | Sysctl setting net.core.somaxconn.                 | int<br/><small>Optional</small>    |
| netCoreWmemDefault             | Sysctl setting net.core.wmem_default.              | int<br/><small>Optional</small>    |
| netCoreWmemMax                 | Sysctl setting net.core.wmem_max.                  | int<br/><small>Optional</small>    |
| netIpv4IpLocalPortRange        | Sysctl setting net.ipv4.ip_local_port_range.       | string<br/><small>Optional</small> |
| netIpv4NeighDefaultGcThresh1   | Sysctl setting net.ipv4.neigh.default.gc_thresh1.  | int<br/><small>Optional</small>    |
| netIpv4NeighDefaultGcThresh2   | Sysctl setting net.ipv4.neigh.default.gc_thresh2.  | int<br/><small>Optional</small>    |
| netIpv4NeighDefaultGcThresh3   | Sysctl setting net.ipv4.neigh.default.gc_thresh3.  | int<br/><small>Optional</small>    |
| netIpv4TcpFinTimeout           | Sysctl setting net.ipv4.tcp_fin_timeout.           | int<br/><small>Optional</small>    |
| netIpv4TcpkeepaliveIntvl       | Sysctl setting net.ipv4.tcp_keepalive_intvl.       | int<br/><small>Optional</small>    |
| netIpv4TcpKeepaliveProbes      | Sysctl setting net.ipv4.tcp_keepalive_probes.      | int<br/><small>Optional</small>    |
| netIpv4TcpKeepaliveTime        | Sysctl setting net.ipv4.tcp_keepalive_time.        | int<br/><small>Optional</small>    |
| netIpv4TcpMaxSynBacklog        | Sysctl setting net.ipv4.tcp_max_syn_backlog.       | int<br/><small>Optional</small>    |
| netIpv4TcpMaxTwBuckets         | Sysctl setting net.ipv4.tcp_max_tw_buckets.        | int<br/><small>Optional</small>    |
| netIpv4TcpTwReuse              | Sysctl setting net.ipv4.tcp_tw_reuse.              | bool<br/><small>Optional</small>   |
| netNetfilterNfConntrackBuckets | Sysctl setting net.netfilter.nf_conntrack_buckets. | int<br/><small>Optional</small>    |
| netNetfilterNfConntrackMax     | Sysctl setting net.netfilter.nf_conntrack_max.     | int<br/><small>Optional</small>    |
| vmMaxMapCount                  | Sysctl setting vm.max_map_count.                   | int<br/><small>Optional</small>    |
| vmSwappiness                   | Sysctl setting vm.swappiness.                      | int<br/><small>Optional</small>    |
| vmVfsCachePressure             | Sysctl setting vm.vfs_cache_pressure.              | int<br/><small>Optional</small>    |

<a id="SysctlConfig_STATUS"></a>SysctlConfig_STATUS
---------------------------------------------------

Sysctl settings for Linux agent nodes.

Used by: [LinuxOSConfig_STATUS](#LinuxOSConfig_STATUS).

| Property                       | Description                                        | Type                               |
|--------------------------------|----------------------------------------------------|------------------------------------|
| fsAioMaxNr                     | Sysctl setting fs.aio-max-nr.                      | int<br/><small>Optional</small>    |
| fsFileMax                      | Sysctl setting fs.file-max.                        | int<br/><small>Optional</small>    |
| fsInotifyMaxUserWatches        | Sysctl setting fs.inotify.max_user_watches.        | int<br/><small>Optional</small>    |
| fsNrOpen                       | Sysctl setting fs.nr_open.                         | int<br/><small>Optional</small>    |
| kernelThreadsMax               | Sysctl setting kernel.threads-max.                 | int<br/><small>Optional</small>    |
| netCoreNetdevMaxBacklog        | Sysctl setting net.core.netdev_max_backlog.        | int<br/><small>Optional</small>    |
| netCoreOptmemMax               | Sysctl setting net.core.optmem_max.                | int<br/><small>Optional</small>    |
| netCoreRmemDefault             | Sysctl setting net.core.rmem_default.              | int<br/><small>Optional</small>    |
| netCoreRmemMax                 | Sysctl setting net.core.rmem_max.                  | int<br/><small>Optional</small>    |
| netCoreSomaxconn               | Sysctl setting net.core.somaxconn.                 | int<br/><small>Optional</small>    |
| netCoreWmemDefault             | Sysctl setting net.core.wmem_default.              | int<br/><small>Optional</small>    |
| netCoreWmemMax                 | Sysctl setting net.core.wmem_max.                  | int<br/><small>Optional</small>    |
| netIpv4IpLocalPortRange        | Sysctl setting net.ipv4.ip_local_port_range.       | string<br/><small>Optional</small> |
| netIpv4NeighDefaultGcThresh1   | Sysctl setting net.ipv4.neigh.default.gc_thresh1.  | int<br/><small>Optional</small>    |
| netIpv4NeighDefaultGcThresh2   | Sysctl setting net.ipv4.neigh.default.gc_thresh2.  | int<br/><small>Optional</small>    |
| netIpv4NeighDefaultGcThresh3   | Sysctl setting net.ipv4.neigh.default.gc_thresh3.  | int<br/><small>Optional</small>    |
| netIpv4TcpFinTimeout           | Sysctl setting net.ipv4.tcp_fin_timeout.           | int<br/><small>Optional</small>    |
| netIpv4TcpkeepaliveIntvl       | Sysctl setting net.ipv4.tcp_keepalive_intvl.       | int<br/><small>Optional</small>    |
| netIpv4TcpKeepaliveProbes      | Sysctl setting net.ipv4.tcp_keepalive_probes.      | int<br/><small>Optional</small>    |
| netIpv4TcpKeepaliveTime        | Sysctl setting net.ipv4.tcp_keepalive_time.        | int<br/><small>Optional</small>    |
| netIpv4TcpMaxSynBacklog        | Sysctl setting net.ipv4.tcp_max_syn_backlog.       | int<br/><small>Optional</small>    |
| netIpv4TcpMaxTwBuckets         | Sysctl setting net.ipv4.tcp_max_tw_buckets.        | int<br/><small>Optional</small>    |
| netIpv4TcpTwReuse              | Sysctl setting net.ipv4.tcp_tw_reuse.              | bool<br/><small>Optional</small>   |
| netNetfilterNfConntrackBuckets | Sysctl setting net.netfilter.nf_conntrack_buckets. | int<br/><small>Optional</small>    |
| netNetfilterNfConntrackMax     | Sysctl setting net.netfilter.nf_conntrack_max.     | int<br/><small>Optional</small>    |
| vmMaxMapCount                  | Sysctl setting vm.max_map_count.                   | int<br/><small>Optional</small>    |
| vmSwappiness                   | Sysctl setting vm.swappiness.                      | int<br/><small>Optional</small>    |
| vmVfsCachePressure             | Sysctl setting vm.vfs_cache_pressure.              | int<br/><small>Optional</small>    |

<a id="UserAssignedIdentityDetails"></a>UserAssignedIdentityDetails
-------------------------------------------------------------------

Information about the user assigned identity for the resource

Used by: [ManagedClusterIdentity](#ManagedClusterIdentity).

| Property  | Description | Type                                                     |
|-----------|-------------|----------------------------------------------------------|
| reference |             | genruntime.ResourceReference<br/><small>Optional</small> |

<a id="ContainerServiceSshPublicKey"></a>ContainerServiceSshPublicKey
---------------------------------------------------------------------

Contains information about SSH certificate public key data.

Used by: [ContainerServiceSshConfiguration](#ContainerServiceSshConfiguration).

| Property | Description                                                                                                                      | Type                               |
|----------|----------------------------------------------------------------------------------------------------------------------------------|------------------------------------|
| keyData  | Certificate public key used to authenticate with VMs through SSH. The certificate must be in PEM format with or without headers. | string<br/><small>Required</small> |

<a id="ContainerServiceSshPublicKey_STATUS"></a>ContainerServiceSshPublicKey_STATUS
-----------------------------------------------------------------------------------

Contains information about SSH certificate public key data.

Used by: [ContainerServiceSshConfiguration_STATUS](#ContainerServiceSshConfiguration_STATUS).

| Property | Description                                                                                                                      | Type                               |
|----------|----------------------------------------------------------------------------------------------------------------------------------|------------------------------------|
| keyData  | Certificate public key used to authenticate with VMs through SSH. The certificate must be in PEM format with or without headers. | string<br/><small>Optional</small> |

<a id="ManagedClusterLoadBalancerProfile_ManagedOutboundIPs"></a>ManagedClusterLoadBalancerProfile_ManagedOutboundIPs
---------------------------------------------------------------------------------------------------------------------

Used by: [ManagedClusterLoadBalancerProfile](#ManagedClusterLoadBalancerProfile).

| Property | Description                                                                                                                                                                     | Type                            |
|----------|---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|---------------------------------|
| count    | The desired number of outbound IPs created/managed by Azure for the cluster load balancer. Allowed values must be in the range of 1 to 100 (inclusive). The default value is 1. | int<br/><small>Optional</small> |

<a id="ManagedClusterLoadBalancerProfile_ManagedOutboundIPs_STATUS"></a>ManagedClusterLoadBalancerProfile_ManagedOutboundIPs_STATUS
-----------------------------------------------------------------------------------------------------------------------------------

Used by: [ManagedClusterLoadBalancerProfile_STATUS](#ManagedClusterLoadBalancerProfile_STATUS).

| Property | Description                                                                                                                                                                     | Type                            |
|----------|---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|---------------------------------|
| count    | The desired number of outbound IPs created/managed by Azure for the cluster load balancer. Allowed values must be in the range of 1 to 100 (inclusive). The default value is 1. | int<br/><small>Optional</small> |

<a id="ManagedClusterLoadBalancerProfile_OutboundIPPrefixes"></a>ManagedClusterLoadBalancerProfile_OutboundIPPrefixes
---------------------------------------------------------------------------------------------------------------------

Used by: [ManagedClusterLoadBalancerProfile](#ManagedClusterLoadBalancerProfile).

| Property         | Description                           | Type                                                                  |
|------------------|---------------------------------------|-----------------------------------------------------------------------|
| publicIPPrefixes | A list of public IP prefix resources. | [ResourceReference[]](#ResourceReference)<br/><small>Optional</small> |

<a id="ManagedClusterLoadBalancerProfile_OutboundIPPrefixes_STATUS"></a>ManagedClusterLoadBalancerProfile_OutboundIPPrefixes_STATUS
-----------------------------------------------------------------------------------------------------------------------------------

Used by: [ManagedClusterLoadBalancerProfile_STATUS](#ManagedClusterLoadBalancerProfile_STATUS).

| Property         | Description                           | Type                                                                                |
|------------------|---------------------------------------|-------------------------------------------------------------------------------------|
| publicIPPrefixes | A list of public IP prefix resources. | [ResourceReference_STATUS[]](#ResourceReference_STATUS)<br/><small>Optional</small> |

<a id="ManagedClusterLoadBalancerProfile_OutboundIPs"></a>ManagedClusterLoadBalancerProfile_OutboundIPs
-------------------------------------------------------------------------------------------------------

Used by: [ManagedClusterLoadBalancerProfile](#ManagedClusterLoadBalancerProfile).

| Property  | Description                    | Type                                                                  |
|-----------|--------------------------------|-----------------------------------------------------------------------|
| publicIPs | A list of public IP resources. | [ResourceReference[]](#ResourceReference)<br/><small>Optional</small> |

<a id="ManagedClusterLoadBalancerProfile_OutboundIPs_STATUS"></a>ManagedClusterLoadBalancerProfile_OutboundIPs_STATUS
---------------------------------------------------------------------------------------------------------------------

Used by: [ManagedClusterLoadBalancerProfile_STATUS](#ManagedClusterLoadBalancerProfile_STATUS).

| Property  | Description                    | Type                                                                                |
|-----------|--------------------------------|-------------------------------------------------------------------------------------|
| publicIPs | A list of public IP resources. | [ResourceReference_STATUS[]](#ResourceReference_STATUS)<br/><small>Optional</small> |

<a id="ManagedClusterPodIdentity_ProvisioningInfo_STATUS"></a>ManagedClusterPodIdentity_ProvisioningInfo_STATUS
---------------------------------------------------------------------------------------------------------------

Used by: [ManagedClusterPodIdentity_STATUS](#ManagedClusterPodIdentity_STATUS).

| Property | Description                             | Type                                                                                                                                |
|----------|-----------------------------------------|-------------------------------------------------------------------------------------------------------------------------------------|
| error    | Pod identity assignment error (if any). | [ManagedClusterPodIdentityProvisioningError_STATUS](#ManagedClusterPodIdentityProvisioningError_STATUS)<br/><small>Optional</small> |

<a id="ManagedClusterPodIdentity_ProvisioningState_STATUS"></a>ManagedClusterPodIdentity_ProvisioningState_STATUS
-----------------------------------------------------------------------------------------------------------------

Used by: [ManagedClusterPodIdentity_STATUS](#ManagedClusterPodIdentity_STATUS).

| Value      | Description |
|------------|-------------|
| "Assigned" |             |
| "Deleting" |             |
| "Failed"   |             |
| "Updating" |             |

<a id="ResourceReference"></a>ResourceReference
-----------------------------------------------

A reference to an Azure resource.

Used by: [ManagedClusterLoadBalancerProfile](#ManagedClusterLoadBalancerProfile), [ManagedClusterLoadBalancerProfile_OutboundIPPrefixes](#ManagedClusterLoadBalancerProfile_OutboundIPPrefixes), and [ManagedClusterLoadBalancerProfile_OutboundIPs](#ManagedClusterLoadBalancerProfile_OutboundIPs).

| Property  | Description                            | Type                                                     |
|-----------|----------------------------------------|----------------------------------------------------------|
| reference | The fully qualified Azure resource id. | genruntime.ResourceReference<br/><small>Optional</small> |

<a id="ResourceReference_STATUS"></a>ResourceReference_STATUS
-------------------------------------------------------------

A reference to an Azure resource.

Used by: [ManagedClusterLoadBalancerProfile_OutboundIPPrefixes_STATUS](#ManagedClusterLoadBalancerProfile_OutboundIPPrefixes_STATUS), [ManagedClusterLoadBalancerProfile_OutboundIPs_STATUS](#ManagedClusterLoadBalancerProfile_OutboundIPs_STATUS), and [ManagedClusterLoadBalancerProfile_STATUS](#ManagedClusterLoadBalancerProfile_STATUS).

| Property | Description                            | Type                               |
|----------|----------------------------------------|------------------------------------|
| id       | The fully qualified Azure resource id. | string<br/><small>Optional</small> |

<a id="ManagedClusterPodIdentityProvisioningError_STATUS"></a>ManagedClusterPodIdentityProvisioningError_STATUS
---------------------------------------------------------------------------------------------------------------

An error response from the pod identity provisioning.

Used by: [ManagedClusterPodIdentity_ProvisioningInfo_STATUS](#ManagedClusterPodIdentity_ProvisioningInfo_STATUS).

| Property | Description              | Type                                                                                                                                        |
|----------|--------------------------|---------------------------------------------------------------------------------------------------------------------------------------------|
| error    | Details about the error. | [ManagedClusterPodIdentityProvisioningErrorBody_STATUS](#ManagedClusterPodIdentityProvisioningErrorBody_STATUS)<br/><small>Optional</small> |

<a id="ManagedClusterPodIdentityProvisioningErrorBody_STATUS"></a>ManagedClusterPodIdentityProvisioningErrorBody_STATUS
-----------------------------------------------------------------------------------------------------------------------

An error response from the pod identity provisioning.

Used by: [ManagedClusterPodIdentityProvisioningError_STATUS](#ManagedClusterPodIdentityProvisioningError_STATUS).

| Property | Description                                                                                        | Type                                                                                                                                                            |
|----------|----------------------------------------------------------------------------------------------------|-----------------------------------------------------------------------------------------------------------------------------------------------------------------|
| code     | An identifier for the error. Codes are invariant and are intended to be consumed programmatically. | string<br/><small>Optional</small>                                                                                                                              |
| details  | A list of additional details about the error.                                                      | [ManagedClusterPodIdentityProvisioningErrorBody_STATUS_Unrolled[]](#ManagedClusterPodIdentityProvisioningErrorBody_STATUS_Unrolled)<br/><small>Optional</small> |
| message  | A message describing the error, intended to be suitable for display in a user interface.           | string<br/><small>Optional</small>                                                                                                                              |
| target   | The target of the particular error. For example, the name of the property in error.                | string<br/><small>Optional</small>                                                                                                                              |

<a id="ManagedClusterPodIdentityProvisioningErrorBody_STATUS_Unrolled"></a>ManagedClusterPodIdentityProvisioningErrorBody_STATUS_Unrolled
-----------------------------------------------------------------------------------------------------------------------------------------

Used by: [ManagedClusterPodIdentityProvisioningErrorBody_STATUS](#ManagedClusterPodIdentityProvisioningErrorBody_STATUS).

| Property | Description                                                                                        | Type                               |
|----------|----------------------------------------------------------------------------------------------------|------------------------------------|
| code     | An identifier for the error. Codes are invariant and are intended to be consumed programmatically. | string<br/><small>Optional</small> |
| message  | A message describing the error, intended to be suitable for display in a user interface.           | string<br/><small>Optional</small> |
| target   | The target of the particular error. For example, the name of the property in error.                | string<br/><small>Optional</small> |
