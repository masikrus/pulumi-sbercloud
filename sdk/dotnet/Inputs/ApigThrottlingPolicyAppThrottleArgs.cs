// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Sbercloud.Inputs
{

    public sealed class ApigThrottlingPolicyAppThrottleArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ID of the special user/application throttling policy.
        /// </summary>
        [Input("id")]
        public Input<string>? Id { get; set; }

        /// <summary>
        /// The maximum number of times an API can be accessed within a specified period.
        /// </summary>
        [Input("maxApiRequests", required: true)]
        public Input<int> MaxApiRequests { get; set; } = null!;

        /// <summary>
        /// The object ID which the special throttling policy belongs.
        /// </summary>
        [Input("throttlingObjectId", required: true)]
        public Input<string> ThrottlingObjectId { get; set; } = null!;

        /// <summary>
        /// The object name which the special user/application throttling policy belongs.
        /// </summary>
        [Input("throttlingObjectName")]
        public Input<string>? ThrottlingObjectName { get; set; }

        public ApigThrottlingPolicyAppThrottleArgs()
        {
        }
        public static new ApigThrottlingPolicyAppThrottleArgs Empty => new ApigThrottlingPolicyAppThrottleArgs();
    }
}
