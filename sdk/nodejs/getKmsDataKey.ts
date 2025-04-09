// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

export function getKmsDataKey(args: GetKmsDataKeyArgs, opts?: pulumi.InvokeOptions): Promise<GetKmsDataKeyResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("sbercloud:index/getKmsDataKey:getKmsDataKey", {
        "datakeyLength": args.datakeyLength,
        "encryptionContext": args.encryptionContext,
        "keyId": args.keyId,
        "region": args.region,
    }, opts);
}

/**
 * A collection of arguments for invoking getKmsDataKey.
 */
export interface GetKmsDataKeyArgs {
    datakeyLength: string;
    encryptionContext?: string;
    keyId: string;
    region?: string;
}

/**
 * A collection of values returned by getKmsDataKey.
 */
export interface GetKmsDataKeyResult {
    readonly cipherText: string;
    readonly datakeyLength: string;
    readonly encryptionContext?: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly keyId: string;
    readonly plainText: string;
    readonly region: string;
}
export function getKmsDataKeyOutput(args: GetKmsDataKeyOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetKmsDataKeyResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("sbercloud:index/getKmsDataKey:getKmsDataKey", {
        "datakeyLength": args.datakeyLength,
        "encryptionContext": args.encryptionContext,
        "keyId": args.keyId,
        "region": args.region,
    }, opts);
}

/**
 * A collection of arguments for invoking getKmsDataKey.
 */
export interface GetKmsDataKeyOutputArgs {
    datakeyLength: pulumi.Input<string>;
    encryptionContext?: pulumi.Input<string>;
    keyId: pulumi.Input<string>;
    region?: pulumi.Input<string>;
}
