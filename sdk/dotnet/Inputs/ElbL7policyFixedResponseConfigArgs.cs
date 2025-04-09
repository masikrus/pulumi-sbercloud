// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Sbercloud.Inputs
{

    public sealed class ElbL7policyFixedResponseConfigArgs : global::Pulumi.ResourceArgs
    {
        [Input("contentType")]
        public Input<string>? ContentType { get; set; }

        [Input("insertHeadersConfig")]
        public Input<Inputs.ElbL7policyFixedResponseConfigInsertHeadersConfigArgs>? InsertHeadersConfig { get; set; }

        [Input("messageBody")]
        public Input<string>? MessageBody { get; set; }

        [Input("removeHeadersConfig")]
        public Input<Inputs.ElbL7policyFixedResponseConfigRemoveHeadersConfigArgs>? RemoveHeadersConfig { get; set; }

        [Input("statusCode", required: true)]
        public Input<string> StatusCode { get; set; } = null!;

        [Input("trafficLimitConfig")]
        public Input<Inputs.ElbL7policyFixedResponseConfigTrafficLimitConfigArgs>? TrafficLimitConfig { get; set; }

        public ElbL7policyFixedResponseConfigArgs()
        {
        }
        public static new ElbL7policyFixedResponseConfigArgs Empty => new ElbL7policyFixedResponseConfigArgs();
    }
}
