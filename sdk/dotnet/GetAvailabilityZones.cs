// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Sbercloud
{
    public static class GetAvailabilityZones
    {
        public static Task<GetAvailabilityZonesResult> InvokeAsync(GetAvailabilityZonesArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetAvailabilityZonesResult>("sbercloud:index/getAvailabilityZones:getAvailabilityZones", args ?? new GetAvailabilityZonesArgs(), options.WithDefaults());

        public static Output<GetAvailabilityZonesResult> Invoke(GetAvailabilityZonesInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetAvailabilityZonesResult>("sbercloud:index/getAvailabilityZones:getAvailabilityZones", args ?? new GetAvailabilityZonesInvokeArgs(), options.WithDefaults());

        public static Output<GetAvailabilityZonesResult> Invoke(GetAvailabilityZonesInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetAvailabilityZonesResult>("sbercloud:index/getAvailabilityZones:getAvailabilityZones", args ?? new GetAvailabilityZonesInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetAvailabilityZonesArgs : global::Pulumi.InvokeArgs
    {
        [Input("region")]
        public string? Region { get; set; }

        [Input("state")]
        public string? State { get; set; }

        public GetAvailabilityZonesArgs()
        {
        }
        public static new GetAvailabilityZonesArgs Empty => new GetAvailabilityZonesArgs();
    }

    public sealed class GetAvailabilityZonesInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("region")]
        public Input<string>? Region { get; set; }

        [Input("state")]
        public Input<string>? State { get; set; }

        public GetAvailabilityZonesInvokeArgs()
        {
        }
        public static new GetAvailabilityZonesInvokeArgs Empty => new GetAvailabilityZonesInvokeArgs();
    }


    [OutputType]
    public sealed class GetAvailabilityZonesResult
    {
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly ImmutableArray<string> Names;
        public readonly string Region;
        public readonly string? State;

        [OutputConstructor]
        private GetAvailabilityZonesResult(
            string id,

            ImmutableArray<string> names,

            string region,

            string? state)
        {
            Id = id;
            Names = names;
            Region = region;
            State = state;
        }
    }
}
