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
    'GetVpcSubnetIdsResult',
    'AwaitableGetVpcSubnetIdsResult',
    'get_vpc_subnet_ids',
    'get_vpc_subnet_ids_output',
]

@pulumi.output_type
class GetVpcSubnetIdsResult:
    """
    A collection of values returned by getVpcSubnetIds.
    """
    def __init__(__self__, id=None, ids=None, region=None, vpc_id=None):
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if ids and not isinstance(ids, list):
            raise TypeError("Expected argument 'ids' to be a list")
        pulumi.set(__self__, "ids", ids)
        if region and not isinstance(region, str):
            raise TypeError("Expected argument 'region' to be a str")
        pulumi.set(__self__, "region", region)
        if vpc_id and not isinstance(vpc_id, str):
            raise TypeError("Expected argument 'vpc_id' to be a str")
        pulumi.set(__self__, "vpc_id", vpc_id)

    @property
    @pulumi.getter
    def id(self) -> builtins.str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def ids(self) -> Sequence[builtins.str]:
        return pulumi.get(self, "ids")

    @property
    @pulumi.getter
    def region(self) -> builtins.str:
        return pulumi.get(self, "region")

    @property
    @pulumi.getter(name="vpcId")
    def vpc_id(self) -> builtins.str:
        return pulumi.get(self, "vpc_id")


class AwaitableGetVpcSubnetIdsResult(GetVpcSubnetIdsResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetVpcSubnetIdsResult(
            id=self.id,
            ids=self.ids,
            region=self.region,
            vpc_id=self.vpc_id)


def get_vpc_subnet_ids(region: Optional[builtins.str] = None,
                       vpc_id: Optional[builtins.str] = None,
                       opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetVpcSubnetIdsResult:
    """
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()
    __args__['region'] = region
    __args__['vpcId'] = vpc_id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('sbercloud:index/getVpcSubnetIds:getVpcSubnetIds', __args__, opts=opts, typ=GetVpcSubnetIdsResult).value

    return AwaitableGetVpcSubnetIdsResult(
        id=pulumi.get(__ret__, 'id'),
        ids=pulumi.get(__ret__, 'ids'),
        region=pulumi.get(__ret__, 'region'),
        vpc_id=pulumi.get(__ret__, 'vpc_id'))
def get_vpc_subnet_ids_output(region: Optional[pulumi.Input[Optional[builtins.str]]] = None,
                              vpc_id: Optional[pulumi.Input[builtins.str]] = None,
                              opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetVpcSubnetIdsResult]:
    """
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()
    __args__['region'] = region
    __args__['vpcId'] = vpc_id
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('sbercloud:index/getVpcSubnetIds:getVpcSubnetIds', __args__, opts=opts, typ=GetVpcSubnetIdsResult)
    return __ret__.apply(lambda __response__: GetVpcSubnetIdsResult(
        id=pulumi.get(__response__, 'id'),
        ids=pulumi.get(__response__, 'ids'),
        region=pulumi.get(__response__, 'region'),
        vpc_id=pulumi.get(__response__, 'vpc_id')))
