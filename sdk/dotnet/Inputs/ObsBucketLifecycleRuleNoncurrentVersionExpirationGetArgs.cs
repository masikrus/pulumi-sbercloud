// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Sbercloud.Inputs
{

    public sealed class ObsBucketLifecycleRuleNoncurrentVersionExpirationGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("days", required: true)]
        public Input<int> Days { get; set; } = null!;

        public ObsBucketLifecycleRuleNoncurrentVersionExpirationGetArgs()
        {
        }
        public static new ObsBucketLifecycleRuleNoncurrentVersionExpirationGetArgs Empty => new ObsBucketLifecycleRuleNoncurrentVersionExpirationGetArgs();
    }
}
