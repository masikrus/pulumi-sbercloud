// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

export class VpcepEndpoint extends pulumi.CustomResource {
    /**
     * Get an existing VpcepEndpoint resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: VpcepEndpointState, opts?: pulumi.CustomResourceOptions): VpcepEndpoint {
        return new VpcepEndpoint(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'sbercloud:index/vpcepEndpoint:VpcepEndpoint';

    /**
     * Returns true if the given object is an instance of VpcepEndpoint.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is VpcepEndpoint {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === VpcepEndpoint.__pulumiType;
    }

    public readonly description!: pulumi.Output<string>;
    public readonly enableDns!: pulumi.Output<boolean | undefined>;
    public readonly enableWhitelist!: pulumi.Output<boolean | undefined>;
    public readonly ipAddress!: pulumi.Output<string>;
    public readonly networkId!: pulumi.Output<string>;
    public /*out*/ readonly packetId!: pulumi.Output<number>;
    public /*out*/ readonly privateDomainName!: pulumi.Output<string>;
    public readonly region!: pulumi.Output<string>;
    public readonly serviceId!: pulumi.Output<string>;
    public /*out*/ readonly serviceName!: pulumi.Output<string>;
    public /*out*/ readonly serviceType!: pulumi.Output<string>;
    public /*out*/ readonly status!: pulumi.Output<string>;
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    public readonly vpcId!: pulumi.Output<string>;
    public readonly whitelists!: pulumi.Output<string[] | undefined>;

    /**
     * Create a VpcepEndpoint resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: VpcepEndpointArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: VpcepEndpointArgs | VpcepEndpointState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as VpcepEndpointState | undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["enableDns"] = state ? state.enableDns : undefined;
            resourceInputs["enableWhitelist"] = state ? state.enableWhitelist : undefined;
            resourceInputs["ipAddress"] = state ? state.ipAddress : undefined;
            resourceInputs["networkId"] = state ? state.networkId : undefined;
            resourceInputs["packetId"] = state ? state.packetId : undefined;
            resourceInputs["privateDomainName"] = state ? state.privateDomainName : undefined;
            resourceInputs["region"] = state ? state.region : undefined;
            resourceInputs["serviceId"] = state ? state.serviceId : undefined;
            resourceInputs["serviceName"] = state ? state.serviceName : undefined;
            resourceInputs["serviceType"] = state ? state.serviceType : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
            resourceInputs["tags"] = state ? state.tags : undefined;
            resourceInputs["vpcId"] = state ? state.vpcId : undefined;
            resourceInputs["whitelists"] = state ? state.whitelists : undefined;
        } else {
            const args = argsOrState as VpcepEndpointArgs | undefined;
            if ((!args || args.networkId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'networkId'");
            }
            if ((!args || args.serviceId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'serviceId'");
            }
            if ((!args || args.vpcId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'vpcId'");
            }
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["enableDns"] = args ? args.enableDns : undefined;
            resourceInputs["enableWhitelist"] = args ? args.enableWhitelist : undefined;
            resourceInputs["ipAddress"] = args ? args.ipAddress : undefined;
            resourceInputs["networkId"] = args ? args.networkId : undefined;
            resourceInputs["region"] = args ? args.region : undefined;
            resourceInputs["serviceId"] = args ? args.serviceId : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["vpcId"] = args ? args.vpcId : undefined;
            resourceInputs["whitelists"] = args ? args.whitelists : undefined;
            resourceInputs["packetId"] = undefined /*out*/;
            resourceInputs["privateDomainName"] = undefined /*out*/;
            resourceInputs["serviceName"] = undefined /*out*/;
            resourceInputs["serviceType"] = undefined /*out*/;
            resourceInputs["status"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(VpcepEndpoint.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering VpcepEndpoint resources.
 */
export interface VpcepEndpointState {
    description?: pulumi.Input<string>;
    enableDns?: pulumi.Input<boolean>;
    enableWhitelist?: pulumi.Input<boolean>;
    ipAddress?: pulumi.Input<string>;
    networkId?: pulumi.Input<string>;
    packetId?: pulumi.Input<number>;
    privateDomainName?: pulumi.Input<string>;
    region?: pulumi.Input<string>;
    serviceId?: pulumi.Input<string>;
    serviceName?: pulumi.Input<string>;
    serviceType?: pulumi.Input<string>;
    status?: pulumi.Input<string>;
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    vpcId?: pulumi.Input<string>;
    whitelists?: pulumi.Input<pulumi.Input<string>[]>;
}

/**
 * The set of arguments for constructing a VpcepEndpoint resource.
 */
export interface VpcepEndpointArgs {
    description?: pulumi.Input<string>;
    enableDns?: pulumi.Input<boolean>;
    enableWhitelist?: pulumi.Input<boolean>;
    ipAddress?: pulumi.Input<string>;
    networkId: pulumi.Input<string>;
    region?: pulumi.Input<string>;
    serviceId: pulumi.Input<string>;
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    vpcId: pulumi.Input<string>;
    whitelists?: pulumi.Input<pulumi.Input<string>[]>;
}
