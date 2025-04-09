// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Sbercloud
{
    [SbercloudResourceType("sbercloud:index/rdsMysqlBinlog:RdsMysqlBinlog")]
    public partial class RdsMysqlBinlog : global::Pulumi.CustomResource
    {
        [Output("binlogRetentionHours")]
        public Output<int> BinlogRetentionHours { get; private set; } = null!;

        [Output("instanceId")]
        public Output<string> InstanceId { get; private set; } = null!;

        [Output("region")]
        public Output<string> Region { get; private set; } = null!;


        /// <summary>
        /// Create a RdsMysqlBinlog resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public RdsMysqlBinlog(string name, RdsMysqlBinlogArgs args, CustomResourceOptions? options = null)
            : base("sbercloud:index/rdsMysqlBinlog:RdsMysqlBinlog", name, args ?? new RdsMysqlBinlogArgs(), MakeResourceOptions(options, ""))
        {
        }

        private RdsMysqlBinlog(string name, Input<string> id, RdsMysqlBinlogState? state = null, CustomResourceOptions? options = null)
            : base("sbercloud:index/rdsMysqlBinlog:RdsMysqlBinlog", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing RdsMysqlBinlog resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static RdsMysqlBinlog Get(string name, Input<string> id, RdsMysqlBinlogState? state = null, CustomResourceOptions? options = null)
        {
            return new RdsMysqlBinlog(name, id, state, options);
        }
    }

    public sealed class RdsMysqlBinlogArgs : global::Pulumi.ResourceArgs
    {
        [Input("binlogRetentionHours", required: true)]
        public Input<int> BinlogRetentionHours { get; set; } = null!;

        [Input("instanceId", required: true)]
        public Input<string> InstanceId { get; set; } = null!;

        [Input("region")]
        public Input<string>? Region { get; set; }

        public RdsMysqlBinlogArgs()
        {
        }
        public static new RdsMysqlBinlogArgs Empty => new RdsMysqlBinlogArgs();
    }

    public sealed class RdsMysqlBinlogState : global::Pulumi.ResourceArgs
    {
        [Input("binlogRetentionHours")]
        public Input<int>? BinlogRetentionHours { get; set; }

        [Input("instanceId")]
        public Input<string>? InstanceId { get; set; }

        [Input("region")]
        public Input<string>? Region { get; set; }

        public RdsMysqlBinlogState()
        {
        }
        public static new RdsMysqlBinlogState Empty => new RdsMysqlBinlogState();
    }
}
