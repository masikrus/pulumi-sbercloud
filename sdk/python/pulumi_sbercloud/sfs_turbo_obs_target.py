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

__all__ = ['SfsTurboObsTargetArgs', 'SfsTurboObsTarget']

@pulumi.input_type
class SfsTurboObsTargetArgs:
    def __init__(__self__, *,
                 file_system_path: pulumi.Input[builtins.str],
                 obs: pulumi.Input['SfsTurboObsTargetObsArgs'],
                 share_id: pulumi.Input[builtins.str],
                 delete_data_in_file_system: Optional[pulumi.Input[builtins.bool]] = None,
                 region: Optional[pulumi.Input[builtins.str]] = None):
        """
        The set of arguments for constructing a SfsTurboObsTarget resource.
        """
        pulumi.set(__self__, "file_system_path", file_system_path)
        pulumi.set(__self__, "obs", obs)
        pulumi.set(__self__, "share_id", share_id)
        if delete_data_in_file_system is not None:
            pulumi.set(__self__, "delete_data_in_file_system", delete_data_in_file_system)
        if region is not None:
            pulumi.set(__self__, "region", region)

    @property
    @pulumi.getter(name="fileSystemPath")
    def file_system_path(self) -> pulumi.Input[builtins.str]:
        return pulumi.get(self, "file_system_path")

    @file_system_path.setter
    def file_system_path(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "file_system_path", value)

    @property
    @pulumi.getter
    def obs(self) -> pulumi.Input['SfsTurboObsTargetObsArgs']:
        return pulumi.get(self, "obs")

    @obs.setter
    def obs(self, value: pulumi.Input['SfsTurboObsTargetObsArgs']):
        pulumi.set(self, "obs", value)

    @property
    @pulumi.getter(name="shareId")
    def share_id(self) -> pulumi.Input[builtins.str]:
        return pulumi.get(self, "share_id")

    @share_id.setter
    def share_id(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "share_id", value)

    @property
    @pulumi.getter(name="deleteDataInFileSystem")
    def delete_data_in_file_system(self) -> Optional[pulumi.Input[builtins.bool]]:
        return pulumi.get(self, "delete_data_in_file_system")

    @delete_data_in_file_system.setter
    def delete_data_in_file_system(self, value: Optional[pulumi.Input[builtins.bool]]):
        pulumi.set(self, "delete_data_in_file_system", value)

    @property
    @pulumi.getter
    def region(self) -> Optional[pulumi.Input[builtins.str]]:
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "region", value)


@pulumi.input_type
class _SfsTurboObsTargetState:
    def __init__(__self__, *,
                 created_at: Optional[pulumi.Input[builtins.str]] = None,
                 delete_data_in_file_system: Optional[pulumi.Input[builtins.bool]] = None,
                 file_system_path: Optional[pulumi.Input[builtins.str]] = None,
                 obs: Optional[pulumi.Input['SfsTurboObsTargetObsArgs']] = None,
                 region: Optional[pulumi.Input[builtins.str]] = None,
                 share_id: Optional[pulumi.Input[builtins.str]] = None,
                 status: Optional[pulumi.Input[builtins.str]] = None):
        """
        Input properties used for looking up and filtering SfsTurboObsTarget resources.
        """
        if created_at is not None:
            pulumi.set(__self__, "created_at", created_at)
        if delete_data_in_file_system is not None:
            pulumi.set(__self__, "delete_data_in_file_system", delete_data_in_file_system)
        if file_system_path is not None:
            pulumi.set(__self__, "file_system_path", file_system_path)
        if obs is not None:
            pulumi.set(__self__, "obs", obs)
        if region is not None:
            pulumi.set(__self__, "region", region)
        if share_id is not None:
            pulumi.set(__self__, "share_id", share_id)
        if status is not None:
            pulumi.set(__self__, "status", status)

    @property
    @pulumi.getter(name="createdAt")
    def created_at(self) -> Optional[pulumi.Input[builtins.str]]:
        return pulumi.get(self, "created_at")

    @created_at.setter
    def created_at(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "created_at", value)

    @property
    @pulumi.getter(name="deleteDataInFileSystem")
    def delete_data_in_file_system(self) -> Optional[pulumi.Input[builtins.bool]]:
        return pulumi.get(self, "delete_data_in_file_system")

    @delete_data_in_file_system.setter
    def delete_data_in_file_system(self, value: Optional[pulumi.Input[builtins.bool]]):
        pulumi.set(self, "delete_data_in_file_system", value)

    @property
    @pulumi.getter(name="fileSystemPath")
    def file_system_path(self) -> Optional[pulumi.Input[builtins.str]]:
        return pulumi.get(self, "file_system_path")

    @file_system_path.setter
    def file_system_path(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "file_system_path", value)

    @property
    @pulumi.getter
    def obs(self) -> Optional[pulumi.Input['SfsTurboObsTargetObsArgs']]:
        return pulumi.get(self, "obs")

    @obs.setter
    def obs(self, value: Optional[pulumi.Input['SfsTurboObsTargetObsArgs']]):
        pulumi.set(self, "obs", value)

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


