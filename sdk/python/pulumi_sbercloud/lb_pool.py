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

__all__ = ['LbPoolArgs', 'LbPool']

@pulumi.input_type
class LbPoolArgs:
    def __init__(__self__, *,
                 lb_method: pulumi.Input[builtins.str],
                 protocol: pulumi.Input[builtins.str],
                 admin_state_up: Optional[pulumi.Input[builtins.bool]] = None,
                 description: Optional[pulumi.Input[builtins.str]] = None,
                 listener_id: Optional[pulumi.Input[builtins.str]] = None,
                 loadbalancer_id: Optional[pulumi.Input[builtins.str]] = None,
                 name: Optional[pulumi.Input[builtins.str]] = None,
                 persistences: Optional[pulumi.Input[Sequence[pulumi.Input['LbPoolPersistenceArgs']]]] = None,
                 protection_reason: Optional[pulumi.Input[builtins.str]] = None,
                 protection_status: Optional[pulumi.Input[builtins.str]] = None,
                 region: Optional[pulumi.Input[builtins.str]] = None,
                 tenant_id: Optional[pulumi.Input[builtins.str]] = None):
        """
        The set of arguments for constructing a LbPool resource.
        """
        pulumi.set(__self__, "lb_method", lb_method)
        pulumi.set(__self__, "protocol", protocol)
        if admin_state_up is not None:
            warnings.warn("""this field is deprecated""", DeprecationWarning)
            pulumi.log.warn("""admin_state_up is deprecated: this field is deprecated""")
        if admin_state_up is not None:
            pulumi.set(__self__, "admin_state_up", admin_state_up)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if listener_id is not None:
            pulumi.set(__self__, "listener_id", listener_id)
        if loadbalancer_id is not None:
            pulumi.set(__self__, "loadbalancer_id", loadbalancer_id)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if persistences is not None:
            pulumi.set(__self__, "persistences", persistences)
        if protection_reason is not None:
            pulumi.set(__self__, "protection_reason", protection_reason)
        if protection_status is not None:
            pulumi.set(__self__, "protection_status", protection_status)
        if region is not None:
            pulumi.set(__self__, "region", region)
        if tenant_id is not None:
            warnings.warn("""tenant_id is deprecated""", DeprecationWarning)
            pulumi.log.warn("""tenant_id is deprecated: tenant_id is deprecated""")
        if tenant_id is not None:
            pulumi.set(__self__, "tenant_id", tenant_id)

    @property
    @pulumi.getter(name="lbMethod")
    def lb_method(self) -> pulumi.Input[builtins.str]:
        return pulumi.get(self, "lb_method")

    @lb_method.setter
    def lb_method(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "lb_method", value)

    @property
    @pulumi.getter
    def protocol(self) -> pulumi.Input[builtins.str]:
        return pulumi.get(self, "protocol")

    @protocol.setter
    def protocol(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "protocol", value)

    @property
    @pulumi.getter(name="adminStateUp")
    @_utilities.deprecated("""this field is deprecated""")
    def admin_state_up(self) -> Optional[pulumi.Input[builtins.bool]]:
        return pulumi.get(self, "admin_state_up")

    @admin_state_up.setter
    def admin_state_up(self, value: Optional[pulumi.Input[builtins.bool]]):
        pulumi.set(self, "admin_state_up", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[builtins.str]]:
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="listenerId")
    def listener_id(self) -> Optional[pulumi.Input[builtins.str]]:
        return pulumi.get(self, "listener_id")

    @listener_id.setter
    def listener_id(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "listener_id", value)

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
    def persistences(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['LbPoolPersistenceArgs']]]]:
        return pulumi.get(self, "persistences")

    @persistences.setter
    def persistences(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['LbPoolPersistenceArgs']]]]):
        pulumi.set(self, "persistences", value)

    @property
    @pulumi.getter(name="protectionReason")
    def protection_reason(self) -> Optional[pulumi.Input[builtins.str]]:
        return pulumi.get(self, "protection_reason")

    @protection_reason.setter
    def protection_reason(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "protection_reason", value)

    @property
    @pulumi.getter(name="protectionStatus")
    def protection_status(self) -> Optional[pulumi.Input[builtins.str]]:
        return pulumi.get(self, "protection_status")

    @protection_status.setter
    def protection_status(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "protection_status", value)

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


