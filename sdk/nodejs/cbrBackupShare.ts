// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

export class CbrBackupShare extends pulumi.CustomResource {
    /**
     * Get an existing CbrBackupShare resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: CbrBackupShareState, opts?: pulumi.CustomResourceOptions): CbrBackupShare {
        return new CbrBackupShare(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'sbercloud:index/cbrBackupShare:CbrBackupShare';

    /**
     * Returns true if the given object is an instance of CbrBackupShare.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is CbrBackupShare {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === CbrBackupShare.__pulumiType;
    }

    /**
     * The backup ID.
     */
    public readonly backupId!: pulumi.Output<string>;
    /**
     * The list of shared members configuration.
     */
    public readonly members!: pulumi.Output<outputs.CbrBackupShareMember[]>;
    /**
     * The region where the shared backup is located.
     */
    public readonly region!: pulumi.Output<string>;

    /**
     * Create a CbrBackupShare resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: CbrBackupShareArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: CbrBackupShareArgs | CbrBackupShareState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as CbrBackupShareState | undefined;
            resourceInputs["backupId"] = state ? state.backupId : undefined;
            resourceInputs["members"] = state ? state.members : undefined;
            resourceInputs["region"] = state ? state.region : undefined;
        } else {
            const args = argsOrState as CbrBackupShareArgs | undefined;
            if ((!args || args.backupId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'backupId'");
            }
            if ((!args || args.members === undefined) && !opts.urn) {
                throw new Error("Missing required property 'members'");
            }
            resourceInputs["backupId"] = args ? args.backupId : undefined;
            resourceInputs["members"] = args ? args.members : undefined;
            resourceInputs["region"] = args ? args.region : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(CbrBackupShare.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering CbrBackupShare resources.
 */
export interface CbrBackupShareState {
    /**
     * The backup ID.
     */
    backupId?: pulumi.Input<string>;
    /**
     * The list of shared members configuration.
     */
    members?: pulumi.Input<pulumi.Input<inputs.CbrBackupShareMember>[]>;
    /**
     * The region where the shared backup is located.
     */
    region?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a CbrBackupShare resource.
 */
export interface CbrBackupShareArgs {
    /**
     * The backup ID.
     */
    backupId: pulumi.Input<string>;
    /**
     * The list of shared members configuration.
     */
    members: pulumi.Input<pulumi.Input<inputs.CbrBackupShareMember>[]>;
    /**
     * The region where the shared backup is located.
     */
    region?: pulumi.Input<string>;
}
