// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

export class ApigAclPolicy extends pulumi.CustomResource {
    /**
     * Get an existing ApigAclPolicy resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ApigAclPolicyState, opts?: pulumi.CustomResourceOptions): ApigAclPolicy {
        return new ApigAclPolicy(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'sbercloud:index/apigAclPolicy:ApigAclPolicy';

    /**
     * Returns true if the given object is an instance of ApigAclPolicy.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ApigAclPolicy {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ApigAclPolicy.__pulumiType;
    }

    /**
     * The entity type of the ACL policy.
     */
    public readonly entityType!: pulumi.Output<string>;
    /**
     * The ID of the dedicated instance to which the ACL policy belongs.
     */
    public readonly instanceId!: pulumi.Output<string>;
    /**
     * The name of the ACL policy.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The region where the ACL policy is located.
     */
    public readonly region!: pulumi.Output<string>;
    /**
     * The type of the ACL policy.
     */
    public readonly type!: pulumi.Output<string>;
    /**
     * The latest update time of the ACL policy.
     */
    public /*out*/ readonly updatedAt!: pulumi.Output<string>;
    /**
     * One or more objects from which the access will be controlled.
     */
    public readonly value!: pulumi.Output<string>;

    /**
     * Create a ApigAclPolicy resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ApigAclPolicyArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ApigAclPolicyArgs | ApigAclPolicyState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ApigAclPolicyState | undefined;
            resourceInputs["entityType"] = state ? state.entityType : undefined;
            resourceInputs["instanceId"] = state ? state.instanceId : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["region"] = state ? state.region : undefined;
            resourceInputs["type"] = state ? state.type : undefined;
            resourceInputs["updatedAt"] = state ? state.updatedAt : undefined;
            resourceInputs["value"] = state ? state.value : undefined;
        } else {
            const args = argsOrState as ApigAclPolicyArgs | undefined;
            if ((!args || args.entityType === undefined) && !opts.urn) {
                throw new Error("Missing required property 'entityType'");
            }
            if ((!args || args.instanceId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'instanceId'");
            }
            if ((!args || args.type === undefined) && !opts.urn) {
                throw new Error("Missing required property 'type'");
            }
            if ((!args || args.value === undefined) && !opts.urn) {
                throw new Error("Missing required property 'value'");
            }
            resourceInputs["entityType"] = args ? args.entityType : undefined;
            resourceInputs["instanceId"] = args ? args.instanceId : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["region"] = args ? args.region : undefined;
            resourceInputs["type"] = args ? args.type : undefined;
            resourceInputs["value"] = args ? args.value : undefined;
            resourceInputs["updatedAt"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(ApigAclPolicy.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ApigAclPolicy resources.
 */
export interface ApigAclPolicyState {
    /**
     * The entity type of the ACL policy.
     */
    entityType?: pulumi.Input<string>;
    /**
     * The ID of the dedicated instance to which the ACL policy belongs.
     */
    instanceId?: pulumi.Input<string>;
    /**
     * The name of the ACL policy.
     */
    name?: pulumi.Input<string>;
    /**
     * The region where the ACL policy is located.
     */
    region?: pulumi.Input<string>;
    /**
     * The type of the ACL policy.
     */
    type?: pulumi.Input<string>;
    /**
     * The latest update time of the ACL policy.
     */
    updatedAt?: pulumi.Input<string>;
    /**
     * One or more objects from which the access will be controlled.
     */
    value?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a ApigAclPolicy resource.
 */
export interface ApigAclPolicyArgs {
    /**
     * The entity type of the ACL policy.
     */
    entityType: pulumi.Input<string>;
    /**
     * The ID of the dedicated instance to which the ACL policy belongs.
     */
    instanceId: pulumi.Input<string>;
    /**
     * The name of the ACL policy.
     */
    name?: pulumi.Input<string>;
    /**
     * The region where the ACL policy is located.
     */
    region?: pulumi.Input<string>;
    /**
     * The type of the ACL policy.
     */
    type: pulumi.Input<string>;
    /**
     * One or more objects from which the access will be controlled.
     */
    value: pulumi.Input<string>;
}
