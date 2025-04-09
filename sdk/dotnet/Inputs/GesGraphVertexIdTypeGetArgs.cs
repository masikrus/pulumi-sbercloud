// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Sbercloud.Inputs
{

    public sealed class GesGraphVertexIdTypeGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The length of ID.
        /// </summary>
        [Input("idLength")]
        public Input<int>? IdLength { get; set; }

        /// <summary>
        /// Vertex ID type.
        /// </summary>
        [Input("idType")]
        public Input<string>? IdType { get; set; }

        public GesGraphVertexIdTypeGetArgs()
        {
        }
        public static new GesGraphVertexIdTypeGetArgs Empty => new GesGraphVertexIdTypeGetArgs();
    }
}
