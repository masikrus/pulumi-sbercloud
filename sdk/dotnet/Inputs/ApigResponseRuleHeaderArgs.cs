// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Sbercloud.Inputs
{

    public sealed class ApigResponseRuleHeaderArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The key name of the response header.
        /// </summary>
        [Input("key", required: true)]
        public Input<string> Key { get; set; } = null!;

        /// <summary>
        /// The value for the specified response header key.
        /// </summary>
        [Input("value", required: true)]
        public Input<string> Value { get; set; } = null!;

        public ApigResponseRuleHeaderArgs()
        {
        }
        public static new ApigResponseRuleHeaderArgs Empty => new ApigResponseRuleHeaderArgs();
    }
}
