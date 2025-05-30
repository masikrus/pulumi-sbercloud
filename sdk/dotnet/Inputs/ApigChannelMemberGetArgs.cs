// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Sbercloud.Inputs
{

    public sealed class ApigChannelMemberGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The group name of the backend server.
        /// </summary>
        [Input("groupName")]
        public Input<string>? GroupName { get; set; }

        /// <summary>
        /// The IP address of the backend server.
        /// </summary>
        [Input("host")]
        public Input<string>? Host { get; set; }

        /// <summary>
        /// The ID of the backend server.
        /// </summary>
        [Input("id")]
        public Input<string>? Id { get; set; }

        /// <summary>
        /// Whether this member is the backup member.
        /// </summary>
        [Input("isBackup")]
        public Input<bool>? IsBackup { get; set; }

        /// <summary>
        /// The name of the backend server.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The port of the backend server.
        /// </summary>
        [Input("port")]
        public Input<int>? Port { get; set; }

        /// <summary>
        /// The status of the backend server.
        /// </summary>
        [Input("status")]
        public Input<int>? Status { get; set; }

        /// <summary>
        /// The weight of current backend server.
        /// </summary>
        [Input("weight")]
        public Input<int>? Weight { get; set; }

        public ApigChannelMemberGetArgs()
        {
        }
        public static new ApigChannelMemberGetArgs Empty => new ApigChannelMemberGetArgs();
    }
}
