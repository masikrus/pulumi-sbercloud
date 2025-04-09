// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Sbercloud
{
    public static class GetRdsFlavors
    {
        public static Task<GetRdsFlavorsResult> InvokeAsync(GetRdsFlavorsArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetRdsFlavorsResult>("sbercloud:index/getRdsFlavors:getRdsFlavors", args ?? new GetRdsFlavorsArgs(), options.WithDefaults());

        public static Output<GetRdsFlavorsResult> Invoke(GetRdsFlavorsInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetRdsFlavorsResult>("sbercloud:index/getRdsFlavors:getRdsFlavors", args ?? new GetRdsFlavorsInvokeArgs(), options.WithDefaults());

        public static Output<GetRdsFlavorsResult> Invoke(GetRdsFlavorsInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetRdsFlavorsResult>("sbercloud:index/getRdsFlavors:getRdsFlavors", args ?? new GetRdsFlavorsInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetRdsFlavorsArgs : global::Pulumi.InvokeArgs
    {
        [Input("availabilityZone")]
        public string? AvailabilityZone { get; set; }

        [Input("dbType", required: true)]
        public string DbType { get; set; } = null!;

        [Input("dbVersion")]
        public string? DbVersion { get; set; }

        [Input("groupType")]
        public string? GroupType { get; set; }

        [Input("instanceMode")]
        public string? InstanceMode { get; set; }

        [Input("memory")]
        public int? Memory { get; set; }

        [Input("region")]
        public string? Region { get; set; }

        [Input("vcpus")]
        public int? Vcpus { get; set; }

        public GetRdsFlavorsArgs()
        {
        }
        public static new GetRdsFlavorsArgs Empty => new GetRdsFlavorsArgs();
    }

    public sealed class GetRdsFlavorsInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("availabilityZone")]
        public Input<string>? AvailabilityZone { get; set; }

        [Input("dbType", required: true)]
        public Input<string> DbType { get; set; } = null!;

        [Input("dbVersion")]
        public Input<string>? DbVersion { get; set; }

        [Input("groupType")]
        public Input<string>? GroupType { get; set; }

        [Input("instanceMode")]
        public Input<string>? InstanceMode { get; set; }

        [Input("memory")]
        public Input<int>? Memory { get; set; }

        [Input("region")]
        public Input<string>? Region { get; set; }

        [Input("vcpus")]
        public Input<int>? Vcpus { get; set; }

        public GetRdsFlavorsInvokeArgs()
        {
        }
        public static new GetRdsFlavorsInvokeArgs Empty => new GetRdsFlavorsInvokeArgs();
    }


    [OutputType]
    public sealed class GetRdsFlavorsResult
    {
        public readonly string? AvailabilityZone;
        public readonly string DbType;
        public readonly string? DbVersion;
        public readonly ImmutableArray<Outputs.GetRdsFlavorsFlavorResult> Flavors;
        public readonly string? GroupType;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string? InstanceMode;
        public readonly int? Memory;
        public readonly string Region;
        public readonly int? Vcpus;

        [OutputConstructor]
        private GetRdsFlavorsResult(
            string? availabilityZone,

            string dbType,

            string? dbVersion,

            ImmutableArray<Outputs.GetRdsFlavorsFlavorResult> flavors,

            string? groupType,

            string id,

            string? instanceMode,

            int? memory,

            string region,

            int? vcpus)
        {
            AvailabilityZone = availabilityZone;
            DbType = dbType;
            DbVersion = dbVersion;
            Flavors = flavors;
            GroupType = groupType;
            Id = id;
            InstanceMode = instanceMode;
            Memory = memory;
            Region = region;
            Vcpus = vcpus;
        }
    }
}
