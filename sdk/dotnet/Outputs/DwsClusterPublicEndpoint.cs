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
    public sealed class DwsClusterPublicEndpoint
    {
        /// <summary>
        /// JDBC URL. Format: jdbc:postgresql://&lt;public_connect_info&gt;/&lt;YOUR_DATABASE_NAME&gt;
        /// </summary>
        public readonly string? JdbcUrl;
        /// <summary>
        /// Public network connection information.
        /// </summary>
        public readonly string? PublicConnectInfo;

        [OutputConstructor]
        private DwsClusterPublicEndpoint(
            string? jdbcUrl,

            string? publicConnectInfo)
        {
            JdbcUrl = jdbcUrl;
            PublicConnectInfo = publicConnectInfo;
        }
    }
}
