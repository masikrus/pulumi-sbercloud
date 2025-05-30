// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Sbercloud
{
    public static class GetApigInstanceFeatures
    {
        public static Task<GetApigInstanceFeaturesResult> InvokeAsync(GetApigInstanceFeaturesArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetApigInstanceFeaturesResult>("sbercloud:index/getApigInstanceFeatures:getApigInstanceFeatures", args ?? new GetApigInstanceFeaturesArgs(), options.WithDefaults());

        public static Output<GetApigInstanceFeaturesResult> Invoke(GetApigInstanceFeaturesInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetApigInstanceFeaturesResult>("sbercloud:index/getApigInstanceFeatures:getApigInstanceFeatures", args ?? new GetApigInstanceFeaturesInvokeArgs(), options.WithDefaults());

        public static Output<GetApigInstanceFeaturesResult> Invoke(GetApigInstanceFeaturesInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetApigInstanceFeaturesResult>("sbercloud:index/getApigInstanceFeatures:getApigInstanceFeatures", args ?? new GetApigInstanceFeaturesInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetApigInstanceFeaturesArgs : global::Pulumi.InvokeArgs
    {
        [Input("instanceId", required: true)]
        public string InstanceId { get; set; } = null!;

        [Input("name")]
        public string? Name { get; set; }

        [Input("region")]
        public string? Region { get; set; }

        public GetApigInstanceFeaturesArgs()
        {
        }
        public static new GetApigInstanceFeaturesArgs Empty => new GetApigInstanceFeaturesArgs();
    }

    public sealed class GetApigInstanceFeaturesInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("instanceId", required: true)]
        public Input<string> InstanceId { get; set; } = null!;

        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("region")]
        public Input<string>? Region { get; set; }

        public GetApigInstanceFeaturesInvokeArgs()
        {
        }
        public static new GetApigInstanceFeaturesInvokeArgs Empty => new GetApigInstanceFeaturesInvokeArgs();
    }


    [OutputType]
    public sealed class GetApigInstanceFeaturesResult
    {
        public readonly ImmutableArray<Outputs.GetApigInstanceFeaturesFeatureResult> Features;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string InstanceId;
        public readonly string? Name;
        public readonly string Region;

        [OutputConstructor]
        private GetApigInstanceFeaturesResult(
            ImmutableArray<Outputs.GetApigInstanceFeaturesFeatureResult> features,

            string id,

            string instanceId,

            string? name,

            string region)
        {
            Features = features;
            Id = id;
            InstanceId = instanceId;
            Name = name;
            Region = region;
        }
    }
}
