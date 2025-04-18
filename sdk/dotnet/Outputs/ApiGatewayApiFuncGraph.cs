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
    public sealed class ApiGatewayApiFuncGraph
    {
        /// <summary>
        /// The ID of the backend custom authorization.
        /// </summary>
        public readonly string? AuthorizerId;
        /// <summary>
        /// The alias URN of the FunctionGraph function.
        /// </summary>
        public readonly string? FunctionAliasUrn;
        /// <summary>
        /// The URN of the FunctionGraph function.
        /// </summary>
        public readonly string FunctionUrn;
        /// <summary>
        /// The invocation type.
        /// </summary>
        public readonly string? InvocationType;
        /// <summary>
        /// The network architecture (framework) type of the FunctionGraph function.
        /// </summary>
        public readonly string? NetworkType;
        /// <summary>
        /// The request protocol of the FunctionGraph function.
        /// </summary>
        public readonly string? RequestProtocol;
        /// <summary>
        /// The timeout for API requests to backend service.
        /// </summary>
        public readonly int? Timeout;
        /// <summary>
        /// The version of the FunctionGraph function.
        /// </summary>
        public readonly string? Version;

        [OutputConstructor]
        private ApiGatewayApiFuncGraph(
            string? authorizerId,

            string? functionAliasUrn,

            string functionUrn,

            string? invocationType,

            string? networkType,

            string? requestProtocol,

            int? timeout,

            string? version)
        {
            AuthorizerId = authorizerId;
            FunctionAliasUrn = functionAliasUrn;
            FunctionUrn = functionUrn;
            InvocationType = invocationType;
            NetworkType = networkType;
            RequestProtocol = requestProtocol;
            Timeout = timeout;
            Version = version;
        }
    }
}
