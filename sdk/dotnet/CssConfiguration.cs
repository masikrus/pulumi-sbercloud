// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Sbercloud
{
    [SbercloudResourceType("sbercloud:index/cssConfiguration:CssConfiguration")]
    public partial class CssConfiguration : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The CSS cluster ID.
        /// </summary>
        [Output("clusterId")]
        public Output<string> ClusterId { get; private set; } = null!;

        /// <summary>
        /// Whether to return the Access-Control-Allow-Credentials of the header during cross-domain access.
        /// </summary>
        [Output("httpCorsAllowCredetials")]
        public Output<string> HttpCorsAllowCredetials { get; private set; } = null!;

        /// <summary>
        /// Headers allowed for cross-domain access.
        /// </summary>
        [Output("httpCorsAllowHeaders")]
        public Output<string> HttpCorsAllowHeaders { get; private set; } = null!;

        /// <summary>
        /// Methods allowed for cross-domain access.
        /// </summary>
        [Output("httpCorsAllowMethods")]
        public Output<string> HttpCorsAllowMethods { get; private set; } = null!;

        /// <summary>
        /// Origin IP address allowed for cross-domain access, for example, **122.122.122.122:9200**.
        /// </summary>
        [Output("httpCorsAllowOrigin")]
        public Output<string> HttpCorsAllowOrigin { get; private set; } = null!;

        /// <summary>
        /// Whether to allow cross-domain access.
        /// </summary>
        [Output("httpCorsEnabled")]
        public Output<string> HttpCorsEnabled { get; private set; } = null!;

        /// <summary>
        /// Cache duration of the browser. The cache is automatically cleared after the time range you specify.
        /// </summary>
        [Output("httpCorsMaxAge")]
        public Output<string> HttpCorsMaxAge { get; private set; } = null!;

        /// <summary>
        /// Cache size in the query phase. Value range: **1** to **100**.
        /// </summary>
        [Output("indicesQueriesCacheSize")]
        public Output<string> IndicesQueriesCacheSize { get; private set; } = null!;

        [Output("region")]
        public Output<string> Region { get; private set; } = null!;

        /// <summary>
        /// Configured for migrating data from the current cluster to the target cluster through the reindex API.
        /// </summary>
        [Output("reindexRemoteWhitelist")]
        public Output<string> ReindexRemoteWhitelist { get; private set; } = null!;

        /// <summary>
        /// Queue size in the force merge thread pool.
        /// </summary>
        [Output("threadPoolForceMergeSize")]
        public Output<string> ThreadPoolForceMergeSize { get; private set; } = null!;


        /// <summary>
        /// Create a CssConfiguration resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public CssConfiguration(string name, CssConfigurationArgs args, CustomResourceOptions? options = null)
            : base("sbercloud:index/cssConfiguration:CssConfiguration", name, args ?? new CssConfigurationArgs(), MakeResourceOptions(options, ""))
        {
        }

        private CssConfiguration(string name, Input<string> id, CssConfigurationState? state = null, CustomResourceOptions? options = null)
            : base("sbercloud:index/cssConfiguration:CssConfiguration", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing CssConfiguration resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static CssConfiguration Get(string name, Input<string> id, CssConfigurationState? state = null, CustomResourceOptions? options = null)
        {
            return new CssConfiguration(name, id, state, options);
        }
    }

    public sealed class CssConfigurationArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The CSS cluster ID.
        /// </summary>
        [Input("clusterId", required: true)]
        public Input<string> ClusterId { get; set; } = null!;

        /// <summary>
        /// Whether to return the Access-Control-Allow-Credentials of the header during cross-domain access.
        /// </summary>
        [Input("httpCorsAllowCredetials")]
        public Input<string>? HttpCorsAllowCredetials { get; set; }

        /// <summary>
        /// Headers allowed for cross-domain access.
        /// </summary>
        [Input("httpCorsAllowHeaders")]
        public Input<string>? HttpCorsAllowHeaders { get; set; }

        /// <summary>
        /// Methods allowed for cross-domain access.
        /// </summary>
        [Input("httpCorsAllowMethods")]
        public Input<string>? HttpCorsAllowMethods { get; set; }

        /// <summary>
        /// Origin IP address allowed for cross-domain access, for example, **122.122.122.122:9200**.
        /// </summary>
        [Input("httpCorsAllowOrigin")]
        public Input<string>? HttpCorsAllowOrigin { get; set; }

        /// <summary>
        /// Whether to allow cross-domain access.
        /// </summary>
        [Input("httpCorsEnabled")]
        public Input<string>? HttpCorsEnabled { get; set; }

        /// <summary>
        /// Cache duration of the browser. The cache is automatically cleared after the time range you specify.
        /// </summary>
        [Input("httpCorsMaxAge")]
        public Input<string>? HttpCorsMaxAge { get; set; }

        /// <summary>
        /// Cache size in the query phase. Value range: **1** to **100**.
        /// </summary>
        [Input("indicesQueriesCacheSize")]
        public Input<string>? IndicesQueriesCacheSize { get; set; }

        [Input("region")]
        public Input<string>? Region { get; set; }

        /// <summary>
        /// Configured for migrating data from the current cluster to the target cluster through the reindex API.
        /// </summary>
        [Input("reindexRemoteWhitelist")]
        public Input<string>? ReindexRemoteWhitelist { get; set; }

        /// <summary>
        /// Queue size in the force merge thread pool.
        /// </summary>
        [Input("threadPoolForceMergeSize")]
        public Input<string>? ThreadPoolForceMergeSize { get; set; }

        public CssConfigurationArgs()
        {
        }
        public static new CssConfigurationArgs Empty => new CssConfigurationArgs();
    }

    public sealed class CssConfigurationState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The CSS cluster ID.
        /// </summary>
        [Input("clusterId")]
        public Input<string>? ClusterId { get; set; }

        /// <summary>
        /// Whether to return the Access-Control-Allow-Credentials of the header during cross-domain access.
        /// </summary>
        [Input("httpCorsAllowCredetials")]
        public Input<string>? HttpCorsAllowCredetials { get; set; }

        /// <summary>
        /// Headers allowed for cross-domain access.
        /// </summary>
        [Input("httpCorsAllowHeaders")]
        public Input<string>? HttpCorsAllowHeaders { get; set; }

        /// <summary>
        /// Methods allowed for cross-domain access.
        /// </summary>
        [Input("httpCorsAllowMethods")]
        public Input<string>? HttpCorsAllowMethods { get; set; }

        /// <summary>
        /// Origin IP address allowed for cross-domain access, for example, **122.122.122.122:9200**.
        /// </summary>
        [Input("httpCorsAllowOrigin")]
        public Input<string>? HttpCorsAllowOrigin { get; set; }

        /// <summary>
        /// Whether to allow cross-domain access.
        /// </summary>
        [Input("httpCorsEnabled")]
        public Input<string>? HttpCorsEnabled { get; set; }

        /// <summary>
        /// Cache duration of the browser. The cache is automatically cleared after the time range you specify.
        /// </summary>
        [Input("httpCorsMaxAge")]
        public Input<string>? HttpCorsMaxAge { get; set; }

        /// <summary>
        /// Cache size in the query phase. Value range: **1** to **100**.
        /// </summary>
        [Input("indicesQueriesCacheSize")]
        public Input<string>? IndicesQueriesCacheSize { get; set; }

        [Input("region")]
        public Input<string>? Region { get; set; }

        /// <summary>
        /// Configured for migrating data from the current cluster to the target cluster through the reindex API.
        /// </summary>
        [Input("reindexRemoteWhitelist")]
        public Input<string>? ReindexRemoteWhitelist { get; set; }

        /// <summary>
        /// Queue size in the force merge thread pool.
        /// </summary>
        [Input("threadPoolForceMergeSize")]
        public Input<string>? ThreadPoolForceMergeSize { get; set; }

        public CssConfigurationState()
        {
        }
        public static new CssConfigurationState Empty => new CssConfigurationState();
    }
}
