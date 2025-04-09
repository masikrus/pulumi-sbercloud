// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

export function getDmsMaintainwindow(args?: GetDmsMaintainwindowArgs, opts?: pulumi.InvokeOptions): Promise<GetDmsMaintainwindowResult> {
    args = args || {};
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("sbercloud:index/getDmsMaintainwindow:getDmsMaintainwindow", {
        "begin": args.begin,
        "default": args.default,
        "end": args.end,
        "region": args.region,
        "seq": args.seq,
    }, opts);
}

/**
 * A collection of arguments for invoking getDmsMaintainwindow.
 */
export interface GetDmsMaintainwindowArgs {
    begin?: string;
    default?: boolean;
    end?: string;
    region?: string;
    seq?: number;
}

/**
 * A collection of values returned by getDmsMaintainwindow.
 */
export interface GetDmsMaintainwindowResult {
    readonly begin: string;
    readonly default: boolean;
    readonly end: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly region: string;
    readonly seq: number;
}
export function getDmsMaintainwindowOutput(args?: GetDmsMaintainwindowOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetDmsMaintainwindowResult> {
    args = args || {};
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("sbercloud:index/getDmsMaintainwindow:getDmsMaintainwindow", {
        "begin": args.begin,
        "default": args.default,
        "end": args.end,
        "region": args.region,
        "seq": args.seq,
    }, opts);
}

/**
 * A collection of arguments for invoking getDmsMaintainwindow.
 */
export interface GetDmsMaintainwindowOutputArgs {
    begin?: pulumi.Input<string>;
    default?: pulumi.Input<boolean>;
    end?: pulumi.Input<string>;
    region?: pulumi.Input<string>;
    seq?: pulumi.Input<number>;
}
