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
    'GetVpcsResult',
    'AwaitableGetVpcsResult',
    'get_vpcs',
    'get_vpcs_output',
]

@pulumi.output_type
class GetVpcsResult:
    """
    A collection of values returned by getVpcs.
    """
    def __init__(__self__, cidr=None, enterprise_project_id=None, id=None, name=None, region=None, status=None, tags=None, vpcs=None):
        if cidr and not isinstance(cidr, str):
            raise TypeError("Expected argument 'cidr' to be a str")
        pulumi.set(__self__, "cidr", cidr)
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
        if status and not isinstance(status, str):
            raise TypeError("Expected argument 'status' to be a str")
        pulumi.set(__self__, "status", status)
        if tags and not isinstance(tags, dict):
            raise TypeError("Expected argument 'tags' to be a dict")
        pulumi.set(__self__, "tags", tags)
        if vpcs and not isinstance(vpcs, list):
            raise TypeError("Expected argument 'vpcs' to be a list")
        pulumi.set(__self__, "vpcs", vpcs)

    @property
    @pulumi.getter
    def cidr(self) -> Optional[builtins.str]:
        return pulumi.get(self, "cidr")

    @property
    @pulumi.getter(name="enterpriseProjectId")
    def enterprise_project_id(self) -> Optional[builtins.str]:
        return pulumi.get(self, "enterprise_project_id")

    @property
    @pulumi.getter
    def id(self) -> builtins.str:
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def name(self) -> Optional[builtins.str]:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def region(self) -> builtins.str:
        return pulumi.get(self, "region")

    @property
    @pulumi.getter
    def status(self) -> Optional[builtins.str]:
        return pulumi.get(self, "status")

    @property
    @pulumi.getter
    def tags(self) -> Optional[Mapping[str, builtins.str]]:
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter
    def vpcs(self) -> Sequence['outputs.GetVpcsVpcResult']:
        return pulumi.get(self, "vpcs")


class AwaitableGetVpcsResult(GetVpcsResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetVpcsResult(
            cidr=self.cidr,
            enterprise_project_id=self.enterprise_project_id,
            id=self.id,
            name=self.name,
            region=self.region,
            status=self.status,
            tags=self.tags,
            vpcs=self.vpcs)


def get_vpcs(cidr: Optional[builtins.str] = None,
             enterprise_project_id: Optional[builtins.str] = None,
             id: Optional[builtins.str] = None,
             name: Optional[builtins.str] = None,
             region: Optional[builtins.str] = None,
             status: Optional[builtins.str] = None,
             tags: Optional[Mapping[str, builtins.str]] = None,
             opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetVpcsResult:
    """
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()
    __args__['cidr'] = cidr
    __args__['enterpriseProjectId'] = enterprise_project_id
    __args__['id'] = id
    __args__['name'] = name
    __args__['region'] = region
    __args__['status'] = status
    __args__['tags'] = tags
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('sbercloud:index/getVpcs:getVpcs', __args__, opts=opts, typ=GetVpcsResult).value

    return AwaitableGetVpcsResult(
        cidr=pulumi.get(__ret__, 'cidr'),
        enterprise_project_id=pulumi.get(__ret__, 'enterprise_project_id'),
        id=pulumi.get(__ret__, 'id'),
        name=pulumi.get(__ret__, 'name'),
        region=pulumi.get(__ret__, 'region'),
        status=pulumi.get(__ret__, 'status'),
        tags=pulumi.get(__ret__, 'tags'),
        vpcs=pulumi.get(__ret__, 'vpcs'))
def get_vpcs_output(cidr: Optional[pulumi.Input[Optional[builtins.str]]] = None,
                    enterprise_project_id: Optional[pulumi.Input[Optional[builtins.str]]] = None,
                    id: Optional[pulumi.Input[Optional[builtins.str]]] = None,
                    name: Optional[pulumi.Input[Optional[builtins.str]]] = None,
                    region: Optional[pulumi.Input[Optional[builtins.str]]] = None,
                    status: Optional[pulumi.Input[Optional[builtins.str]]] = None,
                    tags: Optional[pulumi.Input[Optional[Mapping[str, builtins.str]]]] = None,
                    opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetVpcsResult]:
    """
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()
    __args__['cidr'] = cidr
    __args__['enterpriseProjectId'] = enterprise_project_id
    __args__['id'] = id
    __args__['name'] = name
    __args__['region'] = region
    __args__['status'] = status
    __args__['tags'] = tags
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('sbercloud:index/getVpcs:getVpcs', __args__, opts=opts, typ=GetVpcsResult)
    return __ret__.apply(lambda __response__: GetVpcsResult(
        cidr=pulumi.get(__response__, 'cidr'),
        enterprise_project_id=pulumi.get(__response__, 'enterprise_project_id'),
        id=pulumi.get(__response__, 'id'),
        name=pulumi.get(__response__, 'name'),
        region=pulumi.get(__response__, 'region'),
        status=pulumi.get(__response__, 'status'),
        tags=pulumi.get(__response__, 'tags'),
        vpcs=pulumi.get(__response__, 'vpcs')))
