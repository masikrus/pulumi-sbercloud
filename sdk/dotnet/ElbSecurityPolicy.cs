// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Sbercloud
{
    [SbercloudResourceType("sbercloud:index/elbSecurityPolicy:ElbSecurityPolicy")]
    public partial class ElbSecurityPolicy : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Specifies the cipher suite list of the security policy.
        /// </summary>
        [Output("ciphers")]
        public Output<ImmutableArray<string>> Ciphers { get; private set; } = null!;

        [Output("createdAt")]
        public Output<string> CreatedAt { get; private set; } = null!;

        /// <summary>
        /// Specifies the description of the ELB security policy
        /// </summary>
        [Output("description")]
        public Output<string> Description { get; private set; } = null!;

        /// <summary>
        /// Specifies the enterprise project ID to which the Enterprise router belongs.
        /// </summary>
        [Output("enterpriseProjectId")]
        public Output<string> EnterpriseProjectId { get; private set; } = null!;

        /// <summary>
        /// The listener which the security policy associated with.
        /// </summary>
        [Output("listeners")]
        public Output<ImmutableArray<Outputs.ElbSecurityPolicyListener>> Listeners { get; private set; } = null!;

        /// <summary>
        /// Specifies the ELB security policy name.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Specifies the TSL protocol list which the security policy select.
        /// </summary>
        [Output("protocols")]
        public Output<ImmutableArray<string>> Protocols { get; private set; } = null!;

        [Output("region")]
        public Output<string> Region { get; private set; } = null!;

        [Output("updatedAt")]
        public Output<string> UpdatedAt { get; private set; } = null!;


        /// <summary>
        /// Create a ElbSecurityPolicy resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ElbSecurityPolicy(string name, ElbSecurityPolicyArgs args, CustomResourceOptions? options = null)
            : base("sbercloud:index/elbSecurityPolicy:ElbSecurityPolicy", name, args ?? new ElbSecurityPolicyArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ElbSecurityPolicy(string name, Input<string> id, ElbSecurityPolicyState? state = null, CustomResourceOptions? options = null)
            : base("sbercloud:index/elbSecurityPolicy:ElbSecurityPolicy", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing ElbSecurityPolicy resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ElbSecurityPolicy Get(string name, Input<string> id, ElbSecurityPolicyState? state = null, CustomResourceOptions? options = null)
        {
            return new ElbSecurityPolicy(name, id, state, options);
        }
    }

    public sealed class ElbSecurityPolicyArgs : global::Pulumi.ResourceArgs
    {
        [Input("ciphers", required: true)]
        private InputList<string>? _ciphers;

        /// <summary>
        /// Specifies the cipher suite list of the security policy.
        /// </summary>
        public InputList<string> Ciphers
        {
            get => _ciphers ?? (_ciphers = new InputList<string>());
            set => _ciphers = value;
        }

        /// <summary>
        /// Specifies the description of the ELB security policy
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Specifies the enterprise project ID to which the Enterprise router belongs.
        /// </summary>
        [Input("enterpriseProjectId")]
        public Input<string>? EnterpriseProjectId { get; set; }

        /// <summary>
        /// Specifies the ELB security policy name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("protocols", required: true)]
        private InputList<string>? _protocols;

        /// <summary>
        /// Specifies the TSL protocol list which the security policy select.
        /// </summary>
        public InputList<string> Protocols
        {
            get => _protocols ?? (_protocols = new InputList<string>());
            set => _protocols = value;
        }

        [Input("region")]
        public Input<string>? Region { get; set; }

        public ElbSecurityPolicyArgs()
        {
        }
        public static new ElbSecurityPolicyArgs Empty => new ElbSecurityPolicyArgs();
    }

    public sealed class ElbSecurityPolicyState : global::Pulumi.ResourceArgs
    {
        [Input("ciphers")]
        private InputList<string>? _ciphers;

        /// <summary>
        /// Specifies the cipher suite list of the security policy.
        /// </summary>
        public InputList<string> Ciphers
        {
            get => _ciphers ?? (_ciphers = new InputList<string>());
            set => _ciphers = value;
        }

        [Input("createdAt")]
        public Input<string>? CreatedAt { get; set; }

        /// <summary>
        /// Specifies the description of the ELB security policy
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Specifies the enterprise project ID to which the Enterprise router belongs.
        /// </summary>
        [Input("enterpriseProjectId")]
        public Input<string>? EnterpriseProjectId { get; set; }

        [Input("listeners")]
        private InputList<Inputs.ElbSecurityPolicyListenerGetArgs>? _listeners;

        /// <summary>
        /// The listener which the security policy associated with.
        /// </summary>
        public InputList<Inputs.ElbSecurityPolicyListenerGetArgs> Listeners
        {
            get => _listeners ?? (_listeners = new InputList<Inputs.ElbSecurityPolicyListenerGetArgs>());
            set => _listeners = value;
        }

        /// <summary>
        /// Specifies the ELB security policy name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("protocols")]
        private InputList<string>? _protocols;

        /// <summary>
        /// Specifies the TSL protocol list which the security policy select.
        /// </summary>
        public InputList<string> Protocols
        {
            get => _protocols ?? (_protocols = new InputList<string>());
            set => _protocols = value;
        }

        [Input("region")]
        public Input<string>? Region { get; set; }

        [Input("updatedAt")]
        public Input<string>? UpdatedAt { get; set; }

        public ElbSecurityPolicyState()
        {
        }
        public static new ElbSecurityPolicyState Empty => new ElbSecurityPolicyState();
    }
}
