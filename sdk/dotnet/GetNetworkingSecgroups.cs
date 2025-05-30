// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Sbercloud
{
    public static class GetNetworkingSecgroups
    {
        public static Task<GetNetworkingSecgroupsResult> InvokeAsync(GetNetworkingSecgroupsArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetNetworkingSecgroupsResult>("sbercloud:index/getNetworkingSecgroups:getNetworkingSecgroups", args ?? new GetNetworkingSecgroupsArgs(), options.WithDefaults());

        public static Output<GetNetworkingSecgroupsResult> Invoke(GetNetworkingSecgroupsInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetNetworkingSecgroupsResult>("sbercloud:index/getNetworkingSecgroups:getNetworkingSecgroups", args ?? new GetNetworkingSecgroupsInvokeArgs(), options.WithDefaults());

        public static Output<GetNetworkingSecgroupsResult> Invoke(GetNetworkingSecgroupsInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetNetworkingSecgroupsResult>("sbercloud:index/getNetworkingSecgroups:getNetworkingSecgroups", args ?? new GetNetworkingSecgroupsInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetNetworkingSecgroupsArgs : global::Pulumi.InvokeArgs
    {
        [Input("description")]
        public string? Description { get; set; }

        [Input("enterpriseProjectId")]
        public string? EnterpriseProjectId { get; set; }

        [Input("id")]
        public string? Id { get; set; }

        [Input("name")]
        public string? Name { get; set; }

        [Input("region")]
        public string? Region { get; set; }

        public GetNetworkingSecgroupsArgs()
        {
        }
        public static new GetNetworkingSecgroupsArgs Empty => new GetNetworkingSecgroupsArgs();
    }

    public sealed class GetNetworkingSecgroupsInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("enterpriseProjectId")]
        public Input<string>? EnterpriseProjectId { get; set; }

        [Input("id")]
        public Input<string>? Id { get; set; }

        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("region")]
        public Input<string>? Region { get; set; }

        public GetNetworkingSecgroupsInvokeArgs()
        {
        }
        public static new GetNetworkingSecgroupsInvokeArgs Empty => new GetNetworkingSecgroupsInvokeArgs();
    }


    [OutputType]
    public sealed class GetNetworkingSecgroupsResult
    {
        public readonly string? Description;
        public readonly string? EnterpriseProjectId;
        public readonly string Id;
        public readonly string? Name;
        public readonly string Region;
        public readonly ImmutableArray<Outputs.GetNetworkingSecgroupsSecurityGroupResult> SecurityGroups;

        [OutputConstructor]
        private GetNetworkingSecgroupsResult(
            string? description,

            string? enterpriseProjectId,

            string id,

            string? name,

            string region,

            ImmutableArray<Outputs.GetNetworkingSecgroupsSecurityGroupResult> securityGroups)
        {
            Description = description;
            EnterpriseProjectId = enterpriseProjectId;
            Id = id;
            Name = name;
            Region = region;
            SecurityGroups = securityGroups;
        }
    }
}
