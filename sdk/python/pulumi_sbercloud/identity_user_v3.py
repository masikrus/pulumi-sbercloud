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

__all__ = ['IdentityUserV3Args', 'IdentityUserV3']

@pulumi.input_type
class IdentityUserV3Args:
    def __init__(__self__, *,
                 access_type: Optional[pulumi.Input[builtins.str]] = None,
                 country_code: Optional[pulumi.Input[builtins.str]] = None,
                 description: Optional[pulumi.Input[builtins.str]] = None,
                 email: Optional[pulumi.Input[builtins.str]] = None,
                 enabled: Optional[pulumi.Input[builtins.bool]] = None,
                 external_identity_id: Optional[pulumi.Input[builtins.str]] = None,
                 external_identity_type: Optional[pulumi.Input[builtins.str]] = None,
                 login_protect_verification_method: Optional[pulumi.Input[builtins.str]] = None,
                 name: Optional[pulumi.Input[builtins.str]] = None,
                 password: Optional[pulumi.Input[builtins.str]] = None,
                 phone: Optional[pulumi.Input[builtins.str]] = None,
                 pwd_reset: Optional[pulumi.Input[builtins.bool]] = None):
        """
        The set of arguments for constructing a IdentityUserV3 resource.
        """
        if access_type is not None:
            pulumi.set(__self__, "access_type", access_type)
        if country_code is not None:
            pulumi.set(__self__, "country_code", country_code)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if email is not None:
            pulumi.set(__self__, "email", email)
        if enabled is not None:
            pulumi.set(__self__, "enabled", enabled)
        if external_identity_id is not None:
            pulumi.set(__self__, "external_identity_id", external_identity_id)
        if external_identity_type is not None:
            pulumi.set(__self__, "external_identity_type", external_identity_type)
        if login_protect_verification_method is not None:
            pulumi.set(__self__, "login_protect_verification_method", login_protect_verification_method)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if password is not None:
            pulumi.set(__self__, "password", password)
        if phone is not None:
            pulumi.set(__self__, "phone", phone)
        if pwd_reset is not None:
            pulumi.set(__self__, "pwd_reset", pwd_reset)

    @property
    @pulumi.getter(name="accessType")
    def access_type(self) -> Optional[pulumi.Input[builtins.str]]:
        return pulumi.get(self, "access_type")

    @access_type.setter
    def access_type(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "access_type", value)

    @property
    @pulumi.getter(name="countryCode")
    def country_code(self) -> Optional[pulumi.Input[builtins.str]]:
        return pulumi.get(self, "country_code")

    @country_code.setter
    def country_code(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "country_code", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[builtins.str]]:
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def email(self) -> Optional[pulumi.Input[builtins.str]]:
        return pulumi.get(self, "email")

    @email.setter
    def email(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "email", value)

    @property
    @pulumi.getter
    def enabled(self) -> Optional[pulumi.Input[builtins.bool]]:
        return pulumi.get(self, "enabled")

    @enabled.setter
    def enabled(self, value: Optional[pulumi.Input[builtins.bool]]):
        pulumi.set(self, "enabled", value)

    @property
    @pulumi.getter(name="externalIdentityId")
    def external_identity_id(self) -> Optional[pulumi.Input[builtins.str]]:
        return pulumi.get(self, "external_identity_id")

    @external_identity_id.setter
    def external_identity_id(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "external_identity_id", value)

    @property
    @pulumi.getter(name="externalIdentityType")
    def external_identity_type(self) -> Optional[pulumi.Input[builtins.str]]:
        return pulumi.get(self, "external_identity_type")

    @external_identity_type.setter
    def external_identity_type(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "external_identity_type", value)

    @property
    @pulumi.getter(name="loginProtectVerificationMethod")
    def login_protect_verification_method(self) -> Optional[pulumi.Input[builtins.str]]:
        return pulumi.get(self, "login_protect_verification_method")

    @login_protect_verification_method.setter
    def login_protect_verification_method(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "login_protect_verification_method", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[builtins.str]]:
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def password(self) -> Optional[pulumi.Input[builtins.str]]:
        return pulumi.get(self, "password")

    @password.setter
    def password(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "password", value)

    @property
    @pulumi.getter
    def phone(self) -> Optional[pulumi.Input[builtins.str]]:
        return pulumi.get(self, "phone")

    @phone.setter
    def phone(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "phone", value)

    @property
    @pulumi.getter(name="pwdReset")
    def pwd_reset(self) -> Optional[pulumi.Input[builtins.bool]]:
        return pulumi.get(self, "pwd_reset")

    @pwd_reset.setter
    def pwd_reset(self, value: Optional[pulumi.Input[builtins.bool]]):
        pulumi.set(self, "pwd_reset", value)


