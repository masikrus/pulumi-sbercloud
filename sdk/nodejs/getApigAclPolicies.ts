// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

export function getApigAclPolicies(args: GetApigAclPoliciesArgs, opts?: pulumi.InvokeOptions): Promise<GetApigAclPoliciesResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("sbercloud:index/getApigAclPolicies:getApigAclPolicies", {
        "entityType": args.entityType,
        "instanceId": args.instanceId,
        "name": args.name,
        "policyId": args.policyId,
        "region": args.region,
        "type": args.type,
    }, opts);
}

/**
 * A collection of arguments for invoking getApigAclPolicies.
 */
export interface GetApigAclPoliciesArgs {
    entityType?: string;
    instanceId: string;
    name?: string;
    policyId?: string;
    region?: string;
    type?: string;
}

/**
 * A collection of values returned by getApigAclPolicies.
 */
export interface GetApigAclPoliciesResult {
    readonly entityType?: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly instanceId: string;
    readonly name?: string;
    readonly policies: outputs.GetApigAclPoliciesPolicy[];
    readonly policyId?: string;
    readonly region: string;
    readonly type?: string;
}
export function getApigAclPoliciesOutput(args: GetApigAclPoliciesOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetApigAclPoliciesResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("sbercloud:index/getApigAclPolicies:getApigAclPolicies", {
        "entityType": args.entityType,
        "instanceId": args.instanceId,
        "name": args.name,
        "policyId": args.policyId,
        "region": args.region,
        "type": args.type,
    }, opts);
}

/**
 * A collection of arguments for invoking getApigAclPolicies.
 */
export interface GetApigAclPoliciesOutputArgs {
    entityType?: pulumi.Input<string>;
    instanceId: pulumi.Input<string>;
    name?: pulumi.Input<string>;
    policyId?: pulumi.Input<string>;
    region?: pulumi.Input<string>;
    type?: pulumi.Input<string>;
}
