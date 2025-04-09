// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Sbercloud
{
    public static class GetVpcBandwidth
    {
        public static Task<GetVpcBandwidthResult> InvokeAsync(GetVpcBandwidthArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetVpcBandwidthResult>("sbercloud:index/getVpcBandwidth:getVpcBandwidth", args ?? new GetVpcBandwidthArgs(), options.WithDefaults());

        public static Output<GetVpcBandwidthResult> Invoke(GetVpcBandwidthInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetVpcBandwidthResult>("sbercloud:index/getVpcBandwidth:getVpcBandwidth", args ?? new GetVpcBandwidthInvokeArgs(), options.WithDefaults());

        public static Output<GetVpcBandwidthResult> Invoke(GetVpcBandwidthInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetVpcBandwidthResult>("sbercloud:index/getVpcBandwidth:getVpcBandwidth", args ?? new GetVpcBandwidthInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetVpcBandwidthArgs : global::Pulumi.InvokeArgs
    {
        [Input("enterpriseProjectId")]
        public string? EnterpriseProjectId { get; set; }

        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        [Input("region")]
        public string? Region { get; set; }

        [Input("size")]
        public int? Size { get; set; }

        public GetVpcBandwidthArgs()
        {
        }
        public static new GetVpcBandwidthArgs Empty => new GetVpcBandwidthArgs();
    }

    public sealed class GetVpcBandwidthInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("enterpriseProjectId")]
        public Input<string>? EnterpriseProjectId { get; set; }

        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        [Input("region")]
        public Input<string>? Region { get; set; }

        [Input("size")]
        public Input<int>? Size { get; set; }

        public GetVpcBandwidthInvokeArgs()
        {
        }
        public static new GetVpcBandwidthInvokeArgs Empty => new GetVpcBandwidthInvokeArgs();
    }


    [OutputType]
    public sealed class GetVpcBandwidthResult
    {
        public readonly string BandwidthType;
        public readonly string ChargeMode;
        public readonly string EnterpriseProjectId;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string Name;
        public readonly ImmutableArray<Outputs.GetVpcBandwidthPublicipResult> Publicips;
        public readonly string Region;
        public readonly string ShareType;
        public readonly int Size;
        public readonly string Status;

        [OutputConstructor]
        private GetVpcBandwidthResult(
            string bandwidthType,

            string chargeMode,

            string enterpriseProjectId,

            string id,

            string name,

            ImmutableArray<Outputs.GetVpcBandwidthPublicipResult> publicips,

            string region,

            string shareType,

            int size,

            string status)
        {
            BandwidthType = bandwidthType;
            ChargeMode = chargeMode;
            EnterpriseProjectId = enterpriseProjectId;
            Id = id;
            Name = name;
            Publicips = publicips;
            Region = region;
            ShareType = shareType;
            Size = size;
            Status = status;
        }
    }
}
