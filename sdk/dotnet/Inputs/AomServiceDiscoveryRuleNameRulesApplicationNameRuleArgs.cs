// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Sbercloud.Inputs
{

    public sealed class AomServiceDiscoveryRuleNameRulesApplicationNameRuleArgs : global::Pulumi.ResourceArgs
    {
        [Input("args", required: true)]
        private InputList<string>? _args;
        public InputList<string> Args
        {
            get => _args ?? (_args = new InputList<string>());
            set => _args = value;
        }

        [Input("nameType", required: true)]
        public Input<string> NameType { get; set; } = null!;

        [Input("values")]
        private InputList<string>? _values;
        public InputList<string> Values
        {
            get => _values ?? (_values = new InputList<string>());
            set => _values = value;
        }

        public AomServiceDiscoveryRuleNameRulesApplicationNameRuleArgs()
        {
        }
        public static new AomServiceDiscoveryRuleNameRulesApplicationNameRuleArgs Empty => new AomServiceDiscoveryRuleNameRulesApplicationNameRuleArgs();
    }
}
