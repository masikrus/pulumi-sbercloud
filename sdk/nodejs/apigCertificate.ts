// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

export class ApigCertificate extends pulumi.CustomResource {
    /**
     * Get an existing ApigCertificate resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ApigCertificateState, opts?: pulumi.CustomResourceOptions): ApigCertificate {
        return new ApigCertificate(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'sbercloud:index/apigCertificate:ApigCertificate';

    /**
     * Returns true if the given object is an instance of ApigCertificate.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ApigCertificate {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ApigCertificate.__pulumiType;
    }

    /**
     * The certificate content.
     */
    public readonly content!: pulumi.Output<string>;
    /**
     * The effective time of the certificate.
     */
    public /*out*/ readonly effectedAt!: pulumi.Output<string>;
    /**
     * The expiration time of the certificate.
     */
    public /*out*/ readonly expiresAt!: pulumi.Output<string>;
    /**
     * The dedicated instance ID to which the certificate belongs.
     */
    public readonly instanceId!: pulumi.Output<string>;
    /**
     * The certificate name.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The private key of the certificate.
     */
    public readonly privateKey!: pulumi.Output<string>;
    /**
     * The region where the certificate is located.
     */
    public readonly region!: pulumi.Output<string>;
    /**
     * The SAN (Subject Alternative Names) of the certificate.
     */
    public /*out*/ readonly sans!: pulumi.Output<string[]>;
    /**
     * What signature algorithm the certificate uses.
     */
    public /*out*/ readonly signatureAlgorithm!: pulumi.Output<string>;
    /**
     * The trusted root CA certificate.
     */
    public readonly trustedRootCa!: pulumi.Output<string | undefined>;
    /**
     * The certificate type.
     */
    public readonly type!: pulumi.Output<string>;

    /**
     * Create a ApigCertificate resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ApigCertificateArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ApigCertificateArgs | ApigCertificateState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ApigCertificateState | undefined;
            resourceInputs["content"] = state ? state.content : undefined;
            resourceInputs["effectedAt"] = state ? state.effectedAt : undefined;
            resourceInputs["expiresAt"] = state ? state.expiresAt : undefined;
            resourceInputs["instanceId"] = state ? state.instanceId : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["privateKey"] = state ? state.privateKey : undefined;
            resourceInputs["region"] = state ? state.region : undefined;
            resourceInputs["sans"] = state ? state.sans : undefined;
            resourceInputs["signatureAlgorithm"] = state ? state.signatureAlgorithm : undefined;
            resourceInputs["trustedRootCa"] = state ? state.trustedRootCa : undefined;
            resourceInputs["type"] = state ? state.type : undefined;
        } else {
            const args = argsOrState as ApigCertificateArgs | undefined;
            if ((!args || args.content === undefined) && !opts.urn) {
                throw new Error("Missing required property 'content'");
            }
            if ((!args || args.privateKey === undefined) && !opts.urn) {
                throw new Error("Missing required property 'privateKey'");
            }
            resourceInputs["content"] = args?.content ? pulumi.secret(args.content) : undefined;
            resourceInputs["instanceId"] = args ? args.instanceId : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["privateKey"] = args?.privateKey ? pulumi.secret(args.privateKey) : undefined;
            resourceInputs["region"] = args ? args.region : undefined;
            resourceInputs["trustedRootCa"] = args?.trustedRootCa ? pulumi.secret(args.trustedRootCa) : undefined;
            resourceInputs["type"] = args ? args.type : undefined;
            resourceInputs["effectedAt"] = undefined /*out*/;
            resourceInputs["expiresAt"] = undefined /*out*/;
            resourceInputs["sans"] = undefined /*out*/;
            resourceInputs["signatureAlgorithm"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const secretOpts = { additionalSecretOutputs: ["content", "privateKey", "trustedRootCa"] };
        opts = pulumi.mergeOptions(opts, secretOpts);
        super(ApigCertificate.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ApigCertificate resources.
 */
export interface ApigCertificateState {
    /**
     * The certificate content.
     */
    content?: pulumi.Input<string>;
    /**
     * The effective time of the certificate.
     */
    effectedAt?: pulumi.Input<string>;
    /**
     * The expiration time of the certificate.
     */
    expiresAt?: pulumi.Input<string>;
    /**
     * The dedicated instance ID to which the certificate belongs.
     */
    instanceId?: pulumi.Input<string>;
    /**
     * The certificate name.
     */
    name?: pulumi.Input<string>;
    /**
     * The private key of the certificate.
     */
    privateKey?: pulumi.Input<string>;
    /**
     * The region where the certificate is located.
     */
    region?: pulumi.Input<string>;
    /**
     * The SAN (Subject Alternative Names) of the certificate.
     */
    sans?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * What signature algorithm the certificate uses.
     */
    signatureAlgorithm?: pulumi.Input<string>;
    /**
     * The trusted root CA certificate.
     */
    trustedRootCa?: pulumi.Input<string>;
    /**
     * The certificate type.
     */
    type?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a ApigCertificate resource.
 */
export interface ApigCertificateArgs {
    /**
     * The certificate content.
     */
    content: pulumi.Input<string>;
    /**
     * The dedicated instance ID to which the certificate belongs.
     */
    instanceId?: pulumi.Input<string>;
    /**
     * The certificate name.
     */
    name?: pulumi.Input<string>;
    /**
     * The private key of the certificate.
     */
    privateKey: pulumi.Input<string>;
    /**
     * The region where the certificate is located.
     */
    region?: pulumi.Input<string>;
    /**
     * The trusted root CA certificate.
     */
    trustedRootCa?: pulumi.Input<string>;
    /**
     * The certificate type.
     */
    type?: pulumi.Input<string>;
}
