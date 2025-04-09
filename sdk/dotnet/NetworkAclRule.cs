// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Sbercloud
{
    [SbercloudResourceType("sbercloud:index/networkAclRule:NetworkAclRule")]
    public partial class NetworkAclRule : global::Pulumi.CustomResource
    {
        [Output("action")]
        public Output<string> Action { get; private set; } = null!;

        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        [Output("destinationIpAddress")]
        public Output<string?> DestinationIpAddress { get; private set; } = null!;

        [Output("destinationPort")]
        public Output<string?> DestinationPort { get; private set; } = null!;

        [Output("enabled")]
        public Output<bool?> Enabled { get; private set; } = null!;

        [Output("ipVersion")]
        public Output<int?> IpVersion { get; private set; } = null!;

        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        [Output("protocol")]
        public Output<string> Protocol { get; private set; } = null!;

        [Output("region")]
        public Output<string> Region { get; private set; } = null!;

        [Output("sourceIpAddress")]
        public Output<string?> SourceIpAddress { get; private set; } = null!;

        [Output("sourcePort")]
        public Output<string?> SourcePort { get; private set; } = null!;


        /// <summary>
        /// Create a NetworkAclRule resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public NetworkAclRule(string name, NetworkAclRuleArgs args, CustomResourceOptions? options = null)
            : base("sbercloud:index/networkAclRule:NetworkAclRule", name, args ?? new NetworkAclRuleArgs(), MakeResourceOptions(options, ""))
        {
        }

        private NetworkAclRule(string name, Input<string> id, NetworkAclRuleState? state = null, CustomResourceOptions? options = null)
            : base("sbercloud:index/networkAclRule:NetworkAclRule", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing NetworkAclRule resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static NetworkAclRule Get(string name, Input<string> id, NetworkAclRuleState? state = null, CustomResourceOptions? options = null)
        {
            return new NetworkAclRule(name, id, state, options);
        }
    }

    public sealed class NetworkAclRuleArgs : global::Pulumi.ResourceArgs
    {
        [Input("action", required: true)]
        public Input<string> Action { get; set; } = null!;

        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("destinationIpAddress")]
        public Input<string>? DestinationIpAddress { get; set; }

        [Input("destinationPort")]
        public Input<string>? DestinationPort { get; set; }

        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        [Input("ipVersion")]
        public Input<int>? IpVersion { get; set; }

        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("protocol", required: true)]
        public Input<string> Protocol { get; set; } = null!;

        [Input("region")]
        public Input<string>? Region { get; set; }

        [Input("sourceIpAddress")]
        public Input<string>? SourceIpAddress { get; set; }

        [Input("sourcePort")]
        public Input<string>? SourcePort { get; set; }

        public NetworkAclRuleArgs()
        {
        }
        public static new NetworkAclRuleArgs Empty => new NetworkAclRuleArgs();
    }

    public sealed class NetworkAclRuleState : global::Pulumi.ResourceArgs
    {
        [Input("action")]
        public Input<string>? Action { get; set; }

        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("destinationIpAddress")]
        public Input<string>? DestinationIpAddress { get; set; }

        [Input("destinationPort")]
        public Input<string>? DestinationPort { get; set; }

        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        [Input("ipVersion")]
        public Input<int>? IpVersion { get; set; }

        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("protocol")]
        public Input<string>? Protocol { get; set; }

        [Input("region")]
        public Input<string>? Region { get; set; }

        [Input("sourceIpAddress")]
        public Input<string>? SourceIpAddress { get; set; }

        [Input("sourcePort")]
        public Input<string>? SourcePort { get; set; }

        public NetworkAclRuleState()
        {
        }
        public static new NetworkAclRuleState Empty => new NetworkAclRuleState();
    }
}
