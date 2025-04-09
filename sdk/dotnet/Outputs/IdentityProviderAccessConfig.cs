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
    public sealed class IdentityProviderAccessConfig
    {
        public readonly string AccessType;
        public readonly string? AuthorizationEndpoint;
        public readonly string ClientId;
        public readonly string ProviderUrl;
        public readonly string? ResponseMode;
        public readonly string? ResponseType;
        public readonly ImmutableArray<string> Scopes;
        public readonly string SigningKey;

        [OutputConstructor]
        private IdentityProviderAccessConfig(
            string accessType,

            string? authorizationEndpoint,

            string clientId,

            string providerUrl,

            string? responseMode,

            string? responseType,

            ImmutableArray<string> scopes,

            string signingKey)
        {
            AccessType = accessType;
            AuthorizationEndpoint = authorizationEndpoint;
            ClientId = clientId;
            ProviderUrl = providerUrl;
            ResponseMode = responseMode;
            ResponseType = responseType;
            Scopes = scopes;
            SigningKey = signingKey;
        }
    }
}
