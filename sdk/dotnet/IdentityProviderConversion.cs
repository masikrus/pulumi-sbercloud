// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Sbercloud
{
    [SbercloudResourceType("sbercloud:index/identityProviderConversion:IdentityProviderConversion")]
    public partial class IdentityProviderConversion : global::Pulumi.CustomResource
    {
        [Output("conversionRules")]
        public Output<ImmutableArray<Outputs.IdentityProviderConversionConversionRule>> ConversionRules { get; private set; } = null!;

        [Output("providerId")]
        public Output<string> ProviderId { get; private set; } = null!;


        /// <summary>
        /// Create a IdentityProviderConversion resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public IdentityProviderConversion(string name, IdentityProviderConversionArgs args, CustomResourceOptions? options = null)
            : base("sbercloud:index/identityProviderConversion:IdentityProviderConversion", name, args ?? new IdentityProviderConversionArgs(), MakeResourceOptions(options, ""))
        {
        }

        private IdentityProviderConversion(string name, Input<string> id, IdentityProviderConversionState? state = null, CustomResourceOptions? options = null)
            : base("sbercloud:index/identityProviderConversion:IdentityProviderConversion", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing IdentityProviderConversion resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static IdentityProviderConversion Get(string name, Input<string> id, IdentityProviderConversionState? state = null, CustomResourceOptions? options = null)
        {
            return new IdentityProviderConversion(name, id, state, options);
        }
    }

    public sealed class IdentityProviderConversionArgs : global::Pulumi.ResourceArgs
    {
        [Input("conversionRules", required: true)]
        private InputList<Inputs.IdentityProviderConversionConversionRuleArgs>? _conversionRules;
        public InputList<Inputs.IdentityProviderConversionConversionRuleArgs> ConversionRules
        {
            get => _conversionRules ?? (_conversionRules = new InputList<Inputs.IdentityProviderConversionConversionRuleArgs>());
            set => _conversionRules = value;
        }

        [Input("providerId", required: true)]
        public Input<string> ProviderId { get; set; } = null!;

        public IdentityProviderConversionArgs()
        {
        }
        public static new IdentityProviderConversionArgs Empty => new IdentityProviderConversionArgs();
    }

    public sealed class IdentityProviderConversionState : global::Pulumi.ResourceArgs
    {
        [Input("conversionRules")]
        private InputList<Inputs.IdentityProviderConversionConversionRuleGetArgs>? _conversionRules;
        public InputList<Inputs.IdentityProviderConversionConversionRuleGetArgs> ConversionRules
        {
            get => _conversionRules ?? (_conversionRules = new InputList<Inputs.IdentityProviderConversionConversionRuleGetArgs>());
            set => _conversionRules = value;
        }

        [Input("providerId")]
        public Input<string>? ProviderId { get; set; }

        public IdentityProviderConversionState()
        {
        }
        public static new IdentityProviderConversionState Empty => new IdentityProviderConversionState();
    }
}
