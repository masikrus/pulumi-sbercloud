// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Sbercloud.Inputs
{

    public sealed class CceNodeAttachTaintArgs : global::Pulumi.ResourceArgs
    {
        [Input("effect", required: true)]
        public Input<string> Effect { get; set; } = null!;

        [Input("key", required: true)]
        public Input<string> Key { get; set; } = null!;

        [Input("value", required: true)]
        public Input<string> Value { get; set; } = null!;

        public CceNodeAttachTaintArgs()
        {
        }
        public static new CceNodeAttachTaintArgs Empty => new CceNodeAttachTaintArgs();
    }
}
