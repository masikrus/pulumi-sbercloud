// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Sbercloud
{
    [SbercloudResourceType("sbercloud:index/apigApplicationQuota:ApigApplicationQuota")]
    public partial class ApigApplicationQuota : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The number of bound APPs.
        /// </summary>
        [Output("bindNum")]
        public Output<int> BindNum { get; private set; } = null!;

        /// <summary>
        /// Specifies the access limit of the application quota.
        /// </summary>
        [Output("callLimits")]
        public Output<int> CallLimits { get; private set; } = null!;

        /// <summary>
        /// The creation time of the application quota, in RFC3339 format.
        /// </summary>
        [Output("createdAt")]
        public Output<string> CreatedAt { get; private set; } = null!;

        /// <summary>
        /// Specifies the description of the application quota.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// Specifies the ID of the dedicated instance to which the application quota belongs.
        /// </summary>
        [Output("instanceId")]
        public Output<string> InstanceId { get; private set; } = null!;

        /// <summary>
        /// Specifies the name of the application quota.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        [Output("region")]
        public Output<string> Region { get; private set; } = null!;

        /// <summary>
        /// Specifies the limited time value for flow control of the application quota.
        /// </summary>
        [Output("timeInterval")]
        public Output<int> TimeInterval { get; private set; } = null!;

        /// <summary>
        /// Specifies the limited time unit of the application quota.
        /// </summary>
        [Output("timeUnit")]
        public Output<string> TimeUnit { get; private set; } = null!;


        /// <summary>
        /// Create a ApigApplicationQuota resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ApigApplicationQuota(string name, ApigApplicationQuotaArgs args, CustomResourceOptions? options = null)
            : base("sbercloud:index/apigApplicationQuota:ApigApplicationQuota", name, args ?? new ApigApplicationQuotaArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ApigApplicationQuota(string name, Input<string> id, ApigApplicationQuotaState? state = null, CustomResourceOptions? options = null)
            : base("sbercloud:index/apigApplicationQuota:ApigApplicationQuota", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing ApigApplicationQuota resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ApigApplicationQuota Get(string name, Input<string> id, ApigApplicationQuotaState? state = null, CustomResourceOptions? options = null)
        {
            return new ApigApplicationQuota(name, id, state, options);
        }
    }

    public sealed class ApigApplicationQuotaArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specifies the access limit of the application quota.
        /// </summary>
        [Input("callLimits", required: true)]
        public Input<int> CallLimits { get; set; } = null!;

        /// <summary>
        /// Specifies the description of the application quota.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Specifies the ID of the dedicated instance to which the application quota belongs.
        /// </summary>
        [Input("instanceId", required: true)]
        public Input<string> InstanceId { get; set; } = null!;

        /// <summary>
        /// Specifies the name of the application quota.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("region")]
        public Input<string>? Region { get; set; }

        /// <summary>
        /// Specifies the limited time value for flow control of the application quota.
        /// </summary>
        [Input("timeInterval", required: true)]
        public Input<int> TimeInterval { get; set; } = null!;

        /// <summary>
        /// Specifies the limited time unit of the application quota.
        /// </summary>
        [Input("timeUnit", required: true)]
        public Input<string> TimeUnit { get; set; } = null!;

        public ApigApplicationQuotaArgs()
        {
        }
        public static new ApigApplicationQuotaArgs Empty => new ApigApplicationQuotaArgs();
    }

    public sealed class ApigApplicationQuotaState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The number of bound APPs.
        /// </summary>
        [Input("bindNum")]
        public Input<int>? BindNum { get; set; }

        /// <summary>
        /// Specifies the access limit of the application quota.
        /// </summary>
        [Input("callLimits")]
        public Input<int>? CallLimits { get; set; }

        /// <summary>
        /// The creation time of the application quota, in RFC3339 format.
        /// </summary>
        [Input("createdAt")]
        public Input<string>? CreatedAt { get; set; }

        /// <summary>
        /// Specifies the description of the application quota.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Specifies the ID of the dedicated instance to which the application quota belongs.
        /// </summary>
        [Input("instanceId")]
        public Input<string>? InstanceId { get; set; }

        /// <summary>
        /// Specifies the name of the application quota.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("region")]
        public Input<string>? Region { get; set; }

        /// <summary>
        /// Specifies the limited time value for flow control of the application quota.
        /// </summary>
        [Input("timeInterval")]
        public Input<int>? TimeInterval { get; set; }

        /// <summary>
        /// Specifies the limited time unit of the application quota.
        /// </summary>
        [Input("timeUnit")]
        public Input<string>? TimeUnit { get; set; }

        public ApigApplicationQuotaState()
        {
        }
        public static new ApigApplicationQuotaState Empty => new ApigApplicationQuotaState();
    }
}
