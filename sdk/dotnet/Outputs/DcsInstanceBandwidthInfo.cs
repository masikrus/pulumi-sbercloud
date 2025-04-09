// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Sbercloud.Outputs
{

    [OutputType]
    public sealed class DcsInstanceBandwidthInfo
    {
        public readonly int? Bandwidth;
        public readonly string? BeginTime;
        public readonly string? CurrentTime;
        public readonly string? EndTime;
        public readonly int? ExpandCount;
        public readonly int? ExpandEffectTime;
        public readonly int? ExpandIntervalTime;
        public readonly int? MaxExpandCount;
        public readonly string? NextExpandTime;
        public readonly bool? TaskRunning;

        [OutputConstructor]
        private DcsInstanceBandwidthInfo(
            int? bandwidth,

            string? beginTime,

            string? currentTime,

            string? endTime,

            int? expandCount,

            int? expandEffectTime,

            int? expandIntervalTime,

            int? maxExpandCount,

            string? nextExpandTime,

            bool? taskRunning)
        {
            Bandwidth = bandwidth;
            BeginTime = beginTime;
            CurrentTime = currentTime;
            EndTime = endTime;
            ExpandCount = expandCount;
            ExpandEffectTime = expandEffectTime;
            ExpandIntervalTime = expandIntervalTime;
            MaxExpandCount = maxExpandCount;
            NextExpandTime = nextExpandTime;
            TaskRunning = taskRunning;
        }
    }
}
