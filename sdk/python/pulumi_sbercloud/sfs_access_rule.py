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

__all__ = ['SfsAccessRuleArgs', 'SfsAccessRule']

@pulumi.input_type
class SfsAccessRuleArgs:
    def __init__(__self__, *,
                 access_to: pulumi.Input[builtins.str],
                 sfs_id: pulumi.Input[builtins.str],
                 access_level: Optional[pulumi.Input[builtins.str]] = None,
                 access_type: Optional[pulumi.Input[builtins.str]] = None,
                 region: Optional[pulumi.Input[builtins.str]] = None):
        """
        The set of arguments for constructing a SfsAccessRule resource.
        """
        pulumi.set(__self__, "access_to", access_to)
        pulumi.set(__self__, "sfs_id", sfs_id)
        if access_level is not None:
            pulumi.set(__self__, "access_level", access_level)
        if access_type is not None:
            pulumi.set(__self__, "access_type", access_type)
        if region is not None:
            pulumi.set(__self__, "region", region)

    @property
    @pulumi.getter(name="accessTo")
    def access_to(self) -> pulumi.Input[builtins.str]:
        return pulumi.get(self, "access_to")

    @access_to.setter
    def access_to(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "access_to", value)

    @property
    @pulumi.getter(name="sfsId")
    def sfs_id(self) -> pulumi.Input[builtins.str]:
        return pulumi.get(self, "sfs_id")

    @sfs_id.setter
    def sfs_id(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "sfs_id", value)

    @property
    @pulumi.getter(name="accessLevel")
    def access_level(self) -> Optional[pulumi.Input[builtins.str]]:
        return pulumi.get(self, "access_level")

    @access_level.setter
    def access_level(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "access_level", value)

    @property
    @pulumi.getter(name="accessType")
    def access_type(self) -> Optional[pulumi.Input[builtins.str]]:
        return pulumi.get(self, "access_type")

    @access_type.setter
    def access_type(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "access_type", value)

    @property
    @pulumi.getter
    def region(self) -> Optional[pulumi.Input[builtins.str]]:
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "region", value)


@pulumi.input_type
class _SfsAccessRuleState:
    def __init__(__self__, *,
                 access_level: Optional[pulumi.Input[builtins.str]] = None,
                 access_to: Optional[pulumi.Input[builtins.str]] = None,
                 access_type: Optional[pulumi.Input[builtins.str]] = None,
                 region: Optional[pulumi.Input[builtins.str]] = None,
                 sfs_id: Optional[pulumi.Input[builtins.str]] = None,
                 status: Optional[pulumi.Input[builtins.str]] = None):
        """
        Input properties used for looking up and filtering SfsAccessRule resources.
        """
        if access_level is not None:
            pulumi.set(__self__, "access_level", access_level)
        if access_to is not None:
            pulumi.set(__self__, "access_to", access_to)
        if access_type is not None:
            pulumi.set(__self__, "access_type", access_type)
        if region is not None:
            pulumi.set(__self__, "region", region)
        if sfs_id is not None:
            pulumi.set(__self__, "sfs_id", sfs_id)
        if status is not None:
            pulumi.set(__self__, "status", status)

    @property
    @pulumi.getter(name="accessLevel")
    def access_level(self) -> Optional[pulumi.Input[builtins.str]]:
        return pulumi.get(self, "access_level")

    @access_level.setter
    def access_level(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "access_level", value)

    @property
    @pulumi.getter(name="accessTo")
    def access_to(self) -> Optional[pulumi.Input[builtins.str]]:
        return pulumi.get(self, "access_to")

    @access_to.setter
    def access_to(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "access_to", value)

    @property
    @pulumi.getter(name="accessType")
    def access_type(self) -> Optional[pulumi.Input[builtins.str]]:
        return pulumi.get(self, "access_type")

    @access_type.setter
    def access_type(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "access_type", value)

    @property
    @pulumi.getter
    def region(self) -> Optional[pulumi.Input[builtins.str]]:
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "region", value)

    @property
    @pulumi.getter(name="sfsId")
    def sfs_id(self) -> Optional[pulumi.Input[builtins.str]]:
        return pulumi.get(self, "sfs_id")

    @sfs_id.setter
    def sfs_id(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "sfs_id", value)

    @property
    @pulumi.getter
    def status(self) -> Optional[pulumi.Input[builtins.str]]:
        return pulumi.get(self, "status")

    @status.setter
    def status(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "status", value)


