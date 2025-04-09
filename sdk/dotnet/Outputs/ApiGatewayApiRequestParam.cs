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
    public sealed class ApiGatewayApiRequestParam
    {
        /// <summary>
        /// The default value of the parameter.
        /// </summary>
        public readonly string? Default;
        /// <summary>
        /// The parameter description.
        /// </summary>
        public readonly string? Description;
        /// <summary>
        /// The enumerated value.
        /// </summary>
        public readonly string? Enumeration;
        /// <summary>
        /// The parameter example.
        /// </summary>
        public readonly string? Example;
        /// <summary>
        /// Where this parameter is located.
        /// </summary>
        public readonly string? Location;
        /// <summary>
        /// The maximum value or length (string parameter) for parameter.
        /// </summary>
        public readonly int? Maximum;
        /// <summary>
        /// The minimum value or length (string parameter) for parameter.
        /// </summary>
        public readonly int? Minimum;
        /// <summary>
        /// The name of the request parameter.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The list of orchestration rules that parameter used.
        /// </summary>
        public readonly ImmutableArray<string> Orchestrations;
        /// <summary>
        /// Whether to transparently transfer the parameter.
        /// </summary>
        public readonly bool? Passthrough;
        /// <summary>
        /// Whether this parameter is required.
        /// </summary>
        public readonly bool? Required;
        /// <summary>
        /// The parameter type.
        /// </summary>
        public readonly string? Type;
        /// <summary>
        /// Whether to enable the parameter validation.
        /// </summary>
        public readonly int? ValidEnable;

        [OutputConstructor]
        private ApiGatewayApiRequestParam(
            string? @default,

            string? description,

            string? enumeration,

            string? example,

            string? location,

            int? maximum,

            int? minimum,

            string name,

            ImmutableArray<string> orchestrations,

            bool? passthrough,

            bool? required,

            string? type,

            int? validEnable)
        {
            Default = @default;
            Description = description;
            Enumeration = enumeration;
            Example = example;
            Location = location;
            Maximum = maximum;
            Minimum = minimum;
            Name = name;
            Orchestrations = orchestrations;
            Passthrough = passthrough;
            Required = required;
            Type = type;
            ValidEnable = validEnable;
        }
    }
}
