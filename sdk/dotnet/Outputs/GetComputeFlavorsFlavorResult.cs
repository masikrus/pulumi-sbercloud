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
    public sealed class GetComputeFlavorsFlavorResult
    {
        public readonly int CpuCoreCount;
        public readonly string Generation;
        public readonly string Id;
        public readonly int MemorySize;
        public readonly string PerformanceType;
        public readonly string StorageType;

        [OutputConstructor]
        private GetComputeFlavorsFlavorResult(
            int cpuCoreCount,

            string generation,

            string id,

            int memorySize,

            string performanceType,

            string storageType)
        {
            CpuCoreCount = cpuCoreCount;
            Generation = generation;
            Id = id;
            MemorySize = memorySize;
            PerformanceType = performanceType;
            StorageType = storageType;
        }
    }
}
