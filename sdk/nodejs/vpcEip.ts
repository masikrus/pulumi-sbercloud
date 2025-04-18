// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

export class VpcEip extends pulumi.CustomResource {
    /**
     * Get an existing VpcEip resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: VpcEipState, opts?: pulumi.CustomResourceOptions): VpcEip {
        return new VpcEip(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'sbercloud:index/vpcEip:VpcEip';

    /**
     * Returns true if the given object is an instance of VpcEip.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is VpcEip {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === VpcEip.__pulumiType;
    }

    public /*out*/ readonly address!: pulumi.Output<string>;
    public /*out*/ readonly associateId!: pulumi.Output<string>;
    public /*out*/ readonly associateType!: pulumi.Output<string>;
    /**
     * @deprecated Deprecated
     */
    public readonly autoPay!: pulumi.Output<string | undefined>;
    public readonly autoRenew!: pulumi.Output<string | undefined>;
    /**
     * The bandwidth configuration.
     */
    public readonly bandwidth!: pulumi.Output<outputs.VpcEipBandwidth>;
    public readonly chargingMode!: pulumi.Output<string>;
    public /*out*/ readonly createdAt!: pulumi.Output<string>;
    /**
     * The enterprise project ID to which the EIP belongs.
     */
    public readonly enterpriseProjectId!: pulumi.Output<string>;
    public /*out*/ readonly instanceId!: pulumi.Output<string>;
    public /*out*/ readonly instanceType!: pulumi.Output<string>;
    public /*out*/ readonly ipv6Address!: pulumi.Output<string>;
    /**
     * The name of the EIP.
     */
    public readonly name!: pulumi.Output<string>;
    public readonly period!: pulumi.Output<number | undefined>;
    public readonly periodUnit!: pulumi.Output<string | undefined>;
    public /*out*/ readonly portId!: pulumi.Output<string>;
    public /*out*/ readonly privateIp!: pulumi.Output<string>;
    /**
     * The EIP configuration.
     */
    public readonly publicip!: pulumi.Output<outputs.VpcEipPublicip>;
    /**
     * The region in which to create the EIP resource.
     */
    public readonly region!: pulumi.Output<string>;
    public /*out*/ readonly status!: pulumi.Output<string>;
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    public /*out*/ readonly updatedAt!: pulumi.Output<string>;

