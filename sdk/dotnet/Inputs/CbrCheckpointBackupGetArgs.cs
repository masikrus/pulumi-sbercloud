// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Sbercloud.Inputs
{

    public sealed class CbrCheckpointBackupGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The backup ID.
        /// </summary>
        [Input("id")]
        public Input<string>? Id { get; set; }

        /// <summary>
        /// The backup time.
        /// </summary>
        [Input("protectedAt")]
        public Input<string>? ProtectedAt { get; set; }

        /// <summary>
        /// The ID of backup resource.
        /// </summary>
        [Input("resourceId", required: true)]
        public Input<string> ResourceId { get; set; } = null!;

        /// <summary>
        /// The backup resource size.
        /// </summary>
        [Input("resourceSize")]
        public Input<int>? ResourceSize { get; set; }

        /// <summary>
        /// The backup status.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        /// <summary>
        /// The type of the backup resource.
        /// </summary>
        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        /// <summary>
        /// The latest update time of the backup.
        /// </summary>
        [Input("updatedAt")]
        public Input<string>? UpdatedAt { get; set; }

        public CbrCheckpointBackupGetArgs()
        {
        }
        public static new CbrCheckpointBackupGetArgs Empty => new CbrCheckpointBackupGetArgs();
    }
}
