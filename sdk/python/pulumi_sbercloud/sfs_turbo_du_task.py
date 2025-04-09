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
from ._inputs import *

__all__ = ['SfsTurboDuTaskArgs', 'SfsTurboDuTask']

@pulumi.input_type
class SfsTurboDuTaskArgs:
    def __init__(__self__, *,
                 path: pulumi.Input[builtins.str],
                 share_id: pulumi.Input[builtins.str],
                 region: Optional[pulumi.Input[builtins.str]] = None):
        """
        The set of arguments for constructing a SfsTurboDuTask resource.
        """
        pulumi.set(__self__, "path", path)
        pulumi.set(__self__, "share_id", share_id)
        if region is not None:
            pulumi.set(__self__, "region", region)

    @property
    @pulumi.getter
    def path(self) -> pulumi.Input[builtins.str]:
        return pulumi.get(self, "path")

    @path.setter
    def path(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "path", value)

    @property
    @pulumi.getter(name="shareId")
    def share_id(self) -> pulumi.Input[builtins.str]:
        return pulumi.get(self, "share_id")

    @share_id.setter
    def share_id(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "share_id", value)

    @property
    @pulumi.getter
    def region(self) -> Optional[pulumi.Input[builtins.str]]:
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "region", value)


@pulumi.input_type
class _SfsTurboDuTaskState:
    def __init__(__self__, *,
                 begin_time: Optional[pulumi.Input[builtins.str]] = None,
                 dir_usages: Optional[pulumi.Input[Sequence[pulumi.Input['SfsTurboDuTaskDirUsageArgs']]]] = None,
                 end_time: Optional[pulumi.Input[builtins.str]] = None,
                 path: Optional[pulumi.Input[builtins.str]] = None,
                 region: Optional[pulumi.Input[builtins.str]] = None,
                 share_id: Optional[pulumi.Input[builtins.str]] = None,
                 status: Optional[pulumi.Input[builtins.str]] = None):
        """
        Input properties used for looking up and filtering SfsTurboDuTask resources.
        """
        if begin_time is not None:
            pulumi.set(__self__, "begin_time", begin_time)
        if dir_usages is not None:
            pulumi.set(__self__, "dir_usages", dir_usages)
        if end_time is not None:
            pulumi.set(__self__, "end_time", end_time)
        if path is not None:
            pulumi.set(__self__, "path", path)
        if region is not None:
            pulumi.set(__self__, "region", region)
        if share_id is not None:
            pulumi.set(__self__, "share_id", share_id)
        if status is not None:
            pulumi.set(__self__, "status", status)

    @property
    @pulumi.getter(name="beginTime")
    def begin_time(self) -> Optional[pulumi.Input[builtins.str]]:
        return pulumi.get(self, "begin_time")

    @begin_time.setter
    def begin_time(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "begin_time", value)

    @property
    @pulumi.getter(name="dirUsages")
    def dir_usages(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['SfsTurboDuTaskDirUsageArgs']]]]:
        return pulumi.get(self, "dir_usages")

    @dir_usages.setter
    def dir_usages(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['SfsTurboDuTaskDirUsageArgs']]]]):
        pulumi.set(self, "dir_usages", value)

    @property
    @pulumi.getter(name="endTime")
    def end_time(self) -> Optional[pulumi.Input[builtins.str]]:
        return pulumi.get(self, "end_time")

    @end_time.setter
    def end_time(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "end_time", value)

    @property
    @pulumi.getter
    def path(self) -> Optional[pulumi.Input[builtins.str]]:
        return pulumi.get(self, "path")

    @path.setter
    def path(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "path", value)

    @property
    @pulumi.getter
    def region(self) -> Optional[pulumi.Input[builtins.str]]:
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "region", value)

    @property
    @pulumi.getter(name="shareId")
    def share_id(self) -> Optional[pulumi.Input[builtins.str]]:
        return pulumi.get(self, "share_id")

    @share_id.setter
    def share_id(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "share_id", value)

    @property
    @pulumi.getter
    def status(self) -> Optional[pulumi.Input[builtins.str]]:
        return pulumi.get(self, "status")

    @status.setter
    def status(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "status", value)


