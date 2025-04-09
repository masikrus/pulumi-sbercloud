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

__all__ = ['CssConfigurationArgs', 'CssConfiguration']

@pulumi.input_type
class CssConfigurationArgs:
    def __init__(__self__, *,
                 cluster_id: pulumi.Input[builtins.str],
                 http_cors_allow_credetials: Optional[pulumi.Input[builtins.str]] = None,
                 http_cors_allow_headers: Optional[pulumi.Input[builtins.str]] = None,
                 http_cors_allow_methods: Optional[pulumi.Input[builtins.str]] = None,
                 http_cors_allow_origin: Optional[pulumi.Input[builtins.str]] = None,
                 http_cors_enabled: Optional[pulumi.Input[builtins.str]] = None,
                 http_cors_max_age: Optional[pulumi.Input[builtins.str]] = None,
                 indices_queries_cache_size: Optional[pulumi.Input[builtins.str]] = None,
                 region: Optional[pulumi.Input[builtins.str]] = None,
                 reindex_remote_whitelist: Optional[pulumi.Input[builtins.str]] = None,
                 thread_pool_force_merge_size: Optional[pulumi.Input[builtins.str]] = None):
        """
        The set of arguments for constructing a CssConfiguration resource.
        :param pulumi.Input[builtins.str] cluster_id: The CSS cluster ID.
        :param pulumi.Input[builtins.str] http_cors_allow_credetials: Whether to return the Access-Control-Allow-Credentials of the header during cross-domain access.
        :param pulumi.Input[builtins.str] http_cors_allow_headers: Headers allowed for cross-domain access.
        :param pulumi.Input[builtins.str] http_cors_allow_methods: Methods allowed for cross-domain access.
        :param pulumi.Input[builtins.str] http_cors_allow_origin: Origin IP address allowed for cross-domain access, for example, **122.122.122.122:9200**.
        :param pulumi.Input[builtins.str] http_cors_enabled: Whether to allow cross-domain access.
        :param pulumi.Input[builtins.str] http_cors_max_age: Cache duration of the browser. The cache is automatically cleared after the time range you specify.
        :param pulumi.Input[builtins.str] indices_queries_cache_size: Cache size in the query phase. Value range: **1** to **100**.
        :param pulumi.Input[builtins.str] reindex_remote_whitelist: Configured for migrating data from the current cluster to the target cluster through the reindex API.
        :param pulumi.Input[builtins.str] thread_pool_force_merge_size: Queue size in the force merge thread pool.
        """
        pulumi.set(__self__, "cluster_id", cluster_id)
        if http_cors_allow_credetials is not None:
            pulumi.set(__self__, "http_cors_allow_credetials", http_cors_allow_credetials)
        if http_cors_allow_headers is not None:
            pulumi.set(__self__, "http_cors_allow_headers", http_cors_allow_headers)
        if http_cors_allow_methods is not None:
            pulumi.set(__self__, "http_cors_allow_methods", http_cors_allow_methods)
        if http_cors_allow_origin is not None:
            pulumi.set(__self__, "http_cors_allow_origin", http_cors_allow_origin)
        if http_cors_enabled is not None:
            pulumi.set(__self__, "http_cors_enabled", http_cors_enabled)
        if http_cors_max_age is not None:
            pulumi.set(__self__, "http_cors_max_age", http_cors_max_age)
        if indices_queries_cache_size is not None:
            pulumi.set(__self__, "indices_queries_cache_size", indices_queries_cache_size)
        if region is not None:
            pulumi.set(__self__, "region", region)
        if reindex_remote_whitelist is not None:
            pulumi.set(__self__, "reindex_remote_whitelist", reindex_remote_whitelist)
        if thread_pool_force_merge_size is not None:
            pulumi.set(__self__, "thread_pool_force_merge_size", thread_pool_force_merge_size)

    @property
    @pulumi.getter(name="clusterId")
    def cluster_id(self) -> pulumi.Input[builtins.str]:
        """
        The CSS cluster ID.
        """
        return pulumi.get(self, "cluster_id")

    @cluster_id.setter
    def cluster_id(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "cluster_id", value)

    @property
    @pulumi.getter(name="httpCorsAllowCredetials")
    def http_cors_allow_credetials(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        Whether to return the Access-Control-Allow-Credentials of the header during cross-domain access.
        """
        return pulumi.get(self, "http_cors_allow_credetials")

    @http_cors_allow_credetials.setter
    def http_cors_allow_credetials(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "http_cors_allow_credetials", value)

    @property
    @pulumi.getter(name="httpCorsAllowHeaders")
    def http_cors_allow_headers(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        Headers allowed for cross-domain access.
        """
        return pulumi.get(self, "http_cors_allow_headers")

    @http_cors_allow_headers.setter
    def http_cors_allow_headers(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "http_cors_allow_headers", value)

    @property
    @pulumi.getter(name="httpCorsAllowMethods")
    def http_cors_allow_methods(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        Methods allowed for cross-domain access.
        """
        return pulumi.get(self, "http_cors_allow_methods")

    @http_cors_allow_methods.setter
    def http_cors_allow_methods(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "http_cors_allow_methods", value)

    @property
    @pulumi.getter(name="httpCorsAllowOrigin")
    def http_cors_allow_origin(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        Origin IP address allowed for cross-domain access, for example, **122.122.122.122:9200**.
        """
        return pulumi.get(self, "http_cors_allow_origin")

    @http_cors_allow_origin.setter
    def http_cors_allow_origin(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "http_cors_allow_origin", value)

    @property
    @pulumi.getter(name="httpCorsEnabled")
    def http_cors_enabled(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        Whether to allow cross-domain access.
        """
        return pulumi.get(self, "http_cors_enabled")

    @http_cors_enabled.setter
    def http_cors_enabled(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "http_cors_enabled", value)

    @property
    @pulumi.getter(name="httpCorsMaxAge")
    def http_cors_max_age(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        Cache duration of the browser. The cache is automatically cleared after the time range you specify.
        """
        return pulumi.get(self, "http_cors_max_age")

    @http_cors_max_age.setter
    def http_cors_max_age(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "http_cors_max_age", value)

    @property
    @pulumi.getter(name="indicesQueriesCacheSize")
    def indices_queries_cache_size(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        Cache size in the query phase. Value range: **1** to **100**.
        """
        return pulumi.get(self, "indices_queries_cache_size")

    @indices_queries_cache_size.setter
    def indices_queries_cache_size(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "indices_queries_cache_size", value)

    @property
    @pulumi.getter
    def region(self) -> Optional[pulumi.Input[builtins.str]]:
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "region", value)

    @property
    @pulumi.getter(name="reindexRemoteWhitelist")
    def reindex_remote_whitelist(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        Configured for migrating data from the current cluster to the target cluster through the reindex API.
        """
        return pulumi.get(self, "reindex_remote_whitelist")

    @reindex_remote_whitelist.setter
    def reindex_remote_whitelist(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "reindex_remote_whitelist", value)

    @property
    @pulumi.getter(name="threadPoolForceMergeSize")
    def thread_pool_force_merge_size(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        Queue size in the force merge thread pool.
        """
        return pulumi.get(self, "thread_pool_force_merge_size")

    @thread_pool_force_merge_size.setter
    def thread_pool_force_merge_size(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "thread_pool_force_merge_size", value)


@pulumi.input_type
class _CssConfigurationState:
    def __init__(__self__, *,
                 cluster_id: Optional[pulumi.Input[builtins.str]] = None,
                 http_cors_allow_credetials: Optional[pulumi.Input[builtins.str]] = None,
                 http_cors_allow_headers: Optional[pulumi.Input[builtins.str]] = None,
                 http_cors_allow_methods: Optional[pulumi.Input[builtins.str]] = None,
                 http_cors_allow_origin: Optional[pulumi.Input[builtins.str]] = None,
                 http_cors_enabled: Optional[pulumi.Input[builtins.str]] = None,
                 http_cors_max_age: Optional[pulumi.Input[builtins.str]] = None,
                 indices_queries_cache_size: Optional[pulumi.Input[builtins.str]] = None,
                 region: Optional[pulumi.Input[builtins.str]] = None,
                 reindex_remote_whitelist: Optional[pulumi.Input[builtins.str]] = None,
                 thread_pool_force_merge_size: Optional[pulumi.Input[builtins.str]] = None):
        """
        Input properties used for looking up and filtering CssConfiguration resources.
        :param pulumi.Input[builtins.str] cluster_id: The CSS cluster ID.
        :param pulumi.Input[builtins.str] http_cors_allow_credetials: Whether to return the Access-Control-Allow-Credentials of the header during cross-domain access.
        :param pulumi.Input[builtins.str] http_cors_allow_headers: Headers allowed for cross-domain access.
        :param pulumi.Input[builtins.str] http_cors_allow_methods: Methods allowed for cross-domain access.
        :param pulumi.Input[builtins.str] http_cors_allow_origin: Origin IP address allowed for cross-domain access, for example, **122.122.122.122:9200**.
        :param pulumi.Input[builtins.str] http_cors_enabled: Whether to allow cross-domain access.
        :param pulumi.Input[builtins.str] http_cors_max_age: Cache duration of the browser. The cache is automatically cleared after the time range you specify.
        :param pulumi.Input[builtins.str] indices_queries_cache_size: Cache size in the query phase. Value range: **1** to **100**.
        :param pulumi.Input[builtins.str] reindex_remote_whitelist: Configured for migrating data from the current cluster to the target cluster through the reindex API.
        :param pulumi.Input[builtins.str] thread_pool_force_merge_size: Queue size in the force merge thread pool.
        """
        if cluster_id is not None:
            pulumi.set(__self__, "cluster_id", cluster_id)
        if http_cors_allow_credetials is not None:
            pulumi.set(__self__, "http_cors_allow_credetials", http_cors_allow_credetials)
        if http_cors_allow_headers is not None:
            pulumi.set(__self__, "http_cors_allow_headers", http_cors_allow_headers)
        if http_cors_allow_methods is not None:
            pulumi.set(__self__, "http_cors_allow_methods", http_cors_allow_methods)
        if http_cors_allow_origin is not None:
            pulumi.set(__self__, "http_cors_allow_origin", http_cors_allow_origin)
        if http_cors_enabled is not None:
            pulumi.set(__self__, "http_cors_enabled", http_cors_enabled)
        if http_cors_max_age is not None:
            pulumi.set(__self__, "http_cors_max_age", http_cors_max_age)
        if indices_queries_cache_size is not None:
            pulumi.set(__self__, "indices_queries_cache_size", indices_queries_cache_size)
        if region is not None:
            pulumi.set(__self__, "region", region)
        if reindex_remote_whitelist is not None:
            pulumi.set(__self__, "reindex_remote_whitelist", reindex_remote_whitelist)
        if thread_pool_force_merge_size is not None:
            pulumi.set(__self__, "thread_pool_force_merge_size", thread_pool_force_merge_size)

    @property
    @pulumi.getter(name="clusterId")
    def cluster_id(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The CSS cluster ID.
        """
        return pulumi.get(self, "cluster_id")

    @cluster_id.setter
    def cluster_id(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "cluster_id", value)

    @property
    @pulumi.getter(name="httpCorsAllowCredetials")
    def http_cors_allow_credetials(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        Whether to return the Access-Control-Allow-Credentials of the header during cross-domain access.
        """
        return pulumi.get(self, "http_cors_allow_credetials")

    @http_cors_allow_credetials.setter
    def http_cors_allow_credetials(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "http_cors_allow_credetials", value)

    @property
    @pulumi.getter(name="httpCorsAllowHeaders")
    def http_cors_allow_headers(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        Headers allowed for cross-domain access.
        """
        return pulumi.get(self, "http_cors_allow_headers")

    @http_cors_allow_headers.setter
    def http_cors_allow_headers(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "http_cors_allow_headers", value)

    @property
    @pulumi.getter(name="httpCorsAllowMethods")
    def http_cors_allow_methods(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        Methods allowed for cross-domain access.
        """
        return pulumi.get(self, "http_cors_allow_methods")

    @http_cors_allow_methods.setter
    def http_cors_allow_methods(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "http_cors_allow_methods", value)

    @property
    @pulumi.getter(name="httpCorsAllowOrigin")
    def http_cors_allow_origin(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        Origin IP address allowed for cross-domain access, for example, **122.122.122.122:9200**.
        """
        return pulumi.get(self, "http_cors_allow_origin")

    @http_cors_allow_origin.setter
    def http_cors_allow_origin(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "http_cors_allow_origin", value)

    @property
    @pulumi.getter(name="httpCorsEnabled")
    def http_cors_enabled(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        Whether to allow cross-domain access.
        """
        return pulumi.get(self, "http_cors_enabled")

    @http_cors_enabled.setter
    def http_cors_enabled(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "http_cors_enabled", value)

    @property
    @pulumi.getter(name="httpCorsMaxAge")
    def http_cors_max_age(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        Cache duration of the browser. The cache is automatically cleared after the time range you specify.
        """
        return pulumi.get(self, "http_cors_max_age")

    @http_cors_max_age.setter
    def http_cors_max_age(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "http_cors_max_age", value)

    @property
    @pulumi.getter(name="indicesQueriesCacheSize")
    def indices_queries_cache_size(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        Cache size in the query phase. Value range: **1** to **100**.
        """
        return pulumi.get(self, "indices_queries_cache_size")

    @indices_queries_cache_size.setter
    def indices_queries_cache_size(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "indices_queries_cache_size", value)

    @property
    @pulumi.getter
    def region(self) -> Optional[pulumi.Input[builtins.str]]:
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "region", value)

    @property
    @pulumi.getter(name="reindexRemoteWhitelist")
    def reindex_remote_whitelist(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        Configured for migrating data from the current cluster to the target cluster through the reindex API.
        """
        return pulumi.get(self, "reindex_remote_whitelist")

    @reindex_remote_whitelist.setter
    def reindex_remote_whitelist(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "reindex_remote_whitelist", value)

    @property
    @pulumi.getter(name="threadPoolForceMergeSize")
    def thread_pool_force_merge_size(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        Queue size in the force merge thread pool.
        """
        return pulumi.get(self, "thread_pool_force_merge_size")

    @thread_pool_force_merge_size.setter
    def thread_pool_force_merge_size(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "thread_pool_force_merge_size", value)


class CssConfiguration(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 cluster_id: Optional[pulumi.Input[builtins.str]] = None,
                 http_cors_allow_credetials: Optional[pulumi.Input[builtins.str]] = None,
                 http_cors_allow_headers: Optional[pulumi.Input[builtins.str]] = None,
                 http_cors_allow_methods: Optional[pulumi.Input[builtins.str]] = None,
                 http_cors_allow_origin: Optional[pulumi.Input[builtins.str]] = None,
                 http_cors_enabled: Optional[pulumi.Input[builtins.str]] = None,
                 http_cors_max_age: Optional[pulumi.Input[builtins.str]] = None,
                 indices_queries_cache_size: Optional[pulumi.Input[builtins.str]] = None,
                 region: Optional[pulumi.Input[builtins.str]] = None,
                 reindex_remote_whitelist: Optional[pulumi.Input[builtins.str]] = None,
                 thread_pool_force_merge_size: Optional[pulumi.Input[builtins.str]] = None,
                 __props__=None):
        """
        Create a CssConfiguration resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[builtins.str] cluster_id: The CSS cluster ID.
        :param pulumi.Input[builtins.str] http_cors_allow_credetials: Whether to return the Access-Control-Allow-Credentials of the header during cross-domain access.
        :param pulumi.Input[builtins.str] http_cors_allow_headers: Headers allowed for cross-domain access.
        :param pulumi.Input[builtins.str] http_cors_allow_methods: Methods allowed for cross-domain access.
        :param pulumi.Input[builtins.str] http_cors_allow_origin: Origin IP address allowed for cross-domain access, for example, **122.122.122.122:9200**.
        :param pulumi.Input[builtins.str] http_cors_enabled: Whether to allow cross-domain access.
        :param pulumi.Input[builtins.str] http_cors_max_age: Cache duration of the browser. The cache is automatically cleared after the time range you specify.
        :param pulumi.Input[builtins.str] indices_queries_cache_size: Cache size in the query phase. Value range: **1** to **100**.
        :param pulumi.Input[builtins.str] reindex_remote_whitelist: Configured for migrating data from the current cluster to the target cluster through the reindex API.
        :param pulumi.Input[builtins.str] thread_pool_force_merge_size: Queue size in the force merge thread pool.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: CssConfigurationArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a CssConfiguration resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param CssConfigurationArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(CssConfigurationArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 cluster_id: Optional[pulumi.Input[builtins.str]] = None,
                 http_cors_allow_credetials: Optional[pulumi.Input[builtins.str]] = None,
                 http_cors_allow_headers: Optional[pulumi.Input[builtins.str]] = None,
                 http_cors_allow_methods: Optional[pulumi.Input[builtins.str]] = None,
                 http_cors_allow_origin: Optional[pulumi.Input[builtins.str]] = None,
                 http_cors_enabled: Optional[pulumi.Input[builtins.str]] = None,
                 http_cors_max_age: Optional[pulumi.Input[builtins.str]] = None,
                 indices_queries_cache_size: Optional[pulumi.Input[builtins.str]] = None,
                 region: Optional[pulumi.Input[builtins.str]] = None,
                 reindex_remote_whitelist: Optional[pulumi.Input[builtins.str]] = None,
                 thread_pool_force_merge_size: Optional[pulumi.Input[builtins.str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = CssConfigurationArgs.__new__(CssConfigurationArgs)

            if cluster_id is None and not opts.urn:
                raise TypeError("Missing required property 'cluster_id'")
            __props__.__dict__["cluster_id"] = cluster_id
            __props__.__dict__["http_cors_allow_credetials"] = http_cors_allow_credetials
            __props__.__dict__["http_cors_allow_headers"] = http_cors_allow_headers
            __props__.__dict__["http_cors_allow_methods"] = http_cors_allow_methods
            __props__.__dict__["http_cors_allow_origin"] = http_cors_allow_origin
            __props__.__dict__["http_cors_enabled"] = http_cors_enabled
            __props__.__dict__["http_cors_max_age"] = http_cors_max_age
            __props__.__dict__["indices_queries_cache_size"] = indices_queries_cache_size
            __props__.__dict__["region"] = region
            __props__.__dict__["reindex_remote_whitelist"] = reindex_remote_whitelist
            __props__.__dict__["thread_pool_force_merge_size"] = thread_pool_force_merge_size
        super(CssConfiguration, __self__).__init__(
            'sbercloud:index/cssConfiguration:CssConfiguration',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            cluster_id: Optional[pulumi.Input[builtins.str]] = None,
            http_cors_allow_credetials: Optional[pulumi.Input[builtins.str]] = None,
            http_cors_allow_headers: Optional[pulumi.Input[builtins.str]] = None,
            http_cors_allow_methods: Optional[pulumi.Input[builtins.str]] = None,
            http_cors_allow_origin: Optional[pulumi.Input[builtins.str]] = None,
            http_cors_enabled: Optional[pulumi.Input[builtins.str]] = None,
            http_cors_max_age: Optional[pulumi.Input[builtins.str]] = None,
            indices_queries_cache_size: Optional[pulumi.Input[builtins.str]] = None,
            region: Optional[pulumi.Input[builtins.str]] = None,
            reindex_remote_whitelist: Optional[pulumi.Input[builtins.str]] = None,
            thread_pool_force_merge_size: Optional[pulumi.Input[builtins.str]] = None) -> 'CssConfiguration':
        """
        Get an existing CssConfiguration resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[builtins.str] cluster_id: The CSS cluster ID.
        :param pulumi.Input[builtins.str] http_cors_allow_credetials: Whether to return the Access-Control-Allow-Credentials of the header during cross-domain access.
        :param pulumi.Input[builtins.str] http_cors_allow_headers: Headers allowed for cross-domain access.
        :param pulumi.Input[builtins.str] http_cors_allow_methods: Methods allowed for cross-domain access.
        :param pulumi.Input[builtins.str] http_cors_allow_origin: Origin IP address allowed for cross-domain access, for example, **122.122.122.122:9200**.
        :param pulumi.Input[builtins.str] http_cors_enabled: Whether to allow cross-domain access.
        :param pulumi.Input[builtins.str] http_cors_max_age: Cache duration of the browser. The cache is automatically cleared after the time range you specify.
        :param pulumi.Input[builtins.str] indices_queries_cache_size: Cache size in the query phase. Value range: **1** to **100**.
        :param pulumi.Input[builtins.str] reindex_remote_whitelist: Configured for migrating data from the current cluster to the target cluster through the reindex API.
        :param pulumi.Input[builtins.str] thread_pool_force_merge_size: Queue size in the force merge thread pool.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _CssConfigurationState.__new__(_CssConfigurationState)

        __props__.__dict__["cluster_id"] = cluster_id
        __props__.__dict__["http_cors_allow_credetials"] = http_cors_allow_credetials
        __props__.__dict__["http_cors_allow_headers"] = http_cors_allow_headers
        __props__.__dict__["http_cors_allow_methods"] = http_cors_allow_methods
        __props__.__dict__["http_cors_allow_origin"] = http_cors_allow_origin
        __props__.__dict__["http_cors_enabled"] = http_cors_enabled
        __props__.__dict__["http_cors_max_age"] = http_cors_max_age
        __props__.__dict__["indices_queries_cache_size"] = indices_queries_cache_size
        __props__.__dict__["region"] = region
        __props__.__dict__["reindex_remote_whitelist"] = reindex_remote_whitelist
        __props__.__dict__["thread_pool_force_merge_size"] = thread_pool_force_merge_size
        return CssConfiguration(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="clusterId")
    def cluster_id(self) -> pulumi.Output[builtins.str]:
        """
        The CSS cluster ID.
        """
        return pulumi.get(self, "cluster_id")

    @property
    @pulumi.getter(name="httpCorsAllowCredetials")
    def http_cors_allow_credetials(self) -> pulumi.Output[builtins.str]:
        """
        Whether to return the Access-Control-Allow-Credentials of the header during cross-domain access.
        """
        return pulumi.get(self, "http_cors_allow_credetials")

    @property
    @pulumi.getter(name="httpCorsAllowHeaders")
    def http_cors_allow_headers(self) -> pulumi.Output[builtins.str]:
        """
        Headers allowed for cross-domain access.
        """
        return pulumi.get(self, "http_cors_allow_headers")

    @property
    @pulumi.getter(name="httpCorsAllowMethods")
    def http_cors_allow_methods(self) -> pulumi.Output[builtins.str]:
        """
        Methods allowed for cross-domain access.
        """
        return pulumi.get(self, "http_cors_allow_methods")

    @property
    @pulumi.getter(name="httpCorsAllowOrigin")
    def http_cors_allow_origin(self) -> pulumi.Output[builtins.str]:
        """
        Origin IP address allowed for cross-domain access, for example, **122.122.122.122:9200**.
        """
        return pulumi.get(self, "http_cors_allow_origin")

    @property
    @pulumi.getter(name="httpCorsEnabled")
    def http_cors_enabled(self) -> pulumi.Output[builtins.str]:
        """
        Whether to allow cross-domain access.
        """
        return pulumi.get(self, "http_cors_enabled")

    @property
    @pulumi.getter(name="httpCorsMaxAge")
    def http_cors_max_age(self) -> pulumi.Output[builtins.str]:
        """
        Cache duration of the browser. The cache is automatically cleared after the time range you specify.
        """
        return pulumi.get(self, "http_cors_max_age")

    @property
    @pulumi.getter(name="indicesQueriesCacheSize")
    def indices_queries_cache_size(self) -> pulumi.Output[builtins.str]:
        """
        Cache size in the query phase. Value range: **1** to **100**.
        """
        return pulumi.get(self, "indices_queries_cache_size")

    @property
    @pulumi.getter
    def region(self) -> pulumi.Output[builtins.str]:
        return pulumi.get(self, "region")

    @property
    @pulumi.getter(name="reindexRemoteWhitelist")
    def reindex_remote_whitelist(self) -> pulumi.Output[builtins.str]:
        """
        Configured for migrating data from the current cluster to the target cluster through the reindex API.
        """
        return pulumi.get(self, "reindex_remote_whitelist")

    @property
    @pulumi.getter(name="threadPoolForceMergeSize")
    def thread_pool_force_merge_size(self) -> pulumi.Output[builtins.str]:
        """
        Queue size in the force merge thread pool.
        """
        return pulumi.get(self, "thread_pool_force_merge_size")

