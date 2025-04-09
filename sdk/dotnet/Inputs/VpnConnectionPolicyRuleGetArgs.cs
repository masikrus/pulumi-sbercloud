// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Sbercloud.Inputs
{

    public sealed class VpnConnectionPolicyRuleGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("destinations")]
        private InputList<string>? _destinations;

        /// <summary>
        /// The list of destination CIDRs.
        /// </summary>
        public InputList<string> Destinations
        {
            get => _destinations ?? (_destinations = new InputList<string>());
            set => _destinations = value;
        }

        /// <summary>
        /// The rule index.
        /// </summary>
        [Input("ruleIndex")]
        public Input<int>? RuleIndex { get; set; }

        /// <summary>
        /// The source CIDR.
        /// </summary>
        [Input("source")]
        public Input<string>? Source { get; set; }

        public VpnConnectionPolicyRuleGetArgs()
        {
        }
        public static new VpnConnectionPolicyRuleGetArgs Empty => new VpnConnectionPolicyRuleGetArgs();
    }
}
