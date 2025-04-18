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
    'GetLbPoolsResult',
    'AwaitableGetLbPoolsResult',
    'get_lb_pools',
    'get_lb_pools_output',
]

@pulumi.output_type
class GetLbPoolsResult:
    """
    A collection of values returned by getLbPools.
    """
    def __init__(__self__, description=None, healthmonitor_id=None, id=None, lb_method=None, loadbalancer_id=None, name=None, pool_id=None, pools=None, protocol=None, region=None):
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if healthmonitor_id and not isinstance(healthmonitor_id, str):
            raise TypeError("Expected argument 'healthmonitor_id' to be a str")
        pulumi.set(__self__, "healthmonitor_id", healthmonitor_id)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if lb_method and not isinstance(lb_method, str):
            raise TypeError("Expected argument 'lb_method' to be a str")
        pulumi.set(__self__, "lb_method", lb_method)
        if loadbalancer_id and not isinstance(loadbalancer_id, str):
            raise TypeError("Expected argument 'loadbalancer_id' to be a str")
        pulumi.set(__self__, "loadbalancer_id", loadbalancer_id)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if pool_id and not isinstance(pool_id, str):
            raise TypeError("Expected argument 'pool_id' to be a str")
        pulumi.set(__self__, "pool_id", pool_id)
        if pools and not isinstance(pools, list):
            raise TypeError("Expected argument 'pools' to be a list")
        pulumi.set(__self__, "pools", pools)
        if protocol and not isinstance(protocol, str):
            raise TypeError("Expected argument 'protocol' to be a str")
        pulumi.set(__self__, "protocol", protocol)
        if region and not isinstance(region, str):
            raise TypeError("Expected argument 'region' to be a str")
        pulumi.set(__self__, "region", region)

    @property
    @pulumi.getter
    def description(self) -> Optional[builtins.str]:
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="healthmonitorId")
    def healthmonitor_id(self) -> Optional[builtins.str]:
        return pulumi.get(self, "healthmonitor_id")

    @property
    @pulumi.getter
    def id(self) -> builtins.str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="lbMethod")
    def lb_method(self) -> Optional[builtins.str]:
        return pulumi.get(self, "lb_method")

    @property
    @pulumi.getter(name="loadbalancerId")
    def loadbalancer_id(self) -> Optional[builtins.str]:
        return pulumi.get(self, "loadbalancer_id")

    @property
    @pulumi.getter
    def name(self) -> Optional[builtins.str]:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="poolId")
    def pool_id(self) -> Optional[builtins.str]:
        return pulumi.get(self, "pool_id")

    @property
    @pulumi.getter
    def pools(self) -> Sequence['outputs.GetLbPoolsPoolResult']:
        return pulumi.get(self, "pools")

    @property
    @pulumi.getter
    def protocol(self) -> Optional[builtins.str]:
        return pulumi.get(self, "protocol")

    @property
    @pulumi.getter
    def region(self) -> builtins.str:
        return pulumi.get(self, "region")


class AwaitableGetLbPoolsResult(GetLbPoolsResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetLbPoolsResult(
            description=self.description,
            healthmonitor_id=self.healthmonitor_id,
            id=self.id,
            lb_method=self.lb_method,
            loadbalancer_id=self.loadbalancer_id,
            name=self.name,
            pool_id=self.pool_id,
            pools=self.pools,
            protocol=self.protocol,
            region=self.region)


def get_lb_pools(description: Optional[builtins.str] = None,
                 healthmonitor_id: Optional[builtins.str] = None,
                 lb_method: Optional[builtins.str] = None,
                 loadbalancer_id: Optional[builtins.str] = None,
                 name: Optional[builtins.str] = None,
                 pool_id: Optional[builtins.str] = None,
                 protocol: Optional[builtins.str] = None,
                 region: Optional[builtins.str] = None,
                 opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetLbPoolsResult:
    """
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()
    __args__['description'] = description
    __args__['healthmonitorId'] = healthmonitor_id
    __args__['lbMethod'] = lb_method
    __args__['loadbalancerId'] = loadbalancer_id
    __args__['name'] = name
    __args__['poolId'] = pool_id
    __args__['protocol'] = protocol
    __args__['region'] = region
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('sbercloud:index/getLbPools:getLbPools', __args__, opts=opts, typ=GetLbPoolsResult).value

    return AwaitableGetLbPoolsResult(
        description=pulumi.get(__ret__, 'description'),
        healthmonitor_id=pulumi.get(__ret__, 'healthmonitor_id'),
        id=pulumi.get(__ret__, 'id'),
        lb_method=pulumi.get(__ret__, 'lb_method'),
        loadbalancer_id=pulumi.get(__ret__, 'loadbalancer_id'),
        name=pulumi.get(__ret__, 'name'),
        pool_id=pulumi.get(__ret__, 'pool_id'),
        pools=pulumi.get(__ret__, 'pools'),
        protocol=pulumi.get(__ret__, 'protocol'),
        region=pulumi.get(__ret__, 'region'))
def get_lb_pools_output(description: Optional[pulumi.Input[Optional[builtins.str]]] = None,
                        healthmonitor_id: Optional[pulumi.Input[Optional[builtins.str]]] = None,
                        lb_method: Optional[pulumi.Input[Optional[builtins.str]]] = None,
                        loadbalancer_id: Optional[pulumi.Input[Optional[builtins.str]]] = None,
                        name: Optional[pulumi.Input[Optional[builtins.str]]] = None,
                        pool_id: Optional[pulumi.Input[Optional[builtins.str]]] = None,
                        protocol: Optional[pulumi.Input[Optional[builtins.str]]] = None,
                        region: Optional[pulumi.Input[Optional[builtins.str]]] = None,
                        opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetLbPoolsResult]:
    """
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()
    __args__['description'] = description
    __args__['healthmonitorId'] = healthmonitor_id
    __args__['lbMethod'] = lb_method
    __args__['loadbalancerId'] = loadbalancer_id
    __args__['name'] = name
    __args__['poolId'] = pool_id
    __args__['protocol'] = protocol
    __args__['region'] = region
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('sbercloud:index/getLbPools:getLbPools', __args__, opts=opts, typ=GetLbPoolsResult)
    return __ret__.apply(lambda __response__: GetLbPoolsResult(
        description=pulumi.get(__response__, 'description'),
        healthmonitor_id=pulumi.get(__response__, 'healthmonitor_id'),
        id=pulumi.get(__response__, 'id'),
        lb_method=pulumi.get(__response__, 'lb_method'),
        loadbalancer_id=pulumi.get(__response__, 'loadbalancer_id'),
        name=pulumi.get(__response__, 'name'),
        pool_id=pulumi.get(__response__, 'pool_id'),
        pools=pulumi.get(__response__, 'pools'),
        protocol=pulumi.get(__response__, 'protocol'),
        region=pulumi.get(__response__, 'region')))
