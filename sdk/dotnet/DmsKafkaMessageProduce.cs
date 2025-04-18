// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Sbercloud
{
    [SbercloudResourceType("sbercloud:index/dmsKafkaMessageProduce:DmsKafkaMessageProduce")]
    public partial class DmsKafkaMessageProduce : global::Pulumi.CustomResource
    {
        [Output("body")]
        public Output<string> Body { get; private set; } = null!;

        [Output("instanceId")]
        public Output<string> InstanceId { get; private set; } = null!;

        [Output("propertyLists")]
        public Output<ImmutableArray<Outputs.DmsKafkaMessageProducePropertyList>> PropertyLists { get; private set; } = null!;

        [Output("region")]
        public Output<string> Region { get; private set; } = null!;

        [Output("topic")]
        public Output<string> Topic { get; private set; } = null!;


        /// <summary>
        /// Create a DmsKafkaMessageProduce resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public DmsKafkaMessageProduce(string name, DmsKafkaMessageProduceArgs args, CustomResourceOptions? options = null)
            : base("sbercloud:index/dmsKafkaMessageProduce:DmsKafkaMessageProduce", name, args ?? new DmsKafkaMessageProduceArgs(), MakeResourceOptions(options, ""))
        {
        }

        private DmsKafkaMessageProduce(string name, Input<string> id, DmsKafkaMessageProduceState? state = null, CustomResourceOptions? options = null)
            : base("sbercloud:index/dmsKafkaMessageProduce:DmsKafkaMessageProduce", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing DmsKafkaMessageProduce resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static DmsKafkaMessageProduce Get(string name, Input<string> id, DmsKafkaMessageProduceState? state = null, CustomResourceOptions? options = null)
        {
            return new DmsKafkaMessageProduce(name, id, state, options);
        }
    }

    public sealed class DmsKafkaMessageProduceArgs : global::Pulumi.ResourceArgs
    {
        [Input("body", required: true)]
        public Input<string> Body { get; set; } = null!;

        [Input("instanceId", required: true)]
        public Input<string> InstanceId { get; set; } = null!;

        [Input("propertyLists")]
        private InputList<Inputs.DmsKafkaMessageProducePropertyListArgs>? _propertyLists;
        public InputList<Inputs.DmsKafkaMessageProducePropertyListArgs> PropertyLists
        {
            get => _propertyLists ?? (_propertyLists = new InputList<Inputs.DmsKafkaMessageProducePropertyListArgs>());
            set => _propertyLists = value;
        }

        [Input("region")]
        public Input<string>? Region { get; set; }

        [Input("topic", required: true)]
        public Input<string> Topic { get; set; } = null!;

        public DmsKafkaMessageProduceArgs()
        {
        }
        public static new DmsKafkaMessageProduceArgs Empty => new DmsKafkaMessageProduceArgs();
    }

    public sealed class DmsKafkaMessageProduceState : global::Pulumi.ResourceArgs
    {
        [Input("body")]
        public Input<string>? Body { get; set; }

        [Input("instanceId")]
        public Input<string>? InstanceId { get; set; }

        [Input("propertyLists")]
        private InputList<Inputs.DmsKafkaMessageProducePropertyListGetArgs>? _propertyLists;
        public InputList<Inputs.DmsKafkaMessageProducePropertyListGetArgs> PropertyLists
        {
            get => _propertyLists ?? (_propertyLists = new InputList<Inputs.DmsKafkaMessageProducePropertyListGetArgs>());
            set => _propertyLists = value;
        }

        [Input("region")]
        public Input<string>? Region { get; set; }

        [Input("topic")]
        public Input<string>? Topic { get; set; }

        public DmsKafkaMessageProduceState()
        {
        }
        public static new DmsKafkaMessageProduceState Empty => new DmsKafkaMessageProduceState();
    }
}
