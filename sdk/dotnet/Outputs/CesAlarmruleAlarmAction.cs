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
    public sealed class CesAlarmruleAlarmAction
    {
        public readonly ImmutableArray<string> NotificationLists;
        public readonly string Type;

        [OutputConstructor]
        private CesAlarmruleAlarmAction(
            ImmutableArray<string> notificationLists,

            string type)
        {
            NotificationLists = notificationLists;
            Type = type;
        }
    }
}
