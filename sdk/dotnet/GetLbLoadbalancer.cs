// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Sbercloud
{
    public static class GetLbLoadbalancer
    {
        public static Task<GetLbLoadbalancerResult> InvokeAsync(GetLbLoadbalancerArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetLbLoadbalancerResult>("sbercloud:index/getLbLoadbalancer:getLbLoadbalancer", args ?? new GetLbLoadbalancerArgs(), options.WithDefaults());

        public static Output<GetLbLoadbalancerResult> Invoke(GetLbLoadbalancerInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetLbLoadbalancerResult>("sbercloud:index/getLbLoadbalancer:getLbLoadbalancer", args ?? new GetLbLoadbalancerInvokeArgs(), options.WithDefaults());

        public static Output<GetLbLoadbalancerResult> Invoke(GetLbLoadbalancerInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetLbLoadbalancerResult>("sbercloud:index/getLbLoadbalancer:getLbLoadbalancer", args ?? new GetLbLoadbalancerInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetLbLoadbalancerArgs : global::Pulumi.InvokeArgs
    {
        [Input("description")]
        public string? Description { get; set; }

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

        [Input("vipAddress")]
        public string? VipAddress { get; set; }

        [Input("vipSubnetId")]
        public string? VipSubnetId { get; set; }

        public GetLbLoadbalancerArgs()
        {
        }
        public static new GetLbLoadbalancerArgs Empty => new GetLbLoadbalancerArgs();
    }

    public sealed class GetLbLoadbalancerInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("description")]
        public Input<string>? Description { get; set; }

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

        [Input("vipAddress")]
        public Input<string>? VipAddress { get; set; }

        [Input("vipSubnetId")]
        public Input<string>? VipSubnetId { get; set; }

        public GetLbLoadbalancerInvokeArgs()
        {
        }
        public static new GetLbLoadbalancerInvokeArgs Empty => new GetLbLoadbalancerInvokeArgs();
    }


    [OutputType]
    public sealed class GetLbLoadbalancerResult
    {
        public readonly string Description;
        public readonly string EnterpriseProjectId;
        public readonly string Id;
        public readonly string Name;
        public readonly string PublicIp;
        public readonly string Region;
        public readonly string Status;
        public readonly ImmutableDictionary<string, string> Tags;
        public readonly string VipAddress;
        public readonly string VipPortId;
        public readonly string VipSubnetId;

        [OutputConstructor]
        private GetLbLoadbalancerResult(
            string description,

            string enterpriseProjectId,

            string id,

            string name,

            string publicIp,

            string region,

            string status,

            ImmutableDictionary<string, string> tags,

            string vipAddress,

            string vipPortId,

            string vipSubnetId)
        {
            Description = description;
            EnterpriseProjectId = enterpriseProjectId;
            Id = id;
            Name = name;
            PublicIp = publicIp;
            Region = region;
            Status = status;
            Tags = tags;
            VipAddress = vipAddress;
            VipPortId = vipPortId;
            VipSubnetId = vipSubnetId;
        }
    }
}
