// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Sbercloud.Inputs
{

    public sealed class MapreduceClusterComponentConfigArgs : global::Pulumi.ResourceArgs
    {
        [Input("configs", required: true)]
        private InputList<Inputs.MapreduceClusterComponentConfigConfigArgs>? _configs;
        public InputList<Inputs.MapreduceClusterComponentConfigConfigArgs> Configs
        {
            get => _configs ?? (_configs = new InputList<Inputs.MapreduceClusterComponentConfigConfigArgs>());
            set => _configs = value;
        }

        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        public MapreduceClusterComponentConfigArgs()
        {
        }
        public static new MapreduceClusterComponentConfigArgs Empty => new MapreduceClusterComponentConfigArgs();
    }
}
