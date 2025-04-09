// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Sbercloud.Inputs
{

    public sealed class ElbL7policyRedirectUrlConfigArgs : global::Pulumi.ResourceArgs
    {
        [Input("host")]
        public Input<string>? Host { get; set; }

        [Input("insertHeadersConfig")]
        public Input<Inputs.ElbL7policyRedirectUrlConfigInsertHeadersConfigArgs>? InsertHeadersConfig { get; set; }

        [Input("path")]
        public Input<string>? Path { get; set; }

        [Input("port")]
        public Input<string>? Port { get; set; }

        [Input("protocol")]
        public Input<string>? Protocol { get; set; }

        [Input("query")]
        public Input<string>? Query { get; set; }

        [Input("removeHeadersConfig")]
        public Input<Inputs.ElbL7policyRedirectUrlConfigRemoveHeadersConfigArgs>? RemoveHeadersConfig { get; set; }

        [Input("statusCode", required: true)]
        public Input<string> StatusCode { get; set; } = null!;

        public ElbL7policyRedirectUrlConfigArgs()
        {
        }
        public static new ElbL7policyRedirectUrlConfigArgs Empty => new ElbL7policyRedirectUrlConfigArgs();
    }
}
