// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Sbercloud.Inputs
{

    public sealed class ElbL7policyRedirectPoolsExtendConfigGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("insertHeadersConfig")]
        public Input<Inputs.ElbL7policyRedirectPoolsExtendConfigInsertHeadersConfigGetArgs>? InsertHeadersConfig { get; set; }

        [Input("removeHeadersConfig")]
        public Input<Inputs.ElbL7policyRedirectPoolsExtendConfigRemoveHeadersConfigGetArgs>? RemoveHeadersConfig { get; set; }

        [Input("rewriteUrlConfig")]
        public Input<Inputs.ElbL7policyRedirectPoolsExtendConfigRewriteUrlConfigGetArgs>? RewriteUrlConfig { get; set; }

        [Input("rewriteUrlEnabled")]
        public Input<bool>? RewriteUrlEnabled { get; set; }

        [Input("trafficLimitConfig")]
        public Input<Inputs.ElbL7policyRedirectPoolsExtendConfigTrafficLimitConfigGetArgs>? TrafficLimitConfig { get; set; }

        public ElbL7policyRedirectPoolsExtendConfigGetArgs()
        {
        }
        public static new ElbL7policyRedirectPoolsExtendConfigGetArgs Empty => new ElbL7policyRedirectPoolsExtendConfigGetArgs();
    }
}
