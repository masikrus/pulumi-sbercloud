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
    public sealed class ApiGatewayApiWebPolicyCondition
    {
        /// <summary>
        /// The cookie parameter name.
        /// </summary>
        public readonly string? CookieName;
        /// <summary>
        /// The frontend authentication parameter name.
        /// </summary>
        public readonly string? FrontendAuthorizerName;
        /// <summary>
        /// The location of a parameter generated after orchestration.
        /// </summary>
        public readonly string? MappedParamLocation;
        /// <summary>
        /// The name of a parameter generated after orchestration.
        /// </summary>
        public readonly string? MappedParamName;
        /// <summary>
        /// The request parameter name.
        /// </summary>
        public readonly string? ParamName;
        /// <summary>
        /// The type of the backend policy.
        /// </summary>
        public readonly string? Source;
        /// <summary>
        /// The gateway built-in parameter name.
        /// </summary>
        public readonly string? SysName;
        /// <summary>
        /// The condition type.
        /// </summary>
        public readonly string? Type;
        /// <summary>
        /// The value of the backend policy.
        /// </summary>
        public readonly string Value;

        [OutputConstructor]
        private ApiGatewayApiWebPolicyCondition(
            string? cookieName,

            string? frontendAuthorizerName,

            string? mappedParamLocation,

            string? mappedParamName,

            string? paramName,

            string? source,

            string? sysName,

            string? type,

            string value)
        {
            CookieName = cookieName;
            FrontendAuthorizerName = frontendAuthorizerName;
            MappedParamLocation = mappedParamLocation;
            MappedParamName = mappedParamName;
            ParamName = paramName;
            Source = source;
            SysName = sysName;
            Type = type;
            Value = value;
        }
    }
}
