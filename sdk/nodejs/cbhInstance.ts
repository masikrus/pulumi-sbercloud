// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

export class CbhInstance extends pulumi.CustomResource {
    /**
     * Get an existing CbhInstance resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: CbhInstanceState, opts?: pulumi.CustomResourceOptions): CbhInstance {
        return new CbhInstance(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'sbercloud:index/cbhInstance:CbhInstance';

    /**
     * Returns true if the given object is an instance of CbhInstance.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is CbhInstance {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === CbhInstance.__pulumiType;
    }

    /**
     * Specifies the size of the additional data disk for the CBH instance.
     */
    public readonly attachDiskSize!: pulumi.Output<number | undefined>;
    /**
     * Specifies whether auto renew is enabled.
     */
    public readonly autoRenew!: pulumi.Output<string | undefined>;
    /**
     * Specifies the availability zone name.
     */
    public readonly availabilityZone!: pulumi.Output<string>;
    /**
     * Specifies the charging mode of the CBH instance.
     */
    public readonly chargingMode!: pulumi.Output<string>;
    /**
     * Indicates the data disk size of the instance.
     */
    public /*out*/ readonly dataDiskSize!: pulumi.Output<number>;
    /**
     * Specifies the enterprise project ID to which the CBH instance belongs.
     */
    public readonly enterpriseProjectId!: pulumi.Output<string>;
    /**
     * Specifies the product ID of the CBH server.
     */
    public readonly flavorId!: pulumi.Output<string>;
    /**
     * Specifies whether the IPv6 network is enabled.
     */
    public readonly ipv6Enable!: pulumi.Output<boolean | undefined>;
    /**
     * Specifies the name of the CBH instance.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Specifies the password for logging in to the management console.
     */
    public readonly password!: pulumi.Output<string>;
    /**
     * Specifies the charging period of the CBH instance.
     */
    public readonly period!: pulumi.Output<number>;
    /**
     * Specifies the charging period unit of the instance.
     */
    public readonly periodUnit!: pulumi.Output<string>;
    /**
     * Specifies the power action after the CBH instance is created.
     */
    public readonly powerAction!: pulumi.Output<string | undefined>;
    /**
     * Indicates the private IP of the instance.
     */
    public /*out*/ readonly privateIp!: pulumi.Output<string>;
    /**
     * schema: Computed; The elastic IP address.
     */
    public readonly publicIp!: pulumi.Output<string>;
    /**
     * Specifies the ID of the elastic IP.
     */
    public readonly publicIpId!: pulumi.Output<string | undefined>;
    public readonly region!: pulumi.Output<string>;
    /**
     * Specifies the IDs of the security group.
     */
    public readonly securityGroupId!: pulumi.Output<string>;
    /**
     * Indicates the status of the instance.
     */
    public /*out*/ readonly status!: pulumi.Output<string>;
    /**
     * Specifies the IP address of the subnet.
     */
    public readonly subnetAddress!: pulumi.Output<string>;
    /**
     * Specifies the ID of a subnet.
     */
    public readonly subnetId!: pulumi.Output<string>;
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * Indicates the current version of the instance image.
     */
    public /*out*/ readonly version!: pulumi.Output<string>;
    /**
     * Specifies the ID of a VPC.
     */
    public readonly vpcId!: pulumi.Output<string>;

