// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Sbercloud
{
    public static class GetRdsPgDatabases
    {
        public static Task<GetRdsPgDatabasesResult> InvokeAsync(GetRdsPgDatabasesArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetRdsPgDatabasesResult>("sbercloud:index/getRdsPgDatabases:getRdsPgDatabases", args ?? new GetRdsPgDatabasesArgs(), options.WithDefaults());

        public static Output<GetRdsPgDatabasesResult> Invoke(GetRdsPgDatabasesInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetRdsPgDatabasesResult>("sbercloud:index/getRdsPgDatabases:getRdsPgDatabases", args ?? new GetRdsPgDatabasesInvokeArgs(), options.WithDefaults());

        public static Output<GetRdsPgDatabasesResult> Invoke(GetRdsPgDatabasesInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetRdsPgDatabasesResult>("sbercloud:index/getRdsPgDatabases:getRdsPgDatabases", args ?? new GetRdsPgDatabasesInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetRdsPgDatabasesArgs : global::Pulumi.InvokeArgs
    {
        [Input("characterSet")]
        public string? CharacterSet { get; set; }

        [Input("instanceId", required: true)]
        public string InstanceId { get; set; } = null!;

        [Input("lcCollate")]
        public string? LcCollate { get; set; }

        [Input("name")]
        public string? Name { get; set; }

        [Input("owner")]
        public string? Owner { get; set; }

        [Input("region")]
        public string? Region { get; set; }

        [Input("size")]
        public int? Size { get; set; }

        public GetRdsPgDatabasesArgs()
        {
        }
        public static new GetRdsPgDatabasesArgs Empty => new GetRdsPgDatabasesArgs();
    }

    public sealed class GetRdsPgDatabasesInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("characterSet")]
        public Input<string>? CharacterSet { get; set; }

        [Input("instanceId", required: true)]
        public Input<string> InstanceId { get; set; } = null!;

        [Input("lcCollate")]
        public Input<string>? LcCollate { get; set; }

        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("owner")]
        public Input<string>? Owner { get; set; }

        [Input("region")]
        public Input<string>? Region { get; set; }

        [Input("size")]
        public Input<int>? Size { get; set; }

        public GetRdsPgDatabasesInvokeArgs()
        {
        }
        public static new GetRdsPgDatabasesInvokeArgs Empty => new GetRdsPgDatabasesInvokeArgs();
    }


    [OutputType]
    public sealed class GetRdsPgDatabasesResult
    {
        public readonly string? CharacterSet;
        public readonly ImmutableArray<Outputs.GetRdsPgDatabasesDatabaseResult> Databases;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string InstanceId;
        public readonly string? LcCollate;
        public readonly string? Name;
        public readonly string? Owner;
        public readonly string Region;
        public readonly int? Size;

        [OutputConstructor]
        private GetRdsPgDatabasesResult(
            string? characterSet,

            ImmutableArray<Outputs.GetRdsPgDatabasesDatabaseResult> databases,

            string id,

            string instanceId,

            string? lcCollate,

            string? name,

            string? owner,

            string region,

            int? size)
        {
            CharacterSet = characterSet;
            Databases = databases;
            Id = id;
            InstanceId = instanceId;
            LcCollate = lcCollate;
            Name = name;
            Owner = owner;
            Region = region;
            Size = size;
        }
    }
}