class SfsTurboDuTask(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 path: Optional[pulumi.Input[builtins.str]] = None,
                 region: Optional[pulumi.Input[builtins.str]] = None,
                 share_id: Optional[pulumi.Input[builtins.str]] = None,
                 __props__=None):
        """
        Create a SfsTurboDuTask resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: SfsTurboDuTaskArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a SfsTurboDuTask resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param SfsTurboDuTaskArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(SfsTurboDuTaskArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 path: Optional[pulumi.Input[builtins.str]] = None,
                 region: Optional[pulumi.Input[builtins.str]] = None,
                 share_id: Optional[pulumi.Input[builtins.str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = SfsTurboDuTaskArgs.__new__(SfsTurboDuTaskArgs)

            if path is None and not opts.urn:
                raise TypeError("Missing required property 'path'")
            __props__.__dict__["path"] = path
            __props__.__dict__["region"] = region
            if share_id is None and not opts.urn:
                raise TypeError("Missing required property 'share_id'")
            __props__.__dict__["share_id"] = share_id
            __props__.__dict__["begin_time"] = None
            __props__.__dict__["dir_usages"] = None
            __props__.__dict__["end_time"] = None
            __props__.__dict__["status"] = None
        super(SfsTurboDuTask, __self__).__init__(
            'sbercloud:index/sfsTurboDuTask:SfsTurboDuTask',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            begin_time: Optional[pulumi.Input[builtins.str]] = None,
            dir_usages: Optional[pulumi.Input[Sequence[pulumi.Input[Union['SfsTurboDuTaskDirUsageArgs', 'SfsTurboDuTaskDirUsageArgsDict']]]]] = None,
            end_time: Optional[pulumi.Input[builtins.str]] = None,
            path: Optional[pulumi.Input[builtins.str]] = None,
            region: Optional[pulumi.Input[builtins.str]] = None,
            share_id: Optional[pulumi.Input[builtins.str]] = None,
            status: Optional[pulumi.Input[builtins.str]] = None) -> 'SfsTurboDuTask':
        """
        Get an existing SfsTurboDuTask resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _SfsTurboDuTaskState.__new__(_SfsTurboDuTaskState)

        __props__.__dict__["begin_time"] = begin_time
        __props__.__dict__["dir_usages"] = dir_usages
        __props__.__dict__["end_time"] = end_time
        __props__.__dict__["path"] = path
        __props__.__dict__["region"] = region
        __props__.__dict__["share_id"] = share_id
        __props__.__dict__["status"] = status
        return SfsTurboDuTask(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="beginTime")
    def begin_time(self) -> pulumi.Output[builtins.str]:
        return pulumi.get(self, "begin_time")

    @property
    @pulumi.getter(name="dirUsages")
    def dir_usages(self) -> pulumi.Output[Sequence['outputs.SfsTurboDuTaskDirUsage']]:
        return pulumi.get(self, "dir_usages")

    @property
    @pulumi.getter(name="endTime")
    def end_time(self) -> pulumi.Output[builtins.str]:
        return pulumi.get(self, "end_time")

    @property
    @pulumi.getter
    def path(self) -> pulumi.Output[builtins.str]:
        return pulumi.get(self, "path")

    @property
    @pulumi.getter
    def region(self) -> pulumi.Output[builtins.str]:
        return pulumi.get(self, "region")

    @property
    @pulumi.getter(name="shareId")
    def share_id(self) -> pulumi.Output[builtins.str]:
        return pulumi.get(self, "share_id")

    @property
    @pulumi.getter
    def status(self) -> pulumi.Output[builtins.str]:
        return pulumi.get(self, "status")

