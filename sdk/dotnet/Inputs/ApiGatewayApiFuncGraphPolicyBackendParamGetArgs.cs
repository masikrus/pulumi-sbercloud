// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Sbercloud.Inputs
{

    public sealed class ApiGatewayApiFuncGraphPolicyBackendParamGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The description of the parameter.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Where the parameter is located.
        /// </summary>
        [Input("location", required: true)]
        public Input<string> Location { get; set; } = null!;

        /// <summary>
        /// The parameter name.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        [Input("systemParamType")]
        public Input<string>? SystemParamType { get; set; }

        /// <summary>
        /// The parameter type.
        /// </summary>
        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        /// <summary>
        /// The value of the parameter
        /// </summary>
        [Input("value", required: true)]
        public Input<string> Value { get; set; } = null!;

        public ApiGatewayApiFuncGraphPolicyBackendParamGetArgs()
        {
        }
        public static new ApiGatewayApiFuncGraphPolicyBackendParamGetArgs Empty => new ApiGatewayApiFuncGraphPolicyBackendParamGetArgs();
    }
}
