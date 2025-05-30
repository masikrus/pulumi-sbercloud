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
    'GetVpcPeeringConnectionResult',
    'AwaitableGetVpcPeeringConnectionResult',
    'get_vpc_peering_connection',
    'get_vpc_peering_connection_output',
]

@pulumi.output_type
class GetVpcPeeringConnectionResult:
    """
    A collection of values returned by getVpcPeeringConnection.
    """
    def __init__(__self__, description=None, id=None, name=None, peer_tenant_id=None, peer_vpc_id=None, region=None, status=None, vpc_id=None):
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if peer_tenant_id and not isinstance(peer_tenant_id, str):
            raise TypeError("Expected argument 'peer_tenant_id' to be a str")
        pulumi.set(__self__, "peer_tenant_id", peer_tenant_id)
        if peer_vpc_id and not isinstance(peer_vpc_id, str):
            raise TypeError("Expected argument 'peer_vpc_id' to be a str")
        pulumi.set(__self__, "peer_vpc_id", peer_vpc_id)
        if region and not isinstance(region, str):
            raise TypeError("Expected argument 'region' to be a str")
        pulumi.set(__self__, "region", region)
        if status and not isinstance(status, str):
            raise TypeError("Expected argument 'status' to be a str")
        pulumi.set(__self__, "status", status)
        if vpc_id and not isinstance(vpc_id, str):
            raise TypeError("Expected argument 'vpc_id' to be a str")
        pulumi.set(__self__, "vpc_id", vpc_id)

    @property
    @pulumi.getter
    def description(self) -> builtins.str:
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def id(self) -> builtins.str:
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def name(self) -> builtins.str:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="peerTenantId")
    def peer_tenant_id(self) -> builtins.str:
        return pulumi.get(self, "peer_tenant_id")

    @property
    @pulumi.getter(name="peerVpcId")
    def peer_vpc_id(self) -> builtins.str:
        return pulumi.get(self, "peer_vpc_id")

    @property
    @pulumi.getter
    def region(self) -> builtins.str:
        return pulumi.get(self, "region")

    @property
    @pulumi.getter
    def status(self) -> builtins.str:
        return pulumi.get(self, "status")

    @property
    @pulumi.getter(name="vpcId")
    def vpc_id(self) -> builtins.str:
        return pulumi.get(self, "vpc_id")


class AwaitableGetVpcPeeringConnectionResult(GetVpcPeeringConnectionResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetVpcPeeringConnectionResult(
            description=self.description,
            id=self.id,
            name=self.name,
            peer_tenant_id=self.peer_tenant_id,
            peer_vpc_id=self.peer_vpc_id,
            region=self.region,
            status=self.status,
            vpc_id=self.vpc_id)


def get_vpc_peering_connection(id: Optional[builtins.str] = None,
                               name: Optional[builtins.str] = None,
                               peer_tenant_id: Optional[builtins.str] = None,
                               peer_vpc_id: Optional[builtins.str] = None,
                               region: Optional[builtins.str] = None,
                               status: Optional[builtins.str] = None,
                               vpc_id: Optional[builtins.str] = None,
                               opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetVpcPeeringConnectionResult:
    """
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()
    __args__['id'] = id
    __args__['name'] = name
    __args__['peerTenantId'] = peer_tenant_id
    __args__['peerVpcId'] = peer_vpc_id
    __args__['region'] = region
    __args__['status'] = status
    __args__['vpcId'] = vpc_id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('sbercloud:index/getVpcPeeringConnection:getVpcPeeringConnection', __args__, opts=opts, typ=GetVpcPeeringConnectionResult).value

    return AwaitableGetVpcPeeringConnectionResult(
        description=pulumi.get(__ret__, 'description'),
        id=pulumi.get(__ret__, 'id'),
        name=pulumi.get(__ret__, 'name'),
        peer_tenant_id=pulumi.get(__ret__, 'peer_tenant_id'),
        peer_vpc_id=pulumi.get(__ret__, 'peer_vpc_id'),
        region=pulumi.get(__ret__, 'region'),
        status=pulumi.get(__ret__, 'status'),
        vpc_id=pulumi.get(__ret__, 'vpc_id'))
def get_vpc_peering_connection_output(id: Optional[pulumi.Input[Optional[builtins.str]]] = None,
                                      name: Optional[pulumi.Input[Optional[builtins.str]]] = None,
                                      peer_tenant_id: Optional[pulumi.Input[Optional[builtins.str]]] = None,
                                      peer_vpc_id: Optional[pulumi.Input[Optional[builtins.str]]] = None,
                                      region: Optional[pulumi.Input[Optional[builtins.str]]] = None,
                                      status: Optional[pulumi.Input[Optional[builtins.str]]] = None,
                                      vpc_id: Optional[pulumi.Input[Optional[builtins.str]]] = None,
                                      opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetVpcPeeringConnectionResult]:
    """
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()
    __args__['id'] = id
    __args__['name'] = name
    __args__['peerTenantId'] = peer_tenant_id
    __args__['peerVpcId'] = peer_vpc_id
    __args__['region'] = region
    __args__['status'] = status
    __args__['vpcId'] = vpc_id
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('sbercloud:index/getVpcPeeringConnection:getVpcPeeringConnection', __args__, opts=opts, typ=GetVpcPeeringConnectionResult)
    return __ret__.apply(lambda __response__: GetVpcPeeringConnectionResult(
        description=pulumi.get(__response__, 'description'),
        id=pulumi.get(__response__, 'id'),
        name=pulumi.get(__response__, 'name'),
        peer_tenant_id=pulumi.get(__response__, 'peer_tenant_id'),
        peer_vpc_id=pulumi.get(__response__, 'peer_vpc_id'),
        region=pulumi.get(__response__, 'region'),
        status=pulumi.get(__response__, 'status'),
        vpc_id=pulumi.get(__response__, 'vpc_id')))
