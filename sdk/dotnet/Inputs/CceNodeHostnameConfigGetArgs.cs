// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Sbercloud.Inputs
{

    public sealed class CceNodeHostnameConfigGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        public CceNodeHostnameConfigGetArgs()
        {
        }
        public static new CceNodeHostnameConfigGetArgs Empty => new CceNodeHostnameConfigGetArgs();
    }
}
