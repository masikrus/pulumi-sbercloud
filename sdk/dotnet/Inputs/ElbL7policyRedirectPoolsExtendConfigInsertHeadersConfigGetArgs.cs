// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Sbercloud.Inputs
{

    public sealed class ElbL7policyRedirectPoolsExtendConfigInsertHeadersConfigGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("configs", required: true)]
        private InputList<Inputs.ElbL7policyRedirectPoolsExtendConfigInsertHeadersConfigConfigGetArgs>? _configs;
        public InputList<Inputs.ElbL7policyRedirectPoolsExtendConfigInsertHeadersConfigConfigGetArgs> Configs
        {
            get => _configs ?? (_configs = new InputList<Inputs.ElbL7policyRedirectPoolsExtendConfigInsertHeadersConfigConfigGetArgs>());
            set => _configs = value;
        }

        public ElbL7policyRedirectPoolsExtendConfigInsertHeadersConfigGetArgs()
        {
        }
        public static new ElbL7policyRedirectPoolsExtendConfigInsertHeadersConfigGetArgs Empty => new ElbL7policyRedirectPoolsExtendConfigInsertHeadersConfigGetArgs();
    }
}
