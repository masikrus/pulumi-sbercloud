// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

export class SwrOrganizationPermissions extends pulumi.CustomResource {
    /**
     * Get an existing SwrOrganizationPermissions resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: SwrOrganizationPermissionsState, opts?: pulumi.CustomResourceOptions): SwrOrganizationPermissions {
        return new SwrOrganizationPermissions(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'sbercloud:index/swrOrganizationPermissions:SwrOrganizationPermissions';

    /**
     * Returns true if the given object is an instance of SwrOrganizationPermissions.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is SwrOrganizationPermissions {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === SwrOrganizationPermissions.__pulumiType;
    }

    public /*out*/ readonly creator!: pulumi.Output<string>;
    public readonly organization!: pulumi.Output<string>;
    public readonly region!: pulumi.Output<string>;
    public /*out*/ readonly selfPermissions!: pulumi.Output<outputs.SwrOrganizationPermissionsSelfPermission[]>;
    public readonly users!: pulumi.Output<outputs.SwrOrganizationPermissionsUser[]>;

    /**
     * Create a SwrOrganizationPermissions resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: SwrOrganizationPermissionsArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: SwrOrganizationPermissionsArgs | SwrOrganizationPermissionsState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as SwrOrganizationPermissionsState | undefined;
            resourceInputs["creator"] = state ? state.creator : undefined;
            resourceInputs["organization"] = state ? state.organization : undefined;
            resourceInputs["region"] = state ? state.region : undefined;
            resourceInputs["selfPermissions"] = state ? state.selfPermissions : undefined;
            resourceInputs["users"] = state ? state.users : undefined;
        } else {
            const args = argsOrState as SwrOrganizationPermissionsArgs | undefined;
            if ((!args || args.organization === undefined) && !opts.urn) {
                throw new Error("Missing required property 'organization'");
            }
            if ((!args || args.users === undefined) && !opts.urn) {
                throw new Error("Missing required property 'users'");
            }
            resourceInputs["organization"] = args ? args.organization : undefined;
            resourceInputs["region"] = args ? args.region : undefined;
            resourceInputs["users"] = args ? args.users : undefined;
            resourceInputs["creator"] = undefined /*out*/;
            resourceInputs["selfPermissions"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(SwrOrganizationPermissions.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering SwrOrganizationPermissions resources.
 */
export interface SwrOrganizationPermissionsState {
    creator?: pulumi.Input<string>;
    organization?: pulumi.Input<string>;
    region?: pulumi.Input<string>;
    selfPermissions?: pulumi.Input<pulumi.Input<inputs.SwrOrganizationPermissionsSelfPermission>[]>;
    users?: pulumi.Input<pulumi.Input<inputs.SwrOrganizationPermissionsUser>[]>;
}

/**
 * The set of arguments for constructing a SwrOrganizationPermissions resource.
 */
export interface SwrOrganizationPermissionsArgs {
    organization: pulumi.Input<string>;
    region?: pulumi.Input<string>;
    users: pulumi.Input<pulumi.Input<inputs.SwrOrganizationPermissionsUser>[]>;
}
