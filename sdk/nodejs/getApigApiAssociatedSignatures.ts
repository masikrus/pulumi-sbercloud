// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

export function getApigApiAssociatedSignatures(args: GetApigApiAssociatedSignaturesArgs, opts?: pulumi.InvokeOptions): Promise<GetApigApiAssociatedSignaturesResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("sbercloud:index/getApigApiAssociatedSignatures:getApigApiAssociatedSignatures", {
        "apiId": args.apiId,
        "envId": args.envId,
        "envName": args.envName,
        "instanceId": args.instanceId,
        "name": args.name,
        "region": args.region,
        "signatureId": args.signatureId,
        "type": args.type,
    }, opts);
}

/**
 * A collection of arguments for invoking getApigApiAssociatedSignatures.
 */
export interface GetApigApiAssociatedSignaturesArgs {
    apiId: string;
    envId?: string;
    envName?: string;
    instanceId: string;
    name?: string;
    region?: string;
    signatureId?: string;
    type?: string;
}

/**
 * A collection of values returned by getApigApiAssociatedSignatures.
 */
export interface GetApigApiAssociatedSignaturesResult {
    readonly apiId: string;
    readonly envId?: string;
    readonly envName?: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly instanceId: string;
    readonly name?: string;
    readonly region: string;
    readonly signatureId?: string;
    readonly signatures: outputs.GetApigApiAssociatedSignaturesSignature[];
    readonly type?: string;
}
export function getApigApiAssociatedSignaturesOutput(args: GetApigApiAssociatedSignaturesOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetApigApiAssociatedSignaturesResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("sbercloud:index/getApigApiAssociatedSignatures:getApigApiAssociatedSignatures", {
        "apiId": args.apiId,
        "envId": args.envId,
        "envName": args.envName,
        "instanceId": args.instanceId,
        "name": args.name,
        "region": args.region,
        "signatureId": args.signatureId,
        "type": args.type,
    }, opts);
}

/**
 * A collection of arguments for invoking getApigApiAssociatedSignatures.
 */
export interface GetApigApiAssociatedSignaturesOutputArgs {
    apiId: pulumi.Input<string>;
    envId?: pulumi.Input<string>;
    envName?: pulumi.Input<string>;
    instanceId: pulumi.Input<string>;
    name?: pulumi.Input<string>;
    region?: pulumi.Input<string>;
    signatureId?: pulumi.Input<string>;
    type?: pulumi.Input<string>;
}
