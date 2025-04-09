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
    public sealed class CbrVaultResource
    {
        /// <summary>
        /// The array of disk IDs which will be excluded in the backup.
        /// </summary>
        public readonly ImmutableArray<string> Excludes;
        /// <summary>
        /// The array of disk or SFS file systems which will be included in the backup.
        /// </summary>
        public readonly ImmutableArray<string> Includes;
        /// <summary>
        /// The ID of the ECS instance to be backed up.
        /// </summary>
        public readonly string? ServerId;

        [OutputConstructor]
        private CbrVaultResource(
            ImmutableArray<string> excludes,

            ImmutableArray<string> includes,

            string? serverId)
        {
            Excludes = excludes;
            Includes = includes;
            ServerId = serverId;
        }
    }
}
