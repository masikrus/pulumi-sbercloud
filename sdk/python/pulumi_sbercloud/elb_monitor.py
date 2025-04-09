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

__all__ = ['ElbMonitorArgs', 'ElbMonitor']

@pulumi.input_type
class ElbMonitorArgs:
    def __init__(__self__, *,
                 interval: pulumi.Input[builtins.int],
                 max_retries: pulumi.Input[builtins.int],
                 pool_id: pulumi.Input[builtins.str],
                 protocol: pulumi.Input[builtins.str],
                 timeout: pulumi.Input[builtins.int],
                 domain_name: Optional[pulumi.Input[builtins.str]] = None,
                 port: Optional[pulumi.Input[builtins.int]] = None,
                 region: Optional[pulumi.Input[builtins.str]] = None,
                 status_code: Optional[pulumi.Input[builtins.str]] = None,
                 url_path: Optional[pulumi.Input[builtins.str]] = None):
        """
        The set of arguments for constructing a ElbMonitor resource.
        """
        pulumi.set(__self__, "interval", interval)
        pulumi.set(__self__, "max_retries", max_retries)
        pulumi.set(__self__, "pool_id", pool_id)
        pulumi.set(__self__, "protocol", protocol)
        pulumi.set(__self__, "timeout", timeout)
        if domain_name is not None:
            pulumi.set(__self__, "domain_name", domain_name)
        if port is not None:
            pulumi.set(__self__, "port", port)
        if region is not None:
            pulumi.set(__self__, "region", region)
        if status_code is not None:
            pulumi.set(__self__, "status_code", status_code)
        if url_path is not None:
            pulumi.set(__self__, "url_path", url_path)

    @property
    @pulumi.getter
    def interval(self) -> pulumi.Input[builtins.int]:
        return pulumi.get(self, "interval")

    @interval.setter
    def interval(self, value: pulumi.Input[builtins.int]):
        pulumi.set(self, "interval", value)

    @property
    @pulumi.getter(name="maxRetries")
    def max_retries(self) -> pulumi.Input[builtins.int]:
        return pulumi.get(self, "max_retries")

    @max_retries.setter
    def max_retries(self, value: pulumi.Input[builtins.int]):
        pulumi.set(self, "max_retries", value)

    @property
    @pulumi.getter(name="poolId")
    def pool_id(self) -> pulumi.Input[builtins.str]:
        return pulumi.get(self, "pool_id")

    @pool_id.setter
    def pool_id(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "pool_id", value)

    @property
    @pulumi.getter
    def protocol(self) -> pulumi.Input[builtins.str]:
        return pulumi.get(self, "protocol")

    @protocol.setter
    def protocol(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "protocol", value)

    @property
    @pulumi.getter
    def timeout(self) -> pulumi.Input[builtins.int]:
        return pulumi.get(self, "timeout")

    @timeout.setter
    def timeout(self, value: pulumi.Input[builtins.int]):
        pulumi.set(self, "timeout", value)

    @property
    @pulumi.getter(name="domainName")
    def domain_name(self) -> Optional[pulumi.Input[builtins.str]]:
        return pulumi.get(self, "domain_name")

    @domain_name.setter
    def domain_name(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "domain_name", value)

    @property
    @pulumi.getter
    def port(self) -> Optional[pulumi.Input[builtins.int]]:
        return pulumi.get(self, "port")

    @port.setter
    def port(self, value: Optional[pulumi.Input[builtins.int]]):
        pulumi.set(self, "port", value)

    @property
    @pulumi.getter
    def region(self) -> Optional[pulumi.Input[builtins.str]]:
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "region", value)

    @property
    @pulumi.getter(name="statusCode")
    def status_code(self) -> Optional[pulumi.Input[builtins.str]]:
        return pulumi.get(self, "status_code")

    @status_code.setter
    def status_code(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "status_code", value)

    @property
    @pulumi.getter(name="urlPath")
    def url_path(self) -> Optional[pulumi.Input[builtins.str]]:
        return pulumi.get(self, "url_path")

    @url_path.setter
    def url_path(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "url_path", value)


