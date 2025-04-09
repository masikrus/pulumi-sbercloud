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
    'GetApigAppcodesResult',
    'AwaitableGetApigAppcodesResult',
    'get_apig_appcodes',
    'get_apig_appcodes_output',
]

@pulumi.output_type
class GetApigAppcodesResult:
    """
    A collection of values returned by getApigAppcodes.
    """
    def __init__(__self__, appcodes=None, application_id=None, id=None, instance_id=None, region=None):
        if appcodes and not isinstance(appcodes, list):
            raise TypeError("Expected argument 'appcodes' to be a list")
        pulumi.set(__self__, "appcodes", appcodes)
        if application_id and not isinstance(application_id, str):
            raise TypeError("Expected argument 'application_id' to be a str")
        pulumi.set(__self__, "application_id", application_id)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if instance_id and not isinstance(instance_id, str):
            raise TypeError("Expected argument 'instance_id' to be a str")
        pulumi.set(__self__, "instance_id", instance_id)
        if region and not isinstance(region, str):
            raise TypeError("Expected argument 'region' to be a str")
        pulumi.set(__self__, "region", region)

    @property
    @pulumi.getter
    def appcodes(self) -> Sequence['outputs.GetApigAppcodesAppcodeResult']:
        return pulumi.get(self, "appcodes")

    @property
    @pulumi.getter(name="applicationId")
    def application_id(self) -> builtins.str:
        return pulumi.get(self, "application_id")

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
    def region(self) -> builtins.str:
        return pulumi.get(self, "region")


class AwaitableGetApigAppcodesResult(GetApigAppcodesResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetApigAppcodesResult(
            appcodes=self.appcodes,
            application_id=self.application_id,
            id=self.id,
            instance_id=self.instance_id,
            region=self.region)


def get_apig_appcodes(application_id: Optional[builtins.str] = None,
                      instance_id: Optional[builtins.str] = None,
                      region: Optional[builtins.str] = None,
                      opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetApigAppcodesResult:
    """
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()
    __args__['applicationId'] = application_id
    __args__['instanceId'] = instance_id
    __args__['region'] = region
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('sbercloud:index/getApigAppcodes:getApigAppcodes', __args__, opts=opts, typ=GetApigAppcodesResult).value

    return AwaitableGetApigAppcodesResult(
        appcodes=pulumi.get(__ret__, 'appcodes'),
        application_id=pulumi.get(__ret__, 'application_id'),
        id=pulumi.get(__ret__, 'id'),
        instance_id=pulumi.get(__ret__, 'instance_id'),
        region=pulumi.get(__ret__, 'region'))
def get_apig_appcodes_output(application_id: Optional[pulumi.Input[builtins.str]] = None,
                             instance_id: Optional[pulumi.Input[builtins.str]] = None,
                             region: Optional[pulumi.Input[Optional[builtins.str]]] = None,
                             opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetApigAppcodesResult]:
    """
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()
    __args__['applicationId'] = application_id
    __args__['instanceId'] = instance_id
    __args__['region'] = region
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('sbercloud:index/getApigAppcodes:getApigAppcodes', __args__, opts=opts, typ=GetApigAppcodesResult)
    return __ret__.apply(lambda __response__: GetApigAppcodesResult(
        appcodes=pulumi.get(__response__, 'appcodes'),
        application_id=pulumi.get(__response__, 'application_id'),
        id=pulumi.get(__response__, 'id'),
        instance_id=pulumi.get(__response__, 'instance_id'),
        region=pulumi.get(__response__, 'region')))
