// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Sbercloud
{
    [SbercloudResourceType("sbercloud:index/elbLoadbalancer:ElbLoadbalancer")]
    public partial class ElbLoadbalancer : global::Pulumi.CustomResource
    {
        [Output("autoPay")]
        public Output<string?> AutoPay { get; private set; } = null!;

        [Output("autoRenew")]
        public Output<string?> AutoRenew { get; private set; } = null!;

        [Output("autoscalingEnabled")]
        public Output<bool> AutoscalingEnabled { get; private set; } = null!;

        [Output("availabilityZones")]
        public Output<ImmutableArray<string>> AvailabilityZones { get; private set; } = null!;

        [Output("backendSubnets")]
        public Output<ImmutableArray<string>> BackendSubnets { get; private set; } = null!;

        [Output("bandwidthChargeMode")]
        public Output<string> BandwidthChargeMode { get; private set; } = null!;

        [Output("bandwidthId")]
        public Output<string> BandwidthId { get; private set; } = null!;

        [Output("bandwidthSize")]
        public Output<int> BandwidthSize { get; private set; } = null!;

        [Output("chargeMode")]
        public Output<string> ChargeMode { get; private set; } = null!;

        [Output("chargingMode")]
        public Output<string> ChargingMode { get; private set; } = null!;

        [Output("createdAt")]
        public Output<string> CreatedAt { get; private set; } = null!;

        [Output("crossVpcBackend")]
        public Output<bool> CrossVpcBackend { get; private set; } = null!;

        [Output("deletionProtectionEnable")]
        public Output<bool?> DeletionProtectionEnable { get; private set; } = null!;

        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        [Output("enterpriseProjectId")]
        public Output<string> EnterpriseProjectId { get; private set; } = null!;

        [Output("forceDelete")]
        public Output<bool?> ForceDelete { get; private set; } = null!;

        [Output("guaranteed")]
        public Output<bool> Guaranteed { get; private set; } = null!;

        [Output("gwFlavorId")]
        public Output<string> GwFlavorId { get; private set; } = null!;

        [Output("iptype")]
        public Output<string> Iptype { get; private set; } = null!;

        [Output("ipv4Address")]
        public Output<string> Ipv4Address { get; private set; } = null!;

        [Output("ipv4Eip")]
        public Output<string> Ipv4Eip { get; private set; } = null!;

        [Output("ipv4EipId")]
        public Output<string> Ipv4EipId { get; private set; } = null!;

        [Output("ipv4PortId")]
        public Output<string> Ipv4PortId { get; private set; } = null!;

        /// <summary>
        /// the IPv4 subnet ID of the subnet where the load balancer resides
        /// </summary>
        [Output("ipv4SubnetId")]
        public Output<string?> Ipv4SubnetId { get; private set; } = null!;

        [Output("ipv6Address")]
        public Output<string> Ipv6Address { get; private set; } = null!;

        [Output("ipv6BandwidthId")]
        public Output<string?> Ipv6BandwidthId { get; private set; } = null!;

        [Output("ipv6Eip")]
        public Output<string> Ipv6Eip { get; private set; } = null!;

        [Output("ipv6EipId")]
        public Output<string> Ipv6EipId { get; private set; } = null!;

        /// <summary>
        /// the ID of the subnet where the load balancer resides
        /// </summary>
        [Output("ipv6NetworkId")]
        public Output<string?> Ipv6NetworkId { get; private set; } = null!;

        [Output("l4FlavorId")]
        public Output<string> L4FlavorId { get; private set; } = null!;

        [Output("l7FlavorId")]
        public Output<string> L7FlavorId { get; private set; } = null!;

        [Output("loadbalancerType")]
        public Output<string> LoadbalancerType { get; private set; } = null!;

        [Output("minL7FlavorId")]
        public Output<string> MinL7FlavorId { get; private set; } = null!;

        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        [Output("period")]
        public Output<int?> Period { get; private set; } = null!;

        [Output("periodUnit")]
        public Output<string?> PeriodUnit { get; private set; } = null!;

        [Output("protectionReason")]
        public Output<string?> ProtectionReason { get; private set; } = null!;

        [Output("protectionStatus")]
        public Output<string> ProtectionStatus { get; private set; } = null!;

        [Output("region")]
        public Output<string> Region { get; private set; } = null!;

        [Output("sharetype")]
        public Output<string> Sharetype { get; private set; } = null!;

        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;

        [Output("updatedAt")]
        public Output<string> UpdatedAt { get; private set; } = null!;

        [Output("vpcId")]
        public Output<string> VpcId { get; private set; } = null!;

        [Output("wafFailureAction")]
        public Output<string> WafFailureAction { get; private set; } = null!;


        /// <summary>
        /// Create a ElbLoadbalancer resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ElbLoadbalancer(string name, ElbLoadbalancerArgs args, CustomResourceOptions? options = null)
            : base("sbercloud:index/elbLoadbalancer:ElbLoadbalancer", name, args ?? new ElbLoadbalancerArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ElbLoadbalancer(string name, Input<string> id, ElbLoadbalancerState? state = null, CustomResourceOptions? options = null)
            : base("sbercloud:index/elbLoadbalancer:ElbLoadbalancer", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing ElbLoadbalancer resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ElbLoadbalancer Get(string name, Input<string> id, ElbLoadbalancerState? state = null, CustomResourceOptions? options = null)
        {
            return new ElbLoadbalancer(name, id, state, options);
        }
    }

    public sealed class ElbLoadbalancerArgs : global::Pulumi.ResourceArgs
    {
        [Input("autoPay")]
        public Input<string>? AutoPay { get; set; }

        [Input("autoRenew")]
        public Input<string>? AutoRenew { get; set; }

        [Input("autoscalingEnabled")]
        public Input<bool>? AutoscalingEnabled { get; set; }

        [Input("availabilityZones", required: true)]
        private InputList<string>? _availabilityZones;
        public InputList<string> AvailabilityZones
        {
            get => _availabilityZones ?? (_availabilityZones = new InputList<string>());
            set => _availabilityZones = value;
        }

        [Input("backendSubnets")]
        private InputList<string>? _backendSubnets;
        public InputList<string> BackendSubnets
        {
            get => _backendSubnets ?? (_backendSubnets = new InputList<string>());
            set => _backendSubnets = value;
        }

        [Input("bandwidthChargeMode")]
        public Input<string>? BandwidthChargeMode { get; set; }

        [Input("bandwidthId")]
        public Input<string>? BandwidthId { get; set; }

        [Input("bandwidthSize")]
        public Input<int>? BandwidthSize { get; set; }

        [Input("chargingMode")]
        public Input<string>? ChargingMode { get; set; }

        [Input("crossVpcBackend")]
        public Input<bool>? CrossVpcBackend { get; set; }

        [Input("deletionProtectionEnable")]
        public Input<bool>? DeletionProtectionEnable { get; set; }

        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("enterpriseProjectId")]
        public Input<string>? EnterpriseProjectId { get; set; }

        [Input("forceDelete")]
        public Input<bool>? ForceDelete { get; set; }

        [Input("iptype")]
        public Input<string>? Iptype { get; set; }

        [Input("ipv4Address")]
        public Input<string>? Ipv4Address { get; set; }

        [Input("ipv4EipId")]
        public Input<string>? Ipv4EipId { get; set; }

        /// <summary>
        /// the IPv4 subnet ID of the subnet where the load balancer resides
        /// </summary>
        [Input("ipv4SubnetId")]
        public Input<string>? Ipv4SubnetId { get; set; }

        [Input("ipv6Address")]
        public Input<string>? Ipv6Address { get; set; }

        [Input("ipv6BandwidthId")]
        public Input<string>? Ipv6BandwidthId { get; set; }

        /// <summary>
        /// the ID of the subnet where the load balancer resides
        /// </summary>
        [Input("ipv6NetworkId")]
        public Input<string>? Ipv6NetworkId { get; set; }

        [Input("l4FlavorId")]
        public Input<string>? L4FlavorId { get; set; }

        [Input("l7FlavorId")]
        public Input<string>? L7FlavorId { get; set; }

        [Input("loadbalancerType")]
        public Input<string>? LoadbalancerType { get; set; }

        [Input("minL7FlavorId")]
        public Input<string>? MinL7FlavorId { get; set; }

        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("period")]
        public Input<int>? Period { get; set; }

        [Input("periodUnit")]
        public Input<string>? PeriodUnit { get; set; }

        [Input("protectionReason")]
        public Input<string>? ProtectionReason { get; set; }

        [Input("protectionStatus")]
        public Input<string>? ProtectionStatus { get; set; }

        [Input("region")]
        public Input<string>? Region { get; set; }

        [Input("sharetype")]
        public Input<string>? Sharetype { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        [Input("vpcId")]
        public Input<string>? VpcId { get; set; }

        [Input("wafFailureAction")]
        public Input<string>? WafFailureAction { get; set; }

        public ElbLoadbalancerArgs()
        {
        }
        public static new ElbLoadbalancerArgs Empty => new ElbLoadbalancerArgs();
    }

    public sealed class ElbLoadbalancerState : global::Pulumi.ResourceArgs
    {
        [Input("autoPay")]
        public Input<string>? AutoPay { get; set; }

        [Input("autoRenew")]
        public Input<string>? AutoRenew { get; set; }

        [Input("autoscalingEnabled")]
        public Input<bool>? AutoscalingEnabled { get; set; }

        [Input("availabilityZones")]
        private InputList<string>? _availabilityZones;
        public InputList<string> AvailabilityZones
        {
            get => _availabilityZones ?? (_availabilityZones = new InputList<string>());
            set => _availabilityZones = value;
        }

        [Input("backendSubnets")]
        private InputList<string>? _backendSubnets;
        public InputList<string> BackendSubnets
        {
            get => _backendSubnets ?? (_backendSubnets = new InputList<string>());
            set => _backendSubnets = value;
        }

        [Input("bandwidthChargeMode")]
        public Input<string>? BandwidthChargeMode { get; set; }

        [Input("bandwidthId")]
        public Input<string>? BandwidthId { get; set; }

        [Input("bandwidthSize")]
        public Input<int>? BandwidthSize { get; set; }

        [Input("chargeMode")]
        public Input<string>? ChargeMode { get; set; }

        [Input("chargingMode")]
        public Input<string>? ChargingMode { get; set; }

        [Input("createdAt")]
        public Input<string>? CreatedAt { get; set; }

        [Input("crossVpcBackend")]
        public Input<bool>? CrossVpcBackend { get; set; }

        [Input("deletionProtectionEnable")]
        public Input<bool>? DeletionProtectionEnable { get; set; }

        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("enterpriseProjectId")]
        public Input<string>? EnterpriseProjectId { get; set; }

        [Input("forceDelete")]
        public Input<bool>? ForceDelete { get; set; }

        [Input("guaranteed")]
        public Input<bool>? Guaranteed { get; set; }

        [Input("gwFlavorId")]
        public Input<string>? GwFlavorId { get; set; }

        [Input("iptype")]
        public Input<string>? Iptype { get; set; }

        [Input("ipv4Address")]
        public Input<string>? Ipv4Address { get; set; }

        [Input("ipv4Eip")]
        public Input<string>? Ipv4Eip { get; set; }

        [Input("ipv4EipId")]
        public Input<string>? Ipv4EipId { get; set; }

        [Input("ipv4PortId")]
        public Input<string>? Ipv4PortId { get; set; }

        /// <summary>
        /// the IPv4 subnet ID of the subnet where the load balancer resides
        /// </summary>
        [Input("ipv4SubnetId")]
        public Input<string>? Ipv4SubnetId { get; set; }

        [Input("ipv6Address")]
        public Input<string>? Ipv6Address { get; set; }

        [Input("ipv6BandwidthId")]
        public Input<string>? Ipv6BandwidthId { get; set; }

        [Input("ipv6Eip")]
        public Input<string>? Ipv6Eip { get; set; }

        [Input("ipv6EipId")]
        public Input<string>? Ipv6EipId { get; set; }

        /// <summary>
        /// the ID of the subnet where the load balancer resides
        /// </summary>
        [Input("ipv6NetworkId")]
        public Input<string>? Ipv6NetworkId { get; set; }

        [Input("l4FlavorId")]
        public Input<string>? L4FlavorId { get; set; }

        [Input("l7FlavorId")]
        public Input<string>? L7FlavorId { get; set; }

        [Input("loadbalancerType")]
        public Input<string>? LoadbalancerType { get; set; }

        [Input("minL7FlavorId")]
        public Input<string>? MinL7FlavorId { get; set; }

        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("period")]
        public Input<int>? Period { get; set; }

        [Input("periodUnit")]
        public Input<string>? PeriodUnit { get; set; }

        [Input("protectionReason")]
        public Input<string>? ProtectionReason { get; set; }

        [Input("protectionStatus")]
        public Input<string>? ProtectionStatus { get; set; }

        [Input("region")]
        public Input<string>? Region { get; set; }

        [Input("sharetype")]
        public Input<string>? Sharetype { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        [Input("updatedAt")]
        public Input<string>? UpdatedAt { get; set; }

        [Input("vpcId")]
        public Input<string>? VpcId { get; set; }

        [Input("wafFailureAction")]
        public Input<string>? WafFailureAction { get; set; }

        public ElbLoadbalancerState()
        {
        }
        public static new ElbLoadbalancerState Empty => new ElbLoadbalancerState();
    }
}
