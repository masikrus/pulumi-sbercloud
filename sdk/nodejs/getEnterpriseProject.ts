// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

export function getEnterpriseProject(args?: GetEnterpriseProjectArgs, opts?: pulumi.InvokeOptions): Promise<GetEnterpriseProjectResult> {
    args = args || {};
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("sbercloud:index/getEnterpriseProject:getEnterpriseProject", {
        "id": args.id,
        "name": args.name,
        "status": args.status,
    }, opts);
}

/**
 * A collection of arguments for invoking getEnterpriseProject.
 */
export interface GetEnterpriseProjectArgs {
    id?: string;
    name?: string;
    status?: number;
}

/**
 * A collection of values returned by getEnterpriseProject.
 */
export interface GetEnterpriseProjectResult {
    readonly createdAt: string;
    readonly description: string;
    readonly id: string;
    readonly name: string;
    readonly status: number;
    readonly updatedAt: string;
}
export function getEnterpriseProjectOutput(args?: GetEnterpriseProjectOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetEnterpriseProjectResult> {
    args = args || {};
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("sbercloud:index/getEnterpriseProject:getEnterpriseProject", {
        "id": args.id,
        "name": args.name,
        "status": args.status,
    }, opts);
}

/**
 * A collection of arguments for invoking getEnterpriseProject.
 */
export interface GetEnterpriseProjectOutputArgs {
    id?: pulumi.Input<string>;
    name?: pulumi.Input<string>;
    status?: pulumi.Input<number>;
}
