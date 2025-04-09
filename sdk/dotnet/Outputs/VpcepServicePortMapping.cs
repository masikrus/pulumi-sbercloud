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
    public sealed class VpcepServicePortMapping
    {
        public readonly string? Protocol;
        /// <summary>
        /// schema: Required
        /// </summary>
        public readonly int? ServicePort;
        /// <summary>
        /// schema: Required
        /// </summary>
        public readonly int? TerminalPort;

        [OutputConstructor]
        private VpcepServicePortMapping(
            string? protocol,

            int? servicePort,

            int? terminalPort)
        {
            Protocol = protocol;
            ServicePort = servicePort;
            TerminalPort = terminalPort;
        }
    }
}
