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

__all__ = [
    'GetCceClustersResult',
    'AwaitableGetCceClustersResult',
    'get_cce_clusters',
    'get_cce_clusters_output',
]

@pulumi.output_type
class GetCceClustersResult:
    """
    A collection of values returned by getCceClusters.
    """
    def __init__(__self__, cluster_id=None, cluster_type=None, clusters=None, enterprise_project_id=None, id=None, ids=None, name=None, region=None, status=None, vpc_id=None):
        if cluster_id and not isinstance(cluster_id, str):
            raise TypeError("Expected argument 'cluster_id' to be a str")
        pulumi.set(__self__, "cluster_id", cluster_id)
        if cluster_type and not isinstance(cluster_type, str):
            raise TypeError("Expected argument 'cluster_type' to be a str")
        pulumi.set(__self__, "cluster_type", cluster_type)
        if clusters and not isinstance(clusters, list):
            raise TypeError("Expected argument 'clusters' to be a list")
        pulumi.set(__self__, "clusters", clusters)
        if enterprise_project_id and not isinstance(enterprise_project_id, str):
            raise TypeError("Expected argument 'enterprise_project_id' to be a str")
        pulumi.set(__self__, "enterprise_project_id", enterprise_project_id)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if ids and not isinstance(ids, list):
            raise TypeError("Expected argument 'ids' to be a list")
        pulumi.set(__self__, "ids", ids)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if region and not isinstance(region, str):
            raise TypeError("Expected argument 'region' to be a str")
        pulumi.set(__self__, "region", region)
        if status and not isinstance(status, str):
            raise TypeError("Expected argument 'status' to be a str")
        pulumi.set(__self__, "status", status)
        if vpc_id and not isinstance(vpc_id, str):
            raise TypeError("Expected argument 'vpc_id' to be a str")
        pulumi.set(__self__, "vpc_id", vpc_id)

    @property
    @pulumi.getter(name="clusterId")
    def cluster_id(self) -> Optional[builtins.str]:
        return pulumi.get(self, "cluster_id")

    @property
    @pulumi.getter(name="clusterType")
    def cluster_type(self) -> Optional[builtins.str]:
        return pulumi.get(self, "cluster_type")

    @property
    @pulumi.getter
    def clusters(self) -> Sequence['outputs.GetCceClustersClusterResult']:
        return pulumi.get(self, "clusters")

    @property
    @pulumi.getter(name="enterpriseProjectId")
    def enterprise_project_id(self) -> Optional[builtins.str]:
        return pulumi.get(self, "enterprise_project_id")

    @property
    @pulumi.getter
    def id(self) -> builtins.str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def ids(self) -> Sequence[builtins.str]:
        return pulumi.get(self, "ids")

    @property
    @pulumi.getter
    def name(self) -> Optional[builtins.str]:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def region(self) -> builtins.str:
        return pulumi.get(self, "region")

    @property
    @pulumi.getter
    def status(self) -> Optional[builtins.str]:
        return pulumi.get(self, "status")

    @property
    @pulumi.getter(name="vpcId")
    def vpc_id(self) -> Optional[builtins.str]:
        return pulumi.get(self, "vpc_id")


class AwaitableGetCceClustersResult(GetCceClustersResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetCceClustersResult(
            cluster_id=self.cluster_id,
            cluster_type=self.cluster_type,
            clusters=self.clusters,
            enterprise_project_id=self.enterprise_project_id,
            id=self.id,
            ids=self.ids,
            name=self.name,
            region=self.region,
            status=self.status,
            vpc_id=self.vpc_id)


def get_cce_clusters(cluster_id: Optional[builtins.str] = None,
                     cluster_type: Optional[builtins.str] = None,
                     enterprise_project_id: Optional[builtins.str] = None,
                     name: Optional[builtins.str] = None,
                     region: Optional[builtins.str] = None,
                     status: Optional[builtins.str] = None,
                     vpc_id: Optional[builtins.str] = None,
                     opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetCceClustersResult:
    """
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()
    __args__['clusterId'] = cluster_id
    __args__['clusterType'] = cluster_type
    __args__['enterpriseProjectId'] = enterprise_project_id
    __args__['name'] = name
    __args__['region'] = region
    __args__['status'] = status
    __args__['vpcId'] = vpc_id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('sbercloud:index/getCceClusters:getCceClusters', __args__, opts=opts, typ=GetCceClustersResult).value

    return AwaitableGetCceClustersResult(
        cluster_id=pulumi.get(__ret__, 'cluster_id'),
        cluster_type=pulumi.get(__ret__, 'cluster_type'),
        clusters=pulumi.get(__ret__, 'clusters'),
        enterprise_project_id=pulumi.get(__ret__, 'enterprise_project_id'),
        id=pulumi.get(__ret__, 'id'),
        ids=pulumi.get(__ret__, 'ids'),
        name=pulumi.get(__ret__, 'name'),
        region=pulumi.get(__ret__, 'region'),
        status=pulumi.get(__ret__, 'status'),
        vpc_id=pulumi.get(__ret__, 'vpc_id'))
def get_cce_clusters_output(cluster_id: Optional[pulumi.Input[Optional[builtins.str]]] = None,
                            cluster_type: Optional[pulumi.Input[Optional[builtins.str]]] = None,
                            enterprise_project_id: Optional[pulumi.Input[Optional[builtins.str]]] = None,
                            name: Optional[pulumi.Input[Optional[builtins.str]]] = None,
                            region: Optional[pulumi.Input[Optional[builtins.str]]] = None,
                            status: Optional[pulumi.Input[Optional[builtins.str]]] = None,
                            vpc_id: Optional[pulumi.Input[Optional[builtins.str]]] = None,
                            opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetCceClustersResult]:
    """
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()
    __args__['clusterId'] = cluster_id
    __args__['clusterType'] = cluster_type
    __args__['enterpriseProjectId'] = enterprise_project_id
    __args__['name'] = name
    __args__['region'] = region
    __args__['status'] = status
    __args__['vpcId'] = vpc_id
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('sbercloud:index/getCceClusters:getCceClusters', __args__, opts=opts, typ=GetCceClustersResult)
    return __ret__.apply(lambda __response__: GetCceClustersResult(
        cluster_id=pulumi.get(__response__, 'cluster_id'),
        cluster_type=pulumi.get(__response__, 'cluster_type'),
        clusters=pulumi.get(__response__, 'clusters'),
        enterprise_project_id=pulumi.get(__response__, 'enterprise_project_id'),
        id=pulumi.get(__response__, 'id'),
        ids=pulumi.get(__response__, 'ids'),
        name=pulumi.get(__response__, 'name'),
        region=pulumi.get(__response__, 'region'),
        status=pulumi.get(__response__, 'status'),
        vpc_id=pulumi.get(__response__, 'vpc_id')))
