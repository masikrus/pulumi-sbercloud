// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Sbercloud
{
    [SbercloudResourceType("sbercloud:index/identityGroupMembership:IdentityGroupMembership")]
    public partial class IdentityGroupMembership : global::Pulumi.CustomResource
    {
        [Output("group")]
        public Output<string> Group { get; private set; } = null!;

        [Output("users")]
        public Output<ImmutableArray<string>> Users { get; private set; } = null!;


        /// <summary>
        /// Create a IdentityGroupMembership resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public IdentityGroupMembership(string name, IdentityGroupMembershipArgs args, CustomResourceOptions? options = null)
            : base("sbercloud:index/identityGroupMembership:IdentityGroupMembership", name, args ?? new IdentityGroupMembershipArgs(), MakeResourceOptions(options, ""))
        {
        }

        private IdentityGroupMembership(string name, Input<string> id, IdentityGroupMembershipState? state = null, CustomResourceOptions? options = null)
            : base("sbercloud:index/identityGroupMembership:IdentityGroupMembership", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing IdentityGroupMembership resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static IdentityGroupMembership Get(string name, Input<string> id, IdentityGroupMembershipState? state = null, CustomResourceOptions? options = null)
        {
            return new IdentityGroupMembership(name, id, state, options);
        }
    }

    public sealed class IdentityGroupMembershipArgs : global::Pulumi.ResourceArgs
    {
        [Input("group", required: true)]
        public Input<string> Group { get; set; } = null!;

        [Input("users", required: true)]
        private InputList<string>? _users;
        public InputList<string> Users
        {
            get => _users ?? (_users = new InputList<string>());
            set => _users = value;
        }

        public IdentityGroupMembershipArgs()
        {
        }
        public static new IdentityGroupMembershipArgs Empty => new IdentityGroupMembershipArgs();
    }

    public sealed class IdentityGroupMembershipState : global::Pulumi.ResourceArgs
    {
        [Input("group")]
        public Input<string>? Group { get; set; }

        [Input("users")]
        private InputList<string>? _users;
        public InputList<string> Users
        {
            get => _users ?? (_users = new InputList<string>());
            set => _users = value;
        }

        public IdentityGroupMembershipState()
        {
        }
        public static new IdentityGroupMembershipState Empty => new IdentityGroupMembershipState();
    }
}
