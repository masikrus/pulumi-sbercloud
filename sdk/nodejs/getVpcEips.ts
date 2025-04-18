// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

export function getVpcEips(args?: GetVpcEipsArgs, opts?: pulumi.InvokeOptions): Promise<GetVpcEipsResult> {
    args = args || {};
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("sbercloud:index/getVpcEips:getVpcEips", {
        "enterpriseProjectId": args.enterpriseProjectId,
        "ids": args.ids,
        "ipVersion": args.ipVersion,
        "portIds": args.portIds,
        "privateIps": args.privateIps,
        "publicIps": args.publicIps,
        "region": args.region,
        "tags": args.tags,
    }, opts);
}

/**
 * A collection of arguments for invoking getVpcEips.
 */
export interface GetVpcEipsArgs {
    enterpriseProjectId?: string;
    ids?: string[];
    ipVersion?: number;
    portIds?: string[];
    privateIps?: string[];
    publicIps?: string[];
    region?: string;
    tags?: {[key: string]: string};
}

/**
 * A collection of values returned by getVpcEips.
 */
export interface GetVpcEipsResult {
    readonly eips: outputs.GetVpcEipsEip[];
    readonly enterpriseProjectId?: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly ids?: string[];
    readonly ipVersion?: number;
    readonly portIds?: string[];
    readonly privateIps?: string[];
    readonly publicIps?: string[];
    readonly region: string;
    readonly tags?: {[key: string]: string};
}
export function getVpcEipsOutput(args?: GetVpcEipsOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetVpcEipsResult> {
    args = args || {};
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("sbercloud:index/getVpcEips:getVpcEips", {
        "enterpriseProjectId": args.enterpriseProjectId,
        "ids": args.ids,
        "ipVersion": args.ipVersion,
        "portIds": args.portIds,
        "privateIps": args.privateIps,
        "publicIps": args.publicIps,
        "region": args.region,
        "tags": args.tags,
    }, opts);
}

/**
 * A collection of arguments for invoking getVpcEips.
 */
export interface GetVpcEipsOutputArgs {
    enterpriseProjectId?: pulumi.Input<string>;
    ids?: pulumi.Input<pulumi.Input<string>[]>;
    ipVersion?: pulumi.Input<number>;
    portIds?: pulumi.Input<pulumi.Input<string>[]>;
    privateIps?: pulumi.Input<pulumi.Input<string>[]>;
    publicIps?: pulumi.Input<pulumi.Input<string>[]>;
    region?: pulumi.Input<string>;
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}
