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
    public sealed class RdsMysqlDatabaseTableRestoreRestoreTableTable
    {
        /// <summary>
        /// Specifies the name of the table after restoration.
        /// </summary>
        public readonly string NewName;
        /// <summary>
        /// Specifies the name of the table before restoration.
        /// </summary>
        public readonly string OldName;

        [OutputConstructor]
        private RdsMysqlDatabaseTableRestoreRestoreTableTable(
            string newName,

            string oldName)
        {
            NewName = newName;
            OldName = oldName;
        }
    }
}
