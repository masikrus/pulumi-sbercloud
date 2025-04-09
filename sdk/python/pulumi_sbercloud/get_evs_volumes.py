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
    'GetEvsVolumesResult',
    'AwaitableGetEvsVolumesResult',
    'get_evs_volumes',
    'get_evs_volumes_output',
]

@pulumi.output_type
class GetEvsVolumesResult:
    """
    A collection of values returned by getEvsVolumes.
    """
    def __init__(__self__, availability_zone=None, enterprise_project_id=None, id=None, name=None, region=None, server_id=None, shareable=None, status=None, tags=None, volume_id=None, volume_type_id=None, volumes=None):
        if availability_zone and not isinstance(availability_zone, str):
            raise TypeError("Expected argument 'availability_zone' to be a str")
        pulumi.set(__self__, "availability_zone", availability_zone)
        if enterprise_project_id and not isinstance(enterprise_project_id, str):
            raise TypeError("Expected argument 'enterprise_project_id' to be a str")
        pulumi.set(__self__, "enterprise_project_id", enterprise_project_id)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if region and not isinstance(region, str):
            raise TypeError("Expected argument 'region' to be a str")
        pulumi.set(__self__, "region", region)
        if server_id and not isinstance(server_id, str):
            raise TypeError("Expected argument 'server_id' to be a str")
        pulumi.set(__self__, "server_id", server_id)
        if shareable and not isinstance(shareable, bool):
            raise TypeError("Expected argument 'shareable' to be a bool")
        pulumi.set(__self__, "shareable", shareable)
        if status and not isinstance(status, str):
            raise TypeError("Expected argument 'status' to be a str")
        pulumi.set(__self__, "status", status)
        if tags and not isinstance(tags, dict):
            raise TypeError("Expected argument 'tags' to be a dict")
        pulumi.set(__self__, "tags", tags)
        if volume_id and not isinstance(volume_id, str):
            raise TypeError("Expected argument 'volume_id' to be a str")
        pulumi.set(__self__, "volume_id", volume_id)
        if volume_type_id and not isinstance(volume_type_id, str):
            raise TypeError("Expected argument 'volume_type_id' to be a str")
        pulumi.set(__self__, "volume_type_id", volume_type_id)
        if volumes and not isinstance(volumes, list):
            raise TypeError("Expected argument 'volumes' to be a list")
        pulumi.set(__self__, "volumes", volumes)

    @property
    @pulumi.getter(name="availabilityZone")
    def availability_zone(self) -> Optional[builtins.str]:
        return pulumi.get(self, "availability_zone")

    @property
    @pulumi.getter(name="enterpriseProjectId")
    def enterprise_project_id(self) -> Optional[builtins.str]:
        return pulumi.get(self, "enterprise_project_id")

    @property
    @pulumi.getter
    def id(self) -> builtins.str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def name(self) -> Optional[builtins.str]:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def region(self) -> Optional[builtins.str]:
        return pulumi.get(self, "region")

    @property
    @pulumi.getter(name="serverId")
    def server_id(self) -> Optional[builtins.str]:
        return pulumi.get(self, "server_id")

    @property
    @pulumi.getter
    def shareable(self) -> Optional[builtins.bool]:
        return pulumi.get(self, "shareable")

    @property
    @pulumi.getter
    def status(self) -> Optional[builtins.str]:
        return pulumi.get(self, "status")

    @property
    @pulumi.getter
    def tags(self) -> Optional[Mapping[str, builtins.str]]:
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="volumeId")
    def volume_id(self) -> Optional[builtins.str]:
        return pulumi.get(self, "volume_id")

    @property
    @pulumi.getter(name="volumeTypeId")
    def volume_type_id(self) -> Optional[builtins.str]:
        return pulumi.get(self, "volume_type_id")

    @property
    @pulumi.getter
    def volumes(self) -> Sequence['outputs.GetEvsVolumesVolumeResult']:
        return pulumi.get(self, "volumes")


