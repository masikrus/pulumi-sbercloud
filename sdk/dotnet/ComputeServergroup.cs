// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Sbercloud
{
    [SbercloudResourceType("sbercloud:index/computeServergroup:ComputeServergroup")]
    public partial class ComputeServergroup : global::Pulumi.CustomResource
    {
        /// <summary>
        /// schema: Internal
        /// </summary>
        [Output("faultDomains")]
        public Output<ImmutableArray<string>> FaultDomains { get; private set; } = null!;

        [Output("members")]
        public Output<ImmutableArray<string>> Members { get; private set; } = null!;

        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// schema: Required
        /// </summary>
        [Output("policies")]
        public Output<ImmutableArray<string>> Policies { get; private set; } = null!;

        [Output("region")]
        public Output<string> Region { get; private set; } = null!;


        /// <summary>
        /// Create a ComputeServergroup resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ComputeServergroup(string name, ComputeServergroupArgs? args = null, CustomResourceOptions? options = null)
            : base("sbercloud:index/computeServergroup:ComputeServergroup", name, args ?? new ComputeServergroupArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ComputeServergroup(string name, Input<string> id, ComputeServergroupState? state = null, CustomResourceOptions? options = null)
            : base("sbercloud:index/computeServergroup:ComputeServergroup", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing ComputeServergroup resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ComputeServergroup Get(string name, Input<string> id, ComputeServergroupState? state = null, CustomResourceOptions? options = null)
        {
            return new ComputeServergroup(name, id, state, options);
        }
    }

    public sealed class ComputeServergroupArgs : global::Pulumi.ResourceArgs
    {
        [Input("members")]
        private InputList<string>? _members;
        public InputList<string> Members
        {
            get => _members ?? (_members = new InputList<string>());
            set => _members = value;
        }

        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("policies")]
        private InputList<string>? _policies;

        /// <summary>
        /// schema: Required
        /// </summary>
        public InputList<string> Policies
        {
            get => _policies ?? (_policies = new InputList<string>());
            set => _policies = value;
        }

        [Input("region")]
        public Input<string>? Region { get; set; }

        public ComputeServergroupArgs()
        {
        }
        public static new ComputeServergroupArgs Empty => new ComputeServergroupArgs();
    }

    public sealed class ComputeServergroupState : global::Pulumi.ResourceArgs
    {
        [Input("faultDomains")]
        private InputList<string>? _faultDomains;

        /// <summary>
        /// schema: Internal
        /// </summary>
        public InputList<string> FaultDomains
        {
            get => _faultDomains ?? (_faultDomains = new InputList<string>());
            set => _faultDomains = value;
        }

        [Input("members")]
        private InputList<string>? _members;
        public InputList<string> Members
        {
            get => _members ?? (_members = new InputList<string>());
            set => _members = value;
        }

        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("policies")]
        private InputList<string>? _policies;

        /// <summary>
        /// schema: Required
        /// </summary>
        public InputList<string> Policies
        {
            get => _policies ?? (_policies = new InputList<string>());
            set => _policies = value;
        }

        [Input("region")]
        public Input<string>? Region { get; set; }

        public ComputeServergroupState()
        {
        }
        public static new ComputeServergroupState Empty => new ComputeServergroupState();
    }
}
