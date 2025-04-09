// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Sbercloud.Inputs
{

    public sealed class IdentityProviderAccessConfigArgs : global::Pulumi.ResourceArgs
    {
        [Input("accessType", required: true)]
        public Input<string> AccessType { get; set; } = null!;

        [Input("authorizationEndpoint")]
        public Input<string>? AuthorizationEndpoint { get; set; }

        [Input("clientId", required: true)]
        public Input<string> ClientId { get; set; } = null!;

        [Input("providerUrl", required: true)]
        public Input<string> ProviderUrl { get; set; } = null!;

        [Input("responseMode")]
        public Input<string>? ResponseMode { get; set; }

        [Input("responseType")]
        public Input<string>? ResponseType { get; set; }

        [Input("scopes")]
        private InputList<string>? _scopes;
        public InputList<string> Scopes
        {
            get => _scopes ?? (_scopes = new InputList<string>());
            set => _scopes = value;
        }

        [Input("signingKey", required: true)]
        public Input<string> SigningKey { get; set; } = null!;

        public IdentityProviderAccessConfigArgs()
        {
        }
        public static new IdentityProviderAccessConfigArgs Empty => new IdentityProviderAccessConfigArgs();
    }
}
