// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Sbercloud.Inputs
{

    public sealed class ElbListenerPortRangeGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("endPort", required: true)]
        public Input<int> EndPort { get; set; } = null!;

        [Input("startPort", required: true)]
        public Input<int> StartPort { get; set; } = null!;

        public ElbListenerPortRangeGetArgs()
        {
        }
        public static new ElbListenerPortRangeGetArgs Empty => new ElbListenerPortRangeGetArgs();
    }
}
