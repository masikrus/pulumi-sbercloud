// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Sbercloud.Inputs
{

    public sealed class CssClusterNodeConfigNetworkInfoArgs : global::Pulumi.ResourceArgs
    {
        [Input("securityGroupId", required: true)]
        public Input<string> SecurityGroupId { get; set; } = null!;

        [Input("subnetId", required: true)]
        public Input<string> SubnetId { get; set; } = null!;

        [Input("vpcId", required: true)]
        public Input<string> VpcId { get; set; } = null!;

        public CssClusterNodeConfigNetworkInfoArgs()
        {
        }
        public static new CssClusterNodeConfigNetworkInfoArgs Empty => new CssClusterNodeConfigNetworkInfoArgs();
    }
}
