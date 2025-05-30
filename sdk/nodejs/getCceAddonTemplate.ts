// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

export function getCceAddonTemplate(args: GetCceAddonTemplateArgs, opts?: pulumi.InvokeOptions): Promise<GetCceAddonTemplateResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("sbercloud:index/getCceAddonTemplate:getCceAddonTemplate", {
        "clusterId": args.clusterId,
        "name": args.name,
        "region": args.region,
        "version": args.version,
    }, opts);
}

/**
 * A collection of arguments for invoking getCceAddonTemplate.
 */
export interface GetCceAddonTemplateArgs {
    clusterId: string;
    name: string;
    region?: string;
    version: string;
}

/**
 * A collection of values returned by getCceAddonTemplate.
 */
export interface GetCceAddonTemplateResult {
    readonly clusterId: string;
    readonly description: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly name: string;
    readonly region: string;
    readonly spec: string;
    readonly stable: boolean;
    readonly supportVersions: outputs.GetCceAddonTemplateSupportVersion[];
    readonly version: string;
}
export function getCceAddonTemplateOutput(args: GetCceAddonTemplateOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetCceAddonTemplateResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("sbercloud:index/getCceAddonTemplate:getCceAddonTemplate", {
        "clusterId": args.clusterId,
        "name": args.name,
        "region": args.region,
        "version": args.version,
    }, opts);
}

/**
 * A collection of arguments for invoking getCceAddonTemplate.
 */
export interface GetCceAddonTemplateOutputArgs {
    clusterId: pulumi.Input<string>;
    name: pulumi.Input<string>;
    region?: pulumi.Input<string>;
    version: pulumi.Input<string>;
}
