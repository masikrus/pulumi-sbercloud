// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Sbercloud.Inputs
{

    public sealed class SmnSubscriptionExtensionArgs : global::Pulumi.ResourceArgs
    {
        [Input("clientId")]
        public Input<string>? ClientId { get; set; }

        [Input("clientSecret")]
        public Input<string>? ClientSecret { get; set; }

        [Input("keyword")]
        public Input<string>? Keyword { get; set; }

        [Input("signSecret")]
        public Input<string>? SignSecret { get; set; }

        public SmnSubscriptionExtensionArgs()
        {
        }
        public static new SmnSubscriptionExtensionArgs Empty => new SmnSubscriptionExtensionArgs();
    }
}
