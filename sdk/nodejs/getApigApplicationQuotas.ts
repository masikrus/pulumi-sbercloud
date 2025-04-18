// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

export function getApigApplicationQuotas(args: GetApigApplicationQuotasArgs, opts?: pulumi.InvokeOptions): Promise<GetApigApplicationQuotasResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("sbercloud:index/getApigApplicationQuotas:getApigApplicationQuotas", {
        "instanceId": args.instanceId,
        "name": args.name,
        "quotaId": args.quotaId,
        "region": args.region,
    }, opts);
}

/**
 * A collection of arguments for invoking getApigApplicationQuotas.
 */
export interface GetApigApplicationQuotasArgs {
    instanceId: string;
    name?: string;
    quotaId?: string;
    region?: string;
}

/**
 * A collection of values returned by getApigApplicationQuotas.
 */
export interface GetApigApplicationQuotasResult {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly instanceId: string;
    readonly name?: string;
    readonly quotaId?: string;
    readonly quotas: outputs.GetApigApplicationQuotasQuota[];
    readonly region: string;
}
export function getApigApplicationQuotasOutput(args: GetApigApplicationQuotasOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetApigApplicationQuotasResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("sbercloud:index/getApigApplicationQuotas:getApigApplicationQuotas", {
        "instanceId": args.instanceId,
        "name": args.name,
        "quotaId": args.quotaId,
        "region": args.region,
    }, opts);
}

/**
 * A collection of arguments for invoking getApigApplicationQuotas.
 */
export interface GetApigApplicationQuotasOutputArgs {
    instanceId: pulumi.Input<string>;
    name?: pulumi.Input<string>;
    quotaId?: pulumi.Input<string>;
    region?: pulumi.Input<string>;
}
