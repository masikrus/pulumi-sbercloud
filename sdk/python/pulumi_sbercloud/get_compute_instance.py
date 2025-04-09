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
    'GetComputeInstanceResult',
    'AwaitableGetComputeInstanceResult',
    'get_compute_instance',
    'get_compute_instance_output',
]

@pulumi.output_type
class GetComputeInstanceResult:
    """
    A collection of values returned by getComputeInstance.
    """
    def __init__(__self__, availability_zone=None, charging_mode=None, enterprise_project_id=None, expired_time=None, fixed_ip_v4=None, flavor_id=None, flavor_name=None, id=None, image_id=None, image_name=None, instance_id=None, key_pair=None, name=None, networks=None, public_ip=None, region=None, scheduler_hints=None, security_group_ids=None, security_groups=None, status=None, system_disk_id=None, tags=None, user_data=None, volume_attacheds=None):
        if availability_zone and not isinstance(availability_zone, str):
            raise TypeError("Expected argument 'availability_zone' to be a str")
        pulumi.set(__self__, "availability_zone", availability_zone)
        if charging_mode and not isinstance(charging_mode, str):
            raise TypeError("Expected argument 'charging_mode' to be a str")
        pulumi.set(__self__, "charging_mode", charging_mode)
        if enterprise_project_id and not isinstance(enterprise_project_id, str):
            raise TypeError("Expected argument 'enterprise_project_id' to be a str")
        pulumi.set(__self__, "enterprise_project_id", enterprise_project_id)
        if expired_time and not isinstance(expired_time, str):
            raise TypeError("Expected argument 'expired_time' to be a str")
        pulumi.set(__self__, "expired_time", expired_time)
        if fixed_ip_v4 and not isinstance(fixed_ip_v4, str):
            raise TypeError("Expected argument 'fixed_ip_v4' to be a str")
        pulumi.set(__self__, "fixed_ip_v4", fixed_ip_v4)
        if flavor_id and not isinstance(flavor_id, str):
            raise TypeError("Expected argument 'flavor_id' to be a str")
        pulumi.set(__self__, "flavor_id", flavor_id)
        if flavor_name and not isinstance(flavor_name, str):
            raise TypeError("Expected argument 'flavor_name' to be a str")
        pulumi.set(__self__, "flavor_name", flavor_name)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if image_id and not isinstance(image_id, str):
            raise TypeError("Expected argument 'image_id' to be a str")
        pulumi.set(__self__, "image_id", image_id)
        if image_name and not isinstance(image_name, str):
            raise TypeError("Expected argument 'image_name' to be a str")
        pulumi.set(__self__, "image_name", image_name)
        if instance_id and not isinstance(instance_id, str):
            raise TypeError("Expected argument 'instance_id' to be a str")
        pulumi.set(__self__, "instance_id", instance_id)
        if key_pair and not isinstance(key_pair, str):
            raise TypeError("Expected argument 'key_pair' to be a str")
        pulumi.set(__self__, "key_pair", key_pair)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if networks and not isinstance(networks, list):
            raise TypeError("Expected argument 'networks' to be a list")
        pulumi.set(__self__, "networks", networks)
        if public_ip and not isinstance(public_ip, str):
            raise TypeError("Expected argument 'public_ip' to be a str")
        pulumi.set(__self__, "public_ip", public_ip)
        if region and not isinstance(region, str):
            raise TypeError("Expected argument 'region' to be a str")
        pulumi.set(__self__, "region", region)
        if scheduler_hints and not isinstance(scheduler_hints, list):
            raise TypeError("Expected argument 'scheduler_hints' to be a list")
        pulumi.set(__self__, "scheduler_hints", scheduler_hints)
        if security_group_ids and not isinstance(security_group_ids, list):
            raise TypeError("Expected argument 'security_group_ids' to be a list")
        pulumi.set(__self__, "security_group_ids", security_group_ids)
        if security_groups and not isinstance(security_groups, list):
            raise TypeError("Expected argument 'security_groups' to be a list")
        pulumi.set(__self__, "security_groups", security_groups)
        if status and not isinstance(status, str):
            raise TypeError("Expected argument 'status' to be a str")
        pulumi.set(__self__, "status", status)
        if system_disk_id and not isinstance(system_disk_id, str):
            raise TypeError("Expected argument 'system_disk_id' to be a str")
        pulumi.set(__self__, "system_disk_id", system_disk_id)
        if tags and not isinstance(tags, dict):
            raise TypeError("Expected argument 'tags' to be a dict")
        pulumi.set(__self__, "tags", tags)
        if user_data and not isinstance(user_data, str):
            raise TypeError("Expected argument 'user_data' to be a str")
        pulumi.set(__self__, "user_data", user_data)
        if volume_attacheds and not isinstance(volume_attacheds, list):
            raise TypeError("Expected argument 'volume_attacheds' to be a list")
        pulumi.set(__self__, "volume_attacheds", volume_attacheds)

    @property
    @pulumi.getter(name="availabilityZone")
    def availability_zone(self) -> builtins.str:
        return pulumi.get(self, "availability_zone")

    @property
    @pulumi.getter(name="chargingMode")
    def charging_mode(self) -> builtins.str:
        return pulumi.get(self, "charging_mode")

    @property
    @pulumi.getter(name="enterpriseProjectId")
    def enterprise_project_id(self) -> builtins.str:
        return pulumi.get(self, "enterprise_project_id")

    @property
    @pulumi.getter(name="expiredTime")
    def expired_time(self) -> builtins.str:
        return pulumi.get(self, "expired_time")

    @property
    @pulumi.getter(name="fixedIpV4")
    def fixed_ip_v4(self) -> Optional[builtins.str]:
        return pulumi.get(self, "fixed_ip_v4")

    @property
    @pulumi.getter(name="flavorId")
    def flavor_id(self) -> builtins.str:
        return pulumi.get(self, "flavor_id")

    @property
    @pulumi.getter(name="flavorName")
    def flavor_name(self) -> builtins.str:
        return pulumi.get(self, "flavor_name")

    @property
    @pulumi.getter
    def id(self) -> builtins.str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="imageId")
    def image_id(self) -> builtins.str:
        return pulumi.get(self, "image_id")

    @property
    @pulumi.getter(name="imageName")
    def image_name(self) -> builtins.str:
        return pulumi.get(self, "image_name")

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> Optional[builtins.str]:
        return pulumi.get(self, "instance_id")

    @property
    @pulumi.getter(name="keyPair")
    def key_pair(self) -> builtins.str:
        return pulumi.get(self, "key_pair")

    @property
    @pulumi.getter
    def name(self) -> builtins.str:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def networks(self) -> Sequence['outputs.GetComputeInstanceNetworkResult']:
        return pulumi.get(self, "networks")

    @property
    @pulumi.getter(name="publicIp")
    def public_ip(self) -> builtins.str:
        return pulumi.get(self, "public_ip")

    @property
    @pulumi.getter
    def region(self) -> builtins.str:
        return pulumi.get(self, "region")

    @property
    @pulumi.getter(name="schedulerHints")
    def scheduler_hints(self) -> Sequence['outputs.GetComputeInstanceSchedulerHintResult']:
        return pulumi.get(self, "scheduler_hints")

    @property
    @pulumi.getter(name="securityGroupIds")
    def security_group_ids(self) -> Sequence[builtins.str]:
        return pulumi.get(self, "security_group_ids")

    @property
    @pulumi.getter(name="securityGroups")
    def security_groups(self) -> Sequence[builtins.str]:
        return pulumi.get(self, "security_groups")

    @property
    @pulumi.getter
    def status(self) -> builtins.str:
        return pulumi.get(self, "status")

    @property
    @pulumi.getter(name="systemDiskId")
    def system_disk_id(self) -> builtins.str:
        return pulumi.get(self, "system_disk_id")

    @property
    @pulumi.getter
    def tags(self) -> Optional[Mapping[str, builtins.str]]:
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="userData")
    def user_data(self) -> builtins.str:
        return pulumi.get(self, "user_data")

    @property
    @pulumi.getter(name="volumeAttacheds")
    def volume_attacheds(self) -> Sequence['outputs.GetComputeInstanceVolumeAttachedResult']:
        return pulumi.get(self, "volume_attacheds")