    /**
     * Create a CbhInstance resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: CbhInstanceArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: CbhInstanceArgs | CbhInstanceState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as CbhInstanceState | undefined;
            resourceInputs["attachDiskSize"] = state ? state.attachDiskSize : undefined;
            resourceInputs["autoRenew"] = state ? state.autoRenew : undefined;
            resourceInputs["availabilityZone"] = state ? state.availabilityZone : undefined;
            resourceInputs["chargingMode"] = state ? state.chargingMode : undefined;
            resourceInputs["dataDiskSize"] = state ? state.dataDiskSize : undefined;
            resourceInputs["enterpriseProjectId"] = state ? state.enterpriseProjectId : undefined;
            resourceInputs["flavorId"] = state ? state.flavorId : undefined;
            resourceInputs["ipv6Enable"] = state ? state.ipv6Enable : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["password"] = state ? state.password : undefined;
            resourceInputs["period"] = state ? state.period : undefined;
            resourceInputs["periodUnit"] = state ? state.periodUnit : undefined;
            resourceInputs["powerAction"] = state ? state.powerAction : undefined;
            resourceInputs["privateIp"] = state ? state.privateIp : undefined;
            resourceInputs["publicIp"] = state ? state.publicIp : undefined;
            resourceInputs["publicIpId"] = state ? state.publicIpId : undefined;
            resourceInputs["region"] = state ? state.region : undefined;
            resourceInputs["securityGroupId"] = state ? state.securityGroupId : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
            resourceInputs["subnetAddress"] = state ? state.subnetAddress : undefined;
            resourceInputs["subnetId"] = state ? state.subnetId : undefined;
            resourceInputs["tags"] = state ? state.tags : undefined;
            resourceInputs["version"] = state ? state.version : undefined;
            resourceInputs["vpcId"] = state ? state.vpcId : undefined;
        } else {
            const args = argsOrState as CbhInstanceArgs | undefined;
            if ((!args || args.availabilityZone === undefined) && !opts.urn) {
                throw new Error("Missing required property 'availabilityZone'");
            }
            if ((!args || args.chargingMode === undefined) && !opts.urn) {
                throw new Error("Missing required property 'chargingMode'");
            }
            if ((!args || args.flavorId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'flavorId'");
            }
            if ((!args || args.password === undefined) && !opts.urn) {
                throw new Error("Missing required property 'password'");
            }
            if ((!args || args.period === undefined) && !opts.urn) {
                throw new Error("Missing required property 'period'");
            }
            if ((!args || args.periodUnit === undefined) && !opts.urn) {
                throw new Error("Missing required property 'periodUnit'");
            }
            if ((!args || args.securityGroupId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'securityGroupId'");
            }
            if ((!args || args.subnetId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'subnetId'");
            }
            if ((!args || args.vpcId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'vpcId'");
            }
            resourceInputs["attachDiskSize"] = args ? args.attachDiskSize : undefined;
            resourceInputs["autoRenew"] = args ? args.autoRenew : undefined;
            resourceInputs["availabilityZone"] = args ? args.availabilityZone : undefined;
            resourceInputs["chargingMode"] = args ? args.chargingMode : undefined;
            resourceInputs["enterpriseProjectId"] = args ? args.enterpriseProjectId : undefined;
            resourceInputs["flavorId"] = args ? args.flavorId : undefined;
            resourceInputs["ipv6Enable"] = args ? args.ipv6Enable : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["password"] = args?.password ? pulumi.secret(args.password) : undefined;
            resourceInputs["period"] = args ? args.period : undefined;
            resourceInputs["periodUnit"] = args ? args.periodUnit : undefined;
            resourceInputs["powerAction"] = args ? args.powerAction : undefined;
            resourceInputs["publicIp"] = args ? args.publicIp : undefined;
            resourceInputs["publicIpId"] = args ? args.publicIpId : undefined;
            resourceInputs["region"] = args ? args.region : undefined;
            resourceInputs["securityGroupId"] = args ? args.securityGroupId : undefined;
            resourceInputs["subnetAddress"] = args ? args.subnetAddress : undefined;
            resourceInputs["subnetId"] = args ? args.subnetId : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["vpcId"] = args ? args.vpcId : undefined;
            resourceInputs["dataDiskSize"] = undefined /*out*/;
            resourceInputs["privateIp"] = undefined /*out*/;
            resourceInputs["status"] = undefined /*out*/;
            resourceInputs["version"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const secretOpts = { additionalSecretOutputs: ["password"] };
        opts = pulumi.mergeOptions(opts, secretOpts);
        super(CbhInstance.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering CbhInstance resources.
 */
export interface CbhInstanceState {
    /**
     * Specifies the size of the additional data disk for the CBH instance.
     */
    attachDiskSize?: pulumi.Input<number>;
    /**
     * Specifies whether auto renew is enabled.
     */
    autoRenew?: pulumi.Input<string>;
    /**
     * Specifies the availability zone name.
     */
    availabilityZone?: pulumi.Input<string>;
    /**
     * Specifies the charging mode of the CBH instance.
     */
    chargingMode?: pulumi.Input<string>;
    /**
     * Indicates the data disk size of the instance.
     */
    dataDiskSize?: pulumi.Input<number>;
    /**
     * Specifies the enterprise project ID to which the CBH instance belongs.
     */
    enterpriseProjectId?: pulumi.Input<string>;
    /**
     * Specifies the product ID of the CBH server.
     */
    flavorId?: pulumi.Input<string>;
    /**
     * Specifies whether the IPv6 network is enabled.
     */
    ipv6Enable?: pulumi.Input<boolean>;
    /**
     * Specifies the name of the CBH instance.
     */
    name?: pulumi.Input<string>;
    /**
     * Specifies the password for logging in to the management console.
     */
    password?: pulumi.Input<string>;
    /**
     * Specifies the charging period of the CBH instance.
     */
    period?: pulumi.Input<number>;
    /**
     * Specifies the charging period unit of the instance.
     */
    periodUnit?: pulumi.Input<string>;
    /**
     * Specifies the power action after the CBH instance is created.
     */
    powerAction?: pulumi.Input<string>;
    /**
     * Indicates the private IP of the instance.
     */
    privateIp?: pulumi.Input<string>;
    /**
     * schema: Computed; The elastic IP address.
     */
    publicIp?: pulumi.Input<string>;
    /**
     * Specifies the ID of the elastic IP.
     */
    publicIpId?: pulumi.Input<string>;
    region?: pulumi.Input<string>;
    /**
     * Specifies the IDs of the security group.
     */
    securityGroupId?: pulumi.Input<string>;
    /**
     * Indicates the status of the instance.
     */
    status?: pulumi.Input<string>;
    /**
     * Specifies the IP address of the subnet.
     */
    subnetAddress?: pulumi.Input<string>;
    /**
     * Specifies the ID of a subnet.
     */
    subnetId?: pulumi.Input<string>;
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Indicates the current version of the instance image.
     */
    version?: pulumi.Input<string>;
    /**
     * Specifies the ID of a VPC.
     */
    vpcId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a CbhInstance resource.
 */
export interface CbhInstanceArgs {
    /**
     * Specifies the size of the additional data disk for the CBH instance.
     */
    attachDiskSize?: pulumi.Input<number>;
    /**
     * Specifies whether auto renew is enabled.
     */
    autoRenew?: pulumi.Input<string>;
    /**
     * Specifies the availability zone name.
     */
    availabilityZone: pulumi.Input<string>;
    /**
     * Specifies the charging mode of the CBH instance.
     */
    chargingMode: pulumi.Input<string>;
    /**
     * Specifies the enterprise project ID to which the CBH instance belongs.
     */
    enterpriseProjectId?: pulumi.Input<string>;
    /**
     * Specifies the product ID of the CBH server.
     */
    flavorId: pulumi.Input<string>;
    /**
     * Specifies whether the IPv6 network is enabled.
     */
    ipv6Enable?: pulumi.Input<boolean>;
    /**
     * Specifies the name of the CBH instance.
     */
    name?: pulumi.Input<string>;
    /**
     * Specifies the password for logging in to the management console.
     */
    password: pulumi.Input<string>;
    /**
     * Specifies the charging period of the CBH instance.
     */
    period: pulumi.Input<number>;
    /**
     * Specifies the charging period unit of the instance.
     */
    periodUnit: pulumi.Input<string>;
    /**
     * Specifies the power action after the CBH instance is created.
     */
    powerAction?: pulumi.Input<string>;
    /**
     * schema: Computed; The elastic IP address.
     */
    publicIp?: pulumi.Input<string>;
    /**
     * Specifies the ID of the elastic IP.
     */
    publicIpId?: pulumi.Input<string>;
    region?: pulumi.Input<string>;
    /**
     * Specifies the IDs of the security group.
     */
    securityGroupId: pulumi.Input<string>;
    /**
     * Specifies the IP address of the subnet.
     */
    subnetAddress?: pulumi.Input<string>;
    /**
     * Specifies the ID of a subnet.
     */
    subnetId: pulumi.Input<string>;
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Specifies the ID of a VPC.
     */
    vpcId: pulumi.Input<string>;
}
