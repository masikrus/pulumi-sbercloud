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
    public sealed class RdsInstanceRestore
    {
        public readonly string BackupId;
        public readonly ImmutableDictionary<string, string>? DatabaseName;
        public readonly string InstanceId;

        [OutputConstructor]
        private RdsInstanceRestore(
            string backupId,

            ImmutableDictionary<string, string>? databaseName,

            string instanceId)
        {
            BackupId = backupId;
            DatabaseName = databaseName;
            InstanceId = instanceId;
        }
    }
}
