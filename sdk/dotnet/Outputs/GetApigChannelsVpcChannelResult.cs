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
    public sealed class GetApigChannelsVpcChannelResult
    {
        /// <summary>
        /// The distribution algorithm.
        /// </summary>
        public readonly int BalanceStrategy;
        /// <summary>
        /// The creation time of channel, in RFC3339 format.
        /// </summary>
        public readonly string CreatedAt;
        /// <summary>
        /// The ID of the VPC channel.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The parameter member groups of the VPC channels.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetApigChannelsVpcChannelMemberGroupResult> MemberGroups;
        /// <summary>
        /// The member type of the VPC channel.
        /// </summary>
        public readonly string MemberType;
        /// <summary>
        /// The name of the VPC channel.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The port of the backend server.
        /// </summary>
        public readonly int Port;
        /// <summary>
        /// The type of the VPC channel.
        /// </summary>
        public readonly int Type;

        [OutputConstructor]
        private GetApigChannelsVpcChannelResult(
            int balanceStrategy,

            string createdAt,

            string id,

            ImmutableArray<Outputs.GetApigChannelsVpcChannelMemberGroupResult> memberGroups,

            string memberType,

            string name,

            int port,

            int type)
        {
            BalanceStrategy = balanceStrategy;
            CreatedAt = createdAt;
            Id = id;
            MemberGroups = memberGroups;
            MemberType = memberType;
            Name = name;
            Port = port;
            Type = type;
        }
    }
}
