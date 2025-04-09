# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import builtins
import copy
import warnings
import sys
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
if sys.version_info >= (3, 11):
    from typing import NotRequired, TypedDict, TypeAlias
else:
    from typing_extensions import NotRequired, TypedDict, TypeAlias
from . import _utilities
from . import outputs

__all__ = [
    'GetImagesImagesResult',
    'AwaitableGetImagesImagesResult',
    'get_images_images',
    'get_images_images_output',
]

@pulumi.output_type
class GetImagesImagesResult:
    """
    A collection of values returned by getImagesImages.
    """
    def __init__(__self__, architecture=None, enterprise_project_id=None, flavor_id=None, id=None, image_id=None, image_type=None, images=None, is_whole_image=None, name=None, name_regex=None, os=None, os_version=None, owner=None, region=None, sort_direction=None, sort_key=None, tag=None, visibility=None):
        if architecture and not isinstance(architecture, str):
            raise TypeError("Expected argument 'architecture' to be a str")
        pulumi.set(__self__, "architecture", architecture)
        if enterprise_project_id and not isinstance(enterprise_project_id, str):
            raise TypeError("Expected argument 'enterprise_project_id' to be a str")
        pulumi.set(__self__, "enterprise_project_id", enterprise_project_id)
        if flavor_id and not isinstance(flavor_id, str):
            raise TypeError("Expected argument 'flavor_id' to be a str")
        pulumi.set(__self__, "flavor_id", flavor_id)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if image_id and not isinstance(image_id, str):
            raise TypeError("Expected argument 'image_id' to be a str")
        pulumi.set(__self__, "image_id", image_id)
        if image_type and not isinstance(image_type, str):
            raise TypeError("Expected argument 'image_type' to be a str")
        pulumi.set(__self__, "image_type", image_type)
        if images and not isinstance(images, list):
            raise TypeError("Expected argument 'images' to be a list")
        pulumi.set(__self__, "images", images)
        if is_whole_image and not isinstance(is_whole_image, bool):
            raise TypeError("Expected argument 'is_whole_image' to be a bool")
        pulumi.set(__self__, "is_whole_image", is_whole_image)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if name_regex and not isinstance(name_regex, str):
            raise TypeError("Expected argument 'name_regex' to be a str")
        pulumi.set(__self__, "name_regex", name_regex)
        if os and not isinstance(os, str):
            raise TypeError("Expected argument 'os' to be a str")
        pulumi.set(__self__, "os", os)
        if os_version and not isinstance(os_version, str):
            raise TypeError("Expected argument 'os_version' to be a str")
        pulumi.set(__self__, "os_version", os_version)
        if owner and not isinstance(owner, str):
            raise TypeError("Expected argument 'owner' to be a str")
        pulumi.set(__self__, "owner", owner)
        if region and not isinstance(region, str):
            raise TypeError("Expected argument 'region' to be a str")
        pulumi.set(__self__, "region", region)
        if sort_direction and not isinstance(sort_direction, str):
            raise TypeError("Expected argument 'sort_direction' to be a str")
        pulumi.set(__self__, "sort_direction", sort_direction)
        if sort_key and not isinstance(sort_key, str):
            raise TypeError("Expected argument 'sort_key' to be a str")
        pulumi.set(__self__, "sort_key", sort_key)
        if tag and not isinstance(tag, str):
            raise TypeError("Expected argument 'tag' to be a str")
        pulumi.set(__self__, "tag", tag)
        if visibility and not isinstance(visibility, str):
            raise TypeError("Expected argument 'visibility' to be a str")
        pulumi.set(__self__, "visibility", visibility)

    @property
    @pulumi.getter
    def architecture(self) -> Optional[builtins.str]:
        return pulumi.get(self, "architecture")

    @property
    @pulumi.getter(name="enterpriseProjectId")
    def enterprise_project_id(self) -> Optional[builtins.str]:
        return pulumi.get(self, "enterprise_project_id")

    @property
    @pulumi.getter(name="flavorId")
    def flavor_id(self) -> Optional[builtins.str]:
        return pulumi.get(self, "flavor_id")

    @property
    @pulumi.getter
    def id(self) -> builtins.str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="imageId")
    def image_id(self) -> Optional[builtins.str]:
        return pulumi.get(self, "image_id")

    @property
    @pulumi.getter(name="imageType")
    def image_type(self) -> Optional[builtins.str]:
        return pulumi.get(self, "image_type")

    @property
    @pulumi.getter
    def images(self) -> Sequence['outputs.GetImagesImagesImageResult']:
        return pulumi.get(self, "images")

    @property
    @pulumi.getter(name="isWholeImage")
    def is_whole_image(self) -> Optional[builtins.bool]:
        return pulumi.get(self, "is_whole_image")

    @property
    @pulumi.getter
    def name(self) -> Optional[builtins.str]:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="nameRegex")
    def name_regex(self) -> Optional[builtins.str]:
        return pulumi.get(self, "name_regex")

    @property
    @pulumi.getter
    def os(self) -> Optional[builtins.str]:
        return pulumi.get(self, "os")

    @property
    @pulumi.getter(name="osVersion")
    def os_version(self) -> Optional[builtins.str]:
        return pulumi.get(self, "os_version")

    @property
    @pulumi.getter
    def owner(self) -> Optional[builtins.str]:
        return pulumi.get(self, "owner")

    @property
    @pulumi.getter
    def region(self) -> builtins.str:
        return pulumi.get(self, "region")

    @property
    @pulumi.getter(name="sortDirection")
    def sort_direction(self) -> Optional[builtins.str]:
        return pulumi.get(self, "sort_direction")

    @property
    @pulumi.getter(name="sortKey")
    def sort_key(self) -> Optional[builtins.str]:
        return pulumi.get(self, "sort_key")

    @property
    @pulumi.getter
    def tag(self) -> Optional[builtins.str]:
        return pulumi.get(self, "tag")

    @property
    @pulumi.getter
    def visibility(self) -> Optional[builtins.str]:
        return pulumi.get(self, "visibility")