class AwaitableGetComputeInstanceResult(GetComputeInstanceResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetComputeInstanceResult(
            availability_zone=self.availability_zone,
            charging_mode=self.charging_mode,
            enterprise_project_id=self.enterprise_project_id,
            expired_time=self.expired_time,
            fixed_ip_v4=self.fixed_ip_v4,
            flavor_id=self.flavor_id,
            flavor_name=self.flavor_name,
            id=self.id,
            image_id=self.image_id,
            image_name=self.image_name,
            instance_id=self.instance_id,
            key_pair=self.key_pair,
            name=self.name,
            networks=self.networks,
            public_ip=self.public_ip,
            region=self.region,
            scheduler_hints=self.scheduler_hints,
            security_group_ids=self.security_group_ids,
            security_groups=self.security_groups,
            status=self.status,
            system_disk_id=self.system_disk_id,
            tags=self.tags,
            user_data=self.user_data,
            volume_attacheds=self.volume_attacheds)


def get_compute_instance(enterprise_project_id: Optional[builtins.str] = None,
                         fixed_ip_v4: Optional[builtins.str] = None,
                         flavor_id: Optional[builtins.str] = None,
                         instance_id: Optional[builtins.str] = None,
                         name: Optional[builtins.str] = None,
                         region: Optional[builtins.str] = None,
                         tags: Optional[Mapping[str, builtins.str]] = None,
                         opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetComputeInstanceResult:
    """
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()
    __args__['enterpriseProjectId'] = enterprise_project_id
    __args__['fixedIpV4'] = fixed_ip_v4
    __args__['flavorId'] = flavor_id
    __args__['instanceId'] = instance_id
    __args__['name'] = name
    __args__['region'] = region
    __args__['tags'] = tags
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('sbercloud:index/getComputeInstance:getComputeInstance', __args__, opts=opts, typ=GetComputeInstanceResult).value

    return AwaitableGetComputeInstanceResult(
        availability_zone=pulumi.get(__ret__, 'availability_zone'),
        charging_mode=pulumi.get(__ret__, 'charging_mode'),
        enterprise_project_id=pulumi.get(__ret__, 'enterprise_project_id'),
        expired_time=pulumi.get(__ret__, 'expired_time'),
        fixed_ip_v4=pulumi.get(__ret__, 'fixed_ip_v4'),
        flavor_id=pulumi.get(__ret__, 'flavor_id'),
        flavor_name=pulumi.get(__ret__, 'flavor_name'),
        id=pulumi.get(__ret__, 'id'),
        image_id=pulumi.get(__ret__, 'image_id'),
        image_name=pulumi.get(__ret__, 'image_name'),
        instance_id=pulumi.get(__ret__, 'instance_id'),
        key_pair=pulumi.get(__ret__, 'key_pair'),
        name=pulumi.get(__ret__, 'name'),
        networks=pulumi.get(__ret__, 'networks'),
        public_ip=pulumi.get(__ret__, 'public_ip'),
        region=pulumi.get(__ret__, 'region'),
        scheduler_hints=pulumi.get(__ret__, 'scheduler_hints'),
        security_group_ids=pulumi.get(__ret__, 'security_group_ids'),
        security_groups=pulumi.get(__ret__, 'security_groups'),
        status=pulumi.get(__ret__, 'status'),
        system_disk_id=pulumi.get(__ret__, 'system_disk_id'),
        tags=pulumi.get(__ret__, 'tags'),
        user_data=pulumi.get(__ret__, 'user_data'),
        volume_attacheds=pulumi.get(__ret__, 'volume_attacheds'))
def get_compute_instance_output(enterprise_project_id: Optional[pulumi.Input[Optional[builtins.str]]] = None,
                                fixed_ip_v4: Optional[pulumi.Input[Optional[builtins.str]]] = None,
                                flavor_id: Optional[pulumi.Input[Optional[builtins.str]]] = None,
                                instance_id: Optional[pulumi.Input[Optional[builtins.str]]] = None,
                                name: Optional[pulumi.Input[Optional[builtins.str]]] = None,
                                region: Optional[pulumi.Input[Optional[builtins.str]]] = None,
                                tags: Optional[pulumi.Input[Optional[Mapping[str, builtins.str]]]] = None,
                                opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetComputeInstanceResult]:
    """
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()
    __args__['enterpriseProjectId'] = enterprise_project_id
    __args__['fixedIpV4'] = fixed_ip_v4
    __args__['flavorId'] = flavor_id
    __args__['instanceId'] = instance_id
    __args__['name'] = name
    __args__['region'] = region
    __args__['tags'] = tags
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('sbercloud:index/getComputeInstance:getComputeInstance', __args__, opts=opts, typ=GetComputeInstanceResult)
    return __ret__.apply(lambda __response__: GetComputeInstanceResult(
        availability_zone=pulumi.get(__response__, 'availability_zone'),
        charging_mode=pulumi.get(__response__, 'charging_mode'),
        enterprise_project_id=pulumi.get(__response__, 'enterprise_project_id'),
        expired_time=pulumi.get(__response__, 'expired_time'),
        fixed_ip_v4=pulumi.get(__response__, 'fixed_ip_v4'),
        flavor_id=pulumi.get(__response__, 'flavor_id'),
        flavor_name=pulumi.get(__response__, 'flavor_name'),
        id=pulumi.get(__response__, 'id'),
        image_id=pulumi.get(__response__, 'image_id'),
        image_name=pulumi.get(__response__, 'image_name'),
        instance_id=pulumi.get(__response__, 'instance_id'),
        key_pair=pulumi.get(__response__, 'key_pair'),
        name=pulumi.get(__response__, 'name'),
        networks=pulumi.get(__response__, 'networks'),
        public_ip=pulumi.get(__response__, 'public_ip'),
        region=pulumi.get(__response__, 'region'),
        scheduler_hints=pulumi.get(__response__, 'scheduler_hints'),
        security_group_ids=pulumi.get(__response__, 'security_group_ids'),
        security_groups=pulumi.get(__response__, 'security_groups'),
        status=pulumi.get(__response__, 'status'),
        system_disk_id=pulumi.get(__response__, 'system_disk_id'),
        tags=pulumi.get(__response__, 'tags'),
        user_data=pulumi.get(__response__, 'user_data'),
        volume_attacheds=pulumi.get(__response__, 'volume_attacheds')))
