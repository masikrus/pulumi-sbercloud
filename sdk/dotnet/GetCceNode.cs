// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Sbercloud
{
    public static class GetCceNode
    {
        public static Task<GetCceNodeResult> InvokeAsync(GetCceNodeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetCceNodeResult>("sbercloud:index/getCceNode:getCceNode", args ?? new GetCceNodeArgs(), options.WithDefaults());

        public static Output<GetCceNodeResult> Invoke(GetCceNodeInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetCceNodeResult>("sbercloud:index/getCceNode:getCceNode", args ?? new GetCceNodeInvokeArgs(), options.WithDefaults());

        public static Output<GetCceNodeResult> Invoke(GetCceNodeInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetCceNodeResult>("sbercloud:index/getCceNode:getCceNode", args ?? new GetCceNodeInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetCceNodeArgs : global::Pulumi.InvokeArgs
    {
        [Input("clusterId", required: true)]
        public string ClusterId { get; set; } = null!;

        [Input("name")]
        public string? Name { get; set; }

        [Input("nodeId")]
        public string? NodeId { get; set; }

        [Input("region")]
        public string? Region { get; set; }

        [Input("status")]
        public string? Status { get; set; }

        public GetCceNodeArgs()
        {
        }
        public static new GetCceNodeArgs Empty => new GetCceNodeArgs();
    }

    public sealed class GetCceNodeInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("clusterId", required: true)]
        public Input<string> ClusterId { get; set; } = null!;

        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("nodeId")]
        public Input<string>? NodeId { get; set; }

        [Input("region")]
        public Input<string>? Region { get; set; }

        [Input("status")]
        public Input<string>? Status { get; set; }

        public GetCceNodeInvokeArgs()
        {
        }
        public static new GetCceNodeInvokeArgs Empty => new GetCceNodeInvokeArgs();
    }


    [OutputType]
    public sealed class GetCceNodeResult
    {
        public readonly string AvailabilityZone;
        public readonly int BillingMode;
        public readonly string ClusterId;
        public readonly ImmutableArray<Outputs.GetCceNodeDataVolumeResult> DataVolumes;
        public readonly string EcsGroupId;
        public readonly string EnterpriseProjectId;
        public readonly string FlavorId;
        public readonly ImmutableArray<Outputs.GetCceNodeHostnameConfigResult> HostnameConfigs;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string KeyPair;
        public readonly string Name;
        public readonly string NodeId;
        public readonly string Os;
        public readonly string PrivateIp;
        public readonly string PublicIp;
        public readonly string Region;
        public readonly ImmutableArray<Outputs.GetCceNodeRootVolumeResult> RootVolumes;
        public readonly string ServerId;
        public readonly string Status;
        public readonly string SubnetId;
        public readonly ImmutableDictionary<string, string> Tags;

        [OutputConstructor]
        private GetCceNodeResult(
            string availabilityZone,

            int billingMode,

            string clusterId,

            ImmutableArray<Outputs.GetCceNodeDataVolumeResult> dataVolumes,

            string ecsGroupId,

            string enterpriseProjectId,

            string flavorId,

            ImmutableArray<Outputs.GetCceNodeHostnameConfigResult> hostnameConfigs,

            string id,

            string keyPair,

            string name,

            string nodeId,

            string os,

            string privateIp,

            string publicIp,

            string region,

            ImmutableArray<Outputs.GetCceNodeRootVolumeResult> rootVolumes,

            string serverId,

            string status,

            string subnetId,

            ImmutableDictionary<string, string> tags)
        {
            AvailabilityZone = availabilityZone;
            BillingMode = billingMode;
            ClusterId = clusterId;
            DataVolumes = dataVolumes;
            EcsGroupId = ecsGroupId;
            EnterpriseProjectId = enterpriseProjectId;
            FlavorId = flavorId;
            HostnameConfigs = hostnameConfigs;
            Id = id;
            KeyPair = keyPair;
            Name = name;
            NodeId = nodeId;
            Os = os;
            PrivateIp = privateIp;
            PublicIp = publicIp;
            Region = region;
            RootVolumes = rootVolumes;
            ServerId = serverId;
            Status = status;
            SubnetId = subnetId;
            Tags = tags;
        }
    }
}
