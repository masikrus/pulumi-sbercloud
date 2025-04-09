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
    public sealed class GetObsBucketsBucketResult
    {
        public readonly string Bucket;
        public readonly string CreatedAt;
        public readonly string EnterpriseProjectId;
        public readonly string Region;
        public readonly string StorageClass;

        [OutputConstructor]
        private GetObsBucketsBucketResult(
            string bucket,

            string createdAt,

            string enterpriseProjectId,

            string region,

            string storageClass)
        {
            Bucket = bucket;
            CreatedAt = createdAt;
            EnterpriseProjectId = enterpriseProjectId;
            Region = region;
            StorageClass = storageClass;
        }
    }
}
