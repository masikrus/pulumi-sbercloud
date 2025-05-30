// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

export function getVpcPeeringConnection(args?: GetVpcPeeringConnectionArgs, opts?: pulumi.InvokeOptions): Promise<GetVpcPeeringConnectionResult> {
    args = args || {};
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("sbercloud:index/getVpcPeeringConnection:getVpcPeeringConnection", {
        "id": args.id,
        "name": args.name,
        "peerTenantId": args.peerTenantId,
        "peerVpcId": args.peerVpcId,
        "region": args.region,
        "status": args.status,
        "vpcId": args.vpcId,
    }, opts);
}

/**
 * A collection of arguments for invoking getVpcPeeringConnection.
 */
export interface GetVpcPeeringConnectionArgs {
    id?: string;
    name?: string;
    peerTenantId?: string;
    peerVpcId?: string;
    region?: string;
    status?: string;
    vpcId?: string;
}

/**
 * A collection of values returned by getVpcPeeringConnection.
 */
export interface GetVpcPeeringConnectionResult {
    readonly description: string;
    readonly id: string;
    readonly name: string;
    readonly peerTenantId: string;
    readonly peerVpcId: string;
    readonly region: string;
    readonly status: string;
    readonly vpcId: string;
}
export function getVpcPeeringConnectionOutput(args?: GetVpcPeeringConnectionOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetVpcPeeringConnectionResult> {
    args = args || {};
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("sbercloud:index/getVpcPeeringConnection:getVpcPeeringConnection", {
        "id": args.id,
        "name": args.name,
        "peerTenantId": args.peerTenantId,
        "peerVpcId": args.peerVpcId,
        "region": args.region,
        "status": args.status,
        "vpcId": args.vpcId,
    }, opts);
}

/**
 * A collection of arguments for invoking getVpcPeeringConnection.
 */
export interface GetVpcPeeringConnectionOutputArgs {
    id?: pulumi.Input<string>;
    name?: pulumi.Input<string>;
    peerTenantId?: pulumi.Input<string>;
    peerVpcId?: pulumi.Input<string>;
    region?: pulumi.Input<string>;
    status?: pulumi.Input<string>;
    vpcId?: pulumi.Input<string>;
}
