// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Sbercloud
{
    [SbercloudResourceType("sbercloud:index/lbMonitor:LbMonitor")]
    public partial class LbMonitor : global::Pulumi.CustomResource
    {
        [Output("adminStateUp")]
        public Output<bool?> AdminStateUp { get; private set; } = null!;

        [Output("delay")]
        public Output<int> Delay { get; private set; } = null!;

        [Output("expectedCodes")]
        public Output<string> ExpectedCodes { get; private set; } = null!;

        [Output("httpMethod")]
        public Output<string> HttpMethod { get; private set; } = null!;

        [Output("maxRetries")]
        public Output<int> MaxRetries { get; private set; } = null!;

        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        [Output("poolId")]
        public Output<string> PoolId { get; private set; } = null!;

        [Output("port")]
        public Output<int?> Port { get; private set; } = null!;

        [Output("region")]
        public Output<string> Region { get; private set; } = null!;

        [Output("tenantId")]
        public Output<string> TenantId { get; private set; } = null!;

        [Output("timeout")]
        public Output<int> Timeout { get; private set; } = null!;

        [Output("type")]
        public Output<string> Type { get; private set; } = null!;

        [Output("urlPath")]
        public Output<string> UrlPath { get; private set; } = null!;


        /// <summary>
        /// Create a LbMonitor resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public LbMonitor(string name, LbMonitorArgs args, CustomResourceOptions? options = null)
            : base("sbercloud:index/lbMonitor:LbMonitor", name, args ?? new LbMonitorArgs(), MakeResourceOptions(options, ""))
        {
        }

        private LbMonitor(string name, Input<string> id, LbMonitorState? state = null, CustomResourceOptions? options = null)
            : base("sbercloud:index/lbMonitor:LbMonitor", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing LbMonitor resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static LbMonitor Get(string name, Input<string> id, LbMonitorState? state = null, CustomResourceOptions? options = null)
        {
            return new LbMonitor(name, id, state, options);
        }
    }

    public sealed class LbMonitorArgs : global::Pulumi.ResourceArgs
    {
        [Input("adminStateUp")]
        public Input<bool>? AdminStateUp { get; set; }

        [Input("delay", required: true)]
        public Input<int> Delay { get; set; } = null!;

        [Input("expectedCodes")]
        public Input<string>? ExpectedCodes { get; set; }

        [Input("httpMethod")]
        public Input<string>? HttpMethod { get; set; }

        [Input("maxRetries", required: true)]
        public Input<int> MaxRetries { get; set; } = null!;

        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("poolId", required: true)]
        public Input<string> PoolId { get; set; } = null!;

        [Input("port")]
        public Input<int>? Port { get; set; }

        [Input("region")]
        public Input<string>? Region { get; set; }

        [Input("tenantId")]
        public Input<string>? TenantId { get; set; }

        [Input("timeout", required: true)]
        public Input<int> Timeout { get; set; } = null!;

        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        [Input("urlPath")]
        public Input<string>? UrlPath { get; set; }

        public LbMonitorArgs()
        {
        }
        public static new LbMonitorArgs Empty => new LbMonitorArgs();
    }

    public sealed class LbMonitorState : global::Pulumi.ResourceArgs
    {
        [Input("adminStateUp")]
        public Input<bool>? AdminStateUp { get; set; }

        [Input("delay")]
        public Input<int>? Delay { get; set; }

        [Input("expectedCodes")]
        public Input<string>? ExpectedCodes { get; set; }

        [Input("httpMethod")]
        public Input<string>? HttpMethod { get; set; }

        [Input("maxRetries")]
        public Input<int>? MaxRetries { get; set; }

        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("poolId")]
        public Input<string>? PoolId { get; set; }

        [Input("port")]
        public Input<int>? Port { get; set; }

        [Input("region")]
        public Input<string>? Region { get; set; }

        [Input("tenantId")]
        public Input<string>? TenantId { get; set; }

        [Input("timeout")]
        public Input<int>? Timeout { get; set; }

        [Input("type")]
        public Input<string>? Type { get; set; }

        [Input("urlPath")]
        public Input<string>? UrlPath { get; set; }

        public LbMonitorState()
        {
        }
        public static new LbMonitorState Empty => new LbMonitorState();
    }
}
