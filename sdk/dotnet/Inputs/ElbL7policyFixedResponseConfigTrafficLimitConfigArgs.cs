// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Sbercloud.Inputs
{

    public sealed class ElbL7policyFixedResponseConfigTrafficLimitConfigArgs : global::Pulumi.ResourceArgs
    {
        [Input("burst")]
        public Input<int>? Burst { get; set; }

        [Input("perSourceIpQps")]
        public Input<int>? PerSourceIpQps { get; set; }

        [Input("qps")]
        public Input<int>? Qps { get; set; }

        public ElbL7policyFixedResponseConfigTrafficLimitConfigArgs()
        {
        }
        public static new ElbL7policyFixedResponseConfigTrafficLimitConfigArgs Empty => new ElbL7policyFixedResponseConfigTrafficLimitConfigArgs();
    }
}
