// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

export class VpnConnection extends pulumi.CustomResource {
    /**
     * Get an existing VpnConnection resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: VpnConnectionState, opts?: pulumi.CustomResourceOptions): VpnConnection {
        return new VpnConnection(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'sbercloud:index/vpnConnection:VpnConnection';

    /**
     * Returns true if the given object is an instance of VpnConnection.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is VpnConnection {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === VpnConnection.__pulumiType;
    }

    /**
     * The create time.
     */
    public /*out*/ readonly createdAt!: pulumi.Output<string>;
    /**
     * The customer gateway ID.
     */
    public readonly customerGatewayId!: pulumi.Output<string>;
    /**
     * Whether to enable NQA check.
     */
    public readonly enableNqa!: pulumi.Output<boolean>;
    /**
     * The enterprise project ID.
     */
    public readonly enterpriseProjectId!: pulumi.Output<string>;
    /**
     * The VPN gateway ID.
     */
    public readonly gatewayId!: pulumi.Output<string>;
    /**
     * The VPN gateway IP ID.
     */
    public readonly gatewayIp!: pulumi.Output<string>;
    public readonly haRole!: pulumi.Output<string>;
    public readonly ikepolicy!: pulumi.Output<outputs.VpnConnectionIkepolicy>;
    public readonly ipsecpolicy!: pulumi.Output<outputs.VpnConnectionIpsecpolicy>;
    /**
     * The name of the VPN connection.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The customer subnets.
     */
    public readonly peerSubnets!: pulumi.Output<string[]>;
    /**
     * The policy rules. Only works when vpnType is set to **policy**
     */
    public readonly policyRules!: pulumi.Output<outputs.VpnConnectionPolicyRule[]>;
    /**
     * The pre-shared key.
     */
    public readonly psk!: pulumi.Output<string>;
    public readonly region!: pulumi.Output<string>;
    /**
     * The status of the VPN connection.
     */
    public /*out*/ readonly status!: pulumi.Output<string>;
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * The local tunnel address.
     */
    public readonly tunnelLocalAddress!: pulumi.Output<string>;
    /**
     * The peer tunnel address.
     */
    public readonly tunnelPeerAddress!: pulumi.Output<string>;
    /**
     * The update time.
     */
    public /*out*/ readonly updatedAt!: pulumi.Output<string>;
    /**
     * The connection type. The value can be **policy**, **static** or **bgp**.
     */
    public readonly vpnType!: pulumi.Output<string>;

