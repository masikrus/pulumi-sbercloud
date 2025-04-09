// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Sbercloud
{
    public static class GetVpcSubnet
    {
        public static Task<GetVpcSubnetResult> InvokeAsync(GetVpcSubnetArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetVpcSubnetResult>("sbercloud:index/getVpcSubnet:getVpcSubnet", args ?? new GetVpcSubnetArgs(), options.WithDefaults());

        public static Output<GetVpcSubnetResult> Invoke(GetVpcSubnetInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetVpcSubnetResult>("sbercloud:index/getVpcSubnet:getVpcSubnet", args ?? new GetVpcSubnetInvokeArgs(), options.WithDefaults());

        public static Output<GetVpcSubnetResult> Invoke(GetVpcSubnetInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetVpcSubnetResult>("sbercloud:index/getVpcSubnet:getVpcSubnet", args ?? new GetVpcSubnetInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetVpcSubnetArgs : global::Pulumi.InvokeArgs
    {
        [Input("availabilityZone")]
        public string? AvailabilityZone { get; set; }

        [Input("cidr")]
        public string? Cidr { get; set; }

        [Input("gatewayIp")]
        public string? GatewayIp { get; set; }

        [Input("id")]
        public string? Id { get; set; }

        [Input("name")]
        public string? Name { get; set; }

        [Input("primaryDns")]
        public string? PrimaryDns { get; set; }

        [Input("region")]
        public string? Region { get; set; }

        [Input("secondaryDns")]
        public string? SecondaryDns { get; set; }

        [Input("status")]
        public string? Status { get; set; }

        [Input("vpcId")]
        public string? VpcId { get; set; }

        public GetVpcSubnetArgs()
        {
        }
        public static new GetVpcSubnetArgs Empty => new GetVpcSubnetArgs();
    }

    public sealed class GetVpcSubnetInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("availabilityZone")]
        public Input<string>? AvailabilityZone { get; set; }

        [Input("cidr")]
        public Input<string>? Cidr { get; set; }

        [Input("gatewayIp")]
        public Input<string>? GatewayIp { get; set; }

        [Input("id")]
        public Input<string>? Id { get; set; }

        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("primaryDns")]
        public Input<string>? PrimaryDns { get; set; }

        [Input("region")]
        public Input<string>? Region { get; set; }

        [Input("secondaryDns")]
        public Input<string>? SecondaryDns { get; set; }

        [Input("status")]
        public Input<string>? Status { get; set; }

        [Input("vpcId")]
        public Input<string>? VpcId { get; set; }

        public GetVpcSubnetInvokeArgs()
        {
        }
        public static new GetVpcSubnetInvokeArgs Empty => new GetVpcSubnetInvokeArgs();
    }


    [OutputType]
    public sealed class GetVpcSubnetResult
    {
        public readonly string AvailabilityZone;
        public readonly string Cidr;
        public readonly string Description;
        public readonly bool DhcpEnable;
        public readonly ImmutableArray<string> DnsLists;
        public readonly string GatewayIp;
        public readonly string Id;
        public readonly string Ipv4SubnetId;
        public readonly string Ipv6Cidr;
        public readonly bool Ipv6Enable;
        public readonly string Ipv6Gateway;
        public readonly string Ipv6SubnetId;
        public readonly string Name;
        public readonly string PrimaryDns;
        public readonly string Region;
        public readonly string SecondaryDns;
        public readonly string Status;
        public readonly string SubnetId;
        public readonly string VpcId;

        [OutputConstructor]
        private GetVpcSubnetResult(
            string availabilityZone,

            string cidr,

            string description,

            bool dhcpEnable,

            ImmutableArray<string> dnsLists,

            string gatewayIp,

            string id,

            string ipv4SubnetId,

            string ipv6Cidr,

            bool ipv6Enable,

            string ipv6Gateway,

            string ipv6SubnetId,

            string name,

            string primaryDns,

            string region,

            string secondaryDns,

            string status,

            string subnetId,

            string vpcId)
        {
            AvailabilityZone = availabilityZone;
            Cidr = cidr;
            Description = description;
            DhcpEnable = dhcpEnable;
            DnsLists = dnsLists;
            GatewayIp = gatewayIp;
            Id = id;
            Ipv4SubnetId = ipv4SubnetId;
            Ipv6Cidr = ipv6Cidr;
            Ipv6Enable = ipv6Enable;
            Ipv6Gateway = ipv6Gateway;
            Ipv6SubnetId = ipv6SubnetId;
            Name = name;
            PrimaryDns = primaryDns;
            Region = region;
            SecondaryDns = secondaryDns;
            Status = status;
            SubnetId = subnetId;
            VpcId = vpcId;
        }
    }
}
