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
    public sealed class GesGraphPublicIp
    {
        /// <summary>
        /// The EIP ID.
        /// </summary>
        public readonly string? EipId;
        /// <summary>
        /// The bind type of public IP.
        /// </summary>
        public readonly string? PublicBindType;

        [OutputConstructor]
        private GesGraphPublicIp(
            string? eipId,

            string? publicBindType)
        {
            EipId = eipId;
            PublicBindType = publicBindType;
        }
    }
}
