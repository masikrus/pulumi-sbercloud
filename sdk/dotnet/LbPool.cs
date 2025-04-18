// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Sbercloud
{
    [SbercloudResourceType("sbercloud:index/lbPool:LbPool")]
    public partial class LbPool : global::Pulumi.CustomResource
    {
        [Output("adminStateUp")]
        public Output<bool?> AdminStateUp { get; private set; } = null!;

        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        [Output("lbMethod")]
        public Output<string> LbMethod { get; private set; } = null!;

        [Output("listenerId")]
        public Output<string> ListenerId { get; private set; } = null!;

        [Output("loadbalancerId")]
        public Output<string> LoadbalancerId { get; private set; } = null!;

        [Output("monitorId")]
        public Output<string> MonitorId { get; private set; } = null!;

        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        [Output("persistences")]
        public Output<ImmutableArray<Outputs.LbPoolPersistence>> Persistences { get; private set; } = null!;

        [Output("protectionReason")]
        public Output<string?> ProtectionReason { get; private set; } = null!;

        [Output("protectionStatus")]
        public Output<string> ProtectionStatus { get; private set; } = null!;

        [Output("protocol")]
        public Output<string> Protocol { get; private set; } = null!;

        [Output("region")]
        public Output<string> Region { get; private set; } = null!;

        [Output("tenantId")]
        public Output<string> TenantId { get; private set; } = null!;


        /// <summary>
        /// Create a LbPool resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public LbPool(string name, LbPoolArgs args, CustomResourceOptions? options = null)
            : base("sbercloud:index/lbPool:LbPool", name, args ?? new LbPoolArgs(), MakeResourceOptions(options, ""))
        {
        }

        private LbPool(string name, Input<string> id, LbPoolState? state = null, CustomResourceOptions? options = null)
            : base("sbercloud:index/lbPool:LbPool", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing LbPool resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static LbPool Get(string name, Input<string> id, LbPoolState? state = null, CustomResourceOptions? options = null)
        {
            return new LbPool(name, id, state, options);
        }
    }

    public sealed class LbPoolArgs : global::Pulumi.ResourceArgs
    {
        [Input("adminStateUp")]
        public Input<bool>? AdminStateUp { get; set; }

        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("lbMethod", required: true)]
        public Input<string> LbMethod { get; set; } = null!;

        [Input("listenerId")]
        public Input<string>? ListenerId { get; set; }

        [Input("loadbalancerId")]
        public Input<string>? LoadbalancerId { get; set; }

        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("persistences")]
        private InputList<Inputs.LbPoolPersistenceArgs>? _persistences;
        public InputList<Inputs.LbPoolPersistenceArgs> Persistences
        {
            get => _persistences ?? (_persistences = new InputList<Inputs.LbPoolPersistenceArgs>());
            set => _persistences = value;
        }

        [Input("protectionReason")]
        public Input<string>? ProtectionReason { get; set; }

        [Input("protectionStatus")]
        public Input<string>? ProtectionStatus { get; set; }

        [Input("protocol", required: true)]
        public Input<string> Protocol { get; set; } = null!;

        [Input("region")]
        public Input<string>? Region { get; set; }

        [Input("tenantId")]
        public Input<string>? TenantId { get; set; }

        public LbPoolArgs()
        {
        }
        public static new LbPoolArgs Empty => new LbPoolArgs();
    }

    public sealed class LbPoolState : global::Pulumi.ResourceArgs
    {
        [Input("adminStateUp")]
        public Input<bool>? AdminStateUp { get; set; }

        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("lbMethod")]
        public Input<string>? LbMethod { get; set; }

        [Input("listenerId")]
        public Input<string>? ListenerId { get; set; }

        [Input("loadbalancerId")]
        public Input<string>? LoadbalancerId { get; set; }

        [Input("monitorId")]
        public Input<string>? MonitorId { get; set; }

        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("persistences")]
        private InputList<Inputs.LbPoolPersistenceGetArgs>? _persistences;
        public InputList<Inputs.LbPoolPersistenceGetArgs> Persistences
        {
            get => _persistences ?? (_persistences = new InputList<Inputs.LbPoolPersistenceGetArgs>());
            set => _persistences = value;
        }

        [Input("protectionReason")]
        public Input<string>? ProtectionReason { get; set; }

        [Input("protectionStatus")]
        public Input<string>? ProtectionStatus { get; set; }

        [Input("protocol")]
        public Input<string>? Protocol { get; set; }

        [Input("region")]
        public Input<string>? Region { get; set; }

        [Input("tenantId")]
        public Input<string>? TenantId { get; set; }

        public LbPoolState()
        {
        }
        public static new LbPoolState Empty => new LbPoolState();
    }
}
