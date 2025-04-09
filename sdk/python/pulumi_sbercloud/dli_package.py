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

__all__ = ['DliPackageArgs', 'DliPackage']

@pulumi.input_type
class DliPackageArgs:
    def __init__(__self__, *,
                 object_path: pulumi.Input[builtins.str],
                 type: pulumi.Input[builtins.str],
                 group_name: Optional[pulumi.Input[builtins.str]] = None,
                 is_async: Optional[pulumi.Input[builtins.bool]] = None,
                 owner: Optional[pulumi.Input[builtins.str]] = None,
                 region: Optional[pulumi.Input[builtins.str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]]] = None):
        """
        The set of arguments for constructing a DliPackage resource.
        """
        pulumi.set(__self__, "object_path", object_path)
        pulumi.set(__self__, "type", type)
        if group_name is not None:
            pulumi.set(__self__, "group_name", group_name)
        if is_async is not None:
            pulumi.set(__self__, "is_async", is_async)
        if owner is not None:
            pulumi.set(__self__, "owner", owner)
        if region is not None:
            pulumi.set(__self__, "region", region)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter(name="objectPath")
    def object_path(self) -> pulumi.Input[builtins.str]:
        return pulumi.get(self, "object_path")

    @object_path.setter
    def object_path(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "object_path", value)

    @property
    @pulumi.getter
    def type(self) -> pulumi.Input[builtins.str]:
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "type", value)

    @property
    @pulumi.getter(name="groupName")
    def group_name(self) -> Optional[pulumi.Input[builtins.str]]:
        return pulumi.get(self, "group_name")

    @group_name.setter
    def group_name(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "group_name", value)

    @property
    @pulumi.getter(name="isAsync")
    def is_async(self) -> Optional[pulumi.Input[builtins.bool]]:
        return pulumi.get(self, "is_async")

    @is_async.setter
    def is_async(self, value: Optional[pulumi.Input[builtins.bool]]):
        pulumi.set(self, "is_async", value)

    @property
    @pulumi.getter
    def owner(self) -> Optional[pulumi.Input[builtins.str]]:
        return pulumi.get(self, "owner")

    @owner.setter
    def owner(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "owner", value)

    @property
    @pulumi.getter
    def region(self) -> Optional[pulumi.Input[builtins.str]]:
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "region", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]]]:
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]]]):
        pulumi.set(self, "tags", value)


@pulumi.input_type
class _DliPackageState:
    def __init__(__self__, *,
                 created_at: Optional[pulumi.Input[builtins.str]] = None,
                 group_name: Optional[pulumi.Input[builtins.str]] = None,
                 is_async: Optional[pulumi.Input[builtins.bool]] = None,
                 object_name: Optional[pulumi.Input[builtins.str]] = None,
                 object_path: Optional[pulumi.Input[builtins.str]] = None,
                 owner: Optional[pulumi.Input[builtins.str]] = None,
                 region: Optional[pulumi.Input[builtins.str]] = None,
                 status: Optional[pulumi.Input[builtins.str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]]] = None,
                 type: Optional[pulumi.Input[builtins.str]] = None,
                 updated_at: Optional[pulumi.Input[builtins.str]] = None):
        """
        Input properties used for looking up and filtering DliPackage resources.
        """
        if created_at is not None:
            pulumi.set(__self__, "created_at", created_at)
        if group_name is not None:
            pulumi.set(__self__, "group_name", group_name)
        if is_async is not None:
            pulumi.set(__self__, "is_async", is_async)
        if object_name is not None:
            pulumi.set(__self__, "object_name", object_name)
        if object_path is not None:
            pulumi.set(__self__, "object_path", object_path)
        if owner is not None:
            pulumi.set(__self__, "owner", owner)
        if region is not None:
            pulumi.set(__self__, "region", region)
        if status is not None:
            pulumi.set(__self__, "status", status)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)
        if type is not None:
            pulumi.set(__self__, "type", type)
        if updated_at is not None:
            pulumi.set(__self__, "updated_at", updated_at)

    @property
    @pulumi.getter(name="createdAt")
    def created_at(self) -> Optional[pulumi.Input[builtins.str]]:
        return pulumi.get(self, "created_at")

    @created_at.setter
    def created_at(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "created_at", value)

    @property
    @pulumi.getter(name="groupName")
    def group_name(self) -> Optional[pulumi.Input[builtins.str]]:
        return pulumi.get(self, "group_name")

    @group_name.setter
    def group_name(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "group_name", value)

    @property
    @pulumi.getter(name="isAsync")
    def is_async(self) -> Optional[pulumi.Input[builtins.bool]]:
        return pulumi.get(self, "is_async")

    @is_async.setter
    def is_async(self, value: Optional[pulumi.Input[builtins.bool]]):
        pulumi.set(self, "is_async", value)

    @property
    @pulumi.getter(name="objectName")
    def object_name(self) -> Optional[pulumi.Input[builtins.str]]:
        return pulumi.get(self, "object_name")

    @object_name.setter
    def object_name(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "object_name", value)

    @property
    @pulumi.getter(name="objectPath")
    def object_path(self) -> Optional[pulumi.Input[builtins.str]]:
        return pulumi.get(self, "object_path")

    @object_path.setter
    def object_path(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "object_path", value)

    @property
    @pulumi.getter
    def owner(self) -> Optional[pulumi.Input[builtins.str]]:
        return pulumi.get(self, "owner")

    @owner.setter
    def owner(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "owner", value)

    @property
    @pulumi.getter
    def region(self) -> Optional[pulumi.Input[builtins.str]]:
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "region", value)

    @property
    @pulumi.getter
    def status(self) -> Optional[pulumi.Input[builtins.str]]:
        return pulumi.get(self, "status")

    @status.setter
    def status(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "status", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]]]:
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]]]):
        pulumi.set(self, "tags", value)

    @property
    @pulumi.getter
    def type(self) -> Optional[pulumi.Input[builtins.str]]:
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "type", value)

    @property
    @pulumi.getter(name="updatedAt")
    def updated_at(self) -> Optional[pulumi.Input[builtins.str]]:
        return pulumi.get(self, "updated_at")

    @updated_at.setter
    def updated_at(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "updated_at", value)


