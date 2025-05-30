// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Sbercloud
{
    [SbercloudResourceType("sbercloud:index/apigSignatureAssociate:ApigSignatureAssociate")]
    public partial class ApigSignatureAssociate : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The ID of the dedicated instance to which the APIs and the signature belong.
        /// </summary>
        [Output("instanceId")]
        public Output<string> InstanceId { get; private set; } = null!;

        /// <summary>
        /// The publish IDs corresponding to the APIs bound by the signature.
        /// </summary>
        [Output("publishIds")]
        public Output<ImmutableArray<string>> PublishIds { get; private set; } = null!;

        /// <summary>
        /// The region where the signature and the APIs are located.
        /// </summary>
        [Output("region")]
        public Output<string> Region { get; private set; } = null!;

        /// <summary>
        /// The signature ID for APIs binding.
        /// </summary>
        [Output("signatureId")]
        public Output<string> SignatureId { get; private set; } = null!;


        /// <summary>
        /// Create a ApigSignatureAssociate resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ApigSignatureAssociate(string name, ApigSignatureAssociateArgs args, CustomResourceOptions? options = null)
            : base("sbercloud:index/apigSignatureAssociate:ApigSignatureAssociate", name, args ?? new ApigSignatureAssociateArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ApigSignatureAssociate(string name, Input<string> id, ApigSignatureAssociateState? state = null, CustomResourceOptions? options = null)
            : base("sbercloud:index/apigSignatureAssociate:ApigSignatureAssociate", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing ApigSignatureAssociate resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ApigSignatureAssociate Get(string name, Input<string> id, ApigSignatureAssociateState? state = null, CustomResourceOptions? options = null)
        {
            return new ApigSignatureAssociate(name, id, state, options);
        }
    }

    public sealed class ApigSignatureAssociateArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ID of the dedicated instance to which the APIs and the signature belong.
        /// </summary>
        [Input("instanceId", required: true)]
        public Input<string> InstanceId { get; set; } = null!;

        [Input("publishIds", required: true)]
        private InputList<string>? _publishIds;

        /// <summary>
        /// The publish IDs corresponding to the APIs bound by the signature.
        /// </summary>
        public InputList<string> PublishIds
        {
            get => _publishIds ?? (_publishIds = new InputList<string>());
            set => _publishIds = value;
        }

        /// <summary>
        /// The region where the signature and the APIs are located.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        /// <summary>
        /// The signature ID for APIs binding.
        /// </summary>
        [Input("signatureId", required: true)]
        public Input<string> SignatureId { get; set; } = null!;

        public ApigSignatureAssociateArgs()
        {
        }
        public static new ApigSignatureAssociateArgs Empty => new ApigSignatureAssociateArgs();
    }

    public sealed class ApigSignatureAssociateState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ID of the dedicated instance to which the APIs and the signature belong.
        /// </summary>
        [Input("instanceId")]
        public Input<string>? InstanceId { get; set; }

        [Input("publishIds")]
        private InputList<string>? _publishIds;

        /// <summary>
        /// The publish IDs corresponding to the APIs bound by the signature.
        /// </summary>
        public InputList<string> PublishIds
        {
            get => _publishIds ?? (_publishIds = new InputList<string>());
            set => _publishIds = value;
        }

        /// <summary>
        /// The region where the signature and the APIs are located.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        /// <summary>
        /// The signature ID for APIs binding.
        /// </summary>
        [Input("signatureId")]
        public Input<string>? SignatureId { get; set; }

        public ApigSignatureAssociateState()
        {
        }
        public static new ApigSignatureAssociateState Empty => new ApigSignatureAssociateState();
    }
}
