// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Sbercloud
{
    [SbercloudResourceType("sbercloud:index/obsBucketPolicy:ObsBucketPolicy")]
    public partial class ObsBucketPolicy : global::Pulumi.CustomResource
    {
        [Output("bucket")]
        public Output<string> Bucket { get; private set; } = null!;

        [Output("policy")]
        public Output<string> Policy { get; private set; } = null!;

        [Output("policyFormat")]
        public Output<string?> PolicyFormat { get; private set; } = null!;

        [Output("region")]
        public Output<string> Region { get; private set; } = null!;


        /// <summary>
        /// Create a ObsBucketPolicy resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ObsBucketPolicy(string name, ObsBucketPolicyArgs args, CustomResourceOptions? options = null)
            : base("sbercloud:index/obsBucketPolicy:ObsBucketPolicy", name, args ?? new ObsBucketPolicyArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ObsBucketPolicy(string name, Input<string> id, ObsBucketPolicyState? state = null, CustomResourceOptions? options = null)
            : base("sbercloud:index/obsBucketPolicy:ObsBucketPolicy", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing ObsBucketPolicy resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ObsBucketPolicy Get(string name, Input<string> id, ObsBucketPolicyState? state = null, CustomResourceOptions? options = null)
        {
            return new ObsBucketPolicy(name, id, state, options);
        }
    }

    public sealed class ObsBucketPolicyArgs : global::Pulumi.ResourceArgs
    {
        [Input("bucket", required: true)]
        public Input<string> Bucket { get; set; } = null!;

        [Input("policy", required: true)]
        public Input<string> Policy { get; set; } = null!;

        [Input("policyFormat")]
        public Input<string>? PolicyFormat { get; set; }

        [Input("region")]
        public Input<string>? Region { get; set; }

        public ObsBucketPolicyArgs()
        {
        }
        public static new ObsBucketPolicyArgs Empty => new ObsBucketPolicyArgs();
    }

    public sealed class ObsBucketPolicyState : global::Pulumi.ResourceArgs
    {
        [Input("bucket")]
        public Input<string>? Bucket { get; set; }

        [Input("policy")]
        public Input<string>? Policy { get; set; }

        [Input("policyFormat")]
        public Input<string>? PolicyFormat { get; set; }

        [Input("region")]
        public Input<string>? Region { get; set; }

        public ObsBucketPolicyState()
        {
        }
        public static new ObsBucketPolicyState Empty => new ObsBucketPolicyState();
    }
}
