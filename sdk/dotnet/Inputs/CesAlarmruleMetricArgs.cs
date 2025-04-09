// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Sbercloud.Inputs
{

    public sealed class CesAlarmruleMetricArgs : global::Pulumi.ResourceArgs
    {
        [Input("dimensions")]
        private InputList<Inputs.CesAlarmruleMetricDimensionArgs>? _dimensions;
        public InputList<Inputs.CesAlarmruleMetricDimensionArgs> Dimensions
        {
            get => _dimensions ?? (_dimensions = new InputList<Inputs.CesAlarmruleMetricDimensionArgs>());
            set => _dimensions = value;
        }

        [Input("metricName", required: true)]
        public Input<string> MetricName { get; set; } = null!;

        [Input("namespace", required: true)]
        public Input<string> Namespace { get; set; } = null!;

        public CesAlarmruleMetricArgs()
        {
        }
        public static new CesAlarmruleMetricArgs Empty => new CesAlarmruleMetricArgs();
    }
}
