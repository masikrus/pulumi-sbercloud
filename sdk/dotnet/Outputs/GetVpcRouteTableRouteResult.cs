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
    public sealed class GetVpcRouteTableRouteResult
    {
        public readonly string Description;
        public readonly string Destination;
        public readonly string Nexthop;
        public readonly string Type;

        [OutputConstructor]
        private GetVpcRouteTableRouteResult(
            string description,

            string destination,

            string nexthop,

            string type)
        {
            Description = description;
            Destination = destination;
            Nexthop = nexthop;
            Type = type;
        }
    }
}
