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
    public sealed class GetVpnConnectionsConnectionIkepolicyResult
    {
        public readonly string AuthenticationAlgorithm;
        public readonly string AuthenticationMethod;
        public readonly string DhGroup;
        public readonly ImmutableArray<Outputs.GetVpnConnectionsConnectionIkepolicyDpdResult> Dpds;
        public readonly string EncryptionAlgorithm;
        public readonly string IkeVersion;
        public readonly int LifetimeSeconds;
        public readonly string LocalId;
        public readonly string LocalIdType;
        public readonly string PeerId;
        public readonly string PeerIdType;
        public readonly string Phase1NegotiationMode;

        [OutputConstructor]
        private GetVpnConnectionsConnectionIkepolicyResult(
            string authenticationAlgorithm,

            string authenticationMethod,

            string dhGroup,

            ImmutableArray<Outputs.GetVpnConnectionsConnectionIkepolicyDpdResult> dpds,

            string encryptionAlgorithm,

            string ikeVersion,

            int lifetimeSeconds,

            string localId,

            string localIdType,

            string peerId,

            string peerIdType,

            string phase1NegotiationMode)
        {
            AuthenticationAlgorithm = authenticationAlgorithm;
            AuthenticationMethod = authenticationMethod;
            DhGroup = dhGroup;
            Dpds = dpds;
            EncryptionAlgorithm = encryptionAlgorithm;
            IkeVersion = ikeVersion;
            LifetimeSeconds = lifetimeSeconds;
            LocalId = localId;
            LocalIdType = localIdType;
            PeerId = peerId;
            PeerIdType = peerIdType;
            Phase1NegotiationMode = phase1NegotiationMode;
        }
    }
}
