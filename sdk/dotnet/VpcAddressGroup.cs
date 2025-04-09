// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Sbercloud
{
    [SbercloudResourceType("sbercloud:index/vpcAddressGroup:VpcAddressGroup")]
    public partial class VpcAddressGroup : global::Pulumi.CustomResource
    {
        [Output("addresses")]
        public Output<ImmutableArray<string>> Addresses { get; private set; } = null!;

        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        [Output("enterpriseProjectId")]
        public Output<string> EnterpriseProjectId { get; private set; } = null!;

        [Output("forceDestroy")]
        public Output<bool?> ForceDestroy { get; private set; } = null!;

        [Output("ipExtraSets")]
        public Output<ImmutableArray<Outputs.VpcAddressGroupIpExtraSet>> IpExtraSets { get; private set; } = null!;

        [Output("ipVersion")]
        public Output<int?> IpVersion { get; private set; } = null!;

        [Output("maxCapacity")]
        public Output<int> MaxCapacity { get; private set; } = null!;

        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        [Output("region")]
        public Output<string> Region { get; private set; } = null!;


        /// <summary>
        /// Create a VpcAddressGroup resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public VpcAddressGroup(string name, VpcAddressGroupArgs? args = null, CustomResourceOptions? options = null)
            : base("sbercloud:index/vpcAddressGroup:VpcAddressGroup", name, args ?? new VpcAddressGroupArgs(), MakeResourceOptions(options, ""))
        {
        }

        private VpcAddressGroup(string name, Input<string> id, VpcAddressGroupState? state = null, CustomResourceOptions? options = null)
            : base("sbercloud:index/vpcAddressGroup:VpcAddressGroup", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing VpcAddressGroup resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static VpcAddressGroup Get(string name, Input<string> id, VpcAddressGroupState? state = null, CustomResourceOptions? options = null)
        {
            return new VpcAddressGroup(name, id, state, options);
        }
    }

    public sealed class VpcAddressGroupArgs : global::Pulumi.ResourceArgs
    {
        [Input("addresses")]
        private InputList<string>? _addresses;
        public InputList<string> Addresses
        {
            get => _addresses ?? (_addresses = new InputList<string>());
            set => _addresses = value;
        }

        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("enterpriseProjectId")]
        public Input<string>? EnterpriseProjectId { get; set; }

        [Input("forceDestroy")]
        public Input<bool>? ForceDestroy { get; set; }

        [Input("ipExtraSets")]
        private InputList<Inputs.VpcAddressGroupIpExtraSetArgs>? _ipExtraSets;
        public InputList<Inputs.VpcAddressGroupIpExtraSetArgs> IpExtraSets
        {
            get => _ipExtraSets ?? (_ipExtraSets = new InputList<Inputs.VpcAddressGroupIpExtraSetArgs>());
            set => _ipExtraSets = value;
        }

        [Input("ipVersion")]
        public Input<int>? IpVersion { get; set; }

        [Input("maxCapacity")]
        public Input<int>? MaxCapacity { get; set; }

        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("region")]
        public Input<string>? Region { get; set; }

        public VpcAddressGroupArgs()
        {
        }
        public static new VpcAddressGroupArgs Empty => new VpcAddressGroupArgs();
    }

    public sealed class VpcAddressGroupState : global::Pulumi.ResourceArgs
    {
        [Input("addresses")]
        private InputList<string>? _addresses;
        public InputList<string> Addresses
        {
            get => _addresses ?? (_addresses = new InputList<string>());
            set => _addresses = value;
        }

        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("enterpriseProjectId")]
        public Input<string>? EnterpriseProjectId { get; set; }

        [Input("forceDestroy")]
        public Input<bool>? ForceDestroy { get; set; }

        [Input("ipExtraSets")]
        private InputList<Inputs.VpcAddressGroupIpExtraSetGetArgs>? _ipExtraSets;
        public InputList<Inputs.VpcAddressGroupIpExtraSetGetArgs> IpExtraSets
        {
            get => _ipExtraSets ?? (_ipExtraSets = new InputList<Inputs.VpcAddressGroupIpExtraSetGetArgs>());
            set => _ipExtraSets = value;
        }

        [Input("ipVersion")]
        public Input<int>? IpVersion { get; set; }

        [Input("maxCapacity")]
        public Input<int>? MaxCapacity { get; set; }

        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("region")]
        public Input<string>? Region { get; set; }

        public VpcAddressGroupState()
        {
        }
        public static new VpcAddressGroupState Empty => new VpcAddressGroupState();
    }
}
