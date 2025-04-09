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

__all__ = ['LbMemberArgs', 'LbMember']

@pulumi.input_type
class LbMemberArgs:
    def __init__(__self__, *,
                 address: pulumi.Input[builtins.str],
                 pool_id: pulumi.Input[builtins.str],
                 protocol_port: pulumi.Input[builtins.int],
                 subnet_id: pulumi.Input[builtins.str],
                 admin_state_up: Optional[pulumi.Input[builtins.bool]] = None,
                 name: Optional[pulumi.Input[builtins.str]] = None,
                 region: Optional[pulumi.Input[builtins.str]] = None,
                 tenant_id: Optional[pulumi.Input[builtins.str]] = None,
                 weight: Optional[pulumi.Input[builtins.int]] = None):
        """
        The set of arguments for constructing a LbMember resource.
        :param pulumi.Input[builtins.str] subnet_id: the IPv4 subnet ID of the subnet in which to access the member
        :param pulumi.Input[builtins.bool] admin_state_up: schema: Deprecated
        """
        pulumi.set(__self__, "address", address)
        pulumi.set(__self__, "pool_id", pool_id)
        pulumi.set(__self__, "protocol_port", protocol_port)
        pulumi.set(__self__, "subnet_id", subnet_id)
        if admin_state_up is not None:
            pulumi.set(__self__, "admin_state_up", admin_state_up)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if region is not None:
            pulumi.set(__self__, "region", region)
        if tenant_id is not None:
            warnings.warn("""tenant_id is deprecated""", DeprecationWarning)
            pulumi.log.warn("""tenant_id is deprecated: tenant_id is deprecated""")
        if tenant_id is not None:
            pulumi.set(__self__, "tenant_id", tenant_id)
        if weight is not None:
            pulumi.set(__self__, "weight", weight)

    @property
    @pulumi.getter
    def address(self) -> pulumi.Input[builtins.str]:
        return pulumi.get(self, "address")

    @address.setter
    def address(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "address", value)

    @property
    @pulumi.getter(name="poolId")
    def pool_id(self) -> pulumi.Input[builtins.str]:
        return pulumi.get(self, "pool_id")

    @pool_id.setter
    def pool_id(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "pool_id", value)

    @property
    @pulumi.getter(name="protocolPort")
    def protocol_port(self) -> pulumi.Input[builtins.int]:
        return pulumi.get(self, "protocol_port")

    @protocol_port.setter
    def protocol_port(self, value: pulumi.Input[builtins.int]):
        pulumi.set(self, "protocol_port", value)

    @property
    @pulumi.getter(name="subnetId")
    def subnet_id(self) -> pulumi.Input[builtins.str]:
        """
        the IPv4 subnet ID of the subnet in which to access the member
        """
        return pulumi.get(self, "subnet_id")

    @subnet_id.setter
    def subnet_id(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "subnet_id", value)

    @property
    @pulumi.getter(name="adminStateUp")
    def admin_state_up(self) -> Optional[pulumi.Input[builtins.bool]]:
        """
        schema: Deprecated
        """
        return pulumi.get(self, "admin_state_up")

    @admin_state_up.setter
    def admin_state_up(self, value: Optional[pulumi.Input[builtins.bool]]):
        pulumi.set(self, "admin_state_up", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[builtins.str]]:
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def region(self) -> Optional[pulumi.Input[builtins.str]]:
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "region", value)

    @property
    @pulumi.getter(name="tenantId")
    @_utilities.deprecated("""tenant_id is deprecated""")
    def tenant_id(self) -> Optional[pulumi.Input[builtins.str]]:
        return pulumi.get(self, "tenant_id")

    @tenant_id.setter
    def tenant_id(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "tenant_id", value)

    @property
    @pulumi.getter
    def weight(self) -> Optional[pulumi.Input[builtins.int]]:
        return pulumi.get(self, "weight")

    @weight.setter
    def weight(self, value: Optional[pulumi.Input[builtins.int]]):
        pulumi.set(self, "weight", value)


