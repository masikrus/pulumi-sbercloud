// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

export class ApigGroup extends pulumi.CustomResource {
    /**
     * Get an existing ApigGroup resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ApigGroupState, opts?: pulumi.CustomResourceOptions): ApigGroup {
        return new ApigGroup(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'sbercloud:index/apigGroup:ApigGroup';

    /**
     * Returns true if the given object is an instance of ApigGroup.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ApigGroup {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ApigGroup.__pulumiType;
    }

    /**
     * The creation time of the group, in RFC3339 format.
     */
    public /*out*/ readonly createdAt!: pulumi.Output<string>;
    /**
     * The group description.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * Specifies whether to use the debugging domain name to access the APIs within the group.
     */
    public readonly domainAccessEnabled!: pulumi.Output<boolean>;
    /**
     * The array of one or more environments of the associated group.
     */
    public readonly environments!: pulumi.Output<outputs.ApigGroupEnvironment[]>;
    /**
     * Whether to delete all sub-resources (for API) from this group.
     */
    public readonly forceDestroy!: pulumi.Output<boolean>;
    /**
     * The ID of the dedicated instance to which the group belongs.
     */
    public readonly instanceId!: pulumi.Output<string>;
    /**
     * The group name.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The region where the dedicated instance is located.
     */
    public readonly region!: pulumi.Output<string>;
    /**
     * The registration time.
     */
    public /*out*/ readonly registrationTime!: pulumi.Output<string>;
    /**
     * The latest update time of the group.
     */
    public /*out*/ readonly updateTime!: pulumi.Output<string>;
    /**
     * The latest update time of the group, in RFC3339 format.
     */
    public /*out*/ readonly updatedAt!: pulumi.Output<string>;
    public readonly urlDomains!: pulumi.Output<outputs.ApigGroupUrlDomain[] | undefined>;

    /**
     * Create a ApigGroup resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ApigGroupArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ApigGroupArgs | ApigGroupState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ApigGroupState | undefined;
            resourceInputs["createdAt"] = state ? state.createdAt : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["domainAccessEnabled"] = state ? state.domainAccessEnabled : undefined;
            resourceInputs["environments"] = state ? state.environments : undefined;
            resourceInputs["forceDestroy"] = state ? state.forceDestroy : undefined;
            resourceInputs["instanceId"] = state ? state.instanceId : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["region"] = state ? state.region : undefined;
            resourceInputs["registrationTime"] = state ? state.registrationTime : undefined;
            resourceInputs["updateTime"] = state ? state.updateTime : undefined;
            resourceInputs["updatedAt"] = state ? state.updatedAt : undefined;
            resourceInputs["urlDomains"] = state ? state.urlDomains : undefined;
        } else {
            const args = argsOrState as ApigGroupArgs | undefined;
            if ((!args || args.instanceId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'instanceId'");
            }
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["domainAccessEnabled"] = args ? args.domainAccessEnabled : undefined;
            resourceInputs["environments"] = args ? args.environments : undefined;
            resourceInputs["forceDestroy"] = args ? args.forceDestroy : undefined;
            resourceInputs["instanceId"] = args ? args.instanceId : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["region"] = args ? args.region : undefined;
            resourceInputs["urlDomains"] = args ? args.urlDomains : undefined;
            resourceInputs["createdAt"] = undefined /*out*/;
            resourceInputs["registrationTime"] = undefined /*out*/;
            resourceInputs["updateTime"] = undefined /*out*/;
            resourceInputs["updatedAt"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(ApigGroup.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ApigGroup resources.
 */
export interface ApigGroupState {
    /**
     * The creation time of the group, in RFC3339 format.
     */
    createdAt?: pulumi.Input<string>;
    /**
     * The group description.
     */
    description?: pulumi.Input<string>;
    /**
     * Specifies whether to use the debugging domain name to access the APIs within the group.
     */
    domainAccessEnabled?: pulumi.Input<boolean>;
    /**
     * The array of one or more environments of the associated group.
     */
    environments?: pulumi.Input<pulumi.Input<inputs.ApigGroupEnvironment>[]>;
    /**
     * Whether to delete all sub-resources (for API) from this group.
     */
    forceDestroy?: pulumi.Input<boolean>;
    /**
     * The ID of the dedicated instance to which the group belongs.
     */
    instanceId?: pulumi.Input<string>;
    /**
     * The group name.
     */
    name?: pulumi.Input<string>;
    /**
     * The region where the dedicated instance is located.
     */
    region?: pulumi.Input<string>;
    /**
     * The registration time.
     */
    registrationTime?: pulumi.Input<string>;
    /**
     * The latest update time of the group.
     */
    updateTime?: pulumi.Input<string>;
    /**
     * The latest update time of the group, in RFC3339 format.
     */
    updatedAt?: pulumi.Input<string>;
    urlDomains?: pulumi.Input<pulumi.Input<inputs.ApigGroupUrlDomain>[]>;
}

/**
 * The set of arguments for constructing a ApigGroup resource.
 */
export interface ApigGroupArgs {
    /**
     * The group description.
     */
    description?: pulumi.Input<string>;
    /**
     * Specifies whether to use the debugging domain name to access the APIs within the group.
     */
    domainAccessEnabled?: pulumi.Input<boolean>;
    /**
     * The array of one or more environments of the associated group.
     */
    environments?: pulumi.Input<pulumi.Input<inputs.ApigGroupEnvironment>[]>;
    /**
     * Whether to delete all sub-resources (for API) from this group.
     */
    forceDestroy?: pulumi.Input<boolean>;
    /**
     * The ID of the dedicated instance to which the group belongs.
     */
    instanceId: pulumi.Input<string>;
    /**
     * The group name.
     */
    name?: pulumi.Input<string>;
    /**
     * The region where the dedicated instance is located.
     */
    region?: pulumi.Input<string>;
    urlDomains?: pulumi.Input<pulumi.Input<inputs.ApigGroupUrlDomain>[]>;
}
