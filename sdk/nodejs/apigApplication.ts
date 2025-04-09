// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

export class ApigApplication extends pulumi.CustomResource {
    /**
     * Get an existing ApigApplication resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ApigApplicationState, opts?: pulumi.CustomResourceOptions): ApigApplication {
        return new ApigApplication(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'sbercloud:index/apigApplication:ApigApplication';

    /**
     * Returns true if the given object is an instance of ApigApplication.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ApigApplication {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ApigApplication.__pulumiType;
    }

    /**
     * The array of one or more application codes that the application has.
     */
    public readonly appCodes!: pulumi.Output<string[]>;
    /**
     * The APP key.
     */
    public /*out*/ readonly appKey!: pulumi.Output<string>;
    /**
     * The APP secret.
     */
    public /*out*/ readonly appSecret!: pulumi.Output<string>;
    /**
     * The application description.
     */
    public readonly description!: pulumi.Output<string>;
    /**
     * The ID of the dedicated instance to which the application belongs.
     */
    public readonly instanceId!: pulumi.Output<string>;
    /**
     * The application name.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The region where the application is located.
     */
    public readonly region!: pulumi.Output<string>;
    /**
     * The registration time.
     */
    public /*out*/ readonly registrationTime!: pulumi.Output<string>;
    /**
     * The secret action to be done for the application.
     */
    public readonly secretAction!: pulumi.Output<string | undefined>;
    /**
     * The latest update time of the application.
     */
    public /*out*/ readonly updatedAt!: pulumi.Output<string>;

    /**
     * Create a ApigApplication resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ApigApplicationArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ApigApplicationArgs | ApigApplicationState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ApigApplicationState | undefined;
            resourceInputs["appCodes"] = state ? state.appCodes : undefined;
            resourceInputs["appKey"] = state ? state.appKey : undefined;
            resourceInputs["appSecret"] = state ? state.appSecret : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["instanceId"] = state ? state.instanceId : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["region"] = state ? state.region : undefined;
            resourceInputs["registrationTime"] = state ? state.registrationTime : undefined;
            resourceInputs["secretAction"] = state ? state.secretAction : undefined;
            resourceInputs["updatedAt"] = state ? state.updatedAt : undefined;
        } else {
            const args = argsOrState as ApigApplicationArgs | undefined;
            if ((!args || args.instanceId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'instanceId'");
            }
            resourceInputs["appCodes"] = args ? args.appCodes : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["instanceId"] = args ? args.instanceId : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["region"] = args ? args.region : undefined;
            resourceInputs["secretAction"] = args ? args.secretAction : undefined;
            resourceInputs["appKey"] = undefined /*out*/;
            resourceInputs["appSecret"] = undefined /*out*/;
            resourceInputs["registrationTime"] = undefined /*out*/;
            resourceInputs["updatedAt"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const secretOpts = { additionalSecretOutputs: ["appSecret"] };
        opts = pulumi.mergeOptions(opts, secretOpts);
        super(ApigApplication.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ApigApplication resources.
 */
export interface ApigApplicationState {
    /**
     * The array of one or more application codes that the application has.
     */
    appCodes?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The APP key.
     */
    appKey?: pulumi.Input<string>;
    /**
     * The APP secret.
     */
    appSecret?: pulumi.Input<string>;
    /**
     * The application description.
     */
    description?: pulumi.Input<string>;
    /**
     * The ID of the dedicated instance to which the application belongs.
     */
    instanceId?: pulumi.Input<string>;
    /**
     * The application name.
     */
    name?: pulumi.Input<string>;
    /**
     * The region where the application is located.
     */
    region?: pulumi.Input<string>;
    /**
     * The registration time.
     */
    registrationTime?: pulumi.Input<string>;
    /**
     * The secret action to be done for the application.
     */
    secretAction?: pulumi.Input<string>;
    /**
     * The latest update time of the application.
     */
    updatedAt?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a ApigApplication resource.
 */
export interface ApigApplicationArgs {
    /**
     * The array of one or more application codes that the application has.
     */
    appCodes?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The application description.
     */
    description?: pulumi.Input<string>;
    /**
     * The ID of the dedicated instance to which the application belongs.
     */
    instanceId: pulumi.Input<string>;
    /**
     * The application name.
     */
    name?: pulumi.Input<string>;
    /**
     * The region where the application is located.
     */
    region?: pulumi.Input<string>;
    /**
     * The secret action to be done for the application.
     */
    secretAction?: pulumi.Input<string>;
}
