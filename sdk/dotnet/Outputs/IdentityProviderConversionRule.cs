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
    public sealed class IdentityProviderConversionRule
    {
        public readonly ImmutableArray<Outputs.IdentityProviderConversionRuleLocal> Locals;
        public readonly ImmutableArray<Outputs.IdentityProviderConversionRuleRemote> Remotes;

        [OutputConstructor]
        private IdentityProviderConversionRule(
            ImmutableArray<Outputs.IdentityProviderConversionRuleLocal> locals,

            ImmutableArray<Outputs.IdentityProviderConversionRuleRemote> remotes)
        {
            Locals = locals;
            Remotes = remotes;
        }
    }
}
