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
    public sealed class SfsTurboDuTaskDirUsageFileCount
    {
        public readonly int? Block;
        public readonly int? Char;
        public readonly int? Dir;
        public readonly int? Pipe;
        public readonly int? Regular;
        public readonly int? Socket;
        public readonly int? Symlink;

        [OutputConstructor]
        private SfsTurboDuTaskDirUsageFileCount(
            int? block,

            int? @char,

            int? dir,

            int? pipe,

            int? regular,

            int? socket,

            int? symlink)
        {
            Block = block;
            Char = @char;
            Dir = dir;
            Pipe = pipe;
            Regular = regular;
            Socket = socket;
            Symlink = symlink;
        }
    }
}
