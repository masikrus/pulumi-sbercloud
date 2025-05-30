// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

export class DrsJob extends pulumi.CustomResource {
    /**
     * Get an existing DrsJob resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: DrsJobState, opts?: pulumi.CustomResourceOptions): DrsJob {
        return new DrsJob(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'sbercloud:index/drsJob:DrsJob';

    /**
     * Returns true if the given object is an instance of DrsJob.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is DrsJob {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === DrsJob.__pulumiType;
    }

    public /*out*/ readonly createdAt!: pulumi.Output<string>;
    public readonly description!: pulumi.Output<string | undefined>;
    public readonly destinationDb!: pulumi.Output<outputs.DrsJobDestinationDb>;
    public readonly destinationDbReadnoly!: pulumi.Output<boolean | undefined>;
    public readonly direction!: pulumi.Output<string>;
    public readonly engineType!: pulumi.Output<string>;
    public readonly enterpriseProjectId!: pulumi.Output<string | undefined>;
    public readonly expiredDays!: pulumi.Output<number | undefined>;
    public readonly forceDestroy!: pulumi.Output<boolean | undefined>;
    public readonly limitSpeeds!: pulumi.Output<outputs.DrsJobLimitSpeed[] | undefined>;
    public readonly migrateDefiner!: pulumi.Output<boolean | undefined>;
    public readonly migrationType!: pulumi.Output<string | undefined>;
    public readonly multiWrite!: pulumi.Output<boolean | undefined>;
    public readonly name!: pulumi.Output<string>;
    public readonly netType!: pulumi.Output<string | undefined>;
    public /*out*/ readonly privateIp!: pulumi.Output<string>;
    public /*out*/ readonly publicIp!: pulumi.Output<string>;
    public readonly region!: pulumi.Output<string>;
    public readonly sourceDb!: pulumi.Output<outputs.DrsJobSourceDb>;
    public readonly startTime!: pulumi.Output<string | undefined>;
    public /*out*/ readonly status!: pulumi.Output<string>;
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    public readonly type!: pulumi.Output<string>;

