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
    public sealed class CesAlarmruleMetric
    {
        public readonly ImmutableArray<Outputs.CesAlarmruleMetricDimension> Dimensions;
        public readonly string MetricName;
        public readonly string Namespace;

        [OutputConstructor]
        private CesAlarmruleMetric(
            ImmutableArray<Outputs.CesAlarmruleMetricDimension> dimensions,

            string metricName,

            string @namespace)
        {
            Dimensions = dimensions;
            MetricName = metricName;
            Namespace = @namespace;
        }
    }
}
