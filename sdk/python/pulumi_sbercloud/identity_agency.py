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

__all__ = ['IdentityAgencyArgs', 'IdentityAgency']

@pulumi.input_type
class IdentityAgencyArgs:
    def __init__(__self__, *,
                 all_resources_roles: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]] = None,
                 delegated_domain_name: Optional[pulumi.Input[builtins.str]] = None,
                 delegated_service_name: Optional[pulumi.Input[builtins.str]] = None,
                 description: Optional[pulumi.Input[builtins.str]] = None,
                 domain_roles: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]] = None,
                 duration: Optional[pulumi.Input[builtins.str]] = None,
                 name: Optional[pulumi.Input[builtins.str]] = None,
                 project_roles: Optional[pulumi.Input[Sequence[pulumi.Input['IdentityAgencyProjectRoleArgs']]]] = None):
        """
        The set of arguments for constructing a IdentityAgency resource.
        :param pulumi.Input[builtins.str] delegated_domain_name: schema: Required
        :param pulumi.Input[builtins.str] delegated_service_name: schema: Internal
        """
        if all_resources_roles is not None:
            pulumi.set(__self__, "all_resources_roles", all_resources_roles)
        if delegated_domain_name is not None:
            pulumi.set(__self__, "delegated_domain_name", delegated_domain_name)
        if delegated_service_name is not None:
            pulumi.set(__self__, "delegated_service_name", delegated_service_name)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if domain_roles is not None:
            pulumi.set(__self__, "domain_roles", domain_roles)
        if duration is not None:
            pulumi.set(__self__, "duration", duration)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if project_roles is not None:
            pulumi.set(__self__, "project_roles", project_roles)

    @property
    @pulumi.getter(name="allResourcesRoles")
    def all_resources_roles(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]]:
        return pulumi.get(self, "all_resources_roles")

    @all_resources_roles.setter
    def all_resources_roles(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]]):
        pulumi.set(self, "all_resources_roles", value)

    @property
    @pulumi.getter(name="delegatedDomainName")
    def delegated_domain_name(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        schema: Required
        """
        return pulumi.get(self, "delegated_domain_name")

    @delegated_domain_name.setter
    def delegated_domain_name(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "delegated_domain_name", value)

    @property
    @pulumi.getter(name="delegatedServiceName")
    def delegated_service_name(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        schema: Internal
        """
        return pulumi.get(self, "delegated_service_name")

    @delegated_service_name.setter
    def delegated_service_name(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "delegated_service_name", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[builtins.str]]:
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="domainRoles")
    def domain_roles(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]]:
        return pulumi.get(self, "domain_roles")

    @domain_roles.setter
    def domain_roles(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]]):
        pulumi.set(self, "domain_roles", value)

    @property
    @pulumi.getter
    def duration(self) -> Optional[pulumi.Input[builtins.str]]:
        return pulumi.get(self, "duration")

    @duration.setter
    def duration(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "duration", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[builtins.str]]:
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="projectRoles")
    def project_roles(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['IdentityAgencyProjectRoleArgs']]]]:
        return pulumi.get(self, "project_roles")

    @project_roles.setter
    def project_roles(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['IdentityAgencyProjectRoleArgs']]]]):
        pulumi.set(self, "project_roles", value)


