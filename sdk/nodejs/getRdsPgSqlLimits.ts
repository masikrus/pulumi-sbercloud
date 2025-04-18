// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

export function getRdsPgSqlLimits(args: GetRdsPgSqlLimitsArgs, opts?: pulumi.InvokeOptions): Promise<GetRdsPgSqlLimitsResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("sbercloud:index/getRdsPgSqlLimits:getRdsPgSqlLimits", {
        "dbName": args.dbName,
        "instanceId": args.instanceId,
        "isEffective": args.isEffective,
        "maxConcurrency": args.maxConcurrency,
        "maxWaiting": args.maxWaiting,
        "queryId": args.queryId,
        "queryString": args.queryString,
        "region": args.region,
        "searchPath": args.searchPath,
        "sqlLimitId": args.sqlLimitId,
    }, opts);
}

/**
 * A collection of arguments for invoking getRdsPgSqlLimits.
 */
export interface GetRdsPgSqlLimitsArgs {
    dbName: string;
    instanceId: string;
    isEffective?: string;
    maxConcurrency?: string;
    maxWaiting?: string;
    queryId?: string;
    queryString?: string;
    region?: string;
    searchPath?: string;
    sqlLimitId?: string;
}

/**
 * A collection of values returned by getRdsPgSqlLimits.
 */
export interface GetRdsPgSqlLimitsResult {
    readonly dbName: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly instanceId: string;
    readonly isEffective?: string;
    readonly maxConcurrency?: string;
    readonly maxWaiting?: string;
    readonly queryId?: string;
    readonly queryString?: string;
    readonly region: string;
    readonly searchPath?: string;
    readonly sqlLimitId?: string;
    readonly sqlLimits: outputs.GetRdsPgSqlLimitsSqlLimit[];
}
export function getRdsPgSqlLimitsOutput(args: GetRdsPgSqlLimitsOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetRdsPgSqlLimitsResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("sbercloud:index/getRdsPgSqlLimits:getRdsPgSqlLimits", {
        "dbName": args.dbName,
        "instanceId": args.instanceId,
        "isEffective": args.isEffective,
        "maxConcurrency": args.maxConcurrency,
        "maxWaiting": args.maxWaiting,
        "queryId": args.queryId,
        "queryString": args.queryString,
        "region": args.region,
        "searchPath": args.searchPath,
        "sqlLimitId": args.sqlLimitId,
    }, opts);
}

/**
 * A collection of arguments for invoking getRdsPgSqlLimits.
 */
export interface GetRdsPgSqlLimitsOutputArgs {
    dbName: pulumi.Input<string>;
    instanceId: pulumi.Input<string>;
    isEffective?: pulumi.Input<string>;
    maxConcurrency?: pulumi.Input<string>;
    maxWaiting?: pulumi.Input<string>;
    queryId?: pulumi.Input<string>;
    queryString?: pulumi.Input<string>;
    region?: pulumi.Input<string>;
    searchPath?: pulumi.Input<string>;
    sqlLimitId?: pulumi.Input<string>;
}
