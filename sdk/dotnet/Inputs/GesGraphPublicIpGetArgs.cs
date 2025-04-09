// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Sbercloud.Inputs
{

    public sealed class GesGraphPublicIpGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The EIP ID.
        /// </summary>
        [Input("eipId")]
        public Input<string>? EipId { get; set; }

        /// <summary>
        /// The bind type of public IP.
        /// </summary>
        [Input("publicBindType")]
        public Input<string>? PublicBindType { get; set; }

        public GesGraphPublicIpGetArgs()
        {
        }
        public static new GesGraphPublicIpGetArgs Empty => new GesGraphPublicIpGetArgs();
    }
}
