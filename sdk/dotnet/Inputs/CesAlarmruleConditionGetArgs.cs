// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Sbercloud.Inputs
{

    public sealed class CesAlarmruleConditionGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("comparisonOperator", required: true)]
        public Input<string> ComparisonOperator { get; set; } = null!;

        [Input("count", required: true)]
        public Input<int> Count { get; set; } = null!;

        [Input("filter", required: true)]
        public Input<string> Filter { get; set; } = null!;

        [Input("period", required: true)]
        public Input<int> Period { get; set; } = null!;

        [Input("suppressDuration")]
        public Input<int>? SuppressDuration { get; set; }

        [Input("unit")]
        public Input<string>? Unit { get; set; }

        [Input("value", required: true)]
        public Input<int> Value { get; set; } = null!;

        public CesAlarmruleConditionGetArgs()
        {
        }
        public static new CesAlarmruleConditionGetArgs Empty => new CesAlarmruleConditionGetArgs();
    }
}
