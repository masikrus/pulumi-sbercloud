// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Sbercloud
{
    [SbercloudResourceType("sbercloud:index/elbL7rule:ElbL7rule")]
    public partial class ElbL7rule : global::Pulumi.CustomResource
    {
        [Output("compareType")]
        public Output<string> CompareType { get; private set; } = null!;

        [Output("conditions")]
        public Output<ImmutableArray<Outputs.ElbL7ruleCondition>> Conditions { get; private set; } = null!;

        [Output("createdAt")]
        public Output<string> CreatedAt { get; private set; } = null!;

        [Output("l7policyId")]
        public Output<string> L7policyId { get; private set; } = null!;

        [Output("region")]
        public Output<string> Region { get; private set; } = null!;

        [Output("type")]
        public Output<string> Type { get; private set; } = null!;

        [Output("updatedAt")]
        public Output<string> UpdatedAt { get; private set; } = null!;

        [Output("value")]
        public Output<string> Value { get; private set; } = null!;


        /// <summary>
        /// Create a ElbL7rule resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ElbL7rule(string name, ElbL7ruleArgs args, CustomResourceOptions? options = null)
            : base("sbercloud:index/elbL7rule:ElbL7rule", name, args ?? new ElbL7ruleArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ElbL7rule(string name, Input<string> id, ElbL7ruleState? state = null, CustomResourceOptions? options = null)
            : base("sbercloud:index/elbL7rule:ElbL7rule", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing ElbL7rule resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ElbL7rule Get(string name, Input<string> id, ElbL7ruleState? state = null, CustomResourceOptions? options = null)
        {
            return new ElbL7rule(name, id, state, options);
        }
    }

    public sealed class ElbL7ruleArgs : global::Pulumi.ResourceArgs
    {
        [Input("compareType", required: true)]
        public Input<string> CompareType { get; set; } = null!;

        [Input("conditions")]
        private InputList<Inputs.ElbL7ruleConditionArgs>? _conditions;
        public InputList<Inputs.ElbL7ruleConditionArgs> Conditions
        {
            get => _conditions ?? (_conditions = new InputList<Inputs.ElbL7ruleConditionArgs>());
            set => _conditions = value;
        }

        [Input("l7policyId", required: true)]
        public Input<string> L7policyId { get; set; } = null!;

        [Input("region")]
        public Input<string>? Region { get; set; }

        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        [Input("value")]
        public Input<string>? Value { get; set; }

        public ElbL7ruleArgs()
        {
        }
        public static new ElbL7ruleArgs Empty => new ElbL7ruleArgs();
    }

    public sealed class ElbL7ruleState : global::Pulumi.ResourceArgs
    {
        [Input("compareType")]
        public Input<string>? CompareType { get; set; }

        [Input("conditions")]
        private InputList<Inputs.ElbL7ruleConditionGetArgs>? _conditions;
        public InputList<Inputs.ElbL7ruleConditionGetArgs> Conditions
        {
            get => _conditions ?? (_conditions = new InputList<Inputs.ElbL7ruleConditionGetArgs>());
            set => _conditions = value;
        }

        [Input("createdAt")]
        public Input<string>? CreatedAt { get; set; }

        [Input("l7policyId")]
        public Input<string>? L7policyId { get; set; }

        [Input("region")]
        public Input<string>? Region { get; set; }

        [Input("type")]
        public Input<string>? Type { get; set; }

        [Input("updatedAt")]
        public Input<string>? UpdatedAt { get; set; }

        [Input("value")]
        public Input<string>? Value { get; set; }

        public ElbL7ruleState()
        {
        }
        public static new ElbL7ruleState Empty => new ElbL7ruleState();
    }
}
