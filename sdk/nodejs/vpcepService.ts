// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

export class VpcepService extends pulumi.CustomResource {
    /**
     * Get an existing VpcepService resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: VpcepServiceState, opts?: pulumi.CustomResourceOptions): VpcepService {
        return new VpcepService(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'sbercloud:index/vpcepService:VpcepService';

    /**
     * Returns true if the given object is an instance of VpcepService.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is VpcepService {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === VpcepService.__pulumiType;
    }

    public readonly approval!: pulumi.Output<boolean>;
    public /*out*/ readonly connections!: pulumi.Output<outputs.VpcepServiceConnection[]>;
    public readonly description!: pulumi.Output<string>;
    public readonly enablePolicy!: pulumi.Output<boolean>;
    public readonly name!: pulumi.Output<string>;
    public readonly organizationPermissions!: pulumi.Output<string[]>;
    public readonly permissions!: pulumi.Output<string[] | undefined>;
    public readonly portId!: pulumi.Output<string>;
    public readonly portMappings!: pulumi.Output<outputs.VpcepServicePortMapping[]>;
    public readonly region!: pulumi.Output<string>;
    public readonly serverType!: pulumi.Output<string>;
    public /*out*/ readonly serviceName!: pulumi.Output<string>;
    /**
     * schema: Computed
     */
    public readonly serviceType!: pulumi.Output<string | undefined>;
    public /*out*/ readonly status!: pulumi.Output<string>;
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    public readonly vpcId!: pulumi.Output<string>;

    /**
     * Create a VpcepService resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: VpcepServiceArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: VpcepServiceArgs | VpcepServiceState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as VpcepServiceState | undefined;
            resourceInputs["approval"] = state ? state.approval : undefined;
            resourceInputs["connections"] = state ? state.connections : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["enablePolicy"] = state ? state.enablePolicy : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["organizationPermissions"] = state ? state.organizationPermissions : undefined;
            resourceInputs["permissions"] = state ? state.permissions : undefined;
            resourceInputs["portId"] = state ? state.portId : undefined;
            resourceInputs["portMappings"] = state ? state.portMappings : undefined;
            resourceInputs["region"] = state ? state.region : undefined;
            resourceInputs["serverType"] = state ? state.serverType : undefined;
            resourceInputs["serviceName"] = state ? state.serviceName : undefined;
            resourceInputs["serviceType"] = state ? state.serviceType : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
            resourceInputs["tags"] = state ? state.tags : undefined;
            resourceInputs["vpcId"] = state ? state.vpcId : undefined;
        } else {
            const args = argsOrState as VpcepServiceArgs | undefined;
            if ((!args || args.portId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'portId'");
            }
            if ((!args || args.portMappings === undefined) && !opts.urn) {
                throw new Error("Missing required property 'portMappings'");
            }
            if ((!args || args.serverType === undefined) && !opts.urn) {
                throw new Error("Missing required property 'serverType'");
            }
            if ((!args || args.vpcId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'vpcId'");
            }
            resourceInputs["approval"] = args ? args.approval : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["enablePolicy"] = args ? args.enablePolicy : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["organizationPermissions"] = args ? args.organizationPermissions : undefined;
            resourceInputs["permissions"] = args ? args.permissions : undefined;
            resourceInputs["portId"] = args ? args.portId : undefined;
            resourceInputs["portMappings"] = args ? args.portMappings : undefined;
            resourceInputs["region"] = args ? args.region : undefined;
            resourceInputs["serverType"] = args ? args.serverType : undefined;
            resourceInputs["serviceType"] = args ? args.serviceType : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["vpcId"] = args ? args.vpcId : undefined;
            resourceInputs["connections"] = undefined /*out*/;
            resourceInputs["serviceName"] = undefined /*out*/;
            resourceInputs["status"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(VpcepService.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering VpcepService resources.
 */
export interface VpcepServiceState {
    approval?: pulumi.Input<boolean>;
    connections?: pulumi.Input<pulumi.Input<inputs.VpcepServiceConnection>[]>;
    description?: pulumi.Input<string>;
    enablePolicy?: pulumi.Input<boolean>;
    name?: pulumi.Input<string>;
    organizationPermissions?: pulumi.Input<pulumi.Input<string>[]>;
    permissions?: pulumi.Input<pulumi.Input<string>[]>;
    portId?: pulumi.Input<string>;
    portMappings?: pulumi.Input<pulumi.Input<inputs.VpcepServicePortMapping>[]>;
    region?: pulumi.Input<string>;
    serverType?: pulumi.Input<string>;
    serviceName?: pulumi.Input<string>;
    /**
     * schema: Computed
     */
    serviceType?: pulumi.Input<string>;
    status?: pulumi.Input<string>;
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    vpcId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a VpcepService resource.
 */
export interface VpcepServiceArgs {
    approval?: pulumi.Input<boolean>;
    description?: pulumi.Input<string>;
    enablePolicy?: pulumi.Input<boolean>;
    name?: pulumi.Input<string>;
    organizationPermissions?: pulumi.Input<pulumi.Input<string>[]>;
    permissions?: pulumi.Input<pulumi.Input<string>[]>;
    portId: pulumi.Input<string>;
    portMappings: pulumi.Input<pulumi.Input<inputs.VpcepServicePortMapping>[]>;
    region?: pulumi.Input<string>;
    serverType: pulumi.Input<string>;
    /**
     * schema: Computed
     */
    serviceType?: pulumi.Input<string>;
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    vpcId: pulumi.Input<string>;
}
