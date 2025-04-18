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
    public sealed class FgsFunctionNetworkController
    {
        /// <summary>
        /// Whether to disable the public network access.
        /// </summary>
        public readonly bool? DisablePublicNetwork;
        /// <summary>
        /// The configuration of the VPCs that can trigger the function.
        /// </summary>
        public readonly ImmutableArray<Outputs.FgsFunctionNetworkControllerTriggerAccessVpc> TriggerAccessVpcs;

        [OutputConstructor]
        private FgsFunctionNetworkController(
            bool? disablePublicNetwork,

            ImmutableArray<Outputs.FgsFunctionNetworkControllerTriggerAccessVpc> triggerAccessVpcs)
        {
            DisablePublicNetwork = disablePublicNetwork;
            TriggerAccessVpcs = triggerAccessVpcs;
        }
    }
}
