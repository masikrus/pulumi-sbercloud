// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Sbercloud
{
    [SbercloudResourceType("sbercloud:index/sfsTurboPermRule:SfsTurboPermRule")]
    public partial class SfsTurboPermRule : global::Pulumi.CustomResource
    {
        [Output("ipCidr")]
        public Output<string> IpCidr { get; private set; } = null!;

        [Output("region")]
        public Output<string> Region { get; private set; } = null!;

        [Output("rwType")]
        public Output<string> RwType { get; private set; } = null!;

        [Output("shareId")]
        public Output<string> ShareId { get; private set; } = null!;

        [Output("userType")]
        public Output<string> UserType { get; private set; } = null!;


        /// <summary>
        /// Create a SfsTurboPermRule resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public SfsTurboPermRule(string name, SfsTurboPermRuleArgs args, CustomResourceOptions? options = null)
            : base("sbercloud:index/sfsTurboPermRule:SfsTurboPermRule", name, args ?? new SfsTurboPermRuleArgs(), MakeResourceOptions(options, ""))
        {
        }

        private SfsTurboPermRule(string name, Input<string> id, SfsTurboPermRuleState? state = null, CustomResourceOptions? options = null)
            : base("sbercloud:index/sfsTurboPermRule:SfsTurboPermRule", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing SfsTurboPermRule resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static SfsTurboPermRule Get(string name, Input<string> id, SfsTurboPermRuleState? state = null, CustomResourceOptions? options = null)
        {
            return new SfsTurboPermRule(name, id, state, options);
        }
    }

    public sealed class SfsTurboPermRuleArgs : global::Pulumi.ResourceArgs
    {
        [Input("ipCidr", required: true)]
        public Input<string> IpCidr { get; set; } = null!;

        [Input("region")]
        public Input<string>? Region { get; set; }

        [Input("rwType", required: true)]
        public Input<string> RwType { get; set; } = null!;

        [Input("shareId", required: true)]
        public Input<string> ShareId { get; set; } = null!;

        [Input("userType", required: true)]
        public Input<string> UserType { get; set; } = null!;

        public SfsTurboPermRuleArgs()
        {
        }
        public static new SfsTurboPermRuleArgs Empty => new SfsTurboPermRuleArgs();
    }

    public sealed class SfsTurboPermRuleState : global::Pulumi.ResourceArgs
    {
        [Input("ipCidr")]
        public Input<string>? IpCidr { get; set; }

        [Input("region")]
        public Input<string>? Region { get; set; }

        [Input("rwType")]
        public Input<string>? RwType { get; set; }

        [Input("shareId")]
        public Input<string>? ShareId { get; set; }

        [Input("userType")]
        public Input<string>? UserType { get; set; }

        public SfsTurboPermRuleState()
        {
        }
        public static new SfsTurboPermRuleState Empty => new SfsTurboPermRuleState();
    }
}
