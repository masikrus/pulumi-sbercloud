// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Sbercloud.Inputs
{

    public sealed class RdsPgHbaHostBasedAuthenticationGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specifies the client IP address.
        /// </summary>
        [Input("address", required: true)]
        public Input<string> Address { get; set; } = null!;

        /// <summary>
        /// Specifies the database name.
        /// </summary>
        [Input("database", required: true)]
        public Input<string> Database { get; set; } = null!;

        /// <summary>
        /// Specifies the subnet mask.
        /// </summary>
        [Input("mask")]
        public Input<string>? Mask { get; set; }

        /// <summary>
        /// Specifies the authentication mode.
        /// </summary>
        [Input("method", required: true)]
        public Input<string> Method { get; set; } = null!;

        /// <summary>
        /// Specifies the connection type.
        /// </summary>
        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        /// <summary>
        /// Specifies the Name of a user.
        /// </summary>
        [Input("user", required: true)]
        public Input<string> User { get; set; } = null!;

        public RdsPgHbaHostBasedAuthenticationGetArgs()
        {
        }
        public static new RdsPgHbaHostBasedAuthenticationGetArgs Empty => new RdsPgHbaHostBasedAuthenticationGetArgs();
    }
}
