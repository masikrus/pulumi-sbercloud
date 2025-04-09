// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

export class ApigPluginAssociate extends pulumi.CustomResource {
    /**
     * Get an existing ApigPluginAssociate resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ApigPluginAssociateState, opts?: pulumi.CustomResourceOptions): ApigPluginAssociate {
        return new ApigPluginAssociate(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'sbercloud:index/apigPluginAssociate:ApigPluginAssociate';

    /**
     * Returns true if the given object is an instance of ApigPluginAssociate.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ApigPluginAssociate {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ApigPluginAssociate.__pulumiType;
    }

    /**
     * The APIs bound by the plugin.
     */
    public readonly apiIds!: pulumi.Output<string[]>;
    /**
     * The environment ID where the API was published.
     */
    public readonly envId!: pulumi.Output<string>;
    /**
     * The ID of the dedicated instance to which the plugin belongs.
     */
    public readonly instanceId!: pulumi.Output<string>;
    /**
     * The plugin ID.
     */
    public readonly pluginId!: pulumi.Output<string>;
    /**
     * The region where the plugin is located.
     */
    public readonly region!: pulumi.Output<string>;

    /**
     * Create a ApigPluginAssociate resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ApigPluginAssociateArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ApigPluginAssociateArgs | ApigPluginAssociateState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ApigPluginAssociateState | undefined;
            resourceInputs["apiIds"] = state ? state.apiIds : undefined;
            resourceInputs["envId"] = state ? state.envId : undefined;
            resourceInputs["instanceId"] = state ? state.instanceId : undefined;
            resourceInputs["pluginId"] = state ? state.pluginId : undefined;
            resourceInputs["region"] = state ? state.region : undefined;
        } else {
            const args = argsOrState as ApigPluginAssociateArgs | undefined;
            if ((!args || args.apiIds === undefined) && !opts.urn) {
                throw new Error("Missing required property 'apiIds'");
            }
            if ((!args || args.envId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'envId'");
            }
            if ((!args || args.instanceId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'instanceId'");
            }
            if ((!args || args.pluginId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'pluginId'");
            }
            resourceInputs["apiIds"] = args ? args.apiIds : undefined;
            resourceInputs["envId"] = args ? args.envId : undefined;
            resourceInputs["instanceId"] = args ? args.instanceId : undefined;
            resourceInputs["pluginId"] = args ? args.pluginId : undefined;
            resourceInputs["region"] = args ? args.region : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(ApigPluginAssociate.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ApigPluginAssociate resources.
 */
export interface ApigPluginAssociateState {
    /**
     * The APIs bound by the plugin.
     */
    apiIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The environment ID where the API was published.
     */
    envId?: pulumi.Input<string>;
    /**
     * The ID of the dedicated instance to which the plugin belongs.
     */
    instanceId?: pulumi.Input<string>;
    /**
     * The plugin ID.
     */
    pluginId?: pulumi.Input<string>;
    /**
     * The region where the plugin is located.
     */
    region?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a ApigPluginAssociate resource.
 */
export interface ApigPluginAssociateArgs {
    /**
     * The APIs bound by the plugin.
     */
    apiIds: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The environment ID where the API was published.
     */
    envId: pulumi.Input<string>;
    /**
     * The ID of the dedicated instance to which the plugin belongs.
     */
    instanceId: pulumi.Input<string>;
    /**
     * The plugin ID.
     */
    pluginId: pulumi.Input<string>;
    /**
     * The region where the plugin is located.
     */
    region?: pulumi.Input<string>;
}
