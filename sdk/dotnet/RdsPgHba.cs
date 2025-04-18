// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Sbercloud
{
    [SbercloudResourceType("sbercloud:index/rdsPgHba:RdsPgHba")]
    public partial class RdsPgHba : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Specifies the list of host based authentications.
        /// </summary>
        [Output("hostBasedAuthentications")]
        public Output<ImmutableArray<Outputs.RdsPgHbaHostBasedAuthentication>> HostBasedAuthentications { get; private set; } = null!;

        /// <summary>
        /// Specifies the ID of the RDS PostgreSQL instance.
        /// </summary>
        [Output("instanceId")]
        public Output<string> InstanceId { get; private set; } = null!;

        [Output("region")]
        public Output<string> Region { get; private set; } = null!;


        /// <summary>
        /// Create a RdsPgHba resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public RdsPgHba(string name, RdsPgHbaArgs args, CustomResourceOptions? options = null)
            : base("sbercloud:index/rdsPgHba:RdsPgHba", name, args ?? new RdsPgHbaArgs(), MakeResourceOptions(options, ""))
        {
        }

        private RdsPgHba(string name, Input<string> id, RdsPgHbaState? state = null, CustomResourceOptions? options = null)
            : base("sbercloud:index/rdsPgHba:RdsPgHba", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing RdsPgHba resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static RdsPgHba Get(string name, Input<string> id, RdsPgHbaState? state = null, CustomResourceOptions? options = null)
        {
            return new RdsPgHba(name, id, state, options);
        }
    }

    public sealed class RdsPgHbaArgs : global::Pulumi.ResourceArgs
    {
        [Input("hostBasedAuthentications", required: true)]
        private InputList<Inputs.RdsPgHbaHostBasedAuthenticationArgs>? _hostBasedAuthentications;

        /// <summary>
        /// Specifies the list of host based authentications.
        /// </summary>
        public InputList<Inputs.RdsPgHbaHostBasedAuthenticationArgs> HostBasedAuthentications
        {
            get => _hostBasedAuthentications ?? (_hostBasedAuthentications = new InputList<Inputs.RdsPgHbaHostBasedAuthenticationArgs>());
            set => _hostBasedAuthentications = value;
        }

        /// <summary>
        /// Specifies the ID of the RDS PostgreSQL instance.
        /// </summary>
        [Input("instanceId", required: true)]
        public Input<string> InstanceId { get; set; } = null!;

        [Input("region")]
        public Input<string>? Region { get; set; }

        public RdsPgHbaArgs()
        {
        }
        public static new RdsPgHbaArgs Empty => new RdsPgHbaArgs();
    }

    public sealed class RdsPgHbaState : global::Pulumi.ResourceArgs
    {
        [Input("hostBasedAuthentications")]
        private InputList<Inputs.RdsPgHbaHostBasedAuthenticationGetArgs>? _hostBasedAuthentications;

        /// <summary>
        /// Specifies the list of host based authentications.
        /// </summary>
        public InputList<Inputs.RdsPgHbaHostBasedAuthenticationGetArgs> HostBasedAuthentications
        {
            get => _hostBasedAuthentications ?? (_hostBasedAuthentications = new InputList<Inputs.RdsPgHbaHostBasedAuthenticationGetArgs>());
            set => _hostBasedAuthentications = value;
        }

        /// <summary>
        /// Specifies the ID of the RDS PostgreSQL instance.
        /// </summary>
        [Input("instanceId")]
        public Input<string>? InstanceId { get; set; }

        [Input("region")]
        public Input<string>? Region { get; set; }

        public RdsPgHbaState()
        {
        }
        public static new RdsPgHbaState Empty => new RdsPgHbaState();
    }
}