class AwaitableGetImagesImagesResult(GetImagesImagesResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetImagesImagesResult(
            architecture=self.architecture,
            enterprise_project_id=self.enterprise_project_id,
            flavor_id=self.flavor_id,
            id=self.id,
            image_id=self.image_id,
            image_type=self.image_type,
            images=self.images,
            is_whole_image=self.is_whole_image,
            name=self.name,
            name_regex=self.name_regex,
            os=self.os,
            os_version=self.os_version,
            owner=self.owner,
            region=self.region,
            sort_direction=self.sort_direction,
            sort_key=self.sort_key,
            tag=self.tag,
            visibility=self.visibility)


def get_images_images(architecture: Optional[builtins.str] = None,
                      enterprise_project_id: Optional[builtins.str] = None,
                      flavor_id: Optional[builtins.str] = None,
                      image_id: Optional[builtins.str] = None,
                      image_type: Optional[builtins.str] = None,
                      is_whole_image: Optional[builtins.bool] = None,
                      name: Optional[builtins.str] = None,
                      name_regex: Optional[builtins.str] = None,
                      os: Optional[builtins.str] = None,
                      os_version: Optional[builtins.str] = None,
                      owner: Optional[builtins.str] = None,
                      region: Optional[builtins.str] = None,
                      sort_direction: Optional[builtins.str] = None,
                      sort_key: Optional[builtins.str] = None,
                      tag: Optional[builtins.str] = None,
                      visibility: Optional[builtins.str] = None,
                      opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetImagesImagesResult:
    """
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()
    __args__['architecture'] = architecture
    __args__['enterpriseProjectId'] = enterprise_project_id
    __args__['flavorId'] = flavor_id
    __args__['imageId'] = image_id
    __args__['imageType'] = image_type
    __args__['isWholeImage'] = is_whole_image
    __args__['name'] = name
    __args__['nameRegex'] = name_regex
    __args__['os'] = os
    __args__['osVersion'] = os_version
    __args__['owner'] = owner
    __args__['region'] = region
    __args__['sortDirection'] = sort_direction
    __args__['sortKey'] = sort_key
    __args__['tag'] = tag
    __args__['visibility'] = visibility
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('sbercloud:index/getImagesImages:getImagesImages', __args__, opts=opts, typ=GetImagesImagesResult).value

    return AwaitableGetImagesImagesResult(
        architecture=pulumi.get(__ret__, 'architecture'),
        enterprise_project_id=pulumi.get(__ret__, 'enterprise_project_id'),
        flavor_id=pulumi.get(__ret__, 'flavor_id'),
        id=pulumi.get(__ret__, 'id'),
        image_id=pulumi.get(__ret__, 'image_id'),
        image_type=pulumi.get(__ret__, 'image_type'),
        images=pulumi.get(__ret__, 'images'),
        is_whole_image=pulumi.get(__ret__, 'is_whole_image'),
        name=pulumi.get(__ret__, 'name'),
        name_regex=pulumi.get(__ret__, 'name_regex'),
        os=pulumi.get(__ret__, 'os'),
        os_version=pulumi.get(__ret__, 'os_version'),
        owner=pulumi.get(__ret__, 'owner'),
        region=pulumi.get(__ret__, 'region'),
        sort_direction=pulumi.get(__ret__, 'sort_direction'),
        sort_key=pulumi.get(__ret__, 'sort_key'),
        tag=pulumi.get(__ret__, 'tag'),
        visibility=pulumi.get(__ret__, 'visibility'))
def get_images_images_output(architecture: Optional[pulumi.Input[Optional[builtins.str]]] = None,
                             enterprise_project_id: Optional[pulumi.Input[Optional[builtins.str]]] = None,
                             flavor_id: Optional[pulumi.Input[Optional[builtins.str]]] = None,
                             image_id: Optional[pulumi.Input[Optional[builtins.str]]] = None,
                             image_type: Optional[pulumi.Input[Optional[builtins.str]]] = None,
                             is_whole_image: Optional[pulumi.Input[Optional[builtins.bool]]] = None,
                             name: Optional[pulumi.Input[Optional[builtins.str]]] = None,
                             name_regex: Optional[pulumi.Input[Optional[builtins.str]]] = None,
                             os: Optional[pulumi.Input[Optional[builtins.str]]] = None,
                             os_version: Optional[pulumi.Input[Optional[builtins.str]]] = None,
                             owner: Optional[pulumi.Input[Optional[builtins.str]]] = None,
                             region: Optional[pulumi.Input[Optional[builtins.str]]] = None,
                             sort_direction: Optional[pulumi.Input[Optional[builtins.str]]] = None,
                             sort_key: Optional[pulumi.Input[Optional[builtins.str]]] = None,
                             tag: Optional[pulumi.Input[Optional[builtins.str]]] = None,
                             visibility: Optional[pulumi.Input[Optional[builtins.str]]] = None,
                             opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetImagesImagesResult]:
    """
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()
    __args__['architecture'] = architecture
    __args__['enterpriseProjectId'] = enterprise_project_id
    __args__['flavorId'] = flavor_id
    __args__['imageId'] = image_id
    __args__['imageType'] = image_type
    __args__['isWholeImage'] = is_whole_image
    __args__['name'] = name
    __args__['nameRegex'] = name_regex
    __args__['os'] = os
    __args__['osVersion'] = os_version
    __args__['owner'] = owner
    __args__['region'] = region
    __args__['sortDirection'] = sort_direction
    __args__['sortKey'] = sort_key
    __args__['tag'] = tag
    __args__['visibility'] = visibility
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('sbercloud:index/getImagesImages:getImagesImages', __args__, opts=opts, typ=GetImagesImagesResult)
    return __ret__.apply(lambda __response__: GetImagesImagesResult(
        architecture=pulumi.get(__response__, 'architecture'),
        enterprise_project_id=pulumi.get(__response__, 'enterprise_project_id'),
        flavor_id=pulumi.get(__response__, 'flavor_id'),
        id=pulumi.get(__response__, 'id'),
        image_id=pulumi.get(__response__, 'image_id'),
        image_type=pulumi.get(__response__, 'image_type'),
        images=pulumi.get(__response__, 'images'),
        is_whole_image=pulumi.get(__response__, 'is_whole_image'),
        name=pulumi.get(__response__, 'name'),
        name_regex=pulumi.get(__response__, 'name_regex'),
        os=pulumi.get(__response__, 'os'),
        os_version=pulumi.get(__response__, 'os_version'),
        owner=pulumi.get(__response__, 'owner'),
        region=pulumi.get(__response__, 'region'),
        sort_direction=pulumi.get(__response__, 'sort_direction'),
        sort_key=pulumi.get(__response__, 'sort_key'),
        tag=pulumi.get(__response__, 'tag'),
        visibility=pulumi.get(__response__, 'visibility')))