@pulumi.input_type
class _LbMemberState:
    def __init__(__self__, *,
                 address: Optional[pulumi.Input[builtins.str]] = None,
                 admin_state_up: Optional[pulumi.Input[builtins.bool]] = None,
                 backend_server_status: Optional[pulumi.Input[builtins.bool]] = None,
                 name: Optional[pulumi.Input[builtins.str]] = None,
                 operating_status: Optional[pulumi.Input[builtins.str]] = None,
                 pool_id: Optional[pulumi.Input[builtins.str]] = None,
                 protocol_port: Optional[pulumi.Input[builtins.int]] = None,
                 region: Optional[pulumi.Input[builtins.str]] = None,
                 subnet_id: Optional[pulumi.Input[builtins.str]] = None,
                 tenant_id: Optional[pulumi.Input[builtins.str]] = None,
                 weight: Optional[pulumi.Input[builtins.int]] = None):
        """
        Input properties used for looking up and filtering LbMember resources.
        :param pulumi.Input[builtins.bool] admin_state_up: schema: Deprecated
        :param pulumi.Input[builtins.str] subnet_id: the IPv4 subnet ID of the subnet in which to access the member
        """
        if address is not None:
            pulumi.set(__self__, "address", address)
        if admin_state_up is not None:
            pulumi.set(__self__, "admin_state_up", admin_state_up)
        if backend_server_status is not None:
            pulumi.set(__self__, "backend_server_status", backend_server_status)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if operating_status is not None:
            pulumi.set(__self__, "operating_status", operating_status)
        if pool_id is not None:
            pulumi.set(__self__, "pool_id", pool_id)
        if protocol_port is not None:
            pulumi.set(__self__, "protocol_port", protocol_port)
        if region is not None:
            pulumi.set(__self__, "region", region)
        if subnet_id is not None:
            pulumi.set(__self__, "subnet_id", subnet_id)
        if tenant_id is not None:
            warnings.warn("""tenant_id is deprecated""", DeprecationWarning)
            pulumi.log.warn("""tenant_id is deprecated: tenant_id is deprecated""")
        if tenant_id is not None:
            pulumi.set(__self__, "tenant_id", tenant_id)
        if weight is not None:
            pulumi.set(__self__, "weight", weight)

    @property
    @pulumi.getter
    def address(self) -> Optional[pulumi.Input[builtins.str]]:
        return pulumi.get(self, "address")

    @address.setter
    def address(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "address", value)

    @property
    @pulumi.getter(name="adminStateUp")
    def admin_state_up(self) -> Optional[pulumi.Input[builtins.bool]]:
        """
        schema: Deprecated
        """
        return pulumi.get(self, "admin_state_up")

    @admin_state_up.setter
    def admin_state_up(self, value: Optional[pulumi.Input[builtins.bool]]):
        pulumi.set(self, "admin_state_up", value)

    @property
    @pulumi.getter(name="backendServerStatus")
    def backend_server_status(self) -> Optional[pulumi.Input[builtins.bool]]:
        return pulumi.get(self, "backend_server_status")

    @backend_server_status.setter
    def backend_server_status(self, value: Optional[pulumi.Input[builtins.bool]]):
        pulumi.set(self, "backend_server_status", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[builtins.str]]:
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="operatingStatus")
    def operating_status(self) -> Optional[pulumi.Input[builtins.str]]:
        return pulumi.get(self, "operating_status")

    @operating_status.setter
    def operating_status(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "operating_status", value)

    @property
    @pulumi.getter(name="poolId")
    def pool_id(self) -> Optional[pulumi.Input[builtins.str]]:
        return pulumi.get(self, "pool_id")

    @pool_id.setter
    def pool_id(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "pool_id", value)

    @property
    @pulumi.getter(name="protocolPort")
    def protocol_port(self) -> Optional[pulumi.Input[builtins.int]]:
        return pulumi.get(self, "protocol_port")

    @protocol_port.setter
    def protocol_port(self, value: Optional[pulumi.Input[builtins.int]]):
        pulumi.set(self, "protocol_port", value)

    @property
    @pulumi.getter
    def region(self) -> Optional[pulumi.Input[builtins.str]]:
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "region", value)

    @property
    @pulumi.getter(name="subnetId")
    def subnet_id(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        the IPv4 subnet ID of the subnet in which to access the member
        """
        return pulumi.get(self, "subnet_id")

    @subnet_id.setter
    def subnet_id(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "subnet_id", value)

    @property
    @pulumi.getter(name="tenantId")
    @_utilities.deprecated("""tenant_id is deprecated""")
    def tenant_id(self) -> Optional[pulumi.Input[builtins.str]]:
        return pulumi.get(self, "tenant_id")

    @tenant_id.setter
    def tenant_id(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "tenant_id", value)

    @property
    @pulumi.getter
    def weight(self) -> Optional[pulumi.Input[builtins.int]]:
        return pulumi.get(self, "weight")

    @weight.setter
    def weight(self, value: Optional[pulumi.Input[builtins.int]]):
        pulumi.set(self, "weight", value)


class LbMember(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 address: Optional[pulumi.Input[builtins.str]] = None,
                 admin_state_up: Optional[pulumi.Input[builtins.bool]] = None,
                 name: Optional[pulumi.Input[builtins.str]] = None,
                 pool_id: Optional[pulumi.Input[builtins.str]] = None,
                 protocol_port: Optional[pulumi.Input[builtins.int]] = None,
                 region: Optional[pulumi.Input[builtins.str]] = None,
                 subnet_id: Optional[pulumi.Input[builtins.str]] = None,
                 tenant_id: Optional[pulumi.Input[builtins.str]] = None,
                 weight: Optional[pulumi.Input[builtins.int]] = None,
                 __props__=None):
        """
        Create a LbMember resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[builtins.bool] admin_state_up: schema: Deprecated
        :param pulumi.Input[builtins.str] subnet_id: the IPv4 subnet ID of the subnet in which to access the member
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: LbMemberArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a LbMember resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param LbMemberArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(LbMemberArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 address: Optional[pulumi.Input[builtins.str]] = None,
                 admin_state_up: Optional[pulumi.Input[builtins.bool]] = None,
                 name: Optional[pulumi.Input[builtins.str]] = None,
                 pool_id: Optional[pulumi.Input[builtins.str]] = None,
                 protocol_port: Optional[pulumi.Input[builtins.int]] = None,
                 region: Optional[pulumi.Input[builtins.str]] = None,
                 subnet_id: Optional[pulumi.Input[builtins.str]] = None,
                 tenant_id: Optional[pulumi.Input[builtins.str]] = None,
                 weight: Optional[pulumi.Input[builtins.int]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = LbMemberArgs.__new__(LbMemberArgs)

            if address is None and not opts.urn:
                raise TypeError("Missing required property 'address'")
            __props__.__dict__["address"] = address
            __props__.__dict__["admin_state_up"] = admin_state_up
            __props__.__dict__["name"] = name
            if pool_id is None and not opts.urn:
                raise TypeError("Missing required property 'pool_id'")
            __props__.__dict__["pool_id"] = pool_id
            if protocol_port is None and not opts.urn:
                raise TypeError("Missing required property 'protocol_port'")
            __props__.__dict__["protocol_port"] = protocol_port
            __props__.__dict__["region"] = region
            if subnet_id is None and not opts.urn:
                raise TypeError("Missing required property 'subnet_id'")
            __props__.__dict__["subnet_id"] = subnet_id
            __props__.__dict__["tenant_id"] = tenant_id
            __props__.__dict__["weight"] = weight
            __props__.__dict__["backend_server_status"] = None
            __props__.__dict__["operating_status"] = None
        super(LbMember, __self__).__init__(
            'sbercloud:index/lbMember:LbMember',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            address: Optional[pulumi.Input[builtins.str]] = None,
            admin_state_up: Optional[pulumi.Input[builtins.bool]] = None,
            backend_server_status: Optional[pulumi.Input[builtins.bool]] = None,
            name: Optional[pulumi.Input[builtins.str]] = None,
            operating_status: Optional[pulumi.Input[builtins.str]] = None,
            pool_id: Optional[pulumi.Input[builtins.str]] = None,
            protocol_port: Optional[pulumi.Input[builtins.int]] = None,
            region: Optional[pulumi.Input[builtins.str]] = None,
            subnet_id: Optional[pulumi.Input[builtins.str]] = None,
            tenant_id: Optional[pulumi.Input[builtins.str]] = None,
            weight: Optional[pulumi.Input[builtins.int]] = None) -> 'LbMember':
        """
        Get an existing LbMember resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[builtins.bool] admin_state_up: schema: Deprecated
        :param pulumi.Input[builtins.str] subnet_id: the IPv4 subnet ID of the subnet in which to access the member
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _LbMemberState.__new__(_LbMemberState)

        __props__.__dict__["address"] = address
        __props__.__dict__["admin_state_up"] = admin_state_up
        __props__.__dict__["backend_server_status"] = backend_server_status
        __props__.__dict__["name"] = name
        __props__.__dict__["operating_status"] = operating_status
        __props__.__dict__["pool_id"] = pool_id
        __props__.__dict__["protocol_port"] = protocol_port
        __props__.__dict__["region"] = region
        __props__.__dict__["subnet_id"] = subnet_id
        __props__.__dict__["tenant_id"] = tenant_id
        __props__.__dict__["weight"] = weight
        return LbMember(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def address(self) -> pulumi.Output[builtins.str]:
        return pulumi.get(self, "address")

    @property
    @pulumi.getter(name="adminStateUp")
    def admin_state_up(self) -> pulumi.Output[Optional[builtins.bool]]:
        """
        schema: Deprecated
        """
        return pulumi.get(self, "admin_state_up")

    @property
    @pulumi.getter(name="backendServerStatus")
    def backend_server_status(self) -> pulumi.Output[builtins.bool]:
        return pulumi.get(self, "backend_server_status")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[builtins.str]:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="operatingStatus")
    def operating_status(self) -> pulumi.Output[builtins.str]:
        return pulumi.get(self, "operating_status")

    @property
    @pulumi.getter(name="poolId")
    def pool_id(self) -> pulumi.Output[builtins.str]:
        return pulumi.get(self, "pool_id")

    @property
    @pulumi.getter(name="protocolPort")
    def protocol_port(self) -> pulumi.Output[builtins.int]:
        return pulumi.get(self, "protocol_port")

    @property
    @pulumi.getter
    def region(self) -> pulumi.Output[builtins.str]:
        return pulumi.get(self, "region")

    @property
    @pulumi.getter(name="subnetId")
    def subnet_id(self) -> pulumi.Output[builtins.str]:
        """
        the IPv4 subnet ID of the subnet in which to access the member
        """
        return pulumi.get(self, "subnet_id")

    @property
    @pulumi.getter(name="tenantId")
    @_utilities.deprecated("""tenant_id is deprecated""")
    def tenant_id(self) -> pulumi.Output[builtins.str]:
        return pulumi.get(self, "tenant_id")

    @property
    @pulumi.getter
    def weight(self) -> pulumi.Output[builtins.int]:
        return pulumi.get(self, "weight")

