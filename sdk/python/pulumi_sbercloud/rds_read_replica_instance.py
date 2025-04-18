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
from ._inputs import *

__all__ = ['RdsReadReplicaInstanceArgs', 'RdsReadReplicaInstance']

@pulumi.input_type
class RdsReadReplicaInstanceArgs:
    def __init__(__self__, *,
                 availability_zone: pulumi.Input[builtins.str],
                 flavor: pulumi.Input[builtins.str],
                 primary_instance_id: pulumi.Input[builtins.str],
                 volume: pulumi.Input['RdsReadReplicaInstanceVolumeArgs'],
                 auto_renew: Optional[pulumi.Input[builtins.str]] = None,
                 charging_mode: Optional[pulumi.Input[builtins.str]] = None,
                 db: Optional[pulumi.Input['RdsReadReplicaInstanceDbArgs']] = None,
                 description: Optional[pulumi.Input[builtins.str]] = None,
                 enterprise_project_id: Optional[pulumi.Input[builtins.str]] = None,
                 fixed_ip: Optional[pulumi.Input[builtins.str]] = None,
                 maintain_begin: Optional[pulumi.Input[builtins.str]] = None,
                 maintain_end: Optional[pulumi.Input[builtins.str]] = None,
                 name: Optional[pulumi.Input[builtins.str]] = None,
                 parameters: Optional[pulumi.Input[Sequence[pulumi.Input['RdsReadReplicaInstanceParameterArgs']]]] = None,
                 period: Optional[pulumi.Input[builtins.int]] = None,
                 period_unit: Optional[pulumi.Input[builtins.str]] = None,
                 region: Optional[pulumi.Input[builtins.str]] = None,
                 security_group_id: Optional[pulumi.Input[builtins.str]] = None,
                 ssl_enable: Optional[pulumi.Input[builtins.bool]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]]] = None):
        """
        The set of arguments for constructing a RdsReadReplicaInstance resource.
        """
        pulumi.set(__self__, "availability_zone", availability_zone)
        pulumi.set(__self__, "flavor", flavor)
        pulumi.set(__self__, "primary_instance_id", primary_instance_id)
        pulumi.set(__self__, "volume", volume)
        if auto_renew is not None:
            pulumi.set(__self__, "auto_renew", auto_renew)
        if charging_mode is not None:
            pulumi.set(__self__, "charging_mode", charging_mode)
        if db is not None:
            pulumi.set(__self__, "db", db)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if enterprise_project_id is not None:
            pulumi.set(__self__, "enterprise_project_id", enterprise_project_id)
        if fixed_ip is not None:
            pulumi.set(__self__, "fixed_ip", fixed_ip)
        if maintain_begin is not None:
            pulumi.set(__self__, "maintain_begin", maintain_begin)
        if maintain_end is not None:
            pulumi.set(__self__, "maintain_end", maintain_end)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if parameters is not None:
            pulumi.set(__self__, "parameters", parameters)
        if period is not None:
            pulumi.set(__self__, "period", period)
        if period_unit is not None:
            pulumi.set(__self__, "period_unit", period_unit)
        if region is not None:
            pulumi.set(__self__, "region", region)
        if security_group_id is not None:
            pulumi.set(__self__, "security_group_id", security_group_id)
        if ssl_enable is not None:
            pulumi.set(__self__, "ssl_enable", ssl_enable)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter(name="availabilityZone")
    def availability_zone(self) -> pulumi.Input[builtins.str]:
        return pulumi.get(self, "availability_zone")

    @availability_zone.setter
    def availability_zone(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "availability_zone", value)

    @property
    @pulumi.getter
    def flavor(self) -> pulumi.Input[builtins.str]:
        return pulumi.get(self, "flavor")

    @flavor.setter
    def flavor(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "flavor", value)

    @property
    @pulumi.getter(name="primaryInstanceId")
    def primary_instance_id(self) -> pulumi.Input[builtins.str]:
        return pulumi.get(self, "primary_instance_id")

    @primary_instance_id.setter
    def primary_instance_id(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "primary_instance_id", value)

    @property
    @pulumi.getter
    def volume(self) -> pulumi.Input['RdsReadReplicaInstanceVolumeArgs']:
        return pulumi.get(self, "volume")

    @volume.setter
    def volume(self, value: pulumi.Input['RdsReadReplicaInstanceVolumeArgs']):
        pulumi.set(self, "volume", value)

    @property
    @pulumi.getter(name="autoRenew")
    def auto_renew(self) -> Optional[pulumi.Input[builtins.str]]:
        return pulumi.get(self, "auto_renew")

    @auto_renew.setter
    def auto_renew(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "auto_renew", value)

    @property
    @pulumi.getter(name="chargingMode")
    def charging_mode(self) -> Optional[pulumi.Input[builtins.str]]:
        return pulumi.get(self, "charging_mode")

    @charging_mode.setter
    def charging_mode(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "charging_mode", value)

    @property
    @pulumi.getter
    def db(self) -> Optional[pulumi.Input['RdsReadReplicaInstanceDbArgs']]:
        return pulumi.get(self, "db")

    @db.setter
    def db(self, value: Optional[pulumi.Input['RdsReadReplicaInstanceDbArgs']]):
        pulumi.set(self, "db", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[builtins.str]]:
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="enterpriseProjectId")
    def enterprise_project_id(self) -> Optional[pulumi.Input[builtins.str]]:
        return pulumi.get(self, "enterprise_project_id")

    @enterprise_project_id.setter
    def enterprise_project_id(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "enterprise_project_id", value)

    @property
    @pulumi.getter(name="fixedIp")
    def fixed_ip(self) -> Optional[pulumi.Input[builtins.str]]:
        return pulumi.get(self, "fixed_ip")

    @fixed_ip.setter
    def fixed_ip(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "fixed_ip", value)

    @property
    @pulumi.getter(name="maintainBegin")
    def maintain_begin(self) -> Optional[pulumi.Input[builtins.str]]:
        return pulumi.get(self, "maintain_begin")

    @maintain_begin.setter
    def maintain_begin(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "maintain_begin", value)

    @property
    @pulumi.getter(name="maintainEnd")
    def maintain_end(self) -> Optional[pulumi.Input[builtins.str]]:
        return pulumi.get(self, "maintain_end")

    @maintain_end.setter
    def maintain_end(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "maintain_end", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[builtins.str]]:
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def parameters(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['RdsReadReplicaInstanceParameterArgs']]]]:
        return pulumi.get(self, "parameters")

    @parameters.setter
    def parameters(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['RdsReadReplicaInstanceParameterArgs']]]]):
        pulumi.set(self, "parameters", value)

    @property
    @pulumi.getter
    def period(self) -> Optional[pulumi.Input[builtins.int]]:
        return pulumi.get(self, "period")

    @period.setter
    def period(self, value: Optional[pulumi.Input[builtins.int]]):
        pulumi.set(self, "period", value)

    @property
    @pulumi.getter(name="periodUnit")
    def period_unit(self) -> Optional[pulumi.Input[builtins.str]]:
        return pulumi.get(self, "period_unit")

    @period_unit.setter
    def period_unit(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "period_unit", value)

    @property
    @pulumi.getter
    def region(self) -> Optional[pulumi.Input[builtins.str]]:
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "region", value)

    @property
    @pulumi.getter(name="securityGroupId")
    def security_group_id(self) -> Optional[pulumi.Input[builtins.str]]:
        return pulumi.get(self, "security_group_id")

    @security_group_id.setter
    def security_group_id(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "security_group_id", value)

    @property
    @pulumi.getter(name="sslEnable")
    def ssl_enable(self) -> Optional[pulumi.Input[builtins.bool]]:
        return pulumi.get(self, "ssl_enable")

    @ssl_enable.setter
    def ssl_enable(self, value: Optional[pulumi.Input[builtins.bool]]):
        pulumi.set(self, "ssl_enable", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]]]:
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]]]):
        pulumi.set(self, "tags", value)


