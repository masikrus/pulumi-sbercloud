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
    public sealed class ElbIpgroupIpList
    {
        public readonly string? Description;
        public readonly string Ip;

        [OutputConstructor]
        private ElbIpgroupIpList(
            string? description,

            string ip)
        {
            Description = description;
            Ip = ip;
        }
    }
}
