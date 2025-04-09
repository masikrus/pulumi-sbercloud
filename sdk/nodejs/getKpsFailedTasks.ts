// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

export function getKpsFailedTasks(args?: GetKpsFailedTasksArgs, opts?: pulumi.InvokeOptions): Promise<GetKpsFailedTasksResult> {
    args = args || {};
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("sbercloud:index/getKpsFailedTasks:getKpsFailedTasks", {
        "region": args.region,
    }, opts);
}

/**
 * A collection of arguments for invoking getKpsFailedTasks.
 */
export interface GetKpsFailedTasksArgs {
    region?: string;
}

/**
 * A collection of values returned by getKpsFailedTasks.
 */
export interface GetKpsFailedTasksResult {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly region: string;
    readonly tasks: outputs.GetKpsFailedTasksTask[];
}
export function getKpsFailedTasksOutput(args?: GetKpsFailedTasksOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetKpsFailedTasksResult> {
    args = args || {};
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("sbercloud:index/getKpsFailedTasks:getKpsFailedTasks", {
        "region": args.region,
    }, opts);
}

/**
 * A collection of arguments for invoking getKpsFailedTasks.
 */
export interface GetKpsFailedTasksOutputArgs {
    region?: pulumi.Input<string>;
}
