// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Sbercloud
{
    [SbercloudResourceType("sbercloud:index/cbrVault:CbrVault")]
    public partial class CbrVault : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The allocated capacity, in GB.
        /// </summary>
        [Output("allocated")]
        public Output<double> Allocated { get; private set; } = null!;

        /// <summary>
        /// Whether automatic association is supported.
        /// </summary>
        [Output("autoBind")]
        public Output<bool> AutoBind { get; private set; } = null!;

        /// <summary>
        /// Whether to enable auto capacity expansion for the vault.
        /// </summary>
        [Output("autoExpand")]
        public Output<bool> AutoExpand { get; private set; } = null!;

        [Output("autoPay")]
        public Output<string?> AutoPay { get; private set; } = null!;

        [Output("autoRenew")]
        public Output<string?> AutoRenew { get; private set; } = null!;

        /// <summary>
        /// The backup name prefix.
        /// </summary>
        [Output("backupNamePrefix")]
        public Output<string> BackupNamePrefix { get; private set; } = null!;

        /// <summary>
        /// The rules for automatic association.
        /// </summary>
        [Output("bindRules")]
        public Output<ImmutableDictionary<string, string>?> BindRules { get; private set; } = null!;

        [Output("chargingMode")]
        public Output<string> ChargingMode { get; private set; } = null!;

        /// <summary>
        /// The cloud type of the vault.
        /// </summary>
        [Output("cloudType")]
        public Output<string> CloudType { get; private set; } = null!;

        /// <summary>
        /// The consistent level (specification) of the vault.
        /// </summary>
        [Output("consistentLevel")]
        public Output<string?> ConsistentLevel { get; private set; } = null!;

        /// <summary>
        /// The enterprise project ID to which the vault belongs.
        /// </summary>
        [Output("enterpriseProjectId")]
        public Output<string> EnterpriseProjectId { get; private set; } = null!;

        /// <summary>
        /// Whether multiple availability zones are used for backing up.
        /// </summary>
        [Output("isMultiAz")]
        public Output<bool> IsMultiAz { get; private set; } = null!;

        /// <summary>
        /// The name of the vault.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        [Output("period")]
        public Output<int?> Period { get; private set; } = null!;

        [Output("periodUnit")]
        public Output<string?> PeriodUnit { get; private set; } = null!;

        /// <summary>
        /// The policy details to associate with the CBR vault.
        /// </summary>
        [Output("policies")]
        public Output<ImmutableArray<Outputs.CbrVaultPolicy>> Policies { get; private set; } = null!;

        /// <summary>
        /// schema:Deprecated; Using parameter 'policy' instead.
        /// </summary>
        [Output("policyId")]
        public Output<string?> PolicyId { get; private set; } = null!;

        /// <summary>
        /// The protection type of the vault.
        /// </summary>
        [Output("protectionType")]
        public Output<string> ProtectionType { get; private set; } = null!;

        /// <summary>
        /// The region where the vault is located.
        /// </summary>
        [Output("region")]
        public Output<string> Region { get; private set; } = null!;

        /// <summary>
        /// The array of one or more resources to attach to the CBR vault.
        /// </summary>
        [Output("resources")]
        public Output<ImmutableArray<Outputs.CbrVaultResource>> Resources { get; private set; } = null!;

        /// <summary>
        /// The capacity of the vault, in GB.
        /// </summary>
        [Output("size")]
        public Output<int> Size { get; private set; } = null!;

        /// <summary>
        /// The specification code.
        /// </summary>
        [Output("specCode")]
        public Output<string> SpecCode { get; private set; } = null!;

        /// <summary>
        /// The vault status.
        /// </summary>
        [Output("status")]
        public Output<string> Status { get; private set; } = null!;

        /// <summary>
        /// The name of the bucket for the vault.
        /// </summary>
        [Output("storage")]
        public Output<string> Storage { get; private set; } = null!;

        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;

        /// <summary>
        /// The type of the vault.
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;

        /// <summary>
        /// The used capacity, in GB.
        /// </summary>
        [Output("used")]
        public Output<double> Used { get; private set; } = null!;


        /// <summary>
        /// Create a CbrVault resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public CbrVault(string name, CbrVaultArgs args, CustomResourceOptions? options = null)
            : base("sbercloud:index/cbrVault:CbrVault", name, args ?? new CbrVaultArgs(), MakeResourceOptions(options, ""))
        {
        }

        private CbrVault(string name, Input<string> id, CbrVaultState? state = null, CustomResourceOptions? options = null)
            : base("sbercloud:index/cbrVault:CbrVault", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing CbrVault resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static CbrVault Get(string name, Input<string> id, CbrVaultState? state = null, CustomResourceOptions? options = null)
        {
            return new CbrVault(name, id, state, options);
        }
    }

    public sealed class CbrVaultArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Whether automatic association is supported.
        /// </summary>
        [Input("autoBind")]
        public Input<bool>? AutoBind { get; set; }

        /// <summary>
        /// Whether to enable auto capacity expansion for the vault.
        /// </summary>
        [Input("autoExpand")]
        public Input<bool>? AutoExpand { get; set; }

        [Input("autoPay")]
        public Input<string>? AutoPay { get; set; }

        [Input("autoRenew")]
        public Input<string>? AutoRenew { get; set; }

        /// <summary>
        /// The backup name prefix.
        /// </summary>
        [Input("backupNamePrefix")]
        public Input<string>? BackupNamePrefix { get; set; }

        [Input("bindRules")]
        private InputMap<string>? _bindRules;

        /// <summary>
        /// The rules for automatic association.
        /// </summary>
        public InputMap<string> BindRules
        {
            get => _bindRules ?? (_bindRules = new InputMap<string>());
            set => _bindRules = value;
        }

        [Input("chargingMode")]
        public Input<string>? ChargingMode { get; set; }

        /// <summary>
        /// The cloud type of the vault.
        /// </summary>
        [Input("cloudType")]
        public Input<string>? CloudType { get; set; }

        /// <summary>
        /// The consistent level (specification) of the vault.
        /// </summary>
        [Input("consistentLevel")]
        public Input<string>? ConsistentLevel { get; set; }

        /// <summary>
        /// The enterprise project ID to which the vault belongs.
        /// </summary>
        [Input("enterpriseProjectId")]
        public Input<string>? EnterpriseProjectId { get; set; }

        /// <summary>
        /// Whether multiple availability zones are used for backing up.
        /// </summary>
        [Input("isMultiAz")]
        public Input<bool>? IsMultiAz { get; set; }

        /// <summary>
        /// The name of the vault.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("period")]
        public Input<int>? Period { get; set; }

        [Input("periodUnit")]
        public Input<string>? PeriodUnit { get; set; }

        [Input("policies")]
        private InputList<Inputs.CbrVaultPolicyArgs>? _policies;

        /// <summary>
        /// The policy details to associate with the CBR vault.
        /// </summary>
        public InputList<Inputs.CbrVaultPolicyArgs> Policies
        {
            get => _policies ?? (_policies = new InputList<Inputs.CbrVaultPolicyArgs>());
            set => _policies = value;
        }

        /// <summary>
        /// schema:Deprecated; Using parameter 'policy' instead.
        /// </summary>
        [Input("policyId")]
        public Input<string>? PolicyId { get; set; }

        /// <summary>
        /// The protection type of the vault.
        /// </summary>
        [Input("protectionType", required: true)]
        public Input<string> ProtectionType { get; set; } = null!;

        /// <summary>
        /// The region where the vault is located.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        [Input("resources")]
        private InputList<Inputs.CbrVaultResourceArgs>? _resources;

        /// <summary>
        /// The array of one or more resources to attach to the CBR vault.
        /// </summary>
        public InputList<Inputs.CbrVaultResourceArgs> Resources
        {
            get => _resources ?? (_resources = new InputList<Inputs.CbrVaultResourceArgs>());
            set => _resources = value;
        }

        /// <summary>
        /// The capacity of the vault, in GB.
        /// </summary>
        [Input("size", required: true)]
        public Input<int> Size { get; set; } = null!;

        [Input("tags")]
        private InputMap<string>? _tags;
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        /// <summary>
        /// The type of the vault.
        /// </summary>
        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        public CbrVaultArgs()
        {
        }
        public static new CbrVaultArgs Empty => new CbrVaultArgs();
    }

    public sealed class CbrVaultState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The allocated capacity, in GB.
        /// </summary>
        [Input("allocated")]
        public Input<double>? Allocated { get; set; }

        /// <summary>
        /// Whether automatic association is supported.
        /// </summary>
        [Input("autoBind")]
        public Input<bool>? AutoBind { get; set; }

        /// <summary>
        /// Whether to enable auto capacity expansion for the vault.
        /// </summary>
        [Input("autoExpand")]
        public Input<bool>? AutoExpand { get; set; }

        [Input("autoPay")]
        public Input<string>? AutoPay { get; set; }

        [Input("autoRenew")]
        public Input<string>? AutoRenew { get; set; }

        /// <summary>
        /// The backup name prefix.
        /// </summary>
        [Input("backupNamePrefix")]
        public Input<string>? BackupNamePrefix { get; set; }

        [Input("bindRules")]
        private InputMap<string>? _bindRules;

        /// <summary>
        /// The rules for automatic association.
        /// </summary>
        public InputMap<string> BindRules
        {
            get => _bindRules ?? (_bindRules = new InputMap<string>());
            set => _bindRules = value;
        }

        [Input("chargingMode")]
        public Input<string>? ChargingMode { get; set; }

        /// <summary>
        /// The cloud type of the vault.
        /// </summary>
        [Input("cloudType")]
        public Input<string>? CloudType { get; set; }

        /// <summary>
        /// The consistent level (specification) of the vault.
        /// </summary>
        [Input("consistentLevel")]
        public Input<string>? ConsistentLevel { get; set; }

        /// <summary>
        /// The enterprise project ID to which the vault belongs.
        /// </summary>
        [Input("enterpriseProjectId")]
        public Input<string>? EnterpriseProjectId { get; set; }

        /// <summary>
        /// Whether multiple availability zones are used for backing up.
        /// </summary>
        [Input("isMultiAz")]
        public Input<bool>? IsMultiAz { get; set; }

        /// <summary>
        /// The name of the vault.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("period")]
        public Input<int>? Period { get; set; }

        [Input("periodUnit")]
        public Input<string>? PeriodUnit { get; set; }

        [Input("policies")]
        private InputList<Inputs.CbrVaultPolicyGetArgs>? _policies;

        /// <summary>
        /// The policy details to associate with the CBR vault.
        /// </summary>
        public InputList<Inputs.CbrVaultPolicyGetArgs> Policies
        {
            get => _policies ?? (_policies = new InputList<Inputs.CbrVaultPolicyGetArgs>());
            set => _policies = value;
        }

        /// <summary>
        /// schema:Deprecated; Using parameter 'policy' instead.
        /// </summary>
        [Input("policyId")]
        public Input<string>? PolicyId { get; set; }

        /// <summary>
        /// The protection type of the vault.
        /// </summary>
        [Input("protectionType")]
        public Input<string>? ProtectionType { get; set; }

        /// <summary>
        /// The region where the vault is located.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        [Input("resources")]
        private InputList<Inputs.CbrVaultResourceGetArgs>? _resources;

        /// <summary>
        /// The array of one or more resources to attach to the CBR vault.
        /// </summary>
        public InputList<Inputs.CbrVaultResourceGetArgs> Resources
        {
            get => _resources ?? (_resources = new InputList<Inputs.CbrVaultResourceGetArgs>());
            set => _resources = value;
        }

        /// <summary>
        /// The capacity of the vault, in GB.
        /// </summary>
        [Input("size")]
        public Input<int>? Size { get; set; }

        /// <summary>
        /// The specification code.
        /// </summary>
        [Input("specCode")]
        public Input<string>? SpecCode { get; set; }

        /// <summary>
        /// The vault status.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        /// <summary>
        /// The name of the bucket for the vault.
        /// </summary>
        [Input("storage")]
        public Input<string>? Storage { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        /// <summary>
        /// The type of the vault.
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        /// <summary>
        /// The used capacity, in GB.
        /// </summary>
        [Input("used")]
        public Input<double>? Used { get; set; }

        public CbrVaultState()
        {
        }
        public static new CbrVaultState Empty => new CbrVaultState();
    }
}
