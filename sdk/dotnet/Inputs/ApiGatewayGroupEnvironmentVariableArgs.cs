// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Sbercloud.Inputs
{

    public sealed class ApiGatewayGroupEnvironmentVariableArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ID of the variable that the group has.
        /// </summary>
        [Input("id")]
        public Input<string>? Id { get; set; }

        /// <summary>
        /// The variable name.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// The variable value.
        /// </summary>
        [Input("value", required: true)]
        public Input<string> Value { get; set; } = null!;

        /// <summary>
        /// schema: Deprecated; The ID of the variable that the group has.
        /// </summary>
        [Input("variableId")]
        public Input<string>? VariableId { get; set; }

        public ApiGatewayGroupEnvironmentVariableArgs()
        {
        }
        public static new ApiGatewayGroupEnvironmentVariableArgs Empty => new ApiGatewayGroupEnvironmentVariableArgs();
    }
}
