// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Sbercloud.Inputs
{

    public sealed class CceClusterCertificateClusterGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("certificateAuthorityData")]
        public Input<string>? CertificateAuthorityData { get; set; }

        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("server")]
        public Input<string>? Server { get; set; }

        public CceClusterCertificateClusterGetArgs()
        {
        }
        public static new CceClusterCertificateClusterGetArgs Empty => new CceClusterCertificateClusterGetArgs();
    }
}
