// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Sbercloud.Inputs
{

    public sealed class MapreduceClusterSmnNotifyArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The subscription rule name.
        /// </summary>
        [Input("subscriptionName", required: true)]
        public Input<string> SubscriptionName { get; set; } = null!;

        /// <summary>
        /// The Uniform Resource Name (URN) of the topic.
        /// </summary>
        [Input("topicUrn", required: true)]
        public Input<string> TopicUrn { get; set; } = null!;

        public MapreduceClusterSmnNotifyArgs()
        {
        }
        public static new MapreduceClusterSmnNotifyArgs Empty => new MapreduceClusterSmnNotifyArgs();
    }
}
