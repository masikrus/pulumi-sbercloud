// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Sbercloud.Inputs
{

    public sealed class ElbL7policyRedirectPoolsExtendConfigRewriteUrlConfigArgs : global::Pulumi.ResourceArgs
    {
        [Input("host")]
        public Input<string>? Host { get; set; }

        [Input("path")]
        public Input<string>? Path { get; set; }

        [Input("query")]
        public Input<string>? Query { get; set; }

        public ElbL7policyRedirectPoolsExtendConfigRewriteUrlConfigArgs()
        {
        }
        public static new ElbL7policyRedirectPoolsExtendConfigRewriteUrlConfigArgs Empty => new ElbL7policyRedirectPoolsExtendConfigRewriteUrlConfigArgs();
    }
}
