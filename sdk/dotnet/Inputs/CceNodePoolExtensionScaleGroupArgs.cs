// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Sbercloud.Inputs
{

    public sealed class CceNodePoolExtensionScaleGroupArgs : global::Pulumi.ResourceArgs
    {
        [Input("metadata")]
        public Input<Inputs.CceNodePoolExtensionScaleGroupMetadataArgs>? Metadata { get; set; }

        [Input("spec")]
        public Input<Inputs.CceNodePoolExtensionScaleGroupSpecArgs>? Spec { get; set; }

        public CceNodePoolExtensionScaleGroupArgs()
        {
        }
        public static new CceNodePoolExtensionScaleGroupArgs Empty => new CceNodePoolExtensionScaleGroupArgs();
    }
}
