// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Sbercloud
{
    public static class GetCbhFlavors
    {
        public static Task<GetCbhFlavorsResult> InvokeAsync(GetCbhFlavorsArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetCbhFlavorsResult>("sbercloud:index/getCbhFlavors:getCbhFlavors", args ?? new GetCbhFlavorsArgs(), options.WithDefaults());

        public static Output<GetCbhFlavorsResult> Invoke(GetCbhFlavorsInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetCbhFlavorsResult>("sbercloud:index/getCbhFlavors:getCbhFlavors", args ?? new GetCbhFlavorsInvokeArgs(), options.WithDefaults());

        public static Output<GetCbhFlavorsResult> Invoke(GetCbhFlavorsInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetCbhFlavorsResult>("sbercloud:index/getCbhFlavors:getCbhFlavors", args ?? new GetCbhFlavorsInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetCbhFlavorsArgs : global::Pulumi.InvokeArgs
    {
        [Input("action")]
        public string? Action { get; set; }

        [Input("asset")]
        public int? Asset { get; set; }

        [Input("flavorId")]
        public string? FlavorId { get; set; }

        [Input("maxConnection")]
        public int? MaxConnection { get; set; }

        [Input("memory")]
        public int? Memory { get; set; }

        [Input("region")]
        public string? Region { get; set; }

        [Input("specCode")]
        public string? SpecCode { get; set; }

        [Input("type")]
        public string? Type { get; set; }

        [Input("vcpus")]
        public int? Vcpus { get; set; }

        public GetCbhFlavorsArgs()
        {
        }
        public static new GetCbhFlavorsArgs Empty => new GetCbhFlavorsArgs();
    }

    public sealed class GetCbhFlavorsInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("action")]
        public Input<string>? Action { get; set; }

        [Input("asset")]
        public Input<int>? Asset { get; set; }

        [Input("flavorId")]
        public Input<string>? FlavorId { get; set; }

        [Input("maxConnection")]
        public Input<int>? MaxConnection { get; set; }

        [Input("memory")]
        public Input<int>? Memory { get; set; }

        [Input("region")]
        public Input<string>? Region { get; set; }

        [Input("specCode")]
        public Input<string>? SpecCode { get; set; }

        [Input("type")]
        public Input<string>? Type { get; set; }

        [Input("vcpus")]
        public Input<int>? Vcpus { get; set; }

        public GetCbhFlavorsInvokeArgs()
        {
        }
        public static new GetCbhFlavorsInvokeArgs Empty => new GetCbhFlavorsInvokeArgs();
    }


    [OutputType]
    public sealed class GetCbhFlavorsResult
    {
        public readonly string? Action;
        public readonly int? Asset;
        public readonly string? FlavorId;
        public readonly ImmutableArray<Outputs.GetCbhFlavorsFlavorResult> Flavors;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly int? MaxConnection;
        public readonly int? Memory;
        public readonly string Region;
        public readonly string? SpecCode;
        public readonly string? Type;
        public readonly int? Vcpus;

        [OutputConstructor]
        private GetCbhFlavorsResult(
            string? action,

            int? asset,

            string? flavorId,

            ImmutableArray<Outputs.GetCbhFlavorsFlavorResult> flavors,

            string id,

            int? maxConnection,

            int? memory,

            string region,

            string? specCode,

            string? type,

            int? vcpus)
        {
            Action = action;
            Asset = asset;
            FlavorId = flavorId;
            Flavors = flavors;
            Id = id;
            MaxConnection = maxConnection;
            Memory = memory;
            Region = region;
            SpecCode = specCode;
            Type = type;
            Vcpus = vcpus;
        }
    }
}
