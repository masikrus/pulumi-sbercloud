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
    public sealed class ApiGatewayApiWebPolicy
    {
        /// <summary>
        /// The ID of the backend custom authorization.
        /// </summary>
        public readonly string? AuthorizerId;
        /// <summary>
        /// The backend service address
        /// </summary>
        public readonly string? BackendAddress;
        /// <summary>
        /// The configuration list of the backend parameters.
        /// </summary>
        public readonly ImmutableArray<Outputs.ApiGatewayApiWebPolicyBackendParam> BackendParams;
        /// <summary>
        /// The policy conditions.
        /// </summary>
        public readonly ImmutableArray<Outputs.ApiGatewayApiWebPolicyCondition> Conditions;
        /// <summary>
        /// The effective mode of the backend policy.
        /// </summary>
        public readonly string? EffectiveMode;
        /// <summary>
        /// The proxy host header.
        /// </summary>
        public readonly string? HostHeader;
        /// <summary>
        /// The name of the web policy.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The backend request address.
        /// </summary>
        public readonly string Path;
        /// <summary>
        /// The backend request method of the API.
        /// </summary>
        public readonly string RequestMethod;
        /// <summary>
        /// The backend request protocol.
        /// </summary>
        public readonly string? RequestProtocol;
        /// <summary>
        /// The number of retry attempts to request the backend service.
        /// </summary>
        public readonly int? RetryCount;
        /// <summary>
        /// The timeout for API requests to backend service.
        /// </summary>
        public readonly int? Timeout;
        /// <summary>
        /// The VPC channel ID.
        /// </summary>
        public readonly string? VpcChannelId;

        [OutputConstructor]
        private ApiGatewayApiWebPolicy(
            string? authorizerId,

            string? backendAddress,

            ImmutableArray<Outputs.ApiGatewayApiWebPolicyBackendParam> backendParams,

            ImmutableArray<Outputs.ApiGatewayApiWebPolicyCondition> conditions,

            string? effectiveMode,

            string? hostHeader,

            string name,

            string path,

            string requestMethod,

            string? requestProtocol,

            int? retryCount,

            int? timeout,

            string? vpcChannelId)
        {
            AuthorizerId = authorizerId;
            BackendAddress = backendAddress;
            BackendParams = backendParams;
            Conditions = conditions;
            EffectiveMode = effectiveMode;
            HostHeader = hostHeader;
            Name = name;
            Path = path;
            RequestMethod = requestMethod;
            RequestProtocol = requestProtocol;
            RetryCount = retryCount;
            Timeout = timeout;
            VpcChannelId = vpcChannelId;
        }
    }
}
