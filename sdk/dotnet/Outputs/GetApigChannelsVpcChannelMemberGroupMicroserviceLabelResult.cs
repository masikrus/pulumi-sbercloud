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
    public sealed class GetApigChannelsVpcChannelMemberGroupMicroserviceLabelResult
    {
        /// <summary>
        /// The name of the microservice label.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The value of the microservice label.
        /// </summary>
        public readonly string Value;

        [OutputConstructor]
        private GetApigChannelsVpcChannelMemberGroupMicroserviceLabelResult(
            string name,

            string value)
        {
            Name = name;
            Value = value;
        }
    }
}
