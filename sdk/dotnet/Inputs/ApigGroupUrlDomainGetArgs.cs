// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Sbercloud.Inputs
{

    public sealed class ApigGroupUrlDomainGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("isHttpRedirectToHttps")]
        public Input<bool>? IsHttpRedirectToHttps { get; set; }

        [Input("minSslVersion")]
        public Input<string>? MinSslVersion { get; set; }

        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        public ApigGroupUrlDomainGetArgs()
        {
        }
        public static new ApigGroupUrlDomainGetArgs Empty => new ApigGroupUrlDomainGetArgs();
    }
}
