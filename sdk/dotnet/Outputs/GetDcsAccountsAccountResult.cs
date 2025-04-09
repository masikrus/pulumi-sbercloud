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
    public sealed class GetDcsAccountsAccountResult
    {
        /// <summary>
        /// Account name.
        /// </summary>
        public readonly string AccountName;
        /// <summary>
        /// Account permissions.
        /// </summary>
        public readonly string AccountRole;
        /// <summary>
        /// Account type.
        /// </summary>
        public readonly string AccountType;
        /// <summary>
        /// Account description.
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// Account ID.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Account status.
        /// </summary>
        public readonly string Status;

        [OutputConstructor]
        private GetDcsAccountsAccountResult(
            string accountName,

            string accountRole,

            string accountType,

            string description,

            string id,

            string status)
        {
            AccountName = accountName;
            AccountRole = accountRole;
            AccountType = accountType;
            Description = description;
            Id = id;
            Status = status;
        }
    }
}
