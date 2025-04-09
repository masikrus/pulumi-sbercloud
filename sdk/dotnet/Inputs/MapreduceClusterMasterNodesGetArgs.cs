// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Sbercloud.Inputs
{

    public sealed class MapreduceClusterMasterNodesGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("assignedRoles")]
        private InputList<string>? _assignedRoles;
        public InputList<string> AssignedRoles
        {
            get => _assignedRoles ?? (_assignedRoles = new InputList<string>());
            set => _assignedRoles = value;
        }

        [Input("dataVolumeCount", required: true)]
        public Input<int> DataVolumeCount { get; set; } = null!;

        [Input("dataVolumeSize")]
        public Input<int>? DataVolumeSize { get; set; }

        [Input("dataVolumeType")]
        public Input<string>? DataVolumeType { get; set; }

        [Input("flavor", required: true)]
        public Input<string> Flavor { get; set; } = null!;

        [Input("hostIps")]
        private InputList<string>? _hostIps;
        public InputList<string> HostIps
        {
            get => _hostIps ?? (_hostIps = new InputList<string>());
            set => _hostIps = value;
        }

        [Input("nodeNumber", required: true)]
        public Input<int> NodeNumber { get; set; } = null!;

        [Input("rootVolumeSize", required: true)]
        public Input<int> RootVolumeSize { get; set; } = null!;

        [Input("rootVolumeType", required: true)]
        public Input<string> RootVolumeType { get; set; } = null!;

        public MapreduceClusterMasterNodesGetArgs()
        {
        }
        public static new MapreduceClusterMasterNodesGetArgs Empty => new MapreduceClusterMasterNodesGetArgs();
    }
}
