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
    'GetSfsTurboDataTasksResult',
    'AwaitableGetSfsTurboDataTasksResult',
    'get_sfs_turbo_data_tasks',
    'get_sfs_turbo_data_tasks_output',
]

@pulumi.output_type
class GetSfsTurboDataTasksResult:
    """
    A collection of values returned by getSfsTurboDataTasks.
    """
    def __init__(__self__, id=None, region=None, share_id=None, status=None, tasks=None, type=None):
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
        if tasks and not isinstance(tasks, list):
            raise TypeError("Expected argument 'tasks' to be a list")
        pulumi.set(__self__, "tasks", tasks)
        if type and not isinstance(type, str):
            raise TypeError("Expected argument 'type' to be a str")
        pulumi.set(__self__, "type", type)

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
    @pulumi.getter
    def tasks(self) -> Sequence['outputs.GetSfsTurboDataTasksTaskResult']:
        return pulumi.get(self, "tasks")

    @property
    @pulumi.getter
    def type(self) -> Optional[builtins.str]:
        return pulumi.get(self, "type")


class AwaitableGetSfsTurboDataTasksResult(GetSfsTurboDataTasksResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetSfsTurboDataTasksResult(
            id=self.id,
            region=self.region,
            share_id=self.share_id,
            status=self.status,
            tasks=self.tasks,
            type=self.type)


def get_sfs_turbo_data_tasks(region: Optional[builtins.str] = None,
                             share_id: Optional[builtins.str] = None,
                             status: Optional[builtins.str] = None,
                             type: Optional[builtins.str] = None,
                             opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetSfsTurboDataTasksResult:
    """
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()
    __args__['region'] = region
    __args__['shareId'] = share_id
    __args__['status'] = status
    __args__['type'] = type
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('sbercloud:index/getSfsTurboDataTasks:getSfsTurboDataTasks', __args__, opts=opts, typ=GetSfsTurboDataTasksResult).value

    return AwaitableGetSfsTurboDataTasksResult(
        id=pulumi.get(__ret__, 'id'),
        region=pulumi.get(__ret__, 'region'),
        share_id=pulumi.get(__ret__, 'share_id'),
        status=pulumi.get(__ret__, 'status'),
        tasks=pulumi.get(__ret__, 'tasks'),
        type=pulumi.get(__ret__, 'type'))
def get_sfs_turbo_data_tasks_output(region: Optional[pulumi.Input[Optional[builtins.str]]] = None,
                                    share_id: Optional[pulumi.Input[builtins.str]] = None,
                                    status: Optional[pulumi.Input[Optional[builtins.str]]] = None,
                                    type: Optional[pulumi.Input[Optional[builtins.str]]] = None,
                                    opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetSfsTurboDataTasksResult]:
    """
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()
    __args__['region'] = region
    __args__['shareId'] = share_id
    __args__['status'] = status
    __args__['type'] = type
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('sbercloud:index/getSfsTurboDataTasks:getSfsTurboDataTasks', __args__, opts=opts, typ=GetSfsTurboDataTasksResult)
    return __ret__.apply(lambda __response__: GetSfsTurboDataTasksResult(
        id=pulumi.get(__response__, 'id'),
        region=pulumi.get(__response__, 'region'),
        share_id=pulumi.get(__response__, 'share_id'),
        status=pulumi.get(__response__, 'status'),
        tasks=pulumi.get(__response__, 'tasks'),
        type=pulumi.get(__response__, 'type')))
