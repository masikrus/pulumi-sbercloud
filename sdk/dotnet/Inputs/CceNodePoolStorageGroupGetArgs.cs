// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Sbercloud.Inputs
{

    public sealed class CceNodePoolStorageGroupGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("cceManaged")]
        public Input<bool>? CceManaged { get; set; }

        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        [Input("selectorNames", required: true)]
        private InputList<string>? _selectorNames;
        public InputList<string> SelectorNames
        {
            get => _selectorNames ?? (_selectorNames = new InputList<string>());
            set => _selectorNames = value;
        }

        [Input("virtualSpaces", required: true)]
        private InputList<Inputs.CceNodePoolStorageGroupVirtualSpaceGetArgs>? _virtualSpaces;
        public InputList<Inputs.CceNodePoolStorageGroupVirtualSpaceGetArgs> VirtualSpaces
        {
            get => _virtualSpaces ?? (_virtualSpaces = new InputList<Inputs.CceNodePoolStorageGroupVirtualSpaceGetArgs>());
            set => _virtualSpaces = value;
        }

        public CceNodePoolStorageGroupGetArgs()
        {
        }
        public static new CceNodePoolStorageGroupGetArgs Empty => new CceNodePoolStorageGroupGetArgs();
    }
}