@pulumi.input_type
class _ElbMonitorState:
    def __init__(__self__, *,
                 domain_name: Optional[pulumi.Input[builtins.str]] = None,
                 interval: Optional[pulumi.Input[builtins.int]] = None,
                 max_retries: Optional[pulumi.Input[builtins.int]] = None,
                 pool_id: Optional[pulumi.Input[builtins.str]] = None,
                 port: Optional[pulumi.Input[builtins.int]] = None,
                 protocol: Optional[pulumi.Input[builtins.str]] = None,
                 region: Optional[pulumi.Input[builtins.str]] = None,
                 status_code: Optional[pulumi.Input[builtins.str]] = None,
                 timeout: Optional[pulumi.Input[builtins.int]] = None,
                 url_path: Optional[pulumi.Input[builtins.str]] = None):
        """
        Input properties used for looking up and filtering ElbMonitor resources.
        """
        if domain_name is not None:
            pulumi.set(__self__, "domain_name", domain_name)
        if interval is not None:
            pulumi.set(__self__, "interval", interval)
        if max_retries is not None:
            pulumi.set(__self__, "max_retries", max_retries)
        if pool_id is not None:
            pulumi.set(__self__, "pool_id", pool_id)
        if port is not None:
            pulumi.set(__self__, "port", port)
        if protocol is not None:
            pulumi.set(__self__, "protocol", protocol)
        if region is not None:
            pulumi.set(__self__, "region", region)
        if status_code is not None:
            pulumi.set(__self__, "status_code", status_code)
        if timeout is not None:
            pulumi.set(__self__, "timeout", timeout)
        if url_path is not None:
            pulumi.set(__self__, "url_path", url_path)

    @property
    @pulumi.getter(name="domainName")
    def domain_name(self) -> Optional[pulumi.Input[builtins.str]]:
        return pulumi.get(self, "domain_name")

    @domain_name.setter
    def domain_name(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "domain_name", value)

    @property
    @pulumi.getter
    def interval(self) -> Optional[pulumi.Input[builtins.int]]:
        return pulumi.get(self, "interval")

    @interval.setter
    def interval(self, value: Optional[pulumi.Input[builtins.int]]):
        pulumi.set(self, "interval", value)

    @property
    @pulumi.getter(name="maxRetries")
    def max_retries(self) -> Optional[pulumi.Input[builtins.int]]:
        return pulumi.get(self, "max_retries")

    @max_retries.setter
    def max_retries(self, value: Optional[pulumi.Input[builtins.int]]):
        pulumi.set(self, "max_retries", value)

    @property
    @pulumi.getter(name="poolId")
    def pool_id(self) -> Optional[pulumi.Input[builtins.str]]:
        return pulumi.get(self, "pool_id")

    @pool_id.setter
    def pool_id(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "pool_id", value)

    @property
    @pulumi.getter
    def port(self) -> Optional[pulumi.Input[builtins.int]]:
        return pulumi.get(self, "port")

    @port.setter
    def port(self, value: Optional[pulumi.Input[builtins.int]]):
        pulumi.set(self, "port", value)

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
    @pulumi.getter(name="statusCode")
    def status_code(self) -> Optional[pulumi.Input[builtins.str]]:
        return pulumi.get(self, "status_code")

    @status_code.setter
    def status_code(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "status_code", value)

    @property
    @pulumi.getter
    def timeout(self) -> Optional[pulumi.Input[builtins.int]]:
        return pulumi.get(self, "timeout")

    @timeout.setter
    def timeout(self, value: Optional[pulumi.Input[builtins.int]]):
        pulumi.set(self, "timeout", value)

    @property
    @pulumi.getter(name="urlPath")
    def url_path(self) -> Optional[pulumi.Input[builtins.str]]:
        return pulumi.get(self, "url_path")

    @url_path.setter
    def url_path(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "url_path", value)