@pulumi.input_type
class _IdentityUserV3State:
    def __init__(__self__, *,
                 access_type: Optional[pulumi.Input[builtins.str]] = None,
                 country_code: Optional[pulumi.Input[builtins.str]] = None,
                 create_time: Optional[pulumi.Input[builtins.str]] = None,
                 description: Optional[pulumi.Input[builtins.str]] = None,
                 email: Optional[pulumi.Input[builtins.str]] = None,
                 enabled: Optional[pulumi.Input[builtins.bool]] = None,
                 external_identity_id: Optional[pulumi.Input[builtins.str]] = None,
                 external_identity_type: Optional[pulumi.Input[builtins.str]] = None,
                 last_login: Optional[pulumi.Input[builtins.str]] = None,
                 login_protect_verification_method: Optional[pulumi.Input[builtins.str]] = None,
                 name: Optional[pulumi.Input[builtins.str]] = None,
                 password: Optional[pulumi.Input[builtins.str]] = None,
                 password_strength: Optional[pulumi.Input[builtins.str]] = None,
                 phone: Optional[pulumi.Input[builtins.str]] = None,
                 pwd_reset: Optional[pulumi.Input[builtins.bool]] = None):
        """
        Input properties used for looking up and filtering IdentityUserV3 resources.
        """
        if access_type is not None:
            pulumi.set(__self__, "access_type", access_type)
        if country_code is not None:
            pulumi.set(__self__, "country_code", country_code)
        if create_time is not None:
            pulumi.set(__self__, "create_time", create_time)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if email is not None:
            pulumi.set(__self__, "email", email)
        if enabled is not None:
            pulumi.set(__self__, "enabled", enabled)
        if external_identity_id is not None:
            pulumi.set(__self__, "external_identity_id", external_identity_id)
        if external_identity_type is not None:
            pulumi.set(__self__, "external_identity_type", external_identity_type)
        if last_login is not None:
            pulumi.set(__self__, "last_login", last_login)
        if login_protect_verification_method is not None:
            pulumi.set(__self__, "login_protect_verification_method", login_protect_verification_method)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if password is not None:
            pulumi.set(__self__, "password", password)
        if password_strength is not None:
            pulumi.set(__self__, "password_strength", password_strength)
        if phone is not None:
            pulumi.set(__self__, "phone", phone)
        if pwd_reset is not None:
            pulumi.set(__self__, "pwd_reset", pwd_reset)

    @property
    @pulumi.getter(name="accessType")
    def access_type(self) -> Optional[pulumi.Input[builtins.str]]:
        return pulumi.get(self, "access_type")

    @access_type.setter
    def access_type(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "access_type", value)

    @property
    @pulumi.getter(name="countryCode")
    def country_code(self) -> Optional[pulumi.Input[builtins.str]]:
        return pulumi.get(self, "country_code")

    @country_code.setter
    def country_code(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "country_code", value)

    @property
    @pulumi.getter(name="createTime")
    def create_time(self) -> Optional[pulumi.Input[builtins.str]]:
        return pulumi.get(self, "create_time")

    @create_time.setter
    def create_time(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "create_time", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[builtins.str]]:
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def email(self) -> Optional[pulumi.Input[builtins.str]]:
        return pulumi.get(self, "email")

    @email.setter
    def email(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "email", value)

    @property
    @pulumi.getter
    def enabled(self) -> Optional[pulumi.Input[builtins.bool]]:
        return pulumi.get(self, "enabled")

    @enabled.setter
    def enabled(self, value: Optional[pulumi.Input[builtins.bool]]):
        pulumi.set(self, "enabled", value)

    @property
    @pulumi.getter(name="externalIdentityId")
    def external_identity_id(self) -> Optional[pulumi.Input[builtins.str]]:
        return pulumi.get(self, "external_identity_id")

    @external_identity_id.setter
    def external_identity_id(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "external_identity_id", value)

    @property
    @pulumi.getter(name="externalIdentityType")
    def external_identity_type(self) -> Optional[pulumi.Input[builtins.str]]:
        return pulumi.get(self, "external_identity_type")

    @external_identity_type.setter
    def external_identity_type(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "external_identity_type", value)

    @property
    @pulumi.getter(name="lastLogin")
    def last_login(self) -> Optional[pulumi.Input[builtins.str]]:
        return pulumi.get(self, "last_login")

    @last_login.setter
    def last_login(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "last_login", value)

    @property
    @pulumi.getter(name="loginProtectVerificationMethod")
    def login_protect_verification_method(self) -> Optional[pulumi.Input[builtins.str]]:
        return pulumi.get(self, "login_protect_verification_method")

    @login_protect_verification_method.setter
    def login_protect_verification_method(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "login_protect_verification_method", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[builtins.str]]:
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def password(self) -> Optional[pulumi.Input[builtins.str]]:
        return pulumi.get(self, "password")

    @password.setter
    def password(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "password", value)

    @property
    @pulumi.getter(name="passwordStrength")
    def password_strength(self) -> Optional[pulumi.Input[builtins.str]]:
        return pulumi.get(self, "password_strength")

    @password_strength.setter
    def password_strength(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "password_strength", value)

    @property
    @pulumi.getter
    def phone(self) -> Optional[pulumi.Input[builtins.str]]:
        return pulumi.get(self, "phone")

    @phone.setter
    def phone(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "phone", value)

    @property
    @pulumi.getter(name="pwdReset")
    def pwd_reset(self) -> Optional[pulumi.Input[builtins.bool]]:
        return pulumi.get(self, "pwd_reset")

    @pwd_reset.setter
    def pwd_reset(self, value: Optional[pulumi.Input[builtins.bool]]):
        pulumi.set(self, "pwd_reset", value)


