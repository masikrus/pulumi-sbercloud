// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Sbercloud.Inputs
{

    public sealed class DwsClusterEndpointArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Private network connection information.
        /// </summary>
        [Input("connectInfo")]
        public Input<string>? ConnectInfo { get; set; }

        /// <summary>
        /// JDBC URL. Format: jdbc:postgresql://&lt;connect_info&gt;/&lt;YOUR_DATABASE_NAME&gt;
        /// </summary>
        [Input("jdbcUrl")]
        public Input<string>? JdbcUrl { get; set; }

        public DwsClusterEndpointArgs()
        {
        }
        public static new DwsClusterEndpointArgs Empty => new DwsClusterEndpointArgs();
    }
}