    /**
     * Create a VpcEip resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: VpcEipArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: VpcEipArgs | VpcEipState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as VpcEipState | undefined;
            resourceInputs["address"] = state ? state.address : undefined;
            resourceInputs["associateId"] = state ? state.associateId : undefined;
            resourceInputs["associateType"] = state ? state.associateType : undefined;
            resourceInputs["autoPay"] = state ? state.autoPay : undefined;
            resourceInputs["autoRenew"] = state ? state.autoRenew : undefined;
            resourceInputs["bandwidth"] = state ? state.bandwidth : undefined;
            resourceInputs["chargingMode"] = state ? state.chargingMode : undefined;
            resourceInputs["createdAt"] = state ? state.createdAt : undefined;
            resourceInputs["enterpriseProjectId"] = state ? state.enterpriseProjectId : undefined;
            resourceInputs["instanceId"] = state ? state.instanceId : undefined;
            resourceInputs["instanceType"] = state ? state.instanceType : undefined;
            resourceInputs["ipv6Address"] = state ? state.ipv6Address : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["period"] = state ? state.period : undefined;
            resourceInputs["periodUnit"] = state ? state.periodUnit : undefined;
            resourceInputs["portId"] = state ? state.portId : undefined;
            resourceInputs["privateIp"] = state ? state.privateIp : undefined;
            resourceInputs["publicip"] = state ? state.publicip : undefined;
            resourceInputs["region"] = state ? state.region : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
            resourceInputs["tags"] = state ? state.tags : undefined;
            resourceInputs["updatedAt"] = state ? state.updatedAt : undefined;
        } else {
            const args = argsOrState as VpcEipArgs | undefined;
            if ((!args || args.bandwidth === undefined) && !opts.urn) {
                throw new Error("Missing required property 'bandwidth'");
            }
            if ((!args || args.publicip === undefined) && !opts.urn) {
                throw new Error("Missing required property 'publicip'");
            }
            resourceInputs["autoPay"] = args ? args.autoPay : undefined;
            resourceInputs["autoRenew"] = args ? args.autoRenew : undefined;
            resourceInputs["bandwidth"] = args ? args.bandwidth : undefined;
            resourceInputs["chargingMode"] = args ? args.chargingMode : undefined;
            resourceInputs["enterpriseProjectId"] = args ? args.enterpriseProjectId : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["period"] = args ? args.period : undefined;
            resourceInputs["periodUnit"] = args ? args.periodUnit : undefined;
            resourceInputs["publicip"] = args ? args.publicip : undefined;
            resourceInputs["region"] = args ? args.region : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["address"] = undefined /*out*/;
            resourceInputs["associateId"] = undefined /*out*/;
            resourceInputs["associateType"] = undefined /*out*/;
            resourceInputs["createdAt"] = undefined /*out*/;
            resourceInputs["instanceId"] = undefined /*out*/;
            resourceInputs["instanceType"] = undefined /*out*/;
            resourceInputs["ipv6Address"] = undefined /*out*/;
            resourceInputs["portId"] = undefined /*out*/;
            resourceInputs["privateIp"] = undefined /*out*/;
            resourceInputs["status"] = undefined /*out*/;
            resourceInputs["updatedAt"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(VpcEip.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering VpcEip resources.
 */
export interface VpcEipState {
    address?: pulumi.Input<string>;
    associateId?: pulumi.Input<string>;
    associateType?: pulumi.Input<string>;
    /**
     * @deprecated Deprecated
     */
    autoPay?: pulumi.Input<string>;
    autoRenew?: pulumi.Input<string>;
    /**
     * The bandwidth configuration.
     */
    bandwidth?: pulumi.Input<inputs.VpcEipBandwidth>;
    chargingMode?: pulumi.Input<string>;
    createdAt?: pulumi.Input<string>;
    /**
     * The enterprise project ID to which the EIP belongs.
     */
    enterpriseProjectId?: pulumi.Input<string>;
    instanceId?: pulumi.Input<string>;
    instanceType?: pulumi.Input<string>;
    ipv6Address?: pulumi.Input<string>;
    /**
     * The name of the EIP.
     */
    name?: pulumi.Input<string>;
    period?: pulumi.Input<number>;
    periodUnit?: pulumi.Input<string>;
    portId?: pulumi.Input<string>;
    privateIp?: pulumi.Input<string>;
    /**
     * The EIP configuration.
     */
    publicip?: pulumi.Input<inputs.VpcEipPublicip>;
    /**
     * The region in which to create the EIP resource.
     */
    region?: pulumi.Input<string>;
    status?: pulumi.Input<string>;
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    updatedAt?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a VpcEip resource.
 */
export interface VpcEipArgs {
    /**
     * @deprecated Deprecated
     */
    autoPay?: pulumi.Input<string>;
    autoRenew?: pulumi.Input<string>;
    /**
     * The bandwidth configuration.
     */
    bandwidth: pulumi.Input<inputs.VpcEipBandwidth>;
    chargingMode?: pulumi.Input<string>;
    /**
     * The enterprise project ID to which the EIP belongs.
     */
    enterpriseProjectId?: pulumi.Input<string>;
    /**
     * The name of the EIP.
     */
    name?: pulumi.Input<string>;
    period?: pulumi.Input<number>;
    periodUnit?: pulumi.Input<string>;
    /**
     * The EIP configuration.
     */
    publicip: pulumi.Input<inputs.VpcEipPublicip>;
    /**
     * The region in which to create the EIP resource.
     */
    region?: pulumi.Input<string>;
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}
