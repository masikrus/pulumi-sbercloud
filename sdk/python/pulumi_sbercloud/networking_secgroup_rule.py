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

__all__ = ['NetworkingSecgroupRuleInitArgs', 'NetworkingSecgroupRule']

@pulumi.input_type
class NetworkingSecgroupRuleInitArgs:
    def __init__(__self__, *,
                 direction: pulumi.Input[builtins.str],
                 ethertype: pulumi.Input[builtins.str],
                 security_group_id: pulumi.Input[builtins.str],
                 action: Optional[pulumi.Input[builtins.str]] = None,
                 description: Optional[pulumi.Input[builtins.str]] = None,
                 port_range_max: Optional[pulumi.Input[builtins.int]] = None,
                 port_range_min: Optional[pulumi.Input[builtins.int]] = None,
                 ports: Optional[pulumi.Input[builtins.str]] = None,
                 priority: Optional[pulumi.Input[builtins.int]] = None,
                 protocol: Optional[pulumi.Input[builtins.str]] = None,
                 region: Optional[pulumi.Input[builtins.str]] = None,
                 remote_address_group_id: Optional[pulumi.Input[builtins.str]] = None,
                 remote_group_id: Optional[pulumi.Input[builtins.str]] = None,
                 remote_ip_prefix: Optional[pulumi.Input[builtins.str]] = None):
        """
        The set of arguments for constructing a NetworkingSecgroupRule resource.
        """
        pulumi.set(__self__, "direction", direction)
        pulumi.set(__self__, "ethertype", ethertype)
        pulumi.set(__self__, "security_group_id", security_group_id)
        if action is not None:
            pulumi.set(__self__, "action", action)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if port_range_max is not None:
            pulumi.set(__self__, "port_range_max", port_range_max)
        if port_range_min is not None:
            pulumi.set(__self__, "port_range_min", port_range_min)
        if ports is not None:
            pulumi.set(__self__, "ports", ports)
        if priority is not None:
            pulumi.set(__self__, "priority", priority)
        if protocol is not None:
            pulumi.set(__self__, "protocol", protocol)
        if region is not None:
            pulumi.set(__self__, "region", region)
        if remote_address_group_id is not None:
            pulumi.set(__self__, "remote_address_group_id", remote_address_group_id)
        if remote_group_id is not None:
            pulumi.set(__self__, "remote_group_id", remote_group_id)
        if remote_ip_prefix is not None:
            pulumi.set(__self__, "remote_ip_prefix", remote_ip_prefix)

    @property
    @pulumi.getter
    def direction(self) -> pulumi.Input[builtins.str]:
        return pulumi.get(self, "direction")

    @direction.setter
    def direction(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "direction", value)

    @property
    @pulumi.getter
    def ethertype(self) -> pulumi.Input[builtins.str]:
        return pulumi.get(self, "ethertype")

    @ethertype.setter
    def ethertype(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "ethertype", value)

    @property
    @pulumi.getter(name="securityGroupId")
    def security_group_id(self) -> pulumi.Input[builtins.str]:
        return pulumi.get(self, "security_group_id")

    @security_group_id.setter
    def security_group_id(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "security_group_id", value)

    @property
    @pulumi.getter
    def action(self) -> Optional[pulumi.Input[builtins.str]]:
        return pulumi.get(self, "action")

    @action.setter
    def action(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "action", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[builtins.str]]:
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="portRangeMax")
    def port_range_max(self) -> Optional[pulumi.Input[builtins.int]]:
        return pulumi.get(self, "port_range_max")

    @port_range_max.setter
    def port_range_max(self, value: Optional[pulumi.Input[builtins.int]]):
        pulumi.set(self, "port_range_max", value)

    @property
    @pulumi.getter(name="portRangeMin")
    def port_range_min(self) -> Optional[pulumi.Input[builtins.int]]:
        return pulumi.get(self, "port_range_min")

    @port_range_min.setter
    def port_range_min(self, value: Optional[pulumi.Input[builtins.int]]):
        pulumi.set(self, "port_range_min", value)

    @property
    @pulumi.getter
    def ports(self) -> Optional[pulumi.Input[builtins.str]]:
        return pulumi.get(self, "ports")

    @ports.setter
    def ports(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "ports", value)

    @property
    @pulumi.getter
    def priority(self) -> Optional[pulumi.Input[builtins.int]]:
        return pulumi.get(self, "priority")

    @priority.setter
    def priority(self, value: Optional[pulumi.Input[builtins.int]]):
        pulumi.set(self, "priority", value)

    @property
    @pulumi.getter
    def protocol(self) -> Optional[pulumi.Input[builtins.str]]:
        return pulumi.get(self, "protocol")

    @protocol.setter
    def protocol(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "protocol", value)

    @property
    @pulumi.getter
    def region(self) -> Optional[pulumi.Input[builtins.str]]:
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "region", value)

    @property
    @pulumi.getter(name="remoteAddressGroupId")
    def remote_address_group_id(self) -> Optional[pulumi.Input[builtins.str]]:
        return pulumi.get(self, "remote_address_group_id")

    @remote_address_group_id.setter
    def remote_address_group_id(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "remote_address_group_id", value)

    @property
    @pulumi.getter(name="remoteGroupId")
    def remote_group_id(self) -> Optional[pulumi.Input[builtins.str]]:
        return pulumi.get(self, "remote_group_id")

    @remote_group_id.setter
    def remote_group_id(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "remote_group_id", value)

    @property
    @pulumi.getter(name="remoteIpPrefix")
    def remote_ip_prefix(self) -> Optional[pulumi.Input[builtins.str]]:
        return pulumi.get(self, "remote_ip_prefix")

    @remote_ip_prefix.setter
    def remote_ip_prefix(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "remote_ip_prefix", value)


