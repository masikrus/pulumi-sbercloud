// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

export function getIdentityRoleV3(args?: GetIdentityRoleV3Args, opts?: pulumi.InvokeOptions): Promise<GetIdentityRoleV3Result> {
    args = args || {};
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("sbercloud:index/getIdentityRoleV3:getIdentityRoleV3", {
        "displayName": args.displayName,
        "name": args.name,
    }, opts);
}

/**
 * A collection of arguments for invoking getIdentityRoleV3.
 */
export interface GetIdentityRoleV3Args {
    displayName?: string;
    name?: string;
}

/**
 * A collection of values returned by getIdentityRoleV3.
 */
export interface GetIdentityRoleV3Result {
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
export function getIdentityRoleV3Output(args?: GetIdentityRoleV3OutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetIdentityRoleV3Result> {
    args = args || {};
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("sbercloud:index/getIdentityRoleV3:getIdentityRoleV3", {
        "displayName": args.displayName,
        "name": args.name,
    }, opts);
}

/**
 * A collection of arguments for invoking getIdentityRoleV3.
 */
export interface GetIdentityRoleV3OutputArgs {
    displayName?: pulumi.Input<string>;
    name?: pulumi.Input<string>;
}
