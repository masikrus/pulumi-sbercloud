// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Sbercloud
{
    public static class GetSfsTurboDuTasks
    {
        public static Task<GetSfsTurboDuTasksResult> InvokeAsync(GetSfsTurboDuTasksArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetSfsTurboDuTasksResult>("sbercloud:index/getSfsTurboDuTasks:getSfsTurboDuTasks", args ?? new GetSfsTurboDuTasksArgs(), options.WithDefaults());

        public static Output<GetSfsTurboDuTasksResult> Invoke(GetSfsTurboDuTasksInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetSfsTurboDuTasksResult>("sbercloud:index/getSfsTurboDuTasks:getSfsTurboDuTasks", args ?? new GetSfsTurboDuTasksInvokeArgs(), options.WithDefaults());

        public static Output<GetSfsTurboDuTasksResult> Invoke(GetSfsTurboDuTasksInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetSfsTurboDuTasksResult>("sbercloud:index/getSfsTurboDuTasks:getSfsTurboDuTasks", args ?? new GetSfsTurboDuTasksInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetSfsTurboDuTasksArgs : global::Pulumi.InvokeArgs
    {
        [Input("region")]
        public string? Region { get; set; }

        [Input("shareId", required: true)]
        public string ShareId { get; set; } = null!;

        public GetSfsTurboDuTasksArgs()
        {
        }
        public static new GetSfsTurboDuTasksArgs Empty => new GetSfsTurboDuTasksArgs();
    }

    public sealed class GetSfsTurboDuTasksInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("region")]
        public Input<string>? Region { get; set; }

        [Input("shareId", required: true)]
        public Input<string> ShareId { get; set; } = null!;

        public GetSfsTurboDuTasksInvokeArgs()
        {
        }
        public static new GetSfsTurboDuTasksInvokeArgs Empty => new GetSfsTurboDuTasksInvokeArgs();
    }


    [OutputType]
    public sealed class GetSfsTurboDuTasksResult
    {
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string Region;
        public readonly string ShareId;
        public readonly ImmutableArray<Outputs.GetSfsTurboDuTasksTaskResult> Tasks;

        [OutputConstructor]
        private GetSfsTurboDuTasksResult(
            string id,

            string region,

            string shareId,

            ImmutableArray<Outputs.GetSfsTurboDuTasksTaskResult> tasks)
        {
            Id = id;
            Region = region;
            ShareId = shareId;
            Tasks = tasks;
        }
    }
}
