// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Sbercloud
{
    [SbercloudResourceType("sbercloud:index/rdsSqlserverDatabasePrivilege:RdsSqlserverDatabasePrivilege")]
    public partial class RdsSqlserverDatabasePrivilege : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Specifies the database name.
        /// </summary>
        [Output("dbName")]
        public Output<string> DbName { get; private set; } = null!;

        /// <summary>
        /// Specifies the ID of the RDS SQL Server instance.
        /// </summary>
        [Output("instanceId")]
        public Output<string> InstanceId { get; private set; } = null!;

        [Output("region")]
        public Output<string> Region { get; private set; } = null!;

        /// <summary>
        /// Specifies the account that associated with the database
        /// </summary>
        [Output("users")]
        public Output<ImmutableArray<Outputs.RdsSqlserverDatabasePrivilegeUser>> Users { get; private set; } = null!;


        /// <summary>
        /// Create a RdsSqlserverDatabasePrivilege resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public RdsSqlserverDatabasePrivilege(string name, RdsSqlserverDatabasePrivilegeArgs args, CustomResourceOptions? options = null)
            : base("sbercloud:index/rdsSqlserverDatabasePrivilege:RdsSqlserverDatabasePrivilege", name, args ?? new RdsSqlserverDatabasePrivilegeArgs(), MakeResourceOptions(options, ""))
        {
        }

        private RdsSqlserverDatabasePrivilege(string name, Input<string> id, RdsSqlserverDatabasePrivilegeState? state = null, CustomResourceOptions? options = null)
            : base("sbercloud:index/rdsSqlserverDatabasePrivilege:RdsSqlserverDatabasePrivilege", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing RdsSqlserverDatabasePrivilege resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static RdsSqlserverDatabasePrivilege Get(string name, Input<string> id, RdsSqlserverDatabasePrivilegeState? state = null, CustomResourceOptions? options = null)
        {
            return new RdsSqlserverDatabasePrivilege(name, id, state, options);
        }
    }

    public sealed class RdsSqlserverDatabasePrivilegeArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specifies the database name.
        /// </summary>
        [Input("dbName", required: true)]
        public Input<string> DbName { get; set; } = null!;

        /// <summary>
        /// Specifies the ID of the RDS SQL Server instance.
        /// </summary>
        [Input("instanceId", required: true)]
        public Input<string> InstanceId { get; set; } = null!;

        [Input("region")]
        public Input<string>? Region { get; set; }

        [Input("users", required: true)]
        private InputList<Inputs.RdsSqlserverDatabasePrivilegeUserArgs>? _users;

        /// <summary>
        /// Specifies the account that associated with the database
        /// </summary>
        public InputList<Inputs.RdsSqlserverDatabasePrivilegeUserArgs> Users
        {
            get => _users ?? (_users = new InputList<Inputs.RdsSqlserverDatabasePrivilegeUserArgs>());
            set => _users = value;
        }

        public RdsSqlserverDatabasePrivilegeArgs()
        {
        }
        public static new RdsSqlserverDatabasePrivilegeArgs Empty => new RdsSqlserverDatabasePrivilegeArgs();
    }

    public sealed class RdsSqlserverDatabasePrivilegeState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specifies the database name.
        /// </summary>
        [Input("dbName")]
        public Input<string>? DbName { get; set; }

        /// <summary>
        /// Specifies the ID of the RDS SQL Server instance.
        /// </summary>
        [Input("instanceId")]
        public Input<string>? InstanceId { get; set; }

        [Input("region")]
        public Input<string>? Region { get; set; }

        [Input("users")]
        private InputList<Inputs.RdsSqlserverDatabasePrivilegeUserGetArgs>? _users;

        /// <summary>
        /// Specifies the account that associated with the database
        /// </summary>
        public InputList<Inputs.RdsSqlserverDatabasePrivilegeUserGetArgs> Users
        {
            get => _users ?? (_users = new InputList<Inputs.RdsSqlserverDatabasePrivilegeUserGetArgs>());
            set => _users = value;
        }

        public RdsSqlserverDatabasePrivilegeState()
        {
        }
        public static new RdsSqlserverDatabasePrivilegeState Empty => new RdsSqlserverDatabasePrivilegeState();
    }
}
