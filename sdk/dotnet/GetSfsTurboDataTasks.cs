// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Sbercloud
{
    public static class GetSfsTurboDataTasks
    {
        public static Task<GetSfsTurboDataTasksResult> InvokeAsync(GetSfsTurboDataTasksArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetSfsTurboDataTasksResult>("sbercloud:index/getSfsTurboDataTasks:getSfsTurboDataTasks", args ?? new GetSfsTurboDataTasksArgs(), options.WithDefaults());

        public static Output<GetSfsTurboDataTasksResult> Invoke(GetSfsTurboDataTasksInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetSfsTurboDataTasksResult>("sbercloud:index/getSfsTurboDataTasks:getSfsTurboDataTasks", args ?? new GetSfsTurboDataTasksInvokeArgs(), options.WithDefaults());

        public static Output<GetSfsTurboDataTasksResult> Invoke(GetSfsTurboDataTasksInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetSfsTurboDataTasksResult>("sbercloud:index/getSfsTurboDataTasks:getSfsTurboDataTasks", args ?? new GetSfsTurboDataTasksInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetSfsTurboDataTasksArgs : global::Pulumi.InvokeArgs
    {
        [Input("region")]
        public string? Region { get; set; }

        [Input("shareId", required: true)]
        public string ShareId { get; set; } = null!;

        [Input("status")]
        public string? Status { get; set; }

        [Input("type")]
        public string? Type { get; set; }

        public GetSfsTurboDataTasksArgs()
        {
        }
        public static new GetSfsTurboDataTasksArgs Empty => new GetSfsTurboDataTasksArgs();
    }

    public sealed class GetSfsTurboDataTasksInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("region")]
        public Input<string>? Region { get; set; }

        [Input("shareId", required: true)]
        public Input<string> ShareId { get; set; } = null!;

        [Input("status")]
        public Input<string>? Status { get; set; }

        [Input("type")]
        public Input<string>? Type { get; set; }

        public GetSfsTurboDataTasksInvokeArgs()
        {
        }
        public static new GetSfsTurboDataTasksInvokeArgs Empty => new GetSfsTurboDataTasksInvokeArgs();
    }


    [OutputType]
    public sealed class GetSfsTurboDataTasksResult
    {
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string Region;
        public readonly string ShareId;
        public readonly string? Status;
        public readonly ImmutableArray<Outputs.GetSfsTurboDataTasksTaskResult> Tasks;
        public readonly string? Type;

        [OutputConstructor]
        private GetSfsTurboDataTasksResult(
            string id,

            string region,

            string shareId,

            string? status,

            ImmutableArray<Outputs.GetSfsTurboDataTasksTaskResult> tasks,

            string? type)
        {
            Id = id;
            Region = region;
            ShareId = shareId;
            Status = status;
            Tasks = tasks;
            Type = type;
        }
    }
}
