// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Sbercloud
{
    public static class GetRdsInstances
    {
        public static Task<GetRdsInstancesResult> InvokeAsync(GetRdsInstancesArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetRdsInstancesResult>("sbercloud:index/getRdsInstances:getRdsInstances", args ?? new GetRdsInstancesArgs(), options.WithDefaults());

        public static Output<GetRdsInstancesResult> Invoke(GetRdsInstancesInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetRdsInstancesResult>("sbercloud:index/getRdsInstances:getRdsInstances", args ?? new GetRdsInstancesInvokeArgs(), options.WithDefaults());

        public static Output<GetRdsInstancesResult> Invoke(GetRdsInstancesInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetRdsInstancesResult>("sbercloud:index/getRdsInstances:getRdsInstances", args ?? new GetRdsInstancesInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetRdsInstancesArgs : global::Pulumi.InvokeArgs
    {
        [Input("datastoreType")]
        public string? DatastoreType { get; set; }

        [Input("enterpriseProjectId")]
        public string? EnterpriseProjectId { get; set; }

        [Input("name")]
        public string? Name { get; set; }

        [Input("region")]
        public string? Region { get; set; }

        [Input("subnetId")]
        public string? SubnetId { get; set; }

        [Input("type")]
        public string? Type { get; set; }

        [Input("vpcId")]
        public string? VpcId { get; set; }

        public GetRdsInstancesArgs()
        {
        }
        public static new GetRdsInstancesArgs Empty => new GetRdsInstancesArgs();
    }

    public sealed class GetRdsInstancesInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("datastoreType")]
        public Input<string>? DatastoreType { get; set; }

        [Input("enterpriseProjectId")]
        public Input<string>? EnterpriseProjectId { get; set; }

        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("region")]
        public Input<string>? Region { get; set; }

        [Input("subnetId")]
        public Input<string>? SubnetId { get; set; }

        [Input("type")]
        public Input<string>? Type { get; set; }

        [Input("vpcId")]
        public Input<string>? VpcId { get; set; }

        public GetRdsInstancesInvokeArgs()
        {
        }
        public static new GetRdsInstancesInvokeArgs Empty => new GetRdsInstancesInvokeArgs();
    }


    [OutputType]
    public sealed class GetRdsInstancesResult
    {
        public readonly string? DatastoreType;
        public readonly string? EnterpriseProjectId;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly ImmutableArray<Outputs.GetRdsInstancesInstanceResult> Instances;
        public readonly string? Name;
        public readonly string Region;
        public readonly string? SubnetId;
        public readonly string? Type;
        public readonly string? VpcId;

        [OutputConstructor]
        private GetRdsInstancesResult(
            string? datastoreType,

            string? enterpriseProjectId,

            string id,

            ImmutableArray<Outputs.GetRdsInstancesInstanceResult> instances,

            string? name,

            string region,

            string? subnetId,

            string? type,

            string? vpcId)
        {
            DatastoreType = datastoreType;
            EnterpriseProjectId = enterpriseProjectId;
            Id = id;
            Instances = instances;
            Name = name;
            Region = region;
            SubnetId = subnetId;
            Type = type;
            VpcId = vpcId;
        }
    }
}
