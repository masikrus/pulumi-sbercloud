// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Sbercloud.Inputs
{

    public sealed class NatGatewaySessionConfArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ICMP session expiration time.
        /// </summary>
        [Input("icmpSessionExpireTime")]
        public Input<int>? IcmpSessionExpireTime { get; set; }

        /// <summary>
        /// The TCP session expiration time.
        /// </summary>
        [Input("tcpSessionExpireTime")]
        public Input<int>? TcpSessionExpireTime { get; set; }

        /// <summary>
        /// The duration of TIME_WAIT state when TCP connection is closed.
        /// </summary>
        [Input("tcpTimeWaitTime")]
        public Input<int>? TcpTimeWaitTime { get; set; }

        /// <summary>
        /// The UDP session expiration time.
        /// </summary>
        [Input("udpSessionExpireTime")]
        public Input<int>? UdpSessionExpireTime { get; set; }

        public NatGatewaySessionConfArgs()
        {
        }
        public static new NatGatewaySessionConfArgs Empty => new NatGatewaySessionConfArgs();
    }
}
