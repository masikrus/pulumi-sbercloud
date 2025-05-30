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
    public sealed class GetApigApiFuncGraphPolicyBackendParamResult
    {
        /// <summary>
        /// The description of the constant or system parameter.
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// The ID of the backend parameter.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Where the parameter is located.
        /// </summary>
        public readonly string Location;
        /// <summary>
        /// The name of parameter.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The ID of the corresponding request parameter.
        /// </summary>
        public readonly string RequestId;
        /// <summary>
        /// The type of the system parameter.
        /// </summary>
        public readonly string SystemParamType;
        /// <summary>
        /// The name of parameter.
        /// </summary>
        public readonly string Type;
        /// <summary>
        /// The value of the parameter.
        /// </summary>
        public readonly string Value;

        [OutputConstructor]
        private GetApigApiFuncGraphPolicyBackendParamResult(
            string description,

            string id,

            string location,

            string name,

            string requestId,

            string systemParamType,

            string type,

            string value)
        {
            Description = description;
            Id = id;
            Location = location;
            Name = name;
            RequestId = requestId;
            SystemParamType = systemParamType;
            Type = type;
            Value = value;
        }
    }
}
