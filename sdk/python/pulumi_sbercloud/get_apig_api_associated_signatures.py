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
    'GetApigApiAssociatedSignaturesResult',
    'AwaitableGetApigApiAssociatedSignaturesResult',
    'get_apig_api_associated_signatures',
    'get_apig_api_associated_signatures_output',
]

@pulumi.output_type
class GetApigApiAssociatedSignaturesResult:
    """
    A collection of values returned by getApigApiAssociatedSignatures.
    """
    def __init__(__self__, api_id=None, env_id=None, env_name=None, id=None, instance_id=None, name=None, region=None, signature_id=None, signatures=None, type=None):
        if api_id and not isinstance(api_id, str):
            raise TypeError("Expected argument 'api_id' to be a str")
        pulumi.set(__self__, "api_id", api_id)
        if env_id and not isinstance(env_id, str):
            raise TypeError("Expected argument 'env_id' to be a str")
        pulumi.set(__self__, "env_id", env_id)
        if env_name and not isinstance(env_name, str):
            raise TypeError("Expected argument 'env_name' to be a str")
        pulumi.set(__self__, "env_name", env_name)
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
        if signature_id and not isinstance(signature_id, str):
            raise TypeError("Expected argument 'signature_id' to be a str")
        pulumi.set(__self__, "signature_id", signature_id)
        if signatures and not isinstance(signatures, list):
            raise TypeError("Expected argument 'signatures' to be a list")
        pulumi.set(__self__, "signatures", signatures)
        if type and not isinstance(type, str):
            raise TypeError("Expected argument 'type' to be a str")
        pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter(name="apiId")
    def api_id(self) -> builtins.str:
        return pulumi.get(self, "api_id")

    @property
    @pulumi.getter(name="envId")
    def env_id(self) -> Optional[builtins.str]:
        return pulumi.get(self, "env_id")

    @property
    @pulumi.getter(name="envName")
    def env_name(self) -> Optional[builtins.str]:
        return pulumi.get(self, "env_name")

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
    def name(self) -> Optional[builtins.str]:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def region(self) -> builtins.str:
        return pulumi.get(self, "region")

    @property
    @pulumi.getter(name="signatureId")
    def signature_id(self) -> Optional[builtins.str]:
        return pulumi.get(self, "signature_id")

    @property
    @pulumi.getter
    def signatures(self) -> Sequence['outputs.GetApigApiAssociatedSignaturesSignatureResult']:
        return pulumi.get(self, "signatures")

    @property
    @pulumi.getter
    def type(self) -> Optional[builtins.str]:
        return pulumi.get(self, "type")


class AwaitableGetApigApiAssociatedSignaturesResult(GetApigApiAssociatedSignaturesResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetApigApiAssociatedSignaturesResult(
            api_id=self.api_id,
            env_id=self.env_id,
            env_name=self.env_name,
            id=self.id,
            instance_id=self.instance_id,
            name=self.name,
            region=self.region,
            signature_id=self.signature_id,
            signatures=self.signatures,
            type=self.type)


def get_apig_api_associated_signatures(api_id: Optional[builtins.str] = None,
                                       env_id: Optional[builtins.str] = None,
                                       env_name: Optional[builtins.str] = None,
                                       instance_id: Optional[builtins.str] = None,
                                       name: Optional[builtins.str] = None,
                                       region: Optional[builtins.str] = None,
                                       signature_id: Optional[builtins.str] = None,
                                       type: Optional[builtins.str] = None,
                                       opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetApigApiAssociatedSignaturesResult:
    """
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()
    __args__['apiId'] = api_id
    __args__['envId'] = env_id
    __args__['envName'] = env_name
    __args__['instanceId'] = instance_id
    __args__['name'] = name
    __args__['region'] = region
    __args__['signatureId'] = signature_id
    __args__['type'] = type
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('sbercloud:index/getApigApiAssociatedSignatures:getApigApiAssociatedSignatures', __args__, opts=opts, typ=GetApigApiAssociatedSignaturesResult).value

    return AwaitableGetApigApiAssociatedSignaturesResult(
        api_id=pulumi.get(__ret__, 'api_id'),
        env_id=pulumi.get(__ret__, 'env_id'),
        env_name=pulumi.get(__ret__, 'env_name'),
        id=pulumi.get(__ret__, 'id'),
        instance_id=pulumi.get(__ret__, 'instance_id'),
        name=pulumi.get(__ret__, 'name'),
        region=pulumi.get(__ret__, 'region'),
        signature_id=pulumi.get(__ret__, 'signature_id'),
        signatures=pulumi.get(__ret__, 'signatures'),
        type=pulumi.get(__ret__, 'type'))
def get_apig_api_associated_signatures_output(api_id: Optional[pulumi.Input[builtins.str]] = None,
                                              env_id: Optional[pulumi.Input[Optional[builtins.str]]] = None,
                                              env_name: Optional[pulumi.Input[Optional[builtins.str]]] = None,
                                              instance_id: Optional[pulumi.Input[builtins.str]] = None,
                                              name: Optional[pulumi.Input[Optional[builtins.str]]] = None,
                                              region: Optional[pulumi.Input[Optional[builtins.str]]] = None,
                                              signature_id: Optional[pulumi.Input[Optional[builtins.str]]] = None,
                                              type: Optional[pulumi.Input[Optional[builtins.str]]] = None,
                                              opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetApigApiAssociatedSignaturesResult]:
    """
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()
    __args__['apiId'] = api_id
    __args__['envId'] = env_id
    __args__['envName'] = env_name
    __args__['instanceId'] = instance_id
    __args__['name'] = name
    __args__['region'] = region
    __args__['signatureId'] = signature_id
    __args__['type'] = type
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('sbercloud:index/getApigApiAssociatedSignatures:getApigApiAssociatedSignatures', __args__, opts=opts, typ=GetApigApiAssociatedSignaturesResult)
    return __ret__.apply(lambda __response__: GetApigApiAssociatedSignaturesResult(
        api_id=pulumi.get(__response__, 'api_id'),
        env_id=pulumi.get(__response__, 'env_id'),
        env_name=pulumi.get(__response__, 'env_name'),
        id=pulumi.get(__response__, 'id'),
        instance_id=pulumi.get(__response__, 'instance_id'),
        name=pulumi.get(__response__, 'name'),
        region=pulumi.get(__response__, 'region'),
        signature_id=pulumi.get(__response__, 'signature_id'),
        signatures=pulumi.get(__response__, 'signatures'),
        type=pulumi.get(__response__, 'type')))
