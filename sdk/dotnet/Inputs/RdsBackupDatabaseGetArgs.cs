// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Sbercloud.Inputs
{

    public sealed class RdsBackupDatabaseGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Database to be backed up for Microsoft SQL Server.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        public RdsBackupDatabaseGetArgs()
        {
        }
        public static new RdsBackupDatabaseGetArgs Empty => new RdsBackupDatabaseGetArgs();
    }
}
