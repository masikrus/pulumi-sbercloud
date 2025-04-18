// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Sbercloud
{
    [SbercloudResourceType("sbercloud:index/computeVolumeAttach:ComputeVolumeAttach")]
    public partial class ComputeVolumeAttach : global::Pulumi.CustomResource
    {
        [Output("device")]
        public Output<string> Device { get; private set; } = null!;

        [Output("instanceId")]
        public Output<string> InstanceId { get; private set; } = null!;

        [Output("pciAddress")]
        public Output<string> PciAddress { get; private set; } = null!;

        [Output("region")]
        public Output<string> Region { get; private set; } = null!;

        [Output("volumeId")]
        public Output<string> VolumeId { get; private set; } = null!;


        /// <summary>
        /// Create a ComputeVolumeAttach resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ComputeVolumeAttach(string name, ComputeVolumeAttachArgs args, CustomResourceOptions? options = null)
            : base("sbercloud:index/computeVolumeAttach:ComputeVolumeAttach", name, args ?? new ComputeVolumeAttachArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ComputeVolumeAttach(string name, Input<string> id, ComputeVolumeAttachState? state = null, CustomResourceOptions? options = null)
            : base("sbercloud:index/computeVolumeAttach:ComputeVolumeAttach", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing ComputeVolumeAttach resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ComputeVolumeAttach Get(string name, Input<string> id, ComputeVolumeAttachState? state = null, CustomResourceOptions? options = null)
        {
            return new ComputeVolumeAttach(name, id, state, options);
        }
    }

    public sealed class ComputeVolumeAttachArgs : global::Pulumi.ResourceArgs
    {
        [Input("device")]
        public Input<string>? Device { get; set; }

        [Input("instanceId", required: true)]
        public Input<string> InstanceId { get; set; } = null!;

        [Input("region")]
        public Input<string>? Region { get; set; }

        [Input("volumeId", required: true)]
        public Input<string> VolumeId { get; set; } = null!;

        public ComputeVolumeAttachArgs()
        {
        }
        public static new ComputeVolumeAttachArgs Empty => new ComputeVolumeAttachArgs();
    }

    public sealed class ComputeVolumeAttachState : global::Pulumi.ResourceArgs
    {
        [Input("device")]
        public Input<string>? Device { get; set; }

        [Input("instanceId")]
        public Input<string>? InstanceId { get; set; }

        [Input("pciAddress")]
        public Input<string>? PciAddress { get; set; }

        [Input("region")]
        public Input<string>? Region { get; set; }

        [Input("volumeId")]
        public Input<string>? VolumeId { get; set; }

        public ComputeVolumeAttachState()
        {
        }
        public static new ComputeVolumeAttachState Empty => new ComputeVolumeAttachState();
    }
}
