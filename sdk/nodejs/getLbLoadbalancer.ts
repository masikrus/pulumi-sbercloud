// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

export function getLbLoadbalancer(args?: GetLbLoadbalancerArgs, opts?: pulumi.InvokeOptions): Promise<GetLbLoadbalancerResult> {
    args = args || {};
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("sbercloud:index/getLbLoadbalancer:getLbLoadbalancer", {
        "description": args.description,
        "enterpriseProjectId": args.enterpriseProjectId,
        "id": args.id,
        "name": args.name,
        "region": args.region,
        "status": args.status,
        "vipAddress": args.vipAddress,
        "vipSubnetId": args.vipSubnetId,
    }, opts);
}

/**
 * A collection of arguments for invoking getLbLoadbalancer.
 */
export interface GetLbLoadbalancerArgs {
    description?: string;
    enterpriseProjectId?: string;
    id?: string;
    name?: string;
    region?: string;
    status?: string;
    vipAddress?: string;
    vipSubnetId?: string;
}

/**
 * A collection of values returned by getLbLoadbalancer.
 */
export interface GetLbLoadbalancerResult {
    readonly description: string;
    readonly enterpriseProjectId: string;
    readonly id: string;
    readonly name: string;
    readonly publicIp: string;
    readonly region: string;
    readonly status: string;
    readonly tags: {[key: string]: string};
    readonly vipAddress: string;
    readonly vipPortId: string;
    readonly vipSubnetId: string;
}
export function getLbLoadbalancerOutput(args?: GetLbLoadbalancerOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetLbLoadbalancerResult> {
    args = args || {};
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("sbercloud:index/getLbLoadbalancer:getLbLoadbalancer", {
        "description": args.description,
        "enterpriseProjectId": args.enterpriseProjectId,
        "id": args.id,
        "name": args.name,
        "region": args.region,
        "status": args.status,
        "vipAddress": args.vipAddress,
        "vipSubnetId": args.vipSubnetId,
    }, opts);
}

/**
 * A collection of arguments for invoking getLbLoadbalancer.
 */
export interface GetLbLoadbalancerOutputArgs {
    description?: pulumi.Input<string>;
    enterpriseProjectId?: pulumi.Input<string>;
    id?: pulumi.Input<string>;
    name?: pulumi.Input<string>;
    region?: pulumi.Input<string>;
    status?: pulumi.Input<string>;
    vipAddress?: pulumi.Input<string>;
    vipSubnetId?: pulumi.Input<string>;
}
