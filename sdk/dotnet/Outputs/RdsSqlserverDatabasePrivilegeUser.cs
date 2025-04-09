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
    public sealed class RdsSqlserverDatabasePrivilegeUser
    {
        /// <summary>
        /// Specifies the username of the database account.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Specifies the read-only permission.
        /// </summary>
        public readonly bool? Readonly;

        [OutputConstructor]
        private RdsSqlserverDatabasePrivilegeUser(
            string name,

            bool? @readonly)
        {
            Name = name;
            Readonly = @readonly;
        }
    }
}