@pulumi.input_type
class _RdsReadReplicaInstanceState:
    def __init__(__self__, *,
                 auto_renew: Optional[pulumi.Input[builtins.str]] = None,
                 availability_zone: Optional[pulumi.Input[builtins.str]] = None,
                 charging_mode: Optional[pulumi.Input[builtins.str]] = None,
                 db: Optional[pulumi.Input['RdsReadReplicaInstanceDbArgs']] = None,
                 description: Optional[pulumi.Input[builtins.str]] = None,
                 enterprise_project_id: Optional[pulumi.Input[builtins.str]] = None,
                 fixed_ip: Optional[pulumi.Input[builtins.str]] = None,
                 flavor: Optional[pulumi.Input[builtins.str]] = None,
                 maintain_begin: Optional[pulumi.Input[builtins.str]] = None,
                 maintain_end: Optional[pulumi.Input[builtins.str]] = None,
                 name: Optional[pulumi.Input[builtins.str]] = None,
                 parameters: Optional[pulumi.Input[Sequence[pulumi.Input['RdsReadReplicaInstanceParameterArgs']]]] = None,
                 period: Optional[pulumi.Input[builtins.int]] = None,
                 period_unit: Optional[pulumi.Input[builtins.str]] = None,
                 primary_instance_id: Optional[pulumi.Input[builtins.str]] = None,
                 private_ips: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]] = None,
                 public_ips: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]] = None,
                 region: Optional[pulumi.Input[builtins.str]] = None,
                 security_group_id: Optional[pulumi.Input[builtins.str]] = None,
                 ssl_enable: Optional[pulumi.Input[builtins.bool]] = None,
                 status: Optional[pulumi.Input[builtins.str]] = None,
                 subnet_id: Optional[pulumi.Input[builtins.str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]]] = None,
                 type: Optional[pulumi.Input[builtins.str]] = None,
                 volume: Optional[pulumi.Input['RdsReadReplicaInstanceVolumeArgs']] = None,
                 vpc_id: Optional[pulumi.Input[builtins.str]] = None):
        """
        Input properties used for looking up and filtering RdsReadReplicaInstance resources.
        """
        if auto_renew is not None:
            pulumi.set(__self__, "auto_renew", auto_renew)
        if availability_zone is not None:
            pulumi.set(__self__, "availability_zone", availability_zone)
        if charging_mode is not None:
            pulumi.set(__self__, "charging_mode", charging_mode)
        if db is not None:
            pulumi.set(__self__, "db", db)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if enterprise_project_id is not None:
            pulumi.set(__self__, "enterprise_project_id", enterprise_project_id)
        if fixed_ip is not None:
            pulumi.set(__self__, "fixed_ip", fixed_ip)
        if flavor is not None:
            pulumi.set(__self__, "flavor", flavor)
        if maintain_begin is not None:
            pulumi.set(__self__, "maintain_begin", maintain_begin)
        if maintain_end is not None:
            pulumi.set(__self__, "maintain_end", maintain_end)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if parameters is not None:
            pulumi.set(__self__, "parameters", parameters)
        if period is not None:
            pulumi.set(__self__, "period", period)
        if period_unit is not None:
            pulumi.set(__self__, "period_unit", period_unit)
        if primary_instance_id is not None:
            pulumi.set(__self__, "primary_instance_id", primary_instance_id)
        if private_ips is not None:
            pulumi.set(__self__, "private_ips", private_ips)
        if public_ips is not None:
            pulumi.set(__self__, "public_ips", public_ips)
        if region is not None:
            pulumi.set(__self__, "region", region)
        if security_group_id is not None:
            pulumi.set(__self__, "security_group_id", security_group_id)
        if ssl_enable is not None:
            pulumi.set(__self__, "ssl_enable", ssl_enable)
        if status is not None:
            pulumi.set(__self__, "status", status)
        if subnet_id is not None:
            pulumi.set(__self__, "subnet_id", subnet_id)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)
        if type is not None:
            pulumi.set(__self__, "type", type)
        if volume is not None:
            pulumi.set(__self__, "volume", volume)
        if vpc_id is not None:
            pulumi.set(__self__, "vpc_id", vpc_id)

    @property
    @pulumi.getter(name="autoRenew")
    def auto_renew(self) -> Optional[pulumi.Input[builtins.str]]:
        return pulumi.get(self, "auto_renew")

    @auto_renew.setter
    def auto_renew(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "auto_renew", value)

    @property
    @pulumi.getter(name="availabilityZone")
    def availability_zone(self) -> Optional[pulumi.Input[builtins.str]]:
        return pulumi.get(self, "availability_zone")

    @availability_zone.setter
    def availability_zone(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "availability_zone", value)

    @property
    @pulumi.getter(name="chargingMode")
    def charging_mode(self) -> Optional[pulumi.Input[builtins.str]]:
        return pulumi.get(self, "charging_mode")

    @charging_mode.setter
    def charging_mode(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "charging_mode", value)

    @property
    @pulumi.getter
    def db(self) -> Optional[pulumi.Input['RdsReadReplicaInstanceDbArgs']]:
        return pulumi.get(self, "db")

    @db.setter
    def db(self, value: Optional[pulumi.Input['RdsReadReplicaInstanceDbArgs']]):
        pulumi.set(self, "db", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[builtins.str]]:
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="enterpriseProjectId")
    def enterprise_project_id(self) -> Optional[pulumi.Input[builtins.str]]:
        return pulumi.get(self, "enterprise_project_id")

    @enterprise_project_id.setter
    def enterprise_project_id(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "enterprise_project_id", value)

    @property
    @pulumi.getter(name="fixedIp")
    def fixed_ip(self) -> Optional[pulumi.Input[builtins.str]]:
        return pulumi.get(self, "fixed_ip")

    @fixed_ip.setter
    def fixed_ip(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "fixed_ip", value)

    @property
    @pulumi.getter
    def flavor(self) -> Optional[pulumi.Input[builtins.str]]:
        return pulumi.get(self, "flavor")

    @flavor.setter
    def flavor(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "flavor", value)

    @property
    @pulumi.getter(name="maintainBegin")
    def maintain_begin(self) -> Optional[pulumi.Input[builtins.str]]:
        return pulumi.get(self, "maintain_begin")

    @maintain_begin.setter
    def maintain_begin(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "maintain_begin", value)

    @property
    @pulumi.getter(name="maintainEnd")
    def maintain_end(self) -> Optional[pulumi.Input[builtins.str]]:
        return pulumi.get(self, "maintain_end")

    @maintain_end.setter
    def maintain_end(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "maintain_end", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[builtins.str]]:
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def parameters(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['RdsReadReplicaInstanceParameterArgs']]]]:
        return pulumi.get(self, "parameters")

    @parameters.setter
    def parameters(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['RdsReadReplicaInstanceParameterArgs']]]]):
        pulumi.set(self, "parameters", value)

    @property
    @pulumi.getter
    def period(self) -> Optional[pulumi.Input[builtins.int]]:
        return pulumi.get(self, "period")

    @period.setter
    def period(self, value: Optional[pulumi.Input[builtins.int]]):
        pulumi.set(self, "period", value)

    @property
    @pulumi.getter(name="periodUnit")
    def period_unit(self) -> Optional[pulumi.Input[builtins.str]]:
        return pulumi.get(self, "period_unit")

    @period_unit.setter
    def period_unit(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "period_unit", value)

    @property
    @pulumi.getter(name="primaryInstanceId")
    def primary_instance_id(self) -> Optional[pulumi.Input[builtins.str]]:
        return pulumi.get(self, "primary_instance_id")

    @primary_instance_id.setter
    def primary_instance_id(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "primary_instance_id", value)

    @property
    @pulumi.getter(name="privateIps")
    def private_ips(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]]:
        return pulumi.get(self, "private_ips")

    @private_ips.setter
    def private_ips(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]]):
        pulumi.set(self, "private_ips", value)

    @property
    @pulumi.getter(name="publicIps")
    def public_ips(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]]:
        return pulumi.get(self, "public_ips")

    @public_ips.setter
    def public_ips(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]]):
        pulumi.set(self, "public_ips", value)

    @property
    @pulumi.getter
    def region(self) -> Optional[pulumi.Input[builtins.str]]:
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "region", value)

    @property
    @pulumi.getter(name="securityGroupId")
    def security_group_id(self) -> Optional[pulumi.Input[builtins.str]]:
        return pulumi.get(self, "security_group_id")

    @security_group_id.setter
    def security_group_id(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "security_group_id", value)

    @property
    @pulumi.getter(name="sslEnable")
    def ssl_enable(self) -> Optional[pulumi.Input[builtins.bool]]:
        return pulumi.get(self, "ssl_enable")

    @ssl_enable.setter
    def ssl_enable(self, value: Optional[pulumi.Input[builtins.bool]]):
        pulumi.set(self, "ssl_enable", value)

    @property
    @pulumi.getter
    def status(self) -> Optional[pulumi.Input[builtins.str]]:
        return pulumi.get(self, "status")

    @status.setter
    def status(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "status", value)

    @property
    @pulumi.getter(name="subnetId")
    def subnet_id(self) -> Optional[pulumi.Input[builtins.str]]:
        return pulumi.get(self, "subnet_id")

    @subnet_id.setter
    def subnet_id(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "subnet_id", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]]]:
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]]]):
        pulumi.set(self, "tags", value)

    @property
    @pulumi.getter
    def type(self) -> Optional[pulumi.Input[builtins.str]]:
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "type", value)

    @property
    @pulumi.getter
    def volume(self) -> Optional[pulumi.Input['RdsReadReplicaInstanceVolumeArgs']]:
        return pulumi.get(self, "volume")

    @volume.setter
    def volume(self, value: Optional[pulumi.Input['RdsReadReplicaInstanceVolumeArgs']]):
        pulumi.set(self, "volume", value)

    @property
    @pulumi.getter(name="vpcId")
    def vpc_id(self) -> Optional[pulumi.Input[builtins.str]]:
        return pulumi.get(self, "vpc_id")

    @vpc_id.setter
    def vpc_id(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "vpc_id", value)


