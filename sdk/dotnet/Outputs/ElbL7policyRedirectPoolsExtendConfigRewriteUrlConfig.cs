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
    public sealed class ElbL7policyRedirectPoolsExtendConfigRewriteUrlConfig
    {
        public readonly string? Host;
        public readonly string? Path;
        public readonly string? Query;

        [OutputConstructor]
        private ElbL7policyRedirectPoolsExtendConfigRewriteUrlConfig(
            string? host,

            string? path,

            string? query)
        {
            Host = host;
            Path = path;
            Query = query;
        }
    }
}
