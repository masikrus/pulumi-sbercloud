// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Sbercloud.Inputs
{

    public sealed class RdsMysqlDatabaseTableRestoreRestoreTableArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specifies the database name.
        /// </summary>
        [Input("database", required: true)]
        public Input<string> Database { get; set; } = null!;

        [Input("tables", required: true)]
        private InputList<Inputs.RdsMysqlDatabaseTableRestoreRestoreTableTableArgs>? _tables;

        /// <summary>
        /// Specifies the tables.
        /// </summary>
        public InputList<Inputs.RdsMysqlDatabaseTableRestoreRestoreTableTableArgs> Tables
        {
            get => _tables ?? (_tables = new InputList<Inputs.RdsMysqlDatabaseTableRestoreRestoreTableTableArgs>());
            set => _tables = value;
        }

        public RdsMysqlDatabaseTableRestoreRestoreTableArgs()
        {
        }
        public static new RdsMysqlDatabaseTableRestoreRestoreTableArgs Empty => new RdsMysqlDatabaseTableRestoreRestoreTableArgs();
    }
}
