// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Sbercloud
{
    [SbercloudResourceType("sbercloud:index/networkingEipAssociate:NetworkingEipAssociate")]
    public partial class NetworkingEipAssociate : global::Pulumi.CustomResource
    {
        [Output("fixedIp")]
        public Output<string> FixedIp { get; private set; } = null!;

        [Output("macAddress")]
        public Output<string> MacAddress { get; private set; } = null!;

        [Output("networkId")]
        public Output<string> NetworkId { get; private set; } = null!;

        [Output("portId")]
        public Output<string> PortId { get; private set; } = null!;

        [Output("publicIp")]
        public Output<string> PublicIp { get; private set; } = null!;

        [Output("publicIpv6")]
        public Output<string> PublicIpv6 { get; private set; } = null!;

        [Output("region")]
        public Output<string> Region { get; private set; } = null!;

        [Output("status")]
        public Output<string> Status { get; private set; } = null!;


        /// <summary>
        /// Create a NetworkingEipAssociate resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public NetworkingEipAssociate(string name, NetworkingEipAssociateArgs args, CustomResourceOptions? options = null)
            : base("sbercloud:index/networkingEipAssociate:NetworkingEipAssociate", name, args ?? new NetworkingEipAssociateArgs(), MakeResourceOptions(options, ""))
        {
        }

        private NetworkingEipAssociate(string name, Input<string> id, NetworkingEipAssociateState? state = null, CustomResourceOptions? options = null)
            : base("sbercloud:index/networkingEipAssociate:NetworkingEipAssociate", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing NetworkingEipAssociate resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static NetworkingEipAssociate Get(string name, Input<string> id, NetworkingEipAssociateState? state = null, CustomResourceOptions? options = null)
        {
            return new NetworkingEipAssociate(name, id, state, options);
        }
    }

    public sealed class NetworkingEipAssociateArgs : global::Pulumi.ResourceArgs
    {
        [Input("fixedIp")]
        public Input<string>? FixedIp { get; set; }

        [Input("networkId")]
        public Input<string>? NetworkId { get; set; }

        [Input("portId")]
        public Input<string>? PortId { get; set; }

        [Input("publicIp", required: true)]
        public Input<string> PublicIp { get; set; } = null!;

        [Input("region")]
        public Input<string>? Region { get; set; }

        public NetworkingEipAssociateArgs()
        {
        }
        public static new NetworkingEipAssociateArgs Empty => new NetworkingEipAssociateArgs();
    }

    public sealed class NetworkingEipAssociateState : global::Pulumi.ResourceArgs
    {
        [Input("fixedIp")]
        public Input<string>? FixedIp { get; set; }

        [Input("macAddress")]
        public Input<string>? MacAddress { get; set; }

        [Input("networkId")]
        public Input<string>? NetworkId { get; set; }

        [Input("portId")]
        public Input<string>? PortId { get; set; }

        [Input("publicIp")]
        public Input<string>? PublicIp { get; set; }

        [Input("publicIpv6")]
        public Input<string>? PublicIpv6 { get; set; }

        [Input("region")]
        public Input<string>? Region { get; set; }

        [Input("status")]
        public Input<string>? Status { get; set; }

        public NetworkingEipAssociateState()
        {
        }
        public static new NetworkingEipAssociateState Empty => new NetworkingEipAssociateState();
    }
}
