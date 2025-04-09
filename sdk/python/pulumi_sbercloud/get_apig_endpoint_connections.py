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
    'GetApigEndpointConnectionsResult',
    'AwaitableGetApigEndpointConnectionsResult',
    'get_apig_endpoint_connections',
    'get_apig_endpoint_connections_output',
]

@pulumi.output_type
class GetApigEndpointConnectionsResult:
    """
    A collection of values returned by getApigEndpointConnections.
    """
    def __init__(__self__, connections=None, endpoint_id=None, id=None, instance_id=None, packet_id=None, region=None, status=None):
        if connections and not isinstance(connections, list):
            raise TypeError("Expected argument 'connections' to be a list")
        pulumi.set(__self__, "connections", connections)
        if endpoint_id and not isinstance(endpoint_id, str):
            raise TypeError("Expected argument 'endpoint_id' to be a str")
        pulumi.set(__self__, "endpoint_id", endpoint_id)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if instance_id and not isinstance(instance_id, str):
            raise TypeError("Expected argument 'instance_id' to be a str")
        pulumi.set(__self__, "instance_id", instance_id)
        if packet_id and not isinstance(packet_id, int):
            raise TypeError("Expected argument 'packet_id' to be a int")
        pulumi.set(__self__, "packet_id", packet_id)
        if region and not isinstance(region, str):
            raise TypeError("Expected argument 'region' to be a str")
        pulumi.set(__self__, "region", region)
        if status and not isinstance(status, str):
            raise TypeError("Expected argument 'status' to be a str")
        pulumi.set(__self__, "status", status)

    @property
    @pulumi.getter
    def connections(self) -> Sequence['outputs.GetApigEndpointConnectionsConnectionResult']:
        return pulumi.get(self, "connections")

    @property
    @pulumi.getter(name="endpointId")
    def endpoint_id(self) -> Optional[builtins.str]:
        return pulumi.get(self, "endpoint_id")

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
    @pulumi.getter(name="packetId")
    def packet_id(self) -> Optional[builtins.int]:
        return pulumi.get(self, "packet_id")

    @property
    @pulumi.getter
    def region(self) -> builtins.str:
        return pulumi.get(self, "region")

    @property
    @pulumi.getter
    def status(self) -> Optional[builtins.str]:
        return pulumi.get(self, "status")


class AwaitableGetApigEndpointConnectionsResult(GetApigEndpointConnectionsResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetApigEndpointConnectionsResult(
            connections=self.connections,
            endpoint_id=self.endpoint_id,
            id=self.id,
            instance_id=self.instance_id,
            packet_id=self.packet_id,
            region=self.region,
            status=self.status)


def get_apig_endpoint_connections(endpoint_id: Optional[builtins.str] = None,
                                  instance_id: Optional[builtins.str] = None,
                                  packet_id: Optional[builtins.int] = None,
                                  region: Optional[builtins.str] = None,
                                  status: Optional[builtins.str] = None,
                                  opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetApigEndpointConnectionsResult:
    """
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()
    __args__['endpointId'] = endpoint_id
    __args__['instanceId'] = instance_id
    __args__['packetId'] = packet_id
    __args__['region'] = region
    __args__['status'] = status
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('sbercloud:index/getApigEndpointConnections:getApigEndpointConnections', __args__, opts=opts, typ=GetApigEndpointConnectionsResult).value

    return AwaitableGetApigEndpointConnectionsResult(
        connections=pulumi.get(__ret__, 'connections'),
        endpoint_id=pulumi.get(__ret__, 'endpoint_id'),
        id=pulumi.get(__ret__, 'id'),
        instance_id=pulumi.get(__ret__, 'instance_id'),
        packet_id=pulumi.get(__ret__, 'packet_id'),
        region=pulumi.get(__ret__, 'region'),
        status=pulumi.get(__ret__, 'status'))
def get_apig_endpoint_connections_output(endpoint_id: Optional[pulumi.Input[Optional[builtins.str]]] = None,
                                         instance_id: Optional[pulumi.Input[builtins.str]] = None,
                                         packet_id: Optional[pulumi.Input[Optional[builtins.int]]] = None,
                                         region: Optional[pulumi.Input[Optional[builtins.str]]] = None,
                                         status: Optional[pulumi.Input[Optional[builtins.str]]] = None,
                                         opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetApigEndpointConnectionsResult]:
    """
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()
    __args__['endpointId'] = endpoint_id
    __args__['instanceId'] = instance_id
    __args__['packetId'] = packet_id
    __args__['region'] = region
    __args__['status'] = status
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('sbercloud:index/getApigEndpointConnections:getApigEndpointConnections', __args__, opts=opts, typ=GetApigEndpointConnectionsResult)
    return __ret__.apply(lambda __response__: GetApigEndpointConnectionsResult(
        connections=pulumi.get(__response__, 'connections'),
        endpoint_id=pulumi.get(__response__, 'endpoint_id'),
        id=pulumi.get(__response__, 'id'),
        instance_id=pulumi.get(__response__, 'instance_id'),
        packet_id=pulumi.get(__response__, 'packet_id'),
        region=pulumi.get(__response__, 'region'),
        status=pulumi.get(__response__, 'status')))
