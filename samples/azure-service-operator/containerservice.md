v1api20210501
=============

| Metadata | Value                      |
|----------|----------------------------|
| Group    | containerservice.azure.com |
| Version  | v1api20210501              |

<a id="APIVersion"></a>APIVersion
---------------------------------

| Value        | Description |
|--------------|-------------|
| "2021-05-01" |             |

<a id="AgentPoolMode"></a>AgentPoolMode
---------------------------------------

Used by: [ManagedClusterAgentPoolProfile](#ManagedClusterAgentPoolProfile), [ManagedClusterAgentPoolProfileProperties_ARM](#ManagedClusterAgentPoolProfileProperties_ARM), [ManagedClusterAgentPoolProfile_ARM](#ManagedClusterAgentPoolProfile_ARM), and [ManagedClusters_AgentPool_Spec](#ManagedClusters_AgentPool_Spec).

| Value    | Description |
|----------|-------------|
| "System" |             |
| "User"   |             |

<a id="AgentPoolMode_STATUS"></a>AgentPoolMode_STATUS
-----------------------------------------------------

Used by: [ManagedClusterAgentPoolProfileProperties_STATUS_ARM](#ManagedClusterAgentPoolProfileProperties_STATUS_ARM), [ManagedClusterAgentPoolProfile_STATUS](#ManagedClusterAgentPoolProfile_STATUS), [ManagedClusterAgentPoolProfile_STATUS_ARM](#ManagedClusterAgentPoolProfile_STATUS_ARM), and [ManagedClusters_AgentPool_STATUS](#ManagedClusters_AgentPool_STATUS).

| Value    | Description |
|----------|-------------|
| "System" |             |
| "User"   |             |

<a id="AgentPoolType"></a>AgentPoolType
---------------------------------------

Used by: [ManagedClusterAgentPoolProfile](#ManagedClusterAgentPoolProfile), [ManagedClusterAgentPoolProfileProperties_ARM](#ManagedClusterAgentPoolProfileProperties_ARM), [ManagedClusterAgentPoolProfile_ARM](#ManagedClusterAgentPoolProfile_ARM), and [ManagedClusters_AgentPool_Spec](#ManagedClusters_AgentPool_Spec).

| Value                     | Description |
|---------------------------|-------------|
| "AvailabilitySet"         |             |
| "VirtualMachineScaleSets" |             |

<a id="AgentPoolType_STATUS"></a>AgentPoolType_STATUS
-----------------------------------------------------

Used by: [ManagedClusterAgentPoolProfileProperties_STATUS_ARM](#ManagedClusterAgentPoolProfileProperties_STATUS_ARM), [ManagedClusterAgentPoolProfile_STATUS](#ManagedClusterAgentPoolProfile_STATUS), [ManagedClusterAgentPoolProfile_STATUS_ARM](#ManagedClusterAgentPoolProfile_STATUS_ARM), and [ManagedClusters_AgentPool_STATUS](#ManagedClusters_AgentPool_STATUS).

| Value                     | Description |
|---------------------------|-------------|
| "AvailabilitySet"         |             |
| "VirtualMachineScaleSets" |             |

<a id="AgentPoolUpgradeSettings"></a>AgentPoolUpgradeSettings
-------------------------------------------------------------

Settings for upgrading an agentpool

Used by: [ManagedClusterAgentPoolProfile](#ManagedClusterAgentPoolProfile), and [ManagedClusters_AgentPool_Spec](#ManagedClusters_AgentPool_Spec).

| Property | Description                                                                                                                                                                                                                                                                                                                                                                                                            | Type   |
|----------|------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|--------|
| maxSurge | This can either be set to an integer (e.g. '5') or a percentage (e.g. '50%'). If a percentage is specified, it is the percentage of the total agent pool size at the time of the upgrade. For percentages, fractional nodes are rounded up. If not specified, the default is 1. For more information, including best practices, see: https://docs.microsoft.com/azure/aks/upgrade-cluster#customize-node-surge-upgrade | string |

<a id="AgentPoolUpgradeSettings_ARM"></a>AgentPoolUpgradeSettings_ARM
---------------------------------------------------------------------

Settings for upgrading an agentpool

Used by: [ManagedClusterAgentPoolProfileProperties_ARM](#ManagedClusterAgentPoolProfileProperties_ARM), and [ManagedClusterAgentPoolProfile_ARM](#ManagedClusterAgentPoolProfile_ARM).

| Property | Description                                                                                                                                                                                                                                                                                                                                                                                                            | Type   |
|----------|------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|--------|
| maxSurge | This can either be set to an integer (e.g. '5') or a percentage (e.g. '50%'). If a percentage is specified, it is the percentage of the total agent pool size at the time of the upgrade. For percentages, fractional nodes are rounded up. If not specified, the default is 1. For more information, including best practices, see: https://docs.microsoft.com/azure/aks/upgrade-cluster#customize-node-surge-upgrade | string |

<a id="AgentPoolUpgradeSettings_STATUS"></a>AgentPoolUpgradeSettings_STATUS
---------------------------------------------------------------------------

Settings for upgrading an agentpool

Used by: [ManagedClusterAgentPoolProfile_STATUS](#ManagedClusterAgentPoolProfile_STATUS), and [ManagedClusters_AgentPool_STATUS](#ManagedClusters_AgentPool_STATUS).

| Property | Description                                                                                                                                                                                                                                                                                                                                                                                                            | Type   |
|----------|------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|--------|
| maxSurge | This can either be set to an integer (e.g. '5') or a percentage (e.g. '50%'). If a percentage is specified, it is the percentage of the total agent pool size at the time of the upgrade. For percentages, fractional nodes are rounded up. If not specified, the default is 1. For more information, including best practices, see: https://docs.microsoft.com/azure/aks/upgrade-cluster#customize-node-surge-upgrade | string |

<a id="AgentPoolUpgradeSettings_STATUS_ARM"></a>AgentPoolUpgradeSettings_STATUS_ARM
-----------------------------------------------------------------------------------

Settings for upgrading an agentpool

Used by: [ManagedClusterAgentPoolProfileProperties_STATUS_ARM](#ManagedClusterAgentPoolProfileProperties_STATUS_ARM), and [ManagedClusterAgentPoolProfile_STATUS_ARM](#ManagedClusterAgentPoolProfile_STATUS_ARM).

| Property | Description                                                                                                                                                                                                                                                                                                                                                                                                            | Type   |
|----------|------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|--------|
| maxSurge | This can either be set to an integer (e.g. '5') or a percentage (e.g. '50%'). If a percentage is specified, it is the percentage of the total agent pool size at the time of the upgrade. For percentages, fractional nodes are rounded up. If not specified, the default is 1. For more information, including best practices, see: https://docs.microsoft.com/azure/aks/upgrade-cluster#customize-node-surge-upgrade | string |

<a id="ContainerServiceLinuxProfile"></a>ContainerServiceLinuxProfile
---------------------------------------------------------------------

Profile for Linux VMs in the container service cluster.

Used by: [ManagedCluster_Spec](#ManagedCluster_Spec).

| Property      | Description                                                 | Type                                                                               |
|---------------|-------------------------------------------------------------|------------------------------------------------------------------------------------|
| adminUsername | The administrator username to use for Linux VMs.            | string<br/>Required                                                                |
| ssh           | The SSH configuration for Linux-based VMs running on Azure. | [ContainerServiceSshConfiguration](#ContainerServiceSshConfiguration)<br/>Required |

<a id="ContainerServiceLinuxProfile_ARM"></a>ContainerServiceLinuxProfile_ARM
-----------------------------------------------------------------------------

Profile for Linux VMs in the container service cluster.

Used by: [ManagedClusterProperties_ARM](#ManagedClusterProperties_ARM).

| Property      | Description                                                 | Type                                                                          |
|---------------|-------------------------------------------------------------|-------------------------------------------------------------------------------|
| adminUsername | The administrator username to use for Linux VMs.            | string                                                                        |
| ssh           | The SSH configuration for Linux-based VMs running on Azure. | [ContainerServiceSshConfiguration_ARM](#ContainerServiceSshConfiguration_ARM) |

<a id="ContainerServiceLinuxProfile_STATUS"></a>ContainerServiceLinuxProfile_STATUS
-----------------------------------------------------------------------------------

Profile for Linux VMs in the container service cluster.

Used by: [ManagedCluster_STATUS](#ManagedCluster_STATUS).

| Property      | Description                                                 | Type                                                                                |
|---------------|-------------------------------------------------------------|-------------------------------------------------------------------------------------|
| adminUsername | The administrator username to use for Linux VMs.            | string                                                                              |
| ssh           | The SSH configuration for Linux-based VMs running on Azure. | [ContainerServiceSshConfiguration_STATUS](#ContainerServiceSshConfiguration_STATUS) |

<a id="ContainerServiceLinuxProfile_STATUS_ARM"></a>ContainerServiceLinuxProfile_STATUS_ARM
-------------------------------------------------------------------------------------------

Profile for Linux VMs in the container service cluster.

Used by: [ManagedClusterProperties_STATUS_ARM](#ManagedClusterProperties_STATUS_ARM).

| Property      | Description                                                 | Type                                                                                        |
|---------------|-------------------------------------------------------------|---------------------------------------------------------------------------------------------|
| adminUsername | The administrator username to use for Linux VMs.            | string                                                                                      |
| ssh           | The SSH configuration for Linux-based VMs running on Azure. | [ContainerServiceSshConfiguration_STATUS_ARM](#ContainerServiceSshConfiguration_STATUS_ARM) |

<a id="ContainerServiceNetworkProfile"></a>ContainerServiceNetworkProfile
-------------------------------------------------------------------------

Profile of network configuration.

Used by: [ManagedCluster_Spec](#ManagedCluster_Spec).

| Property            | Description                                                                                                                                                                           | Type                                                                                              |
|---------------------|---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|---------------------------------------------------------------------------------------------------|
| dnsServiceIP        | An IP address assigned to the Kubernetes DNS service. It must be within the Kubernetes service address range specified in serviceCidr.                                                | string                                                                                            |
| dockerBridgeCidr    | A CIDR notation IP range assigned to the Docker bridge network. It must not overlap with any Subnet IP ranges or the Kubernetes service address range.                                | string                                                                                            |
| loadBalancerProfile | Profile of the cluster load balancer.                                                                                                                                                 | [ManagedClusterLoadBalancerProfile](#ManagedClusterLoadBalancerProfile)                           |
| loadBalancerSku     | The default is 'standard'. See [Azure Load Balancer SKUs](https://docs.microsoft.com/azure/load-balancer/skus) for more information about the differences between load balancer SKUs. | [ContainerServiceNetworkProfile_LoadBalancerSku](#ContainerServiceNetworkProfile_LoadBalancerSku) |
| networkMode         | This cannot be specified if networkPlugin is anything other than 'azure'.                                                                                                             | [ContainerServiceNetworkProfile_NetworkMode](#ContainerServiceNetworkProfile_NetworkMode)         |
| networkPlugin       | Network plugin used for building the Kubernetes network.                                                                                                                              | [ContainerServiceNetworkProfile_NetworkPlugin](#ContainerServiceNetworkProfile_NetworkPlugin)     |
| networkPolicy       | Network policy used for building the Kubernetes network.                                                                                                                              | [ContainerServiceNetworkProfile_NetworkPolicy](#ContainerServiceNetworkProfile_NetworkPolicy)     |
| outboundType        | This can only be set at cluster creation time and cannot be changed later. For more information see [egress outbound type](https://docs.microsoft.com/azure/aks/egress-outboundtype). | [ContainerServiceNetworkProfile_OutboundType](#ContainerServiceNetworkProfile_OutboundType)       |
| podCidr             | A CIDR notation IP range from which to assign pod IPs when kubenet is used.                                                                                                           | string                                                                                            |
| serviceCidr         | A CIDR notation IP range from which to assign service cluster IPs. It must not overlap with any Subnet IP ranges.                                                                     | string                                                                                            |

<a id="ContainerServiceNetworkProfile_ARM"></a>ContainerServiceNetworkProfile_ARM
---------------------------------------------------------------------------------

Profile of network configuration.

Used by: [ManagedClusterProperties_ARM](#ManagedClusterProperties_ARM).

| Property            | Description                                                                                                                                                                           | Type                                                                                              |
|---------------------|---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|---------------------------------------------------------------------------------------------------|
| dnsServiceIP        | An IP address assigned to the Kubernetes DNS service. It must be within the Kubernetes service address range specified in serviceCidr.                                                | string                                                                                            |
| dockerBridgeCidr    | A CIDR notation IP range assigned to the Docker bridge network. It must not overlap with any Subnet IP ranges or the Kubernetes service address range.                                | string                                                                                            |
| loadBalancerProfile | Profile of the cluster load balancer.                                                                                                                                                 | [ManagedClusterLoadBalancerProfile_ARM](#ManagedClusterLoadBalancerProfile_ARM)                   |
| loadBalancerSku     | The default is 'standard'. See [Azure Load Balancer SKUs](https://docs.microsoft.com/azure/load-balancer/skus) for more information about the differences between load balancer SKUs. | [ContainerServiceNetworkProfile_LoadBalancerSku](#ContainerServiceNetworkProfile_LoadBalancerSku) |
| networkMode         | This cannot be specified if networkPlugin is anything other than 'azure'.                                                                                                             | [ContainerServiceNetworkProfile_NetworkMode](#ContainerServiceNetworkProfile_NetworkMode)         |
| networkPlugin       | Network plugin used for building the Kubernetes network.                                                                                                                              | [ContainerServiceNetworkProfile_NetworkPlugin](#ContainerServiceNetworkProfile_NetworkPlugin)     |
| networkPolicy       | Network policy used for building the Kubernetes network.                                                                                                                              | [ContainerServiceNetworkProfile_NetworkPolicy](#ContainerServiceNetworkProfile_NetworkPolicy)     |
| outboundType        | This can only be set at cluster creation time and cannot be changed later. For more information see [egress outbound type](https://docs.microsoft.com/azure/aks/egress-outboundtype). | [ContainerServiceNetworkProfile_OutboundType](#ContainerServiceNetworkProfile_OutboundType)       |
| podCidr             | A CIDR notation IP range from which to assign pod IPs when kubenet is used.                                                                                                           | string                                                                                            |
| serviceCidr         | A CIDR notation IP range from which to assign service cluster IPs. It must not overlap with any Subnet IP ranges.                                                                     | string                                                                                            |

<a id="ContainerServiceNetworkProfile_LoadBalancerSku"></a>ContainerServiceNetworkProfile_LoadBalancerSku
---------------------------------------------------------------------------------------------------------

Used by: [ContainerServiceNetworkProfile](#ContainerServiceNetworkProfile), and [ContainerServiceNetworkProfile_ARM](#ContainerServiceNetworkProfile_ARM).

| Value      | Description |
|------------|-------------|
| "basic"    |             |
| "standard" |             |

<a id="ContainerServiceNetworkProfile_LoadBalancerSku_STATUS"></a>ContainerServiceNetworkProfile_LoadBalancerSku_STATUS
-----------------------------------------------------------------------------------------------------------------------

Used by: [ContainerServiceNetworkProfile_STATUS](#ContainerServiceNetworkProfile_STATUS), and [ContainerServiceNetworkProfile_STATUS_ARM](#ContainerServiceNetworkProfile_STATUS_ARM).

| Value      | Description |
|------------|-------------|
| "basic"    |             |
| "standard" |             |

<a id="ContainerServiceNetworkProfile_NetworkMode"></a>ContainerServiceNetworkProfile_NetworkMode
-------------------------------------------------------------------------------------------------

Used by: [ContainerServiceNetworkProfile](#ContainerServiceNetworkProfile), and [ContainerServiceNetworkProfile_ARM](#ContainerServiceNetworkProfile_ARM).

| Value         | Description |
|---------------|-------------|
| "bridge"      |             |
| "transparent" |             |

<a id="ContainerServiceNetworkProfile_NetworkMode_STATUS"></a>ContainerServiceNetworkProfile_NetworkMode_STATUS
---------------------------------------------------------------------------------------------------------------

Used by: [ContainerServiceNetworkProfile_STATUS](#ContainerServiceNetworkProfile_STATUS), and [ContainerServiceNetworkProfile_STATUS_ARM](#ContainerServiceNetworkProfile_STATUS_ARM).

| Value         | Description |
|---------------|-------------|
| "bridge"      |             |
| "transparent" |             |

<a id="ContainerServiceNetworkProfile_NetworkPlugin"></a>ContainerServiceNetworkProfile_NetworkPlugin
-----------------------------------------------------------------------------------------------------

Used by: [ContainerServiceNetworkProfile](#ContainerServiceNetworkProfile), and [ContainerServiceNetworkProfile_ARM](#ContainerServiceNetworkProfile_ARM).

| Value     | Description |
|-----------|-------------|
| "azure"   |             |
| "kubenet" |             |

<a id="ContainerServiceNetworkProfile_NetworkPlugin_STATUS"></a>ContainerServiceNetworkProfile_NetworkPlugin_STATUS
-------------------------------------------------------------------------------------------------------------------

Used by: [ContainerServiceNetworkProfile_STATUS](#ContainerServiceNetworkProfile_STATUS), and [ContainerServiceNetworkProfile_STATUS_ARM](#ContainerServiceNetworkProfile_STATUS_ARM).

| Value     | Description |
|-----------|-------------|
| "azure"   |             |
| "kubenet" |             |

<a id="ContainerServiceNetworkProfile_NetworkPolicy"></a>ContainerServiceNetworkProfile_NetworkPolicy
-----------------------------------------------------------------------------------------------------

Used by: [ContainerServiceNetworkProfile](#ContainerServiceNetworkProfile), and [ContainerServiceNetworkProfile_ARM](#ContainerServiceNetworkProfile_ARM).

| Value    | Description |
|----------|-------------|
| "azure"  |             |
| "calico" |             |

<a id="ContainerServiceNetworkProfile_NetworkPolicy_STATUS"></a>ContainerServiceNetworkProfile_NetworkPolicy_STATUS
-------------------------------------------------------------------------------------------------------------------

Used by: [ContainerServiceNetworkProfile_STATUS](#ContainerServiceNetworkProfile_STATUS), and [ContainerServiceNetworkProfile_STATUS_ARM](#ContainerServiceNetworkProfile_STATUS_ARM).

| Value    | Description |
|----------|-------------|
| "azure"  |             |
| "calico" |             |

<a id="ContainerServiceNetworkProfile_OutboundType"></a>ContainerServiceNetworkProfile_OutboundType
---------------------------------------------------------------------------------------------------

Used by: [ContainerServiceNetworkProfile](#ContainerServiceNetworkProfile), and [ContainerServiceNetworkProfile_ARM](#ContainerServiceNetworkProfile_ARM).

| Value                | Description |
|----------------------|-------------|
| "loadBalancer"       |             |
| "userDefinedRouting" |             |

<a id="ContainerServiceNetworkProfile_OutboundType_STATUS"></a>ContainerServiceNetworkProfile_OutboundType_STATUS
-----------------------------------------------------------------------------------------------------------------

Used by: [ContainerServiceNetworkProfile_STATUS](#ContainerServiceNetworkProfile_STATUS), and [ContainerServiceNetworkProfile_STATUS_ARM](#ContainerServiceNetworkProfile_STATUS_ARM).

| Value                | Description |
|----------------------|-------------|
| "loadBalancer"       |             |
| "userDefinedRouting" |             |

<a id="ContainerServiceNetworkProfile_STATUS"></a>ContainerServiceNetworkProfile_STATUS
---------------------------------------------------------------------------------------

Profile of network configuration.

Used by: [ManagedCluster_STATUS](#ManagedCluster_STATUS).

| Property            | Description                                                                                                                                                                           | Type                                                                                                            |
|---------------------|---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|-----------------------------------------------------------------------------------------------------------------|
| dnsServiceIP        | An IP address assigned to the Kubernetes DNS service. It must be within the Kubernetes service address range specified in serviceCidr.                                                | string                                                                                                          |
| dockerBridgeCidr    | A CIDR notation IP range assigned to the Docker bridge network. It must not overlap with any Subnet IP ranges or the Kubernetes service address range.                                | string                                                                                                          |
| loadBalancerProfile | Profile of the cluster load balancer.                                                                                                                                                 | [ManagedClusterLoadBalancerProfile_STATUS](#ManagedClusterLoadBalancerProfile_STATUS)                           |
| loadBalancerSku     | The default is 'standard'. See [Azure Load Balancer SKUs](https://docs.microsoft.com/azure/load-balancer/skus) for more information about the differences between load balancer SKUs. | [ContainerServiceNetworkProfile_LoadBalancerSku_STATUS](#ContainerServiceNetworkProfile_LoadBalancerSku_STATUS) |
| networkMode         | This cannot be specified if networkPlugin is anything other than 'azure'.                                                                                                             | [ContainerServiceNetworkProfile_NetworkMode_STATUS](#ContainerServiceNetworkProfile_NetworkMode_STATUS)         |
| networkPlugin       | Network plugin used for building the Kubernetes network.                                                                                                                              | [ContainerServiceNetworkProfile_NetworkPlugin_STATUS](#ContainerServiceNetworkProfile_NetworkPlugin_STATUS)     |
| networkPolicy       | Network policy used for building the Kubernetes network.                                                                                                                              | [ContainerServiceNetworkProfile_NetworkPolicy_STATUS](#ContainerServiceNetworkProfile_NetworkPolicy_STATUS)     |
| outboundType        | This can only be set at cluster creation time and cannot be changed later. For more information see [egress outbound type](https://docs.microsoft.com/azure/aks/egress-outboundtype). | [ContainerServiceNetworkProfile_OutboundType_STATUS](#ContainerServiceNetworkProfile_OutboundType_STATUS)       |
| podCidr             | A CIDR notation IP range from which to assign pod IPs when kubenet is used.                                                                                                           | string                                                                                                          |
| serviceCidr         | A CIDR notation IP range from which to assign service cluster IPs. It must not overlap with any Subnet IP ranges.                                                                     | string                                                                                                          |

<a id="ContainerServiceNetworkProfile_STATUS_ARM"></a>ContainerServiceNetworkProfile_STATUS_ARM
-----------------------------------------------------------------------------------------------

Profile of network configuration.

Used by: [ManagedClusterProperties_STATUS_ARM](#ManagedClusterProperties_STATUS_ARM).

| Property            | Description                                                                                                                                                                           | Type                                                                                                            |
|---------------------|---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|-----------------------------------------------------------------------------------------------------------------|
| dnsServiceIP        | An IP address assigned to the Kubernetes DNS service. It must be within the Kubernetes service address range specified in serviceCidr.                                                | string                                                                                                          |
| dockerBridgeCidr    | A CIDR notation IP range assigned to the Docker bridge network. It must not overlap with any Subnet IP ranges or the Kubernetes service address range.                                | string                                                                                                          |
| loadBalancerProfile | Profile of the cluster load balancer.                                                                                                                                                 | [ManagedClusterLoadBalancerProfile_STATUS_ARM](#ManagedClusterLoadBalancerProfile_STATUS_ARM)                   |
| loadBalancerSku     | The default is 'standard'. See [Azure Load Balancer SKUs](https://docs.microsoft.com/azure/load-balancer/skus) for more information about the differences between load balancer SKUs. | [ContainerServiceNetworkProfile_LoadBalancerSku_STATUS](#ContainerServiceNetworkProfile_LoadBalancerSku_STATUS) |
| networkMode         | This cannot be specified if networkPlugin is anything other than 'azure'.                                                                                                             | [ContainerServiceNetworkProfile_NetworkMode_STATUS](#ContainerServiceNetworkProfile_NetworkMode_STATUS)         |
| networkPlugin       | Network plugin used for building the Kubernetes network.                                                                                                                              | [ContainerServiceNetworkProfile_NetworkPlugin_STATUS](#ContainerServiceNetworkProfile_NetworkPlugin_STATUS)     |
| networkPolicy       | Network policy used for building the Kubernetes network.                                                                                                                              | [ContainerServiceNetworkProfile_NetworkPolicy_STATUS](#ContainerServiceNetworkProfile_NetworkPolicy_STATUS)     |
| outboundType        | This can only be set at cluster creation time and cannot be changed later. For more information see [egress outbound type](https://docs.microsoft.com/azure/aks/egress-outboundtype). | [ContainerServiceNetworkProfile_OutboundType_STATUS](#ContainerServiceNetworkProfile_OutboundType_STATUS)       |
| podCidr             | A CIDR notation IP range from which to assign pod IPs when kubenet is used.                                                                                                           | string                                                                                                          |
| serviceCidr         | A CIDR notation IP range from which to assign service cluster IPs. It must not overlap with any Subnet IP ranges.                                                                     | string                                                                                                          |

<a id="ContainerServiceOSDisk"></a>ContainerServiceOSDisk
---------------------------------------------------------

Used by: [ManagedClusterAgentPoolProfile](#ManagedClusterAgentPoolProfile), [ManagedClusterAgentPoolProfileProperties_ARM](#ManagedClusterAgentPoolProfileProperties_ARM), [ManagedClusterAgentPoolProfile_ARM](#ManagedClusterAgentPoolProfile_ARM), and [ManagedClusters_AgentPool_Spec](#ManagedClusters_AgentPool_Spec).

<a id="ContainerServiceSshConfiguration"></a>ContainerServiceSshConfiguration
-----------------------------------------------------------------------------

SSH configuration for Linux-based VMs running on Azure.

Used by: [ContainerServiceLinuxProfile](#ContainerServiceLinuxProfile).

| Property   | Description                                                                                                 | Type                                                                         |
|------------|-------------------------------------------------------------------------------------------------------------|------------------------------------------------------------------------------|
| publicKeys | The list of SSH public keys used to authenticate with Linux-based VMs. A maximum of 1 key may be specified. | [ContainerServiceSshPublicKey[]](#ContainerServiceSshPublicKey)<br/>Required |

<a id="ContainerServiceSshConfiguration_ARM"></a>ContainerServiceSshConfiguration_ARM
-------------------------------------------------------------------------------------

SSH configuration for Linux-based VMs running on Azure.

Used by: [ContainerServiceLinuxProfile_ARM](#ContainerServiceLinuxProfile_ARM).

| Property   | Description                                                                                                 | Type                                                                    |
|------------|-------------------------------------------------------------------------------------------------------------|-------------------------------------------------------------------------|
| publicKeys | The list of SSH public keys used to authenticate with Linux-based VMs. A maximum of 1 key may be specified. | [ContainerServiceSshPublicKey_ARM[]](#ContainerServiceSshPublicKey_ARM) |

<a id="ContainerServiceSshConfiguration_STATUS"></a>ContainerServiceSshConfiguration_STATUS
-------------------------------------------------------------------------------------------

SSH configuration for Linux-based VMs running on Azure.

Used by: [ContainerServiceLinuxProfile_STATUS](#ContainerServiceLinuxProfile_STATUS).

| Property   | Description                                                                                                 | Type                                                                          |
|------------|-------------------------------------------------------------------------------------------------------------|-------------------------------------------------------------------------------|
| publicKeys | The list of SSH public keys used to authenticate with Linux-based VMs. A maximum of 1 key may be specified. | [ContainerServiceSshPublicKey_STATUS[]](#ContainerServiceSshPublicKey_STATUS) |

<a id="ContainerServiceSshConfiguration_STATUS_ARM"></a>ContainerServiceSshConfiguration_STATUS_ARM
---------------------------------------------------------------------------------------------------

SSH configuration for Linux-based VMs running on Azure.

Used by: [ContainerServiceLinuxProfile_STATUS_ARM](#ContainerServiceLinuxProfile_STATUS_ARM).

| Property   | Description                                                                                                 | Type                                                                                  |
|------------|-------------------------------------------------------------------------------------------------------------|---------------------------------------------------------------------------------------|
| publicKeys | The list of SSH public keys used to authenticate with Linux-based VMs. A maximum of 1 key may be specified. | [ContainerServiceSshPublicKey_STATUS_ARM[]](#ContainerServiceSshPublicKey_STATUS_ARM) |

<a id="ContainerServiceSshPublicKey"></a>ContainerServiceSshPublicKey
---------------------------------------------------------------------

Contains information about SSH certificate public key data.

Used by: [ContainerServiceSshConfiguration](#ContainerServiceSshConfiguration).

| Property | Description                                                                                                                      | Type                |
|----------|----------------------------------------------------------------------------------------------------------------------------------|---------------------|
| keyData  | Certificate public key used to authenticate with VMs through SSH. The certificate must be in PEM format with or without headers. | string<br/>Required |

<a id="ContainerServiceSshPublicKey_ARM"></a>ContainerServiceSshPublicKey_ARM
-----------------------------------------------------------------------------

Contains information about SSH certificate public key data.

Used by: [ContainerServiceSshConfiguration_ARM](#ContainerServiceSshConfiguration_ARM).

| Property | Description                                                                                                                      | Type   |
|----------|----------------------------------------------------------------------------------------------------------------------------------|--------|
| keyData  | Certificate public key used to authenticate with VMs through SSH. The certificate must be in PEM format with or without headers. | string |

<a id="ContainerServiceSshPublicKey_STATUS"></a>ContainerServiceSshPublicKey_STATUS
-----------------------------------------------------------------------------------

Contains information about SSH certificate public key data.

Used by: [ContainerServiceSshConfiguration_STATUS](#ContainerServiceSshConfiguration_STATUS).

| Property | Description                                                                                                                      | Type   |
|----------|----------------------------------------------------------------------------------------------------------------------------------|--------|
| keyData  | Certificate public key used to authenticate with VMs through SSH. The certificate must be in PEM format with or without headers. | string |

<a id="ContainerServiceSshPublicKey_STATUS_ARM"></a>ContainerServiceSshPublicKey_STATUS_ARM
-------------------------------------------------------------------------------------------

Contains information about SSH certificate public key data.

Used by: [ContainerServiceSshConfiguration_STATUS_ARM](#ContainerServiceSshConfiguration_STATUS_ARM).

| Property | Description                                                                                                                      | Type   |
|----------|----------------------------------------------------------------------------------------------------------------------------------|--------|
| keyData  | Certificate public key used to authenticate with VMs through SSH. The certificate must be in PEM format with or without headers. | string |

<a id="ExtendedLocation"></a>ExtendedLocation
---------------------------------------------

The complex type of the extended location.

Used by: [ManagedCluster_Spec](#ManagedCluster_Spec).

| Property | Description                        | Type                                          |
|----------|------------------------------------|-----------------------------------------------|
| name     | The name of the extended location. | string                                        |
| type     | The type of the extended location. | [ExtendedLocationType](#ExtendedLocationType) |

<a id="ExtendedLocationType"></a>ExtendedLocationType
-----------------------------------------------------

Used by: [ExtendedLocation](#ExtendedLocation), and [ExtendedLocation_ARM](#ExtendedLocation_ARM).

| Value      | Description |
|------------|-------------|
| "EdgeZone" |             |

<a id="ExtendedLocationType_STATUS"></a>ExtendedLocationType_STATUS
-------------------------------------------------------------------

Used by: [ExtendedLocation_STATUS](#ExtendedLocation_STATUS), and [ExtendedLocation_STATUS_ARM](#ExtendedLocation_STATUS_ARM).

| Value      | Description |
|------------|-------------|
| "EdgeZone" |             |

<a id="ExtendedLocation_ARM"></a>ExtendedLocation_ARM
-----------------------------------------------------

The complex type of the extended location.

Used by: [ManagedCluster_Spec_ARM](#ManagedCluster_Spec_ARM).

| Property | Description                        | Type                                          |
|----------|------------------------------------|-----------------------------------------------|
| name     | The name of the extended location. | string                                        |
| type     | The type of the extended location. | [ExtendedLocationType](#ExtendedLocationType) |

<a id="ExtendedLocation_STATUS"></a>ExtendedLocation_STATUS
-----------------------------------------------------------

The complex type of the extended location.

Used by: [ManagedCluster_STATUS](#ManagedCluster_STATUS).

| Property | Description                        | Type                                                        |
|----------|------------------------------------|-------------------------------------------------------------|
| name     | The name of the extended location. | string                                                      |
| type     | The type of the extended location. | [ExtendedLocationType_STATUS](#ExtendedLocationType_STATUS) |

<a id="ExtendedLocation_STATUS_ARM"></a>ExtendedLocation_STATUS_ARM
-------------------------------------------------------------------

The complex type of the extended location.

Used by: [ManagedCluster_STATUS_ARM](#ManagedCluster_STATUS_ARM).

| Property | Description                        | Type                                                        |
|----------|------------------------------------|-------------------------------------------------------------|
| name     | The name of the extended location. | string                                                      |
| type     | The type of the extended location. | [ExtendedLocationType_STATUS](#ExtendedLocationType_STATUS) |

<a id="GPUInstanceProfile"></a>GPUInstanceProfile
-------------------------------------------------

Used by: [ManagedClusterAgentPoolProfile](#ManagedClusterAgentPoolProfile), [ManagedClusterAgentPoolProfileProperties_ARM](#ManagedClusterAgentPoolProfileProperties_ARM), [ManagedClusterAgentPoolProfile_ARM](#ManagedClusterAgentPoolProfile_ARM), and [ManagedClusters_AgentPool_Spec](#ManagedClusters_AgentPool_Spec).

| Value   | Description |
|---------|-------------|
| "MIG1g" |             |
| "MIG2g" |             |
| "MIG3g" |             |
| "MIG4g" |             |
| "MIG7g" |             |

<a id="GPUInstanceProfile_STATUS"></a>GPUInstanceProfile_STATUS
---------------------------------------------------------------

Used by: [ManagedClusterAgentPoolProfileProperties_STATUS_ARM](#ManagedClusterAgentPoolProfileProperties_STATUS_ARM), [ManagedClusterAgentPoolProfile_STATUS](#ManagedClusterAgentPoolProfile_STATUS), [ManagedClusterAgentPoolProfile_STATUS_ARM](#ManagedClusterAgentPoolProfile_STATUS_ARM), and [ManagedClusters_AgentPool_STATUS](#ManagedClusters_AgentPool_STATUS).

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

<a id="KubeletConfig_ARM"></a>KubeletConfig_ARM
-----------------------------------------------

See [AKS custom node configuration](https://docs.microsoft.com/azure/aks/custom-node-configuration) for more details.

Used by: [ManagedClusterAgentPoolProfileProperties_ARM](#ManagedClusterAgentPoolProfileProperties_ARM), and [ManagedClusterAgentPoolProfile_ARM](#ManagedClusterAgentPoolProfile_ARM).

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

<a id="KubeletConfig_STATUS"></a>KubeletConfig_STATUS
-----------------------------------------------------

See [AKS custom node configuration](https://docs.microsoft.com/azure/aks/custom-node-configuration) for more details.

Used by: [ManagedClusterAgentPoolProfile_STATUS](#ManagedClusterAgentPoolProfile_STATUS), and [ManagedClusters_AgentPool_STATUS](#ManagedClusters_AgentPool_STATUS).

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

<a id="KubeletConfig_STATUS_ARM"></a>KubeletConfig_STATUS_ARM
-------------------------------------------------------------

See [AKS custom node configuration](https://docs.microsoft.com/azure/aks/custom-node-configuration) for more details.

Used by: [ManagedClusterAgentPoolProfileProperties_STATUS_ARM](#ManagedClusterAgentPoolProfileProperties_STATUS_ARM), and [ManagedClusterAgentPoolProfile_STATUS_ARM](#ManagedClusterAgentPoolProfile_STATUS_ARM).

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

<a id="KubeletDiskType"></a>KubeletDiskType
-------------------------------------------

Used by: [ManagedClusterAgentPoolProfile](#ManagedClusterAgentPoolProfile), [ManagedClusterAgentPoolProfileProperties_ARM](#ManagedClusterAgentPoolProfileProperties_ARM), [ManagedClusterAgentPoolProfile_ARM](#ManagedClusterAgentPoolProfile_ARM), and [ManagedClusters_AgentPool_Spec](#ManagedClusters_AgentPool_Spec).

| Value       | Description |
|-------------|-------------|
| "OS"        |             |
| "Temporary" |             |

<a id="KubeletDiskType_STATUS"></a>KubeletDiskType_STATUS
---------------------------------------------------------

Used by: [ManagedClusterAgentPoolProfileProperties_STATUS_ARM](#ManagedClusterAgentPoolProfileProperties_STATUS_ARM), [ManagedClusterAgentPoolProfile_STATUS](#ManagedClusterAgentPoolProfile_STATUS), [ManagedClusterAgentPoolProfile_STATUS_ARM](#ManagedClusterAgentPoolProfile_STATUS_ARM), and [ManagedClusters_AgentPool_STATUS](#ManagedClusters_AgentPool_STATUS).

| Value       | Description |
|-------------|-------------|
| "OS"        |             |
| "Temporary" |             |

<a id="LinuxOSConfig"></a>LinuxOSConfig
---------------------------------------

See [AKS custom node configuration](https://docs.microsoft.com/azure/aks/custom-node-configuration) for more details.

Used by: [ManagedClusterAgentPoolProfile](#ManagedClusterAgentPoolProfile), and [ManagedClusters_AgentPool_Spec](#ManagedClusters_AgentPool_Spec).

| Property                   | Description                                                                                                                                                                                                                                         | Type                          |
|----------------------------|-----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|-------------------------------|
| swapFileSizeMB             | The size in MB of a swap file that will be created on each node.                                                                                                                                                                                    | int                           |
| sysctls                    | Sysctl settings for Linux agent nodes.                                                                                                                                                                                                              | [SysctlConfig](#SysctlConfig) |
| transparentHugePageDefrag  | Valid values are 'always', 'defer', 'defer+madvise', 'madvise' and 'never'. The default is 'madvise'. For more information see [Transparent Hugepages](https://www.kernel.org/doc/html/latest/admin-guide/mm/transhuge.html#admin-guide-transhuge). | string                        |
| transparentHugePageEnabled | Valid values are 'always', 'madvise', and 'never'. The default is 'always'. For more information see [Transparent Hugepages](https://www.kernel.org/doc/html/latest/admin-guide/mm/transhuge.html#admin-guide-transhuge).                           | string                        |

<a id="LinuxOSConfig_ARM"></a>LinuxOSConfig_ARM
-----------------------------------------------

See [AKS custom node configuration](https://docs.microsoft.com/azure/aks/custom-node-configuration) for more details.

Used by: [ManagedClusterAgentPoolProfileProperties_ARM](#ManagedClusterAgentPoolProfileProperties_ARM), and [ManagedClusterAgentPoolProfile_ARM](#ManagedClusterAgentPoolProfile_ARM).

| Property                   | Description                                                                                                                                                                                                                                         | Type                                  |
|----------------------------|-----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|---------------------------------------|
| swapFileSizeMB             | The size in MB of a swap file that will be created on each node.                                                                                                                                                                                    | int                                   |
| sysctls                    | Sysctl settings for Linux agent nodes.                                                                                                                                                                                                              | [SysctlConfig_ARM](#SysctlConfig_ARM) |
| transparentHugePageDefrag  | Valid values are 'always', 'defer', 'defer+madvise', 'madvise' and 'never'. The default is 'madvise'. For more information see [Transparent Hugepages](https://www.kernel.org/doc/html/latest/admin-guide/mm/transhuge.html#admin-guide-transhuge). | string                                |
| transparentHugePageEnabled | Valid values are 'always', 'madvise', and 'never'. The default is 'always'. For more information see [Transparent Hugepages](https://www.kernel.org/doc/html/latest/admin-guide/mm/transhuge.html#admin-guide-transhuge).                           | string                                |

<a id="LinuxOSConfig_STATUS"></a>LinuxOSConfig_STATUS
-----------------------------------------------------

See [AKS custom node configuration](https://docs.microsoft.com/azure/aks/custom-node-configuration) for more details.

Used by: [ManagedClusterAgentPoolProfile_STATUS](#ManagedClusterAgentPoolProfile_STATUS), and [ManagedClusters_AgentPool_STATUS](#ManagedClusters_AgentPool_STATUS).

| Property                   | Description                                                                                                                                                                                                                                         | Type                                        |
|----------------------------|-----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|---------------------------------------------|
| swapFileSizeMB             | The size in MB of a swap file that will be created on each node.                                                                                                                                                                                    | int                                         |
| sysctls                    | Sysctl settings for Linux agent nodes.                                                                                                                                                                                                              | [SysctlConfig_STATUS](#SysctlConfig_STATUS) |
| transparentHugePageDefrag  | Valid values are 'always', 'defer', 'defer+madvise', 'madvise' and 'never'. The default is 'madvise'. For more information see [Transparent Hugepages](https://www.kernel.org/doc/html/latest/admin-guide/mm/transhuge.html#admin-guide-transhuge). | string                                      |
| transparentHugePageEnabled | Valid values are 'always', 'madvise', and 'never'. The default is 'always'. For more information see [Transparent Hugepages](https://www.kernel.org/doc/html/latest/admin-guide/mm/transhuge.html#admin-guide-transhuge).                           | string                                      |

<a id="LinuxOSConfig_STATUS_ARM"></a>LinuxOSConfig_STATUS_ARM
-------------------------------------------------------------

See [AKS custom node configuration](https://docs.microsoft.com/azure/aks/custom-node-configuration) for more details.

Used by: [ManagedClusterAgentPoolProfileProperties_STATUS_ARM](#ManagedClusterAgentPoolProfileProperties_STATUS_ARM), and [ManagedClusterAgentPoolProfile_STATUS_ARM](#ManagedClusterAgentPoolProfile_STATUS_ARM).

| Property                   | Description                                                                                                                                                                                                                                         | Type                                                |
|----------------------------|-----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|-----------------------------------------------------|
| swapFileSizeMB             | The size in MB of a swap file that will be created on each node.                                                                                                                                                                                    | int                                                 |
| sysctls                    | Sysctl settings for Linux agent nodes.                                                                                                                                                                                                              | [SysctlConfig_STATUS_ARM](#SysctlConfig_STATUS_ARM) |
| transparentHugePageDefrag  | Valid values are 'always', 'defer', 'defer+madvise', 'madvise' and 'never'. The default is 'madvise'. For more information see [Transparent Hugepages](https://www.kernel.org/doc/html/latest/admin-guide/mm/transhuge.html#admin-guide-transhuge). | string                                              |
| transparentHugePageEnabled | Valid values are 'always', 'madvise', and 'never'. The default is 'always'. For more information see [Transparent Hugepages](https://www.kernel.org/doc/html/latest/admin-guide/mm/transhuge.html#admin-guide-transhuge).                           | string                                              |

<a id="ManagedCluster"></a>ManagedCluster
-----------------------------------------

Generator information: - Generated from: /containerservice/resource-manager/Microsoft.ContainerService/aks/stable/2021-05-01/managedClusters.json - ARM URI: /&ZeroWidthSpace;subscriptions/&ZeroWidthSpace;{subscriptionId}/&ZeroWidthSpace;resourceGroups/&ZeroWidthSpace;{resourceGroupName}/&ZeroWidthSpace;providers/&ZeroWidthSpace;Microsoft.ContainerService/managedClusters/{resourceName}

Used by: [ManagedClusterList](#ManagedClusterList).

| Property          | Description | Type                                            |
|-------------------|-------------|-------------------------------------------------|
| metav1.TypeMeta   |             |                                                 |
| metav1.ObjectMeta |             |                                                 |
| spec              |             | [ManagedCluster_Spec](#ManagedCluster_Spec)     |
| status            |             | [ManagedCluster_STATUS](#ManagedCluster_STATUS) |

### <a id="ManagedCluster_Spec"></a>ManagedCluster_Spec

| Property                     | Description                                                                                                                                                                                                                                                                                                                                                                                 | Type                                                                                      |
|------------------------------|---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|-------------------------------------------------------------------------------------------|
| aadProfile                   | The Azure Active Directory configuration.                                                                                                                                                                                                                                                                                                                                                   | [ManagedClusterAADProfile](#ManagedClusterAADProfile)                                     |
| addonProfiles                | The profile of managed cluster add-on.                                                                                                                                                                                                                                                                                                                                                      | [map[string]ManagedClusterAddonProfile](#ManagedClusterAddonProfile)                      |
| agentPoolProfiles            | The agent pool properties.                                                                                                                                                                                                                                                                                                                                                                  | [ManagedClusterAgentPoolProfile[]](#ManagedClusterAgentPoolProfile)                       |
| apiServerAccessProfile       | The access profile for managed cluster API server.                                                                                                                                                                                                                                                                                                                                          | [ManagedClusterAPIServerAccessProfile](#ManagedClusterAPIServerAccessProfile)             |
| autoScalerProfile            | Parameters to be applied to the cluster-autoscaler when enabled                                                                                                                                                                                                                                                                                                                             | [ManagedClusterProperties_AutoScalerProfile](#ManagedClusterProperties_AutoScalerProfile) |
| autoUpgradeProfile           | The auto upgrade configuration.                                                                                                                                                                                                                                                                                                                                                             | [ManagedClusterAutoUpgradeProfile](#ManagedClusterAutoUpgradeProfile)                     |
| azureName                    | The name of the resource in Azure. This is often the same as the name of the resource in Kubernetes but it doesn't have to be.                                                                                                                                                                                                                                                              | string                                                                                    |
| disableLocalAccounts         | If set to true, getting static credentials will be disabled for this cluster. This must only be used on Managed Clusters that are AAD enabled. For more details see [disable local accounts](https://docs.microsoft.com/azure/aks/managed-aad#disable-local-accounts-preview).                                                                                                              | bool                                                                                      |
| diskEncryptionSetIDReference | This is of the form: '/&ZeroWidthSpace;subscriptions/&ZeroWidthSpace;{subscriptionId}/&ZeroWidthSpace;resourceGroups/&ZeroWidthSpace;{resourceGroupName}/&ZeroWidthSpace;providers/&ZeroWidthSpace;Microsoft.Compute/diskEncryptionSets/{encryptionSetName}'                                                                                                                                | genruntime.ResourceReference                                                              |
| dnsPrefix                    | This cannot be updated once the Managed Cluster has been created.                                                                                                                                                                                                                                                                                                                           | string                                                                                    |
| enablePodSecurityPolicy      | (DEPRECATING) Whether to enable Kubernetes pod security policy (preview). This feature is set for removal on October 15th, 2020. Learn more at aka.ms/aks/azpodpolicy.                                                                                                                                                                                                                      | bool                                                                                      |
| enableRBAC                   | Whether to enable Kubernetes Role-Based Access Control.                                                                                                                                                                                                                                                                                                                                     | bool                                                                                      |
| extendedLocation             | The extended location of the Virtual Machine.                                                                                                                                                                                                                                                                                                                                               | [ExtendedLocation](#ExtendedLocation)                                                     |
| fqdnSubdomain                | This cannot be updated once the Managed Cluster has been created.                                                                                                                                                                                                                                                                                                                           | string                                                                                    |
| httpProxyConfig              | Configurations for provisioning the cluster with HTTP proxy servers.                                                                                                                                                                                                                                                                                                                        | [ManagedClusterHTTPProxyConfig](#ManagedClusterHTTPProxyConfig)                           |
| identity                     | The identity of the managed cluster, if configured.                                                                                                                                                                                                                                                                                                                                         | [ManagedClusterIdentity](#ManagedClusterIdentity)                                         |
| identityProfile              | Identities associated with the cluster.                                                                                                                                                                                                                                                                                                                                                     | [map[string]UserAssignedIdentity](#UserAssignedIdentity)                                  |
| kubernetesVersion            | When you upgrade a supported AKS cluster, Kubernetes minor versions cannot be skipped. All upgrades must be performed sequentially by major version number. For example, upgrades between 1.14.x -> 1.15.x or 1.15.x -> 1.16.x are allowed, however 1.14.x -> 1.16.x is not allowed. See [upgrading an AKS cluster](https://docs.microsoft.com/azure/aks/upgrade-cluster) for more details. | string                                                                                    |
| linuxProfile                 | The profile for Linux VMs in the Managed Cluster.                                                                                                                                                                                                                                                                                                                                           | [ContainerServiceLinuxProfile](#ContainerServiceLinuxProfile)                             |
| location                     | Resource location                                                                                                                                                                                                                                                                                                                                                                           | string<br/>Required                                                                       |
| networkProfile               | The network configuration profile.                                                                                                                                                                                                                                                                                                                                                          | [ContainerServiceNetworkProfile](#ContainerServiceNetworkProfile)                         |
| nodeResourceGroup            | The name of the resource group containing agent pool nodes.                                                                                                                                                                                                                                                                                                                                 | string                                                                                    |
| operatorSpec                 | The specification for configuring operator behavior. This field is interpreted by the operator and not passed directly to Azure                                                                                                                                                                                                                                                             | [ManagedClusterOperatorSpec](#ManagedClusterOperatorSpec)                                 |
| owner                        | The owner of the resource. The owner controls where the resource goes when it is deployed. The owner also controls the resources lifecycle. When the owner is deleted the resource will also be deleted. Owner is expected to be a reference to a resources.azure.com/ResourceGroup resource                                                                                                | genruntime.KnownResourceReference<br/>Required                                            |
| podIdentityProfile           | See [use AAD pod identity](https://docs.microsoft.com/azure/aks/use-azure-ad-pod-identity) for more details on AAD pod identity integration.                                                                                                                                                                                                                                                | [ManagedClusterPodIdentityProfile](#ManagedClusterPodIdentityProfile)                     |
| privateLinkResources         | Private link resources associated with the cluster.                                                                                                                                                                                                                                                                                                                                         | [PrivateLinkResource[]](#PrivateLinkResource)                                             |
| servicePrincipalProfile      | Information about a service principal identity for the cluster to use for manipulating Azure APIs.                                                                                                                                                                                                                                                                                          | [ManagedClusterServicePrincipalProfile](#ManagedClusterServicePrincipalProfile)           |
| sku                          | The managed cluster SKU.                                                                                                                                                                                                                                                                                                                                                                    | [ManagedClusterSKU](#ManagedClusterSKU)                                                   |
| tags                         | Resource tags                                                                                                                                                                                                                                                                                                                                                                               | map[string]string                                                                         |
| windowsProfile               | The profile for Windows VMs in the Managed Cluster.                                                                                                                                                                                                                                                                                                                                         | [ManagedClusterWindowsProfile](#ManagedClusterWindowsProfile)                             |

### <a id="ManagedCluster_STATUS"></a>ManagedCluster_STATUS

| Property                | Description                                                                                                                                                                                                                                                                                                                                                                                 | Type                                                                                                    |
|-------------------------|---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|---------------------------------------------------------------------------------------------------------|
| aadProfile              | The Azure Active Directory configuration.                                                                                                                                                                                                                                                                                                                                                   | [ManagedClusterAADProfile_STATUS](#ManagedClusterAADProfile_STATUS)                                     |
| addonProfiles           | The profile of managed cluster add-on.                                                                                                                                                                                                                                                                                                                                                      | [map[string]ManagedClusterAddonProfile_STATUS](#ManagedClusterAddonProfile_STATUS)                      |
| agentPoolProfiles       | The agent pool properties.                                                                                                                                                                                                                                                                                                                                                                  | [ManagedClusterAgentPoolProfile_STATUS[]](#ManagedClusterAgentPoolProfile_STATUS)                       |
| apiServerAccessProfile  | The access profile for managed cluster API server.                                                                                                                                                                                                                                                                                                                                          | [ManagedClusterAPIServerAccessProfile_STATUS](#ManagedClusterAPIServerAccessProfile_STATUS)             |
| autoScalerProfile       | Parameters to be applied to the cluster-autoscaler when enabled                                                                                                                                                                                                                                                                                                                             | [ManagedClusterProperties_AutoScalerProfile_STATUS](#ManagedClusterProperties_AutoScalerProfile_STATUS) |
| autoUpgradeProfile      | The auto upgrade configuration.                                                                                                                                                                                                                                                                                                                                                             | [ManagedClusterAutoUpgradeProfile_STATUS](#ManagedClusterAutoUpgradeProfile_STATUS)                     |
| azurePortalFQDN         | The Azure Portal requires certain Cross-Origin Resource Sharing (CORS) headers to be sent in some responses, which Kubernetes APIServer doesn't handle by default. This special FQDN supports CORS, allowing the Azure Portal to function properly.                                                                                                                                         | string                                                                                                  |
| conditions              | The observed state of the resource                                                                                                                                                                                                                                                                                                                                                          | conditions.Condition[]                                                                                  |
| disableLocalAccounts    | If set to true, getting static credentials will be disabled for this cluster. This must only be used on Managed Clusters that are AAD enabled. For more details see [disable local accounts](https://docs.microsoft.com/azure/aks/managed-aad#disable-local-accounts-preview).                                                                                                              | bool                                                                                                    |
| diskEncryptionSetID     | This is of the form: '/&ZeroWidthSpace;subscriptions/&ZeroWidthSpace;{subscriptionId}/&ZeroWidthSpace;resourceGroups/&ZeroWidthSpace;{resourceGroupName}/&ZeroWidthSpace;providers/&ZeroWidthSpace;Microsoft.Compute/diskEncryptionSets/{encryptionSetName}'                                                                                                                                | string                                                                                                  |
| dnsPrefix               | This cannot be updated once the Managed Cluster has been created.                                                                                                                                                                                                                                                                                                                           | string                                                                                                  |
| enablePodSecurityPolicy | (DEPRECATING) Whether to enable Kubernetes pod security policy (preview). This feature is set for removal on October 15th, 2020. Learn more at aka.ms/aks/azpodpolicy.                                                                                                                                                                                                                      | bool                                                                                                    |
| enableRBAC              | Whether to enable Kubernetes Role-Based Access Control.                                                                                                                                                                                                                                                                                                                                     | bool                                                                                                    |
| extendedLocation        | The extended location of the Virtual Machine.                                                                                                                                                                                                                                                                                                                                               | [ExtendedLocation_STATUS](#ExtendedLocation_STATUS)                                                     |
| fqdn                    | The FQDN of the master pool.                                                                                                                                                                                                                                                                                                                                                                | string                                                                                                  |
| fqdnSubdomain           | This cannot be updated once the Managed Cluster has been created.                                                                                                                                                                                                                                                                                                                           | string                                                                                                  |
| httpProxyConfig         | Configurations for provisioning the cluster with HTTP proxy servers.                                                                                                                                                                                                                                                                                                                        | [ManagedClusterHTTPProxyConfig_STATUS](#ManagedClusterHTTPProxyConfig_STATUS)                           |
| id                      | Resource Id                                                                                                                                                                                                                                                                                                                                                                                 | string                                                                                                  |
| identity                | The identity of the managed cluster, if configured.                                                                                                                                                                                                                                                                                                                                         | [ManagedClusterIdentity_STATUS](#ManagedClusterIdentity_STATUS)                                         |
| identityProfile         | Identities associated with the cluster.                                                                                                                                                                                                                                                                                                                                                     | [map[string]UserAssignedIdentity_STATUS](#UserAssignedIdentity_STATUS)                                  |
| kubernetesVersion       | When you upgrade a supported AKS cluster, Kubernetes minor versions cannot be skipped. All upgrades must be performed sequentially by major version number. For example, upgrades between 1.14.x -> 1.15.x or 1.15.x -> 1.16.x are allowed, however 1.14.x -> 1.16.x is not allowed. See [upgrading an AKS cluster](https://docs.microsoft.com/azure/aks/upgrade-cluster) for more details. | string                                                                                                  |
| linuxProfile            | The profile for Linux VMs in the Managed Cluster.                                                                                                                                                                                                                                                                                                                                           | [ContainerServiceLinuxProfile_STATUS](#ContainerServiceLinuxProfile_STATUS)                             |
| location                | Resource location                                                                                                                                                                                                                                                                                                                                                                           | string                                                                                                  |
| maxAgentPools           | The max number of agent pools for the managed cluster.                                                                                                                                                                                                                                                                                                                                      | int                                                                                                     |
| name                    | Resource name                                                                                                                                                                                                                                                                                                                                                                               | string                                                                                                  |
| networkProfile          | The network configuration profile.                                                                                                                                                                                                                                                                                                                                                          | [ContainerServiceNetworkProfile_STATUS](#ContainerServiceNetworkProfile_STATUS)                         |
| nodeResourceGroup       | The name of the resource group containing agent pool nodes.                                                                                                                                                                                                                                                                                                                                 | string                                                                                                  |
| podIdentityProfile      | See [use AAD pod identity](https://docs.microsoft.com/azure/aks/use-azure-ad-pod-identity) for more details on AAD pod identity integration.                                                                                                                                                                                                                                                | [ManagedClusterPodIdentityProfile_STATUS](#ManagedClusterPodIdentityProfile_STATUS)                     |
| powerState              | The Power State of the cluster.                                                                                                                                                                                                                                                                                                                                                             | [PowerState_STATUS](#PowerState_STATUS)                                                                 |
| privateFQDN             | The FQDN of private cluster.                                                                                                                                                                                                                                                                                                                                                                | string                                                                                                  |
| privateLinkResources    | Private link resources associated with the cluster.                                                                                                                                                                                                                                                                                                                                         | [PrivateLinkResource_STATUS[]](#PrivateLinkResource_STATUS)                                             |
| provisioningState       | The current provisioning state.                                                                                                                                                                                                                                                                                                                                                             | string                                                                                                  |
| servicePrincipalProfile | Information about a service principal identity for the cluster to use for manipulating Azure APIs.                                                                                                                                                                                                                                                                                          | [ManagedClusterServicePrincipalProfile_STATUS](#ManagedClusterServicePrincipalProfile_STATUS)           |
| sku                     | The managed cluster SKU.                                                                                                                                                                                                                                                                                                                                                                    | [ManagedClusterSKU_STATUS](#ManagedClusterSKU_STATUS)                                                   |
| tags                    | Resource tags                                                                                                                                                                                                                                                                                                                                                                               | map[string]string                                                                                       |
| type                    | Resource type                                                                                                                                                                                                                                                                                                                                                                               | string                                                                                                  |
| windowsProfile          | The profile for Windows VMs in the Managed Cluster.                                                                                                                                                                                                                                                                                                                                         | [ManagedClusterWindowsProfile_STATUS](#ManagedClusterWindowsProfile_STATUS)                             |

<a id="ManagedClusterAADProfile"></a>ManagedClusterAADProfile
-------------------------------------------------------------

For more details see [managed AAD on AKS](https://docs.microsoft.com/azure/aks/managed-aad).

Used by: [ManagedCluster_Spec](#ManagedCluster_Spec).

| Property            | Description                                                                                                        | Type     |
|---------------------|--------------------------------------------------------------------------------------------------------------------|----------|
| adminGroupObjectIDs | The list of AAD group object IDs that will have admin role of the cluster.                                         | string[] |
| clientAppID         | The client AAD application ID.                                                                                     | string   |
| enableAzureRBAC     | Whether to enable Azure RBAC for Kubernetes authorization.                                                         | bool     |
| managed             | Whether to enable managed AAD.                                                                                     | bool     |
| serverAppID         | The server AAD application ID.                                                                                     | string   |
| serverAppSecret     | The server AAD application secret.                                                                                 | string   |
| tenantID            | The AAD tenant ID to use for authentication. If not specified, will use the tenant of the deployment subscription. | string   |

<a id="ManagedClusterAADProfile_ARM"></a>ManagedClusterAADProfile_ARM
---------------------------------------------------------------------

For more details see [managed AAD on AKS](https://docs.microsoft.com/azure/aks/managed-aad).

Used by: [ManagedClusterProperties_ARM](#ManagedClusterProperties_ARM).

| Property            | Description                                                                                                        | Type     |
|---------------------|--------------------------------------------------------------------------------------------------------------------|----------|
| adminGroupObjectIDs | The list of AAD group object IDs that will have admin role of the cluster.                                         | string[] |
| clientAppID         | The client AAD application ID.                                                                                     | string   |
| enableAzureRBAC     | Whether to enable Azure RBAC for Kubernetes authorization.                                                         | bool     |
| managed             | Whether to enable managed AAD.                                                                                     | bool     |
| serverAppID         | The server AAD application ID.                                                                                     | string   |
| serverAppSecret     | The server AAD application secret.                                                                                 | string   |
| tenantID            | The AAD tenant ID to use for authentication. If not specified, will use the tenant of the deployment subscription. | string   |

<a id="ManagedClusterAADProfile_STATUS"></a>ManagedClusterAADProfile_STATUS
---------------------------------------------------------------------------

For more details see [managed AAD on AKS](https://docs.microsoft.com/azure/aks/managed-aad).

Used by: [ManagedCluster_STATUS](#ManagedCluster_STATUS).

| Property            | Description                                                                                                        | Type     |
|---------------------|--------------------------------------------------------------------------------------------------------------------|----------|
| adminGroupObjectIDs | The list of AAD group object IDs that will have admin role of the cluster.                                         | string[] |
| clientAppID         | The client AAD application ID.                                                                                     | string   |
| enableAzureRBAC     | Whether to enable Azure RBAC for Kubernetes authorization.                                                         | bool     |
| managed             | Whether to enable managed AAD.                                                                                     | bool     |
| serverAppID         | The server AAD application ID.                                                                                     | string   |
| serverAppSecret     | The server AAD application secret.                                                                                 | string   |
| tenantID            | The AAD tenant ID to use for authentication. If not specified, will use the tenant of the deployment subscription. | string   |

<a id="ManagedClusterAADProfile_STATUS_ARM"></a>ManagedClusterAADProfile_STATUS_ARM
-----------------------------------------------------------------------------------

For more details see [managed AAD on AKS](https://docs.microsoft.com/azure/aks/managed-aad).

Used by: [ManagedClusterProperties_STATUS_ARM](#ManagedClusterProperties_STATUS_ARM).

| Property            | Description                                                                                                        | Type     |
|---------------------|--------------------------------------------------------------------------------------------------------------------|----------|
| adminGroupObjectIDs | The list of AAD group object IDs that will have admin role of the cluster.                                         | string[] |
| clientAppID         | The client AAD application ID.                                                                                     | string   |
| enableAzureRBAC     | Whether to enable Azure RBAC for Kubernetes authorization.                                                         | bool     |
| managed             | Whether to enable managed AAD.                                                                                     | bool     |
| serverAppID         | The server AAD application ID.                                                                                     | string   |
| serverAppSecret     | The server AAD application secret.                                                                                 | string   |
| tenantID            | The AAD tenant ID to use for authentication. If not specified, will use the tenant of the deployment subscription. | string   |

<a id="ManagedClusterAPIServerAccessProfile"></a>ManagedClusterAPIServerAccessProfile
-------------------------------------------------------------------------------------

Access profile for managed cluster API server.

Used by: [ManagedCluster_Spec](#ManagedCluster_Spec).

| Property                       | Description                                                                                                                                                                                                                                                                                                                   | Type     |
|--------------------------------|-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|----------|
| authorizedIPRanges             | IP ranges are specified in CIDR format, e.g. 137.117.106.88/29. This feature is not compatible with clusters that use Public IP Per Node, or clusters that are using a Basic Load Balancer. For more information see [API server authorized IP ranges](https://docs.microsoft.com/azure/aks/api-server-authorized-ip-ranges). | string[] |
| enablePrivateCluster           | For more details, see [Creating a private AKS cluster](https://docs.microsoft.com/azure/aks/private-clusters).                                                                                                                                                                                                                | bool     |
| enablePrivateClusterPublicFQDN | Whether to create additional public FQDN for private cluster or not.                                                                                                                                                                                                                                                          | bool     |
| privateDNSZone                 | The default is System. For more details see [configure private DNS zone](https://docs.microsoft.com/azure/aks/private-clusters#configure-private-dns-zone). Allowed values are 'system' and 'none'.                                                                                                                           | string   |

<a id="ManagedClusterAPIServerAccessProfile_ARM"></a>ManagedClusterAPIServerAccessProfile_ARM
---------------------------------------------------------------------------------------------

Access profile for managed cluster API server.

Used by: [ManagedClusterProperties_ARM](#ManagedClusterProperties_ARM).

| Property                       | Description                                                                                                                                                                                                                                                                                                                   | Type     |
|--------------------------------|-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|----------|
| authorizedIPRanges             | IP ranges are specified in CIDR format, e.g. 137.117.106.88/29. This feature is not compatible with clusters that use Public IP Per Node, or clusters that are using a Basic Load Balancer. For more information see [API server authorized IP ranges](https://docs.microsoft.com/azure/aks/api-server-authorized-ip-ranges). | string[] |
| enablePrivateCluster           | For more details, see [Creating a private AKS cluster](https://docs.microsoft.com/azure/aks/private-clusters).                                                                                                                                                                                                                | bool     |
| enablePrivateClusterPublicFQDN | Whether to create additional public FQDN for private cluster or not.                                                                                                                                                                                                                                                          | bool     |
| privateDNSZone                 | The default is System. For more details see [configure private DNS zone](https://docs.microsoft.com/azure/aks/private-clusters#configure-private-dns-zone). Allowed values are 'system' and 'none'.                                                                                                                           | string   |

<a id="ManagedClusterAPIServerAccessProfile_STATUS"></a>ManagedClusterAPIServerAccessProfile_STATUS
---------------------------------------------------------------------------------------------------

Access profile for managed cluster API server.

Used by: [ManagedCluster_STATUS](#ManagedCluster_STATUS).

| Property                       | Description                                                                                                                                                                                                                                                                                                                   | Type     |
|--------------------------------|-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|----------|
| authorizedIPRanges             | IP ranges are specified in CIDR format, e.g. 137.117.106.88/29. This feature is not compatible with clusters that use Public IP Per Node, or clusters that are using a Basic Load Balancer. For more information see [API server authorized IP ranges](https://docs.microsoft.com/azure/aks/api-server-authorized-ip-ranges). | string[] |
| enablePrivateCluster           | For more details, see [Creating a private AKS cluster](https://docs.microsoft.com/azure/aks/private-clusters).                                                                                                                                                                                                                | bool     |
| enablePrivateClusterPublicFQDN | Whether to create additional public FQDN for private cluster or not.                                                                                                                                                                                                                                                          | bool     |
| privateDNSZone                 | The default is System. For more details see [configure private DNS zone](https://docs.microsoft.com/azure/aks/private-clusters#configure-private-dns-zone). Allowed values are 'system' and 'none'.                                                                                                                           | string   |

<a id="ManagedClusterAPIServerAccessProfile_STATUS_ARM"></a>ManagedClusterAPIServerAccessProfile_STATUS_ARM
-----------------------------------------------------------------------------------------------------------

Access profile for managed cluster API server.

Used by: [ManagedClusterProperties_STATUS_ARM](#ManagedClusterProperties_STATUS_ARM).

| Property                       | Description                                                                                                                                                                                                                                                                                                                   | Type     |
|--------------------------------|-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|----------|
| authorizedIPRanges             | IP ranges are specified in CIDR format, e.g. 137.117.106.88/29. This feature is not compatible with clusters that use Public IP Per Node, or clusters that are using a Basic Load Balancer. For more information see [API server authorized IP ranges](https://docs.microsoft.com/azure/aks/api-server-authorized-ip-ranges). | string[] |
| enablePrivateCluster           | For more details, see [Creating a private AKS cluster](https://docs.microsoft.com/azure/aks/private-clusters).                                                                                                                                                                                                                | bool     |
| enablePrivateClusterPublicFQDN | Whether to create additional public FQDN for private cluster or not.                                                                                                                                                                                                                                                          | bool     |
| privateDNSZone                 | The default is System. For more details see [configure private DNS zone](https://docs.microsoft.com/azure/aks/private-clusters#configure-private-dns-zone). Allowed values are 'system' and 'none'.                                                                                                                           | string   |

<a id="ManagedClusterAddonProfile"></a>ManagedClusterAddonProfile
-----------------------------------------------------------------

A Kubernetes add-on profile for a managed cluster.

Used by: [ManagedCluster_Spec](#ManagedCluster_Spec).

| Property | Description                                | Type              |
|----------|--------------------------------------------|-------------------|
| config   | Key-value pairs for configuring an add-on. | map[string]string |
| enabled  | Whether the add-on is enabled or not.      | bool<br/>Required |

<a id="ManagedClusterAddonProfile_ARM"></a>ManagedClusterAddonProfile_ARM
-------------------------------------------------------------------------

A Kubernetes add-on profile for a managed cluster.

Used by: [ManagedClusterProperties_ARM](#ManagedClusterProperties_ARM).

| Property | Description                                | Type              |
|----------|--------------------------------------------|-------------------|
| config   | Key-value pairs for configuring an add-on. | map[string]string |
| enabled  | Whether the add-on is enabled or not.      | bool              |

<a id="ManagedClusterAddonProfile_STATUS"></a>ManagedClusterAddonProfile_STATUS
-------------------------------------------------------------------------------

A Kubernetes add-on profile for a managed cluster.

Used by: [ManagedCluster_STATUS](#ManagedCluster_STATUS).

| Property | Description                                                | Type                                                        |
|----------|------------------------------------------------------------|-------------------------------------------------------------|
| config   | Key-value pairs for configuring an add-on.                 | map[string]string                                           |
| enabled  | Whether the add-on is enabled or not.                      | bool                                                        |
| identity | Information of user assigned identity used by this add-on. | [UserAssignedIdentity_STATUS](#UserAssignedIdentity_STATUS) |

<a id="ManagedClusterAddonProfile_STATUS_ARM"></a>ManagedClusterAddonProfile_STATUS_ARM
---------------------------------------------------------------------------------------

A Kubernetes add-on profile for a managed cluster.

Used by: [ManagedClusterProperties_STATUS_ARM](#ManagedClusterProperties_STATUS_ARM).

| Property | Description                                                | Type                                                                |
|----------|------------------------------------------------------------|---------------------------------------------------------------------|
| config   | Key-value pairs for configuring an add-on.                 | map[string]string                                                   |
| enabled  | Whether the add-on is enabled or not.                      | bool                                                                |
| identity | Information of user assigned identity used by this add-on. | [UserAssignedIdentity_STATUS_ARM](#UserAssignedIdentity_STATUS_ARM) |

<a id="ManagedClusterAgentPoolProfile"></a>ManagedClusterAgentPoolProfile
-------------------------------------------------------------------------

Profile for the container service agent pool.

Used by: [ManagedCluster_Spec](#ManagedCluster_Spec).

| Property                      | Description                                                                                                                                                                                                                                                                                                                                                                                                                                                                                  | Type                                                  |
|-------------------------------|----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|-------------------------------------------------------|
| availabilityZones             | The list of Availability zones to use for nodes. This can only be specified if the AgentPoolType property is 'VirtualMachineScaleSets'.                                                                                                                                                                                                                                                                                                                                                      | string[]                                              |
| count                         | Number of agents (VMs) to host docker containers. Allowed values must be in the range of 0 to 1000 (inclusive) for user pools and in the range of 1 to 1000 (inclusive) for system pools. The default value is 1.                                                                                                                                                                                                                                                                            | int                                                   |
| enableAutoScaling             | Whether to enable auto-scaler                                                                                                                                                                                                                                                                                                                                                                                                                                                                | bool                                                  |
| enableEncryptionAtHost        | This is only supported on certain VM sizes and in certain Azure regions. For more information, see: https://docs.microsoft.com/azure/aks/enable-host-encryption                                                                                                                                                                                                                                                                                                                              | bool                                                  |
| enableFIPS                    | See [Add a FIPS-enabled node pool](https://docs.microsoft.com/azure/aks/use-multiple-node-pools#add-a-fips-enabled-node-pool-preview) for more details.                                                                                                                                                                                                                                                                                                                                      | bool                                                  |
| enableNodePublicIP            | Some scenarios may require nodes in a node pool to receive their own dedicated public IP addresses. A common scenario is for gaming workloads, where a console needs to make a direct connection to a cloud virtual machine to minimize hops. For more information see [assigning a public IP per node](https://docs.microsoft.com/azure/aks/use-multiple-node-pools#assign-a-public-ip-per-node-for-your-node-pools). The default is false.                                                 | bool                                                  |
| enableUltraSSD                | Whether to enable UltraSSD                                                                                                                                                                                                                                                                                                                                                                                                                                                                   | bool                                                  |
| gpuInstanceProfile            | GPUInstanceProfile to be used to specify GPU MIG instance profile for supported GPU VM SKU.                                                                                                                                                                                                                                                                                                                                                                                                  | [GPUInstanceProfile](#GPUInstanceProfile)             |
| kubeletConfig                 | The Kubelet configuration on the agent pool nodes.                                                                                                                                                                                                                                                                                                                                                                                                                                           | [KubeletConfig](#KubeletConfig)                       |
| kubeletDiskType               | Determines the placement of emptyDir volumes, container runtime data root, and Kubelet ephemeral storage.                                                                                                                                                                                                                                                                                                                                                                                    | [KubeletDiskType](#KubeletDiskType)                   |
| linuxOSConfig                 | The OS configuration of Linux agent nodes.                                                                                                                                                                                                                                                                                                                                                                                                                                                   | [LinuxOSConfig](#LinuxOSConfig)                       |
| maxCount                      | The maximum number of nodes for auto-scaling                                                                                                                                                                                                                                                                                                                                                                                                                                                 | int                                                   |
| maxPods                       | The maximum number of pods that can run on a node.                                                                                                                                                                                                                                                                                                                                                                                                                                           | int                                                   |
| minCount                      | The minimum number of nodes for auto-scaling                                                                                                                                                                                                                                                                                                                                                                                                                                                 | int                                                   |
| mode                          | A cluster must have at least one 'System' Agent Pool at all times. For additional information on agent pool restrictions and best practices, see: https://docs.microsoft.com/azure/aks/use-system-pools                                                                                                                                                                                                                                                                                      | [AgentPoolMode](#AgentPoolMode)                       |
| name                          | Windows agent pool names must be 6 characters or less.                                                                                                                                                                                                                                                                                                                                                                                                                                       | string                                                |
| nodeLabels                    | The node labels to be persisted across all nodes in agent pool.                                                                                                                                                                                                                                                                                                                                                                                                                              | map[string]string                                     |
| nodePublicIPPrefixIDReference | This is of the form: /&ZeroWidthSpace;subscriptions/&ZeroWidthSpace;{subscriptionId}/&ZeroWidthSpace;resourceGroups/&ZeroWidthSpace;{resourceGroupName}/&ZeroWidthSpace;providers/&ZeroWidthSpace;Microsoft.Network/publicIPPrefixes/{publicIPPrefixName}                                                                                                                                                                                                                                    | genruntime.ResourceReference                          |
| nodeTaints                    | The taints added to new nodes during node pool create and scale. For example, key=value:NoSchedule.                                                                                                                                                                                                                                                                                                                                                                                          | string[]                                              |
| orchestratorVersion           | As a best practice, you should upgrade all node pools in an AKS cluster to the same Kubernetes version. The node pool version must have the same major version as the control plane. The node pool minor version must be within two minor versions of the control plane version. The node pool version cannot be greater than the control plane version. For more information see [upgrading a node pool](https://docs.microsoft.com/azure/aks/use-multiple-node-pools#upgrade-a-node-pool). | string                                                |
| osDiskSizeGB                  |                                                                                                                                                                                                                                                                                                                                                                                                                                                                                              | [ContainerServiceOSDisk](#ContainerServiceOSDisk)     |
| osDiskType                    | The default is 'Ephemeral' if the VM supports it and has a cache disk larger than the requested OSDiskSizeGB. Otherwise, defaults to 'Managed'. May not be changed after creation. For more information see [Ephemeral OS](https://docs.microsoft.com/azure/aks/cluster-configuration#ephemeral-os).                                                                                                                                                                                         | [OSDiskType](#OSDiskType)                             |
| osSKU                         | Specifies an OS SKU. This value must not be specified if OSType is Windows.                                                                                                                                                                                                                                                                                                                                                                                                                  | [OSSKU](#OSSKU)                                       |
| osType                        | The operating system type. The default is Linux.                                                                                                                                                                                                                                                                                                                                                                                                                                             | [OSType](#OSType)                                     |
| podSubnetIDReference          | If omitted, pod IPs are statically assigned on the node subnet (see vnetSubnetID for more details). This is of the form: /&ZeroWidthSpace;subscriptions/&ZeroWidthSpace;{subscriptionId}/&ZeroWidthSpace;resourceGroups/&ZeroWidthSpace;{resourceGroupName}/&ZeroWidthSpace;providers/&ZeroWidthSpace;Microsoft.Network/virtualNetworks/{virtualNetworkName}/subnets/{subnetName}                                                                                                            | genruntime.ResourceReference                          |
| proximityPlacementGroupID     | The ID for Proximity Placement Group.                                                                                                                                                                                                                                                                                                                                                                                                                                                        | string                                                |
| scaleSetEvictionPolicy        | This cannot be specified unless the scaleSetPriority is 'Spot'. If not specified, the default is 'Delete'.                                                                                                                                                                                                                                                                                                                                                                                   | [ScaleSetEvictionPolicy](#ScaleSetEvictionPolicy)     |
| scaleSetPriority              | The Virtual Machine Scale Set priority. If not specified, the default is 'Regular'.                                                                                                                                                                                                                                                                                                                                                                                                          | [ScaleSetPriority](#ScaleSetPriority)                 |
| spotMaxPrice                  | Possible values are any decimal value greater than zero or -1 which indicates the willingness to pay any on-demand price. For more details on spot pricing, see [spot VMs pricing](https://docs.microsoft.com/azure/virtual-machines/spot-vms#pricing)                                                                                                                                                                                                                                       | float64                                               |
| tags                          | The tags to be persisted on the agent pool virtual machine scale set.                                                                                                                                                                                                                                                                                                                                                                                                                        | map[string]string                                     |
| type                          | The type of Agent Pool.                                                                                                                                                                                                                                                                                                                                                                                                                                                                      | [AgentPoolType](#AgentPoolType)                       |
| upgradeSettings               | Settings for upgrading the agentpool                                                                                                                                                                                                                                                                                                                                                                                                                                                         | [AgentPoolUpgradeSettings](#AgentPoolUpgradeSettings) |
| vmSize                        | VM size availability varies by region. If a node contains insufficient compute resources (memory, cpu, etc) pods might fail to run correctly. For more details on restricted VM sizes, see: https://docs.microsoft.com/azure/aks/quotas-skus-regions                                                                                                                                                                                                                                         | string                                                |
| vnetSubnetIDReference         | If this is not specified, a VNET and subnet will be generated and used. If no podSubnetID is specified, this applies to nodes and pods, otherwise it applies to just nodes. This is of the form: /&ZeroWidthSpace;subscriptions/&ZeroWidthSpace;{subscriptionId}/&ZeroWidthSpace;resourceGroups/&ZeroWidthSpace;{resourceGroupName}/&ZeroWidthSpace;providers/&ZeroWidthSpace;Microsoft.Network/virtualNetworks/{virtualNetworkName}/subnets/{subnetName}                                    | genruntime.ResourceReference                          |

<a id="ManagedClusterAgentPoolProfileProperties_ARM"></a>ManagedClusterAgentPoolProfileProperties_ARM
-----------------------------------------------------------------------------------------------------

Properties for the container service agent pool profile.

Used by: [ManagedClusters_AgentPool_Spec_ARM](#ManagedClusters_AgentPool_Spec_ARM).

| Property                  | Description                                                                                                                                                                                                                                                                                                                                                                                                                                                                                  | Type                                                          |
|---------------------------|----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|---------------------------------------------------------------|
| availabilityZones         | The list of Availability zones to use for nodes. This can only be specified if the AgentPoolType property is 'VirtualMachineScaleSets'.                                                                                                                                                                                                                                                                                                                                                      | string[]                                                      |
| count                     | Number of agents (VMs) to host docker containers. Allowed values must be in the range of 0 to 1000 (inclusive) for user pools and in the range of 1 to 1000 (inclusive) for system pools. The default value is 1.                                                                                                                                                                                                                                                                            | int                                                           |
| enableAutoScaling         | Whether to enable auto-scaler                                                                                                                                                                                                                                                                                                                                                                                                                                                                | bool                                                          |
| enableEncryptionAtHost    | This is only supported on certain VM sizes and in certain Azure regions. For more information, see: https://docs.microsoft.com/azure/aks/enable-host-encryption                                                                                                                                                                                                                                                                                                                              | bool                                                          |
| enableFIPS                | See [Add a FIPS-enabled node pool](https://docs.microsoft.com/azure/aks/use-multiple-node-pools#add-a-fips-enabled-node-pool-preview) for more details.                                                                                                                                                                                                                                                                                                                                      | bool                                                          |
| enableNodePublicIP        | Some scenarios may require nodes in a node pool to receive their own dedicated public IP addresses. A common scenario is for gaming workloads, where a console needs to make a direct connection to a cloud virtual machine to minimize hops. For more information see [assigning a public IP per node](https://docs.microsoft.com/azure/aks/use-multiple-node-pools#assign-a-public-ip-per-node-for-your-node-pools). The default is false.                                                 | bool                                                          |
| enableUltraSSD            | Whether to enable UltraSSD                                                                                                                                                                                                                                                                                                                                                                                                                                                                   | bool                                                          |
| gpuInstanceProfile        | GPUInstanceProfile to be used to specify GPU MIG instance profile for supported GPU VM SKU.                                                                                                                                                                                                                                                                                                                                                                                                  | [GPUInstanceProfile](#GPUInstanceProfile)                     |
| kubeletConfig             | The Kubelet configuration on the agent pool nodes.                                                                                                                                                                                                                                                                                                                                                                                                                                           | [KubeletConfig_ARM](#KubeletConfig_ARM)                       |
| kubeletDiskType           | Determines the placement of emptyDir volumes, container runtime data root, and Kubelet ephemeral storage.                                                                                                                                                                                                                                                                                                                                                                                    | [KubeletDiskType](#KubeletDiskType)                           |
| linuxOSConfig             | The OS configuration of Linux agent nodes.                                                                                                                                                                                                                                                                                                                                                                                                                                                   | [LinuxOSConfig_ARM](#LinuxOSConfig_ARM)                       |
| maxCount                  | The maximum number of nodes for auto-scaling                                                                                                                                                                                                                                                                                                                                                                                                                                                 | int                                                           |
| maxPods                   | The maximum number of pods that can run on a node.                                                                                                                                                                                                                                                                                                                                                                                                                                           | int                                                           |
| minCount                  | The minimum number of nodes for auto-scaling                                                                                                                                                                                                                                                                                                                                                                                                                                                 | int                                                           |
| mode                      | A cluster must have at least one 'System' Agent Pool at all times. For additional information on agent pool restrictions and best practices, see: https://docs.microsoft.com/azure/aks/use-system-pools                                                                                                                                                                                                                                                                                      | [AgentPoolMode](#AgentPoolMode)                               |
| nodeLabels                | The node labels to be persisted across all nodes in agent pool.                                                                                                                                                                                                                                                                                                                                                                                                                              | map[string]string                                             |
| nodePublicIPPrefixID      |                                                                                                                                                                                                                                                                                                                                                                                                                                                                                              | string                                                        |
| nodeTaints                | The taints added to new nodes during node pool create and scale. For example, key=value:NoSchedule.                                                                                                                                                                                                                                                                                                                                                                                          | string[]                                                      |
| orchestratorVersion       | As a best practice, you should upgrade all node pools in an AKS cluster to the same Kubernetes version. The node pool version must have the same major version as the control plane. The node pool minor version must be within two minor versions of the control plane version. The node pool version cannot be greater than the control plane version. For more information see [upgrading a node pool](https://docs.microsoft.com/azure/aks/use-multiple-node-pools#upgrade-a-node-pool). | string                                                        |
| osDiskSizeGB              |                                                                                                                                                                                                                                                                                                                                                                                                                                                                                              | [ContainerServiceOSDisk](#ContainerServiceOSDisk)             |
| osDiskType                | The default is 'Ephemeral' if the VM supports it and has a cache disk larger than the requested OSDiskSizeGB. Otherwise, defaults to 'Managed'. May not be changed after creation. For more information see [Ephemeral OS](https://docs.microsoft.com/azure/aks/cluster-configuration#ephemeral-os).                                                                                                                                                                                         | [OSDiskType](#OSDiskType)                                     |
| osSKU                     | Specifies an OS SKU. This value must not be specified if OSType is Windows.                                                                                                                                                                                                                                                                                                                                                                                                                  | [OSSKU](#OSSKU)                                               |
| osType                    | The operating system type. The default is Linux.                                                                                                                                                                                                                                                                                                                                                                                                                                             | [OSType](#OSType)                                             |
| podSubnetID               |                                                                                                                                                                                                                                                                                                                                                                                                                                                                                              | string                                                        |
| proximityPlacementGroupID | The ID for Proximity Placement Group.                                                                                                                                                                                                                                                                                                                                                                                                                                                        | string                                                        |
| scaleSetEvictionPolicy    | This cannot be specified unless the scaleSetPriority is 'Spot'. If not specified, the default is 'Delete'.                                                                                                                                                                                                                                                                                                                                                                                   | [ScaleSetEvictionPolicy](#ScaleSetEvictionPolicy)             |
| scaleSetPriority          | The Virtual Machine Scale Set priority. If not specified, the default is 'Regular'.                                                                                                                                                                                                                                                                                                                                                                                                          | [ScaleSetPriority](#ScaleSetPriority)                         |
| spotMaxPrice              | Possible values are any decimal value greater than zero or -1 which indicates the willingness to pay any on-demand price. For more details on spot pricing, see [spot VMs pricing](https://docs.microsoft.com/azure/virtual-machines/spot-vms#pricing)                                                                                                                                                                                                                                       | float64                                                       |
| tags                      | The tags to be persisted on the agent pool virtual machine scale set.                                                                                                                                                                                                                                                                                                                                                                                                                        | map[string]string                                             |
| type                      | The type of Agent Pool.                                                                                                                                                                                                                                                                                                                                                                                                                                                                      | [AgentPoolType](#AgentPoolType)                               |
| upgradeSettings           | Settings for upgrading the agentpool                                                                                                                                                                                                                                                                                                                                                                                                                                                         | [AgentPoolUpgradeSettings_ARM](#AgentPoolUpgradeSettings_ARM) |
| vmSize                    | VM size availability varies by region. If a node contains insufficient compute resources (memory, cpu, etc) pods might fail to run correctly. For more details on restricted VM sizes, see: https://docs.microsoft.com/azure/aks/quotas-skus-regions                                                                                                                                                                                                                                         | string                                                        |
| vnetSubnetID              |                                                                                                                                                                                                                                                                                                                                                                                                                                                                                              | string                                                        |

<a id="ManagedClusterAgentPoolProfileProperties_STATUS_ARM"></a>ManagedClusterAgentPoolProfileProperties_STATUS_ARM
-------------------------------------------------------------------------------------------------------------------

Properties for the container service agent pool profile.

Used by: [ManagedClusters_AgentPool_STATUS_ARM](#ManagedClusters_AgentPool_STATUS_ARM).

| Property                  | Description                                                                                                                                                                                                                                                                                                                                                                                                                                                                                  | Type                                                                        |
|---------------------------|----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|-----------------------------------------------------------------------------|
| availabilityZones         | The list of Availability zones to use for nodes. This can only be specified if the AgentPoolType property is 'VirtualMachineScaleSets'.                                                                                                                                                                                                                                                                                                                                                      | string[]                                                                    |
| count                     | Number of agents (VMs) to host docker containers. Allowed values must be in the range of 0 to 1000 (inclusive) for user pools and in the range of 1 to 1000 (inclusive) for system pools. The default value is 1.                                                                                                                                                                                                                                                                            | int                                                                         |
| enableAutoScaling         | Whether to enable auto-scaler                                                                                                                                                                                                                                                                                                                                                                                                                                                                | bool                                                                        |
| enableEncryptionAtHost    | This is only supported on certain VM sizes and in certain Azure regions. For more information, see: https://docs.microsoft.com/azure/aks/enable-host-encryption                                                                                                                                                                                                                                                                                                                              | bool                                                                        |
| enableFIPS                | See [Add a FIPS-enabled node pool](https://docs.microsoft.com/azure/aks/use-multiple-node-pools#add-a-fips-enabled-node-pool-preview) for more details.                                                                                                                                                                                                                                                                                                                                      | bool                                                                        |
| enableNodePublicIP        | Some scenarios may require nodes in a node pool to receive their own dedicated public IP addresses. A common scenario is for gaming workloads, where a console needs to make a direct connection to a cloud virtual machine to minimize hops. For more information see [assigning a public IP per node](https://docs.microsoft.com/azure/aks/use-multiple-node-pools#assign-a-public-ip-per-node-for-your-node-pools). The default is false.                                                 | bool                                                                        |
| enableUltraSSD            | Whether to enable UltraSSD                                                                                                                                                                                                                                                                                                                                                                                                                                                                   | bool                                                                        |
| gpuInstanceProfile        | GPUInstanceProfile to be used to specify GPU MIG instance profile for supported GPU VM SKU.                                                                                                                                                                                                                                                                                                                                                                                                  | [GPUInstanceProfile_STATUS](#GPUInstanceProfile_STATUS)                     |
| kubeletConfig             | The Kubelet configuration on the agent pool nodes.                                                                                                                                                                                                                                                                                                                                                                                                                                           | [KubeletConfig_STATUS_ARM](#KubeletConfig_STATUS_ARM)                       |
| kubeletDiskType           | Determines the placement of emptyDir volumes, container runtime data root, and Kubelet ephemeral storage.                                                                                                                                                                                                                                                                                                                                                                                    | [KubeletDiskType_STATUS](#KubeletDiskType_STATUS)                           |
| linuxOSConfig             | The OS configuration of Linux agent nodes.                                                                                                                                                                                                                                                                                                                                                                                                                                                   | [LinuxOSConfig_STATUS_ARM](#LinuxOSConfig_STATUS_ARM)                       |
| maxCount                  | The maximum number of nodes for auto-scaling                                                                                                                                                                                                                                                                                                                                                                                                                                                 | int                                                                         |
| maxPods                   | The maximum number of pods that can run on a node.                                                                                                                                                                                                                                                                                                                                                                                                                                           | int                                                                         |
| minCount                  | The minimum number of nodes for auto-scaling                                                                                                                                                                                                                                                                                                                                                                                                                                                 | int                                                                         |
| mode                      | A cluster must have at least one 'System' Agent Pool at all times. For additional information on agent pool restrictions and best practices, see: https://docs.microsoft.com/azure/aks/use-system-pools                                                                                                                                                                                                                                                                                      | [AgentPoolMode_STATUS](#AgentPoolMode_STATUS)                               |
| nodeImageVersion          | The version of node image                                                                                                                                                                                                                                                                                                                                                                                                                                                                    | string                                                                      |
| nodeLabels                | The node labels to be persisted across all nodes in agent pool.                                                                                                                                                                                                                                                                                                                                                                                                                              | map[string]string                                                           |
| nodePublicIPPrefixID      | This is of the form: /&ZeroWidthSpace;subscriptions/&ZeroWidthSpace;{subscriptionId}/&ZeroWidthSpace;resourceGroups/&ZeroWidthSpace;{resourceGroupName}/&ZeroWidthSpace;providers/&ZeroWidthSpace;Microsoft.Network/publicIPPrefixes/{publicIPPrefixName}                                                                                                                                                                                                                                    | string                                                                      |
| nodeTaints                | The taints added to new nodes during node pool create and scale. For example, key=value:NoSchedule.                                                                                                                                                                                                                                                                                                                                                                                          | string[]                                                                    |
| orchestratorVersion       | As a best practice, you should upgrade all node pools in an AKS cluster to the same Kubernetes version. The node pool version must have the same major version as the control plane. The node pool minor version must be within two minor versions of the control plane version. The node pool version cannot be greater than the control plane version. For more information see [upgrading a node pool](https://docs.microsoft.com/azure/aks/use-multiple-node-pools#upgrade-a-node-pool). | string                                                                      |
| osDiskSizeGB              |                                                                                                                                                                                                                                                                                                                                                                                                                                                                                              | int                                                                         |
| osDiskType                | The default is 'Ephemeral' if the VM supports it and has a cache disk larger than the requested OSDiskSizeGB. Otherwise, defaults to 'Managed'. May not be changed after creation. For more information see [Ephemeral OS](https://docs.microsoft.com/azure/aks/cluster-configuration#ephemeral-os).                                                                                                                                                                                         | [OSDiskType_STATUS](#OSDiskType_STATUS)                                     |
| osSKU                     | Specifies an OS SKU. This value must not be specified if OSType is Windows.                                                                                                                                                                                                                                                                                                                                                                                                                  | [OSSKU_STATUS](#OSSKU_STATUS)                                               |
| osType                    | The operating system type. The default is Linux.                                                                                                                                                                                                                                                                                                                                                                                                                                             | [OSType_STATUS](#OSType_STATUS)                                             |
| podSubnetID               | If omitted, pod IPs are statically assigned on the node subnet (see vnetSubnetID for more details). This is of the form: /&ZeroWidthSpace;subscriptions/&ZeroWidthSpace;{subscriptionId}/&ZeroWidthSpace;resourceGroups/&ZeroWidthSpace;{resourceGroupName}/&ZeroWidthSpace;providers/&ZeroWidthSpace;Microsoft.Network/virtualNetworks/{virtualNetworkName}/subnets/{subnetName}                                                                                                            | string                                                                      |
| powerState                | Describes whether the Agent Pool is Running or Stopped                                                                                                                                                                                                                                                                                                                                                                                                                                       | [PowerState_STATUS_ARM](#PowerState_STATUS_ARM)                             |
| provisioningState         | The current deployment or provisioning state.                                                                                                                                                                                                                                                                                                                                                                                                                                                | string                                                                      |
| proximityPlacementGroupID | The ID for Proximity Placement Group.                                                                                                                                                                                                                                                                                                                                                                                                                                                        | string                                                                      |
| scaleSetEvictionPolicy    | This cannot be specified unless the scaleSetPriority is 'Spot'. If not specified, the default is 'Delete'.                                                                                                                                                                                                                                                                                                                                                                                   | [ScaleSetEvictionPolicy_STATUS](#ScaleSetEvictionPolicy_STATUS)             |
| scaleSetPriority          | The Virtual Machine Scale Set priority. If not specified, the default is 'Regular'.                                                                                                                                                                                                                                                                                                                                                                                                          | [ScaleSetPriority_STATUS](#ScaleSetPriority_STATUS)                         |
| spotMaxPrice              | Possible values are any decimal value greater than zero or -1 which indicates the willingness to pay any on-demand price. For more details on spot pricing, see [spot VMs pricing](https://docs.microsoft.com/azure/virtual-machines/spot-vms#pricing)                                                                                                                                                                                                                                       | float64                                                                     |
| tags                      | The tags to be persisted on the agent pool virtual machine scale set.                                                                                                                                                                                                                                                                                                                                                                                                                        | map[string]string                                                           |
| type                      | The type of Agent Pool.                                                                                                                                                                                                                                                                                                                                                                                                                                                                      | [AgentPoolType_STATUS](#AgentPoolType_STATUS)                               |
| upgradeSettings           | Settings for upgrading the agentpool                                                                                                                                                                                                                                                                                                                                                                                                                                                         | [AgentPoolUpgradeSettings_STATUS_ARM](#AgentPoolUpgradeSettings_STATUS_ARM) |
| vmSize                    | VM size availability varies by region. If a node contains insufficient compute resources (memory, cpu, etc) pods might fail to run correctly. For more details on restricted VM sizes, see: https://docs.microsoft.com/azure/aks/quotas-skus-regions                                                                                                                                                                                                                                         | string                                                                      |
| vnetSubnetID              | If this is not specified, a VNET and subnet will be generated and used. If no podSubnetID is specified, this applies to nodes and pods, otherwise it applies to just nodes. This is of the form: /&ZeroWidthSpace;subscriptions/&ZeroWidthSpace;{subscriptionId}/&ZeroWidthSpace;resourceGroups/&ZeroWidthSpace;{resourceGroupName}/&ZeroWidthSpace;providers/&ZeroWidthSpace;Microsoft.Network/virtualNetworks/{virtualNetworkName}/subnets/{subnetName}                                    | string                                                                      |

<a id="ManagedClusterAgentPoolProfile_ARM"></a>ManagedClusterAgentPoolProfile_ARM
---------------------------------------------------------------------------------

Profile for the container service agent pool.

Used by: [ManagedClusterProperties_ARM](#ManagedClusterProperties_ARM).

| Property                  | Description                                                                                                                                                                                                                                                                                                                                                                                                                                                                                  | Type                                                          |
|---------------------------|----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|---------------------------------------------------------------|
| availabilityZones         | The list of Availability zones to use for nodes. This can only be specified if the AgentPoolType property is 'VirtualMachineScaleSets'.                                                                                                                                                                                                                                                                                                                                                      | string[]                                                      |
| count                     | Number of agents (VMs) to host docker containers. Allowed values must be in the range of 0 to 1000 (inclusive) for user pools and in the range of 1 to 1000 (inclusive) for system pools. The default value is 1.                                                                                                                                                                                                                                                                            | int                                                           |
| enableAutoScaling         | Whether to enable auto-scaler                                                                                                                                                                                                                                                                                                                                                                                                                                                                | bool                                                          |
| enableEncryptionAtHost    | This is only supported on certain VM sizes and in certain Azure regions. For more information, see: https://docs.microsoft.com/azure/aks/enable-host-encryption                                                                                                                                                                                                                                                                                                                              | bool                                                          |
| enableFIPS                | See [Add a FIPS-enabled node pool](https://docs.microsoft.com/azure/aks/use-multiple-node-pools#add-a-fips-enabled-node-pool-preview) for more details.                                                                                                                                                                                                                                                                                                                                      | bool                                                          |
| enableNodePublicIP        | Some scenarios may require nodes in a node pool to receive their own dedicated public IP addresses. A common scenario is for gaming workloads, where a console needs to make a direct connection to a cloud virtual machine to minimize hops. For more information see [assigning a public IP per node](https://docs.microsoft.com/azure/aks/use-multiple-node-pools#assign-a-public-ip-per-node-for-your-node-pools). The default is false.                                                 | bool                                                          |
| enableUltraSSD            | Whether to enable UltraSSD                                                                                                                                                                                                                                                                                                                                                                                                                                                                   | bool                                                          |
| gpuInstanceProfile        | GPUInstanceProfile to be used to specify GPU MIG instance profile for supported GPU VM SKU.                                                                                                                                                                                                                                                                                                                                                                                                  | [GPUInstanceProfile](#GPUInstanceProfile)                     |
| kubeletConfig             | The Kubelet configuration on the agent pool nodes.                                                                                                                                                                                                                                                                                                                                                                                                                                           | [KubeletConfig_ARM](#KubeletConfig_ARM)                       |
| kubeletDiskType           | Determines the placement of emptyDir volumes, container runtime data root, and Kubelet ephemeral storage.                                                                                                                                                                                                                                                                                                                                                                                    | [KubeletDiskType](#KubeletDiskType)                           |
| linuxOSConfig             | The OS configuration of Linux agent nodes.                                                                                                                                                                                                                                                                                                                                                                                                                                                   | [LinuxOSConfig_ARM](#LinuxOSConfig_ARM)                       |
| maxCount                  | The maximum number of nodes for auto-scaling                                                                                                                                                                                                                                                                                                                                                                                                                                                 | int                                                           |
| maxPods                   | The maximum number of pods that can run on a node.                                                                                                                                                                                                                                                                                                                                                                                                                                           | int                                                           |
| minCount                  | The minimum number of nodes for auto-scaling                                                                                                                                                                                                                                                                                                                                                                                                                                                 | int                                                           |
| mode                      | A cluster must have at least one 'System' Agent Pool at all times. For additional information on agent pool restrictions and best practices, see: https://docs.microsoft.com/azure/aks/use-system-pools                                                                                                                                                                                                                                                                                      | [AgentPoolMode](#AgentPoolMode)                               |
| name                      | Windows agent pool names must be 6 characters or less.                                                                                                                                                                                                                                                                                                                                                                                                                                       | string                                                        |
| nodeLabels                | The node labels to be persisted across all nodes in agent pool.                                                                                                                                                                                                                                                                                                                                                                                                                              | map[string]string                                             |
| nodePublicIPPrefixID      |                                                                                                                                                                                                                                                                                                                                                                                                                                                                                              | string                                                        |
| nodeTaints                | The taints added to new nodes during node pool create and scale. For example, key=value:NoSchedule.                                                                                                                                                                                                                                                                                                                                                                                          | string[]                                                      |
| orchestratorVersion       | As a best practice, you should upgrade all node pools in an AKS cluster to the same Kubernetes version. The node pool version must have the same major version as the control plane. The node pool minor version must be within two minor versions of the control plane version. The node pool version cannot be greater than the control plane version. For more information see [upgrading a node pool](https://docs.microsoft.com/azure/aks/use-multiple-node-pools#upgrade-a-node-pool). | string                                                        |
| osDiskSizeGB              |                                                                                                                                                                                                                                                                                                                                                                                                                                                                                              | [ContainerServiceOSDisk](#ContainerServiceOSDisk)             |
| osDiskType                | The default is 'Ephemeral' if the VM supports it and has a cache disk larger than the requested OSDiskSizeGB. Otherwise, defaults to 'Managed'. May not be changed after creation. For more information see [Ephemeral OS](https://docs.microsoft.com/azure/aks/cluster-configuration#ephemeral-os).                                                                                                                                                                                         | [OSDiskType](#OSDiskType)                                     |
| osSKU                     | Specifies an OS SKU. This value must not be specified if OSType is Windows.                                                                                                                                                                                                                                                                                                                                                                                                                  | [OSSKU](#OSSKU)                                               |
| osType                    | The operating system type. The default is Linux.                                                                                                                                                                                                                                                                                                                                                                                                                                             | [OSType](#OSType)                                             |
| podSubnetID               |                                                                                                                                                                                                                                                                                                                                                                                                                                                                                              | string                                                        |
| proximityPlacementGroupID | The ID for Proximity Placement Group.                                                                                                                                                                                                                                                                                                                                                                                                                                                        | string                                                        |
| scaleSetEvictionPolicy    | This cannot be specified unless the scaleSetPriority is 'Spot'. If not specified, the default is 'Delete'.                                                                                                                                                                                                                                                                                                                                                                                   | [ScaleSetEvictionPolicy](#ScaleSetEvictionPolicy)             |
| scaleSetPriority          | The Virtual Machine Scale Set priority. If not specified, the default is 'Regular'.                                                                                                                                                                                                                                                                                                                                                                                                          | [ScaleSetPriority](#ScaleSetPriority)                         |
| spotMaxPrice              | Possible values are any decimal value greater than zero or -1 which indicates the willingness to pay any on-demand price. For more details on spot pricing, see [spot VMs pricing](https://docs.microsoft.com/azure/virtual-machines/spot-vms#pricing)                                                                                                                                                                                                                                       | float64                                                       |
| tags                      | The tags to be persisted on the agent pool virtual machine scale set.                                                                                                                                                                                                                                                                                                                                                                                                                        | map[string]string                                             |
| type                      | The type of Agent Pool.                                                                                                                                                                                                                                                                                                                                                                                                                                                                      | [AgentPoolType](#AgentPoolType)                               |
| upgradeSettings           | Settings for upgrading the agentpool                                                                                                                                                                                                                                                                                                                                                                                                                                                         | [AgentPoolUpgradeSettings_ARM](#AgentPoolUpgradeSettings_ARM) |
| vmSize                    | VM size availability varies by region. If a node contains insufficient compute resources (memory, cpu, etc) pods might fail to run correctly. For more details on restricted VM sizes, see: https://docs.microsoft.com/azure/aks/quotas-skus-regions                                                                                                                                                                                                                                         | string                                                        |
| vnetSubnetID              |                                                                                                                                                                                                                                                                                                                                                                                                                                                                                              | string                                                        |

<a id="ManagedClusterAgentPoolProfile_STATUS"></a>ManagedClusterAgentPoolProfile_STATUS
---------------------------------------------------------------------------------------

Profile for the container service agent pool.

Used by: [ManagedCluster_STATUS](#ManagedCluster_STATUS).

| Property                  | Description                                                                                                                                                                                                                                                                                                                                                                                                                                                                                  | Type                                                                |
|---------------------------|----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|---------------------------------------------------------------------|
| availabilityZones         | The list of Availability zones to use for nodes. This can only be specified if the AgentPoolType property is 'VirtualMachineScaleSets'.                                                                                                                                                                                                                                                                                                                                                      | string[]                                                            |
| count                     | Number of agents (VMs) to host docker containers. Allowed values must be in the range of 0 to 1000 (inclusive) for user pools and in the range of 1 to 1000 (inclusive) for system pools. The default value is 1.                                                                                                                                                                                                                                                                            | int                                                                 |
| enableAutoScaling         | Whether to enable auto-scaler                                                                                                                                                                                                                                                                                                                                                                                                                                                                | bool                                                                |
| enableEncryptionAtHost    | This is only supported on certain VM sizes and in certain Azure regions. For more information, see: https://docs.microsoft.com/azure/aks/enable-host-encryption                                                                                                                                                                                                                                                                                                                              | bool                                                                |
| enableFIPS                | See [Add a FIPS-enabled node pool](https://docs.microsoft.com/azure/aks/use-multiple-node-pools#add-a-fips-enabled-node-pool-preview) for more details.                                                                                                                                                                                                                                                                                                                                      | bool                                                                |
| enableNodePublicIP        | Some scenarios may require nodes in a node pool to receive their own dedicated public IP addresses. A common scenario is for gaming workloads, where a console needs to make a direct connection to a cloud virtual machine to minimize hops. For more information see [assigning a public IP per node](https://docs.microsoft.com/azure/aks/use-multiple-node-pools#assign-a-public-ip-per-node-for-your-node-pools). The default is false.                                                 | bool                                                                |
| enableUltraSSD            | Whether to enable UltraSSD                                                                                                                                                                                                                                                                                                                                                                                                                                                                   | bool                                                                |
| gpuInstanceProfile        | GPUInstanceProfile to be used to specify GPU MIG instance profile for supported GPU VM SKU.                                                                                                                                                                                                                                                                                                                                                                                                  | [GPUInstanceProfile_STATUS](#GPUInstanceProfile_STATUS)             |
| kubeletConfig             | The Kubelet configuration on the agent pool nodes.                                                                                                                                                                                                                                                                                                                                                                                                                                           | [KubeletConfig_STATUS](#KubeletConfig_STATUS)                       |
| kubeletDiskType           | Determines the placement of emptyDir volumes, container runtime data root, and Kubelet ephemeral storage.                                                                                                                                                                                                                                                                                                                                                                                    | [KubeletDiskType_STATUS](#KubeletDiskType_STATUS)                   |
| linuxOSConfig             | The OS configuration of Linux agent nodes.                                                                                                                                                                                                                                                                                                                                                                                                                                                   | [LinuxOSConfig_STATUS](#LinuxOSConfig_STATUS)                       |
| maxCount                  | The maximum number of nodes for auto-scaling                                                                                                                                                                                                                                                                                                                                                                                                                                                 | int                                                                 |
| maxPods                   | The maximum number of pods that can run on a node.                                                                                                                                                                                                                                                                                                                                                                                                                                           | int                                                                 |
| minCount                  | The minimum number of nodes for auto-scaling                                                                                                                                                                                                                                                                                                                                                                                                                                                 | int                                                                 |
| mode                      | A cluster must have at least one 'System' Agent Pool at all times. For additional information on agent pool restrictions and best practices, see: https://docs.microsoft.com/azure/aks/use-system-pools                                                                                                                                                                                                                                                                                      | [AgentPoolMode_STATUS](#AgentPoolMode_STATUS)                       |
| name                      | Windows agent pool names must be 6 characters or less.                                                                                                                                                                                                                                                                                                                                                                                                                                       | string                                                              |
| nodeImageVersion          | The version of node image                                                                                                                                                                                                                                                                                                                                                                                                                                                                    | string                                                              |
| nodeLabels                | The node labels to be persisted across all nodes in agent pool.                                                                                                                                                                                                                                                                                                                                                                                                                              | map[string]string                                                   |
| nodePublicIPPrefixID      | This is of the form: /&ZeroWidthSpace;subscriptions/&ZeroWidthSpace;{subscriptionId}/&ZeroWidthSpace;resourceGroups/&ZeroWidthSpace;{resourceGroupName}/&ZeroWidthSpace;providers/&ZeroWidthSpace;Microsoft.Network/publicIPPrefixes/{publicIPPrefixName}                                                                                                                                                                                                                                    | string                                                              |
| nodeTaints                | The taints added to new nodes during node pool create and scale. For example, key=value:NoSchedule.                                                                                                                                                                                                                                                                                                                                                                                          | string[]                                                            |
| orchestratorVersion       | As a best practice, you should upgrade all node pools in an AKS cluster to the same Kubernetes version. The node pool version must have the same major version as the control plane. The node pool minor version must be within two minor versions of the control plane version. The node pool version cannot be greater than the control plane version. For more information see [upgrading a node pool](https://docs.microsoft.com/azure/aks/use-multiple-node-pools#upgrade-a-node-pool). | string                                                              |
| osDiskSizeGB              |                                                                                                                                                                                                                                                                                                                                                                                                                                                                                              | int                                                                 |
| osDiskType                | The default is 'Ephemeral' if the VM supports it and has a cache disk larger than the requested OSDiskSizeGB. Otherwise, defaults to 'Managed'. May not be changed after creation. For more information see [Ephemeral OS](https://docs.microsoft.com/azure/aks/cluster-configuration#ephemeral-os).                                                                                                                                                                                         | [OSDiskType_STATUS](#OSDiskType_STATUS)                             |
| osSKU                     | Specifies an OS SKU. This value must not be specified if OSType is Windows.                                                                                                                                                                                                                                                                                                                                                                                                                  | [OSSKU_STATUS](#OSSKU_STATUS)                                       |
| osType                    | The operating system type. The default is Linux.                                                                                                                                                                                                                                                                                                                                                                                                                                             | [OSType_STATUS](#OSType_STATUS)                                     |
| podSubnetID               | If omitted, pod IPs are statically assigned on the node subnet (see vnetSubnetID for more details). This is of the form: /&ZeroWidthSpace;subscriptions/&ZeroWidthSpace;{subscriptionId}/&ZeroWidthSpace;resourceGroups/&ZeroWidthSpace;{resourceGroupName}/&ZeroWidthSpace;providers/&ZeroWidthSpace;Microsoft.Network/virtualNetworks/{virtualNetworkName}/subnets/{subnetName}                                                                                                            | string                                                              |
| powerState                | Describes whether the Agent Pool is Running or Stopped                                                                                                                                                                                                                                                                                                                                                                                                                                       | [PowerState_STATUS](#PowerState_STATUS)                             |
| provisioningState         | The current deployment or provisioning state.                                                                                                                                                                                                                                                                                                                                                                                                                                                | string                                                              |
| proximityPlacementGroupID | The ID for Proximity Placement Group.                                                                                                                                                                                                                                                                                                                                                                                                                                                        | string                                                              |
| scaleSetEvictionPolicy    | This cannot be specified unless the scaleSetPriority is 'Spot'. If not specified, the default is 'Delete'.                                                                                                                                                                                                                                                                                                                                                                                   | [ScaleSetEvictionPolicy_STATUS](#ScaleSetEvictionPolicy_STATUS)     |
| scaleSetPriority          | The Virtual Machine Scale Set priority. If not specified, the default is 'Regular'.                                                                                                                                                                                                                                                                                                                                                                                                          | [ScaleSetPriority_STATUS](#ScaleSetPriority_STATUS)                 |
| spotMaxPrice              | Possible values are any decimal value greater than zero or -1 which indicates the willingness to pay any on-demand price. For more details on spot pricing, see [spot VMs pricing](https://docs.microsoft.com/azure/virtual-machines/spot-vms#pricing)                                                                                                                                                                                                                                       | float64                                                             |
| tags                      | The tags to be persisted on the agent pool virtual machine scale set.                                                                                                                                                                                                                                                                                                                                                                                                                        | map[string]string                                                   |
| type                      | The type of Agent Pool.                                                                                                                                                                                                                                                                                                                                                                                                                                                                      | [AgentPoolType_STATUS](#AgentPoolType_STATUS)                       |
| upgradeSettings           | Settings for upgrading the agentpool                                                                                                                                                                                                                                                                                                                                                                                                                                                         | [AgentPoolUpgradeSettings_STATUS](#AgentPoolUpgradeSettings_STATUS) |
| vmSize                    | VM size availability varies by region. If a node contains insufficient compute resources (memory, cpu, etc) pods might fail to run correctly. For more details on restricted VM sizes, see: https://docs.microsoft.com/azure/aks/quotas-skus-regions                                                                                                                                                                                                                                         | string                                                              |
| vnetSubnetID              | If this is not specified, a VNET and subnet will be generated and used. If no podSubnetID is specified, this applies to nodes and pods, otherwise it applies to just nodes. This is of the form: /&ZeroWidthSpace;subscriptions/&ZeroWidthSpace;{subscriptionId}/&ZeroWidthSpace;resourceGroups/&ZeroWidthSpace;{resourceGroupName}/&ZeroWidthSpace;providers/&ZeroWidthSpace;Microsoft.Network/virtualNetworks/{virtualNetworkName}/subnets/{subnetName}                                    | string                                                              |

<a id="ManagedClusterAgentPoolProfile_STATUS_ARM"></a>ManagedClusterAgentPoolProfile_STATUS_ARM
-----------------------------------------------------------------------------------------------

Profile for the container service agent pool.

Used by: [ManagedClusterProperties_STATUS_ARM](#ManagedClusterProperties_STATUS_ARM).

| Property                  | Description                                                                                                                                                                                                                                                                                                                                                                                                                                                                                  | Type                                                                        |
|---------------------------|----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|-----------------------------------------------------------------------------|
| availabilityZones         | The list of Availability zones to use for nodes. This can only be specified if the AgentPoolType property is 'VirtualMachineScaleSets'.                                                                                                                                                                                                                                                                                                                                                      | string[]                                                                    |
| count                     | Number of agents (VMs) to host docker containers. Allowed values must be in the range of 0 to 1000 (inclusive) for user pools and in the range of 1 to 1000 (inclusive) for system pools. The default value is 1.                                                                                                                                                                                                                                                                            | int                                                                         |
| enableAutoScaling         | Whether to enable auto-scaler                                                                                                                                                                                                                                                                                                                                                                                                                                                                | bool                                                                        |
| enableEncryptionAtHost    | This is only supported on certain VM sizes and in certain Azure regions. For more information, see: https://docs.microsoft.com/azure/aks/enable-host-encryption                                                                                                                                                                                                                                                                                                                              | bool                                                                        |
| enableFIPS                | See [Add a FIPS-enabled node pool](https://docs.microsoft.com/azure/aks/use-multiple-node-pools#add-a-fips-enabled-node-pool-preview) for more details.                                                                                                                                                                                                                                                                                                                                      | bool                                                                        |
| enableNodePublicIP        | Some scenarios may require nodes in a node pool to receive their own dedicated public IP addresses. A common scenario is for gaming workloads, where a console needs to make a direct connection to a cloud virtual machine to minimize hops. For more information see [assigning a public IP per node](https://docs.microsoft.com/azure/aks/use-multiple-node-pools#assign-a-public-ip-per-node-for-your-node-pools). The default is false.                                                 | bool                                                                        |
| enableUltraSSD            | Whether to enable UltraSSD                                                                                                                                                                                                                                                                                                                                                                                                                                                                   | bool                                                                        |
| gpuInstanceProfile        | GPUInstanceProfile to be used to specify GPU MIG instance profile for supported GPU VM SKU.                                                                                                                                                                                                                                                                                                                                                                                                  | [GPUInstanceProfile_STATUS](#GPUInstanceProfile_STATUS)                     |
| kubeletConfig             | The Kubelet configuration on the agent pool nodes.                                                                                                                                                                                                                                                                                                                                                                                                                                           | [KubeletConfig_STATUS_ARM](#KubeletConfig_STATUS_ARM)                       |
| kubeletDiskType           | Determines the placement of emptyDir volumes, container runtime data root, and Kubelet ephemeral storage.                                                                                                                                                                                                                                                                                                                                                                                    | [KubeletDiskType_STATUS](#KubeletDiskType_STATUS)                           |
| linuxOSConfig             | The OS configuration of Linux agent nodes.                                                                                                                                                                                                                                                                                                                                                                                                                                                   | [LinuxOSConfig_STATUS_ARM](#LinuxOSConfig_STATUS_ARM)                       |
| maxCount                  | The maximum number of nodes for auto-scaling                                                                                                                                                                                                                                                                                                                                                                                                                                                 | int                                                                         |
| maxPods                   | The maximum number of pods that can run on a node.                                                                                                                                                                                                                                                                                                                                                                                                                                           | int                                                                         |
| minCount                  | The minimum number of nodes for auto-scaling                                                                                                                                                                                                                                                                                                                                                                                                                                                 | int                                                                         |
| mode                      | A cluster must have at least one 'System' Agent Pool at all times. For additional information on agent pool restrictions and best practices, see: https://docs.microsoft.com/azure/aks/use-system-pools                                                                                                                                                                                                                                                                                      | [AgentPoolMode_STATUS](#AgentPoolMode_STATUS)                               |
| name                      | Windows agent pool names must be 6 characters or less.                                                                                                                                                                                                                                                                                                                                                                                                                                       | string                                                                      |
| nodeImageVersion          | The version of node image                                                                                                                                                                                                                                                                                                                                                                                                                                                                    | string                                                                      |
| nodeLabels                | The node labels to be persisted across all nodes in agent pool.                                                                                                                                                                                                                                                                                                                                                                                                                              | map[string]string                                                           |
| nodePublicIPPrefixID      | This is of the form: /&ZeroWidthSpace;subscriptions/&ZeroWidthSpace;{subscriptionId}/&ZeroWidthSpace;resourceGroups/&ZeroWidthSpace;{resourceGroupName}/&ZeroWidthSpace;providers/&ZeroWidthSpace;Microsoft.Network/publicIPPrefixes/{publicIPPrefixName}                                                                                                                                                                                                                                    | string                                                                      |
| nodeTaints                | The taints added to new nodes during node pool create and scale. For example, key=value:NoSchedule.                                                                                                                                                                                                                                                                                                                                                                                          | string[]                                                                    |
| orchestratorVersion       | As a best practice, you should upgrade all node pools in an AKS cluster to the same Kubernetes version. The node pool version must have the same major version as the control plane. The node pool minor version must be within two minor versions of the control plane version. The node pool version cannot be greater than the control plane version. For more information see [upgrading a node pool](https://docs.microsoft.com/azure/aks/use-multiple-node-pools#upgrade-a-node-pool). | string                                                                      |
| osDiskSizeGB              |                                                                                                                                                                                                                                                                                                                                                                                                                                                                                              | int                                                                         |
| osDiskType                | The default is 'Ephemeral' if the VM supports it and has a cache disk larger than the requested OSDiskSizeGB. Otherwise, defaults to 'Managed'. May not be changed after creation. For more information see [Ephemeral OS](https://docs.microsoft.com/azure/aks/cluster-configuration#ephemeral-os).                                                                                                                                                                                         | [OSDiskType_STATUS](#OSDiskType_STATUS)                                     |
| osSKU                     | Specifies an OS SKU. This value must not be specified if OSType is Windows.                                                                                                                                                                                                                                                                                                                                                                                                                  | [OSSKU_STATUS](#OSSKU_STATUS)                                               |
| osType                    | The operating system type. The default is Linux.                                                                                                                                                                                                                                                                                                                                                                                                                                             | [OSType_STATUS](#OSType_STATUS)                                             |
| podSubnetID               | If omitted, pod IPs are statically assigned on the node subnet (see vnetSubnetID for more details). This is of the form: /&ZeroWidthSpace;subscriptions/&ZeroWidthSpace;{subscriptionId}/&ZeroWidthSpace;resourceGroups/&ZeroWidthSpace;{resourceGroupName}/&ZeroWidthSpace;providers/&ZeroWidthSpace;Microsoft.Network/virtualNetworks/{virtualNetworkName}/subnets/{subnetName}                                                                                                            | string                                                                      |
| powerState                | Describes whether the Agent Pool is Running or Stopped                                                                                                                                                                                                                                                                                                                                                                                                                                       | [PowerState_STATUS_ARM](#PowerState_STATUS_ARM)                             |
| provisioningState         | The current deployment or provisioning state.                                                                                                                                                                                                                                                                                                                                                                                                                                                | string                                                                      |
| proximityPlacementGroupID | The ID for Proximity Placement Group.                                                                                                                                                                                                                                                                                                                                                                                                                                                        | string                                                                      |
| scaleSetEvictionPolicy    | This cannot be specified unless the scaleSetPriority is 'Spot'. If not specified, the default is 'Delete'.                                                                                                                                                                                                                                                                                                                                                                                   | [ScaleSetEvictionPolicy_STATUS](#ScaleSetEvictionPolicy_STATUS)             |
| scaleSetPriority          | The Virtual Machine Scale Set priority. If not specified, the default is 'Regular'.                                                                                                                                                                                                                                                                                                                                                                                                          | [ScaleSetPriority_STATUS](#ScaleSetPriority_STATUS)                         |
| spotMaxPrice              | Possible values are any decimal value greater than zero or -1 which indicates the willingness to pay any on-demand price. For more details on spot pricing, see [spot VMs pricing](https://docs.microsoft.com/azure/virtual-machines/spot-vms#pricing)                                                                                                                                                                                                                                       | float64                                                                     |
| tags                      | The tags to be persisted on the agent pool virtual machine scale set.                                                                                                                                                                                                                                                                                                                                                                                                                        | map[string]string                                                           |
| type                      | The type of Agent Pool.                                                                                                                                                                                                                                                                                                                                                                                                                                                                      | [AgentPoolType_STATUS](#AgentPoolType_STATUS)                               |
| upgradeSettings           | Settings for upgrading the agentpool                                                                                                                                                                                                                                                                                                                                                                                                                                                         | [AgentPoolUpgradeSettings_STATUS_ARM](#AgentPoolUpgradeSettings_STATUS_ARM) |
| vmSize                    | VM size availability varies by region. If a node contains insufficient compute resources (memory, cpu, etc) pods might fail to run correctly. For more details on restricted VM sizes, see: https://docs.microsoft.com/azure/aks/quotas-skus-regions                                                                                                                                                                                                                                         | string                                                                      |
| vnetSubnetID              | If this is not specified, a VNET and subnet will be generated and used. If no podSubnetID is specified, this applies to nodes and pods, otherwise it applies to just nodes. This is of the form: /&ZeroWidthSpace;subscriptions/&ZeroWidthSpace;{subscriptionId}/&ZeroWidthSpace;resourceGroups/&ZeroWidthSpace;{resourceGroupName}/&ZeroWidthSpace;providers/&ZeroWidthSpace;Microsoft.Network/virtualNetworks/{virtualNetworkName}/subnets/{subnetName}                                    | string                                                                      |

<a id="ManagedClusterAutoUpgradeProfile"></a>ManagedClusterAutoUpgradeProfile
-----------------------------------------------------------------------------

Auto upgrade profile for a managed cluster.

Used by: [ManagedCluster_Spec](#ManagedCluster_Spec).

| Property       | Description                                                                                                                                             | Type                                                                                                |
|----------------|---------------------------------------------------------------------------------------------------------------------------------------------------------|-----------------------------------------------------------------------------------------------------|
| upgradeChannel | For more information see [setting the AKS cluster auto-upgrade channel](https://docs.microsoft.com/azure/aks/upgrade-cluster#set-auto-upgrade-channel). | [ManagedClusterAutoUpgradeProfile_UpgradeChannel](#ManagedClusterAutoUpgradeProfile_UpgradeChannel) |

<a id="ManagedClusterAutoUpgradeProfile_ARM"></a>ManagedClusterAutoUpgradeProfile_ARM
-------------------------------------------------------------------------------------

Auto upgrade profile for a managed cluster.

Used by: [ManagedClusterProperties_ARM](#ManagedClusterProperties_ARM).

| Property       | Description                                                                                                                                             | Type                                                                                                |
|----------------|---------------------------------------------------------------------------------------------------------------------------------------------------------|-----------------------------------------------------------------------------------------------------|
| upgradeChannel | For more information see [setting the AKS cluster auto-upgrade channel](https://docs.microsoft.com/azure/aks/upgrade-cluster#set-auto-upgrade-channel). | [ManagedClusterAutoUpgradeProfile_UpgradeChannel](#ManagedClusterAutoUpgradeProfile_UpgradeChannel) |

<a id="ManagedClusterAutoUpgradeProfile_STATUS"></a>ManagedClusterAutoUpgradeProfile_STATUS
-------------------------------------------------------------------------------------------

Auto upgrade profile for a managed cluster.

Used by: [ManagedCluster_STATUS](#ManagedCluster_STATUS).

| Property       | Description                                                                                                                                             | Type                                                                                                              |
|----------------|---------------------------------------------------------------------------------------------------------------------------------------------------------|-------------------------------------------------------------------------------------------------------------------|
| upgradeChannel | For more information see [setting the AKS cluster auto-upgrade channel](https://docs.microsoft.com/azure/aks/upgrade-cluster#set-auto-upgrade-channel). | [ManagedClusterAutoUpgradeProfile_UpgradeChannel_STATUS](#ManagedClusterAutoUpgradeProfile_UpgradeChannel_STATUS) |

<a id="ManagedClusterAutoUpgradeProfile_STATUS_ARM"></a>ManagedClusterAutoUpgradeProfile_STATUS_ARM
---------------------------------------------------------------------------------------------------

Auto upgrade profile for a managed cluster.

Used by: [ManagedClusterProperties_STATUS_ARM](#ManagedClusterProperties_STATUS_ARM).

| Property       | Description                                                                                                                                             | Type                                                                                                              |
|----------------|---------------------------------------------------------------------------------------------------------------------------------------------------------|-------------------------------------------------------------------------------------------------------------------|
| upgradeChannel | For more information see [setting the AKS cluster auto-upgrade channel](https://docs.microsoft.com/azure/aks/upgrade-cluster#set-auto-upgrade-channel). | [ManagedClusterAutoUpgradeProfile_UpgradeChannel_STATUS](#ManagedClusterAutoUpgradeProfile_UpgradeChannel_STATUS) |

<a id="ManagedClusterAutoUpgradeProfile_UpgradeChannel"></a>ManagedClusterAutoUpgradeProfile_UpgradeChannel
-----------------------------------------------------------------------------------------------------------

Used by: [ManagedClusterAutoUpgradeProfile](#ManagedClusterAutoUpgradeProfile), and [ManagedClusterAutoUpgradeProfile_ARM](#ManagedClusterAutoUpgradeProfile_ARM).

| Value        | Description |
|--------------|-------------|
| "node-image" |             |
| "none"       |             |
| "patch"      |             |
| "rapid"      |             |
| "stable"     |             |

<a id="ManagedClusterAutoUpgradeProfile_UpgradeChannel_STATUS"></a>ManagedClusterAutoUpgradeProfile_UpgradeChannel_STATUS
-------------------------------------------------------------------------------------------------------------------------

Used by: [ManagedClusterAutoUpgradeProfile_STATUS](#ManagedClusterAutoUpgradeProfile_STATUS), and [ManagedClusterAutoUpgradeProfile_STATUS_ARM](#ManagedClusterAutoUpgradeProfile_STATUS_ARM).

| Value        | Description |
|--------------|-------------|
| "node-image" |             |
| "none"       |             |
| "patch"      |             |
| "rapid"      |             |
| "stable"     |             |

<a id="ManagedClusterHTTPProxyConfig"></a>ManagedClusterHTTPProxyConfig
-----------------------------------------------------------------------

Cluster HTTP proxy configuration.

Used by: [ManagedCluster_Spec](#ManagedCluster_Spec).

| Property   | Description                                                 | Type     |
|------------|-------------------------------------------------------------|----------|
| httpProxy  | The HTTP proxy server endpoint to use.                      | string   |
| httpsProxy | The HTTPS proxy server endpoint to use.                     | string   |
| noProxy    | The endpoints that should not go through proxy.             | string[] |
| trustedCa  | Alternative CA cert to use for connecting to proxy servers. | string   |

<a id="ManagedClusterHTTPProxyConfig_ARM"></a>ManagedClusterHTTPProxyConfig_ARM
-------------------------------------------------------------------------------

Cluster HTTP proxy configuration.

Used by: [ManagedClusterProperties_ARM](#ManagedClusterProperties_ARM).

| Property   | Description                                                 | Type     |
|------------|-------------------------------------------------------------|----------|
| httpProxy  | The HTTP proxy server endpoint to use.                      | string   |
| httpsProxy | The HTTPS proxy server endpoint to use.                     | string   |
| noProxy    | The endpoints that should not go through proxy.             | string[] |
| trustedCa  | Alternative CA cert to use for connecting to proxy servers. | string   |

<a id="ManagedClusterHTTPProxyConfig_STATUS"></a>ManagedClusterHTTPProxyConfig_STATUS
-------------------------------------------------------------------------------------

Cluster HTTP proxy configuration.

Used by: [ManagedCluster_STATUS](#ManagedCluster_STATUS).

| Property   | Description                                                 | Type     |
|------------|-------------------------------------------------------------|----------|
| httpProxy  | The HTTP proxy server endpoint to use.                      | string   |
| httpsProxy | The HTTPS proxy server endpoint to use.                     | string   |
| noProxy    | The endpoints that should not go through proxy.             | string[] |
| trustedCa  | Alternative CA cert to use for connecting to proxy servers. | string   |

<a id="ManagedClusterHTTPProxyConfig_STATUS_ARM"></a>ManagedClusterHTTPProxyConfig_STATUS_ARM
---------------------------------------------------------------------------------------------

Cluster HTTP proxy configuration.

Used by: [ManagedClusterProperties_STATUS_ARM](#ManagedClusterProperties_STATUS_ARM).

| Property   | Description                                                 | Type     |
|------------|-------------------------------------------------------------|----------|
| httpProxy  | The HTTP proxy server endpoint to use.                      | string   |
| httpsProxy | The HTTPS proxy server endpoint to use.                     | string   |
| noProxy    | The endpoints that should not go through proxy.             | string[] |
| trustedCa  | Alternative CA cert to use for connecting to proxy servers. | string   |

<a id="ManagedClusterIdentity"></a>ManagedClusterIdentity
---------------------------------------------------------

Identity for the managed cluster.

Used by: [ManagedCluster_Spec](#ManagedCluster_Spec).

| Property               | Description                                                                                                                                                                                                                                                                                    | Type                                                          |
|------------------------|------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|---------------------------------------------------------------|
| type                   | For more information see [use managed identities in AKS](https://docs.microsoft.com/azure/aks/use-managed-identity).                                                                                                                                                                           | [ManagedClusterIdentity_Type](#ManagedClusterIdentity_Type)   |
| userAssignedIdentities | The keys must be ARM resource IDs in the form: '/&ZeroWidthSpace;subscriptions/&ZeroWidthSpace;{subscriptionId}/&ZeroWidthSpace;resourceGroups/&ZeroWidthSpace;{resourceGroupName}/&ZeroWidthSpace;providers/&ZeroWidthSpace;Microsoft.ManagedIdentity/userAssignedIdentities/{identityName}'. | [UserAssignedIdentityDetails[]](#UserAssignedIdentityDetails) |

<a id="ManagedClusterIdentity_ARM"></a>ManagedClusterIdentity_ARM
-----------------------------------------------------------------

Identity for the managed cluster.

Used by: [ManagedCluster_Spec_ARM](#ManagedCluster_Spec_ARM).

| Property               | Description                                                                                                          | Type                                                        |
|------------------------|----------------------------------------------------------------------------------------------------------------------|-------------------------------------------------------------|
| type                   | For more information see [use managed identities in AKS](https://docs.microsoft.com/azure/aks/use-managed-identity). | [ManagedClusterIdentity_Type](#ManagedClusterIdentity_Type) |
| userAssignedIdentities |                                                                                                                      | map[string]UserAssignedIdentityDetails_ARM                  |

<a id="ManagedClusterIdentity_STATUS"></a>ManagedClusterIdentity_STATUS
-----------------------------------------------------------------------

Identity for the managed cluster.

Used by: [ManagedCluster_STATUS](#ManagedCluster_STATUS).

| Property               | Description                                                                                                                                                                                                                                                                                    | Type                                                                                                                     |
|------------------------|------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|--------------------------------------------------------------------------------------------------------------------------|
| principalId            | The principal id of the system assigned identity which is used by master components.                                                                                                                                                                                                           | string                                                                                                                   |
| tenantId               | The tenant id of the system assigned identity which is used by master components.                                                                                                                                                                                                              | string                                                                                                                   |
| type                   | For more information see [use managed identities in AKS](https://docs.microsoft.com/azure/aks/use-managed-identity).                                                                                                                                                                           | [ManagedClusterIdentity_Type_STATUS](#ManagedClusterIdentity_Type_STATUS)                                                |
| userAssignedIdentities | The keys must be ARM resource IDs in the form: '/&ZeroWidthSpace;subscriptions/&ZeroWidthSpace;{subscriptionId}/&ZeroWidthSpace;resourceGroups/&ZeroWidthSpace;{resourceGroupName}/&ZeroWidthSpace;providers/&ZeroWidthSpace;Microsoft.ManagedIdentity/userAssignedIdentities/{identityName}'. | [map[string]ManagedClusterIdentity_UserAssignedIdentities_STATUS](#ManagedClusterIdentity_UserAssignedIdentities_STATUS) |

<a id="ManagedClusterIdentity_STATUS_ARM"></a>ManagedClusterIdentity_STATUS_ARM
-------------------------------------------------------------------------------

Identity for the managed cluster.

Used by: [ManagedCluster_STATUS_ARM](#ManagedCluster_STATUS_ARM).

| Property               | Description                                                                                                                                                                                                                                                                                    | Type                                                                                                                             |
|------------------------|------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|----------------------------------------------------------------------------------------------------------------------------------|
| principalId            | The principal id of the system assigned identity which is used by master components.                                                                                                                                                                                                           | string                                                                                                                           |
| tenantId               | The tenant id of the system assigned identity which is used by master components.                                                                                                                                                                                                              | string                                                                                                                           |
| type                   | For more information see [use managed identities in AKS](https://docs.microsoft.com/azure/aks/use-managed-identity).                                                                                                                                                                           | [ManagedClusterIdentity_Type_STATUS](#ManagedClusterIdentity_Type_STATUS)                                                        |
| userAssignedIdentities | The keys must be ARM resource IDs in the form: '/&ZeroWidthSpace;subscriptions/&ZeroWidthSpace;{subscriptionId}/&ZeroWidthSpace;resourceGroups/&ZeroWidthSpace;{resourceGroupName}/&ZeroWidthSpace;providers/&ZeroWidthSpace;Microsoft.ManagedIdentity/userAssignedIdentities/{identityName}'. | [map[string]ManagedClusterIdentity_UserAssignedIdentities_STATUS_ARM](#ManagedClusterIdentity_UserAssignedIdentities_STATUS_ARM) |

<a id="ManagedClusterIdentity_Type"></a>ManagedClusterIdentity_Type
-------------------------------------------------------------------

Used by: [ManagedClusterIdentity](#ManagedClusterIdentity), and [ManagedClusterIdentity_ARM](#ManagedClusterIdentity_ARM).

| Value            | Description |
|------------------|-------------|
| "None"           |             |
| "SystemAssigned" |             |
| "UserAssigned"   |             |

<a id="ManagedClusterIdentity_Type_STATUS"></a>ManagedClusterIdentity_Type_STATUS
---------------------------------------------------------------------------------

Used by: [ManagedClusterIdentity_STATUS](#ManagedClusterIdentity_STATUS), and [ManagedClusterIdentity_STATUS_ARM](#ManagedClusterIdentity_STATUS_ARM).

| Value            | Description |
|------------------|-------------|
| "None"           |             |
| "SystemAssigned" |             |
| "UserAssigned"   |             |

<a id="ManagedClusterIdentity_UserAssignedIdentities_STATUS"></a>ManagedClusterIdentity_UserAssignedIdentities_STATUS
---------------------------------------------------------------------------------------------------------------------

Used by: [ManagedClusterIdentity_STATUS](#ManagedClusterIdentity_STATUS).

| Property    | Description                                 | Type   |
|-------------|---------------------------------------------|--------|
| clientId    | The client id of user assigned identity.    | string |
| principalId | The principal id of user assigned identity. | string |

<a id="ManagedClusterIdentity_UserAssignedIdentities_STATUS_ARM"></a>ManagedClusterIdentity_UserAssignedIdentities_STATUS_ARM
-----------------------------------------------------------------------------------------------------------------------------

Used by: [ManagedClusterIdentity_STATUS_ARM](#ManagedClusterIdentity_STATUS_ARM).

| Property    | Description                                 | Type   |
|-------------|---------------------------------------------|--------|
| clientId    | The client id of user assigned identity.    | string |
| principalId | The principal id of user assigned identity. | string |

<a id="ManagedClusterList"></a>ManagedClusterList
-------------------------------------------------

Generator information: - Generated from: /containerservice/resource-manager/Microsoft.ContainerService/aks/stable/2021-05-01/managedClusters.json - ARM URI: /&ZeroWidthSpace;subscriptions/&ZeroWidthSpace;{subscriptionId}/&ZeroWidthSpace;resourceGroups/&ZeroWidthSpace;{resourceGroupName}/&ZeroWidthSpace;providers/&ZeroWidthSpace;Microsoft.ContainerService/managedClusters/{resourceName}

| Property        | Description | Type                                |
|-----------------|-------------|-------------------------------------|
| metav1.TypeMeta |             |                                     |
| metav1.ListMeta |             |                                     |
| items           |             | [ManagedCluster[]](#ManagedCluster) |

<a id="ManagedClusterLoadBalancerProfile"></a>ManagedClusterLoadBalancerProfile
-------------------------------------------------------------------------------

Profile of the managed cluster load balancer.

Used by: [ContainerServiceNetworkProfile](#ContainerServiceNetworkProfile).

| Property               | Description                                                                                                                                                                               | Type                                                                                                          |
|------------------------|-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|---------------------------------------------------------------------------------------------------------------|
| allocatedOutboundPorts | The desired number of allocated SNAT ports per VM. Allowed values are in the range of 0 to 64000 (inclusive). The default value is 0 which results in Azure dynamically allocating ports. | int                                                                                                           |
| effectiveOutboundIPs   | The effective outbound IP resources of the cluster load balancer.                                                                                                                         | [ResourceReference[]](#ResourceReference)                                                                     |
| idleTimeoutInMinutes   | Desired outbound flow idle timeout in minutes. Allowed values are in the range of 4 to 120 (inclusive). The default value is 30 minutes.                                                  | int                                                                                                           |
| managedOutboundIPs     | Desired managed outbound IPs for the cluster load balancer.                                                                                                                               | [ManagedClusterLoadBalancerProfile_ManagedOutboundIPs](#ManagedClusterLoadBalancerProfile_ManagedOutboundIPs) |
| outboundIPPrefixes     | Desired outbound IP Prefix resources for the cluster load balancer.                                                                                                                       | [ManagedClusterLoadBalancerProfile_OutboundIPPrefixes](#ManagedClusterLoadBalancerProfile_OutboundIPPrefixes) |
| outboundIPs            | Desired outbound IP resources for the cluster load balancer.                                                                                                                              | [ManagedClusterLoadBalancerProfile_OutboundIPs](#ManagedClusterLoadBalancerProfile_OutboundIPs)               |

<a id="ManagedClusterLoadBalancerProfile_ARM"></a>ManagedClusterLoadBalancerProfile_ARM
---------------------------------------------------------------------------------------

Profile of the managed cluster load balancer.

Used by: [ContainerServiceNetworkProfile_ARM](#ContainerServiceNetworkProfile_ARM).

| Property               | Description                                                                                                                                                                               | Type                                                                                                                  |
|------------------------|-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|-----------------------------------------------------------------------------------------------------------------------|
| allocatedOutboundPorts | The desired number of allocated SNAT ports per VM. Allowed values are in the range of 0 to 64000 (inclusive). The default value is 0 which results in Azure dynamically allocating ports. | int                                                                                                                   |
| effectiveOutboundIPs   | The effective outbound IP resources of the cluster load balancer.                                                                                                                         | [ResourceReference_ARM[]](#ResourceReference_ARM)                                                                     |
| idleTimeoutInMinutes   | Desired outbound flow idle timeout in minutes. Allowed values are in the range of 4 to 120 (inclusive). The default value is 30 minutes.                                                  | int                                                                                                                   |
| managedOutboundIPs     | Desired managed outbound IPs for the cluster load balancer.                                                                                                                               | [ManagedClusterLoadBalancerProfile_ManagedOutboundIPs_ARM](#ManagedClusterLoadBalancerProfile_ManagedOutboundIPs_ARM) |
| outboundIPPrefixes     | Desired outbound IP Prefix resources for the cluster load balancer.                                                                                                                       | [ManagedClusterLoadBalancerProfile_OutboundIPPrefixes_ARM](#ManagedClusterLoadBalancerProfile_OutboundIPPrefixes_ARM) |
| outboundIPs            | Desired outbound IP resources for the cluster load balancer.                                                                                                                              | [ManagedClusterLoadBalancerProfile_OutboundIPs_ARM](#ManagedClusterLoadBalancerProfile_OutboundIPs_ARM)               |

<a id="ManagedClusterLoadBalancerProfile_ManagedOutboundIPs"></a>ManagedClusterLoadBalancerProfile_ManagedOutboundIPs
---------------------------------------------------------------------------------------------------------------------

Used by: [ManagedClusterLoadBalancerProfile](#ManagedClusterLoadBalancerProfile).

| Property | Description                                                                                                                                                                     | Type |
|----------|---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|------|
| count    | The desired number of outbound IPs created/managed by Azure for the cluster load balancer. Allowed values must be in the range of 1 to 100 (inclusive). The default value is 1. | int  |

<a id="ManagedClusterLoadBalancerProfile_ManagedOutboundIPs_ARM"></a>ManagedClusterLoadBalancerProfile_ManagedOutboundIPs_ARM
-----------------------------------------------------------------------------------------------------------------------------

Used by: [ManagedClusterLoadBalancerProfile_ARM](#ManagedClusterLoadBalancerProfile_ARM).

| Property | Description                                                                                                                                                                     | Type |
|----------|---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|------|
| count    | The desired number of outbound IPs created/managed by Azure for the cluster load balancer. Allowed values must be in the range of 1 to 100 (inclusive). The default value is 1. | int  |

<a id="ManagedClusterLoadBalancerProfile_ManagedOutboundIPs_STATUS"></a>ManagedClusterLoadBalancerProfile_ManagedOutboundIPs_STATUS
-----------------------------------------------------------------------------------------------------------------------------------

Used by: [ManagedClusterLoadBalancerProfile_STATUS](#ManagedClusterLoadBalancerProfile_STATUS).

| Property | Description                                                                                                                                                                     | Type |
|----------|---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|------|
| count    | The desired number of outbound IPs created/managed by Azure for the cluster load balancer. Allowed values must be in the range of 1 to 100 (inclusive). The default value is 1. | int  |

<a id="ManagedClusterLoadBalancerProfile_ManagedOutboundIPs_STATUS_ARM"></a>ManagedClusterLoadBalancerProfile_ManagedOutboundIPs_STATUS_ARM
-------------------------------------------------------------------------------------------------------------------------------------------

Used by: [ManagedClusterLoadBalancerProfile_STATUS_ARM](#ManagedClusterLoadBalancerProfile_STATUS_ARM).

| Property | Description                                                                                                                                                                     | Type |
|----------|---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|------|
| count    | The desired number of outbound IPs created/managed by Azure for the cluster load balancer. Allowed values must be in the range of 1 to 100 (inclusive). The default value is 1. | int  |

<a id="ManagedClusterLoadBalancerProfile_OutboundIPPrefixes"></a>ManagedClusterLoadBalancerProfile_OutboundIPPrefixes
---------------------------------------------------------------------------------------------------------------------

Used by: [ManagedClusterLoadBalancerProfile](#ManagedClusterLoadBalancerProfile).

| Property         | Description                           | Type                                      |
|------------------|---------------------------------------|-------------------------------------------|
| publicIPPrefixes | A list of public IP prefix resources. | [ResourceReference[]](#ResourceReference) |

<a id="ManagedClusterLoadBalancerProfile_OutboundIPPrefixes_ARM"></a>ManagedClusterLoadBalancerProfile_OutboundIPPrefixes_ARM
-----------------------------------------------------------------------------------------------------------------------------

Used by: [ManagedClusterLoadBalancerProfile_ARM](#ManagedClusterLoadBalancerProfile_ARM).

| Property         | Description                           | Type                                              |
|------------------|---------------------------------------|---------------------------------------------------|
| publicIPPrefixes | A list of public IP prefix resources. | [ResourceReference_ARM[]](#ResourceReference_ARM) |

<a id="ManagedClusterLoadBalancerProfile_OutboundIPPrefixes_STATUS"></a>ManagedClusterLoadBalancerProfile_OutboundIPPrefixes_STATUS
-----------------------------------------------------------------------------------------------------------------------------------

Used by: [ManagedClusterLoadBalancerProfile_STATUS](#ManagedClusterLoadBalancerProfile_STATUS).

| Property         | Description                           | Type                                                    |
|------------------|---------------------------------------|---------------------------------------------------------|
| publicIPPrefixes | A list of public IP prefix resources. | [ResourceReference_STATUS[]](#ResourceReference_STATUS) |

<a id="ManagedClusterLoadBalancerProfile_OutboundIPPrefixes_STATUS_ARM"></a>ManagedClusterLoadBalancerProfile_OutboundIPPrefixes_STATUS_ARM
-------------------------------------------------------------------------------------------------------------------------------------------

Used by: [ManagedClusterLoadBalancerProfile_STATUS_ARM](#ManagedClusterLoadBalancerProfile_STATUS_ARM).

| Property         | Description                           | Type                                                            |
|------------------|---------------------------------------|-----------------------------------------------------------------|
| publicIPPrefixes | A list of public IP prefix resources. | [ResourceReference_STATUS_ARM[]](#ResourceReference_STATUS_ARM) |

<a id="ManagedClusterLoadBalancerProfile_OutboundIPs"></a>ManagedClusterLoadBalancerProfile_OutboundIPs
-------------------------------------------------------------------------------------------------------

Used by: [ManagedClusterLoadBalancerProfile](#ManagedClusterLoadBalancerProfile).

| Property  | Description                    | Type                                      |
|-----------|--------------------------------|-------------------------------------------|
| publicIPs | A list of public IP resources. | [ResourceReference[]](#ResourceReference) |

<a id="ManagedClusterLoadBalancerProfile_OutboundIPs_ARM"></a>ManagedClusterLoadBalancerProfile_OutboundIPs_ARM
---------------------------------------------------------------------------------------------------------------

Used by: [ManagedClusterLoadBalancerProfile_ARM](#ManagedClusterLoadBalancerProfile_ARM).

| Property  | Description                    | Type                                              |
|-----------|--------------------------------|---------------------------------------------------|
| publicIPs | A list of public IP resources. | [ResourceReference_ARM[]](#ResourceReference_ARM) |

<a id="ManagedClusterLoadBalancerProfile_OutboundIPs_STATUS"></a>ManagedClusterLoadBalancerProfile_OutboundIPs_STATUS
---------------------------------------------------------------------------------------------------------------------

Used by: [ManagedClusterLoadBalancerProfile_STATUS](#ManagedClusterLoadBalancerProfile_STATUS).

| Property  | Description                    | Type                                                    |
|-----------|--------------------------------|---------------------------------------------------------|
| publicIPs | A list of public IP resources. | [ResourceReference_STATUS[]](#ResourceReference_STATUS) |

<a id="ManagedClusterLoadBalancerProfile_OutboundIPs_STATUS_ARM"></a>ManagedClusterLoadBalancerProfile_OutboundIPs_STATUS_ARM
-----------------------------------------------------------------------------------------------------------------------------

Used by: [ManagedClusterLoadBalancerProfile_STATUS_ARM](#ManagedClusterLoadBalancerProfile_STATUS_ARM).

| Property  | Description                    | Type                                                            |
|-----------|--------------------------------|-----------------------------------------------------------------|
| publicIPs | A list of public IP resources. | [ResourceReference_STATUS_ARM[]](#ResourceReference_STATUS_ARM) |

<a id="ManagedClusterLoadBalancerProfile_STATUS"></a>ManagedClusterLoadBalancerProfile_STATUS
---------------------------------------------------------------------------------------------

Profile of the managed cluster load balancer.

Used by: [ContainerServiceNetworkProfile_STATUS](#ContainerServiceNetworkProfile_STATUS).

| Property               | Description                                                                                                                                                                               | Type                                                                                                                        |
|------------------------|-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|-----------------------------------------------------------------------------------------------------------------------------|
| allocatedOutboundPorts | The desired number of allocated SNAT ports per VM. Allowed values are in the range of 0 to 64000 (inclusive). The default value is 0 which results in Azure dynamically allocating ports. | int                                                                                                                         |
| effectiveOutboundIPs   | The effective outbound IP resources of the cluster load balancer.                                                                                                                         | [ResourceReference_STATUS[]](#ResourceReference_STATUS)                                                                     |
| idleTimeoutInMinutes   | Desired outbound flow idle timeout in minutes. Allowed values are in the range of 4 to 120 (inclusive). The default value is 30 minutes.                                                  | int                                                                                                                         |
| managedOutboundIPs     | Desired managed outbound IPs for the cluster load balancer.                                                                                                                               | [ManagedClusterLoadBalancerProfile_ManagedOutboundIPs_STATUS](#ManagedClusterLoadBalancerProfile_ManagedOutboundIPs_STATUS) |
| outboundIPPrefixes     | Desired outbound IP Prefix resources for the cluster load balancer.                                                                                                                       | [ManagedClusterLoadBalancerProfile_OutboundIPPrefixes_STATUS](#ManagedClusterLoadBalancerProfile_OutboundIPPrefixes_STATUS) |
| outboundIPs            | Desired outbound IP resources for the cluster load balancer.                                                                                                                              | [ManagedClusterLoadBalancerProfile_OutboundIPs_STATUS](#ManagedClusterLoadBalancerProfile_OutboundIPs_STATUS)               |

<a id="ManagedClusterLoadBalancerProfile_STATUS_ARM"></a>ManagedClusterLoadBalancerProfile_STATUS_ARM
-----------------------------------------------------------------------------------------------------

Profile of the managed cluster load balancer.

Used by: [ContainerServiceNetworkProfile_STATUS_ARM](#ContainerServiceNetworkProfile_STATUS_ARM).

| Property               | Description                                                                                                                                                                               | Type                                                                                                                                |
|------------------------|-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|-------------------------------------------------------------------------------------------------------------------------------------|
| allocatedOutboundPorts | The desired number of allocated SNAT ports per VM. Allowed values are in the range of 0 to 64000 (inclusive). The default value is 0 which results in Azure dynamically allocating ports. | int                                                                                                                                 |
| effectiveOutboundIPs   | The effective outbound IP resources of the cluster load balancer.                                                                                                                         | [ResourceReference_STATUS_ARM[]](#ResourceReference_STATUS_ARM)                                                                     |
| idleTimeoutInMinutes   | Desired outbound flow idle timeout in minutes. Allowed values are in the range of 4 to 120 (inclusive). The default value is 30 minutes.                                                  | int                                                                                                                                 |
| managedOutboundIPs     | Desired managed outbound IPs for the cluster load balancer.                                                                                                                               | [ManagedClusterLoadBalancerProfile_ManagedOutboundIPs_STATUS_ARM](#ManagedClusterLoadBalancerProfile_ManagedOutboundIPs_STATUS_ARM) |
| outboundIPPrefixes     | Desired outbound IP Prefix resources for the cluster load balancer.                                                                                                                       | [ManagedClusterLoadBalancerProfile_OutboundIPPrefixes_STATUS_ARM](#ManagedClusterLoadBalancerProfile_OutboundIPPrefixes_STATUS_ARM) |
| outboundIPs            | Desired outbound IP resources for the cluster load balancer.                                                                                                                              | [ManagedClusterLoadBalancerProfile_OutboundIPs_STATUS_ARM](#ManagedClusterLoadBalancerProfile_OutboundIPs_STATUS_ARM)               |

<a id="ManagedClusterOperatorSecrets"></a>ManagedClusterOperatorSecrets
-----------------------------------------------------------------------

Used by: [ManagedClusterOperatorSpec](#ManagedClusterOperatorSpec).

| Property         | Description                                                                                                            | Type                         |
|------------------|------------------------------------------------------------------------------------------------------------------------|------------------------------|
| adminCredentials | indicates where the AdminCredentials secret should be placed. If omitted, the secret will not be retrieved from Azure. | genruntime.SecretDestination |
| userCredentials  | indicates where the UserCredentials secret should be placed. If omitted, the secret will not be retrieved from Azure.  | genruntime.SecretDestination |

<a id="ManagedClusterOperatorSpec"></a>ManagedClusterOperatorSpec
-----------------------------------------------------------------

Details for configuring operator behavior. Fields in this struct are interpreted by the operator directly rather than being passed to Azure

Used by: [ManagedCluster_Spec](#ManagedCluster_Spec).

| Property | Description                                        | Type                                                            |
|----------|----------------------------------------------------|-----------------------------------------------------------------|
| secrets  | configures where to place Azure generated secrets. | [ManagedClusterOperatorSecrets](#ManagedClusterOperatorSecrets) |

<a id="ManagedClusterPodIdentity"></a>ManagedClusterPodIdentity
---------------------------------------------------------------

Details about the pod identity assigned to the Managed Cluster.

Used by: [ManagedClusterPodIdentityProfile](#ManagedClusterPodIdentityProfile).

| Property        | Description                                                        | Type                                                       |
|-----------------|--------------------------------------------------------------------|------------------------------------------------------------|
| bindingSelector | The binding selector to use for the AzureIdentityBinding resource. | string                                                     |
| identity        | The user assigned identity details.                                | [UserAssignedIdentity](#UserAssignedIdentity)<br/>Required |
| name            | The name of the pod identity.                                      | string<br/>Required                                        |
| namespace       | The namespace of the pod identity.                                 | string<br/>Required                                        |

<a id="ManagedClusterPodIdentityException"></a>ManagedClusterPodIdentityException
---------------------------------------------------------------------------------

See [disable AAD Pod Identity for a specific Pod/Application](https://azure.github.io/aad-pod-identity/docs/configure/application_exception/) for more details.

Used by: [ManagedClusterPodIdentityProfile](#ManagedClusterPodIdentityProfile).

| Property  | Description                                  | Type                           |
|-----------|----------------------------------------------|--------------------------------|
| name      | The name of the pod identity exception.      | string<br/>Required            |
| namespace | The namespace of the pod identity exception. | string<br/>Required            |
| podLabels | The pod labels to match.                     | map[string]string<br/>Required |

<a id="ManagedClusterPodIdentityException_ARM"></a>ManagedClusterPodIdentityException_ARM
-----------------------------------------------------------------------------------------

See [disable AAD Pod Identity for a specific Pod/Application](https://azure.github.io/aad-pod-identity/docs/configure/application_exception/) for more details.

Used by: [ManagedClusterPodIdentityProfile_ARM](#ManagedClusterPodIdentityProfile_ARM).

| Property  | Description                                  | Type              |
|-----------|----------------------------------------------|-------------------|
| name      | The name of the pod identity exception.      | string            |
| namespace | The namespace of the pod identity exception. | string            |
| podLabels | The pod labels to match.                     | map[string]string |

<a id="ManagedClusterPodIdentityException_STATUS"></a>ManagedClusterPodIdentityException_STATUS
-----------------------------------------------------------------------------------------------

See [disable AAD Pod Identity for a specific Pod/Application](https://azure.github.io/aad-pod-identity/docs/configure/application_exception/) for more details.

Used by: [ManagedClusterPodIdentityProfile_STATUS](#ManagedClusterPodIdentityProfile_STATUS).

| Property  | Description                                  | Type              |
|-----------|----------------------------------------------|-------------------|
| name      | The name of the pod identity exception.      | string            |
| namespace | The namespace of the pod identity exception. | string            |
| podLabels | The pod labels to match.                     | map[string]string |

<a id="ManagedClusterPodIdentityException_STATUS_ARM"></a>ManagedClusterPodIdentityException_STATUS_ARM
-------------------------------------------------------------------------------------------------------

See [disable AAD Pod Identity for a specific Pod/Application](https://azure.github.io/aad-pod-identity/docs/configure/application_exception/) for more details.

Used by: [ManagedClusterPodIdentityProfile_STATUS_ARM](#ManagedClusterPodIdentityProfile_STATUS_ARM).

| Property  | Description                                  | Type              |
|-----------|----------------------------------------------|-------------------|
| name      | The name of the pod identity exception.      | string            |
| namespace | The namespace of the pod identity exception. | string            |
| podLabels | The pod labels to match.                     | map[string]string |

<a id="ManagedClusterPodIdentityProfile"></a>ManagedClusterPodIdentityProfile
-----------------------------------------------------------------------------

See [use AAD pod identity](https://docs.microsoft.com/azure/aks/use-azure-ad-pod-identity) for more details on pod identity integration.

Used by: [ManagedCluster_Spec](#ManagedCluster_Spec).

| Property                       | Description                                                                                                                                                                                                                                                                                                                                                   | Type                                                                        |
|--------------------------------|---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|-----------------------------------------------------------------------------|
| allowNetworkPluginKubenet      | Running in Kubenet is disabled by default due to the security related nature of AAD Pod Identity and the risks of IP spoofing. See [using Kubenet network plugin with AAD Pod Identity](https://docs.microsoft.com/azure/aks/use-azure-ad-pod-identity#using-kubenet-network-plugin-with-azure-active-directory-pod-managed-identities) for more information. | bool                                                                        |
| enabled                        | Whether the pod identity addon is enabled.                                                                                                                                                                                                                                                                                                                    | bool                                                                        |
| userAssignedIdentities         | The pod identities to use in the cluster.                                                                                                                                                                                                                                                                                                                     | [ManagedClusterPodIdentity[]](#ManagedClusterPodIdentity)                   |
| userAssignedIdentityExceptions | The pod identity exceptions to allow.                                                                                                                                                                                                                                                                                                                         | [ManagedClusterPodIdentityException[]](#ManagedClusterPodIdentityException) |

<a id="ManagedClusterPodIdentityProfile_ARM"></a>ManagedClusterPodIdentityProfile_ARM
-------------------------------------------------------------------------------------

See [use AAD pod identity](https://docs.microsoft.com/azure/aks/use-azure-ad-pod-identity) for more details on pod identity integration.

Used by: [ManagedClusterProperties_ARM](#ManagedClusterProperties_ARM).

| Property                       | Description                                                                                                                                                                                                                                                                                                                                                   | Type                                                                                |
|--------------------------------|---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|-------------------------------------------------------------------------------------|
| allowNetworkPluginKubenet      | Running in Kubenet is disabled by default due to the security related nature of AAD Pod Identity and the risks of IP spoofing. See [using Kubenet network plugin with AAD Pod Identity](https://docs.microsoft.com/azure/aks/use-azure-ad-pod-identity#using-kubenet-network-plugin-with-azure-active-directory-pod-managed-identities) for more information. | bool                                                                                |
| enabled                        | Whether the pod identity addon is enabled.                                                                                                                                                                                                                                                                                                                    | bool                                                                                |
| userAssignedIdentities         | The pod identities to use in the cluster.                                                                                                                                                                                                                                                                                                                     | [ManagedClusterPodIdentity_ARM[]](#ManagedClusterPodIdentity_ARM)                   |
| userAssignedIdentityExceptions | The pod identity exceptions to allow.                                                                                                                                                                                                                                                                                                                         | [ManagedClusterPodIdentityException_ARM[]](#ManagedClusterPodIdentityException_ARM) |

<a id="ManagedClusterPodIdentityProfile_STATUS"></a>ManagedClusterPodIdentityProfile_STATUS
-------------------------------------------------------------------------------------------

See [use AAD pod identity](https://docs.microsoft.com/azure/aks/use-azure-ad-pod-identity) for more details on pod identity integration.

Used by: [ManagedCluster_STATUS](#ManagedCluster_STATUS).

| Property                       | Description                                                                                                                                                                                                                                                                                                                                                   | Type                                                                                      |
|--------------------------------|---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|-------------------------------------------------------------------------------------------|
| allowNetworkPluginKubenet      | Running in Kubenet is disabled by default due to the security related nature of AAD Pod Identity and the risks of IP spoofing. See [using Kubenet network plugin with AAD Pod Identity](https://docs.microsoft.com/azure/aks/use-azure-ad-pod-identity#using-kubenet-network-plugin-with-azure-active-directory-pod-managed-identities) for more information. | bool                                                                                      |
| enabled                        | Whether the pod identity addon is enabled.                                                                                                                                                                                                                                                                                                                    | bool                                                                                      |
| userAssignedIdentities         | The pod identities to use in the cluster.                                                                                                                                                                                                                                                                                                                     | [ManagedClusterPodIdentity_STATUS[]](#ManagedClusterPodIdentity_STATUS)                   |
| userAssignedIdentityExceptions | The pod identity exceptions to allow.                                                                                                                                                                                                                                                                                                                         | [ManagedClusterPodIdentityException_STATUS[]](#ManagedClusterPodIdentityException_STATUS) |

<a id="ManagedClusterPodIdentityProfile_STATUS_ARM"></a>ManagedClusterPodIdentityProfile_STATUS_ARM
---------------------------------------------------------------------------------------------------

See [use AAD pod identity](https://docs.microsoft.com/azure/aks/use-azure-ad-pod-identity) for more details on pod identity integration.

Used by: [ManagedClusterProperties_STATUS_ARM](#ManagedClusterProperties_STATUS_ARM).

| Property                       | Description                                                                                                                                                                                                                                                                                                                                                   | Type                                                                                              |
|--------------------------------|---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|---------------------------------------------------------------------------------------------------|
| allowNetworkPluginKubenet      | Running in Kubenet is disabled by default due to the security related nature of AAD Pod Identity and the risks of IP spoofing. See [using Kubenet network plugin with AAD Pod Identity](https://docs.microsoft.com/azure/aks/use-azure-ad-pod-identity#using-kubenet-network-plugin-with-azure-active-directory-pod-managed-identities) for more information. | bool                                                                                              |
| enabled                        | Whether the pod identity addon is enabled.                                                                                                                                                                                                                                                                                                                    | bool                                                                                              |
| userAssignedIdentities         | The pod identities to use in the cluster.                                                                                                                                                                                                                                                                                                                     | [ManagedClusterPodIdentity_STATUS_ARM[]](#ManagedClusterPodIdentity_STATUS_ARM)                   |
| userAssignedIdentityExceptions | The pod identity exceptions to allow.                                                                                                                                                                                                                                                                                                                         | [ManagedClusterPodIdentityException_STATUS_ARM[]](#ManagedClusterPodIdentityException_STATUS_ARM) |

<a id="ManagedClusterPodIdentityProvisioningErrorBody_STATUS"></a>ManagedClusterPodIdentityProvisioningErrorBody_STATUS
-----------------------------------------------------------------------------------------------------------------------

An error response from the pod identity provisioning.

Used by: [ManagedClusterPodIdentityProvisioningError_STATUS](#ManagedClusterPodIdentityProvisioningError_STATUS).

| Property | Description                                                                                        | Type                                                                                                                                |
|----------|----------------------------------------------------------------------------------------------------|-------------------------------------------------------------------------------------------------------------------------------------|
| code     | An identifier for the error. Codes are invariant and are intended to be consumed programmatically. | string                                                                                                                              |
| details  | A list of additional details about the error.                                                      | [ManagedClusterPodIdentityProvisioningErrorBody_STATUS_Unrolled[]](#ManagedClusterPodIdentityProvisioningErrorBody_STATUS_Unrolled) |
| message  | A message describing the error, intended to be suitable for display in a user interface.           | string                                                                                                                              |
| target   | The target of the particular error. For example, the name of the property in error.                | string                                                                                                                              |

<a id="ManagedClusterPodIdentityProvisioningErrorBody_STATUS_ARM"></a>ManagedClusterPodIdentityProvisioningErrorBody_STATUS_ARM
-------------------------------------------------------------------------------------------------------------------------------

An error response from the pod identity provisioning.

Used by: [ManagedClusterPodIdentityProvisioningError_STATUS_ARM](#ManagedClusterPodIdentityProvisioningError_STATUS_ARM).

| Property | Description                                                                                        | Type                                                                                                                                        |
|----------|----------------------------------------------------------------------------------------------------|---------------------------------------------------------------------------------------------------------------------------------------------|
| code     | An identifier for the error. Codes are invariant and are intended to be consumed programmatically. | string                                                                                                                                      |
| details  | A list of additional details about the error.                                                      | [ManagedClusterPodIdentityProvisioningErrorBody_STATUS_Unrolled_ARM[]](#ManagedClusterPodIdentityProvisioningErrorBody_STATUS_Unrolled_ARM) |
| message  | A message describing the error, intended to be suitable for display in a user interface.           | string                                                                                                                                      |
| target   | The target of the particular error. For example, the name of the property in error.                | string                                                                                                                                      |

<a id="ManagedClusterPodIdentityProvisioningErrorBody_STATUS_Unrolled"></a>ManagedClusterPodIdentityProvisioningErrorBody_STATUS_Unrolled
-----------------------------------------------------------------------------------------------------------------------------------------

Used by: [ManagedClusterPodIdentityProvisioningErrorBody_STATUS](#ManagedClusterPodIdentityProvisioningErrorBody_STATUS).

| Property | Description                                                                                        | Type   |
|----------|----------------------------------------------------------------------------------------------------|--------|
| code     | An identifier for the error. Codes are invariant and are intended to be consumed programmatically. | string |
| message  | A message describing the error, intended to be suitable for display in a user interface.           | string |
| target   | The target of the particular error. For example, the name of the property in error.                | string |

<a id="ManagedClusterPodIdentityProvisioningErrorBody_STATUS_Unrolled_ARM"></a>ManagedClusterPodIdentityProvisioningErrorBody_STATUS_Unrolled_ARM
-------------------------------------------------------------------------------------------------------------------------------------------------

Used by: [ManagedClusterPodIdentityProvisioningErrorBody_STATUS_ARM](#ManagedClusterPodIdentityProvisioningErrorBody_STATUS_ARM).

| Property | Description                                                                                        | Type   |
|----------|----------------------------------------------------------------------------------------------------|--------|
| code     | An identifier for the error. Codes are invariant and are intended to be consumed programmatically. | string |
| message  | A message describing the error, intended to be suitable for display in a user interface.           | string |
| target   | The target of the particular error. For example, the name of the property in error.                | string |

<a id="ManagedClusterPodIdentityProvisioningError_STATUS"></a>ManagedClusterPodIdentityProvisioningError_STATUS
---------------------------------------------------------------------------------------------------------------

An error response from the pod identity provisioning.

Used by: [ManagedClusterPodIdentity_ProvisioningInfo_STATUS](#ManagedClusterPodIdentity_ProvisioningInfo_STATUS).

| Property | Description              | Type                                                                                                            |
|----------|--------------------------|-----------------------------------------------------------------------------------------------------------------|
| error    | Details about the error. | [ManagedClusterPodIdentityProvisioningErrorBody_STATUS](#ManagedClusterPodIdentityProvisioningErrorBody_STATUS) |

<a id="ManagedClusterPodIdentityProvisioningError_STATUS_ARM"></a>ManagedClusterPodIdentityProvisioningError_STATUS_ARM
-----------------------------------------------------------------------------------------------------------------------

An error response from the pod identity provisioning.

Used by: [ManagedClusterPodIdentity_ProvisioningInfo_STATUS_ARM](#ManagedClusterPodIdentity_ProvisioningInfo_STATUS_ARM).

| Property | Description              | Type                                                                                                                    |
|----------|--------------------------|-------------------------------------------------------------------------------------------------------------------------|
| error    | Details about the error. | [ManagedClusterPodIdentityProvisioningErrorBody_STATUS_ARM](#ManagedClusterPodIdentityProvisioningErrorBody_STATUS_ARM) |

<a id="ManagedClusterPodIdentity_ARM"></a>ManagedClusterPodIdentity_ARM
-----------------------------------------------------------------------

Details about the pod identity assigned to the Managed Cluster.

Used by: [ManagedClusterPodIdentityProfile_ARM](#ManagedClusterPodIdentityProfile_ARM).

| Property        | Description                                                        | Type                                                  |
|-----------------|--------------------------------------------------------------------|-------------------------------------------------------|
| bindingSelector | The binding selector to use for the AzureIdentityBinding resource. | string                                                |
| identity        | The user assigned identity details.                                | [UserAssignedIdentity_ARM](#UserAssignedIdentity_ARM) |
| name            | The name of the pod identity.                                      | string                                                |
| namespace       | The namespace of the pod identity.                                 | string                                                |

<a id="ManagedClusterPodIdentity_ProvisioningInfo_STATUS"></a>ManagedClusterPodIdentity_ProvisioningInfo_STATUS
---------------------------------------------------------------------------------------------------------------

Used by: [ManagedClusterPodIdentity_STATUS](#ManagedClusterPodIdentity_STATUS).

| Property | Description                             | Type                                                                                                    |
|----------|-----------------------------------------|---------------------------------------------------------------------------------------------------------|
| error    | Pod identity assignment error (if any). | [ManagedClusterPodIdentityProvisioningError_STATUS](#ManagedClusterPodIdentityProvisioningError_STATUS) |

<a id="ManagedClusterPodIdentity_ProvisioningInfo_STATUS_ARM"></a>ManagedClusterPodIdentity_ProvisioningInfo_STATUS_ARM
-----------------------------------------------------------------------------------------------------------------------

Used by: [ManagedClusterPodIdentity_STATUS_ARM](#ManagedClusterPodIdentity_STATUS_ARM).

| Property | Description                             | Type                                                                                                            |
|----------|-----------------------------------------|-----------------------------------------------------------------------------------------------------------------|
| error    | Pod identity assignment error (if any). | [ManagedClusterPodIdentityProvisioningError_STATUS_ARM](#ManagedClusterPodIdentityProvisioningError_STATUS_ARM) |

<a id="ManagedClusterPodIdentity_ProvisioningState_STATUS"></a>ManagedClusterPodIdentity_ProvisioningState_STATUS
-----------------------------------------------------------------------------------------------------------------

Used by: [ManagedClusterPodIdentity_STATUS](#ManagedClusterPodIdentity_STATUS), and [ManagedClusterPodIdentity_STATUS_ARM](#ManagedClusterPodIdentity_STATUS_ARM).

| Value      | Description |
|------------|-------------|
| "Assigned" |             |
| "Deleting" |             |
| "Failed"   |             |
| "Updating" |             |

<a id="ManagedClusterPodIdentity_STATUS"></a>ManagedClusterPodIdentity_STATUS
-----------------------------------------------------------------------------

Details about the pod identity assigned to the Managed Cluster.

Used by: [ManagedClusterPodIdentityProfile_STATUS](#ManagedClusterPodIdentityProfile_STATUS).

| Property          | Description                                                        | Type                                                                                                      |
|-------------------|--------------------------------------------------------------------|-----------------------------------------------------------------------------------------------------------|
| bindingSelector   | The binding selector to use for the AzureIdentityBinding resource. | string                                                                                                    |
| identity          | The user assigned identity details.                                | [UserAssignedIdentity_STATUS](#UserAssignedIdentity_STATUS)                                               |
| name              | The name of the pod identity.                                      | string                                                                                                    |
| namespace         | The namespace of the pod identity.                                 | string                                                                                                    |
| provisioningInfo  |                                                                    | [ManagedClusterPodIdentity_ProvisioningInfo_STATUS](#ManagedClusterPodIdentity_ProvisioningInfo_STATUS)   |
| provisioningState | The current provisioning state of the pod identity.                | [ManagedClusterPodIdentity_ProvisioningState_STATUS](#ManagedClusterPodIdentity_ProvisioningState_STATUS) |

<a id="ManagedClusterPodIdentity_STATUS_ARM"></a>ManagedClusterPodIdentity_STATUS_ARM
-------------------------------------------------------------------------------------

Details about the pod identity assigned to the Managed Cluster.

Used by: [ManagedClusterPodIdentityProfile_STATUS_ARM](#ManagedClusterPodIdentityProfile_STATUS_ARM).

| Property          | Description                                                        | Type                                                                                                            |
|-------------------|--------------------------------------------------------------------|-----------------------------------------------------------------------------------------------------------------|
| bindingSelector   | The binding selector to use for the AzureIdentityBinding resource. | string                                                                                                          |
| identity          | The user assigned identity details.                                | [UserAssignedIdentity_STATUS_ARM](#UserAssignedIdentity_STATUS_ARM)                                             |
| name              | The name of the pod identity.                                      | string                                                                                                          |
| namespace         | The namespace of the pod identity.                                 | string                                                                                                          |
| provisioningInfo  |                                                                    | [ManagedClusterPodIdentity_ProvisioningInfo_STATUS_ARM](#ManagedClusterPodIdentity_ProvisioningInfo_STATUS_ARM) |
| provisioningState | The current provisioning state of the pod identity.                | [ManagedClusterPodIdentity_ProvisioningState_STATUS](#ManagedClusterPodIdentity_ProvisioningState_STATUS)       |

<a id="ManagedClusterProperties_ARM"></a>ManagedClusterProperties_ARM
---------------------------------------------------------------------

Properties of the managed cluster.

Used by: [ManagedCluster_Spec_ARM](#ManagedCluster_Spec_ARM).

| Property                | Description                                                                                                                                                                                                                                                                                                                                                                                 | Type                                                                                              |
|-------------------------|---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|---------------------------------------------------------------------------------------------------|
| aadProfile              | The Azure Active Directory configuration.                                                                                                                                                                                                                                                                                                                                                   | [ManagedClusterAADProfile_ARM](#ManagedClusterAADProfile_ARM)                                     |
| addonProfiles           | The profile of managed cluster add-on.                                                                                                                                                                                                                                                                                                                                                      | [map[string]ManagedClusterAddonProfile_ARM](#ManagedClusterAddonProfile_ARM)                      |
| agentPoolProfiles       | The agent pool properties.                                                                                                                                                                                                                                                                                                                                                                  | [ManagedClusterAgentPoolProfile_ARM[]](#ManagedClusterAgentPoolProfile_ARM)                       |
| apiServerAccessProfile  | The access profile for managed cluster API server.                                                                                                                                                                                                                                                                                                                                          | [ManagedClusterAPIServerAccessProfile_ARM](#ManagedClusterAPIServerAccessProfile_ARM)             |
| autoScalerProfile       | Parameters to be applied to the cluster-autoscaler when enabled                                                                                                                                                                                                                                                                                                                             | [ManagedClusterProperties_AutoScalerProfile_ARM](#ManagedClusterProperties_AutoScalerProfile_ARM) |
| autoUpgradeProfile      | The auto upgrade configuration.                                                                                                                                                                                                                                                                                                                                                             | [ManagedClusterAutoUpgradeProfile_ARM](#ManagedClusterAutoUpgradeProfile_ARM)                     |
| disableLocalAccounts    | If set to true, getting static credentials will be disabled for this cluster. This must only be used on Managed Clusters that are AAD enabled. For more details see [disable local accounts](https://docs.microsoft.com/azure/aks/managed-aad#disable-local-accounts-preview).                                                                                                              | bool                                                                                              |
| diskEncryptionSetID     |                                                                                                                                                                                                                                                                                                                                                                                             | string                                                                                            |
| dnsPrefix               | This cannot be updated once the Managed Cluster has been created.                                                                                                                                                                                                                                                                                                                           | string                                                                                            |
| enablePodSecurityPolicy | (DEPRECATING) Whether to enable Kubernetes pod security policy (preview). This feature is set for removal on October 15th, 2020. Learn more at aka.ms/aks/azpodpolicy.                                                                                                                                                                                                                      | bool                                                                                              |
| enableRBAC              | Whether to enable Kubernetes Role-Based Access Control.                                                                                                                                                                                                                                                                                                                                     | bool                                                                                              |
| fqdnSubdomain           | This cannot be updated once the Managed Cluster has been created.                                                                                                                                                                                                                                                                                                                           | string                                                                                            |
| httpProxyConfig         | Configurations for provisioning the cluster with HTTP proxy servers.                                                                                                                                                                                                                                                                                                                        | [ManagedClusterHTTPProxyConfig_ARM](#ManagedClusterHTTPProxyConfig_ARM)                           |
| identityProfile         | Identities associated with the cluster.                                                                                                                                                                                                                                                                                                                                                     | [map[string]UserAssignedIdentity_ARM](#UserAssignedIdentity_ARM)                                  |
| kubernetesVersion       | When you upgrade a supported AKS cluster, Kubernetes minor versions cannot be skipped. All upgrades must be performed sequentially by major version number. For example, upgrades between 1.14.x -> 1.15.x or 1.15.x -> 1.16.x are allowed, however 1.14.x -> 1.16.x is not allowed. See [upgrading an AKS cluster](https://docs.microsoft.com/azure/aks/upgrade-cluster) for more details. | string                                                                                            |
| linuxProfile            | The profile for Linux VMs in the Managed Cluster.                                                                                                                                                                                                                                                                                                                                           | [ContainerServiceLinuxProfile_ARM](#ContainerServiceLinuxProfile_ARM)                             |
| networkProfile          | The network configuration profile.                                                                                                                                                                                                                                                                                                                                                          | [ContainerServiceNetworkProfile_ARM](#ContainerServiceNetworkProfile_ARM)                         |
| nodeResourceGroup       | The name of the resource group containing agent pool nodes.                                                                                                                                                                                                                                                                                                                                 | string                                                                                            |
| podIdentityProfile      | See [use AAD pod identity](https://docs.microsoft.com/azure/aks/use-azure-ad-pod-identity) for more details on AAD pod identity integration.                                                                                                                                                                                                                                                | [ManagedClusterPodIdentityProfile_ARM](#ManagedClusterPodIdentityProfile_ARM)                     |
| privateLinkResources    | Private link resources associated with the cluster.                                                                                                                                                                                                                                                                                                                                         | [PrivateLinkResource_ARM[]](#PrivateLinkResource_ARM)                                             |
| servicePrincipalProfile | Information about a service principal identity for the cluster to use for manipulating Azure APIs.                                                                                                                                                                                                                                                                                          | [ManagedClusterServicePrincipalProfile_ARM](#ManagedClusterServicePrincipalProfile_ARM)           |
| windowsProfile          | The profile for Windows VMs in the Managed Cluster.                                                                                                                                                                                                                                                                                                                                         | [ManagedClusterWindowsProfile_ARM](#ManagedClusterWindowsProfile_ARM)                             |

<a id="ManagedClusterProperties_AutoScalerProfile"></a>ManagedClusterProperties_AutoScalerProfile
-------------------------------------------------------------------------------------------------

Used by: [ManagedCluster_Spec](#ManagedCluster_Spec).

| Property                         | Description                                                                                                                                                                                                                                                                                                                         | Type                                                                                                        |
|----------------------------------|-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|-------------------------------------------------------------------------------------------------------------|
| balance-similar-node-groups      | Valid values are 'true' and 'false'                                                                                                                                                                                                                                                                                                 | string                                                                                                      |
| expander                         | If not specified, the default is 'random'. See [expanders](https://github.com/kubernetes/autoscaler/blob/master/cluster-autoscaler/FAQ.md#what-are-expanders) for more information.                                                                                                                                                 | [ManagedClusterProperties_AutoScalerProfile_Expander](#ManagedClusterProperties_AutoScalerProfile_Expander) |
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

<a id="ManagedClusterProperties_AutoScalerProfile_ARM"></a>ManagedClusterProperties_AutoScalerProfile_ARM
---------------------------------------------------------------------------------------------------------

Used by: [ManagedClusterProperties_ARM](#ManagedClusterProperties_ARM).

| Property                         | Description                                                                                                                                                                                                                                                                                                                         | Type                                                                                                        |
|----------------------------------|-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|-------------------------------------------------------------------------------------------------------------|
| balance-similar-node-groups      | Valid values are 'true' and 'false'                                                                                                                                                                                                                                                                                                 | string                                                                                                      |
| expander                         | If not specified, the default is 'random'. See [expanders](https://github.com/kubernetes/autoscaler/blob/master/cluster-autoscaler/FAQ.md#what-are-expanders) for more information.                                                                                                                                                 | [ManagedClusterProperties_AutoScalerProfile_Expander](#ManagedClusterProperties_AutoScalerProfile_Expander) |
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

<a id="ManagedClusterProperties_AutoScalerProfile_Expander"></a>ManagedClusterProperties_AutoScalerProfile_Expander
-------------------------------------------------------------------------------------------------------------------

Used by: [ManagedClusterProperties_AutoScalerProfile](#ManagedClusterProperties_AutoScalerProfile), and [ManagedClusterProperties_AutoScalerProfile_ARM](#ManagedClusterProperties_AutoScalerProfile_ARM).

| Value         | Description |
|---------------|-------------|
| "least-waste" |             |
| "most-pods"   |             |
| "priority"    |             |
| "random"      |             |

<a id="ManagedClusterProperties_AutoScalerProfile_Expander_STATUS"></a>ManagedClusterProperties_AutoScalerProfile_Expander_STATUS
---------------------------------------------------------------------------------------------------------------------------------

Used by: [ManagedClusterProperties_AutoScalerProfile_STATUS](#ManagedClusterProperties_AutoScalerProfile_STATUS), and [ManagedClusterProperties_AutoScalerProfile_STATUS_ARM](#ManagedClusterProperties_AutoScalerProfile_STATUS_ARM).

| Value         | Description |
|---------------|-------------|
| "least-waste" |             |
| "most-pods"   |             |
| "priority"    |             |
| "random"      |             |

<a id="ManagedClusterProperties_AutoScalerProfile_STATUS"></a>ManagedClusterProperties_AutoScalerProfile_STATUS
---------------------------------------------------------------------------------------------------------------

Used by: [ManagedCluster_STATUS](#ManagedCluster_STATUS).

| Property                         | Description                                                                                                                                                                                                                                                                                                                         | Type                                                                                                                      |
|----------------------------------|-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|---------------------------------------------------------------------------------------------------------------------------|
| balance-similar-node-groups      | Valid values are 'true' and 'false'                                                                                                                                                                                                                                                                                                 | string                                                                                                                    |
| expander                         | If not specified, the default is 'random'. See [expanders](https://github.com/kubernetes/autoscaler/blob/master/cluster-autoscaler/FAQ.md#what-are-expanders) for more information.                                                                                                                                                 | [ManagedClusterProperties_AutoScalerProfile_Expander_STATUS](#ManagedClusterProperties_AutoScalerProfile_Expander_STATUS) |
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

<a id="ManagedClusterProperties_AutoScalerProfile_STATUS_ARM"></a>ManagedClusterProperties_AutoScalerProfile_STATUS_ARM
-----------------------------------------------------------------------------------------------------------------------

Used by: [ManagedClusterProperties_STATUS_ARM](#ManagedClusterProperties_STATUS_ARM).

| Property                         | Description                                                                                                                                                                                                                                                                                                                         | Type                                                                                                                      |
|----------------------------------|-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|---------------------------------------------------------------------------------------------------------------------------|
| balance-similar-node-groups      | Valid values are 'true' and 'false'                                                                                                                                                                                                                                                                                                 | string                                                                                                                    |
| expander                         | If not specified, the default is 'random'. See [expanders](https://github.com/kubernetes/autoscaler/blob/master/cluster-autoscaler/FAQ.md#what-are-expanders) for more information.                                                                                                                                                 | [ManagedClusterProperties_AutoScalerProfile_Expander_STATUS](#ManagedClusterProperties_AutoScalerProfile_Expander_STATUS) |
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

<a id="ManagedClusterProperties_STATUS_ARM"></a>ManagedClusterProperties_STATUS_ARM
-----------------------------------------------------------------------------------

Properties of the managed cluster.

Used by: [ManagedCluster_STATUS_ARM](#ManagedCluster_STATUS_ARM).

| Property                | Description                                                                                                                                                                                                                                                                                                                                                                                 | Type                                                                                                            |
|-------------------------|---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|-----------------------------------------------------------------------------------------------------------------|
| aadProfile              | The Azure Active Directory configuration.                                                                                                                                                                                                                                                                                                                                                   | [ManagedClusterAADProfile_STATUS_ARM](#ManagedClusterAADProfile_STATUS_ARM)                                     |
| addonProfiles           | The profile of managed cluster add-on.                                                                                                                                                                                                                                                                                                                                                      | [map[string]ManagedClusterAddonProfile_STATUS_ARM](#ManagedClusterAddonProfile_STATUS_ARM)                      |
| agentPoolProfiles       | The agent pool properties.                                                                                                                                                                                                                                                                                                                                                                  | [ManagedClusterAgentPoolProfile_STATUS_ARM[]](#ManagedClusterAgentPoolProfile_STATUS_ARM)                       |
| apiServerAccessProfile  | The access profile for managed cluster API server.                                                                                                                                                                                                                                                                                                                                          | [ManagedClusterAPIServerAccessProfile_STATUS_ARM](#ManagedClusterAPIServerAccessProfile_STATUS_ARM)             |
| autoScalerProfile       | Parameters to be applied to the cluster-autoscaler when enabled                                                                                                                                                                                                                                                                                                                             | [ManagedClusterProperties_AutoScalerProfile_STATUS_ARM](#ManagedClusterProperties_AutoScalerProfile_STATUS_ARM) |
| autoUpgradeProfile      | The auto upgrade configuration.                                                                                                                                                                                                                                                                                                                                                             | [ManagedClusterAutoUpgradeProfile_STATUS_ARM](#ManagedClusterAutoUpgradeProfile_STATUS_ARM)                     |
| azurePortalFQDN         | The Azure Portal requires certain Cross-Origin Resource Sharing (CORS) headers to be sent in some responses, which Kubernetes APIServer doesn't handle by default. This special FQDN supports CORS, allowing the Azure Portal to function properly.                                                                                                                                         | string                                                                                                          |
| disableLocalAccounts    | If set to true, getting static credentials will be disabled for this cluster. This must only be used on Managed Clusters that are AAD enabled. For more details see [disable local accounts](https://docs.microsoft.com/azure/aks/managed-aad#disable-local-accounts-preview).                                                                                                              | bool                                                                                                            |
| diskEncryptionSetID     | This is of the form: '/&ZeroWidthSpace;subscriptions/&ZeroWidthSpace;{subscriptionId}/&ZeroWidthSpace;resourceGroups/&ZeroWidthSpace;{resourceGroupName}/&ZeroWidthSpace;providers/&ZeroWidthSpace;Microsoft.Compute/diskEncryptionSets/{encryptionSetName}'                                                                                                                                | string                                                                                                          |
| dnsPrefix               | This cannot be updated once the Managed Cluster has been created.                                                                                                                                                                                                                                                                                                                           | string                                                                                                          |
| enablePodSecurityPolicy | (DEPRECATING) Whether to enable Kubernetes pod security policy (preview). This feature is set for removal on October 15th, 2020. Learn more at aka.ms/aks/azpodpolicy.                                                                                                                                                                                                                      | bool                                                                                                            |
| enableRBAC              | Whether to enable Kubernetes Role-Based Access Control.                                                                                                                                                                                                                                                                                                                                     | bool                                                                                                            |
| fqdn                    | The FQDN of the master pool.                                                                                                                                                                                                                                                                                                                                                                | string                                                                                                          |
| fqdnSubdomain           | This cannot be updated once the Managed Cluster has been created.                                                                                                                                                                                                                                                                                                                           | string                                                                                                          |
| httpProxyConfig         | Configurations for provisioning the cluster with HTTP proxy servers.                                                                                                                                                                                                                                                                                                                        | [ManagedClusterHTTPProxyConfig_STATUS_ARM](#ManagedClusterHTTPProxyConfig_STATUS_ARM)                           |
| identityProfile         | Identities associated with the cluster.                                                                                                                                                                                                                                                                                                                                                     | [map[string]UserAssignedIdentity_STATUS_ARM](#UserAssignedIdentity_STATUS_ARM)                                  |
| kubernetesVersion       | When you upgrade a supported AKS cluster, Kubernetes minor versions cannot be skipped. All upgrades must be performed sequentially by major version number. For example, upgrades between 1.14.x -> 1.15.x or 1.15.x -> 1.16.x are allowed, however 1.14.x -> 1.16.x is not allowed. See [upgrading an AKS cluster](https://docs.microsoft.com/azure/aks/upgrade-cluster) for more details. | string                                                                                                          |
| linuxProfile            | The profile for Linux VMs in the Managed Cluster.                                                                                                                                                                                                                                                                                                                                           | [ContainerServiceLinuxProfile_STATUS_ARM](#ContainerServiceLinuxProfile_STATUS_ARM)                             |
| maxAgentPools           | The max number of agent pools for the managed cluster.                                                                                                                                                                                                                                                                                                                                      | int                                                                                                             |
| networkProfile          | The network configuration profile.                                                                                                                                                                                                                                                                                                                                                          | [ContainerServiceNetworkProfile_STATUS_ARM](#ContainerServiceNetworkProfile_STATUS_ARM)                         |
| nodeResourceGroup       | The name of the resource group containing agent pool nodes.                                                                                                                                                                                                                                                                                                                                 | string                                                                                                          |
| podIdentityProfile      | See [use AAD pod identity](https://docs.microsoft.com/azure/aks/use-azure-ad-pod-identity) for more details on AAD pod identity integration.                                                                                                                                                                                                                                                | [ManagedClusterPodIdentityProfile_STATUS_ARM](#ManagedClusterPodIdentityProfile_STATUS_ARM)                     |
| powerState              | The Power State of the cluster.                                                                                                                                                                                                                                                                                                                                                             | [PowerState_STATUS_ARM](#PowerState_STATUS_ARM)                                                                 |
| privateFQDN             | The FQDN of private cluster.                                                                                                                                                                                                                                                                                                                                                                | string                                                                                                          |
| privateLinkResources    | Private link resources associated with the cluster.                                                                                                                                                                                                                                                                                                                                         | [PrivateLinkResource_STATUS_ARM[]](#PrivateLinkResource_STATUS_ARM)                                             |
| provisioningState       | The current provisioning state.                                                                                                                                                                                                                                                                                                                                                             | string                                                                                                          |
| servicePrincipalProfile | Information about a service principal identity for the cluster to use for manipulating Azure APIs.                                                                                                                                                                                                                                                                                          | [ManagedClusterServicePrincipalProfile_STATUS_ARM](#ManagedClusterServicePrincipalProfile_STATUS_ARM)           |
| windowsProfile          | The profile for Windows VMs in the Managed Cluster.                                                                                                                                                                                                                                                                                                                                         | [ManagedClusterWindowsProfile_STATUS_ARM](#ManagedClusterWindowsProfile_STATUS_ARM)                             |

<a id="ManagedClusterSKU"></a>ManagedClusterSKU
-----------------------------------------------

The SKU of a Managed Cluster.

Used by: [ManagedCluster_Spec](#ManagedCluster_Spec).

| Property | Description                                                                                                                  | Type                                              |
|----------|------------------------------------------------------------------------------------------------------------------------------|---------------------------------------------------|
| name     | The name of a managed cluster SKU.                                                                                           | [ManagedClusterSKU_Name](#ManagedClusterSKU_Name) |
| tier     | If not specified, the default is 'Free'. See [uptime SLA](https://docs.microsoft.com/azure/aks/uptime-sla) for more details. | [ManagedClusterSKU_Tier](#ManagedClusterSKU_Tier) |

<a id="ManagedClusterSKU_ARM"></a>ManagedClusterSKU_ARM
-------------------------------------------------------

The SKU of a Managed Cluster.

Used by: [ManagedCluster_Spec_ARM](#ManagedCluster_Spec_ARM).

| Property | Description                                                                                                                  | Type                                              |
|----------|------------------------------------------------------------------------------------------------------------------------------|---------------------------------------------------|
| name     | The name of a managed cluster SKU.                                                                                           | [ManagedClusterSKU_Name](#ManagedClusterSKU_Name) |
| tier     | If not specified, the default is 'Free'. See [uptime SLA](https://docs.microsoft.com/azure/aks/uptime-sla) for more details. | [ManagedClusterSKU_Tier](#ManagedClusterSKU_Tier) |

<a id="ManagedClusterSKU_Name"></a>ManagedClusterSKU_Name
---------------------------------------------------------

Used by: [ManagedClusterSKU](#ManagedClusterSKU), and [ManagedClusterSKU_ARM](#ManagedClusterSKU_ARM).

| Value   | Description |
|---------|-------------|
| "Basic" |             |

<a id="ManagedClusterSKU_Name_STATUS"></a>ManagedClusterSKU_Name_STATUS
-----------------------------------------------------------------------

Used by: [ManagedClusterSKU_STATUS](#ManagedClusterSKU_STATUS), and [ManagedClusterSKU_STATUS_ARM](#ManagedClusterSKU_STATUS_ARM).

| Value   | Description |
|---------|-------------|
| "Basic" |             |

<a id="ManagedClusterSKU_STATUS"></a>ManagedClusterSKU_STATUS
-------------------------------------------------------------

The SKU of a Managed Cluster.

Used by: [ManagedCluster_STATUS](#ManagedCluster_STATUS).

| Property | Description                                                                                                                  | Type                                                            |
|----------|------------------------------------------------------------------------------------------------------------------------------|-----------------------------------------------------------------|
| name     | The name of a managed cluster SKU.                                                                                           | [ManagedClusterSKU_Name_STATUS](#ManagedClusterSKU_Name_STATUS) |
| tier     | If not specified, the default is 'Free'. See [uptime SLA](https://docs.microsoft.com/azure/aks/uptime-sla) for more details. | [ManagedClusterSKU_Tier_STATUS](#ManagedClusterSKU_Tier_STATUS) |

<a id="ManagedClusterSKU_STATUS_ARM"></a>ManagedClusterSKU_STATUS_ARM
---------------------------------------------------------------------

The SKU of a Managed Cluster.

Used by: [ManagedCluster_STATUS_ARM](#ManagedCluster_STATUS_ARM).

| Property | Description                                                                                                                  | Type                                                            |
|----------|------------------------------------------------------------------------------------------------------------------------------|-----------------------------------------------------------------|
| name     | The name of a managed cluster SKU.                                                                                           | [ManagedClusterSKU_Name_STATUS](#ManagedClusterSKU_Name_STATUS) |
| tier     | If not specified, the default is 'Free'. See [uptime SLA](https://docs.microsoft.com/azure/aks/uptime-sla) for more details. | [ManagedClusterSKU_Tier_STATUS](#ManagedClusterSKU_Tier_STATUS) |

<a id="ManagedClusterSKU_Tier"></a>ManagedClusterSKU_Tier
---------------------------------------------------------

Used by: [ManagedClusterSKU](#ManagedClusterSKU), and [ManagedClusterSKU_ARM](#ManagedClusterSKU_ARM).

| Value  | Description |
|--------|-------------|
| "Free" |             |
| "Paid" |             |

<a id="ManagedClusterSKU_Tier_STATUS"></a>ManagedClusterSKU_Tier_STATUS
-----------------------------------------------------------------------

Used by: [ManagedClusterSKU_STATUS](#ManagedClusterSKU_STATUS), and [ManagedClusterSKU_STATUS_ARM](#ManagedClusterSKU_STATUS_ARM).

| Value  | Description |
|--------|-------------|
| "Free" |             |
| "Paid" |             |

<a id="ManagedClusterServicePrincipalProfile"></a>ManagedClusterServicePrincipalProfile
---------------------------------------------------------------------------------------

Information about a service principal identity for the cluster to use for manipulating Azure APIs.

Used by: [ManagedCluster_Spec](#ManagedCluster_Spec).

| Property | Description                                                              | Type                       |
|----------|--------------------------------------------------------------------------|----------------------------|
| clientId | The ID for the service principal.                                        | string<br/>Required        |
| secret   | The secret password associated with the service principal in plain text. | genruntime.SecretReference |

<a id="ManagedClusterServicePrincipalProfile_ARM"></a>ManagedClusterServicePrincipalProfile_ARM
-----------------------------------------------------------------------------------------------

Information about a service principal identity for the cluster to use for manipulating Azure APIs.

Used by: [ManagedClusterProperties_ARM](#ManagedClusterProperties_ARM).

| Property | Description                                                              | Type   |
|----------|--------------------------------------------------------------------------|--------|
| clientId | The ID for the service principal.                                        | string |
| secret   | The secret password associated with the service principal in plain text. | string |

<a id="ManagedClusterServicePrincipalProfile_STATUS"></a>ManagedClusterServicePrincipalProfile_STATUS
-----------------------------------------------------------------------------------------------------

Information about a service principal identity for the cluster to use for manipulating Azure APIs.

Used by: [ManagedCluster_STATUS](#ManagedCluster_STATUS).

| Property | Description                       | Type   |
|----------|-----------------------------------|--------|
| clientId | The ID for the service principal. | string |

<a id="ManagedClusterServicePrincipalProfile_STATUS_ARM"></a>ManagedClusterServicePrincipalProfile_STATUS_ARM
-------------------------------------------------------------------------------------------------------------

Information about a service principal identity for the cluster to use for manipulating Azure APIs.

Used by: [ManagedClusterProperties_STATUS_ARM](#ManagedClusterProperties_STATUS_ARM).

| Property | Description                       | Type   |
|----------|-----------------------------------|--------|
| clientId | The ID for the service principal. | string |

<a id="ManagedClusterWindowsProfile"></a>ManagedClusterWindowsProfile
---------------------------------------------------------------------

Profile for Windows VMs in the managed cluster.

Used by: [ManagedCluster_Spec](#ManagedCluster_Spec).

| Property       | Description                                                                                                                                                                                                                                                                                                                                                                                                                                                        | Type                                                                                  |
|----------------|--------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|---------------------------------------------------------------------------------------|
| adminPassword  | Specifies the password of the administrator account. Minimum-length: 8 characters Max-length: 123 characters Complexity requirements: 3 out of 4 conditions below need to be fulfilled Has lower characters Has upper characters Has a digit Has a special character (Regex match \[\W_]) Disallowed values: "abc@123", "P@$$w0rd", "P@ssw0rd", "P@ssword123", "Pa$$word", "pass@word1", "Password!", "Password1", "Password22", "iloveyou!"                       | string                                                                                |
| adminUsername  | Specifies the name of the administrator account. Restriction: Cannot end in "." Disallowed values: "administrator", "admin", "user", "user1", "test", "user2", "test1", "user3", "admin1", "1", "123", "a", "actuser", "adm", "admin2", "aspnet", "backup", "console", "david", "guest", "john", "owner", "root", "server", "sql", "support", "support_388945a0", "sys", "test2", "test3", "user4", "user5". Minimum-length: 1 character Max-length: 20 characters | string<br/>Required                                                                   |
| enableCSIProxy | For more details on CSI proxy, see the [CSI proxy GitHub repo](https://github.com/kubernetes-csi/csi-proxy).                                                                                                                                                                                                                                                                                                                                                       | bool                                                                                  |
| licenseType    | The license type to use for Windows VMs. See [Azure Hybrid User Benefits](https://azure.microsoft.com/pricing/hybrid-benefit/faq/) for more details.                                                                                                                                                                                                                                                                                                               | [ManagedClusterWindowsProfile_LicenseType](#ManagedClusterWindowsProfile_LicenseType) |

<a id="ManagedClusterWindowsProfile_ARM"></a>ManagedClusterWindowsProfile_ARM
-----------------------------------------------------------------------------

Profile for Windows VMs in the managed cluster.

Used by: [ManagedClusterProperties_ARM](#ManagedClusterProperties_ARM).

| Property       | Description                                                                                                                                                                                                                                                                                                                                                                                                                                                        | Type                                                                                  |
|----------------|--------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|---------------------------------------------------------------------------------------|
| adminPassword  | Specifies the password of the administrator account. Minimum-length: 8 characters Max-length: 123 characters Complexity requirements: 3 out of 4 conditions below need to be fulfilled Has lower characters Has upper characters Has a digit Has a special character (Regex match \[\W_]) Disallowed values: "abc@123", "P@$$w0rd", "P@ssw0rd", "P@ssword123", "Pa$$word", "pass@word1", "Password!", "Password1", "Password22", "iloveyou!"                       | string                                                                                |
| adminUsername  | Specifies the name of the administrator account. Restriction: Cannot end in "." Disallowed values: "administrator", "admin", "user", "user1", "test", "user2", "test1", "user3", "admin1", "1", "123", "a", "actuser", "adm", "admin2", "aspnet", "backup", "console", "david", "guest", "john", "owner", "root", "server", "sql", "support", "support_388945a0", "sys", "test2", "test3", "user4", "user5". Minimum-length: 1 character Max-length: 20 characters | string                                                                                |
| enableCSIProxy | For more details on CSI proxy, see the [CSI proxy GitHub repo](https://github.com/kubernetes-csi/csi-proxy).                                                                                                                                                                                                                                                                                                                                                       | bool                                                                                  |
| licenseType    | The license type to use for Windows VMs. See [Azure Hybrid User Benefits](https://azure.microsoft.com/pricing/hybrid-benefit/faq/) for more details.                                                                                                                                                                                                                                                                                                               | [ManagedClusterWindowsProfile_LicenseType](#ManagedClusterWindowsProfile_LicenseType) |

<a id="ManagedClusterWindowsProfile_LicenseType"></a>ManagedClusterWindowsProfile_LicenseType
---------------------------------------------------------------------------------------------

Used by: [ManagedClusterWindowsProfile](#ManagedClusterWindowsProfile), and [ManagedClusterWindowsProfile_ARM](#ManagedClusterWindowsProfile_ARM).

| Value            | Description |
|------------------|-------------|
| "None"           |             |
| "Windows_Server" |             |

<a id="ManagedClusterWindowsProfile_LicenseType_STATUS"></a>ManagedClusterWindowsProfile_LicenseType_STATUS
-----------------------------------------------------------------------------------------------------------

Used by: [ManagedClusterWindowsProfile_STATUS](#ManagedClusterWindowsProfile_STATUS), and [ManagedClusterWindowsProfile_STATUS_ARM](#ManagedClusterWindowsProfile_STATUS_ARM).

| Value            | Description |
|------------------|-------------|
| "None"           |             |
| "Windows_Server" |             |

<a id="ManagedClusterWindowsProfile_STATUS"></a>ManagedClusterWindowsProfile_STATUS
-----------------------------------------------------------------------------------

Profile for Windows VMs in the managed cluster.

Used by: [ManagedCluster_STATUS](#ManagedCluster_STATUS).

| Property       | Description                                                                                                                                                                                                                                                                                                                                                                                                                                                        | Type                                                                                                |
|----------------|--------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|-----------------------------------------------------------------------------------------------------|
| adminPassword  | Specifies the password of the administrator account. Minimum-length: 8 characters Max-length: 123 characters Complexity requirements: 3 out of 4 conditions below need to be fulfilled Has lower characters Has upper characters Has a digit Has a special character (Regex match \[\W_]) Disallowed values: "abc@123", "P@$$w0rd", "P@ssw0rd", "P@ssword123", "Pa$$word", "pass@word1", "Password!", "Password1", "Password22", "iloveyou!"                       | string                                                                                              |
| adminUsername  | Specifies the name of the administrator account. Restriction: Cannot end in "." Disallowed values: "administrator", "admin", "user", "user1", "test", "user2", "test1", "user3", "admin1", "1", "123", "a", "actuser", "adm", "admin2", "aspnet", "backup", "console", "david", "guest", "john", "owner", "root", "server", "sql", "support", "support_388945a0", "sys", "test2", "test3", "user4", "user5". Minimum-length: 1 character Max-length: 20 characters | string                                                                                              |
| enableCSIProxy | For more details on CSI proxy, see the [CSI proxy GitHub repo](https://github.com/kubernetes-csi/csi-proxy).                                                                                                                                                                                                                                                                                                                                                       | bool                                                                                                |
| licenseType    | The license type to use for Windows VMs. See [Azure Hybrid User Benefits](https://azure.microsoft.com/pricing/hybrid-benefit/faq/) for more details.                                                                                                                                                                                                                                                                                                               | [ManagedClusterWindowsProfile_LicenseType_STATUS](#ManagedClusterWindowsProfile_LicenseType_STATUS) |

<a id="ManagedClusterWindowsProfile_STATUS_ARM"></a>ManagedClusterWindowsProfile_STATUS_ARM
-------------------------------------------------------------------------------------------

Profile for Windows VMs in the managed cluster.

Used by: [ManagedClusterProperties_STATUS_ARM](#ManagedClusterProperties_STATUS_ARM).

| Property       | Description                                                                                                                                                                                                                                                                                                                                                                                                                                                        | Type                                                                                                |
|----------------|--------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|-----------------------------------------------------------------------------------------------------|
| adminPassword  | Specifies the password of the administrator account. Minimum-length: 8 characters Max-length: 123 characters Complexity requirements: 3 out of 4 conditions below need to be fulfilled Has lower characters Has upper characters Has a digit Has a special character (Regex match \[\W_]) Disallowed values: "abc@123", "P@$$w0rd", "P@ssw0rd", "P@ssword123", "Pa$$word", "pass@word1", "Password!", "Password1", "Password22", "iloveyou!"                       | string                                                                                              |
| adminUsername  | Specifies the name of the administrator account. Restriction: Cannot end in "." Disallowed values: "administrator", "admin", "user", "user1", "test", "user2", "test1", "user3", "admin1", "1", "123", "a", "actuser", "adm", "admin2", "aspnet", "backup", "console", "david", "guest", "john", "owner", "root", "server", "sql", "support", "support_388945a0", "sys", "test2", "test3", "user4", "user5". Minimum-length: 1 character Max-length: 20 characters | string                                                                                              |
| enableCSIProxy | For more details on CSI proxy, see the [CSI proxy GitHub repo](https://github.com/kubernetes-csi/csi-proxy).                                                                                                                                                                                                                                                                                                                                                       | bool                                                                                                |
| licenseType    | The license type to use for Windows VMs. See [Azure Hybrid User Benefits](https://azure.microsoft.com/pricing/hybrid-benefit/faq/) for more details.                                                                                                                                                                                                                                                                                                               | [ManagedClusterWindowsProfile_LicenseType_STATUS](#ManagedClusterWindowsProfile_LicenseType_STATUS) |

<a id="ManagedCluster_STATUS"></a>ManagedCluster_STATUS
-------------------------------------------------------

Managed cluster.

Used by: [ManagedCluster](#ManagedCluster).

| Property                | Description                                                                                                                                                                                                                                                                                                                                                                                 | Type                                                                                                    |
|-------------------------|---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|---------------------------------------------------------------------------------------------------------|
| aadProfile              | The Azure Active Directory configuration.                                                                                                                                                                                                                                                                                                                                                   | [ManagedClusterAADProfile_STATUS](#ManagedClusterAADProfile_STATUS)                                     |
| addonProfiles           | The profile of managed cluster add-on.                                                                                                                                                                                                                                                                                                                                                      | [map[string]ManagedClusterAddonProfile_STATUS](#ManagedClusterAddonProfile_STATUS)                      |
| agentPoolProfiles       | The agent pool properties.                                                                                                                                                                                                                                                                                                                                                                  | [ManagedClusterAgentPoolProfile_STATUS[]](#ManagedClusterAgentPoolProfile_STATUS)                       |
| apiServerAccessProfile  | The access profile for managed cluster API server.                                                                                                                                                                                                                                                                                                                                          | [ManagedClusterAPIServerAccessProfile_STATUS](#ManagedClusterAPIServerAccessProfile_STATUS)             |
| autoScalerProfile       | Parameters to be applied to the cluster-autoscaler when enabled                                                                                                                                                                                                                                                                                                                             | [ManagedClusterProperties_AutoScalerProfile_STATUS](#ManagedClusterProperties_AutoScalerProfile_STATUS) |
| autoUpgradeProfile      | The auto upgrade configuration.                                                                                                                                                                                                                                                                                                                                                             | [ManagedClusterAutoUpgradeProfile_STATUS](#ManagedClusterAutoUpgradeProfile_STATUS)                     |
| azurePortalFQDN         | The Azure Portal requires certain Cross-Origin Resource Sharing (CORS) headers to be sent in some responses, which Kubernetes APIServer doesn't handle by default. This special FQDN supports CORS, allowing the Azure Portal to function properly.                                                                                                                                         | string                                                                                                  |
| conditions              | The observed state of the resource                                                                                                                                                                                                                                                                                                                                                          | conditions.Condition[]                                                                                  |
| disableLocalAccounts    | If set to true, getting static credentials will be disabled for this cluster. This must only be used on Managed Clusters that are AAD enabled. For more details see [disable local accounts](https://docs.microsoft.com/azure/aks/managed-aad#disable-local-accounts-preview).                                                                                                              | bool                                                                                                    |
| diskEncryptionSetID     | This is of the form: '/&ZeroWidthSpace;subscriptions/&ZeroWidthSpace;{subscriptionId}/&ZeroWidthSpace;resourceGroups/&ZeroWidthSpace;{resourceGroupName}/&ZeroWidthSpace;providers/&ZeroWidthSpace;Microsoft.Compute/diskEncryptionSets/{encryptionSetName}'                                                                                                                                | string                                                                                                  |
| dnsPrefix               | This cannot be updated once the Managed Cluster has been created.                                                                                                                                                                                                                                                                                                                           | string                                                                                                  |
| enablePodSecurityPolicy | (DEPRECATING) Whether to enable Kubernetes pod security policy (preview). This feature is set for removal on October 15th, 2020. Learn more at aka.ms/aks/azpodpolicy.                                                                                                                                                                                                                      | bool                                                                                                    |
| enableRBAC              | Whether to enable Kubernetes Role-Based Access Control.                                                                                                                                                                                                                                                                                                                                     | bool                                                                                                    |
| extendedLocation        | The extended location of the Virtual Machine.                                                                                                                                                                                                                                                                                                                                               | [ExtendedLocation_STATUS](#ExtendedLocation_STATUS)                                                     |
| fqdn                    | The FQDN of the master pool.                                                                                                                                                                                                                                                                                                                                                                | string                                                                                                  |
| fqdnSubdomain           | This cannot be updated once the Managed Cluster has been created.                                                                                                                                                                                                                                                                                                                           | string                                                                                                  |
| httpProxyConfig         | Configurations for provisioning the cluster with HTTP proxy servers.                                                                                                                                                                                                                                                                                                                        | [ManagedClusterHTTPProxyConfig_STATUS](#ManagedClusterHTTPProxyConfig_STATUS)                           |
| id                      | Resource Id                                                                                                                                                                                                                                                                                                                                                                                 | string                                                                                                  |
| identity                | The identity of the managed cluster, if configured.                                                                                                                                                                                                                                                                                                                                         | [ManagedClusterIdentity_STATUS](#ManagedClusterIdentity_STATUS)                                         |
| identityProfile         | Identities associated with the cluster.                                                                                                                                                                                                                                                                                                                                                     | [map[string]UserAssignedIdentity_STATUS](#UserAssignedIdentity_STATUS)                                  |
| kubernetesVersion       | When you upgrade a supported AKS cluster, Kubernetes minor versions cannot be skipped. All upgrades must be performed sequentially by major version number. For example, upgrades between 1.14.x -> 1.15.x or 1.15.x -> 1.16.x are allowed, however 1.14.x -> 1.16.x is not allowed. See [upgrading an AKS cluster](https://docs.microsoft.com/azure/aks/upgrade-cluster) for more details. | string                                                                                                  |
| linuxProfile            | The profile for Linux VMs in the Managed Cluster.                                                                                                                                                                                                                                                                                                                                           | [ContainerServiceLinuxProfile_STATUS](#ContainerServiceLinuxProfile_STATUS)                             |
| location                | Resource location                                                                                                                                                                                                                                                                                                                                                                           | string                                                                                                  |
| maxAgentPools           | The max number of agent pools for the managed cluster.                                                                                                                                                                                                                                                                                                                                      | int                                                                                                     |
| name                    | Resource name                                                                                                                                                                                                                                                                                                                                                                               | string                                                                                                  |
| networkProfile          | The network configuration profile.                                                                                                                                                                                                                                                                                                                                                          | [ContainerServiceNetworkProfile_STATUS](#ContainerServiceNetworkProfile_STATUS)                         |
| nodeResourceGroup       | The name of the resource group containing agent pool nodes.                                                                                                                                                                                                                                                                                                                                 | string                                                                                                  |
| podIdentityProfile      | See [use AAD pod identity](https://docs.microsoft.com/azure/aks/use-azure-ad-pod-identity) for more details on AAD pod identity integration.                                                                                                                                                                                                                                                | [ManagedClusterPodIdentityProfile_STATUS](#ManagedClusterPodIdentityProfile_STATUS)                     |
| powerState              | The Power State of the cluster.                                                                                                                                                                                                                                                                                                                                                             | [PowerState_STATUS](#PowerState_STATUS)                                                                 |
| privateFQDN             | The FQDN of private cluster.                                                                                                                                                                                                                                                                                                                                                                | string                                                                                                  |
| privateLinkResources    | Private link resources associated with the cluster.                                                                                                                                                                                                                                                                                                                                         | [PrivateLinkResource_STATUS[]](#PrivateLinkResource_STATUS)                                             |
| provisioningState       | The current provisioning state.                                                                                                                                                                                                                                                                                                                                                             | string                                                                                                  |
| servicePrincipalProfile | Information about a service principal identity for the cluster to use for manipulating Azure APIs.                                                                                                                                                                                                                                                                                          | [ManagedClusterServicePrincipalProfile_STATUS](#ManagedClusterServicePrincipalProfile_STATUS)           |
| sku                     | The managed cluster SKU.                                                                                                                                                                                                                                                                                                                                                                    | [ManagedClusterSKU_STATUS](#ManagedClusterSKU_STATUS)                                                   |
| tags                    | Resource tags                                                                                                                                                                                                                                                                                                                                                                               | map[string]string                                                                                       |
| type                    | Resource type                                                                                                                                                                                                                                                                                                                                                                               | string                                                                                                  |
| windowsProfile          | The profile for Windows VMs in the Managed Cluster.                                                                                                                                                                                                                                                                                                                                         | [ManagedClusterWindowsProfile_STATUS](#ManagedClusterWindowsProfile_STATUS)                             |

<a id="ManagedCluster_STATUS_ARM"></a>ManagedCluster_STATUS_ARM
---------------------------------------------------------------

Managed cluster.

| Property         | Description                                         | Type                                                                        |
|------------------|-----------------------------------------------------|-----------------------------------------------------------------------------|
| extendedLocation | The extended location of the Virtual Machine.       | [ExtendedLocation_STATUS_ARM](#ExtendedLocation_STATUS_ARM)                 |
| id               | Resource Id                                         | string                                                                      |
| identity         | The identity of the managed cluster, if configured. | [ManagedClusterIdentity_STATUS_ARM](#ManagedClusterIdentity_STATUS_ARM)     |
| location         | Resource location                                   | string                                                                      |
| name             | Resource name                                       | string                                                                      |
| properties       | Properties of a managed cluster.                    | [ManagedClusterProperties_STATUS_ARM](#ManagedClusterProperties_STATUS_ARM) |
| sku              | The managed cluster SKU.                            | [ManagedClusterSKU_STATUS_ARM](#ManagedClusterSKU_STATUS_ARM)               |
| tags             | Resource tags                                       | map[string]string                                                           |
| type             | Resource type                                       | string                                                                      |

<a id="ManagedCluster_Spec"></a>ManagedCluster_Spec
---------------------------------------------------

Used by: [ManagedCluster](#ManagedCluster).

| Property                     | Description                                                                                                                                                                                                                                                                                                                                                                                 | Type                                                                                      |
|------------------------------|---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|-------------------------------------------------------------------------------------------|
| aadProfile                   | The Azure Active Directory configuration.                                                                                                                                                                                                                                                                                                                                                   | [ManagedClusterAADProfile](#ManagedClusterAADProfile)                                     |
| addonProfiles                | The profile of managed cluster add-on.                                                                                                                                                                                                                                                                                                                                                      | [map[string]ManagedClusterAddonProfile](#ManagedClusterAddonProfile)                      |
| agentPoolProfiles            | The agent pool properties.                                                                                                                                                                                                                                                                                                                                                                  | [ManagedClusterAgentPoolProfile[]](#ManagedClusterAgentPoolProfile)                       |
| apiServerAccessProfile       | The access profile for managed cluster API server.                                                                                                                                                                                                                                                                                                                                          | [ManagedClusterAPIServerAccessProfile](#ManagedClusterAPIServerAccessProfile)             |
| autoScalerProfile            | Parameters to be applied to the cluster-autoscaler when enabled                                                                                                                                                                                                                                                                                                                             | [ManagedClusterProperties_AutoScalerProfile](#ManagedClusterProperties_AutoScalerProfile) |
| autoUpgradeProfile           | The auto upgrade configuration.                                                                                                                                                                                                                                                                                                                                                             | [ManagedClusterAutoUpgradeProfile](#ManagedClusterAutoUpgradeProfile)                     |
| azureName                    | The name of the resource in Azure. This is often the same as the name of the resource in Kubernetes but it doesn't have to be.                                                                                                                                                                                                                                                              | string                                                                                    |
| disableLocalAccounts         | If set to true, getting static credentials will be disabled for this cluster. This must only be used on Managed Clusters that are AAD enabled. For more details see [disable local accounts](https://docs.microsoft.com/azure/aks/managed-aad#disable-local-accounts-preview).                                                                                                              | bool                                                                                      |
| diskEncryptionSetIDReference | This is of the form: '/&ZeroWidthSpace;subscriptions/&ZeroWidthSpace;{subscriptionId}/&ZeroWidthSpace;resourceGroups/&ZeroWidthSpace;{resourceGroupName}/&ZeroWidthSpace;providers/&ZeroWidthSpace;Microsoft.Compute/diskEncryptionSets/{encryptionSetName}'                                                                                                                                | genruntime.ResourceReference                                                              |
| dnsPrefix                    | This cannot be updated once the Managed Cluster has been created.                                                                                                                                                                                                                                                                                                                           | string                                                                                    |
| enablePodSecurityPolicy      | (DEPRECATING) Whether to enable Kubernetes pod security policy (preview). This feature is set for removal on October 15th, 2020. Learn more at aka.ms/aks/azpodpolicy.                                                                                                                                                                                                                      | bool                                                                                      |
| enableRBAC                   | Whether to enable Kubernetes Role-Based Access Control.                                                                                                                                                                                                                                                                                                                                     | bool                                                                                      |
| extendedLocation             | The extended location of the Virtual Machine.                                                                                                                                                                                                                                                                                                                                               | [ExtendedLocation](#ExtendedLocation)                                                     |
| fqdnSubdomain                | This cannot be updated once the Managed Cluster has been created.                                                                                                                                                                                                                                                                                                                           | string                                                                                    |
| httpProxyConfig              | Configurations for provisioning the cluster with HTTP proxy servers.                                                                                                                                                                                                                                                                                                                        | [ManagedClusterHTTPProxyConfig](#ManagedClusterHTTPProxyConfig)                           |
| identity                     | The identity of the managed cluster, if configured.                                                                                                                                                                                                                                                                                                                                         | [ManagedClusterIdentity](#ManagedClusterIdentity)                                         |
| identityProfile              | Identities associated with the cluster.                                                                                                                                                                                                                                                                                                                                                     | [map[string]UserAssignedIdentity](#UserAssignedIdentity)                                  |
| kubernetesVersion            | When you upgrade a supported AKS cluster, Kubernetes minor versions cannot be skipped. All upgrades must be performed sequentially by major version number. For example, upgrades between 1.14.x -> 1.15.x or 1.15.x -> 1.16.x are allowed, however 1.14.x -> 1.16.x is not allowed. See [upgrading an AKS cluster](https://docs.microsoft.com/azure/aks/upgrade-cluster) for more details. | string                                                                                    |
| linuxProfile                 | The profile for Linux VMs in the Managed Cluster.                                                                                                                                                                                                                                                                                                                                           | [ContainerServiceLinuxProfile](#ContainerServiceLinuxProfile)                             |
| location                     | Resource location                                                                                                                                                                                                                                                                                                                                                                           | string<br/>Required                                                                       |
| networkProfile               | The network configuration profile.                                                                                                                                                                                                                                                                                                                                                          | [ContainerServiceNetworkProfile](#ContainerServiceNetworkProfile)                         |
| nodeResourceGroup            | The name of the resource group containing agent pool nodes.                                                                                                                                                                                                                                                                                                                                 | string                                                                                    |
| operatorSpec                 | The specification for configuring operator behavior. This field is interpreted by the operator and not passed directly to Azure                                                                                                                                                                                                                                                             | [ManagedClusterOperatorSpec](#ManagedClusterOperatorSpec)                                 |
| owner                        | The owner of the resource. The owner controls where the resource goes when it is deployed. The owner also controls the resources lifecycle. When the owner is deleted the resource will also be deleted. Owner is expected to be a reference to a resources.azure.com/ResourceGroup resource                                                                                                | genruntime.KnownResourceReference<br/>Required                                            |
| podIdentityProfile           | See [use AAD pod identity](https://docs.microsoft.com/azure/aks/use-azure-ad-pod-identity) for more details on AAD pod identity integration.                                                                                                                                                                                                                                                | [ManagedClusterPodIdentityProfile](#ManagedClusterPodIdentityProfile)                     |
| privateLinkResources         | Private link resources associated with the cluster.                                                                                                                                                                                                                                                                                                                                         | [PrivateLinkResource[]](#PrivateLinkResource)                                             |
| servicePrincipalProfile      | Information about a service principal identity for the cluster to use for manipulating Azure APIs.                                                                                                                                                                                                                                                                                          | [ManagedClusterServicePrincipalProfile](#ManagedClusterServicePrincipalProfile)           |
| sku                          | The managed cluster SKU.                                                                                                                                                                                                                                                                                                                                                                    | [ManagedClusterSKU](#ManagedClusterSKU)                                                   |
| tags                         | Resource tags                                                                                                                                                                                                                                                                                                                                                                               | map[string]string                                                                         |
| windowsProfile               | The profile for Windows VMs in the Managed Cluster.                                                                                                                                                                                                                                                                                                                                         | [ManagedClusterWindowsProfile](#ManagedClusterWindowsProfile)                             |

<a id="ManagedCluster_Spec_ARM"></a>ManagedCluster_Spec_ARM
-----------------------------------------------------------

| Property         | Description                                         | Type                                                          |
|------------------|-----------------------------------------------------|---------------------------------------------------------------|
| extendedLocation | The extended location of the Virtual Machine.       | [ExtendedLocation_ARM](#ExtendedLocation_ARM)                 |
| identity         | The identity of the managed cluster, if configured. | [ManagedClusterIdentity_ARM](#ManagedClusterIdentity_ARM)     |
| location         | Resource location                                   | string                                                        |
| name             |                                                     | string                                                        |
| properties       | Properties of a managed cluster.                    | [ManagedClusterProperties_ARM](#ManagedClusterProperties_ARM) |
| sku              | The managed cluster SKU.                            | [ManagedClusterSKU_ARM](#ManagedClusterSKU_ARM)               |
| tags             | Resource tags                                       | map[string]string                                             |

<a id="ManagedClustersAgentPool"></a>ManagedClustersAgentPool
-------------------------------------------------------------

Generator information: - Generated from: /containerservice/resource-manager/Microsoft.ContainerService/aks/stable/2021-05-01/managedClusters.json - ARM URI: /&ZeroWidthSpace;subscriptions/&ZeroWidthSpace;{subscriptionId}/&ZeroWidthSpace;resourceGroups/&ZeroWidthSpace;{resourceGroupName}/&ZeroWidthSpace;providers/&ZeroWidthSpace;Microsoft.ContainerService/managedClusters/{resourceName}/agentPools/{agentPoolName}

Used by: [ManagedClustersAgentPoolList](#ManagedClustersAgentPoolList).

| Property          | Description | Type                                                                  |
|-------------------|-------------|-----------------------------------------------------------------------|
| metav1.TypeMeta   |             |                                                                       |
| metav1.ObjectMeta |             |                                                                       |
| spec              |             | [ManagedClusters_AgentPool_Spec](#ManagedClusters_AgentPool_Spec)     |
| status            |             | [ManagedClusters_AgentPool_STATUS](#ManagedClusters_AgentPool_STATUS) |

### <a id="ManagedClusters_AgentPool_Spec"></a>ManagedClusters_AgentPool_Spec

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
| gpuInstanceProfile            | GPUInstanceProfile to be used to specify GPU MIG instance profile for supported GPU VM SKU.                                                                                                                                                                                                                                                                                                                                                                                                  | [GPUInstanceProfile](#GPUInstanceProfile)             |
| kubeletConfig                 | The Kubelet configuration on the agent pool nodes.                                                                                                                                                                                                                                                                                                                                                                                                                                           | [KubeletConfig](#KubeletConfig)                       |
| kubeletDiskType               | Determines the placement of emptyDir volumes, container runtime data root, and Kubelet ephemeral storage.                                                                                                                                                                                                                                                                                                                                                                                    | [KubeletDiskType](#KubeletDiskType)                   |
| linuxOSConfig                 | The OS configuration of Linux agent nodes.                                                                                                                                                                                                                                                                                                                                                                                                                                                   | [LinuxOSConfig](#LinuxOSConfig)                       |
| maxCount                      | The maximum number of nodes for auto-scaling                                                                                                                                                                                                                                                                                                                                                                                                                                                 | int                                                   |
| maxPods                       | The maximum number of pods that can run on a node.                                                                                                                                                                                                                                                                                                                                                                                                                                           | int                                                   |
| minCount                      | The minimum number of nodes for auto-scaling                                                                                                                                                                                                                                                                                                                                                                                                                                                 | int                                                   |
| mode                          | A cluster must have at least one 'System' Agent Pool at all times. For additional information on agent pool restrictions and best practices, see: https://docs.microsoft.com/azure/aks/use-system-pools                                                                                                                                                                                                                                                                                      | [AgentPoolMode](#AgentPoolMode)                       |
| nodeLabels                    | The node labels to be persisted across all nodes in agent pool.                                                                                                                                                                                                                                                                                                                                                                                                                              | map[string]string                                     |
| nodePublicIPPrefixIDReference | This is of the form: /&ZeroWidthSpace;subscriptions/&ZeroWidthSpace;{subscriptionId}/&ZeroWidthSpace;resourceGroups/&ZeroWidthSpace;{resourceGroupName}/&ZeroWidthSpace;providers/&ZeroWidthSpace;Microsoft.Network/publicIPPrefixes/{publicIPPrefixName}                                                                                                                                                                                                                                    | genruntime.ResourceReference                          |
| nodeTaints                    | The taints added to new nodes during node pool create and scale. For example, key=value:NoSchedule.                                                                                                                                                                                                                                                                                                                                                                                          | string[]                                              |
| orchestratorVersion           | As a best practice, you should upgrade all node pools in an AKS cluster to the same Kubernetes version. The node pool version must have the same major version as the control plane. The node pool minor version must be within two minor versions of the control plane version. The node pool version cannot be greater than the control plane version. For more information see [upgrading a node pool](https://docs.microsoft.com/azure/aks/use-multiple-node-pools#upgrade-a-node-pool). | string                                                |
| osDiskSizeGB                  |                                                                                                                                                                                                                                                                                                                                                                                                                                                                                              | [ContainerServiceOSDisk](#ContainerServiceOSDisk)     |
| osDiskType                    | The default is 'Ephemeral' if the VM supports it and has a cache disk larger than the requested OSDiskSizeGB. Otherwise, defaults to 'Managed'. May not be changed after creation. For more information see [Ephemeral OS](https://docs.microsoft.com/azure/aks/cluster-configuration#ephemeral-os).                                                                                                                                                                                         | [OSDiskType](#OSDiskType)                             |
| osSKU                         | Specifies an OS SKU. This value must not be specified if OSType is Windows.                                                                                                                                                                                                                                                                                                                                                                                                                  | [OSSKU](#OSSKU)                                       |
| osType                        | The operating system type. The default is Linux.                                                                                                                                                                                                                                                                                                                                                                                                                                             | [OSType](#OSType)                                     |
| owner                         | The owner of the resource. The owner controls where the resource goes when it is deployed. The owner also controls the resources lifecycle. When the owner is deleted the resource will also be deleted. Owner is expected to be a reference to a containerservice.azure.com/ManagedCluster resource                                                                                                                                                                                         | genruntime.KnownResourceReference<br/>Required        |
| podSubnetIDReference          | If omitted, pod IPs are statically assigned on the node subnet (see vnetSubnetID for more details). This is of the form: /&ZeroWidthSpace;subscriptions/&ZeroWidthSpace;{subscriptionId}/&ZeroWidthSpace;resourceGroups/&ZeroWidthSpace;{resourceGroupName}/&ZeroWidthSpace;providers/&ZeroWidthSpace;Microsoft.Network/virtualNetworks/{virtualNetworkName}/subnets/{subnetName}                                                                                                            | genruntime.ResourceReference                          |
| proximityPlacementGroupID     | The ID for Proximity Placement Group.                                                                                                                                                                                                                                                                                                                                                                                                                                                        | string                                                |
| scaleSetEvictionPolicy        | This cannot be specified unless the scaleSetPriority is 'Spot'. If not specified, the default is 'Delete'.                                                                                                                                                                                                                                                                                                                                                                                   | [ScaleSetEvictionPolicy](#ScaleSetEvictionPolicy)     |
| scaleSetPriority              | The Virtual Machine Scale Set priority. If not specified, the default is 'Regular'.                                                                                                                                                                                                                                                                                                                                                                                                          | [ScaleSetPriority](#ScaleSetPriority)                 |
| spotMaxPrice                  | Possible values are any decimal value greater than zero or -1 which indicates the willingness to pay any on-demand price. For more details on spot pricing, see [spot VMs pricing](https://docs.microsoft.com/azure/virtual-machines/spot-vms#pricing)                                                                                                                                                                                                                                       | float64                                               |
| tags                          | The tags to be persisted on the agent pool virtual machine scale set.                                                                                                                                                                                                                                                                                                                                                                                                                        | map[string]string                                     |
| type                          | The type of Agent Pool.                                                                                                                                                                                                                                                                                                                                                                                                                                                                      | [AgentPoolType](#AgentPoolType)                       |
| upgradeSettings               | Settings for upgrading the agentpool                                                                                                                                                                                                                                                                                                                                                                                                                                                         | [AgentPoolUpgradeSettings](#AgentPoolUpgradeSettings) |
| vmSize                        | VM size availability varies by region. If a node contains insufficient compute resources (memory, cpu, etc) pods might fail to run correctly. For more details on restricted VM sizes, see: https://docs.microsoft.com/azure/aks/quotas-skus-regions                                                                                                                                                                                                                                         | string                                                |
| vnetSubnetIDReference         | If this is not specified, a VNET and subnet will be generated and used. If no podSubnetID is specified, this applies to nodes and pods, otherwise it applies to just nodes. This is of the form: /&ZeroWidthSpace;subscriptions/&ZeroWidthSpace;{subscriptionId}/&ZeroWidthSpace;resourceGroups/&ZeroWidthSpace;{resourceGroupName}/&ZeroWidthSpace;providers/&ZeroWidthSpace;Microsoft.Network/virtualNetworks/{virtualNetworkName}/subnets/{subnetName}                                    | genruntime.ResourceReference                          |

### <a id="ManagedClusters_AgentPool_STATUS"></a>ManagedClusters_AgentPool_STATUS

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
| gpuInstanceProfile        | GPUInstanceProfile to be used to specify GPU MIG instance profile for supported GPU VM SKU.                                                                                                                                                                                                                                                                                                                                                                                                  | [GPUInstanceProfile_STATUS](#GPUInstanceProfile_STATUS)             |
| id                        | Resource ID.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                 | string                                                              |
| kubeletConfig             | The Kubelet configuration on the agent pool nodes.                                                                                                                                                                                                                                                                                                                                                                                                                                           | [KubeletConfig_STATUS](#KubeletConfig_STATUS)                       |
| kubeletDiskType           | Determines the placement of emptyDir volumes, container runtime data root, and Kubelet ephemeral storage.                                                                                                                                                                                                                                                                                                                                                                                    | [KubeletDiskType_STATUS](#KubeletDiskType_STATUS)                   |
| linuxOSConfig             | The OS configuration of Linux agent nodes.                                                                                                                                                                                                                                                                                                                                                                                                                                                   | [LinuxOSConfig_STATUS](#LinuxOSConfig_STATUS)                       |
| maxCount                  | The maximum number of nodes for auto-scaling                                                                                                                                                                                                                                                                                                                                                                                                                                                 | int                                                                 |
| maxPods                   | The maximum number of pods that can run on a node.                                                                                                                                                                                                                                                                                                                                                                                                                                           | int                                                                 |
| minCount                  | The minimum number of nodes for auto-scaling                                                                                                                                                                                                                                                                                                                                                                                                                                                 | int                                                                 |
| mode                      | A cluster must have at least one 'System' Agent Pool at all times. For additional information on agent pool restrictions and best practices, see: https://docs.microsoft.com/azure/aks/use-system-pools                                                                                                                                                                                                                                                                                      | [AgentPoolMode_STATUS](#AgentPoolMode_STATUS)                       |
| name                      | The name of the resource that is unique within a resource group. This name can be used to access the resource.                                                                                                                                                                                                                                                                                                                                                                               | string                                                              |
| nodeImageVersion          | The version of node image                                                                                                                                                                                                                                                                                                                                                                                                                                                                    | string                                                              |
| nodeLabels                | The node labels to be persisted across all nodes in agent pool.                                                                                                                                                                                                                                                                                                                                                                                                                              | map[string]string                                                   |
| nodePublicIPPrefixID      | This is of the form: /&ZeroWidthSpace;subscriptions/&ZeroWidthSpace;{subscriptionId}/&ZeroWidthSpace;resourceGroups/&ZeroWidthSpace;{resourceGroupName}/&ZeroWidthSpace;providers/&ZeroWidthSpace;Microsoft.Network/publicIPPrefixes/{publicIPPrefixName}                                                                                                                                                                                                                                    | string                                                              |
| nodeTaints                | The taints added to new nodes during node pool create and scale. For example, key=value:NoSchedule.                                                                                                                                                                                                                                                                                                                                                                                          | string[]                                                            |
| orchestratorVersion       | As a best practice, you should upgrade all node pools in an AKS cluster to the same Kubernetes version. The node pool version must have the same major version as the control plane. The node pool minor version must be within two minor versions of the control plane version. The node pool version cannot be greater than the control plane version. For more information see [upgrading a node pool](https://docs.microsoft.com/azure/aks/use-multiple-node-pools#upgrade-a-node-pool). | string                                                              |
| osDiskSizeGB              |                                                                                                                                                                                                                                                                                                                                                                                                                                                                                              | int                                                                 |
| osDiskType                | The default is 'Ephemeral' if the VM supports it and has a cache disk larger than the requested OSDiskSizeGB. Otherwise, defaults to 'Managed'. May not be changed after creation. For more information see [Ephemeral OS](https://docs.microsoft.com/azure/aks/cluster-configuration#ephemeral-os).                                                                                                                                                                                         | [OSDiskType_STATUS](#OSDiskType_STATUS)                             |
| osSKU                     | Specifies an OS SKU. This value must not be specified if OSType is Windows.                                                                                                                                                                                                                                                                                                                                                                                                                  | [OSSKU_STATUS](#OSSKU_STATUS)                                       |
| osType                    | The operating system type. The default is Linux.                                                                                                                                                                                                                                                                                                                                                                                                                                             | [OSType_STATUS](#OSType_STATUS)                                     |
| podSubnetID               | If omitted, pod IPs are statically assigned on the node subnet (see vnetSubnetID for more details). This is of the form: /&ZeroWidthSpace;subscriptions/&ZeroWidthSpace;{subscriptionId}/&ZeroWidthSpace;resourceGroups/&ZeroWidthSpace;{resourceGroupName}/&ZeroWidthSpace;providers/&ZeroWidthSpace;Microsoft.Network/virtualNetworks/{virtualNetworkName}/subnets/{subnetName}                                                                                                            | string                                                              |
| powerState                | Describes whether the Agent Pool is Running or Stopped                                                                                                                                                                                                                                                                                                                                                                                                                                       | [PowerState_STATUS](#PowerState_STATUS)                             |
| properties_type           | The type of Agent Pool.                                                                                                                                                                                                                                                                                                                                                                                                                                                                      | [AgentPoolType_STATUS](#AgentPoolType_STATUS)                       |
| provisioningState         | The current deployment or provisioning state.                                                                                                                                                                                                                                                                                                                                                                                                                                                | string                                                              |
| proximityPlacementGroupID | The ID for Proximity Placement Group.                                                                                                                                                                                                                                                                                                                                                                                                                                                        | string                                                              |
| scaleSetEvictionPolicy    | This cannot be specified unless the scaleSetPriority is 'Spot'. If not specified, the default is 'Delete'.                                                                                                                                                                                                                                                                                                                                                                                   | [ScaleSetEvictionPolicy_STATUS](#ScaleSetEvictionPolicy_STATUS)     |
| scaleSetPriority          | The Virtual Machine Scale Set priority. If not specified, the default is 'Regular'.                                                                                                                                                                                                                                                                                                                                                                                                          | [ScaleSetPriority_STATUS](#ScaleSetPriority_STATUS)                 |
| spotMaxPrice              | Possible values are any decimal value greater than zero or -1 which indicates the willingness to pay any on-demand price. For more details on spot pricing, see [spot VMs pricing](https://docs.microsoft.com/azure/virtual-machines/spot-vms#pricing)                                                                                                                                                                                                                                       | float64                                                             |
| tags                      | The tags to be persisted on the agent pool virtual machine scale set.                                                                                                                                                                                                                                                                                                                                                                                                                        | map[string]string                                                   |
| type                      | Resource type                                                                                                                                                                                                                                                                                                                                                                                                                                                                                | string                                                              |
| upgradeSettings           | Settings for upgrading the agentpool                                                                                                                                                                                                                                                                                                                                                                                                                                                         | [AgentPoolUpgradeSettings_STATUS](#AgentPoolUpgradeSettings_STATUS) |
| vmSize                    | VM size availability varies by region. If a node contains insufficient compute resources (memory, cpu, etc) pods might fail to run correctly. For more details on restricted VM sizes, see: https://docs.microsoft.com/azure/aks/quotas-skus-regions                                                                                                                                                                                                                                         | string                                                              |
| vnetSubnetID              | If this is not specified, a VNET and subnet will be generated and used. If no podSubnetID is specified, this applies to nodes and pods, otherwise it applies to just nodes. This is of the form: /&ZeroWidthSpace;subscriptions/&ZeroWidthSpace;{subscriptionId}/&ZeroWidthSpace;resourceGroups/&ZeroWidthSpace;{resourceGroupName}/&ZeroWidthSpace;providers/&ZeroWidthSpace;Microsoft.Network/virtualNetworks/{virtualNetworkName}/subnets/{subnetName}                                    | string                                                              |

<a id="ManagedClustersAgentPoolList"></a>ManagedClustersAgentPoolList
---------------------------------------------------------------------

Generator information: - Generated from: /containerservice/resource-manager/Microsoft.ContainerService/aks/stable/2021-05-01/managedClusters.json - ARM URI: /&ZeroWidthSpace;subscriptions/&ZeroWidthSpace;{subscriptionId}/&ZeroWidthSpace;resourceGroups/&ZeroWidthSpace;{resourceGroupName}/&ZeroWidthSpace;providers/&ZeroWidthSpace;Microsoft.ContainerService/managedClusters/{resourceName}/agentPools/{agentPoolName}

| Property        | Description | Type                                                    |
|-----------------|-------------|---------------------------------------------------------|
| metav1.TypeMeta |             |                                                         |
| metav1.ListMeta |             |                                                         |
| items           |             | [ManagedClustersAgentPool[]](#ManagedClustersAgentPool) |

<a id="ManagedClusters_AgentPool_STATUS"></a>ManagedClusters_AgentPool_STATUS
-----------------------------------------------------------------------------

Used by: [ManagedClustersAgentPool](#ManagedClustersAgentPool).

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
| gpuInstanceProfile        | GPUInstanceProfile to be used to specify GPU MIG instance profile for supported GPU VM SKU.                                                                                                                                                                                                                                                                                                                                                                                                  | [GPUInstanceProfile_STATUS](#GPUInstanceProfile_STATUS)             |
| id                        | Resource ID.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                 | string                                                              |
| kubeletConfig             | The Kubelet configuration on the agent pool nodes.                                                                                                                                                                                                                                                                                                                                                                                                                                           | [KubeletConfig_STATUS](#KubeletConfig_STATUS)                       |
| kubeletDiskType           | Determines the placement of emptyDir volumes, container runtime data root, and Kubelet ephemeral storage.                                                                                                                                                                                                                                                                                                                                                                                    | [KubeletDiskType_STATUS](#KubeletDiskType_STATUS)                   |
| linuxOSConfig             | The OS configuration of Linux agent nodes.                                                                                                                                                                                                                                                                                                                                                                                                                                                   | [LinuxOSConfig_STATUS](#LinuxOSConfig_STATUS)                       |
| maxCount                  | The maximum number of nodes for auto-scaling                                                                                                                                                                                                                                                                                                                                                                                                                                                 | int                                                                 |
| maxPods                   | The maximum number of pods that can run on a node.                                                                                                                                                                                                                                                                                                                                                                                                                                           | int                                                                 |
| minCount                  | The minimum number of nodes for auto-scaling                                                                                                                                                                                                                                                                                                                                                                                                                                                 | int                                                                 |
| mode                      | A cluster must have at least one 'System' Agent Pool at all times. For additional information on agent pool restrictions and best practices, see: https://docs.microsoft.com/azure/aks/use-system-pools                                                                                                                                                                                                                                                                                      | [AgentPoolMode_STATUS](#AgentPoolMode_STATUS)                       |
| name                      | The name of the resource that is unique within a resource group. This name can be used to access the resource.                                                                                                                                                                                                                                                                                                                                                                               | string                                                              |
| nodeImageVersion          | The version of node image                                                                                                                                                                                                                                                                                                                                                                                                                                                                    | string                                                              |
| nodeLabels                | The node labels to be persisted across all nodes in agent pool.                                                                                                                                                                                                                                                                                                                                                                                                                              | map[string]string                                                   |
| nodePublicIPPrefixID      | This is of the form: /&ZeroWidthSpace;subscriptions/&ZeroWidthSpace;{subscriptionId}/&ZeroWidthSpace;resourceGroups/&ZeroWidthSpace;{resourceGroupName}/&ZeroWidthSpace;providers/&ZeroWidthSpace;Microsoft.Network/publicIPPrefixes/{publicIPPrefixName}                                                                                                                                                                                                                                    | string                                                              |
| nodeTaints                | The taints added to new nodes during node pool create and scale. For example, key=value:NoSchedule.                                                                                                                                                                                                                                                                                                                                                                                          | string[]                                                            |
| orchestratorVersion       | As a best practice, you should upgrade all node pools in an AKS cluster to the same Kubernetes version. The node pool version must have the same major version as the control plane. The node pool minor version must be within two minor versions of the control plane version. The node pool version cannot be greater than the control plane version. For more information see [upgrading a node pool](https://docs.microsoft.com/azure/aks/use-multiple-node-pools#upgrade-a-node-pool). | string                                                              |
| osDiskSizeGB              |                                                                                                                                                                                                                                                                                                                                                                                                                                                                                              | int                                                                 |
| osDiskType                | The default is 'Ephemeral' if the VM supports it and has a cache disk larger than the requested OSDiskSizeGB. Otherwise, defaults to 'Managed'. May not be changed after creation. For more information see [Ephemeral OS](https://docs.microsoft.com/azure/aks/cluster-configuration#ephemeral-os).                                                                                                                                                                                         | [OSDiskType_STATUS](#OSDiskType_STATUS)                             |
| osSKU                     | Specifies an OS SKU. This value must not be specified if OSType is Windows.                                                                                                                                                                                                                                                                                                                                                                                                                  | [OSSKU_STATUS](#OSSKU_STATUS)                                       |
| osType                    | The operating system type. The default is Linux.                                                                                                                                                                                                                                                                                                                                                                                                                                             | [OSType_STATUS](#OSType_STATUS)                                     |
| podSubnetID               | If omitted, pod IPs are statically assigned on the node subnet (see vnetSubnetID for more details). This is of the form: /&ZeroWidthSpace;subscriptions/&ZeroWidthSpace;{subscriptionId}/&ZeroWidthSpace;resourceGroups/&ZeroWidthSpace;{resourceGroupName}/&ZeroWidthSpace;providers/&ZeroWidthSpace;Microsoft.Network/virtualNetworks/{virtualNetworkName}/subnets/{subnetName}                                                                                                            | string                                                              |
| powerState                | Describes whether the Agent Pool is Running or Stopped                                                                                                                                                                                                                                                                                                                                                                                                                                       | [PowerState_STATUS](#PowerState_STATUS)                             |
| properties_type           | The type of Agent Pool.                                                                                                                                                                                                                                                                                                                                                                                                                                                                      | [AgentPoolType_STATUS](#AgentPoolType_STATUS)                       |
| provisioningState         | The current deployment or provisioning state.                                                                                                                                                                                                                                                                                                                                                                                                                                                | string                                                              |
| proximityPlacementGroupID | The ID for Proximity Placement Group.                                                                                                                                                                                                                                                                                                                                                                                                                                                        | string                                                              |
| scaleSetEvictionPolicy    | This cannot be specified unless the scaleSetPriority is 'Spot'. If not specified, the default is 'Delete'.                                                                                                                                                                                                                                                                                                                                                                                   | [ScaleSetEvictionPolicy_STATUS](#ScaleSetEvictionPolicy_STATUS)     |
| scaleSetPriority          | The Virtual Machine Scale Set priority. If not specified, the default is 'Regular'.                                                                                                                                                                                                                                                                                                                                                                                                          | [ScaleSetPriority_STATUS](#ScaleSetPriority_STATUS)                 |
| spotMaxPrice              | Possible values are any decimal value greater than zero or -1 which indicates the willingness to pay any on-demand price. For more details on spot pricing, see [spot VMs pricing](https://docs.microsoft.com/azure/virtual-machines/spot-vms#pricing)                                                                                                                                                                                                                                       | float64                                                             |
| tags                      | The tags to be persisted on the agent pool virtual machine scale set.                                                                                                                                                                                                                                                                                                                                                                                                                        | map[string]string                                                   |
| type                      | Resource type                                                                                                                                                                                                                                                                                                                                                                                                                                                                                | string                                                              |
| upgradeSettings           | Settings for upgrading the agentpool                                                                                                                                                                                                                                                                                                                                                                                                                                                         | [AgentPoolUpgradeSettings_STATUS](#AgentPoolUpgradeSettings_STATUS) |
| vmSize                    | VM size availability varies by region. If a node contains insufficient compute resources (memory, cpu, etc) pods might fail to run correctly. For more details on restricted VM sizes, see: https://docs.microsoft.com/azure/aks/quotas-skus-regions                                                                                                                                                                                                                                         | string                                                              |
| vnetSubnetID              | If this is not specified, a VNET and subnet will be generated and used. If no podSubnetID is specified, this applies to nodes and pods, otherwise it applies to just nodes. This is of the form: /&ZeroWidthSpace;subscriptions/&ZeroWidthSpace;{subscriptionId}/&ZeroWidthSpace;resourceGroups/&ZeroWidthSpace;{resourceGroupName}/&ZeroWidthSpace;providers/&ZeroWidthSpace;Microsoft.Network/virtualNetworks/{virtualNetworkName}/subnets/{subnetName}                                    | string                                                              |

<a id="ManagedClusters_AgentPool_STATUS_ARM"></a>ManagedClusters_AgentPool_STATUS_ARM
-------------------------------------------------------------------------------------

| Property   | Description                                                                                                    | Type                                                                                                        |
|------------|----------------------------------------------------------------------------------------------------------------|-------------------------------------------------------------------------------------------------------------|
| id         | Resource ID.                                                                                                   | string                                                                                                      |
| name       | The name of the resource that is unique within a resource group. This name can be used to access the resource. | string                                                                                                      |
| properties | Properties of an agent pool.                                                                                   | [ManagedClusterAgentPoolProfileProperties_STATUS_ARM](#ManagedClusterAgentPoolProfileProperties_STATUS_ARM) |
| type       | Resource type                                                                                                  | string                                                                                                      |

<a id="ManagedClusters_AgentPool_Spec"></a>ManagedClusters_AgentPool_Spec
-------------------------------------------------------------------------

Used by: [ManagedClustersAgentPool](#ManagedClustersAgentPool).

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
| gpuInstanceProfile            | GPUInstanceProfile to be used to specify GPU MIG instance profile for supported GPU VM SKU.                                                                                                                                                                                                                                                                                                                                                                                                  | [GPUInstanceProfile](#GPUInstanceProfile)             |
| kubeletConfig                 | The Kubelet configuration on the agent pool nodes.                                                                                                                                                                                                                                                                                                                                                                                                                                           | [KubeletConfig](#KubeletConfig)                       |
| kubeletDiskType               | Determines the placement of emptyDir volumes, container runtime data root, and Kubelet ephemeral storage.                                                                                                                                                                                                                                                                                                                                                                                    | [KubeletDiskType](#KubeletDiskType)                   |
| linuxOSConfig                 | The OS configuration of Linux agent nodes.                                                                                                                                                                                                                                                                                                                                                                                                                                                   | [LinuxOSConfig](#LinuxOSConfig)                       |
| maxCount                      | The maximum number of nodes for auto-scaling                                                                                                                                                                                                                                                                                                                                                                                                                                                 | int                                                   |
| maxPods                       | The maximum number of pods that can run on a node.                                                                                                                                                                                                                                                                                                                                                                                                                                           | int                                                   |
| minCount                      | The minimum number of nodes for auto-scaling                                                                                                                                                                                                                                                                                                                                                                                                                                                 | int                                                   |
| mode                          | A cluster must have at least one 'System' Agent Pool at all times. For additional information on agent pool restrictions and best practices, see: https://docs.microsoft.com/azure/aks/use-system-pools                                                                                                                                                                                                                                                                                      | [AgentPoolMode](#AgentPoolMode)                       |
| nodeLabels                    | The node labels to be persisted across all nodes in agent pool.                                                                                                                                                                                                                                                                                                                                                                                                                              | map[string]string                                     |
| nodePublicIPPrefixIDReference | This is of the form: /&ZeroWidthSpace;subscriptions/&ZeroWidthSpace;{subscriptionId}/&ZeroWidthSpace;resourceGroups/&ZeroWidthSpace;{resourceGroupName}/&ZeroWidthSpace;providers/&ZeroWidthSpace;Microsoft.Network/publicIPPrefixes/{publicIPPrefixName}                                                                                                                                                                                                                                    | genruntime.ResourceReference                          |
| nodeTaints                    | The taints added to new nodes during node pool create and scale. For example, key=value:NoSchedule.                                                                                                                                                                                                                                                                                                                                                                                          | string[]                                              |
| orchestratorVersion           | As a best practice, you should upgrade all node pools in an AKS cluster to the same Kubernetes version. The node pool version must have the same major version as the control plane. The node pool minor version must be within two minor versions of the control plane version. The node pool version cannot be greater than the control plane version. For more information see [upgrading a node pool](https://docs.microsoft.com/azure/aks/use-multiple-node-pools#upgrade-a-node-pool). | string                                                |
| osDiskSizeGB                  |                                                                                                                                                                                                                                                                                                                                                                                                                                                                                              | [ContainerServiceOSDisk](#ContainerServiceOSDisk)     |
| osDiskType                    | The default is 'Ephemeral' if the VM supports it and has a cache disk larger than the requested OSDiskSizeGB. Otherwise, defaults to 'Managed'. May not be changed after creation. For more information see [Ephemeral OS](https://docs.microsoft.com/azure/aks/cluster-configuration#ephemeral-os).                                                                                                                                                                                         | [OSDiskType](#OSDiskType)                             |
| osSKU                         | Specifies an OS SKU. This value must not be specified if OSType is Windows.                                                                                                                                                                                                                                                                                                                                                                                                                  | [OSSKU](#OSSKU)                                       |
| osType                        | The operating system type. The default is Linux.                                                                                                                                                                                                                                                                                                                                                                                                                                             | [OSType](#OSType)                                     |
| owner                         | The owner of the resource. The owner controls where the resource goes when it is deployed. The owner also controls the resources lifecycle. When the owner is deleted the resource will also be deleted. Owner is expected to be a reference to a containerservice.azure.com/ManagedCluster resource                                                                                                                                                                                         | genruntime.KnownResourceReference<br/>Required        |
| podSubnetIDReference          | If omitted, pod IPs are statically assigned on the node subnet (see vnetSubnetID for more details). This is of the form: /&ZeroWidthSpace;subscriptions/&ZeroWidthSpace;{subscriptionId}/&ZeroWidthSpace;resourceGroups/&ZeroWidthSpace;{resourceGroupName}/&ZeroWidthSpace;providers/&ZeroWidthSpace;Microsoft.Network/virtualNetworks/{virtualNetworkName}/subnets/{subnetName}                                                                                                            | genruntime.ResourceReference                          |
| proximityPlacementGroupID     | The ID for Proximity Placement Group.                                                                                                                                                                                                                                                                                                                                                                                                                                                        | string                                                |
| scaleSetEvictionPolicy        | This cannot be specified unless the scaleSetPriority is 'Spot'. If not specified, the default is 'Delete'.                                                                                                                                                                                                                                                                                                                                                                                   | [ScaleSetEvictionPolicy](#ScaleSetEvictionPolicy)     |
| scaleSetPriority              | The Virtual Machine Scale Set priority. If not specified, the default is 'Regular'.                                                                                                                                                                                                                                                                                                                                                                                                          | [ScaleSetPriority](#ScaleSetPriority)                 |
| spotMaxPrice                  | Possible values are any decimal value greater than zero or -1 which indicates the willingness to pay any on-demand price. For more details on spot pricing, see [spot VMs pricing](https://docs.microsoft.com/azure/virtual-machines/spot-vms#pricing)                                                                                                                                                                                                                                       | float64                                               |
| tags                          | The tags to be persisted on the agent pool virtual machine scale set.                                                                                                                                                                                                                                                                                                                                                                                                                        | map[string]string                                     |
| type                          | The type of Agent Pool.                                                                                                                                                                                                                                                                                                                                                                                                                                                                      | [AgentPoolType](#AgentPoolType)                       |
| upgradeSettings               | Settings for upgrading the agentpool                                                                                                                                                                                                                                                                                                                                                                                                                                                         | [AgentPoolUpgradeSettings](#AgentPoolUpgradeSettings) |
| vmSize                        | VM size availability varies by region. If a node contains insufficient compute resources (memory, cpu, etc) pods might fail to run correctly. For more details on restricted VM sizes, see: https://docs.microsoft.com/azure/aks/quotas-skus-regions                                                                                                                                                                                                                                         | string                                                |
| vnetSubnetIDReference         | If this is not specified, a VNET and subnet will be generated and used. If no podSubnetID is specified, this applies to nodes and pods, otherwise it applies to just nodes. This is of the form: /&ZeroWidthSpace;subscriptions/&ZeroWidthSpace;{subscriptionId}/&ZeroWidthSpace;resourceGroups/&ZeroWidthSpace;{resourceGroupName}/&ZeroWidthSpace;providers/&ZeroWidthSpace;Microsoft.Network/virtualNetworks/{virtualNetworkName}/subnets/{subnetName}                                    | genruntime.ResourceReference                          |

<a id="ManagedClusters_AgentPool_Spec_ARM"></a>ManagedClusters_AgentPool_Spec_ARM
---------------------------------------------------------------------------------

| Property   | Description                  | Type                                                                                          |
|------------|------------------------------|-----------------------------------------------------------------------------------------------|
| name       |                              | string                                                                                        |
| properties | Properties of an agent pool. | [ManagedClusterAgentPoolProfileProperties_ARM](#ManagedClusterAgentPoolProfileProperties_ARM) |

<a id="OSDiskType"></a>OSDiskType
---------------------------------

Used by: [ManagedClusterAgentPoolProfile](#ManagedClusterAgentPoolProfile), [ManagedClusterAgentPoolProfileProperties_ARM](#ManagedClusterAgentPoolProfileProperties_ARM), [ManagedClusterAgentPoolProfile_ARM](#ManagedClusterAgentPoolProfile_ARM), and [ManagedClusters_AgentPool_Spec](#ManagedClusters_AgentPool_Spec).

| Value       | Description |
|-------------|-------------|
| "Ephemeral" |             |
| "Managed"   |             |

<a id="OSDiskType_STATUS"></a>OSDiskType_STATUS
-----------------------------------------------

Used by: [ManagedClusterAgentPoolProfileProperties_STATUS_ARM](#ManagedClusterAgentPoolProfileProperties_STATUS_ARM), [ManagedClusterAgentPoolProfile_STATUS](#ManagedClusterAgentPoolProfile_STATUS), [ManagedClusterAgentPoolProfile_STATUS_ARM](#ManagedClusterAgentPoolProfile_STATUS_ARM), and [ManagedClusters_AgentPool_STATUS](#ManagedClusters_AgentPool_STATUS).

| Value       | Description |
|-------------|-------------|
| "Ephemeral" |             |
| "Managed"   |             |

<a id="OSSKU"></a>OSSKU
-----------------------

Used by: [ManagedClusterAgentPoolProfile](#ManagedClusterAgentPoolProfile), [ManagedClusterAgentPoolProfileProperties_ARM](#ManagedClusterAgentPoolProfileProperties_ARM), [ManagedClusterAgentPoolProfile_ARM](#ManagedClusterAgentPoolProfile_ARM), and [ManagedClusters_AgentPool_Spec](#ManagedClusters_AgentPool_Spec).

| Value        | Description |
|--------------|-------------|
| "CBLMariner" |             |
| "Ubuntu"     |             |

<a id="OSSKU_STATUS"></a>OSSKU_STATUS
-------------------------------------

Used by: [ManagedClusterAgentPoolProfileProperties_STATUS_ARM](#ManagedClusterAgentPoolProfileProperties_STATUS_ARM), [ManagedClusterAgentPoolProfile_STATUS](#ManagedClusterAgentPoolProfile_STATUS), [ManagedClusterAgentPoolProfile_STATUS_ARM](#ManagedClusterAgentPoolProfile_STATUS_ARM), and [ManagedClusters_AgentPool_STATUS](#ManagedClusters_AgentPool_STATUS).

| Value        | Description |
|--------------|-------------|
| "CBLMariner" |             |
| "Ubuntu"     |             |

<a id="OSType"></a>OSType
-------------------------

Used by: [ManagedClusterAgentPoolProfile](#ManagedClusterAgentPoolProfile), [ManagedClusterAgentPoolProfileProperties_ARM](#ManagedClusterAgentPoolProfileProperties_ARM), [ManagedClusterAgentPoolProfile_ARM](#ManagedClusterAgentPoolProfile_ARM), and [ManagedClusters_AgentPool_Spec](#ManagedClusters_AgentPool_Spec).

| Value     | Description |
|-----------|-------------|
| "Linux"   |             |
| "Windows" |             |

<a id="OSType_STATUS"></a>OSType_STATUS
---------------------------------------

Used by: [ManagedClusterAgentPoolProfileProperties_STATUS_ARM](#ManagedClusterAgentPoolProfileProperties_STATUS_ARM), [ManagedClusterAgentPoolProfile_STATUS](#ManagedClusterAgentPoolProfile_STATUS), [ManagedClusterAgentPoolProfile_STATUS_ARM](#ManagedClusterAgentPoolProfile_STATUS_ARM), and [ManagedClusters_AgentPool_STATUS](#ManagedClusters_AgentPool_STATUS).

| Value     | Description |
|-----------|-------------|
| "Linux"   |             |
| "Windows" |             |

<a id="PowerState_Code_STATUS"></a>PowerState_Code_STATUS
---------------------------------------------------------

Used by: [PowerState_STATUS](#PowerState_STATUS), and [PowerState_STATUS_ARM](#PowerState_STATUS_ARM).

| Value     | Description |
|-----------|-------------|
| "Running" |             |
| "Stopped" |             |

<a id="PowerState_STATUS"></a>PowerState_STATUS
-----------------------------------------------

Describes the Power State of the cluster

Used by: [ManagedClusterAgentPoolProfile_STATUS](#ManagedClusterAgentPoolProfile_STATUS), [ManagedCluster_STATUS](#ManagedCluster_STATUS), and [ManagedClusters_AgentPool_STATUS](#ManagedClusters_AgentPool_STATUS).

| Property | Description                                     | Type                                              |
|----------|-------------------------------------------------|---------------------------------------------------|
| code     | Tells whether the cluster is Running or Stopped | [PowerState_Code_STATUS](#PowerState_Code_STATUS) |

<a id="PowerState_STATUS_ARM"></a>PowerState_STATUS_ARM
-------------------------------------------------------

Describes the Power State of the cluster

Used by: [ManagedClusterAgentPoolProfileProperties_STATUS_ARM](#ManagedClusterAgentPoolProfileProperties_STATUS_ARM), [ManagedClusterAgentPoolProfile_STATUS_ARM](#ManagedClusterAgentPoolProfile_STATUS_ARM), and [ManagedClusterProperties_STATUS_ARM](#ManagedClusterProperties_STATUS_ARM).

| Property | Description                                     | Type                                              |
|----------|-------------------------------------------------|---------------------------------------------------|
| code     | Tells whether the cluster is Running or Stopped | [PowerState_Code_STATUS](#PowerState_Code_STATUS) |

<a id="PrivateLinkResource"></a>PrivateLinkResource
---------------------------------------------------

A private link resource

Used by: [ManagedCluster_Spec](#ManagedCluster_Spec).

| Property        | Description                            | Type                         |
|-----------------|----------------------------------------|------------------------------|
| groupId         | The group ID of the resource.          | string                       |
| name            | The name of the private link resource. | string                       |
| reference       | The ID of the private link resource.   | genruntime.ResourceReference |
| requiredMembers | The RequiredMembers of the resource    | string[]                     |
| type            | The resource type.                     | string                       |

<a id="PrivateLinkResource_ARM"></a>PrivateLinkResource_ARM
-----------------------------------------------------------

A private link resource

Used by: [ManagedClusterProperties_ARM](#ManagedClusterProperties_ARM).

| Property        | Description                            | Type     |
|-----------------|----------------------------------------|----------|
| groupId         | The group ID of the resource.          | string   |
| id              |                                        | string   |
| name            | The name of the private link resource. | string   |
| requiredMembers | The RequiredMembers of the resource    | string[] |
| type            | The resource type.                     | string   |

<a id="PrivateLinkResource_STATUS"></a>PrivateLinkResource_STATUS
-----------------------------------------------------------------

A private link resource

Used by: [ManagedCluster_STATUS](#ManagedCluster_STATUS).

| Property             | Description                                                                                | Type     |
|----------------------|--------------------------------------------------------------------------------------------|----------|
| groupId              | The group ID of the resource.                                                              | string   |
| id                   | The ID of the private link resource.                                                       | string   |
| name                 | The name of the private link resource.                                                     | string   |
| privateLinkServiceID | The private link service ID of the resource, this field is exposed only to NRP internally. | string   |
| requiredMembers      | The RequiredMembers of the resource                                                        | string[] |
| type                 | The resource type.                                                                         | string   |

<a id="PrivateLinkResource_STATUS_ARM"></a>PrivateLinkResource_STATUS_ARM
-------------------------------------------------------------------------

A private link resource

Used by: [ManagedClusterProperties_STATUS_ARM](#ManagedClusterProperties_STATUS_ARM).

| Property             | Description                                                                                | Type     |
|----------------------|--------------------------------------------------------------------------------------------|----------|
| groupId              | The group ID of the resource.                                                              | string   |
| id                   | The ID of the private link resource.                                                       | string   |
| name                 | The name of the private link resource.                                                     | string   |
| privateLinkServiceID | The private link service ID of the resource, this field is exposed only to NRP internally. | string   |
| requiredMembers      | The RequiredMembers of the resource                                                        | string[] |
| type                 | The resource type.                                                                         | string   |

<a id="ResourceReference"></a>ResourceReference
-----------------------------------------------

A reference to an Azure resource.

Used by: [ManagedClusterLoadBalancerProfile](#ManagedClusterLoadBalancerProfile), [ManagedClusterLoadBalancerProfile_OutboundIPPrefixes](#ManagedClusterLoadBalancerProfile_OutboundIPPrefixes), and [ManagedClusterLoadBalancerProfile_OutboundIPs](#ManagedClusterLoadBalancerProfile_OutboundIPs).

| Property  | Description                            | Type                         |
|-----------|----------------------------------------|------------------------------|
| reference | The fully qualified Azure resource id. | genruntime.ResourceReference |

<a id="ResourceReference_ARM"></a>ResourceReference_ARM
-------------------------------------------------------

A reference to an Azure resource.

Used by: [ManagedClusterLoadBalancerProfile_ARM](#ManagedClusterLoadBalancerProfile_ARM), [ManagedClusterLoadBalancerProfile_OutboundIPPrefixes_ARM](#ManagedClusterLoadBalancerProfile_OutboundIPPrefixes_ARM), and [ManagedClusterLoadBalancerProfile_OutboundIPs_ARM](#ManagedClusterLoadBalancerProfile_OutboundIPs_ARM).

| Property | Description | Type   |
|----------|-------------|--------|
| id       |             | string |

<a id="ResourceReference_STATUS"></a>ResourceReference_STATUS
-------------------------------------------------------------

A reference to an Azure resource.

Used by: [ManagedClusterLoadBalancerProfile_OutboundIPPrefixes_STATUS](#ManagedClusterLoadBalancerProfile_OutboundIPPrefixes_STATUS), [ManagedClusterLoadBalancerProfile_OutboundIPs_STATUS](#ManagedClusterLoadBalancerProfile_OutboundIPs_STATUS), and [ManagedClusterLoadBalancerProfile_STATUS](#ManagedClusterLoadBalancerProfile_STATUS).

| Property | Description                            | Type   |
|----------|----------------------------------------|--------|
| id       | The fully qualified Azure resource id. | string |

<a id="ResourceReference_STATUS_ARM"></a>ResourceReference_STATUS_ARM
---------------------------------------------------------------------

A reference to an Azure resource.

Used by: [ManagedClusterLoadBalancerProfile_OutboundIPPrefixes_STATUS_ARM](#ManagedClusterLoadBalancerProfile_OutboundIPPrefixes_STATUS_ARM), [ManagedClusterLoadBalancerProfile_OutboundIPs_STATUS_ARM](#ManagedClusterLoadBalancerProfile_OutboundIPs_STATUS_ARM), and [ManagedClusterLoadBalancerProfile_STATUS_ARM](#ManagedClusterLoadBalancerProfile_STATUS_ARM).

| Property | Description                            | Type   |
|----------|----------------------------------------|--------|
| id       | The fully qualified Azure resource id. | string |

<a id="ScaleSetEvictionPolicy"></a>ScaleSetEvictionPolicy
---------------------------------------------------------

Used by: [ManagedClusterAgentPoolProfile](#ManagedClusterAgentPoolProfile), [ManagedClusterAgentPoolProfileProperties_ARM](#ManagedClusterAgentPoolProfileProperties_ARM), [ManagedClusterAgentPoolProfile_ARM](#ManagedClusterAgentPoolProfile_ARM), and [ManagedClusters_AgentPool_Spec](#ManagedClusters_AgentPool_Spec).

| Value        | Description |
|--------------|-------------|
| "Deallocate" |             |
| "Delete"     |             |

<a id="ScaleSetEvictionPolicy_STATUS"></a>ScaleSetEvictionPolicy_STATUS
-----------------------------------------------------------------------

Used by: [ManagedClusterAgentPoolProfileProperties_STATUS_ARM](#ManagedClusterAgentPoolProfileProperties_STATUS_ARM), [ManagedClusterAgentPoolProfile_STATUS](#ManagedClusterAgentPoolProfile_STATUS), [ManagedClusterAgentPoolProfile_STATUS_ARM](#ManagedClusterAgentPoolProfile_STATUS_ARM), and [ManagedClusters_AgentPool_STATUS](#ManagedClusters_AgentPool_STATUS).

| Value        | Description |
|--------------|-------------|
| "Deallocate" |             |
| "Delete"     |             |

<a id="ScaleSetPriority"></a>ScaleSetPriority
---------------------------------------------

Used by: [ManagedClusterAgentPoolProfile](#ManagedClusterAgentPoolProfile), [ManagedClusterAgentPoolProfileProperties_ARM](#ManagedClusterAgentPoolProfileProperties_ARM), [ManagedClusterAgentPoolProfile_ARM](#ManagedClusterAgentPoolProfile_ARM), and [ManagedClusters_AgentPool_Spec](#ManagedClusters_AgentPool_Spec).

| Value     | Description |
|-----------|-------------|
| "Regular" |             |
| "Spot"    |             |

<a id="ScaleSetPriority_STATUS"></a>ScaleSetPriority_STATUS
-----------------------------------------------------------

Used by: [ManagedClusterAgentPoolProfileProperties_STATUS_ARM](#ManagedClusterAgentPoolProfileProperties_STATUS_ARM), [ManagedClusterAgentPoolProfile_STATUS](#ManagedClusterAgentPoolProfile_STATUS), [ManagedClusterAgentPoolProfile_STATUS_ARM](#ManagedClusterAgentPoolProfile_STATUS_ARM), and [ManagedClusters_AgentPool_STATUS](#ManagedClusters_AgentPool_STATUS).

| Value     | Description |
|-----------|-------------|
| "Regular" |             |
| "Spot"    |             |

<a id="SysctlConfig"></a>SysctlConfig
-------------------------------------

Sysctl settings for Linux agent nodes.

Used by: [LinuxOSConfig](#LinuxOSConfig).

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

<a id="SysctlConfig_ARM"></a>SysctlConfig_ARM
---------------------------------------------

Sysctl settings for Linux agent nodes.

Used by: [LinuxOSConfig_ARM](#LinuxOSConfig_ARM).

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

<a id="SysctlConfig_STATUS"></a>SysctlConfig_STATUS
---------------------------------------------------

Sysctl settings for Linux agent nodes.

Used by: [LinuxOSConfig_STATUS](#LinuxOSConfig_STATUS).

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

<a id="SysctlConfig_STATUS_ARM"></a>SysctlConfig_STATUS_ARM
-----------------------------------------------------------

Sysctl settings for Linux agent nodes.

Used by: [LinuxOSConfig_STATUS_ARM](#LinuxOSConfig_STATUS_ARM).

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

<a id="UserAssignedIdentity"></a>UserAssignedIdentity
-----------------------------------------------------

Details about a user assigned identity.

Used by: [ManagedClusterPodIdentity](#ManagedClusterPodIdentity), and [ManagedCluster_Spec](#ManagedCluster_Spec).

| Property          | Description                                    | Type                         |
|-------------------|------------------------------------------------|------------------------------|
| clientId          | The client ID of the user assigned identity.   | string                       |
| objectId          | The object ID of the user assigned identity.   | string                       |
| resourceReference | The resource ID of the user assigned identity. | genruntime.ResourceReference |

<a id="UserAssignedIdentityDetails"></a>UserAssignedIdentityDetails
-------------------------------------------------------------------

Information about the user assigned identity for the resource

Used by: [ManagedClusterIdentity](#ManagedClusterIdentity).

| Property  | Description | Type                         |
|-----------|-------------|------------------------------|
| reference |             | genruntime.ResourceReference |

<a id="UserAssignedIdentity_ARM"></a>UserAssignedIdentity_ARM
-------------------------------------------------------------

Details about a user assigned identity.

Used by: [ManagedClusterPodIdentity_ARM](#ManagedClusterPodIdentity_ARM), and [ManagedClusterProperties_ARM](#ManagedClusterProperties_ARM).

| Property   | Description                                  | Type   |
|------------|----------------------------------------------|--------|
| clientId   | The client ID of the user assigned identity. | string |
| objectId   | The object ID of the user assigned identity. | string |
| resourceId |                                              | string |

<a id="UserAssignedIdentity_STATUS"></a>UserAssignedIdentity_STATUS
-------------------------------------------------------------------

Details about a user assigned identity.

Used by: [ManagedClusterAddonProfile_STATUS](#ManagedClusterAddonProfile_STATUS), [ManagedClusterPodIdentity_STATUS](#ManagedClusterPodIdentity_STATUS), and [ManagedCluster_STATUS](#ManagedCluster_STATUS).

| Property   | Description                                    | Type   |
|------------|------------------------------------------------|--------|
| clientId   | The client ID of the user assigned identity.   | string |
| objectId   | The object ID of the user assigned identity.   | string |
| resourceId | The resource ID of the user assigned identity. | string |

<a id="UserAssignedIdentity_STATUS_ARM"></a>UserAssignedIdentity_STATUS_ARM
---------------------------------------------------------------------------

Details about a user assigned identity.

Used by: [ManagedClusterAddonProfile_STATUS_ARM](#ManagedClusterAddonProfile_STATUS_ARM), [ManagedClusterPodIdentity_STATUS_ARM](#ManagedClusterPodIdentity_STATUS_ARM), and [ManagedClusterProperties_STATUS_ARM](#ManagedClusterProperties_STATUS_ARM).

| Property   | Description                                    | Type   |
|------------|------------------------------------------------|--------|
| clientId   | The client ID of the user assigned identity.   | string |
| objectId   | The object ID of the user assigned identity.   | string |
| resourceId | The resource ID of the user assigned identity. | string |
