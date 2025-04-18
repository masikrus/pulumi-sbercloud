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
    public sealed class DwsClusterEndpoint
    {
        /// <summary>
        /// Private network connection information.
        /// </summary>
        public readonly string? ConnectInfo;
        /// <summary>
        /// JDBC URL. Format: jdbc:postgresql://&lt;connect_info&gt;/&lt;YOUR_DATABASE_NAME&gt;
        /// </summary>
        public readonly string? JdbcUrl;

        [OutputConstructor]
        private DwsClusterEndpoint(
            string? connectInfo,

            string? jdbcUrl)
        {
            ConnectInfo = connectInfo;
            JdbcUrl = jdbcUrl;
        }
    }
}
