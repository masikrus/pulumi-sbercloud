// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Sbercloud.Inputs
{

    public sealed class CceNodeExtensionNicArgs : global::Pulumi.ResourceArgs
    {
        [Input("subnetId", required: true)]
        public Input<string> SubnetId { get; set; } = null!;

        public CceNodeExtensionNicArgs()
        {
        }
        public static new CceNodeExtensionNicArgs Empty => new CceNodeExtensionNicArgs();
    }
}
