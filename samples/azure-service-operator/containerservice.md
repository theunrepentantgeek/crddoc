v1api20210501
=============

<a id="apiversion"></a>APIVersion
---------------------------------

.

<a id="agentpoolmode"></a>AgentPoolMode
---------------------------------------

Used by: [ManagedClusterAgentPoolProfile](#managedclusteragentpoolprofile), and [ManagedClusters_AgentPool_Spec](#managedclusters_agentpool_spec).

<a id="agentpoolmode_status"></a>AgentPoolMode_STATUS
-----------------------------------------------------

Used by: [ManagedClusterAgentPoolProfile_STATUS](#managedclusteragentpoolprofile_status), and [ManagedClusters_AgentPool_STATUS](#managedclusters_agentpool_status).

<a id="agentpooltype"></a>AgentPoolType
---------------------------------------

Used by: [ManagedClusterAgentPoolProfile](#managedclusteragentpoolprofile), and [ManagedClusters_AgentPool_Spec](#managedclusters_agentpool_spec).

<a id="agentpooltype_status"></a>AgentPoolType_STATUS
-----------------------------------------------------

Used by: [ManagedClusterAgentPoolProfile_STATUS](#managedclusteragentpoolprofile_status), and [ManagedClusters_AgentPool_STATUS](#managedclusters_agentpool_status).

<a id="agentpoolupgradesettings"></a>AgentPoolUpgradeSettings
-------------------------------------------------------------

Settings for upgrading an agentpool

Used by: [ManagedClusterAgentPoolProfile](#managedclusteragentpoolprofile), and [ManagedClusters_AgentPool_Spec](#managedclusters_agentpool_spec).

| Property | Description                                                                                                                                                                                                                                                                                                                                                                                                            | Type   |
|----------|------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|--------|
| maxSurge | This can either be set to an integer (e.g. '5') or a percentage (e.g. '50%'). If a percentage is specified, it is the percentage of the total agent pool size at the time of the upgrade. For percentages, fractional nodes are rounded up. If not specified, the default is 1. For more information, including best practices, see: https://docs.microsoft.com/azure/aks/upgrade-cluster#customize-node-surge-upgrade | string |

<a id="agentpoolupgradesettings_status"></a>AgentPoolUpgradeSettings_STATUS
---------------------------------------------------------------------------

Settings for upgrading an agentpool

Used by: [ManagedClusterAgentPoolProfile_STATUS](#managedclusteragentpoolprofile_status), and [ManagedClusters_AgentPool_STATUS](#managedclusters_agentpool_status).

| Property | Description                                                                                                                                                                                                                                                                                                                                                                                                            | Type   |
|----------|------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|--------|
| maxSurge | This can either be set to an integer (e.g. '5') or a percentage (e.g. '50%'). If a percentage is specified, it is the percentage of the total agent pool size at the time of the upgrade. For percentages, fractional nodes are rounded up. If not specified, the default is 1. For more information, including best practices, see: https://docs.microsoft.com/azure/aks/upgrade-cluster#customize-node-surge-upgrade | string |

<a id="containerservicelinuxprofile"></a>ContainerServiceLinuxProfile
---------------------------------------------------------------------

Profile for Linux VMs in the container service cluster.

Used by: [ManagedCluster_Spec](#managedcluster_spec).

| Property      | Description                                                 | Type                                                                               |
|---------------|-------------------------------------------------------------|------------------------------------------------------------------------------------|
| adminUsername | The administrator username to use for Linux VMs.            | string<br/>Required                                                                |
| ssh           | The SSH configuration for Linux-based VMs running on Azure. | [ContainerServiceSshConfiguration](#containerservicesshconfiguration)<br/>Required |

<a id="containerservicelinuxprofile_status"></a>ContainerServiceLinuxProfile_STATUS
-----------------------------------------------------------------------------------

Profile for Linux VMs in the container service cluster.

Used by: [ManagedCluster_STATUS](#managedcluster_status).

| Property      | Description                                                 | Type                                                                                |
|---------------|-------------------------------------------------------------|-------------------------------------------------------------------------------------|
| adminUsername | The administrator username to use for Linux VMs.            | string                                                                              |
| ssh           | The SSH configuration for Linux-based VMs running on Azure. | [ContainerServiceSshConfiguration_STATUS](#containerservicesshconfiguration_status) |

<a id="containerservicenetworkprofile"></a>ContainerServiceNetworkProfile
-------------------------------------------------------------------------

Profile of network configuration.

Used by: [ManagedCluster_Spec](#managedcluster_spec).

| Property            | Description                                                                                                                                                                           | Type                                                                                              |
|---------------------|---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|---------------------------------------------------------------------------------------------------|
| dnsServiceIP        | An IP address assigned to the Kubernetes DNS service. It must be within the Kubernetes service address range specified in serviceCidr.                                                | string                                                                                            |
| dockerBridgeCidr    | A CIDR notation IP range assigned to the Docker bridge network. It must not overlap with any Subnet IP ranges or the Kubernetes service address range.                                | string                                                                                            |
| loadBalancerProfile | Profile of the cluster load balancer.                                                                                                                                                 | [ManagedClusterLoadBalancerProfile](#managedclusterloadbalancerprofile)                           |
| loadBalancerSku     | The default is 'standard'. See [Azure Load Balancer SKUs](https://docs.microsoft.com/azure/load-balancer/skus) for more information about the differences between load balancer SKUs. | [ContainerServiceNetworkProfile_LoadBalancerSku](#containerservicenetworkprofile_loadbalancersku) |
| networkMode         | This cannot be specified if networkPlugin is anything other than 'azure'.                                                                                                             | [ContainerServiceNetworkProfile_NetworkMode](#containerservicenetworkprofile_networkmode)         |
| networkPlugin       | Network plugin used for building the Kubernetes network.                                                                                                                              | [ContainerServiceNetworkProfile_NetworkPlugin](#containerservicenetworkprofile_networkplugin)     |
| networkPolicy       | Network policy used for building the Kubernetes network.                                                                                                                              | [ContainerServiceNetworkProfile_NetworkPolicy](#containerservicenetworkprofile_networkpolicy)     |
| outboundType        | This can only be set at cluster creation time and cannot be changed later. For more information see [egress outbound type](https://docs.microsoft.com/azure/aks/egress-outboundtype). | [ContainerServiceNetworkProfile_OutboundType](#containerservicenetworkprofile_outboundtype)       |
| podCidr             | A CIDR notation IP range from which to assign pod IPs when kubenet is used.                                                                                                           | string                                                                                            |
| serviceCidr         | A CIDR notation IP range from which to assign service cluster IPs. It must not overlap with any Subnet IP ranges.                                                                     | string                                                                                            |

<a id="containerservicenetworkprofile_loadbalancersku"></a>ContainerServiceNetworkProfile_LoadBalancerSku
---------------------------------------------------------------------------------------------------------

Used by: [ContainerServiceNetworkProfile](#containerservicenetworkprofile).

<a id="containerservicenetworkprofile_loadbalancersku_status"></a>ContainerServiceNetworkProfile_LoadBalancerSku_STATUS
-----------------------------------------------------------------------------------------------------------------------

Used by: [ContainerServiceNetworkProfile_STATUS](#containerservicenetworkprofile_status).

<a id="containerservicenetworkprofile_networkmode"></a>ContainerServiceNetworkProfile_NetworkMode
-------------------------------------------------------------------------------------------------

Used by: [ContainerServiceNetworkProfile](#containerservicenetworkprofile).

<a id="containerservicenetworkprofile_networkmode_status"></a>ContainerServiceNetworkProfile_NetworkMode_STATUS
---------------------------------------------------------------------------------------------------------------

Used by: [ContainerServiceNetworkProfile_STATUS](#containerservicenetworkprofile_status).

<a id="containerservicenetworkprofile_networkplugin"></a>ContainerServiceNetworkProfile_NetworkPlugin
-----------------------------------------------------------------------------------------------------

Used by: [ContainerServiceNetworkProfile](#containerservicenetworkprofile).

<a id="containerservicenetworkprofile_networkplugin_status"></a>ContainerServiceNetworkProfile_NetworkPlugin_STATUS
-------------------------------------------------------------------------------------------------------------------

Used by: [ContainerServiceNetworkProfile_STATUS](#containerservicenetworkprofile_status).

<a id="containerservicenetworkprofile_networkpolicy"></a>ContainerServiceNetworkProfile_NetworkPolicy
-----------------------------------------------------------------------------------------------------

Used by: [ContainerServiceNetworkProfile](#containerservicenetworkprofile).

<a id="containerservicenetworkprofile_networkpolicy_status"></a>ContainerServiceNetworkProfile_NetworkPolicy_STATUS
-------------------------------------------------------------------------------------------------------------------

Used by: [ContainerServiceNetworkProfile_STATUS](#containerservicenetworkprofile_status).

<a id="containerservicenetworkprofile_outboundtype"></a>ContainerServiceNetworkProfile_OutboundType
---------------------------------------------------------------------------------------------------

Used by: [ContainerServiceNetworkProfile](#containerservicenetworkprofile).

<a id="containerservicenetworkprofile_outboundtype_status"></a>ContainerServiceNetworkProfile_OutboundType_STATUS
-----------------------------------------------------------------------------------------------------------------

Used by: [ContainerServiceNetworkProfile_STATUS](#containerservicenetworkprofile_status).

<a id="containerservicenetworkprofile_status"></a>ContainerServiceNetworkProfile_STATUS
---------------------------------------------------------------------------------------

Profile of network configuration.

Used by: [ManagedCluster_STATUS](#managedcluster_status).

| Property            | Description                                                                                                                                                                           | Type                                                                                                            |
|---------------------|---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|-----------------------------------------------------------------------------------------------------------------|
| dnsServiceIP        | An IP address assigned to the Kubernetes DNS service. It must be within the Kubernetes service address range specified in serviceCidr.                                                | string                                                                                                          |
| dockerBridgeCidr    | A CIDR notation IP range assigned to the Docker bridge network. It must not overlap with any Subnet IP ranges or the Kubernetes service address range.                                | string                                                                                                          |
| loadBalancerProfile | Profile of the cluster load balancer.                                                                                                                                                 | [ManagedClusterLoadBalancerProfile_STATUS](#managedclusterloadbalancerprofile_status)                           |
| loadBalancerSku     | The default is 'standard'. See [Azure Load Balancer SKUs](https://docs.microsoft.com/azure/load-balancer/skus) for more information about the differences between load balancer SKUs. | [ContainerServiceNetworkProfile_LoadBalancerSku_STATUS](#containerservicenetworkprofile_loadbalancersku_status) |
| networkMode         | This cannot be specified if networkPlugin is anything other than 'azure'.                                                                                                             | [ContainerServiceNetworkProfile_NetworkMode_STATUS](#containerservicenetworkprofile_networkmode_status)         |
| networkPlugin       | Network plugin used for building the Kubernetes network.                                                                                                                              | [ContainerServiceNetworkProfile_NetworkPlugin_STATUS](#containerservicenetworkprofile_networkplugin_status)     |
| networkPolicy       | Network policy used for building the Kubernetes network.                                                                                                                              | [ContainerServiceNetworkProfile_NetworkPolicy_STATUS](#containerservicenetworkprofile_networkpolicy_status)     |
| outboundType        | This can only be set at cluster creation time and cannot be changed later. For more information see [egress outbound type](https://docs.microsoft.com/azure/aks/egress-outboundtype). | [ContainerServiceNetworkProfile_OutboundType_STATUS](#containerservicenetworkprofile_outboundtype_status)       |
| podCidr             | A CIDR notation IP range from which to assign pod IPs when kubenet is used.                                                                                                           | string                                                                                                          |
| serviceCidr         | A CIDR notation IP range from which to assign service cluster IPs. It must not overlap with any Subnet IP ranges.                                                                     | string                                                                                                          |

<a id="containerserviceosdisk"></a>ContainerServiceOSDisk
---------------------------------------------------------

Used by: [ManagedClusterAgentPoolProfile](#managedclusteragentpoolprofile), and [ManagedClusters_AgentPool_Spec](#managedclusters_agentpool_spec).

<a id="containerservicesshconfiguration"></a>ContainerServiceSshConfiguration
-----------------------------------------------------------------------------

SSH configuration for Linux-based VMs running on Azure.

Used by: [ContainerServiceLinuxProfile](#containerservicelinuxprofile).

| Property   | Description                                                                                                 | Type                                                                         |
|------------|-------------------------------------------------------------------------------------------------------------|------------------------------------------------------------------------------|
| publicKeys | The list of SSH public keys used to authenticate with Linux-based VMs. A maximum of 1 key may be specified. | [ContainerServiceSshPublicKey[]](#containerservicesshpublickey)<br/>Required |

<a id="containerservicesshconfiguration_status"></a>ContainerServiceSshConfiguration_STATUS
-------------------------------------------------------------------------------------------

SSH configuration for Linux-based VMs running on Azure.

Used by: [ContainerServiceLinuxProfile_STATUS](#containerservicelinuxprofile_status).

| Property   | Description                                                                                                 | Type                                                                          |
|------------|-------------------------------------------------------------------------------------------------------------|-------------------------------------------------------------------------------|
| publicKeys | The list of SSH public keys used to authenticate with Linux-based VMs. A maximum of 1 key may be specified. | [ContainerServiceSshPublicKey_STATUS[]](#containerservicesshpublickey_status) |

<a id="containerservicesshpublickey"></a>ContainerServiceSshPublicKey
---------------------------------------------------------------------

Contains information about SSH certificate public key data.

Used by: [ContainerServiceSshConfiguration](#containerservicesshconfiguration).

| Property | Description                                                                                                                      | Type                |
|----------|----------------------------------------------------------------------------------------------------------------------------------|---------------------|
| keyData  | Certificate public key used to authenticate with VMs through SSH. The certificate must be in PEM format with or without headers. | string<br/>Required |

<a id="containerservicesshpublickey_status"></a>ContainerServiceSshPublicKey_STATUS
-----------------------------------------------------------------------------------

Contains information about SSH certificate public key data.

Used by: [ContainerServiceSshConfiguration_STATUS](#containerservicesshconfiguration_status).

| Property | Description                                                                                                                      | Type   |
|----------|----------------------------------------------------------------------------------------------------------------------------------|--------|
| keyData  | Certificate public key used to authenticate with VMs through SSH. The certificate must be in PEM format with or without headers. | string |

<a id="extendedlocation"></a>ExtendedLocation
---------------------------------------------

The complex type of the extended location.

Used by: [ManagedCluster_Spec](#managedcluster_spec).

| Property | Description                        | Type                                          |
|----------|------------------------------------|-----------------------------------------------|
| name     | The name of the extended location. | string                                        |
| type     | The type of the extended location. | [ExtendedLocationType](#extendedlocationtype) |

<a id="extendedlocationtype"></a>ExtendedLocationType
-----------------------------------------------------

Used by: [ExtendedLocation](#extendedlocation).

<a id="extendedlocationtype_status"></a>ExtendedLocationType_STATUS
-------------------------------------------------------------------

Used by: [ExtendedLocation_STATUS](#extendedlocation_status).

<a id="extendedlocation_status"></a>ExtendedLocation_STATUS
-----------------------------------------------------------

The complex type of the extended location.

Used by: [ManagedCluster_STATUS](#managedcluster_status).

| Property | Description                        | Type                                                        |
|----------|------------------------------------|-------------------------------------------------------------|
| name     | The name of the extended location. | string                                                      |
| type     | The type of the extended location. | [ExtendedLocationType_STATUS](#extendedlocationtype_status) |

<a id="gpuinstanceprofile"></a>GPUInstanceProfile
-------------------------------------------------

Used by: [ManagedClusterAgentPoolProfile](#managedclusteragentpoolprofile), and [ManagedClusters_AgentPool_Spec](#managedclusters_agentpool_spec).

<a id="gpuinstanceprofile_status"></a>GPUInstanceProfile_STATUS
---------------------------------------------------------------

Used by: [ManagedClusterAgentPoolProfile_STATUS](#managedclusteragentpoolprofile_status), and [ManagedClusters_AgentPool_STATUS](#managedclusters_agentpool_status).

<a id="kubeletconfig"></a>KubeletConfig
---------------------------------------

See [AKS custom node configuration](https://docs.microsoft.com/azure/aks/custom-node-configuration) for more details.

Used by: [ManagedClusterAgentPoolProfile](#managedclusteragentpoolprofile), and [ManagedClusters_AgentPool_Spec](#managedclusters_agentpool_spec).

| Property              | Description                                                                                                                                                                                                                          | Type     |
|-----------------------|--------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|----------|
| allowedUnsafeSysctls  | Allowed list of unsafe sysctls or unsafe sysctl patterns (ending in `*`).                                                                                                                                                            | string[] |
| containerLogMaxFiles  | The maximum number of container log files that can be present for a container. The number must be ≥ 2.                                                                                                                               | int      |
| containerLogMaxSizeMB | The maximum size (e.g. 10Mi) of container log file before it is rotated.                                                                                                                                                             | int      |
| cpuCfsQuota           | The default is true.                                                                                                                                                                                                                 | bool     |
| cpuCfsQuotaPeriod     | The default is '100ms.' Valid values are a sequence of decimal numbers with an optional fraction and a unit suffix. For example: '300ms', '2h45m'. Supported units are 'ns', 'us', 'ms', 's', 'm', and 'h'.                          | string   |
| cpuManagerPolicy      | The default is 'none'. See [Kubernetes CPU management policies](https://kubernetes.io/docs/tasks/administer-cluster/cpu-management-policies/#cpu-management-policies) for more information. Allowed values are 'none' and 'static'.  | string   |
| failSwapOn            | If set to true it will make the Kubelet fail to start if swap is enabled on the node.                                                                                                                                                | bool     |
| imageGcHighThreshold  | To disable image garbage collection, set to 100. The default is 85%                                                                                                                                                                  | int      |
| imageGcLowThreshold   | This cannot be set higher than imageGcHighThreshold. The default is 80%                                                                                                                                                              | int      |
| podMaxPids            | The maximum number of processes per pod.                                                                                                                                                                                             | int      |
| topologyManagerPolicy | For more information see [Kubernetes Topology Manager](https://kubernetes.io/docs/tasks/administer-cluster/topology-manager). The default is 'none'. Allowed values are 'none', 'best-effort', 'restricted', and 'single-numa-node'. | string   |

<a id="kubeletconfig_status"></a>KubeletConfig_STATUS
-----------------------------------------------------

See [AKS custom node configuration](https://docs.microsoft.com/azure/aks/custom-node-configuration) for more details.

Used by: [ManagedClusterAgentPoolProfile_STATUS](#managedclusteragentpoolprofile_status), and [ManagedClusters_AgentPool_STATUS](#managedclusters_agentpool_status).

| Property              | Description                                                                                                                                                                                                                          | Type     |
|-----------------------|--------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|----------|
| allowedUnsafeSysctls  | Allowed list of unsafe sysctls or unsafe sysctl patterns (ending in `*`).                                                                                                                                                            | string[] |
| containerLogMaxFiles  | The maximum number of container log files that can be present for a container. The number must be ≥ 2.                                                                                                                               | int      |
| containerLogMaxSizeMB | The maximum size (e.g. 10Mi) of container log file before it is rotated.                                                                                                                                                             | int      |
| cpuCfsQuota           | The default is true.                                                                                                                                                                                                                 | bool     |
| cpuCfsQuotaPeriod     | The default is '100ms.' Valid values are a sequence of decimal numbers with an optional fraction and a unit suffix. For example: '300ms', '2h45m'. Supported units are 'ns', 'us', 'ms', 's', 'm', and 'h'.                          | string   |
| cpuManagerPolicy      | The default is 'none'. See [Kubernetes CPU management policies](https://kubernetes.io/docs/tasks/administer-cluster/cpu-management-policies/#cpu-management-policies) for more information. Allowed values are 'none' and 'static'.  | string   |
| failSwapOn            | If set to true it will make the Kubelet fail to start if swap is enabled on the node.                                                                                                                                                | bool     |
| imageGcHighThreshold  | To disable image garbage collection, set to 100. The default is 85%                                                                                                                                                                  | int      |
| imageGcLowThreshold   | This cannot be set higher than imageGcHighThreshold. The default is 80%                                                                                                                                                              | int      |
| podMaxPids            | The maximum number of processes per pod.                                                                                                                                                                                             | int      |
| topologyManagerPolicy | For more information see [Kubernetes Topology Manager](https://kubernetes.io/docs/tasks/administer-cluster/topology-manager). The default is 'none'. Allowed values are 'none', 'best-effort', 'restricted', and 'single-numa-node'. | string   |

<a id="kubeletdisktype"></a>KubeletDiskType
-------------------------------------------

Used by: [ManagedClusterAgentPoolProfile](#managedclusteragentpoolprofile), and [ManagedClusters_AgentPool_Spec](#managedclusters_agentpool_spec).

<a id="kubeletdisktype_status"></a>KubeletDiskType_STATUS
---------------------------------------------------------

Used by: [ManagedClusterAgentPoolProfile_STATUS](#managedclusteragentpoolprofile_status), and [ManagedClusters_AgentPool_STATUS](#managedclusters_agentpool_status).

<a id="linuxosconfig"></a>LinuxOSConfig
---------------------------------------

See [AKS custom node configuration](https://docs.microsoft.com/azure/aks/custom-node-configuration) for more details.

Used by: [ManagedClusterAgentPoolProfile](#managedclusteragentpoolprofile), and [ManagedClusters_AgentPool_Spec](#managedclusters_agentpool_spec).

| Property                   | Description                                                                                                                                                                                                                                         | Type                          |
|----------------------------|-----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|-------------------------------|
| swapFileSizeMB             | The size in MB of a swap file that will be created on each node.                                                                                                                                                                                    | int                           |
| sysctls                    | Sysctl settings for Linux agent nodes.                                                                                                                                                                                                              | [SysctlConfig](#sysctlconfig) |
| transparentHugePageDefrag  | Valid values are 'always', 'defer', 'defer+madvise', 'madvise' and 'never'. The default is 'madvise'. For more information see [Transparent Hugepages](https://www.kernel.org/doc/html/latest/admin-guide/mm/transhuge.html#admin-guide-transhuge). | string                        |
| transparentHugePageEnabled | Valid values are 'always', 'madvise', and 'never'. The default is 'always'. For more information see [Transparent Hugepages](https://www.kernel.org/doc/html/latest/admin-guide/mm/transhuge.html#admin-guide-transhuge).                           | string                        |

<a id="linuxosconfig_status"></a>LinuxOSConfig_STATUS
-----------------------------------------------------

See [AKS custom node configuration](https://docs.microsoft.com/azure/aks/custom-node-configuration) for more details.

Used by: [ManagedClusterAgentPoolProfile_STATUS](#managedclusteragentpoolprofile_status), and [ManagedClusters_AgentPool_STATUS](#managedclusters_agentpool_status).

| Property                   | Description                                                                                                                                                                                                                                         | Type                                        |
|----------------------------|-----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|---------------------------------------------|
| swapFileSizeMB             | The size in MB of a swap file that will be created on each node.                                                                                                                                                                                    | int                                         |
| sysctls                    | Sysctl settings for Linux agent nodes.                                                                                                                                                                                                              | [SysctlConfig_STATUS](#sysctlconfig_status) |
| transparentHugePageDefrag  | Valid values are 'always', 'defer', 'defer+madvise', 'madvise' and 'never'. The default is 'madvise'. For more information see [Transparent Hugepages](https://www.kernel.org/doc/html/latest/admin-guide/mm/transhuge.html#admin-guide-transhuge). | string                                      |
| transparentHugePageEnabled | Valid values are 'always', 'madvise', and 'never'. The default is 'always'. For more information see [Transparent Hugepages](https://www.kernel.org/doc/html/latest/admin-guide/mm/transhuge.html#admin-guide-transhuge).                           | string                                      |

<a id="managedcluster"></a>ManagedCluster
-----------------------------------------

Generator information: - Generated from: /containerservice/resource-manager/Microsoft.ContainerService/aks/stable/2021-05-01/managedClusters.json - ARM URI: /&ZeroWidthSpace;subscriptions/&ZeroWidthSpace;{subscriptionId}/&ZeroWidthSpace;resourceGroups/&ZeroWidthSpace;{resourceGroupName}/&ZeroWidthSpace;providers/&ZeroWidthSpace;Microsoft.ContainerService/managedClusters/{resourceName}

Used by: [ManagedClusterList](#managedclusterlist).

| Property          | Description | Type                                            |
|-------------------|-------------|-------------------------------------------------|
| metav1.TypeMeta   |             |                                                 |
| metav1.ObjectMeta |             |                                                 |
| spec              |             | [ManagedCluster_Spec](#managedcluster_spec)     |
| status            |             | [ManagedCluster_STATUS](#managedcluster_status) |

<a id="managedclusteraadprofile"></a>ManagedClusterAADProfile
-------------------------------------------------------------

For more details see [managed AAD on AKS](https://docs.microsoft.com/azure/aks/managed-aad).

Used by: [ManagedCluster_Spec](#managedcluster_spec).

| Property            | Description                                                                                                        | Type     |
|---------------------|--------------------------------------------------------------------------------------------------------------------|----------|
| adminGroupObjectIDs | The list of AAD group object IDs that will have admin role of the cluster.                                         | string[] |
| clientAppID         | The client AAD application ID.                                                                                     | string   |
| enableAzureRBAC     | Whether to enable Azure RBAC for Kubernetes authorization.                                                         | bool     |
| managed             | Whether to enable managed AAD.                                                                                     | bool     |
| serverAppID         | The server AAD application ID.                                                                                     | string   |
| serverAppSecret     | The server AAD application secret.                                                                                 | string   |
| tenantID            | The AAD tenant ID to use for authentication. If not specified, will use the tenant of the deployment subscription. | string   |

<a id="managedclusteraadprofile_status"></a>ManagedClusterAADProfile_STATUS
---------------------------------------------------------------------------

For more details see [managed AAD on AKS](https://docs.microsoft.com/azure/aks/managed-aad).

Used by: [ManagedCluster_STATUS](#managedcluster_status).

| Property            | Description                                                                                                        | Type     |
|---------------------|--------------------------------------------------------------------------------------------------------------------|----------|
| adminGroupObjectIDs | The list of AAD group object IDs that will have admin role of the cluster.                                         | string[] |
| clientAppID         | The client AAD application ID.                                                                                     | string   |
| enableAzureRBAC     | Whether to enable Azure RBAC for Kubernetes authorization.                                                         | bool     |
| managed             | Whether to enable managed AAD.                                                                                     | bool     |
| serverAppID         | The server AAD application ID.                                                                                     | string   |
| serverAppSecret     | The server AAD application secret.                                                                                 | string   |
| tenantID            | The AAD tenant ID to use for authentication. If not specified, will use the tenant of the deployment subscription. | string   |

<a id="managedclusterapiserveraccessprofile"></a>ManagedClusterAPIServerAccessProfile
-------------------------------------------------------------------------------------

Access profile for managed cluster API server.

Used by: [ManagedCluster_Spec](#managedcluster_spec).

| Property                       | Description                                                                                                                                                                                                                                                                                                                   | Type     |
|--------------------------------|-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|----------|
| authorizedIPRanges             | IP ranges are specified in CIDR format, e.g. 137.117.106.88/29. This feature is not compatible with clusters that use Public IP Per Node, or clusters that are using a Basic Load Balancer. For more information see [API server authorized IP ranges](https://docs.microsoft.com/azure/aks/api-server-authorized-ip-ranges). | string[] |
| enablePrivateCluster           | For more details, see [Creating a private AKS cluster](https://docs.microsoft.com/azure/aks/private-clusters).                                                                                                                                                                                                                | bool     |
| enablePrivateClusterPublicFQDN | Whether to create additional public FQDN for private cluster or not.                                                                                                                                                                                                                                                          | bool     |
| privateDNSZone                 | The default is System. For more details see [configure private DNS zone](https://docs.microsoft.com/azure/aks/private-clusters#configure-private-dns-zone). Allowed values are 'system' and 'none'.                                                                                                                           | string   |

<a id="managedclusterapiserveraccessprofile_status"></a>ManagedClusterAPIServerAccessProfile_STATUS
---------------------------------------------------------------------------------------------------

Access profile for managed cluster API server.

Used by: [ManagedCluster_STATUS](#managedcluster_status).

| Property                       | Description                                                                                                                                                                                                                                                                                                                   | Type     |
|--------------------------------|-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|----------|
| authorizedIPRanges             | IP ranges are specified in CIDR format, e.g. 137.117.106.88/29. This feature is not compatible with clusters that use Public IP Per Node, or clusters that are using a Basic Load Balancer. For more information see [API server authorized IP ranges](https://docs.microsoft.com/azure/aks/api-server-authorized-ip-ranges). | string[] |
| enablePrivateCluster           | For more details, see [Creating a private AKS cluster](https://docs.microsoft.com/azure/aks/private-clusters).                                                                                                                                                                                                                | bool     |
| enablePrivateClusterPublicFQDN | Whether to create additional public FQDN for private cluster or not.                                                                                                                                                                                                                                                          | bool     |
| privateDNSZone                 | The default is System. For more details see [configure private DNS zone](https://docs.microsoft.com/azure/aks/private-clusters#configure-private-dns-zone). Allowed values are 'system' and 'none'.                                                                                                                           | string   |

<a id="managedclusteraddonprofile"></a>ManagedClusterAddonProfile
-----------------------------------------------------------------

A Kubernetes add-on profile for a managed cluster.

Used by: [ManagedCluster_Spec](#managedcluster_spec).

| Property | Description                                | Type              |
|----------|--------------------------------------------|-------------------|
| config   | Key-value pairs for configuring an add-on. | map[string]string |
| enabled  | Whether the add-on is enabled or not.      | bool<br/>Required |

<a id="managedclusteraddonprofile_status"></a>ManagedClusterAddonProfile_STATUS
-------------------------------------------------------------------------------

A Kubernetes add-on profile for a managed cluster.

Used by: [ManagedCluster_STATUS](#managedcluster_status).

| Property | Description                                                | Type                                                        |
|----------|------------------------------------------------------------|-------------------------------------------------------------|
| config   | Key-value pairs for configuring an add-on.                 | map[string]string                                           |
| enabled  | Whether the add-on is enabled or not.                      | bool                                                        |
| identity | Information of user assigned identity used by this add-on. | [UserAssignedIdentity_STATUS](#userassignedidentity_status) |

<a id="managedclusteragentpoolprofile"></a>ManagedClusterAgentPoolProfile
-------------------------------------------------------------------------

Profile for the container service agent pool.

Used by: [ManagedCluster_Spec](#managedcluster_spec).

| Property                      | Description                                                                                                                                                                                                                                                                                                                                                                                                                                                                                  | Type                                                  |
|-------------------------------|----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|-------------------------------------------------------|
| availabilityZones             | The list of Availability zones to use for nodes. This can only be specified if the AgentPoolType property is 'VirtualMachineScaleSets'.                                                                                                                                                                                                                                                                                                                                                      | string[]                                              |
| count                         | Number of agents (VMs) to host docker containers. Allowed values must be in the range of 0 to 1000 (inclusive) for user pools and in the range of 1 to 1000 (inclusive) for system pools. The default value is 1.                                                                                                                                                                                                                                                                            | int                                                   |
| enableAutoScaling             | Whether to enable auto-scaler                                                                                                                                                                                                                                                                                                                                                                                                                                                                | bool                                                  |
| enableEncryptionAtHost        | This is only supported on certain VM sizes and in certain Azure regions. For more information, see: https://docs.microsoft.com/azure/aks/enable-host-encryption                                                                                                                                                                                                                                                                                                                              | bool                                                  |
| enableFIPS                    | See [Add a FIPS-enabled node pool](https://docs.microsoft.com/azure/aks/use-multiple-node-pools#add-a-fips-enabled-node-pool-preview) for more details.                                                                                                                                                                                                                                                                                                                                      | bool                                                  |
| enableNodePublicIP            | Some scenarios may require nodes in a node pool to receive their own dedicated public IP addresses. A common scenario is for gaming workloads, where a console needs to make a direct connection to a cloud virtual machine to minimize hops. For more information see [assigning a public IP per node](https://docs.microsoft.com/azure/aks/use-multiple-node-pools#assign-a-public-ip-per-node-for-your-node-pools). The default is false.                                                 | bool                                                  |
| enableUltraSSD                | Whether to enable UltraSSD                                                                                                                                                                                                                                                                                                                                                                                                                                                                   | bool                                                  |
| gpuInstanceProfile            | GPUInstanceProfile to be used to specify GPU MIG instance profile for supported GPU VM SKU.                                                                                                                                                                                                                                                                                                                                                                                                  | [GPUInstanceProfile](#gpuinstanceprofile)             |
| kubeletConfig                 | The Kubelet configuration on the agent pool nodes.                                                                                                                                                                                                                                                                                                                                                                                                                                           | [KubeletConfig](#kubeletconfig)                       |
| kubeletDiskType               | Determines the placement of emptyDir volumes, container runtime data root, and Kubelet ephemeral storage.                                                                                                                                                                                                                                                                                                                                                                                    | [KubeletDiskType](#kubeletdisktype)                   |
| linuxOSConfig                 | The OS configuration of Linux agent nodes.                                                                                                                                                                                                                                                                                                                                                                                                                                                   | [LinuxOSConfig](#linuxosconfig)                       |
| maxCount                      | The maximum number of nodes for auto-scaling                                                                                                                                                                                                                                                                                                                                                                                                                                                 | int                                                   |
| maxPods                       | The maximum number of pods that can run on a node.                                                                                                                                                                                                                                                                                                                                                                                                                                           | int                                                   |
| minCount                      | The minimum number of nodes for auto-scaling                                                                                                                                                                                                                                                                                                                                                                                                                                                 | int                                                   |
| mode                          | A cluster must have at least one 'System' Agent Pool at all times. For additional information on agent pool restrictions and best practices, see: https://docs.microsoft.com/azure/aks/use-system-pools                                                                                                                                                                                                                                                                                      | [AgentPoolMode](#agentpoolmode)                       |
| name                          | Windows agent pool names must be 6 characters or less.                                                                                                                                                                                                                                                                                                                                                                                                                                       | string                                                |
| nodeLabels                    | The node labels to be persisted across all nodes in agent pool.                                                                                                                                                                                                                                                                                                                                                                                                                              | map[string]string                                     |
| nodePublicIPPrefixIDReference | This is of the form: /&ZeroWidthSpace;subscriptions/&ZeroWidthSpace;{subscriptionId}/&ZeroWidthSpace;resourceGroups/&ZeroWidthSpace;{resourceGroupName}/&ZeroWidthSpace;providers/&ZeroWidthSpace;Microsoft.Network/publicIPPrefixes/{publicIPPrefixName}                                                                                                                                                                                                                                    | genruntime.ResourceReference                          |
| nodeTaints                    | The taints added to new nodes during node pool create and scale. For example, key=value:NoSchedule.                                                                                                                                                                                                                                                                                                                                                                                          | string[]                                              |
| orchestratorVersion           | As a best practice, you should upgrade all node pools in an AKS cluster to the same Kubernetes version. The node pool version must have the same major version as the control plane. The node pool minor version must be within two minor versions of the control plane version. The node pool version cannot be greater than the control plane version. For more information see [upgrading a node pool](https://docs.microsoft.com/azure/aks/use-multiple-node-pools#upgrade-a-node-pool). | string                                                |
| osDiskSizeGB                  |                                                                                                                                                                                                                                                                                                                                                                                                                                                                                              | [ContainerServiceOSDisk](#containerserviceosdisk)     |
| osDiskType                    | The default is 'Ephemeral' if the VM supports it and has a cache disk larger than the requested OSDiskSizeGB. Otherwise, defaults to 'Managed'. May not be changed after creation. For more information see [Ephemeral OS](https://docs.microsoft.com/azure/aks/cluster-configuration#ephemeral-os).                                                                                                                                                                                         | [OSDiskType](#osdisktype)                             |
| osSKU                         | Specifies an OS SKU. This value must not be specified if OSType is Windows.                                                                                                                                                                                                                                                                                                                                                                                                                  | [OSSKU](#ossku)                                       |
| osType                        | The operating system type. The default is Linux.                                                                                                                                                                                                                                                                                                                                                                                                                                             | [OSType](#ostype)                                     |
| podSubnetIDReference          | If omitted, pod IPs are statically assigned on the node subnet (see vnetSubnetID for more details). This is of the form: /&ZeroWidthSpace;subscriptions/&ZeroWidthSpace;{subscriptionId}/&ZeroWidthSpace;resourceGroups/&ZeroWidthSpace;{resourceGroupName}/&ZeroWidthSpace;providers/&ZeroWidthSpace;Microsoft.Network/virtualNetworks/{virtualNetworkName}/subnets/{subnetName}                                                                                                            | genruntime.ResourceReference                          |
| proximityPlacementGroupID     | The ID for Proximity Placement Group.                                                                                                                                                                                                                                                                                                                                                                                                                                                        | string                                                |
| scaleSetEvictionPolicy        | This cannot be specified unless the scaleSetPriority is 'Spot'. If not specified, the default is 'Delete'.                                                                                                                                                                                                                                                                                                                                                                                   | [ScaleSetEvictionPolicy](#scalesetevictionpolicy)     |
| scaleSetPriority              | The Virtual Machine Scale Set priority. If not specified, the default is 'Regular'.                                                                                                                                                                                                                                                                                                                                                                                                          | [ScaleSetPriority](#scalesetpriority)                 |
| spotMaxPrice                  | Possible values are any decimal value greater than zero or -1 which indicates the willingness to pay any on-demand price. For more details on spot pricing, see [spot VMs pricing](https://docs.microsoft.com/azure/virtual-machines/spot-vms#pricing)                                                                                                                                                                                                                                       | float64                                               |
| tags                          | The tags to be persisted on the agent pool virtual machine scale set.                                                                                                                                                                                                                                                                                                                                                                                                                        | map[string]string                                     |
| type                          | The type of Agent Pool.                                                                                                                                                                                                                                                                                                                                                                                                                                                                      | [AgentPoolType](#agentpooltype)                       |
| upgradeSettings               | Settings for upgrading the agentpool                                                                                                                                                                                                                                                                                                                                                                                                                                                         | [AgentPoolUpgradeSettings](#agentpoolupgradesettings) |
| vmSize                        | VM size availability varies by region. If a node contains insufficient compute resources (memory, cpu, etc) pods might fail to run correctly. For more details on restricted VM sizes, see: https://docs.microsoft.com/azure/aks/quotas-skus-regions                                                                                                                                                                                                                                         | string                                                |
| vnetSubnetIDReference         | If this is not specified, a VNET and subnet will be generated and used. If no podSubnetID is specified, this applies to nodes and pods, otherwise it applies to just nodes. This is of the form: /&ZeroWidthSpace;subscriptions/&ZeroWidthSpace;{subscriptionId}/&ZeroWidthSpace;resourceGroups/&ZeroWidthSpace;{resourceGroupName}/&ZeroWidthSpace;providers/&ZeroWidthSpace;Microsoft.Network/virtualNetworks/{virtualNetworkName}/subnets/{subnetName}                                    | genruntime.ResourceReference                          |

<a id="managedclusteragentpoolprofile_status"></a>ManagedClusterAgentPoolProfile_STATUS
---------------------------------------------------------------------------------------

Profile for the container service agent pool.

Used by: [ManagedCluster_STATUS](#managedcluster_status).

| Property                  | Description                                                                                                                                                                                                                                                                                                                                                                                                                                                                                  | Type                                                                |
|---------------------------|----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|---------------------------------------------------------------------|
| availabilityZones         | The list of Availability zones to use for nodes. This can only be specified if the AgentPoolType property is 'VirtualMachineScaleSets'.                                                                                                                                                                                                                                                                                                                                                      | string[]                                                            |
| count                     | Number of agents (VMs) to host docker containers. Allowed values must be in the range of 0 to 1000 (inclusive) for user pools and in the range of 1 to 1000 (inclusive) for system pools. The default value is 1.                                                                                                                                                                                                                                                                            | int                                                                 |
| enableAutoScaling         | Whether to enable auto-scaler                                                                                                                                                                                                                                                                                                                                                                                                                                                                | bool                                                                |
| enableEncryptionAtHost    | This is only supported on certain VM sizes and in certain Azure regions. For more information, see: https://docs.microsoft.com/azure/aks/enable-host-encryption                                                                                                                                                                                                                                                                                                                              | bool                                                                |
| enableFIPS                | See [Add a FIPS-enabled node pool](https://docs.microsoft.com/azure/aks/use-multiple-node-pools#add-a-fips-enabled-node-pool-preview) for more details.                                                                                                                                                                                                                                                                                                                                      | bool                                                                |
| enableNodePublicIP        | Some scenarios may require nodes in a node pool to receive their own dedicated public IP addresses. A common scenario is for gaming workloads, where a console needs to make a direct connection to a cloud virtual machine to minimize hops. For more information see [assigning a public IP per node](https://docs.microsoft.com/azure/aks/use-multiple-node-pools#assign-a-public-ip-per-node-for-your-node-pools). The default is false.                                                 | bool                                                                |
| enableUltraSSD            | Whether to enable UltraSSD                                                                                                                                                                                                                                                                                                                                                                                                                                                                   | bool                                                                |
| gpuInstanceProfile        | GPUInstanceProfile to be used to specify GPU MIG instance profile for supported GPU VM SKU.                                                                                                                                                                                                                                                                                                                                                                                                  | [GPUInstanceProfile_STATUS](#gpuinstanceprofile_status)             |
| kubeletConfig             | The Kubelet configuration on the agent pool nodes.                                                                                                                                                                                                                                                                                                                                                                                                                                           | [KubeletConfig_STATUS](#kubeletconfig_status)                       |
| kubeletDiskType           | Determines the placement of emptyDir volumes, container runtime data root, and Kubelet ephemeral storage.                                                                                                                                                                                                                                                                                                                                                                                    | [KubeletDiskType_STATUS](#kubeletdisktype_status)                   |
| linuxOSConfig             | The OS configuration of Linux agent nodes.                                                                                                                                                                                                                                                                                                                                                                                                                                                   | [LinuxOSConfig_STATUS](#linuxosconfig_status)                       |
| maxCount                  | The maximum number of nodes for auto-scaling                                                                                                                                                                                                                                                                                                                                                                                                                                                 | int                                                                 |
| maxPods                   | The maximum number of pods that can run on a node.                                                                                                                                                                                                                                                                                                                                                                                                                                           | int                                                                 |
| minCount                  | The minimum number of nodes for auto-scaling                                                                                                                                                                                                                                                                                                                                                                                                                                                 | int                                                                 |
| mode                      | A cluster must have at least one 'System' Agent Pool at all times. For additional information on agent pool restrictions and best practices, see: https://docs.microsoft.com/azure/aks/use-system-pools                                                                                                                                                                                                                                                                                      | [AgentPoolMode_STATUS](#agentpoolmode_status)                       |
| name                      | Windows agent pool names must be 6 characters or less.                                                                                                                                                                                                                                                                                                                                                                                                                                       | string                                                              |
| nodeImageVersion          | The version of node image                                                                                                                                                                                                                                                                                                                                                                                                                                                                    | string                                                              |
| nodeLabels                | The node labels to be persisted across all nodes in agent pool.                                                                                                                                                                                                                                                                                                                                                                                                                              | map[string]string                                                   |
| nodePublicIPPrefixID      | This is of the form: /&ZeroWidthSpace;subscriptions/&ZeroWidthSpace;{subscriptionId}/&ZeroWidthSpace;resourceGroups/&ZeroWidthSpace;{resourceGroupName}/&ZeroWidthSpace;providers/&ZeroWidthSpace;Microsoft.Network/publicIPPrefixes/{publicIPPrefixName}                                                                                                                                                                                                                                    | string                                                              |
| nodeTaints                | The taints added to new nodes during node pool create and scale. For example, key=value:NoSchedule.                                                                                                                                                                                                                                                                                                                                                                                          | string[]                                                            |
| orchestratorVersion       | As a best practice, you should upgrade all node pools in an AKS cluster to the same Kubernetes version. The node pool version must have the same major version as the control plane. The node pool minor version must be within two minor versions of the control plane version. The node pool version cannot be greater than the control plane version. For more information see [upgrading a node pool](https://docs.microsoft.com/azure/aks/use-multiple-node-pools#upgrade-a-node-pool). | string                                                              |
| osDiskSizeGB              |                                                                                                                                                                                                                                                                                                                                                                                                                                                                                              | int                                                                 |
| osDiskType                | The default is 'Ephemeral' if the VM supports it and has a cache disk larger than the requested OSDiskSizeGB. Otherwise, defaults to 'Managed'. May not be changed after creation. For more information see [Ephemeral OS](https://docs.microsoft.com/azure/aks/cluster-configuration#ephemeral-os).                                                                                                                                                                                         | [OSDiskType_STATUS](#osdisktype_status)                             |
| osSKU                     | Specifies an OS SKU. This value must not be specified if OSType is Windows.                                                                                                                                                                                                                                                                                                                                                                                                                  | [OSSKU_STATUS](#ossku_status)                                       |
| osType                    | The operating system type. The default is Linux.                                                                                                                                                                                                                                                                                                                                                                                                                                             | [OSType_STATUS](#ostype_status)                                     |
| podSubnetID               | If omitted, pod IPs are statically assigned on the node subnet (see vnetSubnetID for more details). This is of the form: /&ZeroWidthSpace;subscriptions/&ZeroWidthSpace;{subscriptionId}/&ZeroWidthSpace;resourceGroups/&ZeroWidthSpace;{resourceGroupName}/&ZeroWidthSpace;providers/&ZeroWidthSpace;Microsoft.Network/virtualNetworks/{virtualNetworkName}/subnets/{subnetName}                                                                                                            | string                                                              |
| powerState                | Describes whether the Agent Pool is Running or Stopped                                                                                                                                                                                                                                                                                                                                                                                                                                       | [PowerState_STATUS](#powerstate_status)                             |
| provisioningState         | The current deployment or provisioning state.                                                                                                                                                                                                                                                                                                                                                                                                                                                | string                                                              |
| proximityPlacementGroupID | The ID for Proximity Placement Group.                                                                                                                                                                                                                                                                                                                                                                                                                                                        | string                                                              |
| scaleSetEvictionPolicy    | This cannot be specified unless the scaleSetPriority is 'Spot'. If not specified, the default is 'Delete'.                                                                                                                                                                                                                                                                                                                                                                                   | [ScaleSetEvictionPolicy_STATUS](#scalesetevictionpolicy_status)     |
| scaleSetPriority          | The Virtual Machine Scale Set priority. If not specified, the default is 'Regular'.                                                                                                                                                                                                                                                                                                                                                                                                          | [ScaleSetPriority_STATUS](#scalesetpriority_status)                 |
| spotMaxPrice              | Possible values are any decimal value greater than zero or -1 which indicates the willingness to pay any on-demand price. For more details on spot pricing, see [spot VMs pricing](https://docs.microsoft.com/azure/virtual-machines/spot-vms#pricing)                                                                                                                                                                                                                                       | float64                                                             |
| tags                      | The tags to be persisted on the agent pool virtual machine scale set.                                                                                                                                                                                                                                                                                                                                                                                                                        | map[string]string                                                   |
| type                      | The type of Agent Pool.                                                                                                                                                                                                                                                                                                                                                                                                                                                                      | [AgentPoolType_STATUS](#agentpooltype_status)                       |
| upgradeSettings           | Settings for upgrading the agentpool                                                                                                                                                                                                                                                                                                                                                                                                                                                         | [AgentPoolUpgradeSettings_STATUS](#agentpoolupgradesettings_status) |
| vmSize                    | VM size availability varies by region. If a node contains insufficient compute resources (memory, cpu, etc) pods might fail to run correctly. For more details on restricted VM sizes, see: https://docs.microsoft.com/azure/aks/quotas-skus-regions                                                                                                                                                                                                                                         | string                                                              |
| vnetSubnetID              | If this is not specified, a VNET and subnet will be generated and used. If no podSubnetID is specified, this applies to nodes and pods, otherwise it applies to just nodes. This is of the form: /&ZeroWidthSpace;subscriptions/&ZeroWidthSpace;{subscriptionId}/&ZeroWidthSpace;resourceGroups/&ZeroWidthSpace;{resourceGroupName}/&ZeroWidthSpace;providers/&ZeroWidthSpace;Microsoft.Network/virtualNetworks/{virtualNetworkName}/subnets/{subnetName}                                    | string                                                              |

<a id="managedclusterautoupgradeprofile"></a>ManagedClusterAutoUpgradeProfile
-----------------------------------------------------------------------------

Auto upgrade profile for a managed cluster.

Used by: [ManagedCluster_Spec](#managedcluster_spec).

| Property       | Description                                                                                                                                             | Type                                                                                                |
|----------------|---------------------------------------------------------------------------------------------------------------------------------------------------------|-----------------------------------------------------------------------------------------------------|
| upgradeChannel | For more information see [setting the AKS cluster auto-upgrade channel](https://docs.microsoft.com/azure/aks/upgrade-cluster#set-auto-upgrade-channel). | [ManagedClusterAutoUpgradeProfile_UpgradeChannel](#managedclusterautoupgradeprofile_upgradechannel) |

<a id="managedclusterautoupgradeprofile_status"></a>ManagedClusterAutoUpgradeProfile_STATUS
-------------------------------------------------------------------------------------------

Auto upgrade profile for a managed cluster.

Used by: [ManagedCluster_STATUS](#managedcluster_status).

| Property       | Description                                                                                                                                             | Type                                                                                                              |
|----------------|---------------------------------------------------------------------------------------------------------------------------------------------------------|-------------------------------------------------------------------------------------------------------------------|
| upgradeChannel | For more information see [setting the AKS cluster auto-upgrade channel](https://docs.microsoft.com/azure/aks/upgrade-cluster#set-auto-upgrade-channel). | [ManagedClusterAutoUpgradeProfile_UpgradeChannel_STATUS](#managedclusterautoupgradeprofile_upgradechannel_status) |

<a id="managedclusterautoupgradeprofile_upgradechannel"></a>ManagedClusterAutoUpgradeProfile_UpgradeChannel
-----------------------------------------------------------------------------------------------------------

Used by: [ManagedClusterAutoUpgradeProfile](#managedclusterautoupgradeprofile).

<a id="managedclusterautoupgradeprofile_upgradechannel_status"></a>ManagedClusterAutoUpgradeProfile_UpgradeChannel_STATUS
-------------------------------------------------------------------------------------------------------------------------

Used by: [ManagedClusterAutoUpgradeProfile_STATUS](#managedclusterautoupgradeprofile_status).

<a id="managedclusterhttpproxyconfig"></a>ManagedClusterHTTPProxyConfig
-----------------------------------------------------------------------

Cluster HTTP proxy configuration.

Used by: [ManagedCluster_Spec](#managedcluster_spec).

| Property   | Description                                                 | Type     |
|------------|-------------------------------------------------------------|----------|
| httpProxy  | The HTTP proxy server endpoint to use.                      | string   |
| httpsProxy | The HTTPS proxy server endpoint to use.                     | string   |
| noProxy    | The endpoints that should not go through proxy.             | string[] |
| trustedCa  | Alternative CA cert to use for connecting to proxy servers. | string   |

<a id="managedclusterhttpproxyconfig_status"></a>ManagedClusterHTTPProxyConfig_STATUS
-------------------------------------------------------------------------------------

Cluster HTTP proxy configuration.

Used by: [ManagedCluster_STATUS](#managedcluster_status).

| Property   | Description                                                 | Type     |
|------------|-------------------------------------------------------------|----------|
| httpProxy  | The HTTP proxy server endpoint to use.                      | string   |
| httpsProxy | The HTTPS proxy server endpoint to use.                     | string   |
| noProxy    | The endpoints that should not go through proxy.             | string[] |
| trustedCa  | Alternative CA cert to use for connecting to proxy servers. | string   |

<a id="managedclusteridentity"></a>ManagedClusterIdentity
---------------------------------------------------------

Identity for the managed cluster.

Used by: [ManagedCluster_Spec](#managedcluster_spec).

| Property               | Description                                                                                                                                                                                                                                                                                    | Type                                                          |
|------------------------|------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|---------------------------------------------------------------|
| type                   | For more information see [use managed identities in AKS](https://docs.microsoft.com/azure/aks/use-managed-identity).                                                                                                                                                                           | [ManagedClusterIdentity_Type](#managedclusteridentity_type)   |
| userAssignedIdentities | The keys must be ARM resource IDs in the form: '/&ZeroWidthSpace;subscriptions/&ZeroWidthSpace;{subscriptionId}/&ZeroWidthSpace;resourceGroups/&ZeroWidthSpace;{resourceGroupName}/&ZeroWidthSpace;providers/&ZeroWidthSpace;Microsoft.ManagedIdentity/userAssignedIdentities/{identityName}'. | [UserAssignedIdentityDetails[]](#userassignedidentitydetails) |

<a id="managedclusteridentity_status"></a>ManagedClusterIdentity_STATUS
-----------------------------------------------------------------------

Identity for the managed cluster.

Used by: [ManagedCluster_STATUS](#managedcluster_status).

| Property               | Description                                                                                                                                                                                                                                                                                    | Type                                                                                                                     |
|------------------------|------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|--------------------------------------------------------------------------------------------------------------------------|
| principalId            | The principal id of the system assigned identity which is used by master components.                                                                                                                                                                                                           | string                                                                                                                   |
| tenantId               | The tenant id of the system assigned identity which is used by master components.                                                                                                                                                                                                              | string                                                                                                                   |
| type                   | For more information see [use managed identities in AKS](https://docs.microsoft.com/azure/aks/use-managed-identity).                                                                                                                                                                           | [ManagedClusterIdentity_Type_STATUS](#managedclusteridentity_type_status)                                                |
| userAssignedIdentities | The keys must be ARM resource IDs in the form: '/&ZeroWidthSpace;subscriptions/&ZeroWidthSpace;{subscriptionId}/&ZeroWidthSpace;resourceGroups/&ZeroWidthSpace;{resourceGroupName}/&ZeroWidthSpace;providers/&ZeroWidthSpace;Microsoft.ManagedIdentity/userAssignedIdentities/{identityName}'. | [map[string]ManagedClusterIdentity_UserAssignedIdentities_STATUS](#managedclusteridentity_userassignedidentities_status) |

<a id="managedclusteridentity_type"></a>ManagedClusterIdentity_Type
-------------------------------------------------------------------

Used by: [ManagedClusterIdentity](#managedclusteridentity).

<a id="managedclusteridentity_type_status"></a>ManagedClusterIdentity_Type_STATUS
---------------------------------------------------------------------------------

Used by: [ManagedClusterIdentity_STATUS](#managedclusteridentity_status).

<a id="managedclusteridentity_userassignedidentities_status"></a>ManagedClusterIdentity_UserAssignedIdentities_STATUS
---------------------------------------------------------------------------------------------------------------------

Used by: [ManagedClusterIdentity_STATUS](#managedclusteridentity_status).

| Property    | Description                                 | Type   |
|-------------|---------------------------------------------|--------|
| clientId    | The client id of user assigned identity.    | string |
| principalId | The principal id of user assigned identity. | string |

<a id="managedclusterlist"></a>ManagedClusterList
-------------------------------------------------

Generator information: - Generated from: /containerservice/resource-manager/Microsoft.ContainerService/aks/stable/2021-05-01/managedClusters.json - ARM URI: /&ZeroWidthSpace;subscriptions/&ZeroWidthSpace;{subscriptionId}/&ZeroWidthSpace;resourceGroups/&ZeroWidthSpace;{resourceGroupName}/&ZeroWidthSpace;providers/&ZeroWidthSpace;Microsoft.ContainerService/managedClusters/{resourceName}

.

| Property        | Description | Type                                |
|-----------------|-------------|-------------------------------------|
| metav1.TypeMeta |             |                                     |
| metav1.ListMeta |             |                                     |
| items           |             | [ManagedCluster[]](#managedcluster) |

<a id="managedclusterloadbalancerprofile"></a>ManagedClusterLoadBalancerProfile
-------------------------------------------------------------------------------

Profile of the managed cluster load balancer.

Used by: [ContainerServiceNetworkProfile](#containerservicenetworkprofile).

| Property               | Description                                                                                                                                                                               | Type                                                                                                          |
|------------------------|-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|---------------------------------------------------------------------------------------------------------------|
| allocatedOutboundPorts | The desired number of allocated SNAT ports per VM. Allowed values are in the range of 0 to 64000 (inclusive). The default value is 0 which results in Azure dynamically allocating ports. | int                                                                                                           |
| effectiveOutboundIPs   | The effective outbound IP resources of the cluster load balancer.                                                                                                                         | [ResourceReference[]](#resourcereference)                                                                     |
| idleTimeoutInMinutes   | Desired outbound flow idle timeout in minutes. Allowed values are in the range of 4 to 120 (inclusive). The default value is 30 minutes.                                                  | int                                                                                                           |
| managedOutboundIPs     | Desired managed outbound IPs for the cluster load balancer.                                                                                                                               | [ManagedClusterLoadBalancerProfile_ManagedOutboundIPs](#managedclusterloadbalancerprofile_managedoutboundips) |
| outboundIPPrefixes     | Desired outbound IP Prefix resources for the cluster load balancer.                                                                                                                       | [ManagedClusterLoadBalancerProfile_OutboundIPPrefixes](#managedclusterloadbalancerprofile_outboundipprefixes) |
| outboundIPs            | Desired outbound IP resources for the cluster load balancer.                                                                                                                              | [ManagedClusterLoadBalancerProfile_OutboundIPs](#managedclusterloadbalancerprofile_outboundips)               |

<a id="managedclusterloadbalancerprofile_managedoutboundips"></a>ManagedClusterLoadBalancerProfile_ManagedOutboundIPs
---------------------------------------------------------------------------------------------------------------------

Used by: [ManagedClusterLoadBalancerProfile](#managedclusterloadbalancerprofile).

| Property | Description                                                                                                                                                                     | Type |
|----------|---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|------|
| count    | The desired number of outbound IPs created/managed by Azure for the cluster load balancer. Allowed values must be in the range of 1 to 100 (inclusive). The default value is 1. | int  |

<a id="managedclusterloadbalancerprofile_managedoutboundips_status"></a>ManagedClusterLoadBalancerProfile_ManagedOutboundIPs_STATUS
-----------------------------------------------------------------------------------------------------------------------------------

Used by: [ManagedClusterLoadBalancerProfile_STATUS](#managedclusterloadbalancerprofile_status).

| Property | Description                                                                                                                                                                     | Type |
|----------|---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|------|
| count    | The desired number of outbound IPs created/managed by Azure for the cluster load balancer. Allowed values must be in the range of 1 to 100 (inclusive). The default value is 1. | int  |

<a id="managedclusterloadbalancerprofile_outboundipprefixes"></a>ManagedClusterLoadBalancerProfile_OutboundIPPrefixes
---------------------------------------------------------------------------------------------------------------------

Used by: [ManagedClusterLoadBalancerProfile](#managedclusterloadbalancerprofile).

| Property         | Description                           | Type                                      |
|------------------|---------------------------------------|-------------------------------------------|
| publicIPPrefixes | A list of public IP prefix resources. | [ResourceReference[]](#resourcereference) |

<a id="managedclusterloadbalancerprofile_outboundipprefixes_status"></a>ManagedClusterLoadBalancerProfile_OutboundIPPrefixes_STATUS
-----------------------------------------------------------------------------------------------------------------------------------

Used by: [ManagedClusterLoadBalancerProfile_STATUS](#managedclusterloadbalancerprofile_status).

| Property         | Description                           | Type                                                    |
|------------------|---------------------------------------|---------------------------------------------------------|
| publicIPPrefixes | A list of public IP prefix resources. | [ResourceReference_STATUS[]](#resourcereference_status) |

<a id="managedclusterloadbalancerprofile_outboundips"></a>ManagedClusterLoadBalancerProfile_OutboundIPs
-------------------------------------------------------------------------------------------------------

Used by: [ManagedClusterLoadBalancerProfile](#managedclusterloadbalancerprofile).

| Property  | Description                    | Type                                      |
|-----------|--------------------------------|-------------------------------------------|
| publicIPs | A list of public IP resources. | [ResourceReference[]](#resourcereference) |

<a id="managedclusterloadbalancerprofile_outboundips_status"></a>ManagedClusterLoadBalancerProfile_OutboundIPs_STATUS
---------------------------------------------------------------------------------------------------------------------

Used by: [ManagedClusterLoadBalancerProfile_STATUS](#managedclusterloadbalancerprofile_status).

| Property  | Description                    | Type                                                    |
|-----------|--------------------------------|---------------------------------------------------------|
| publicIPs | A list of public IP resources. | [ResourceReference_STATUS[]](#resourcereference_status) |

<a id="managedclusterloadbalancerprofile_status"></a>ManagedClusterLoadBalancerProfile_STATUS
---------------------------------------------------------------------------------------------

Profile of the managed cluster load balancer.

Used by: [ContainerServiceNetworkProfile_STATUS](#containerservicenetworkprofile_status).

| Property               | Description                                                                                                                                                                               | Type                                                                                                                        |
|------------------------|-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|-----------------------------------------------------------------------------------------------------------------------------|
| allocatedOutboundPorts | The desired number of allocated SNAT ports per VM. Allowed values are in the range of 0 to 64000 (inclusive). The default value is 0 which results in Azure dynamically allocating ports. | int                                                                                                                         |
| effectiveOutboundIPs   | The effective outbound IP resources of the cluster load balancer.                                                                                                                         | [ResourceReference_STATUS[]](#resourcereference_status)                                                                     |
| idleTimeoutInMinutes   | Desired outbound flow idle timeout in minutes. Allowed values are in the range of 4 to 120 (inclusive). The default value is 30 minutes.                                                  | int                                                                                                                         |
| managedOutboundIPs     | Desired managed outbound IPs for the cluster load balancer.                                                                                                                               | [ManagedClusterLoadBalancerProfile_ManagedOutboundIPs_STATUS](#managedclusterloadbalancerprofile_managedoutboundips_status) |
| outboundIPPrefixes     | Desired outbound IP Prefix resources for the cluster load balancer.                                                                                                                       | [ManagedClusterLoadBalancerProfile_OutboundIPPrefixes_STATUS](#managedclusterloadbalancerprofile_outboundipprefixes_status) |
| outboundIPs            | Desired outbound IP resources for the cluster load balancer.                                                                                                                              | [ManagedClusterLoadBalancerProfile_OutboundIPs_STATUS](#managedclusterloadbalancerprofile_outboundips_status)               |

<a id="managedclusteroperatorsecrets"></a>ManagedClusterOperatorSecrets
-----------------------------------------------------------------------

Used by: [ManagedClusterOperatorSpec](#managedclusteroperatorspec).

| Property         | Description                                                                                                            | Type                         |
|------------------|------------------------------------------------------------------------------------------------------------------------|------------------------------|
| adminCredentials | indicates where the AdminCredentials secret should be placed. If omitted, the secret will not be retrieved from Azure. | genruntime.SecretDestination |
| userCredentials  | indicates where the UserCredentials secret should be placed. If omitted, the secret will not be retrieved from Azure.  | genruntime.SecretDestination |

<a id="managedclusteroperatorspec"></a>ManagedClusterOperatorSpec
-----------------------------------------------------------------

Details for configuring operator behavior. Fields in this struct are interpreted by the operator directly rather than being passed to Azure

Used by: [ManagedCluster_Spec](#managedcluster_spec).

| Property | Description                                        | Type                                                            |
|----------|----------------------------------------------------|-----------------------------------------------------------------|
| secrets  | configures where to place Azure generated secrets. | [ManagedClusterOperatorSecrets](#managedclusteroperatorsecrets) |

<a id="managedclusterpodidentity"></a>ManagedClusterPodIdentity
---------------------------------------------------------------

Details about the pod identity assigned to the Managed Cluster.

Used by: [ManagedClusterPodIdentityProfile](#managedclusterpodidentityprofile).

| Property        | Description                                                        | Type                                                       |
|-----------------|--------------------------------------------------------------------|------------------------------------------------------------|
| bindingSelector | The binding selector to use for the AzureIdentityBinding resource. | string                                                     |
| identity        | The user assigned identity details.                                | [UserAssignedIdentity](#userassignedidentity)<br/>Required |
| name            | The name of the pod identity.                                      | string<br/>Required                                        |
| namespace       | The namespace of the pod identity.                                 | string<br/>Required                                        |

<a id="managedclusterpodidentityexception"></a>ManagedClusterPodIdentityException
---------------------------------------------------------------------------------

See [disable AAD Pod Identity for a specific Pod/Application](https://azure.github.io/aad-pod-identity/docs/configure/application_exception/) for more details.

Used by: [ManagedClusterPodIdentityProfile](#managedclusterpodidentityprofile).

| Property  | Description                                  | Type                           |
|-----------|----------------------------------------------|--------------------------------|
| name      | The name of the pod identity exception.      | string<br/>Required            |
| namespace | The namespace of the pod identity exception. | string<br/>Required            |
| podLabels | The pod labels to match.                     | map[string]string<br/>Required |

<a id="managedclusterpodidentityexception_status"></a>ManagedClusterPodIdentityException_STATUS
-----------------------------------------------------------------------------------------------

See [disable AAD Pod Identity for a specific Pod/Application](https://azure.github.io/aad-pod-identity/docs/configure/application_exception/) for more details.

Used by: [ManagedClusterPodIdentityProfile_STATUS](#managedclusterpodidentityprofile_status).

| Property  | Description                                  | Type              |
|-----------|----------------------------------------------|-------------------|
| name      | The name of the pod identity exception.      | string            |
| namespace | The namespace of the pod identity exception. | string            |
| podLabels | The pod labels to match.                     | map[string]string |

<a id="managedclusterpodidentityprofile"></a>ManagedClusterPodIdentityProfile
-----------------------------------------------------------------------------

See [use AAD pod identity](https://docs.microsoft.com/azure/aks/use-azure-ad-pod-identity) for more details on pod identity integration.

Used by: [ManagedCluster_Spec](#managedcluster_spec).

| Property                       | Description                                                                                                                                                                                                                                                                                                                                                   | Type                                                                        |
|--------------------------------|---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|-----------------------------------------------------------------------------|
| allowNetworkPluginKubenet      | Running in Kubenet is disabled by default due to the security related nature of AAD Pod Identity and the risks of IP spoofing. See [using Kubenet network plugin with AAD Pod Identity](https://docs.microsoft.com/azure/aks/use-azure-ad-pod-identity#using-kubenet-network-plugin-with-azure-active-directory-pod-managed-identities) for more information. | bool                                                                        |
| enabled                        | Whether the pod identity addon is enabled.                                                                                                                                                                                                                                                                                                                    | bool                                                                        |
| userAssignedIdentities         | The pod identities to use in the cluster.                                                                                                                                                                                                                                                                                                                     | [ManagedClusterPodIdentity[]](#managedclusterpodidentity)                   |
| userAssignedIdentityExceptions | The pod identity exceptions to allow.                                                                                                                                                                                                                                                                                                                         | [ManagedClusterPodIdentityException[]](#managedclusterpodidentityexception) |

<a id="managedclusterpodidentityprofile_status"></a>ManagedClusterPodIdentityProfile_STATUS
-------------------------------------------------------------------------------------------

See [use AAD pod identity](https://docs.microsoft.com/azure/aks/use-azure-ad-pod-identity) for more details on pod identity integration.

Used by: [ManagedCluster_STATUS](#managedcluster_status).

| Property                       | Description                                                                                                                                                                                                                                                                                                                                                   | Type                                                                                      |
|--------------------------------|---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|-------------------------------------------------------------------------------------------|
| allowNetworkPluginKubenet      | Running in Kubenet is disabled by default due to the security related nature of AAD Pod Identity and the risks of IP spoofing. See [using Kubenet network plugin with AAD Pod Identity](https://docs.microsoft.com/azure/aks/use-azure-ad-pod-identity#using-kubenet-network-plugin-with-azure-active-directory-pod-managed-identities) for more information. | bool                                                                                      |
| enabled                        | Whether the pod identity addon is enabled.                                                                                                                                                                                                                                                                                                                    | bool                                                                                      |
| userAssignedIdentities         | The pod identities to use in the cluster.                                                                                                                                                                                                                                                                                                                     | [ManagedClusterPodIdentity_STATUS[]](#managedclusterpodidentity_status)                   |
| userAssignedIdentityExceptions | The pod identity exceptions to allow.                                                                                                                                                                                                                                                                                                                         | [ManagedClusterPodIdentityException_STATUS[]](#managedclusterpodidentityexception_status) |

<a id="managedclusterpodidentityprovisioningerrorbody_status"></a>ManagedClusterPodIdentityProvisioningErrorBody_STATUS
-----------------------------------------------------------------------------------------------------------------------

An error response from the pod identity provisioning.

Used by: [ManagedClusterPodIdentityProvisioningError_STATUS](#managedclusterpodidentityprovisioningerror_status).

| Property | Description                                                                                        | Type                                                                                                                                |
|----------|----------------------------------------------------------------------------------------------------|-------------------------------------------------------------------------------------------------------------------------------------|
| code     | An identifier for the error. Codes are invariant and are intended to be consumed programmatically. | string                                                                                                                              |
| details  | A list of additional details about the error.                                                      | [ManagedClusterPodIdentityProvisioningErrorBody_STATUS_Unrolled[]](#managedclusterpodidentityprovisioningerrorbody_status_unrolled) |
| message  | A message describing the error, intended to be suitable for display in a user interface.           | string                                                                                                                              |
| target   | The target of the particular error. For example, the name of the property in error.                | string                                                                                                                              |

<a id="managedclusterpodidentityprovisioningerrorbody_status_unrolled"></a>ManagedClusterPodIdentityProvisioningErrorBody_STATUS_Unrolled
-----------------------------------------------------------------------------------------------------------------------------------------

Used by: [ManagedClusterPodIdentityProvisioningErrorBody_STATUS](#managedclusterpodidentityprovisioningerrorbody_status).

| Property | Description                                                                                        | Type   |
|----------|----------------------------------------------------------------------------------------------------|--------|
| code     | An identifier for the error. Codes are invariant and are intended to be consumed programmatically. | string |
| message  | A message describing the error, intended to be suitable for display in a user interface.           | string |
| target   | The target of the particular error. For example, the name of the property in error.                | string |

<a id="managedclusterpodidentityprovisioningerror_status"></a>ManagedClusterPodIdentityProvisioningError_STATUS
---------------------------------------------------------------------------------------------------------------

An error response from the pod identity provisioning.

Used by: [ManagedClusterPodIdentity_ProvisioningInfo_STATUS](#managedclusterpodidentity_provisioninginfo_status).

| Property | Description              | Type                                                                                                            |
|----------|--------------------------|-----------------------------------------------------------------------------------------------------------------|
| error    | Details about the error. | [ManagedClusterPodIdentityProvisioningErrorBody_STATUS](#managedclusterpodidentityprovisioningerrorbody_status) |

<a id="managedclusterpodidentity_provisioninginfo_status"></a>ManagedClusterPodIdentity_ProvisioningInfo_STATUS
---------------------------------------------------------------------------------------------------------------

Used by: [ManagedClusterPodIdentity_STATUS](#managedclusterpodidentity_status).

| Property | Description                             | Type                                                                                                    |
|----------|-----------------------------------------|---------------------------------------------------------------------------------------------------------|
| error    | Pod identity assignment error (if any). | [ManagedClusterPodIdentityProvisioningError_STATUS](#managedclusterpodidentityprovisioningerror_status) |

<a id="managedclusterpodidentity_provisioningstate_status"></a>ManagedClusterPodIdentity_ProvisioningState_STATUS
-----------------------------------------------------------------------------------------------------------------

Used by: [ManagedClusterPodIdentity_STATUS](#managedclusterpodidentity_status).

<a id="managedclusterpodidentity_status"></a>ManagedClusterPodIdentity_STATUS
-----------------------------------------------------------------------------

Details about the pod identity assigned to the Managed Cluster.

Used by: [ManagedClusterPodIdentityProfile_STATUS](#managedclusterpodidentityprofile_status).

| Property          | Description                                                        | Type                                                                                                      |
|-------------------|--------------------------------------------------------------------|-----------------------------------------------------------------------------------------------------------|
| bindingSelector   | The binding selector to use for the AzureIdentityBinding resource. | string                                                                                                    |
| identity          | The user assigned identity details.                                | [UserAssignedIdentity_STATUS](#userassignedidentity_status)                                               |
| name              | The name of the pod identity.                                      | string                                                                                                    |
| namespace         | The namespace of the pod identity.                                 | string                                                                                                    |
| provisioningInfo  |                                                                    | [ManagedClusterPodIdentity_ProvisioningInfo_STATUS](#managedclusterpodidentity_provisioninginfo_status)   |
| provisioningState | The current provisioning state of the pod identity.                | [ManagedClusterPodIdentity_ProvisioningState_STATUS](#managedclusterpodidentity_provisioningstate_status) |

<a id="managedclusterproperties_autoscalerprofile"></a>ManagedClusterProperties_AutoScalerProfile
-------------------------------------------------------------------------------------------------

Used by: [ManagedCluster_Spec](#managedcluster_spec).

| Property                         | Description                                                                                                                                                                                                                                                                                                                         | Type                                                                                                        |
|----------------------------------|-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|-------------------------------------------------------------------------------------------------------------|
| balance-similar-node-groups      | Valid values are 'true' and 'false'                                                                                                                                                                                                                                                                                                 | string                                                                                                      |
| expander                         | If not specified, the default is 'random'. See [expanders](https://github.com/kubernetes/autoscaler/blob/master/cluster-autoscaler/FAQ.md#what-are-expanders) for more information.                                                                                                                                                 | [ManagedClusterProperties_AutoScalerProfile_Expander](#managedclusterproperties_autoscalerprofile_expander) |
| max-empty-bulk-delete            | The default is 10.                                                                                                                                                                                                                                                                                                                  | string                                                                                                      |
| max-graceful-termination-sec     | The default is 600.                                                                                                                                                                                                                                                                                                                 | string                                                                                                      |
| max-node-provision-time          | The default is '15m'. Values must be an integer followed by an 'm'. No unit of time other than minutes (m) is supported.                                                                                                                                                                                                            | string                                                                                                      |
| max-total-unready-percentage     | The default is 45. The maximum is 100 and the minimum is 0.                                                                                                                                                                                                                                                                         | string                                                                                                      |
| new-pod-scale-up-delay           | For scenarios like burst/batch scale where you don't want CA to act before the kubernetes scheduler could schedule all the pods, you can tell CA to ignore unscheduled pods before they're a certain age. The default is '0s'. Values must be an integer followed by a unit ('s' for seconds, 'm' for minutes, 'h' for hours, etc). | string                                                                                                      |
| ok-total-unready-count           | This must be an integer. The default is 3.                                                                                                                                                                                                                                                                                          | string                                                                                                      |
| scale-down-delay-after-add       | The default is '10m'. Values must be an integer followed by an 'm'. No unit of time other than minutes (m) is supported.                                                                                                                                                                                                            | string                                                                                                      |
| scale-down-delay-after-delete    | The default is the scan-interval. Values must be an integer followed by an 'm'. No unit of time other than minutes (m) is supported.                                                                                                                                                                                                | string                                                                                                      |
| scale-down-delay-after-failure   | The default is '3m'. Values must be an integer followed by an 'm'. No unit of time other than minutes (m) is supported.                                                                                                                                                                                                             | string                                                                                                      |
| scale-down-unneeded-time         | The default is '10m'. Values must be an integer followed by an 'm'. No unit of time other than minutes (m) is supported.                                                                                                                                                                                                            | string                                                                                                      |
| scale-down-unready-time          | The default is '20m'. Values must be an integer followed by an 'm'. No unit of time other than minutes (m) is supported.                                                                                                                                                                                                            | string                                                                                                      |
| scale-down-utilization-threshold | The default is '0.5'.                                                                                                                                                                                                                                                                                                               | string                                                                                                      |
| scan-interval                    | The default is '10'. Values must be an integer number of seconds.                                                                                                                                                                                                                                                                   | string                                                                                                      |
| skip-nodes-with-local-storage    | The default is true.                                                                                                                                                                                                                                                                                                                | string                                                                                                      |
| skip-nodes-with-system-pods      | The default is true.                                                                                                                                                                                                                                                                                                                | string                                                                                                      |

<a id="managedclusterproperties_autoscalerprofile_expander"></a>ManagedClusterProperties_AutoScalerProfile_Expander
-------------------------------------------------------------------------------------------------------------------

Used by: [ManagedClusterProperties_AutoScalerProfile](#managedclusterproperties_autoscalerprofile).

<a id="managedclusterproperties_autoscalerprofile_expander_status"></a>ManagedClusterProperties_AutoScalerProfile_Expander_STATUS
---------------------------------------------------------------------------------------------------------------------------------

Used by: [ManagedClusterProperties_AutoScalerProfile_STATUS](#managedclusterproperties_autoscalerprofile_status).

<a id="managedclusterproperties_autoscalerprofile_status"></a>ManagedClusterProperties_AutoScalerProfile_STATUS
---------------------------------------------------------------------------------------------------------------

Used by: [ManagedCluster_STATUS](#managedcluster_status).

| Property                         | Description                                                                                                                                                                                                                                                                                                                         | Type                                                                                                                      |
|----------------------------------|-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|---------------------------------------------------------------------------------------------------------------------------|
| balance-similar-node-groups      | Valid values are 'true' and 'false'                                                                                                                                                                                                                                                                                                 | string                                                                                                                    |
| expander                         | If not specified, the default is 'random'. See [expanders](https://github.com/kubernetes/autoscaler/blob/master/cluster-autoscaler/FAQ.md#what-are-expanders) for more information.                                                                                                                                                 | [ManagedClusterProperties_AutoScalerProfile_Expander_STATUS](#managedclusterproperties_autoscalerprofile_expander_status) |
| max-empty-bulk-delete            | The default is 10.                                                                                                                                                                                                                                                                                                                  | string                                                                                                                    |
| max-graceful-termination-sec     | The default is 600.                                                                                                                                                                                                                                                                                                                 | string                                                                                                                    |
| max-node-provision-time          | The default is '15m'. Values must be an integer followed by an 'm'. No unit of time other than minutes (m) is supported.                                                                                                                                                                                                            | string                                                                                                                    |
| max-total-unready-percentage     | The default is 45. The maximum is 100 and the minimum is 0.                                                                                                                                                                                                                                                                         | string                                                                                                                    |
| new-pod-scale-up-delay           | For scenarios like burst/batch scale where you don't want CA to act before the kubernetes scheduler could schedule all the pods, you can tell CA to ignore unscheduled pods before they're a certain age. The default is '0s'. Values must be an integer followed by a unit ('s' for seconds, 'm' for minutes, 'h' for hours, etc). | string                                                                                                                    |
| ok-total-unready-count           | This must be an integer. The default is 3.                                                                                                                                                                                                                                                                                          | string                                                                                                                    |
| scale-down-delay-after-add       | The default is '10m'. Values must be an integer followed by an 'm'. No unit of time other than minutes (m) is supported.                                                                                                                                                                                                            | string                                                                                                                    |
| scale-down-delay-after-delete    | The default is the scan-interval. Values must be an integer followed by an 'm'. No unit of time other than minutes (m) is supported.                                                                                                                                                                                                | string                                                                                                                    |
| scale-down-delay-after-failure   | The default is '3m'. Values must be an integer followed by an 'm'. No unit of time other than minutes (m) is supported.                                                                                                                                                                                                             | string                                                                                                                    |
| scale-down-unneeded-time         | The default is '10m'. Values must be an integer followed by an 'm'. No unit of time other than minutes (m) is supported.                                                                                                                                                                                                            | string                                                                                                                    |
| scale-down-unready-time          | The default is '20m'. Values must be an integer followed by an 'm'. No unit of time other than minutes (m) is supported.                                                                                                                                                                                                            | string                                                                                                                    |
| scale-down-utilization-threshold | The default is '0.5'.                                                                                                                                                                                                                                                                                                               | string                                                                                                                    |
| scan-interval                    | The default is '10'. Values must be an integer number of seconds.                                                                                                                                                                                                                                                                   | string                                                                                                                    |
| skip-nodes-with-local-storage    | The default is true.                                                                                                                                                                                                                                                                                                                | string                                                                                                                    |
| skip-nodes-with-system-pods      | The default is true.                                                                                                                                                                                                                                                                                                                | string                                                                                                                    |

<a id="managedclustersku"></a>ManagedClusterSKU
-----------------------------------------------

The SKU of a Managed Cluster.

Used by: [ManagedCluster_Spec](#managedcluster_spec).

| Property | Description                                                                                                                  | Type                                              |
|----------|------------------------------------------------------------------------------------------------------------------------------|---------------------------------------------------|
| name     | The name of a managed cluster SKU.                                                                                           | [ManagedClusterSKU_Name](#managedclustersku_name) |
| tier     | If not specified, the default is 'Free'. See [uptime SLA](https://docs.microsoft.com/azure/aks/uptime-sla) for more details. | [ManagedClusterSKU_Tier](#managedclustersku_tier) |

<a id="managedclustersku_name"></a>ManagedClusterSKU_Name
---------------------------------------------------------

Used by: [ManagedClusterSKU](#managedclustersku).

<a id="managedclustersku_name_status"></a>ManagedClusterSKU_Name_STATUS
-----------------------------------------------------------------------

Used by: [ManagedClusterSKU_STATUS](#managedclustersku_status).

<a id="managedclustersku_status"></a>ManagedClusterSKU_STATUS
-------------------------------------------------------------

The SKU of a Managed Cluster.

Used by: [ManagedCluster_STATUS](#managedcluster_status).

| Property | Description                                                                                                                  | Type                                                            |
|----------|------------------------------------------------------------------------------------------------------------------------------|-----------------------------------------------------------------|
| name     | The name of a managed cluster SKU.                                                                                           | [ManagedClusterSKU_Name_STATUS](#managedclustersku_name_status) |
| tier     | If not specified, the default is 'Free'. See [uptime SLA](https://docs.microsoft.com/azure/aks/uptime-sla) for more details. | [ManagedClusterSKU_Tier_STATUS](#managedclustersku_tier_status) |

<a id="managedclustersku_tier"></a>ManagedClusterSKU_Tier
---------------------------------------------------------

Used by: [ManagedClusterSKU](#managedclustersku).

<a id="managedclustersku_tier_status"></a>ManagedClusterSKU_Tier_STATUS
-----------------------------------------------------------------------

Used by: [ManagedClusterSKU_STATUS](#managedclustersku_status).

<a id="managedclusterserviceprincipalprofile"></a>ManagedClusterServicePrincipalProfile
---------------------------------------------------------------------------------------

Information about a service principal identity for the cluster to use for manipulating Azure APIs.

Used by: [ManagedCluster_Spec](#managedcluster_spec).

| Property | Description                                                              | Type                       |
|----------|--------------------------------------------------------------------------|----------------------------|
| clientId | The ID for the service principal.                                        | string<br/>Required        |
| secret   | The secret password associated with the service principal in plain text. | genruntime.SecretReference |

<a id="managedclusterserviceprincipalprofile_status"></a>ManagedClusterServicePrincipalProfile_STATUS
-----------------------------------------------------------------------------------------------------

Information about a service principal identity for the cluster to use for manipulating Azure APIs.

Used by: [ManagedCluster_STATUS](#managedcluster_status).

| Property | Description                       | Type   |
|----------|-----------------------------------|--------|
| clientId | The ID for the service principal. | string |

<a id="managedclusterwindowsprofile"></a>ManagedClusterWindowsProfile
---------------------------------------------------------------------

Profile for Windows VMs in the managed cluster.

Used by: [ManagedCluster_Spec](#managedcluster_spec).

| Property       | Description                                                                                                                                                                                                                                                                                                                                                                                                                                                        | Type                                                                                  |
|----------------|--------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|---------------------------------------------------------------------------------------|
| adminPassword  | Specifies the password of the administrator account. Minimum-length: 8 characters Max-length: 123 characters Complexity requirements: 3 out of 4 conditions below need to be fulfilled Has lower characters Has upper characters Has a digit Has a special character (Regex match \[\W_]) Disallowed values: "abc@123", "P@$$w0rd", "P@ssw0rd", "P@ssword123", "Pa$$word", "pass@word1", "Password!", "Password1", "Password22", "iloveyou!"                       | string                                                                                |
| adminUsername  | Specifies the name of the administrator account. Restriction: Cannot end in "." Disallowed values: "administrator", "admin", "user", "user1", "test", "user2", "test1", "user3", "admin1", "1", "123", "a", "actuser", "adm", "admin2", "aspnet", "backup", "console", "david", "guest", "john", "owner", "root", "server", "sql", "support", "support_388945a0", "sys", "test2", "test3", "user4", "user5". Minimum-length: 1 character Max-length: 20 characters | string<br/>Required                                                                   |
| enableCSIProxy | For more details on CSI proxy, see the [CSI proxy GitHub repo](https://github.com/kubernetes-csi/csi-proxy).                                                                                                                                                                                                                                                                                                                                                       | bool                                                                                  |
| licenseType    | The license type to use for Windows VMs. See [Azure Hybrid User Benefits](https://azure.microsoft.com/pricing/hybrid-benefit/faq/) for more details.                                                                                                                                                                                                                                                                                                               | [ManagedClusterWindowsProfile_LicenseType](#managedclusterwindowsprofile_licensetype) |

<a id="managedclusterwindowsprofile_licensetype"></a>ManagedClusterWindowsProfile_LicenseType
---------------------------------------------------------------------------------------------

Used by: [ManagedClusterWindowsProfile](#managedclusterwindowsprofile).

<a id="managedclusterwindowsprofile_licensetype_status"></a>ManagedClusterWindowsProfile_LicenseType_STATUS
-----------------------------------------------------------------------------------------------------------

Used by: [ManagedClusterWindowsProfile_STATUS](#managedclusterwindowsprofile_status).

<a id="managedclusterwindowsprofile_status"></a>ManagedClusterWindowsProfile_STATUS
-----------------------------------------------------------------------------------

Profile for Windows VMs in the managed cluster.

Used by: [ManagedCluster_STATUS](#managedcluster_status).

| Property       | Description                                                                                                                                                                                                                                                                                                                                                                                                                                                        | Type                                                                                                |
|----------------|--------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|-----------------------------------------------------------------------------------------------------|
| adminPassword  | Specifies the password of the administrator account. Minimum-length: 8 characters Max-length: 123 characters Complexity requirements: 3 out of 4 conditions below need to be fulfilled Has lower characters Has upper characters Has a digit Has a special character (Regex match \[\W_]) Disallowed values: "abc@123", "P@$$w0rd", "P@ssw0rd", "P@ssword123", "Pa$$word", "pass@word1", "Password!", "Password1", "Password22", "iloveyou!"                       | string                                                                                              |
| adminUsername  | Specifies the name of the administrator account. Restriction: Cannot end in "." Disallowed values: "administrator", "admin", "user", "user1", "test", "user2", "test1", "user3", "admin1", "1", "123", "a", "actuser", "adm", "admin2", "aspnet", "backup", "console", "david", "guest", "john", "owner", "root", "server", "sql", "support", "support_388945a0", "sys", "test2", "test3", "user4", "user5". Minimum-length: 1 character Max-length: 20 characters | string                                                                                              |
| enableCSIProxy | For more details on CSI proxy, see the [CSI proxy GitHub repo](https://github.com/kubernetes-csi/csi-proxy).                                                                                                                                                                                                                                                                                                                                                       | bool                                                                                                |
| licenseType    | The license type to use for Windows VMs. See [Azure Hybrid User Benefits](https://azure.microsoft.com/pricing/hybrid-benefit/faq/) for more details.                                                                                                                                                                                                                                                                                                               | [ManagedClusterWindowsProfile_LicenseType_STATUS](#managedclusterwindowsprofile_licensetype_status) |

<a id="managedcluster_status"></a>ManagedCluster_STATUS
-------------------------------------------------------

Managed cluster.

Used by: [ManagedCluster](#managedcluster).

| Property                | Description                                                                                                                                                                                                                                                                                                                                                                                 | Type                                                                                                    |
|-------------------------|---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|---------------------------------------------------------------------------------------------------------|
| aadProfile              | The Azure Active Directory configuration.                                                                                                                                                                                                                                                                                                                                                   | [ManagedClusterAADProfile_STATUS](#managedclusteraadprofile_status)                                     |
| addonProfiles           | The profile of managed cluster add-on.                                                                                                                                                                                                                                                                                                                                                      | [map[string]ManagedClusterAddonProfile_STATUS](#managedclusteraddonprofile_status)                      |
| agentPoolProfiles       | The agent pool properties.                                                                                                                                                                                                                                                                                                                                                                  | [ManagedClusterAgentPoolProfile_STATUS[]](#managedclusteragentpoolprofile_status)                       |
| apiServerAccessProfile  | The access profile for managed cluster API server.                                                                                                                                                                                                                                                                                                                                          | [ManagedClusterAPIServerAccessProfile_STATUS](#managedclusterapiserveraccessprofile_status)             |
| autoScalerProfile       | Parameters to be applied to the cluster-autoscaler when enabled                                                                                                                                                                                                                                                                                                                             | [ManagedClusterProperties_AutoScalerProfile_STATUS](#managedclusterproperties_autoscalerprofile_status) |
| autoUpgradeProfile      | The auto upgrade configuration.                                                                                                                                                                                                                                                                                                                                                             | [ManagedClusterAutoUpgradeProfile_STATUS](#managedclusterautoupgradeprofile_status)                     |
| azurePortalFQDN         | The Azure Portal requires certain Cross-Origin Resource Sharing (CORS) headers to be sent in some responses, which Kubernetes APIServer doesn't handle by default. This special FQDN supports CORS, allowing the Azure Portal to function properly.                                                                                                                                         | string                                                                                                  |
| conditions              | The observed state of the resource                                                                                                                                                                                                                                                                                                                                                          | conditions.Condition[]                                                                                  |
| disableLocalAccounts    | If set to true, getting static credentials will be disabled for this cluster. This must only be used on Managed Clusters that are AAD enabled. For more details see [disable local accounts](https://docs.microsoft.com/azure/aks/managed-aad#disable-local-accounts-preview).                                                                                                              | bool                                                                                                    |
| diskEncryptionSetID     | This is of the form: '/&ZeroWidthSpace;subscriptions/&ZeroWidthSpace;{subscriptionId}/&ZeroWidthSpace;resourceGroups/&ZeroWidthSpace;{resourceGroupName}/&ZeroWidthSpace;providers/&ZeroWidthSpace;Microsoft.Compute/diskEncryptionSets/{encryptionSetName}'                                                                                                                                | string                                                                                                  |
| dnsPrefix               | This cannot be updated once the Managed Cluster has been created.                                                                                                                                                                                                                                                                                                                           | string                                                                                                  |
| enablePodSecurityPolicy | (DEPRECATING) Whether to enable Kubernetes pod security policy (preview). This feature is set for removal on October 15th, 2020. Learn more at aka.ms/aks/azpodpolicy.                                                                                                                                                                                                                      | bool                                                                                                    |
| enableRBAC              | Whether to enable Kubernetes Role-Based Access Control.                                                                                                                                                                                                                                                                                                                                     | bool                                                                                                    |
| extendedLocation        | The extended location of the Virtual Machine.                                                                                                                                                                                                                                                                                                                                               | [ExtendedLocation_STATUS](#extendedlocation_status)                                                     |
| fqdn                    | The FQDN of the master pool.                                                                                                                                                                                                                                                                                                                                                                | string                                                                                                  |
| fqdnSubdomain           | This cannot be updated once the Managed Cluster has been created.                                                                                                                                                                                                                                                                                                                           | string                                                                                                  |
| httpProxyConfig         | Configurations for provisioning the cluster with HTTP proxy servers.                                                                                                                                                                                                                                                                                                                        | [ManagedClusterHTTPProxyConfig_STATUS](#managedclusterhttpproxyconfig_status)                           |
| id                      | Resource Id                                                                                                                                                                                                                                                                                                                                                                                 | string                                                                                                  |
| identity                | The identity of the managed cluster, if configured.                                                                                                                                                                                                                                                                                                                                         | [ManagedClusterIdentity_STATUS](#managedclusteridentity_status)                                         |
| identityProfile         | Identities associated with the cluster.                                                                                                                                                                                                                                                                                                                                                     | [map[string]UserAssignedIdentity_STATUS](#userassignedidentity_status)                                  |
| kubernetesVersion       | When you upgrade a supported AKS cluster, Kubernetes minor versions cannot be skipped. All upgrades must be performed sequentially by major version number. For example, upgrades between 1.14.x -> 1.15.x or 1.15.x -> 1.16.x are allowed, however 1.14.x -> 1.16.x is not allowed. See [upgrading an AKS cluster](https://docs.microsoft.com/azure/aks/upgrade-cluster) for more details. | string                                                                                                  |
| linuxProfile            | The profile for Linux VMs in the Managed Cluster.                                                                                                                                                                                                                                                                                                                                           | [ContainerServiceLinuxProfile_STATUS](#containerservicelinuxprofile_status)                             |
| location                | Resource location                                                                                                                                                                                                                                                                                                                                                                           | string                                                                                                  |
| maxAgentPools           | The max number of agent pools for the managed cluster.                                                                                                                                                                                                                                                                                                                                      | int                                                                                                     |
| name                    | Resource name                                                                                                                                                                                                                                                                                                                                                                               | string                                                                                                  |
| networkProfile          | The network configuration profile.                                                                                                                                                                                                                                                                                                                                                          | [ContainerServiceNetworkProfile_STATUS](#containerservicenetworkprofile_status)                         |
| nodeResourceGroup       | The name of the resource group containing agent pool nodes.                                                                                                                                                                                                                                                                                                                                 | string                                                                                                  |
| podIdentityProfile      | See [use AAD pod identity](https://docs.microsoft.com/azure/aks/use-azure-ad-pod-identity) for more details on AAD pod identity integration.                                                                                                                                                                                                                                                | [ManagedClusterPodIdentityProfile_STATUS](#managedclusterpodidentityprofile_status)                     |
| powerState              | The Power State of the cluster.                                                                                                                                                                                                                                                                                                                                                             | [PowerState_STATUS](#powerstate_status)                                                                 |
| privateFQDN             | The FQDN of private cluster.                                                                                                                                                                                                                                                                                                                                                                | string                                                                                                  |
| privateLinkResources    | Private link resources associated with the cluster.                                                                                                                                                                                                                                                                                                                                         | [PrivateLinkResource_STATUS[]](#privatelinkresource_status)                                             |
| provisioningState       | The current provisioning state.                                                                                                                                                                                                                                                                                                                                                             | string                                                                                                  |
| servicePrincipalProfile | Information about a service principal identity for the cluster to use for manipulating Azure APIs.                                                                                                                                                                                                                                                                                          | [ManagedClusterServicePrincipalProfile_STATUS](#managedclusterserviceprincipalprofile_status)           |
| sku                     | The managed cluster SKU.                                                                                                                                                                                                                                                                                                                                                                    | [ManagedClusterSKU_STATUS](#managedclustersku_status)                                                   |
| tags                    | Resource tags                                                                                                                                                                                                                                                                                                                                                                               | map[string]string                                                                                       |
| type                    | Resource type                                                                                                                                                                                                                                                                                                                                                                               | string                                                                                                  |
| windowsProfile          | The profile for Windows VMs in the Managed Cluster.                                                                                                                                                                                                                                                                                                                                         | [ManagedClusterWindowsProfile_STATUS](#managedclusterwindowsprofile_status)                             |

<a id="managedcluster_spec"></a>ManagedCluster_Spec
---------------------------------------------------

Used by: [ManagedCluster](#managedcluster).

| Property                     | Description                                                                                                                                                                                                                                                                                                                                                                                 | Type                                                                                      |
|------------------------------|---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|-------------------------------------------------------------------------------------------|
| aadProfile                   | The Azure Active Directory configuration.                                                                                                                                                                                                                                                                                                                                                   | [ManagedClusterAADProfile](#managedclusteraadprofile)                                     |
| addonProfiles                | The profile of managed cluster add-on.                                                                                                                                                                                                                                                                                                                                                      | [map[string]ManagedClusterAddonProfile](#managedclusteraddonprofile)                      |
| agentPoolProfiles            | The agent pool properties.                                                                                                                                                                                                                                                                                                                                                                  | [ManagedClusterAgentPoolProfile[]](#managedclusteragentpoolprofile)                       |
| apiServerAccessProfile       | The access profile for managed cluster API server.                                                                                                                                                                                                                                                                                                                                          | [ManagedClusterAPIServerAccessProfile](#managedclusterapiserveraccessprofile)             |
| autoScalerProfile            | Parameters to be applied to the cluster-autoscaler when enabled                                                                                                                                                                                                                                                                                                                             | [ManagedClusterProperties_AutoScalerProfile](#managedclusterproperties_autoscalerprofile) |
| autoUpgradeProfile           | The auto upgrade configuration.                                                                                                                                                                                                                                                                                                                                                             | [ManagedClusterAutoUpgradeProfile](#managedclusterautoupgradeprofile)                     |
| azureName                    | The name of the resource in Azure. This is often the same as the name of the resource in Kubernetes but it doesn't have to be.                                                                                                                                                                                                                                                              | string                                                                                    |
| disableLocalAccounts         | If set to true, getting static credentials will be disabled for this cluster. This must only be used on Managed Clusters that are AAD enabled. For more details see [disable local accounts](https://docs.microsoft.com/azure/aks/managed-aad#disable-local-accounts-preview).                                                                                                              | bool                                                                                      |
| diskEncryptionSetIDReference | This is of the form: '/&ZeroWidthSpace;subscriptions/&ZeroWidthSpace;{subscriptionId}/&ZeroWidthSpace;resourceGroups/&ZeroWidthSpace;{resourceGroupName}/&ZeroWidthSpace;providers/&ZeroWidthSpace;Microsoft.Compute/diskEncryptionSets/{encryptionSetName}'                                                                                                                                | genruntime.ResourceReference                                                              |
| dnsPrefix                    | This cannot be updated once the Managed Cluster has been created.                                                                                                                                                                                                                                                                                                                           | string                                                                                    |
| enablePodSecurityPolicy      | (DEPRECATING) Whether to enable Kubernetes pod security policy (preview). This feature is set for removal on October 15th, 2020. Learn more at aka.ms/aks/azpodpolicy.                                                                                                                                                                                                                      | bool                                                                                      |
| enableRBAC                   | Whether to enable Kubernetes Role-Based Access Control.                                                                                                                                                                                                                                                                                                                                     | bool                                                                                      |
| extendedLocation             | The extended location of the Virtual Machine.                                                                                                                                                                                                                                                                                                                                               | [ExtendedLocation](#extendedlocation)                                                     |
| fqdnSubdomain                | This cannot be updated once the Managed Cluster has been created.                                                                                                                                                                                                                                                                                                                           | string                                                                                    |
| httpProxyConfig              | Configurations for provisioning the cluster with HTTP proxy servers.                                                                                                                                                                                                                                                                                                                        | [ManagedClusterHTTPProxyConfig](#managedclusterhttpproxyconfig)                           |
| identity                     | The identity of the managed cluster, if configured.                                                                                                                                                                                                                                                                                                                                         | [ManagedClusterIdentity](#managedclusteridentity)                                         |
| identityProfile              | Identities associated with the cluster.                                                                                                                                                                                                                                                                                                                                                     | [map[string]UserAssignedIdentity](#userassignedidentity)                                  |
| kubernetesVersion            | When you upgrade a supported AKS cluster, Kubernetes minor versions cannot be skipped. All upgrades must be performed sequentially by major version number. For example, upgrades between 1.14.x -> 1.15.x or 1.15.x -> 1.16.x are allowed, however 1.14.x -> 1.16.x is not allowed. See [upgrading an AKS cluster](https://docs.microsoft.com/azure/aks/upgrade-cluster) for more details. | string                                                                                    |
| linuxProfile                 | The profile for Linux VMs in the Managed Cluster.                                                                                                                                                                                                                                                                                                                                           | [ContainerServiceLinuxProfile](#containerservicelinuxprofile)                             |
| location                     | Resource location                                                                                                                                                                                                                                                                                                                                                                           | string<br/>Required                                                                       |
| networkProfile               | The network configuration profile.                                                                                                                                                                                                                                                                                                                                                          | [ContainerServiceNetworkProfile](#containerservicenetworkprofile)                         |
| nodeResourceGroup            | The name of the resource group containing agent pool nodes.                                                                                                                                                                                                                                                                                                                                 | string                                                                                    |
| operatorSpec                 | The specification for configuring operator behavior. This field is interpreted by the operator and not passed directly to Azure                                                                                                                                                                                                                                                             | [ManagedClusterOperatorSpec](#managedclusteroperatorspec)                                 |
| owner                        | The owner of the resource. The owner controls where the resource goes when it is deployed. The owner also controls the resources lifecycle. When the owner is deleted the resource will also be deleted. Owner is expected to be a reference to a resources.azure.com/ResourceGroup resource                                                                                                | genruntime.KnownResourceReference<br/>Required                                            |
| podIdentityProfile           | See [use AAD pod identity](https://docs.microsoft.com/azure/aks/use-azure-ad-pod-identity) for more details on AAD pod identity integration.                                                                                                                                                                                                                                                | [ManagedClusterPodIdentityProfile](#managedclusterpodidentityprofile)                     |
| privateLinkResources         | Private link resources associated with the cluster.                                                                                                                                                                                                                                                                                                                                         | [PrivateLinkResource[]](#privatelinkresource)                                             |
| servicePrincipalProfile      | Information about a service principal identity for the cluster to use for manipulating Azure APIs.                                                                                                                                                                                                                                                                                          | [ManagedClusterServicePrincipalProfile](#managedclusterserviceprincipalprofile)           |
| sku                          | The managed cluster SKU.                                                                                                                                                                                                                                                                                                                                                                    | [ManagedClusterSKU](#managedclustersku)                                                   |
| tags                         | Resource tags                                                                                                                                                                                                                                                                                                                                                                               | map[string]string                                                                         |
| windowsProfile               | The profile for Windows VMs in the Managed Cluster.                                                                                                                                                                                                                                                                                                                                         | [ManagedClusterWindowsProfile](#managedclusterwindowsprofile)                             |

<a id="managedclustersagentpool"></a>ManagedClustersAgentPool
-------------------------------------------------------------

Generator information: - Generated from: /containerservice/resource-manager/Microsoft.ContainerService/aks/stable/2021-05-01/managedClusters.json - ARM URI: /&ZeroWidthSpace;subscriptions/&ZeroWidthSpace;{subscriptionId}/&ZeroWidthSpace;resourceGroups/&ZeroWidthSpace;{resourceGroupName}/&ZeroWidthSpace;providers/&ZeroWidthSpace;Microsoft.ContainerService/managedClusters/{resourceName}/agentPools/{agentPoolName}

Used by: [ManagedClustersAgentPoolList](#managedclustersagentpoollist).

| Property          | Description | Type                                                                  |
|-------------------|-------------|-----------------------------------------------------------------------|
| metav1.TypeMeta   |             |                                                                       |
| metav1.ObjectMeta |             |                                                                       |
| spec              |             | [ManagedClusters_AgentPool_Spec](#managedclusters_agentpool_spec)     |
| status            |             | [ManagedClusters_AgentPool_STATUS](#managedclusters_agentpool_status) |

<a id="managedclustersagentpoollist"></a>ManagedClustersAgentPoolList
---------------------------------------------------------------------

Generator information: - Generated from: /containerservice/resource-manager/Microsoft.ContainerService/aks/stable/2021-05-01/managedClusters.json - ARM URI: /&ZeroWidthSpace;subscriptions/&ZeroWidthSpace;{subscriptionId}/&ZeroWidthSpace;resourceGroups/&ZeroWidthSpace;{resourceGroupName}/&ZeroWidthSpace;providers/&ZeroWidthSpace;Microsoft.ContainerService/managedClusters/{resourceName}/agentPools/{agentPoolName}

.

| Property        | Description | Type                                                    |
|-----------------|-------------|---------------------------------------------------------|
| metav1.TypeMeta |             |                                                         |
| metav1.ListMeta |             |                                                         |
| items           |             | [ManagedClustersAgentPool[]](#managedclustersagentpool) |

<a id="managedclusters_agentpool_status"></a>ManagedClusters_AgentPool_STATUS
-----------------------------------------------------------------------------

Used by: [ManagedClustersAgentPool](#managedclustersagentpool).

| Property                  | Description                                                                                                                                                                                                                                                                                                                                                                                                                                                                                  | Type                                                                |
|---------------------------|----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|---------------------------------------------------------------------|
| availabilityZones         | The list of Availability zones to use for nodes. This can only be specified if the AgentPoolType property is 'VirtualMachineScaleSets'.                                                                                                                                                                                                                                                                                                                                                      | string[]                                                            |
| conditions                | The observed state of the resource                                                                                                                                                                                                                                                                                                                                                                                                                                                           | conditions.Condition[]                                              |
| count                     | Number of agents (VMs) to host docker containers. Allowed values must be in the range of 0 to 1000 (inclusive) for user pools and in the range of 1 to 1000 (inclusive) for system pools. The default value is 1.                                                                                                                                                                                                                                                                            | int                                                                 |
| enableAutoScaling         | Whether to enable auto-scaler                                                                                                                                                                                                                                                                                                                                                                                                                                                                | bool                                                                |
| enableEncryptionAtHost    | This is only supported on certain VM sizes and in certain Azure regions. For more information, see: https://docs.microsoft.com/azure/aks/enable-host-encryption                                                                                                                                                                                                                                                                                                                              | bool                                                                |
| enableFIPS                | See [Add a FIPS-enabled node pool](https://docs.microsoft.com/azure/aks/use-multiple-node-pools#add-a-fips-enabled-node-pool-preview) for more details.                                                                                                                                                                                                                                                                                                                                      | bool                                                                |
| enableNodePublicIP        | Some scenarios may require nodes in a node pool to receive their own dedicated public IP addresses. A common scenario is for gaming workloads, where a console needs to make a direct connection to a cloud virtual machine to minimize hops. For more information see [assigning a public IP per node](https://docs.microsoft.com/azure/aks/use-multiple-node-pools#assign-a-public-ip-per-node-for-your-node-pools). The default is false.                                                 | bool                                                                |
| enableUltraSSD            | Whether to enable UltraSSD                                                                                                                                                                                                                                                                                                                                                                                                                                                                   | bool                                                                |
| gpuInstanceProfile        | GPUInstanceProfile to be used to specify GPU MIG instance profile for supported GPU VM SKU.                                                                                                                                                                                                                                                                                                                                                                                                  | [GPUInstanceProfile_STATUS](#gpuinstanceprofile_status)             |
| id                        | Resource ID.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                 | string                                                              |
| kubeletConfig             | The Kubelet configuration on the agent pool nodes.                                                                                                                                                                                                                                                                                                                                                                                                                                           | [KubeletConfig_STATUS](#kubeletconfig_status)                       |
| kubeletDiskType           | Determines the placement of emptyDir volumes, container runtime data root, and Kubelet ephemeral storage.                                                                                                                                                                                                                                                                                                                                                                                    | [KubeletDiskType_STATUS](#kubeletdisktype_status)                   |
| linuxOSConfig             | The OS configuration of Linux agent nodes.                                                                                                                                                                                                                                                                                                                                                                                                                                                   | [LinuxOSConfig_STATUS](#linuxosconfig_status)                       |
| maxCount                  | The maximum number of nodes for auto-scaling                                                                                                                                                                                                                                                                                                                                                                                                                                                 | int                                                                 |
| maxPods                   | The maximum number of pods that can run on a node.                                                                                                                                                                                                                                                                                                                                                                                                                                           | int                                                                 |
| minCount                  | The minimum number of nodes for auto-scaling                                                                                                                                                                                                                                                                                                                                                                                                                                                 | int                                                                 |
| mode                      | A cluster must have at least one 'System' Agent Pool at all times. For additional information on agent pool restrictions and best practices, see: https://docs.microsoft.com/azure/aks/use-system-pools                                                                                                                                                                                                                                                                                      | [AgentPoolMode_STATUS](#agentpoolmode_status)                       |
| name                      | The name of the resource that is unique within a resource group. This name can be used to access the resource.                                                                                                                                                                                                                                                                                                                                                                               | string                                                              |
| nodeImageVersion          | The version of node image                                                                                                                                                                                                                                                                                                                                                                                                                                                                    | string                                                              |
| nodeLabels                | The node labels to be persisted across all nodes in agent pool.                                                                                                                                                                                                                                                                                                                                                                                                                              | map[string]string                                                   |
| nodePublicIPPrefixID      | This is of the form: /&ZeroWidthSpace;subscriptions/&ZeroWidthSpace;{subscriptionId}/&ZeroWidthSpace;resourceGroups/&ZeroWidthSpace;{resourceGroupName}/&ZeroWidthSpace;providers/&ZeroWidthSpace;Microsoft.Network/publicIPPrefixes/{publicIPPrefixName}                                                                                                                                                                                                                                    | string                                                              |
| nodeTaints                | The taints added to new nodes during node pool create and scale. For example, key=value:NoSchedule.                                                                                                                                                                                                                                                                                                                                                                                          | string[]                                                            |
| orchestratorVersion       | As a best practice, you should upgrade all node pools in an AKS cluster to the same Kubernetes version. The node pool version must have the same major version as the control plane. The node pool minor version must be within two minor versions of the control plane version. The node pool version cannot be greater than the control plane version. For more information see [upgrading a node pool](https://docs.microsoft.com/azure/aks/use-multiple-node-pools#upgrade-a-node-pool). | string                                                              |
| osDiskSizeGB              |                                                                                                                                                                                                                                                                                                                                                                                                                                                                                              | int                                                                 |
| osDiskType                | The default is 'Ephemeral' if the VM supports it and has a cache disk larger than the requested OSDiskSizeGB. Otherwise, defaults to 'Managed'. May not be changed after creation. For more information see [Ephemeral OS](https://docs.microsoft.com/azure/aks/cluster-configuration#ephemeral-os).                                                                                                                                                                                         | [OSDiskType_STATUS](#osdisktype_status)                             |
| osSKU                     | Specifies an OS SKU. This value must not be specified if OSType is Windows.                                                                                                                                                                                                                                                                                                                                                                                                                  | [OSSKU_STATUS](#ossku_status)                                       |
| osType                    | The operating system type. The default is Linux.                                                                                                                                                                                                                                                                                                                                                                                                                                             | [OSType_STATUS](#ostype_status)                                     |
| podSubnetID               | If omitted, pod IPs are statically assigned on the node subnet (see vnetSubnetID for more details). This is of the form: /&ZeroWidthSpace;subscriptions/&ZeroWidthSpace;{subscriptionId}/&ZeroWidthSpace;resourceGroups/&ZeroWidthSpace;{resourceGroupName}/&ZeroWidthSpace;providers/&ZeroWidthSpace;Microsoft.Network/virtualNetworks/{virtualNetworkName}/subnets/{subnetName}                                                                                                            | string                                                              |
| powerState                | Describes whether the Agent Pool is Running or Stopped                                                                                                                                                                                                                                                                                                                                                                                                                                       | [PowerState_STATUS](#powerstate_status)                             |
| properties_type           | The type of Agent Pool.                                                                                                                                                                                                                                                                                                                                                                                                                                                                      | [AgentPoolType_STATUS](#agentpooltype_status)                       |
| provisioningState         | The current deployment or provisioning state.                                                                                                                                                                                                                                                                                                                                                                                                                                                | string                                                              |
| proximityPlacementGroupID | The ID for Proximity Placement Group.                                                                                                                                                                                                                                                                                                                                                                                                                                                        | string                                                              |
| scaleSetEvictionPolicy    | This cannot be specified unless the scaleSetPriority is 'Spot'. If not specified, the default is 'Delete'.                                                                                                                                                                                                                                                                                                                                                                                   | [ScaleSetEvictionPolicy_STATUS](#scalesetevictionpolicy_status)     |
| scaleSetPriority          | The Virtual Machine Scale Set priority. If not specified, the default is 'Regular'.                                                                                                                                                                                                                                                                                                                                                                                                          | [ScaleSetPriority_STATUS](#scalesetpriority_status)                 |
| spotMaxPrice              | Possible values are any decimal value greater than zero or -1 which indicates the willingness to pay any on-demand price. For more details on spot pricing, see [spot VMs pricing](https://docs.microsoft.com/azure/virtual-machines/spot-vms#pricing)                                                                                                                                                                                                                                       | float64                                                             |
| tags                      | The tags to be persisted on the agent pool virtual machine scale set.                                                                                                                                                                                                                                                                                                                                                                                                                        | map[string]string                                                   |
| type                      | Resource type                                                                                                                                                                                                                                                                                                                                                                                                                                                                                | string                                                              |
| upgradeSettings           | Settings for upgrading the agentpool                                                                                                                                                                                                                                                                                                                                                                                                                                                         | [AgentPoolUpgradeSettings_STATUS](#agentpoolupgradesettings_status) |
| vmSize                    | VM size availability varies by region. If a node contains insufficient compute resources (memory, cpu, etc) pods might fail to run correctly. For more details on restricted VM sizes, see: https://docs.microsoft.com/azure/aks/quotas-skus-regions                                                                                                                                                                                                                                         | string                                                              |
| vnetSubnetID              | If this is not specified, a VNET and subnet will be generated and used. If no podSubnetID is specified, this applies to nodes and pods, otherwise it applies to just nodes. This is of the form: /&ZeroWidthSpace;subscriptions/&ZeroWidthSpace;{subscriptionId}/&ZeroWidthSpace;resourceGroups/&ZeroWidthSpace;{resourceGroupName}/&ZeroWidthSpace;providers/&ZeroWidthSpace;Microsoft.Network/virtualNetworks/{virtualNetworkName}/subnets/{subnetName}                                    | string                                                              |

<a id="managedclusters_agentpool_spec"></a>ManagedClusters_AgentPool_Spec
-------------------------------------------------------------------------

Used by: [ManagedClustersAgentPool](#managedclustersagentpool).

| Property                      | Description                                                                                                                                                                                                                                                                                                                                                                                                                                                                                  | Type                                                  |
|-------------------------------|----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|-------------------------------------------------------|
| availabilityZones             | The list of Availability zones to use for nodes. This can only be specified if the AgentPoolType property is 'VirtualMachineScaleSets'.                                                                                                                                                                                                                                                                                                                                                      | string[]                                              |
| azureName                     | The name of the resource in Azure. This is often the same as the name of the resource in Kubernetes but it doesn't have to be.                                                                                                                                                                                                                                                                                                                                                               | string                                                |
| count                         | Number of agents (VMs) to host docker containers. Allowed values must be in the range of 0 to 1000 (inclusive) for user pools and in the range of 1 to 1000 (inclusive) for system pools. The default value is 1.                                                                                                                                                                                                                                                                            | int                                                   |
| enableAutoScaling             | Whether to enable auto-scaler                                                                                                                                                                                                                                                                                                                                                                                                                                                                | bool                                                  |
| enableEncryptionAtHost        | This is only supported on certain VM sizes and in certain Azure regions. For more information, see: https://docs.microsoft.com/azure/aks/enable-host-encryption                                                                                                                                                                                                                                                                                                                              | bool                                                  |
| enableFIPS                    | See [Add a FIPS-enabled node pool](https://docs.microsoft.com/azure/aks/use-multiple-node-pools#add-a-fips-enabled-node-pool-preview) for more details.                                                                                                                                                                                                                                                                                                                                      | bool                                                  |
| enableNodePublicIP            | Some scenarios may require nodes in a node pool to receive their own dedicated public IP addresses. A common scenario is for gaming workloads, where a console needs to make a direct connection to a cloud virtual machine to minimize hops. For more information see [assigning a public IP per node](https://docs.microsoft.com/azure/aks/use-multiple-node-pools#assign-a-public-ip-per-node-for-your-node-pools). The default is false.                                                 | bool                                                  |
| enableUltraSSD                | Whether to enable UltraSSD                                                                                                                                                                                                                                                                                                                                                                                                                                                                   | bool                                                  |
| gpuInstanceProfile            | GPUInstanceProfile to be used to specify GPU MIG instance profile for supported GPU VM SKU.                                                                                                                                                                                                                                                                                                                                                                                                  | [GPUInstanceProfile](#gpuinstanceprofile)             |
| kubeletConfig                 | The Kubelet configuration on the agent pool nodes.                                                                                                                                                                                                                                                                                                                                                                                                                                           | [KubeletConfig](#kubeletconfig)                       |
| kubeletDiskType               | Determines the placement of emptyDir volumes, container runtime data root, and Kubelet ephemeral storage.                                                                                                                                                                                                                                                                                                                                                                                    | [KubeletDiskType](#kubeletdisktype)                   |
| linuxOSConfig                 | The OS configuration of Linux agent nodes.                                                                                                                                                                                                                                                                                                                                                                                                                                                   | [LinuxOSConfig](#linuxosconfig)                       |
| maxCount                      | The maximum number of nodes for auto-scaling                                                                                                                                                                                                                                                                                                                                                                                                                                                 | int                                                   |
| maxPods                       | The maximum number of pods that can run on a node.                                                                                                                                                                                                                                                                                                                                                                                                                                           | int                                                   |
| minCount                      | The minimum number of nodes for auto-scaling                                                                                                                                                                                                                                                                                                                                                                                                                                                 | int                                                   |
| mode                          | A cluster must have at least one 'System' Agent Pool at all times. For additional information on agent pool restrictions and best practices, see: https://docs.microsoft.com/azure/aks/use-system-pools                                                                                                                                                                                                                                                                                      | [AgentPoolMode](#agentpoolmode)                       |
| nodeLabels                    | The node labels to be persisted across all nodes in agent pool.                                                                                                                                                                                                                                                                                                                                                                                                                              | map[string]string                                     |
| nodePublicIPPrefixIDReference | This is of the form: /&ZeroWidthSpace;subscriptions/&ZeroWidthSpace;{subscriptionId}/&ZeroWidthSpace;resourceGroups/&ZeroWidthSpace;{resourceGroupName}/&ZeroWidthSpace;providers/&ZeroWidthSpace;Microsoft.Network/publicIPPrefixes/{publicIPPrefixName}                                                                                                                                                                                                                                    | genruntime.ResourceReference                          |
| nodeTaints                    | The taints added to new nodes during node pool create and scale. For example, key=value:NoSchedule.                                                                                                                                                                                                                                                                                                                                                                                          | string[]                                              |
| orchestratorVersion           | As a best practice, you should upgrade all node pools in an AKS cluster to the same Kubernetes version. The node pool version must have the same major version as the control plane. The node pool minor version must be within two minor versions of the control plane version. The node pool version cannot be greater than the control plane version. For more information see [upgrading a node pool](https://docs.microsoft.com/azure/aks/use-multiple-node-pools#upgrade-a-node-pool). | string                                                |
| osDiskSizeGB                  |                                                                                                                                                                                                                                                                                                                                                                                                                                                                                              | [ContainerServiceOSDisk](#containerserviceosdisk)     |
| osDiskType                    | The default is 'Ephemeral' if the VM supports it and has a cache disk larger than the requested OSDiskSizeGB. Otherwise, defaults to 'Managed'. May not be changed after creation. For more information see [Ephemeral OS](https://docs.microsoft.com/azure/aks/cluster-configuration#ephemeral-os).                                                                                                                                                                                         | [OSDiskType](#osdisktype)                             |
| osSKU                         | Specifies an OS SKU. This value must not be specified if OSType is Windows.                                                                                                                                                                                                                                                                                                                                                                                                                  | [OSSKU](#ossku)                                       |
| osType                        | The operating system type. The default is Linux.                                                                                                                                                                                                                                                                                                                                                                                                                                             | [OSType](#ostype)                                     |
| owner                         | The owner of the resource. The owner controls where the resource goes when it is deployed. The owner also controls the resources lifecycle. When the owner is deleted the resource will also be deleted. Owner is expected to be a reference to a containerservice.azure.com/ManagedCluster resource                                                                                                                                                                                         | genruntime.KnownResourceReference<br/>Required        |
| podSubnetIDReference          | If omitted, pod IPs are statically assigned on the node subnet (see vnetSubnetID for more details). This is of the form: /&ZeroWidthSpace;subscriptions/&ZeroWidthSpace;{subscriptionId}/&ZeroWidthSpace;resourceGroups/&ZeroWidthSpace;{resourceGroupName}/&ZeroWidthSpace;providers/&ZeroWidthSpace;Microsoft.Network/virtualNetworks/{virtualNetworkName}/subnets/{subnetName}                                                                                                            | genruntime.ResourceReference                          |
| proximityPlacementGroupID     | The ID for Proximity Placement Group.                                                                                                                                                                                                                                                                                                                                                                                                                                                        | string                                                |
| scaleSetEvictionPolicy        | This cannot be specified unless the scaleSetPriority is 'Spot'. If not specified, the default is 'Delete'.                                                                                                                                                                                                                                                                                                                                                                                   | [ScaleSetEvictionPolicy](#scalesetevictionpolicy)     |
| scaleSetPriority              | The Virtual Machine Scale Set priority. If not specified, the default is 'Regular'.                                                                                                                                                                                                                                                                                                                                                                                                          | [ScaleSetPriority](#scalesetpriority)                 |
| spotMaxPrice                  | Possible values are any decimal value greater than zero or -1 which indicates the willingness to pay any on-demand price. For more details on spot pricing, see [spot VMs pricing](https://docs.microsoft.com/azure/virtual-machines/spot-vms#pricing)                                                                                                                                                                                                                                       | float64                                               |
| tags                          | The tags to be persisted on the agent pool virtual machine scale set.                                                                                                                                                                                                                                                                                                                                                                                                                        | map[string]string                                     |
| type                          | The type of Agent Pool.                                                                                                                                                                                                                                                                                                                                                                                                                                                                      | [AgentPoolType](#agentpooltype)                       |
| upgradeSettings               | Settings for upgrading the agentpool                                                                                                                                                                                                                                                                                                                                                                                                                                                         | [AgentPoolUpgradeSettings](#agentpoolupgradesettings) |
| vmSize                        | VM size availability varies by region. If a node contains insufficient compute resources (memory, cpu, etc) pods might fail to run correctly. For more details on restricted VM sizes, see: https://docs.microsoft.com/azure/aks/quotas-skus-regions                                                                                                                                                                                                                                         | string                                                |
| vnetSubnetIDReference         | If this is not specified, a VNET and subnet will be generated and used. If no podSubnetID is specified, this applies to nodes and pods, otherwise it applies to just nodes. This is of the form: /&ZeroWidthSpace;subscriptions/&ZeroWidthSpace;{subscriptionId}/&ZeroWidthSpace;resourceGroups/&ZeroWidthSpace;{resourceGroupName}/&ZeroWidthSpace;providers/&ZeroWidthSpace;Microsoft.Network/virtualNetworks/{virtualNetworkName}/subnets/{subnetName}                                    | genruntime.ResourceReference                          |

<a id="osdisktype"></a>OSDiskType
---------------------------------

Used by: [ManagedClusterAgentPoolProfile](#managedclusteragentpoolprofile), and [ManagedClusters_AgentPool_Spec](#managedclusters_agentpool_spec).

<a id="osdisktype_status"></a>OSDiskType_STATUS
-----------------------------------------------

Used by: [ManagedClusterAgentPoolProfile_STATUS](#managedclusteragentpoolprofile_status), and [ManagedClusters_AgentPool_STATUS](#managedclusters_agentpool_status).

<a id="ossku"></a>OSSKU
-----------------------

Used by: [ManagedClusterAgentPoolProfile](#managedclusteragentpoolprofile), and [ManagedClusters_AgentPool_Spec](#managedclusters_agentpool_spec).

<a id="ossku_status"></a>OSSKU_STATUS
-------------------------------------

Used by: [ManagedClusterAgentPoolProfile_STATUS](#managedclusteragentpoolprofile_status), and [ManagedClusters_AgentPool_STATUS](#managedclusters_agentpool_status).

<a id="ostype"></a>OSType
-------------------------

Used by: [ManagedClusterAgentPoolProfile](#managedclusteragentpoolprofile), and [ManagedClusters_AgentPool_Spec](#managedclusters_agentpool_spec).

<a id="ostype_status"></a>OSType_STATUS
---------------------------------------

Used by: [ManagedClusterAgentPoolProfile_STATUS](#managedclusteragentpoolprofile_status), and [ManagedClusters_AgentPool_STATUS](#managedclusters_agentpool_status).

<a id="powerstate_code_status"></a>PowerState_Code_STATUS
---------------------------------------------------------

Used by: [PowerState_STATUS](#powerstate_status).

<a id="powerstate_status"></a>PowerState_STATUS
-----------------------------------------------

Describes the Power State of the cluster

Used by: [ManagedClusterAgentPoolProfile_STATUS](#managedclusteragentpoolprofile_status), [ManagedCluster_STATUS](#managedcluster_status), and [ManagedClusters_AgentPool_STATUS](#managedclusters_agentpool_status).

| Property | Description                                     | Type                                              |
|----------|-------------------------------------------------|---------------------------------------------------|
| code     | Tells whether the cluster is Running or Stopped | [PowerState_Code_STATUS](#powerstate_code_status) |

<a id="privatelinkresource"></a>PrivateLinkResource
---------------------------------------------------

A private link resource

Used by: [ManagedCluster_Spec](#managedcluster_spec).

| Property        | Description                            | Type                         |
|-----------------|----------------------------------------|------------------------------|
| groupId         | The group ID of the resource.          | string                       |
| name            | The name of the private link resource. | string                       |
| reference       | The ID of the private link resource.   | genruntime.ResourceReference |
| requiredMembers | The RequiredMembers of the resource    | string[]                     |
| type            | The resource type.                     | string                       |

<a id="privatelinkresource_status"></a>PrivateLinkResource_STATUS
-----------------------------------------------------------------

A private link resource

Used by: [ManagedCluster_STATUS](#managedcluster_status).

| Property             | Description                                                                                | Type     |
|----------------------|--------------------------------------------------------------------------------------------|----------|
| groupId              | The group ID of the resource.                                                              | string   |
| id                   | The ID of the private link resource.                                                       | string   |
| name                 | The name of the private link resource.                                                     | string   |
| privateLinkServiceID | The private link service ID of the resource, this field is exposed only to NRP internally. | string   |
| requiredMembers      | The RequiredMembers of the resource                                                        | string[] |
| type                 | The resource type.                                                                         | string   |

<a id="resourcereference"></a>ResourceReference
-----------------------------------------------

A reference to an Azure resource.

Used by: [ManagedClusterLoadBalancerProfile](#managedclusterloadbalancerprofile), [ManagedClusterLoadBalancerProfile_OutboundIPPrefixes](#managedclusterloadbalancerprofile_outboundipprefixes), and [ManagedClusterLoadBalancerProfile_OutboundIPs](#managedclusterloadbalancerprofile_outboundips).

| Property  | Description                            | Type                         |
|-----------|----------------------------------------|------------------------------|
| reference | The fully qualified Azure resource id. | genruntime.ResourceReference |

<a id="resourcereference_status"></a>ResourceReference_STATUS
-------------------------------------------------------------

A reference to an Azure resource.

Used by: [ManagedClusterLoadBalancerProfile_OutboundIPPrefixes_STATUS](#managedclusterloadbalancerprofile_outboundipprefixes_status), [ManagedClusterLoadBalancerProfile_OutboundIPs_STATUS](#managedclusterloadbalancerprofile_outboundips_status), and [ManagedClusterLoadBalancerProfile_STATUS](#managedclusterloadbalancerprofile_status).

| Property | Description                            | Type   |
|----------|----------------------------------------|--------|
| id       | The fully qualified Azure resource id. | string |

<a id="scalesetevictionpolicy"></a>ScaleSetEvictionPolicy
---------------------------------------------------------

Used by: [ManagedClusterAgentPoolProfile](#managedclusteragentpoolprofile), and [ManagedClusters_AgentPool_Spec](#managedclusters_agentpool_spec).

<a id="scalesetevictionpolicy_status"></a>ScaleSetEvictionPolicy_STATUS
-----------------------------------------------------------------------

Used by: [ManagedClusterAgentPoolProfile_STATUS](#managedclusteragentpoolprofile_status), and [ManagedClusters_AgentPool_STATUS](#managedclusters_agentpool_status).

<a id="scalesetpriority"></a>ScaleSetPriority
---------------------------------------------

Used by: [ManagedClusterAgentPoolProfile](#managedclusteragentpoolprofile), and [ManagedClusters_AgentPool_Spec](#managedclusters_agentpool_spec).

<a id="scalesetpriority_status"></a>ScaleSetPriority_STATUS
-----------------------------------------------------------

Used by: [ManagedClusterAgentPoolProfile_STATUS](#managedclusteragentpoolprofile_status), and [ManagedClusters_AgentPool_STATUS](#managedclusters_agentpool_status).

<a id="sysctlconfig"></a>SysctlConfig
-------------------------------------

Sysctl settings for Linux agent nodes.

Used by: [LinuxOSConfig](#linuxosconfig).

| Property                       | Description                                        | Type   |
|--------------------------------|----------------------------------------------------|--------|
| fsAioMaxNr                     | Sysctl setting fs.aio-max-nr.                      | int    |
| fsFileMax                      | Sysctl setting fs.file-max.                        | int    |
| fsInotifyMaxUserWatches        | Sysctl setting fs.inotify.max_user_watches.        | int    |
| fsNrOpen                       | Sysctl setting fs.nr_open.                         | int    |
| kernelThreadsMax               | Sysctl setting kernel.threads-max.                 | int    |
| netCoreNetdevMaxBacklog        | Sysctl setting net.core.netdev_max_backlog.        | int    |
| netCoreOptmemMax               | Sysctl setting net.core.optmem_max.                | int    |
| netCoreRmemDefault             | Sysctl setting net.core.rmem_default.              | int    |
| netCoreRmemMax                 | Sysctl setting net.core.rmem_max.                  | int    |
| netCoreSomaxconn               | Sysctl setting net.core.somaxconn.                 | int    |
| netCoreWmemDefault             | Sysctl setting net.core.wmem_default.              | int    |
| netCoreWmemMax                 | Sysctl setting net.core.wmem_max.                  | int    |
| netIpv4IpLocalPortRange        | Sysctl setting net.ipv4.ip_local_port_range.       | string |
| netIpv4NeighDefaultGcThresh1   | Sysctl setting net.ipv4.neigh.default.gc_thresh1.  | int    |
| netIpv4NeighDefaultGcThresh2   | Sysctl setting net.ipv4.neigh.default.gc_thresh2.  | int    |
| netIpv4NeighDefaultGcThresh3   | Sysctl setting net.ipv4.neigh.default.gc_thresh3.  | int    |
| netIpv4TcpFinTimeout           | Sysctl setting net.ipv4.tcp_fin_timeout.           | int    |
| netIpv4TcpKeepaliveProbes      | Sysctl setting net.ipv4.tcp_keepalive_probes.      | int    |
| netIpv4TcpKeepaliveTime        | Sysctl setting net.ipv4.tcp_keepalive_time.        | int    |
| netIpv4TcpMaxSynBacklog        | Sysctl setting net.ipv4.tcp_max_syn_backlog.       | int    |
| netIpv4TcpMaxTwBuckets         | Sysctl setting net.ipv4.tcp_max_tw_buckets.        | int    |
| netIpv4TcpTwReuse              | Sysctl setting net.ipv4.tcp_tw_reuse.              | bool   |
| netIpv4TcpkeepaliveIntvl       | Sysctl setting net.ipv4.tcp_keepalive_intvl.       | int    |
| netNetfilterNfConntrackBuckets | Sysctl setting net.netfilter.nf_conntrack_buckets. | int    |
| netNetfilterNfConntrackMax     | Sysctl setting net.netfilter.nf_conntrack_max.     | int    |
| vmMaxMapCount                  | Sysctl setting vm.max_map_count.                   | int    |
| vmSwappiness                   | Sysctl setting vm.swappiness.                      | int    |
| vmVfsCachePressure             | Sysctl setting vm.vfs_cache_pressure.              | int    |

<a id="sysctlconfig_status"></a>SysctlConfig_STATUS
---------------------------------------------------

Sysctl settings for Linux agent nodes.

Used by: [LinuxOSConfig_STATUS](#linuxosconfig_status).

| Property                       | Description                                        | Type   |
|--------------------------------|----------------------------------------------------|--------|
| fsAioMaxNr                     | Sysctl setting fs.aio-max-nr.                      | int    |
| fsFileMax                      | Sysctl setting fs.file-max.                        | int    |
| fsInotifyMaxUserWatches        | Sysctl setting fs.inotify.max_user_watches.        | int    |
| fsNrOpen                       | Sysctl setting fs.nr_open.                         | int    |
| kernelThreadsMax               | Sysctl setting kernel.threads-max.                 | int    |
| netCoreNetdevMaxBacklog        | Sysctl setting net.core.netdev_max_backlog.        | int    |
| netCoreOptmemMax               | Sysctl setting net.core.optmem_max.                | int    |
| netCoreRmemDefault             | Sysctl setting net.core.rmem_default.              | int    |
| netCoreRmemMax                 | Sysctl setting net.core.rmem_max.                  | int    |
| netCoreSomaxconn               | Sysctl setting net.core.somaxconn.                 | int    |
| netCoreWmemDefault             | Sysctl setting net.core.wmem_default.              | int    |
| netCoreWmemMax                 | Sysctl setting net.core.wmem_max.                  | int    |
| netIpv4IpLocalPortRange        | Sysctl setting net.ipv4.ip_local_port_range.       | string |
| netIpv4NeighDefaultGcThresh1   | Sysctl setting net.ipv4.neigh.default.gc_thresh1.  | int    |
| netIpv4NeighDefaultGcThresh2   | Sysctl setting net.ipv4.neigh.default.gc_thresh2.  | int    |
| netIpv4NeighDefaultGcThresh3   | Sysctl setting net.ipv4.neigh.default.gc_thresh3.  | int    |
| netIpv4TcpFinTimeout           | Sysctl setting net.ipv4.tcp_fin_timeout.           | int    |
| netIpv4TcpKeepaliveProbes      | Sysctl setting net.ipv4.tcp_keepalive_probes.      | int    |
| netIpv4TcpKeepaliveTime        | Sysctl setting net.ipv4.tcp_keepalive_time.        | int    |
| netIpv4TcpMaxSynBacklog        | Sysctl setting net.ipv4.tcp_max_syn_backlog.       | int    |
| netIpv4TcpMaxTwBuckets         | Sysctl setting net.ipv4.tcp_max_tw_buckets.        | int    |
| netIpv4TcpTwReuse              | Sysctl setting net.ipv4.tcp_tw_reuse.              | bool   |
| netIpv4TcpkeepaliveIntvl       | Sysctl setting net.ipv4.tcp_keepalive_intvl.       | int    |
| netNetfilterNfConntrackBuckets | Sysctl setting net.netfilter.nf_conntrack_buckets. | int    |
| netNetfilterNfConntrackMax     | Sysctl setting net.netfilter.nf_conntrack_max.     | int    |
| vmMaxMapCount                  | Sysctl setting vm.max_map_count.                   | int    |
| vmSwappiness                   | Sysctl setting vm.swappiness.                      | int    |
| vmVfsCachePressure             | Sysctl setting vm.vfs_cache_pressure.              | int    |

<a id="userassignedidentity"></a>UserAssignedIdentity
-----------------------------------------------------

Details about a user assigned identity.

Used by: [ManagedClusterPodIdentity](#managedclusterpodidentity), and [ManagedCluster_Spec](#managedcluster_spec).

| Property          | Description                                    | Type                         |
|-------------------|------------------------------------------------|------------------------------|
| clientId          | The client ID of the user assigned identity.   | string                       |
| objectId          | The object ID of the user assigned identity.   | string                       |
| resourceReference | The resource ID of the user assigned identity. | genruntime.ResourceReference |

<a id="userassignedidentitydetails"></a>UserAssignedIdentityDetails
-------------------------------------------------------------------

Information about the user assigned identity for the resource

Used by: [ManagedClusterIdentity](#managedclusteridentity).

| Property  | Description | Type                         |
|-----------|-------------|------------------------------|
| reference |             | genruntime.ResourceReference |

<a id="userassignedidentity_status"></a>UserAssignedIdentity_STATUS
-------------------------------------------------------------------

Details about a user assigned identity.

Used by: [ManagedClusterAddonProfile_STATUS](#managedclusteraddonprofile_status), [ManagedClusterPodIdentity_STATUS](#managedclusterpodidentity_status), and [ManagedCluster_STATUS](#managedcluster_status).

| Property   | Description                                    | Type   |
|------------|------------------------------------------------|--------|
| clientId   | The client ID of the user assigned identity.   | string |
| objectId   | The object ID of the user assigned identity.   | string |
| resourceId | The resource ID of the user assigned identity. | string |
