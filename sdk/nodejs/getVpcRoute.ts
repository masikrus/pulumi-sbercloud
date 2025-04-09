// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

export function getVpcRoute(args?: GetVpcRouteArgs, opts?: pulumi.InvokeOptions): Promise<GetVpcRouteResult> {
    args = args || {};
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("sbercloud:index/getVpcRoute:getVpcRoute", {
        "destination": args.destination,
        "id": args.id,
        "region": args.region,
        "tenantId": args.tenantId,
        "type": args.type,
        "vpcId": args.vpcId,
    }, opts);
}

/**
 * A collection of arguments for invoking getVpcRoute.
 */
export interface GetVpcRouteArgs {
    destination?: string;
    id?: string;
    region?: string;
    tenantId?: string;
    type?: string;
    vpcId?: string;
}

/**
 * A collection of values returned by getVpcRoute.
 */
export interface GetVpcRouteResult {
    readonly destination: string;
    readonly id: string;
    readonly nexthop: string;
    readonly region: string;
    readonly tenantId: string;
    readonly type: string;
    readonly vpcId: string;
}
export function getVpcRouteOutput(args?: GetVpcRouteOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetVpcRouteResult> {
    args = args || {};
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("sbercloud:index/getVpcRoute:getVpcRoute", {
        "destination": args.destination,
        "id": args.id,
        "region": args.region,
        "tenantId": args.tenantId,
        "type": args.type,
        "vpcId": args.vpcId,
    }, opts);
}

/**
 * A collection of arguments for invoking getVpcRoute.
 */
export interface GetVpcRouteOutputArgs {
    destination?: pulumi.Input<string>;
    id?: pulumi.Input<string>;
    region?: pulumi.Input<string>;
    tenantId?: pulumi.Input<string>;
    type?: pulumi.Input<string>;
    vpcId?: pulumi.Input<string>;
}
