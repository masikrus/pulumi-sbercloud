// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Sbercloud.Inputs
{

    public sealed class CceNodeStorageGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("groups", required: true)]
        private InputList<Inputs.CceNodeStorageGroupGetArgs>? _groups;
        public InputList<Inputs.CceNodeStorageGroupGetArgs> Groups
        {
            get => _groups ?? (_groups = new InputList<Inputs.CceNodeStorageGroupGetArgs>());
            set => _groups = value;
        }

        [Input("selectors", required: true)]
        private InputList<Inputs.CceNodeStorageSelectorGetArgs>? _selectors;
        public InputList<Inputs.CceNodeStorageSelectorGetArgs> Selectors
        {
            get => _selectors ?? (_selectors = new InputList<Inputs.CceNodeStorageSelectorGetArgs>());
            set => _selectors = value;
        }

        public CceNodeStorageGetArgs()
        {
        }
        public static new CceNodeStorageGetArgs Empty => new CceNodeStorageGetArgs();
    }
}
