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
    public sealed class DliSparkJobDependentPackage
    {
        public readonly string GroupName;
        public readonly ImmutableArray<Outputs.DliSparkJobDependentPackagePackage> Packages;

        [OutputConstructor]
        private DliSparkJobDependentPackage(
            string groupName,

            ImmutableArray<Outputs.DliSparkJobDependentPackagePackage> packages)
        {
            GroupName = groupName;
            Packages = packages;
        }
    }
}
