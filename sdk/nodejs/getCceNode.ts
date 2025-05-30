// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

export function getCceNode(args: GetCceNodeArgs, opts?: pulumi.InvokeOptions): Promise<GetCceNodeResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("sbercloud:index/getCceNode:getCceNode", {
        "clusterId": args.clusterId,
        "name": args.name,
        "nodeId": args.nodeId,
        "region": args.region,
        "status": args.status,
    }, opts);
}

/**
 * A collection of arguments for invoking getCceNode.
 */
export interface GetCceNodeArgs {
    clusterId: string;
    name?: string;
    nodeId?: string;
    region?: string;
    status?: string;
}

/**
 * A collection of values returned by getCceNode.
 */
export interface GetCceNodeResult {
    readonly availabilityZone: string;
    readonly billingMode: number;
    readonly clusterId: string;
    readonly dataVolumes: outputs.GetCceNodeDataVolume[];
    readonly ecsGroupId: string;
    readonly enterpriseProjectId: string;
    readonly flavorId: string;
    readonly hostnameConfigs: outputs.GetCceNodeHostnameConfig[];
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly keyPair: string;
    readonly name: string;
    readonly nodeId: string;
    readonly os: string;
    readonly privateIp: string;
    readonly publicIp: string;
    readonly region: string;
    readonly rootVolumes: outputs.GetCceNodeRootVolume[];
    readonly serverId: string;
    readonly status: string;
    readonly subnetId: string;
    readonly tags: {[key: string]: string};
}
export function getCceNodeOutput(args: GetCceNodeOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetCceNodeResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("sbercloud:index/getCceNode:getCceNode", {
        "clusterId": args.clusterId,
        "name": args.name,
        "nodeId": args.nodeId,
        "region": args.region,
        "status": args.status,
    }, opts);
}

/**
 * A collection of arguments for invoking getCceNode.
 */
export interface GetCceNodeOutputArgs {
    clusterId: pulumi.Input<string>;
    name?: pulumi.Input<string>;
    nodeId?: pulumi.Input<string>;
    region?: pulumi.Input<string>;
    status?: pulumi.Input<string>;
}
