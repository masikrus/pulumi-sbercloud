// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

export function getImagesImage(args?: GetImagesImageArgs, opts?: pulumi.InvokeOptions): Promise<GetImagesImageResult> {
    args = args || {};
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("sbercloud:index/getImagesImage:getImagesImage", {
        "architecture": args.architecture,
        "enterpriseProjectId": args.enterpriseProjectId,
        "flavorId": args.flavorId,
        "imageId": args.imageId,
        "imageType": args.imageType,
        "isWholeImage": args.isWholeImage,
        "mostRecent": args.mostRecent,
        "name": args.name,
        "nameRegex": args.nameRegex,
        "os": args.os,
        "osVersion": args.osVersion,
        "owner": args.owner,
        "region": args.region,
        "sizeMax": args.sizeMax,
        "sizeMin": args.sizeMin,
        "sortDirection": args.sortDirection,
        "sortKey": args.sortKey,
        "tag": args.tag,
        "visibility": args.visibility,
    }, opts);
}

/**
 * A collection of arguments for invoking getImagesImage.
 */
export interface GetImagesImageArgs {
    architecture?: string;
    enterpriseProjectId?: string;
    flavorId?: string;
    imageId?: string;
    imageType?: string;
    isWholeImage?: boolean;
    mostRecent?: boolean;
    name?: string;
    nameRegex?: string;
    os?: string;
    osVersion?: string;
    owner?: string;
    region?: string;
    /**
     * @deprecated size_max is deprecated
     */
    sizeMax?: number;
    /**
     * @deprecated size_min is deprecated
     */
    sizeMin?: number;
    sortDirection?: string;
    sortKey?: string;
    tag?: string;
    visibility?: string;
}

/**
 * A collection of values returned by getImagesImage.
 */
export interface GetImagesImageResult {
    readonly activeAt: string;
    readonly architecture: string;
    readonly backupId: string;
    readonly checksum: string;
    readonly containerFormat: string;
    readonly createdAt: string;
    readonly dataOrigin: string;
    readonly description: string;
    readonly diskFormat: string;
    readonly enterpriseProjectId: string;
    readonly file: string;
    readonly flavorId?: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly imageId: string;
    readonly imageType: string;
    readonly isWholeImage?: boolean;
    readonly maxRamMb: number;
    readonly metadata: {[key: string]: string};
    readonly minDiskGb: number;
    readonly minRamMb: number;
    readonly mostRecent?: boolean;
    readonly name: string;
    readonly nameRegex?: string;
    readonly os: string;
    readonly osVersion: string;
    readonly owner: string;
    readonly protected: boolean;
    readonly region: string;
    readonly schema: string;
    readonly sizeBytes: number;
    /**
     * @deprecated size_max is deprecated
     */
    readonly sizeMax?: number;
    /**
     * @deprecated size_min is deprecated
     */
    readonly sizeMin?: number;
    readonly sortDirection?: string;
    readonly sortKey?: string;
    readonly status: string;
    readonly tag?: string;
    readonly updatedAt: string;
    readonly visibility: string;
}
export function getImagesImageOutput(args?: GetImagesImageOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetImagesImageResult> {
    args = args || {};
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("sbercloud:index/getImagesImage:getImagesImage", {
        "architecture": args.architecture,
        "enterpriseProjectId": args.enterpriseProjectId,
        "flavorId": args.flavorId,
        "imageId": args.imageId,
        "imageType": args.imageType,
        "isWholeImage": args.isWholeImage,
        "mostRecent": args.mostRecent,
        "name": args.name,
        "nameRegex": args.nameRegex,
        "os": args.os,
        "osVersion": args.osVersion,
        "owner": args.owner,
        "region": args.region,
        "sizeMax": args.sizeMax,
        "sizeMin": args.sizeMin,
        "sortDirection": args.sortDirection,
        "sortKey": args.sortKey,
        "tag": args.tag,
        "visibility": args.visibility,
    }, opts);
}

/**
 * A collection of arguments for invoking getImagesImage.
 */
export interface GetImagesImageOutputArgs {
    architecture?: pulumi.Input<string>;
    enterpriseProjectId?: pulumi.Input<string>;
    flavorId?: pulumi.Input<string>;
    imageId?: pulumi.Input<string>;
    imageType?: pulumi.Input<string>;
    isWholeImage?: pulumi.Input<boolean>;
    mostRecent?: pulumi.Input<boolean>;
    name?: pulumi.Input<string>;
    nameRegex?: pulumi.Input<string>;
    os?: pulumi.Input<string>;
    osVersion?: pulumi.Input<string>;
    owner?: pulumi.Input<string>;
    region?: pulumi.Input<string>;
    /**
     * @deprecated size_max is deprecated
     */
    sizeMax?: pulumi.Input<number>;
    /**
     * @deprecated size_min is deprecated
     */
    sizeMin?: pulumi.Input<number>;
    sortDirection?: pulumi.Input<string>;
    sortKey?: pulumi.Input<string>;
    tag?: pulumi.Input<string>;
    visibility?: pulumi.Input<string>;
}
