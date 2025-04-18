// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Sbercloud
{
    [SbercloudResourceType("sbercloud:index/gesGraph:GesGraph")]
    public partial class GesGraph : global::Pulumi.CustomResource
    {
        /// <summary>
        /// AZ code
        /// </summary>
        [Output("azCode")]
        public Output<string> AzCode { get; private set; } = null!;

        /// <summary>
        /// Graph instance's CPU architecture type.
        /// </summary>
        [Output("cpuArch")]
        public Output<string> CpuArch { get; private set; } = null!;

        /// <summary>
        /// Graph instance cryptography algorithm.
        /// </summary>
        [Output("cryptAlgorithm")]
        public Output<string> CryptAlgorithm { get; private set; } = null!;

        /// <summary>
        /// Whether to enable full-text index control for the created graph.
        /// </summary>
        [Output("enableFullTextIndex")]
        public Output<bool> EnableFullTextIndex { get; private set; } = null!;

        /// <summary>
        /// Whether to enable the security mode. This mode may damage GES performance greatly.
        /// </summary>
        [Output("enableHttps")]
        public Output<bool> EnableHttps { get; private set; } = null!;

        /// <summary>
        /// Whether to enable HyG for the graph. This parameter is available for database edition graphs only.
        /// </summary>
        [Output("enableHyg")]
        public Output<bool> EnableHyg { get; private set; } = null!;

        /// <summary>
        /// Whether the created graph supports the cross-AZ mode. The default value is false.
        /// </summary>
        [Output("enableMultiAz")]
        public Output<bool> EnableMultiAz { get; private set; } = null!;

        /// <summary>
        /// Whether to enable granular permission control for the created graph.
        /// </summary>
        [Output("enableRbac")]
        public Output<bool> EnableRbac { get; private set; } = null!;

        /// <summary>
        /// Whether to enable data encryption The value can be true or false.
        /// </summary>
        [Output("encryption")]
        public Output<Outputs.GesGraphEncryption> Encryption { get; private set; } = null!;

        /// <summary>
        /// The enterprise project ID.
        /// </summary>
        [Output("enterpriseProjectId")]
        public Output<string> EnterpriseProjectId { get; private set; } = null!;

        /// <summary>
        /// Graph size type index.
        /// </summary>
        [Output("graphSizeTypeIndex")]
        public Output<string> GraphSizeTypeIndex { get; private set; } = null!;

        /// <summary>
        /// Whether to retain the backups of a graph after it is deleted.
        /// </summary>
        [Output("keepBackup")]
        public Output<bool> KeepBackup { get; private set; } = null!;

        [Output("ltsOperationTrace")]
        public Output<Outputs.GesGraphLtsOperationTrace> LtsOperationTrace { get; private set; } = null!;

        /// <summary>
        /// The graph name.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Floating IP address of a graph instance.
        /// </summary>
        [Output("privateIp")]
        public Output<string> PrivateIp { get; private set; } = null!;

        /// <summary>
        /// Graph product type
        /// </summary>
        [Output("productType")]
        public Output<string> ProductType { get; private set; } = null!;

        /// <summary>
        /// The information about public IP.
        /// </summary>
        [Output("publicIp")]
        public Output<Outputs.GesGraphPublicIp> PublicIp { get; private set; } = null!;

        [Output("region")]
        public Output<string> Region { get; private set; } = null!;

        [Output("replication")]
        public Output<int> Replication { get; private set; } = null!;

        /// <summary>
        /// The security group ID.
        /// </summary>
        [Output("securityGroupId")]
        public Output<string> SecurityGroupId { get; private set; } = null!;

        /// <summary>
        /// Status of a graph.
        /// </summary>
        [Output("status")]
        public Output<string> Status { get; private set; } = null!;

        /// <summary>
        /// The subnet ID.
        /// </summary>
        [Output("subnetId")]
        public Output<string> SubnetId { get; private set; } = null!;

        /// <summary>
        /// The key/value pairs to associate with the graph.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>> Tags { get; private set; } = null!;

        /// <summary>
        /// Physical addresses of a graph instance for access from private networks.
        /// </summary>
        [Output("trafficIpLists")]
        public Output<ImmutableArray<string>> TrafficIpLists { get; private set; } = null!;

        /// <summary>
        /// ID type of vertices. This parameter is mandatory only for database edition graphs.
        /// </summary>
        [Output("vertexIdType")]
        public Output<Outputs.GesGraphVertexIdType> VertexIdType { get; private set; } = null!;

        /// <summary>
        /// The VPC ID.
        /// </summary>
        [Output("vpcId")]
        public Output<string> VpcId { get; private set; } = null!;


        /// <summary>
        /// Create a GesGraph resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public GesGraph(string name, GesGraphArgs args, CustomResourceOptions? options = null)
            : base("sbercloud:index/gesGraph:GesGraph", name, args ?? new GesGraphArgs(), MakeResourceOptions(options, ""))
        {
        }

        private GesGraph(string name, Input<string> id, GesGraphState? state = null, CustomResourceOptions? options = null)
            : base("sbercloud:index/gesGraph:GesGraph", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing GesGraph resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static GesGraph Get(string name, Input<string> id, GesGraphState? state = null, CustomResourceOptions? options = null)
        {
            return new GesGraph(name, id, state, options);
        }
    }

    public sealed class GesGraphArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Graph instance's CPU architecture type.
        /// </summary>
        [Input("cpuArch")]
        public Input<string>? CpuArch { get; set; }

        /// <summary>
        /// Graph instance cryptography algorithm.
        /// </summary>
        [Input("cryptAlgorithm", required: true)]
        public Input<string> CryptAlgorithm { get; set; } = null!;

        /// <summary>
        /// Whether to enable full-text index control for the created graph.
        /// </summary>
        [Input("enableFullTextIndex")]
        public Input<bool>? EnableFullTextIndex { get; set; }

        /// <summary>
        /// Whether to enable the security mode. This mode may damage GES performance greatly.
        /// </summary>
        [Input("enableHttps", required: true)]
        public Input<bool> EnableHttps { get; set; } = null!;

        /// <summary>
        /// Whether to enable HyG for the graph. This parameter is available for database edition graphs only.
        /// </summary>
        [Input("enableHyg")]
        public Input<bool>? EnableHyg { get; set; }

        /// <summary>
        /// Whether the created graph supports the cross-AZ mode. The default value is false.
        /// </summary>
        [Input("enableMultiAz")]
        public Input<bool>? EnableMultiAz { get; set; }

        /// <summary>
        /// Whether to enable granular permission control for the created graph.
        /// </summary>
        [Input("enableRbac")]
        public Input<bool>? EnableRbac { get; set; }

        /// <summary>
        /// Whether to enable data encryption The value can be true or false.
        /// </summary>
        [Input("encryption")]
        public Input<Inputs.GesGraphEncryptionArgs>? Encryption { get; set; }

        /// <summary>
        /// The enterprise project ID.
        /// </summary>
        [Input("enterpriseProjectId")]
        public Input<string>? EnterpriseProjectId { get; set; }

        /// <summary>
        /// Graph size type index.
        /// </summary>
        [Input("graphSizeTypeIndex", required: true)]
        public Input<string> GraphSizeTypeIndex { get; set; } = null!;

        /// <summary>
        /// Whether to retain the backups of a graph after it is deleted.
        /// </summary>
        [Input("keepBackup")]
        public Input<bool>? KeepBackup { get; set; }

        [Input("ltsOperationTrace")]
        public Input<Inputs.GesGraphLtsOperationTraceArgs>? LtsOperationTrace { get; set; }

        /// <summary>
        /// The graph name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Graph product type
        /// </summary>
        [Input("productType")]
        public Input<string>? ProductType { get; set; }

        /// <summary>
        /// The information about public IP.
        /// </summary>
        [Input("publicIp")]
        public Input<Inputs.GesGraphPublicIpArgs>? PublicIp { get; set; }

        [Input("region")]
        public Input<string>? Region { get; set; }

        [Input("replication")]
        public Input<int>? Replication { get; set; }

        /// <summary>
        /// The security group ID.
        /// </summary>
        [Input("securityGroupId", required: true)]
        public Input<string> SecurityGroupId { get; set; } = null!;

        /// <summary>
        /// The subnet ID.
        /// </summary>
        [Input("subnetId", required: true)]
        public Input<string> SubnetId { get; set; } = null!;

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// The key/value pairs to associate with the graph.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        /// <summary>
        /// ID type of vertices. This parameter is mandatory only for database edition graphs.
        /// </summary>
        [Input("vertexIdType")]
        public Input<Inputs.GesGraphVertexIdTypeArgs>? VertexIdType { get; set; }

        /// <summary>
        /// The VPC ID.
        /// </summary>
        [Input("vpcId", required: true)]
        public Input<string> VpcId { get; set; } = null!;

        public GesGraphArgs()
        {
        }
        public static new GesGraphArgs Empty => new GesGraphArgs();
    }

    public sealed class GesGraphState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// AZ code
        /// </summary>
        [Input("azCode")]
        public Input<string>? AzCode { get; set; }

        /// <summary>
        /// Graph instance's CPU architecture type.
        /// </summary>
        [Input("cpuArch")]
        public Input<string>? CpuArch { get; set; }

        /// <summary>
        /// Graph instance cryptography algorithm.
        /// </summary>
        [Input("cryptAlgorithm")]
        public Input<string>? CryptAlgorithm { get; set; }

        /// <summary>
        /// Whether to enable full-text index control for the created graph.
        /// </summary>
        [Input("enableFullTextIndex")]
        public Input<bool>? EnableFullTextIndex { get; set; }

        /// <summary>
        /// Whether to enable the security mode. This mode may damage GES performance greatly.
        /// </summary>
        [Input("enableHttps")]
        public Input<bool>? EnableHttps { get; set; }

        /// <summary>
        /// Whether to enable HyG for the graph. This parameter is available for database edition graphs only.
        /// </summary>
        [Input("enableHyg")]
        public Input<bool>? EnableHyg { get; set; }

        /// <summary>
        /// Whether the created graph supports the cross-AZ mode. The default value is false.
        /// </summary>
        [Input("enableMultiAz")]
        public Input<bool>? EnableMultiAz { get; set; }

        /// <summary>
        /// Whether to enable granular permission control for the created graph.
        /// </summary>
        [Input("enableRbac")]
        public Input<bool>? EnableRbac { get; set; }

        /// <summary>
        /// Whether to enable data encryption The value can be true or false.
        /// </summary>
        [Input("encryption")]
        public Input<Inputs.GesGraphEncryptionGetArgs>? Encryption { get; set; }

        /// <summary>
        /// The enterprise project ID.
        /// </summary>
        [Input("enterpriseProjectId")]
        public Input<string>? EnterpriseProjectId { get; set; }

        /// <summary>
        /// Graph size type index.
        /// </summary>
        [Input("graphSizeTypeIndex")]
        public Input<string>? GraphSizeTypeIndex { get; set; }

        /// <summary>
        /// Whether to retain the backups of a graph after it is deleted.
        /// </summary>
        [Input("keepBackup")]
        public Input<bool>? KeepBackup { get; set; }

        [Input("ltsOperationTrace")]
        public Input<Inputs.GesGraphLtsOperationTraceGetArgs>? LtsOperationTrace { get; set; }

        /// <summary>
        /// The graph name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Floating IP address of a graph instance.
        /// </summary>
        [Input("privateIp")]
        public Input<string>? PrivateIp { get; set; }

        /// <summary>
        /// Graph product type
        /// </summary>
        [Input("productType")]
        public Input<string>? ProductType { get; set; }

        /// <summary>
        /// The information about public IP.
        /// </summary>
        [Input("publicIp")]
        public Input<Inputs.GesGraphPublicIpGetArgs>? PublicIp { get; set; }

        [Input("region")]
        public Input<string>? Region { get; set; }

        [Input("replication")]
        public Input<int>? Replication { get; set; }

        /// <summary>
        /// The security group ID.
        /// </summary>
        [Input("securityGroupId")]
        public Input<string>? SecurityGroupId { get; set; }

        /// <summary>
        /// Status of a graph.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        /// <summary>
        /// The subnet ID.
        /// </summary>
        [Input("subnetId")]
        public Input<string>? SubnetId { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// The key/value pairs to associate with the graph.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        [Input("trafficIpLists")]
        private InputList<string>? _trafficIpLists;

        /// <summary>
        /// Physical addresses of a graph instance for access from private networks.
        /// </summary>
        public InputList<string> TrafficIpLists
        {
            get => _trafficIpLists ?? (_trafficIpLists = new InputList<string>());
            set => _trafficIpLists = value;
        }

        /// <summary>
        /// ID type of vertices. This parameter is mandatory only for database edition graphs.
        /// </summary>
        [Input("vertexIdType")]
        public Input<Inputs.GesGraphVertexIdTypeGetArgs>? VertexIdType { get; set; }

        /// <summary>
        /// The VPC ID.
        /// </summary>
        [Input("vpcId")]
        public Input<string>? VpcId { get; set; }

        public GesGraphState()
        {
        }
        public static new GesGraphState Empty => new GesGraphState();
    }
}
