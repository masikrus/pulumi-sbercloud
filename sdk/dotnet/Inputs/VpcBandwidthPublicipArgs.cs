// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Sbercloud.Inputs
{

    public sealed class VpcBandwidthPublicipArgs : global::Pulumi.ResourceArgs
    {
        [Input("id")]
        public Input<string>? Id { get; set; }

        [Input("ipAddress")]
        public Input<string>? IpAddress { get; set; }

        [Input("ipVersion")]
        public Input<int>? IpVersion { get; set; }

        [Input("type")]
        public Input<string>? Type { get; set; }

        public VpcBandwidthPublicipArgs()
        {
        }
        public static new VpcBandwidthPublicipArgs Empty => new VpcBandwidthPublicipArgs();
    }
}
