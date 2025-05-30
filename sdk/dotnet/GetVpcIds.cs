// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Sbercloud
{
    public static class GetVpcIds
    {
        public static Task<GetVpcIdsResult> InvokeAsync(GetVpcIdsArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetVpcIdsResult>("sbercloud:index/getVpcIds:getVpcIds", args ?? new GetVpcIdsArgs(), options.WithDefaults());

        public static Output<GetVpcIdsResult> Invoke(GetVpcIdsInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetVpcIdsResult>("sbercloud:index/getVpcIds:getVpcIds", args ?? new GetVpcIdsInvokeArgs(), options.WithDefaults());

        public static Output<GetVpcIdsResult> Invoke(GetVpcIdsInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetVpcIdsResult>("sbercloud:index/getVpcIds:getVpcIds", args ?? new GetVpcIdsInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetVpcIdsArgs : global::Pulumi.InvokeArgs
    {
        [Input("region")]
        public string? Region { get; set; }

        public GetVpcIdsArgs()
        {
        }
        public static new GetVpcIdsArgs Empty => new GetVpcIdsArgs();
    }

    public sealed class GetVpcIdsInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("region")]
        public Input<string>? Region { get; set; }

        public GetVpcIdsInvokeArgs()
        {
        }
        public static new GetVpcIdsInvokeArgs Empty => new GetVpcIdsInvokeArgs();
    }


    [OutputType]
    public sealed class GetVpcIdsResult
    {
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly ImmutableArray<string> Ids;
        public readonly string Region;

        [OutputConstructor]
        private GetVpcIdsResult(
            string id,

            ImmutableArray<string> ids,

            string region)
        {
            Id = id;
            Ids = ids;
            Region = region;
        }
    }
}
