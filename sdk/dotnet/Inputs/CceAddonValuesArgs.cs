// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Sbercloud.Inputs
{

    public sealed class CceAddonValuesArgs : global::Pulumi.ResourceArgs
    {
        [Input("basic")]
        private InputMap<string>? _basic;
        public InputMap<string> Basic
        {
            get => _basic ?? (_basic = new InputMap<string>());
            set => _basic = value;
        }

        [Input("basicJson")]
        public Input<string>? BasicJson { get; set; }

        [Input("custom")]
        private InputMap<string>? _custom;
        public InputMap<string> Custom
        {
            get => _custom ?? (_custom = new InputMap<string>());
            set => _custom = value;
        }

        [Input("customJson")]
        public Input<string>? CustomJson { get; set; }

        [Input("flavor")]
        private InputMap<string>? _flavor;
        public InputMap<string> Flavor
        {
            get => _flavor ?? (_flavor = new InputMap<string>());
            set => _flavor = value;
        }

        [Input("flavorJson")]
        public Input<string>? FlavorJson { get; set; }

        public CceAddonValuesArgs()
        {
        }
        public static new CceAddonValuesArgs Empty => new CceAddonValuesArgs();
    }
}
