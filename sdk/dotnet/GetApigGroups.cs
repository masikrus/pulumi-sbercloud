// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Sbercloud
{
    public static class GetApigGroups
    {
        public static Task<GetApigGroupsResult> InvokeAsync(GetApigGroupsArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetApigGroupsResult>("sbercloud:index/getApigGroups:getApigGroups", args ?? new GetApigGroupsArgs(), options.WithDefaults());

        public static Output<GetApigGroupsResult> Invoke(GetApigGroupsInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetApigGroupsResult>("sbercloud:index/getApigGroups:getApigGroups", args ?? new GetApigGroupsInvokeArgs(), options.WithDefaults());

        public static Output<GetApigGroupsResult> Invoke(GetApigGroupsInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetApigGroupsResult>("sbercloud:index/getApigGroups:getApigGroups", args ?? new GetApigGroupsInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetApigGroupsArgs : global::Pulumi.InvokeArgs
    {
        [Input("groupId")]
        public string? GroupId { get; set; }

        [Input("instanceId", required: true)]
        public string InstanceId { get; set; } = null!;

        [Input("name")]
        public string? Name { get; set; }

        [Input("region")]
        public string? Region { get; set; }

        public GetApigGroupsArgs()
        {
        }
        public static new GetApigGroupsArgs Empty => new GetApigGroupsArgs();
    }

    public sealed class GetApigGroupsInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("groupId")]
        public Input<string>? GroupId { get; set; }

        [Input("instanceId", required: true)]
        public Input<string> InstanceId { get; set; } = null!;

        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("region")]
        public Input<string>? Region { get; set; }

        public GetApigGroupsInvokeArgs()
        {
        }
        public static new GetApigGroupsInvokeArgs Empty => new GetApigGroupsInvokeArgs();
    }


    [OutputType]
    public sealed class GetApigGroupsResult
    {
        public readonly string? GroupId;
        public readonly ImmutableArray<Outputs.GetApigGroupsGroupResult> Groups;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string InstanceId;
        public readonly string? Name;
        public readonly string Region;

        [OutputConstructor]
        private GetApigGroupsResult(
            string? groupId,

            ImmutableArray<Outputs.GetApigGroupsGroupResult> groups,

            string id,

            string instanceId,

            string? name,

            string region)
        {
            GroupId = groupId;
            Groups = groups;
            Id = id;
            InstanceId = instanceId;
            Name = name;
            Region = region;
        }
    }
}
