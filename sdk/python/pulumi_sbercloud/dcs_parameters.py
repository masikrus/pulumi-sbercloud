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

__all__ = ['DcsParametersArgs', 'DcsParameters']

@pulumi.input_type
class DcsParametersArgs:
    def __init__(__self__, *,
                 instance_id: pulumi.Input[builtins.str],
                 parameters: pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]],
                 project_id: pulumi.Input[builtins.str]):
        """
        The set of arguments for constructing a DcsParameters resource.
        """
        pulumi.set(__self__, "instance_id", instance_id)
        pulumi.set(__self__, "parameters", parameters)
        pulumi.set(__self__, "project_id", project_id)

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> pulumi.Input[builtins.str]:
        return pulumi.get(self, "instance_id")

    @instance_id.setter
    def instance_id(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "instance_id", value)

    @property
    @pulumi.getter
    def parameters(self) -> pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]]:
        return pulumi.get(self, "parameters")

    @parameters.setter
    def parameters(self, value: pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]]):
        pulumi.set(self, "parameters", value)

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> pulumi.Input[builtins.str]:
        return pulumi.get(self, "project_id")

    @project_id.setter
    def project_id(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "project_id", value)


@pulumi.input_type
class _DcsParametersState:
    def __init__(__self__, *,
                 configuration_parameters: Optional[pulumi.Input[Sequence[pulumi.Input['DcsParametersConfigurationParameterArgs']]]] = None,
                 instance_id: Optional[pulumi.Input[builtins.str]] = None,
                 parameters: Optional[pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]]] = None,
                 project_id: Optional[pulumi.Input[builtins.str]] = None):
        """
        Input properties used for looking up and filtering DcsParameters resources.
        """
        if configuration_parameters is not None:
            pulumi.set(__self__, "configuration_parameters", configuration_parameters)
        if instance_id is not None:
            pulumi.set(__self__, "instance_id", instance_id)
        if parameters is not None:
            pulumi.set(__self__, "parameters", parameters)
        if project_id is not None:
            pulumi.set(__self__, "project_id", project_id)

    @property
    @pulumi.getter(name="configurationParameters")
    def configuration_parameters(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['DcsParametersConfigurationParameterArgs']]]]:
        return pulumi.get(self, "configuration_parameters")

    @configuration_parameters.setter
    def configuration_parameters(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['DcsParametersConfigurationParameterArgs']]]]):
        pulumi.set(self, "configuration_parameters", value)

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> Optional[pulumi.Input[builtins.str]]:
        return pulumi.get(self, "instance_id")

    @instance_id.setter
    def instance_id(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "instance_id", value)

    @property
    @pulumi.getter
    def parameters(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]]]:
        return pulumi.get(self, "parameters")

    @parameters.setter
    def parameters(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]]]):
        pulumi.set(self, "parameters", value)

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> Optional[pulumi.Input[builtins.str]]:
        return pulumi.get(self, "project_id")

    @project_id.setter
    def project_id(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "project_id", value)


class DcsParameters(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 instance_id: Optional[pulumi.Input[builtins.str]] = None,
                 parameters: Optional[pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]]] = None,
                 project_id: Optional[pulumi.Input[builtins.str]] = None,
                 __props__=None):
        """
        Create a DcsParameters resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: DcsParametersArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a DcsParameters resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param DcsParametersArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(DcsParametersArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 instance_id: Optional[pulumi.Input[builtins.str]] = None,
                 parameters: Optional[pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]]] = None,
                 project_id: Optional[pulumi.Input[builtins.str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = DcsParametersArgs.__new__(DcsParametersArgs)

            if instance_id is None and not opts.urn:
                raise TypeError("Missing required property 'instance_id'")
            __props__.__dict__["instance_id"] = instance_id
            if parameters is None and not opts.urn:
                raise TypeError("Missing required property 'parameters'")
            __props__.__dict__["parameters"] = parameters
            if project_id is None and not opts.urn:
                raise TypeError("Missing required property 'project_id'")
            __props__.__dict__["project_id"] = project_id
            __props__.__dict__["configuration_parameters"] = None
        super(DcsParameters, __self__).__init__(
            'sbercloud:index/dcsParameters:DcsParameters',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            configuration_parameters: Optional[pulumi.Input[Sequence[pulumi.Input[Union['DcsParametersConfigurationParameterArgs', 'DcsParametersConfigurationParameterArgsDict']]]]] = None,
            instance_id: Optional[pulumi.Input[builtins.str]] = None,
            parameters: Optional[pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]]] = None,
            project_id: Optional[pulumi.Input[builtins.str]] = None) -> 'DcsParameters':
        """
        Get an existing DcsParameters resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _DcsParametersState.__new__(_DcsParametersState)

        __props__.__dict__["configuration_parameters"] = configuration_parameters
        __props__.__dict__["instance_id"] = instance_id
        __props__.__dict__["parameters"] = parameters
        __props__.__dict__["project_id"] = project_id
        return DcsParameters(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="configurationParameters")
    def configuration_parameters(self) -> pulumi.Output[Sequence['outputs.DcsParametersConfigurationParameter']]:
        return pulumi.get(self, "configuration_parameters")

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> pulumi.Output[builtins.str]:
        return pulumi.get(self, "instance_id")

    @property
    @pulumi.getter
    def parameters(self) -> pulumi.Output[Mapping[str, builtins.str]]:
        return pulumi.get(self, "parameters")

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> pulumi.Output[builtins.str]:
        return pulumi.get(self, "project_id")

