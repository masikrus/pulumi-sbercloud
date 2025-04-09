// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Sbercloud
{
    [SbercloudResourceType("sbercloud:index/identityGroupMembershipV3:IdentityGroupMembershipV3")]
    public partial class IdentityGroupMembershipV3 : global::Pulumi.CustomResource
    {
        [Output("group")]
        public Output<string> Group { get; private set; } = null!;

        [Output("users")]
        public Output<ImmutableArray<string>> Users { get; private set; } = null!;


        /// <summary>
        /// Create a IdentityGroupMembershipV3 resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public IdentityGroupMembershipV3(string name, IdentityGroupMembershipV3Args args, CustomResourceOptions? options = null)
            : base("sbercloud:index/identityGroupMembershipV3:IdentityGroupMembershipV3", name, args ?? new IdentityGroupMembershipV3Args(), MakeResourceOptions(options, ""))
        {
        }

        private IdentityGroupMembershipV3(string name, Input<string> id, IdentityGroupMembershipV3State? state = null, CustomResourceOptions? options = null)
            : base("sbercloud:index/identityGroupMembershipV3:IdentityGroupMembershipV3", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing IdentityGroupMembershipV3 resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static IdentityGroupMembershipV3 Get(string name, Input<string> id, IdentityGroupMembershipV3State? state = null, CustomResourceOptions? options = null)
        {
            return new IdentityGroupMembershipV3(name, id, state, options);
        }
    }

    public sealed class IdentityGroupMembershipV3Args : global::Pulumi.ResourceArgs
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

        public IdentityGroupMembershipV3Args()
        {
        }
        public static new IdentityGroupMembershipV3Args Empty => new IdentityGroupMembershipV3Args();
    }

    public sealed class IdentityGroupMembershipV3State : global::Pulumi.ResourceArgs
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

        public IdentityGroupMembershipV3State()
        {
        }
        public static new IdentityGroupMembershipV3State Empty => new IdentityGroupMembershipV3State();
    }
}
