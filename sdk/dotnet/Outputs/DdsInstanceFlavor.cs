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
    public sealed class DdsInstanceFlavor
    {
        public readonly int Num;
        public readonly int? Size;
        public readonly string SpecCode;
        public readonly string? Storage;
        public readonly string Type;

        [OutputConstructor]
        private DdsInstanceFlavor(
            int num,

            int? size,

            string specCode,

            string? storage,

            string type)
        {
            Num = num;
            Size = size;
            SpecCode = specCode;
            Storage = storage;
            Type = type;
        }
    }
}
