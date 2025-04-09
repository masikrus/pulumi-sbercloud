// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Sbercloud
{
    public static class GetVpc
    {
        public static Task<GetVpcResult> InvokeAsync(GetVpcArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetVpcResult>("sbercloud:index/getVpc:getVpc", args ?? new GetVpcArgs(), options.WithDefaults());

        public static Output<GetVpcResult> Invoke(GetVpcInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetVpcResult>("sbercloud:index/getVpc:getVpc", args ?? new GetVpcInvokeArgs(), options.WithDefaults());

        public static Output<GetVpcResult> Invoke(GetVpcInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetVpcResult>("sbercloud:index/getVpc:getVpc", args ?? new GetVpcInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetVpcArgs : global::Pulumi.InvokeArgs
    {
        [Input("cidr")]
        public string? Cidr { get; set; }

        [Input("enterpriseProjectId")]
        public string? EnterpriseProjectId { get; set; }

        [Input("id")]
        public string? Id { get; set; }

        [Input("name")]
        public string? Name { get; set; }

        [Input("region")]
        public string? Region { get; set; }

        [Input("status")]
        public string? Status { get; set; }

        public GetVpcArgs()
        {
        }
        public static new GetVpcArgs Empty => new GetVpcArgs();
    }

    public sealed class GetVpcInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("cidr")]
        public Input<string>? Cidr { get; set; }

        [Input("enterpriseProjectId")]
        public Input<string>? EnterpriseProjectId { get; set; }

        [Input("id")]
        public Input<string>? Id { get; set; }

        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("region")]
        public Input<string>? Region { get; set; }

        [Input("status")]
        public Input<string>? Status { get; set; }

        public GetVpcInvokeArgs()
        {
        }
        public static new GetVpcInvokeArgs Empty => new GetVpcInvokeArgs();
    }


    [OutputType]
    public sealed class GetVpcResult
    {
        public readonly string Cidr;
        public readonly string Description;
        public readonly string EnterpriseProjectId;
        public readonly string Id;
        public readonly string Name;
        public readonly string Region;
        public readonly ImmutableArray<Outputs.GetVpcRouteResult> Routes;
        public readonly ImmutableArray<string> SecondaryCidrs;
        public readonly string Status;
        public readonly ImmutableDictionary<string, string> Tags;

        [OutputConstructor]
        private GetVpcResult(
            string cidr,

            string description,

            string enterpriseProjectId,

            string id,

            string name,

            string region,

            ImmutableArray<Outputs.GetVpcRouteResult> routes,

            ImmutableArray<string> secondaryCidrs,

            string status,

            ImmutableDictionary<string, string> tags)
        {
            Cidr = cidr;
            Description = description;
            EnterpriseProjectId = enterpriseProjectId;
            Id = id;
            Name = name;
            Region = region;
            Routes = routes;
            SecondaryCidrs = secondaryCidrs;
            Status = status;
            Tags = tags;
        }
    }
}
