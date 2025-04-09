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
    'GetVpnCustomerGatewaysResult',
    'AwaitableGetVpnCustomerGatewaysResult',
    'get_vpn_customer_gateways',
    'get_vpn_customer_gateways_output',
]

@pulumi.output_type
class GetVpnCustomerGatewaysResult:
    """
    A collection of values returned by getVpnCustomerGateways.
    """
    def __init__(__self__, asn=None, customer_gateway_id=None, customer_gateways=None, id=None, ip=None, name=None, region=None, route_mode=None):
        if asn and not isinstance(asn, int):
            raise TypeError("Expected argument 'asn' to be a int")
        pulumi.set(__self__, "asn", asn)
        if customer_gateway_id and not isinstance(customer_gateway_id, str):
            raise TypeError("Expected argument 'customer_gateway_id' to be a str")
        pulumi.set(__self__, "customer_gateway_id", customer_gateway_id)
        if customer_gateways and not isinstance(customer_gateways, list):
            raise TypeError("Expected argument 'customer_gateways' to be a list")
        pulumi.set(__self__, "customer_gateways", customer_gateways)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if ip and not isinstance(ip, str):
            raise TypeError("Expected argument 'ip' to be a str")
        pulumi.set(__self__, "ip", ip)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if region and not isinstance(region, str):
            raise TypeError("Expected argument 'region' to be a str")
        pulumi.set(__self__, "region", region)
        if route_mode and not isinstance(route_mode, str):
            raise TypeError("Expected argument 'route_mode' to be a str")
        pulumi.set(__self__, "route_mode", route_mode)

    @property
    @pulumi.getter
    def asn(self) -> Optional[builtins.int]:
        return pulumi.get(self, "asn")

    @property
    @pulumi.getter(name="customerGatewayId")
    def customer_gateway_id(self) -> Optional[builtins.str]:
        return pulumi.get(self, "customer_gateway_id")

    @property
    @pulumi.getter(name="customerGateways")
    def customer_gateways(self) -> Sequence['outputs.GetVpnCustomerGatewaysCustomerGatewayResult']:
        return pulumi.get(self, "customer_gateways")

    @property
    @pulumi.getter
    def id(self) -> builtins.str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def ip(self) -> Optional[builtins.str]:
        return pulumi.get(self, "ip")

    @property
    @pulumi.getter
    def name(self) -> Optional[builtins.str]:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def region(self) -> builtins.str:
        return pulumi.get(self, "region")

    @property
    @pulumi.getter(name="routeMode")
    def route_mode(self) -> Optional[builtins.str]:
        return pulumi.get(self, "route_mode")


class AwaitableGetVpnCustomerGatewaysResult(GetVpnCustomerGatewaysResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetVpnCustomerGatewaysResult(
            asn=self.asn,
            customer_gateway_id=self.customer_gateway_id,
            customer_gateways=self.customer_gateways,
            id=self.id,
            ip=self.ip,
            name=self.name,
            region=self.region,
            route_mode=self.route_mode)


def get_vpn_customer_gateways(asn: Optional[builtins.int] = None,
                              customer_gateway_id: Optional[builtins.str] = None,
                              ip: Optional[builtins.str] = None,
                              name: Optional[builtins.str] = None,
                              region: Optional[builtins.str] = None,
                              route_mode: Optional[builtins.str] = None,
                              opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetVpnCustomerGatewaysResult:
    """
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()
    __args__['asn'] = asn
    __args__['customerGatewayId'] = customer_gateway_id
    __args__['ip'] = ip
    __args__['name'] = name
    __args__['region'] = region
    __args__['routeMode'] = route_mode
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('sbercloud:index/getVpnCustomerGateways:getVpnCustomerGateways', __args__, opts=opts, typ=GetVpnCustomerGatewaysResult).value

    return AwaitableGetVpnCustomerGatewaysResult(
        asn=pulumi.get(__ret__, 'asn'),
        customer_gateway_id=pulumi.get(__ret__, 'customer_gateway_id'),
        customer_gateways=pulumi.get(__ret__, 'customer_gateways'),
        id=pulumi.get(__ret__, 'id'),
        ip=pulumi.get(__ret__, 'ip'),
        name=pulumi.get(__ret__, 'name'),
        region=pulumi.get(__ret__, 'region'),
        route_mode=pulumi.get(__ret__, 'route_mode'))
def get_vpn_customer_gateways_output(asn: Optional[pulumi.Input[Optional[builtins.int]]] = None,
                                     customer_gateway_id: Optional[pulumi.Input[Optional[builtins.str]]] = None,
                                     ip: Optional[pulumi.Input[Optional[builtins.str]]] = None,
                                     name: Optional[pulumi.Input[Optional[builtins.str]]] = None,
                                     region: Optional[pulumi.Input[Optional[builtins.str]]] = None,
                                     route_mode: Optional[pulumi.Input[Optional[builtins.str]]] = None,
                                     opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetVpnCustomerGatewaysResult]:
    """
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()
    __args__['asn'] = asn
    __args__['customerGatewayId'] = customer_gateway_id
    __args__['ip'] = ip
    __args__['name'] = name
    __args__['region'] = region
    __args__['routeMode'] = route_mode
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('sbercloud:index/getVpnCustomerGateways:getVpnCustomerGateways', __args__, opts=opts, typ=GetVpnCustomerGatewaysResult)
    return __ret__.apply(lambda __response__: GetVpnCustomerGatewaysResult(
        asn=pulumi.get(__response__, 'asn'),
        customer_gateway_id=pulumi.get(__response__, 'customer_gateway_id'),
        customer_gateways=pulumi.get(__response__, 'customer_gateways'),
        id=pulumi.get(__response__, 'id'),
        ip=pulumi.get(__response__, 'ip'),
        name=pulumi.get(__response__, 'name'),
        region=pulumi.get(__response__, 'region'),
        route_mode=pulumi.get(__response__, 'route_mode')))