class ElbMonitor(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 domain_name: Optional[pulumi.Input[builtins.str]] = None,
                 interval: Optional[pulumi.Input[builtins.int]] = None,
                 max_retries: Optional[pulumi.Input[builtins.int]] = None,
                 pool_id: Optional[pulumi.Input[builtins.str]] = None,
                 port: Optional[pulumi.Input[builtins.int]] = None,
                 protocol: Optional[pulumi.Input[builtins.str]] = None,
                 region: Optional[pulumi.Input[builtins.str]] = None,
                 status_code: Optional[pulumi.Input[builtins.str]] = None,
                 timeout: Optional[pulumi.Input[builtins.int]] = None,
                 url_path: Optional[pulumi.Input[builtins.str]] = None,
                 __props__=None):
        """
        Create a ElbMonitor resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ElbMonitorArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a ElbMonitor resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param ElbMonitorArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ElbMonitorArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 domain_name: Optional[pulumi.Input[builtins.str]] = None,
                 interval: Optional[pulumi.Input[builtins.int]] = None,
                 max_retries: Optional[pulumi.Input[builtins.int]] = None,
                 pool_id: Optional[pulumi.Input[builtins.str]] = None,
                 port: Optional[pulumi.Input[builtins.int]] = None,
                 protocol: Optional[pulumi.Input[builtins.str]] = None,
                 region: Optional[pulumi.Input[builtins.str]] = None,
                 status_code: Optional[pulumi.Input[builtins.str]] = None,
                 timeout: Optional[pulumi.Input[builtins.int]] = None,
                 url_path: Optional[pulumi.Input[builtins.str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ElbMonitorArgs.__new__(ElbMonitorArgs)

            __props__.__dict__["domain_name"] = domain_name
            if interval is None and not opts.urn:
                raise TypeError("Missing required property 'interval'")
            __props__.__dict__["interval"] = interval
            if max_retries is None and not opts.urn:
                raise TypeError("Missing required property 'max_retries'")
            __props__.__dict__["max_retries"] = max_retries
            if pool_id is None and not opts.urn:
                raise TypeError("Missing required property 'pool_id'")
            __props__.__dict__["pool_id"] = pool_id
            __props__.__dict__["port"] = port
            if protocol is None and not opts.urn:
                raise TypeError("Missing required property 'protocol'")
            __props__.__dict__["protocol"] = protocol
            __props__.__dict__["region"] = region
            __props__.__dict__["status_code"] = status_code
            if timeout is None and not opts.urn:
                raise TypeError("Missing required property 'timeout'")
            __props__.__dict__["timeout"] = timeout
            __props__.__dict__["url_path"] = url_path
        super(ElbMonitor, __self__).__init__(
            'sbercloud:index/elbMonitor:ElbMonitor',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            domain_name: Optional[pulumi.Input[builtins.str]] = None,
            interval: Optional[pulumi.Input[builtins.int]] = None,
            max_retries: Optional[pulumi.Input[builtins.int]] = None,
            pool_id: Optional[pulumi.Input[builtins.str]] = None,
            port: Optional[pulumi.Input[builtins.int]] = None,
            protocol: Optional[pulumi.Input[builtins.str]] = None,
            region: Optional[pulumi.Input[builtins.str]] = None,
            status_code: Optional[pulumi.Input[builtins.str]] = None,
            timeout: Optional[pulumi.Input[builtins.int]] = None,
            url_path: Optional[pulumi.Input[builtins.str]] = None) -> 'ElbMonitor':
        """
        Get an existing ElbMonitor resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _ElbMonitorState.__new__(_ElbMonitorState)

        __props__.__dict__["domain_name"] = domain_name
        __props__.__dict__["interval"] = interval
        __props__.__dict__["max_retries"] = max_retries
        __props__.__dict__["pool_id"] = pool_id
        __props__.__dict__["port"] = port
        __props__.__dict__["protocol"] = protocol
        __props__.__dict__["region"] = region
        __props__.__dict__["status_code"] = status_code
        __props__.__dict__["timeout"] = timeout
        __props__.__dict__["url_path"] = url_path
        return ElbMonitor(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="domainName")
    def domain_name(self) -> pulumi.Output[builtins.str]:
        return pulumi.get(self, "domain_name")

    @property
    @pulumi.getter
    def interval(self) -> pulumi.Output[builtins.int]:
        return pulumi.get(self, "interval")

    @property
    @pulumi.getter(name="maxRetries")
    def max_retries(self) -> pulumi.Output[builtins.int]:
        return pulumi.get(self, "max_retries")

    @property
    @pulumi.getter(name="poolId")
    def pool_id(self) -> pulumi.Output[builtins.str]:
        return pulumi.get(self, "pool_id")

    @property
    @pulumi.getter
    def port(self) -> pulumi.Output[builtins.int]:
        return pulumi.get(self, "port")

    @property
    @pulumi.getter
    def protocol(self) -> pulumi.Output[builtins.str]:
        return pulumi.get(self, "protocol")

    @property
    @pulumi.getter
    def region(self) -> pulumi.Output[builtins.str]:
        return pulumi.get(self, "region")

    @property
    @pulumi.getter(name="statusCode")
    def status_code(self) -> pulumi.Output[builtins.str]:
        return pulumi.get(self, "status_code")

    @property
    @pulumi.getter
    def timeout(self) -> pulumi.Output[builtins.int]:
        return pulumi.get(self, "timeout")

    @property
    @pulumi.getter(name="urlPath")
    def url_path(self) -> pulumi.Output[builtins.str]:
        return pulumi.get(self, "url_path")

