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
    public sealed class ApigResponseRuleHeader
    {
        /// <summary>
        /// The key name of the response header.
        /// </summary>
        public readonly string Key;
        /// <summary>
        /// The value for the specified response header key.
        /// </summary>
        public readonly string Value;

        [OutputConstructor]
        private ApigResponseRuleHeader(
            string key,

            string value)
        {
            Key = key;
            Value = value;
        }
    }
}
