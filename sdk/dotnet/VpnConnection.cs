// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Sbercloud
{
    [SbercloudResourceType("sbercloud:index/vpnConnection:VpnConnection")]
    public partial class VpnConnection : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The create time.
        /// </summary>
        [Output("createdAt")]
        public Output<string> CreatedAt { get; private set; } = null!;

        /// <summary>
        /// The customer gateway ID.
        /// </summary>
        [Output("customerGatewayId")]
        public Output<string> CustomerGatewayId { get; private set; } = null!;

        /// <summary>
        /// Whether to enable NQA check.
        /// </summary>
        [Output("enableNqa")]
        public Output<bool> EnableNqa { get; private set; } = null!;

        /// <summary>
        /// The enterprise project ID.
        /// </summary>
        [Output("enterpriseProjectId")]
        public Output<string> EnterpriseProjectId { get; private set; } = null!;

        /// <summary>
        /// The VPN gateway ID.
        /// </summary>
        [Output("gatewayId")]
        public Output<string> GatewayId { get; private set; } = null!;

        /// <summary>
        /// The VPN gateway IP ID.
        /// </summary>
        [Output("gatewayIp")]
        public Output<string> GatewayIp { get; private set; } = null!;

        [Output("haRole")]
        public Output<string> HaRole { get; private set; } = null!;

        [Output("ikepolicy")]
        public Output<Outputs.VpnConnectionIkepolicy> Ikepolicy { get; private set; } = null!;

        [Output("ipsecpolicy")]
        public Output<Outputs.VpnConnectionIpsecpolicy> Ipsecpolicy { get; private set; } = null!;

        /// <summary>
        /// The name of the VPN connection.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The customer subnets.
        /// </summary>
        [Output("peerSubnets")]
        public Output<ImmutableArray<string>> PeerSubnets { get; private set; } = null!;

        /// <summary>
        /// The policy rules. Only works when vpn_type is set to **policy**
        /// </summary>
        [Output("policyRules")]
        public Output<ImmutableArray<Outputs.VpnConnectionPolicyRule>> PolicyRules { get; private set; } = null!;

        /// <summary>
        /// The pre-shared key.
        /// </summary>
        [Output("psk")]
        public Output<string> Psk { get; private set; } = null!;

        [Output("region")]
        public Output<string> Region { get; private set; } = null!;

        /// <summary>
        /// The status of the VPN connection.
        /// </summary>
        [Output("status")]
        public Output<string> Status { get; private set; } = null!;

        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;

        /// <summary>
        /// The local tunnel address.
        /// </summary>
        [Output("tunnelLocalAddress")]
        public Output<string> TunnelLocalAddress { get; private set; } = null!;

        /// <summary>
        /// The peer tunnel address.
        /// </summary>
        [Output("tunnelPeerAddress")]
        public Output<string> TunnelPeerAddress { get; private set; } = null!;

        /// <summary>
        /// The update time.
        /// </summary>
        [Output("updatedAt")]
        public Output<string> UpdatedAt { get; private set; } = null!;

        /// <summary>
        /// The connection type. The value can be **policy**, **static** or **bgp**.
        /// </summary>
        [Output("vpnType")]
        public Output<string> VpnType { get; private set; } = null!;


        /// <summary>
        /// Create a VpnConnection resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public VpnConnection(string name, VpnConnectionArgs args, CustomResourceOptions? options = null)
            : base("sbercloud:index/vpnConnection:VpnConnection", name, args ?? new VpnConnectionArgs(), MakeResourceOptions(options, ""))
        {
        }

        private VpnConnection(string name, Input<string> id, VpnConnectionState? state = null, CustomResourceOptions? options = null)
            : base("sbercloud:index/vpnConnection:VpnConnection", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing VpnConnection resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static VpnConnection Get(string name, Input<string> id, VpnConnectionState? state = null, CustomResourceOptions? options = null)
        {
            return new VpnConnection(name, id, state, options);
        }
    }

    public sealed class VpnConnectionArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The customer gateway ID.
        /// </summary>
        [Input("customerGatewayId", required: true)]
        public Input<string> CustomerGatewayId { get; set; } = null!;

        /// <summary>
        /// Whether to enable NQA check.
        /// </summary>
        [Input("enableNqa")]
        public Input<bool>? EnableNqa { get; set; }

        /// <summary>
        /// The enterprise project ID.
        /// </summary>
        [Input("enterpriseProjectId")]
        public Input<string>? EnterpriseProjectId { get; set; }

        /// <summary>
        /// The VPN gateway ID.
        /// </summary>
        [Input("gatewayId", required: true)]
        public Input<string> GatewayId { get; set; } = null!;

        /// <summary>
        /// The VPN gateway IP ID.
        /// </summary>
        [Input("gatewayIp", required: true)]
        public Input<string> GatewayIp { get; set; } = null!;

        [Input("haRole")]
        public Input<string>? HaRole { get; set; }

        [Input("ikepolicy")]
        public Input<Inputs.VpnConnectionIkepolicyArgs>? Ikepolicy { get; set; }

        [Input("ipsecpolicy")]
        public Input<Inputs.VpnConnectionIpsecpolicyArgs>? Ipsecpolicy { get; set; }

        /// <summary>
        /// The name of the VPN connection.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("peerSubnets")]
        private InputList<string>? _peerSubnets;

        /// <summary>
        /// The customer subnets.
        /// </summary>
        public InputList<string> PeerSubnets
        {
            get => _peerSubnets ?? (_peerSubnets = new InputList<string>());
            set => _peerSubnets = value;
        }

        [Input("policyRules")]
        private InputList<Inputs.VpnConnectionPolicyRuleArgs>? _policyRules;

        /// <summary>
        /// The policy rules. Only works when vpn_type is set to **policy**
        /// </summary>
        public InputList<Inputs.VpnConnectionPolicyRuleArgs> PolicyRules
        {
            get => _policyRules ?? (_policyRules = new InputList<Inputs.VpnConnectionPolicyRuleArgs>());
            set => _policyRules = value;
        }

        /// <summary>
        /// The pre-shared key.
        /// </summary>
        [Input("psk", required: true)]
        public Input<string> Psk { get; set; } = null!;

        [Input("region")]
        public Input<string>? Region { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        /// <summary>
        /// The local tunnel address.
        /// </summary>
        [Input("tunnelLocalAddress")]
        public Input<string>? TunnelLocalAddress { get; set; }

        /// <summary>
        /// The peer tunnel address.
        /// </summary>
        [Input("tunnelPeerAddress")]
        public Input<string>? TunnelPeerAddress { get; set; }

        /// <summary>
        /// The connection type. The value can be **policy**, **static** or **bgp**.
        /// </summary>
        [Input("vpnType", required: true)]
        public Input<string> VpnType { get; set; } = null!;

        public VpnConnectionArgs()
        {
        }
        public static new VpnConnectionArgs Empty => new VpnConnectionArgs();
    }

    public sealed class VpnConnectionState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The create time.
        /// </summary>
        [Input("createdAt")]
        public Input<string>? CreatedAt { get; set; }

        /// <summary>
        /// The customer gateway ID.
        /// </summary>
        [Input("customerGatewayId")]
        public Input<string>? CustomerGatewayId { get; set; }

        /// <summary>
        /// Whether to enable NQA check.
        /// </summary>
        [Input("enableNqa")]
        public Input<bool>? EnableNqa { get; set; }

        /// <summary>
        /// The enterprise project ID.
        /// </summary>
        [Input("enterpriseProjectId")]
        public Input<string>? EnterpriseProjectId { get; set; }

        /// <summary>
        /// The VPN gateway ID.
        /// </summary>
        [Input("gatewayId")]
        public Input<string>? GatewayId { get; set; }

        /// <summary>
        /// The VPN gateway IP ID.
        /// </summary>
        [Input("gatewayIp")]
        public Input<string>? GatewayIp { get; set; }

        [Input("haRole")]
        public Input<string>? HaRole { get; set; }

        [Input("ikepolicy")]
        public Input<Inputs.VpnConnectionIkepolicyGetArgs>? Ikepolicy { get; set; }

        [Input("ipsecpolicy")]
        public Input<Inputs.VpnConnectionIpsecpolicyGetArgs>? Ipsecpolicy { get; set; }

        /// <summary>
        /// The name of the VPN connection.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("peerSubnets")]
        private InputList<string>? _peerSubnets;

        /// <summary>
        /// The customer subnets.
        /// </summary>
        public InputList<string> PeerSubnets
        {
            get => _peerSubnets ?? (_peerSubnets = new InputList<string>());
            set => _peerSubnets = value;
        }

        [Input("policyRules")]
        private InputList<Inputs.VpnConnectionPolicyRuleGetArgs>? _policyRules;

        /// <summary>
        /// The policy rules. Only works when vpn_type is set to **policy**
        /// </summary>
        public InputList<Inputs.VpnConnectionPolicyRuleGetArgs> PolicyRules
        {
            get => _policyRules ?? (_policyRules = new InputList<Inputs.VpnConnectionPolicyRuleGetArgs>());
            set => _policyRules = value;
        }

        /// <summary>
        /// The pre-shared key.
        /// </summary>
        [Input("psk")]
        public Input<string>? Psk { get; set; }

        [Input("region")]
        public Input<string>? Region { get; set; }

        /// <summary>
        /// The status of the VPN connection.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        /// <summary>
        /// The local tunnel address.
        /// </summary>
        [Input("tunnelLocalAddress")]
        public Input<string>? TunnelLocalAddress { get; set; }

        /// <summary>
        /// The peer tunnel address.
        /// </summary>
        [Input("tunnelPeerAddress")]
        public Input<string>? TunnelPeerAddress { get; set; }

        /// <summary>
        /// The update time.
        /// </summary>
        [Input("updatedAt")]
        public Input<string>? UpdatedAt { get; set; }

        /// <summary>
        /// The connection type. The value can be **policy**, **static** or **bgp**.
        /// </summary>
        [Input("vpnType")]
        public Input<string>? VpnType { get; set; }

        public VpnConnectionState()
        {
        }
        public static new VpnConnectionState Empty => new VpnConnectionState();
    }
}
