// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

export function getDcsAccounts(args: GetDcsAccountsArgs, opts?: pulumi.InvokeOptions): Promise<GetDcsAccountsResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("sbercloud:index/getDcsAccounts:getDcsAccounts", {
        "accountName": args.accountName,
        "accountRole": args.accountRole,
        "accountType": args.accountType,
        "description": args.description,
        "instanceId": args.instanceId,
        "region": args.region,
        "status": args.status,
    }, opts);
}

/**
 * A collection of arguments for invoking getDcsAccounts.
 */
export interface GetDcsAccountsArgs {
    accountName?: string;
    accountRole?: string;
    accountType?: string;
    description?: string;
    instanceId: string;
    region?: string;
    status?: string;
}

/**
 * A collection of values returned by getDcsAccounts.
 */
export interface GetDcsAccountsResult {
    readonly accountName?: string;
    readonly accountRole?: string;
    readonly accountType?: string;
    readonly accounts: outputs.GetDcsAccountsAccount[];
    readonly description?: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly instanceId: string;
    readonly region: string;
    readonly status?: string;
}
export function getDcsAccountsOutput(args: GetDcsAccountsOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetDcsAccountsResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("sbercloud:index/getDcsAccounts:getDcsAccounts", {
        "accountName": args.accountName,
        "accountRole": args.accountRole,
        "accountType": args.accountType,
        "description": args.description,
        "instanceId": args.instanceId,
        "region": args.region,
        "status": args.status,
    }, opts);
}

/**
 * A collection of arguments for invoking getDcsAccounts.
 */
export interface GetDcsAccountsOutputArgs {
    accountName?: pulumi.Input<string>;
    accountRole?: pulumi.Input<string>;
    accountType?: pulumi.Input<string>;
    description?: pulumi.Input<string>;
    instanceId: pulumi.Input<string>;
    region?: pulumi.Input<string>;
    status?: pulumi.Input<string>;
}
