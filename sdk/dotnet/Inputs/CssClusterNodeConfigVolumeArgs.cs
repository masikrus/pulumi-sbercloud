// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Sbercloud.Inputs
{

    public sealed class CssClusterNodeConfigVolumeArgs : global::Pulumi.ResourceArgs
    {
        [Input("size", required: true)]
        public Input<int> Size { get; set; } = null!;

        [Input("volumeType", required: true)]
        public Input<string> VolumeType { get; set; } = null!;

        public CssClusterNodeConfigVolumeArgs()
        {
        }
        public static new CssClusterNodeConfigVolumeArgs Empty => new CssClusterNodeConfigVolumeArgs();
    }
}