class SfsAccessRule(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 access_level: Optional[pulumi.Input[builtins.str]] = None,
                 access_to: Optional[pulumi.Input[builtins.str]] = None,
                 access_type: Optional[pulumi.Input[builtins.str]] = None,
                 region: Optional[pulumi.Input[builtins.str]] = None,
                 sfs_id: Optional[pulumi.Input[builtins.str]] = None,
                 __props__=None):
        """
        Create a SfsAccessRule resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: SfsAccessRuleArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a SfsAccessRule resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param SfsAccessRuleArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(SfsAccessRuleArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 access_level: Optional[pulumi.Input[builtins.str]] = None,
                 access_to: Optional[pulumi.Input[builtins.str]] = None,
                 access_type: Optional[pulumi.Input[builtins.str]] = None,
                 region: Optional[pulumi.Input[builtins.str]] = None,
                 sfs_id: Optional[pulumi.Input[builtins.str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = SfsAccessRuleArgs.__new__(SfsAccessRuleArgs)

            __props__.__dict__["access_level"] = access_level
            if access_to is None and not opts.urn:
                raise TypeError("Missing required property 'access_to'")
            __props__.__dict__["access_to"] = access_to
            __props__.__dict__["access_type"] = access_type
            __props__.__dict__["region"] = region
            if sfs_id is None and not opts.urn:
                raise TypeError("Missing required property 'sfs_id'")
            __props__.__dict__["sfs_id"] = sfs_id
            __props__.__dict__["status"] = None
        super(SfsAccessRule, __self__).__init__(
            'sbercloud:index/sfsAccessRule:SfsAccessRule',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            access_level: Optional[pulumi.Input[builtins.str]] = None,
            access_to: Optional[pulumi.Input[builtins.str]] = None,
            access_type: Optional[pulumi.Input[builtins.str]] = None,
            region: Optional[pulumi.Input[builtins.str]] = None,
            sfs_id: Optional[pulumi.Input[builtins.str]] = None,
            status: Optional[pulumi.Input[builtins.str]] = None) -> 'SfsAccessRule':
        """
        Get an existing SfsAccessRule resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _SfsAccessRuleState.__new__(_SfsAccessRuleState)

        __props__.__dict__["access_level"] = access_level
        __props__.__dict__["access_to"] = access_to
        __props__.__dict__["access_type"] = access_type
        __props__.__dict__["region"] = region
        __props__.__dict__["sfs_id"] = sfs_id
        __props__.__dict__["status"] = status
        return SfsAccessRule(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="accessLevel")
    def access_level(self) -> pulumi.Output[Optional[builtins.str]]:
        return pulumi.get(self, "access_level")

    @property
    @pulumi.getter(name="accessTo")
    def access_to(self) -> pulumi.Output[builtins.str]:
        return pulumi.get(self, "access_to")

    @property
    @pulumi.getter(name="accessType")
    def access_type(self) -> pulumi.Output[Optional[builtins.str]]:
        return pulumi.get(self, "access_type")

    @property
    @pulumi.getter
    def region(self) -> pulumi.Output[builtins.str]:
        return pulumi.get(self, "region")

    @property
    @pulumi.getter(name="sfsId")
    def sfs_id(self) -> pulumi.Output[builtins.str]:
        return pulumi.get(self, "sfs_id")

    @property
    @pulumi.getter
    def status(self) -> pulumi.Output[builtins.str]:
        return pulumi.get(self, "status")

