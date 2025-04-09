// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

export class NetworkAclRule extends pulumi.CustomResource {
    /**
     * Get an existing NetworkAclRule resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: NetworkAclRuleState, opts?: pulumi.CustomResourceOptions): NetworkAclRule {
        return new NetworkAclRule(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'sbercloud:index/networkAclRule:NetworkAclRule';

    /**
     * Returns true if the given object is an instance of NetworkAclRule.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is NetworkAclRule {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === NetworkAclRule.__pulumiType;
    }

    public readonly action!: pulumi.Output<string>;
    public readonly description!: pulumi.Output<string | undefined>;
    public readonly destinationIpAddress!: pulumi.Output<string | undefined>;
    public readonly destinationPort!: pulumi.Output<string | undefined>;
    public readonly enabled!: pulumi.Output<boolean | undefined>;
    public readonly ipVersion!: pulumi.Output<number | undefined>;
    public readonly name!: pulumi.Output<string>;
    public readonly protocol!: pulumi.Output<string>;
    public readonly region!: pulumi.Output<string>;
    public readonly sourceIpAddress!: pulumi.Output<string | undefined>;
    public readonly sourcePort!: pulumi.Output<string | undefined>;

    /**
     * Create a NetworkAclRule resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: NetworkAclRuleArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: NetworkAclRuleArgs | NetworkAclRuleState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as NetworkAclRuleState | undefined;
            resourceInputs["action"] = state ? state.action : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["destinationIpAddress"] = state ? state.destinationIpAddress : undefined;
            resourceInputs["destinationPort"] = state ? state.destinationPort : undefined;
            resourceInputs["enabled"] = state ? state.enabled : undefined;
            resourceInputs["ipVersion"] = state ? state.ipVersion : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["protocol"] = state ? state.protocol : undefined;
            resourceInputs["region"] = state ? state.region : undefined;
            resourceInputs["sourceIpAddress"] = state ? state.sourceIpAddress : undefined;
            resourceInputs["sourcePort"] = state ? state.sourcePort : undefined;
        } else {
            const args = argsOrState as NetworkAclRuleArgs | undefined;
            if ((!args || args.action === undefined) && !opts.urn) {
                throw new Error("Missing required property 'action'");
            }
            if ((!args || args.protocol === undefined) && !opts.urn) {
                throw new Error("Missing required property 'protocol'");
            }
            resourceInputs["action"] = args ? args.action : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["destinationIpAddress"] = args ? args.destinationIpAddress : undefined;
            resourceInputs["destinationPort"] = args ? args.destinationPort : undefined;
            resourceInputs["enabled"] = args ? args.enabled : undefined;
            resourceInputs["ipVersion"] = args ? args.ipVersion : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["protocol"] = args ? args.protocol : undefined;
            resourceInputs["region"] = args ? args.region : undefined;
            resourceInputs["sourceIpAddress"] = args ? args.sourceIpAddress : undefined;
            resourceInputs["sourcePort"] = args ? args.sourcePort : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(NetworkAclRule.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering NetworkAclRule resources.
 */
export interface NetworkAclRuleState {
    action?: pulumi.Input<string>;
    description?: pulumi.Input<string>;
    destinationIpAddress?: pulumi.Input<string>;
    destinationPort?: pulumi.Input<string>;
    enabled?: pulumi.Input<boolean>;
    ipVersion?: pulumi.Input<number>;
    name?: pulumi.Input<string>;
    protocol?: pulumi.Input<string>;
    region?: pulumi.Input<string>;
    sourceIpAddress?: pulumi.Input<string>;
    sourcePort?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a NetworkAclRule resource.
 */
export interface NetworkAclRuleArgs {
    action: pulumi.Input<string>;
    description?: pulumi.Input<string>;
    destinationIpAddress?: pulumi.Input<string>;
    destinationPort?: pulumi.Input<string>;
    enabled?: pulumi.Input<boolean>;
    ipVersion?: pulumi.Input<number>;
    name?: pulumi.Input<string>;
    protocol: pulumi.Input<string>;
    region?: pulumi.Input<string>;
    sourceIpAddress?: pulumi.Input<string>;
    sourcePort?: pulumi.Input<string>;
}
