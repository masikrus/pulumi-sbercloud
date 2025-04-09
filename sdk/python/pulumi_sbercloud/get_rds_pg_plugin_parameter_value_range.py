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

__all__ = [
    'GetRdsPgPluginParameterValueRangeResult',
    'AwaitableGetRdsPgPluginParameterValueRangeResult',
    'get_rds_pg_plugin_parameter_value_range',
    'get_rds_pg_plugin_parameter_value_range_output',
]

@pulumi.output_type
class GetRdsPgPluginParameterValueRangeResult:
    """
    A collection of values returned by getRdsPgPluginParameterValueRange.
    """
    def __init__(__self__, default_values=None, id=None, instance_id=None, name=None, region=None, restart_required=None, values=None):
        if default_values and not isinstance(default_values, list):
            raise TypeError("Expected argument 'default_values' to be a list")
        pulumi.set(__self__, "default_values", default_values)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if instance_id and not isinstance(instance_id, str):
            raise TypeError("Expected argument 'instance_id' to be a str")
        pulumi.set(__self__, "instance_id", instance_id)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if region and not isinstance(region, str):
            raise TypeError("Expected argument 'region' to be a str")
        pulumi.set(__self__, "region", region)
        if restart_required and not isinstance(restart_required, bool):
            raise TypeError("Expected argument 'restart_required' to be a bool")
        pulumi.set(__self__, "restart_required", restart_required)
        if values and not isinstance(values, list):
            raise TypeError("Expected argument 'values' to be a list")
        pulumi.set(__self__, "values", values)

    @property
    @pulumi.getter(name="defaultValues")
    def default_values(self) -> Sequence[builtins.str]:
        return pulumi.get(self, "default_values")

    @property
    @pulumi.getter
    def id(self) -> builtins.str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> builtins.str:
        return pulumi.get(self, "instance_id")

    @property
    @pulumi.getter
    def name(self) -> builtins.str:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def region(self) -> builtins.str:
        return pulumi.get(self, "region")

    @property
    @pulumi.getter(name="restartRequired")
    def restart_required(self) -> builtins.bool:
        return pulumi.get(self, "restart_required")

    @property
    @pulumi.getter
    def values(self) -> Sequence[builtins.str]:
        return pulumi.get(self, "values")


class AwaitableGetRdsPgPluginParameterValueRangeResult(GetRdsPgPluginParameterValueRangeResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetRdsPgPluginParameterValueRangeResult(
            default_values=self.default_values,
            id=self.id,
            instance_id=self.instance_id,
            name=self.name,
            region=self.region,
            restart_required=self.restart_required,
            values=self.values)


def get_rds_pg_plugin_parameter_value_range(instance_id: Optional[builtins.str] = None,
                                            name: Optional[builtins.str] = None,
                                            region: Optional[builtins.str] = None,
                                            opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetRdsPgPluginParameterValueRangeResult:
    """
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()
    __args__['instanceId'] = instance_id
    __args__['name'] = name
    __args__['region'] = region
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('sbercloud:index/getRdsPgPluginParameterValueRange:getRdsPgPluginParameterValueRange', __args__, opts=opts, typ=GetRdsPgPluginParameterValueRangeResult).value

    return AwaitableGetRdsPgPluginParameterValueRangeResult(
        default_values=pulumi.get(__ret__, 'default_values'),
        id=pulumi.get(__ret__, 'id'),
        instance_id=pulumi.get(__ret__, 'instance_id'),
        name=pulumi.get(__ret__, 'name'),
        region=pulumi.get(__ret__, 'region'),
        restart_required=pulumi.get(__ret__, 'restart_required'),
        values=pulumi.get(__ret__, 'values'))
def get_rds_pg_plugin_parameter_value_range_output(instance_id: Optional[pulumi.Input[builtins.str]] = None,
                                                   name: Optional[pulumi.Input[builtins.str]] = None,
                                                   region: Optional[pulumi.Input[Optional[builtins.str]]] = None,
                                                   opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetRdsPgPluginParameterValueRangeResult]:
    """
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()
    __args__['instanceId'] = instance_id
    __args__['name'] = name
    __args__['region'] = region
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('sbercloud:index/getRdsPgPluginParameterValueRange:getRdsPgPluginParameterValueRange', __args__, opts=opts, typ=GetRdsPgPluginParameterValueRangeResult)
    return __ret__.apply(lambda __response__: GetRdsPgPluginParameterValueRangeResult(
        default_values=pulumi.get(__response__, 'default_values'),
        id=pulumi.get(__response__, 'id'),
        instance_id=pulumi.get(__response__, 'instance_id'),
        name=pulumi.get(__response__, 'name'),
        region=pulumi.get(__response__, 'region'),
        restart_required=pulumi.get(__response__, 'restart_required'),
        values=pulumi.get(__response__, 'values')))
