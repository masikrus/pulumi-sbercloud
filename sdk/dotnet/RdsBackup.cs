// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Sbercloud
{
    [SbercloudResourceType("sbercloud:index/rdsBackup:RdsBackup")]
    public partial class RdsBackup : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Whether a DDM instance has been associated.
        /// </summary>
        [Output("associatedWithDdm")]
        public Output<bool> AssociatedWithDdm { get; private set; } = null!;

        /// <summary>
        /// Backup start time in the "yyyy-mm-ddThh:mm:ssZ" format.
        /// </summary>
        [Output("beginTime")]
        public Output<string> BeginTime { get; private set; } = null!;

        /// <summary>
        /// List of self-built Microsoft SQL Server databases that are partially backed up.
        /// </summary>
        [Output("databases")]
        public Output<ImmutableArray<Outputs.RdsBackupDatabase>> Databases { get; private set; } = null!;

        /// <summary>
        /// The description about the backup.
        /// </summary>
        [Output("description")]
        public Output<string> Description { get; private set; } = null!;

        /// <summary>
        /// Backup end time in the "yyyy-mm-ddThh:mm:ssZ" format.
        /// </summary>
        [Output("endTime")]
        public Output<string> EndTime { get; private set; } = null!;

        /// <summary>
        /// Instance ID.
        /// </summary>
        [Output("instanceId")]
        public Output<string> InstanceId { get; private set; } = null!;

        /// <summary>
        /// Backup name.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        [Output("region")]
        public Output<string> Region { get; private set; } = null!;

        /// <summary>
        /// Backup size in KB.
        /// </summary>
        [Output("size")]
        public Output<int> Size { get; private set; } = null!;

        /// <summary>
        /// Backup status.
        /// </summary>
        [Output("status")]
        public Output<string> Status { get; private set; } = null!;


        /// <summary>
        /// Create a RdsBackup resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public RdsBackup(string name, RdsBackupArgs args, CustomResourceOptions? options = null)
            : base("sbercloud:index/rdsBackup:RdsBackup", name, args ?? new RdsBackupArgs(), MakeResourceOptions(options, ""))
        {
        }

        private RdsBackup(string name, Input<string> id, RdsBackupState? state = null, CustomResourceOptions? options = null)
            : base("sbercloud:index/rdsBackup:RdsBackup", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing RdsBackup resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static RdsBackup Get(string name, Input<string> id, RdsBackupState? state = null, CustomResourceOptions? options = null)
        {
            return new RdsBackup(name, id, state, options);
        }
    }

    public sealed class RdsBackupArgs : global::Pulumi.ResourceArgs
    {
        [Input("databases")]
        private InputList<Inputs.RdsBackupDatabaseArgs>? _databases;

        /// <summary>
        /// List of self-built Microsoft SQL Server databases that are partially backed up.
        /// </summary>
        public InputList<Inputs.RdsBackupDatabaseArgs> Databases
        {
            get => _databases ?? (_databases = new InputList<Inputs.RdsBackupDatabaseArgs>());
            set => _databases = value;
        }

        /// <summary>
        /// The description about the backup.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Instance ID.
        /// </summary>
        [Input("instanceId", required: true)]
        public Input<string> InstanceId { get; set; } = null!;

        /// <summary>
        /// Backup name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("region")]
        public Input<string>? Region { get; set; }

        public RdsBackupArgs()
        {
        }
        public static new RdsBackupArgs Empty => new RdsBackupArgs();
    }

    public sealed class RdsBackupState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Whether a DDM instance has been associated.
        /// </summary>
        [Input("associatedWithDdm")]
        public Input<bool>? AssociatedWithDdm { get; set; }

        /// <summary>
        /// Backup start time in the "yyyy-mm-ddThh:mm:ssZ" format.
        /// </summary>
        [Input("beginTime")]
        public Input<string>? BeginTime { get; set; }

        [Input("databases")]
        private InputList<Inputs.RdsBackupDatabaseGetArgs>? _databases;

        /// <summary>
        /// List of self-built Microsoft SQL Server databases that are partially backed up.
        /// </summary>
        public InputList<Inputs.RdsBackupDatabaseGetArgs> Databases
        {
            get => _databases ?? (_databases = new InputList<Inputs.RdsBackupDatabaseGetArgs>());
            set => _databases = value;
        }

        /// <summary>
        /// The description about the backup.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Backup end time in the "yyyy-mm-ddThh:mm:ssZ" format.
        /// </summary>
        [Input("endTime")]
        public Input<string>? EndTime { get; set; }

        /// <summary>
        /// Instance ID.
        /// </summary>
        [Input("instanceId")]
        public Input<string>? InstanceId { get; set; }

        /// <summary>
        /// Backup name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("region")]
        public Input<string>? Region { get; set; }

        /// <summary>
        /// Backup size in KB.
        /// </summary>
        [Input("size")]
        public Input<int>? Size { get; set; }

        /// <summary>
        /// Backup status.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        public RdsBackupState()
        {
        }
        public static new RdsBackupState Empty => new RdsBackupState();
    }
}
