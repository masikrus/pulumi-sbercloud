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

__all__ = ['IdentityGroupMembershipV3Args', 'IdentityGroupMembershipV3']

@pulumi.input_type
class IdentityGroupMembershipV3Args:
    def __init__(__self__, *,
                 group: pulumi.Input[builtins.str],
                 users: pulumi.Input[Sequence[pulumi.Input[builtins.str]]]):
        """
        The set of arguments for constructing a IdentityGroupMembershipV3 resource.
        """
        pulumi.set(__self__, "group", group)
        pulumi.set(__self__, "users", users)

    @property
    @pulumi.getter
    def group(self) -> pulumi.Input[builtins.str]:
        return pulumi.get(self, "group")

    @group.setter
    def group(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "group", value)

    @property
    @pulumi.getter
    def users(self) -> pulumi.Input[Sequence[pulumi.Input[builtins.str]]]:
        return pulumi.get(self, "users")

    @users.setter
    def users(self, value: pulumi.Input[Sequence[pulumi.Input[builtins.str]]]):
        pulumi.set(self, "users", value)


@pulumi.input_type
class _IdentityGroupMembershipV3State:
    def __init__(__self__, *,
                 group: Optional[pulumi.Input[builtins.str]] = None,
                 users: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]] = None):
        """
        Input properties used for looking up and filtering IdentityGroupMembershipV3 resources.
        """
        if group is not None:
            pulumi.set(__self__, "group", group)
        if users is not None:
            pulumi.set(__self__, "users", users)

    @property
    @pulumi.getter
    def group(self) -> Optional[pulumi.Input[builtins.str]]:
        return pulumi.get(self, "group")

    @group.setter
    def group(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "group", value)

    @property
    @pulumi.getter
    def users(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]]:
        return pulumi.get(self, "users")

    @users.setter
    def users(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]]):
        pulumi.set(self, "users", value)


class IdentityGroupMembershipV3(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 group: Optional[pulumi.Input[builtins.str]] = None,
                 users: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]] = None,
                 __props__=None):
        """
        Create a IdentityGroupMembershipV3 resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: IdentityGroupMembershipV3Args,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a IdentityGroupMembershipV3 resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param IdentityGroupMembershipV3Args args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(IdentityGroupMembershipV3Args, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 group: Optional[pulumi.Input[builtins.str]] = None,
                 users: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = IdentityGroupMembershipV3Args.__new__(IdentityGroupMembershipV3Args)

            if group is None and not opts.urn:
                raise TypeError("Missing required property 'group'")
            __props__.__dict__["group"] = group
            if users is None and not opts.urn:
                raise TypeError("Missing required property 'users'")
            __props__.__dict__["users"] = users
        super(IdentityGroupMembershipV3, __self__).__init__(
            'sbercloud:index/identityGroupMembershipV3:IdentityGroupMembershipV3',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            group: Optional[pulumi.Input[builtins.str]] = None,
            users: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]] = None) -> 'IdentityGroupMembershipV3':
        """
        Get an existing IdentityGroupMembershipV3 resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _IdentityGroupMembershipV3State.__new__(_IdentityGroupMembershipV3State)

        __props__.__dict__["group"] = group
        __props__.__dict__["users"] = users
        return IdentityGroupMembershipV3(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def group(self) -> pulumi.Output[builtins.str]:
        return pulumi.get(self, "group")

    @property
    @pulumi.getter
    def users(self) -> pulumi.Output[Sequence[builtins.str]]:
        return pulumi.get(self, "users")

