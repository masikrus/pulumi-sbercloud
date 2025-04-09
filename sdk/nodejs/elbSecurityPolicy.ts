// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

export class ElbSecurityPolicy extends pulumi.CustomResource {
    /**
     * Get an existing ElbSecurityPolicy resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ElbSecurityPolicyState, opts?: pulumi.CustomResourceOptions): ElbSecurityPolicy {
        return new ElbSecurityPolicy(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'sbercloud:index/elbSecurityPolicy:ElbSecurityPolicy';

    /**
     * Returns true if the given object is an instance of ElbSecurityPolicy.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ElbSecurityPolicy {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ElbSecurityPolicy.__pulumiType;
    }

    /**
     * Specifies the cipher suite list of the security policy.
     */
    public readonly ciphers!: pulumi.Output<string[]>;
    public /*out*/ readonly createdAt!: pulumi.Output<string>;
    /**
     * Specifies the description of the ELB security policy
     */
    public readonly description!: pulumi.Output<string>;
    /**
     * Specifies the enterprise project ID to which the Enterprise router belongs.
     */
    public readonly enterpriseProjectId!: pulumi.Output<string>;
    /**
     * The listener which the security policy associated with.
     */
    public /*out*/ readonly listeners!: pulumi.Output<outputs.ElbSecurityPolicyListener[]>;
    /**
     * Specifies the ELB security policy name.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Specifies the TSL protocol list which the security policy select.
     */
    public readonly protocols!: pulumi.Output<string[]>;
    public readonly region!: pulumi.Output<string>;
    public /*out*/ readonly updatedAt!: pulumi.Output<string>;

    /**
     * Create a ElbSecurityPolicy resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ElbSecurityPolicyArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ElbSecurityPolicyArgs | ElbSecurityPolicyState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ElbSecurityPolicyState | undefined;
            resourceInputs["ciphers"] = state ? state.ciphers : undefined;
            resourceInputs["createdAt"] = state ? state.createdAt : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["enterpriseProjectId"] = state ? state.enterpriseProjectId : undefined;
            resourceInputs["listeners"] = state ? state.listeners : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["protocols"] = state ? state.protocols : undefined;
            resourceInputs["region"] = state ? state.region : undefined;
            resourceInputs["updatedAt"] = state ? state.updatedAt : undefined;
        } else {
            const args = argsOrState as ElbSecurityPolicyArgs | undefined;
            if ((!args || args.ciphers === undefined) && !opts.urn) {
                throw new Error("Missing required property 'ciphers'");
            }
            if ((!args || args.protocols === undefined) && !opts.urn) {
                throw new Error("Missing required property 'protocols'");
            }
            resourceInputs["ciphers"] = args ? args.ciphers : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["enterpriseProjectId"] = args ? args.enterpriseProjectId : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["protocols"] = args ? args.protocols : undefined;
            resourceInputs["region"] = args ? args.region : undefined;
            resourceInputs["createdAt"] = undefined /*out*/;
            resourceInputs["listeners"] = undefined /*out*/;
            resourceInputs["updatedAt"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(ElbSecurityPolicy.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ElbSecurityPolicy resources.
 */
export interface ElbSecurityPolicyState {
    /**
     * Specifies the cipher suite list of the security policy.
     */
    ciphers?: pulumi.Input<pulumi.Input<string>[]>;
    createdAt?: pulumi.Input<string>;
    /**
     * Specifies the description of the ELB security policy
     */
    description?: pulumi.Input<string>;
    /**
     * Specifies the enterprise project ID to which the Enterprise router belongs.
     */
    enterpriseProjectId?: pulumi.Input<string>;
    /**
     * The listener which the security policy associated with.
     */
    listeners?: pulumi.Input<pulumi.Input<inputs.ElbSecurityPolicyListener>[]>;
    /**
     * Specifies the ELB security policy name.
     */
    name?: pulumi.Input<string>;
    /**
     * Specifies the TSL protocol list which the security policy select.
     */
    protocols?: pulumi.Input<pulumi.Input<string>[]>;
    region?: pulumi.Input<string>;
    updatedAt?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a ElbSecurityPolicy resource.
 */
export interface ElbSecurityPolicyArgs {
    /**
     * Specifies the cipher suite list of the security policy.
     */
    ciphers: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Specifies the description of the ELB security policy
     */
    description?: pulumi.Input<string>;
    /**
     * Specifies the enterprise project ID to which the Enterprise router belongs.
     */
    enterpriseProjectId?: pulumi.Input<string>;
    /**
     * Specifies the ELB security policy name.
     */
    name?: pulumi.Input<string>;
    /**
     * Specifies the TSL protocol list which the security policy select.
     */
    protocols: pulumi.Input<pulumi.Input<string>[]>;
    region?: pulumi.Input<string>;
}
