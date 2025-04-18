// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Sbercloud
{
    public static class GetIdentityProjects
    {
        public static Task<GetIdentityProjectsResult> InvokeAsync(GetIdentityProjectsArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetIdentityProjectsResult>("sbercloud:index/getIdentityProjects:getIdentityProjects", args ?? new GetIdentityProjectsArgs(), options.WithDefaults());

        public static Output<GetIdentityProjectsResult> Invoke(GetIdentityProjectsInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetIdentityProjectsResult>("sbercloud:index/getIdentityProjects:getIdentityProjects", args ?? new GetIdentityProjectsInvokeArgs(), options.WithDefaults());

        public static Output<GetIdentityProjectsResult> Invoke(GetIdentityProjectsInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetIdentityProjectsResult>("sbercloud:index/getIdentityProjects:getIdentityProjects", args ?? new GetIdentityProjectsInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetIdentityProjectsArgs : global::Pulumi.InvokeArgs
    {
        [Input("name")]
        public string? Name { get; set; }

        public GetIdentityProjectsArgs()
        {
        }
        public static new GetIdentityProjectsArgs Empty => new GetIdentityProjectsArgs();
    }

    public sealed class GetIdentityProjectsInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("name")]
        public Input<string>? Name { get; set; }

        public GetIdentityProjectsInvokeArgs()
        {
        }
        public static new GetIdentityProjectsInvokeArgs Empty => new GetIdentityProjectsInvokeArgs();
    }


    [OutputType]
    public sealed class GetIdentityProjectsResult
    {
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string? Name;
        public readonly ImmutableArray<Outputs.GetIdentityProjectsProjectResult> Projects;

        [OutputConstructor]
        private GetIdentityProjectsResult(
            string id,

            string? name,

            ImmutableArray<Outputs.GetIdentityProjectsProjectResult> projects)
        {
            Id = id;
            Name = name;
            Projects = projects;
        }
    }
}
