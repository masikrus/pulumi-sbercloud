// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

export class DliQueue extends pulumi.CustomResource {
    /**
     * Get an existing DliQueue resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: DliQueueState, opts?: pulumi.CustomResourceOptions): DliQueue {
        return new DliQueue(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'sbercloud:index/dliQueue:DliQueue';

    /**
     * Returns true if the given object is an instance of DliQueue.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is DliQueue {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === DliQueue.__pulumiType;
    }

    public /*out*/ readonly createTime!: pulumi.Output<number>;
    public readonly cuCount!: pulumi.Output<number>;
    public readonly description!: pulumi.Output<string>;
    /**
     * The name of the elastic resource pool to which the queue belongs.
     */
    public readonly elasticResourcePoolName!: pulumi.Output<string>;
    public readonly enterpriseProjectId!: pulumi.Output<string>;
    public readonly feature!: pulumi.Output<string | undefined>;
    /**
     * @deprecated management_subnet_cidr is Deprecated
     */
    public readonly managementSubnetCidr!: pulumi.Output<string | undefined>;
    public readonly name!: pulumi.Output<string>;
    public readonly platform!: pulumi.Output<string | undefined>;
    public readonly queueType!: pulumi.Output<string | undefined>;
    public readonly region!: pulumi.Output<string>;
    /**
     * The queue resource mode.
     */
    public readonly resourceMode!: pulumi.Output<number>;
    public readonly scalingPolicies!: pulumi.Output<outputs.DliQueueScalingPolicy[]>;
    public readonly sparkDriver!: pulumi.Output<outputs.DliQueueSparkDriver | undefined>;
    /**
     * @deprecated subnet_cidr is Deprecated
     */
    public readonly subnetCidr!: pulumi.Output<string | undefined>;
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * The CIDR block of the queue.
     */
    public readonly vpcCidr!: pulumi.Output<string>;

    /**
     * Create a DliQueue resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: DliQueueArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: DliQueueArgs | DliQueueState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as DliQueueState | undefined;
            resourceInputs["createTime"] = state ? state.createTime : undefined;
            resourceInputs["cuCount"] = state ? state.cuCount : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["elasticResourcePoolName"] = state ? state.elasticResourcePoolName : undefined;
            resourceInputs["enterpriseProjectId"] = state ? state.enterpriseProjectId : undefined;
            resourceInputs["feature"] = state ? state.feature : undefined;
            resourceInputs["managementSubnetCidr"] = state ? state.managementSubnetCidr : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["platform"] = state ? state.platform : undefined;
            resourceInputs["queueType"] = state ? state.queueType : undefined;
            resourceInputs["region"] = state ? state.region : undefined;
            resourceInputs["resourceMode"] = state ? state.resourceMode : undefined;
            resourceInputs["scalingPolicies"] = state ? state.scalingPolicies : undefined;
            resourceInputs["sparkDriver"] = state ? state.sparkDriver : undefined;
            resourceInputs["subnetCidr"] = state ? state.subnetCidr : undefined;
            resourceInputs["tags"] = state ? state.tags : undefined;
            resourceInputs["vpcCidr"] = state ? state.vpcCidr : undefined;
        } else {
            const args = argsOrState as DliQueueArgs | undefined;
            if ((!args || args.cuCount === undefined) && !opts.urn) {
                throw new Error("Missing required property 'cuCount'");
            }
            resourceInputs["cuCount"] = args ? args.cuCount : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["elasticResourcePoolName"] = args ? args.elasticResourcePoolName : undefined;
            resourceInputs["enterpriseProjectId"] = args ? args.enterpriseProjectId : undefined;
            resourceInputs["feature"] = args ? args.feature : undefined;
            resourceInputs["managementSubnetCidr"] = args ? args.managementSubnetCidr : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["platform"] = args ? args.platform : undefined;
            resourceInputs["queueType"] = args ? args.queueType : undefined;
            resourceInputs["region"] = args ? args.region : undefined;
            resourceInputs["resourceMode"] = args ? args.resourceMode : undefined;
            resourceInputs["scalingPolicies"] = args ? args.scalingPolicies : undefined;
            resourceInputs["sparkDriver"] = args ? args.sparkDriver : undefined;
            resourceInputs["subnetCidr"] = args ? args.subnetCidr : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["vpcCidr"] = args ? args.vpcCidr : undefined;
            resourceInputs["createTime"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(DliQueue.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering DliQueue resources.
 */
export interface DliQueueState {
    createTime?: pulumi.Input<number>;
    cuCount?: pulumi.Input<number>;
    description?: pulumi.Input<string>;
    /**
     * The name of the elastic resource pool to which the queue belongs.
     */
    elasticResourcePoolName?: pulumi.Input<string>;
    enterpriseProjectId?: pulumi.Input<string>;
    feature?: pulumi.Input<string>;
    /**
     * @deprecated management_subnet_cidr is Deprecated
     */
    managementSubnetCidr?: pulumi.Input<string>;
    name?: pulumi.Input<string>;
    platform?: pulumi.Input<string>;
    queueType?: pulumi.Input<string>;
    region?: pulumi.Input<string>;
    /**
     * The queue resource mode.
     */
    resourceMode?: pulumi.Input<number>;
    scalingPolicies?: pulumi.Input<pulumi.Input<inputs.DliQueueScalingPolicy>[]>;
    sparkDriver?: pulumi.Input<inputs.DliQueueSparkDriver>;
    /**
     * @deprecated subnet_cidr is Deprecated
     */
    subnetCidr?: pulumi.Input<string>;
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The CIDR block of the queue.
     */
    vpcCidr?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a DliQueue resource.
 */
export interface DliQueueArgs {
    cuCount: pulumi.Input<number>;
    description?: pulumi.Input<string>;
    /**
     * The name of the elastic resource pool to which the queue belongs.
     */
    elasticResourcePoolName?: pulumi.Input<string>;
    enterpriseProjectId?: pulumi.Input<string>;
    feature?: pulumi.Input<string>;
    /**
     * @deprecated management_subnet_cidr is Deprecated
     */
    managementSubnetCidr?: pulumi.Input<string>;
    name?: pulumi.Input<string>;
    platform?: pulumi.Input<string>;
    queueType?: pulumi.Input<string>;
    region?: pulumi.Input<string>;
    /**
     * The queue resource mode.
     */
    resourceMode?: pulumi.Input<number>;
    scalingPolicies?: pulumi.Input<pulumi.Input<inputs.DliQueueScalingPolicy>[]>;
    sparkDriver?: pulumi.Input<inputs.DliQueueSparkDriver>;
    /**
     * @deprecated subnet_cidr is Deprecated
     */
    subnetCidr?: pulumi.Input<string>;
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The CIDR block of the queue.
     */
    vpcCidr?: pulumi.Input<string>;
}
