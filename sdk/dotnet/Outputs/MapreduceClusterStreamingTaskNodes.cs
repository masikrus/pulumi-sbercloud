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
    public sealed class MapreduceClusterStreamingTaskNodes
    {
        public readonly ImmutableArray<string> AssignedRoles;
        public readonly int DataVolumeCount;
        public readonly int? DataVolumeSize;
        public readonly string? DataVolumeType;
        public readonly string Flavor;
        public readonly ImmutableArray<string> HostIps;
        public readonly int NodeNumber;
        public readonly int RootVolumeSize;
        public readonly string RootVolumeType;

        [OutputConstructor]
        private MapreduceClusterStreamingTaskNodes(
            ImmutableArray<string> assignedRoles,

            int dataVolumeCount,

            int? dataVolumeSize,

            string? dataVolumeType,

            string flavor,

            ImmutableArray<string> hostIps,

            int nodeNumber,

            int rootVolumeSize,

            string rootVolumeType)
        {
            AssignedRoles = assignedRoles;
            DataVolumeCount = dataVolumeCount;
            DataVolumeSize = dataVolumeSize;
            DataVolumeType = dataVolumeType;
            Flavor = flavor;
            HostIps = hostIps;
            NodeNumber = nodeNumber;
            RootVolumeSize = rootVolumeSize;
            RootVolumeType = rootVolumeType;
        }
    }
}
