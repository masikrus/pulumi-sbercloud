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
    public sealed class ApigApiFuncGraphPolicy
    {
        /// <summary>
        /// The ID of the backend custom authorization.
        /// </summary>
        public readonly string? AuthorizerId;
        /// <summary>
        /// The configaiton list of the backend parameters.
        /// </summary>
        public readonly ImmutableArray<Outputs.ApigApiFuncGraphPolicyBackendParam> BackendParams;
        /// <summary>
        /// The policy conditions.
        /// </summary>
        public readonly ImmutableArray<Outputs.ApigApiFuncGraphPolicyCondition> Conditions;
        /// <summary>
        /// The effective mode of the backend policy.
        /// </summary>
        public readonly string? EffectiveMode;
        /// <summary>
        /// The alias URN of the FunctionGraph function.
        /// </summary>
        public readonly string? FunctionAliasUrn;
        /// <summary>
        /// The URN of the FunctionGraph function.
        /// </summary>
        public readonly string FunctionUrn;
        /// <summary>
        /// The invocation mode of the FunctionGraph function.
        /// </summary>
        public readonly string? InvocationMode;
        /// <summary>
        /// The invocation mode of the FunctionGraph function.
        /// </summary>
        public readonly string? InvocationType;
        /// <summary>
        /// The name of the backend policy.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The network (framework) type of the FunctionGraph function.
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
        private ApigApiFuncGraphPolicy(
            string? authorizerId,

            ImmutableArray<Outputs.ApigApiFuncGraphPolicyBackendParam> backendParams,

            ImmutableArray<Outputs.ApigApiFuncGraphPolicyCondition> conditions,

            string? effectiveMode,

            string? functionAliasUrn,

            string functionUrn,

            string? invocationMode,

            string? invocationType,

            string name,

            string? networkType,

            string? requestProtocol,

            int? timeout,

            string? version)
        {
            AuthorizerId = authorizerId;
            BackendParams = backendParams;
            Conditions = conditions;
            EffectiveMode = effectiveMode;
            FunctionAliasUrn = functionAliasUrn;
            FunctionUrn = functionUrn;
            InvocationMode = invocationMode;
            InvocationType = invocationType;
            Name = name;
            NetworkType = networkType;
            RequestProtocol = requestProtocol;
            Timeout = timeout;
            Version = version;
        }
    }
}
