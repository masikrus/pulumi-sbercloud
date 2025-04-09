// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Sbercloud
{
    [SbercloudResourceType("sbercloud:index/rdsInstance:RdsInstance")]
    public partial class RdsInstance : global::Pulumi.CustomResource
    {
        [Output("autoPay")]
        public Output<string?> AutoPay { get; private set; } = null!;

        [Output("autoRenew")]
        public Output<string?> AutoRenew { get; private set; } = null!;

        [Output("availabilityZones")]
        public Output<ImmutableArray<string>> AvailabilityZones { get; private set; } = null!;

        [Output("backupStrategy")]
        public Output<Outputs.RdsInstanceBackupStrategy> BackupStrategy { get; private set; } = null!;

        [Output("binlogRetentionHours")]
        public Output<int?> BinlogRetentionHours { get; private set; } = null!;

        [Output("chargingMode")]
        public Output<string> ChargingMode { get; private set; } = null!;

        [Output("collation")]
        public Output<string> Collation { get; private set; } = null!;

        [Output("created")]
        public Output<string> Created { get; private set; } = null!;

        [Output("db")]
        public Output<Outputs.RdsInstanceDb> Db { get; private set; } = null!;

        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        [Output("dssPoolId")]
        public Output<string?> DssPoolId { get; private set; } = null!;

        [Output("enterpriseProjectId")]
        public Output<string> EnterpriseProjectId { get; private set; } = null!;

        [Output("fixedIp")]
        public Output<string> FixedIp { get; private set; } = null!;

        [Output("flavor")]
        public Output<string> Flavor { get; private set; } = null!;

        [Output("haReplicationMode")]
        public Output<string> HaReplicationMode { get; private set; } = null!;

        [Output("lowerCaseTableNames")]
        public Output<string?> LowerCaseTableNames { get; private set; } = null!;

        [Output("maintainBegin")]
        public Output<string> MaintainBegin { get; private set; } = null!;

        [Output("maintainEnd")]
        public Output<string> MaintainEnd { get; private set; } = null!;

        [Output("msdtcHosts")]
        public Output<ImmutableArray<Outputs.RdsInstanceMsdtcHost>> MsdtcHosts { get; private set; } = null!;

        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        [Output("nodes")]
        public Output<ImmutableArray<Outputs.RdsInstanceNode>> Nodes { get; private set; } = null!;

        [Output("paramGroupId")]
        public Output<string?> ParamGroupId { get; private set; } = null!;

        [Output("parameters")]
        public Output<ImmutableArray<Outputs.RdsInstanceParameter>> Parameters { get; private set; } = null!;

        [Output("period")]
        public Output<int?> Period { get; private set; } = null!;

        [Output("periodUnit")]
        public Output<string?> PeriodUnit { get; private set; } = null!;

        [Output("powerAction")]
        public Output<string> PowerAction { get; private set; } = null!;

        [Output("privateDnsNamePrefix")]
        public Output<string> PrivateDnsNamePrefix { get; private set; } = null!;

        [Output("privateDnsNames")]
        public Output<ImmutableArray<string>> PrivateDnsNames { get; private set; } = null!;

        [Output("privateIps")]
        public Output<ImmutableArray<string>> PrivateIps { get; private set; } = null!;

        [Output("publicIps")]
        public Output<ImmutableArray<string>> PublicIps { get; private set; } = null!;

        [Output("readWritePermissions")]
        public Output<string?> ReadWritePermissions { get; private set; } = null!;

        [Output("region")]
        public Output<string> Region { get; private set; } = null!;

        [Output("restore")]
        public Output<Outputs.RdsInstanceRestore?> Restore { get; private set; } = null!;

        [Output("rotateDay")]
        public Output<int?> RotateDay { get; private set; } = null!;

        [Output("secondsLevelMonitoringEnabled")]
        public Output<bool> SecondsLevelMonitoringEnabled { get; private set; } = null!;

        [Output("secondsLevelMonitoringInterval")]
        public Output<int> SecondsLevelMonitoringInterval { get; private set; } = null!;

        [Output("secretId")]
        public Output<string?> SecretId { get; private set; } = null!;

        [Output("secretName")]
        public Output<string?> SecretName { get; private set; } = null!;

        [Output("secretVersion")]
        public Output<string?> SecretVersion { get; private set; } = null!;

        [Output("securityGroupId")]
        public Output<string> SecurityGroupId { get; private set; } = null!;

        [Output("slowLogShowOriginalStatus")]
        public Output<string?> SlowLogShowOriginalStatus { get; private set; } = null!;

        [Output("sslEnable")]
        public Output<bool> SslEnable { get; private set; } = null!;

        [Output("status")]
        public Output<string> Status { get; private set; } = null!;

        [Output("subnetId")]
        public Output<string> SubnetId { get; private set; } = null!;

        [Output("switchStrategy")]
        public Output<string> SwitchStrategy { get; private set; } = null!;

        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;

        [Output("tdeEnabled")]
        public Output<bool> TdeEnabled { get; private set; } = null!;

        [Output("timeZone")]
        public Output<string> TimeZone { get; private set; } = null!;

        [Output("volume")]
        public Output<Outputs.RdsInstanceVolume> Volume { get; private set; } = null!;

        [Output("vpcId")]
        public Output<string> VpcId { get; private set; } = null!;


        /// <summary>
        /// Create a RdsInstance resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public RdsInstance(string name, RdsInstanceArgs args, CustomResourceOptions? options = null)
            : base("sbercloud:index/rdsInstance:RdsInstance", name, args ?? new RdsInstanceArgs(), MakeResourceOptions(options, ""))
        {
        }

        private RdsInstance(string name, Input<string> id, RdsInstanceState? state = null, CustomResourceOptions? options = null)
            : base("sbercloud:index/rdsInstance:RdsInstance", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing RdsInstance resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static RdsInstance Get(string name, Input<string> id, RdsInstanceState? state = null, CustomResourceOptions? options = null)
        {
            return new RdsInstance(name, id, state, options);
        }
    }

    public sealed class RdsInstanceArgs : global::Pulumi.ResourceArgs
    {
        [Input("autoPay")]
        public Input<string>? AutoPay { get; set; }

        [Input("autoRenew")]
        public Input<string>? AutoRenew { get; set; }

        [Input("availabilityZones", required: true)]
        private InputList<string>? _availabilityZones;
        public InputList<string> AvailabilityZones
        {
            get => _availabilityZones ?? (_availabilityZones = new InputList<string>());
            set => _availabilityZones = value;
        }

        [Input("backupStrategy")]
        public Input<Inputs.RdsInstanceBackupStrategyArgs>? BackupStrategy { get; set; }

        [Input("binlogRetentionHours")]
        public Input<int>? BinlogRetentionHours { get; set; }

        [Input("chargingMode")]
        public Input<string>? ChargingMode { get; set; }

        [Input("collation")]
        public Input<string>? Collation { get; set; }

        [Input("db", required: true)]
        public Input<Inputs.RdsInstanceDbArgs> Db { get; set; } = null!;

        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("dssPoolId")]
        public Input<string>? DssPoolId { get; set; }

        [Input("enterpriseProjectId")]
        public Input<string>? EnterpriseProjectId { get; set; }

        [Input("fixedIp")]
        public Input<string>? FixedIp { get; set; }

        [Input("flavor", required: true)]
        public Input<string> Flavor { get; set; } = null!;

        [Input("haReplicationMode")]
        public Input<string>? HaReplicationMode { get; set; }

        [Input("lowerCaseTableNames")]
        public Input<string>? LowerCaseTableNames { get; set; }

        [Input("maintainBegin")]
        public Input<string>? MaintainBegin { get; set; }

        [Input("maintainEnd")]
        public Input<string>? MaintainEnd { get; set; }

        [Input("msdtcHosts")]
        private InputList<Inputs.RdsInstanceMsdtcHostArgs>? _msdtcHosts;
        public InputList<Inputs.RdsInstanceMsdtcHostArgs> MsdtcHosts
        {
            get => _msdtcHosts ?? (_msdtcHosts = new InputList<Inputs.RdsInstanceMsdtcHostArgs>());
            set => _msdtcHosts = value;
        }

        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("paramGroupId")]
        public Input<string>? ParamGroupId { get; set; }

        [Input("parameters")]
        private InputList<Inputs.RdsInstanceParameterArgs>? _parameters;
        public InputList<Inputs.RdsInstanceParameterArgs> Parameters
        {
            get => _parameters ?? (_parameters = new InputList<Inputs.RdsInstanceParameterArgs>());
            set => _parameters = value;
        }

        [Input("period")]
        public Input<int>? Period { get; set; }

        [Input("periodUnit")]
        public Input<string>? PeriodUnit { get; set; }

        [Input("powerAction")]
        public Input<string>? PowerAction { get; set; }

        [Input("privateDnsNamePrefix")]
        public Input<string>? PrivateDnsNamePrefix { get; set; }

        [Input("readWritePermissions")]
        public Input<string>? ReadWritePermissions { get; set; }

        [Input("region")]
        public Input<string>? Region { get; set; }

        [Input("restore")]
        public Input<Inputs.RdsInstanceRestoreArgs>? Restore { get; set; }

        [Input("rotateDay")]
        public Input<int>? RotateDay { get; set; }

        [Input("secondsLevelMonitoringEnabled")]
        public Input<bool>? SecondsLevelMonitoringEnabled { get; set; }

        [Input("secondsLevelMonitoringInterval")]
        public Input<int>? SecondsLevelMonitoringInterval { get; set; }

        [Input("secretId")]
        public Input<string>? SecretId { get; set; }

        [Input("secretName")]
        public Input<string>? SecretName { get; set; }

        [Input("secretVersion")]
        public Input<string>? SecretVersion { get; set; }

        [Input("securityGroupId", required: true)]
        public Input<string> SecurityGroupId { get; set; } = null!;

        [Input("slowLogShowOriginalStatus")]
        public Input<string>? SlowLogShowOriginalStatus { get; set; }

        [Input("sslEnable")]
        public Input<bool>? SslEnable { get; set; }

        [Input("subnetId", required: true)]
        public Input<string> SubnetId { get; set; } = null!;

        [Input("switchStrategy")]
        public Input<string>? SwitchStrategy { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        [Input("tdeEnabled")]
        public Input<bool>? TdeEnabled { get; set; }

        [Input("timeZone")]
        public Input<string>? TimeZone { get; set; }

        [Input("volume", required: true)]
        public Input<Inputs.RdsInstanceVolumeArgs> Volume { get; set; } = null!;

        [Input("vpcId", required: true)]
        public Input<string> VpcId { get; set; } = null!;

        public RdsInstanceArgs()
        {
        }
        public static new RdsInstanceArgs Empty => new RdsInstanceArgs();
    }

    public sealed class RdsInstanceState : global::Pulumi.ResourceArgs
    {
        [Input("autoPay")]
        public Input<string>? AutoPay { get; set; }

        [Input("autoRenew")]
        public Input<string>? AutoRenew { get; set; }

        [Input("availabilityZones")]
        private InputList<string>? _availabilityZones;
        public InputList<string> AvailabilityZones
        {
            get => _availabilityZones ?? (_availabilityZones = new InputList<string>());
            set => _availabilityZones = value;
        }

        [Input("backupStrategy")]
        public Input<Inputs.RdsInstanceBackupStrategyGetArgs>? BackupStrategy { get; set; }

        [Input("binlogRetentionHours")]
        public Input<int>? BinlogRetentionHours { get; set; }

        [Input("chargingMode")]
        public Input<string>? ChargingMode { get; set; }

        [Input("collation")]
        public Input<string>? Collation { get; set; }

        [Input("created")]
        public Input<string>? Created { get; set; }

        [Input("db")]
        public Input<Inputs.RdsInstanceDbGetArgs>? Db { get; set; }

        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("dssPoolId")]
        public Input<string>? DssPoolId { get; set; }

        [Input("enterpriseProjectId")]
        public Input<string>? EnterpriseProjectId { get; set; }

        [Input("fixedIp")]
        public Input<string>? FixedIp { get; set; }

        [Input("flavor")]
        public Input<string>? Flavor { get; set; }

        [Input("haReplicationMode")]
        public Input<string>? HaReplicationMode { get; set; }

        [Input("lowerCaseTableNames")]
        public Input<string>? LowerCaseTableNames { get; set; }

        [Input("maintainBegin")]
        public Input<string>? MaintainBegin { get; set; }

        [Input("maintainEnd")]
        public Input<string>? MaintainEnd { get; set; }

        [Input("msdtcHosts")]
        private InputList<Inputs.RdsInstanceMsdtcHostGetArgs>? _msdtcHosts;
        public InputList<Inputs.RdsInstanceMsdtcHostGetArgs> MsdtcHosts
        {
            get => _msdtcHosts ?? (_msdtcHosts = new InputList<Inputs.RdsInstanceMsdtcHostGetArgs>());
            set => _msdtcHosts = value;
        }

        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("nodes")]
        private InputList<Inputs.RdsInstanceNodeGetArgs>? _nodes;
        public InputList<Inputs.RdsInstanceNodeGetArgs> Nodes
        {
            get => _nodes ?? (_nodes = new InputList<Inputs.RdsInstanceNodeGetArgs>());
            set => _nodes = value;
        }

        [Input("paramGroupId")]
        public Input<string>? ParamGroupId { get; set; }

        [Input("parameters")]
        private InputList<Inputs.RdsInstanceParameterGetArgs>? _parameters;
        public InputList<Inputs.RdsInstanceParameterGetArgs> Parameters
        {
            get => _parameters ?? (_parameters = new InputList<Inputs.RdsInstanceParameterGetArgs>());
            set => _parameters = value;
        }

        [Input("period")]
        public Input<int>? Period { get; set; }

        [Input("periodUnit")]
        public Input<string>? PeriodUnit { get; set; }

        [Input("powerAction")]
        public Input<string>? PowerAction { get; set; }

        [Input("privateDnsNamePrefix")]
        public Input<string>? PrivateDnsNamePrefix { get; set; }

        [Input("privateDnsNames")]
        private InputList<string>? _privateDnsNames;
        public InputList<string> PrivateDnsNames
        {
            get => _privateDnsNames ?? (_privateDnsNames = new InputList<string>());
            set => _privateDnsNames = value;
        }

        [Input("privateIps")]
        private InputList<string>? _privateIps;
        public InputList<string> PrivateIps
        {
            get => _privateIps ?? (_privateIps = new InputList<string>());
            set => _privateIps = value;
        }

        [Input("publicIps")]
        private InputList<string>? _publicIps;
        public InputList<string> PublicIps
        {
            get => _publicIps ?? (_publicIps = new InputList<string>());
            set => _publicIps = value;
        }

        [Input("readWritePermissions")]
        public Input<string>? ReadWritePermissions { get; set; }

        [Input("region")]
        public Input<string>? Region { get; set; }

        [Input("restore")]
        public Input<Inputs.RdsInstanceRestoreGetArgs>? Restore { get; set; }

        [Input("rotateDay")]
        public Input<int>? RotateDay { get; set; }

        [Input("secondsLevelMonitoringEnabled")]
        public Input<bool>? SecondsLevelMonitoringEnabled { get; set; }

        [Input("secondsLevelMonitoringInterval")]
        public Input<int>? SecondsLevelMonitoringInterval { get; set; }

        [Input("secretId")]
        public Input<string>? SecretId { get; set; }

        [Input("secretName")]
        public Input<string>? SecretName { get; set; }

        [Input("secretVersion")]
        public Input<string>? SecretVersion { get; set; }

        [Input("securityGroupId")]
        public Input<string>? SecurityGroupId { get; set; }

        [Input("slowLogShowOriginalStatus")]
        public Input<string>? SlowLogShowOriginalStatus { get; set; }

        [Input("sslEnable")]
        public Input<bool>? SslEnable { get; set; }

        [Input("status")]
        public Input<string>? Status { get; set; }

        [Input("subnetId")]
        public Input<string>? SubnetId { get; set; }

        [Input("switchStrategy")]
        public Input<string>? SwitchStrategy { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        [Input("tdeEnabled")]
        public Input<bool>? TdeEnabled { get; set; }

        [Input("timeZone")]
        public Input<string>? TimeZone { get; set; }

        [Input("volume")]
        public Input<Inputs.RdsInstanceVolumeGetArgs>? Volume { get; set; }

        [Input("vpcId")]
        public Input<string>? VpcId { get; set; }

        public RdsInstanceState()
        {
        }
        public static new RdsInstanceState Empty => new RdsInstanceState();
    }
}
