// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Sbercloud.Inputs
{

    public sealed class ApigApiRequestParamArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The default value of the parameter.
        /// </summary>
        [Input("default")]
        public Input<string>? Default { get; set; }

        /// <summary>
        /// The parameter description.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The enumerated value.
        /// </summary>
        [Input("enumeration")]
        public Input<string>? Enumeration { get; set; }

        /// <summary>
        /// The parameter example.
        /// </summary>
        [Input("example")]
        public Input<string>? Example { get; set; }

        /// <summary>
        /// Where this parameter is located.
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// The maximum value or length (string parameter) for parameter.
        /// </summary>
        [Input("maximum")]
        public Input<int>? Maximum { get; set; }

        /// <summary>
        /// The minimum value or length (string parameter) for parameter.
        /// </summary>
        [Input("minimum")]
        public Input<int>? Minimum { get; set; }

        /// <summary>
        /// The name of the request parameter.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        [Input("orchestrations")]
        private InputList<string>? _orchestrations;

        /// <summary>
        /// The list of orchestration rules that parameter used.
        /// </summary>
        public InputList<string> Orchestrations
        {
            get => _orchestrations ?? (_orchestrations = new InputList<string>());
            set => _orchestrations = value;
        }

        /// <summary>
        /// Whether to transparently transfer the parameter.
        /// </summary>
        [Input("passthrough")]
        public Input<bool>? Passthrough { get; set; }

        /// <summary>
        /// Whether this parameter is required.
        /// </summary>
        [Input("required")]
        public Input<bool>? Required { get; set; }

        /// <summary>
        /// The parameter type.
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        /// <summary>
        /// Whether to enable the parameter validation.
        /// </summary>
        [Input("validEnable")]
        public Input<int>? ValidEnable { get; set; }

        public ApigApiRequestParamArgs()
        {
        }
        public static new ApigApiRequestParamArgs Empty => new ApigApiRequestParamArgs();
    }
}
