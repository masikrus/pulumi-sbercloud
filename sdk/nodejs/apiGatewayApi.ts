// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

export class ApiGatewayApi extends pulumi.CustomResource {
    /**
     * Get an existing ApiGatewayApi resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ApiGatewayApiState, opts?: pulumi.CustomResourceOptions): ApiGatewayApi {
        return new ApiGatewayApi(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'sbercloud:index/apiGatewayApi:ApiGatewayApi';

    /**
     * Returns true if the given object is an instance of ApiGatewayApi.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ApiGatewayApi {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ApiGatewayApi.__pulumiType;
    }

    /**
     * The ID of the authorizer to which the API request used.
     */
    public readonly authorizerId!: pulumi.Output<string | undefined>;
    /**
     * The configurations of the backend parameters.
     */
    public readonly backendParams!: pulumi.Output<outputs.ApiGatewayApiBackendParam[] | undefined>;
    /**
     * The description of the API request body, which can be an example request body, media type or parameters.
     */
    public readonly bodyDescription!: pulumi.Output<string | undefined>;
    /**
     * The content type of the request body.
     */
    public readonly contentType!: pulumi.Output<string>;
    /**
     * Whether CORS is supported.
     */
    public readonly cors!: pulumi.Output<boolean | undefined>;
    /**
     * The API description.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * The example response for a failure request.
     */
    public readonly failureResponse!: pulumi.Output<string | undefined>;
    /**
     * The FunctionGraph backend details.
     */
    public readonly funcGraph!: pulumi.Output<outputs.ApiGatewayApiFuncGraph>;
    /**
     * The policy backends of the FunctionGraph function.
     */
    public readonly funcGraphPolicies!: pulumi.Output<outputs.ApiGatewayApiFuncGraphPolicy[] | undefined>;
    /**
     * The ID of the API group to which the API belongs.
     */
    public readonly groupId!: pulumi.Output<string>;
    /**
     * The ID of the instance to which the API belongs.
     */
    public readonly instanceId!: pulumi.Output<string>;
    /**
     * Whether to perform Base64 encoding on the body for interaction with FunctionGraph.
     */
    public readonly isSendFgBodyBase64!: pulumi.Output<boolean | undefined>;
    /**
     * The matching mode of the API.
     */
    public readonly matching!: pulumi.Output<string | undefined>;
    /**
     * The mock backend details.
     */
    public readonly mock!: pulumi.Output<outputs.ApiGatewayApiMock>;
    /**
     * The mock policy backends.
     */
    public readonly mockPolicies!: pulumi.Output<outputs.ApiGatewayApiMockPolicy[] | undefined>;
    /**
     * The API name.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The region where the API is located.
     */
    public readonly region!: pulumi.Output<string>;
    /**
     * The registered time of the API.
     */
    public /*out*/ readonly registeredAt!: pulumi.Output<string>;
    /**
     * The request method of the API.
     */
    public readonly requestMethod!: pulumi.Output<string>;
    /**
     * The configurations of the front-end parameters.
     */
    public readonly requestParams!: pulumi.Output<outputs.ApiGatewayApiRequestParam[]>;
    /**
     * The request address.
     */
    public readonly requestPath!: pulumi.Output<string>;
    /**
     * The request protocol of the API request.
     */
    public readonly requestProtocol!: pulumi.Output<string>;
    /**
     * The ID of the custom response that API used.
     */
    public readonly responseId!: pulumi.Output<string | undefined>;
    /**
     * The security authentication mode of the API request.
     */
    public readonly securityAuthentication!: pulumi.Output<string | undefined>;
    /**
     * Whether the authentication of the application code is enabled.
     */
    public readonly simpleAuthentication!: pulumi.Output<boolean>;
    /**
     * The example response for a successful request.
     */
    public readonly successResponse!: pulumi.Output<string | undefined>;
    /**
     * The list of tags configuration.
     */
    public readonly tags!: pulumi.Output<string[] | undefined>;
    /**
     * The API type.
     */
    public readonly type!: pulumi.Output<string>;
    /**
     * The latest update time of the API.
     */
    public /*out*/ readonly updatedAt!: pulumi.Output<string>;
    /**
     * The web backend details.
     */
    public readonly web!: pulumi.Output<outputs.ApiGatewayApiWeb>;
    /**
     * The web policy backends.
     */
    public readonly webPolicies!: pulumi.Output<outputs.ApiGatewayApiWebPolicy[] | undefined>;