@pulumi.input_type
class _IdentityAgencyState:
    def __init__(__self__, *,
                 all_resources_roles: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]] = None,
                 create_time: Optional[pulumi.Input[builtins.str]] = None,
                 delegated_domain_name: Optional[pulumi.Input[builtins.str]] = None,
                 delegated_service_name: Optional[pulumi.Input[builtins.str]] = None,
                 description: Optional[pulumi.Input[builtins.str]] = None,
                 domain_roles: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]] = None,
                 duration: Optional[pulumi.Input[builtins.str]] = None,
                 expire_time: Optional[pulumi.Input[builtins.str]] = None,
                 name: Optional[pulumi.Input[builtins.str]] = None,
                 project_roles: Optional[pulumi.Input[Sequence[pulumi.Input['IdentityAgencyProjectRoleArgs']]]] = None):
        """
        Input properties used for looking up and filtering IdentityAgency resources.
        :param pulumi.Input[builtins.str] delegated_domain_name: schema: Required
        :param pulumi.Input[builtins.str] delegated_service_name: schema: Internal
        """
        if all_resources_roles is not None:
            pulumi.set(__self__, "all_resources_roles", all_resources_roles)
        if create_time is not None:
            pulumi.set(__self__, "create_time", create_time)
        if delegated_domain_name is not None:
            pulumi.set(__self__, "delegated_domain_name", delegated_domain_name)
        if delegated_service_name is not None:
            pulumi.set(__self__, "delegated_service_name", delegated_service_name)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if domain_roles is not None:
            pulumi.set(__self__, "domain_roles", domain_roles)
        if duration is not None:
            pulumi.set(__self__, "duration", duration)
        if expire_time is not None:
            pulumi.set(__self__, "expire_time", expire_time)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if project_roles is not None:
            pulumi.set(__self__, "project_roles", project_roles)

    @property
    @pulumi.getter(name="allResourcesRoles")
    def all_resources_roles(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]]:
        return pulumi.get(self, "all_resources_roles")

    @all_resources_roles.setter
    def all_resources_roles(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]]):
        pulumi.set(self, "all_resources_roles", value)

    @property
    @pulumi.getter(name="createTime")
    def create_time(self) -> Optional[pulumi.Input[builtins.str]]:
        return pulumi.get(self, "create_time")

    @create_time.setter
    def create_time(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "create_time", value)

    @property
    @pulumi.getter(name="delegatedDomainName")
    def delegated_domain_name(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        schema: Required
        """
        return pulumi.get(self, "delegated_domain_name")

    @delegated_domain_name.setter
    def delegated_domain_name(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "delegated_domain_name", value)

    @property
    @pulumi.getter(name="delegatedServiceName")
    def delegated_service_name(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        schema: Internal
        """
        return pulumi.get(self, "delegated_service_name")

    @delegated_service_name.setter
    def delegated_service_name(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "delegated_service_name", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[builtins.str]]:
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="domainRoles")
    def domain_roles(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]]:
        return pulumi.get(self, "domain_roles")

    @domain_roles.setter
    def domain_roles(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]]):
        pulumi.set(self, "domain_roles", value)

    @property
    @pulumi.getter
    def duration(self) -> Optional[pulumi.Input[builtins.str]]:
        return pulumi.get(self, "duration")

    @duration.setter
    def duration(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "duration", value)

    @property
    @pulumi.getter(name="expireTime")
    def expire_time(self) -> Optional[pulumi.Input[builtins.str]]:
        return pulumi.get(self, "expire_time")

    @expire_time.setter
    def expire_time(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "expire_time", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[builtins.str]]:
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="projectRoles")
    def project_roles(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['IdentityAgencyProjectRoleArgs']]]]:
        return pulumi.get(self, "project_roles")

    @project_roles.setter
    def project_roles(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['IdentityAgencyProjectRoleArgs']]]]):
        pulumi.set(self, "project_roles", value)


