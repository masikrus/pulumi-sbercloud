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
    public sealed class RdsInstanceMsdtcHost
    {
        public readonly string HostName;
        public readonly string? Id;
        public readonly string Ip;

        [OutputConstructor]
        private RdsInstanceMsdtcHost(
            string hostName,

            string? id,

            string ip)
        {
            HostName = hostName;
            Id = id;
            Ip = ip;
        }
    }
}
