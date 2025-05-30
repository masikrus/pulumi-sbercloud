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
    public sealed class ElbL7policyRedirectUrlConfig
    {
        public readonly string? Host;
        public readonly Outputs.ElbL7policyRedirectUrlConfigInsertHeadersConfig? InsertHeadersConfig;
        public readonly string? Path;
        public readonly string? Port;
        public readonly string? Protocol;
        public readonly string? Query;
        public readonly Outputs.ElbL7policyRedirectUrlConfigRemoveHeadersConfig? RemoveHeadersConfig;
        public readonly string StatusCode;

        [OutputConstructor]
        private ElbL7policyRedirectUrlConfig(
            string? host,

            Outputs.ElbL7policyRedirectUrlConfigInsertHeadersConfig? insertHeadersConfig,

            string? path,

            string? port,

            string? protocol,

            string? query,

            Outputs.ElbL7policyRedirectUrlConfigRemoveHeadersConfig? removeHeadersConfig,

            string statusCode)
        {
            Host = host;
            InsertHeadersConfig = insertHeadersConfig;
            Path = path;
            Port = port;
            Protocol = protocol;
            Query = query;
            RemoveHeadersConfig = removeHeadersConfig;
            StatusCode = statusCode;
        }
    }
}
