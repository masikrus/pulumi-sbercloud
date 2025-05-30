// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Sbercloud
{
    [SbercloudResourceType("sbercloud:index/vpcBandwidth:VpcBandwidth")]
    public partial class VpcBandwidth : global::Pulumi.CustomResource
    {
        [Output("autoRenew")]
        public Output<string?> AutoRenew { get; private set; } = null!;

        [Output("bandwidthType")]
        public Output<string> BandwidthType { get; private set; } = null!;

        [Output("chargeMode")]
        public Output<string> ChargeMode { get; private set; } = null!;

        [Output("chargingMode")]
        public Output<string> ChargingMode { get; private set; } = null!;

        [Output("createdAt")]
        public Output<string> CreatedAt { get; private set; } = null!;

        [Output("enterpriseProjectId")]
        public Output<string> EnterpriseProjectId { get; private set; } = null!;

        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        [Output("period")]
        public Output<int?> Period { get; private set; } = null!;

        [Output("periodUnit")]
        public Output<string?> PeriodUnit { get; private set; } = null!;

        [Output("publicBorderGroup")]
        public Output<string> PublicBorderGroup { get; private set; } = null!;

        [Output("publicips")]
        public Output<ImmutableArray<Outputs.VpcBandwidthPublicip>> Publicips { get; private set; } = null!;

        [Output("region")]
        public Output<string> Region { get; private set; } = null!;

        [Output("shareType")]
        public Output<string> ShareType { get; private set; } = null!;

        [Output("size")]
        public Output<int> Size { get; private set; } = null!;

        [Output("status")]
        public Output<string> Status { get; private set; } = null!;

        [Output("updatedAt")]
        public Output<string> UpdatedAt { get; private set; } = null!;


        /// <summary>
        /// Create a VpcBandwidth resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public VpcBandwidth(string name, VpcBandwidthArgs args, CustomResourceOptions? options = null)
            : base("sbercloud:index/vpcBandwidth:VpcBandwidth", name, args ?? new VpcBandwidthArgs(), MakeResourceOptions(options, ""))
        {
        }

        private VpcBandwidth(string name, Input<string> id, VpcBandwidthState? state = null, CustomResourceOptions? options = null)
            : base("sbercloud:index/vpcBandwidth:VpcBandwidth", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing VpcBandwidth resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static VpcBandwidth Get(string name, Input<string> id, VpcBandwidthState? state = null, CustomResourceOptions? options = null)
        {
            return new VpcBandwidth(name, id, state, options);
        }
    }

    public sealed class VpcBandwidthArgs : global::Pulumi.ResourceArgs
    {
        [Input("autoRenew")]
        public Input<string>? AutoRenew { get; set; }

        [Input("bandwidthType")]
        public Input<string>? BandwidthType { get; set; }

        [Input("chargeMode")]
        public Input<string>? ChargeMode { get; set; }

        [Input("chargingMode")]
        public Input<string>? ChargingMode { get; set; }

        [Input("enterpriseProjectId")]
        public Input<string>? EnterpriseProjectId { get; set; }

        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("period")]
        public Input<int>? Period { get; set; }

        [Input("periodUnit")]
        public Input<string>? PeriodUnit { get; set; }

        [Input("publicBorderGroup")]
        public Input<string>? PublicBorderGroup { get; set; }

        [Input("region")]
        public Input<string>? Region { get; set; }

        [Input("size", required: true)]
        public Input<int> Size { get; set; } = null!;

        public VpcBandwidthArgs()
        {
        }
        public static new VpcBandwidthArgs Empty => new VpcBandwidthArgs();
    }

    public sealed class VpcBandwidthState : global::Pulumi.ResourceArgs
    {
        [Input("autoRenew")]
        public Input<string>? AutoRenew { get; set; }

        [Input("bandwidthType")]
        public Input<string>? BandwidthType { get; set; }

        [Input("chargeMode")]
        public Input<string>? ChargeMode { get; set; }

        [Input("chargingMode")]
        public Input<string>? ChargingMode { get; set; }

        [Input("createdAt")]
        public Input<string>? CreatedAt { get; set; }

        [Input("enterpriseProjectId")]
        public Input<string>? EnterpriseProjectId { get; set; }

        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("period")]
        public Input<int>? Period { get; set; }

        [Input("periodUnit")]
        public Input<string>? PeriodUnit { get; set; }

        [Input("publicBorderGroup")]
        public Input<string>? PublicBorderGroup { get; set; }

        [Input("publicips")]
        private InputList<Inputs.VpcBandwidthPublicipGetArgs>? _publicips;
        public InputList<Inputs.VpcBandwidthPublicipGetArgs> Publicips
        {
            get => _publicips ?? (_publicips = new InputList<Inputs.VpcBandwidthPublicipGetArgs>());
            set => _publicips = value;
        }

        [Input("region")]
        public Input<string>? Region { get; set; }

        [Input("shareType")]
        public Input<string>? ShareType { get; set; }

        [Input("size")]
        public Input<int>? Size { get; set; }

        [Input("status")]
        public Input<string>? Status { get; set; }

        [Input("updatedAt")]
        public Input<string>? UpdatedAt { get; set; }

        public VpcBandwidthState()
        {
        }
        public static new VpcBandwidthState Empty => new VpcBandwidthState();
    }
}
