// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Sbercloud.Inputs
{

    public sealed class ElbL7policyRedirectPoolsExtendConfigRemoveHeadersConfigGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("configs", required: true)]
        private InputList<Inputs.ElbL7policyRedirectPoolsExtendConfigRemoveHeadersConfigConfigGetArgs>? _configs;
        public InputList<Inputs.ElbL7policyRedirectPoolsExtendConfigRemoveHeadersConfigConfigGetArgs> Configs
        {
            get => _configs ?? (_configs = new InputList<Inputs.ElbL7policyRedirectPoolsExtendConfigRemoveHeadersConfigConfigGetArgs>());
            set => _configs = value;
        }

        public ElbL7policyRedirectPoolsExtendConfigRemoveHeadersConfigGetArgs()
        {
        }
        public static new ElbL7policyRedirectPoolsExtendConfigRemoveHeadersConfigGetArgs Empty => new ElbL7policyRedirectPoolsExtendConfigRemoveHeadersConfigGetArgs();
    }
}
