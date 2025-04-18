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
    'GetApigApplicationsResult',
    'AwaitableGetApigApplicationsResult',
    'get_apig_applications',
    'get_apig_applications_output',
]

@pulumi.output_type
class GetApigApplicationsResult:
    """
    A collection of values returned by getApigApplications.
    """
    def __init__(__self__, app_key=None, application_id=None, applications=None, created_by=None, id=None, instance_id=None, name=None, region=None):
        if app_key and not isinstance(app_key, str):
            raise TypeError("Expected argument 'app_key' to be a str")
        pulumi.set(__self__, "app_key", app_key)
        if application_id and not isinstance(application_id, str):
            raise TypeError("Expected argument 'application_id' to be a str")
        pulumi.set(__self__, "application_id", application_id)
        if applications and not isinstance(applications, list):
            raise TypeError("Expected argument 'applications' to be a list")
        pulumi.set(__self__, "applications", applications)
        if created_by and not isinstance(created_by, str):
            raise TypeError("Expected argument 'created_by' to be a str")
        pulumi.set(__self__, "created_by", created_by)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if instance_id and not isinstance(instance_id, str):
            raise TypeError("Expected argument 'instance_id' to be a str")
        pulumi.set(__self__, "instance_id", instance_id)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if region and not isinstance(region, str):
            raise TypeError("Expected argument 'region' to be a str")
        pulumi.set(__self__, "region", region)

    @property
    @pulumi.getter(name="appKey")
    def app_key(self) -> Optional[builtins.str]:
        return pulumi.get(self, "app_key")

    @property
    @pulumi.getter(name="applicationId")
    def application_id(self) -> Optional[builtins.str]:
        return pulumi.get(self, "application_id")

    @property
    @pulumi.getter
    def applications(self) -> Sequence['outputs.GetApigApplicationsApplicationResult']:
        return pulumi.get(self, "applications")

    @property
    @pulumi.getter(name="createdBy")
    def created_by(self) -> Optional[builtins.str]:
        return pulumi.get(self, "created_by")

    @property
    @pulumi.getter
    def id(self) -> builtins.str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> builtins.str:
        return pulumi.get(self, "instance_id")

    @property
    @pulumi.getter
    def name(self) -> Optional[builtins.str]:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def region(self) -> builtins.str:
        return pulumi.get(self, "region")


class AwaitableGetApigApplicationsResult(GetApigApplicationsResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetApigApplicationsResult(
            app_key=self.app_key,
            application_id=self.application_id,
            applications=self.applications,
            created_by=self.created_by,
            id=self.id,
            instance_id=self.instance_id,
            name=self.name,
            region=self.region)


def get_apig_applications(app_key: Optional[builtins.str] = None,
                          application_id: Optional[builtins.str] = None,
                          created_by: Optional[builtins.str] = None,
                          instance_id: Optional[builtins.str] = None,
                          name: Optional[builtins.str] = None,
                          region: Optional[builtins.str] = None,
                          opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetApigApplicationsResult:
    """
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()
    __args__['appKey'] = app_key
    __args__['applicationId'] = application_id
    __args__['createdBy'] = created_by
    __args__['instanceId'] = instance_id
    __args__['name'] = name
    __args__['region'] = region
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('sbercloud:index/getApigApplications:getApigApplications', __args__, opts=opts, typ=GetApigApplicationsResult).value

    return AwaitableGetApigApplicationsResult(
        app_key=pulumi.get(__ret__, 'app_key'),
        application_id=pulumi.get(__ret__, 'application_id'),
        applications=pulumi.get(__ret__, 'applications'),
        created_by=pulumi.get(__ret__, 'created_by'),
        id=pulumi.get(__ret__, 'id'),
        instance_id=pulumi.get(__ret__, 'instance_id'),
        name=pulumi.get(__ret__, 'name'),
        region=pulumi.get(__ret__, 'region'))
def get_apig_applications_output(app_key: Optional[pulumi.Input[Optional[builtins.str]]] = None,
                                 application_id: Optional[pulumi.Input[Optional[builtins.str]]] = None,
                                 created_by: Optional[pulumi.Input[Optional[builtins.str]]] = None,
                                 instance_id: Optional[pulumi.Input[builtins.str]] = None,
                                 name: Optional[pulumi.Input[Optional[builtins.str]]] = None,
                                 region: Optional[pulumi.Input[Optional[builtins.str]]] = None,
                                 opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetApigApplicationsResult]:
    """
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()
    __args__['appKey'] = app_key
    __args__['applicationId'] = application_id
    __args__['createdBy'] = created_by
    __args__['instanceId'] = instance_id
    __args__['name'] = name
    __args__['region'] = region
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('sbercloud:index/getApigApplications:getApigApplications', __args__, opts=opts, typ=GetApigApplicationsResult)
    return __ret__.apply(lambda __response__: GetApigApplicationsResult(
        app_key=pulumi.get(__response__, 'app_key'),
        application_id=pulumi.get(__response__, 'application_id'),
        applications=pulumi.get(__response__, 'applications'),
        created_by=pulumi.get(__response__, 'created_by'),
        id=pulumi.get(__response__, 'id'),
        instance_id=pulumi.get(__response__, 'instance_id'),
        name=pulumi.get(__response__, 'name'),
        region=pulumi.get(__response__, 'region')))
