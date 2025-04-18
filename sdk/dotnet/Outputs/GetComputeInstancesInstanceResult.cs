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
    public sealed class GetComputeInstancesInstanceResult
    {
        public readonly string AvailabilityZone;
        public readonly string ChargingMode;
        public readonly string EnterpriseProjectId;
        public readonly string ExpiredTime;
        public readonly string FlavorId;
        public readonly string FlavorName;
        public readonly string Id;
        public readonly string ImageId;
        public readonly string ImageName;
        public readonly string KeyPair;
        public readonly string Name;
        public readonly ImmutableArray<Outputs.GetComputeInstancesInstanceNetworkResult> Networks;
        public readonly string PublicIp;
        public readonly ImmutableArray<Outputs.GetComputeInstancesInstanceSchedulerHintResult> SchedulerHints;
        public readonly ImmutableArray<string> SecurityGroupIds;
        public readonly string Status;
        public readonly string SystemDiskId;
        public readonly ImmutableDictionary<string, string> Tags;
        public readonly string UserData;
        public readonly ImmutableArray<Outputs.GetComputeInstancesInstanceVolumeAttachedResult> VolumeAttacheds;

        [OutputConstructor]
        private GetComputeInstancesInstanceResult(
            string availabilityZone,

            string chargingMode,

            string enterpriseProjectId,

            string expiredTime,

            string flavorId,

            string flavorName,

            string id,

            string imageId,

            string imageName,

            string keyPair,

            string name,

            ImmutableArray<Outputs.GetComputeInstancesInstanceNetworkResult> networks,

            string publicIp,

            ImmutableArray<Outputs.GetComputeInstancesInstanceSchedulerHintResult> schedulerHints,

            ImmutableArray<string> securityGroupIds,

            string status,

            string systemDiskId,

            ImmutableDictionary<string, string> tags,

            string userData,

            ImmutableArray<Outputs.GetComputeInstancesInstanceVolumeAttachedResult> volumeAttacheds)
        {
            AvailabilityZone = availabilityZone;
            ChargingMode = chargingMode;
            EnterpriseProjectId = enterpriseProjectId;
            ExpiredTime = expiredTime;
            FlavorId = flavorId;
            FlavorName = flavorName;
            Id = id;
            ImageId = imageId;
            ImageName = imageName;
            KeyPair = keyPair;
            Name = name;
            Networks = networks;
            PublicIp = publicIp;
            SchedulerHints = schedulerHints;
            SecurityGroupIds = securityGroupIds;
            Status = status;
            SystemDiskId = systemDiskId;
            Tags = tags;
            UserData = userData;
            VolumeAttacheds = volumeAttacheds;
        }
    }
}
