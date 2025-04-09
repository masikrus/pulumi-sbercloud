// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

export class ElbMonitor extends pulumi.CustomResource {
    /**
     * Get an existing ElbMonitor resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ElbMonitorState, opts?: pulumi.CustomResourceOptions): ElbMonitor {
        return new ElbMonitor(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'sbercloud:index/elbMonitor:ElbMonitor';

    /**
     * Returns true if the given object is an instance of ElbMonitor.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ElbMonitor {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ElbMonitor.__pulumiType;
    }

    public readonly domainName!: pulumi.Output<string>;
    public readonly interval!: pulumi.Output<number>;
    public readonly maxRetries!: pulumi.Output<number>;
    public readonly poolId!: pulumi.Output<string>;
    public readonly port!: pulumi.Output<number>;
    public readonly protocol!: pulumi.Output<string>;
    public readonly region!: pulumi.Output<string>;
    public readonly statusCode!: pulumi.Output<string>;
    public readonly timeout!: pulumi.Output<number>;
    public readonly urlPath!: pulumi.Output<string>;

    /**
     * Create a ElbMonitor resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ElbMonitorArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ElbMonitorArgs | ElbMonitorState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ElbMonitorState | undefined;
            resourceInputs["domainName"] = state ? state.domainName : undefined;
            resourceInputs["interval"] = state ? state.interval : undefined;
            resourceInputs["maxRetries"] = state ? state.maxRetries : undefined;
            resourceInputs["poolId"] = state ? state.poolId : undefined;
            resourceInputs["port"] = state ? state.port : undefined;
            resourceInputs["protocol"] = state ? state.protocol : undefined;
            resourceInputs["region"] = state ? state.region : undefined;
            resourceInputs["statusCode"] = state ? state.statusCode : undefined;
            resourceInputs["timeout"] = state ? state.timeout : undefined;
            resourceInputs["urlPath"] = state ? state.urlPath : undefined;
        } else {
            const args = argsOrState as ElbMonitorArgs | undefined;
            if ((!args || args.interval === undefined) && !opts.urn) {
                throw new Error("Missing required property 'interval'");
            }
            if ((!args || args.maxRetries === undefined) && !opts.urn) {
                throw new Error("Missing required property 'maxRetries'");
            }
            if ((!args || args.poolId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'poolId'");
            }
            if ((!args || args.protocol === undefined) && !opts.urn) {
                throw new Error("Missing required property 'protocol'");
            }
            if ((!args || args.timeout === undefined) && !opts.urn) {
                throw new Error("Missing required property 'timeout'");
            }
            resourceInputs["domainName"] = args ? args.domainName : undefined;
            resourceInputs["interval"] = args ? args.interval : undefined;
            resourceInputs["maxRetries"] = args ? args.maxRetries : undefined;
            resourceInputs["poolId"] = args ? args.poolId : undefined;
            resourceInputs["port"] = args ? args.port : undefined;
            resourceInputs["protocol"] = args ? args.protocol : undefined;
            resourceInputs["region"] = args ? args.region : undefined;
            resourceInputs["statusCode"] = args ? args.statusCode : undefined;
            resourceInputs["timeout"] = args ? args.timeout : undefined;
            resourceInputs["urlPath"] = args ? args.urlPath : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(ElbMonitor.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ElbMonitor resources.
 */
export interface ElbMonitorState {
    domainName?: pulumi.Input<string>;
    interval?: pulumi.Input<number>;
    maxRetries?: pulumi.Input<number>;
    poolId?: pulumi.Input<string>;
    port?: pulumi.Input<number>;
    protocol?: pulumi.Input<string>;
    region?: pulumi.Input<string>;
    statusCode?: pulumi.Input<string>;
    timeout?: pulumi.Input<number>;
    urlPath?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a ElbMonitor resource.
 */
export interface ElbMonitorArgs {
    domainName?: pulumi.Input<string>;
    interval: pulumi.Input<number>;
    maxRetries: pulumi.Input<number>;
    poolId: pulumi.Input<string>;
    port?: pulumi.Input<number>;
    protocol: pulumi.Input<string>;
    region?: pulumi.Input<string>;
    statusCode?: pulumi.Input<string>;
    timeout: pulumi.Input<number>;
    urlPath?: pulumi.Input<string>;
}
