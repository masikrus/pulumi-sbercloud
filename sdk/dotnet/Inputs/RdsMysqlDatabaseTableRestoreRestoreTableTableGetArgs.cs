// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Sbercloud.Inputs
{

    public sealed class RdsMysqlDatabaseTableRestoreRestoreTableTableGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specifies the name of the table after restoration.
        /// </summary>
        [Input("newName", required: true)]
        public Input<string> NewName { get; set; } = null!;

        /// <summary>
        /// Specifies the name of the table before restoration.
        /// </summary>
        [Input("oldName", required: true)]
        public Input<string> OldName { get; set; } = null!;

        public RdsMysqlDatabaseTableRestoreRestoreTableTableGetArgs()
        {
        }
        public static new RdsMysqlDatabaseTableRestoreRestoreTableTableGetArgs Empty => new RdsMysqlDatabaseTableRestoreRestoreTableTableGetArgs();
    }
}
