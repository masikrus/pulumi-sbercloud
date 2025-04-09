// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

export function getDmsKafkaMessages(args: GetDmsKafkaMessagesArgs, opts?: pulumi.InvokeOptions): Promise<GetDmsKafkaMessagesResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("sbercloud:index/getDmsKafkaMessages:getDmsKafkaMessages", {
        "download": args.download,
        "endTime": args.endTime,
        "instanceId": args.instanceId,
        "keyword": args.keyword,
        "messageOffset": args.messageOffset,
        "partition": args.partition,
        "region": args.region,
        "startTime": args.startTime,
        "topic": args.topic,
    }, opts);
}

/**
 * A collection of arguments for invoking getDmsKafkaMessages.
 */
export interface GetDmsKafkaMessagesArgs {
    download?: boolean;
    endTime?: string;
    instanceId: string;
    keyword?: string;
    messageOffset?: string;
    partition?: string;
    region?: string;
    startTime?: string;
    topic: string;
}

/**
 * A collection of values returned by getDmsKafkaMessages.
 */
export interface GetDmsKafkaMessagesResult {
    readonly download?: boolean;
    readonly endTime?: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly instanceId: string;
    readonly keyword?: string;
    readonly messageOffset?: string;
    readonly messages: outputs.GetDmsKafkaMessagesMessage[];
    readonly partition?: string;
    readonly region: string;
    readonly startTime?: string;
    readonly topic: string;
}
export function getDmsKafkaMessagesOutput(args: GetDmsKafkaMessagesOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetDmsKafkaMessagesResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("sbercloud:index/getDmsKafkaMessages:getDmsKafkaMessages", {
        "download": args.download,
        "endTime": args.endTime,
        "instanceId": args.instanceId,
        "keyword": args.keyword,
        "messageOffset": args.messageOffset,
        "partition": args.partition,
        "region": args.region,
        "startTime": args.startTime,
        "topic": args.topic,
    }, opts);
}

/**
 * A collection of arguments for invoking getDmsKafkaMessages.
 */
export interface GetDmsKafkaMessagesOutputArgs {
    download?: pulumi.Input<boolean>;
    endTime?: pulumi.Input<string>;
    instanceId: pulumi.Input<string>;
    keyword?: pulumi.Input<string>;
    messageOffset?: pulumi.Input<string>;
    partition?: pulumi.Input<string>;
    region?: pulumi.Input<string>;
    startTime?: pulumi.Input<string>;
    topic: pulumi.Input<string>;
}
