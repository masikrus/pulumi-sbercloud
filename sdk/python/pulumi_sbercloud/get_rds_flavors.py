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
    'GetRdsFlavorsResult',
    'AwaitableGetRdsFlavorsResult',
    'get_rds_flavors',
    'get_rds_flavors_output',
]

@pulumi.output_type
class GetRdsFlavorsResult:
    """
    A collection of values returned by getRdsFlavors.
    """
    def __init__(__self__, availability_zone=None, db_type=None, db_version=None, flavors=None, group_type=None, id=None, instance_mode=None, memory=None, region=None, vcpus=None):
        if availability_zone and not isinstance(availability_zone, str):
            raise TypeError("Expected argument 'availability_zone' to be a str")
        pulumi.set(__self__, "availability_zone", availability_zone)
        if db_type and not isinstance(db_type, str):
            raise TypeError("Expected argument 'db_type' to be a str")
        pulumi.set(__self__, "db_type", db_type)
        if db_version and not isinstance(db_version, str):
            raise TypeError("Expected argument 'db_version' to be a str")
        pulumi.set(__self__, "db_version", db_version)
        if flavors and not isinstance(flavors, list):
            raise TypeError("Expected argument 'flavors' to be a list")
        pulumi.set(__self__, "flavors", flavors)
        if group_type and not isinstance(group_type, str):
            raise TypeError("Expected argument 'group_type' to be a str")
        pulumi.set(__self__, "group_type", group_type)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if instance_mode and not isinstance(instance_mode, str):
            raise TypeError("Expected argument 'instance_mode' to be a str")
        pulumi.set(__self__, "instance_mode", instance_mode)
        if memory and not isinstance(memory, int):
            raise TypeError("Expected argument 'memory' to be a int")
        pulumi.set(__self__, "memory", memory)
        if region and not isinstance(region, str):
            raise TypeError("Expected argument 'region' to be a str")
        pulumi.set(__self__, "region", region)
        if vcpus and not isinstance(vcpus, int):
            raise TypeError("Expected argument 'vcpus' to be a int")
        pulumi.set(__self__, "vcpus", vcpus)

    @property
    @pulumi.getter(name="availabilityZone")
    def availability_zone(self) -> Optional[builtins.str]:
        return pulumi.get(self, "availability_zone")

    @property
    @pulumi.getter(name="dbType")
    def db_type(self) -> builtins.str:
        return pulumi.get(self, "db_type")

    @property
    @pulumi.getter(name="dbVersion")
    def db_version(self) -> Optional[builtins.str]:
        return pulumi.get(self, "db_version")

    @property
    @pulumi.getter
    def flavors(self) -> Sequence['outputs.GetRdsFlavorsFlavorResult']:
        return pulumi.get(self, "flavors")

    @property
    @pulumi.getter(name="groupType")
    def group_type(self) -> Optional[builtins.str]:
        return pulumi.get(self, "group_type")

    @property
    @pulumi.getter
    def id(self) -> builtins.str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="instanceMode")
    def instance_mode(self) -> Optional[builtins.str]:
        return pulumi.get(self, "instance_mode")

    @property
    @pulumi.getter
    def memory(self) -> Optional[builtins.int]:
        return pulumi.get(self, "memory")

    @property
    @pulumi.getter
    def region(self) -> builtins.str:
        return pulumi.get(self, "region")

    @property
    @pulumi.getter
    def vcpus(self) -> Optional[builtins.int]:
        return pulumi.get(self, "vcpus")


class AwaitableGetRdsFlavorsResult(GetRdsFlavorsResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetRdsFlavorsResult(
            availability_zone=self.availability_zone,
            db_type=self.db_type,
            db_version=self.db_version,
            flavors=self.flavors,
            group_type=self.group_type,
            id=self.id,
            instance_mode=self.instance_mode,
            memory=self.memory,
            region=self.region,
            vcpus=self.vcpus)


def get_rds_flavors(availability_zone: Optional[builtins.str] = None,
                    db_type: Optional[builtins.str] = None,
                    db_version: Optional[builtins.str] = None,
                    group_type: Optional[builtins.str] = None,
                    instance_mode: Optional[builtins.str] = None,
                    memory: Optional[builtins.int] = None,
                    region: Optional[builtins.str] = None,
                    vcpus: Optional[builtins.int] = None,
                    opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetRdsFlavorsResult:
    """
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()
    __args__['availabilityZone'] = availability_zone
    __args__['dbType'] = db_type
    __args__['dbVersion'] = db_version
    __args__['groupType'] = group_type
    __args__['instanceMode'] = instance_mode
    __args__['memory'] = memory
    __args__['region'] = region
    __args__['vcpus'] = vcpus
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('sbercloud:index/getRdsFlavors:getRdsFlavors', __args__, opts=opts, typ=GetRdsFlavorsResult).value

    return AwaitableGetRdsFlavorsResult(
        availability_zone=pulumi.get(__ret__, 'availability_zone'),
        db_type=pulumi.get(__ret__, 'db_type'),
        db_version=pulumi.get(__ret__, 'db_version'),
        flavors=pulumi.get(__ret__, 'flavors'),
        group_type=pulumi.get(__ret__, 'group_type'),
        id=pulumi.get(__ret__, 'id'),
        instance_mode=pulumi.get(__ret__, 'instance_mode'),
        memory=pulumi.get(__ret__, 'memory'),
        region=pulumi.get(__ret__, 'region'),
        vcpus=pulumi.get(__ret__, 'vcpus'))
def get_rds_flavors_output(availability_zone: Optional[pulumi.Input[Optional[builtins.str]]] = None,
                           db_type: Optional[pulumi.Input[builtins.str]] = None,
                           db_version: Optional[pulumi.Input[Optional[builtins.str]]] = None,
                           group_type: Optional[pulumi.Input[Optional[builtins.str]]] = None,
                           instance_mode: Optional[pulumi.Input[Optional[builtins.str]]] = None,
                           memory: Optional[pulumi.Input[Optional[builtins.int]]] = None,
                           region: Optional[pulumi.Input[Optional[builtins.str]]] = None,
                           vcpus: Optional[pulumi.Input[Optional[builtins.int]]] = None,
                           opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetRdsFlavorsResult]:
    """
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()
    __args__['availabilityZone'] = availability_zone
    __args__['dbType'] = db_type
    __args__['dbVersion'] = db_version
    __args__['groupType'] = group_type
    __args__['instanceMode'] = instance_mode
    __args__['memory'] = memory
    __args__['region'] = region
    __args__['vcpus'] = vcpus
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('sbercloud:index/getRdsFlavors:getRdsFlavors', __args__, opts=opts, typ=GetRdsFlavorsResult)
    return __ret__.apply(lambda __response__: GetRdsFlavorsResult(
        availability_zone=pulumi.get(__response__, 'availability_zone'),
        db_type=pulumi.get(__response__, 'db_type'),
        db_version=pulumi.get(__response__, 'db_version'),
        flavors=pulumi.get(__response__, 'flavors'),
        group_type=pulumi.get(__response__, 'group_type'),
        id=pulumi.get(__response__, 'id'),
        instance_mode=pulumi.get(__response__, 'instance_mode'),
        memory=pulumi.get(__response__, 'memory'),
        region=pulumi.get(__response__, 'region'),
        vcpus=pulumi.get(__response__, 'vcpus')))