class IdentityUserV3(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 access_type: Optional[pulumi.Input[builtins.str]] = None,
                 country_code: Optional[pulumi.Input[builtins.str]] = None,
                 description: Optional[pulumi.Input[builtins.str]] = None,
                 email: Optional[pulumi.Input[builtins.str]] = None,
                 enabled: Optional[pulumi.Input[builtins.bool]] = None,
                 external_identity_id: Optional[pulumi.Input[builtins.str]] = None,
                 external_identity_type: Optional[pulumi.Input[builtins.str]] = None,
                 login_protect_verification_method: Optional[pulumi.Input[builtins.str]] = None,
                 name: Optional[pulumi.Input[builtins.str]] = None,
                 password: Optional[pulumi.Input[builtins.str]] = None,
                 phone: Optional[pulumi.Input[builtins.str]] = None,
                 pwd_reset: Optional[pulumi.Input[builtins.bool]] = None,
                 __props__=None):
        """
        Create a IdentityUserV3 resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[IdentityUserV3Args] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a IdentityUserV3 resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param IdentityUserV3Args args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(IdentityUserV3Args, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 access_type: Optional[pulumi.Input[builtins.str]] = None,
                 country_code: Optional[pulumi.Input[builtins.str]] = None,
                 description: Optional[pulumi.Input[builtins.str]] = None,
                 email: Optional[pulumi.Input[builtins.str]] = None,
                 enabled: Optional[pulumi.Input[builtins.bool]] = None,
                 external_identity_id: Optional[pulumi.Input[builtins.str]] = None,
                 external_identity_type: Optional[pulumi.Input[builtins.str]] = None,
                 login_protect_verification_method: Optional[pulumi.Input[builtins.str]] = None,
                 name: Optional[pulumi.Input[builtins.str]] = None,
                 password: Optional[pulumi.Input[builtins.str]] = None,
                 phone: Optional[pulumi.Input[builtins.str]] = None,
                 pwd_reset: Optional[pulumi.Input[builtins.bool]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = IdentityUserV3Args.__new__(IdentityUserV3Args)

            __props__.__dict__["access_type"] = access_type
            __props__.__dict__["country_code"] = country_code
            __props__.__dict__["description"] = description
            __props__.__dict__["email"] = email
            __props__.__dict__["enabled"] = enabled
            __props__.__dict__["external_identity_id"] = external_identity_id
            __props__.__dict__["external_identity_type"] = external_identity_type
            __props__.__dict__["login_protect_verification_method"] = login_protect_verification_method
            __props__.__dict__["name"] = name
            __props__.__dict__["password"] = None if password is None else pulumi.Output.secret(password)
            __props__.__dict__["phone"] = phone
            __props__.__dict__["pwd_reset"] = pwd_reset
            __props__.__dict__["create_time"] = None
            __props__.__dict__["last_login"] = None
            __props__.__dict__["password_strength"] = None
        secret_opts = pulumi.ResourceOptions(additional_secret_outputs=["password"])
        opts = pulumi.ResourceOptions.merge(opts, secret_opts)
        super(IdentityUserV3, __self__).__init__(
            'sbercloud:index/identityUserV3:IdentityUserV3',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            access_type: Optional[pulumi.Input[builtins.str]] = None,
            country_code: Optional[pulumi.Input[builtins.str]] = None,
            create_time: Optional[pulumi.Input[builtins.str]] = None,
            description: Optional[pulumi.Input[builtins.str]] = None,
            email: Optional[pulumi.Input[builtins.str]] = None,
            enabled: Optional[pulumi.Input[builtins.bool]] = None,
            external_identity_id: Optional[pulumi.Input[builtins.str]] = None,
            external_identity_type: Optional[pulumi.Input[builtins.str]] = None,
            last_login: Optional[pulumi.Input[builtins.str]] = None,
            login_protect_verification_method: Optional[pulumi.Input[builtins.str]] = None,
            name: Optional[pulumi.Input[builtins.str]] = None,
            password: Optional[pulumi.Input[builtins.str]] = None,
            password_strength: Optional[pulumi.Input[builtins.str]] = None,
            phone: Optional[pulumi.Input[builtins.str]] = None,
            pwd_reset: Optional[pulumi.Input[builtins.bool]] = None) -> 'IdentityUserV3':
        """
        Get an existing IdentityUserV3 resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _IdentityUserV3State.__new__(_IdentityUserV3State)

        __props__.__dict__["access_type"] = access_type
        __props__.__dict__["country_code"] = country_code
        __props__.__dict__["create_time"] = create_time
        __props__.__dict__["description"] = description
        __props__.__dict__["email"] = email
        __props__.__dict__["enabled"] = enabled
        __props__.__dict__["external_identity_id"] = external_identity_id
        __props__.__dict__["external_identity_type"] = external_identity_type
        __props__.__dict__["last_login"] = last_login
        __props__.__dict__["login_protect_verification_method"] = login_protect_verification_method
        __props__.__dict__["name"] = name
        __props__.__dict__["password"] = password
        __props__.__dict__["password_strength"] = password_strength
        __props__.__dict__["phone"] = phone
        __props__.__dict__["pwd_reset"] = pwd_reset
        return IdentityUserV3(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="accessType")
    def access_type(self) -> pulumi.Output[builtins.str]:
        return pulumi.get(self, "access_type")

    @property
    @pulumi.getter(name="countryCode")
    def country_code(self) -> pulumi.Output[Optional[builtins.str]]:
        return pulumi.get(self, "country_code")

    @property
    @pulumi.getter(name="createTime")
    def create_time(self) -> pulumi.Output[builtins.str]:
        return pulumi.get(self, "create_time")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[builtins.str]]:
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def email(self) -> pulumi.Output[Optional[builtins.str]]:
        return pulumi.get(self, "email")

    @property
    @pulumi.getter
    def enabled(self) -> pulumi.Output[Optional[builtins.bool]]:
        return pulumi.get(self, "enabled")

    @property
    @pulumi.getter(name="externalIdentityId")
    def external_identity_id(self) -> pulumi.Output[Optional[builtins.str]]:
        return pulumi.get(self, "external_identity_id")

    @property
    @pulumi.getter(name="externalIdentityType")
    def external_identity_type(self) -> pulumi.Output[builtins.str]:
        return pulumi.get(self, "external_identity_type")

    @property
    @pulumi.getter(name="lastLogin")
    def last_login(self) -> pulumi.Output[builtins.str]:
        return pulumi.get(self, "last_login")

    @property
    @pulumi.getter(name="loginProtectVerificationMethod")
    def login_protect_verification_method(self) -> pulumi.Output[Optional[builtins.str]]:
        return pulumi.get(self, "login_protect_verification_method")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[builtins.str]:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def password(self) -> pulumi.Output[Optional[builtins.str]]:
        return pulumi.get(self, "password")

    @property
    @pulumi.getter(name="passwordStrength")
    def password_strength(self) -> pulumi.Output[builtins.str]:
        return pulumi.get(self, "password_strength")

    @property
    @pulumi.getter
    def phone(self) -> pulumi.Output[Optional[builtins.str]]:
        return pulumi.get(self, "phone")

    @property
    @pulumi.getter(name="pwdReset")
    def pwd_reset(self) -> pulumi.Output[Optional[builtins.bool]]:
        return pulumi.get(self, "pwd_reset")

