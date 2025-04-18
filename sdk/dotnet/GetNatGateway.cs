// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Sbercloud
{
    public static class GetNatGateway
    {
        public static Task<GetNatGatewayResult> InvokeAsync(GetNatGatewayArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetNatGatewayResult>("sbercloud:index/getNatGateway:getNatGateway", args ?? new GetNatGatewayArgs(), options.WithDefaults());

        public static Output<GetNatGatewayResult> Invoke(GetNatGatewayInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetNatGatewayResult>("sbercloud:index/getNatGateway:getNatGateway", args ?? new GetNatGatewayInvokeArgs(), options.WithDefaults());

        public static Output<GetNatGatewayResult> Invoke(GetNatGatewayInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetNatGatewayResult>("sbercloud:index/getNatGateway:getNatGateway", args ?? new GetNatGatewayInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetNatGatewayArgs : global::Pulumi.InvokeArgs
    {
        [Input("description")]
        public string? Description { get; set; }

        [Input("enterpriseProjectId")]
        public string? EnterpriseProjectId { get; set; }

        [Input("id")]
        public string? Id { get; set; }

        [Input("internalNetworkId")]
        public string? InternalNetworkId { get; set; }

        [Input("name")]
        public string? Name { get; set; }

        [Input("region")]
        public string? Region { get; set; }

        [Input("routerId")]
        public string? RouterId { get; set; }

        [Input("spec")]
        public string? Spec { get; set; }

        [Input("status")]
        public string? Status { get; set; }

        [Input("subnetId")]
        public string? SubnetId { get; set; }

        [Input("vpcId")]
        public string? VpcId { get; set; }

        public GetNatGatewayArgs()
        {
        }
        public static new GetNatGatewayArgs Empty => new GetNatGatewayArgs();
    }

    public sealed class GetNatGatewayInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("enterpriseProjectId")]
        public Input<string>? EnterpriseProjectId { get; set; }

        [Input("id")]
        public Input<string>? Id { get; set; }

        [Input("internalNetworkId")]
        public Input<string>? InternalNetworkId { get; set; }

        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("region")]
        public Input<string>? Region { get; set; }

        [Input("routerId")]
        public Input<string>? RouterId { get; set; }

        [Input("spec")]
        public Input<string>? Spec { get; set; }

        [Input("status")]
        public Input<string>? Status { get; set; }

        [Input("subnetId")]
        public Input<string>? SubnetId { get; set; }

        [Input("vpcId")]
        public Input<string>? VpcId { get; set; }

        public GetNatGatewayInvokeArgs()
        {
        }
        public static new GetNatGatewayInvokeArgs Empty => new GetNatGatewayInvokeArgs();
    }


    [OutputType]
    public sealed class GetNatGatewayResult
    {
        public readonly string Description;
        public readonly string EnterpriseProjectId;
        public readonly string Id;
        public readonly string? InternalNetworkId;
        public readonly string Name;
        public readonly string Region;
        public readonly string? RouterId;
        public readonly string Spec;
        public readonly string Status;
        public readonly string SubnetId;
        public readonly string VpcId;

        [OutputConstructor]
        private GetNatGatewayResult(
            string description,

            string enterpriseProjectId,

            string id,

            string? internalNetworkId,

            string name,

            string region,

            string? routerId,

            string spec,

            string status,

            string subnetId,

            string vpcId)
        {
            Description = description;
            EnterpriseProjectId = enterpriseProjectId;
            Id = id;
            InternalNetworkId = internalNetworkId;
            Name = name;
            Region = region;
            RouterId = routerId;
            Spec = spec;
            Status = status;
            SubnetId = subnetId;
            VpcId = vpcId;
        }
    }
}