@pulumi.input_type
class _LbPoolState:
    def __init__(__self__, *,
                 admin_state_up: Optional[pulumi.Input[builtins.bool]] = None,
                 description: Optional[pulumi.Input[builtins.str]] = None,
                 lb_method: Optional[pulumi.Input[builtins.str]] = None,
                 listener_id: Optional[pulumi.Input[builtins.str]] = None,
                 loadbalancer_id: Optional[pulumi.Input[builtins.str]] = None,
                 monitor_id: Optional[pulumi.Input[builtins.str]] = None,
                 name: Optional[pulumi.Input[builtins.str]] = None,
                 persistences: Optional[pulumi.Input[Sequence[pulumi.Input['LbPoolPersistenceArgs']]]] = None,
                 protection_reason: Optional[pulumi.Input[builtins.str]] = None,
                 protection_status: Optional[pulumi.Input[builtins.str]] = None,
                 protocol: Optional[pulumi.Input[builtins.str]] = None,
                 region: Optional[pulumi.Input[builtins.str]] = None,
                 tenant_id: Optional[pulumi.Input[builtins.str]] = None):
        """
        Input properties used for looking up and filtering LbPool resources.
        """
        if admin_state_up is not None:
            warnings.warn("""this field is deprecated""", DeprecationWarning)
            pulumi.log.warn("""admin_state_up is deprecated: this field is deprecated""")
        if admin_state_up is not None:
            pulumi.set(__self__, "admin_state_up", admin_state_up)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if lb_method is not None:
            pulumi.set(__self__, "lb_method", lb_method)
        if listener_id is not None:
            pulumi.set(__self__, "listener_id", listener_id)
        if loadbalancer_id is not None:
            pulumi.set(__self__, "loadbalancer_id", loadbalancer_id)
        if monitor_id is not None:
            pulumi.set(__self__, "monitor_id", monitor_id)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if persistences is not None:
            pulumi.set(__self__, "persistences", persistences)
        if protection_reason is not None:
            pulumi.set(__self__, "protection_reason", protection_reason)
        if protection_status is not None:
            pulumi.set(__self__, "protection_status", protection_status)
        if protocol is not None:
            pulumi.set(__self__, "protocol", protocol)
        if region is not None:
            pulumi.set(__self__, "region", region)
        if tenant_id is not None:
            warnings.warn("""tenant_id is deprecated""", DeprecationWarning)
            pulumi.log.warn("""tenant_id is deprecated: tenant_id is deprecated""")
        if tenant_id is not None:
            pulumi.set(__self__, "tenant_id", tenant_id)

    @property
    @pulumi.getter(name="adminStateUp")
    @_utilities.deprecated("""this field is deprecated""")
    def admin_state_up(self) -> Optional[pulumi.Input[builtins.bool]]:
        return pulumi.get(self, "admin_state_up")

    @admin_state_up.setter
    def admin_state_up(self, value: Optional[pulumi.Input[builtins.bool]]):
        pulumi.set(self, "admin_state_up", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[builtins.str]]:
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="lbMethod")
    def lb_method(self) -> Optional[pulumi.Input[builtins.str]]:
        return pulumi.get(self, "lb_method")

    @lb_method.setter
    def lb_method(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "lb_method", value)

    @property
    @pulumi.getter(name="listenerId")
    def listener_id(self) -> Optional[pulumi.Input[builtins.str]]:
        return pulumi.get(self, "listener_id")

    @listener_id.setter
    def listener_id(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "listener_id", value)

    @property
    @pulumi.getter(name="loadbalancerId")
    def loadbalancer_id(self) -> Optional[pulumi.Input[builtins.str]]:
        return pulumi.get(self, "loadbalancer_id")

    @loadbalancer_id.setter
    def loadbalancer_id(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "loadbalancer_id", value)

    @property
    @pulumi.getter(name="monitorId")
    def monitor_id(self) -> Optional[pulumi.Input[builtins.str]]:
        return pulumi.get(self, "monitor_id")

    @monitor_id.setter
    def monitor_id(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "monitor_id", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[builtins.str]]:
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def persistences(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['LbPoolPersistenceArgs']]]]:
        return pulumi.get(self, "persistences")

    @persistences.setter
    def persistences(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['LbPoolPersistenceArgs']]]]):
        pulumi.set(self, "persistences", value)

    @property
    @pulumi.getter(name="protectionReason")
    def protection_reason(self) -> Optional[pulumi.Input[builtins.str]]:
        return pulumi.get(self, "protection_reason")

    @protection_reason.setter
    def protection_reason(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "protection_reason", value)

    @property
    @pulumi.getter(name="protectionStatus")
    def protection_status(self) -> Optional[pulumi.Input[builtins.str]]:
        return pulumi.get(self, "protection_status")

    @protection_status.setter
    def protection_status(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "protection_status", value)

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
    @pulumi.getter(name="tenantId")
    @_utilities.deprecated("""tenant_id is deprecated""")
    def tenant_id(self) -> Optional[pulumi.Input[builtins.str]]:
        return pulumi.get(self, "tenant_id")

    @tenant_id.setter
    def tenant_id(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "tenant_id", value)


