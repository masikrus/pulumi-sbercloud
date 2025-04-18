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
    public sealed class SmnSubscriptionExtension
    {
        public readonly string? ClientId;
        public readonly string? ClientSecret;
        public readonly string? Keyword;
        public readonly string? SignSecret;

        [OutputConstructor]
        private SmnSubscriptionExtension(
            string? clientId,

            string? clientSecret,

            string? keyword,

            string? signSecret)
        {
            ClientId = clientId;
            ClientSecret = clientSecret;
            Keyword = keyword;
            SignSecret = signSecret;
        }
    }
}
