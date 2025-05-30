// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

export function getApigApi(args: GetApigApiArgs, opts?: pulumi.InvokeOptions): Promise<GetApigApiResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("sbercloud:index/getApigApi:getApigApi", {
        "apiId": args.apiId,
        "instanceId": args.instanceId,
        "region": args.region,
    }, opts);
}

/**
 * A collection of arguments for invoking getApigApi.
 */
export interface GetApigApiArgs {
    apiId: string;
    instanceId: string;
    region?: string;
}

/**
 * A collection of values returned by getApigApi.
 */
export interface GetApigApiResult {
    readonly apiId: string;
    readonly authorizerId: string;
    readonly backendParams: outputs.GetApigApiBackendParam[];
    readonly backendType: string;
    readonly bodyDescription: string;
    readonly cors: boolean;
    readonly description: string;
    readonly envId: string;
    readonly envName: string;
    readonly failureResponse: string;
    readonly funcGraphPolicies: outputs.GetApigApiFuncGraphPolicy[];
    readonly funcGraphs: outputs.GetApigApiFuncGraph[];
    readonly groupId: string;
    readonly groupName: string;
    readonly groupVersion: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly instanceId: string;
    readonly matching: string;
    readonly mockPolicies: outputs.GetApigApiMockPolicy[];
    readonly mocks: outputs.GetApigApiMock[];
    readonly name: string;
    readonly publishId: string;
    readonly publishedAt: string;
    readonly region: string;
    readonly registeredAt: string;
    readonly requestMethod: string;
    readonly requestParams: outputs.GetApigApiRequestParam[];
    readonly requestPath: string;
    readonly requestProtocol: string;
    readonly responseId: string;
    readonly securityAuthentication: string;
    readonly simpleAuthentication: boolean;
    readonly successResponse: string;
    readonly tags: string[];
    readonly type: string;
    readonly updatedAt: string;
    readonly webPolicies: outputs.GetApigApiWebPolicy[];
    readonly webs: outputs.GetApigApiWeb[];
}
export function getApigApiOutput(args: GetApigApiOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetApigApiResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("sbercloud:index/getApigApi:getApigApi", {
        "apiId": args.apiId,
        "instanceId": args.instanceId,
        "region": args.region,
    }, opts);
}

/**
 * A collection of arguments for invoking getApigApi.
 */
export interface GetApigApiOutputArgs {
    apiId: pulumi.Input<string>;
    instanceId: pulumi.Input<string>;
    region?: pulumi.Input<string>;
}
