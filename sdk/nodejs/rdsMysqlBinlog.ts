// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

export class RdsMysqlBinlog extends pulumi.CustomResource {
    /**
     * Get an existing RdsMysqlBinlog resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: RdsMysqlBinlogState, opts?: pulumi.CustomResourceOptions): RdsMysqlBinlog {
        return new RdsMysqlBinlog(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'sbercloud:index/rdsMysqlBinlog:RdsMysqlBinlog';

    /**
     * Returns true if the given object is an instance of RdsMysqlBinlog.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is RdsMysqlBinlog {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === RdsMysqlBinlog.__pulumiType;
    }

    public readonly binlogRetentionHours!: pulumi.Output<number>;
    public readonly instanceId!: pulumi.Output<string>;
    public readonly region!: pulumi.Output<string>;

    /**
     * Create a RdsMysqlBinlog resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: RdsMysqlBinlogArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: RdsMysqlBinlogArgs | RdsMysqlBinlogState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as RdsMysqlBinlogState | undefined;
            resourceInputs["binlogRetentionHours"] = state ? state.binlogRetentionHours : undefined;
            resourceInputs["instanceId"] = state ? state.instanceId : undefined;
            resourceInputs["region"] = state ? state.region : undefined;
        } else {
            const args = argsOrState as RdsMysqlBinlogArgs | undefined;
            if ((!args || args.binlogRetentionHours === undefined) && !opts.urn) {
                throw new Error("Missing required property 'binlogRetentionHours'");
            }
            if ((!args || args.instanceId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'instanceId'");
            }
            resourceInputs["binlogRetentionHours"] = args ? args.binlogRetentionHours : undefined;
            resourceInputs["instanceId"] = args ? args.instanceId : undefined;
            resourceInputs["region"] = args ? args.region : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(RdsMysqlBinlog.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering RdsMysqlBinlog resources.
 */
export interface RdsMysqlBinlogState {
    binlogRetentionHours?: pulumi.Input<number>;
    instanceId?: pulumi.Input<string>;
    region?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a RdsMysqlBinlog resource.
 */
export interface RdsMysqlBinlogArgs {
    binlogRetentionHours: pulumi.Input<number>;
    instanceId: pulumi.Input<string>;
    region?: pulumi.Input<string>;
}
