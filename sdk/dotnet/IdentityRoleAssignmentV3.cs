// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Sbercloud
{
    [SbercloudResourceType("sbercloud:index/identityRoleAssignmentV3:IdentityRoleAssignmentV3")]
    public partial class IdentityRoleAssignmentV3 : global::Pulumi.CustomResource
    {
        [Output("domainId")]
        public Output<string?> DomainId { get; private set; } = null!;

        [Output("enterpriseProjectId")]
        public Output<string?> EnterpriseProjectId { get; private set; } = null!;

        [Output("groupId")]
        public Output<string> GroupId { get; private set; } = null!;

        [Output("projectId")]
        public Output<string?> ProjectId { get; private set; } = null!;

        [Output("roleId")]
        public Output<string> RoleId { get; private set; } = null!;


        /// <summary>
        /// Create a IdentityRoleAssignmentV3 resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public IdentityRoleAssignmentV3(string name, IdentityRoleAssignmentV3Args args, CustomResourceOptions? options = null)
            : base("sbercloud:index/identityRoleAssignmentV3:IdentityRoleAssignmentV3", name, args ?? new IdentityRoleAssignmentV3Args(), MakeResourceOptions(options, ""))
        {
        }

        private IdentityRoleAssignmentV3(string name, Input<string> id, IdentityRoleAssignmentV3State? state = null, CustomResourceOptions? options = null)
            : base("sbercloud:index/identityRoleAssignmentV3:IdentityRoleAssignmentV3", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing IdentityRoleAssignmentV3 resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static IdentityRoleAssignmentV3 Get(string name, Input<string> id, IdentityRoleAssignmentV3State? state = null, CustomResourceOptions? options = null)
        {
            return new IdentityRoleAssignmentV3(name, id, state, options);
        }
    }

    public sealed class IdentityRoleAssignmentV3Args : global::Pulumi.ResourceArgs
    {
        [Input("domainId")]
        public Input<string>? DomainId { get; set; }

        [Input("enterpriseProjectId")]
        public Input<string>? EnterpriseProjectId { get; set; }

        [Input("groupId", required: true)]
        public Input<string> GroupId { get; set; } = null!;

        [Input("projectId")]
        public Input<string>? ProjectId { get; set; }

        [Input("roleId", required: true)]
        public Input<string> RoleId { get; set; } = null!;

        public IdentityRoleAssignmentV3Args()
        {
        }
        public static new IdentityRoleAssignmentV3Args Empty => new IdentityRoleAssignmentV3Args();
    }

    public sealed class IdentityRoleAssignmentV3State : global::Pulumi.ResourceArgs
    {
        [Input("domainId")]
        public Input<string>? DomainId { get; set; }

        [Input("enterpriseProjectId")]
        public Input<string>? EnterpriseProjectId { get; set; }

        [Input("groupId")]
        public Input<string>? GroupId { get; set; }

        [Input("projectId")]
        public Input<string>? ProjectId { get; set; }

        [Input("roleId")]
        public Input<string>? RoleId { get; set; }

        public IdentityRoleAssignmentV3State()
        {
        }
        public static new IdentityRoleAssignmentV3State Empty => new IdentityRoleAssignmentV3State();
    }
}
