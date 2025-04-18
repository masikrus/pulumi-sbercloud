// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

export class ComputeInstance extends pulumi.CustomResource {
    /**
     * Get an existing ComputeInstance resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ComputeInstanceState, opts?: pulumi.CustomResourceOptions): ComputeInstance {
        return new ComputeInstance(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'sbercloud:index/computeInstance:ComputeInstance';

    /**
     * Returns true if the given object is an instance of ComputeInstance.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ComputeInstance {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ComputeInstance.__pulumiType;
    }

    public /*out*/ readonly accessIpV4!: pulumi.Output<string>;
    public /*out*/ readonly accessIpV6!: pulumi.Output<string>;
    public readonly adminPass!: pulumi.Output<string | undefined>;
    public readonly agencyName!: pulumi.Output<string>;
    public readonly agentList!: pulumi.Output<string>;
    /**
     * @deprecated Deprecated
     */
    public readonly autoPay!: pulumi.Output<string | undefined>;
    public readonly autoRenew!: pulumi.Output<string | undefined>;
    public readonly autoTerminateTime!: pulumi.Output<string | undefined>;
    public readonly availabilityZone!: pulumi.Output<string>;
    public readonly bandwidth!: pulumi.Output<outputs.ComputeInstanceBandwidth | undefined>;
    public readonly chargingMode!: pulumi.Output<string>;
    public /*out*/ readonly createdAt!: pulumi.Output<string>;
    public readonly dataDisks!: pulumi.Output<outputs.ComputeInstanceDataDisk[] | undefined>;
    public readonly deleteDisksOnTermination!: pulumi.Output<boolean | undefined>;
    public readonly deleteEipOnTermination!: pulumi.Output<boolean | undefined>;
    public readonly description!: pulumi.Output<string>;
    public readonly eipId!: pulumi.Output<string | undefined>;
    public readonly eipType!: pulumi.Output<string | undefined>;
    public readonly enterpriseProjectId!: pulumi.Output<string>;
    public /*out*/ readonly expiredTime!: pulumi.Output<string>;
    /**
     * schema: Required
     */
    public readonly flavorId!: pulumi.Output<string>;
    /**
     * schema: Computed
     */
    public readonly flavorName!: pulumi.Output<string>;
    public readonly hostname!: pulumi.Output<string>;
    public readonly imageId!: pulumi.Output<string>;
    public readonly imageName!: pulumi.Output<string>;
    public readonly keyPair!: pulumi.Output<string | undefined>;
    public readonly metadata!: pulumi.Output<{[key: string]: string} | undefined>;
    public readonly name!: pulumi.Output<string>;
    public readonly networks!: pulumi.Output<outputs.ComputeInstanceNetwork[]>;
    public readonly period!: pulumi.Output<number | undefined>;
    public readonly periodUnit!: pulumi.Output<string | undefined>;
    public readonly powerAction!: pulumi.Output<string>;
    public readonly privateKey!: pulumi.Output<string | undefined>;
    public /*out*/ readonly publicIp!: pulumi.Output<string>;
    public readonly region!: pulumi.Output<string>;
    public readonly schedulerHints!: pulumi.Output<outputs.ComputeInstanceSchedulerHint[]>;
    public readonly securityGroupIds!: pulumi.Output<string[]>;
    /**
     * schema: Computed
     */
    public readonly securityGroups!: pulumi.Output<string[]>;
    public readonly spotDuration!: pulumi.Output<number | undefined>;
    public readonly spotDurationCount!: pulumi.Output<number>;
    public readonly spotMaximumPrice!: pulumi.Output<string | undefined>;
    public /*out*/ readonly status!: pulumi.Output<string>;
    public readonly stopBeforeDestroy!: pulumi.Output<boolean | undefined>;
    public readonly systemDiskDssPoolId!: pulumi.Output<string | undefined>;
    public /*out*/ readonly systemDiskId!: pulumi.Output<string>;
    public readonly systemDiskIops!: pulumi.Output<number>;
    public readonly systemDiskKmsKeyId!: pulumi.Output<string>;
    public readonly systemDiskSize!: pulumi.Output<number>;
    public readonly systemDiskThroughput!: pulumi.Output<number>;
    public readonly systemDiskType!: pulumi.Output<string>;
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    public /*out*/ readonly updatedAt!: pulumi.Output<string>;
    public readonly userData!: pulumi.Output<string | undefined>;
    public readonly userId!: pulumi.Output<string | undefined>;
    public /*out*/ readonly volumeAttacheds!: pulumi.Output<outputs.ComputeInstanceVolumeAttached[]>;

