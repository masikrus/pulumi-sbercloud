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
    public sealed class FgsFunctionFuncMount
    {
        /// <summary>
        /// The function access path.
        /// </summary>
        public readonly string LocalMountPath;
        /// <summary>
        /// The ID of the mounted resource (corresponding cloud service).
        /// </summary>
        public readonly string MountResource;
        /// <summary>
        /// The remote mount path.
        /// </summary>
        public readonly string MountSharePath;
        /// <summary>
        /// The mount type.
        /// </summary>
        public readonly string MountType;
        /// <summary>
        /// The mount status.
        /// </summary>
        public readonly string? Status;

        [OutputConstructor]
        private FgsFunctionFuncMount(
            string localMountPath,

            string mountResource,

            string mountSharePath,

            string mountType,

            string? status)
        {
            LocalMountPath = localMountPath;
            MountResource = mountResource;
            MountSharePath = mountSharePath;
            MountType = mountType;
            Status = status;
        }
    }
}
