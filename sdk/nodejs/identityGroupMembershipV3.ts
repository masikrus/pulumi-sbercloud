// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

export class IdentityGroupMembershipV3 extends pulumi.CustomResource {
    /**
     * Get an existing IdentityGroupMembershipV3 resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: IdentityGroupMembershipV3State, opts?: pulumi.CustomResourceOptions): IdentityGroupMembershipV3 {
        return new IdentityGroupMembershipV3(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'sbercloud:index/identityGroupMembershipV3:IdentityGroupMembershipV3';

    /**
     * Returns true if the given object is an instance of IdentityGroupMembershipV3.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is IdentityGroupMembershipV3 {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === IdentityGroupMembershipV3.__pulumiType;
    }

    public readonly group!: pulumi.Output<string>;
    public readonly users!: pulumi.Output<string[]>;

    /**
     * Create a IdentityGroupMembershipV3 resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: IdentityGroupMembershipV3Args, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: IdentityGroupMembershipV3Args | IdentityGroupMembershipV3State, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as IdentityGroupMembershipV3State | undefined;
            resourceInputs["group"] = state ? state.group : undefined;
            resourceInputs["users"] = state ? state.users : undefined;
        } else {
            const args = argsOrState as IdentityGroupMembershipV3Args | undefined;
            if ((!args || args.group === undefined) && !opts.urn) {
                throw new Error("Missing required property 'group'");
            }
            if ((!args || args.users === undefined) && !opts.urn) {
                throw new Error("Missing required property 'users'");
            }
            resourceInputs["group"] = args ? args.group : undefined;
            resourceInputs["users"] = args ? args.users : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(IdentityGroupMembershipV3.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering IdentityGroupMembershipV3 resources.
 */
export interface IdentityGroupMembershipV3State {
    group?: pulumi.Input<string>;
    users?: pulumi.Input<pulumi.Input<string>[]>;
}

/**
 * The set of arguments for constructing a IdentityGroupMembershipV3 resource.
 */
export interface IdentityGroupMembershipV3Args {
    group: pulumi.Input<string>;
    users: pulumi.Input<pulumi.Input<string>[]>;
}
