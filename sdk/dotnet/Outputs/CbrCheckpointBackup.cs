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
    public sealed class CbrCheckpointBackup
    {
        /// <summary>
        /// The backup ID.
        /// </summary>
        public readonly string? Id;
        /// <summary>
        /// The backup time.
        /// </summary>
        public readonly string? ProtectedAt;
        /// <summary>
        /// The ID of backup resource.
        /// </summary>
        public readonly string ResourceId;
        /// <summary>
        /// The backup resource size.
        /// </summary>
        public readonly int? ResourceSize;
        /// <summary>
        /// The backup status.
        /// </summary>
        public readonly string? Status;
        /// <summary>
        /// The type of the backup resource.
        /// </summary>
        public readonly string Type;
        /// <summary>
        /// The latest update time of the backup.
        /// </summary>
        public readonly string? UpdatedAt;

        [OutputConstructor]
        private CbrCheckpointBackup(
            string? id,

            string? protectedAt,

            string resourceId,

            int? resourceSize,

            string? status,

            string type,

            string? updatedAt)
        {
            Id = id;
            ProtectedAt = protectedAt;
            ResourceId = resourceId;
            ResourceSize = resourceSize;
            Status = status;
            Type = type;
            UpdatedAt = updatedAt;
        }
    }
}
