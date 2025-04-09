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
    public sealed class ApigApiBackendParam
    {
        /// <summary>
        /// The description of the parameter.
        /// </summary>
        public readonly string? Description;
        /// <summary>
        /// Where the parameter is located.
        /// </summary>
        public readonly string Location;
        /// <summary>
        /// The parameter name.
        /// </summary>
        public readonly string Name;
        public readonly string? SystemParamType;
        /// <summary>
        /// The parameter type.
        /// </summary>
        public readonly string Type;
        /// <summary>
        /// The value of the parameter
        /// </summary>
        public readonly string Value;

        [OutputConstructor]
        private ApigApiBackendParam(
            string? description,

            string location,

            string name,

            string? systemParamType,

            string type,

            string value)
        {
            Description = description;
            Location = location;
            Name = name;
            SystemParamType = systemParamType;
            Type = type;
            Value = value;
        }
    }
}
