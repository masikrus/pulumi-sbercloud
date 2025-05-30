// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Sbercloud
{
    public static class GetVpcepPublicServices
    {
        public static Task<GetVpcepPublicServicesResult> InvokeAsync(GetVpcepPublicServicesArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetVpcepPublicServicesResult>("sbercloud:index/getVpcepPublicServices:getVpcepPublicServices", args ?? new GetVpcepPublicServicesArgs(), options.WithDefaults());

        public static Output<GetVpcepPublicServicesResult> Invoke(GetVpcepPublicServicesInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetVpcepPublicServicesResult>("sbercloud:index/getVpcepPublicServices:getVpcepPublicServices", args ?? new GetVpcepPublicServicesInvokeArgs(), options.WithDefaults());

        public static Output<GetVpcepPublicServicesResult> Invoke(GetVpcepPublicServicesInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetVpcepPublicServicesResult>("sbercloud:index/getVpcepPublicServices:getVpcepPublicServices", args ?? new GetVpcepPublicServicesInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetVpcepPublicServicesArgs : global::Pulumi.InvokeArgs
    {
        [Input("region")]
        public string? Region { get; set; }

        [Input("serviceId")]
        public string? ServiceId { get; set; }

        [Input("serviceName")]
        public string? ServiceName { get; set; }

        public GetVpcepPublicServicesArgs()
        {
        }
        public static new GetVpcepPublicServicesArgs Empty => new GetVpcepPublicServicesArgs();
    }

    public sealed class GetVpcepPublicServicesInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("region")]
        public Input<string>? Region { get; set; }

        [Input("serviceId")]
        public Input<string>? ServiceId { get; set; }

        [Input("serviceName")]
        public Input<string>? ServiceName { get; set; }

        public GetVpcepPublicServicesInvokeArgs()
        {
        }
        public static new GetVpcepPublicServicesInvokeArgs Empty => new GetVpcepPublicServicesInvokeArgs();
    }


    [OutputType]
    public sealed class GetVpcepPublicServicesResult
    {
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string Region;
        public readonly string? ServiceId;
        public readonly string? ServiceName;
        public readonly ImmutableArray<Outputs.GetVpcepPublicServicesServiceResult> Services;

        [OutputConstructor]
        private GetVpcepPublicServicesResult(
            string id,

            string region,

            string? serviceId,

            string? serviceName,

            ImmutableArray<Outputs.GetVpcepPublicServicesServiceResult> services)
        {
            Id = id;
            Region = region;
            ServiceId = serviceId;
            ServiceName = serviceName;
            Services = services;
        }
    }
}
