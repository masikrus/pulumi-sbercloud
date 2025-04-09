// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Sbercloud.Inputs
{

    public sealed class ApigChannelMemberGroupArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The description of the member group.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("microserviceLabels")]
        private InputMap<string>? _microserviceLabels;

        /// <summary>
        /// The microservice tags of the backend server group.
        /// </summary>
        public InputMap<string> MicroserviceLabels
        {
            get => _microserviceLabels ?? (_microserviceLabels = new InputMap<string>());
            set => _microserviceLabels = value;
        }

        /// <summary>
        /// The microservice port of the backend server group.
        /// </summary>
        [Input("microservicePort")]
        public Input<int>? MicroservicePort { get; set; }

        /// <summary>
        /// The microservice version of the backend server group.
        /// </summary>
        [Input("microserviceVersion")]
        public Input<string>? MicroserviceVersion { get; set; }

        /// <summary>
        /// The name of the member group.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// The ID of the reference load balance channel.
        /// </summary>
        [Input("referenceVpcChannelId")]
        public Input<string>? ReferenceVpcChannelId { get; set; }

        /// <summary>
        /// The weight of the current member group.
        /// </summary>
        [Input("weight")]
        public Input<int>? Weight { get; set; }

        public ApigChannelMemberGroupArgs()
        {
        }
        public static new ApigChannelMemberGroupArgs Empty => new ApigChannelMemberGroupArgs();
    }
}
