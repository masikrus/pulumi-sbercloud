// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Sbercloud
{
    public static class GetCceCluster
    {
        public static Task<GetCceClusterResult> InvokeAsync(GetCceClusterArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetCceClusterResult>("sbercloud:index/getCceCluster:getCceCluster", args ?? new GetCceClusterArgs(), options.WithDefaults());

        public static Output<GetCceClusterResult> Invoke(GetCceClusterInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetCceClusterResult>("sbercloud:index/getCceCluster:getCceCluster", args ?? new GetCceClusterInvokeArgs(), options.WithDefaults());

        public static Output<GetCceClusterResult> Invoke(GetCceClusterInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetCceClusterResult>("sbercloud:index/getCceCluster:getCceCluster", args ?? new GetCceClusterInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetCceClusterArgs : global::Pulumi.InvokeArgs
    {
        [Input("clusterType")]
        public string? ClusterType { get; set; }

        [Input("id")]
        public string? Id { get; set; }

        [Input("name")]
        public string? Name { get; set; }

        [Input("region")]
        public string? Region { get; set; }

        [Input("status")]
        public string? Status { get; set; }

        [Input("vpcId")]
        public string? VpcId { get; set; }

        public GetCceClusterArgs()
        {
        }
        public static new GetCceClusterArgs Empty => new GetCceClusterArgs();
    }

    public sealed class GetCceClusterInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("clusterType")]
        public Input<string>? ClusterType { get; set; }

        [Input("id")]
        public Input<string>? Id { get; set; }

        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("region")]
        public Input<string>? Region { get; set; }

        [Input("status")]
        public Input<string>? Status { get; set; }

        [Input("vpcId")]
        public Input<string>? VpcId { get; set; }

        public GetCceClusterInvokeArgs()
        {
        }
        public static new GetCceClusterInvokeArgs Empty => new GetCceClusterInvokeArgs();
    }


    [OutputType]
    public sealed class GetCceClusterResult
    {
        public readonly string AuthenticationMode;
        public readonly int BillingMode;
        public readonly ImmutableArray<Outputs.GetCceClusterCertificateClusterResult> CertificateClusters;
        public readonly ImmutableArray<Outputs.GetCceClusterCertificateUserResult> CertificateUsers;
        public readonly string ClusterType;
        public readonly string ClusterVersion;
        public readonly string ContainerNetworkCidr;
        public readonly string ContainerNetworkType;
        public readonly string Description;
        public readonly ImmutableArray<Outputs.GetCceClusterEndpointResult> Endpoints;
        public readonly string EniSubnetCidr;
        public readonly string EniSubnetId;
        public readonly string EnterpriseProjectId;
        public readonly string FlavorId;
        public readonly string HighwaySubnetId;
        public readonly string Id;
        public readonly string KubeConfigRaw;
        public readonly ImmutableArray<Outputs.GetCceClusterMasterResult> Masters;
        public readonly string Name;
        public readonly string Region;
        public readonly string SecurityGroupId;
        public readonly string ServiceNetworkCidr;
        public readonly string Status;
        public readonly string SubnetId;
        public readonly string VpcId;

        [OutputConstructor]
        private GetCceClusterResult(
            string authenticationMode,

            int billingMode,

            ImmutableArray<Outputs.GetCceClusterCertificateClusterResult> certificateClusters,

            ImmutableArray<Outputs.GetCceClusterCertificateUserResult> certificateUsers,

            string clusterType,

            string clusterVersion,

            string containerNetworkCidr,

            string containerNetworkType,

            string description,

            ImmutableArray<Outputs.GetCceClusterEndpointResult> endpoints,

            string eniSubnetCidr,

            string eniSubnetId,

            string enterpriseProjectId,

            string flavorId,

            string highwaySubnetId,

            string id,

            string kubeConfigRaw,

            ImmutableArray<Outputs.GetCceClusterMasterResult> masters,

            string name,

            string region,

            string securityGroupId,

            string serviceNetworkCidr,

            string status,

            string subnetId,

            string vpcId)
        {
            AuthenticationMode = authenticationMode;
            BillingMode = billingMode;
            CertificateClusters = certificateClusters;
            CertificateUsers = certificateUsers;
            ClusterType = clusterType;
            ClusterVersion = clusterVersion;
            ContainerNetworkCidr = containerNetworkCidr;
            ContainerNetworkType = containerNetworkType;
            Description = description;
            Endpoints = endpoints;
            EniSubnetCidr = eniSubnetCidr;
            EniSubnetId = eniSubnetId;
            EnterpriseProjectId = enterpriseProjectId;
            FlavorId = flavorId;
            HighwaySubnetId = highwaySubnetId;
            Id = id;
            KubeConfigRaw = kubeConfigRaw;
            Masters = masters;
            Name = name;
            Region = region;
            SecurityGroupId = securityGroupId;
            ServiceNetworkCidr = serviceNetworkCidr;
            Status = status;
            SubnetId = subnetId;
            VpcId = vpcId;
        }
    }
}