class SfsTurboObsTarget(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 delete_data_in_file_system: Optional[pulumi.Input[builtins.bool]] = None,
                 file_system_path: Optional[pulumi.Input[builtins.str]] = None,
                 obs: Optional[pulumi.Input[Union['SfsTurboObsTargetObsArgs', 'SfsTurboObsTargetObsArgsDict']]] = None,
                 region: Optional[pulumi.Input[builtins.str]] = None,
                 share_id: Optional[pulumi.Input[builtins.str]] = None,
                 __props__=None):
        """
        Create a SfsTurboObsTarget resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: SfsTurboObsTargetArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a SfsTurboObsTarget resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param SfsTurboObsTargetArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(SfsTurboObsTargetArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 delete_data_in_file_system: Optional[pulumi.Input[builtins.bool]] = None,
                 file_system_path: Optional[pulumi.Input[builtins.str]] = None,
                 obs: Optional[pulumi.Input[Union['SfsTurboObsTargetObsArgs', 'SfsTurboObsTargetObsArgsDict']]] = None,
                 region: Optional[pulumi.Input[builtins.str]] = None,
                 share_id: Optional[pulumi.Input[builtins.str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = SfsTurboObsTargetArgs.__new__(SfsTurboObsTargetArgs)

            __props__.__dict__["delete_data_in_file_system"] = delete_data_in_file_system
            if file_system_path is None and not opts.urn:
                raise TypeError("Missing required property 'file_system_path'")
            __props__.__dict__["file_system_path"] = file_system_path
            if obs is None and not opts.urn:
                raise TypeError("Missing required property 'obs'")
            __props__.__dict__["obs"] = obs
            __props__.__dict__["region"] = region
            if share_id is None and not opts.urn:
                raise TypeError("Missing required property 'share_id'")
            __props__.__dict__["share_id"] = share_id
            __props__.__dict__["created_at"] = None
            __props__.__dict__["status"] = None
        super(SfsTurboObsTarget, __self__).__init__(
            'sbercloud:index/sfsTurboObsTarget:SfsTurboObsTarget',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            created_at: Optional[pulumi.Input[builtins.str]] = None,
            delete_data_in_file_system: Optional[pulumi.Input[builtins.bool]] = None,
            file_system_path: Optional[pulumi.Input[builtins.str]] = None,
            obs: Optional[pulumi.Input[Union['SfsTurboObsTargetObsArgs', 'SfsTurboObsTargetObsArgsDict']]] = None,
            region: Optional[pulumi.Input[builtins.str]] = None,
            share_id: Optional[pulumi.Input[builtins.str]] = None,
            status: Optional[pulumi.Input[builtins.str]] = None) -> 'SfsTurboObsTarget':
        """
        Get an existing SfsTurboObsTarget resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _SfsTurboObsTargetState.__new__(_SfsTurboObsTargetState)

        __props__.__dict__["created_at"] = created_at
        __props__.__dict__["delete_data_in_file_system"] = delete_data_in_file_system
        __props__.__dict__["file_system_path"] = file_system_path
        __props__.__dict__["obs"] = obs
        __props__.__dict__["region"] = region
        __props__.__dict__["share_id"] = share_id
        __props__.__dict__["status"] = status
        return SfsTurboObsTarget(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="createdAt")
    def created_at(self) -> pulumi.Output[builtins.str]:
        return pulumi.get(self, "created_at")

    @property
    @pulumi.getter(name="deleteDataInFileSystem")
    def delete_data_in_file_system(self) -> pulumi.Output[Optional[builtins.bool]]:
        return pulumi.get(self, "delete_data_in_file_system")

    @property
    @pulumi.getter(name="fileSystemPath")
    def file_system_path(self) -> pulumi.Output[builtins.str]:
        return pulumi.get(self, "file_system_path")

    @property
    @pulumi.getter
    def obs(self) -> pulumi.Output['outputs.SfsTurboObsTargetObs']:
        return pulumi.get(self, "obs")

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

