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
    public sealed class GetElbFlavorsFlavorResult
    {
        public readonly int Bandwidth;
        public readonly int Cps;
        public readonly string Id;
        public readonly int MaxConnections;
        public readonly string Name;
        public readonly int Qps;
        public readonly string Type;

        [OutputConstructor]
        private GetElbFlavorsFlavorResult(
            int bandwidth,

            int cps,

            string id,

            int maxConnections,

            string name,

            int qps,

            string type)
        {
            Bandwidth = bandwidth;
            Cps = cps;
            Id = id;
            MaxConnections = maxConnections;
            Name = name;
            Qps = qps;
            Type = type;
        }
    }
}
