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
    public sealed class AsPolicyScalingPolicyAction
    {
        public readonly int? InstanceNumber;
        public readonly int? InstancePercentage;
        public readonly string? Operation;

        [OutputConstructor]
        private AsPolicyScalingPolicyAction(
            int? instanceNumber,

            int? instancePercentage,

            string? operation)
        {
            InstanceNumber = instanceNumber;
            InstancePercentage = instancePercentage;
            Operation = operation;
        }
    }
}
