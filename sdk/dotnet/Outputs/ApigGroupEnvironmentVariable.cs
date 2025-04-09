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
    public sealed class ApigGroupEnvironmentVariable
    {
        /// <summary>
        /// The ID of the variable that the group has.
        /// </summary>
        public readonly string? Id;
        /// <summary>
        /// The variable name.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The variable value.
        /// </summary>
        public readonly string Value;
        /// <summary>
        /// schema: Deprecated; The ID of the variable that the group has.
        /// </summary>
        public readonly string? VariableId;

        [OutputConstructor]
        private ApigGroupEnvironmentVariable(
            string? id,

            string name,

            string value,

            string? variableId)
        {
            Id = id;
            Name = name;
            Value = value;
            VariableId = variableId;
        }
    }
}
