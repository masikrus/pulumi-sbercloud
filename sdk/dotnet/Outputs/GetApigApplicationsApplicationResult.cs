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
    public sealed class GetApigApplicationsApplicationResult
    {
        /// <summary>
        /// The key of the application.
        /// </summary>
        public readonly string AppKey;
        /// <summary>
        /// The secret of the application.
        /// </summary>
        public readonly string AppSecret;
        /// <summary>
        /// The type of the application.
        /// </summary>
        public readonly string AppType;
        /// <summary>
        /// The number of bound APIs.
        /// </summary>
        public readonly int BindNum;
        /// <summary>
        /// The creation time of the application.
        /// </summary>
        public readonly string CreatedAt;
        /// <summary>
        /// The creator of the application.
        /// </summary>
        public readonly string CreatedBy;
        /// <summary>
        /// The description of the application.
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// The ID of the application.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The name of the application.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The status of the application.
        /// </summary>
        public readonly int Status;
        /// <summary>
        /// The latest update time of the application.
        /// </summary>
        public readonly string UpdatedAt;

        [OutputConstructor]
        private GetApigApplicationsApplicationResult(
            string appKey,

            string appSecret,

            string appType,

            int bindNum,

            string createdAt,

            string createdBy,

            string description,

            string id,

            string name,

            int status,

            string updatedAt)
        {
            AppKey = appKey;
            AppSecret = appSecret;
            AppType = appType;
            BindNum = bindNum;
            CreatedAt = createdAt;
            CreatedBy = createdBy;
            Description = description;
            Id = id;
            Name = name;
            Status = status;
            UpdatedAt = updatedAt;
        }
    }
}
