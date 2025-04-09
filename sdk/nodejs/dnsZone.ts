// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

export class DnsZone extends pulumi.CustomResource {
    /**
     * Get an existing DnsZone resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: DnsZoneState, opts?: pulumi.CustomResourceOptions): DnsZone {
        return new DnsZone(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'sbercloud:index/dnsZone:DnsZone';

    /**
     * Returns true if the given object is an instance of DnsZone.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is DnsZone {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === DnsZone.__pulumiType;
    }

    /**
     * The description of the zone.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * The email address of the administrator managing the zone.
     */
    public readonly email!: pulumi.Output<string>;
    /**
     * The enterprise project ID of the zone.
     */
    public readonly enterpriseProjectId!: pulumi.Output<string>;
    /**
     * The list of the masters of the DNS server.
     */
    public /*out*/ readonly masters!: pulumi.Output<string[]>;
    /**
     * The name of the zone.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The recursive resolution proxy mode for subdomains of the private zone.
     */
    public readonly proxyPattern!: pulumi.Output<string>;
    public readonly region!: pulumi.Output<string>;
    public readonly routers!: pulumi.Output<outputs.DnsZoneRouter[] | undefined>;
    /**
     * The status of the zone.
     */
    public readonly status!: pulumi.Output<string>;
    /**
     * The key/value pairs to associate with the zone.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * The time to live (TTL) of the zone.
     */
    public readonly ttl!: pulumi.Output<number | undefined>;
    /**
     * The type of zone.
     */
    public readonly zoneType!: pulumi.Output<string | undefined>;

    /**
     * Create a DnsZone resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: DnsZoneArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: DnsZoneArgs | DnsZoneState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as DnsZoneState | undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["email"] = state ? state.email : undefined;
            resourceInputs["enterpriseProjectId"] = state ? state.enterpriseProjectId : undefined;
            resourceInputs["masters"] = state ? state.masters : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["proxyPattern"] = state ? state.proxyPattern : undefined;
            resourceInputs["region"] = state ? state.region : undefined;
            resourceInputs["routers"] = state ? state.routers : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
            resourceInputs["tags"] = state ? state.tags : undefined;
            resourceInputs["ttl"] = state ? state.ttl : undefined;
            resourceInputs["zoneType"] = state ? state.zoneType : undefined;
        } else {
            const args = argsOrState as DnsZoneArgs | undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["email"] = args ? args.email : undefined;
            resourceInputs["enterpriseProjectId"] = args ? args.enterpriseProjectId : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["proxyPattern"] = args ? args.proxyPattern : undefined;
            resourceInputs["region"] = args ? args.region : undefined;
            resourceInputs["routers"] = args ? args.routers : undefined;
            resourceInputs["status"] = args ? args.status : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["ttl"] = args ? args.ttl : undefined;
            resourceInputs["zoneType"] = args ? args.zoneType : undefined;
            resourceInputs["masters"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(DnsZone.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering DnsZone resources.
 */
export interface DnsZoneState {
    /**
     * The description of the zone.
     */
    description?: pulumi.Input<string>;
    /**
     * The email address of the administrator managing the zone.
     */
    email?: pulumi.Input<string>;
    /**
     * The enterprise project ID of the zone.
     */
    enterpriseProjectId?: pulumi.Input<string>;
    /**
     * The list of the masters of the DNS server.
     */
    masters?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The name of the zone.
     */
    name?: pulumi.Input<string>;
    /**
     * The recursive resolution proxy mode for subdomains of the private zone.
     */
    proxyPattern?: pulumi.Input<string>;
    region?: pulumi.Input<string>;
    routers?: pulumi.Input<pulumi.Input<inputs.DnsZoneRouter>[]>;
    /**
     * The status of the zone.
     */
    status?: pulumi.Input<string>;
    /**
     * The key/value pairs to associate with the zone.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The time to live (TTL) of the zone.
     */
    ttl?: pulumi.Input<number>;
    /**
     * The type of zone.
     */
    zoneType?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a DnsZone resource.
 */
export interface DnsZoneArgs {
    /**
     * The description of the zone.
     */
    description?: pulumi.Input<string>;
    /**
     * The email address of the administrator managing the zone.
     */
    email?: pulumi.Input<string>;
    /**
     * The enterprise project ID of the zone.
     */
    enterpriseProjectId?: pulumi.Input<string>;
    /**
     * The name of the zone.
     */
    name?: pulumi.Input<string>;
    /**
     * The recursive resolution proxy mode for subdomains of the private zone.
     */
    proxyPattern?: pulumi.Input<string>;
    region?: pulumi.Input<string>;
    routers?: pulumi.Input<pulumi.Input<inputs.DnsZoneRouter>[]>;
    /**
     * The status of the zone.
     */
    status?: pulumi.Input<string>;
    /**
     * The key/value pairs to associate with the zone.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The time to live (TTL) of the zone.
     */
    ttl?: pulumi.Input<number>;
    /**
     * The type of zone.
     */
    zoneType?: pulumi.Input<string>;
}
