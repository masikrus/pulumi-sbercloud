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
    public sealed class RdsParametergroupDatastore
    {
        public readonly string Type;
        public readonly string Version;

        [OutputConstructor]
        private RdsParametergroupDatastore(
            string type,

            string version)
        {
            Type = type;
            Version = version;
        }
    }
}