class AwaitableGetEvsVolumesResult(GetEvsVolumesResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetEvsVolumesResult(
            availability_zone=self.availability_zone,
            enterprise_project_id=self.enterprise_project_id,
            id=self.id,
            name=self.name,
            region=self.region,
            server_id=self.server_id,
            shareable=self.shareable,
            status=self.status,
            tags=self.tags,
            volume_id=self.volume_id,
            volume_type_id=self.volume_type_id,
            volumes=self.volumes)


def get_evs_volumes(availability_zone: Optional[builtins.str] = None,
                    enterprise_project_id: Optional[builtins.str] = None,
                    name: Optional[builtins.str] = None,
                    region: Optional[builtins.str] = None,
                    server_id: Optional[builtins.str] = None,
                    shareable: Optional[builtins.bool] = None,
                    status: Optional[builtins.str] = None,
                    tags: Optional[Mapping[str, builtins.str]] = None,
                    volume_id: Optional[builtins.str] = None,
                    volume_type_id: Optional[builtins.str] = None,
                    opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetEvsVolumesResult:
    """
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()
    __args__['availabilityZone'] = availability_zone
    __args__['enterpriseProjectId'] = enterprise_project_id
    __args__['name'] = name
    __args__['region'] = region
    __args__['serverId'] = server_id
    __args__['shareable'] = shareable
    __args__['status'] = status
    __args__['tags'] = tags
    __args__['volumeId'] = volume_id
    __args__['volumeTypeId'] = volume_type_id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('sbercloud:index/getEvsVolumes:getEvsVolumes', __args__, opts=opts, typ=GetEvsVolumesResult).value

    return AwaitableGetEvsVolumesResult(
        availability_zone=pulumi.get(__ret__, 'availability_zone'),
        enterprise_project_id=pulumi.get(__ret__, 'enterprise_project_id'),
        id=pulumi.get(__ret__, 'id'),
        name=pulumi.get(__ret__, 'name'),
        region=pulumi.get(__ret__, 'region'),
        server_id=pulumi.get(__ret__, 'server_id'),
        shareable=pulumi.get(__ret__, 'shareable'),
        status=pulumi.get(__ret__, 'status'),
        tags=pulumi.get(__ret__, 'tags'),
        volume_id=pulumi.get(__ret__, 'volume_id'),
        volume_type_id=pulumi.get(__ret__, 'volume_type_id'),
        volumes=pulumi.get(__ret__, 'volumes'))
def get_evs_volumes_output(availability_zone: Optional[pulumi.Input[Optional[builtins.str]]] = None,
                           enterprise_project_id: Optional[pulumi.Input[Optional[builtins.str]]] = None,
                           name: Optional[pulumi.Input[Optional[builtins.str]]] = None,
                           region: Optional[pulumi.Input[Optional[builtins.str]]] = None,
                           server_id: Optional[pulumi.Input[Optional[builtins.str]]] = None,
                           shareable: Optional[pulumi.Input[Optional[builtins.bool]]] = None,
                           status: Optional[pulumi.Input[Optional[builtins.str]]] = None,
                           tags: Optional[pulumi.Input[Optional[Mapping[str, builtins.str]]]] = None,
                           volume_id: Optional[pulumi.Input[Optional[builtins.str]]] = None,
                           volume_type_id: Optional[pulumi.Input[Optional[builtins.str]]] = None,
                           opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetEvsVolumesResult]:
    """
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()
    __args__['availabilityZone'] = availability_zone
    __args__['enterpriseProjectId'] = enterprise_project_id
    __args__['name'] = name
    __args__['region'] = region
    __args__['serverId'] = server_id
    __args__['shareable'] = shareable
    __args__['status'] = status
    __args__['tags'] = tags
    __args__['volumeId'] = volume_id
    __args__['volumeTypeId'] = volume_type_id
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('sbercloud:index/getEvsVolumes:getEvsVolumes', __args__, opts=opts, typ=GetEvsVolumesResult)
    return __ret__.apply(lambda __response__: GetEvsVolumesResult(
        availability_zone=pulumi.get(__response__, 'availability_zone'),
        enterprise_project_id=pulumi.get(__response__, 'enterprise_project_id'),
        id=pulumi.get(__response__, 'id'),
        name=pulumi.get(__response__, 'name'),
        region=pulumi.get(__response__, 'region'),
        server_id=pulumi.get(__response__, 'server_id'),
        shareable=pulumi.get(__response__, 'shareable'),
        status=pulumi.get(__response__, 'status'),
        tags=pulumi.get(__response__, 'tags'),
        volume_id=pulumi.get(__response__, 'volume_id'),
        volume_type_id=pulumi.get(__response__, 'volume_type_id'),
        volumes=pulumi.get(__response__, 'volumes')))
