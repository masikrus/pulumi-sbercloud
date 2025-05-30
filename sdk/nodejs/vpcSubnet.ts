// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

export class VpcSubnet extends pulumi.CustomResource {
    /**
     * Get an existing VpcSubnet resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: VpcSubnetState, opts?: pulumi.CustomResourceOptions): VpcSubnet {
        return new VpcSubnet(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'sbercloud:index/vpcSubnet:VpcSubnet';

    /**
     * Returns true if the given object is an instance of VpcSubnet.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is VpcSubnet {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === VpcSubnet.__pulumiType;
    }

    public readonly availabilityZone!: pulumi.Output<string>;
    public readonly cidr!: pulumi.Output<string>;
    public readonly description!: pulumi.Output<string | undefined>;
    public readonly dhcpEnable!: pulumi.Output<boolean | undefined>;
    public readonly dhcpLeaseTime!: pulumi.Output<string>;
    public readonly dnsLists!: pulumi.Output<string[]>;
    public readonly gatewayIp!: pulumi.Output<string>;
    public /*out*/ readonly ipv4SubnetId!: pulumi.Output<string>;
    public /*out*/ readonly ipv6Cidr!: pulumi.Output<string>;
    public readonly ipv6Enable!: pulumi.Output<boolean | undefined>;
    public /*out*/ readonly ipv6Gateway!: pulumi.Output<string>;
    public /*out*/ readonly ipv6SubnetId!: pulumi.Output<string>;
    public readonly name!: pulumi.Output<string>;
    public readonly ntpServerAddress!: pulumi.Output<string | undefined>;
    public readonly primaryDns!: pulumi.Output<string>;
    public readonly region!: pulumi.Output<string>;
    public readonly secondaryDns!: pulumi.Output<string>;
    /**
     * schema: Deprecated
     */
    public /*out*/ readonly subnetId!: pulumi.Output<string>;
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    public readonly vpcId!: pulumi.Output<string>;

    /**
     * Create a VpcSubnet resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: VpcSubnetArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: VpcSubnetArgs | VpcSubnetState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as VpcSubnetState | undefined;
            resourceInputs["availabilityZone"] = state ? state.availabilityZone : undefined;
            resourceInputs["cidr"] = state ? state.cidr : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["dhcpEnable"] = state ? state.dhcpEnable : undefined;
            resourceInputs["dhcpLeaseTime"] = state ? state.dhcpLeaseTime : undefined;
            resourceInputs["dnsLists"] = state ? state.dnsLists : undefined;
            resourceInputs["gatewayIp"] = state ? state.gatewayIp : undefined;
            resourceInputs["ipv4SubnetId"] = state ? state.ipv4SubnetId : undefined;
            resourceInputs["ipv6Cidr"] = state ? state.ipv6Cidr : undefined;
            resourceInputs["ipv6Enable"] = state ? state.ipv6Enable : undefined;
            resourceInputs["ipv6Gateway"] = state ? state.ipv6Gateway : undefined;
            resourceInputs["ipv6SubnetId"] = state ? state.ipv6SubnetId : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["ntpServerAddress"] = state ? state.ntpServerAddress : undefined;
            resourceInputs["primaryDns"] = state ? state.primaryDns : undefined;
            resourceInputs["region"] = state ? state.region : undefined;
            resourceInputs["secondaryDns"] = state ? state.secondaryDns : undefined;
            resourceInputs["subnetId"] = state ? state.subnetId : undefined;
            resourceInputs["tags"] = state ? state.tags : undefined;
            resourceInputs["vpcId"] = state ? state.vpcId : undefined;
        } else {
            const args = argsOrState as VpcSubnetArgs | undefined;
            if ((!args || args.cidr === undefined) && !opts.urn) {
                throw new Error("Missing required property 'cidr'");
            }
            if ((!args || args.gatewayIp === undefined) && !opts.urn) {
                throw new Error("Missing required property 'gatewayIp'");
            }
            if ((!args || args.vpcId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'vpcId'");
            }
            resourceInputs["availabilityZone"] = args ? args.availabilityZone : undefined;
            resourceInputs["cidr"] = args ? args.cidr : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["dhcpEnable"] = args ? args.dhcpEnable : undefined;
            resourceInputs["dhcpLeaseTime"] = args ? args.dhcpLeaseTime : undefined;
            resourceInputs["dnsLists"] = args ? args.dnsLists : undefined;
            resourceInputs["gatewayIp"] = args ? args.gatewayIp : undefined;
            resourceInputs["ipv6Enable"] = args ? args.ipv6Enable : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["ntpServerAddress"] = args ? args.ntpServerAddress : undefined;
            resourceInputs["primaryDns"] = args ? args.primaryDns : undefined;
            resourceInputs["region"] = args ? args.region : undefined;
            resourceInputs["secondaryDns"] = args ? args.secondaryDns : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["vpcId"] = args ? args.vpcId : undefined;
            resourceInputs["ipv4SubnetId"] = undefined /*out*/;
            resourceInputs["ipv6Cidr"] = undefined /*out*/;
            resourceInputs["ipv6Gateway"] = undefined /*out*/;
            resourceInputs["ipv6SubnetId"] = undefined /*out*/;
            resourceInputs["subnetId"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(VpcSubnet.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering VpcSubnet resources.
 */
export interface VpcSubnetState {
    availabilityZone?: pulumi.Input<string>;
    cidr?: pulumi.Input<string>;
    description?: pulumi.Input<string>;
    dhcpEnable?: pulumi.Input<boolean>;
    dhcpLeaseTime?: pulumi.Input<string>;
    dnsLists?: pulumi.Input<pulumi.Input<string>[]>;
    gatewayIp?: pulumi.Input<string>;
    ipv4SubnetId?: pulumi.Input<string>;
    ipv6Cidr?: pulumi.Input<string>;
    ipv6Enable?: pulumi.Input<boolean>;
    ipv6Gateway?: pulumi.Input<string>;
    ipv6SubnetId?: pulumi.Input<string>;
    name?: pulumi.Input<string>;
    ntpServerAddress?: pulumi.Input<string>;
    primaryDns?: pulumi.Input<string>;
    region?: pulumi.Input<string>;
    secondaryDns?: pulumi.Input<string>;
    /**
     * schema: Deprecated
     */
    subnetId?: pulumi.Input<string>;
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    vpcId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a VpcSubnet resource.
 */
export interface VpcSubnetArgs {
    availabilityZone?: pulumi.Input<string>;
    cidr: pulumi.Input<string>;
    description?: pulumi.Input<string>;
    dhcpEnable?: pulumi.Input<boolean>;
    dhcpLeaseTime?: pulumi.Input<string>;
    dnsLists?: pulumi.Input<pulumi.Input<string>[]>;
    gatewayIp: pulumi.Input<string>;
    ipv6Enable?: pulumi.Input<boolean>;
    name?: pulumi.Input<string>;
    ntpServerAddress?: pulumi.Input<string>;
    primaryDns?: pulumi.Input<string>;
    region?: pulumi.Input<string>;
    secondaryDns?: pulumi.Input<string>;
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    vpcId: pulumi.Input<string>;
}
