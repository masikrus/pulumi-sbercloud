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
    public sealed class ApiGatewayApiMock
    {
        /// <summary>
        /// The ID of the backend custom authorization.
        /// </summary>
        public readonly string? AuthorizerId;
        /// <summary>
        /// The response content of the mock.
        /// </summary>
        public readonly string? Response;
        /// <summary>
        /// The custom status code of the mock response.
        /// </summary>
        public readonly int? StatusCode;

        [OutputConstructor]
        private ApiGatewayApiMock(
            string? authorizerId,

            string? response,

            int? statusCode)
        {
            AuthorizerId = authorizerId;
            Response = response;
            StatusCode = statusCode;
        }
    }
}