@pulumi.input_type
class _NetworkingSecgroupRuleState:
    def __init__(__self__, *,
                 action: Optional[pulumi.Input[builtins.str]] = None,
                 description: Optional[pulumi.Input[builtins.str]] = None,
                 direction: Optional[pulumi.Input[builtins.str]] = None,
                 ethertype: Optional[pulumi.Input[builtins.str]] = None,
                 port_range_max: Optional[pulumi.Input[builtins.int]] = None,
                 port_range_min: Optional[pulumi.Input[builtins.int]] = None,
                 ports: Optional[pulumi.Input[builtins.str]] = None,
                 priority: Optional[pulumi.Input[builtins.int]] = None,
                 protocol: Optional[pulumi.Input[builtins.str]] = None,
                 region: Optional[pulumi.Input[builtins.str]] = None,
                 remote_address_group_id: Optional[pulumi.Input[builtins.str]] = None,
                 remote_group_id: Optional[pulumi.Input[builtins.str]] = None,
                 remote_ip_prefix: Optional[pulumi.Input[builtins.str]] = None,
                 security_group_id: Optional[pulumi.Input[builtins.str]] = None):
        """
        Input properties used for looking up and filtering NetworkingSecgroupRule resources.
        """
        if action is not None:
            pulumi.set(__self__, "action", action)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if direction is not None:
            pulumi.set(__self__, "direction", direction)
        if ethertype is not None:
            pulumi.set(__self__, "ethertype", ethertype)
        if port_range_max is not None:
            pulumi.set(__self__, "port_range_max", port_range_max)
        if port_range_min is not None:
            pulumi.set(__self__, "port_range_min", port_range_min)
        if ports is not None:
            pulumi.set(__self__, "ports", ports)
        if priority is not None:
            pulumi.set(__self__, "priority", priority)
        if protocol is not None:
            pulumi.set(__self__, "protocol", protocol)
        if region is not None:
            pulumi.set(__self__, "region", region)
        if remote_address_group_id is not None:
            pulumi.set(__self__, "remote_address_group_id", remote_address_group_id)
        if remote_group_id is not None:
            pulumi.set(__self__, "remote_group_id", remote_group_id)
        if remote_ip_prefix is not None:
            pulumi.set(__self__, "remote_ip_prefix", remote_ip_prefix)
        if security_group_id is not None:
            pulumi.set(__self__, "security_group_id", security_group_id)

    @property
    @pulumi.getter
    def action(self) -> Optional[pulumi.Input[builtins.str]]:
        return pulumi.get(self, "action")

    @action.setter
    def action(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "action", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[builtins.str]]:
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def direction(self) -> Optional[pulumi.Input[builtins.str]]:
        return pulumi.get(self, "direction")

    @direction.setter
    def direction(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "direction", value)

    @property
    @pulumi.getter
    def ethertype(self) -> Optional[pulumi.Input[builtins.str]]:
        return pulumi.get(self, "ethertype")

    @ethertype.setter
    def ethertype(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "ethertype", value)

    @property
    @pulumi.getter(name="portRangeMax")
    def port_range_max(self) -> Optional[pulumi.Input[builtins.int]]:
        return pulumi.get(self, "port_range_max")

    @port_range_max.setter
    def port_range_max(self, value: Optional[pulumi.Input[builtins.int]]):
        pulumi.set(self, "port_range_max", value)

    @property
    @pulumi.getter(name="portRangeMin")
    def port_range_min(self) -> Optional[pulumi.Input[builtins.int]]:
        return pulumi.get(self, "port_range_min")

    @port_range_min.setter
    def port_range_min(self, value: Optional[pulumi.Input[builtins.int]]):
        pulumi.set(self, "port_range_min", value)

    @property
    @pulumi.getter
    def ports(self) -> Optional[pulumi.Input[builtins.str]]:
        return pulumi.get(self, "ports")

    @ports.setter
    def ports(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "ports", value)

    @property
    @pulumi.getter
    def priority(self) -> Optional[pulumi.Input[builtins.int]]:
        return pulumi.get(self, "priority")

    @priority.setter
    def priority(self, value: Optional[pulumi.Input[builtins.int]]):
        pulumi.set(self, "priority", value)

    @property
    @pulumi.getter
    def protocol(self) -> Optional[pulumi.Input[builtins.str]]:
        return pulumi.get(self, "protocol")

    @protocol.setter
    def protocol(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "protocol", value)

    @property
    @pulumi.getter
    def region(self) -> Optional[pulumi.Input[builtins.str]]:
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "region", value)

    @property
    @pulumi.getter(name="remoteAddressGroupId")
    def remote_address_group_id(self) -> Optional[pulumi.Input[builtins.str]]:
        return pulumi.get(self, "remote_address_group_id")

    @remote_address_group_id.setter
    def remote_address_group_id(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "remote_address_group_id", value)

    @property
    @pulumi.getter(name="remoteGroupId")
    def remote_group_id(self) -> Optional[pulumi.Input[builtins.str]]:
        return pulumi.get(self, "remote_group_id")

    @remote_group_id.setter
    def remote_group_id(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "remote_group_id", value)

    @property
    @pulumi.getter(name="remoteIpPrefix")
    def remote_ip_prefix(self) -> Optional[pulumi.Input[builtins.str]]:
        return pulumi.get(self, "remote_ip_prefix")

    @remote_ip_prefix.setter
    def remote_ip_prefix(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "remote_ip_prefix", value)

    @property
    @pulumi.getter(name="securityGroupId")
    def security_group_id(self) -> Optional[pulumi.Input[builtins.str]]:
        return pulumi.get(self, "security_group_id")

    @security_group_id.setter
    def security_group_id(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "security_group_id", value)


