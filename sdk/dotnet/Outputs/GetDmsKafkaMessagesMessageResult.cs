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
    public sealed class GetDmsKafkaMessagesMessageResult
    {
        /// <summary>
        /// Indicates the application ID.
        /// </summary>
        public readonly string AppId;
        /// <summary>
        /// Indicates the big data flag.
        /// </summary>
        public readonly bool HugeMessage;
        /// <summary>
        /// Indicates the message key.
        /// </summary>
        public readonly string Key;
        /// <summary>
        /// Indicates the message ID.
        /// </summary>
        public readonly string MessageId;
        /// <summary>
        /// Indicates the message offset.
        /// </summary>
        public readonly int MessageOffset;
        /// <summary>
        /// Indicates the partition where the message is located.
        /// </summary>
        public readonly int Partition;
        /// <summary>
        /// Indicates the message size.
        /// </summary>
        public readonly int Size;
        /// <summary>
        /// Indicates the message label.
        /// </summary>
        public readonly string Tag;
        /// <summary>
        /// Indicates the message production time.
        /// </summary>
        public readonly string Timestamp;
        /// <summary>
        /// Indicates the message content.
        /// </summary>
        public readonly string Value;

        [OutputConstructor]
        private GetDmsKafkaMessagesMessageResult(
            string appId,

            bool hugeMessage,

            string key,

            string messageId,

            int messageOffset,

            int partition,

            int size,

            string tag,

            string timestamp,

            string value)
        {
            AppId = appId;
            HugeMessage = hugeMessage;
            Key = key;
            MessageId = messageId;
            MessageOffset = messageOffset;
            Partition = partition;
            Size = size;
            Tag = tag;
            Timestamp = timestamp;
            Value = value;
        }
    }
}
