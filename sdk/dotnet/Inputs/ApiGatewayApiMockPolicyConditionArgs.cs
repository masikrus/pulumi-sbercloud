// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Sbercloud.Inputs
{

    public sealed class ApiGatewayApiMockPolicyConditionArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The cookie parameter name.
        /// </summary>
        [Input("cookieName")]
        public Input<string>? CookieName { get; set; }

        /// <summary>
        /// The frontend authentication parameter name.
        /// </summary>
        [Input("frontendAuthorizerName")]
        public Input<string>? FrontendAuthorizerName { get; set; }

        /// <summary>
        /// The location of a parameter generated after orchestration.
        /// </summary>
        [Input("mappedParamLocation")]
        public Input<string>? MappedParamLocation { get; set; }

        /// <summary>
        /// The name of a parameter generated after orchestration.
        /// </summary>
        [Input("mappedParamName")]
        public Input<string>? MappedParamName { get; set; }

        /// <summary>
        /// The request parameter name.
        /// </summary>
        [Input("paramName")]
        public Input<string>? ParamName { get; set; }

        /// <summary>
        /// The type of the backend policy.
        /// </summary>
        [Input("source")]
        public Input<string>? Source { get; set; }

        /// <summary>
        /// The gateway built-in parameter name.
        /// </summary>
        [Input("sysName")]
        public Input<string>? SysName { get; set; }

        /// <summary>
        /// The condition type.
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        /// <summary>
        /// The value of the backend policy.
        /// </summary>
        [Input("value", required: true)]
        public Input<string> Value { get; set; } = null!;

        public ApiGatewayApiMockPolicyConditionArgs()
        {
        }
        public static new ApiGatewayApiMockPolicyConditionArgs Empty => new ApiGatewayApiMockPolicyConditionArgs();
    }
}
