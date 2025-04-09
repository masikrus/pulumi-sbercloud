// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Sbercloud
{
    [SbercloudResourceType("sbercloud:index/dmsKafkaUser:DmsKafkaUser")]
    public partial class DmsKafkaUser : global::Pulumi.CustomResource
    {
        [Output("createdAt")]
        public Output<string> CreatedAt { get; private set; } = null!;

        [Output("defaultApp")]
        public Output<bool> DefaultApp { get; private set; } = null!;

        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        [Output("instanceId")]
        public Output<string> InstanceId { get; private set; } = null!;

        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        [Output("password")]
        public Output<string> Password { get; private set; } = null!;

        [Output("region")]
        public Output<string> Region { get; private set; } = null!;

        [Output("role")]
        public Output<string> Role { get; private set; } = null!;


        /// <summary>
        /// Create a DmsKafkaUser resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public DmsKafkaUser(string name, DmsKafkaUserArgs args, CustomResourceOptions? options = null)
            : base("sbercloud:index/dmsKafkaUser:DmsKafkaUser", name, args ?? new DmsKafkaUserArgs(), MakeResourceOptions(options, ""))
        {
        }

        private DmsKafkaUser(string name, Input<string> id, DmsKafkaUserState? state = null, CustomResourceOptions? options = null)
            : base("sbercloud:index/dmsKafkaUser:DmsKafkaUser", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                AdditionalSecretOutputs =
                {
                    "password",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing DmsKafkaUser resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static DmsKafkaUser Get(string name, Input<string> id, DmsKafkaUserState? state = null, CustomResourceOptions? options = null)
        {
            return new DmsKafkaUser(name, id, state, options);
        }
    }

    public sealed class DmsKafkaUserArgs : global::Pulumi.ResourceArgs
    {
        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("instanceId", required: true)]
        public Input<string> InstanceId { get; set; } = null!;

        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("password", required: true)]
        private Input<string>? _password;
        public Input<string>? Password
        {
            get => _password;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _password = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        [Input("region")]
        public Input<string>? Region { get; set; }

        public DmsKafkaUserArgs()
        {
        }
        public static new DmsKafkaUserArgs Empty => new DmsKafkaUserArgs();
    }

    public sealed class DmsKafkaUserState : global::Pulumi.ResourceArgs
    {
        [Input("createdAt")]
        public Input<string>? CreatedAt { get; set; }

        [Input("defaultApp")]
        public Input<bool>? DefaultApp { get; set; }

        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("instanceId")]
        public Input<string>? InstanceId { get; set; }

        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("password")]
        private Input<string>? _password;
        public Input<string>? Password
        {
            get => _password;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _password = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        [Input("region")]
        public Input<string>? Region { get; set; }

        [Input("role")]
        public Input<string>? Role { get; set; }

        public DmsKafkaUserState()
        {
        }
        public static new DmsKafkaUserState Empty => new DmsKafkaUserState();
    }
}
