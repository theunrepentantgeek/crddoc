v1beta1
=======

| Metadata | Value                           |
|----------|---------------------------------|
| Group    | infrastructure.cluster.x-k8s.io |
| Version  | v1beta1                         |

<a id="AADProfile"></a>AADProfile
---------------------------------

AADProfile - AAD integration managed by AKS. See also [AKS doc](https://learn.microsoft.com/azure/aks/managed-aad). <br/>

Used by: [AzureManagedControlPlaneSpec](#AzureManagedControlPlaneSpec).

| Property            | Description                                                                          | Type                  |
|---------------------|--------------------------------------------------------------------------------------|-----------------------|
| adminGroupObjectIDs | AdminGroupObjectIDs - AAD group object IDs that will have admin role of the cluster. | string[]<br/>Required |
| managed             | Managed - Whether to enable managed AAD.                                             | bool<br/>Required     |

<a id="AKSSku"></a>AKSSku
-------------------------

AKSSku - AKS SKU.

Used by: [AzureManagedControlPlaneSpec](#AzureManagedControlPlaneSpec).

| Property | Description                    | Type                                                                |
|----------|--------------------------------|---------------------------------------------------------------------|
| tier     | Tier - Tier of an AKS cluster. | [AzureManagedControlPlaneSkuTier](#AzureManagedControlPlaneSkuTier) |

<a id="APIServerAccessProfile"></a>APIServerAccessProfile
---------------------------------------------------------

APIServerAccessProfile tunes the accessibility of the cluster's control plane. See also [AKS doc](https://learn.microsoft.com/azure/aks/api-server-authorized-ip-ranges). <br/>

Used by: [AzureManagedControlPlaneSpec](#AzureManagedControlPlaneSpec).

| Property                       | Description                                                                                           | Type     |
|--------------------------------|-------------------------------------------------------------------------------------------------------|----------|
| authorizedIPRanges             | AuthorizedIPRanges - Authorized IP Ranges to kubernetes API server.                                   | string[] |
| enablePrivateCluster           | EnablePrivateCluster - Whether to create the cluster as a private cluster or not.                     | bool     |
| enablePrivateClusterPublicFQDN | EnablePrivateClusterPublicFQDN - Whether to create additional public FQDN for private cluster or not. | bool     |
| privateDNSZone                 | PrivateDNSZone - Private dns zone mode for private cluster.                                           | string   |

<a id="AdditionalCapabilities"></a>AdditionalCapabilities
---------------------------------------------------------

AdditionalCapabilities enables or disables a capability on the virtual machine.

Used by: [AzureMachineSpec](#AzureMachineSpec).

| Property        | Description                                                                                                                                                                                       | Type |
|-----------------|---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|------|
| ultraSSDEnabled | UltraSSDEnabled enables or disables Azure UltraSSD capability for the virtual machine. Defaults to true if Ultra SSD data disks are specified, otherwise it doesn't set the capability on the VM. | bool |

<a id="AddonProfile"></a>AddonProfile
-------------------------------------

AddonProfile represents a managed cluster add-on.

Used by: [AzureManagedControlPlaneSpec](#AzureManagedControlPlaneSpec).

| Property | Description                                          | Type              |
|----------|------------------------------------------------------|-------------------|
| config   | Config - Key-value pairs for configuring the add-on. | map[string]string |
| enabled  | Enabled - Whether the add-on is enabled or not.      | bool              |
| name     | Name - The name of the managed cluster add-on.       | string            |

<a id="AddressRecord"></a>AddressRecord
---------------------------------------

AddressRecord specifies a DNS record mapping a hostname to an IPV4 or IPv6 address.

| Property | Description | Type   |
|----------|-------------|--------|
| Hostname |             | string |
| IP       |             | string |

<a id="AllowedNamespaces"></a>AllowedNamespaces
-----------------------------------------------

AllowedNamespaces defines the namespaces the clusters are allowed to use the identity from NamespaceList takes precedence over the Selector.

Used by: [AzureClusterIdentitySpec](#AzureClusterIdentitySpec).

| Property | Description                                                                                                                                                                                                                                                                                                                                        | Type                 |
|----------|----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|----------------------|
| list     | A nil or empty list indicates that AzureCluster cannot use the identity from any namespace. <br/>                                                                                                                                                                                                                                                  | string[]             |
| selector | Selector is a selector of namespaces that AzureCluster can use this Identity from. This is a standard Kubernetes LabelSelector, a label query over a set of resources. The result of matchLabels and matchExpressions are ANDed. <br/>A nil or empty selector indicates that AzureCluster cannot use this AzureClusterIdentity from any namespace. | metav1.LabelSelector |

<a id="AutoScalerProfile"></a>AutoScalerProfile
-----------------------------------------------

AutoScalerProfile parameters to be applied to the cluster-autoscaler. See also [AKS doc](https://learn.microsoft.com/azure/aks/cluster-autoscaler#use-the-cluster-autoscaler-profile), [K8s doc](https://github.com/kubernetes/autoscaler/blob/master/cluster-autoscaler/FAQ.md#what-are-the-parameters-to-ca). <br/>

Used by: [AzureManagedControlPlaneSpec](#AzureManagedControlPlaneSpec).

| Property                      | Description                                                                                                                                                                                                                                                                                                                                              | Type                                                    |
|-------------------------------|----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|---------------------------------------------------------|
| balanceSimilarNodeGroups      | BalanceSimilarNodeGroups - Valid values are 'true' and 'false'. The default is false.                                                                                                                                                                                                                                                                    | [BalanceSimilarNodeGroups](#BalanceSimilarNodeGroups)   |
| expander                      | Expander - If not specified, the default is 'random'. See [expanders](https://github.com/kubernetes/autoscaler/blob/master/cluster-autoscaler/FAQ.md#what-are-expanders) for more information.                                                                                                                                                           | [Expander](#Expander)                                   |
| maxEmptyBulkDelete            | MaxEmptyBulkDelete - The default is 10.                                                                                                                                                                                                                                                                                                                  | string                                                  |
| maxGracefulTerminationSec     | MaxGracefulTerminationSec - The default is 600.                                                                                                                                                                                                                                                                                                          | string                                                  |
| maxNodeProvisionTime          | MaxNodeProvisionTime - The default is '15m'. Values must be an integer followed by an 'm'. No unit of time other than minutes (m) is supported.                                                                                                                                                                                                          | string                                                  |
| maxTotalUnreadyPercentage     | MaxTotalUnreadyPercentage - The default is 45. The maximum is 100 and the minimum is 0.                                                                                                                                                                                                                                                                  | string                                                  |
| newPodScaleUpDelay            | NewPodScaleUpDelay - For scenarios like burst/batch scale where you don't want CA to act before the kubernetes scheduler could schedule all the pods, you can tell CA to ignore unscheduled pods before they're a certain age. The default is '0s'. Values must be an integer followed by a unit ('s' for seconds, 'm' for minutes, 'h' for hours, etc). | string                                                  |
| okTotalUnreadyCount           | OkTotalUnreadyCount - This must be an integer. The default is 3.                                                                                                                                                                                                                                                                                         | string                                                  |
| scaleDownDelayAfterAdd        | ScaleDownDelayAfterAdd - The default is '10m'. Values must be an integer followed by an 'm'. No unit of time other than minutes (m) is supported.                                                                                                                                                                                                        | string                                                  |
| scaleDownDelayAfterDelete     | ScaleDownDelayAfterDelete - The default is the scan-interval. Values must be an integer followed by an 's'. No unit of time other than seconds (s) is supported.                                                                                                                                                                                         | string                                                  |
| scaleDownDelayAfterFailure    | ScaleDownDelayAfterFailure - The default is '3m'. Values must be an integer followed by an 'm'. No unit of time other than minutes (m) is supported.                                                                                                                                                                                                     | string                                                  |
| scaleDownUnneededTime         | ScaleDownUnneededTime - The default is '10m'. Values must be an integer followed by an 'm'. No unit of time other than minutes (m) is supported.                                                                                                                                                                                                         | string                                                  |
| scaleDownUnreadyTime          | ScaleDownUnreadyTime - The default is '20m'. Values must be an integer followed by an 'm'. No unit of time other than minutes (m) is supported.                                                                                                                                                                                                          | string                                                  |
| scaleDownUtilizationThreshold | ScaleDownUtilizationThreshold - The default is '0.5'.                                                                                                                                                                                                                                                                                                    | string                                                  |
| scanInterval                  | ScanInterval - How often cluster is reevaluated for scale up or down. The default is '10s'.                                                                                                                                                                                                                                                              | string                                                  |
| skipNodesWithLocalStorage     | SkipNodesWithLocalStorage - The default is false.                                                                                                                                                                                                                                                                                                        | [SkipNodesWithLocalStorage](#SkipNodesWithLocalStorage) |
| skipNodesWithSystemPods       | SkipNodesWithSystemPods - The default is true.                                                                                                                                                                                                                                                                                                           | [SkipNodesWithSystemPods](#SkipNodesWithSystemPods)     |

<a id="AzureBastion"></a>AzureBastion
-------------------------------------

AzureBastion specifies how the Azure Bastion cloud component should be configured.

Used by: [BastionSpec](#BastionSpec).

| Property        | Description                                                                                                           | Type                                      |
|-----------------|-----------------------------------------------------------------------------------------------------------------------|-------------------------------------------|
| enableTunneling | EnableTunneling enables the native client support feature for the Azure Bastion Host. Defaults to false.              | bool                                      |
| name            |                                                                                                                       | string                                    |
| publicIP        |                                                                                                                       | [PublicIPSpec](#PublicIPSpec)             |
| sku             | BastionHostSkuName configures the tier of the Azure Bastion Host. Can be either Basic or Standard. Defaults to Basic. | [BastionHostSkuName](#BastionHostSkuName) |
| subnet          |                                                                                                                       | [SubnetSpec](#SubnetSpec)                 |

<a id="AzureBastionTemplateSpec"></a>AzureBastionTemplateSpec
-------------------------------------------------------------

AzureBastionTemplateSpec specifies a template for an Azure Bastion host.

Used by: [BastionTemplateSpec](#BastionTemplateSpec).

| Property | Description | Type                                      |
|----------|-------------|-------------------------------------------|
| subnet   |             | [SubnetTemplateSpec](#SubnetTemplateSpec) |

<a id="AzureCluster"></a>AzureCluster
-------------------------------------

AzureCluster is the Schema for the azureclusters API.

Used by: [AzureClusterList](#AzureClusterList).

| Property          | Description | Type                                      |
|-------------------|-------------|-------------------------------------------|
| metav1.TypeMeta   |             |                                           |
| metav1.ObjectMeta |             |                                           |
| spec              |             | [AzureClusterSpec](#AzureClusterSpec)     |
| status            |             | [AzureClusterStatus](#AzureClusterStatus) |

### <a id="AzureClusterSpec"></a>AzureClusterSpec

| Property                                        | Description                                                                                                                                                                                                                             | Type                        |
|-------------------------------------------------|-----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|-----------------------------|
| [AzureClusterClassSpec](#AzureClusterClassSpec) |                                                                                                                                                                                                                                         |                             |
| bastionSpec                                     | BastionSpec encapsulates all things related to the Bastions in the cluster.                                                                                                                                                             | [BastionSpec](#BastionSpec) |
| controlPlaneEndpoint                            | ControlPlaneEndpoint represents the endpoint used to communicate with the control plane. It is not recommended to set this when creating an AzureCluster as CAPZ will set this for you. However, if it is set, CAPZ will not change it. | clusterv1.APIEndpoint       |
| networkSpec                                     | NetworkSpec encapsulates all things related to Azure network.                                                                                                                                                                           | [NetworkSpec](#NetworkSpec) |
| resourceGroup                                   |                                                                                                                                                                                                                                         | string                      |

### <a id="AzureClusterStatus"></a>AzureClusterStatus

| Property                   | Description                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                              | Type                     |
|----------------------------|--------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|--------------------------|
| conditions                 | Conditions defines current service state of the AzureCluster.                                                                                                                                                                                                                                                                                                                                                                                                                                                            | clusterv1.Conditions     |
| failureDomains             | FailureDomains specifies the list of unique failure domains for the location/region of the cluster. A FailureDomain maps to Availability Zone with an Azure Region (if the region support them). An Availability Zone is a separate data center within a region and they can be used to ensure the cluster is more resilient to failure. See: https://learn.microsoft.com/azure/reliability/availability-zones-overview This list will be used by Cluster API to try and spread the machines across the failure domains. | clusterv1.FailureDomains |
| longRunningOperationStates | LongRunningOperationStates saves the states for Azure long-running operations so they can be continued on the next reconciliation loop.                                                                                                                                                                                                                                                                                                                                                                                  | Futures                  |
| ready                      | Ready is true when the provider resource is ready.                                                                                                                                                                                                                                                                                                                                                                                                                                                                       | bool                     |

<a id="AzureClusterClassSpec"></a>AzureClusterClassSpec
-------------------------------------------------------

AzureClusterClassSpec defines the AzureCluster properties that may be shared across several Azure clusters.

| Property                     | Description                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                  | Type                                                          |
|------------------------------|------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|---------------------------------------------------------------|
| additionalTags               | AdditionalTags is an optional set of tags to add to Azure resources managed by the Azure provider, in addition to the ones added by default.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                 | Tags                                                          |
| azureEnvironment             | AzureEnvironment is the name of the AzureCloud to be used. The default value that would be used by most users is "AzurePublicCloud", other values are: - ChinaCloud: "AzureChinaCloud" - GermanCloud: "AzureGermanCloud" - PublicCloud: "AzurePublicCloud" - USGovernmentCloud: "AzureUSGovernmentCloud"                                                                                                                                                                                                                                                                                                                     | string                                                        |
| cloudProviderConfigOverrides | CloudProviderConfigOverrides is an optional set of configuration values that can be overridden in azure cloud provider config. This is only a subset of options that are available in azure cloud provider config. Some values for the cloud provider config are inferred from other parts of cluster api provider azure spec, and may not be available for overrides. See: https://cloud-provider-azure.sigs.k8s.io/install/configs Note: All cloud provider config values can be customized by creating the secret beforehand. CloudProviderConfigOverrides is only used when the secret is managed by the Azure Provider. | [CloudProviderConfigOverrides](#CloudProviderConfigOverrides) |
| extendedLocation             | ExtendedLocation is an optional set of ExtendedLocation properties for clusters on Azure public MEC.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                         | [ExtendedLocationSpec](#ExtendedLocationSpec)                 |
| identityRef                  | IdentityRef is a reference to an AzureIdentity to be used when reconciling this cluster                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                      | corev1.ObjectReference                                        |
| location                     |                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                              | string                                                        |
| subscriptionID               |                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                              | string                                                        |

<a id="AzureClusterIdentity"></a>AzureClusterIdentity
-----------------------------------------------------

AzureClusterIdentity is the Schema for the azureclustersidentities API.

Used by: [AzureClusterIdentityList](#AzureClusterIdentityList).

| Property          | Description | Type                                                      |
|-------------------|-------------|-----------------------------------------------------------|
| metav1.TypeMeta   |             |                                                           |
| metav1.ObjectMeta |             |                                                           |
| spec              |             | [AzureClusterIdentitySpec](#AzureClusterIdentitySpec)     |
| status            |             | [AzureClusterIdentityStatus](#AzureClusterIdentityStatus) |

### <a id="AzureClusterIdentitySpec"></a>AzureClusterIdentitySpec

| Property          | Description                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                         | Type                                    |
|-------------------|---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|-----------------------------------------|
| allowedNamespaces | AllowedNamespaces is used to identify the namespaces the clusters are allowed to use the identity from. Namespaces can be selected either using an array of namespaces or with label selector. An empty allowedNamespaces object indicates that AzureClusters can use this identity from any namespace. If this object is nil, no namespaces will be allowed (default behaviour, if this field is not provided) A namespace should be either in the NamespaceList or match with Selector to use the identity. <br/> | [AllowedNamespaces](#AllowedNamespaces) |
| clientID          | ClientID is the service principal client ID. Both User Assigned MSI and SP can use this field.                                                                                                                                                                                                                                                                                                                                                                                                                      | string                                  |
| clientSecret      | ClientSecret is a secret reference which should contain either a Service Principal password or certificate secret.                                                                                                                                                                                                                                                                                                                                                                                                  | corev1.SecretReference                  |
| resourceID        | ResourceID is the Azure resource ID for the User Assigned MSI resource. Only applicable when type is UserAssignedMSI.                                                                                                                                                                                                                                                                                                                                                                                               | string                                  |
| tenantID          | TenantID is the service principal primary tenant id.                                                                                                                                                                                                                                                                                                                                                                                                                                                                | string                                  |
| type              | Type is the type of Azure Identity used. ServicePrincipal, ServicePrincipalCertificate, UserAssignedMSI, ManualServicePrincipal or WorkloadIdentity.                                                                                                                                                                                                                                                                                                                                                                | [IdentityType](#IdentityType)           |

### <a id="AzureClusterIdentityStatus"></a>AzureClusterIdentityStatus

| Property   | Description                                                           | Type                 |
|------------|-----------------------------------------------------------------------|----------------------|
| conditions | Conditions defines current service state of the AzureClusterIdentity. | clusterv1.Conditions |

<a id="AzureClusterIdentityList"></a>AzureClusterIdentityList
-------------------------------------------------------------

AzureClusterIdentityList contains a list of AzureClusterIdentity.

| Property        | Description | Type                                            |
|-----------------|-------------|-------------------------------------------------|
| metav1.TypeMeta |             |                                                 |
| metav1.ListMeta |             |                                                 |
| items           |             | [AzureClusterIdentity[]](#AzureClusterIdentity) |

<a id="AzureClusterIdentitySpec"></a>AzureClusterIdentitySpec
-------------------------------------------------------------

AzureClusterIdentitySpec defines the parameters that are used to create an AzureIdentity.

Used by: [AzureClusterIdentity](#AzureClusterIdentity).

| Property          | Description                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                         | Type                                    |
|-------------------|---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|-----------------------------------------|
| allowedNamespaces | AllowedNamespaces is used to identify the namespaces the clusters are allowed to use the identity from. Namespaces can be selected either using an array of namespaces or with label selector. An empty allowedNamespaces object indicates that AzureClusters can use this identity from any namespace. If this object is nil, no namespaces will be allowed (default behaviour, if this field is not provided) A namespace should be either in the NamespaceList or match with Selector to use the identity. <br/> | [AllowedNamespaces](#AllowedNamespaces) |
| clientID          | ClientID is the service principal client ID. Both User Assigned MSI and SP can use this field.                                                                                                                                                                                                                                                                                                                                                                                                                      | string                                  |
| clientSecret      | ClientSecret is a secret reference which should contain either a Service Principal password or certificate secret.                                                                                                                                                                                                                                                                                                                                                                                                  | corev1.SecretReference                  |
| resourceID        | ResourceID is the Azure resource ID for the User Assigned MSI resource. Only applicable when type is UserAssignedMSI.                                                                                                                                                                                                                                                                                                                                                                                               | string                                  |
| tenantID          | TenantID is the service principal primary tenant id.                                                                                                                                                                                                                                                                                                                                                                                                                                                                | string                                  |
| type              | Type is the type of Azure Identity used. ServicePrincipal, ServicePrincipalCertificate, UserAssignedMSI, ManualServicePrincipal or WorkloadIdentity.                                                                                                                                                                                                                                                                                                                                                                | [IdentityType](#IdentityType)           |

<a id="AzureClusterIdentityStatus"></a>AzureClusterIdentityStatus
-----------------------------------------------------------------

AzureClusterIdentityStatus defines the observed state of AzureClusterIdentity.

Used by: [AzureClusterIdentity](#AzureClusterIdentity).

| Property   | Description                                                           | Type                 |
|------------|-----------------------------------------------------------------------|----------------------|
| conditions | Conditions defines current service state of the AzureClusterIdentity. | clusterv1.Conditions |

<a id="AzureClusterList"></a>AzureClusterList
---------------------------------------------

AzureClusterList contains a list of AzureClusters.

| Property        | Description | Type                            |
|-----------------|-------------|---------------------------------|
| metav1.TypeMeta |             |                                 |
| metav1.ListMeta |             |                                 |
| items           |             | [AzureCluster[]](#AzureCluster) |

<a id="AzureClusterSpec"></a>AzureClusterSpec
---------------------------------------------

AzureClusterSpec defines the desired state of AzureCluster.

Used by: [AzureCluster](#AzureCluster).

| Property                                        | Description                                                                                                                                                                                                                             | Type                        |
|-------------------------------------------------|-----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|-----------------------------|
| [AzureClusterClassSpec](#AzureClusterClassSpec) |                                                                                                                                                                                                                                         |                             |
| bastionSpec                                     | BastionSpec encapsulates all things related to the Bastions in the cluster.                                                                                                                                                             | [BastionSpec](#BastionSpec) |
| controlPlaneEndpoint                            | ControlPlaneEndpoint represents the endpoint used to communicate with the control plane. It is not recommended to set this when creating an AzureCluster as CAPZ will set this for you. However, if it is set, CAPZ will not change it. | clusterv1.APIEndpoint       |
| networkSpec                                     | NetworkSpec encapsulates all things related to Azure network.                                                                                                                                                                           | [NetworkSpec](#NetworkSpec) |
| resourceGroup                                   |                                                                                                                                                                                                                                         | string                      |

<a id="AzureClusterStatus"></a>AzureClusterStatus
-------------------------------------------------

AzureClusterStatus defines the observed state of AzureCluster.

Used by: [AzureCluster](#AzureCluster).

| Property                   | Description                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                              | Type                     |
|----------------------------|--------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|--------------------------|
| conditions                 | Conditions defines current service state of the AzureCluster.                                                                                                                                                                                                                                                                                                                                                                                                                                                            | clusterv1.Conditions     |
| failureDomains             | FailureDomains specifies the list of unique failure domains for the location/region of the cluster. A FailureDomain maps to Availability Zone with an Azure Region (if the region support them). An Availability Zone is a separate data center within a region and they can be used to ensure the cluster is more resilient to failure. See: https://learn.microsoft.com/azure/reliability/availability-zones-overview This list will be used by Cluster API to try and spread the machines across the failure domains. | clusterv1.FailureDomains |
| longRunningOperationStates | LongRunningOperationStates saves the states for Azure long-running operations so they can be continued on the next reconciliation loop.                                                                                                                                                                                                                                                                                                                                                                                  | Futures                  |
| ready                      | Ready is true when the provider resource is ready.                                                                                                                                                                                                                                                                                                                                                                                                                                                                       | bool                     |

<a id="AzureClusterTemplate"></a>AzureClusterTemplate
-----------------------------------------------------

AzureClusterTemplate is the Schema for the azureclustertemplates API.

Used by: [AzureClusterTemplateList](#AzureClusterTemplateList).

| Property          | Description | Type                                                  |
|-------------------|-------------|-------------------------------------------------------|
| metav1.TypeMeta   |             |                                                       |
| metav1.ObjectMeta |             |                                                       |
| spec              |             | [AzureClusterTemplateSpec](#AzureClusterTemplateSpec) |

<a id="AzureClusterTemplateList"></a>AzureClusterTemplateList
-------------------------------------------------------------

AzureClusterTemplateList contains a list of AzureClusterTemplate.

| Property        | Description | Type                                            |
|-----------------|-------------|-------------------------------------------------|
| metav1.TypeMeta |             |                                                 |
| metav1.ListMeta |             |                                                 |
| items           |             | [AzureClusterTemplate[]](#AzureClusterTemplate) |

<a id="AzureClusterTemplateResource"></a>AzureClusterTemplateResource
---------------------------------------------------------------------

AzureClusterTemplateResource describes the data needed to create an AzureCluster from a template.

Used by: [AzureClusterTemplateSpec](#AzureClusterTemplateSpec).

| Property | Description | Type                                                                  |
|----------|-------------|-----------------------------------------------------------------------|
| spec     |             | [AzureClusterTemplateResourceSpec](#AzureClusterTemplateResourceSpec) |

<a id="AzureClusterTemplateResourceSpec"></a>AzureClusterTemplateResourceSpec
-----------------------------------------------------------------------------

AzureClusterTemplateResourceSpec specifies an Azure cluster template resource.

Used by: [AzureClusterTemplateResource](#AzureClusterTemplateResource).

| Property                                        | Description                                                                 | Type                                        |
|-------------------------------------------------|-----------------------------------------------------------------------------|---------------------------------------------|
| [AzureClusterClassSpec](#AzureClusterClassSpec) |                                                                             |                                             |
| bastionSpec                                     | BastionSpec encapsulates all things related to the Bastions in the cluster. | [BastionTemplateSpec](#BastionTemplateSpec) |
| networkSpec                                     | NetworkSpec encapsulates all things related to Azure network.               | [NetworkTemplateSpec](#NetworkTemplateSpec) |

<a id="AzureClusterTemplateSpec"></a>AzureClusterTemplateSpec
-------------------------------------------------------------

AzureClusterTemplateSpec defines the desired state of AzureClusterTemplate.

Used by: [AzureClusterTemplate](#AzureClusterTemplate).

| Property | Description | Type                                                          |
|----------|-------------|---------------------------------------------------------------|
| template |             | [AzureClusterTemplateResource](#AzureClusterTemplateResource) |

<a id="AzureComputeGalleryImage"></a>AzureComputeGalleryImage
-------------------------------------------------------------

AzureComputeGalleryImage defines an image in the Azure Compute Gallery to use for VM creation.

Used by: [Image](#Image).

| Property       | Description                                                                                                                                                                                                                                                                                                                                                                  | Type                    |
|----------------|------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|-------------------------|
| gallery        | Gallery specifies the name of the compute image gallery that contains the image                                                                                                                                                                                                                                                                                              | string                  |
| name           | Name is the name of the image                                                                                                                                                                                                                                                                                                                                                | string                  |
| plan           | Plan contains plan information.                                                                                                                                                                                                                                                                                                                                              | [ImagePlan](#ImagePlan) |
| resourceGroup  | ResourceGroup specifies the resource group containing the private compute gallery.                                                                                                                                                                                                                                                                                           | string                  |
| subscriptionID | SubscriptionID is the identifier of the subscription that contains the private compute gallery.                                                                                                                                                                                                                                                                              | string                  |
| version        | Version specifies the version of the marketplace image. The allowed formats are Major.Minor.Build or 'latest'. Major, Minor, and Build are decimal numbers. Specify 'latest' to use the latest version of an image available at deploy time. Even if you use 'latest', the VM image will not automatically update after deploy time even if a new version becomes available. | string                  |

<a id="AzureMachine"></a>AzureMachine
-------------------------------------

AzureMachine is the Schema for the azuremachines API.

Used by: [AzureMachineList](#AzureMachineList).

| Property          | Description | Type                                      |
|-------------------|-------------|-------------------------------------------|
| metav1.TypeMeta   |             |                                           |
| metav1.ObjectMeta |             |                                           |
| spec              |             | [AzureMachineSpec](#AzureMachineSpec)     |
| status            |             | [AzureMachineStatus](#AzureMachineStatus) |

### <a id="AzureMachineSpec"></a>AzureMachineSpec

| Property                   | Description                                                                                                                                                                                                                                                                                                          | Type                                                      |
|----------------------------|----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|-----------------------------------------------------------|
| acceleratedNetworking      | Deprecated: AcceleratedNetworking should be set in the networkInterfaces field.                                                                                                                                                                                                                                      | bool                                                      |
| additionalCapabilities     | AdditionalCapabilities specifies additional capabilities enabled or disabled on the virtual machine.                                                                                                                                                                                                                 | [AdditionalCapabilities](#AdditionalCapabilities)         |
| additionalTags             | AdditionalTags is an optional set of tags to add to an instance, in addition to the ones added by default by the Azure provider. If both the AzureCluster and the AzureMachine specify the same tag name with different values, the AzureMachine's value takes precedence.                                           | Tags                                                      |
| allocatePublicIP           | AllocatePublicIP allows the ability to create dynamic public ips for machines where this value is true.                                                                                                                                                                                                              | bool                                                      |
| dataDisks                  | DataDisk specifies the parameters that are used to add one or more data disks to the machine                                                                                                                                                                                                                         | [DataDisk[]](#DataDisk)                                   |
| diagnostics                | Diagnostics specifies the diagnostics settings for a virtual machine. If not specified then Boot diagnostics (Managed) will be enabled.                                                                                                                                                                              | [Diagnostics](#Diagnostics)                               |
| dnsServers                 | DNSServers adds a list of DNS Server IP addresses to the VM NICs.                                                                                                                                                                                                                                                    | string[]                                                  |
| enableIPForwarding         | EnableIPForwarding enables IP Forwarding in Azure which is required for some CNI's to send traffic from a pods on one machine to another. This is required for IpV6 with Calico in combination with User Defined Routes (set by the Azure Cloud Controller manager). Default is false for disabled.                  | bool                                                      |
| failureDomain              | FailureDomain is the failure domain unique identifier this Machine should be attached to, as defined in Cluster API. This relates to an Azure Availability Zone                                                                                                                                                      | string                                                    |
| identity                   | Identity is the type of identity used for the virtual machine. The type 'SystemAssigned' is an implicitly created identity. The generated identity will be assigned a Subscription contributor role. The type 'UserAssigned' is a standalone Azure resource provided by the user and assigned to the VM              | [VMIdentity](#VMIdentity)                                 |
| image                      | Image is used to provide details of an image to use during VM creation. If image details are omitted the image will default the Azure Marketplace "capi" offer, which is based on Ubuntu.                                                                                                                            | [Image](#Image)                                           |
| networkInterfaces          | NetworkInterfaces specifies a list of network interface configurations. If left unspecified, the VM will get a single network interface with a single IPConfig in the subnet specified in the cluster's node subnet field. The primary interface will be the first networkInterface specified (index 0) in the list. | [NetworkInterface[]](#NetworkInterface)                   |
| osDisk                     | OSDisk specifies the parameters for the operating system disk of the machine                                                                                                                                                                                                                                         | [OSDisk](#OSDisk)                                         |
| providerID                 | ProviderID is the unique identifier as specified by the cloud provider.                                                                                                                                                                                                                                              | string                                                    |
| roleAssignmentName         | Deprecated: RoleAssignmentName should be set in the systemAssignedIdentityRole field.                                                                                                                                                                                                                                | string                                                    |
| securityProfile            | SecurityProfile specifies the Security profile settings for a virtual machine.                                                                                                                                                                                                                                       | [SecurityProfile](#SecurityProfile)                       |
| spotVMOptions              | SpotVMOptions allows the ability to specify the Machine should use a Spot VM                                                                                                                                                                                                                                         | [SpotVMOptions](#SpotVMOptions)                           |
| sshPublicKey               | SSHPublicKey is the SSH public key string, base64-encoded to add to a Virtual Machine. Linux only. Refer to documentation on how to set up SSH access on Windows instances.                                                                                                                                          | string                                                    |
| subnetName                 | Deprecated: SubnetName should be set in the networkInterfaces field.                                                                                                                                                                                                                                                 | string                                                    |
| systemAssignedIdentityRole | SystemAssignedIdentityRole defines the role and scope to assign to the system-assigned identity.                                                                                                                                                                                                                     | [SystemAssignedIdentityRole](#SystemAssignedIdentityRole) |
| userAssignedIdentities     | UserAssignedIdentities is a list of standalone Azure identities provided by the user The lifecycle of a user-assigned identity is managed separately from the lifecycle of the AzureMachine. See https://learn.microsoft.com/azure/active-directory/managed-identities-azure-resources/how-to-manage-ua-identity-cli | [UserAssignedIdentity[]](#UserAssignedIdentity)           |
| vmExtensions               | VMExtensions specifies a list of extensions to be added to the virtual machine.                                                                                                                                                                                                                                      | [VMExtension[]](#VMExtension)                             |
| vmSize                     |                                                                                                                                                                                                                                                                                                                      | string                                                    |

### <a id="AzureMachineStatus"></a>AzureMachineStatus

| Property                   | Description                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                   | Type                                    |
|----------------------------|-----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|-----------------------------------------|
| addresses                  | Addresses contains the Azure instance associated addresses.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                   | corev1.NodeAddress[]                    |
| conditions                 | Conditions defines current service state of the AzureMachine.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                 | clusterv1.Conditions                    |
| failureMessage             | ErrorMessage will be set in the event that there is a terminal problem reconciling the Machine and will contain a more verbose string suitable for logging and human consumption. <br/>This field should not be set for transitive errors that a controller faces that are expected to be fixed automatically over time (like service outages), but instead indicate that something is fundamentally wrong with the Machine's spec or the configuration of the controller, and that manual intervention is required. Examples of terminal errors would be invalid combinations of settings in the spec, values that are unsupported by the controller, or the responsible controller itself being critically misconfigured. <br/>Any transient errors that occur during the reconciliation of Machines can be added as events to the Machine object and/or logged in the controller's output. | string                                  |
| failureReason              | ErrorReason will be set in the event that there is a terminal problem reconciling the Machine and will contain a succinct value suitable for machine interpretation. <br/>This field should not be set for transitive errors that a controller faces that are expected to be fixed automatically over time (like service outages), but instead indicate that something is fundamentally wrong with the Machine's spec or the configuration of the controller, and that manual intervention is required. Examples of terminal errors would be invalid combinations of settings in the spec, values that are unsupported by the controller, or the responsible controller itself being critically misconfigured. <br/>Any transient errors that occur during the reconciliation of Machines can be added as events to the Machine object and/or logged in the controller's output.              | errors.MachineStatusError               |
| longRunningOperationStates | LongRunningOperationStates saves the states for Azure long-running operations so they can be continued on the next reconciliation loop.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                       | Futures                                 |
| ready                      | Ready is true when the provider resource is ready.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                            | bool                                    |
| vmState                    | VMState is the provisioning state of the Azure virtual machine.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                               | [ProvisioningState](#ProvisioningState) |

<a id="AzureMachineList"></a>AzureMachineList
---------------------------------------------

AzureMachineList contains a list of AzureMachine.

| Property        | Description | Type                            |
|-----------------|-------------|---------------------------------|
| metav1.TypeMeta |             |                                 |
| metav1.ListMeta |             |                                 |
| items           |             | [AzureMachine[]](#AzureMachine) |

<a id="AzureMachineSpec"></a>AzureMachineSpec
---------------------------------------------

AzureMachineSpec defines the desired state of AzureMachine.

Used by: [AzureMachine](#AzureMachine), and [AzureMachineTemplateResource](#AzureMachineTemplateResource).

| Property                   | Description                                                                                                                                                                                                                                                                                                          | Type                                                      |
|----------------------------|----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|-----------------------------------------------------------|
| acceleratedNetworking      | Deprecated: AcceleratedNetworking should be set in the networkInterfaces field.                                                                                                                                                                                                                                      | bool                                                      |
| additionalCapabilities     | AdditionalCapabilities specifies additional capabilities enabled or disabled on the virtual machine.                                                                                                                                                                                                                 | [AdditionalCapabilities](#AdditionalCapabilities)         |
| additionalTags             | AdditionalTags is an optional set of tags to add to an instance, in addition to the ones added by default by the Azure provider. If both the AzureCluster and the AzureMachine specify the same tag name with different values, the AzureMachine's value takes precedence.                                           | Tags                                                      |
| allocatePublicIP           | AllocatePublicIP allows the ability to create dynamic public ips for machines where this value is true.                                                                                                                                                                                                              | bool                                                      |
| dataDisks                  | DataDisk specifies the parameters that are used to add one or more data disks to the machine                                                                                                                                                                                                                         | [DataDisk[]](#DataDisk)                                   |
| diagnostics                | Diagnostics specifies the diagnostics settings for a virtual machine. If not specified then Boot diagnostics (Managed) will be enabled.                                                                                                                                                                              | [Diagnostics](#Diagnostics)                               |
| dnsServers                 | DNSServers adds a list of DNS Server IP addresses to the VM NICs.                                                                                                                                                                                                                                                    | string[]                                                  |
| enableIPForwarding         | EnableIPForwarding enables IP Forwarding in Azure which is required for some CNI's to send traffic from a pods on one machine to another. This is required for IpV6 with Calico in combination with User Defined Routes (set by the Azure Cloud Controller manager). Default is false for disabled.                  | bool                                                      |
| failureDomain              | FailureDomain is the failure domain unique identifier this Machine should be attached to, as defined in Cluster API. This relates to an Azure Availability Zone                                                                                                                                                      | string                                                    |
| identity                   | Identity is the type of identity used for the virtual machine. The type 'SystemAssigned' is an implicitly created identity. The generated identity will be assigned a Subscription contributor role. The type 'UserAssigned' is a standalone Azure resource provided by the user and assigned to the VM              | [VMIdentity](#VMIdentity)                                 |
| image                      | Image is used to provide details of an image to use during VM creation. If image details are omitted the image will default the Azure Marketplace "capi" offer, which is based on Ubuntu.                                                                                                                            | [Image](#Image)                                           |
| networkInterfaces          | NetworkInterfaces specifies a list of network interface configurations. If left unspecified, the VM will get a single network interface with a single IPConfig in the subnet specified in the cluster's node subnet field. The primary interface will be the first networkInterface specified (index 0) in the list. | [NetworkInterface[]](#NetworkInterface)                   |
| osDisk                     | OSDisk specifies the parameters for the operating system disk of the machine                                                                                                                                                                                                                                         | [OSDisk](#OSDisk)                                         |
| providerID                 | ProviderID is the unique identifier as specified by the cloud provider.                                                                                                                                                                                                                                              | string                                                    |
| roleAssignmentName         | Deprecated: RoleAssignmentName should be set in the systemAssignedIdentityRole field.                                                                                                                                                                                                                                | string                                                    |
| securityProfile            | SecurityProfile specifies the Security profile settings for a virtual machine.                                                                                                                                                                                                                                       | [SecurityProfile](#SecurityProfile)                       |
| spotVMOptions              | SpotVMOptions allows the ability to specify the Machine should use a Spot VM                                                                                                                                                                                                                                         | [SpotVMOptions](#SpotVMOptions)                           |
| sshPublicKey               | SSHPublicKey is the SSH public key string, base64-encoded to add to a Virtual Machine. Linux only. Refer to documentation on how to set up SSH access on Windows instances.                                                                                                                                          | string                                                    |
| subnetName                 | Deprecated: SubnetName should be set in the networkInterfaces field.                                                                                                                                                                                                                                                 | string                                                    |
| systemAssignedIdentityRole | SystemAssignedIdentityRole defines the role and scope to assign to the system-assigned identity.                                                                                                                                                                                                                     | [SystemAssignedIdentityRole](#SystemAssignedIdentityRole) |
| userAssignedIdentities     | UserAssignedIdentities is a list of standalone Azure identities provided by the user The lifecycle of a user-assigned identity is managed separately from the lifecycle of the AzureMachine. See https://learn.microsoft.com/azure/active-directory/managed-identities-azure-resources/how-to-manage-ua-identity-cli | [UserAssignedIdentity[]](#UserAssignedIdentity)           |
| vmExtensions               | VMExtensions specifies a list of extensions to be added to the virtual machine.                                                                                                                                                                                                                                      | [VMExtension[]](#VMExtension)                             |
| vmSize                     |                                                                                                                                                                                                                                                                                                                      | string                                                    |

<a id="AzureMachineStatus"></a>AzureMachineStatus
-------------------------------------------------

AzureMachineStatus defines the observed state of AzureMachine.

Used by: [AzureMachine](#AzureMachine).

| Property                   | Description                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                   | Type                                    |
|----------------------------|-----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|-----------------------------------------|
| addresses                  | Addresses contains the Azure instance associated addresses.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                   | corev1.NodeAddress[]                    |
| conditions                 | Conditions defines current service state of the AzureMachine.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                 | clusterv1.Conditions                    |
| failureMessage             | ErrorMessage will be set in the event that there is a terminal problem reconciling the Machine and will contain a more verbose string suitable for logging and human consumption. <br/>This field should not be set for transitive errors that a controller faces that are expected to be fixed automatically over time (like service outages), but instead indicate that something is fundamentally wrong with the Machine's spec or the configuration of the controller, and that manual intervention is required. Examples of terminal errors would be invalid combinations of settings in the spec, values that are unsupported by the controller, or the responsible controller itself being critically misconfigured. <br/>Any transient errors that occur during the reconciliation of Machines can be added as events to the Machine object and/or logged in the controller's output. | string                                  |
| failureReason              | ErrorReason will be set in the event that there is a terminal problem reconciling the Machine and will contain a succinct value suitable for machine interpretation. <br/>This field should not be set for transitive errors that a controller faces that are expected to be fixed automatically over time (like service outages), but instead indicate that something is fundamentally wrong with the Machine's spec or the configuration of the controller, and that manual intervention is required. Examples of terminal errors would be invalid combinations of settings in the spec, values that are unsupported by the controller, or the responsible controller itself being critically misconfigured. <br/>Any transient errors that occur during the reconciliation of Machines can be added as events to the Machine object and/or logged in the controller's output.              | errors.MachineStatusError               |
| longRunningOperationStates | LongRunningOperationStates saves the states for Azure long-running operations so they can be continued on the next reconciliation loop.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                       | Futures                                 |
| ready                      | Ready is true when the provider resource is ready.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                            | bool                                    |
| vmState                    | VMState is the provisioning state of the Azure virtual machine.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                               | [ProvisioningState](#ProvisioningState) |

<a id="AzureMachineTemplate"></a>AzureMachineTemplate
-----------------------------------------------------

AzureMachineTemplate is the Schema for the azuremachinetemplates API.

Used by: [AzureMachineTemplateList](#AzureMachineTemplateList).

| Property          | Description | Type                                                  |
|-------------------|-------------|-------------------------------------------------------|
| metav1.TypeMeta   |             |                                                       |
| metav1.ObjectMeta |             |                                                       |
| spec              |             | [AzureMachineTemplateSpec](#AzureMachineTemplateSpec) |

<a id="AzureMachineTemplateList"></a>AzureMachineTemplateList
-------------------------------------------------------------

AzureMachineTemplateList contains a list of AzureMachineTemplates.

| Property        | Description | Type                                            |
|-----------------|-------------|-------------------------------------------------|
| metav1.TypeMeta |             |                                                 |
| metav1.ListMeta |             |                                                 |
| items           |             | [AzureMachineTemplate[]](#AzureMachineTemplate) |

<a id="AzureMachineTemplateResource"></a>AzureMachineTemplateResource
---------------------------------------------------------------------

AzureMachineTemplateResource describes the data needed to create an AzureMachine from a template.

Used by: [AzureMachineTemplateSpec](#AzureMachineTemplateSpec).

| Property | Description                                                       | Type                                  |
|----------|-------------------------------------------------------------------|---------------------------------------|
| metadata |                                                                   | clusterv1.ObjectMeta                  |
| spec     | Spec is the specification of the desired behavior of the machine. | [AzureMachineSpec](#AzureMachineSpec) |

<a id="AzureMachineTemplateSpec"></a>AzureMachineTemplateSpec
-------------------------------------------------------------

AzureMachineTemplateSpec defines the desired state of AzureMachineTemplate.

Used by: [AzureMachineTemplate](#AzureMachineTemplate).

| Property | Description | Type                                                          |
|----------|-------------|---------------------------------------------------------------|
| template |             | [AzureMachineTemplateResource](#AzureMachineTemplateResource) |

<a id="AzureManagedCluster"></a>AzureManagedCluster
---------------------------------------------------

AzureManagedCluster is the Schema for the azuremanagedclusters API.

Used by: [AzureManagedClusterList](#AzureManagedClusterList).

| Property          | Description | Type                                                    |
|-------------------|-------------|---------------------------------------------------------|
| metav1.TypeMeta   |             |                                                         |
| metav1.ObjectMeta |             |                                                         |
| spec              |             | [AzureManagedClusterSpec](#AzureManagedClusterSpec)     |
| status            |             | [AzureManagedClusterStatus](#AzureManagedClusterStatus) |

### <a id="AzureManagedClusterSpec"></a>AzureManagedClusterSpec

| Property             | Description                                                                                                                             | Type                  |
|----------------------|-----------------------------------------------------------------------------------------------------------------------------------------|-----------------------|
| controlPlaneEndpoint | ControlPlaneEndpoint represents the endpoint used to communicate with the control plane. Immutable, populated by the AKS API at create. | clusterv1.APIEndpoint |

### <a id="AzureManagedClusterStatus"></a>AzureManagedClusterStatus

| Property | Description                                        | Type |
|----------|----------------------------------------------------|------|
| ready    | Ready is true when the provider resource is ready. | bool |

<a id="AzureManagedClusterList"></a>AzureManagedClusterList
-----------------------------------------------------------

AzureManagedClusterList contains a list of AzureManagedClusters.

| Property        | Description | Type                                          |
|-----------------|-------------|-----------------------------------------------|
| metav1.TypeMeta |             |                                               |
| metav1.ListMeta |             |                                               |
| items           |             | [AzureManagedCluster[]](#AzureManagedCluster) |

<a id="AzureManagedClusterSpec"></a>AzureManagedClusterSpec
-----------------------------------------------------------

AzureManagedClusterSpec defines the desired state of AzureManagedCluster.

Used by: [AzureManagedCluster](#AzureManagedCluster).

| Property             | Description                                                                                                                             | Type                  |
|----------------------|-----------------------------------------------------------------------------------------------------------------------------------------|-----------------------|
| controlPlaneEndpoint | ControlPlaneEndpoint represents the endpoint used to communicate with the control plane. Immutable, populated by the AKS API at create. | clusterv1.APIEndpoint |

<a id="AzureManagedClusterStatus"></a>AzureManagedClusterStatus
---------------------------------------------------------------

AzureManagedClusterStatus defines the observed state of AzureManagedCluster.

Used by: [AzureManagedCluster](#AzureManagedCluster).

| Property | Description                                        | Type |
|----------|----------------------------------------------------|------|
| ready    | Ready is true when the provider resource is ready. | bool |

<a id="AzureManagedControlPlane"></a>AzureManagedControlPlane
-------------------------------------------------------------

AzureManagedControlPlane is the Schema for the azuremanagedcontrolplanes API.

Used by: [AzureManagedControlPlaneList](#AzureManagedControlPlaneList).

| Property          | Description | Type                                                              |
|-------------------|-------------|-------------------------------------------------------------------|
| metav1.TypeMeta   |             |                                                                   |
| metav1.ObjectMeta |             |                                                                   |
| spec              |             | [AzureManagedControlPlaneSpec](#AzureManagedControlPlaneSpec)     |
| status            |             | [AzureManagedControlPlaneStatus](#AzureManagedControlPlaneStatus) |

### <a id="AzureManagedControlPlaneSpec"></a>AzureManagedControlPlaneSpec

| Property                    | Description                                                                                                                                                                                                                                                            | Type                                                                    |
|-----------------------------|------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|-------------------------------------------------------------------------|
| aadProfile                  | AadProfile is Azure Active Directory configuration to integrate with AKS for aad authentication.                                                                                                                                                                       | [AADProfile](#AADProfile)                                               |
| additionalTags              | AdditionalTags is an optional set of tags to add to Azure resources managed by the Azure provider, in addition to the ones added by default.                                                                                                                           | Tags                                                                    |
| addonProfiles               | AddonProfiles are the profiles of managed cluster add-on.                                                                                                                                                                                                              | [AddonProfile[]](#AddonProfile)                                         |
| apiServerAccessProfile      | APIServerAccessProfile is the access profile for AKS API server. Immutable except for `authorizedIPRanges`.                                                                                                                                                            | [APIServerAccessProfile](#APIServerAccessProfile)                       |
| autoscalerProfile           | AutoscalerProfile is the parameters to be applied to the cluster-autoscaler when enabled                                                                                                                                                                               | [AutoScalerProfile](#AutoScalerProfile)                                 |
| azureEnvironment            | AzureEnvironment is the name of the AzureCloud to be used. The default value that would be used by most users is "AzurePublicCloud", other values are: - ChinaCloud: "AzureChinaCloud" - PublicCloud: "AzurePublicCloud" - USGovernmentCloud: "AzureUSGovernmentCloud" | string                                                                  |
| controlPlaneEndpoint        | ControlPlaneEndpoint represents the endpoint used to communicate with the control plane. Immutable, populated by the AKS API at create.                                                                                                                                | clusterv1.APIEndpoint                                                   |
| dnsServiceIP                | DNSServiceIP is an IP address assigned to the Kubernetes DNS service. It must be within the Kubernetes service address range specified in serviceCidr. Immutable.                                                                                                      | string                                                                  |
| httpProxyConfig             | HTTPProxyConfig is the HTTP proxy configuration for the cluster. Immutable.                                                                                                                                                                                            | [HTTPProxyConfig](#HTTPProxyConfig)                                     |
| identity                    | Identity configuration used by the AKS control plane.                                                                                                                                                                                                                  | [Identity](#Identity)                                                   |
| identityRef                 | IdentityRef is a reference to a AzureClusterIdentity to be used when reconciling this cluster                                                                                                                                                                          | corev1.ObjectReference                                                  |
| kubeletUserAssignedIdentity | KubeletUserAssignedIdentity is the user-assigned identity for kubelet. For authentication with Azure Container Registry.                                                                                                                                               | string                                                                  |
| loadBalancerProfile         | LoadBalancerProfile is the profile of the cluster load balancer.                                                                                                                                                                                                       | [LoadBalancerProfile](#LoadBalancerProfile)                             |
| loadBalancerSKU             | LoadBalancerSKU is the SKU of the loadBalancer to be provisioned. Immutable.                                                                                                                                                                                           | string                                                                  |
| location                    | Location is a string matching one of the canonical Azure region names. Examples: "westus2", "eastus". Immutable.                                                                                                                                                       | string                                                                  |
| networkPlugin               | NetworkPlugin used for building Kubernetes network. Allowed values are "azure", "kubenet". Immutable.                                                                                                                                                                  | string                                                                  |
| networkPluginMode           | NetworkPluginMode is the mode the network plugin should use. Allowed value is "overlay".                                                                                                                                                                               | [NetworkPluginMode](#NetworkPluginMode)                                 |
| networkPolicy               | NetworkPolicy used for building Kubernetes network. Allowed values are "azure", "calico". Immutable.                                                                                                                                                                   | string                                                                  |
| nodeResourceGroupName       | NodeResourceGroupName is the name of the resource group containing cluster IaaS resources. Will be populated to default in webhook. Immutable.                                                                                                                         | string                                                                  |
| oidcIssuerProfile           | OIDCIssuerProfile is the OIDC issuer profile of the Managed Cluster.                                                                                                                                                                                                   | [OIDCIssuerProfile](#OIDCIssuerProfile)                                 |
| outboundType                | Outbound configuration used by Nodes. Immutable.                                                                                                                                                                                                                       | [ManagedControlPlaneOutboundType](#ManagedControlPlaneOutboundType)     |
| resourceGroupName           | ResourceGroupName is the name of the Azure resource group for this AKS Cluster. Immutable.                                                                                                                                                                             | string                                                                  |
| sku                         | SKU is the SKU of the AKS to be provisioned.                                                                                                                                                                                                                           | [AKSSku](#AKSSku)                                                       |
| sshPublicKey                | SSHPublicKey is a string literal containing an ssh public key base64 encoded. Use empty string to autogenerate new key. Use null value to not set key. Immutable.                                                                                                      | string                                                                  |
| subscriptionID              | SubscriptionID is the GUID of the Azure subscription to hold this cluster. Immutable.                                                                                                                                                                                  | string                                                                  |
| version                     | Version defines the desired Kubernetes version.                                                                                                                                                                                                                        | string                                                                  |
| virtualNetwork              | VirtualNetwork describes the vnet for the AKS cluster. Will be created if it does not exist. Immutable except for `subnet`.                                                                                                                                            | [ManagedControlPlaneVirtualNetwork](#ManagedControlPlaneVirtualNetwork) |

### <a id="AzureManagedControlPlaneStatus"></a>AzureManagedControlPlaneStatus

| Property                   | Description                                                                                                                                                                                               | Type                                                |
|----------------------------|-----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|-----------------------------------------------------|
| conditions                 | Conditions defines current service state of the AzureManagedControlPlane.                                                                                                                                 | clusterv1.Conditions                                |
| initialized                | Initialized is true when the control plane is available for initial contact. This may occur before the control plane is fully ready. In the AzureManagedControlPlane implementation, these are identical. | bool                                                |
| longRunningOperationStates | LongRunningOperationStates saves the states for Azure long-running operations so they can be continued on the next reconciliation loop.                                                                   | Futures                                             |
| oidcIssuerProfile          | OIDCIssuerProfile is the OIDC issuer profile of the Managed Cluster.                                                                                                                                      | [OIDCIssuerProfileStatus](#OIDCIssuerProfileStatus) |
| ready                      | Ready is true when the provider resource is ready.                                                                                                                                                        | bool                                                |

<a id="AzureManagedControlPlaneList"></a>AzureManagedControlPlaneList
---------------------------------------------------------------------

AzureManagedControlPlaneList contains a list of AzureManagedControlPlane.

| Property        | Description | Type                                                    |
|-----------------|-------------|---------------------------------------------------------|
| metav1.TypeMeta |             |                                                         |
| metav1.ListMeta |             |                                                         |
| items           |             | [AzureManagedControlPlane[]](#AzureManagedControlPlane) |

<a id="AzureManagedControlPlaneSkuTier"></a>AzureManagedControlPlaneSkuTier
---------------------------------------------------------------------------

Used by: [AKSSku](#AKSSku).

<a id="AzureManagedControlPlaneSpec"></a>AzureManagedControlPlaneSpec
---------------------------------------------------------------------

AzureManagedControlPlaneSpec defines the desired state of AzureManagedControlPlane.

Used by: [AzureManagedControlPlane](#AzureManagedControlPlane).

| Property                    | Description                                                                                                                                                                                                                                                            | Type                                                                    |
|-----------------------------|------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|-------------------------------------------------------------------------|
| aadProfile                  | AadProfile is Azure Active Directory configuration to integrate with AKS for aad authentication.                                                                                                                                                                       | [AADProfile](#AADProfile)                                               |
| additionalTags              | AdditionalTags is an optional set of tags to add to Azure resources managed by the Azure provider, in addition to the ones added by default.                                                                                                                           | Tags                                                                    |
| addonProfiles               | AddonProfiles are the profiles of managed cluster add-on.                                                                                                                                                                                                              | [AddonProfile[]](#AddonProfile)                                         |
| apiServerAccessProfile      | APIServerAccessProfile is the access profile for AKS API server. Immutable except for `authorizedIPRanges`.                                                                                                                                                            | [APIServerAccessProfile](#APIServerAccessProfile)                       |
| autoscalerProfile           | AutoscalerProfile is the parameters to be applied to the cluster-autoscaler when enabled                                                                                                                                                                               | [AutoScalerProfile](#AutoScalerProfile)                                 |
| azureEnvironment            | AzureEnvironment is the name of the AzureCloud to be used. The default value that would be used by most users is "AzurePublicCloud", other values are: - ChinaCloud: "AzureChinaCloud" - PublicCloud: "AzurePublicCloud" - USGovernmentCloud: "AzureUSGovernmentCloud" | string                                                                  |
| controlPlaneEndpoint        | ControlPlaneEndpoint represents the endpoint used to communicate with the control plane. Immutable, populated by the AKS API at create.                                                                                                                                | clusterv1.APIEndpoint                                                   |
| dnsServiceIP                | DNSServiceIP is an IP address assigned to the Kubernetes DNS service. It must be within the Kubernetes service address range specified in serviceCidr. Immutable.                                                                                                      | string                                                                  |
| httpProxyConfig             | HTTPProxyConfig is the HTTP proxy configuration for the cluster. Immutable.                                                                                                                                                                                            | [HTTPProxyConfig](#HTTPProxyConfig)                                     |
| identity                    | Identity configuration used by the AKS control plane.                                                                                                                                                                                                                  | [Identity](#Identity)                                                   |
| identityRef                 | IdentityRef is a reference to a AzureClusterIdentity to be used when reconciling this cluster                                                                                                                                                                          | corev1.ObjectReference                                                  |
| kubeletUserAssignedIdentity | KubeletUserAssignedIdentity is the user-assigned identity for kubelet. For authentication with Azure Container Registry.                                                                                                                                               | string                                                                  |
| loadBalancerProfile         | LoadBalancerProfile is the profile of the cluster load balancer.                                                                                                                                                                                                       | [LoadBalancerProfile](#LoadBalancerProfile)                             |
| loadBalancerSKU             | LoadBalancerSKU is the SKU of the loadBalancer to be provisioned. Immutable.                                                                                                                                                                                           | string                                                                  |
| location                    | Location is a string matching one of the canonical Azure region names. Examples: "westus2", "eastus". Immutable.                                                                                                                                                       | string                                                                  |
| networkPlugin               | NetworkPlugin used for building Kubernetes network. Allowed values are "azure", "kubenet". Immutable.                                                                                                                                                                  | string                                                                  |
| networkPluginMode           | NetworkPluginMode is the mode the network plugin should use. Allowed value is "overlay".                                                                                                                                                                               | [NetworkPluginMode](#NetworkPluginMode)                                 |
| networkPolicy               | NetworkPolicy used for building Kubernetes network. Allowed values are "azure", "calico". Immutable.                                                                                                                                                                   | string                                                                  |
| nodeResourceGroupName       | NodeResourceGroupName is the name of the resource group containing cluster IaaS resources. Will be populated to default in webhook. Immutable.                                                                                                                         | string                                                                  |
| oidcIssuerProfile           | OIDCIssuerProfile is the OIDC issuer profile of the Managed Cluster.                                                                                                                                                                                                   | [OIDCIssuerProfile](#OIDCIssuerProfile)                                 |
| outboundType                | Outbound configuration used by Nodes. Immutable.                                                                                                                                                                                                                       | [ManagedControlPlaneOutboundType](#ManagedControlPlaneOutboundType)     |
| resourceGroupName           | ResourceGroupName is the name of the Azure resource group for this AKS Cluster. Immutable.                                                                                                                                                                             | string                                                                  |
| sku                         | SKU is the SKU of the AKS to be provisioned.                                                                                                                                                                                                                           | [AKSSku](#AKSSku)                                                       |
| sshPublicKey                | SSHPublicKey is a string literal containing an ssh public key base64 encoded. Use empty string to autogenerate new key. Use null value to not set key. Immutable.                                                                                                      | string                                                                  |
| subscriptionID              | SubscriptionID is the GUID of the Azure subscription to hold this cluster. Immutable.                                                                                                                                                                                  | string                                                                  |
| version                     | Version defines the desired Kubernetes version.                                                                                                                                                                                                                        | string                                                                  |
| virtualNetwork              | VirtualNetwork describes the vnet for the AKS cluster. Will be created if it does not exist. Immutable except for `subnet`.                                                                                                                                            | [ManagedControlPlaneVirtualNetwork](#ManagedControlPlaneVirtualNetwork) |

<a id="AzureManagedControlPlaneStatus"></a>AzureManagedControlPlaneStatus
-------------------------------------------------------------------------

AzureManagedControlPlaneStatus defines the observed state of AzureManagedControlPlane.

Used by: [AzureManagedControlPlane](#AzureManagedControlPlane).

| Property                   | Description                                                                                                                                                                                               | Type                                                |
|----------------------------|-----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|-----------------------------------------------------|
| conditions                 | Conditions defines current service state of the AzureManagedControlPlane.                                                                                                                                 | clusterv1.Conditions                                |
| initialized                | Initialized is true when the control plane is available for initial contact. This may occur before the control plane is fully ready. In the AzureManagedControlPlane implementation, these are identical. | bool                                                |
| longRunningOperationStates | LongRunningOperationStates saves the states for Azure long-running operations so they can be continued on the next reconciliation loop.                                                                   | Futures                                             |
| oidcIssuerProfile          | OIDCIssuerProfile is the OIDC issuer profile of the Managed Cluster.                                                                                                                                      | [OIDCIssuerProfileStatus](#OIDCIssuerProfileStatus) |
| ready                      | Ready is true when the provider resource is ready.                                                                                                                                                        | bool                                                |

<a id="AzureManagedMachinePool"></a>AzureManagedMachinePool
-----------------------------------------------------------

AzureManagedMachinePool is the Schema for the azuremanagedmachinepools API.

Used by: [AzureManagedMachinePoolList](#AzureManagedMachinePoolList).

| Property          | Description | Type                                                            |
|-------------------|-------------|-----------------------------------------------------------------|
| metav1.TypeMeta   |             |                                                                 |
| metav1.ObjectMeta |             |                                                                 |
| spec              |             | [AzureManagedMachinePoolSpec](#AzureManagedMachinePoolSpec)     |
| status            |             | [AzureManagedMachinePoolStatus](#AzureManagedMachinePoolStatus) |

### <a id="AzureManagedMachinePoolSpec"></a>AzureManagedMachinePoolSpec

| Property             | Description                                                                                                                                                                                                                                                                                                                                                     | Type                                                    |
|----------------------|-----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|---------------------------------------------------------|
| additionalTags       | AdditionalTags is an optional set of tags to add to Azure resources managed by the Azure provider, in addition to the ones added by default.                                                                                                                                                                                                                    | Tags                                                    |
| availabilityZones    | AvailabilityZones - Availability zones for nodes. Must use VirtualMachineScaleSets AgentPoolType. Immutable.                                                                                                                                                                                                                                                    | string[]                                                |
| enableFIPS           | EnableFIPS indicates whether FIPS is enabled on the node pool. Immutable.                                                                                                                                                                                                                                                                                       | bool                                                    |
| enableNodePublicIP   | EnableNodePublicIP controls whether or not nodes in the pool each have a public IP address. Immutable.                                                                                                                                                                                                                                                          | bool                                                    |
| enableUltraSSD       | EnableUltraSSD enables the storage type UltraSSD_LRS for the agent pool. Immutable.                                                                                                                                                                                                                                                                             | bool                                                    |
| kubeletConfig        | KubeletConfig specifies the kubelet configurations for nodes. Immutable.                                                                                                                                                                                                                                                                                        | [KubeletConfig](#KubeletConfig)                         |
| kubeletDiskType      | KubeletDiskType specifies the kubelet disk type. Default to OS. Possible values include: 'OS', 'Temporary'. Requires Microsoft.ContainerService/KubeletDisk preview feature to be set. Immutable. See also [AKS doc](https://learn.microsoft.com/rest/api/aks/agent-pools/create-or-update?tabs=HTTP#kubeletdisktype). <br/>                                    | [KubeletDiskType](#KubeletDiskType)                     |
| linuxOSConfig        | LinuxOSConfig specifies the custom Linux OS settings and configurations. Immutable.                                                                                                                                                                                                                                                                             | [LinuxOSConfig](#LinuxOSConfig)                         |
| maxPods              | MaxPods specifies the kubelet `--max-pods` configuration for the node pool. Immutable. See also [AKS doc](https://learn.microsoft.com/azure/aks/configure-azure-cni#configure-maximum---new-clusters), [K8s doc](https://kubernetes.io/docs/reference/command-line-tools-reference/kubelet/). <br/>                                                             | int32                                                   |
| mode                 | Mode - represents mode of an agent pool. Possible values include: System, User.                                                                                                                                                                                                                                                                                 | string                                                  |
| name                 | Name - name of the agent pool. If not specified, CAPZ uses the name of the CR as the agent pool name. Immutable.                                                                                                                                                                                                                                                | string                                                  |
| nodeLabels           | Node labels - labels for all of the nodes present in node pool. See also [AKS doc](https://learn.microsoft.com/azure/aks/use-labels). <br/>                                                                                                                                                                                                                     | map[string]string                                       |
| nodePublicIPPrefixID | NodePublicIPPrefixID specifies the public IP prefix resource ID which VM nodes should use IPs from. Immutable.                                                                                                                                                                                                                                                  | string                                                  |
| osDiskSizeGB         | OSDiskSizeGB is the disk size for every machine in this agent pool. If you specify 0, it will apply the default osDisk size according to the vmSize specified. Immutable.                                                                                                                                                                                       | int32                                                   |
| osDiskType           | OsDiskType specifies the OS disk type for each node in the pool. Allowed values are 'Ephemeral' and 'Managed' (default). Immutable. See also [AKS doc](https://learn.microsoft.com/azure/aks/cluster-configuration#ephemeral-os). <br/>                                                                                                                         | string                                                  |
| osType               | OSType specifies the virtual machine operating system. Default to Linux. Possible values include: 'Linux', 'Windows'. 'Windows' requires the AzureManagedControlPlane's `spec.networkPlugin` to be `azure`. Immutable. See also [AKS doc](https://learn.microsoft.com/rest/api/aks/agent-pools/create-or-update?tabs=HTTP#ostype). <br/>                        | string                                                  |
| providerIDList       | ProviderIDList is the unique identifier as specified by the cloud provider.                                                                                                                                                                                                                                                                                     | string[]                                                |
| scaleDownMode        | ScaleDownMode affects the cluster autoscaler behavior. Default to Delete. Possible values include: 'Deallocate', 'Delete'                                                                                                                                                                                                                                       | string                                                  |
| scaleSetPriority     | ScaleSetPriority specifies the ScaleSetPriority value. Default to Regular. Possible values include: 'Regular', 'Spot' Immutable.                                                                                                                                                                                                                                | string                                                  |
| scaling              | Scaling specifies the autoscaling parameters for the node pool.                                                                                                                                                                                                                                                                                                 | [ManagedMachinePoolScaling](#ManagedMachinePoolScaling) |
| sku                  | SKU is the size of the VMs in the node pool. Immutable.                                                                                                                                                                                                                                                                                                         | string                                                  |
| spotMaxPrice         | SpotMaxPrice defines max price to pay for spot instance. Possible values are any decimal value greater than zero or -1. If you set the max price to be -1, the VM won't be evicted based on price. The price for the VM will be the current price for spot or the price for a standard VM, which ever is less, as long as there's capacity and quota available. | resource.Quantity                                       |
| subnetName           | SubnetName specifies the Subnet where the MachinePool will be placed Immutable.                                                                                                                                                                                                                                                                                 | string                                                  |
| taints               | Taints specifies the taints for nodes present in this agent pool. See also [AKS doc](https://learn.microsoft.com/azure/aks/use-multiple-node-pools#setting-node-pool-taints). <br/>                                                                                                                                                                             | Taints                                                  |

### <a id="AzureManagedMachinePoolStatus"></a>AzureManagedMachinePoolStatus

| Property                   | Description                                                                                                                                                  | Type                          |
|----------------------------|--------------------------------------------------------------------------------------------------------------------------------------------------------------|-------------------------------|
| conditions                 | Conditions defines current service state of the AzureManagedControlPlane.                                                                                    | clusterv1.Conditions          |
| errorMessage               | Any transient errors that occur during the reconciliation of Machines can be added as events to the Machine object and/or logged in the controller's output. | string                        |
| errorReason                | Any transient errors that occur during the reconciliation of Machines can be added as events to the Machine object and/or logged in the controller's output. | capierrors.MachineStatusError |
| longRunningOperationStates | LongRunningOperationStates saves the states for Azure long-running operations so they can be continued on the next reconciliation loop.                      | Futures                       |
| ready                      | Ready is true when the provider resource is ready.                                                                                                           | bool                          |
| replicas                   | Replicas is the most recently observed number of replicas.                                                                                                   | int32                         |

<a id="AzureManagedMachinePoolList"></a>AzureManagedMachinePoolList
-------------------------------------------------------------------

AzureManagedMachinePoolList contains a list of AzureManagedMachinePools.

| Property        | Description | Type                                                  |
|-----------------|-------------|-------------------------------------------------------|
| metav1.TypeMeta |             |                                                       |
| metav1.ListMeta |             |                                                       |
| items           |             | [AzureManagedMachinePool[]](#AzureManagedMachinePool) |

<a id="AzureManagedMachinePoolSpec"></a>AzureManagedMachinePoolSpec
-------------------------------------------------------------------

AzureManagedMachinePoolSpec defines the desired state of AzureManagedMachinePool.

Used by: [AzureManagedMachinePool](#AzureManagedMachinePool).

| Property             | Description                                                                                                                                                                                                                                                                                                                                                     | Type                                                    |
|----------------------|-----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|---------------------------------------------------------|
| additionalTags       | AdditionalTags is an optional set of tags to add to Azure resources managed by the Azure provider, in addition to the ones added by default.                                                                                                                                                                                                                    | Tags                                                    |
| availabilityZones    | AvailabilityZones - Availability zones for nodes. Must use VirtualMachineScaleSets AgentPoolType. Immutable.                                                                                                                                                                                                                                                    | string[]                                                |
| enableFIPS           | EnableFIPS indicates whether FIPS is enabled on the node pool. Immutable.                                                                                                                                                                                                                                                                                       | bool                                                    |
| enableNodePublicIP   | EnableNodePublicIP controls whether or not nodes in the pool each have a public IP address. Immutable.                                                                                                                                                                                                                                                          | bool                                                    |
| enableUltraSSD       | EnableUltraSSD enables the storage type UltraSSD_LRS for the agent pool. Immutable.                                                                                                                                                                                                                                                                             | bool                                                    |
| kubeletConfig        | KubeletConfig specifies the kubelet configurations for nodes. Immutable.                                                                                                                                                                                                                                                                                        | [KubeletConfig](#KubeletConfig)                         |
| kubeletDiskType      | KubeletDiskType specifies the kubelet disk type. Default to OS. Possible values include: 'OS', 'Temporary'. Requires Microsoft.ContainerService/KubeletDisk preview feature to be set. Immutable. See also [AKS doc](https://learn.microsoft.com/rest/api/aks/agent-pools/create-or-update?tabs=HTTP#kubeletdisktype). <br/>                                    | [KubeletDiskType](#KubeletDiskType)                     |
| linuxOSConfig        | LinuxOSConfig specifies the custom Linux OS settings and configurations. Immutable.                                                                                                                                                                                                                                                                             | [LinuxOSConfig](#LinuxOSConfig)                         |
| maxPods              | MaxPods specifies the kubelet `--max-pods` configuration for the node pool. Immutable. See also [AKS doc](https://learn.microsoft.com/azure/aks/configure-azure-cni#configure-maximum---new-clusters), [K8s doc](https://kubernetes.io/docs/reference/command-line-tools-reference/kubelet/). <br/>                                                             | int32                                                   |
| mode                 | Mode - represents mode of an agent pool. Possible values include: System, User.                                                                                                                                                                                                                                                                                 | string                                                  |
| name                 | Name - name of the agent pool. If not specified, CAPZ uses the name of the CR as the agent pool name. Immutable.                                                                                                                                                                                                                                                | string                                                  |
| nodeLabels           | Node labels - labels for all of the nodes present in node pool. See also [AKS doc](https://learn.microsoft.com/azure/aks/use-labels). <br/>                                                                                                                                                                                                                     | map[string]string                                       |
| nodePublicIPPrefixID | NodePublicIPPrefixID specifies the public IP prefix resource ID which VM nodes should use IPs from. Immutable.                                                                                                                                                                                                                                                  | string                                                  |
| osDiskSizeGB         | OSDiskSizeGB is the disk size for every machine in this agent pool. If you specify 0, it will apply the default osDisk size according to the vmSize specified. Immutable.                                                                                                                                                                                       | int32                                                   |
| osDiskType           | OsDiskType specifies the OS disk type for each node in the pool. Allowed values are 'Ephemeral' and 'Managed' (default). Immutable. See also [AKS doc](https://learn.microsoft.com/azure/aks/cluster-configuration#ephemeral-os). <br/>                                                                                                                         | string                                                  |
| osType               | OSType specifies the virtual machine operating system. Default to Linux. Possible values include: 'Linux', 'Windows'. 'Windows' requires the AzureManagedControlPlane's `spec.networkPlugin` to be `azure`. Immutable. See also [AKS doc](https://learn.microsoft.com/rest/api/aks/agent-pools/create-or-update?tabs=HTTP#ostype). <br/>                        | string                                                  |
| providerIDList       | ProviderIDList is the unique identifier as specified by the cloud provider.                                                                                                                                                                                                                                                                                     | string[]                                                |
| scaleDownMode        | ScaleDownMode affects the cluster autoscaler behavior. Default to Delete. Possible values include: 'Deallocate', 'Delete'                                                                                                                                                                                                                                       | string                                                  |
| scaleSetPriority     | ScaleSetPriority specifies the ScaleSetPriority value. Default to Regular. Possible values include: 'Regular', 'Spot' Immutable.                                                                                                                                                                                                                                | string                                                  |
| scaling              | Scaling specifies the autoscaling parameters for the node pool.                                                                                                                                                                                                                                                                                                 | [ManagedMachinePoolScaling](#ManagedMachinePoolScaling) |
| sku                  | SKU is the size of the VMs in the node pool. Immutable.                                                                                                                                                                                                                                                                                                         | string                                                  |
| spotMaxPrice         | SpotMaxPrice defines max price to pay for spot instance. Possible values are any decimal value greater than zero or -1. If you set the max price to be -1, the VM won't be evicted based on price. The price for the VM will be the current price for spot or the price for a standard VM, which ever is less, as long as there's capacity and quota available. | resource.Quantity                                       |
| subnetName           | SubnetName specifies the Subnet where the MachinePool will be placed Immutable.                                                                                                                                                                                                                                                                                 | string                                                  |
| taints               | Taints specifies the taints for nodes present in this agent pool. See also [AKS doc](https://learn.microsoft.com/azure/aks/use-multiple-node-pools#setting-node-pool-taints). <br/>                                                                                                                                                                             | Taints                                                  |

<a id="AzureManagedMachinePoolStatus"></a>AzureManagedMachinePoolStatus
-----------------------------------------------------------------------

AzureManagedMachinePoolStatus defines the observed state of AzureManagedMachinePool.

Used by: [AzureManagedMachinePool](#AzureManagedMachinePool).

| Property                   | Description                                                                                                                                                  | Type                          |
|----------------------------|--------------------------------------------------------------------------------------------------------------------------------------------------------------|-------------------------------|
| conditions                 | Conditions defines current service state of the AzureManagedControlPlane.                                                                                    | clusterv1.Conditions          |
| errorMessage               | Any transient errors that occur during the reconciliation of Machines can be added as events to the Machine object and/or logged in the controller's output. | string                        |
| errorReason                | Any transient errors that occur during the reconciliation of Machines can be added as events to the Machine object and/or logged in the controller's output. | capierrors.MachineStatusError |
| longRunningOperationStates | LongRunningOperationStates saves the states for Azure long-running operations so they can be continued on the next reconciliation loop.                      | Futures                       |
| ready                      | Ready is true when the provider resource is ready.                                                                                                           | bool                          |
| replicas                   | Replicas is the most recently observed number of replicas.                                                                                                   | int32                         |

<a id="AzureMarketplaceImage"></a>AzureMarketplaceImage
-------------------------------------------------------

AzureMarketplaceImage defines an image in the Azure Marketplace to use for VM creation.

Used by: [Image](#Image).

| Property                | Description                                                                                                                                                                                                                                                                                                                                                         | Type   |
|-------------------------|---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|--------|
| [ImagePlan](#ImagePlan) |                                                                                                                                                                                                                                                                                                                                                                     |        |
| thirdPartyImage         | ThirdPartyImage indicates the image is published by a third party publisher and a Plan will be generated for it.                                                                                                                                                                                                                                                    | bool   |
| version                 | Version specifies the version of an image sku. The allowed formats are Major.Minor.Build or 'latest'. Major, Minor, and Build are decimal numbers. Specify 'latest' to use the latest version of an image available at deploy time. Even if you use 'latest', the VM image will not automatically update after deploy time even if a new version becomes available. | string |

<a id="AzureSharedGalleryImage"></a>AzureSharedGalleryImage
-----------------------------------------------------------

AzureSharedGalleryImage defines an image in a Shared Image Gallery to use for VM creation.

Used by: [Image](#Image).

| Property       | Description                                                                                                                                                                                                                                                                                                                                                                  | Type   |
|----------------|------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|--------|
| gallery        | Gallery specifies the name of the shared image gallery that contains the image                                                                                                                                                                                                                                                                                               | string |
| name           | Name is the name of the image                                                                                                                                                                                                                                                                                                                                                | string |
| offer          | Offer specifies the name of a group of related images created by the publisher. For example, UbuntuServer, WindowsServer This value will be used to add a `Plan` in the API request when creating the VM/VMSS resource. This is needed when the source image from which this SIG image was built requires the `Plan` to be used.                                             | string |
| publisher      | Publisher is the name of the organization that created the image. This value will be used to add a `Plan` in the API request when creating the VM/VMSS resource. This is needed when the source image from which this SIG image was built requires the `Plan` to be used.                                                                                                    | string |
| resourceGroup  | ResourceGroup specifies the resource group containing the shared image gallery                                                                                                                                                                                                                                                                                               | string |
| sku            | SKU specifies an instance of an offer, such as a major release of a distribution. For example, 18.04-LTS, 2019-Datacenter This value will be used to add a `Plan` in the API request when creating the VM/VMSS resource. This is needed when the source image from which this SIG image was built requires the `Plan` to be used.                                            | string |
| subscriptionID | SubscriptionID is the identifier of the subscription that contains the shared image gallery                                                                                                                                                                                                                                                                                  | string |
| version        | Version specifies the version of the marketplace image. The allowed formats are Major.Minor.Build or 'latest'. Major, Minor, and Build are decimal numbers. Specify 'latest' to use the latest version of an image available at deploy time. Even if you use 'latest', the VM image will not automatically update after deploy time even if a new version becomes available. | string |

<a id="BackOffConfig"></a>BackOffConfig
---------------------------------------

BackOffConfig indicates the back-off config options.

Used by: [CloudProviderConfigOverrides](#CloudProviderConfigOverrides).

| Property                     | Description | Type              |
|------------------------------|-------------|-------------------|
| cloudProviderBackoff         |             | bool              |
| cloudProviderBackoffDuration |             | int               |
| cloudProviderBackoffExponent |             | resource.Quantity |
| cloudProviderBackoffJitter   |             | resource.Quantity |
| cloudProviderBackoffRetries  |             | int               |

<a id="BackendPool"></a>BackendPool
-----------------------------------

BackendPool describes the backend pool of the load balancer.

Used by: [LoadBalancerSpec](#LoadBalancerSpec).

| Property | Description                                                                                                                                         | Type   |
|----------|-----------------------------------------------------------------------------------------------------------------------------------------------------|--------|
| name     | Name specifies the name of backend pool for the load balancer. If not specified, the default name will be set, depending on the load balancer role. | string |

<a id="BalanceSimilarNodeGroups"></a>BalanceSimilarNodeGroups
-------------------------------------------------------------

Used by: [AutoScalerProfile](#AutoScalerProfile).

<a id="BastionHostSkuName"></a>BastionHostSkuName
-------------------------------------------------

Used by: [AzureBastion](#AzureBastion).

<a id="BastionSpec"></a>BastionSpec
-----------------------------------

BastionSpec specifies how the Bastion feature should be set up for the cluster.

Used by: [AzureClusterSpec](#AzureClusterSpec).

| Property     | Description | Type                          |
|--------------|-------------|-------------------------------|
| azureBastion |             | [AzureBastion](#AzureBastion) |

<a id="BastionTemplateSpec"></a>BastionTemplateSpec
---------------------------------------------------

BastionTemplateSpec specifies a template for a bastion host.

Used by: [AzureClusterTemplateResourceSpec](#AzureClusterTemplateResourceSpec).

| Property     | Description | Type                                                  |
|--------------|-------------|-------------------------------------------------------|
| azureBastion |             | [AzureBastionTemplateSpec](#AzureBastionTemplateSpec) |

<a id="BootDiagnostics"></a>BootDiagnostics
-------------------------------------------

BootDiagnostics configures the boot diagnostics settings for the virtual machine. This allows you to configure capturing serial output from the virtual machine on boot. This is useful for debugging software based launch issues.

Used by: [Diagnostics](#Diagnostics).

| Property           | Description                                                                                                                                                                       | Type                                                                                 |
|--------------------|-----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|--------------------------------------------------------------------------------------|
| storageAccountType | StorageAccountType determines if the storage account for storing the diagnostics data should be disabled (Disabled), provisioned by Azure (Managed) or by the user (UserManaged). | [BootDiagnosticsStorageAccountType](#BootDiagnosticsStorageAccountType)<br/>Required |
| userManaged        | UserManaged provides a reference to the user-managed storage account.                                                                                                             | [UserManagedBootDiagnostics](#UserManagedBootDiagnostics)                            |

<a id="BootDiagnosticsStorageAccountType"></a>BootDiagnosticsStorageAccountType
-------------------------------------------------------------------------------

Used by: [BootDiagnostics](#BootDiagnostics).

<a id="BuildParams"></a>BuildParams
-----------------------------------

BuildParams is used to build tags around an azure resource.

| Property    | Description                                                                | Type                                    |
|-------------|----------------------------------------------------------------------------|-----------------------------------------|
| Additional  | Any additional tags to be added to the resource.                           | Tags                                    |
| ClusterName | ClusterName is the cluster associated with the resource.                   | string                                  |
| Lifecycle   | Lifecycle determines the resource lifecycle.                               | [ResourceLifecycle](#ResourceLifecycle) |
| Name        | Name is the name of the resource, it's applied as the tag "Name" on Azure. | string                                  |
| ResourceID  | ResourceID is the unique identifier of the resource to be tagged.          | string                                  |
| Role        | Role is the role associated to the resource.                               | string                                  |

<a id="CPUManagerPolicy"></a>CPUManagerPolicy
---------------------------------------------

Used by: [KubeletConfig](#KubeletConfig).

<a id="CloudProviderConfigOverrides"></a>CloudProviderConfigOverrides
---------------------------------------------------------------------

CloudProviderConfigOverrides represents the fields that can be overridden in azure cloud provider config.

Used by: [AzureClusterClassSpec](#AzureClusterClassSpec).

| Property   | Description | Type                              |
|------------|-------------|-----------------------------------|
| backOffs   |             | [BackOffConfig](#BackOffConfig)   |
| rateLimits |             | [RateLimitSpec[]](#RateLimitSpec) |

<a id="DataDisk"></a>DataDisk
-----------------------------

DataDisk specifies the parameters that are used to add one or more data disks to the machine.

Used by: [AzureMachineSpec](#AzureMachineSpec).

| Property    | Description                                                                                                                                                                                                           | Type                                            |
|-------------|-----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|-------------------------------------------------|
| cachingType | CachingType specifies the caching requirements.                                                                                                                                                                       | string                                          |
| diskSizeGB  | DiskSizeGB is the size in GB to assign to the data disk.                                                                                                                                                              | int32                                           |
| lun         | Lun Specifies the logical unit number of the data disk. This value is used to identify data disks within the VM and therefore must be unique for each data disk attached to a VM. The value must be between 0 and 63. | int32                                           |
| managedDisk | ManagedDisk specifies the Managed Disk parameters for the data disk.                                                                                                                                                  | [ManagedDiskParameters](#ManagedDiskParameters) |
| nameSuffix  | NameSuffix is the suffix to be appended to the machine name to generate the disk name. Each disk name will be in format <machineName>\_<nameSuffix>.                                                                  | string                                          |

<a id="Diagnostics"></a>Diagnostics
-----------------------------------

Diagnostics is used to configure the diagnostic settings of the virtual machine.

Used by: [AzureMachineSpec](#AzureMachineSpec).

| Property | Description                                                                                                                                                                                                                                                                            | Type                                |
|----------|----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|-------------------------------------|
| boot     | Boot configures the boot diagnostics settings for the virtual machine. This allows to configure capturing serial output from the virtual machine on boot. This is useful for debugging software based launch issues. If not specified then Boot diagnostics (Managed) will be enabled. | [BootDiagnostics](#BootDiagnostics) |

<a id="DiffDiskSettings"></a>DiffDiskSettings
---------------------------------------------

DiffDiskSettings describe ephemeral disk settings for the os disk.

Used by: [OSDisk](#OSDisk).

| Property | Description                                                                                                                                | Type   |
|----------|--------------------------------------------------------------------------------------------------------------------------------------------|--------|
| option   | Option enables ephemeral OS when set to "Local" See https://learn.microsoft.com/azure/virtual-machines/ephemeral-os-disks for full details | string |

<a id="DiskEncryptionSetParameters"></a>DiskEncryptionSetParameters
-------------------------------------------------------------------

DiskEncryptionSetParameters defines disk encryption options.

Used by: [ManagedDiskParameters](#ManagedDiskParameters), and [VMDiskSecurityProfile](#VMDiskSecurityProfile).

| Property | Description                                                                               | Type   |
|----------|-------------------------------------------------------------------------------------------|--------|
| id       | ID defines resourceID for diskEncryptionSet resource. It must be in the same subscription | string |

<a id="Expander"></a>Expander
-----------------------------

Used by: [AutoScalerProfile](#AutoScalerProfile).

<a id="ExtendedLocationSpec"></a>ExtendedLocationSpec
-----------------------------------------------------

ExtendedLocationSpec defines the ExtendedLocation properties to enable CAPZ for Azure public MEC.

Used by: [AzureClusterClassSpec](#AzureClusterClassSpec).

| Property | Description                                      | Type   |
|----------|--------------------------------------------------|--------|
| name     | Name defines the name for the extended location. | string |
| type     | Type defines the type for the extended location. | string |

<a id="FrontendIP"></a>FrontendIP
---------------------------------

FrontendIP defines a load balancer frontend IP configuration.

Used by: [LoadBalancerSpec](#LoadBalancerSpec).

| Property                            | Description | Type                          |
|-------------------------------------|-------------|-------------------------------|
| [FrontendIPClass](#FrontendIPClass) |             |                               |
| name                                |             | string                        |
| publicIP                            |             | [PublicIPSpec](#PublicIPSpec) |

<a id="FrontendIPClass"></a>FrontendIPClass
-------------------------------------------

FrontendIPClass defines the FrontendIP properties that may be shared across several Azure clusters.

| Property  | Description | Type   |
|-----------|-------------|--------|
| privateIP |             | string |

<a id="Future"></a>Future
-------------------------

Future contains the data needed for an Azure long-running operation to continue across reconcile loops.

| Property      | Description                                                                                                                            | Type   |
|---------------|----------------------------------------------------------------------------------------------------------------------------------------|--------|
| data          | Data is the base64 url encoded json Azure AutoRest Future.                                                                             | string |
| name          | Name is the name of the Azure resource. Together with the service name, this forms the unique identifier for the future.               | string |
| resourceGroup | ResourceGroup is the Azure resource group for the resource.                                                                            | string |
| serviceName   | ServiceName is the name of the Azure service. Together with the name of the resource, this forms the unique identifier for the future. | string |
| type          | Type describes the type of future, such as update, create, delete, etc.                                                                | string |

<a id="HTTPProxyConfig"></a>HTTPProxyConfig
-------------------------------------------

HTTPProxyConfig is the HTTP proxy configuration for the cluster.

Used by: [AzureManagedControlPlaneSpec](#AzureManagedControlPlaneSpec).

| Property   | Description                                                                  | Type     |
|------------|------------------------------------------------------------------------------|----------|
| httpProxy  | HTTPProxy is the HTTP proxy server endpoint to use.                          | string   |
| httpsProxy | HTTPSProxy is the HTTPS proxy server endpoint to use.                        | string   |
| noProxy    | NoProxy indicates the endpoints that should not go through proxy.            | string[] |
| trustedCa  | TrustedCA is the alternative CA cert to use for connecting to proxy servers. | string   |

<a id="IPTag"></a>IPTag
-----------------------

IPTag contains the IpTag associated with the object.

Used by: [PublicIPSpec](#PublicIPSpec).

| Property | Description                                                                        | Type   |
|----------|------------------------------------------------------------------------------------|--------|
| tag      | Tag specifies the value of the IP tag associated with the public IP. Example: SQL. | string |
| type     | Type specifies the IP tag type. Example: FirstPartyUsage.                          | string |

<a id="Identity"></a>Identity
-----------------------------

Identity represents the Identity configuration for an AKS control plane. See also [AKS doc](https://learn.microsoft.com/en-us/azure/aks/use-managed-identity). <br/>

Used by: [AzureManagedControlPlaneSpec](#AzureManagedControlPlaneSpec).

| Property                       | Description                                                                                  | Type                                                                |
|--------------------------------|----------------------------------------------------------------------------------------------|---------------------------------------------------------------------|
| type                           | Type - The Identity type to use.                                                             | [ManagedControlPlaneIdentityType](#ManagedControlPlaneIdentityType) |
| userAssignedIdentityResourceID | UserAssignedIdentityResourceID - Identity ARM resource ID when using user-assigned identity. | string                                                              |

<a id="IdentityType"></a>IdentityType
-------------------------------------

Used by: [AzureClusterIdentitySpec](#AzureClusterIdentitySpec).

<a id="Image"></a>Image
-----------------------

Image defines information about the image to use for VM creation. There are three ways to specify an image: by ID, Marketplace Image or SharedImageGallery One of ID, SharedImage or Marketplace should be set.

Used by: [AzureMachineSpec](#AzureMachineSpec).

| Property       | Description                                                                                                        | Type                                                  |
|----------------|--------------------------------------------------------------------------------------------------------------------|-------------------------------------------------------|
| computeGallery | ComputeGallery specifies an image to use from the Azure Compute Gallery                                            | [AzureComputeGalleryImage](#AzureComputeGalleryImage) |
| id             | ID specifies an image to use by ID                                                                                 | string                                                |
| marketplace    | Marketplace specifies an image to use from the Azure Marketplace                                                   | [AzureMarketplaceImage](#AzureMarketplaceImage)       |
| sharedGallery  | SharedGallery specifies an image to use from an Azure Shared Image Gallery Deprecated: use ComputeGallery instead. | [AzureSharedGalleryImage](#AzureSharedGalleryImage)   |

<a id="ImagePlan"></a>ImagePlan
-------------------------------

ImagePlan contains plan information for marketplace images.

Used by: [AzureComputeGalleryImage](#AzureComputeGalleryImage).

| Property  | Description                                                                                                               | Type   |
|-----------|---------------------------------------------------------------------------------------------------------------------------|--------|
| offer     | Offer specifies the name of a group of related images created by the publisher. For example, UbuntuServer, WindowsServer  | string |
| publisher | Publisher is the name of the organization that created the image                                                          | string |
| sku       | SKU specifies an instance of an offer, such as a major release of a distribution. For example, 18.04-LTS, 2019-Datacenter | string |

<a id="KubeletConfig"></a>KubeletConfig
---------------------------------------

KubeletConfig defines the supported subset of kubelet configurations for nodes in pools. See also [AKS doc](https://learn.microsoft.com/azure/aks/custom-node-configuration), [K8s doc](https://kubernetes.io/docs/reference/config-api/kubelet-config.v1beta1/). <br/>

Used by: [AzureManagedMachinePoolSpec](#AzureManagedMachinePoolSpec).

| Property              | Description                                                                                                                                                                             | Type                                            |
|-----------------------|-----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|-------------------------------------------------|
| allowedUnsafeSysctls  | AllowedUnsafeSysctls - Allowlist of unsafe sysctls or unsafe sysctl patterns (ending in `*`). Valid values match `kernel.shm*`, `kernel.msg*`, `kernel.sem`, `fs.mqueue.*`, or `net.*`. | string[]                                        |
| containerLogMaxFiles  | ContainerLogMaxFiles - The maximum number of container log files that can be present for a container. The number must be ≥ 2.                                                           | int32                                           |
| containerLogMaxSizeMB | ContainerLogMaxSizeMB - The maximum size in MB of a container log file before it is rotated.                                                                                            | int32                                           |
| cpuCfsQuota           | CPUCfsQuota - Enable CPU CFS quota enforcement for containers that specify CPU limits.                                                                                                  | bool                                            |
| cpuCfsQuotaPeriod     | CPUCfsQuotaPeriod - Sets CPU CFS quota period value. Must end in "ms", e.g. "100ms"                                                                                                     | string                                          |
| cpuManagerPolicy      | CPUManagerPolicy - CPU Manager policy to use.                                                                                                                                           | [CPUManagerPolicy](#CPUManagerPolicy)           |
| failSwapOn            | FailSwapOn - If set to true it will make the Kubelet fail to start if swap is enabled on the node.                                                                                      | bool                                            |
| imageGcHighThreshold  | ImageGcHighThreshold - The percent of disk usage after which image garbage collection is always run. Valid values are 0-100 (inclusive).                                                | int32                                           |
| imageGcLowThreshold   | ImageGcLowThreshold - The percent of disk usage before which image garbage collection is never run. Valid values are 0-100 (inclusive) and must be less than `imageGcHighThreshold`.    | int32                                           |
| podMaxPids            | PodMaxPids - The maximum number of processes per pod. Must not exceed kernel PID limit. -1 disables the limit.                                                                          | int32                                           |
| topologyManagerPolicy | TopologyManagerPolicy - Topology Manager policy to use.                                                                                                                                 | [TopologyManagerPolicy](#TopologyManagerPolicy) |

<a id="KubeletDiskType"></a>KubeletDiskType
-------------------------------------------

Used by: [AzureManagedMachinePoolSpec](#AzureManagedMachinePoolSpec).

<a id="LBType"></a>LBType
-------------------------

Used by: [LoadBalancerClassSpec](#LoadBalancerClassSpec).

| Value      | Description |
|------------|-------------|
| "Internal" |             |
| "Public"   |             |

<a id="LinuxOSConfig"></a>LinuxOSConfig
---------------------------------------

LinuxOSConfig specifies the custom Linux OS settings and configurations. See also [AKS doc](https://learn.microsoft.com/azure/aks/custom-node-configuration#linux-os-custom-configuration). <br/>

Used by: [AzureManagedMachinePoolSpec](#AzureManagedMachinePoolSpec).

| Property                   | Description                                                                                                                                                                                                                                                                                                                    | Type                                                    |
|----------------------------|--------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|---------------------------------------------------------|
| swapFileSizeMB             | SwapFileSizeMB specifies size in MB of a swap file will be created on the agent nodes from this node pool. Max value of SwapFileSizeMB should be the size of temporary disk(/dev/sdb). Must be at least 1. See also [AKS doc](https://learn.microsoft.com/azure/virtual-machines/managed-disks-overview#temporary-disk). <br/> | int32                                                   |
| sysctls                    | Sysctl specifies the settings for Linux agent nodes.                                                                                                                                                                                                                                                                           | [SysctlConfig](#SysctlConfig)                           |
| transparentHugePageDefrag  | TransparentHugePageDefrag specifies whether the kernel should make aggressive use of memory compaction to make more hugepages available. See also [Linux doc](https://www.kernel.org/doc/html/latest/admin-guide/mm/transhuge.html#admin-guide-transhuge for more details.). <br/>                                             | [TransparentHugePageOption](#TransparentHugePageOption) |
| transparentHugePageEnabled | TransparentHugePageEnabled specifies various modes of Transparent Hugepages. See also [Linux doc](https://www.kernel.org/doc/html/latest/admin-guide/mm/transhuge.html#admin-guide-transhuge for more details.). <br/>                                                                                                         | [TransparentHugePageOption](#TransparentHugePageOption) |

<a id="LoadBalancerClassSpec"></a>LoadBalancerClassSpec
-------------------------------------------------------

LoadBalancerClassSpec defines the LoadBalancerSpec properties that may be shared across several Azure clusters.

Used by: [NetworkTemplateSpec](#NetworkTemplateSpec), [NetworkTemplateSpec](#NetworkTemplateSpec), and [NetworkTemplateSpec](#NetworkTemplateSpec).

| Property             | Description                                                             | Type              |
|----------------------|-------------------------------------------------------------------------|-------------------|
| idleTimeoutInMinutes | IdleTimeoutInMinutes specifies the timeout for the TCP idle connection. | int32             |
| sku                  |                                                                         | [SKU](#SKU)       |
| type                 |                                                                         | [LBType](#LBType) |

<a id="LoadBalancerProfile"></a>LoadBalancerProfile
---------------------------------------------------

LoadBalancerProfile - Profile of the cluster load balancer. At most one of `managedOutboundIPs`, `outboundIPPrefixes`, or `outboundIPs` may be specified. See also [AKS doc](https://learn.microsoft.com/azure/aks/load-balancer-standard). <br/>

Used by: [AzureManagedControlPlaneSpec](#AzureManagedControlPlaneSpec).

| Property               | Description                                                                                                                                                                                                        | Type     |
|------------------------|--------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|----------|
| allocatedOutboundPorts | AllocatedOutboundPorts - Desired number of allocated SNAT ports per VM. Allowed values must be in the range of 0 to 64000 (inclusive). The default value is 0 which results in Azure dynamically allocating ports. | int32    |
| idleTimeoutInMinutes   | IdleTimeoutInMinutes - Desired outbound flow idle timeout in minutes. Allowed values must be in the range of 4 to 120 (inclusive). The default value is 30 minutes.                                                | int32    |
| managedOutboundIPs     | ManagedOutboundIPs - Desired managed outbound IPs for the cluster load balancer.                                                                                                                                   | int32    |
| outboundIPPrefixes     | OutboundIPPrefixes - Desired outbound IP Prefix resources for the cluster load balancer.                                                                                                                           | string[] |
| outboundIPs            | OutboundIPs - Desired outbound IP resources for the cluster load balancer.                                                                                                                                         | string[] |

<a id="LoadBalancerSpec"></a>LoadBalancerSpec
---------------------------------------------

LoadBalancerSpec defines an Azure load balancer.

Used by: [NetworkSpec](#NetworkSpec), [NetworkSpec](#NetworkSpec), and [NetworkSpec](#NetworkSpec).

| Property                                        | Description                                                                           | Type                        |
|-------------------------------------------------|---------------------------------------------------------------------------------------|-----------------------------|
| [LoadBalancerClassSpec](#LoadBalancerClassSpec) |                                                                                       |                             |
| backendPool                                     | BackendPool describes the backend pool of the load balancer.                          | [BackendPool](#BackendPool) |
| frontendIPs                                     |                                                                                       | [FrontendIP[]](#FrontendIP) |
| frontendIPsCount                                | FrontendIPsCount specifies the number of frontend IP addresses for the load balancer. | int32                       |
| id                                              | ID is the Azure resource ID of the load balancer. READ-ONLY                           | string                      |
| name                                            |                                                                                       | string                      |

<a id="ManagedControlPlaneIdentityType"></a>ManagedControlPlaneIdentityType
---------------------------------------------------------------------------

Used by: [Identity](#Identity).

<a id="ManagedControlPlaneOutboundType"></a>ManagedControlPlaneOutboundType
---------------------------------------------------------------------------

Used by: [AzureManagedControlPlaneSpec](#AzureManagedControlPlaneSpec).

<a id="ManagedControlPlaneSubnet"></a>ManagedControlPlaneSubnet
---------------------------------------------------------------

ManagedControlPlaneSubnet describes a subnet for an AKS cluster.

Used by: [ManagedControlPlaneVirtualNetwork](#ManagedControlPlaneVirtualNetwork).

| Property         | Description                                                                                 | Type             |
|------------------|---------------------------------------------------------------------------------------------|------------------|
| cidrBlock        |                                                                                             | string           |
| name             |                                                                                             | string           |
| privateEndpoints | PrivateEndpoints is a slice of Virtual Network private endpoints to create for the subnets. | PrivateEndpoints |
| serviceEndpoints | ServiceEndpoints is a slice of Virtual Network service endpoints to enable for the subnets. | ServiceEndpoints |

<a id="ManagedControlPlaneVirtualNetwork"></a>ManagedControlPlaneVirtualNetwork
-------------------------------------------------------------------------------

ManagedControlPlaneVirtualNetwork describes a virtual network required to provision AKS clusters.

Used by: [AzureManagedControlPlaneSpec](#AzureManagedControlPlaneSpec).

| Property      | Description                                                                    | Type                                                    |
|---------------|--------------------------------------------------------------------------------|---------------------------------------------------------|
| cidrBlock     |                                                                                | string                                                  |
| name          |                                                                                | string                                                  |
| resourceGroup | ResourceGroup is the name of the Azure resource group for the VNet and Subnet. | string                                                  |
| subnet        | Immutable except for `serviceEndpoints`.                                       | [ManagedControlPlaneSubnet](#ManagedControlPlaneSubnet) |

<a id="ManagedDiskParameters"></a>ManagedDiskParameters
-------------------------------------------------------

ManagedDiskParameters defines the parameters of a managed disk.

Used by: [DataDisk](#DataDisk), and [OSDisk](#OSDisk).

| Property           | Description                                                                                            | Type                                                        |
|--------------------|--------------------------------------------------------------------------------------------------------|-------------------------------------------------------------|
| diskEncryptionSet  | DiskEncryptionSet specifies the customer-managed disk encryption set resource id for the managed disk. | [DiskEncryptionSetParameters](#DiskEncryptionSetParameters) |
| securityProfile    | SecurityProfile specifies the security profile for the managed disk.                                   | [VMDiskSecurityProfile](#VMDiskSecurityProfile)             |
| storageAccountType |                                                                                                        | string                                                      |

<a id="ManagedMachinePoolScaling"></a>ManagedMachinePoolScaling
---------------------------------------------------------------

ManagedMachinePoolScaling specifies scaling options.

Used by: [AzureManagedMachinePoolSpec](#AzureManagedMachinePoolSpec).

| Property | Description                                              | Type  |
|----------|----------------------------------------------------------|-------|
| maxSize  | MaxSize is the maximum number of nodes for auto-scaling. | int32 |
| minSize  | MinSize is the minimum number of nodes for auto-scaling. | int32 |

<a id="NatGateway"></a>NatGateway
---------------------------------

NatGateway defines an Azure NAT gateway. NAT gateway resources are part of Vnet NAT and provide outbound Internet connectivity for subnets of a virtual network.

Used by: [SubnetSpec](#SubnetSpec).

| Property                                    | Description                                               | Type                          |
|---------------------------------------------|-----------------------------------------------------------|-------------------------------|
| [NatGatewayClassSpec](#NatGatewayClassSpec) |                                                           |                               |
| id                                          | ID is the Azure resource ID of the NAT gateway. READ-ONLY | string                        |
| ip                                          |                                                           | [PublicIPSpec](#PublicIPSpec) |

<a id="NatGatewayClassSpec"></a>NatGatewayClassSpec
---------------------------------------------------

NatGatewayClassSpec defines a NAT gateway class specification.

Used by: [SubnetTemplateSpec](#SubnetTemplateSpec).

| Property | Description | Type   |
|----------|-------------|--------|
| name     |             | string |

<a id="NetworkClassSpec"></a>NetworkClassSpec
---------------------------------------------

NetworkClassSpec defines the NetworkSpec properties that may be shared across several Azure clusters.

| Property           | Description                                                         | Type   |
|--------------------|---------------------------------------------------------------------|--------|
| privateDNSZoneName | PrivateDNSZoneName defines the zone name for the Azure Private DNS. | string |

<a id="NetworkInterface"></a>NetworkInterface
---------------------------------------------

NetworkInterface defines a network interface.

Used by: [AzureMachineSpec](#AzureMachineSpec).

| Property              | Description                                                                                                                                                                                                                                                                           | Type   |
|-----------------------|---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|--------|
| acceleratedNetworking | AcceleratedNetworking enables or disables Azure accelerated networking. If omitted, it will be set based on whether the requested VMSize supports accelerated networking. If AcceleratedNetworking is set to true with a VMSize that does not support it, Azure will return an error. | bool   |
| privateIPConfigs      | PrivateIPConfigs specifies the number of private IP addresses to attach to the interface. Defaults to 1 if not specified.                                                                                                                                                             | int    |
| subnetName            | SubnetName specifies the subnet in which the new network interface will be placed.                                                                                                                                                                                                    | string |

<a id="NetworkPluginMode"></a>NetworkPluginMode
-----------------------------------------------

Used by: [AzureManagedControlPlaneSpec](#AzureManagedControlPlaneSpec).

<a id="NetworkSpec"></a>NetworkSpec
-----------------------------------

NetworkSpec specifies what the Azure networking resources should look like.

Used by: [AzureClusterSpec](#AzureClusterSpec).

| Property                              | Description                                                                                                                                                                                                    | Type                                  |
|---------------------------------------|----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|---------------------------------------|
| [NetworkClassSpec](#NetworkClassSpec) |                                                                                                                                                                                                                |                                       |
| apiServerLB                           | APIServerLB is the configuration for the control-plane load balancer.                                                                                                                                          | [LoadBalancerSpec](#LoadBalancerSpec) |
| controlPlaneOutboundLB                | ControlPlaneOutboundLB is the configuration for the control-plane outbound load balancer. This is different from APIServerLB, and is used only in private clusters (optionally) for enabling outbound traffic. | [LoadBalancerSpec](#LoadBalancerSpec) |
| nodeOutboundLB                        | NodeOutboundLB is the configuration for the node outbound load balancer.                                                                                                                                       | [LoadBalancerSpec](#LoadBalancerSpec) |
| subnets                               | Subnets is the configuration for the control-plane subnet and the node subnet.                                                                                                                                 | Subnets                               |
| vnet                                  | Vnet is the configuration for the Azure virtual network.                                                                                                                                                       | [VnetSpec](#VnetSpec)                 |

<a id="NetworkTemplateSpec"></a>NetworkTemplateSpec
---------------------------------------------------

NetworkTemplateSpec specifies a network template.

Used by: [AzureClusterTemplateResourceSpec](#AzureClusterTemplateResourceSpec).

| Property                              | Description                                                                                                                                                                                                    | Type                                            |
|---------------------------------------|----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|-------------------------------------------------|
| [NetworkClassSpec](#NetworkClassSpec) |                                                                                                                                                                                                                |                                                 |
| apiServerLB                           | APIServerLB is the configuration for the control-plane load balancer.                                                                                                                                          | [LoadBalancerClassSpec](#LoadBalancerClassSpec) |
| controlPlaneOutboundLB                | ControlPlaneOutboundLB is the configuration for the control-plane outbound load balancer. This is different from APIServerLB, and is used only in private clusters (optionally) for enabling outbound traffic. | [LoadBalancerClassSpec](#LoadBalancerClassSpec) |
| nodeOutboundLB                        | NodeOutboundLB is the configuration for the node outbound load balancer.                                                                                                                                       | [LoadBalancerClassSpec](#LoadBalancerClassSpec) |
| subnets                               | Subnets is the configuration for the control-plane subnet and the node subnet.                                                                                                                                 | SubnetTemplatesSpec                             |
| vnet                                  | Vnet is the configuration for the Azure virtual network.                                                                                                                                                       | [VnetTemplateSpec](#VnetTemplateSpec)           |

<a id="NodePoolMode"></a>NodePoolMode
-------------------------------------

<a id="OIDCIssuerProfile"></a>OIDCIssuerProfile
-----------------------------------------------

OIDCIssuerProfile is the OIDC issuer profile of the Managed Cluster. See also [AKS doc](https://learn.microsoft.com/en-us/azure/aks/use-oidc-issuer). <br/>

Used by: [AzureManagedControlPlaneSpec](#AzureManagedControlPlaneSpec).

| Property | Description                                    | Type |
|----------|------------------------------------------------|------|
| enabled  | Enabled is whether the OIDC issuer is enabled. | bool |

<a id="OIDCIssuerProfileStatus"></a>OIDCIssuerProfileStatus
-----------------------------------------------------------

OIDCIssuerProfileStatus is the OIDC issuer profile of the Managed Cluster.

Used by: [AzureManagedControlPlaneStatus](#AzureManagedControlPlaneStatus).

| Property  | Description                                              | Type   |
|-----------|----------------------------------------------------------|--------|
| issuerURL | IssuerURL is the OIDC issuer url of the Managed Cluster. | string |

<a id="OSDisk"></a>OSDisk
-------------------------

OSDisk defines the operating system disk for a VM. <br/>WARNING: this requires any updates to ManagedDisk to be manually converted. This is due to the odd issue with conversion-gen where the warning message generated uses a relative directory import rather than the fully qualified import when generating outside of the GOPATH.

Used by: [AzureMachineSpec](#AzureMachineSpec), and [osDiskTestInput](#osDiskTestInput).

| Property         | Description                                                                                        | Type                                            |
|------------------|----------------------------------------------------------------------------------------------------|-------------------------------------------------|
| cachingType      | CachingType specifies the caching requirements.                                                    | string                                          |
| diffDiskSettings |                                                                                                    | [DiffDiskSettings](#DiffDiskSettings)           |
| diskSizeGB       | DiskSizeGB is the size in GB to assign to the OS disk. Will have a default of 30GB if not provided | int32                                           |
| managedDisk      | ManagedDisk specifies the Managed Disk parameters for the OS disk.                                 | [ManagedDiskParameters](#ManagedDiskParameters) |
| osType           |                                                                                                    | string                                          |

<a id="OrchestrationModeType"></a>OrchestrationModeType
-------------------------------------------------------

<a id="PrivateEndpointSpec"></a>PrivateEndpointSpec
---------------------------------------------------

PrivateEndpointSpec configures an Azure Private Endpoint.

| Property                      | Description                                                                                                                                                                                                     | Type                                                            |
|-------------------------------|-----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|-----------------------------------------------------------------|
| applicationSecurityGroups     | ApplicationSecurityGroups specifies the Application security group in which the private endpoint IP configuration is included.                                                                                  | string[]                                                        |
| customNetworkInterfaceName    | CustomNetworkInterfaceName specifies the network interface name associated with the private endpoint.                                                                                                           | string                                                          |
| location                      | Location specifies the region to create the private endpoint.                                                                                                                                                   | string                                                          |
| manualApproval                | ManualApproval specifies if the connection approval needs to be done manually or not. Set it true when the network admin does not have access to approve connections to the remote resource. Defaults to false. | bool                                                            |
| name                          | Name specifies the name of the private endpoint.                                                                                                                                                                | string                                                          |
| privateIPAddresses            | PrivateIPAddresses specifies the IP addresses for the network interface associated with the private endpoint. They have to be part of the subnet where the private endpoint is linked.                          | string[]                                                        |
| privateLinkServiceConnections | PrivateLinkServiceConnections specifies Private Link Service Connections of the private endpoint.                                                                                                               | [PrivateLinkServiceConnection[]](#PrivateLinkServiceConnection) |

<a id="PrivateLinkServiceConnection"></a>PrivateLinkServiceConnection
---------------------------------------------------------------------

PrivateLinkServiceConnection defines the specification for a private link service connection associated with a private endpoint.

Used by: [PrivateEndpointSpec](#PrivateEndpointSpec).

| Property             | Description                                                                                                                  | Type     |
|----------------------|------------------------------------------------------------------------------------------------------------------------------|----------|
| groupIDs             | GroupIDs specifies the ID(s) of the group(s) obtained from the remote resource that this private endpoint should connect to. | string[] |
| name                 | Name specifies the name of the private link service.                                                                         | string   |
| privateLinkServiceID | PrivateLinkServiceID specifies the resource ID of the private link service.                                                  | string   |
| requestMessage       | RequestMessage specifies a message passed to the owner of the remote resource with the private endpoint connection request.  | string   |

<a id="ProvisioningState"></a>ProvisioningState
-----------------------------------------------

Used by: [AzureMachineStatus](#AzureMachineStatus).

<a id="PublicIPSpec"></a>PublicIPSpec
-------------------------------------

PublicIPSpec defines the inputs to create an Azure public IP address.

Used by: [AzureBastion](#AzureBastion), [FrontendIP](#FrontendIP), and [NatGateway](#NatGateway).

| Property | Description | Type              |
|----------|-------------|-------------------|
| dnsName  |             | string            |
| ipTags   |             | [IPTag[]](#IPTag) |
| name     |             | string            |

<a id="RateLimitConfig"></a>RateLimitConfig
-------------------------------------------

RateLimitConfig indicates the rate limit config options.

Used by: [RateLimitSpec](#RateLimitSpec).

| Property                          | Description | Type              |
|-----------------------------------|-------------|-------------------|
| cloudProviderRateLimit            |             | bool              |
| cloudProviderRateLimitBucket      |             | int               |
| cloudProviderRateLimitBucketWrite |             | int               |
| cloudProviderRateLimitQPS         |             | resource.Quantity |
| cloudProviderRateLimitQPSWrite    |             | resource.Quantity |

<a id="RateLimitSpec"></a>RateLimitSpec
---------------------------------------

RateLimitSpec represents the rate limit configuration for a particular kind of resource. Eg. loadBalancerRateLimit is used to configure rate limits for load balancers. This eventually gets converted to CloudProviderRateLimitConfig that cloud-provider-azure expects. See: https://github.com/kubernetes-sigs/cloud-provider-azure/blob/d585c2031925b39c925624302f22f8856e29e352/pkg/provider/azure_ratelimit.go#L25 We cannot use CloudProviderRateLimitConfig directly because floating point values are not supported in controller-tools. See: https://github.com/kubernetes-sigs/controller-tools/issues/245

Used by: [CloudProviderConfigOverrides](#CloudProviderConfigOverrides).

| Property | Description                              | Type                                |
|----------|------------------------------------------|-------------------------------------|
| config   |                                          | [RateLimitConfig](#RateLimitConfig) |
| name     | Name is the name of the rate limit spec. | string                              |

<a id="ResourceLifecycle"></a>ResourceLifecycle
-----------------------------------------------

Used by: [BuildParams](#BuildParams).

| Value    | Description |
|----------|-------------|
| "owned"  |             |
| "shared" |             |

<a id="RouteTable"></a>RouteTable
---------------------------------

RouteTable defines an Azure route table.

Used by: [SubnetSpec](#SubnetSpec).

| Property | Description                                               | Type   |
|----------|-----------------------------------------------------------|--------|
| id       | ID is the Azure resource ID of the route table. READ-ONLY | string |
| name     |                                                           | string |

<a id="SKU"></a>SKU
-------------------

Used by: [LoadBalancerClassSpec](#LoadBalancerClassSpec).

| Value      | Description |
|------------|-------------|
| "Standard" |             |

<a id="SecurityEncryptionType"></a>SecurityEncryptionType
---------------------------------------------------------

Used by: [VMDiskSecurityProfile](#VMDiskSecurityProfile).

<a id="SecurityGroup"></a>SecurityGroup
---------------------------------------

SecurityGroup defines an Azure security group.

Used by: [SubnetSpec](#SubnetSpec).

| Property                                  | Description                                                  | Type   |
|-------------------------------------------|--------------------------------------------------------------|--------|
| [SecurityGroupClass](#SecurityGroupClass) |                                                              |        |
| id                                        | ID is the Azure resource ID of the security group. READ-ONLY | string |
| name                                      |                                                              | string |

<a id="SecurityGroupClass"></a>SecurityGroupClass
-------------------------------------------------

SecurityGroupClass defines the SecurityGroup properties that may be shared across several Azure clusters.

Used by: [SubnetTemplateSpec](#SubnetTemplateSpec).

| Property      | Description | Type          |
|---------------|-------------|---------------|
| securityRules |             | SecurityRules |
| tags          |             | Tags          |

<a id="SecurityGroupProtocol"></a>SecurityGroupProtocol
-------------------------------------------------------

Used by: [SecurityRule](#SecurityRule).

| Value  | Description |
|--------|-------------|
| "*"    |             |
| "Tcp"  |             |
| "Udp"  |             |
| "Icmp" |             |

<a id="SecurityProfile"></a>SecurityProfile
-------------------------------------------

SecurityProfile specifies the Security profile settings for a virtual machine or virtual machine scale set.

Used by: [AzureMachineSpec](#AzureMachineSpec).

| Property         | Description                                                                                                                                                                                                                             | Type                            |
|------------------|-----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|---------------------------------|
| encryptionAtHost | This field indicates whether Host Encryption should be enabled or disabled for a virtual machine or virtual machine scale set. This should be disabled when SecurityEncryptionType is set to DiskWithVMGuestState. Default is disabled. | bool                            |
| securityType     | SecurityType specifies the SecurityType of the virtual machine. It has to be set to any specified value to enable UefiSettings. The default behavior is: UefiSettings will not be enabled unless this property is set.                  | [SecurityTypes](#SecurityTypes) |
| uefiSettings     | UefiSettings specifies the security settings like secure boot and vTPM used while creating the virtual machine.                                                                                                                         | [UefiSettings](#UefiSettings)   |

<a id="SecurityRule"></a>SecurityRule
-------------------------------------

SecurityRule defines an Azure security rule for security groups.

| Property         | Description                                                                                                                                                                                                                                                                  | Type                                            |
|------------------|------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|-------------------------------------------------|
| action           | Action specifies whether network traffic is allowed or denied. Can either be "Allow" or "Deny". Defaults to "Allow".                                                                                                                                                         | [SecurityRuleAccess](#SecurityRuleAccess)       |
| description      | A description for this rule. Restricted to 140 chars.                                                                                                                                                                                                                        | string                                          |
| destination      | Destination is the destination address prefix. CIDR or destination IP range. Asterix '*' can also be used to match all source IPs. Default tags such as 'VirtualNetwork', 'AzureLoadBalancer' and 'Internet' can also be used.                                               | string                                          |
| destinationPorts | DestinationPorts specifies the destination port or range. Integer or range between 0 and 65535. Asterix '*' can also be used to match all ports.                                                                                                                             | string                                          |
| direction        | Direction indicates whether the rule applies to inbound, or outbound traffic. "Inbound" or "Outbound".                                                                                                                                                                       | [SecurityRuleDirection](#SecurityRuleDirection) |
| name             | Name is a unique name within the network security group.                                                                                                                                                                                                                     | string                                          |
| priority         | Priority is a number between 100 and 4096. Each rule should have a unique value for priority. Rules are processed in priority order, with lower numbers processed before higher numbers. Once traffic matches a rule, processing stops.                                      | int32                                           |
| protocol         | Protocol specifies the protocol type. "Tcp", "Udp", "Icmp", or "*".                                                                                                                                                                                                          | [SecurityGroupProtocol](#SecurityGroupProtocol) |
| source           | Source specifies the CIDR or source IP range. Asterix '*' can also be used to match all source IPs. Default tags such as 'VirtualNetwork', 'AzureLoadBalancer' and 'Internet' can also be used. If this is an ingress rule, specifies where network traffic originates from. | string                                          |
| sourcePorts      | SourcePorts specifies source port or range. Integer or range between 0 and 65535. Asterix '*' can also be used to match all ports.                                                                                                                                           | string                                          |

<a id="SecurityRuleAccess"></a>SecurityRuleAccess
-------------------------------------------------

Used by: [SecurityRule](#SecurityRule).

<a id="SecurityRuleDirection"></a>SecurityRuleDirection
-------------------------------------------------------

Used by: [SecurityRule](#SecurityRule).

| Value      | Description |
|------------|-------------|
| "Inbound"  |             |
| "Outbound" |             |

<a id="SecurityTypes"></a>SecurityTypes
---------------------------------------

Used by: [SecurityProfile](#SecurityProfile).

<a id="ServiceEndpointSpec"></a>ServiceEndpointSpec
---------------------------------------------------

ServiceEndpointSpec configures an Azure Service Endpoint.

| Property  | Description | Type     |
|-----------|-------------|----------|
| locations |             | string[] |
| service   |             | string   |

<a id="SkipNodesWithLocalStorage"></a>SkipNodesWithLocalStorage
---------------------------------------------------------------

Used by: [AutoScalerProfile](#AutoScalerProfile).

<a id="SkipNodesWithSystemPods"></a>SkipNodesWithSystemPods
-----------------------------------------------------------

Used by: [AutoScalerProfile](#AutoScalerProfile).

<a id="SpotEvictionPolicy"></a>SpotEvictionPolicy
-------------------------------------------------

Used by: [SpotVMOptions](#SpotVMOptions).

<a id="SpotVMOptions"></a>SpotVMOptions
---------------------------------------

SpotVMOptions defines the options relevant to running the Machine on Spot VMs.

Used by: [AzureMachineSpec](#AzureMachineSpec).

| Property       | Description                                                                                                           | Type                                      |
|----------------|-----------------------------------------------------------------------------------------------------------------------|-------------------------------------------|
| evictionPolicy | EvictionPolicy defines the behavior of the virtual machine when it is evicted. It can be either Delete or Deallocate. | [SpotEvictionPolicy](#SpotEvictionPolicy) |
| maxPrice       | MaxPrice defines the maximum price the user is willing to pay for Spot VM instances                                   | resource.Quantity                         |

<a id="SubnetClassSpec"></a>SubnetClassSpec
-------------------------------------------

SubnetClassSpec defines the SubnetSpec properties that may be shared across several Azure clusters.

| Property         | Description                                                                                                | Type                      |
|------------------|------------------------------------------------------------------------------------------------------------|---------------------------|
| cidrBlocks       | CIDRBlocks defines the subnet's address space, specified as one or more address prefixes in CIDR notation. | string[]                  |
| name             | Name defines a name for the subnet resource.                                                               | string                    |
| privateEndpoints | PrivateEndpoints defines a list of private endpoints that should be attached to this subnet.               | PrivateEndpoints          |
| role             | Role defines the subnet role (eg. Node, ControlPlane)                                                      | [SubnetRole](#SubnetRole) |
| serviceEndpoints | ServiceEndpoints is a slice of Virtual Network service endpoints to enable for the subnets.                | ServiceEndpoints          |

<a id="SubnetRole"></a>SubnetRole
---------------------------------

Used by: [SubnetClassSpec](#SubnetClassSpec).

<a id="SubnetSpec"></a>SubnetSpec
---------------------------------

SubnetSpec configures an Azure subnet.

Used by: [AzureBastion](#AzureBastion).

| Property                            | Description                                                                                    | Type                            |
|-------------------------------------|------------------------------------------------------------------------------------------------|---------------------------------|
| [SubnetClassSpec](#SubnetClassSpec) |                                                                                                |                                 |
| id                                  | ID is the Azure resource ID of the subnet. READ-ONLY                                           | string                          |
| natGateway                          | NatGateway associated with this subnet.                                                        | [NatGateway](#NatGateway)       |
| routeTable                          | RouteTable defines the route table that should be attached to this subnet.                     | [RouteTable](#RouteTable)       |
| securityGroup                       | SecurityGroup defines the NSG (network security group) that should be attached to this subnet. | [SecurityGroup](#SecurityGroup) |

<a id="SubnetTemplateSpec"></a>SubnetTemplateSpec
-------------------------------------------------

SubnetTemplateSpec specifies a template for a subnet.

Used by: [AzureBastionTemplateSpec](#AzureBastionTemplateSpec).

| Property                            | Description                                                                                    | Type                                        |
|-------------------------------------|------------------------------------------------------------------------------------------------|---------------------------------------------|
| [SubnetClassSpec](#SubnetClassSpec) |                                                                                                |                                             |
| natGateway                          | NatGateway associated with this subnet.                                                        | [NatGatewayClassSpec](#NatGatewayClassSpec) |
| securityGroup                       | SecurityGroup defines the NSG (network security group) that should be attached to this subnet. | [SecurityGroupClass](#SecurityGroupClass)   |

<a id="SysctlConfig"></a>SysctlConfig
-------------------------------------

SysctlConfig specifies the settings for Linux agent nodes.

Used by: [LinuxOSConfig](#LinuxOSConfig).

| Property                       | Description                                                                                                                                                                                                                                                                                                                                                                                                                                | Type   |
|--------------------------------|--------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|--------|
| fsAioMaxNr                     | FsAioMaxNr specifies the maximum number of system-wide asynchronous io requests. Valid values are 65536-6553500 (inclusive). Maps to fs.aio-max-nr.                                                                                                                                                                                                                                                                                        | int32  |
| fsFileMax                      | FsFileMax specifies the max number of file-handles that the Linux kernel will allocate, by increasing increases the maximum number of open files permitted. Valid values are 8192-12000500 (inclusive). Maps to fs.file-max.                                                                                                                                                                                                               | int32  |
| fsInotifyMaxUserWatches        | FsInotifyMaxUserWatches specifies the number of file watches allowed by the system. Each watch is roughly 90 bytes on a 32-bit kernel, and roughly 160 bytes on a 64-bit kernel. Valid values are 781250-2097152 (inclusive). Maps to fs.inotify.max_user_watches.                                                                                                                                                                         | int32  |
| fsNrOpen                       | FsNrOpen specifies the maximum number of file-handles a process can allocate. Valid values are 8192-20000500 (inclusive). Maps to fs.nr_open.                                                                                                                                                                                                                                                                                              | int32  |
| kernelThreadsMax               | KernelThreadsMax specifies the maximum number of all threads that can be created. Valid values are 20-513785 (inclusive). Maps to kernel.threads-max.                                                                                                                                                                                                                                                                                      | int32  |
| netCoreNetdevMaxBacklog        | NetCoreNetdevMaxBacklog specifies maximum number of packets, queued on the INPUT side, when the interface receives packets faster than kernel can process them. Valid values are 1000-3240000 (inclusive). Maps to net.core.netdev_max_backlog.                                                                                                                                                                                            | int32  |
| netCoreOptmemMax               | NetCoreOptmemMax specifies the maximum ancillary buffer size (option memory buffer) allowed per socket. Socket option memory is used in a few cases to store extra structures relating to usage of the socket. Valid values are 20480-4194304 (inclusive). Maps to net.core.optmem_max.                                                                                                                                                    | int32  |
| netCoreRmemDefault             | NetCoreRmemDefault specifies the default receive socket buffer size in bytes. Valid values are 212992-134217728 (inclusive). Maps to net.core.rmem_default.                                                                                                                                                                                                                                                                                | int32  |
| netCoreRmemMax                 | NetCoreRmemMax specifies the maximum receive socket buffer size in bytes. Valid values are 212992-134217728 (inclusive). Maps to net.core.rmem_max.                                                                                                                                                                                                                                                                                        | int32  |
| netCoreSomaxconn               | NetCoreSomaxconn specifies maximum number of connection requests that can be queued for any given listening socket. An upper limit for the value of the backlog parameter passed to the listen(2)(https://man7.org/linux/man-pages/man2/listen.2.html) function. If the backlog argument is greater than the somaxconn, then it's silently truncated to this limit. Valid values are 4096-3240000 (inclusive). Maps to net.core.somaxconn. | int32  |
| netCoreWmemDefault             | NetCoreWmemDefault specifies the default send socket buffer size in bytes. Valid values are 212992-134217728 (inclusive). Maps to net.core.wmem_default.                                                                                                                                                                                                                                                                                   | int32  |
| netCoreWmemMax                 | NetCoreWmemMax specifies the maximum send socket buffer size in bytes. Valid values are 212992-134217728 (inclusive). Maps to net.core.wmem_max.                                                                                                                                                                                                                                                                                           | int32  |
| netIpv4IPLocalPortRange        | NetIpv4IPLocalPortRange is used by TCP and UDP traffic to choose the local port on the agent node. PortRange should be specified in the format "first last". First, being an integer, must be between [1024 - 60999]. Last, being an integer, must be between [32768 - 65000]. Maps to net.ipv4.ip_local_port_range.                                                                                                                       | string |
| netIpv4NeighDefaultGcThresh1   | NetIpv4NeighDefaultGcThresh1 specifies the minimum number of entries that may be in the ARP cache. Garbage collection won't be triggered if the number of entries is below this setting. Valid values are 128-80000 (inclusive). Maps to net.ipv4.neigh.default.gc_thresh1.                                                                                                                                                                | int32  |
| netIpv4NeighDefaultGcThresh2   | NetIpv4NeighDefaultGcThresh2 specifies soft maximum number of entries that may be in the ARP cache. ARP garbage collection will be triggered about 5 seconds after reaching this soft maximum. Valid values are 512-90000 (inclusive). Maps to net.ipv4.neigh.default.gc_thresh2.                                                                                                                                                          | int32  |
| netIpv4NeighDefaultGcThresh3   | NetIpv4NeighDefaultGcThresh3 specified hard maximum number of entries in the ARP cache. Valid values are 1024-100000 (inclusive). Maps to net.ipv4.neigh.default.gc_thresh3.                                                                                                                                                                                                                                                               | int32  |
| netIpv4TCPFinTimeout           | NetIpv4TCPFinTimeout specifies the length of time an orphaned connection will remain in the FIN_WAIT_2 state before it's aborted at the local end. Valid values are 5-120 (inclusive). Maps to net.ipv4.tcp_fin_timeout.                                                                                                                                                                                                                   | int32  |
| netIpv4TCPKeepaliveProbes      | NetIpv4TCPKeepaliveProbes specifies the number of keepalive probes TCP sends out, until it decides the connection is broken. Valid values are 1-15 (inclusive). Maps to net.ipv4.tcp_keepalive_probes.                                                                                                                                                                                                                                     | int32  |
| netIpv4TCPKeepaliveTime        | NetIpv4TCPKeepaliveTime specifies the rate at which TCP sends out a keepalive message when keepalive is enabled. Valid values are 30-432000 (inclusive). Maps to net.ipv4.tcp_keepalive_time.                                                                                                                                                                                                                                              | int32  |
| netIpv4TCPMaxSynBacklog        | NetIpv4TCPMaxSynBacklog specifies the maximum number of queued connection requests that have still not received an acknowledgment from the connecting client. If this number is exceeded, the kernel will begin dropping requests. Valid values are 128-3240000 (inclusive). Maps to net.ipv4.tcp_max_syn_backlog.                                                                                                                         | int32  |
| netIpv4TCPMaxTwBuckets         | NetIpv4TCPMaxTwBuckets specifies maximal number of timewait sockets held by system simultaneously. If this number is exceeded, time-wait socket is immediately destroyed and warning is printed. Valid values are 8000-1440000 (inclusive). Maps to net.ipv4.tcp_max_tw_buckets.                                                                                                                                                           | int32  |
| netIpv4TCPTwReuse              | NetIpv4TCPTwReuse is used to allow to reuse TIME-WAIT sockets for new connections when it's safe from protocol viewpoint. Maps to net.ipv4.tcp_tw_reuse.                                                                                                                                                                                                                                                                                   | bool   |
| netIpv4TCPkeepaliveIntvl       | NetIpv4TCPkeepaliveIntvl specifies the frequency of the probes sent out. Multiplied by tcpKeepaliveprobes, it makes up the time to kill a connection that isn't responding, after probes started. Valid values are 1-75 (inclusive). Maps to net.ipv4.tcp_keepalive_intvl.                                                                                                                                                                 | int32  |
| netNetfilterNfConntrackBuckets | NetNetfilterNfConntrackBuckets specifies the size of hash table used by nf_conntrack module to record the established connection record of the TCP protocol. Valid values are 65536-147456 (inclusive). Maps to net.netfilter.nf_conntrack_buckets.                                                                                                                                                                                        | int32  |
| netNetfilterNfConntrackMax     | NetNetfilterNfConntrackMax specifies the maximum number of connections supported by the nf_conntrack module or the size of connection tracking table. Valid values are 131072-1048576 (inclusive). Maps to net.netfilter.nf_conntrack_max.                                                                                                                                                                                                 | int32  |
| vmMaxMapCount                  | VMMaxMapCount specifies the maximum number of memory map areas a process may have. Maps to vm.max_map_count. Valid values are 65530-262144 (inclusive).                                                                                                                                                                                                                                                                                    | int32  |
| vmSwappiness                   | VMSwappiness specifies aggressiveness of the kernel in swapping memory pages. Higher values will increase aggressiveness, lower values decrease the amount of swap. Valid values are 0-100 (inclusive). Maps to vm.swappiness.                                                                                                                                                                                                             | int32  |
| vmVfsCachePressure             | VMVfsCachePressure specifies the percentage value that controls tendency of the kernel to reclaim the memory, which is used for caching of directory and inode objects. Valid values are 1-500 (inclusive). Maps to vm.vfs_cache_pressure.                                                                                                                                                                                                 | int32  |

<a id="SystemAssignedIdentityRole"></a>SystemAssignedIdentityRole
-----------------------------------------------------------------

SystemAssignedIdentityRole defines the role and scope to assign to the system assigned identity.

Used by: [AzureMachineSpec](#AzureMachineSpec).

| Property     | Description                                                                                                                                                                                                                                          | Type   |
|--------------|------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|--------|
| definitionID | DefinitionID is the ID of the role definition to create for a system assigned identity. It can be an Azure built-in role or a custom role. Refer to built-in roles: https://learn.microsoft.com/en-us/azure/role-based-access-control/built-in-roles | string |
| name         | Name is the name of the role assignment to create for a system assigned identity. It can be any valid UUID. If not specified, a random UUID will be generated.                                                                                       | string |
| scope        | Scope is the scope that the role assignment or definition applies to. The scope can be any REST resource instance. If not specified, the scope will be the subscription.                                                                             | string |

<a id="Taint"></a>Taint
-----------------------

Taint represents a Kubernetes taint.

| Property | Description                               | Type                        |
|----------|-------------------------------------------|-----------------------------|
| effect   | Effect specifies the effect for the taint | [TaintEffect](#TaintEffect) |
| key      | Key is the key of the taint               | string                      |
| value    | Value is the value of the taint           | string                      |

<a id="TaintEffect"></a>TaintEffect
-----------------------------------

Used by: [Taint](#Taint).

<a id="TopologyManagerPolicy"></a>TopologyManagerPolicy
-------------------------------------------------------

Used by: [KubeletConfig](#KubeletConfig).

<a id="TransparentHugePageOption"></a>TransparentHugePageOption
---------------------------------------------------------------

Used by: [LinuxOSConfig](#LinuxOSConfig), and [LinuxOSConfig](#LinuxOSConfig).

<a id="UefiSettings"></a>UefiSettings
-------------------------------------

UefiSettings specifies the security settings like secure boot and vTPM used while creating the virtual machine.

Used by: [SecurityProfile](#SecurityProfile).

| Property          | Description                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                        | Type |
|-------------------|----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|------|
| secureBootEnabled | SecureBootEnabled specifies whether secure boot should be enabled on the virtual machine. Secure Boot verifies the digital signature of all boot components and halts the boot process if signature verification fails. If omitted, the platform chooses a default, which is subject to change over time, currently that default is false.                                                                                                                                                                                                         | bool |
| vTpmEnabled       | VTpmEnabled specifies whether vTPM should be enabled on the virtual machine. When true it enables the virtualized trusted platform module measurements to create a known good boot integrity policy baseline. The integrity policy baseline is used for comparison with measurements from subsequent VM boots to determine if anything has changed. This is required to be set to Enabled if SecurityEncryptionType is defined. If omitted, the platform chooses a default, which is subject to change over time, currently that default is false. | bool |

<a id="UserAssignedIdentity"></a>UserAssignedIdentity
-----------------------------------------------------

UserAssignedIdentity defines the user-assigned identities provided by the user to be assigned to Azure resources.

Used by: [AzureMachineSpec](#AzureMachineSpec).

| Property   | Description                                                                                                                                                                                                                                             | Type   |
|------------|---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|--------|
| providerID | ProviderID is the identification ID of the user-assigned Identity, the format of an identity is: 'azure:///subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ManagedIdentity/userAssignedIdentities/{identityName}' | string |

<a id="UserManagedBootDiagnostics"></a>UserManagedBootDiagnostics
-----------------------------------------------------------------

UserManagedBootDiagnostics provides a reference to a user-managed storage account.

Used by: [BootDiagnostics](#BootDiagnostics).

| Property          | Description                                                                                                                                                                                                                                                                                                                                                                                                                                                                                 | Type                |
|-------------------|---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|---------------------|
| storageAccountURI | StorageAccountURI is the URI of the user-managed storage account. The URI typically will be `https://<mystorageaccountname>.blob.core.windows.net/` but may differ if you are using Azure DNS zone endpoints. You can find the correct endpoint by looking for the Blob Primary Endpoint in the endpoints tab in the Azure console or with the CLI by issuing `az storage account list --query='[].{name: name, "resource group": resourceGroup, "blob endpoint": primaryEndpoints.blob}'`. | string<br/>Required |

<a id="VMDiskSecurityProfile"></a>VMDiskSecurityProfile
-------------------------------------------------------

VMDiskSecurityProfile specifies the security profile settings for the managed disk. It can be set only for Confidential VMs.

Used by: [ManagedDiskParameters](#ManagedDiskParameters).

| Property               | Description                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                 | Type                                                        |
|------------------------|-----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|-------------------------------------------------------------|
| diskEncryptionSet      | DiskEncryptionSet specifies the customer-managed disk encryption set resource id for the managed disk that is used for Customer Managed Key encrypted ConfidentialVM OS Disk and VMGuest blob.                                                                                                                                                                                                                                                                                                                              | [DiskEncryptionSetParameters](#DiskEncryptionSetParameters) |
| securityEncryptionType | SecurityEncryptionType specifies the encryption type of the managed disk. It is set to DiskWithVMGuestState to encrypt the managed disk along with the VMGuestState blob, and to VMGuestStateOnly to encrypt the VMGuestState blob only. When set to VMGuestStateOnly, VirtualizedTrustedPlatformModule should be set to Enabled. When set to DiskWithVMGuestState, EncryptionAtHost should be disabled, SecureBoot and VirtualizedTrustedPlatformModule should be set to Enabled. It can be set only for Confidential VMs. | [SecurityEncryptionType](#SecurityEncryptionType)           |

<a id="VMExtension"></a>VMExtension
-----------------------------------

VMExtension specifies the parameters for a custom VM extension.

Used by: [AzureMachineSpec](#AzureMachineSpec).

| Property          | Description                                                                 | Type   |
|-------------------|-----------------------------------------------------------------------------|--------|
| name              | Name is the name of the extension.                                          | string |
| protectedSettings | ProtectedSettings is a JSON formatted protected settings for the extension. | Tags   |
| publisher         | Publisher is the name of the extension handler publisher.                   | string |
| settings          | Settings is a JSON formatted public settings for the extension.             | Tags   |
| version           | Version specifies the version of the script handler.                        | string |

<a id="VMIdentity"></a>VMIdentity
---------------------------------

Used by: [AzureMachineSpec](#AzureMachineSpec).

<a id="VMState"></a>VMState
---------------------------

<a id="VnetClassSpec"></a>VnetClassSpec
---------------------------------------

VnetClassSpec defines the VnetSpec properties that may be shared across several Azure clusters.

| Property   | Description                                                                                                         | Type     |
|------------|---------------------------------------------------------------------------------------------------------------------|----------|
| cidrBlocks | CIDRBlocks defines the virtual network's address space, specified as one or more address prefixes in CIDR notation. | string[] |
| tags       | Tags is a collection of tags describing the resource.                                                               | Tags     |

<a id="VnetPeeringClassSpec"></a>VnetPeeringClassSpec
-----------------------------------------------------

VnetPeeringClassSpec specifies a virtual network peering class.

| Property                 | Description                                                                                                                            | Type                                            |
|--------------------------|----------------------------------------------------------------------------------------------------------------------------------------|-------------------------------------------------|
| forwardPeeringProperties | ForwardPeeringProperties specifies VnetPeeringProperties for peering from the cluster's virtual network to the remote virtual network. | [VnetPeeringProperties](#VnetPeeringProperties) |
| remoteVnetName           | RemoteVnetName defines name of the remote virtual network.                                                                             | string                                          |
| resourceGroup            | ResourceGroup is the resource group name of the remote virtual network.                                                                | string                                          |
| reversePeeringProperties | ReversePeeringProperties specifies VnetPeeringProperties for peering from the remote virtual network to the cluster's virtual network. | [VnetPeeringProperties](#VnetPeeringProperties) |

<a id="VnetPeeringProperties"></a>VnetPeeringProperties
-------------------------------------------------------

VnetPeeringProperties specifies virtual network peering properties.

Used by: [VnetPeeringClassSpec](#VnetPeeringClassSpec), and [VnetPeeringClassSpec](#VnetPeeringClassSpec).

| Property                  | Description                                                                                                                                                                                                                                                                                                                                                                            | Type |
|---------------------------|----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|------|
| allowForwardedTraffic     | AllowForwardedTraffic specifies whether the forwarded traffic from the VMs in the local virtual network will be allowed/disallowed in remote virtual network.                                                                                                                                                                                                                          | bool |
| allowGatewayTransit       | AllowGatewayTransit specifies if gateway links can be used in remote virtual networking to link to this virtual network.                                                                                                                                                                                                                                                               | bool |
| allowVirtualNetworkAccess | AllowVirtualNetworkAccess specifies whether the VMs in the local virtual network space would be able to access the VMs in remote virtual network space.                                                                                                                                                                                                                                | bool |
| useRemoteGateways         | UseRemoteGateways specifies if remote gateways can be used on this virtual network. If the flag is set to true, and allowGatewayTransit on remote peering is also set to true, the virtual network will use the gateways of the remote virtual network for transit. Only one peering can have this flag set to true. This flag cannot be set if virtual network already has a gateway. | bool |

<a id="VnetPeeringSpec"></a>VnetPeeringSpec
-------------------------------------------

VnetPeeringSpec specifies an existing remote virtual network to peer with the AzureCluster's virtual network.

<a id="VnetSpec"></a>VnetSpec
-----------------------------

VnetSpec configures an Azure virtual network.

Used by: [NetworkSpec](#NetworkSpec).

| Property                        | Description                                                                                                                                              | Type         |
|---------------------------------|----------------------------------------------------------------------------------------------------------------------------------------------------------|--------------|
| [VnetClassSpec](#VnetClassSpec) |                                                                                                                                                          |              |
| id                              | ID is the Azure resource ID of the virtual network. READ-ONLY                                                                                            | string       |
| name                            | Name defines a name for the virtual network resource.                                                                                                    | string       |
| peerings                        | Peerings defines a list of peerings of the newly created virtual network with existing virtual networks.                                                 | VnetPeerings |
| resourceGroup                   | ResourceGroup is the name of the resource group of the existing virtual network or the resource group where a managed virtual network should be created. | string       |

<a id="VnetTemplateSpec"></a>VnetTemplateSpec
---------------------------------------------

VnetTemplateSpec defines the desired state of a virtual network.

Used by: [NetworkTemplateSpec](#NetworkTemplateSpec).

| Property                        | Description                                                                                              | Type                     |
|---------------------------------|----------------------------------------------------------------------------------------------------------|--------------------------|
| [VnetClassSpec](#VnetClassSpec) |                                                                                                          |                          |
| peerings                        | Peerings defines a list of peerings of the newly created virtual network with existing virtual networks. | VnetPeeringsTemplateSpec |

<a id="azureMachineWebhook"></a>azureMachineWebhook
---------------------------------------------------

azureMachineWebhook implements a validating and defaulting webhook for AzureMachines.

| Property | Description | Type          |
|----------|-------------|---------------|
| Client   |             | client.Client |

<a id="azureManagedControlPlaneWebhook"></a>azureManagedControlPlaneWebhook
---------------------------------------------------------------------------

azureManagedControlPlaneWebhook implements a validating and defaulting webhook for AzureManagedControlPlane.

| Property | Description | Type          |
|----------|-------------|---------------|
| Client   |             | client.Client |

<a id="azureManagedMachinePoolWebhook"></a>azureManagedMachinePoolWebhook
-------------------------------------------------------------------------

azureManagedMachinePoolWebhook implements a validating and defaulting webhook for AzureManagedMachinePool.

| Property | Description | Type          |
|----------|-------------|---------------|
| Client   |             | client.Client |

<a id="mockClient"></a>mockClient
---------------------------------

| Property      | Description | Type |
|---------------|-------------|------|
| client.Client |             |      |
| ReturnError   |             | bool |

<a id="mockDefaultClient"></a>mockDefaultClient
-----------------------------------------------

| Property       | Description | Type   |
|----------------|-------------|--------|
| client.Client  |             |        |
| SubscriptionID |             | string |

<a id="osDiskTestInput"></a>osDiskTestInput
-------------------------------------------

| Property | Description | Type              |
|----------|-------------|-------------------|
| name     |             | string            |
| osDisk   |             | [OSDisk](#OSDisk) |
| wantErr  |             | bool              |
