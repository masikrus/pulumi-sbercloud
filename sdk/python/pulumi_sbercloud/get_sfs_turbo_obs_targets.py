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
    'GetSfsTurboObsTargetsResult',
    'AwaitableGetSfsTurboObsTargetsResult',
    'get_sfs_turbo_obs_targets',
    'get_sfs_turbo_obs_targets_output',
]

@pulumi.output_type
class GetSfsTurboObsTargetsResult:
    """
    A collection of values returned by getSfsTurboObsTargets.
    """
    def __init__(__self__, bucket=None, id=None, region=None, share_id=None, status=None, target_id=None, targets=None):
        if bucket and not isinstance(bucket, str):
            raise TypeError("Expected argument 'bucket' to be a str")
        pulumi.set(__self__, "bucket", bucket)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if region and not isinstance(region, str):
            raise TypeError("Expected argument 'region' to be a str")
        pulumi.set(__self__, "region", region)
        if share_id and not isinstance(share_id, str):
            raise TypeError("Expected argument 'share_id' to be a str")
        pulumi.set(__self__, "share_id", share_id)
        if status and not isinstance(status, str):
            raise TypeError("Expected argument 'status' to be a str")
        pulumi.set(__self__, "status", status)
        if target_id and not isinstance(target_id, str):
            raise TypeError("Expected argument 'target_id' to be a str")
        pulumi.set(__self__, "target_id", target_id)
        if targets and not isinstance(targets, list):
            raise TypeError("Expected argument 'targets' to be a list")
        pulumi.set(__self__, "targets", targets)

    @property
    @pulumi.getter
    def bucket(self) -> Optional[builtins.str]:
        return pulumi.get(self, "bucket")

    @property
    @pulumi.getter
    def id(self) -> builtins.str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def region(self) -> builtins.str:
        return pulumi.get(self, "region")

    @property
    @pulumi.getter(name="shareId")
    def share_id(self) -> builtins.str:
        return pulumi.get(self, "share_id")

    @property
    @pulumi.getter
    def status(self) -> Optional[builtins.str]:
        return pulumi.get(self, "status")

    @property
    @pulumi.getter(name="targetId")
    def target_id(self) -> Optional[builtins.str]:
        return pulumi.get(self, "target_id")

    @property
    @pulumi.getter
    def targets(self) -> Sequence['outputs.GetSfsTurboObsTargetsTargetResult']:
        return pulumi.get(self, "targets")


class AwaitableGetSfsTurboObsTargetsResult(GetSfsTurboObsTargetsResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetSfsTurboObsTargetsResult(
            bucket=self.bucket,
            id=self.id,
            region=self.region,
            share_id=self.share_id,
            status=self.status,
            target_id=self.target_id,
            targets=self.targets)


def get_sfs_turbo_obs_targets(bucket: Optional[builtins.str] = None,
                              region: Optional[builtins.str] = None,
                              share_id: Optional[builtins.str] = None,
                              status: Optional[builtins.str] = None,
                              target_id: Optional[builtins.str] = None,
                              opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetSfsTurboObsTargetsResult:
    """
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()
    __args__['bucket'] = bucket
    __args__['region'] = region
    __args__['shareId'] = share_id
    __args__['status'] = status
    __args__['targetId'] = target_id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('sbercloud:index/getSfsTurboObsTargets:getSfsTurboObsTargets', __args__, opts=opts, typ=GetSfsTurboObsTargetsResult).value

    return AwaitableGetSfsTurboObsTargetsResult(
        bucket=pulumi.get(__ret__, 'bucket'),
        id=pulumi.get(__ret__, 'id'),
        region=pulumi.get(__ret__, 'region'),
        share_id=pulumi.get(__ret__, 'share_id'),
        status=pulumi.get(__ret__, 'status'),
        target_id=pulumi.get(__ret__, 'target_id'),
        targets=pulumi.get(__ret__, 'targets'))
def get_sfs_turbo_obs_targets_output(bucket: Optional[pulumi.Input[Optional[builtins.str]]] = None,
                                     region: Optional[pulumi.Input[Optional[builtins.str]]] = None,
                                     share_id: Optional[pulumi.Input[builtins.str]] = None,
                                     status: Optional[pulumi.Input[Optional[builtins.str]]] = None,
                                     target_id: Optional[pulumi.Input[Optional[builtins.str]]] = None,
                                     opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetSfsTurboObsTargetsResult]:
    """
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()
    __args__['bucket'] = bucket
    __args__['region'] = region
    __args__['shareId'] = share_id
    __args__['status'] = status
    __args__['targetId'] = target_id
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('sbercloud:index/getSfsTurboObsTargets:getSfsTurboObsTargets', __args__, opts=opts, typ=GetSfsTurboObsTargetsResult)
    return __ret__.apply(lambda __response__: GetSfsTurboObsTargetsResult(
        bucket=pulumi.get(__response__, 'bucket'),
        id=pulumi.get(__response__, 'id'),
        region=pulumi.get(__response__, 'region'),
        share_id=pulumi.get(__response__, 'share_id'),
        status=pulumi.get(__response__, 'status'),
        target_id=pulumi.get(__response__, 'target_id'),
        targets=pulumi.get(__response__, 'targets')))
