// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Sbercloud
{
    public static class GetVpcRoute
    {
        public static Task<GetVpcRouteResult> InvokeAsync(GetVpcRouteArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetVpcRouteResult>("sbercloud:index/getVpcRoute:getVpcRoute", args ?? new GetVpcRouteArgs(), options.WithDefaults());

        public static Output<GetVpcRouteResult> Invoke(GetVpcRouteInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetVpcRouteResult>("sbercloud:index/getVpcRoute:getVpcRoute", args ?? new GetVpcRouteInvokeArgs(), options.WithDefaults());

        public static Output<GetVpcRouteResult> Invoke(GetVpcRouteInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetVpcRouteResult>("sbercloud:index/getVpcRoute:getVpcRoute", args ?? new GetVpcRouteInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetVpcRouteArgs : global::Pulumi.InvokeArgs
    {
        [Input("destination")]
        public string? Destination { get; set; }

        [Input("id")]
        public string? Id { get; set; }

        [Input("region")]
        public string? Region { get; set; }

        [Input("tenantId")]
        public string? TenantId { get; set; }

        [Input("type")]
        public string? Type { get; set; }

        [Input("vpcId")]
        public string? VpcId { get; set; }

        public GetVpcRouteArgs()
        {
        }
        public static new GetVpcRouteArgs Empty => new GetVpcRouteArgs();
    }

    public sealed class GetVpcRouteInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("destination")]
        public Input<string>? Destination { get; set; }

        [Input("id")]
        public Input<string>? Id { get; set; }

        [Input("region")]
        public Input<string>? Region { get; set; }

        [Input("tenantId")]
        public Input<string>? TenantId { get; set; }

        [Input("type")]
        public Input<string>? Type { get; set; }

        [Input("vpcId")]
        public Input<string>? VpcId { get; set; }

        public GetVpcRouteInvokeArgs()
        {
        }
        public static new GetVpcRouteInvokeArgs Empty => new GetVpcRouteInvokeArgs();
    }


    [OutputType]
    public sealed class GetVpcRouteResult
    {
        public readonly string Destination;
        public readonly string Id;
        public readonly string Nexthop;
        public readonly string Region;
        public readonly string TenantId;
        public readonly string Type;
        public readonly string VpcId;

        [OutputConstructor]
        private GetVpcRouteResult(
            string destination,

            string id,

            string nexthop,

            string region,

            string tenantId,

            string type,

            string vpcId)
        {
            Destination = destination;
            Id = id;
            Nexthop = nexthop;
            Region = region;
            TenantId = tenantId;
            Type = type;
            VpcId = vpcId;
        }
    }
}
