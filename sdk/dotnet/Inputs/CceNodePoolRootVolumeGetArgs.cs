// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Sbercloud.Inputs
{

    public sealed class CceNodePoolRootVolumeGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("dssPoolId")]
        public Input<string>? DssPoolId { get; set; }

        [Input("extendParam")]
        public Input<string>? ExtendParam { get; set; }

        [Input("extendParams")]
        private InputMap<string>? _extendParams;
        public InputMap<string> ExtendParams
        {
            get => _extendParams ?? (_extendParams = new InputMap<string>());
            set => _extendParams = value;
        }

        /// <summary>
        /// schema: Internal
        /// </summary>
        [Input("hwPassthrough")]
        public Input<bool>? HwPassthrough { get; set; }

        [Input("iops")]
        public Input<int>? Iops { get; set; }

        [Input("kmsKeyId")]
        public Input<string>? KmsKeyId { get; set; }

        [Input("size", required: true)]
        public Input<int> Size { get; set; } = null!;

        [Input("throughput")]
        public Input<int>? Throughput { get; set; }

        [Input("volumetype", required: true)]
        public Input<string> Volumetype { get; set; } = null!;

        public CceNodePoolRootVolumeGetArgs()
        {
        }
        public static new CceNodePoolRootVolumeGetArgs Empty => new CceNodePoolRootVolumeGetArgs();
    }
}
