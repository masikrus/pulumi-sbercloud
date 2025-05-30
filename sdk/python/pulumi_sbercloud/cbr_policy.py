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

__all__ = ['CbrPolicyArgs', 'CbrPolicy']

@pulumi.input_type
class CbrPolicyArgs:
    def __init__(__self__, *,
                 backup_cycle: pulumi.Input['CbrPolicyBackupCycleArgs'],
                 type: pulumi.Input[builtins.str],
                 backup_quantity: Optional[pulumi.Input[builtins.int]] = None,
                 destination_project_id: Optional[pulumi.Input[builtins.str]] = None,
                 destination_region: Optional[pulumi.Input[builtins.str]] = None,
                 enable_acceleration: Optional[pulumi.Input[builtins.bool]] = None,
                 enabled: Optional[pulumi.Input[builtins.bool]] = None,
                 long_term_retention: Optional[pulumi.Input['CbrPolicyLongTermRetentionArgs']] = None,
                 name: Optional[pulumi.Input[builtins.str]] = None,
                 region: Optional[pulumi.Input[builtins.str]] = None,
                 time_period: Optional[pulumi.Input[builtins.int]] = None,
                 time_zone: Optional[pulumi.Input[builtins.str]] = None):
        """
        The set of arguments for constructing a CbrPolicy resource.
        :param pulumi.Input['CbrPolicyBackupCycleArgs'] backup_cycle: The scheduling rule for the CBR policy backup execution.
        :param pulumi.Input[builtins.str] type: The protection type of the CBR policy.
        :param pulumi.Input[builtins.int] backup_quantity: The maximum number of retained backups.
        :param pulumi.Input[builtins.str] destination_project_id: The ID of the replication destination project.
        :param pulumi.Input[builtins.str] destination_region: The name of the replication destination region.
        :param pulumi.Input[builtins.bool] enable_acceleration: Whether to enable the acceleration function to shorten the replication time for cross-region
        :param pulumi.Input[builtins.bool] enabled: Whether to enable the CBR policy.
        :param pulumi.Input['CbrPolicyLongTermRetentionArgs'] long_term_retention: The long-term retention rules.
        :param pulumi.Input[builtins.str] name: The policy name.
        :param pulumi.Input[builtins.str] region: The region where the policy is located.
        :param pulumi.Input[builtins.int] time_period: The duration (in days) for retained backups.
        :param pulumi.Input[builtins.str] time_zone: The UTC time zone.
        """
        pulumi.set(__self__, "backup_cycle", backup_cycle)
        pulumi.set(__self__, "type", type)
        if backup_quantity is not None:
            pulumi.set(__self__, "backup_quantity", backup_quantity)
        if destination_project_id is not None:
            pulumi.set(__self__, "destination_project_id", destination_project_id)
        if destination_region is not None:
            pulumi.set(__self__, "destination_region", destination_region)
        if enable_acceleration is not None:
            pulumi.set(__self__, "enable_acceleration", enable_acceleration)
        if enabled is not None:
            pulumi.set(__self__, "enabled", enabled)
        if long_term_retention is not None:
            pulumi.set(__self__, "long_term_retention", long_term_retention)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if region is not None:
            pulumi.set(__self__, "region", region)
        if time_period is not None:
            pulumi.set(__self__, "time_period", time_period)
        if time_zone is not None:
            pulumi.set(__self__, "time_zone", time_zone)

    @property
    @pulumi.getter(name="backupCycle")
    def backup_cycle(self) -> pulumi.Input['CbrPolicyBackupCycleArgs']:
        """
        The scheduling rule for the CBR policy backup execution.
        """
        return pulumi.get(self, "backup_cycle")

    @backup_cycle.setter
    def backup_cycle(self, value: pulumi.Input['CbrPolicyBackupCycleArgs']):
        pulumi.set(self, "backup_cycle", value)

    @property
    @pulumi.getter
    def type(self) -> pulumi.Input[builtins.str]:
        """
        The protection type of the CBR policy.
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "type", value)

    @property
    @pulumi.getter(name="backupQuantity")
    def backup_quantity(self) -> Optional[pulumi.Input[builtins.int]]:
        """
        The maximum number of retained backups.
        """
        return pulumi.get(self, "backup_quantity")

    @backup_quantity.setter
    def backup_quantity(self, value: Optional[pulumi.Input[builtins.int]]):
        pulumi.set(self, "backup_quantity", value)

    @property
    @pulumi.getter(name="destinationProjectId")
    def destination_project_id(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The ID of the replication destination project.
        """
        return pulumi.get(self, "destination_project_id")

    @destination_project_id.setter
    def destination_project_id(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "destination_project_id", value)

    @property
    @pulumi.getter(name="destinationRegion")
    def destination_region(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The name of the replication destination region.
        """
        return pulumi.get(self, "destination_region")

    @destination_region.setter
    def destination_region(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "destination_region", value)

    @property
    @pulumi.getter(name="enableAcceleration")
    def enable_acceleration(self) -> Optional[pulumi.Input[builtins.bool]]:
        """
        Whether to enable the acceleration function to shorten the replication time for cross-region
        """
        return pulumi.get(self, "enable_acceleration")

    @enable_acceleration.setter
    def enable_acceleration(self, value: Optional[pulumi.Input[builtins.bool]]):
        pulumi.set(self, "enable_acceleration", value)

    @property
    @pulumi.getter
    def enabled(self) -> Optional[pulumi.Input[builtins.bool]]:
        """
        Whether to enable the CBR policy.
        """
        return pulumi.get(self, "enabled")

    @enabled.setter
    def enabled(self, value: Optional[pulumi.Input[builtins.bool]]):
        pulumi.set(self, "enabled", value)

    @property
    @pulumi.getter(name="longTermRetention")
    def long_term_retention(self) -> Optional[pulumi.Input['CbrPolicyLongTermRetentionArgs']]:
        """
        The long-term retention rules.
        """
        return pulumi.get(self, "long_term_retention")

    @long_term_retention.setter
    def long_term_retention(self, value: Optional[pulumi.Input['CbrPolicyLongTermRetentionArgs']]):
        pulumi.set(self, "long_term_retention", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The policy name.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def region(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The region where the policy is located.
        """
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "region", value)

    @property
    @pulumi.getter(name="timePeriod")
    def time_period(self) -> Optional[pulumi.Input[builtins.int]]:
        """
        The duration (in days) for retained backups.
        """
        return pulumi.get(self, "time_period")

    @time_period.setter
    def time_period(self, value: Optional[pulumi.Input[builtins.int]]):
        pulumi.set(self, "time_period", value)

    @property
    @pulumi.getter(name="timeZone")
    def time_zone(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The UTC time zone.
        """
        return pulumi.get(self, "time_zone")

    @time_zone.setter
    def time_zone(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "time_zone", value)


@pulumi.input_type
class _CbrPolicyState:
    def __init__(__self__, *,
                 backup_cycle: Optional[pulumi.Input['CbrPolicyBackupCycleArgs']] = None,
                 backup_quantity: Optional[pulumi.Input[builtins.int]] = None,
                 destination_project_id: Optional[pulumi.Input[builtins.str]] = None,
                 destination_region: Optional[pulumi.Input[builtins.str]] = None,
                 enable_acceleration: Optional[pulumi.Input[builtins.bool]] = None,
                 enabled: Optional[pulumi.Input[builtins.bool]] = None,
                 long_term_retention: Optional[pulumi.Input['CbrPolicyLongTermRetentionArgs']] = None,
                 name: Optional[pulumi.Input[builtins.str]] = None,
                 region: Optional[pulumi.Input[builtins.str]] = None,
                 time_period: Optional[pulumi.Input[builtins.int]] = None,
                 time_zone: Optional[pulumi.Input[builtins.str]] = None,
                 type: Optional[pulumi.Input[builtins.str]] = None):
        """
        Input properties used for looking up and filtering CbrPolicy resources.
        :param pulumi.Input['CbrPolicyBackupCycleArgs'] backup_cycle: The scheduling rule for the CBR policy backup execution.
        :param pulumi.Input[builtins.int] backup_quantity: The maximum number of retained backups.
        :param pulumi.Input[builtins.str] destination_project_id: The ID of the replication destination project.
        :param pulumi.Input[builtins.str] destination_region: The name of the replication destination region.
        :param pulumi.Input[builtins.bool] enable_acceleration: Whether to enable the acceleration function to shorten the replication time for cross-region
        :param pulumi.Input[builtins.bool] enabled: Whether to enable the CBR policy.
        :param pulumi.Input['CbrPolicyLongTermRetentionArgs'] long_term_retention: The long-term retention rules.
        :param pulumi.Input[builtins.str] name: The policy name.
        :param pulumi.Input[builtins.str] region: The region where the policy is located.
        :param pulumi.Input[builtins.int] time_period: The duration (in days) for retained backups.
        :param pulumi.Input[builtins.str] time_zone: The UTC time zone.
        :param pulumi.Input[builtins.str] type: The protection type of the CBR policy.
        """
        if backup_cycle is not None:
            pulumi.set(__self__, "backup_cycle", backup_cycle)
        if backup_quantity is not None:
            pulumi.set(__self__, "backup_quantity", backup_quantity)
        if destination_project_id is not None:
            pulumi.set(__self__, "destination_project_id", destination_project_id)
        if destination_region is not None:
            pulumi.set(__self__, "destination_region", destination_region)
        if enable_acceleration is not None:
            pulumi.set(__self__, "enable_acceleration", enable_acceleration)
        if enabled is not None:
            pulumi.set(__self__, "enabled", enabled)
        if long_term_retention is not None:
            pulumi.set(__self__, "long_term_retention", long_term_retention)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if region is not None:
            pulumi.set(__self__, "region", region)
        if time_period is not None:
            pulumi.set(__self__, "time_period", time_period)
        if time_zone is not None:
            pulumi.set(__self__, "time_zone", time_zone)
        if type is not None:
            pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter(name="backupCycle")
    def backup_cycle(self) -> Optional[pulumi.Input['CbrPolicyBackupCycleArgs']]:
        """
        The scheduling rule for the CBR policy backup execution.
        """
        return pulumi.get(self, "backup_cycle")

    @backup_cycle.setter
    def backup_cycle(self, value: Optional[pulumi.Input['CbrPolicyBackupCycleArgs']]):
        pulumi.set(self, "backup_cycle", value)

    @property
    @pulumi.getter(name="backupQuantity")
    def backup_quantity(self) -> Optional[pulumi.Input[builtins.int]]:
        """
        The maximum number of retained backups.
        """
        return pulumi.get(self, "backup_quantity")

    @backup_quantity.setter
    def backup_quantity(self, value: Optional[pulumi.Input[builtins.int]]):
        pulumi.set(self, "backup_quantity", value)

    @property
    @pulumi.getter(name="destinationProjectId")
    def destination_project_id(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The ID of the replication destination project.
        """
        return pulumi.get(self, "destination_project_id")

    @destination_project_id.setter
    def destination_project_id(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "destination_project_id", value)

    @property
    @pulumi.getter(name="destinationRegion")
    def destination_region(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The name of the replication destination region.
        """
        return pulumi.get(self, "destination_region")

    @destination_region.setter
    def destination_region(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "destination_region", value)

    @property
    @pulumi.getter(name="enableAcceleration")
    def enable_acceleration(self) -> Optional[pulumi.Input[builtins.bool]]:
        """
        Whether to enable the acceleration function to shorten the replication time for cross-region
        """
        return pulumi.get(self, "enable_acceleration")

    @enable_acceleration.setter
    def enable_acceleration(self, value: Optional[pulumi.Input[builtins.bool]]):
        pulumi.set(self, "enable_acceleration", value)

    @property
    @pulumi.getter
    def enabled(self) -> Optional[pulumi.Input[builtins.bool]]:
        """
        Whether to enable the CBR policy.
        """
        return pulumi.get(self, "enabled")

    @enabled.setter
    def enabled(self, value: Optional[pulumi.Input[builtins.bool]]):
        pulumi.set(self, "enabled", value)

    @property
    @pulumi.getter(name="longTermRetention")
    def long_term_retention(self) -> Optional[pulumi.Input['CbrPolicyLongTermRetentionArgs']]:
        """
        The long-term retention rules.
        """
        return pulumi.get(self, "long_term_retention")

    @long_term_retention.setter
    def long_term_retention(self, value: Optional[pulumi.Input['CbrPolicyLongTermRetentionArgs']]):
        pulumi.set(self, "long_term_retention", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The policy name.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def region(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The region where the policy is located.
        """
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "region", value)

    @property
    @pulumi.getter(name="timePeriod")
    def time_period(self) -> Optional[pulumi.Input[builtins.int]]:
        """
        The duration (in days) for retained backups.
        """
        return pulumi.get(self, "time_period")

    @time_period.setter
    def time_period(self, value: Optional[pulumi.Input[builtins.int]]):
        pulumi.set(self, "time_period", value)

    @property
    @pulumi.getter(name="timeZone")
    def time_zone(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The UTC time zone.
        """
        return pulumi.get(self, "time_zone")

    @time_zone.setter
    def time_zone(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "time_zone", value)

    @property
    @pulumi.getter
    def type(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The protection type of the CBR policy.
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "type", value)


class CbrPolicy(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 backup_cycle: Optional[pulumi.Input[Union['CbrPolicyBackupCycleArgs', 'CbrPolicyBackupCycleArgsDict']]] = None,
                 backup_quantity: Optional[pulumi.Input[builtins.int]] = None,
                 destination_project_id: Optional[pulumi.Input[builtins.str]] = None,
                 destination_region: Optional[pulumi.Input[builtins.str]] = None,
                 enable_acceleration: Optional[pulumi.Input[builtins.bool]] = None,
                 enabled: Optional[pulumi.Input[builtins.bool]] = None,
                 long_term_retention: Optional[pulumi.Input[Union['CbrPolicyLongTermRetentionArgs', 'CbrPolicyLongTermRetentionArgsDict']]] = None,
                 name: Optional[pulumi.Input[builtins.str]] = None,
                 region: Optional[pulumi.Input[builtins.str]] = None,
                 time_period: Optional[pulumi.Input[builtins.int]] = None,
                 time_zone: Optional[pulumi.Input[builtins.str]] = None,
                 type: Optional[pulumi.Input[builtins.str]] = None,
                 __props__=None):
        """
        Create a CbrPolicy resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Union['CbrPolicyBackupCycleArgs', 'CbrPolicyBackupCycleArgsDict']] backup_cycle: The scheduling rule for the CBR policy backup execution.
        :param pulumi.Input[builtins.int] backup_quantity: The maximum number of retained backups.
        :param pulumi.Input[builtins.str] destination_project_id: The ID of the replication destination project.
        :param pulumi.Input[builtins.str] destination_region: The name of the replication destination region.
        :param pulumi.Input[builtins.bool] enable_acceleration: Whether to enable the acceleration function to shorten the replication time for cross-region
        :param pulumi.Input[builtins.bool] enabled: Whether to enable the CBR policy.
        :param pulumi.Input[Union['CbrPolicyLongTermRetentionArgs', 'CbrPolicyLongTermRetentionArgsDict']] long_term_retention: The long-term retention rules.
        :param pulumi.Input[builtins.str] name: The policy name.
        :param pulumi.Input[builtins.str] region: The region where the policy is located.
        :param pulumi.Input[builtins.int] time_period: The duration (in days) for retained backups.
        :param pulumi.Input[builtins.str] time_zone: The UTC time zone.
        :param pulumi.Input[builtins.str] type: The protection type of the CBR policy.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: CbrPolicyArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a CbrPolicy resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param CbrPolicyArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(CbrPolicyArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 backup_cycle: Optional[pulumi.Input[Union['CbrPolicyBackupCycleArgs', 'CbrPolicyBackupCycleArgsDict']]] = None,
                 backup_quantity: Optional[pulumi.Input[builtins.int]] = None,
                 destination_project_id: Optional[pulumi.Input[builtins.str]] = None,
                 destination_region: Optional[pulumi.Input[builtins.str]] = None,
                 enable_acceleration: Optional[pulumi.Input[builtins.bool]] = None,
                 enabled: Optional[pulumi.Input[builtins.bool]] = None,
                 long_term_retention: Optional[pulumi.Input[Union['CbrPolicyLongTermRetentionArgs', 'CbrPolicyLongTermRetentionArgsDict']]] = None,
                 name: Optional[pulumi.Input[builtins.str]] = None,
                 region: Optional[pulumi.Input[builtins.str]] = None,
                 time_period: Optional[pulumi.Input[builtins.int]] = None,
                 time_zone: Optional[pulumi.Input[builtins.str]] = None,
                 type: Optional[pulumi.Input[builtins.str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = CbrPolicyArgs.__new__(CbrPolicyArgs)

            if backup_cycle is None and not opts.urn:
                raise TypeError("Missing required property 'backup_cycle'")
            __props__.__dict__["backup_cycle"] = backup_cycle
            __props__.__dict__["backup_quantity"] = backup_quantity
            __props__.__dict__["destination_project_id"] = destination_project_id
            __props__.__dict__["destination_region"] = destination_region
            __props__.__dict__["enable_acceleration"] = enable_acceleration
            __props__.__dict__["enabled"] = enabled
            __props__.__dict__["long_term_retention"] = long_term_retention
            __props__.__dict__["name"] = name
            __props__.__dict__["region"] = region
            __props__.__dict__["time_period"] = time_period
            __props__.__dict__["time_zone"] = time_zone
            if type is None and not opts.urn:
                raise TypeError("Missing required property 'type'")
            __props__.__dict__["type"] = type
        super(CbrPolicy, __self__).__init__(
            'sbercloud:index/cbrPolicy:CbrPolicy',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            backup_cycle: Optional[pulumi.Input[Union['CbrPolicyBackupCycleArgs', 'CbrPolicyBackupCycleArgsDict']]] = None,
            backup_quantity: Optional[pulumi.Input[builtins.int]] = None,
            destination_project_id: Optional[pulumi.Input[builtins.str]] = None,
            destination_region: Optional[pulumi.Input[builtins.str]] = None,
            enable_acceleration: Optional[pulumi.Input[builtins.bool]] = None,
            enabled: Optional[pulumi.Input[builtins.bool]] = None,
            long_term_retention: Optional[pulumi.Input[Union['CbrPolicyLongTermRetentionArgs', 'CbrPolicyLongTermRetentionArgsDict']]] = None,
            name: Optional[pulumi.Input[builtins.str]] = None,
            region: Optional[pulumi.Input[builtins.str]] = None,
            time_period: Optional[pulumi.Input[builtins.int]] = None,
            time_zone: Optional[pulumi.Input[builtins.str]] = None,
            type: Optional[pulumi.Input[builtins.str]] = None) -> 'CbrPolicy':
        """
        Get an existing CbrPolicy resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Union['CbrPolicyBackupCycleArgs', 'CbrPolicyBackupCycleArgsDict']] backup_cycle: The scheduling rule for the CBR policy backup execution.
        :param pulumi.Input[builtins.int] backup_quantity: The maximum number of retained backups.
        :param pulumi.Input[builtins.str] destination_project_id: The ID of the replication destination project.
        :param pulumi.Input[builtins.str] destination_region: The name of the replication destination region.
        :param pulumi.Input[builtins.bool] enable_acceleration: Whether to enable the acceleration function to shorten the replication time for cross-region
        :param pulumi.Input[builtins.bool] enabled: Whether to enable the CBR policy.
        :param pulumi.Input[Union['CbrPolicyLongTermRetentionArgs', 'CbrPolicyLongTermRetentionArgsDict']] long_term_retention: The long-term retention rules.
        :param pulumi.Input[builtins.str] name: The policy name.
        :param pulumi.Input[builtins.str] region: The region where the policy is located.
        :param pulumi.Input[builtins.int] time_period: The duration (in days) for retained backups.
        :param pulumi.Input[builtins.str] time_zone: The UTC time zone.
        :param pulumi.Input[builtins.str] type: The protection type of the CBR policy.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _CbrPolicyState.__new__(_CbrPolicyState)

        __props__.__dict__["backup_cycle"] = backup_cycle
        __props__.__dict__["backup_quantity"] = backup_quantity
        __props__.__dict__["destination_project_id"] = destination_project_id
        __props__.__dict__["destination_region"] = destination_region
        __props__.__dict__["enable_acceleration"] = enable_acceleration
        __props__.__dict__["enabled"] = enabled
        __props__.__dict__["long_term_retention"] = long_term_retention
        __props__.__dict__["name"] = name
        __props__.__dict__["region"] = region
        __props__.__dict__["time_period"] = time_period
        __props__.__dict__["time_zone"] = time_zone
        __props__.__dict__["type"] = type
        return CbrPolicy(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="backupCycle")
    def backup_cycle(self) -> pulumi.Output['outputs.CbrPolicyBackupCycle']:
        """
        The scheduling rule for the CBR policy backup execution.
        """
        return pulumi.get(self, "backup_cycle")

    @property
    @pulumi.getter(name="backupQuantity")
    def backup_quantity(self) -> pulumi.Output[Optional[builtins.int]]:
        """
        The maximum number of retained backups.
        """
        return pulumi.get(self, "backup_quantity")

    @property
    @pulumi.getter(name="destinationProjectId")
    def destination_project_id(self) -> pulumi.Output[Optional[builtins.str]]:
        """
        The ID of the replication destination project.
        """
        return pulumi.get(self, "destination_project_id")

    @property
    @pulumi.getter(name="destinationRegion")
    def destination_region(self) -> pulumi.Output[Optional[builtins.str]]:
        """
        The name of the replication destination region.
        """
        return pulumi.get(self, "destination_region")

    @property
    @pulumi.getter(name="enableAcceleration")
    def enable_acceleration(self) -> pulumi.Output[Optional[builtins.bool]]:
        """
        Whether to enable the acceleration function to shorten the replication time for cross-region
        """
        return pulumi.get(self, "enable_acceleration")

    @property
    @pulumi.getter
    def enabled(self) -> pulumi.Output[Optional[builtins.bool]]:
        """
        Whether to enable the CBR policy.
        """
        return pulumi.get(self, "enabled")

    @property
    @pulumi.getter(name="longTermRetention")
    def long_term_retention(self) -> pulumi.Output[Optional['outputs.CbrPolicyLongTermRetention']]:
        """
        The long-term retention rules.
        """
        return pulumi.get(self, "long_term_retention")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[builtins.str]:
        """
        The policy name.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def region(self) -> pulumi.Output[builtins.str]:
        """
        The region where the policy is located.
        """
        return pulumi.get(self, "region")

    @property
    @pulumi.getter(name="timePeriod")
    def time_period(self) -> pulumi.Output[Optional[builtins.int]]:
        """
        The duration (in days) for retained backups.
        """
        return pulumi.get(self, "time_period")

    @property
    @pulumi.getter(name="timeZone")
    def time_zone(self) -> pulumi.Output[builtins.str]:
        """
        The UTC time zone.
        """
        return pulumi.get(self, "time_zone")

    @property
    @pulumi.getter
    def type(self) -> pulumi.Output[builtins.str]:
        """
        The protection type of the CBR policy.
        """
        return pulumi.get(self, "type")

