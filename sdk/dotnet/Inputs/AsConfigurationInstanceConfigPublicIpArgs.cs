// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Sbercloud.Inputs
{

    public sealed class AsConfigurationInstanceConfigPublicIpArgs : global::Pulumi.ResourceArgs
    {
        [Input("eip", required: true)]
        public Input<Inputs.AsConfigurationInstanceConfigPublicIpEipArgs> Eip { get; set; } = null!;

        public AsConfigurationInstanceConfigPublicIpArgs()
        {
        }
        public static new AsConfigurationInstanceConfigPublicIpArgs Empty => new AsConfigurationInstanceConfigPublicIpArgs();
    }
}
