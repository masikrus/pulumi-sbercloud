// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Sbercloud
{
    [SbercloudResourceType("sbercloud:index/kmsKey:KmsKey")]
    public partial class KmsKey : global::Pulumi.CustomResource
    {
        [Output("creationDate")]
        public Output<string> CreationDate { get; private set; } = null!;

        [Output("defaultKeyFlag")]
        public Output<string> DefaultKeyFlag { get; private set; } = null!;

        [Output("domainId")]
        public Output<string> DomainId { get; private set; } = null!;

        [Output("enterpriseProjectId")]
        public Output<string> EnterpriseProjectId { get; private set; } = null!;

        [Output("expirationTime")]
        public Output<string> ExpirationTime { get; private set; } = null!;

        [Output("isEnabled")]
        public Output<bool?> IsEnabled { get; private set; } = null!;

        [Output("keyAlgorithm")]
        public Output<string> KeyAlgorithm { get; private set; } = null!;

        [Output("keyAlias")]
        public Output<string> KeyAlias { get; private set; } = null!;

        [Output("keyDescription")]
        public Output<string?> KeyDescription { get; private set; } = null!;

        [Output("keyId")]
        public Output<string> KeyId { get; private set; } = null!;

        [Output("keyState")]
        public Output<string> KeyState { get; private set; } = null!;

        [Output("keyUsage")]
        public Output<string> KeyUsage { get; private set; } = null!;

        [Output("keystoreId")]
        public Output<string> KeystoreId { get; private set; } = null!;

        [Output("origin")]
        public Output<string> Origin { get; private set; } = null!;

        [Output("pendingDays")]
        public Output<string?> PendingDays { get; private set; } = null!;

        [Output("region")]
        public Output<string> Region { get; private set; } = null!;

        [Output("rotationEnabled")]
        public Output<bool?> RotationEnabled { get; private set; } = null!;

        [Output("rotationInterval")]
        public Output<int> RotationInterval { get; private set; } = null!;

        [Output("rotationNumber")]
        public Output<int> RotationNumber { get; private set; } = null!;

        [Output("scheduledDeletionDate")]
        public Output<string> ScheduledDeletionDate { get; private set; } = null!;

        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;


        /// <summary>
        /// Create a KmsKey resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public KmsKey(string name, KmsKeyArgs args, CustomResourceOptions? options = null)
            : base("sbercloud:index/kmsKey:KmsKey", name, args ?? new KmsKeyArgs(), MakeResourceOptions(options, ""))
        {
        }

        private KmsKey(string name, Input<string> id, KmsKeyState? state = null, CustomResourceOptions? options = null)
            : base("sbercloud:index/kmsKey:KmsKey", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing KmsKey resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static KmsKey Get(string name, Input<string> id, KmsKeyState? state = null, CustomResourceOptions? options = null)
        {
            return new KmsKey(name, id, state, options);
        }
    }

    public sealed class KmsKeyArgs : global::Pulumi.ResourceArgs
    {
        [Input("enterpriseProjectId")]
        public Input<string>? EnterpriseProjectId { get; set; }

        [Input("isEnabled")]
        public Input<bool>? IsEnabled { get; set; }

        [Input("keyAlgorithm")]
        public Input<string>? KeyAlgorithm { get; set; }

        [Input("keyAlias", required: true)]
        public Input<string> KeyAlias { get; set; } = null!;

        [Input("keyDescription")]
        public Input<string>? KeyDescription { get; set; }

        [Input("keyUsage")]
        public Input<string>? KeyUsage { get; set; }

        [Input("keystoreId")]
        public Input<string>? KeystoreId { get; set; }

        [Input("origin")]
        public Input<string>? Origin { get; set; }

        [Input("pendingDays")]
        public Input<string>? PendingDays { get; set; }

        [Input("region")]
        public Input<string>? Region { get; set; }

        [Input("rotationEnabled")]
        public Input<bool>? RotationEnabled { get; set; }

        [Input("rotationInterval")]
        public Input<int>? RotationInterval { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        public KmsKeyArgs()
        {
        }
        public static new KmsKeyArgs Empty => new KmsKeyArgs();
    }

    public sealed class KmsKeyState : global::Pulumi.ResourceArgs
    {
        [Input("creationDate")]
        public Input<string>? CreationDate { get; set; }

        [Input("defaultKeyFlag")]
        public Input<string>? DefaultKeyFlag { get; set; }

        [Input("domainId")]
        public Input<string>? DomainId { get; set; }

        [Input("enterpriseProjectId")]
        public Input<string>? EnterpriseProjectId { get; set; }

        [Input("expirationTime")]
        public Input<string>? ExpirationTime { get; set; }

        [Input("isEnabled")]
        public Input<bool>? IsEnabled { get; set; }

        [Input("keyAlgorithm")]
        public Input<string>? KeyAlgorithm { get; set; }

        [Input("keyAlias")]
        public Input<string>? KeyAlias { get; set; }

        [Input("keyDescription")]
        public Input<string>? KeyDescription { get; set; }

        [Input("keyId")]
        public Input<string>? KeyId { get; set; }

        [Input("keyState")]
        public Input<string>? KeyState { get; set; }

        [Input("keyUsage")]
        public Input<string>? KeyUsage { get; set; }

        [Input("keystoreId")]
        public Input<string>? KeystoreId { get; set; }

        [Input("origin")]
        public Input<string>? Origin { get; set; }

        [Input("pendingDays")]
        public Input<string>? PendingDays { get; set; }

        [Input("region")]
        public Input<string>? Region { get; set; }

        [Input("rotationEnabled")]
        public Input<bool>? RotationEnabled { get; set; }

        [Input("rotationInterval")]
        public Input<int>? RotationInterval { get; set; }

        [Input("rotationNumber")]
        public Input<int>? RotationNumber { get; set; }

        [Input("scheduledDeletionDate")]
        public Input<string>? ScheduledDeletionDate { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        public KmsKeyState()
        {
        }
        public static new KmsKeyState Empty => new KmsKeyState();
    }
}
