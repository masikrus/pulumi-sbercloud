// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Sbercloud
{
    [SbercloudResourceType("sbercloud:index/identityProvider:IdentityProvider")]
    public partial class IdentityProvider : global::Pulumi.CustomResource
    {
        [Output("accessConfig")]
        public Output<Outputs.IdentityProviderAccessConfig?> AccessConfig { get; private set; } = null!;

        [Output("conversionRules")]
        public Output<ImmutableArray<Outputs.IdentityProviderConversionRule>> ConversionRules { get; private set; } = null!;

        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        [Output("loginLink")]
        public Output<string> LoginLink { get; private set; } = null!;

        [Output("metadata")]
        public Output<string?> Metadata { get; private set; } = null!;

        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        [Output("protocol")]
        public Output<string> Protocol { get; private set; } = null!;

        [Output("ssoType")]
        public Output<string> SsoType { get; private set; } = null!;

        [Output("status")]
        public Output<bool?> Status { get; private set; } = null!;


        /// <summary>
        /// Create a IdentityProvider resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public IdentityProvider(string name, IdentityProviderArgs args, CustomResourceOptions? options = null)
            : base("sbercloud:index/identityProvider:IdentityProvider", name, args ?? new IdentityProviderArgs(), MakeResourceOptions(options, ""))
        {
        }

        private IdentityProvider(string name, Input<string> id, IdentityProviderState? state = null, CustomResourceOptions? options = null)
            : base("sbercloud:index/identityProvider:IdentityProvider", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing IdentityProvider resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static IdentityProvider Get(string name, Input<string> id, IdentityProviderState? state = null, CustomResourceOptions? options = null)
        {
            return new IdentityProvider(name, id, state, options);
        }
    }

    public sealed class IdentityProviderArgs : global::Pulumi.ResourceArgs
    {
        [Input("accessConfig")]
        public Input<Inputs.IdentityProviderAccessConfigArgs>? AccessConfig { get; set; }

        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("metadata")]
        public Input<string>? Metadata { get; set; }

        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("protocol", required: true)]
        public Input<string> Protocol { get; set; } = null!;

        [Input("ssoType")]
        public Input<string>? SsoType { get; set; }

        [Input("status")]
        public Input<bool>? Status { get; set; }

        public IdentityProviderArgs()
        {
        }
        public static new IdentityProviderArgs Empty => new IdentityProviderArgs();
    }

    public sealed class IdentityProviderState : global::Pulumi.ResourceArgs
    {
        [Input("accessConfig")]
        public Input<Inputs.IdentityProviderAccessConfigGetArgs>? AccessConfig { get; set; }

        [Input("conversionRules")]
        private InputList<Inputs.IdentityProviderConversionRuleGetArgs>? _conversionRules;
        public InputList<Inputs.IdentityProviderConversionRuleGetArgs> ConversionRules
        {
            get => _conversionRules ?? (_conversionRules = new InputList<Inputs.IdentityProviderConversionRuleGetArgs>());
            set => _conversionRules = value;
        }

        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("loginLink")]
        public Input<string>? LoginLink { get; set; }

        [Input("metadata")]
        public Input<string>? Metadata { get; set; }

        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("protocol")]
        public Input<string>? Protocol { get; set; }

        [Input("ssoType")]
        public Input<string>? SsoType { get; set; }

        [Input("status")]
        public Input<bool>? Status { get; set; }

        public IdentityProviderState()
        {
        }
        public static new IdentityProviderState Empty => new IdentityProviderState();
    }
}
