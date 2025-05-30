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
    'GetDmsMaintainwindowResult',
    'AwaitableGetDmsMaintainwindowResult',
    'get_dms_maintainwindow',
    'get_dms_maintainwindow_output',
]

@pulumi.output_type
class GetDmsMaintainwindowResult:
    """
    A collection of values returned by getDmsMaintainwindow.
    """
    def __init__(__self__, begin=None, default=None, end=None, id=None, region=None, seq=None):
        if begin and not isinstance(begin, str):
            raise TypeError("Expected argument 'begin' to be a str")
        pulumi.set(__self__, "begin", begin)
        if default and not isinstance(default, bool):
            raise TypeError("Expected argument 'default' to be a bool")
        pulumi.set(__self__, "default", default)
        if end and not isinstance(end, str):
            raise TypeError("Expected argument 'end' to be a str")
        pulumi.set(__self__, "end", end)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if region and not isinstance(region, str):
            raise TypeError("Expected argument 'region' to be a str")
        pulumi.set(__self__, "region", region)
        if seq and not isinstance(seq, int):
            raise TypeError("Expected argument 'seq' to be a int")
        pulumi.set(__self__, "seq", seq)

    @property
    @pulumi.getter
    def begin(self) -> builtins.str:
        return pulumi.get(self, "begin")

    @property
    @pulumi.getter
    def default(self) -> builtins.bool:
        return pulumi.get(self, "default")

    @property
    @pulumi.getter
    def end(self) -> builtins.str:
        return pulumi.get(self, "end")

    @property
    @pulumi.getter
    def id(self) -> builtins.str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def region(self) -> builtins.str:
        return pulumi.get(self, "region")

    @property
    @pulumi.getter
    def seq(self) -> builtins.int:
        return pulumi.get(self, "seq")


class AwaitableGetDmsMaintainwindowResult(GetDmsMaintainwindowResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetDmsMaintainwindowResult(
            begin=self.begin,
            default=self.default,
            end=self.end,
            id=self.id,
            region=self.region,
            seq=self.seq)


def get_dms_maintainwindow(begin: Optional[builtins.str] = None,
                           default: Optional[builtins.bool] = None,
                           end: Optional[builtins.str] = None,
                           region: Optional[builtins.str] = None,
                           seq: Optional[builtins.int] = None,
                           opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetDmsMaintainwindowResult:
    """
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()
    __args__['begin'] = begin
    __args__['default'] = default
    __args__['end'] = end
    __args__['region'] = region
    __args__['seq'] = seq
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('sbercloud:index/getDmsMaintainwindow:getDmsMaintainwindow', __args__, opts=opts, typ=GetDmsMaintainwindowResult).value

    return AwaitableGetDmsMaintainwindowResult(
        begin=pulumi.get(__ret__, 'begin'),
        default=pulumi.get(__ret__, 'default'),
        end=pulumi.get(__ret__, 'end'),
        id=pulumi.get(__ret__, 'id'),
        region=pulumi.get(__ret__, 'region'),
        seq=pulumi.get(__ret__, 'seq'))
def get_dms_maintainwindow_output(begin: Optional[pulumi.Input[Optional[builtins.str]]] = None,
                                  default: Optional[pulumi.Input[Optional[builtins.bool]]] = None,
                                  end: Optional[pulumi.Input[Optional[builtins.str]]] = None,
                                  region: Optional[pulumi.Input[Optional[builtins.str]]] = None,
                                  seq: Optional[pulumi.Input[Optional[builtins.int]]] = None,
                                  opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetDmsMaintainwindowResult]:
    """
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()
    __args__['begin'] = begin
    __args__['default'] = default
    __args__['end'] = end
    __args__['region'] = region
    __args__['seq'] = seq
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('sbercloud:index/getDmsMaintainwindow:getDmsMaintainwindow', __args__, opts=opts, typ=GetDmsMaintainwindowResult)
    return __ret__.apply(lambda __response__: GetDmsMaintainwindowResult(
        begin=pulumi.get(__response__, 'begin'),
        default=pulumi.get(__response__, 'default'),
        end=pulumi.get(__response__, 'end'),
        id=pulumi.get(__response__, 'id'),
        region=pulumi.get(__response__, 'region'),
        seq=pulumi.get(__response__, 'seq')))
