// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Sbercloud.Inputs
{

    public sealed class CceNodePoolStorageGroupVirtualSpaceArgs : global::Pulumi.ResourceArgs
    {
        [Input("lvmLvType")]
        public Input<string>? LvmLvType { get; set; }

        [Input("lvmPath")]
        public Input<string>? LvmPath { get; set; }

        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        [Input("runtimeLvType")]
        public Input<string>? RuntimeLvType { get; set; }

        [Input("size", required: true)]
        public Input<string> Size { get; set; } = null!;

        public CceNodePoolStorageGroupVirtualSpaceArgs()
        {
        }
        public static new CceNodePoolStorageGroupVirtualSpaceArgs Empty => new CceNodePoolStorageGroupVirtualSpaceArgs();
    }
}
