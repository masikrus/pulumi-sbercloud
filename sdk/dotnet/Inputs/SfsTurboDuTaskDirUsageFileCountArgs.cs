// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Sbercloud.Inputs
{

    public sealed class SfsTurboDuTaskDirUsageFileCountArgs : global::Pulumi.ResourceArgs
    {
        [Input("block")]
        public Input<int>? Block { get; set; }

        [Input("char")]
        public Input<int>? Char { get; set; }

        [Input("dir")]
        public Input<int>? Dir { get; set; }

        [Input("pipe")]
        public Input<int>? Pipe { get; set; }

        [Input("regular")]
        public Input<int>? Regular { get; set; }

        [Input("socket")]
        public Input<int>? Socket { get; set; }

        [Input("symlink")]
        public Input<int>? Symlink { get; set; }

        public SfsTurboDuTaskDirUsageFileCountArgs()
        {
        }
        public static new SfsTurboDuTaskDirUsageFileCountArgs Empty => new SfsTurboDuTaskDirUsageFileCountArgs();
    }
}
