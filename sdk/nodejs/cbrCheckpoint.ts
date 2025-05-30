// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

export class CbrCheckpoint extends pulumi.CustomResource {
    /**
     * Get an existing CbrCheckpoint resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: CbrCheckpointState, opts?: pulumi.CustomResourceOptions): CbrCheckpoint {
        return new CbrCheckpoint(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'sbercloud:index/cbrCheckpoint:CbrCheckpoint';

    /**
     * Returns true if the given object is an instance of CbrCheckpoint.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is CbrCheckpoint {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === CbrCheckpoint.__pulumiType;
    }

    /**
     * The list of backups configuration.
     */
    public readonly backups!: pulumi.Output<outputs.CbrCheckpointBackup[]>;
    /**
     * The creation time of the checkpoint.
     */
    public /*out*/ readonly createdAt!: pulumi.Output<string>;
    /**
     * The description of the checkpoint.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * Whether the backups are incremental backups.
     */
    public readonly incremental!: pulumi.Output<boolean | undefined>;
    /**
     * The name of the checkpoint.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The region where the vault and backup resources are located.
     */
    public readonly region!: pulumi.Output<string>;
    /**
     * The status of the checkpoint.
     */
    public /*out*/ readonly status!: pulumi.Output<string>;
    /**
     * The ID of the vault where the checkpoint to create.
     */
    public readonly vaultId!: pulumi.Output<string>;

    /**
     * Create a CbrCheckpoint resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: CbrCheckpointArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: CbrCheckpointArgs | CbrCheckpointState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as CbrCheckpointState | undefined;
            resourceInputs["backups"] = state ? state.backups : undefined;
            resourceInputs["createdAt"] = state ? state.createdAt : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["incremental"] = state ? state.incremental : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["region"] = state ? state.region : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
            resourceInputs["vaultId"] = state ? state.vaultId : undefined;
        } else {
            const args = argsOrState as CbrCheckpointArgs | undefined;
            if ((!args || args.backups === undefined) && !opts.urn) {
                throw new Error("Missing required property 'backups'");
            }
            if ((!args || args.vaultId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'vaultId'");
            }
            resourceInputs["backups"] = args ? args.backups : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["incremental"] = args ? args.incremental : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["region"] = args ? args.region : undefined;
            resourceInputs["vaultId"] = args ? args.vaultId : undefined;
            resourceInputs["createdAt"] = undefined /*out*/;
            resourceInputs["status"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(CbrCheckpoint.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering CbrCheckpoint resources.
 */
export interface CbrCheckpointState {
    /**
     * The list of backups configuration.
     */
    backups?: pulumi.Input<pulumi.Input<inputs.CbrCheckpointBackup>[]>;
    /**
     * The creation time of the checkpoint.
     */
    createdAt?: pulumi.Input<string>;
    /**
     * The description of the checkpoint.
     */
    description?: pulumi.Input<string>;
    /**
     * Whether the backups are incremental backups.
     */
    incremental?: pulumi.Input<boolean>;
    /**
     * The name of the checkpoint.
     */
    name?: pulumi.Input<string>;
    /**
     * The region where the vault and backup resources are located.
     */
    region?: pulumi.Input<string>;
    /**
     * The status of the checkpoint.
     */
    status?: pulumi.Input<string>;
    /**
     * The ID of the vault where the checkpoint to create.
     */
    vaultId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a CbrCheckpoint resource.
 */
export interface CbrCheckpointArgs {
    /**
     * The list of backups configuration.
     */
    backups: pulumi.Input<pulumi.Input<inputs.CbrCheckpointBackup>[]>;
    /**
     * The description of the checkpoint.
     */
    description?: pulumi.Input<string>;
    /**
     * Whether the backups are incremental backups.
     */
    incremental?: pulumi.Input<boolean>;
    /**
     * The name of the checkpoint.
     */
    name?: pulumi.Input<string>;
    /**
     * The region where the vault and backup resources are located.
     */
    region?: pulumi.Input<string>;
    /**
     * The ID of the vault where the checkpoint to create.
     */
    vaultId: pulumi.Input<string>;
}