class LbPool(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 admin_state_up: Optional[pulumi.Input[builtins.bool]] = None,
                 description: Optional[pulumi.Input[builtins.str]] = None,
                 lb_method: Optional[pulumi.Input[builtins.str]] = None,
                 listener_id: Optional[pulumi.Input[builtins.str]] = None,
                 loadbalancer_id: Optional[pulumi.Input[builtins.str]] = None,
                 name: Optional[pulumi.Input[builtins.str]] = None,
                 persistences: Optional[pulumi.Input[Sequence[pulumi.Input[Union['LbPoolPersistenceArgs', 'LbPoolPersistenceArgsDict']]]]] = None,
                 protection_reason: Optional[pulumi.Input[builtins.str]] = None,
                 protection_status: Optional[pulumi.Input[builtins.str]] = None,
                 protocol: Optional[pulumi.Input[builtins.str]] = None,
                 region: Optional[pulumi.Input[builtins.str]] = None,
                 tenant_id: Optional[pulumi.Input[builtins.str]] = None,
                 __props__=None):
        """
        Create a LbPool resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: LbPoolArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a LbPool resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param LbPoolArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(LbPoolArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 admin_state_up: Optional[pulumi.Input[builtins.bool]] = None,
                 description: Optional[pulumi.Input[builtins.str]] = None,
                 lb_method: Optional[pulumi.Input[builtins.str]] = None,
                 listener_id: Optional[pulumi.Input[builtins.str]] = None,
                 loadbalancer_id: Optional[pulumi.Input[builtins.str]] = None,
                 name: Optional[pulumi.Input[builtins.str]] = None,
                 persistences: Optional[pulumi.Input[Sequence[pulumi.Input[Union['LbPoolPersistenceArgs', 'LbPoolPersistenceArgsDict']]]]] = None,
                 protection_reason: Optional[pulumi.Input[builtins.str]] = None,
                 protection_status: Optional[pulumi.Input[builtins.str]] = None,
                 protocol: Optional[pulumi.Input[builtins.str]] = None,
                 region: Optional[pulumi.Input[builtins.str]] = None,
                 tenant_id: Optional[pulumi.Input[builtins.str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = LbPoolArgs.__new__(LbPoolArgs)

            __props__.__dict__["admin_state_up"] = admin_state_up
            __props__.__dict__["description"] = description
            if lb_method is None and not opts.urn:
                raise TypeError("Missing required property 'lb_method'")
            __props__.__dict__["lb_method"] = lb_method
            __props__.__dict__["listener_id"] = listener_id
            __props__.__dict__["loadbalancer_id"] = loadbalancer_id
            __props__.__dict__["name"] = name
            __props__.__dict__["persistences"] = persistences
            __props__.__dict__["protection_reason"] = protection_reason
            __props__.__dict__["protection_status"] = protection_status
            if protocol is None and not opts.urn:
                raise TypeError("Missing required property 'protocol'")
            __props__.__dict__["protocol"] = protocol
            __props__.__dict__["region"] = region
            __props__.__dict__["tenant_id"] = tenant_id
            __props__.__dict__["monitor_id"] = None
        super(LbPool, __self__).__init__(
            'sbercloud:index/lbPool:LbPool',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            admin_state_up: Optional[pulumi.Input[builtins.bool]] = None,
            description: Optional[pulumi.Input[builtins.str]] = None,
            lb_method: Optional[pulumi.Input[builtins.str]] = None,
            listener_id: Optional[pulumi.Input[builtins.str]] = None,
            loadbalancer_id: Optional[pulumi.Input[builtins.str]] = None,
            monitor_id: Optional[pulumi.Input[builtins.str]] = None,
            name: Optional[pulumi.Input[builtins.str]] = None,
            persistences: Optional[pulumi.Input[Sequence[pulumi.Input[Union['LbPoolPersistenceArgs', 'LbPoolPersistenceArgsDict']]]]] = None,
            protection_reason: Optional[pulumi.Input[builtins.str]] = None,
            protection_status: Optional[pulumi.Input[builtins.str]] = None,
            protocol: Optional[pulumi.Input[builtins.str]] = None,
            region: Optional[pulumi.Input[builtins.str]] = None,
            tenant_id: Optional[pulumi.Input[builtins.str]] = None) -> 'LbPool':
        """
        Get an existing LbPool resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _LbPoolState.__new__(_LbPoolState)

        __props__.__dict__["admin_state_up"] = admin_state_up
        __props__.__dict__["description"] = description
        __props__.__dict__["lb_method"] = lb_method
        __props__.__dict__["listener_id"] = listener_id
        __props__.__dict__["loadbalancer_id"] = loadbalancer_id
        __props__.__dict__["monitor_id"] = monitor_id
        __props__.__dict__["name"] = name
        __props__.__dict__["persistences"] = persistences
        __props__.__dict__["protection_reason"] = protection_reason
        __props__.__dict__["protection_status"] = protection_status
        __props__.__dict__["protocol"] = protocol
        __props__.__dict__["region"] = region
        __props__.__dict__["tenant_id"] = tenant_id
        return LbPool(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="adminStateUp")
    @_utilities.deprecated("""this field is deprecated""")
    def admin_state_up(self) -> pulumi.Output[Optional[builtins.bool]]:
        return pulumi.get(self, "admin_state_up")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[builtins.str]]:
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="lbMethod")
    def lb_method(self) -> pulumi.Output[builtins.str]:
        return pulumi.get(self, "lb_method")

    @property
    @pulumi.getter(name="listenerId")
    def listener_id(self) -> pulumi.Output[builtins.str]:
        return pulumi.get(self, "listener_id")

    @property
    @pulumi.getter(name="loadbalancerId")
    def loadbalancer_id(self) -> pulumi.Output[builtins.str]:
        return pulumi.get(self, "loadbalancer_id")

    @property
    @pulumi.getter(name="monitorId")
    def monitor_id(self) -> pulumi.Output[builtins.str]:
        return pulumi.get(self, "monitor_id")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[builtins.str]:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def persistences(self) -> pulumi.Output[Optional[Sequence['outputs.LbPoolPersistence']]]:
        return pulumi.get(self, "persistences")

    @property
    @pulumi.getter(name="protectionReason")
    def protection_reason(self) -> pulumi.Output[Optional[builtins.str]]:
        return pulumi.get(self, "protection_reason")

    @property
    @pulumi.getter(name="protectionStatus")
    def protection_status(self) -> pulumi.Output[builtins.str]:
        return pulumi.get(self, "protection_status")

    @property
    @pulumi.getter
    def protocol(self) -> pulumi.Output[builtins.str]:
        return pulumi.get(self, "protocol")

    @property
    @pulumi.getter
    def region(self) -> pulumi.Output[builtins.str]:
        return pulumi.get(self, "region")

    @property
    @pulumi.getter(name="tenantId")
    @_utilities.deprecated("""tenant_id is deprecated""")
    def tenant_id(self) -> pulumi.Output[builtins.str]:
        return pulumi.get(self, "tenant_id")

