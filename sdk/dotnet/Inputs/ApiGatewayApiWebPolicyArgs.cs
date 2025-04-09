// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Sbercloud.Inputs
{

    public sealed class ApiGatewayApiWebPolicyArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ID of the backend custom authorization.
        /// </summary>
        [Input("authorizerId")]
        public Input<string>? AuthorizerId { get; set; }

        /// <summary>
        /// The backend service address
        /// </summary>
        [Input("backendAddress")]
        public Input<string>? BackendAddress { get; set; }

        [Input("backendParams")]
        private InputList<Inputs.ApiGatewayApiWebPolicyBackendParamArgs>? _backendParams;

        /// <summary>
        /// The configuration list of the backend parameters.
        /// </summary>
        public InputList<Inputs.ApiGatewayApiWebPolicyBackendParamArgs> BackendParams
        {
            get => _backendParams ?? (_backendParams = new InputList<Inputs.ApiGatewayApiWebPolicyBackendParamArgs>());
            set => _backendParams = value;
        }

        [Input("conditions", required: true)]
        private InputList<Inputs.ApiGatewayApiWebPolicyConditionArgs>? _conditions;

        /// <summary>
        /// The policy conditions.
        /// </summary>
        public InputList<Inputs.ApiGatewayApiWebPolicyConditionArgs> Conditions
        {
            get => _conditions ?? (_conditions = new InputList<Inputs.ApiGatewayApiWebPolicyConditionArgs>());
            set => _conditions = value;
        }

        /// <summary>
        /// The effective mode of the backend policy.
        /// </summary>
        [Input("effectiveMode")]
        public Input<string>? EffectiveMode { get; set; }

        /// <summary>
        /// The proxy host header.
        /// </summary>
        [Input("hostHeader")]
        public Input<string>? HostHeader { get; set; }

        /// <summary>
        /// The name of the web policy.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// The backend request address.
        /// </summary>
        [Input("path", required: true)]
        public Input<string> Path { get; set; } = null!;

        /// <summary>
        /// The backend request method of the API.
        /// </summary>
        [Input("requestMethod", required: true)]
        public Input<string> RequestMethod { get; set; } = null!;

        /// <summary>
        /// The backend request protocol.
        /// </summary>
        [Input("requestProtocol")]
        public Input<string>? RequestProtocol { get; set; }

        /// <summary>
        /// The number of retry attempts to request the backend service.
        /// </summary>
        [Input("retryCount")]
        public Input<int>? RetryCount { get; set; }

        /// <summary>
        /// The timeout for API requests to backend service.
        /// </summary>
        [Input("timeout")]
        public Input<int>? Timeout { get; set; }

        /// <summary>
        /// The VPC channel ID.
        /// </summary>
        [Input("vpcChannelId")]
        public Input<string>? VpcChannelId { get; set; }

        public ApiGatewayApiWebPolicyArgs()
        {
        }
        public static new ApiGatewayApiWebPolicyArgs Empty => new ApiGatewayApiWebPolicyArgs();
    }
}
