// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Sbercloud
{
    public static class GetVpcSubnetIds
    {
        public static Task<GetVpcSubnetIdsResult> InvokeAsync(GetVpcSubnetIdsArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetVpcSubnetIdsResult>("sbercloud:index/getVpcSubnetIds:getVpcSubnetIds", args ?? new GetVpcSubnetIdsArgs(), options.WithDefaults());

        public static Output<GetVpcSubnetIdsResult> Invoke(GetVpcSubnetIdsInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetVpcSubnetIdsResult>("sbercloud:index/getVpcSubnetIds:getVpcSubnetIds", args ?? new GetVpcSubnetIdsInvokeArgs(), options.WithDefaults());

        public static Output<GetVpcSubnetIdsResult> Invoke(GetVpcSubnetIdsInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetVpcSubnetIdsResult>("sbercloud:index/getVpcSubnetIds:getVpcSubnetIds", args ?? new GetVpcSubnetIdsInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetVpcSubnetIdsArgs : global::Pulumi.InvokeArgs
    {
        [Input("region")]
        public string? Region { get; set; }

        [Input("vpcId", required: true)]
        public string VpcId { get; set; } = null!;

        public GetVpcSubnetIdsArgs()
        {
        }
        public static new GetVpcSubnetIdsArgs Empty => new GetVpcSubnetIdsArgs();
    }

    public sealed class GetVpcSubnetIdsInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("region")]
        public Input<string>? Region { get; set; }

        [Input("vpcId", required: true)]
        public Input<string> VpcId { get; set; } = null!;

        public GetVpcSubnetIdsInvokeArgs()
        {
        }
        public static new GetVpcSubnetIdsInvokeArgs Empty => new GetVpcSubnetIdsInvokeArgs();
    }


    [OutputType]
    public sealed class GetVpcSubnetIdsResult
    {
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly ImmutableArray<string> Ids;
        public readonly string Region;
        public readonly string VpcId;

        [OutputConstructor]
        private GetVpcSubnetIdsResult(
            string id,

            ImmutableArray<string> ids,

            string region,

            string vpcId)
        {
            Id = id;
            Ids = ids;
            Region = region;
            VpcId = vpcId;
        }
    }
}
