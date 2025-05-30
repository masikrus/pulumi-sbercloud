// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Sbercloud
{
    [SbercloudResourceType("sbercloud:index/dwsCluster:DwsCluster")]
    public partial class DwsCluster : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The availability zone in which to create the cluster instance.
        /// </summary>
        [Output("availabilityZone")]
        public Output<string> AvailabilityZone { get; private set; } = null!;

        /// <summary>
        /// The creation time of the cluster.
        /// </summary>
        [Output("created")]
        public Output<string> Created { get; private set; } = null!;

        /// <summary>
        /// The description of the cluster.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// Dedicated storage pool ID.
        /// </summary>
        [Output("dssPoolId")]
        public Output<string> DssPoolId { get; private set; } = null!;

        /// <summary>
        /// The ID of the ELB load balancer.
        /// </summary>
        [Output("elbId")]
        public Output<string?> ElbId { get; private set; } = null!;

        /// <summary>
        /// The ELB information bound to the cluster.
        /// </summary>
        [Output("elbs")]
        public Output<ImmutableArray<Outputs.DwsClusterElb>> Elbs { get; private set; } = null!;

        /// <summary>
        /// Private network connection information about the cluster.
        /// </summary>
        [Output("endpoints")]
        public Output<ImmutableArray<Outputs.DwsClusterEndpoint>> Endpoints { get; private set; } = null!;

        /// <summary>
        /// The enterprise project ID.
        /// </summary>
        [Output("enterpriseProjectId")]
        public Output<string> EnterpriseProjectId { get; private set; } = null!;

        /// <summary>
        /// Whether to automatically execute snapshot when shrinking the number of nodes.
        /// </summary>
        [Output("forceBackup")]
        public Output<bool?> ForceBackup { get; private set; } = null!;

        /// <summary>
        /// The number of latest manual snapshots that need to be retained when deleting the cluster.
        /// </summary>
        [Output("keepLastManualSnapshot")]
        public Output<int?> KeepLastManualSnapshot { get; private set; } = null!;

        /// <summary>
        /// The KMS key ID.
        /// </summary>
        [Output("kmsKeyId")]
        public Output<string> KmsKeyId { get; private set; } = null!;

        /// <summary>
        /// Whether to enable logical cluster.
        /// </summary>
        [Output("logicalClusterEnable")]
        public Output<bool?> LogicalClusterEnable { get; private set; } = null!;

        /// <summary>
        /// Whether to enable LTS.
        /// </summary>
        [Output("ltsEnable")]
        public Output<bool?> LtsEnable { get; private set; } = null!;

        /// <summary>
        /// Cluster maintenance window.
        /// </summary>
        [Output("maintainWindows")]
        public Output<ImmutableArray<Outputs.DwsClusterMaintainWindow>> MaintainWindows { get; private set; } = null!;

        /// <summary>
        /// The cluster name.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The subnet ID.
        /// </summary>
        [Output("networkId")]
        public Output<string> NetworkId { get; private set; } = null!;

        /// <summary>
        /// The flavor of the cluster.
        /// </summary>
        [Output("nodeType")]
        public Output<string> NodeType { get; private set; } = null!;

        /// <summary>
        /// schema: Required
        /// </summary>
        [Output("numberOfCn")]
        public Output<int?> NumberOfCn { get; private set; } = null!;

        /// <summary>
        /// Number of nodes in a cluster.
        /// </summary>
        [Output("numberOfNode")]
        public Output<int> NumberOfNode { get; private set; } = null!;

        /// <summary>
        /// Service port of a cluster (8000 to 10000). The default value is 8000.
        /// </summary>
        [Output("port")]
        public Output<int?> Port { get; private set; } = null!;

        /// <summary>
        /// List of private network IP addresses.
        /// </summary>
        [Output("privateIps")]
        public Output<ImmutableArray<string>> PrivateIps { get; private set; } = null!;

        /// <summary>
        /// Public network connection information about the cluster.
        /// </summary>
        [Output("publicEndpoints")]
        public Output<ImmutableArray<Outputs.DwsClusterPublicEndpoint>> PublicEndpoints { get; private set; } = null!;

        [Output("publicIp")]
        public Output<Outputs.DwsClusterPublicIp> PublicIp { get; private set; } = null!;

        [Output("recentEvent")]
        public Output<int> RecentEvent { get; private set; } = null!;

        [Output("region")]
        public Output<string> Region { get; private set; } = null!;

        /// <summary>
        /// The security group ID.
        /// </summary>
        [Output("securityGroupId")]
        public Output<string> SecurityGroupId { get; private set; } = null!;

        /// <summary>
        /// The cluster status.
        /// </summary>
        [Output("status")]
        public Output<string> Status { get; private set; } = null!;

        /// <summary>
        /// Sub-status of clusters in the AVAILABLE state.
        /// </summary>
        [Output("subStatus")]
        public Output<string> SubStatus { get; private set; } = null!;

        /// <summary>
        /// The key/value pairs to associate with the cluster.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;

        /// <summary>
        /// Cluster management task.
        /// </summary>
        [Output("taskStatus")]
        public Output<string> TaskStatus { get; private set; } = null!;

        /// <summary>
        /// The updated time of the cluster.
        /// </summary>
        [Output("updated")]
        public Output<string> Updated { get; private set; } = null!;

        /// <summary>
        /// Administrator username for logging in to a data warehouse cluster.
        /// </summary>
        [Output("userName")]
        public Output<string> UserName { get; private set; } = null!;

        /// <summary>
        /// Administrator password for logging in to a data warehouse cluster.
        /// </summary>
        [Output("userPwd")]
        public Output<string> UserPwd { get; private set; } = null!;

        /// <summary>
        /// schema: Required
        /// </summary>
        [Output("version")]
        public Output<string> Version { get; private set; } = null!;

        [Output("volume")]
        public Output<Outputs.DwsClusterVolume> Volume { get; private set; } = null!;

        /// <summary>
        /// The VPC ID.
        /// </summary>
        [Output("vpcId")]
        public Output<string> VpcId { get; private set; } = null!;


        /// <summary>
        /// Create a DwsCluster resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public DwsCluster(string name, DwsClusterArgs args, CustomResourceOptions? options = null)
            : base("sbercloud:index/dwsCluster:DwsCluster", name, args ?? new DwsClusterArgs(), MakeResourceOptions(options, ""))
        {
        }

        private DwsCluster(string name, Input<string> id, DwsClusterState? state = null, CustomResourceOptions? options = null)
            : base("sbercloud:index/dwsCluster:DwsCluster", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                AdditionalSecretOutputs =
                {
                    "userPwd",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing DwsCluster resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static DwsCluster Get(string name, Input<string> id, DwsClusterState? state = null, CustomResourceOptions? options = null)
        {
            return new DwsCluster(name, id, state, options);
        }
    }

    public sealed class DwsClusterArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The availability zone in which to create the cluster instance.
        /// </summary>
        [Input("availabilityZone", required: true)]
        public Input<string> AvailabilityZone { get; set; } = null!;

        /// <summary>
        /// The description of the cluster.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Dedicated storage pool ID.
        /// </summary>
        [Input("dssPoolId")]
        public Input<string>? DssPoolId { get; set; }

        /// <summary>
        /// The ID of the ELB load balancer.
        /// </summary>
        [Input("elbId")]
        public Input<string>? ElbId { get; set; }

        /// <summary>
        /// The enterprise project ID.
        /// </summary>
        [Input("enterpriseProjectId")]
        public Input<string>? EnterpriseProjectId { get; set; }

        /// <summary>
        /// Whether to automatically execute snapshot when shrinking the number of nodes.
        /// </summary>
        [Input("forceBackup")]
        public Input<bool>? ForceBackup { get; set; }

        /// <summary>
        /// The number of latest manual snapshots that need to be retained when deleting the cluster.
        /// </summary>
        [Input("keepLastManualSnapshot")]
        public Input<int>? KeepLastManualSnapshot { get; set; }

        /// <summary>
        /// The KMS key ID.
        /// </summary>
        [Input("kmsKeyId")]
        public Input<string>? KmsKeyId { get; set; }

        /// <summary>
        /// Whether to enable logical cluster.
        /// </summary>
        [Input("logicalClusterEnable")]
        public Input<bool>? LogicalClusterEnable { get; set; }

        /// <summary>
        /// Whether to enable LTS.
        /// </summary>
        [Input("ltsEnable")]
        public Input<bool>? LtsEnable { get; set; }

        /// <summary>
        /// The cluster name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The subnet ID.
        /// </summary>
        [Input("networkId", required: true)]
        public Input<string> NetworkId { get; set; } = null!;

        /// <summary>
        /// The flavor of the cluster.
        /// </summary>
        [Input("nodeType", required: true)]
        public Input<string> NodeType { get; set; } = null!;

        /// <summary>
        /// schema: Required
        /// </summary>
        [Input("numberOfCn")]
        public Input<int>? NumberOfCn { get; set; }

        /// <summary>
        /// Number of nodes in a cluster.
        /// </summary>
        [Input("numberOfNode", required: true)]
        public Input<int> NumberOfNode { get; set; } = null!;

        /// <summary>
        /// Service port of a cluster (8000 to 10000). The default value is 8000.
        /// </summary>
        [Input("port")]
        public Input<int>? Port { get; set; }

        [Input("publicIp")]
        public Input<Inputs.DwsClusterPublicIpArgs>? PublicIp { get; set; }

        [Input("region")]
        public Input<string>? Region { get; set; }

        /// <summary>
        /// The security group ID.
        /// </summary>
        [Input("securityGroupId", required: true)]
        public Input<string> SecurityGroupId { get; set; } = null!;

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// The key/value pairs to associate with the cluster.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        /// <summary>
        /// Administrator username for logging in to a data warehouse cluster.
        /// </summary>
        [Input("userName", required: true)]
        public Input<string> UserName { get; set; } = null!;

        [Input("userPwd", required: true)]
        private Input<string>? _userPwd;

        /// <summary>
        /// Administrator password for logging in to a data warehouse cluster.
        /// </summary>
        public Input<string>? UserPwd
        {
            get => _userPwd;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _userPwd = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        /// <summary>
        /// schema: Required
        /// </summary>
        [Input("version")]
        public Input<string>? Version { get; set; }

        [Input("volume")]
        public Input<Inputs.DwsClusterVolumeArgs>? Volume { get; set; }

        /// <summary>
        /// The VPC ID.
        /// </summary>
        [Input("vpcId", required: true)]
        public Input<string> VpcId { get; set; } = null!;

        public DwsClusterArgs()
        {
        }
        public static new DwsClusterArgs Empty => new DwsClusterArgs();
    }

    public sealed class DwsClusterState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The availability zone in which to create the cluster instance.
        /// </summary>
        [Input("availabilityZone")]
        public Input<string>? AvailabilityZone { get; set; }

        /// <summary>
        /// The creation time of the cluster.
        /// </summary>
        [Input("created")]
        public Input<string>? Created { get; set; }

        /// <summary>
        /// The description of the cluster.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Dedicated storage pool ID.
        /// </summary>
        [Input("dssPoolId")]
        public Input<string>? DssPoolId { get; set; }

        /// <summary>
        /// The ID of the ELB load balancer.
        /// </summary>
        [Input("elbId")]
        public Input<string>? ElbId { get; set; }

        [Input("elbs")]
        private InputList<Inputs.DwsClusterElbGetArgs>? _elbs;

        /// <summary>
        /// The ELB information bound to the cluster.
        /// </summary>
        public InputList<Inputs.DwsClusterElbGetArgs> Elbs
        {
            get => _elbs ?? (_elbs = new InputList<Inputs.DwsClusterElbGetArgs>());
            set => _elbs = value;
        }

        [Input("endpoints")]
        private InputList<Inputs.DwsClusterEndpointGetArgs>? _endpoints;

        /// <summary>
        /// Private network connection information about the cluster.
        /// </summary>
        public InputList<Inputs.DwsClusterEndpointGetArgs> Endpoints
        {
            get => _endpoints ?? (_endpoints = new InputList<Inputs.DwsClusterEndpointGetArgs>());
            set => _endpoints = value;
        }

        /// <summary>
        /// The enterprise project ID.
        /// </summary>
        [Input("enterpriseProjectId")]
        public Input<string>? EnterpriseProjectId { get; set; }

        /// <summary>
        /// Whether to automatically execute snapshot when shrinking the number of nodes.
        /// </summary>
        [Input("forceBackup")]
        public Input<bool>? ForceBackup { get; set; }

        /// <summary>
        /// The number of latest manual snapshots that need to be retained when deleting the cluster.
        /// </summary>
        [Input("keepLastManualSnapshot")]
        public Input<int>? KeepLastManualSnapshot { get; set; }

        /// <summary>
        /// The KMS key ID.
        /// </summary>
        [Input("kmsKeyId")]
        public Input<string>? KmsKeyId { get; set; }

        /// <summary>
        /// Whether to enable logical cluster.
        /// </summary>
        [Input("logicalClusterEnable")]
        public Input<bool>? LogicalClusterEnable { get; set; }

        /// <summary>
        /// Whether to enable LTS.
        /// </summary>
        [Input("ltsEnable")]
        public Input<bool>? LtsEnable { get; set; }

        [Input("maintainWindows")]
        private InputList<Inputs.DwsClusterMaintainWindowGetArgs>? _maintainWindows;

        /// <summary>
        /// Cluster maintenance window.
        /// </summary>
        public InputList<Inputs.DwsClusterMaintainWindowGetArgs> MaintainWindows
        {
            get => _maintainWindows ?? (_maintainWindows = new InputList<Inputs.DwsClusterMaintainWindowGetArgs>());
            set => _maintainWindows = value;
        }

        /// <summary>
        /// The cluster name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The subnet ID.
        /// </summary>
        [Input("networkId")]
        public Input<string>? NetworkId { get; set; }

        /// <summary>
        /// The flavor of the cluster.
        /// </summary>
        [Input("nodeType")]
        public Input<string>? NodeType { get; set; }

        /// <summary>
        /// schema: Required
        /// </summary>
        [Input("numberOfCn")]
        public Input<int>? NumberOfCn { get; set; }

        /// <summary>
        /// Number of nodes in a cluster.
        /// </summary>
        [Input("numberOfNode")]
        public Input<int>? NumberOfNode { get; set; }

        /// <summary>
        /// Service port of a cluster (8000 to 10000). The default value is 8000.
        /// </summary>
        [Input("port")]
        public Input<int>? Port { get; set; }

        [Input("privateIps")]
        private InputList<string>? _privateIps;

        /// <summary>
        /// List of private network IP addresses.
        /// </summary>
        public InputList<string> PrivateIps
        {
            get => _privateIps ?? (_privateIps = new InputList<string>());
            set => _privateIps = value;
        }

        [Input("publicEndpoints")]
        private InputList<Inputs.DwsClusterPublicEndpointGetArgs>? _publicEndpoints;

        /// <summary>
        /// Public network connection information about the cluster.
        /// </summary>
        public InputList<Inputs.DwsClusterPublicEndpointGetArgs> PublicEndpoints
        {
            get => _publicEndpoints ?? (_publicEndpoints = new InputList<Inputs.DwsClusterPublicEndpointGetArgs>());
            set => _publicEndpoints = value;
        }

        [Input("publicIp")]
        public Input<Inputs.DwsClusterPublicIpGetArgs>? PublicIp { get; set; }

        [Input("recentEvent")]
        public Input<int>? RecentEvent { get; set; }

        [Input("region")]
        public Input<string>? Region { get; set; }

        /// <summary>
        /// The security group ID.
        /// </summary>
        [Input("securityGroupId")]
        public Input<string>? SecurityGroupId { get; set; }

        /// <summary>
        /// The cluster status.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        /// <summary>
        /// Sub-status of clusters in the AVAILABLE state.
        /// </summary>
        [Input("subStatus")]
        public Input<string>? SubStatus { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// The key/value pairs to associate with the cluster.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        /// <summary>
        /// Cluster management task.
        /// </summary>
        [Input("taskStatus")]
        public Input<string>? TaskStatus { get; set; }

        /// <summary>
        /// The updated time of the cluster.
        /// </summary>
        [Input("updated")]
        public Input<string>? Updated { get; set; }

        /// <summary>
        /// Administrator username for logging in to a data warehouse cluster.
        /// </summary>
        [Input("userName")]
        public Input<string>? UserName { get; set; }

        [Input("userPwd")]
        private Input<string>? _userPwd;

        /// <summary>
        /// Administrator password for logging in to a data warehouse cluster.
        /// </summary>
        public Input<string>? UserPwd
        {
            get => _userPwd;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _userPwd = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        /// <summary>
        /// schema: Required
        /// </summary>
        [Input("version")]
        public Input<string>? Version { get; set; }

        [Input("volume")]
        public Input<Inputs.DwsClusterVolumeGetArgs>? Volume { get; set; }

        /// <summary>
        /// The VPC ID.
        /// </summary>
        [Input("vpcId")]
        public Input<string>? VpcId { get; set; }

        public DwsClusterState()
        {
        }
        public static new DwsClusterState Empty => new DwsClusterState();
    }
}
