// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Sbercloud
{
    [SbercloudResourceType("sbercloud:index/apigEndpointWhitelist:ApigEndpointWhitelist")]
    public partial class ApigEndpointWhitelist : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The ID of the dedicated instance to which the endpoint service belongs.
        /// </summary>
        [Output("instanceId")]
        public Output<string> InstanceId { get; private set; } = null!;

        /// <summary>
        /// The region where the endpoint service is located.
        /// </summary>
        [Output("region")]
        public Output<string> Region { get; private set; } = null!;

        /// <summary>
        /// The whitelist records of the endpoint service.
        /// </summary>
        [Output("whitelists")]
        public Output<ImmutableArray<string>> Whitelists { get; private set; } = null!;


        /// <summary>
        /// Create a ApigEndpointWhitelist resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ApigEndpointWhitelist(string name, ApigEndpointWhitelistArgs args, CustomResourceOptions? options = null)
            : base("sbercloud:index/apigEndpointWhitelist:ApigEndpointWhitelist", name, args ?? new ApigEndpointWhitelistArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ApigEndpointWhitelist(string name, Input<string> id, ApigEndpointWhitelistState? state = null, CustomResourceOptions? options = null)
            : base("sbercloud:index/apigEndpointWhitelist:ApigEndpointWhitelist", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing ApigEndpointWhitelist resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ApigEndpointWhitelist Get(string name, Input<string> id, ApigEndpointWhitelistState? state = null, CustomResourceOptions? options = null)
        {
            return new ApigEndpointWhitelist(name, id, state, options);
        }
    }

    public sealed class ApigEndpointWhitelistArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ID of the dedicated instance to which the endpoint service belongs.
        /// </summary>
        [Input("instanceId", required: true)]
        public Input<string> InstanceId { get; set; } = null!;

        /// <summary>
        /// The region where the endpoint service is located.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        [Input("whitelists", required: true)]
        private InputList<string>? _whitelists;

        /// <summary>
        /// The whitelist records of the endpoint service.
        /// </summary>
        public InputList<string> Whitelists
        {
            get => _whitelists ?? (_whitelists = new InputList<string>());
            set => _whitelists = value;
        }

        public ApigEndpointWhitelistArgs()
        {
        }
        public static new ApigEndpointWhitelistArgs Empty => new ApigEndpointWhitelistArgs();
    }

    public sealed class ApigEndpointWhitelistState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ID of the dedicated instance to which the endpoint service belongs.
        /// </summary>
        [Input("instanceId")]
        public Input<string>? InstanceId { get; set; }

        /// <summary>
        /// The region where the endpoint service is located.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        [Input("whitelists")]
        private InputList<string>? _whitelists;

        /// <summary>
        /// The whitelist records of the endpoint service.
        /// </summary>
        public InputList<string> Whitelists
        {
            get => _whitelists ?? (_whitelists = new InputList<string>());
            set => _whitelists = value;
        }

        public ApigEndpointWhitelistState()
        {
        }
        public static new ApigEndpointWhitelistState Empty => new ApigEndpointWhitelistState();
    }
}
