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
    public sealed class ElbL7policyRedirectPoolsExtendConfigTrafficLimitConfig
    {
        public readonly int? Burst;
        public readonly int? PerSourceIpQps;
        public readonly int? Qps;

        [OutputConstructor]
        private ElbL7policyRedirectPoolsExtendConfigTrafficLimitConfig(
            int? burst,

            int? perSourceIpQps,

            int? qps)
        {
            Burst = burst;
            PerSourceIpQps = perSourceIpQps;
            Qps = qps;
        }
    }
}
