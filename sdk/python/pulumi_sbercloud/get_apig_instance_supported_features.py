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
    'GetApigInstanceSupportedFeaturesResult',
    'AwaitableGetApigInstanceSupportedFeaturesResult',
    'get_apig_instance_supported_features',
    'get_apig_instance_supported_features_output',
]

@pulumi.output_type
class GetApigInstanceSupportedFeaturesResult:
    """
    A collection of values returned by getApigInstanceSupportedFeatures.
    """
    def __init__(__self__, features=None, id=None, instance_id=None, region=None):
        if features and not isinstance(features, list):
            raise TypeError("Expected argument 'features' to be a list")
        pulumi.set(__self__, "features", features)
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
    def features(self) -> Sequence[builtins.str]:
        return pulumi.get(self, "features")

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


class AwaitableGetApigInstanceSupportedFeaturesResult(GetApigInstanceSupportedFeaturesResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetApigInstanceSupportedFeaturesResult(
            features=self.features,
            id=self.id,
            instance_id=self.instance_id,
            region=self.region)


def get_apig_instance_supported_features(instance_id: Optional[builtins.str] = None,
                                         region: Optional[builtins.str] = None,
                                         opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetApigInstanceSupportedFeaturesResult:
    """
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()
    __args__['instanceId'] = instance_id
    __args__['region'] = region
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('sbercloud:index/getApigInstanceSupportedFeatures:getApigInstanceSupportedFeatures', __args__, opts=opts, typ=GetApigInstanceSupportedFeaturesResult).value

    return AwaitableGetApigInstanceSupportedFeaturesResult(
        features=pulumi.get(__ret__, 'features'),
        id=pulumi.get(__ret__, 'id'),
        instance_id=pulumi.get(__ret__, 'instance_id'),
        region=pulumi.get(__ret__, 'region'))
def get_apig_instance_supported_features_output(instance_id: Optional[pulumi.Input[builtins.str]] = None,
                                                region: Optional[pulumi.Input[Optional[builtins.str]]] = None,
                                                opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetApigInstanceSupportedFeaturesResult]:
    """
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()
    __args__['instanceId'] = instance_id
    __args__['region'] = region
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('sbercloud:index/getApigInstanceSupportedFeatures:getApigInstanceSupportedFeatures', __args__, opts=opts, typ=GetApigInstanceSupportedFeaturesResult)
    return __ret__.apply(lambda __response__: GetApigInstanceSupportedFeaturesResult(
        features=pulumi.get(__response__, 'features'),
        id=pulumi.get(__response__, 'id'),
        instance_id=pulumi.get(__response__, 'instance_id'),
        region=pulumi.get(__response__, 'region')))
