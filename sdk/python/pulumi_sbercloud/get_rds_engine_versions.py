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
    'GetRdsEngineVersionsResult',
    'AwaitableGetRdsEngineVersionsResult',
    'get_rds_engine_versions',
    'get_rds_engine_versions_output',
]

@pulumi.output_type
class GetRdsEngineVersionsResult:
    """
    A collection of values returned by getRdsEngineVersions.
    """
    def __init__(__self__, id=None, region=None, type=None, versions=None):
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if region and not isinstance(region, str):
            raise TypeError("Expected argument 'region' to be a str")
        pulumi.set(__self__, "region", region)
        if type and not isinstance(type, str):
            raise TypeError("Expected argument 'type' to be a str")
        pulumi.set(__self__, "type", type)
        if versions and not isinstance(versions, list):
            raise TypeError("Expected argument 'versions' to be a list")
        pulumi.set(__self__, "versions", versions)

    @property
    @pulumi.getter
    def id(self) -> builtins.str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def region(self) -> Optional[builtins.str]:
        return pulumi.get(self, "region")

    @property
    @pulumi.getter
    def type(self) -> Optional[builtins.str]:
        return pulumi.get(self, "type")

    @property
    @pulumi.getter
    def versions(self) -> Sequence['outputs.GetRdsEngineVersionsVersionResult']:
        return pulumi.get(self, "versions")


class AwaitableGetRdsEngineVersionsResult(GetRdsEngineVersionsResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetRdsEngineVersionsResult(
            id=self.id,
            region=self.region,
            type=self.type,
            versions=self.versions)


def get_rds_engine_versions(region: Optional[builtins.str] = None,
                            type: Optional[builtins.str] = None,
                            opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetRdsEngineVersionsResult:
    """
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()
    __args__['region'] = region
    __args__['type'] = type
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('sbercloud:index/getRdsEngineVersions:getRdsEngineVersions', __args__, opts=opts, typ=GetRdsEngineVersionsResult).value

    return AwaitableGetRdsEngineVersionsResult(
        id=pulumi.get(__ret__, 'id'),
        region=pulumi.get(__ret__, 'region'),
        type=pulumi.get(__ret__, 'type'),
        versions=pulumi.get(__ret__, 'versions'))
def get_rds_engine_versions_output(region: Optional[pulumi.Input[Optional[builtins.str]]] = None,
                                   type: Optional[pulumi.Input[Optional[builtins.str]]] = None,
                                   opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetRdsEngineVersionsResult]:
    """
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()
    __args__['region'] = region
    __args__['type'] = type
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('sbercloud:index/getRdsEngineVersions:getRdsEngineVersions', __args__, opts=opts, typ=GetRdsEngineVersionsResult)
    return __ret__.apply(lambda __response__: GetRdsEngineVersionsResult(
        id=pulumi.get(__response__, 'id'),
        region=pulumi.get(__response__, 'region'),
        type=pulumi.get(__response__, 'type'),
        versions=pulumi.get(__response__, 'versions')))
