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

__all__ = ['ApigEnvironmentVariableArgs', 'ApigEnvironmentVariable']

@pulumi.input_type
class ApigEnvironmentVariableArgs:
    def __init__(__self__, *,
                 env_id: pulumi.Input[builtins.str],
                 group_id: pulumi.Input[builtins.str],
                 instance_id: pulumi.Input[builtins.str],
                 value: pulumi.Input[builtins.str],
                 name: Optional[pulumi.Input[builtins.str]] = None,
                 region: Optional[pulumi.Input[builtins.str]] = None):
        """
        The set of arguments for constructing a ApigEnvironmentVariable resource.
        :param pulumi.Input[builtins.str] env_id: Specifies the ID of the environment to which the environment variable belongs.
        :param pulumi.Input[builtins.str] group_id: Specifies the ID of the group to which the environment variable belongs.
        :param pulumi.Input[builtins.str] instance_id: Specifies the ID of the dedicated instance to which the environment variable belongs.
        :param pulumi.Input[builtins.str] value: Specifies the value of the environment variable.
        :param pulumi.Input[builtins.str] name: Specifies the name of the environment variable.
        :param pulumi.Input[builtins.str] region: Specifies the region in which to create the resource.
        """
        pulumi.set(__self__, "env_id", env_id)
        pulumi.set(__self__, "group_id", group_id)
        pulumi.set(__self__, "instance_id", instance_id)
        pulumi.set(__self__, "value", value)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if region is not None:
            pulumi.set(__self__, "region", region)

    @property
    @pulumi.getter(name="envId")
    def env_id(self) -> pulumi.Input[builtins.str]:
        """
        Specifies the ID of the environment to which the environment variable belongs.
        """
        return pulumi.get(self, "env_id")

    @env_id.setter
    def env_id(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "env_id", value)

    @property
    @pulumi.getter(name="groupId")
    def group_id(self) -> pulumi.Input[builtins.str]:
        """
        Specifies the ID of the group to which the environment variable belongs.
        """
        return pulumi.get(self, "group_id")

    @group_id.setter
    def group_id(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "group_id", value)

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> pulumi.Input[builtins.str]:
        """
        Specifies the ID of the dedicated instance to which the environment variable belongs.
        """
        return pulumi.get(self, "instance_id")

    @instance_id.setter
    def instance_id(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "instance_id", value)

    @property
    @pulumi.getter
    def value(self) -> pulumi.Input[builtins.str]:
        """
        Specifies the value of the environment variable.
        """
        return pulumi.get(self, "value")

    @value.setter
    def value(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "value", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        Specifies the name of the environment variable.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def region(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        Specifies the region in which to create the resource.
        """
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "region", value)


@pulumi.input_type
class _ApigEnvironmentVariableState:
    def __init__(__self__, *,
                 env_id: Optional[pulumi.Input[builtins.str]] = None,
                 group_id: Optional[pulumi.Input[builtins.str]] = None,
                 instance_id: Optional[pulumi.Input[builtins.str]] = None,
                 name: Optional[pulumi.Input[builtins.str]] = None,
                 region: Optional[pulumi.Input[builtins.str]] = None,
                 value: Optional[pulumi.Input[builtins.str]] = None):
        """
        Input properties used for looking up and filtering ApigEnvironmentVariable resources.
        :param pulumi.Input[builtins.str] env_id: Specifies the ID of the environment to which the environment variable belongs.
        :param pulumi.Input[builtins.str] group_id: Specifies the ID of the group to which the environment variable belongs.
        :param pulumi.Input[builtins.str] instance_id: Specifies the ID of the dedicated instance to which the environment variable belongs.
        :param pulumi.Input[builtins.str] name: Specifies the name of the environment variable.
        :param pulumi.Input[builtins.str] region: Specifies the region in which to create the resource.
        :param pulumi.Input[builtins.str] value: Specifies the value of the environment variable.
        """
        if env_id is not None:
            pulumi.set(__self__, "env_id", env_id)
        if group_id is not None:
            pulumi.set(__self__, "group_id", group_id)
        if instance_id is not None:
            pulumi.set(__self__, "instance_id", instance_id)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if region is not None:
            pulumi.set(__self__, "region", region)
        if value is not None:
            pulumi.set(__self__, "value", value)

    @property
    @pulumi.getter(name="envId")
    def env_id(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        Specifies the ID of the environment to which the environment variable belongs.
        """
        return pulumi.get(self, "env_id")

    @env_id.setter
    def env_id(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "env_id", value)

    @property
    @pulumi.getter(name="groupId")
    def group_id(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        Specifies the ID of the group to which the environment variable belongs.
        """
        return pulumi.get(self, "group_id")

    @group_id.setter
    def group_id(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "group_id", value)

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        Specifies the ID of the dedicated instance to which the environment variable belongs.
        """
        return pulumi.get(self, "instance_id")

    @instance_id.setter
    def instance_id(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "instance_id", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        Specifies the name of the environment variable.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def region(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        Specifies the region in which to create the resource.
        """
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "region", value)

    @property
    @pulumi.getter
    def value(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        Specifies the value of the environment variable.
        """
        return pulumi.get(self, "value")

    @value.setter
    def value(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "value", value)


class ApigEnvironmentVariable(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 env_id: Optional[pulumi.Input[builtins.str]] = None,
                 group_id: Optional[pulumi.Input[builtins.str]] = None,
                 instance_id: Optional[pulumi.Input[builtins.str]] = None,
                 name: Optional[pulumi.Input[builtins.str]] = None,
                 region: Optional[pulumi.Input[builtins.str]] = None,
                 value: Optional[pulumi.Input[builtins.str]] = None,
                 __props__=None):
        """
        Create a ApigEnvironmentVariable resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[builtins.str] env_id: Specifies the ID of the environment to which the environment variable belongs.
        :param pulumi.Input[builtins.str] group_id: Specifies the ID of the group to which the environment variable belongs.
        :param pulumi.Input[builtins.str] instance_id: Specifies the ID of the dedicated instance to which the environment variable belongs.
        :param pulumi.Input[builtins.str] name: Specifies the name of the environment variable.
        :param pulumi.Input[builtins.str] region: Specifies the region in which to create the resource.
        :param pulumi.Input[builtins.str] value: Specifies the value of the environment variable.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ApigEnvironmentVariableArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a ApigEnvironmentVariable resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param ApigEnvironmentVariableArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ApigEnvironmentVariableArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 env_id: Optional[pulumi.Input[builtins.str]] = None,
                 group_id: Optional[pulumi.Input[builtins.str]] = None,
                 instance_id: Optional[pulumi.Input[builtins.str]] = None,
                 name: Optional[pulumi.Input[builtins.str]] = None,
                 region: Optional[pulumi.Input[builtins.str]] = None,
                 value: Optional[pulumi.Input[builtins.str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ApigEnvironmentVariableArgs.__new__(ApigEnvironmentVariableArgs)

            if env_id is None and not opts.urn:
                raise TypeError("Missing required property 'env_id'")
            __props__.__dict__["env_id"] = env_id
            if group_id is None and not opts.urn:
                raise TypeError("Missing required property 'group_id'")
            __props__.__dict__["group_id"] = group_id
            if instance_id is None and not opts.urn:
                raise TypeError("Missing required property 'instance_id'")
            __props__.__dict__["instance_id"] = instance_id
            __props__.__dict__["name"] = name
            __props__.__dict__["region"] = region
            if value is None and not opts.urn:
                raise TypeError("Missing required property 'value'")
            __props__.__dict__["value"] = value
        super(ApigEnvironmentVariable, __self__).__init__(
            'sbercloud:index/apigEnvironmentVariable:ApigEnvironmentVariable',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            env_id: Optional[pulumi.Input[builtins.str]] = None,
            group_id: Optional[pulumi.Input[builtins.str]] = None,
            instance_id: Optional[pulumi.Input[builtins.str]] = None,
            name: Optional[pulumi.Input[builtins.str]] = None,
            region: Optional[pulumi.Input[builtins.str]] = None,
            value: Optional[pulumi.Input[builtins.str]] = None) -> 'ApigEnvironmentVariable':
        """
        Get an existing ApigEnvironmentVariable resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[builtins.str] env_id: Specifies the ID of the environment to which the environment variable belongs.
        :param pulumi.Input[builtins.str] group_id: Specifies the ID of the group to which the environment variable belongs.
        :param pulumi.Input[builtins.str] instance_id: Specifies the ID of the dedicated instance to which the environment variable belongs.
        :param pulumi.Input[builtins.str] name: Specifies the name of the environment variable.
        :param pulumi.Input[builtins.str] region: Specifies the region in which to create the resource.
        :param pulumi.Input[builtins.str] value: Specifies the value of the environment variable.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _ApigEnvironmentVariableState.__new__(_ApigEnvironmentVariableState)

        __props__.__dict__["env_id"] = env_id
        __props__.__dict__["group_id"] = group_id
        __props__.__dict__["instance_id"] = instance_id
        __props__.__dict__["name"] = name
        __props__.__dict__["region"] = region
        __props__.__dict__["value"] = value
        return ApigEnvironmentVariable(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="envId")
    def env_id(self) -> pulumi.Output[builtins.str]:
        """
        Specifies the ID of the environment to which the environment variable belongs.
        """
        return pulumi.get(self, "env_id")

    @property
    @pulumi.getter(name="groupId")
    def group_id(self) -> pulumi.Output[builtins.str]:
        """
        Specifies the ID of the group to which the environment variable belongs.
        """
        return pulumi.get(self, "group_id")

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> pulumi.Output[builtins.str]:
        """
        Specifies the ID of the dedicated instance to which the environment variable belongs.
        """
        return pulumi.get(self, "instance_id")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[builtins.str]:
        """
        Specifies the name of the environment variable.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def region(self) -> pulumi.Output[builtins.str]:
        """
        Specifies the region in which to create the resource.
        """
        return pulumi.get(self, "region")

    @property
    @pulumi.getter
    def value(self) -> pulumi.Output[builtins.str]:
        """
        Specifies the value of the environment variable.
        """
        return pulumi.get(self, "value")

