// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

export class DwsCluster extends pulumi.CustomResource {
    /**
     * Get an existing DwsCluster resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: DwsClusterState, opts?: pulumi.CustomResourceOptions): DwsCluster {
        return new DwsCluster(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'sbercloud:index/dwsCluster:DwsCluster';

    /**
     * Returns true if the given object is an instance of DwsCluster.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is DwsCluster {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === DwsCluster.__pulumiType;
    }

    /**
     * The availability zone in which to create the cluster instance.
     */
    public readonly availabilityZone!: pulumi.Output<string>;
    /**
     * The creation time of the cluster.
     */
    public /*out*/ readonly created!: pulumi.Output<string>;
    /**
     * The description of the cluster.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * Dedicated storage pool ID.
     */
    public readonly dssPoolId!: pulumi.Output<string>;
    /**
     * The ID of the ELB load balancer.
     */
    public readonly elbId!: pulumi.Output<string | undefined>;
    /**
     * The ELB information bound to the cluster.
     */
    public /*out*/ readonly elbs!: pulumi.Output<outputs.DwsClusterElb[]>;
    /**
     * Private network connection information about the cluster.
     */
    public /*out*/ readonly endpoints!: pulumi.Output<outputs.DwsClusterEndpoint[]>;
    /**
     * The enterprise project ID.
     */
    public readonly enterpriseProjectId!: pulumi.Output<string>;
    /**
     * Whether to automatically execute snapshot when shrinking the number of nodes.
     */
    public readonly forceBackup!: pulumi.Output<boolean | undefined>;
    /**
     * The number of latest manual snapshots that need to be retained when deleting the cluster.
     */
    public readonly keepLastManualSnapshot!: pulumi.Output<number | undefined>;
    /**
     * The KMS key ID.
     */
    public readonly kmsKeyId!: pulumi.Output<string>;
    /**
     * Whether to enable logical cluster.
     */
    public readonly logicalClusterEnable!: pulumi.Output<boolean | undefined>;
    /**
     * Whether to enable LTS.
     */
    public readonly ltsEnable!: pulumi.Output<boolean | undefined>;
    /**
     * Cluster maintenance window.
     */
    public /*out*/ readonly maintainWindows!: pulumi.Output<outputs.DwsClusterMaintainWindow[]>;
    /**
     * The cluster name.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The subnet ID.
     */
    public readonly networkId!: pulumi.Output<string>;
    /**
     * The flavor of the cluster.
     */
    public readonly nodeType!: pulumi.Output<string>;
    /**
     * schema: Required
     */
    public readonly numberOfCn!: pulumi.Output<number | undefined>;
    /**
     * Number of nodes in a cluster.
     */
    public readonly numberOfNode!: pulumi.Output<number>;
    /**
     * Service port of a cluster (8000 to 10000). The default value is 8000.
     */
    public readonly port!: pulumi.Output<number | undefined>;
    /**
     * List of private network IP addresses.
     */
    public /*out*/ readonly privateIps!: pulumi.Output<string[]>;
    /**
     * Public network connection information about the cluster.
     */
    public /*out*/ readonly publicEndpoints!: pulumi.Output<outputs.DwsClusterPublicEndpoint[]>;
    public readonly publicIp!: pulumi.Output<outputs.DwsClusterPublicIp>;
    public /*out*/ readonly recentEvent!: pulumi.Output<number>;
    public readonly region!: pulumi.Output<string>;
    /**
     * The security group ID.
     */
    public readonly securityGroupId!: pulumi.Output<string>;
    /**
     * The cluster status.
     */
    public /*out*/ readonly status!: pulumi.Output<string>;
    /**
     * Sub-status of clusters in the AVAILABLE state.
     */
    public /*out*/ readonly subStatus!: pulumi.Output<string>;
    /**
     * The key/value pairs to associate with the cluster.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * Cluster management task.
     */
    public /*out*/ readonly taskStatus!: pulumi.Output<string>;
    /**
     * The updated time of the cluster.
     */
    public /*out*/ readonly updated!: pulumi.Output<string>;
    /**
     * Administrator username for logging in to a data warehouse cluster.
     */
    public readonly userName!: pulumi.Output<string>;
    /**
     * Administrator password for logging in to a data warehouse cluster.
     */
    public readonly userPwd!: pulumi.Output<string>;
    /**
     * schema: Required
     */
    public readonly version!: pulumi.Output<string>;
    public readonly volume!: pulumi.Output<outputs.DwsClusterVolume>;
    /**
     * The VPC ID.
     */
    public readonly vpcId!: pulumi.Output<string>;

