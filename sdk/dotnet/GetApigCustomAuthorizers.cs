// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Sbercloud
{
    public static class GetApigCustomAuthorizers
    {
        public static Task<GetApigCustomAuthorizersResult> InvokeAsync(GetApigCustomAuthorizersArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetApigCustomAuthorizersResult>("sbercloud:index/getApigCustomAuthorizers:getApigCustomAuthorizers", args ?? new GetApigCustomAuthorizersArgs(), options.WithDefaults());

        public static Output<GetApigCustomAuthorizersResult> Invoke(GetApigCustomAuthorizersInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetApigCustomAuthorizersResult>("sbercloud:index/getApigCustomAuthorizers:getApigCustomAuthorizers", args ?? new GetApigCustomAuthorizersInvokeArgs(), options.WithDefaults());

        public static Output<GetApigCustomAuthorizersResult> Invoke(GetApigCustomAuthorizersInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetApigCustomAuthorizersResult>("sbercloud:index/getApigCustomAuthorizers:getApigCustomAuthorizers", args ?? new GetApigCustomAuthorizersInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetApigCustomAuthorizersArgs : global::Pulumi.InvokeArgs
    {
        [Input("authorizerId")]
        public string? AuthorizerId { get; set; }

        [Input("instanceId", required: true)]
        public string InstanceId { get; set; } = null!;

        [Input("name")]
        public string? Name { get; set; }

        [Input("region")]
        public string? Region { get; set; }

        [Input("type")]
        public string? Type { get; set; }

        public GetApigCustomAuthorizersArgs()
        {
        }
        public static new GetApigCustomAuthorizersArgs Empty => new GetApigCustomAuthorizersArgs();
    }

    public sealed class GetApigCustomAuthorizersInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("authorizerId")]
        public Input<string>? AuthorizerId { get; set; }

        [Input("instanceId", required: true)]
        public Input<string> InstanceId { get; set; } = null!;

        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("region")]
        public Input<string>? Region { get; set; }

        [Input("type")]
        public Input<string>? Type { get; set; }

        public GetApigCustomAuthorizersInvokeArgs()
        {
        }
        public static new GetApigCustomAuthorizersInvokeArgs Empty => new GetApigCustomAuthorizersInvokeArgs();
    }


    [OutputType]
    public sealed class GetApigCustomAuthorizersResult
    {
        public readonly string? AuthorizerId;
        public readonly ImmutableArray<Outputs.GetApigCustomAuthorizersAuthorizerResult> Authorizers;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string InstanceId;
        public readonly string? Name;
        public readonly string Region;
        public readonly string? Type;

        [OutputConstructor]
        private GetApigCustomAuthorizersResult(
            string? authorizerId,

            ImmutableArray<Outputs.GetApigCustomAuthorizersAuthorizerResult> authorizers,

            string id,

            string instanceId,

            string? name,

            string region,

            string? type)
        {
            AuthorizerId = authorizerId;
            Authorizers = authorizers;
            Id = id;
            InstanceId = instanceId;
            Name = name;
            Region = region;
            Type = type;
        }
    }
}