    /**
     * Create a ApiGatewayApi resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ApiGatewayApiArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ApiGatewayApiArgs | ApiGatewayApiState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ApiGatewayApiState | undefined;
            resourceInputs["authorizerId"] = state ? state.authorizerId : undefined;
            resourceInputs["backendParams"] = state ? state.backendParams : undefined;
            resourceInputs["bodyDescription"] = state ? state.bodyDescription : undefined;
            resourceInputs["contentType"] = state ? state.contentType : undefined;
            resourceInputs["cors"] = state ? state.cors : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["failureResponse"] = state ? state.failureResponse : undefined;
            resourceInputs["funcGraph"] = state ? state.funcGraph : undefined;
            resourceInputs["funcGraphPolicies"] = state ? state.funcGraphPolicies : undefined;
            resourceInputs["groupId"] = state ? state.groupId : undefined;
            resourceInputs["instanceId"] = state ? state.instanceId : undefined;
            resourceInputs["isSendFgBodyBase64"] = state ? state.isSendFgBodyBase64 : undefined;
            resourceInputs["matching"] = state ? state.matching : undefined;
            resourceInputs["mock"] = state ? state.mock : undefined;
            resourceInputs["mockPolicies"] = state ? state.mockPolicies : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["region"] = state ? state.region : undefined;
            resourceInputs["registeredAt"] = state ? state.registeredAt : undefined;
            resourceInputs["requestMethod"] = state ? state.requestMethod : undefined;
            resourceInputs["requestParams"] = state ? state.requestParams : undefined;
            resourceInputs["requestPath"] = state ? state.requestPath : undefined;
            resourceInputs["requestProtocol"] = state ? state.requestProtocol : undefined;
            resourceInputs["responseId"] = state ? state.responseId : undefined;
            resourceInputs["securityAuthentication"] = state ? state.securityAuthentication : undefined;
            resourceInputs["simpleAuthentication"] = state ? state.simpleAuthentication : undefined;
            resourceInputs["successResponse"] = state ? state.successResponse : undefined;
            resourceInputs["tags"] = state ? state.tags : undefined;
            resourceInputs["type"] = state ? state.type : undefined;
            resourceInputs["updatedAt"] = state ? state.updatedAt : undefined;
            resourceInputs["web"] = state ? state.web : undefined;
            resourceInputs["webPolicies"] = state ? state.webPolicies : undefined;
        } else {
            const args = argsOrState as ApiGatewayApiArgs | undefined;
            if ((!args || args.groupId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'groupId'");
            }
            if ((!args || args.instanceId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'instanceId'");
            }
            if ((!args || args.requestMethod === undefined) && !opts.urn) {
                throw new Error("Missing required property 'requestMethod'");
            }
            if ((!args || args.requestPath === undefined) && !opts.urn) {
                throw new Error("Missing required property 'requestPath'");
            }
            if ((!args || args.requestProtocol === undefined) && !opts.urn) {
                throw new Error("Missing required property 'requestProtocol'");
            }
            if ((!args || args.type === undefined) && !opts.urn) {
                throw new Error("Missing required property 'type'");
            }
            resourceInputs["authorizerId"] = args ? args.authorizerId : undefined;
            resourceInputs["backendParams"] = args ? args.backendParams : undefined;
            resourceInputs["bodyDescription"] = args ? args.bodyDescription : undefined;
            resourceInputs["contentType"] = args ? args.contentType : undefined;
            resourceInputs["cors"] = args ? args.cors : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["failureResponse"] = args ? args.failureResponse : undefined;
            resourceInputs["funcGraph"] = args ? args.funcGraph : undefined;
            resourceInputs["funcGraphPolicies"] = args ? args.funcGraphPolicies : undefined;
            resourceInputs["groupId"] = args ? args.groupId : undefined;
            resourceInputs["instanceId"] = args ? args.instanceId : undefined;
            resourceInputs["isSendFgBodyBase64"] = args ? args.isSendFgBodyBase64 : undefined;
            resourceInputs["matching"] = args ? args.matching : undefined;
            resourceInputs["mock"] = args ? args.mock : undefined;
            resourceInputs["mockPolicies"] = args ? args.mockPolicies : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["region"] = args ? args.region : undefined;
            resourceInputs["requestMethod"] = args ? args.requestMethod : undefined;
            resourceInputs["requestParams"] = args ? args.requestParams : undefined;
            resourceInputs["requestPath"] = args ? args.requestPath : undefined;
            resourceInputs["requestProtocol"] = args ? args.requestProtocol : undefined;
            resourceInputs["responseId"] = args ? args.responseId : undefined;
            resourceInputs["securityAuthentication"] = args ? args.securityAuthentication : undefined;
            resourceInputs["simpleAuthentication"] = args ? args.simpleAuthentication : undefined;
            resourceInputs["successResponse"] = args ? args.successResponse : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["type"] = args ? args.type : undefined;
            resourceInputs["web"] = args ? args.web : undefined;
            resourceInputs["webPolicies"] = args ? args.webPolicies : undefined;
            resourceInputs["registeredAt"] = undefined /*out*/;
            resourceInputs["updatedAt"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(ApiGatewayApi.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ApiGatewayApi resources.
 */
export interface ApiGatewayApiState {
    /**
     * The ID of the authorizer to which the API request used.
     */
    authorizerId?: pulumi.Input<string>;
    /**
     * The configurations of the backend parameters.
     */
    backendParams?: pulumi.Input<pulumi.Input<inputs.ApiGatewayApiBackendParam>[]>;
    /**
     * The description of the API request body, which can be an example request body, media type or parameters.
     */
    bodyDescription?: pulumi.Input<string>;
    /**
     * The content type of the request body.
     */
    contentType?: pulumi.Input<string>;
    /**
     * Whether CORS is supported.
     */
    cors?: pulumi.Input<boolean>;
    /**
     * The API description.
     */
    description?: pulumi.Input<string>;
    /**
     * The example response for a failure request.
     */
    failureResponse?: pulumi.Input<string>;
    /**
     * The FunctionGraph backend details.
     */
    funcGraph?: pulumi.Input<inputs.ApiGatewayApiFuncGraph>;
    /**
     * The policy backends of the FunctionGraph function.
     */
    funcGraphPolicies?: pulumi.Input<pulumi.Input<inputs.ApiGatewayApiFuncGraphPolicy>[]>;
    /**
     * The ID of the API group to which the API belongs.
     */
    groupId?: pulumi.Input<string>;
    /**
     * The ID of the instance to which the API belongs.
     */
    instanceId?: pulumi.Input<string>;
    /**
     * Whether to perform Base64 encoding on the body for interaction with FunctionGraph.
     */
    isSendFgBodyBase64?: pulumi.Input<boolean>;
    /**
     * The matching mode of the API.
     */
    matching?: pulumi.Input<string>;
    /**
     * The mock backend details.
     */
    mock?: pulumi.Input<inputs.ApiGatewayApiMock>;
    /**
     * The mock policy backends.
     */
    mockPolicies?: pulumi.Input<pulumi.Input<inputs.ApiGatewayApiMockPolicy>[]>;
    /**
     * The API name.
     */
    name?: pulumi.Input<string>;
    /**
     * The region where the API is located.
     */
    region?: pulumi.Input<string>;
    /**
     * The registered time of the API.
     */
    registeredAt?: pulumi.Input<string>;
    /**
     * The request method of the API.
     */
    requestMethod?: pulumi.Input<string>;
    /**
     * The configurations of the front-end parameters.
     */
    requestParams?: pulumi.Input<pulumi.Input<inputs.ApiGatewayApiRequestParam>[]>;
    /**
     * The request address.
     */
    requestPath?: pulumi.Input<string>;
    /**
     * The request protocol of the API request.
     */
    requestProtocol?: pulumi.Input<string>;
    /**
     * The ID of the custom response that API used.
     */
    responseId?: pulumi.Input<string>;
    /**
     * The security authentication mode of the API request.
     */
    securityAuthentication?: pulumi.Input<string>;
    /**
     * Whether the authentication of the application code is enabled.
     */
    simpleAuthentication?: pulumi.Input<boolean>;
    /**
     * The example response for a successful request.
     */
    successResponse?: pulumi.Input<string>;
    /**
     * The list of tags configuration.
     */
    tags?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The API type.
     */
    type?: pulumi.Input<string>;
    /**
     * The latest update time of the API.
     */
    updatedAt?: pulumi.Input<string>;
    /**
     * The web backend details.
     */
    web?: pulumi.Input<inputs.ApiGatewayApiWeb>;
    /**
     * The web policy backends.
     */
    webPolicies?: pulumi.Input<pulumi.Input<inputs.ApiGatewayApiWebPolicy>[]>;
}

/**
 * The set of arguments for constructing a ApiGatewayApi resource.
 */
export interface ApiGatewayApiArgs {
    /**
     * The ID of the authorizer to which the API request used.
     */
    authorizerId?: pulumi.Input<string>;
    /**
     * The configurations of the backend parameters.
     */
    backendParams?: pulumi.Input<pulumi.Input<inputs.ApiGatewayApiBackendParam>[]>;
    /**
     * The description of the API request body, which can be an example request body, media type or parameters.
     */
    bodyDescription?: pulumi.Input<string>;
    /**
     * The content type of the request body.
     */
    contentType?: pulumi.Input<string>;
    /**
     * Whether CORS is supported.
     */
    cors?: pulumi.Input<boolean>;
    /**
     * The API description.
     */
    description?: pulumi.Input<string>;
    /**
     * The example response for a failure request.
     */
    failureResponse?: pulumi.Input<string>;
    /**
     * The FunctionGraph backend details.
     */
    funcGraph?: pulumi.Input<inputs.ApiGatewayApiFuncGraph>;
    /**
     * The policy backends of the FunctionGraph function.
     */
    funcGraphPolicies?: pulumi.Input<pulumi.Input<inputs.ApiGatewayApiFuncGraphPolicy>[]>;
    /**
     * The ID of the API group to which the API belongs.
     */
    groupId: pulumi.Input<string>;
    /**
     * The ID of the instance to which the API belongs.
     */
    instanceId: pulumi.Input<string>;
    /**
     * Whether to perform Base64 encoding on the body for interaction with FunctionGraph.
     */
    isSendFgBodyBase64?: pulumi.Input<boolean>;
    /**
     * The matching mode of the API.
     */
    matching?: pulumi.Input<string>;
    /**
     * The mock backend details.
     */
    mock?: pulumi.Input<inputs.ApiGatewayApiMock>;
    /**
     * The mock policy backends.
     */
    mockPolicies?: pulumi.Input<pulumi.Input<inputs.ApiGatewayApiMockPolicy>[]>;
    /**
     * The API name.
     */
    name?: pulumi.Input<string>;
    /**
     * The region where the API is located.
     */
    region?: pulumi.Input<string>;
    /**
     * The request method of the API.
     */
    requestMethod: pulumi.Input<string>;
    /**
     * The configurations of the front-end parameters.
     */
    requestParams?: pulumi.Input<pulumi.Input<inputs.ApiGatewayApiRequestParam>[]>;
    /**
     * The request address.
     */
    requestPath: pulumi.Input<string>;
    /**
     * The request protocol of the API request.
     */
    requestProtocol: pulumi.Input<string>;
    /**
     * The ID of the custom response that API used.
     */
    responseId?: pulumi.Input<string>;
    /**
     * The security authentication mode of the API request.
     */
    securityAuthentication?: pulumi.Input<string>;
    /**
     * Whether the authentication of the application code is enabled.
     */
    simpleAuthentication?: pulumi.Input<boolean>;
    /**
     * The example response for a successful request.
     */
    successResponse?: pulumi.Input<string>;
    /**
     * The list of tags configuration.
     */
    tags?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The API type.
     */
    type: pulumi.Input<string>;
    /**
     * The web backend details.
     */
    web?: pulumi.Input<inputs.ApiGatewayApiWeb>;
    /**
     * The web policy backends.
     */
    webPolicies?: pulumi.Input<pulumi.Input<inputs.ApiGatewayApiWebPolicy>[]>;
}