    /**
     * Create a DwsCluster resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: DwsClusterArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: DwsClusterArgs | DwsClusterState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as DwsClusterState | undefined;
            resourceInputs["availabilityZone"] = state ? state.availabilityZone : undefined;
            resourceInputs["created"] = state ? state.created : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["dssPoolId"] = state ? state.dssPoolId : undefined;
            resourceInputs["elbId"] = state ? state.elbId : undefined;
            resourceInputs["elbs"] = state ? state.elbs : undefined;
            resourceInputs["endpoints"] = state ? state.endpoints : undefined;
            resourceInputs["enterpriseProjectId"] = state ? state.enterpriseProjectId : undefined;
            resourceInputs["forceBackup"] = state ? state.forceBackup : undefined;
            resourceInputs["keepLastManualSnapshot"] = state ? state.keepLastManualSnapshot : undefined;
            resourceInputs["kmsKeyId"] = state ? state.kmsKeyId : undefined;
            resourceInputs["logicalClusterEnable"] = state ? state.logicalClusterEnable : undefined;
            resourceInputs["ltsEnable"] = state ? state.ltsEnable : undefined;
            resourceInputs["maintainWindows"] = state ? state.maintainWindows : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["networkId"] = state ? state.networkId : undefined;
            resourceInputs["nodeType"] = state ? state.nodeType : undefined;
            resourceInputs["numberOfCn"] = state ? state.numberOfCn : undefined;
            resourceInputs["numberOfNode"] = state ? state.numberOfNode : undefined;
            resourceInputs["port"] = state ? state.port : undefined;
            resourceInputs["privateIps"] = state ? state.privateIps : undefined;
            resourceInputs["publicEndpoints"] = state ? state.publicEndpoints : undefined;
            resourceInputs["publicIp"] = state ? state.publicIp : undefined;
            resourceInputs["recentEvent"] = state ? state.recentEvent : undefined;
            resourceInputs["region"] = state ? state.region : undefined;
            resourceInputs["securityGroupId"] = state ? state.securityGroupId : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
            resourceInputs["subStatus"] = state ? state.subStatus : undefined;
            resourceInputs["tags"] = state ? state.tags : undefined;
            resourceInputs["taskStatus"] = state ? state.taskStatus : undefined;
            resourceInputs["updated"] = state ? state.updated : undefined;
            resourceInputs["userName"] = state ? state.userName : undefined;
            resourceInputs["userPwd"] = state ? state.userPwd : undefined;
            resourceInputs["version"] = state ? state.version : undefined;
            resourceInputs["volume"] = state ? state.volume : undefined;
            resourceInputs["vpcId"] = state ? state.vpcId : undefined;
        } else {
            const args = argsOrState as DwsClusterArgs | undefined;
            if ((!args || args.availabilityZone === undefined) && !opts.urn) {
                throw new Error("Missing required property 'availabilityZone'");
            }
            if ((!args || args.networkId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'networkId'");
            }
            if ((!args || args.nodeType === undefined) && !opts.urn) {
                throw new Error("Missing required property 'nodeType'");
            }
            if ((!args || args.numberOfNode === undefined) && !opts.urn) {
                throw new Error("Missing required property 'numberOfNode'");
            }
            if ((!args || args.securityGroupId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'securityGroupId'");
            }
            if ((!args || args.userName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'userName'");
            }
            if ((!args || args.userPwd === undefined) && !opts.urn) {
                throw new Error("Missing required property 'userPwd'");
            }
            if ((!args || args.vpcId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'vpcId'");
            }
            resourceInputs["availabilityZone"] = args ? args.availabilityZone : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["dssPoolId"] = args ? args.dssPoolId : undefined;
            resourceInputs["elbId"] = args ? args.elbId : undefined;
            resourceInputs["enterpriseProjectId"] = args ? args.enterpriseProjectId : undefined;
            resourceInputs["forceBackup"] = args ? args.forceBackup : undefined;
            resourceInputs["keepLastManualSnapshot"] = args ? args.keepLastManualSnapshot : undefined;
            resourceInputs["kmsKeyId"] = args ? args.kmsKeyId : undefined;
            resourceInputs["logicalClusterEnable"] = args ? args.logicalClusterEnable : undefined;
            resourceInputs["ltsEnable"] = args ? args.ltsEnable : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["networkId"] = args ? args.networkId : undefined;
            resourceInputs["nodeType"] = args ? args.nodeType : undefined;
            resourceInputs["numberOfCn"] = args ? args.numberOfCn : undefined;
            resourceInputs["numberOfNode"] = args ? args.numberOfNode : undefined;
            resourceInputs["port"] = args ? args.port : undefined;
            resourceInputs["publicIp"] = args ? args.publicIp : undefined;
            resourceInputs["region"] = args ? args.region : undefined;
            resourceInputs["securityGroupId"] = args ? args.securityGroupId : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["userName"] = args ? args.userName : undefined;
            resourceInputs["userPwd"] = args?.userPwd ? pulumi.secret(args.userPwd) : undefined;
            resourceInputs["version"] = args ? args.version : undefined;
            resourceInputs["volume"] = args ? args.volume : undefined;
            resourceInputs["vpcId"] = args ? args.vpcId : undefined;
            resourceInputs["created"] = undefined /*out*/;
            resourceInputs["elbs"] = undefined /*out*/;
            resourceInputs["endpoints"] = undefined /*out*/;
            resourceInputs["maintainWindows"] = undefined /*out*/;
            resourceInputs["privateIps"] = undefined /*out*/;
            resourceInputs["publicEndpoints"] = undefined /*out*/;
            resourceInputs["recentEvent"] = undefined /*out*/;
            resourceInputs["status"] = undefined /*out*/;
            resourceInputs["subStatus"] = undefined /*out*/;
            resourceInputs["taskStatus"] = undefined /*out*/;
            resourceInputs["updated"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const secretOpts = { additionalSecretOutputs: ["userPwd"] };
        opts = pulumi.mergeOptions(opts, secretOpts);
        super(DwsCluster.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering DwsCluster resources.
 */
export interface DwsClusterState {
    /**
     * The availability zone in which to create the cluster instance.
     */
    availabilityZone?: pulumi.Input<string>;
    /**
     * The creation time of the cluster.
     */
    created?: pulumi.Input<string>;
    /**
     * The description of the cluster.
     */
    description?: pulumi.Input<string>;
    /**
     * Dedicated storage pool ID.
     */
    dssPoolId?: pulumi.Input<string>;
    /**
     * The ID of the ELB load balancer.
     */
    elbId?: pulumi.Input<string>;
    /**
     * The ELB information bound to the cluster.
     */
    elbs?: pulumi.Input<pulumi.Input<inputs.DwsClusterElb>[]>;
    /**
     * Private network connection information about the cluster.
     */
    endpoints?: pulumi.Input<pulumi.Input<inputs.DwsClusterEndpoint>[]>;
    /**
     * The enterprise project ID.
     */
    enterpriseProjectId?: pulumi.Input<string>;
    /**
     * Whether to automatically execute snapshot when shrinking the number of nodes.
     */
    forceBackup?: pulumi.Input<boolean>;
    /**
     * The number of latest manual snapshots that need to be retained when deleting the cluster.
     */
    keepLastManualSnapshot?: pulumi.Input<number>;
    /**
     * The KMS key ID.
     */
    kmsKeyId?: pulumi.Input<string>;
    /**
     * Whether to enable logical cluster.
     */
    logicalClusterEnable?: pulumi.Input<boolean>;
    /**
     * Whether to enable LTS.
     */
    ltsEnable?: pulumi.Input<boolean>;
    /**
     * Cluster maintenance window.
     */
    maintainWindows?: pulumi.Input<pulumi.Input<inputs.DwsClusterMaintainWindow>[]>;
    /**
     * The cluster name.
     */
    name?: pulumi.Input<string>;
    /**
     * The subnet ID.
     */
    networkId?: pulumi.Input<string>;
    /**
     * The flavor of the cluster.
     */
    nodeType?: pulumi.Input<string>;
    /**
     * schema: Required
     */
    numberOfCn?: pulumi.Input<number>;
    /**
     * Number of nodes in a cluster.
     */
    numberOfNode?: pulumi.Input<number>;
    /**
     * Service port of a cluster (8000 to 10000). The default value is 8000.
     */
    port?: pulumi.Input<number>;
    /**
     * List of private network IP addresses.
     */
    privateIps?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Public network connection information about the cluster.
     */
    publicEndpoints?: pulumi.Input<pulumi.Input<inputs.DwsClusterPublicEndpoint>[]>;
    publicIp?: pulumi.Input<inputs.DwsClusterPublicIp>;
    recentEvent?: pulumi.Input<number>;
    region?: pulumi.Input<string>;
    /**
     * The security group ID.
     */
    securityGroupId?: pulumi.Input<string>;
    /**
     * The cluster status.
     */
    status?: pulumi.Input<string>;
    /**
     * Sub-status of clusters in the AVAILABLE state.
     */
    subStatus?: pulumi.Input<string>;
    /**
     * The key/value pairs to associate with the cluster.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Cluster management task.
     */
    taskStatus?: pulumi.Input<string>;
    /**
     * The updated time of the cluster.
     */
    updated?: pulumi.Input<string>;
    /**
     * Administrator username for logging in to a data warehouse cluster.
     */
    userName?: pulumi.Input<string>;
    /**
     * Administrator password for logging in to a data warehouse cluster.
     */
    userPwd?: pulumi.Input<string>;
    /**
     * schema: Required
     */
    version?: pulumi.Input<string>;
    volume?: pulumi.Input<inputs.DwsClusterVolume>;
    /**
     * The VPC ID.
     */
    vpcId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a DwsCluster resource.
 */
export interface DwsClusterArgs {
    /**
     * The availability zone in which to create the cluster instance.
     */
    availabilityZone: pulumi.Input<string>;
    /**
     * The description of the cluster.
     */
    description?: pulumi.Input<string>;
    /**
     * Dedicated storage pool ID.
     */
    dssPoolId?: pulumi.Input<string>;
    /**
     * The ID of the ELB load balancer.
     */
    elbId?: pulumi.Input<string>;
    /**
     * The enterprise project ID.
     */
    enterpriseProjectId?: pulumi.Input<string>;
    /**
     * Whether to automatically execute snapshot when shrinking the number of nodes.
     */
    forceBackup?: pulumi.Input<boolean>;
    /**
     * The number of latest manual snapshots that need to be retained when deleting the cluster.
     */
    keepLastManualSnapshot?: pulumi.Input<number>;
    /**
     * The KMS key ID.
     */
    kmsKeyId?: pulumi.Input<string>;
    /**
     * Whether to enable logical cluster.
     */
    logicalClusterEnable?: pulumi.Input<boolean>;
    /**
     * Whether to enable LTS.
     */
    ltsEnable?: pulumi.Input<boolean>;
    /**
     * The cluster name.
     */
    name?: pulumi.Input<string>;
    /**
     * The subnet ID.
     */
    networkId: pulumi.Input<string>;
    /**
     * The flavor of the cluster.
     */
    nodeType: pulumi.Input<string>;
    /**
     * schema: Required
     */
    numberOfCn?: pulumi.Input<number>;
    /**
     * Number of nodes in a cluster.
     */
    numberOfNode: pulumi.Input<number>;
    /**
     * Service port of a cluster (8000 to 10000). The default value is 8000.
     */
    port?: pulumi.Input<number>;
    publicIp?: pulumi.Input<inputs.DwsClusterPublicIp>;
    region?: pulumi.Input<string>;
    /**
     * The security group ID.
     */
    securityGroupId: pulumi.Input<string>;
    /**
     * The key/value pairs to associate with the cluster.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Administrator username for logging in to a data warehouse cluster.
     */
    userName: pulumi.Input<string>;
    /**
     * Administrator password for logging in to a data warehouse cluster.
     */
    userPwd: pulumi.Input<string>;
    /**
     * schema: Required
     */
    version?: pulumi.Input<string>;
    volume?: pulumi.Input<inputs.DwsClusterVolume>;
    /**
     * The VPC ID.
     */
    vpcId: pulumi.Input<string>;
}
