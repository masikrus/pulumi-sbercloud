// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Sbercloud.Inputs
{

    public sealed class FgsFunctionNetworkControllerGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Whether to disable the public network access.
        /// </summary>
        [Input("disablePublicNetwork")]
        public Input<bool>? DisablePublicNetwork { get; set; }

        [Input("triggerAccessVpcs", required: true)]
        private InputList<Inputs.FgsFunctionNetworkControllerTriggerAccessVpcGetArgs>? _triggerAccessVpcs;

        /// <summary>
        /// The configuration of the VPCs that can trigger the function.
        /// </summary>
        public InputList<Inputs.FgsFunctionNetworkControllerTriggerAccessVpcGetArgs> TriggerAccessVpcs
        {
            get => _triggerAccessVpcs ?? (_triggerAccessVpcs = new InputList<Inputs.FgsFunctionNetworkControllerTriggerAccessVpcGetArgs>());
            set => _triggerAccessVpcs = value;
        }

        public FgsFunctionNetworkControllerGetArgs()
        {
        }
        public static new FgsFunctionNetworkControllerGetArgs Empty => new FgsFunctionNetworkControllerGetArgs();
    }
}
