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
    'GetVpnGatewaysResult',
    'AwaitableGetVpnGatewaysResult',
    'get_vpn_gateways',
    'get_vpn_gateways_output',
]

@pulumi.output_type
class GetVpnGatewaysResult:
    """
    A collection of values returned by getVpnGateways.
    """
    def __init__(__self__, attachment_type=None, enterprise_project_id=None, gateway_id=None, gateways=None, id=None, name=None, network_type=None, region=None):
        if attachment_type and not isinstance(attachment_type, str):
            raise TypeError("Expected argument 'attachment_type' to be a str")
        pulumi.set(__self__, "attachment_type", attachment_type)
        if enterprise_project_id and not isinstance(enterprise_project_id, str):
            raise TypeError("Expected argument 'enterprise_project_id' to be a str")
        pulumi.set(__self__, "enterprise_project_id", enterprise_project_id)
        if gateway_id and not isinstance(gateway_id, str):
            raise TypeError("Expected argument 'gateway_id' to be a str")
        pulumi.set(__self__, "gateway_id", gateway_id)
        if gateways and not isinstance(gateways, list):
            raise TypeError("Expected argument 'gateways' to be a list")
        pulumi.set(__self__, "gateways", gateways)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if network_type and not isinstance(network_type, str):
            raise TypeError("Expected argument 'network_type' to be a str")
        pulumi.set(__self__, "network_type", network_type)
        if region and not isinstance(region, str):
            raise TypeError("Expected argument 'region' to be a str")
        pulumi.set(__self__, "region", region)

    @property
    @pulumi.getter(name="attachmentType")
    def attachment_type(self) -> Optional[builtins.str]:
        return pulumi.get(self, "attachment_type")

    @property
    @pulumi.getter(name="enterpriseProjectId")
    def enterprise_project_id(self) -> Optional[builtins.str]:
        return pulumi.get(self, "enterprise_project_id")

    @property
    @pulumi.getter(name="gatewayId")
    def gateway_id(self) -> Optional[builtins.str]:
        return pulumi.get(self, "gateway_id")

    @property
    @pulumi.getter
    def gateways(self) -> Sequence['outputs.GetVpnGatewaysGatewayResult']:
        return pulumi.get(self, "gateways")

    @property
    @pulumi.getter
    def id(self) -> builtins.str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def name(self) -> Optional[builtins.str]:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="networkType")
    def network_type(self) -> Optional[builtins.str]:
        return pulumi.get(self, "network_type")

    @property
    @pulumi.getter
    def region(self) -> builtins.str:
        return pulumi.get(self, "region")


class AwaitableGetVpnGatewaysResult(GetVpnGatewaysResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetVpnGatewaysResult(
            attachment_type=self.attachment_type,
            enterprise_project_id=self.enterprise_project_id,
            gateway_id=self.gateway_id,
            gateways=self.gateways,
            id=self.id,
            name=self.name,
            network_type=self.network_type,
            region=self.region)


def get_vpn_gateways(attachment_type: Optional[builtins.str] = None,
                     enterprise_project_id: Optional[builtins.str] = None,
                     gateway_id: Optional[builtins.str] = None,
                     name: Optional[builtins.str] = None,
                     network_type: Optional[builtins.str] = None,
                     region: Optional[builtins.str] = None,
                     opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetVpnGatewaysResult:
    """
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()
    __args__['attachmentType'] = attachment_type
    __args__['enterpriseProjectId'] = enterprise_project_id
    __args__['gatewayId'] = gateway_id
    __args__['name'] = name
    __args__['networkType'] = network_type
    __args__['region'] = region
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('sbercloud:index/getVpnGateways:getVpnGateways', __args__, opts=opts, typ=GetVpnGatewaysResult).value

    return AwaitableGetVpnGatewaysResult(
        attachment_type=pulumi.get(__ret__, 'attachment_type'),
        enterprise_project_id=pulumi.get(__ret__, 'enterprise_project_id'),
        gateway_id=pulumi.get(__ret__, 'gateway_id'),
        gateways=pulumi.get(__ret__, 'gateways'),
        id=pulumi.get(__ret__, 'id'),
        name=pulumi.get(__ret__, 'name'),
        network_type=pulumi.get(__ret__, 'network_type'),
        region=pulumi.get(__ret__, 'region'))
def get_vpn_gateways_output(attachment_type: Optional[pulumi.Input[Optional[builtins.str]]] = None,
                            enterprise_project_id: Optional[pulumi.Input[Optional[builtins.str]]] = None,
                            gateway_id: Optional[pulumi.Input[Optional[builtins.str]]] = None,
                            name: Optional[pulumi.Input[Optional[builtins.str]]] = None,
                            network_type: Optional[pulumi.Input[Optional[builtins.str]]] = None,
                            region: Optional[pulumi.Input[Optional[builtins.str]]] = None,
                            opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetVpnGatewaysResult]:
    """
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()
    __args__['attachmentType'] = attachment_type
    __args__['enterpriseProjectId'] = enterprise_project_id
    __args__['gatewayId'] = gateway_id
    __args__['name'] = name
    __args__['networkType'] = network_type
    __args__['region'] = region
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('sbercloud:index/getVpnGateways:getVpnGateways', __args__, opts=opts, typ=GetVpnGatewaysResult)
    return __ret__.apply(lambda __response__: GetVpnGatewaysResult(
        attachment_type=pulumi.get(__response__, 'attachment_type'),
        enterprise_project_id=pulumi.get(__response__, 'enterprise_project_id'),
        gateway_id=pulumi.get(__response__, 'gateway_id'),
        gateways=pulumi.get(__response__, 'gateways'),
        id=pulumi.get(__response__, 'id'),
        name=pulumi.get(__response__, 'name'),
        network_type=pulumi.get(__response__, 'network_type'),
        region=pulumi.get(__response__, 'region')))
