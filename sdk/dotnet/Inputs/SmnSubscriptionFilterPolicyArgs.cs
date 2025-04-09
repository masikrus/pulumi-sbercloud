// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Sbercloud.Inputs
{

    public sealed class SmnSubscriptionFilterPolicyArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The filter policy name. The policy name must be unique.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("stringEquals")]
        private InputList<string>? _stringEquals;

        /// <summary>
        /// The string array for exact match.
        /// </summary>
        public InputList<string> StringEquals
        {
            get => _stringEquals ?? (_stringEquals = new InputList<string>());
            set => _stringEquals = value;
        }

        public SmnSubscriptionFilterPolicyArgs()
        {
        }
        public static new SmnSubscriptionFilterPolicyArgs Empty => new SmnSubscriptionFilterPolicyArgs();
    }
}
