// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Sbercloud
{
    public static class GetIdentityRole
    {
        public static Task<GetIdentityRoleResult> InvokeAsync(GetIdentityRoleArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetIdentityRoleResult>("sbercloud:index/getIdentityRole:getIdentityRole", args ?? new GetIdentityRoleArgs(), options.WithDefaults());

        public static Output<GetIdentityRoleResult> Invoke(GetIdentityRoleInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetIdentityRoleResult>("sbercloud:index/getIdentityRole:getIdentityRole", args ?? new GetIdentityRoleInvokeArgs(), options.WithDefaults());

        public static Output<GetIdentityRoleResult> Invoke(GetIdentityRoleInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetIdentityRoleResult>("sbercloud:index/getIdentityRole:getIdentityRole", args ?? new GetIdentityRoleInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetIdentityRoleArgs : global::Pulumi.InvokeArgs
    {
        [Input("displayName")]
        public string? DisplayName { get; set; }

        [Input("name")]
        public string? Name { get; set; }

        public GetIdentityRoleArgs()
        {
        }
        public static new GetIdentityRoleArgs Empty => new GetIdentityRoleArgs();
    }

    public sealed class GetIdentityRoleInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("displayName")]
        public Input<string>? DisplayName { get; set; }

        [Input("name")]
        public Input<string>? Name { get; set; }

        public GetIdentityRoleInvokeArgs()
        {
        }
        public static new GetIdentityRoleInvokeArgs Empty => new GetIdentityRoleInvokeArgs();
    }


    [OutputType]
    public sealed class GetIdentityRoleResult
    {
        public readonly string Catalog;
        public readonly string Description;
        public readonly string DisplayName;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string Name;
        public readonly string Policy;
        public readonly string Type;

        [OutputConstructor]
        private GetIdentityRoleResult(
            string catalog,

            string description,

            string displayName,

            string id,

            string name,

            string policy,

            string type)
        {
            Catalog = catalog;
            Description = description;
            DisplayName = displayName;
            Id = id;
            Name = name;
            Policy = policy;
            Type = type;
        }
    }
}
