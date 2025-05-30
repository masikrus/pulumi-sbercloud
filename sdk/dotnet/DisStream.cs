// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Sbercloud
{
    [SbercloudResourceType("sbercloud:index/disStream:DisStream")]
    public partial class DisStream : global::Pulumi.CustomResource
    {
        [Output("autoScaleMaxPartitionCount")]
        public Output<int> AutoScaleMaxPartitionCount { get; private set; } = null!;

        [Output("autoScaleMinPartitionCount")]
        public Output<int> AutoScaleMinPartitionCount { get; private set; } = null!;

        [Output("compressionFormat")]
        public Output<string> CompressionFormat { get; private set; } = null!;

        [Output("created")]
        public Output<int> Created { get; private set; } = null!;

        [Output("csvDelimiter")]
        public Output<string> CsvDelimiter { get; private set; } = null!;

        [Output("dataSchema")]
        public Output<string> DataSchema { get; private set; } = null!;

        [Output("dataType")]
        public Output<string> DataType { get; private set; } = null!;

        [Output("enterpriseProjectId")]
        public Output<string> EnterpriseProjectId { get; private set; } = null!;

        [Output("partitionCount")]
        public Output<int> PartitionCount { get; private set; } = null!;

        [Output("partitions")]
        public Output<ImmutableArray<Outputs.DisStreamPartition>> Partitions { get; private set; } = null!;

        [Output("readablePartitionCount")]
        public Output<int> ReadablePartitionCount { get; private set; } = null!;

        [Output("region")]
        public Output<string> Region { get; private set; } = null!;

        [Output("retentionPeriod")]
        public Output<int?> RetentionPeriod { get; private set; } = null!;

        [Output("status")]
        public Output<string> Status { get; private set; } = null!;

        [Output("streamId")]
        public Output<string> StreamId { get; private set; } = null!;

        [Output("streamName")]
        public Output<string> StreamName { get; private set; } = null!;

        [Output("streamType")]
        public Output<string> StreamType { get; private set; } = null!;

        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;

        [Output("writablePartitionCount")]
        public Output<int> WritablePartitionCount { get; private set; } = null!;


        /// <summary>
        /// Create a DisStream resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public DisStream(string name, DisStreamArgs args, CustomResourceOptions? options = null)
            : base("sbercloud:index/disStream:DisStream", name, args ?? new DisStreamArgs(), MakeResourceOptions(options, ""))
        {
        }

        private DisStream(string name, Input<string> id, DisStreamState? state = null, CustomResourceOptions? options = null)
            : base("sbercloud:index/disStream:DisStream", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing DisStream resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static DisStream Get(string name, Input<string> id, DisStreamState? state = null, CustomResourceOptions? options = null)
        {
            return new DisStream(name, id, state, options);
        }
    }

    public sealed class DisStreamArgs : global::Pulumi.ResourceArgs
    {
        [Input("autoScaleMaxPartitionCount")]
        public Input<int>? AutoScaleMaxPartitionCount { get; set; }

        [Input("autoScaleMinPartitionCount")]
        public Input<int>? AutoScaleMinPartitionCount { get; set; }

        [Input("compressionFormat")]
        public Input<string>? CompressionFormat { get; set; }

        [Input("csvDelimiter")]
        public Input<string>? CsvDelimiter { get; set; }

        [Input("dataSchema")]
        public Input<string>? DataSchema { get; set; }

        [Input("dataType")]
        public Input<string>? DataType { get; set; }

        [Input("enterpriseProjectId")]
        public Input<string>? EnterpriseProjectId { get; set; }

        [Input("partitionCount", required: true)]
        public Input<int> PartitionCount { get; set; } = null!;

        [Input("region")]
        public Input<string>? Region { get; set; }

        [Input("retentionPeriod")]
        public Input<int>? RetentionPeriod { get; set; }

        [Input("streamName", required: true)]
        public Input<string> StreamName { get; set; } = null!;

        [Input("streamType")]
        public Input<string>? StreamType { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        public DisStreamArgs()
        {
        }
        public static new DisStreamArgs Empty => new DisStreamArgs();
    }

    public sealed class DisStreamState : global::Pulumi.ResourceArgs
    {
        [Input("autoScaleMaxPartitionCount")]
        public Input<int>? AutoScaleMaxPartitionCount { get; set; }

        [Input("autoScaleMinPartitionCount")]
        public Input<int>? AutoScaleMinPartitionCount { get; set; }

        [Input("compressionFormat")]
        public Input<string>? CompressionFormat { get; set; }

        [Input("created")]
        public Input<int>? Created { get; set; }

        [Input("csvDelimiter")]
        public Input<string>? CsvDelimiter { get; set; }

        [Input("dataSchema")]
        public Input<string>? DataSchema { get; set; }

        [Input("dataType")]
        public Input<string>? DataType { get; set; }

        [Input("enterpriseProjectId")]
        public Input<string>? EnterpriseProjectId { get; set; }

        [Input("partitionCount")]
        public Input<int>? PartitionCount { get; set; }

        [Input("partitions")]
        private InputList<Inputs.DisStreamPartitionGetArgs>? _partitions;
        public InputList<Inputs.DisStreamPartitionGetArgs> Partitions
        {
            get => _partitions ?? (_partitions = new InputList<Inputs.DisStreamPartitionGetArgs>());
            set => _partitions = value;
        }

        [Input("readablePartitionCount")]
        public Input<int>? ReadablePartitionCount { get; set; }

        [Input("region")]
        public Input<string>? Region { get; set; }

        [Input("retentionPeriod")]
        public Input<int>? RetentionPeriod { get; set; }

        [Input("status")]
        public Input<string>? Status { get; set; }

        [Input("streamId")]
        public Input<string>? StreamId { get; set; }

        [Input("streamName")]
        public Input<string>? StreamName { get; set; }

        [Input("streamType")]
        public Input<string>? StreamType { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        [Input("writablePartitionCount")]
        public Input<int>? WritablePartitionCount { get; set; }

        public DisStreamState()
        {
        }
        public static new DisStreamState Empty => new DisStreamState();
    }
}
