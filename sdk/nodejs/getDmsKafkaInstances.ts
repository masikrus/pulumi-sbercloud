// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

export function getDmsKafkaInstances(args?: GetDmsKafkaInstancesArgs, opts?: pulumi.InvokeOptions): Promise<GetDmsKafkaInstancesResult> {
    args = args || {};
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("sbercloud:index/getDmsKafkaInstances:getDmsKafkaInstances", {
        "enterpriseProjectId": args.enterpriseProjectId,
        "fuzzyMatch": args.fuzzyMatch,
        "includeFailure": args.includeFailure,
        "instanceId": args.instanceId,
        "name": args.name,
        "region": args.region,
        "status": args.status,
    }, opts);
}

/**
 * A collection of arguments for invoking getDmsKafkaInstances.
 */
export interface GetDmsKafkaInstancesArgs {
    enterpriseProjectId?: string;
    fuzzyMatch?: boolean;
    includeFailure?: boolean;
    instanceId?: string;
    name?: string;
    region?: string;
    status?: string;
}

/**
 * A collection of values returned by getDmsKafkaInstances.
 */
export interface GetDmsKafkaInstancesResult {
    readonly enterpriseProjectId?: string;
    readonly fuzzyMatch?: boolean;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly includeFailure?: boolean;
    readonly instanceId?: string;
    readonly instances: outputs.GetDmsKafkaInstancesInstance[];
    readonly name?: string;
    readonly region?: string;
    readonly status?: string;
}
export function getDmsKafkaInstancesOutput(args?: GetDmsKafkaInstancesOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetDmsKafkaInstancesResult> {
    args = args || {};
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("sbercloud:index/getDmsKafkaInstances:getDmsKafkaInstances", {
        "enterpriseProjectId": args.enterpriseProjectId,
        "fuzzyMatch": args.fuzzyMatch,
        "includeFailure": args.includeFailure,
        "instanceId": args.instanceId,
        "name": args.name,
        "region": args.region,
        "status": args.status,
    }, opts);
}

/**
 * A collection of arguments for invoking getDmsKafkaInstances.
 */
export interface GetDmsKafkaInstancesOutputArgs {
    enterpriseProjectId?: pulumi.Input<string>;
    fuzzyMatch?: pulumi.Input<boolean>;
    includeFailure?: pulumi.Input<boolean>;
    instanceId?: pulumi.Input<string>;
    name?: pulumi.Input<string>;
    region?: pulumi.Input<string>;
    status?: pulumi.Input<string>;
}
