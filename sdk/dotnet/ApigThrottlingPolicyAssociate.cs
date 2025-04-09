// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Sbercloud
{
    [SbercloudResourceType("sbercloud:index/apigThrottlingPolicyAssociate:ApigThrottlingPolicyAssociate")]
    public partial class ApigThrottlingPolicyAssociate : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The ID of the dedicated instance to which the APIs and the throttling policy belongs.
        /// </summary>
        [Output("instanceId")]
        public Output<string> InstanceId { get; private set; } = null!;

        /// <summary>
        /// The ID of the throttling policy.
        /// </summary>
        [Output("policyId")]
        public Output<string> PolicyId { get; private set; } = null!;

        /// <summary>
        /// The publish IDs corresponding to the APIs bound by the throttling policy.
        /// </summary>
        [Output("publishIds")]
        public Output<ImmutableArray<string>> PublishIds { get; private set; } = null!;

        /// <summary>
        /// The region where the dedicated instance and the throttling policy are located.
        /// </summary>
        [Output("region")]
        public Output<string> Region { get; private set; } = null!;


        /// <summary>
        /// Create a ApigThrottlingPolicyAssociate resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ApigThrottlingPolicyAssociate(string name, ApigThrottlingPolicyAssociateArgs args, CustomResourceOptions? options = null)
            : base("sbercloud:index/apigThrottlingPolicyAssociate:ApigThrottlingPolicyAssociate", name, args ?? new ApigThrottlingPolicyAssociateArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ApigThrottlingPolicyAssociate(string name, Input<string> id, ApigThrottlingPolicyAssociateState? state = null, CustomResourceOptions? options = null)
            : base("sbercloud:index/apigThrottlingPolicyAssociate:ApigThrottlingPolicyAssociate", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing ApigThrottlingPolicyAssociate resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ApigThrottlingPolicyAssociate Get(string name, Input<string> id, ApigThrottlingPolicyAssociateState? state = null, CustomResourceOptions? options = null)
        {
            return new ApigThrottlingPolicyAssociate(name, id, state, options);
        }
    }

    public sealed class ApigThrottlingPolicyAssociateArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ID of the dedicated instance to which the APIs and the throttling policy belongs.
        /// </summary>
        [Input("instanceId", required: true)]
        public Input<string> InstanceId { get; set; } = null!;

        /// <summary>
        /// The ID of the throttling policy.
        /// </summary>
        [Input("policyId", required: true)]
        public Input<string> PolicyId { get; set; } = null!;

        [Input("publishIds", required: true)]
        private InputList<string>? _publishIds;

        /// <summary>
        /// The publish IDs corresponding to the APIs bound by the throttling policy.
        /// </summary>
        public InputList<string> PublishIds
        {
            get => _publishIds ?? (_publishIds = new InputList<string>());
            set => _publishIds = value;
        }

        /// <summary>
        /// The region where the dedicated instance and the throttling policy are located.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        public ApigThrottlingPolicyAssociateArgs()
        {
        }
        public static new ApigThrottlingPolicyAssociateArgs Empty => new ApigThrottlingPolicyAssociateArgs();
    }

    public sealed class ApigThrottlingPolicyAssociateState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ID of the dedicated instance to which the APIs and the throttling policy belongs.
        /// </summary>
        [Input("instanceId")]
        public Input<string>? InstanceId { get; set; }

        /// <summary>
        /// The ID of the throttling policy.
        /// </summary>
        [Input("policyId")]
        public Input<string>? PolicyId { get; set; }

        [Input("publishIds")]
        private InputList<string>? _publishIds;

        /// <summary>
        /// The publish IDs corresponding to the APIs bound by the throttling policy.
        /// </summary>
        public InputList<string> PublishIds
        {
            get => _publishIds ?? (_publishIds = new InputList<string>());
            set => _publishIds = value;
        }

        /// <summary>
        /// The region where the dedicated instance and the throttling policy are located.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        public ApigThrottlingPolicyAssociateState()
        {
        }
        public static new ApigThrottlingPolicyAssociateState Empty => new ApigThrottlingPolicyAssociateState();
    }
}
