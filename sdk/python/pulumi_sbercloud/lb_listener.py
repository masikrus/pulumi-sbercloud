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

__all__ = ['LbListenerArgs', 'LbListener']

@pulumi.input_type
class LbListenerArgs:
    def __init__(__self__, *,
                 loadbalancer_id: pulumi.Input[builtins.str],
                 protocol: pulumi.Input[builtins.str],
                 protocol_port: pulumi.Input[builtins.int],
                 admin_state_up: Optional[pulumi.Input[builtins.bool]] = None,
                 connection_limit: Optional[pulumi.Input[builtins.int]] = None,
                 default_pool_id: Optional[pulumi.Input[builtins.str]] = None,
                 default_tls_container_ref: Optional[pulumi.Input[builtins.str]] = None,
                 description: Optional[pulumi.Input[builtins.str]] = None,
                 http2_enable: Optional[pulumi.Input[builtins.bool]] = None,
                 name: Optional[pulumi.Input[builtins.str]] = None,
                 region: Optional[pulumi.Input[builtins.str]] = None,
                 sni_container_refs: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]]] = None,
                 tenant_id: Optional[pulumi.Input[builtins.str]] = None):
        """
        The set of arguments for constructing a LbListener resource.
        """
        pulumi.set(__self__, "loadbalancer_id", loadbalancer_id)
        pulumi.set(__self__, "protocol", protocol)
        pulumi.set(__self__, "protocol_port", protocol_port)
        if admin_state_up is not None:
            pulumi.set(__self__, "admin_state_up", admin_state_up)
        if connection_limit is not None:
            pulumi.set(__self__, "connection_limit", connection_limit)
        if default_pool_id is not None:
            pulumi.set(__self__, "default_pool_id", default_pool_id)
        if default_tls_container_ref is not None:
            pulumi.set(__self__, "default_tls_container_ref", default_tls_container_ref)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if http2_enable is not None:
            pulumi.set(__self__, "http2_enable", http2_enable)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if region is not None:
            pulumi.set(__self__, "region", region)
        if sni_container_refs is not None:
            pulumi.set(__self__, "sni_container_refs", sni_container_refs)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)
        if tenant_id is not None:
            warnings.warn("""tenant_id is deprecated""", DeprecationWarning)
            pulumi.log.warn("""tenant_id is deprecated: tenant_id is deprecated""")
        if tenant_id is not None:
            pulumi.set(__self__, "tenant_id", tenant_id)

    @property
    @pulumi.getter(name="loadbalancerId")
    def loadbalancer_id(self) -> pulumi.Input[builtins.str]:
        return pulumi.get(self, "loadbalancer_id")

    @loadbalancer_id.setter
    def loadbalancer_id(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "loadbalancer_id", value)

    @property
    @pulumi.getter
    def protocol(self) -> pulumi.Input[builtins.str]:
        return pulumi.get(self, "protocol")

    @protocol.setter
    def protocol(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "protocol", value)

    @property
    @pulumi.getter(name="protocolPort")
    def protocol_port(self) -> pulumi.Input[builtins.int]:
        return pulumi.get(self, "protocol_port")

    @protocol_port.setter
    def protocol_port(self, value: pulumi.Input[builtins.int]):
        pulumi.set(self, "protocol_port", value)

    @property
    @pulumi.getter(name="adminStateUp")
    def admin_state_up(self) -> Optional[pulumi.Input[builtins.bool]]:
        return pulumi.get(self, "admin_state_up")

    @admin_state_up.setter
    def admin_state_up(self, value: Optional[pulumi.Input[builtins.bool]]):
        pulumi.set(self, "admin_state_up", value)

    @property
    @pulumi.getter(name="connectionLimit")
    def connection_limit(self) -> Optional[pulumi.Input[builtins.int]]:
        return pulumi.get(self, "connection_limit")

    @connection_limit.setter
    def connection_limit(self, value: Optional[pulumi.Input[builtins.int]]):
        pulumi.set(self, "connection_limit", value)

    @property
    @pulumi.getter(name="defaultPoolId")
    def default_pool_id(self) -> Optional[pulumi.Input[builtins.str]]:
        return pulumi.get(self, "default_pool_id")

    @default_pool_id.setter
    def default_pool_id(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "default_pool_id", value)

    @property
    @pulumi.getter(name="defaultTlsContainerRef")
    def default_tls_container_ref(self) -> Optional[pulumi.Input[builtins.str]]:
        return pulumi.get(self, "default_tls_container_ref")

    @default_tls_container_ref.setter
    def default_tls_container_ref(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "default_tls_container_ref", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[builtins.str]]:
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="http2Enable")
    def http2_enable(self) -> Optional[pulumi.Input[builtins.bool]]:
        return pulumi.get(self, "http2_enable")

    @http2_enable.setter
    def http2_enable(self, value: Optional[pulumi.Input[builtins.bool]]):
        pulumi.set(self, "http2_enable", value)

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
    @pulumi.getter(name="sniContainerRefs")
    def sni_container_refs(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]]:
        return pulumi.get(self, "sni_container_refs")

    @sni_container_refs.setter
    def sni_container_refs(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]]):
        pulumi.set(self, "sni_container_refs", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]]]:
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]]]):
        pulumi.set(self, "tags", value)

    @property
    @pulumi.getter(name="tenantId")
    @_utilities.deprecated("""tenant_id is deprecated""")
    def tenant_id(self) -> Optional[pulumi.Input[builtins.str]]:
        return pulumi.get(self, "tenant_id")

    @tenant_id.setter
    def tenant_id(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "tenant_id", value)


