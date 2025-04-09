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
    public sealed class GetComputeInstanceVolumeAttachedResult
    {
        public readonly int BootIndex;
        public readonly bool IsSysVolume;
        public readonly string PciAddress;
        public readonly int Size;
        public readonly string Type;
        public readonly string VolumeId;

        [OutputConstructor]
        private GetComputeInstanceVolumeAttachedResult(
            int bootIndex,

            bool isSysVolume,

            string pciAddress,

            int size,

            string type,

            string volumeId)
        {
            BootIndex = bootIndex;
            IsSysVolume = isSysVolume;
            PciAddress = pciAddress;
            Size = size;
            Type = type;
            VolumeId = volumeId;
        }
    }
}
