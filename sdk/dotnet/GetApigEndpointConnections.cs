// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Sbercloud
{
    public static class GetApigEndpointConnections
    {
        public static Task<GetApigEndpointConnectionsResult> InvokeAsync(GetApigEndpointConnectionsArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetApigEndpointConnectionsResult>("sbercloud:index/getApigEndpointConnections:getApigEndpointConnections", args ?? new GetApigEndpointConnectionsArgs(), options.WithDefaults());

        public static Output<GetApigEndpointConnectionsResult> Invoke(GetApigEndpointConnectionsInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetApigEndpointConnectionsResult>("sbercloud:index/getApigEndpointConnections:getApigEndpointConnections", args ?? new GetApigEndpointConnectionsInvokeArgs(), options.WithDefaults());

        public static Output<GetApigEndpointConnectionsResult> Invoke(GetApigEndpointConnectionsInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetApigEndpointConnectionsResult>("sbercloud:index/getApigEndpointConnections:getApigEndpointConnections", args ?? new GetApigEndpointConnectionsInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetApigEndpointConnectionsArgs : global::Pulumi.InvokeArgs
    {
        [Input("endpointId")]
        public string? EndpointId { get; set; }

        [Input("instanceId", required: true)]
        public string InstanceId { get; set; } = null!;

        [Input("packetId")]
        public int? PacketId { get; set; }

        [Input("region")]
        public string? Region { get; set; }

        [Input("status")]
        public string? Status { get; set; }

        public GetApigEndpointConnectionsArgs()
        {
        }
        public static new GetApigEndpointConnectionsArgs Empty => new GetApigEndpointConnectionsArgs();
    }

    public sealed class GetApigEndpointConnectionsInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("endpointId")]
        public Input<string>? EndpointId { get; set; }

        [Input("instanceId", required: true)]
        public Input<string> InstanceId { get; set; } = null!;

        [Input("packetId")]
        public Input<int>? PacketId { get; set; }

        [Input("region")]
        public Input<string>? Region { get; set; }

        [Input("status")]
        public Input<string>? Status { get; set; }

        public GetApigEndpointConnectionsInvokeArgs()
        {
        }
        public static new GetApigEndpointConnectionsInvokeArgs Empty => new GetApigEndpointConnectionsInvokeArgs();
    }


    [OutputType]
    public sealed class GetApigEndpointConnectionsResult
    {
        public readonly ImmutableArray<Outputs.GetApigEndpointConnectionsConnectionResult> Connections;
        public readonly string? EndpointId;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string InstanceId;
        public readonly int? PacketId;
        public readonly string Region;
        public readonly string? Status;

        [OutputConstructor]
        private GetApigEndpointConnectionsResult(
            ImmutableArray<Outputs.GetApigEndpointConnectionsConnectionResult> connections,

            string? endpointId,

            string id,

            string instanceId,

            int? packetId,

            string region,

            string? status)
        {
            Connections = connections;
            EndpointId = endpointId;
            Id = id;
            InstanceId = instanceId;
            PacketId = packetId;
            Region = region;
            Status = status;
        }
    }
}
