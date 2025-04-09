// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Sbercloud
{
    [SbercloudResourceType("sbercloud:index/dliQueue:DliQueue")]
    public partial class DliQueue : global::Pulumi.CustomResource
    {
        [Output("createTime")]
        public Output<int> CreateTime { get; private set; } = null!;

        [Output("cuCount")]
        public Output<int> CuCount { get; private set; } = null!;

        [Output("description")]
        public Output<string> Description { get; private set; } = null!;

        /// <summary>
        /// The name of the elastic resource pool to which the queue belongs.
        /// </summary>
        [Output("elasticResourcePoolName")]
        public Output<string> ElasticResourcePoolName { get; private set; } = null!;

        [Output("enterpriseProjectId")]
        public Output<string> EnterpriseProjectId { get; private set; } = null!;

        [Output("feature")]
        public Output<string?> Feature { get; private set; } = null!;

        [Output("managementSubnetCidr")]
        public Output<string?> ManagementSubnetCidr { get; private set; } = null!;

        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        [Output("platform")]
        public Output<string?> Platform { get; private set; } = null!;

        [Output("queueType")]
        public Output<string?> QueueType { get; private set; } = null!;

        [Output("region")]
        public Output<string> Region { get; private set; } = null!;

        /// <summary>
        /// The queue resource mode.
        /// </summary>
        [Output("resourceMode")]
        public Output<int> ResourceMode { get; private set; } = null!;

        [Output("scalingPolicies")]
        public Output<ImmutableArray<Outputs.DliQueueScalingPolicy>> ScalingPolicies { get; private set; } = null!;

        [Output("sparkDriver")]
        public Output<Outputs.DliQueueSparkDriver?> SparkDriver { get; private set; } = null!;

        [Output("subnetCidr")]
        public Output<string?> SubnetCidr { get; private set; } = null!;

        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;

        /// <summary>
        /// The CIDR block of the queue.
        /// </summary>
        [Output("vpcCidr")]
        public Output<string> VpcCidr { get; private set; } = null!;


        /// <summary>
        /// Create a DliQueue resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public DliQueue(string name, DliQueueArgs args, CustomResourceOptions? options = null)
            : base("sbercloud:index/dliQueue:DliQueue", name, args ?? new DliQueueArgs(), MakeResourceOptions(options, ""))
        {
        }

        private DliQueue(string name, Input<string> id, DliQueueState? state = null, CustomResourceOptions? options = null)
            : base("sbercloud:index/dliQueue:DliQueue", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing DliQueue resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static DliQueue Get(string name, Input<string> id, DliQueueState? state = null, CustomResourceOptions? options = null)
        {
            return new DliQueue(name, id, state, options);
        }
    }

    public sealed class DliQueueArgs : global::Pulumi.ResourceArgs
    {
        [Input("cuCount", required: true)]
        public Input<int> CuCount { get; set; } = null!;

        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The name of the elastic resource pool to which the queue belongs.
        /// </summary>
        [Input("elasticResourcePoolName")]
        public Input<string>? ElasticResourcePoolName { get; set; }

        [Input("enterpriseProjectId")]
        public Input<string>? EnterpriseProjectId { get; set; }

        [Input("feature")]
        public Input<string>? Feature { get; set; }

        [Input("managementSubnetCidr")]
        public Input<string>? ManagementSubnetCidr { get; set; }

        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("platform")]
        public Input<string>? Platform { get; set; }

        [Input("queueType")]
        public Input<string>? QueueType { get; set; }

        [Input("region")]
        public Input<string>? Region { get; set; }

        /// <summary>
        /// The queue resource mode.
        /// </summary>
        [Input("resourceMode")]
        public Input<int>? ResourceMode { get; set; }

        [Input("scalingPolicies")]
        private InputList<Inputs.DliQueueScalingPolicyArgs>? _scalingPolicies;
        public InputList<Inputs.DliQueueScalingPolicyArgs> ScalingPolicies
        {
            get => _scalingPolicies ?? (_scalingPolicies = new InputList<Inputs.DliQueueScalingPolicyArgs>());
            set => _scalingPolicies = value;
        }

        [Input("sparkDriver")]
        public Input<Inputs.DliQueueSparkDriverArgs>? SparkDriver { get; set; }

        [Input("subnetCidr")]
        public Input<string>? SubnetCidr { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        /// <summary>
        /// The CIDR block of the queue.
        /// </summary>
        [Input("vpcCidr")]
        public Input<string>? VpcCidr { get; set; }

        public DliQueueArgs()
        {
        }
        public static new DliQueueArgs Empty => new DliQueueArgs();
    }

    public sealed class DliQueueState : global::Pulumi.ResourceArgs
    {
        [Input("createTime")]
        public Input<int>? CreateTime { get; set; }

        [Input("cuCount")]
        public Input<int>? CuCount { get; set; }

        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The name of the elastic resource pool to which the queue belongs.
        /// </summary>
        [Input("elasticResourcePoolName")]
        public Input<string>? ElasticResourcePoolName { get; set; }

        [Input("enterpriseProjectId")]
        public Input<string>? EnterpriseProjectId { get; set; }

        [Input("feature")]
        public Input<string>? Feature { get; set; }

        [Input("managementSubnetCidr")]
        public Input<string>? ManagementSubnetCidr { get; set; }

        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("platform")]
        public Input<string>? Platform { get; set; }

        [Input("queueType")]
        public Input<string>? QueueType { get; set; }

        [Input("region")]
        public Input<string>? Region { get; set; }

        /// <summary>
        /// The queue resource mode.
        /// </summary>
        [Input("resourceMode")]
        public Input<int>? ResourceMode { get; set; }

        [Input("scalingPolicies")]
        private InputList<Inputs.DliQueueScalingPolicyGetArgs>? _scalingPolicies;
        public InputList<Inputs.DliQueueScalingPolicyGetArgs> ScalingPolicies
        {
            get => _scalingPolicies ?? (_scalingPolicies = new InputList<Inputs.DliQueueScalingPolicyGetArgs>());
            set => _scalingPolicies = value;
        }

        [Input("sparkDriver")]
        public Input<Inputs.DliQueueSparkDriverGetArgs>? SparkDriver { get; set; }

        [Input("subnetCidr")]
        public Input<string>? SubnetCidr { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        /// <summary>
        /// The CIDR block of the queue.
        /// </summary>
        [Input("vpcCidr")]
        public Input<string>? VpcCidr { get; set; }

        public DliQueueState()
        {
        }
        public static new DliQueueState Empty => new DliQueueState();
    }
}
