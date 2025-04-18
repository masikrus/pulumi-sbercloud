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
    public sealed class GetCceClusterCertificateClusterResult
    {
        public readonly string CertificateAuthorityData;
        public readonly bool InsecureSkipTlsVerify;
        public readonly string Name;
        public readonly string Server;

        [OutputConstructor]
        private GetCceClusterCertificateClusterResult(
            string certificateAuthorityData,

            bool insecureSkipTlsVerify,

            string name,

            string server)
        {
            CertificateAuthorityData = certificateAuthorityData;
            InsecureSkipTlsVerify = insecureSkipTlsVerify;
            Name = name;
            Server = server;
        }
    }
}
