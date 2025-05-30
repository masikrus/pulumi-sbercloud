// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Sbercloud.Inputs
{

    public sealed class DcsInstanceBandwidthInfoArgs : global::Pulumi.ResourceArgs
    {
        [Input("bandwidth")]
        public Input<int>? Bandwidth { get; set; }

        [Input("beginTime")]
        public Input<string>? BeginTime { get; set; }

        [Input("currentTime")]
        public Input<string>? CurrentTime { get; set; }

        [Input("endTime")]
        public Input<string>? EndTime { get; set; }

        [Input("expandCount")]
        public Input<int>? ExpandCount { get; set; }

        [Input("expandEffectTime")]
        public Input<int>? ExpandEffectTime { get; set; }

        [Input("expandIntervalTime")]
        public Input<int>? ExpandIntervalTime { get; set; }

        [Input("maxExpandCount")]
        public Input<int>? MaxExpandCount { get; set; }

        [Input("nextExpandTime")]
        public Input<string>? NextExpandTime { get; set; }

        [Input("taskRunning")]
        public Input<bool>? TaskRunning { get; set; }

        public DcsInstanceBandwidthInfoArgs()
        {
        }
        public static new DcsInstanceBandwidthInfoArgs Empty => new DcsInstanceBandwidthInfoArgs();
    }
}