@pulumi.input_type
class _LbListenerState:
    def __init__(__self__, *,
                 admin_state_up: Optional[pulumi.Input[builtins.bool]] = None,
                 connection_limit: Optional[pulumi.Input[builtins.int]] = None,
                 default_pool_id: Optional[pulumi.Input[builtins.str]] = None,
                 default_tls_container_ref: Optional[pulumi.Input[builtins.str]] = None,
                 description: Optional[pulumi.Input[builtins.str]] = None,
                 http2_enable: Optional[pulumi.Input[builtins.bool]] = None,
                 loadbalancer_id: Optional[pulumi.Input[builtins.str]] = None,
                 name: Optional[pulumi.Input[builtins.str]] = None,
                 protocol: Optional[pulumi.Input[builtins.str]] = None,
                 protocol_port: Optional[pulumi.Input[builtins.int]] = None,
                 region: Optional[pulumi.Input[builtins.str]] = None,
                 sni_container_refs: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]]] = None,
                 tenant_id: Optional[pulumi.Input[builtins.str]] = None):
        """
        Input properties used for looking up and filtering LbListener resources.
        """
        if admin_state_up is not None:
            pulumi.set(__self__, "admin_state_up", admin_state_up)
        if connection_limit is not None:
            pulumi.set(__self__, "connection_limit", connection_limit)
        if default_pool_id is not None:
            pulumi.set(__self__, "default_pool_id", default_pool_id)
        if default_tls_container_ref is not None:
            pulumi.set(__self__, "default_tls_container_ref", default_tls_container_ref)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if http2_enable is not None:
            pulumi.set(__self__, "http2_enable", http2_enable)
        if loadbalancer_id is not None:
            pulumi.set(__self__, "loadbalancer_id", loadbalancer_id)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if protocol is not None:
            pulumi.set(__self__, "protocol", protocol)
        if protocol_port is not None:
            pulumi.set(__self__, "protocol_port", protocol_port)
        if region is not None:
            pulumi.set(__self__, "region", region)
        if sni_container_refs is not None:
            pulumi.set(__self__, "sni_container_refs", sni_container_refs)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)
        if tenant_id is not None:
            warnings.warn("""tenant_id is deprecated""", DeprecationWarning)
            pulumi.log.warn("""tenant_id is deprecated: tenant_id is deprecated""")
        if tenant_id is not None:
            pulumi.set(__self__, "tenant_id", tenant_id)

    @property
    @pulumi.getter(name="adminStateUp")
    def admin_state_up(self) -> Optional[pulumi.Input[builtins.bool]]:
        return pulumi.get(self, "admin_state_up")

    @admin_state_up.setter
    def admin_state_up(self, value: Optional[pulumi.Input[builtins.bool]]):
        pulumi.set(self, "admin_state_up", value)

    @property
    @pulumi.getter(name="connectionLimit")
    def connection_limit(self) -> Optional[pulumi.Input[builtins.int]]:
        return pulumi.get(self, "connection_limit")

    @connection_limit.setter
    def connection_limit(self, value: Optional[pulumi.Input[builtins.int]]):
        pulumi.set(self, "connection_limit", value)

    @property
    @pulumi.getter(name="defaultPoolId")
    def default_pool_id(self) -> Optional[pulumi.Input[builtins.str]]:
        return pulumi.get(self, "default_pool_id")

    @default_pool_id.setter
    def default_pool_id(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "default_pool_id", value)

    @property
    @pulumi.getter(name="defaultTlsContainerRef")
    def default_tls_container_ref(self) -> Optional[pulumi.Input[builtins.str]]:
        return pulumi.get(self, "default_tls_container_ref")

    @default_tls_container_ref.setter
    def default_tls_container_ref(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "default_tls_container_ref", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[builtins.str]]:
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="http2Enable")
    def http2_enable(self) -> Optional[pulumi.Input[builtins.bool]]:
        return pulumi.get(self, "http2_enable")

    @http2_enable.setter
    def http2_enable(self, value: Optional[pulumi.Input[builtins.bool]]):
        pulumi.set(self, "http2_enable", value)

    @property
    @pulumi.getter(name="loadbalancerId")
    def loadbalancer_id(self) -> Optional[pulumi.Input[builtins.str]]:
        return pulumi.get(self, "loadbalancer_id")

    @loadbalancer_id.setter
    def loadbalancer_id(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "loadbalancer_id", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[builtins.str]]:
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def protocol(self) -> Optional[pulumi.Input[builtins.str]]:
        return pulumi.get(self, "protocol")

    @protocol.setter
    def protocol(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "protocol", value)

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
    @pulumi.getter(name="sniContainerRefs")
    def sni_container_refs(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]]:
        return pulumi.get(self, "sni_container_refs")

    @sni_container_refs.setter
    def sni_container_refs(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]]):
        pulumi.set(self, "sni_container_refs", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]]]:
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]]]):
        pulumi.set(self, "tags", value)

    @property
    @pulumi.getter(name="tenantId")
    @_utilities.deprecated("""tenant_id is deprecated""")
    def tenant_id(self) -> Optional[pulumi.Input[builtins.str]]:
        return pulumi.get(self, "tenant_id")

    @tenant_id.setter
    def tenant_id(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "tenant_id", value)


