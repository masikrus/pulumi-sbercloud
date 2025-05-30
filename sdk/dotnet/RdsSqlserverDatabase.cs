// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Sbercloud
{
    [SbercloudResourceType("sbercloud:index/rdsSqlserverDatabase:RdsSqlserverDatabase")]
    public partial class RdsSqlserverDatabase : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Indicates the character set used by the database.
        /// </summary>
        [Output("characterSet")]
        public Output<string> CharacterSet { get; private set; } = null!;

        /// <summary>
        /// Specifies the ID of the RDS SQLServer instance.
        /// </summary>
        [Output("instanceId")]
        public Output<string> InstanceId { get; private set; } = null!;

        /// <summary>
        /// Specifies the database name.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        [Output("region")]
        public Output<string> Region { get; private set; } = null!;

        /// <summary>
        /// Indicates the database status.
        /// </summary>
        [Output("state")]
        public Output<string> State { get; private set; } = null!;


        /// <summary>
        /// Create a RdsSqlserverDatabase resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public RdsSqlserverDatabase(string name, RdsSqlserverDatabaseArgs args, CustomResourceOptions? options = null)
            : base("sbercloud:index/rdsSqlserverDatabase:RdsSqlserverDatabase", name, args ?? new RdsSqlserverDatabaseArgs(), MakeResourceOptions(options, ""))
        {
        }

        private RdsSqlserverDatabase(string name, Input<string> id, RdsSqlserverDatabaseState? state = null, CustomResourceOptions? options = null)
            : base("sbercloud:index/rdsSqlserverDatabase:RdsSqlserverDatabase", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing RdsSqlserverDatabase resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static RdsSqlserverDatabase Get(string name, Input<string> id, RdsSqlserverDatabaseState? state = null, CustomResourceOptions? options = null)
        {
            return new RdsSqlserverDatabase(name, id, state, options);
        }
    }

    public sealed class RdsSqlserverDatabaseArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specifies the ID of the RDS SQLServer instance.
        /// </summary>
        [Input("instanceId", required: true)]
        public Input<string> InstanceId { get; set; } = null!;

        /// <summary>
        /// Specifies the database name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("region")]
        public Input<string>? Region { get; set; }

        public RdsSqlserverDatabaseArgs()
        {
        }
        public static new RdsSqlserverDatabaseArgs Empty => new RdsSqlserverDatabaseArgs();
    }

    public sealed class RdsSqlserverDatabaseState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Indicates the character set used by the database.
        /// </summary>
        [Input("characterSet")]
        public Input<string>? CharacterSet { get; set; }

        /// <summary>
        /// Specifies the ID of the RDS SQLServer instance.
        /// </summary>
        [Input("instanceId")]
        public Input<string>? InstanceId { get; set; }

        /// <summary>
        /// Specifies the database name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("region")]
        public Input<string>? Region { get; set; }

        /// <summary>
        /// Indicates the database status.
        /// </summary>
        [Input("state")]
        public Input<string>? State { get; set; }

        public RdsSqlserverDatabaseState()
        {
        }
        public static new RdsSqlserverDatabaseState Empty => new RdsSqlserverDatabaseState();
    }
}
