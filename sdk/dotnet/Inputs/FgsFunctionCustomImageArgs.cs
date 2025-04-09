// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Sbercloud.Inputs
{

    public sealed class FgsFunctionCustomImageArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The command line arguments used to start the SWR image.
        /// </summary>
        [Input("args")]
        public Input<string>? Args { get; set; }

        /// <summary>
        /// The startup commands of the SWR image.
        /// </summary>
        [Input("command")]
        public Input<string>? Command { get; set; }

        /// <summary>
        /// The URL of SWR image.
        /// </summary>
        [Input("url", required: true)]
        public Input<string> Url { get; set; } = null!;

        /// <summary>
        /// The user group ID that used to run SWR image.
        /// </summary>
        [Input("userGroupId")]
        public Input<string>? UserGroupId { get; set; }

        /// <summary>
        /// The user ID that used to run SWR image.
        /// </summary>
        [Input("userId")]
        public Input<string>? UserId { get; set; }

        /// <summary>
        /// The working directory of the SWR image.
        /// </summary>
        [Input("workingDir")]
        public Input<string>? WorkingDir { get; set; }

        public FgsFunctionCustomImageArgs()
        {
        }
        public static new FgsFunctionCustomImageArgs Empty => new FgsFunctionCustomImageArgs();
    }
}
