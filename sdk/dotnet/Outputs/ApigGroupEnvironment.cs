// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Sbercloud.Outputs
{

    [OutputType]
    public sealed class ApigGroupEnvironment
    {
        /// <summary>
        /// The ID of the environment to which the variables belongs.
        /// </summary>
        public readonly string EnvironmentId;
        /// <summary>
        /// The array of one or more environment variables.
        /// </summary>
        public readonly ImmutableArray<Outputs.ApigGroupEnvironmentVariable> Variables;

        [OutputConstructor]
        private ApigGroupEnvironment(
            string environmentId,

            ImmutableArray<Outputs.ApigGroupEnvironmentVariable> variables)
        {
            EnvironmentId = environmentId;
            Variables = variables;
        }
    }
}
