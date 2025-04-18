// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Sbercloud
{
    [SbercloudResourceType("sbercloud:index/computeInterfaceAttach:ComputeInterfaceAttach")]
    public partial class ComputeInterfaceAttach : global::Pulumi.CustomResource
    {
        [Output("fixedIp")]
        public Output<string> FixedIp { get; private set; } = null!;

        [Output("fixedIpv6")]
        public Output<string> FixedIpv6 { get; private set; } = null!;

        [Output("instanceId")]
        public Output<string> InstanceId { get; private set; } = null!;

        [Output("ipv6BandwidthId")]
        public Output<string?> Ipv6BandwidthId { get; private set; } = null!;

        [Output("ipv6Enable")]
        public Output<bool> Ipv6Enable { get; private set; } = null!;

        [Output("mac")]
        public Output<string> Mac { get; private set; } = null!;

        [Output("networkId")]
        public Output<string> NetworkId { get; private set; } = null!;

        [Output("portId")]
        public Output<string> PortId { get; private set; } = null!;

        [Output("region")]
        public Output<string> Region { get; private set; } = null!;

        [Output("securityGroupIds")]
        public Output<ImmutableArray<string>> SecurityGroupIds { get; private set; } = null!;

        [Output("sourceDestCheck")]
        public Output<bool?> SourceDestCheck { get; private set; } = null!;


        /// <summary>
        /// Create a ComputeInterfaceAttach resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ComputeInterfaceAttach(string name, ComputeInterfaceAttachArgs args, CustomResourceOptions? options = null)
            : base("sbercloud:index/computeInterfaceAttach:ComputeInterfaceAttach", name, args ?? new ComputeInterfaceAttachArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ComputeInterfaceAttach(string name, Input<string> id, ComputeInterfaceAttachState? state = null, CustomResourceOptions? options = null)
            : base("sbercloud:index/computeInterfaceAttach:ComputeInterfaceAttach", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing ComputeInterfaceAttach resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ComputeInterfaceAttach Get(string name, Input<string> id, ComputeInterfaceAttachState? state = null, CustomResourceOptions? options = null)
        {
            return new ComputeInterfaceAttach(name, id, state, options);
        }
    }

    public sealed class ComputeInterfaceAttachArgs : global::Pulumi.ResourceArgs
    {
        [Input("fixedIp")]
        public Input<string>? FixedIp { get; set; }

        [Input("instanceId", required: true)]
        public Input<string> InstanceId { get; set; } = null!;

        [Input("ipv6BandwidthId")]
        public Input<string>? Ipv6BandwidthId { get; set; }

        [Input("ipv6Enable")]
        public Input<bool>? Ipv6Enable { get; set; }

        [Input("networkId")]
        public Input<string>? NetworkId { get; set; }

        [Input("portId")]
        public Input<string>? PortId { get; set; }

        [Input("region")]
        public Input<string>? Region { get; set; }

        [Input("securityGroupIds")]
        private InputList<string>? _securityGroupIds;
        public InputList<string> SecurityGroupIds
        {
            get => _securityGroupIds ?? (_securityGroupIds = new InputList<string>());
            set => _securityGroupIds = value;
        }

        [Input("sourceDestCheck")]
        public Input<bool>? SourceDestCheck { get; set; }

        public ComputeInterfaceAttachArgs()
        {
        }
        public static new ComputeInterfaceAttachArgs Empty => new ComputeInterfaceAttachArgs();
    }

    public sealed class ComputeInterfaceAttachState : global::Pulumi.ResourceArgs
    {
        [Input("fixedIp")]
        public Input<string>? FixedIp { get; set; }

        [Input("fixedIpv6")]
        public Input<string>? FixedIpv6 { get; set; }

        [Input("instanceId")]
        public Input<string>? InstanceId { get; set; }

        [Input("ipv6BandwidthId")]
        public Input<string>? Ipv6BandwidthId { get; set; }

        [Input("ipv6Enable")]
        public Input<bool>? Ipv6Enable { get; set; }

        [Input("mac")]
        public Input<string>? Mac { get; set; }

        [Input("networkId")]
        public Input<string>? NetworkId { get; set; }

        [Input("portId")]
        public Input<string>? PortId { get; set; }

        [Input("region")]
        public Input<string>? Region { get; set; }

        [Input("securityGroupIds")]
        private InputList<string>? _securityGroupIds;
        public InputList<string> SecurityGroupIds
        {
            get => _securityGroupIds ?? (_securityGroupIds = new InputList<string>());
            set => _securityGroupIds = value;
        }

        [Input("sourceDestCheck")]
        public Input<bool>? SourceDestCheck { get; set; }

        public ComputeInterfaceAttachState()
        {
        }
        public static new ComputeInterfaceAttachState Empty => new ComputeInterfaceAttachState();
    }
}
