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
    public sealed class CceNodePoolStorageGroupVirtualSpace
    {
        public readonly string? LvmLvType;
        public readonly string? LvmPath;
        public readonly string Name;
        public readonly string? RuntimeLvType;
        public readonly string Size;

        [OutputConstructor]
        private CceNodePoolStorageGroupVirtualSpace(
            string? lvmLvType,

            string? lvmPath,

            string name,

            string? runtimeLvType,

            string size)
        {
            LvmLvType = lvmLvType;
            LvmPath = lvmPath;
            Name = name;
            RuntimeLvType = runtimeLvType;
            Size = size;
        }
    }
}
