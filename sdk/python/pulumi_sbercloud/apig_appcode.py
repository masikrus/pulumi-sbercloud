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

__all__ = ['ApigAppcodeArgs', 'ApigAppcode']

@pulumi.input_type
class ApigAppcodeArgs:
    def __init__(__self__, *,
                 application_id: pulumi.Input[builtins.str],
                 instance_id: pulumi.Input[builtins.str],
                 region: Optional[pulumi.Input[builtins.str]] = None,
                 value: Optional[pulumi.Input[builtins.str]] = None):
        """
        The set of arguments for constructing a ApigAppcode resource.
        :param pulumi.Input[builtins.str] application_id: The ID of the application to which the APPCODE belongs.
        :param pulumi.Input[builtins.str] instance_id: The ID of the dedicated instance to which the application and APPCODE belong.
        :param pulumi.Input[builtins.str] region: The region where the application and APPCODE are located.
        :param pulumi.Input[builtins.str] value: The APPCODE value (content).
        """
        pulumi.set(__self__, "application_id", application_id)
        pulumi.set(__self__, "instance_id", instance_id)
        if region is not None:
            pulumi.set(__self__, "region", region)
        if value is not None:
            pulumi.set(__self__, "value", value)

    @property
    @pulumi.getter(name="applicationId")
    def application_id(self) -> pulumi.Input[builtins.str]:
        """
        The ID of the application to which the APPCODE belongs.
        """
        return pulumi.get(self, "application_id")

    @application_id.setter
    def application_id(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "application_id", value)

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> pulumi.Input[builtins.str]:
        """
        The ID of the dedicated instance to which the application and APPCODE belong.
        """
        return pulumi.get(self, "instance_id")

    @instance_id.setter
    def instance_id(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "instance_id", value)

    @property
    @pulumi.getter
    def region(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The region where the application and APPCODE are located.
        """
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "region", value)

    @property
    @pulumi.getter
    def value(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The APPCODE value (content).
        """
        return pulumi.get(self, "value")

    @value.setter
    def value(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "value", value)


@pulumi.input_type
class _ApigAppcodeState:
    def __init__(__self__, *,
                 application_id: Optional[pulumi.Input[builtins.str]] = None,
                 created_at: Optional[pulumi.Input[builtins.str]] = None,
                 instance_id: Optional[pulumi.Input[builtins.str]] = None,
                 region: Optional[pulumi.Input[builtins.str]] = None,
                 value: Optional[pulumi.Input[builtins.str]] = None):
        """
        Input properties used for looking up and filtering ApigAppcode resources.
        :param pulumi.Input[builtins.str] application_id: The ID of the application to which the APPCODE belongs.
        :param pulumi.Input[builtins.str] created_at: The creation time of the APPCODE.
        :param pulumi.Input[builtins.str] instance_id: The ID of the dedicated instance to which the application and APPCODE belong.
        :param pulumi.Input[builtins.str] region: The region where the application and APPCODE are located.
        :param pulumi.Input[builtins.str] value: The APPCODE value (content).
        """
        if application_id is not None:
            pulumi.set(__self__, "application_id", application_id)
        if created_at is not None:
            pulumi.set(__self__, "created_at", created_at)
        if instance_id is not None:
            pulumi.set(__self__, "instance_id", instance_id)
        if region is not None:
            pulumi.set(__self__, "region", region)
        if value is not None:
            pulumi.set(__self__, "value", value)

    @property
    @pulumi.getter(name="applicationId")
    def application_id(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The ID of the application to which the APPCODE belongs.
        """
        return pulumi.get(self, "application_id")

    @application_id.setter
    def application_id(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "application_id", value)

    @property
    @pulumi.getter(name="createdAt")
    def created_at(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The creation time of the APPCODE.
        """
        return pulumi.get(self, "created_at")

    @created_at.setter
    def created_at(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "created_at", value)

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The ID of the dedicated instance to which the application and APPCODE belong.
        """
        return pulumi.get(self, "instance_id")

    @instance_id.setter
    def instance_id(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "instance_id", value)

    @property
    @pulumi.getter
    def region(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The region where the application and APPCODE are located.
        """
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "region", value)

    @property
    @pulumi.getter
    def value(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The APPCODE value (content).
        """
        return pulumi.get(self, "value")

    @value.setter
    def value(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "value", value)


class ApigAppcode(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 application_id: Optional[pulumi.Input[builtins.str]] = None,
                 instance_id: Optional[pulumi.Input[builtins.str]] = None,
                 region: Optional[pulumi.Input[builtins.str]] = None,
                 value: Optional[pulumi.Input[builtins.str]] = None,
                 __props__=None):
        """
        Create a ApigAppcode resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[builtins.str] application_id: The ID of the application to which the APPCODE belongs.
        :param pulumi.Input[builtins.str] instance_id: The ID of the dedicated instance to which the application and APPCODE belong.
        :param pulumi.Input[builtins.str] region: The region where the application and APPCODE are located.
        :param pulumi.Input[builtins.str] value: The APPCODE value (content).
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ApigAppcodeArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a ApigAppcode resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param ApigAppcodeArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ApigAppcodeArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 application_id: Optional[pulumi.Input[builtins.str]] = None,
                 instance_id: Optional[pulumi.Input[builtins.str]] = None,
                 region: Optional[pulumi.Input[builtins.str]] = None,
                 value: Optional[pulumi.Input[builtins.str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ApigAppcodeArgs.__new__(ApigAppcodeArgs)

            if application_id is None and not opts.urn:
                raise TypeError("Missing required property 'application_id'")
            __props__.__dict__["application_id"] = application_id
            if instance_id is None and not opts.urn:
                raise TypeError("Missing required property 'instance_id'")
            __props__.__dict__["instance_id"] = instance_id
            __props__.__dict__["region"] = region
            __props__.__dict__["value"] = value
            __props__.__dict__["created_at"] = None
        super(ApigAppcode, __self__).__init__(
            'sbercloud:index/apigAppcode:ApigAppcode',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            application_id: Optional[pulumi.Input[builtins.str]] = None,
            created_at: Optional[pulumi.Input[builtins.str]] = None,
            instance_id: Optional[pulumi.Input[builtins.str]] = None,
            region: Optional[pulumi.Input[builtins.str]] = None,
            value: Optional[pulumi.Input[builtins.str]] = None) -> 'ApigAppcode':
        """
        Get an existing ApigAppcode resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[builtins.str] application_id: The ID of the application to which the APPCODE belongs.
        :param pulumi.Input[builtins.str] created_at: The creation time of the APPCODE.
        :param pulumi.Input[builtins.str] instance_id: The ID of the dedicated instance to which the application and APPCODE belong.
        :param pulumi.Input[builtins.str] region: The region where the application and APPCODE are located.
        :param pulumi.Input[builtins.str] value: The APPCODE value (content).
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _ApigAppcodeState.__new__(_ApigAppcodeState)

        __props__.__dict__["application_id"] = application_id
        __props__.__dict__["created_at"] = created_at
        __props__.__dict__["instance_id"] = instance_id
        __props__.__dict__["region"] = region
        __props__.__dict__["value"] = value
        return ApigAppcode(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="applicationId")
    def application_id(self) -> pulumi.Output[builtins.str]:
        """
        The ID of the application to which the APPCODE belongs.
        """
        return pulumi.get(self, "application_id")

    @property
    @pulumi.getter(name="createdAt")
    def created_at(self) -> pulumi.Output[builtins.str]:
        """
        The creation time of the APPCODE.
        """
        return pulumi.get(self, "created_at")

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> pulumi.Output[builtins.str]:
        """
        The ID of the dedicated instance to which the application and APPCODE belong.
        """
        return pulumi.get(self, "instance_id")

    @property
    @pulumi.getter
    def region(self) -> pulumi.Output[builtins.str]:
        """
        The region where the application and APPCODE are located.
        """
        return pulumi.get(self, "region")

    @property
    @pulumi.getter
    def value(self) -> pulumi.Output[builtins.str]:
        """
        The APPCODE value (content).
        """
        return pulumi.get(self, "value")

