// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Sbercloud
{
    public static class GetImagesImages
    {
        public static Task<GetImagesImagesResult> InvokeAsync(GetImagesImagesArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetImagesImagesResult>("sbercloud:index/getImagesImages:getImagesImages", args ?? new GetImagesImagesArgs(), options.WithDefaults());

        public static Output<GetImagesImagesResult> Invoke(GetImagesImagesInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetImagesImagesResult>("sbercloud:index/getImagesImages:getImagesImages", args ?? new GetImagesImagesInvokeArgs(), options.WithDefaults());

        public static Output<GetImagesImagesResult> Invoke(GetImagesImagesInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetImagesImagesResult>("sbercloud:index/getImagesImages:getImagesImages", args ?? new GetImagesImagesInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetImagesImagesArgs : global::Pulumi.InvokeArgs
    {
        [Input("architecture")]
        public string? Architecture { get; set; }

        [Input("enterpriseProjectId")]
        public string? EnterpriseProjectId { get; set; }

        [Input("flavorId")]
        public string? FlavorId { get; set; }

        [Input("imageId")]
        public string? ImageId { get; set; }

        [Input("imageType")]
        public string? ImageType { get; set; }

        [Input("isWholeImage")]
        public bool? IsWholeImage { get; set; }

        [Input("name")]
        public string? Name { get; set; }

        [Input("nameRegex")]
        public string? NameRegex { get; set; }

        [Input("os")]
        public string? Os { get; set; }

        [Input("osVersion")]
        public string? OsVersion { get; set; }

        [Input("owner")]
        public string? Owner { get; set; }

        [Input("region")]
        public string? Region { get; set; }

        [Input("sortDirection")]
        public string? SortDirection { get; set; }

        [Input("sortKey")]
        public string? SortKey { get; set; }

        [Input("tag")]
        public string? Tag { get; set; }

        [Input("visibility")]
        public string? Visibility { get; set; }

        public GetImagesImagesArgs()
        {
        }
        public static new GetImagesImagesArgs Empty => new GetImagesImagesArgs();
    }

    public sealed class GetImagesImagesInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("architecture")]
        public Input<string>? Architecture { get; set; }

        [Input("enterpriseProjectId")]
        public Input<string>? EnterpriseProjectId { get; set; }

        [Input("flavorId")]
        public Input<string>? FlavorId { get; set; }

        [Input("imageId")]
        public Input<string>? ImageId { get; set; }

        [Input("imageType")]
        public Input<string>? ImageType { get; set; }

        [Input("isWholeImage")]
        public Input<bool>? IsWholeImage { get; set; }

        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("nameRegex")]
        public Input<string>? NameRegex { get; set; }

        [Input("os")]
        public Input<string>? Os { get; set; }

        [Input("osVersion")]
        public Input<string>? OsVersion { get; set; }

        [Input("owner")]
        public Input<string>? Owner { get; set; }

        [Input("region")]
        public Input<string>? Region { get; set; }

        [Input("sortDirection")]
        public Input<string>? SortDirection { get; set; }

        [Input("sortKey")]
        public Input<string>? SortKey { get; set; }

        [Input("tag")]
        public Input<string>? Tag { get; set; }

        [Input("visibility")]
        public Input<string>? Visibility { get; set; }

        public GetImagesImagesInvokeArgs()
        {
        }
        public static new GetImagesImagesInvokeArgs Empty => new GetImagesImagesInvokeArgs();
    }


    [OutputType]
    public sealed class GetImagesImagesResult
    {
        public readonly string? Architecture;
        public readonly string? EnterpriseProjectId;
        public readonly string? FlavorId;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string? ImageId;
        public readonly string? ImageType;
        public readonly ImmutableArray<Outputs.GetImagesImagesImageResult> Images;
        public readonly bool? IsWholeImage;
        public readonly string? Name;
        public readonly string? NameRegex;
        public readonly string? Os;
        public readonly string? OsVersion;
        public readonly string? Owner;
        public readonly string Region;
        public readonly string? SortDirection;
        public readonly string? SortKey;
        public readonly string? Tag;
        public readonly string? Visibility;

        [OutputConstructor]
        private GetImagesImagesResult(
            string? architecture,

            string? enterpriseProjectId,

            string? flavorId,

            string id,

            string? imageId,

            string? imageType,

            ImmutableArray<Outputs.GetImagesImagesImageResult> images,

            bool? isWholeImage,

            string? name,

            string? nameRegex,

            string? os,

            string? osVersion,

            string? owner,

            string region,

            string? sortDirection,

            string? sortKey,

            string? tag,

            string? visibility)
        {
            Architecture = architecture;
            EnterpriseProjectId = enterpriseProjectId;
            FlavorId = flavorId;
            Id = id;
            ImageId = imageId;
            ImageType = imageType;
            Images = images;
            IsWholeImage = isWholeImage;
            Name = name;
            NameRegex = nameRegex;
            Os = os;
            OsVersion = osVersion;
            Owner = owner;
            Region = region;
            SortDirection = sortDirection;
            SortKey = sortKey;
            Tag = tag;
            Visibility = visibility;
        }
    }
}
