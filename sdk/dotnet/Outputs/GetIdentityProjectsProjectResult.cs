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
    public sealed class GetIdentityProjectsProjectResult
    {
        public readonly bool Enabled;
        public readonly string Id;
        public readonly string Name;

        [OutputConstructor]
        private GetIdentityProjectsProjectResult(
            bool enabled,

            string id,

            string name)
        {
            Enabled = enabled;
            Id = id;
            Name = name;
        }
    }
}
