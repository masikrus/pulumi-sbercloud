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
    public sealed class ApigApiMockPolicy
    {
        /// <summary>
        /// The ID of the backend custom authorization.
        /// </summary>
        public readonly string? AuthorizerId;
        /// <summary>
        /// The configuration list of backend parameters.
        /// </summary>
        public readonly ImmutableArray<Outputs.ApigApiMockPolicyBackendParam> BackendParams;
        /// <summary>
        /// The policy conditions.
        /// </summary>
        public readonly ImmutableArray<Outputs.ApigApiMockPolicyCondition> Conditions;
        /// <summary>
        /// The effective mode of the backend policy.
        /// </summary>
        public readonly string? EffectiveMode;
        /// <summary>
        /// The backend policy name.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The response content of the mock.
        /// </summary>
        public readonly string? Response;
        /// <summary>
        /// The custom status code of the mock response.
        /// </summary>
        public readonly int? StatusCode;

        [OutputConstructor]
        private ApigApiMockPolicy(
            string? authorizerId,

            ImmutableArray<Outputs.ApigApiMockPolicyBackendParam> backendParams,

            ImmutableArray<Outputs.ApigApiMockPolicyCondition> conditions,

            string? effectiveMode,

            string name,

            string? response,

            int? statusCode)
        {
            AuthorizerId = authorizerId;
            BackendParams = backendParams;
            Conditions = conditions;
            EffectiveMode = effectiveMode;
            Name = name;
            Response = response;
            StatusCode = statusCode;
        }
    }
}
