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
    public sealed class GetLbPoolsPoolResult
    {
        /// <summary>
        /// The description of pool.
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// Specifies the health monitor ID of the LB pool.
        /// </summary>
        public readonly string HealthmonitorId;
        /// <summary>
        /// The pool ID.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The method of the LB pool.
        /// </summary>
        public readonly string LbMethod;
        /// <summary>
        /// Listener list. For details, see Data structure of the listener field.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetLbPoolsPoolListenerResult> Listeners;
        /// <summary>
        /// Loadbalancer list. For details, see Data structure of the loadbalancer field.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetLbPoolsPoolLoadbalancerResult> Loadbalancers;
        /// <summary>
        /// Loadbalancer list. For details, see Data structure of the members field.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetLbPoolsPoolMemberResult> Members;
        /// <summary>
        /// The pool name.
        /// </summary>
        public readonly string Name;
        public readonly ImmutableArray<Outputs.GetLbPoolsPoolPersistenceResult> Persistences;
        /// <summary>
        /// The protocol of pool.
        /// </summary>
        public readonly string Protocol;

        [OutputConstructor]
        private GetLbPoolsPoolResult(
            string description,

            string healthmonitorId,

            string id,

            string lbMethod,

            ImmutableArray<Outputs.GetLbPoolsPoolListenerResult> listeners,

            ImmutableArray<Outputs.GetLbPoolsPoolLoadbalancerResult> loadbalancers,

            ImmutableArray<Outputs.GetLbPoolsPoolMemberResult> members,

            string name,

            ImmutableArray<Outputs.GetLbPoolsPoolPersistenceResult> persistences,

            string protocol)
        {
            Description = description;
            HealthmonitorId = healthmonitorId;
            Id = id;
            LbMethod = lbMethod;
            Listeners = listeners;
            Loadbalancers = loadbalancers;
            Members = members;
            Name = name;
            Persistences = persistences;
            Protocol = protocol;
        }
    }
}
