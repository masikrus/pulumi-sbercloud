// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Sbercloud.Inputs
{

    public sealed class DwsClusterPublicEndpointGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// JDBC URL. Format: jdbc:postgresql://&lt;public_connect_info&gt;/&lt;YOUR_DATABASE_NAME&gt;
        /// </summary>
        [Input("jdbcUrl")]
        public Input<string>? JdbcUrl { get; set; }

        /// <summary>
        /// Public network connection information.
        /// </summary>
        [Input("publicConnectInfo")]
        public Input<string>? PublicConnectInfo { get; set; }

        public DwsClusterPublicEndpointGetArgs()
        {
        }
        public static new DwsClusterPublicEndpointGetArgs Empty => new DwsClusterPublicEndpointGetArgs();
    }
}
