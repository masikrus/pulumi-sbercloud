// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Sbercloud.Inputs
{

    public sealed class AomServiceDiscoveryRuleDiscoveryRuleGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("checkContents", required: true)]
        private InputList<string>? _checkContents;
        public InputList<string> CheckContents
        {
            get => _checkContents ?? (_checkContents = new InputList<string>());
            set => _checkContents = value;
        }

        [Input("checkMode", required: true)]
        public Input<string> CheckMode { get; set; } = null!;

        [Input("checkType", required: true)]
        public Input<string> CheckType { get; set; } = null!;

        public AomServiceDiscoveryRuleDiscoveryRuleGetArgs()
        {
        }
        public static new AomServiceDiscoveryRuleDiscoveryRuleGetArgs Empty => new AomServiceDiscoveryRuleDiscoveryRuleGetArgs();
    }
}