class RdsReadReplicaInstance(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 auto_renew: Optional[pulumi.Input[builtins.str]] = None,
                 availability_zone: Optional[pulumi.Input[builtins.str]] = None,
                 charging_mode: Optional[pulumi.Input[builtins.str]] = None,
                 db: Optional[pulumi.Input[Union['RdsReadReplicaInstanceDbArgs', 'RdsReadReplicaInstanceDbArgsDict']]] = None,
                 description: Optional[pulumi.Input[builtins.str]] = None,
                 enterprise_project_id: Optional[pulumi.Input[builtins.str]] = None,
                 fixed_ip: Optional[pulumi.Input[builtins.str]] = None,
                 flavor: Optional[pulumi.Input[builtins.str]] = None,
                 maintain_begin: Optional[pulumi.Input[builtins.str]] = None,
                 maintain_end: Optional[pulumi.Input[builtins.str]] = None,
                 name: Optional[pulumi.Input[builtins.str]] = None,
                 parameters: Optional[pulumi.Input[Sequence[pulumi.Input[Union['RdsReadReplicaInstanceParameterArgs', 'RdsReadReplicaInstanceParameterArgsDict']]]]] = None,
                 period: Optional[pulumi.Input[builtins.int]] = None,
                 period_unit: Optional[pulumi.Input[builtins.str]] = None,
                 primary_instance_id: Optional[pulumi.Input[builtins.str]] = None,
                 region: Optional[pulumi.Input[builtins.str]] = None,
                 security_group_id: Optional[pulumi.Input[builtins.str]] = None,
                 ssl_enable: Optional[pulumi.Input[builtins.bool]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]]] = None,
                 volume: Optional[pulumi.Input[Union['RdsReadReplicaInstanceVolumeArgs', 'RdsReadReplicaInstanceVolumeArgsDict']]] = None,
                 __props__=None):
        """
        Create a RdsReadReplicaInstance resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: RdsReadReplicaInstanceArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a RdsReadReplicaInstance resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param RdsReadReplicaInstanceArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(RdsReadReplicaInstanceArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 auto_renew: Optional[pulumi.Input[builtins.str]] = None,
                 availability_zone: Optional[pulumi.Input[builtins.str]] = None,
                 charging_mode: Optional[pulumi.Input[builtins.str]] = None,
                 db: Optional[pulumi.Input[Union['RdsReadReplicaInstanceDbArgs', 'RdsReadReplicaInstanceDbArgsDict']]] = None,
                 description: Optional[pulumi.Input[builtins.str]] = None,
                 enterprise_project_id: Optional[pulumi.Input[builtins.str]] = None,
                 fixed_ip: Optional[pulumi.Input[builtins.str]] = None,
                 flavor: Optional[pulumi.Input[builtins.str]] = None,
                 maintain_begin: Optional[pulumi.Input[builtins.str]] = None,
                 maintain_end: Optional[pulumi.Input[builtins.str]] = None,
                 name: Optional[pulumi.Input[builtins.str]] = None,
                 parameters: Optional[pulumi.Input[Sequence[pulumi.Input[Union['RdsReadReplicaInstanceParameterArgs', 'RdsReadReplicaInstanceParameterArgsDict']]]]] = None,
                 period: Optional[pulumi.Input[builtins.int]] = None,
                 period_unit: Optional[pulumi.Input[builtins.str]] = None,
                 primary_instance_id: Optional[pulumi.Input[builtins.str]] = None,
                 region: Optional[pulumi.Input[builtins.str]] = None,
                 security_group_id: Optional[pulumi.Input[builtins.str]] = None,
                 ssl_enable: Optional[pulumi.Input[builtins.bool]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]]] = None,
                 volume: Optional[pulumi.Input[Union['RdsReadReplicaInstanceVolumeArgs', 'RdsReadReplicaInstanceVolumeArgsDict']]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = RdsReadReplicaInstanceArgs.__new__(RdsReadReplicaInstanceArgs)

            __props__.__dict__["auto_renew"] = auto_renew
            if availability_zone is None and not opts.urn:
                raise TypeError("Missing required property 'availability_zone'")
            __props__.__dict__["availability_zone"] = availability_zone
            __props__.__dict__["charging_mode"] = charging_mode
            __props__.__dict__["db"] = db
            __props__.__dict__["description"] = description
            __props__.__dict__["enterprise_project_id"] = enterprise_project_id
            __props__.__dict__["fixed_ip"] = fixed_ip
            if flavor is None and not opts.urn:
                raise TypeError("Missing required property 'flavor'")
            __props__.__dict__["flavor"] = flavor
            __props__.__dict__["maintain_begin"] = maintain_begin
            __props__.__dict__["maintain_end"] = maintain_end
            __props__.__dict__["name"] = name
            __props__.__dict__["parameters"] = parameters
            __props__.__dict__["period"] = period
            __props__.__dict__["period_unit"] = period_unit
            if primary_instance_id is None and not opts.urn:
                raise TypeError("Missing required property 'primary_instance_id'")
            __props__.__dict__["primary_instance_id"] = primary_instance_id
            __props__.__dict__["region"] = region
            __props__.__dict__["security_group_id"] = security_group_id
            __props__.__dict__["ssl_enable"] = ssl_enable
            __props__.__dict__["tags"] = tags
            if volume is None and not opts.urn:
                raise TypeError("Missing required property 'volume'")
            __props__.__dict__["volume"] = volume
            __props__.__dict__["private_ips"] = None
            __props__.__dict__["public_ips"] = None
            __props__.__dict__["status"] = None
            __props__.__dict__["subnet_id"] = None
            __props__.__dict__["type"] = None
            __props__.__dict__["vpc_id"] = None
        super(RdsReadReplicaInstance, __self__).__init__(
            'sbercloud:index/rdsReadReplicaInstance:RdsReadReplicaInstance',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            auto_renew: Optional[pulumi.Input[builtins.str]] = None,
            availability_zone: Optional[pulumi.Input[builtins.str]] = None,
            charging_mode: Optional[pulumi.Input[builtins.str]] = None,
            db: Optional[pulumi.Input[Union['RdsReadReplicaInstanceDbArgs', 'RdsReadReplicaInstanceDbArgsDict']]] = None,
            description: Optional[pulumi.Input[builtins.str]] = None,
            enterprise_project_id: Optional[pulumi.Input[builtins.str]] = None,
            fixed_ip: Optional[pulumi.Input[builtins.str]] = None,
            flavor: Optional[pulumi.Input[builtins.str]] = None,
            maintain_begin: Optional[pulumi.Input[builtins.str]] = None,
            maintain_end: Optional[pulumi.Input[builtins.str]] = None,
            name: Optional[pulumi.Input[builtins.str]] = None,
            parameters: Optional[pulumi.Input[Sequence[pulumi.Input[Union['RdsReadReplicaInstanceParameterArgs', 'RdsReadReplicaInstanceParameterArgsDict']]]]] = None,
            period: Optional[pulumi.Input[builtins.int]] = None,
            period_unit: Optional[pulumi.Input[builtins.str]] = None,
            primary_instance_id: Optional[pulumi.Input[builtins.str]] = None,
            private_ips: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]] = None,
            public_ips: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]] = None,
            region: Optional[pulumi.Input[builtins.str]] = None,
            security_group_id: Optional[pulumi.Input[builtins.str]] = None,
            ssl_enable: Optional[pulumi.Input[builtins.bool]] = None,
            status: Optional[pulumi.Input[builtins.str]] = None,
            subnet_id: Optional[pulumi.Input[builtins.str]] = None,
            tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]]] = None,
            type: Optional[pulumi.Input[builtins.str]] = None,
            volume: Optional[pulumi.Input[Union['RdsReadReplicaInstanceVolumeArgs', 'RdsReadReplicaInstanceVolumeArgsDict']]] = None,
            vpc_id: Optional[pulumi.Input[builtins.str]] = None) -> 'RdsReadReplicaInstance':
        """
        Get an existing RdsReadReplicaInstance resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _RdsReadReplicaInstanceState.__new__(_RdsReadReplicaInstanceState)

        __props__.__dict__["auto_renew"] = auto_renew
        __props__.__dict__["availability_zone"] = availability_zone
        __props__.__dict__["charging_mode"] = charging_mode
        __props__.__dict__["db"] = db
        __props__.__dict__["description"] = description
        __props__.__dict__["enterprise_project_id"] = enterprise_project_id
        __props__.__dict__["fixed_ip"] = fixed_ip
        __props__.__dict__["flavor"] = flavor
        __props__.__dict__["maintain_begin"] = maintain_begin
        __props__.__dict__["maintain_end"] = maintain_end
        __props__.__dict__["name"] = name
        __props__.__dict__["parameters"] = parameters
        __props__.__dict__["period"] = period
        __props__.__dict__["period_unit"] = period_unit
        __props__.__dict__["primary_instance_id"] = primary_instance_id
        __props__.__dict__["private_ips"] = private_ips
        __props__.__dict__["public_ips"] = public_ips
        __props__.__dict__["region"] = region
        __props__.__dict__["security_group_id"] = security_group_id
        __props__.__dict__["ssl_enable"] = ssl_enable
        __props__.__dict__["status"] = status
        __props__.__dict__["subnet_id"] = subnet_id
        __props__.__dict__["tags"] = tags
        __props__.__dict__["type"] = type
        __props__.__dict__["volume"] = volume
        __props__.__dict__["vpc_id"] = vpc_id
        return RdsReadReplicaInstance(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="autoRenew")
    def auto_renew(self) -> pulumi.Output[Optional[builtins.str]]:
        return pulumi.get(self, "auto_renew")

    @property
    @pulumi.getter(name="availabilityZone")
    def availability_zone(self) -> pulumi.Output[builtins.str]:
        return pulumi.get(self, "availability_zone")

    @property
    @pulumi.getter(name="chargingMode")
    def charging_mode(self) -> pulumi.Output[builtins.str]:
        return pulumi.get(self, "charging_mode")

    @property
    @pulumi.getter
    def db(self) -> pulumi.Output['outputs.RdsReadReplicaInstanceDb']:
        return pulumi.get(self, "db")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[builtins.str]]:
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="enterpriseProjectId")
    def enterprise_project_id(self) -> pulumi.Output[builtins.str]:
        return pulumi.get(self, "enterprise_project_id")

    @property
    @pulumi.getter(name="fixedIp")
    def fixed_ip(self) -> pulumi.Output[builtins.str]:
        return pulumi.get(self, "fixed_ip")

    @property
    @pulumi.getter
    def flavor(self) -> pulumi.Output[builtins.str]:
        return pulumi.get(self, "flavor")

    @property
    @pulumi.getter(name="maintainBegin")
    def maintain_begin(self) -> pulumi.Output[builtins.str]:
        return pulumi.get(self, "maintain_begin")

    @property
    @pulumi.getter(name="maintainEnd")
    def maintain_end(self) -> pulumi.Output[builtins.str]:
        return pulumi.get(self, "maintain_end")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[builtins.str]:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def parameters(self) -> pulumi.Output[Sequence['outputs.RdsReadReplicaInstanceParameter']]:
        return pulumi.get(self, "parameters")

    @property
    @pulumi.getter
    def period(self) -> pulumi.Output[Optional[builtins.int]]:
        return pulumi.get(self, "period")

    @property
    @pulumi.getter(name="periodUnit")
    def period_unit(self) -> pulumi.Output[Optional[builtins.str]]:
        return pulumi.get(self, "period_unit")

    @property
    @pulumi.getter(name="primaryInstanceId")
    def primary_instance_id(self) -> pulumi.Output[builtins.str]:
        return pulumi.get(self, "primary_instance_id")

    @property
    @pulumi.getter(name="privateIps")
    def private_ips(self) -> pulumi.Output[Sequence[builtins.str]]:
        return pulumi.get(self, "private_ips")

    @property
    @pulumi.getter(name="publicIps")
    def public_ips(self) -> pulumi.Output[Sequence[builtins.str]]:
        return pulumi.get(self, "public_ips")

    @property
    @pulumi.getter
    def region(self) -> pulumi.Output[builtins.str]:
        return pulumi.get(self, "region")

    @property
    @pulumi.getter(name="securityGroupId")
    def security_group_id(self) -> pulumi.Output[builtins.str]:
        return pulumi.get(self, "security_group_id")

    @property
    @pulumi.getter(name="sslEnable")
    def ssl_enable(self) -> pulumi.Output[builtins.bool]:
        return pulumi.get(self, "ssl_enable")

    @property
    @pulumi.getter
    def status(self) -> pulumi.Output[builtins.str]:
        return pulumi.get(self, "status")

    @property
    @pulumi.getter(name="subnetId")
    def subnet_id(self) -> pulumi.Output[builtins.str]:
        return pulumi.get(self, "subnet_id")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Mapping[str, builtins.str]]]:
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter
    def type(self) -> pulumi.Output[builtins.str]:
        return pulumi.get(self, "type")

    @property
    @pulumi.getter
    def volume(self) -> pulumi.Output['outputs.RdsReadReplicaInstanceVolume']:
        return pulumi.get(self, "volume")

    @property
    @pulumi.getter(name="vpcId")
    def vpc_id(self) -> pulumi.Output[builtins.str]:
        return pulumi.get(self, "vpc_id")

