// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

export function getIdentityRole(args?: GetIdentityRoleArgs, opts?: pulumi.InvokeOptions): Promise<GetIdentityRoleResult> {
    args = args || {};
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("sbercloud:index/getIdentityRole:getIdentityRole", {
        "displayName": args.displayName,
        "name": args.name,
    }, opts);
}

/**
 * A collection of arguments for invoking getIdentityRole.
 */
export interface GetIdentityRoleArgs {
    displayName?: string;
    name?: string;
}

/**
 * A collection of values returned by getIdentityRole.
 */
export interface GetIdentityRoleResult {
    readonly catalog: string;
    readonly description: string;
    readonly displayName: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly name: string;
    readonly policy: string;
    readonly type: string;
}
export function getIdentityRoleOutput(args?: GetIdentityRoleOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetIdentityRoleResult> {
    args = args || {};
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("sbercloud:index/getIdentityRole:getIdentityRole", {
        "displayName": args.displayName,
        "name": args.name,
    }, opts);
}

/**
 * A collection of arguments for invoking getIdentityRole.
 */
export interface GetIdentityRoleOutputArgs {
    displayName?: pulumi.Input<string>;
    name?: pulumi.Input<string>;
}
