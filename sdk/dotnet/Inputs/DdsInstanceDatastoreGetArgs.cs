// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Sbercloud.Inputs
{

    public sealed class DdsInstanceDatastoreGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("storageEngine")]
        public Input<string>? StorageEngine { get; set; }

        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        [Input("version", required: true)]
        public Input<string> Version { get; set; } = null!;

        public DdsInstanceDatastoreGetArgs()
        {
        }
        public static new DdsInstanceDatastoreGetArgs Empty => new DdsInstanceDatastoreGetArgs();
    }
}
