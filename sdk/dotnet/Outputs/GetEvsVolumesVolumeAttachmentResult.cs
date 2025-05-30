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
    public sealed class GetEvsVolumesVolumeAttachmentResult
    {
        public readonly string AttachedAt;
        public readonly string AttachedMode;
        public readonly string DeviceName;
        public readonly string Id;
        public readonly string ServerId;

        [OutputConstructor]
        private GetEvsVolumesVolumeAttachmentResult(
            string attachedAt,

            string attachedMode,

            string deviceName,

            string id,

            string serverId)
        {
            AttachedAt = attachedAt;
            AttachedMode = attachedMode;
            DeviceName = deviceName;
            Id = id;
            ServerId = serverId;
        }
    }
}