    /**
     * Create a VpnConnection resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: VpnConnectionArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: VpnConnectionArgs | VpnConnectionState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as VpnConnectionState | undefined;
            resourceInputs["createdAt"] = state ? state.createdAt : undefined;
            resourceInputs["customerGatewayId"] = state ? state.customerGatewayId : undefined;
            resourceInputs["enableNqa"] = state ? state.enableNqa : undefined;
            resourceInputs["enterpriseProjectId"] = state ? state.enterpriseProjectId : undefined;
            resourceInputs["gatewayId"] = state ? state.gatewayId : undefined;
            resourceInputs["gatewayIp"] = state ? state.gatewayIp : undefined;
            resourceInputs["haRole"] = state ? state.haRole : undefined;
            resourceInputs["ikepolicy"] = state ? state.ikepolicy : undefined;
            resourceInputs["ipsecpolicy"] = state ? state.ipsecpolicy : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["peerSubnets"] = state ? state.peerSubnets : undefined;
            resourceInputs["policyRules"] = state ? state.policyRules : undefined;
            resourceInputs["psk"] = state ? state.psk : undefined;
            resourceInputs["region"] = state ? state.region : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
            resourceInputs["tags"] = state ? state.tags : undefined;
            resourceInputs["tunnelLocalAddress"] = state ? state.tunnelLocalAddress : undefined;
            resourceInputs["tunnelPeerAddress"] = state ? state.tunnelPeerAddress : undefined;
            resourceInputs["updatedAt"] = state ? state.updatedAt : undefined;
            resourceInputs["vpnType"] = state ? state.vpnType : undefined;
        } else {
            const args = argsOrState as VpnConnectionArgs | undefined;
            if ((!args || args.customerGatewayId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'customerGatewayId'");
            }
            if ((!args || args.gatewayId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'gatewayId'");
            }
            if ((!args || args.gatewayIp === undefined) && !opts.urn) {
                throw new Error("Missing required property 'gatewayIp'");
            }
            if ((!args || args.psk === undefined) && !opts.urn) {
                throw new Error("Missing required property 'psk'");
            }
            if ((!args || args.vpnType === undefined) && !opts.urn) {
                throw new Error("Missing required property 'vpnType'");
            }
            resourceInputs["customerGatewayId"] = args ? args.customerGatewayId : undefined;
            resourceInputs["enableNqa"] = args ? args.enableNqa : undefined;
            resourceInputs["enterpriseProjectId"] = args ? args.enterpriseProjectId : undefined;
            resourceInputs["gatewayId"] = args ? args.gatewayId : undefined;
            resourceInputs["gatewayIp"] = args ? args.gatewayIp : undefined;
            resourceInputs["haRole"] = args ? args.haRole : undefined;
            resourceInputs["ikepolicy"] = args ? args.ikepolicy : undefined;
            resourceInputs["ipsecpolicy"] = args ? args.ipsecpolicy : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["peerSubnets"] = args ? args.peerSubnets : undefined;
            resourceInputs["policyRules"] = args ? args.policyRules : undefined;
            resourceInputs["psk"] = args ? args.psk : undefined;
            resourceInputs["region"] = args ? args.region : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["tunnelLocalAddress"] = args ? args.tunnelLocalAddress : undefined;
            resourceInputs["tunnelPeerAddress"] = args ? args.tunnelPeerAddress : undefined;
            resourceInputs["vpnType"] = args ? args.vpnType : undefined;
            resourceInputs["createdAt"] = undefined /*out*/;
            resourceInputs["status"] = undefined /*out*/;
            resourceInputs["updatedAt"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(VpnConnection.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering VpnConnection resources.
 */
export interface VpnConnectionState {
    /**
     * The create time.
     */
    createdAt?: pulumi.Input<string>;
    /**
     * The customer gateway ID.
     */
    customerGatewayId?: pulumi.Input<string>;
    /**
     * Whether to enable NQA check.
     */
    enableNqa?: pulumi.Input<boolean>;
    /**
     * The enterprise project ID.
     */
    enterpriseProjectId?: pulumi.Input<string>;
    /**
     * The VPN gateway ID.
     */
    gatewayId?: pulumi.Input<string>;
    /**
     * The VPN gateway IP ID.
     */
    gatewayIp?: pulumi.Input<string>;
    haRole?: pulumi.Input<string>;
    ikepolicy?: pulumi.Input<inputs.VpnConnectionIkepolicy>;
    ipsecpolicy?: pulumi.Input<inputs.VpnConnectionIpsecpolicy>;
    /**
     * The name of the VPN connection.
     */
    name?: pulumi.Input<string>;
    /**
     * The customer subnets.
     */
    peerSubnets?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The policy rules. Only works when vpnType is set to **policy**
     */
    policyRules?: pulumi.Input<pulumi.Input<inputs.VpnConnectionPolicyRule>[]>;
    /**
     * The pre-shared key.
     */
    psk?: pulumi.Input<string>;
    region?: pulumi.Input<string>;
    /**
     * The status of the VPN connection.
     */
    status?: pulumi.Input<string>;
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The local tunnel address.
     */
    tunnelLocalAddress?: pulumi.Input<string>;
    /**
     * The peer tunnel address.
     */
    tunnelPeerAddress?: pulumi.Input<string>;
    /**
     * The update time.
     */
    updatedAt?: pulumi.Input<string>;
    /**
     * The connection type. The value can be **policy**, **static** or **bgp**.
     */
    vpnType?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a VpnConnection resource.
 */
export interface VpnConnectionArgs {
    /**
     * The customer gateway ID.
     */
    customerGatewayId: pulumi.Input<string>;
    /**
     * Whether to enable NQA check.
     */
    enableNqa?: pulumi.Input<boolean>;
    /**
     * The enterprise project ID.
     */
    enterpriseProjectId?: pulumi.Input<string>;
    /**
     * The VPN gateway ID.
     */
    gatewayId: pulumi.Input<string>;
    /**
     * The VPN gateway IP ID.
     */
    gatewayIp: pulumi.Input<string>;
    haRole?: pulumi.Input<string>;
    ikepolicy?: pulumi.Input<inputs.VpnConnectionIkepolicy>;
    ipsecpolicy?: pulumi.Input<inputs.VpnConnectionIpsecpolicy>;
    /**
     * The name of the VPN connection.
     */
    name?: pulumi.Input<string>;
    /**
     * The customer subnets.
     */
    peerSubnets?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The policy rules. Only works when vpnType is set to **policy**
     */
    policyRules?: pulumi.Input<pulumi.Input<inputs.VpnConnectionPolicyRule>[]>;
    /**
     * The pre-shared key.
     */
    psk: pulumi.Input<string>;
    region?: pulumi.Input<string>;
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The local tunnel address.
     */
    tunnelLocalAddress?: pulumi.Input<string>;
    /**
     * The peer tunnel address.
     */
    tunnelPeerAddress?: pulumi.Input<string>;
    /**
     * The connection type. The value can be **policy**, **static** or **bgp**.
     */
    vpnType: pulumi.Input<string>;
}
