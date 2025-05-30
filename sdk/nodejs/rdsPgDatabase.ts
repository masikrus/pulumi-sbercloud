// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

export class RdsPgDatabase extends pulumi.CustomResource {
    /**
     * Get an existing RdsPgDatabase resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: RdsPgDatabaseState, opts?: pulumi.CustomResourceOptions): RdsPgDatabase {
        return new RdsPgDatabase(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'sbercloud:index/rdsPgDatabase:RdsPgDatabase';

    /**
     * Returns true if the given object is an instance of RdsPgDatabase.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is RdsPgDatabase {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === RdsPgDatabase.__pulumiType;
    }

    /**
     * Specifies the database character set.
     */
    public readonly characterSet!: pulumi.Output<string>;
    /**
     * Specifies the database description.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * Specifies the ID of the RDS PostgreSQL instance.
     */
    public readonly instanceId!: pulumi.Output<string>;
    /**
     * Specifies whether to revoke the PUBLIC CREATE permission of the public schema.
     */
    public readonly isRevokePublicPrivilege!: pulumi.Output<boolean>;
    /**
     * Specifies the database collocation.
     */
    public readonly lcCollate!: pulumi.Output<string>;
    /**
     * Specifies the database classification.
     */
    public readonly lcCtype!: pulumi.Output<string>;
    /**
     * Specifies the database name.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Specifies the database user.
     */
    public readonly owner!: pulumi.Output<string>;
    public readonly region!: pulumi.Output<string>;
    /**
     * Indicates the database size, in bytes.
     */
    public /*out*/ readonly size!: pulumi.Output<number>;
    /**
     * Specifies the name of the database template.
     */
    public readonly template!: pulumi.Output<string>;

    /**
     * Create a RdsPgDatabase resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: RdsPgDatabaseArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: RdsPgDatabaseArgs | RdsPgDatabaseState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as RdsPgDatabaseState | undefined;
            resourceInputs["characterSet"] = state ? state.characterSet : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["instanceId"] = state ? state.instanceId : undefined;
            resourceInputs["isRevokePublicPrivilege"] = state ? state.isRevokePublicPrivilege : undefined;
            resourceInputs["lcCollate"] = state ? state.lcCollate : undefined;
            resourceInputs["lcCtype"] = state ? state.lcCtype : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["owner"] = state ? state.owner : undefined;
            resourceInputs["region"] = state ? state.region : undefined;
            resourceInputs["size"] = state ? state.size : undefined;
            resourceInputs["template"] = state ? state.template : undefined;
        } else {
            const args = argsOrState as RdsPgDatabaseArgs | undefined;
            if ((!args || args.instanceId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'instanceId'");
            }
            resourceInputs["characterSet"] = args ? args.characterSet : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["instanceId"] = args ? args.instanceId : undefined;
            resourceInputs["isRevokePublicPrivilege"] = args ? args.isRevokePublicPrivilege : undefined;
            resourceInputs["lcCollate"] = args ? args.lcCollate : undefined;
            resourceInputs["lcCtype"] = args ? args.lcCtype : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["owner"] = args ? args.owner : undefined;
            resourceInputs["region"] = args ? args.region : undefined;
            resourceInputs["template"] = args ? args.template : undefined;
            resourceInputs["size"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(RdsPgDatabase.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering RdsPgDatabase resources.
 */
export interface RdsPgDatabaseState {
    /**
     * Specifies the database character set.
     */
    characterSet?: pulumi.Input<string>;
    /**
     * Specifies the database description.
     */
    description?: pulumi.Input<string>;
    /**
     * Specifies the ID of the RDS PostgreSQL instance.
     */
    instanceId?: pulumi.Input<string>;
    /**
     * Specifies whether to revoke the PUBLIC CREATE permission of the public schema.
     */
    isRevokePublicPrivilege?: pulumi.Input<boolean>;
    /**
     * Specifies the database collocation.
     */
    lcCollate?: pulumi.Input<string>;
    /**
     * Specifies the database classification.
     */
    lcCtype?: pulumi.Input<string>;
    /**
     * Specifies the database name.
     */
    name?: pulumi.Input<string>;
    /**
     * Specifies the database user.
     */
    owner?: pulumi.Input<string>;
    region?: pulumi.Input<string>;
    /**
     * Indicates the database size, in bytes.
     */
    size?: pulumi.Input<number>;
    /**
     * Specifies the name of the database template.
     */
    template?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a RdsPgDatabase resource.
 */
export interface RdsPgDatabaseArgs {
    /**
     * Specifies the database character set.
     */
    characterSet?: pulumi.Input<string>;
    /**
     * Specifies the database description.
     */
    description?: pulumi.Input<string>;
    /**
     * Specifies the ID of the RDS PostgreSQL instance.
     */
    instanceId: pulumi.Input<string>;
    /**
     * Specifies whether to revoke the PUBLIC CREATE permission of the public schema.
     */
    isRevokePublicPrivilege?: pulumi.Input<boolean>;
    /**
     * Specifies the database collocation.
     */
    lcCollate?: pulumi.Input<string>;
    /**
     * Specifies the database classification.
     */
    lcCtype?: pulumi.Input<string>;
    /**
     * Specifies the database name.
     */
    name?: pulumi.Input<string>;
    /**
     * Specifies the database user.
     */
    owner?: pulumi.Input<string>;
    region?: pulumi.Input<string>;
    /**
     * Specifies the name of the database template.
     */
    template?: pulumi.Input<string>;
}