class LbListener(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 admin_state_up: Optional[pulumi.Input[builtins.bool]] = None,
                 connection_limit: Optional[pulumi.Input[builtins.int]] = None,
                 default_pool_id: Optional[pulumi.Input[builtins.str]] = None,
                 default_tls_container_ref: Optional[pulumi.Input[builtins.str]] = None,
                 description: Optional[pulumi.Input[builtins.str]] = None,
                 http2_enable: Optional[pulumi.Input[builtins.bool]] = None,
                 loadbalancer_id: Optional[pulumi.Input[builtins.str]] = None,
                 name: Optional[pulumi.Input[builtins.str]] = None,
                 protocol: Optional[pulumi.Input[builtins.str]] = None,
                 protocol_port: Optional[pulumi.Input[builtins.int]] = None,
                 region: Optional[pulumi.Input[builtins.str]] = None,
                 sni_container_refs: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]]] = None,
                 tenant_id: Optional[pulumi.Input[builtins.str]] = None,
                 __props__=None):
        """
        Create a LbListener resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: LbListenerArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a LbListener resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param LbListenerArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(LbListenerArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 admin_state_up: Optional[pulumi.Input[builtins.bool]] = None,
                 connection_limit: Optional[pulumi.Input[builtins.int]] = None,
                 default_pool_id: Optional[pulumi.Input[builtins.str]] = None,
                 default_tls_container_ref: Optional[pulumi.Input[builtins.str]] = None,
                 description: Optional[pulumi.Input[builtins.str]] = None,
                 http2_enable: Optional[pulumi.Input[builtins.bool]] = None,
                 loadbalancer_id: Optional[pulumi.Input[builtins.str]] = None,
                 name: Optional[pulumi.Input[builtins.str]] = None,
                 protocol: Optional[pulumi.Input[builtins.str]] = None,
                 protocol_port: Optional[pulumi.Input[builtins.int]] = None,
                 region: Optional[pulumi.Input[builtins.str]] = None,
                 sni_container_refs: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]]] = None,
                 tenant_id: Optional[pulumi.Input[builtins.str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = LbListenerArgs.__new__(LbListenerArgs)

            __props__.__dict__["admin_state_up"] = admin_state_up
            __props__.__dict__["connection_limit"] = connection_limit
            __props__.__dict__["default_pool_id"] = default_pool_id
            __props__.__dict__["default_tls_container_ref"] = default_tls_container_ref
            __props__.__dict__["description"] = description
            __props__.__dict__["http2_enable"] = http2_enable
            if loadbalancer_id is None and not opts.urn:
                raise TypeError("Missing required property 'loadbalancer_id'")
            __props__.__dict__["loadbalancer_id"] = loadbalancer_id
            __props__.__dict__["name"] = name
            if protocol is None and not opts.urn:
                raise TypeError("Missing required property 'protocol'")
            __props__.__dict__["protocol"] = protocol
            if protocol_port is None and not opts.urn:
                raise TypeError("Missing required property 'protocol_port'")
            __props__.__dict__["protocol_port"] = protocol_port
            __props__.__dict__["region"] = region
            __props__.__dict__["sni_container_refs"] = sni_container_refs
            __props__.__dict__["tags"] = tags
            __props__.__dict__["tenant_id"] = tenant_id
        super(LbListener, __self__).__init__(
            'sbercloud:index/lbListener:LbListener',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            admin_state_up: Optional[pulumi.Input[builtins.bool]] = None,
            connection_limit: Optional[pulumi.Input[builtins.int]] = None,
            default_pool_id: Optional[pulumi.Input[builtins.str]] = None,
            default_tls_container_ref: Optional[pulumi.Input[builtins.str]] = None,
            description: Optional[pulumi.Input[builtins.str]] = None,
            http2_enable: Optional[pulumi.Input[builtins.bool]] = None,
            loadbalancer_id: Optional[pulumi.Input[builtins.str]] = None,
            name: Optional[pulumi.Input[builtins.str]] = None,
            protocol: Optional[pulumi.Input[builtins.str]] = None,
            protocol_port: Optional[pulumi.Input[builtins.int]] = None,
            region: Optional[pulumi.Input[builtins.str]] = None,
            sni_container_refs: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]] = None,
            tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]]] = None,
            tenant_id: Optional[pulumi.Input[builtins.str]] = None) -> 'LbListener':
        """
        Get an existing LbListener resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _LbListenerState.__new__(_LbListenerState)

        __props__.__dict__["admin_state_up"] = admin_state_up
        __props__.__dict__["connection_limit"] = connection_limit
        __props__.__dict__["default_pool_id"] = default_pool_id
        __props__.__dict__["default_tls_container_ref"] = default_tls_container_ref
        __props__.__dict__["description"] = description
        __props__.__dict__["http2_enable"] = http2_enable
        __props__.__dict__["loadbalancer_id"] = loadbalancer_id
        __props__.__dict__["name"] = name
        __props__.__dict__["protocol"] = protocol
        __props__.__dict__["protocol_port"] = protocol_port
        __props__.__dict__["region"] = region
        __props__.__dict__["sni_container_refs"] = sni_container_refs
        __props__.__dict__["tags"] = tags
        __props__.__dict__["tenant_id"] = tenant_id
        return LbListener(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="adminStateUp")
    def admin_state_up(self) -> pulumi.Output[Optional[builtins.bool]]:
        return pulumi.get(self, "admin_state_up")

    @property
    @pulumi.getter(name="connectionLimit")
    def connection_limit(self) -> pulumi.Output[builtins.int]:
        return pulumi.get(self, "connection_limit")

    @property
    @pulumi.getter(name="defaultPoolId")
    def default_pool_id(self) -> pulumi.Output[builtins.str]:
        return pulumi.get(self, "default_pool_id")

    @property
    @pulumi.getter(name="defaultTlsContainerRef")
    def default_tls_container_ref(self) -> pulumi.Output[Optional[builtins.str]]:
        return pulumi.get(self, "default_tls_container_ref")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[builtins.str]]:
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="http2Enable")
    def http2_enable(self) -> pulumi.Output[Optional[builtins.bool]]:
        return pulumi.get(self, "http2_enable")

    @property
    @pulumi.getter(name="loadbalancerId")
    def loadbalancer_id(self) -> pulumi.Output[builtins.str]:
        return pulumi.get(self, "loadbalancer_id")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[builtins.str]:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def protocol(self) -> pulumi.Output[builtins.str]:
        return pulumi.get(self, "protocol")

    @property
    @pulumi.getter(name="protocolPort")
    def protocol_port(self) -> pulumi.Output[builtins.int]:
        return pulumi.get(self, "protocol_port")

    @property
    @pulumi.getter
    def region(self) -> pulumi.Output[builtins.str]:
        return pulumi.get(self, "region")

    @property
    @pulumi.getter(name="sniContainerRefs")
    def sni_container_refs(self) -> pulumi.Output[Optional[Sequence[builtins.str]]]:
        return pulumi.get(self, "sni_container_refs")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Mapping[str, builtins.str]]]:
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="tenantId")
    @_utilities.deprecated("""tenant_id is deprecated""")
    def tenant_id(self) -> pulumi.Output[builtins.str]:
        return pulumi.get(self, "tenant_id")

