// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Sbercloud.Outputs
{

    [OutputType]
    public sealed class AsGroupNetwork
    {
        public readonly string Id;
        public readonly string? Ipv6BandwidthId;
        public readonly bool? Ipv6Enable;
        public readonly bool? SourceDestCheck;

        [OutputConstructor]
        private AsGroupNetwork(
            string id,

            string? ipv6BandwidthId,

            bool? ipv6Enable,

            bool? sourceDestCheck)
        {
            Id = id;
            Ipv6BandwidthId = ipv6BandwidthId;
            Ipv6Enable = ipv6Enable;
            SourceDestCheck = sourceDestCheck;
        }
    }
}
