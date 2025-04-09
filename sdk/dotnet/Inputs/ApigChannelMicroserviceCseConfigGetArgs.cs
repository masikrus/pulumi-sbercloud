// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Sbercloud.Inputs
{

    public sealed class ApigChannelMicroserviceCseConfigGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// schema:Internal; The microservice engine ID.
        /// </summary>
        [Input("engineId", required: true)]
        public Input<string> EngineId { get; set; } = null!;

        /// <summary>
        /// schema:Internal; The microservice ID.
        /// </summary>
        [Input("serviceId", required: true)]
        public Input<string> ServiceId { get; set; } = null!;

        public ApigChannelMicroserviceCseConfigGetArgs()
        {
        }
        public static new ApigChannelMicroserviceCseConfigGetArgs Empty => new ApigChannelMicroserviceCseConfigGetArgs();
    }
}
