// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Sbercloud
{
    [SbercloudResourceType("sbercloud:index/apigApplicationQuotaAssociate:ApigApplicationQuotaAssociate")]
    public partial class ApigApplicationQuotaAssociate : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The configuration of applications bound to the quota.
        /// </summary>
        [Output("applications")]
        public Output<ImmutableArray<Outputs.ApigApplicationQuotaAssociateApplication>> Applications { get; private set; } = null!;

        /// <summary>
        /// The ID of the dedicated instance to which the application quota (policy) belongs.
        /// </summary>
        [Output("instanceId")]
        public Output<string> InstanceId { get; private set; } = null!;

        /// <summary>
        /// The ID of the application quota (policy).
        /// </summary>
        [Output("quotaId")]
        public Output<string> QuotaId { get; private set; } = null!;

        /// <summary>
        /// The region where the application quota (policy) is located.
        /// </summary>
        [Output("region")]
        public Output<string> Region { get; private set; } = null!;


        /// <summary>
        /// Create a ApigApplicationQuotaAssociate resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ApigApplicationQuotaAssociate(string name, ApigApplicationQuotaAssociateArgs args, CustomResourceOptions? options = null)
            : base("sbercloud:index/apigApplicationQuotaAssociate:ApigApplicationQuotaAssociate", name, args ?? new ApigApplicationQuotaAssociateArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ApigApplicationQuotaAssociate(string name, Input<string> id, ApigApplicationQuotaAssociateState? state = null, CustomResourceOptions? options = null)
            : base("sbercloud:index/apigApplicationQuotaAssociate:ApigApplicationQuotaAssociate", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing ApigApplicationQuotaAssociate resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ApigApplicationQuotaAssociate Get(string name, Input<string> id, ApigApplicationQuotaAssociateState? state = null, CustomResourceOptions? options = null)
        {
            return new ApigApplicationQuotaAssociate(name, id, state, options);
        }
    }

    public sealed class ApigApplicationQuotaAssociateArgs : global::Pulumi.ResourceArgs
    {
        [Input("applications", required: true)]
        private InputList<Inputs.ApigApplicationQuotaAssociateApplicationArgs>? _applications;

        /// <summary>
        /// The configuration of applications bound to the quota.
        /// </summary>
        public InputList<Inputs.ApigApplicationQuotaAssociateApplicationArgs> Applications
        {
            get => _applications ?? (_applications = new InputList<Inputs.ApigApplicationQuotaAssociateApplicationArgs>());
            set => _applications = value;
        }

        /// <summary>
        /// The ID of the dedicated instance to which the application quota (policy) belongs.
        /// </summary>
        [Input("instanceId", required: true)]
        public Input<string> InstanceId { get; set; } = null!;

        /// <summary>
        /// The ID of the application quota (policy).
        /// </summary>
        [Input("quotaId", required: true)]
        public Input<string> QuotaId { get; set; } = null!;

        /// <summary>
        /// The region where the application quota (policy) is located.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        public ApigApplicationQuotaAssociateArgs()
        {
        }
        public static new ApigApplicationQuotaAssociateArgs Empty => new ApigApplicationQuotaAssociateArgs();
    }

    public sealed class ApigApplicationQuotaAssociateState : global::Pulumi.ResourceArgs
    {
        [Input("applications")]
        private InputList<Inputs.ApigApplicationQuotaAssociateApplicationGetArgs>? _applications;

        /// <summary>
        /// The configuration of applications bound to the quota.
        /// </summary>
        public InputList<Inputs.ApigApplicationQuotaAssociateApplicationGetArgs> Applications
        {
            get => _applications ?? (_applications = new InputList<Inputs.ApigApplicationQuotaAssociateApplicationGetArgs>());
            set => _applications = value;
        }

        /// <summary>
        /// The ID of the dedicated instance to which the application quota (policy) belongs.
        /// </summary>
        [Input("instanceId")]
        public Input<string>? InstanceId { get; set; }

        /// <summary>
        /// The ID of the application quota (policy).
        /// </summary>
        [Input("quotaId")]
        public Input<string>? QuotaId { get; set; }

        /// <summary>
        /// The region where the application quota (policy) is located.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        public ApigApplicationQuotaAssociateState()
        {
        }
        public static new ApigApplicationQuotaAssociateState Empty => new ApigApplicationQuotaAssociateState();
    }
}
