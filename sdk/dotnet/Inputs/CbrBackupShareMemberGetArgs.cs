// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Sbercloud.Inputs
{

    public sealed class CbrBackupShareMemberGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The creation time of the backup shared member.
        /// </summary>
        [Input("createdAt")]
        public Input<string>? CreatedAt { get; set; }

        /// <summary>
        /// The ID of the project with which the backup is shared.
        /// </summary>
        [Input("destProjectId", required: true)]
        public Input<string> DestProjectId { get; set; } = null!;

        /// <summary>
        /// The ID of the backup shared member record.
        /// </summary>
        [Input("id")]
        public Input<string>? Id { get; set; }

        /// <summary>
        /// The ID of the image registered with the shared backup copy.
        /// </summary>
        [Input("imageId")]
        public Input<string>? ImageId { get; set; }

        /// <summary>
        /// The backup shared status.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        /// <summary>
        /// The latest update time of the backup shared member.
        /// </summary>
        [Input("updatedAt")]
        public Input<string>? UpdatedAt { get; set; }

        /// <summary>
        /// The ID of the vault where the shared backup is stored.
        /// </summary>
        [Input("vaultId")]
        public Input<string>? VaultId { get; set; }

        public CbrBackupShareMemberGetArgs()
        {
        }
        public static new CbrBackupShareMemberGetArgs Empty => new CbrBackupShareMemberGetArgs();
    }
}
