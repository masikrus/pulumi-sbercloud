// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

export function getNatGateway(args?: GetNatGatewayArgs, opts?: pulumi.InvokeOptions): Promise<GetNatGatewayResult> {
    args = args || {};
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("sbercloud:index/getNatGateway:getNatGateway", {
        "description": args.description,
        "enterpriseProjectId": args.enterpriseProjectId,
        "id": args.id,
        "internalNetworkId": args.internalNetworkId,
        "name": args.name,
        "region": args.region,
        "routerId": args.routerId,
        "spec": args.spec,
        "status": args.status,
        "subnetId": args.subnetId,
        "vpcId": args.vpcId,
    }, opts);
}

/**
 * A collection of arguments for invoking getNatGateway.
 */
export interface GetNatGatewayArgs {
    description?: string;
    enterpriseProjectId?: string;
    id?: string;
    /**
     * @deprecated use subnetId instead
     */
    internalNetworkId?: string;
    name?: string;
    region?: string;
    /**
     * @deprecated use vpcId instead
     */
    routerId?: string;
    spec?: string;
    status?: string;
    subnetId?: string;
    vpcId?: string;
}

/**
 * A collection of values returned by getNatGateway.
 */
export interface GetNatGatewayResult {
    readonly description: string;
    readonly enterpriseProjectId: string;
    readonly id: string;
    /**
     * @deprecated use subnetId instead
     */
    readonly internalNetworkId?: string;
    readonly name: string;
    readonly region: string;
    /**
     * @deprecated use vpcId instead
     */
    readonly routerId?: string;
    readonly spec: string;
    readonly status: string;
    readonly subnetId: string;
    readonly vpcId: string;
}
export function getNatGatewayOutput(args?: GetNatGatewayOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetNatGatewayResult> {
    args = args || {};
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("sbercloud:index/getNatGateway:getNatGateway", {
        "description": args.description,
        "enterpriseProjectId": args.enterpriseProjectId,
        "id": args.id,
        "internalNetworkId": args.internalNetworkId,
        "name": args.name,
        "region": args.region,
        "routerId": args.routerId,
        "spec": args.spec,
        "status": args.status,
        "subnetId": args.subnetId,
        "vpcId": args.vpcId,
    }, opts);
}

/**
 * A collection of arguments for invoking getNatGateway.
 */
export interface GetNatGatewayOutputArgs {
    description?: pulumi.Input<string>;
    enterpriseProjectId?: pulumi.Input<string>;
    id?: pulumi.Input<string>;
    /**
     * @deprecated use subnetId instead
     */
    internalNetworkId?: pulumi.Input<string>;
    name?: pulumi.Input<string>;
    region?: pulumi.Input<string>;
    /**
     * @deprecated use vpcId instead
     */
    routerId?: pulumi.Input<string>;
    spec?: pulumi.Input<string>;
    status?: pulumi.Input<string>;
    subnetId?: pulumi.Input<string>;
    vpcId?: pulumi.Input<string>;
}
