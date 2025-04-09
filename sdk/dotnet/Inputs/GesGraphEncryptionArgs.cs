// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Sbercloud.Inputs
{

    public sealed class GesGraphEncryptionArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Whether to enable data encryption The value can be true or false. The default value is false.
        /// </summary>
        [Input("enable")]
        public Input<bool>? Enable { get; set; }

        /// <summary>
        /// ID of the customer master key created by DEW in the project corresponding to the graph creation.
        /// </summary>
        [Input("masterKeyId")]
        public Input<string>? MasterKeyId { get; set; }

        public GesGraphEncryptionArgs()
        {
        }
        public static new GesGraphEncryptionArgs Empty => new GesGraphEncryptionArgs();
    }
}
