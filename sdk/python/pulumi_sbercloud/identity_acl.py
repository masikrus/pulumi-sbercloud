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

__all__ = ['IdentityAclArgs', 'IdentityAcl']

@pulumi.input_type
class IdentityAclArgs:
    def __init__(__self__, *,
                 type: pulumi.Input[builtins.str],
                 ip_cidrs: Optional[pulumi.Input[Sequence[pulumi.Input['IdentityAclIpCidrArgs']]]] = None,
                 ip_ranges: Optional[pulumi.Input[Sequence[pulumi.Input['IdentityAclIpRangeArgs']]]] = None):
        """
        The set of arguments for constructing a IdentityAcl resource.
        """
        pulumi.set(__self__, "type", type)
        if ip_cidrs is not None:
            pulumi.set(__self__, "ip_cidrs", ip_cidrs)
        if ip_ranges is not None:
            pulumi.set(__self__, "ip_ranges", ip_ranges)

    @property
    @pulumi.getter
    def type(self) -> pulumi.Input[builtins.str]:
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "type", value)

    @property
    @pulumi.getter(name="ipCidrs")
    def ip_cidrs(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['IdentityAclIpCidrArgs']]]]:
        return pulumi.get(self, "ip_cidrs")

    @ip_cidrs.setter
    def ip_cidrs(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['IdentityAclIpCidrArgs']]]]):
        pulumi.set(self, "ip_cidrs", value)

    @property
    @pulumi.getter(name="ipRanges")
    def ip_ranges(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['IdentityAclIpRangeArgs']]]]:
        return pulumi.get(self, "ip_ranges")

    @ip_ranges.setter
    def ip_ranges(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['IdentityAclIpRangeArgs']]]]):
        pulumi.set(self, "ip_ranges", value)


@pulumi.input_type
class _IdentityAclState:
    def __init__(__self__, *,
                 ip_cidrs: Optional[pulumi.Input[Sequence[pulumi.Input['IdentityAclIpCidrArgs']]]] = None,
                 ip_ranges: Optional[pulumi.Input[Sequence[pulumi.Input['IdentityAclIpRangeArgs']]]] = None,
                 type: Optional[pulumi.Input[builtins.str]] = None):
        """
        Input properties used for looking up and filtering IdentityAcl resources.
        """
        if ip_cidrs is not None:
            pulumi.set(__self__, "ip_cidrs", ip_cidrs)
        if ip_ranges is not None:
            pulumi.set(__self__, "ip_ranges", ip_ranges)
        if type is not None:
            pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter(name="ipCidrs")
    def ip_cidrs(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['IdentityAclIpCidrArgs']]]]:
        return pulumi.get(self, "ip_cidrs")

    @ip_cidrs.setter
    def ip_cidrs(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['IdentityAclIpCidrArgs']]]]):
        pulumi.set(self, "ip_cidrs", value)

    @property
    @pulumi.getter(name="ipRanges")
    def ip_ranges(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['IdentityAclIpRangeArgs']]]]:
        return pulumi.get(self, "ip_ranges")

    @ip_ranges.setter
    def ip_ranges(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['IdentityAclIpRangeArgs']]]]):
        pulumi.set(self, "ip_ranges", value)

    @property
    @pulumi.getter
    def type(self) -> Optional[pulumi.Input[builtins.str]]:
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "type", value)


class IdentityAcl(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 ip_cidrs: Optional[pulumi.Input[Sequence[pulumi.Input[Union['IdentityAclIpCidrArgs', 'IdentityAclIpCidrArgsDict']]]]] = None,
                 ip_ranges: Optional[pulumi.Input[Sequence[pulumi.Input[Union['IdentityAclIpRangeArgs', 'IdentityAclIpRangeArgsDict']]]]] = None,
                 type: Optional[pulumi.Input[builtins.str]] = None,
                 __props__=None):
        """
        Create a IdentityAcl resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: IdentityAclArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a IdentityAcl resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param IdentityAclArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(IdentityAclArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 ip_cidrs: Optional[pulumi.Input[Sequence[pulumi.Input[Union['IdentityAclIpCidrArgs', 'IdentityAclIpCidrArgsDict']]]]] = None,
                 ip_ranges: Optional[pulumi.Input[Sequence[pulumi.Input[Union['IdentityAclIpRangeArgs', 'IdentityAclIpRangeArgsDict']]]]] = None,
                 type: Optional[pulumi.Input[builtins.str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = IdentityAclArgs.__new__(IdentityAclArgs)

            __props__.__dict__["ip_cidrs"] = ip_cidrs
            __props__.__dict__["ip_ranges"] = ip_ranges
            if type is None and not opts.urn:
                raise TypeError("Missing required property 'type'")
            __props__.__dict__["type"] = type
        super(IdentityAcl, __self__).__init__(
            'sbercloud:index/identityAcl:IdentityAcl',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            ip_cidrs: Optional[pulumi.Input[Sequence[pulumi.Input[Union['IdentityAclIpCidrArgs', 'IdentityAclIpCidrArgsDict']]]]] = None,
            ip_ranges: Optional[pulumi.Input[Sequence[pulumi.Input[Union['IdentityAclIpRangeArgs', 'IdentityAclIpRangeArgsDict']]]]] = None,
            type: Optional[pulumi.Input[builtins.str]] = None) -> 'IdentityAcl':
        """
        Get an existing IdentityAcl resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _IdentityAclState.__new__(_IdentityAclState)

        __props__.__dict__["ip_cidrs"] = ip_cidrs
        __props__.__dict__["ip_ranges"] = ip_ranges
        __props__.__dict__["type"] = type
        return IdentityAcl(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="ipCidrs")
    def ip_cidrs(self) -> pulumi.Output[Optional[Sequence['outputs.IdentityAclIpCidr']]]:
        return pulumi.get(self, "ip_cidrs")

    @property
    @pulumi.getter(name="ipRanges")
    def ip_ranges(self) -> pulumi.Output[Optional[Sequence['outputs.IdentityAclIpRange']]]:
        return pulumi.get(self, "ip_ranges")

    @property
    @pulumi.getter
    def type(self) -> pulumi.Output[builtins.str]:
        return pulumi.get(self, "type")