class NetworkingSecgroupRule(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 action: Optional[pulumi.Input[builtins.str]] = None,
                 description: Optional[pulumi.Input[builtins.str]] = None,
                 direction: Optional[pulumi.Input[builtins.str]] = None,
                 ethertype: Optional[pulumi.Input[builtins.str]] = None,
                 port_range_max: Optional[pulumi.Input[builtins.int]] = None,
                 port_range_min: Optional[pulumi.Input[builtins.int]] = None,
                 ports: Optional[pulumi.Input[builtins.str]] = None,
                 priority: Optional[pulumi.Input[builtins.int]] = None,
                 protocol: Optional[pulumi.Input[builtins.str]] = None,
                 region: Optional[pulumi.Input[builtins.str]] = None,
                 remote_address_group_id: Optional[pulumi.Input[builtins.str]] = None,
                 remote_group_id: Optional[pulumi.Input[builtins.str]] = None,
                 remote_ip_prefix: Optional[pulumi.Input[builtins.str]] = None,
                 security_group_id: Optional[pulumi.Input[builtins.str]] = None,
                 __props__=None):
        """
        Create a NetworkingSecgroupRule resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: NetworkingSecgroupRuleInitArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a NetworkingSecgroupRule resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param NetworkingSecgroupRuleInitArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(NetworkingSecgroupRuleInitArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 action: Optional[pulumi.Input[builtins.str]] = None,
                 description: Optional[pulumi.Input[builtins.str]] = None,
                 direction: Optional[pulumi.Input[builtins.str]] = None,
                 ethertype: Optional[pulumi.Input[builtins.str]] = None,
                 port_range_max: Optional[pulumi.Input[builtins.int]] = None,
                 port_range_min: Optional[pulumi.Input[builtins.int]] = None,
                 ports: Optional[pulumi.Input[builtins.str]] = None,
                 priority: Optional[pulumi.Input[builtins.int]] = None,
                 protocol: Optional[pulumi.Input[builtins.str]] = None,
                 region: Optional[pulumi.Input[builtins.str]] = None,
                 remote_address_group_id: Optional[pulumi.Input[builtins.str]] = None,
                 remote_group_id: Optional[pulumi.Input[builtins.str]] = None,
                 remote_ip_prefix: Optional[pulumi.Input[builtins.str]] = None,
                 security_group_id: Optional[pulumi.Input[builtins.str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = NetworkingSecgroupRuleInitArgs.__new__(NetworkingSecgroupRuleInitArgs)

            __props__.__dict__["action"] = action
            __props__.__dict__["description"] = description
            if direction is None and not opts.urn:
                raise TypeError("Missing required property 'direction'")
            __props__.__dict__["direction"] = direction
            if ethertype is None and not opts.urn:
                raise TypeError("Missing required property 'ethertype'")
            __props__.__dict__["ethertype"] = ethertype
            __props__.__dict__["port_range_max"] = port_range_max
            __props__.__dict__["port_range_min"] = port_range_min
            __props__.__dict__["ports"] = ports
            __props__.__dict__["priority"] = priority
            __props__.__dict__["protocol"] = protocol
            __props__.__dict__["region"] = region
            __props__.__dict__["remote_address_group_id"] = remote_address_group_id
            __props__.__dict__["remote_group_id"] = remote_group_id
            __props__.__dict__["remote_ip_prefix"] = remote_ip_prefix
            if security_group_id is None and not opts.urn:
                raise TypeError("Missing required property 'security_group_id'")
            __props__.__dict__["security_group_id"] = security_group_id
        super(NetworkingSecgroupRule, __self__).__init__(
            'sbercloud:index/networkingSecgroupRule:NetworkingSecgroupRule',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            action: Optional[pulumi.Input[builtins.str]] = None,
            description: Optional[pulumi.Input[builtins.str]] = None,
            direction: Optional[pulumi.Input[builtins.str]] = None,
            ethertype: Optional[pulumi.Input[builtins.str]] = None,
            port_range_max: Optional[pulumi.Input[builtins.int]] = None,
            port_range_min: Optional[pulumi.Input[builtins.int]] = None,
            ports: Optional[pulumi.Input[builtins.str]] = None,
            priority: Optional[pulumi.Input[builtins.int]] = None,
            protocol: Optional[pulumi.Input[builtins.str]] = None,
            region: Optional[pulumi.Input[builtins.str]] = None,
            remote_address_group_id: Optional[pulumi.Input[builtins.str]] = None,
            remote_group_id: Optional[pulumi.Input[builtins.str]] = None,
            remote_ip_prefix: Optional[pulumi.Input[builtins.str]] = None,
            security_group_id: Optional[pulumi.Input[builtins.str]] = None) -> 'NetworkingSecgroupRule':
        """
        Get an existing NetworkingSecgroupRule resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _NetworkingSecgroupRuleState.__new__(_NetworkingSecgroupRuleState)

        __props__.__dict__["action"] = action
        __props__.__dict__["description"] = description
        __props__.__dict__["direction"] = direction
        __props__.__dict__["ethertype"] = ethertype
        __props__.__dict__["port_range_max"] = port_range_max
        __props__.__dict__["port_range_min"] = port_range_min
        __props__.__dict__["ports"] = ports
        __props__.__dict__["priority"] = priority
        __props__.__dict__["protocol"] = protocol
        __props__.__dict__["region"] = region
        __props__.__dict__["remote_address_group_id"] = remote_address_group_id
        __props__.__dict__["remote_group_id"] = remote_group_id
        __props__.__dict__["remote_ip_prefix"] = remote_ip_prefix
        __props__.__dict__["security_group_id"] = security_group_id
        return NetworkingSecgroupRule(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def action(self) -> pulumi.Output[builtins.str]:
        return pulumi.get(self, "action")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[builtins.str]]:
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def direction(self) -> pulumi.Output[builtins.str]:
        return pulumi.get(self, "direction")

    @property
    @pulumi.getter
    def ethertype(self) -> pulumi.Output[builtins.str]:
        return pulumi.get(self, "ethertype")

    @property
    @pulumi.getter(name="portRangeMax")
    def port_range_max(self) -> pulumi.Output[builtins.int]:
        return pulumi.get(self, "port_range_max")

    @property
    @pulumi.getter(name="portRangeMin")
    def port_range_min(self) -> pulumi.Output[builtins.int]:
        return pulumi.get(self, "port_range_min")

    @property
    @pulumi.getter
    def ports(self) -> pulumi.Output[builtins.str]:
        return pulumi.get(self, "ports")

    @property
    @pulumi.getter
    def priority(self) -> pulumi.Output[builtins.int]:
        return pulumi.get(self, "priority")

    @property
    @pulumi.getter
    def protocol(self) -> pulumi.Output[builtins.str]:
        return pulumi.get(self, "protocol")

    @property
    @pulumi.getter
    def region(self) -> pulumi.Output[builtins.str]:
        return pulumi.get(self, "region")

    @property
    @pulumi.getter(name="remoteAddressGroupId")
    def remote_address_group_id(self) -> pulumi.Output[builtins.str]:
        return pulumi.get(self, "remote_address_group_id")

    @property
    @pulumi.getter(name="remoteGroupId")
    def remote_group_id(self) -> pulumi.Output[builtins.str]:
        return pulumi.get(self, "remote_group_id")

    @property
    @pulumi.getter(name="remoteIpPrefix")
    def remote_ip_prefix(self) -> pulumi.Output[builtins.str]:
        return pulumi.get(self, "remote_ip_prefix")

    @property
    @pulumi.getter(name="securityGroupId")
    def security_group_id(self) -> pulumi.Output[builtins.str]:
        return pulumi.get(self, "security_group_id")