class DliPackage(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 group_name: Optional[pulumi.Input[builtins.str]] = None,
                 is_async: Optional[pulumi.Input[builtins.bool]] = None,
                 object_path: Optional[pulumi.Input[builtins.str]] = None,
                 owner: Optional[pulumi.Input[builtins.str]] = None,
                 region: Optional[pulumi.Input[builtins.str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]]] = None,
                 type: Optional[pulumi.Input[builtins.str]] = None,
                 __props__=None):
        """
        Create a DliPackage resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: DliPackageArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a DliPackage resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param DliPackageArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(DliPackageArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 group_name: Optional[pulumi.Input[builtins.str]] = None,
                 is_async: Optional[pulumi.Input[builtins.bool]] = None,
                 object_path: Optional[pulumi.Input[builtins.str]] = None,
                 owner: Optional[pulumi.Input[builtins.str]] = None,
                 region: Optional[pulumi.Input[builtins.str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]]] = None,
                 type: Optional[pulumi.Input[builtins.str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = DliPackageArgs.__new__(DliPackageArgs)

            __props__.__dict__["group_name"] = group_name
            __props__.__dict__["is_async"] = is_async
            if object_path is None and not opts.urn:
                raise TypeError("Missing required property 'object_path'")
            __props__.__dict__["object_path"] = object_path
            __props__.__dict__["owner"] = owner
            __props__.__dict__["region"] = region
            __props__.__dict__["tags"] = tags
            if type is None and not opts.urn:
                raise TypeError("Missing required property 'type'")
            __props__.__dict__["type"] = type
            __props__.__dict__["created_at"] = None
            __props__.__dict__["object_name"] = None
            __props__.__dict__["status"] = None
            __props__.__dict__["updated_at"] = None
        super(DliPackage, __self__).__init__(
            'sbercloud:index/dliPackage:DliPackage',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            created_at: Optional[pulumi.Input[builtins.str]] = None,
            group_name: Optional[pulumi.Input[builtins.str]] = None,
            is_async: Optional[pulumi.Input[builtins.bool]] = None,
            object_name: Optional[pulumi.Input[builtins.str]] = None,
            object_path: Optional[pulumi.Input[builtins.str]] = None,
            owner: Optional[pulumi.Input[builtins.str]] = None,
            region: Optional[pulumi.Input[builtins.str]] = None,
            status: Optional[pulumi.Input[builtins.str]] = None,
            tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]]] = None,
            type: Optional[pulumi.Input[builtins.str]] = None,
            updated_at: Optional[pulumi.Input[builtins.str]] = None) -> 'DliPackage':
        """
        Get an existing DliPackage resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _DliPackageState.__new__(_DliPackageState)

        __props__.__dict__["created_at"] = created_at
        __props__.__dict__["group_name"] = group_name
        __props__.__dict__["is_async"] = is_async
        __props__.__dict__["object_name"] = object_name
        __props__.__dict__["object_path"] = object_path
        __props__.__dict__["owner"] = owner
        __props__.__dict__["region"] = region
        __props__.__dict__["status"] = status
        __props__.__dict__["tags"] = tags
        __props__.__dict__["type"] = type
        __props__.__dict__["updated_at"] = updated_at
        return DliPackage(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="createdAt")
    def created_at(self) -> pulumi.Output[builtins.str]:
        return pulumi.get(self, "created_at")

    @property
    @pulumi.getter(name="groupName")
    def group_name(self) -> pulumi.Output[Optional[builtins.str]]:
        return pulumi.get(self, "group_name")

    @property
    @pulumi.getter(name="isAsync")
    def is_async(self) -> pulumi.Output[builtins.bool]:
        return pulumi.get(self, "is_async")

    @property
    @pulumi.getter(name="objectName")
    def object_name(self) -> pulumi.Output[builtins.str]:
        return pulumi.get(self, "object_name")

    @property
    @pulumi.getter(name="objectPath")
    def object_path(self) -> pulumi.Output[builtins.str]:
        return pulumi.get(self, "object_path")

    @property
    @pulumi.getter
    def owner(self) -> pulumi.Output[builtins.str]:
        return pulumi.get(self, "owner")

    @property
    @pulumi.getter
    def region(self) -> pulumi.Output[builtins.str]:
        return pulumi.get(self, "region")

    @property
    @pulumi.getter
    def status(self) -> pulumi.Output[builtins.str]:
        return pulumi.get(self, "status")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Mapping[str, builtins.str]]]:
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter
    def type(self) -> pulumi.Output[builtins.str]:
        return pulumi.get(self, "type")

    @property
    @pulumi.getter(name="updatedAt")
    def updated_at(self) -> pulumi.Output[builtins.str]:
        return pulumi.get(self, "updated_at")

