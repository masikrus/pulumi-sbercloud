// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

export function getKpsKeypairs(args?: GetKpsKeypairsArgs, opts?: pulumi.InvokeOptions): Promise<GetKpsKeypairsResult> {
    args = args || {};
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("sbercloud:index/getKpsKeypairs:getKpsKeypairs", {
        "fingerprint": args.fingerprint,
        "isManaged": args.isManaged,
        "name": args.name,
        "publicKey": args.publicKey,
        "region": args.region,
    }, opts);
}

/**
 * A collection of arguments for invoking getKpsKeypairs.
 */
export interface GetKpsKeypairsArgs {
    fingerprint?: string;
    isManaged?: boolean;
    name?: string;
    publicKey?: string;
    region?: string;
}

/**
 * A collection of values returned by getKpsKeypairs.
 */
export interface GetKpsKeypairsResult {
    readonly fingerprint?: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly isManaged?: boolean;
    readonly keypairs: outputs.GetKpsKeypairsKeypair[];
    readonly name?: string;
    readonly publicKey?: string;
    readonly region?: string;
}
export function getKpsKeypairsOutput(args?: GetKpsKeypairsOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetKpsKeypairsResult> {
    args = args || {};
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("sbercloud:index/getKpsKeypairs:getKpsKeypairs", {
        "fingerprint": args.fingerprint,
        "isManaged": args.isManaged,
        "name": args.name,
        "publicKey": args.publicKey,
        "region": args.region,
    }, opts);
}

/**
 * A collection of arguments for invoking getKpsKeypairs.
 */
export interface GetKpsKeypairsOutputArgs {
    fingerprint?: pulumi.Input<string>;
    isManaged?: pulumi.Input<boolean>;
    name?: pulumi.Input<string>;
    publicKey?: pulumi.Input<string>;
    region?: pulumi.Input<string>;
}