    /**
     * Create a ComputeInstance resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ComputeInstanceArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ComputeInstanceArgs | ComputeInstanceState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ComputeInstanceState | undefined;
            resourceInputs["accessIpV4"] = state ? state.accessIpV4 : undefined;
            resourceInputs["accessIpV6"] = state ? state.accessIpV6 : undefined;
            resourceInputs["adminPass"] = state ? state.adminPass : undefined;
            resourceInputs["agencyName"] = state ? state.agencyName : undefined;
            resourceInputs["agentList"] = state ? state.agentList : undefined;
            resourceInputs["autoPay"] = state ? state.autoPay : undefined;
            resourceInputs["autoRenew"] = state ? state.autoRenew : undefined;
            resourceInputs["autoTerminateTime"] = state ? state.autoTerminateTime : undefined;
            resourceInputs["availabilityZone"] = state ? state.availabilityZone : undefined;
            resourceInputs["bandwidth"] = state ? state.bandwidth : undefined;
            resourceInputs["chargingMode"] = state ? state.chargingMode : undefined;
            resourceInputs["createdAt"] = state ? state.createdAt : undefined;
            resourceInputs["dataDisks"] = state ? state.dataDisks : undefined;
            resourceInputs["deleteDisksOnTermination"] = state ? state.deleteDisksOnTermination : undefined;
            resourceInputs["deleteEipOnTermination"] = state ? state.deleteEipOnTermination : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["eipId"] = state ? state.eipId : undefined;
            resourceInputs["eipType"] = state ? state.eipType : undefined;
            resourceInputs["enterpriseProjectId"] = state ? state.enterpriseProjectId : undefined;
            resourceInputs["expiredTime"] = state ? state.expiredTime : undefined;
            resourceInputs["flavorId"] = state ? state.flavorId : undefined;
            resourceInputs["flavorName"] = state ? state.flavorName : undefined;
            resourceInputs["hostname"] = state ? state.hostname : undefined;
            resourceInputs["imageId"] = state ? state.imageId : undefined;
            resourceInputs["imageName"] = state ? state.imageName : undefined;
            resourceInputs["keyPair"] = state ? state.keyPair : undefined;
            resourceInputs["metadata"] = state ? state.metadata : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["networks"] = state ? state.networks : undefined;
            resourceInputs["period"] = state ? state.period : undefined;
            resourceInputs["periodUnit"] = state ? state.periodUnit : undefined;
            resourceInputs["powerAction"] = state ? state.powerAction : undefined;
            resourceInputs["privateKey"] = state ? state.privateKey : undefined;
            resourceInputs["publicIp"] = state ? state.publicIp : undefined;
            resourceInputs["region"] = state ? state.region : undefined;
            resourceInputs["schedulerHints"] = state ? state.schedulerHints : undefined;
            resourceInputs["securityGroupIds"] = state ? state.securityGroupIds : undefined;
            resourceInputs["securityGroups"] = state ? state.securityGroups : undefined;
            resourceInputs["spotDuration"] = state ? state.spotDuration : undefined;
            resourceInputs["spotDurationCount"] = state ? state.spotDurationCount : undefined;
            resourceInputs["spotMaximumPrice"] = state ? state.spotMaximumPrice : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
            resourceInputs["stopBeforeDestroy"] = state ? state.stopBeforeDestroy : undefined;
            resourceInputs["systemDiskDssPoolId"] = state ? state.systemDiskDssPoolId : undefined;
            resourceInputs["systemDiskId"] = state ? state.systemDiskId : undefined;
            resourceInputs["systemDiskIops"] = state ? state.systemDiskIops : undefined;
            resourceInputs["systemDiskKmsKeyId"] = state ? state.systemDiskKmsKeyId : undefined;
            resourceInputs["systemDiskSize"] = state ? state.systemDiskSize : undefined;
            resourceInputs["systemDiskThroughput"] = state ? state.systemDiskThroughput : undefined;
            resourceInputs["systemDiskType"] = state ? state.systemDiskType : undefined;
            resourceInputs["tags"] = state ? state.tags : undefined;
            resourceInputs["updatedAt"] = state ? state.updatedAt : undefined;
            resourceInputs["userData"] = state ? state.userData : undefined;
            resourceInputs["userId"] = state ? state.userId : undefined;
            resourceInputs["volumeAttacheds"] = state ? state.volumeAttacheds : undefined;
        } else {
            const args = argsOrState as ComputeInstanceArgs | undefined;
            if ((!args || args.networks === undefined) && !opts.urn) {
                throw new Error("Missing required property 'networks'");
            }
            resourceInputs["adminPass"] = args?.adminPass ? pulumi.secret(args.adminPass) : undefined;
            resourceInputs["agencyName"] = args ? args.agencyName : undefined;
            resourceInputs["agentList"] = args ? args.agentList : undefined;
            resourceInputs["autoPay"] = args ? args.autoPay : undefined;
            resourceInputs["autoRenew"] = args ? args.autoRenew : undefined;
            resourceInputs["autoTerminateTime"] = args ? args.autoTerminateTime : undefined;
            resourceInputs["availabilityZone"] = args ? args.availabilityZone : undefined;
            resourceInputs["bandwidth"] = args ? args.bandwidth : undefined;
            resourceInputs["chargingMode"] = args ? args.chargingMode : undefined;
            resourceInputs["dataDisks"] = args ? args.dataDisks : undefined;
            resourceInputs["deleteDisksOnTermination"] = args ? args.deleteDisksOnTermination : undefined;
            resourceInputs["deleteEipOnTermination"] = args ? args.deleteEipOnTermination : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["eipId"] = args ? args.eipId : undefined;
            resourceInputs["eipType"] = args ? args.eipType : undefined;
            resourceInputs["enterpriseProjectId"] = args ? args.enterpriseProjectId : undefined;
            resourceInputs["flavorId"] = args ? args.flavorId : undefined;
            resourceInputs["flavorName"] = args ? args.flavorName : undefined;
            resourceInputs["hostname"] = args ? args.hostname : undefined;
            resourceInputs["imageId"] = args ? args.imageId : undefined;
            resourceInputs["imageName"] = args ? args.imageName : undefined;
            resourceInputs["keyPair"] = args ? args.keyPair : undefined;
            resourceInputs["metadata"] = args ? args.metadata : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["networks"] = args ? args.networks : undefined;
            resourceInputs["period"] = args ? args.period : undefined;
            resourceInputs["periodUnit"] = args ? args.periodUnit : undefined;
            resourceInputs["powerAction"] = args ? args.powerAction : undefined;
            resourceInputs["privateKey"] = args?.privateKey ? pulumi.secret(args.privateKey) : undefined;
            resourceInputs["region"] = args ? args.region : undefined;
            resourceInputs["schedulerHints"] = args ? args.schedulerHints : undefined;
            resourceInputs["securityGroupIds"] = args ? args.securityGroupIds : undefined;
            resourceInputs["securityGroups"] = args ? args.securityGroups : undefined;
            resourceInputs["spotDuration"] = args ? args.spotDuration : undefined;
            resourceInputs["spotDurationCount"] = args ? args.spotDurationCount : undefined;
            resourceInputs["spotMaximumPrice"] = args ? args.spotMaximumPrice : undefined;
            resourceInputs["stopBeforeDestroy"] = args ? args.stopBeforeDestroy : undefined;
            resourceInputs["systemDiskDssPoolId"] = args ? args.systemDiskDssPoolId : undefined;
            resourceInputs["systemDiskIops"] = args ? args.systemDiskIops : undefined;
            resourceInputs["systemDiskKmsKeyId"] = args ? args.systemDiskKmsKeyId : undefined;
            resourceInputs["systemDiskSize"] = args ? args.systemDiskSize : undefined;
            resourceInputs["systemDiskThroughput"] = args ? args.systemDiskThroughput : undefined;
            resourceInputs["systemDiskType"] = args ? args.systemDiskType : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["userData"] = args ? args.userData : undefined;
            resourceInputs["userId"] = args ? args.userId : undefined;
            resourceInputs["accessIpV4"] = undefined /*out*/;
            resourceInputs["accessIpV6"] = undefined /*out*/;
            resourceInputs["createdAt"] = undefined /*out*/;
            resourceInputs["expiredTime"] = undefined /*out*/;
            resourceInputs["publicIp"] = undefined /*out*/;
            resourceInputs["status"] = undefined /*out*/;
            resourceInputs["systemDiskId"] = undefined /*out*/;
            resourceInputs["updatedAt"] = undefined /*out*/;
            resourceInputs["volumeAttacheds"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const secretOpts = { additionalSecretOutputs: ["adminPass", "privateKey"] };
        opts = pulumi.mergeOptions(opts, secretOpts);
        super(ComputeInstance.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ComputeInstance resources.
 */
export interface ComputeInstanceState {
    accessIpV4?: pulumi.Input<string>;
    accessIpV6?: pulumi.Input<string>;
    adminPass?: pulumi.Input<string>;
    agencyName?: pulumi.Input<string>;
    agentList?: pulumi.Input<string>;
    /**
     * @deprecated Deprecated
     */
    autoPay?: pulumi.Input<string>;
    autoRenew?: pulumi.Input<string>;
    autoTerminateTime?: pulumi.Input<string>;
    availabilityZone?: pulumi.Input<string>;
    bandwidth?: pulumi.Input<inputs.ComputeInstanceBandwidth>;
    chargingMode?: pulumi.Input<string>;
    createdAt?: pulumi.Input<string>;
    dataDisks?: pulumi.Input<pulumi.Input<inputs.ComputeInstanceDataDisk>[]>;
    deleteDisksOnTermination?: pulumi.Input<boolean>;
    deleteEipOnTermination?: pulumi.Input<boolean>;
    description?: pulumi.Input<string>;
    eipId?: pulumi.Input<string>;
    eipType?: pulumi.Input<string>;
    enterpriseProjectId?: pulumi.Input<string>;
    expiredTime?: pulumi.Input<string>;
    /**
     * schema: Required
     */
    flavorId?: pulumi.Input<string>;
    /**
     * schema: Computed
     */
    flavorName?: pulumi.Input<string>;
    hostname?: pulumi.Input<string>;
    imageId?: pulumi.Input<string>;
    imageName?: pulumi.Input<string>;
    keyPair?: pulumi.Input<string>;
    metadata?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    name?: pulumi.Input<string>;
    networks?: pulumi.Input<pulumi.Input<inputs.ComputeInstanceNetwork>[]>;
    period?: pulumi.Input<number>;
    periodUnit?: pulumi.Input<string>;
    powerAction?: pulumi.Input<string>;
    privateKey?: pulumi.Input<string>;
    publicIp?: pulumi.Input<string>;
    region?: pulumi.Input<string>;
    schedulerHints?: pulumi.Input<pulumi.Input<inputs.ComputeInstanceSchedulerHint>[]>;
    securityGroupIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * schema: Computed
     */
    securityGroups?: pulumi.Input<pulumi.Input<string>[]>;
    spotDuration?: pulumi.Input<number>;
    spotDurationCount?: pulumi.Input<number>;
    spotMaximumPrice?: pulumi.Input<string>;
    status?: pulumi.Input<string>;
    stopBeforeDestroy?: pulumi.Input<boolean>;
    systemDiskDssPoolId?: pulumi.Input<string>;
    systemDiskId?: pulumi.Input<string>;
    systemDiskIops?: pulumi.Input<number>;
    systemDiskKmsKeyId?: pulumi.Input<string>;
    systemDiskSize?: pulumi.Input<number>;
    systemDiskThroughput?: pulumi.Input<number>;
    systemDiskType?: pulumi.Input<string>;
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    updatedAt?: pulumi.Input<string>;
    userData?: pulumi.Input<string>;
    userId?: pulumi.Input<string>;
    volumeAttacheds?: pulumi.Input<pulumi.Input<inputs.ComputeInstanceVolumeAttached>[]>;
}

/**
 * The set of arguments for constructing a ComputeInstance resource.
 */
export interface ComputeInstanceArgs {
    adminPass?: pulumi.Input<string>;
    agencyName?: pulumi.Input<string>;
    agentList?: pulumi.Input<string>;
    /**
     * @deprecated Deprecated
     */
    autoPay?: pulumi.Input<string>;
    autoRenew?: pulumi.Input<string>;
    autoTerminateTime?: pulumi.Input<string>;
    availabilityZone?: pulumi.Input<string>;
    bandwidth?: pulumi.Input<inputs.ComputeInstanceBandwidth>;
    chargingMode?: pulumi.Input<string>;
    dataDisks?: pulumi.Input<pulumi.Input<inputs.ComputeInstanceDataDisk>[]>;
    deleteDisksOnTermination?: pulumi.Input<boolean>;
    deleteEipOnTermination?: pulumi.Input<boolean>;
    description?: pulumi.Input<string>;
    eipId?: pulumi.Input<string>;
    eipType?: pulumi.Input<string>;
    enterpriseProjectId?: pulumi.Input<string>;
    /**
     * schema: Required
     */
    flavorId?: pulumi.Input<string>;
    /**
     * schema: Computed
     */
    flavorName?: pulumi.Input<string>;
    hostname?: pulumi.Input<string>;
    imageId?: pulumi.Input<string>;
    imageName?: pulumi.Input<string>;
    keyPair?: pulumi.Input<string>;
    metadata?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    name?: pulumi.Input<string>;
    networks: pulumi.Input<pulumi.Input<inputs.ComputeInstanceNetwork>[]>;
    period?: pulumi.Input<number>;
    periodUnit?: pulumi.Input<string>;
    powerAction?: pulumi.Input<string>;
    privateKey?: pulumi.Input<string>;
    region?: pulumi.Input<string>;
    schedulerHints?: pulumi.Input<pulumi.Input<inputs.ComputeInstanceSchedulerHint>[]>;
    securityGroupIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * schema: Computed
     */
    securityGroups?: pulumi.Input<pulumi.Input<string>[]>;
    spotDuration?: pulumi.Input<number>;
    spotDurationCount?: pulumi.Input<number>;
    spotMaximumPrice?: pulumi.Input<string>;
    stopBeforeDestroy?: pulumi.Input<boolean>;
    systemDiskDssPoolId?: pulumi.Input<string>;
    systemDiskIops?: pulumi.Input<number>;
    systemDiskKmsKeyId?: pulumi.Input<string>;
    systemDiskSize?: pulumi.Input<number>;
    systemDiskThroughput?: pulumi.Input<number>;
    systemDiskType?: pulumi.Input<string>;
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    userData?: pulumi.Input<string>;
    userId?: pulumi.Input<string>;
}
