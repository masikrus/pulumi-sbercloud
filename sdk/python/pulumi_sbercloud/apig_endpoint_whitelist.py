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

__all__ = ['ApigEndpointWhitelistArgs', 'ApigEndpointWhitelist']

@pulumi.input_type
class ApigEndpointWhitelistArgs:
    def __init__(__self__, *,
                 instance_id: pulumi.Input[builtins.str],
                 whitelists: pulumi.Input[Sequence[pulumi.Input[builtins.str]]],
                 region: Optional[pulumi.Input[builtins.str]] = None):
        """
        The set of arguments for constructing a ApigEndpointWhitelist resource.
        :param pulumi.Input[builtins.str] instance_id: The ID of the dedicated instance to which the endpoint service belongs.
        :param pulumi.Input[Sequence[pulumi.Input[builtins.str]]] whitelists: The whitelist records of the endpoint service.
        :param pulumi.Input[builtins.str] region: The region where the endpoint service is located.
        """
        pulumi.set(__self__, "instance_id", instance_id)
        pulumi.set(__self__, "whitelists", whitelists)
        if region is not None:
            pulumi.set(__self__, "region", region)

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> pulumi.Input[builtins.str]:
        """
        The ID of the dedicated instance to which the endpoint service belongs.
        """
        return pulumi.get(self, "instance_id")

    @instance_id.setter
    def instance_id(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "instance_id", value)

    @property
    @pulumi.getter
    def whitelists(self) -> pulumi.Input[Sequence[pulumi.Input[builtins.str]]]:
        """
        The whitelist records of the endpoint service.
        """
        return pulumi.get(self, "whitelists")

    @whitelists.setter
    def whitelists(self, value: pulumi.Input[Sequence[pulumi.Input[builtins.str]]]):
        pulumi.set(self, "whitelists", value)

    @property
    @pulumi.getter
    def region(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The region where the endpoint service is located.
        """
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "region", value)


@pulumi.input_type
class _ApigEndpointWhitelistState:
    def __init__(__self__, *,
                 instance_id: Optional[pulumi.Input[builtins.str]] = None,
                 region: Optional[pulumi.Input[builtins.str]] = None,
                 whitelists: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]] = None):
        """
        Input properties used for looking up and filtering ApigEndpointWhitelist resources.
        :param pulumi.Input[builtins.str] instance_id: The ID of the dedicated instance to which the endpoint service belongs.
        :param pulumi.Input[builtins.str] region: The region where the endpoint service is located.
        :param pulumi.Input[Sequence[pulumi.Input[builtins.str]]] whitelists: The whitelist records of the endpoint service.
        """
        if instance_id is not None:
            pulumi.set(__self__, "instance_id", instance_id)
        if region is not None:
            pulumi.set(__self__, "region", region)
        if whitelists is not None:
            pulumi.set(__self__, "whitelists", whitelists)

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The ID of the dedicated instance to which the endpoint service belongs.
        """
        return pulumi.get(self, "instance_id")

    @instance_id.setter
    def instance_id(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "instance_id", value)

    @property
    @pulumi.getter
    def region(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The region where the endpoint service is located.
        """
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "region", value)

    @property
    @pulumi.getter
    def whitelists(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]]:
        """
        The whitelist records of the endpoint service.
        """
        return pulumi.get(self, "whitelists")

    @whitelists.setter
    def whitelists(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]]):
        pulumi.set(self, "whitelists", value)


class ApigEndpointWhitelist(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 instance_id: Optional[pulumi.Input[builtins.str]] = None,
                 region: Optional[pulumi.Input[builtins.str]] = None,
                 whitelists: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]] = None,
                 __props__=None):
        """
        Create a ApigEndpointWhitelist resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[builtins.str] instance_id: The ID of the dedicated instance to which the endpoint service belongs.
        :param pulumi.Input[builtins.str] region: The region where the endpoint service is located.
        :param pulumi.Input[Sequence[pulumi.Input[builtins.str]]] whitelists: The whitelist records of the endpoint service.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ApigEndpointWhitelistArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a ApigEndpointWhitelist resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param ApigEndpointWhitelistArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ApigEndpointWhitelistArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 instance_id: Optional[pulumi.Input[builtins.str]] = None,
                 region: Optional[pulumi.Input[builtins.str]] = None,
                 whitelists: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ApigEndpointWhitelistArgs.__new__(ApigEndpointWhitelistArgs)

            if instance_id is None and not opts.urn:
                raise TypeError("Missing required property 'instance_id'")
            __props__.__dict__["instance_id"] = instance_id
            __props__.__dict__["region"] = region
            if whitelists is None and not opts.urn:
                raise TypeError("Missing required property 'whitelists'")
            __props__.__dict__["whitelists"] = whitelists
        super(ApigEndpointWhitelist, __self__).__init__(
            'sbercloud:index/apigEndpointWhitelist:ApigEndpointWhitelist',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            instance_id: Optional[pulumi.Input[builtins.str]] = None,
            region: Optional[pulumi.Input[builtins.str]] = None,
            whitelists: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]] = None) -> 'ApigEndpointWhitelist':
        """
        Get an existing ApigEndpointWhitelist resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[builtins.str] instance_id: The ID of the dedicated instance to which the endpoint service belongs.
        :param pulumi.Input[builtins.str] region: The region where the endpoint service is located.
        :param pulumi.Input[Sequence[pulumi.Input[builtins.str]]] whitelists: The whitelist records of the endpoint service.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _ApigEndpointWhitelistState.__new__(_ApigEndpointWhitelistState)

        __props__.__dict__["instance_id"] = instance_id
        __props__.__dict__["region"] = region
        __props__.__dict__["whitelists"] = whitelists
        return ApigEndpointWhitelist(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> pulumi.Output[builtins.str]:
        """
        The ID of the dedicated instance to which the endpoint service belongs.
        """
        return pulumi.get(self, "instance_id")

    @property
    @pulumi.getter
    def region(self) -> pulumi.Output[builtins.str]:
        """
        The region where the endpoint service is located.
        """
        return pulumi.get(self, "region")

    @property
    @pulumi.getter
    def whitelists(self) -> pulumi.Output[Sequence[builtins.str]]:
        """
        The whitelist records of the endpoint service.
        """
        return pulumi.get(self, "whitelists")

