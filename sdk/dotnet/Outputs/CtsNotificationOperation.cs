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
    public sealed class CtsNotificationOperation
    {
        public readonly string Resource;
        public readonly string Service;
        public readonly ImmutableArray<string> TraceNames;

        [OutputConstructor]
        private CtsNotificationOperation(
            string resource,

            string service,

            ImmutableArray<string> traceNames)
        {
            Resource = resource;
            Service = service;
            TraceNames = traceNames;
        }
    }
}
