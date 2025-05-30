// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Sbercloud.Inputs
{

    public sealed class CceNodeStorageSelectorGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("matchLabelCount")]
        public Input<string>? MatchLabelCount { get; set; }

        [Input("matchLabelMetadataCmkid")]
        public Input<string>? MatchLabelMetadataCmkid { get; set; }

        [Input("matchLabelMetadataEncrypted")]
        public Input<string>? MatchLabelMetadataEncrypted { get; set; }

        [Input("matchLabelSize")]
        public Input<string>? MatchLabelSize { get; set; }

        [Input("matchLabelVolumeType")]
        public Input<string>? MatchLabelVolumeType { get; set; }

        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        [Input("type")]
        public Input<string>? Type { get; set; }

        public CceNodeStorageSelectorGetArgs()
        {
        }
        public static new CceNodeStorageSelectorGetArgs Empty => new CceNodeStorageSelectorGetArgs();
    }
}