    /**
     * Create a DrsJob resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: DrsJobArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: DrsJobArgs | DrsJobState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as DrsJobState | undefined;
            resourceInputs["createdAt"] = state ? state.createdAt : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["destinationDb"] = state ? state.destinationDb : undefined;
            resourceInputs["destinationDbReadnoly"] = state ? state.destinationDbReadnoly : undefined;
            resourceInputs["direction"] = state ? state.direction : undefined;
            resourceInputs["engineType"] = state ? state.engineType : undefined;
            resourceInputs["enterpriseProjectId"] = state ? state.enterpriseProjectId : undefined;
            resourceInputs["expiredDays"] = state ? state.expiredDays : undefined;
            resourceInputs["forceDestroy"] = state ? state.forceDestroy : undefined;
            resourceInputs["limitSpeeds"] = state ? state.limitSpeeds : undefined;
            resourceInputs["migrateDefiner"] = state ? state.migrateDefiner : undefined;
            resourceInputs["migrationType"] = state ? state.migrationType : undefined;
            resourceInputs["multiWrite"] = state ? state.multiWrite : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["netType"] = state ? state.netType : undefined;
            resourceInputs["privateIp"] = state ? state.privateIp : undefined;
            resourceInputs["publicIp"] = state ? state.publicIp : undefined;
            resourceInputs["region"] = state ? state.region : undefined;
            resourceInputs["sourceDb"] = state ? state.sourceDb : undefined;
            resourceInputs["startTime"] = state ? state.startTime : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
            resourceInputs["tags"] = state ? state.tags : undefined;
            resourceInputs["type"] = state ? state.type : undefined;
        } else {
            const args = argsOrState as DrsJobArgs | undefined;
            if ((!args || args.destinationDb === undefined) && !opts.urn) {
                throw new Error("Missing required property 'destinationDb'");
            }
            if ((!args || args.direction === undefined) && !opts.urn) {
                throw new Error("Missing required property 'direction'");
            }
            if ((!args || args.engineType === undefined) && !opts.urn) {
                throw new Error("Missing required property 'engineType'");
            }
            if ((!args || args.sourceDb === undefined) && !opts.urn) {
                throw new Error("Missing required property 'sourceDb'");
            }
            if ((!args || args.type === undefined) && !opts.urn) {
                throw new Error("Missing required property 'type'");
            }
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["destinationDb"] = args ? args.destinationDb : undefined;
            resourceInputs["destinationDbReadnoly"] = args ? args.destinationDbReadnoly : undefined;
            resourceInputs["direction"] = args ? args.direction : undefined;
            resourceInputs["engineType"] = args ? args.engineType : undefined;
            resourceInputs["enterpriseProjectId"] = args ? args.enterpriseProjectId : undefined;
            resourceInputs["expiredDays"] = args ? args.expiredDays : undefined;
            resourceInputs["forceDestroy"] = args ? args.forceDestroy : undefined;
            resourceInputs["limitSpeeds"] = args ? args.limitSpeeds : undefined;
            resourceInputs["migrateDefiner"] = args ? args.migrateDefiner : undefined;
            resourceInputs["migrationType"] = args ? args.migrationType : undefined;
            resourceInputs["multiWrite"] = args ? args.multiWrite : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["netType"] = args ? args.netType : undefined;
            resourceInputs["region"] = args ? args.region : undefined;
            resourceInputs["sourceDb"] = args ? args.sourceDb : undefined;
            resourceInputs["startTime"] = args ? args.startTime : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["type"] = args ? args.type : undefined;
            resourceInputs["createdAt"] = undefined /*out*/;
            resourceInputs["privateIp"] = undefined /*out*/;
            resourceInputs["publicIp"] = undefined /*out*/;
            resourceInputs["status"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(DrsJob.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering DrsJob resources.
 */
export interface DrsJobState {
    createdAt?: pulumi.Input<string>;
    description?: pulumi.Input<string>;
    destinationDb?: pulumi.Input<inputs.DrsJobDestinationDb>;
    destinationDbReadnoly?: pulumi.Input<boolean>;
    direction?: pulumi.Input<string>;
    engineType?: pulumi.Input<string>;
    enterpriseProjectId?: pulumi.Input<string>;
    expiredDays?: pulumi.Input<number>;
    forceDestroy?: pulumi.Input<boolean>;
    limitSpeeds?: pulumi.Input<pulumi.Input<inputs.DrsJobLimitSpeed>[]>;
    migrateDefiner?: pulumi.Input<boolean>;
    migrationType?: pulumi.Input<string>;
    multiWrite?: pulumi.Input<boolean>;
    name?: pulumi.Input<string>;
    netType?: pulumi.Input<string>;
    privateIp?: pulumi.Input<string>;
    publicIp?: pulumi.Input<string>;
    region?: pulumi.Input<string>;
    sourceDb?: pulumi.Input<inputs.DrsJobSourceDb>;
    startTime?: pulumi.Input<string>;
    status?: pulumi.Input<string>;
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    type?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a DrsJob resource.
 */
export interface DrsJobArgs {
    description?: pulumi.Input<string>;
    destinationDb: pulumi.Input<inputs.DrsJobDestinationDb>;
    destinationDbReadnoly?: pulumi.Input<boolean>;
    direction: pulumi.Input<string>;
    engineType: pulumi.Input<string>;
    enterpriseProjectId?: pulumi.Input<string>;
    expiredDays?: pulumi.Input<number>;
    forceDestroy?: pulumi.Input<boolean>;
    limitSpeeds?: pulumi.Input<pulumi.Input<inputs.DrsJobLimitSpeed>[]>;
    migrateDefiner?: pulumi.Input<boolean>;
    migrationType?: pulumi.Input<string>;
    multiWrite?: pulumi.Input<boolean>;
    name?: pulumi.Input<string>;
    netType?: pulumi.Input<string>;
    region?: pulumi.Input<string>;
    sourceDb: pulumi.Input<inputs.DrsJobSourceDb>;
    startTime?: pulumi.Input<string>;
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    type: pulumi.Input<string>;
}
