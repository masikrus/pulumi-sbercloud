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
    public sealed class VpcEipBandwidth
    {
        /// <summary>
        /// Whether the bandwidth is billed by traffic or by bandwidth size.
        /// </summary>
        public readonly string? ChargeMode;
        /// <summary>
        /// The shared bandwidth ID.
        /// </summary>
        public readonly string? Id;
        /// <summary>
        /// The dedicated bandwidth name.
        /// </summary>
        public readonly string? Name;
        /// <summary>
        /// Whether the bandwidth is dedicated or shared.
        /// </summary>
        public readonly string ShareType;
        /// <summary>
        /// The dedicated bandwidth size.
        /// </summary>
        public readonly int? Size;

        [OutputConstructor]
        private VpcEipBandwidth(
            string? chargeMode,

            string? id,

            string? name,

            string shareType,

            int? size)
        {
            ChargeMode = chargeMode;
            Id = id;
            Name = name;
            ShareType = shareType;
            Size = size;
        }
    }
}