class IdentityAgency(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 all_resources_roles: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]] = None,
                 delegated_domain_name: Optional[pulumi.Input[builtins.str]] = None,
                 delegated_service_name: Optional[pulumi.Input[builtins.str]] = None,
                 description: Optional[pulumi.Input[builtins.str]] = None,
                 domain_roles: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]] = None,
                 duration: Optional[pulumi.Input[builtins.str]] = None,
                 name: Optional[pulumi.Input[builtins.str]] = None,
                 project_roles: Optional[pulumi.Input[Sequence[pulumi.Input[Union['IdentityAgencyProjectRoleArgs', 'IdentityAgencyProjectRoleArgsDict']]]]] = None,
                 __props__=None):
        """
        Create a IdentityAgency resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[builtins.str] delegated_domain_name: schema: Required
        :param pulumi.Input[builtins.str] delegated_service_name: schema: Internal
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[IdentityAgencyArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a IdentityAgency resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param IdentityAgencyArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(IdentityAgencyArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 all_resources_roles: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]] = None,
                 delegated_domain_name: Optional[pulumi.Input[builtins.str]] = None,
                 delegated_service_name: Optional[pulumi.Input[builtins.str]] = None,
                 description: Optional[pulumi.Input[builtins.str]] = None,
                 domain_roles: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]] = None,
                 duration: Optional[pulumi.Input[builtins.str]] = None,
                 name: Optional[pulumi.Input[builtins.str]] = None,
                 project_roles: Optional[pulumi.Input[Sequence[pulumi.Input[Union['IdentityAgencyProjectRoleArgs', 'IdentityAgencyProjectRoleArgsDict']]]]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = IdentityAgencyArgs.__new__(IdentityAgencyArgs)

            __props__.__dict__["all_resources_roles"] = all_resources_roles
            __props__.__dict__["delegated_domain_name"] = delegated_domain_name
            __props__.__dict__["delegated_service_name"] = delegated_service_name
            __props__.__dict__["description"] = description
            __props__.__dict__["domain_roles"] = domain_roles
            __props__.__dict__["duration"] = duration
            __props__.__dict__["name"] = name
            __props__.__dict__["project_roles"] = project_roles
            __props__.__dict__["create_time"] = None
            __props__.__dict__["expire_time"] = None
        super(IdentityAgency, __self__).__init__(
            'sbercloud:index/identityAgency:IdentityAgency',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            all_resources_roles: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]] = None,
            create_time: Optional[pulumi.Input[builtins.str]] = None,
            delegated_domain_name: Optional[pulumi.Input[builtins.str]] = None,
            delegated_service_name: Optional[pulumi.Input[builtins.str]] = None,
            description: Optional[pulumi.Input[builtins.str]] = None,
            domain_roles: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]] = None,
            duration: Optional[pulumi.Input[builtins.str]] = None,
            expire_time: Optional[pulumi.Input[builtins.str]] = None,
            name: Optional[pulumi.Input[builtins.str]] = None,
            project_roles: Optional[pulumi.Input[Sequence[pulumi.Input[Union['IdentityAgencyProjectRoleArgs', 'IdentityAgencyProjectRoleArgsDict']]]]] = None) -> 'IdentityAgency':
        """
        Get an existing IdentityAgency resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[builtins.str] delegated_domain_name: schema: Required
        :param pulumi.Input[builtins.str] delegated_service_name: schema: Internal
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _IdentityAgencyState.__new__(_IdentityAgencyState)

        __props__.__dict__["all_resources_roles"] = all_resources_roles
        __props__.__dict__["create_time"] = create_time
        __props__.__dict__["delegated_domain_name"] = delegated_domain_name
        __props__.__dict__["delegated_service_name"] = delegated_service_name
        __props__.__dict__["description"] = description
        __props__.__dict__["domain_roles"] = domain_roles
        __props__.__dict__["duration"] = duration
        __props__.__dict__["expire_time"] = expire_time
        __props__.__dict__["name"] = name
        __props__.__dict__["project_roles"] = project_roles
        return IdentityAgency(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="allResourcesRoles")
    def all_resources_roles(self) -> pulumi.Output[Optional[Sequence[builtins.str]]]:
        return pulumi.get(self, "all_resources_roles")

    @property
    @pulumi.getter(name="createTime")
    def create_time(self) -> pulumi.Output[builtins.str]:
        return pulumi.get(self, "create_time")

    @property
    @pulumi.getter(name="delegatedDomainName")
    def delegated_domain_name(self) -> pulumi.Output[Optional[builtins.str]]:
        """
        schema: Required
        """
        return pulumi.get(self, "delegated_domain_name")

    @property
    @pulumi.getter(name="delegatedServiceName")
    def delegated_service_name(self) -> pulumi.Output[Optional[builtins.str]]:
        """
        schema: Internal
        """
        return pulumi.get(self, "delegated_service_name")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[builtins.str]:
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="domainRoles")
    def domain_roles(self) -> pulumi.Output[Optional[Sequence[builtins.str]]]:
        return pulumi.get(self, "domain_roles")

    @property
    @pulumi.getter
    def duration(self) -> pulumi.Output[Optional[builtins.str]]:
        return pulumi.get(self, "duration")

    @property
    @pulumi.getter(name="expireTime")
    def expire_time(self) -> pulumi.Output[builtins.str]:
        return pulumi.get(self, "expire_time")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[builtins.str]:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="projectRoles")
    def project_roles(self) -> pulumi.Output[Optional[Sequence['outputs.IdentityAgencyProjectRole']]]:
        return pulumi.get(self, "project_roles")

