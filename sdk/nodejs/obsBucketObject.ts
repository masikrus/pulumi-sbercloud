// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

export class ObsBucketObject extends pulumi.CustomResource {
    /**
     * Get an existing ObsBucketObject resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ObsBucketObjectState, opts?: pulumi.CustomResourceOptions): ObsBucketObject {
        return new ObsBucketObject(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'sbercloud:index/obsBucketObject:ObsBucketObject';

    /**
     * Returns true if the given object is an instance of ObsBucketObject.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ObsBucketObject {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ObsBucketObject.__pulumiType;
    }

    public readonly acl!: pulumi.Output<string | undefined>;
    public readonly bucket!: pulumi.Output<string>;
    public readonly content!: pulumi.Output<string | undefined>;
    public readonly contentType!: pulumi.Output<string>;
    public readonly encryption!: pulumi.Output<boolean | undefined>;
    public readonly etag!: pulumi.Output<string>;
    public readonly key!: pulumi.Output<string>;
    public readonly kmsKeyId!: pulumi.Output<string | undefined>;
    public readonly region!: pulumi.Output<string>;
    public /*out*/ readonly size!: pulumi.Output<number>;
    public readonly source!: pulumi.Output<string | undefined>;
    public readonly storageClass!: pulumi.Output<string>;
    public /*out*/ readonly versionId!: pulumi.Output<string>;

    /**
     * Create a ObsBucketObject resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ObsBucketObjectArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ObsBucketObjectArgs | ObsBucketObjectState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ObsBucketObjectState | undefined;
            resourceInputs["acl"] = state ? state.acl : undefined;
            resourceInputs["bucket"] = state ? state.bucket : undefined;
            resourceInputs["content"] = state ? state.content : undefined;
            resourceInputs["contentType"] = state ? state.contentType : undefined;
            resourceInputs["encryption"] = state ? state.encryption : undefined;
            resourceInputs["etag"] = state ? state.etag : undefined;
            resourceInputs["key"] = state ? state.key : undefined;
            resourceInputs["kmsKeyId"] = state ? state.kmsKeyId : undefined;
            resourceInputs["region"] = state ? state.region : undefined;
            resourceInputs["size"] = state ? state.size : undefined;
            resourceInputs["source"] = state ? state.source : undefined;
            resourceInputs["storageClass"] = state ? state.storageClass : undefined;
            resourceInputs["versionId"] = state ? state.versionId : undefined;
        } else {
            const args = argsOrState as ObsBucketObjectArgs | undefined;
            if ((!args || args.bucket === undefined) && !opts.urn) {
                throw new Error("Missing required property 'bucket'");
            }
            if ((!args || args.key === undefined) && !opts.urn) {
                throw new Error("Missing required property 'key'");
            }
            resourceInputs["acl"] = args ? args.acl : undefined;
            resourceInputs["bucket"] = args ? args.bucket : undefined;
            resourceInputs["content"] = args ? args.content : undefined;
            resourceInputs["contentType"] = args ? args.contentType : undefined;
            resourceInputs["encryption"] = args ? args.encryption : undefined;
            resourceInputs["etag"] = args ? args.etag : undefined;
            resourceInputs["key"] = args ? args.key : undefined;
            resourceInputs["kmsKeyId"] = args ? args.kmsKeyId : undefined;
            resourceInputs["region"] = args ? args.region : undefined;
            resourceInputs["source"] = args ? args.source : undefined;
            resourceInputs["storageClass"] = args ? args.storageClass : undefined;
            resourceInputs["size"] = undefined /*out*/;
            resourceInputs["versionId"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(ObsBucketObject.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ObsBucketObject resources.
 */
export interface ObsBucketObjectState {
    acl?: pulumi.Input<string>;
    bucket?: pulumi.Input<string>;
    content?: pulumi.Input<string>;
    contentType?: pulumi.Input<string>;
    encryption?: pulumi.Input<boolean>;
    etag?: pulumi.Input<string>;
    key?: pulumi.Input<string>;
    kmsKeyId?: pulumi.Input<string>;
    region?: pulumi.Input<string>;
    size?: pulumi.Input<number>;
    source?: pulumi.Input<string>;
    storageClass?: pulumi.Input<string>;
    versionId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a ObsBucketObject resource.
 */
export interface ObsBucketObjectArgs {
    acl?: pulumi.Input<string>;
    bucket: pulumi.Input<string>;
    content?: pulumi.Input<string>;
    contentType?: pulumi.Input<string>;
    encryption?: pulumi.Input<boolean>;
    etag?: pulumi.Input<string>;
    key: pulumi.Input<string>;
    kmsKeyId?: pulumi.Input<string>;
    region?: pulumi.Input<string>;
    source?: pulumi.Input<string>;
    storageClass?: pulumi.Input<string>;
}
