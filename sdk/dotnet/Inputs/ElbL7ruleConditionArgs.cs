// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Sbercloud.Inputs
{

    public sealed class ElbL7ruleConditionArgs : global::Pulumi.ResourceArgs
    {
        [Input("key")]
        public Input<string>? Key { get; set; }

        [Input("value", required: true)]
        public Input<string> Value { get; set; } = null!;

        public ElbL7ruleConditionArgs()
        {
        }
        public static new ElbL7ruleConditionArgs Empty => new ElbL7ruleConditionArgs();
    }
}
