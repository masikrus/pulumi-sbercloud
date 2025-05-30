// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

export class ApigApiPublishment extends pulumi.CustomResource {
    /**
     * Get an existing ApigApiPublishment resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ApigApiPublishmentState, opts?: pulumi.CustomResourceOptions): ApigApiPublishment {
        return new ApigApiPublishment(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'sbercloud:index/apigApiPublishment:ApigApiPublishment';

    /**
     * Returns true if the given object is an instance of ApigApiPublishment.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ApigApiPublishment {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ApigApiPublishment.__pulumiType;
    }

    /**
     * The ID of the API to be published or already published.
     */
    public readonly apiId!: pulumi.Output<string>;
    /**
     * The description of the current publishment.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * The ID of the environment to which the current version of the API will be published or has been published.
     */
    public readonly envId!: pulumi.Output<string>;
    /**
     * The name of the environment to which the current version of the API is published.
     */
    public /*out*/ readonly envName!: pulumi.Output<string>;
    /**
     * All publish informations of the API.
     */
    public /*out*/ readonly histories!: pulumi.Output<outputs.ApigApiPublishmentHistory[]>;
    /**
     * The ID of the dedicated instance to which the API and the environment belongs.
     */
    public readonly instanceId!: pulumi.Output<string>;
    /**
     * The publish ID of the API in current environment.
     */
    public /*out*/ readonly publishId!: pulumi.Output<string>;
    /**
     * Time when the current version was published.
     */
    public /*out*/ readonly publishedAt!: pulumi.Output<string>;
    /**
     * The region in which to publish API.
     */
    public readonly region!: pulumi.Output<string>;
    /**
     * The version ID of the current publishment.
     */
    public readonly versionId!: pulumi.Output<string | undefined>;

    /**
     * Create a ApigApiPublishment resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ApigApiPublishmentArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ApigApiPublishmentArgs | ApigApiPublishmentState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ApigApiPublishmentState | undefined;
            resourceInputs["apiId"] = state ? state.apiId : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["envId"] = state ? state.envId : undefined;
            resourceInputs["envName"] = state ? state.envName : undefined;
            resourceInputs["histories"] = state ? state.histories : undefined;
            resourceInputs["instanceId"] = state ? state.instanceId : undefined;
            resourceInputs["publishId"] = state ? state.publishId : undefined;
            resourceInputs["publishedAt"] = state ? state.publishedAt : undefined;
            resourceInputs["region"] = state ? state.region : undefined;
            resourceInputs["versionId"] = state ? state.versionId : undefined;
        } else {
            const args = argsOrState as ApigApiPublishmentArgs | undefined;
            if ((!args || args.apiId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'apiId'");
            }
            if ((!args || args.envId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'envId'");
            }
            if ((!args || args.instanceId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'instanceId'");
            }
            resourceInputs["apiId"] = args ? args.apiId : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["envId"] = args ? args.envId : undefined;
            resourceInputs["instanceId"] = args ? args.instanceId : undefined;
            resourceInputs["region"] = args ? args.region : undefined;
            resourceInputs["versionId"] = args ? args.versionId : undefined;
            resourceInputs["envName"] = undefined /*out*/;
            resourceInputs["histories"] = undefined /*out*/;
            resourceInputs["publishId"] = undefined /*out*/;
            resourceInputs["publishedAt"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(ApigApiPublishment.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ApigApiPublishment resources.
 */
export interface ApigApiPublishmentState {
    /**
     * The ID of the API to be published or already published.
     */
    apiId?: pulumi.Input<string>;
    /**
     * The description of the current publishment.
     */
    description?: pulumi.Input<string>;
    /**
     * The ID of the environment to which the current version of the API will be published or has been published.
     */
    envId?: pulumi.Input<string>;
    /**
     * The name of the environment to which the current version of the API is published.
     */
    envName?: pulumi.Input<string>;
    /**
     * All publish informations of the API.
     */
    histories?: pulumi.Input<pulumi.Input<inputs.ApigApiPublishmentHistory>[]>;
    /**
     * The ID of the dedicated instance to which the API and the environment belongs.
     */
    instanceId?: pulumi.Input<string>;
    /**
     * The publish ID of the API in current environment.
     */
    publishId?: pulumi.Input<string>;
    /**
     * Time when the current version was published.
     */
    publishedAt?: pulumi.Input<string>;
    /**
     * The region in which to publish API.
     */
    region?: pulumi.Input<string>;
    /**
     * The version ID of the current publishment.
     */
    versionId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a ApigApiPublishment resource.
 */
export interface ApigApiPublishmentArgs {
    /**
     * The ID of the API to be published or already published.
     */
    apiId: pulumi.Input<string>;
    /**
     * The description of the current publishment.
     */
    description?: pulumi.Input<string>;
    /**
     * The ID of the environment to which the current version of the API will be published or has been published.
     */
    envId: pulumi.Input<string>;
    /**
     * The ID of the dedicated instance to which the API and the environment belongs.
     */
    instanceId: pulumi.Input<string>;
    /**
     * The region in which to publish API.
     */
    region?: pulumi.Input<string>;
    /**
     * The version ID of the current publishment.
     */
    versionId?: pulumi.Input<string>;
}
