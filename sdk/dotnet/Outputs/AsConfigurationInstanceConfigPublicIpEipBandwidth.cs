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
    public sealed class AsConfigurationInstanceConfigPublicIpEipBandwidth
    {
        public readonly string? ChargingMode;
        public readonly string? Id;
        public readonly string ShareType;
        public readonly int? Size;

        [OutputConstructor]
        private AsConfigurationInstanceConfigPublicIpEipBandwidth(
            string? chargingMode,

            string? id,

            string shareType,

            int? size)
        {
            ChargingMode = chargingMode;
            Id = id;
            ShareType = shareType;
            Size = size;
        }
    }
}
