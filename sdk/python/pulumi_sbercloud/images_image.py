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

__all__ = ['ImagesImageArgs', 'ImagesImage']

@pulumi.input_type
class ImagesImageArgs:
    def __init__(__self__, *,
                 backup_id: Optional[pulumi.Input[builtins.str]] = None,
                 cmk_id: Optional[pulumi.Input[builtins.str]] = None,
                 description: Optional[pulumi.Input[builtins.str]] = None,
                 enterprise_project_id: Optional[pulumi.Input[builtins.str]] = None,
                 image_url: Optional[pulumi.Input[builtins.str]] = None,
                 instance_id: Optional[pulumi.Input[builtins.str]] = None,
                 is_config: Optional[pulumi.Input[builtins.bool]] = None,
                 max_ram: Optional[pulumi.Input[builtins.int]] = None,
                 min_disk: Optional[pulumi.Input[builtins.int]] = None,
                 min_ram: Optional[pulumi.Input[builtins.int]] = None,
                 name: Optional[pulumi.Input[builtins.str]] = None,
                 os_version: Optional[pulumi.Input[builtins.str]] = None,
                 region: Optional[pulumi.Input[builtins.str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]]] = None,
                 type: Optional[pulumi.Input[builtins.str]] = None,
                 vault_id: Optional[pulumi.Input[builtins.str]] = None,
                 volume_id: Optional[pulumi.Input[builtins.str]] = None):
        """
        The set of arguments for constructing a ImagesImage resource.
        """
        if backup_id is not None:
            pulumi.set(__self__, "backup_id", backup_id)
        if cmk_id is not None:
            pulumi.set(__self__, "cmk_id", cmk_id)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if enterprise_project_id is not None:
            pulumi.set(__self__, "enterprise_project_id", enterprise_project_id)
        if image_url is not None:
            pulumi.set(__self__, "image_url", image_url)
        if instance_id is not None:
            pulumi.set(__self__, "instance_id", instance_id)
        if is_config is not None:
            pulumi.set(__self__, "is_config", is_config)
        if max_ram is not None:
            pulumi.set(__self__, "max_ram", max_ram)
        if min_disk is not None:
            pulumi.set(__self__, "min_disk", min_disk)
        if min_ram is not None:
            pulumi.set(__self__, "min_ram", min_ram)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if os_version is not None:
            pulumi.set(__self__, "os_version", os_version)
        if region is not None:
            pulumi.set(__self__, "region", region)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)
        if type is not None:
            pulumi.set(__self__, "type", type)
        if vault_id is not None:
            pulumi.set(__self__, "vault_id", vault_id)
        if volume_id is not None:
            pulumi.set(__self__, "volume_id", volume_id)

    @property
    @pulumi.getter(name="backupId")
    def backup_id(self) -> Optional[pulumi.Input[builtins.str]]:
        return pulumi.get(self, "backup_id")

    @backup_id.setter
    def backup_id(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "backup_id", value)

    @property
    @pulumi.getter(name="cmkId")
    def cmk_id(self) -> Optional[pulumi.Input[builtins.str]]:
        return pulumi.get(self, "cmk_id")

    @cmk_id.setter
    def cmk_id(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "cmk_id", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[builtins.str]]:
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="enterpriseProjectId")
    def enterprise_project_id(self) -> Optional[pulumi.Input[builtins.str]]:
        return pulumi.get(self, "enterprise_project_id")

    @enterprise_project_id.setter
    def enterprise_project_id(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "enterprise_project_id", value)

    @property
    @pulumi.getter(name="imageUrl")
    def image_url(self) -> Optional[pulumi.Input[builtins.str]]:
        return pulumi.get(self, "image_url")

    @image_url.setter
    def image_url(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "image_url", value)

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> Optional[pulumi.Input[builtins.str]]:
        return pulumi.get(self, "instance_id")

    @instance_id.setter
    def instance_id(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "instance_id", value)

    @property
    @pulumi.getter(name="isConfig")
    def is_config(self) -> Optional[pulumi.Input[builtins.bool]]:
        return pulumi.get(self, "is_config")

    @is_config.setter
    def is_config(self, value: Optional[pulumi.Input[builtins.bool]]):
        pulumi.set(self, "is_config", value)

    @property
    @pulumi.getter(name="maxRam")
    def max_ram(self) -> Optional[pulumi.Input[builtins.int]]:
        return pulumi.get(self, "max_ram")

    @max_ram.setter
    def max_ram(self, value: Optional[pulumi.Input[builtins.int]]):
        pulumi.set(self, "max_ram", value)

    @property
    @pulumi.getter(name="minDisk")
    def min_disk(self) -> Optional[pulumi.Input[builtins.int]]:
        return pulumi.get(self, "min_disk")

    @min_disk.setter
    def min_disk(self, value: Optional[pulumi.Input[builtins.int]]):
        pulumi.set(self, "min_disk", value)

    @property
    @pulumi.getter(name="minRam")
    def min_ram(self) -> Optional[pulumi.Input[builtins.int]]:
        return pulumi.get(self, "min_ram")

    @min_ram.setter
    def min_ram(self, value: Optional[pulumi.Input[builtins.int]]):
        pulumi.set(self, "min_ram", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[builtins.str]]:
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="osVersion")
    def os_version(self) -> Optional[pulumi.Input[builtins.str]]:
        return pulumi.get(self, "os_version")

    @os_version.setter
    def os_version(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "os_version", value)

    @property
    @pulumi.getter
    def region(self) -> Optional[pulumi.Input[builtins.str]]:
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "region", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]]]:
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]]]):
        pulumi.set(self, "tags", value)

    @property
    @pulumi.getter
    def type(self) -> Optional[pulumi.Input[builtins.str]]:
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "type", value)

    @property
    @pulumi.getter(name="vaultId")
    def vault_id(self) -> Optional[pulumi.Input[builtins.str]]:
        return pulumi.get(self, "vault_id")

    @vault_id.setter
    def vault_id(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "vault_id", value)

    @property
    @pulumi.getter(name="volumeId")
    def volume_id(self) -> Optional[pulumi.Input[builtins.str]]:
        return pulumi.get(self, "volume_id")

    @volume_id.setter
    def volume_id(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "volume_id", value)


@pulumi.input_type
class _ImagesImageState:
    def __init__(__self__, *,
                 backup_id: Optional[pulumi.Input[builtins.str]] = None,
                 checksum: Optional[pulumi.Input[builtins.str]] = None,
                 cmk_id: Optional[pulumi.Input[builtins.str]] = None,
                 data_origin: Optional[pulumi.Input[builtins.str]] = None,
                 description: Optional[pulumi.Input[builtins.str]] = None,
                 disk_format: Optional[pulumi.Input[builtins.str]] = None,
                 enterprise_project_id: Optional[pulumi.Input[builtins.str]] = None,
                 image_size: Optional[pulumi.Input[builtins.str]] = None,
                 image_url: Optional[pulumi.Input[builtins.str]] = None,
                 instance_id: Optional[pulumi.Input[builtins.str]] = None,
                 is_config: Optional[pulumi.Input[builtins.bool]] = None,
                 max_ram: Optional[pulumi.Input[builtins.int]] = None,
                 min_disk: Optional[pulumi.Input[builtins.int]] = None,
                 min_ram: Optional[pulumi.Input[builtins.int]] = None,
                 name: Optional[pulumi.Input[builtins.str]] = None,
                 os_version: Optional[pulumi.Input[builtins.str]] = None,
                 region: Optional[pulumi.Input[builtins.str]] = None,
                 status: Optional[pulumi.Input[builtins.str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]]] = None,
                 type: Optional[pulumi.Input[builtins.str]] = None,
                 vault_id: Optional[pulumi.Input[builtins.str]] = None,
                 visibility: Optional[pulumi.Input[builtins.str]] = None,
                 volume_id: Optional[pulumi.Input[builtins.str]] = None):
        """
        Input properties used for looking up and filtering ImagesImage resources.
        """
        if backup_id is not None:
            pulumi.set(__self__, "backup_id", backup_id)
        if checksum is not None:
            pulumi.set(__self__, "checksum", checksum)
        if cmk_id is not None:
            pulumi.set(__self__, "cmk_id", cmk_id)
        if data_origin is not None:
            pulumi.set(__self__, "data_origin", data_origin)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if disk_format is not None:
            pulumi.set(__self__, "disk_format", disk_format)
        if enterprise_project_id is not None:
            pulumi.set(__self__, "enterprise_project_id", enterprise_project_id)
        if image_size is not None:
            pulumi.set(__self__, "image_size", image_size)
        if image_url is not None:
            pulumi.set(__self__, "image_url", image_url)
        if instance_id is not None:
            pulumi.set(__self__, "instance_id", instance_id)
        if is_config is not None:
            pulumi.set(__self__, "is_config", is_config)
        if max_ram is not None:
            pulumi.set(__self__, "max_ram", max_ram)
        if min_disk is not None:
            pulumi.set(__self__, "min_disk", min_disk)
        if min_ram is not None:
            pulumi.set(__self__, "min_ram", min_ram)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if os_version is not None:
            pulumi.set(__self__, "os_version", os_version)
        if region is not None:
            pulumi.set(__self__, "region", region)
        if status is not None:
            pulumi.set(__self__, "status", status)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)
        if type is not None:
            pulumi.set(__self__, "type", type)
        if vault_id is not None:
            pulumi.set(__self__, "vault_id", vault_id)
        if visibility is not None:
            pulumi.set(__self__, "visibility", visibility)
        if volume_id is not None:
            pulumi.set(__self__, "volume_id", volume_id)

    @property
    @pulumi.getter(name="backupId")
    def backup_id(self) -> Optional[pulumi.Input[builtins.str]]:
        return pulumi.get(self, "backup_id")

    @backup_id.setter
    def backup_id(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "backup_id", value)

    @property
    @pulumi.getter
    def checksum(self) -> Optional[pulumi.Input[builtins.str]]:
        return pulumi.get(self, "checksum")

    @checksum.setter
    def checksum(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "checksum", value)

    @property
    @pulumi.getter(name="cmkId")
    def cmk_id(self) -> Optional[pulumi.Input[builtins.str]]:
        return pulumi.get(self, "cmk_id")

    @cmk_id.setter
    def cmk_id(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "cmk_id", value)

    @property
    @pulumi.getter(name="dataOrigin")
    def data_origin(self) -> Optional[pulumi.Input[builtins.str]]:
        return pulumi.get(self, "data_origin")

    @data_origin.setter
    def data_origin(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "data_origin", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[builtins.str]]:
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="diskFormat")
    def disk_format(self) -> Optional[pulumi.Input[builtins.str]]:
        return pulumi.get(self, "disk_format")

    @disk_format.setter
    def disk_format(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "disk_format", value)

    @property
    @pulumi.getter(name="enterpriseProjectId")
    def enterprise_project_id(self) -> Optional[pulumi.Input[builtins.str]]:
        return pulumi.get(self, "enterprise_project_id")

    @enterprise_project_id.setter
    def enterprise_project_id(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "enterprise_project_id", value)

    @property
    @pulumi.getter(name="imageSize")
    def image_size(self) -> Optional[pulumi.Input[builtins.str]]:
        return pulumi.get(self, "image_size")

    @image_size.setter
    def image_size(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "image_size", value)

    @property
    @pulumi.getter(name="imageUrl")
    def image_url(self) -> Optional[pulumi.Input[builtins.str]]:
        return pulumi.get(self, "image_url")

    @image_url.setter
    def image_url(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "image_url", value)

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> Optional[pulumi.Input[builtins.str]]:
        return pulumi.get(self, "instance_id")

    @instance_id.setter
    def instance_id(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "instance_id", value)

    @property
    @pulumi.getter(name="isConfig")
    def is_config(self) -> Optional[pulumi.Input[builtins.bool]]:
        return pulumi.get(self, "is_config")

    @is_config.setter
    def is_config(self, value: Optional[pulumi.Input[builtins.bool]]):
        pulumi.set(self, "is_config", value)

    @property
    @pulumi.getter(name="maxRam")
    def max_ram(self) -> Optional[pulumi.Input[builtins.int]]:
        return pulumi.get(self, "max_ram")

    @max_ram.setter
    def max_ram(self, value: Optional[pulumi.Input[builtins.int]]):
        pulumi.set(self, "max_ram", value)

    @property
    @pulumi.getter(name="minDisk")
    def min_disk(self) -> Optional[pulumi.Input[builtins.int]]:
        return pulumi.get(self, "min_disk")

    @min_disk.setter
    def min_disk(self, value: Optional[pulumi.Input[builtins.int]]):
        pulumi.set(self, "min_disk", value)

    @property
    @pulumi.getter(name="minRam")
    def min_ram(self) -> Optional[pulumi.Input[builtins.int]]:
        return pulumi.get(self, "min_ram")

    @min_ram.setter
    def min_ram(self, value: Optional[pulumi.Input[builtins.int]]):
        pulumi.set(self, "min_ram", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[builtins.str]]:
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="osVersion")
    def os_version(self) -> Optional[pulumi.Input[builtins.str]]:
        return pulumi.get(self, "os_version")

    @os_version.setter
    def os_version(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "os_version", value)

    @property
    @pulumi.getter
    def region(self) -> Optional[pulumi.Input[builtins.str]]:
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "region", value)

    @property
    @pulumi.getter
    def status(self) -> Optional[pulumi.Input[builtins.str]]:
        return pulumi.get(self, "status")

    @status.setter
    def status(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "status", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]]]:
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]]]):
        pulumi.set(self, "tags", value)

    @property
    @pulumi.getter
    def type(self) -> Optional[pulumi.Input[builtins.str]]:
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "type", value)

    @property
    @pulumi.getter(name="vaultId")
    def vault_id(self) -> Optional[pulumi.Input[builtins.str]]:
        return pulumi.get(self, "vault_id")

    @vault_id.setter
    def vault_id(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "vault_id", value)

    @property
    @pulumi.getter
    def visibility(self) -> Optional[pulumi.Input[builtins.str]]:
        return pulumi.get(self, "visibility")

    @visibility.setter
    def visibility(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "visibility", value)

    @property
    @pulumi.getter(name="volumeId")
    def volume_id(self) -> Optional[pulumi.Input[builtins.str]]:
        return pulumi.get(self, "volume_id")

    @volume_id.setter
    def volume_id(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "volume_id", value)


class ImagesImage(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 backup_id: Optional[pulumi.Input[builtins.str]] = None,
                 cmk_id: Optional[pulumi.Input[builtins.str]] = None,
                 description: Optional[pulumi.Input[builtins.str]] = None,
                 enterprise_project_id: Optional[pulumi.Input[builtins.str]] = None,
                 image_url: Optional[pulumi.Input[builtins.str]] = None,
                 instance_id: Optional[pulumi.Input[builtins.str]] = None,
                 is_config: Optional[pulumi.Input[builtins.bool]] = None,
                 max_ram: Optional[pulumi.Input[builtins.int]] = None,
                 min_disk: Optional[pulumi.Input[builtins.int]] = None,
                 min_ram: Optional[pulumi.Input[builtins.int]] = None,
                 name: Optional[pulumi.Input[builtins.str]] = None,
                 os_version: Optional[pulumi.Input[builtins.str]] = None,
                 region: Optional[pulumi.Input[builtins.str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]]] = None,
                 type: Optional[pulumi.Input[builtins.str]] = None,
                 vault_id: Optional[pulumi.Input[builtins.str]] = None,
                 volume_id: Optional[pulumi.Input[builtins.str]] = None,
                 __props__=None):
        """
        Create a ImagesImage resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[ImagesImageArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a ImagesImage resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param ImagesImageArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ImagesImageArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 backup_id: Optional[pulumi.Input[builtins.str]] = None,
                 cmk_id: Optional[pulumi.Input[builtins.str]] = None,
                 description: Optional[pulumi.Input[builtins.str]] = None,
                 enterprise_project_id: Optional[pulumi.Input[builtins.str]] = None,
                 image_url: Optional[pulumi.Input[builtins.str]] = None,
                 instance_id: Optional[pulumi.Input[builtins.str]] = None,
                 is_config: Optional[pulumi.Input[builtins.bool]] = None,
                 max_ram: Optional[pulumi.Input[builtins.int]] = None,
                 min_disk: Optional[pulumi.Input[builtins.int]] = None,
                 min_ram: Optional[pulumi.Input[builtins.int]] = None,
                 name: Optional[pulumi.Input[builtins.str]] = None,
                 os_version: Optional[pulumi.Input[builtins.str]] = None,
                 region: Optional[pulumi.Input[builtins.str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]]] = None,
                 type: Optional[pulumi.Input[builtins.str]] = None,
                 vault_id: Optional[pulumi.Input[builtins.str]] = None,
                 volume_id: Optional[pulumi.Input[builtins.str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ImagesImageArgs.__new__(ImagesImageArgs)

            __props__.__dict__["backup_id"] = backup_id
            __props__.__dict__["cmk_id"] = cmk_id
            __props__.__dict__["description"] = description
            __props__.__dict__["enterprise_project_id"] = enterprise_project_id
            __props__.__dict__["image_url"] = image_url
            __props__.__dict__["instance_id"] = instance_id
            __props__.__dict__["is_config"] = is_config
            __props__.__dict__["max_ram"] = max_ram
            __props__.__dict__["min_disk"] = min_disk
            __props__.__dict__["min_ram"] = min_ram
            __props__.__dict__["name"] = name
            __props__.__dict__["os_version"] = os_version
            __props__.__dict__["region"] = region
            __props__.__dict__["tags"] = tags
            __props__.__dict__["type"] = type
            __props__.__dict__["vault_id"] = vault_id
            __props__.__dict__["volume_id"] = volume_id
            __props__.__dict__["checksum"] = None
            __props__.__dict__["data_origin"] = None
            __props__.__dict__["disk_format"] = None
            __props__.__dict__["image_size"] = None
            __props__.__dict__["status"] = None
            __props__.__dict__["visibility"] = None
        super(ImagesImage, __self__).__init__(
            'sbercloud:index/imagesImage:ImagesImage',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            backup_id: Optional[pulumi.Input[builtins.str]] = None,
            checksum: Optional[pulumi.Input[builtins.str]] = None,
            cmk_id: Optional[pulumi.Input[builtins.str]] = None,
            data_origin: Optional[pulumi.Input[builtins.str]] = None,
            description: Optional[pulumi.Input[builtins.str]] = None,
            disk_format: Optional[pulumi.Input[builtins.str]] = None,
            enterprise_project_id: Optional[pulumi.Input[builtins.str]] = None,
            image_size: Optional[pulumi.Input[builtins.str]] = None,
            image_url: Optional[pulumi.Input[builtins.str]] = None,
            instance_id: Optional[pulumi.Input[builtins.str]] = None,
            is_config: Optional[pulumi.Input[builtins.bool]] = None,
            max_ram: Optional[pulumi.Input[builtins.int]] = None,
            min_disk: Optional[pulumi.Input[builtins.int]] = None,
            min_ram: Optional[pulumi.Input[builtins.int]] = None,
            name: Optional[pulumi.Input[builtins.str]] = None,
            os_version: Optional[pulumi.Input[builtins.str]] = None,
            region: Optional[pulumi.Input[builtins.str]] = None,
            status: Optional[pulumi.Input[builtins.str]] = None,
            tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]]] = None,
            type: Optional[pulumi.Input[builtins.str]] = None,
            vault_id: Optional[pulumi.Input[builtins.str]] = None,
            visibility: Optional[pulumi.Input[builtins.str]] = None,
            volume_id: Optional[pulumi.Input[builtins.str]] = None) -> 'ImagesImage':
        """
        Get an existing ImagesImage resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _ImagesImageState.__new__(_ImagesImageState)

        __props__.__dict__["backup_id"] = backup_id
        __props__.__dict__["checksum"] = checksum
        __props__.__dict__["cmk_id"] = cmk_id
        __props__.__dict__["data_origin"] = data_origin
        __props__.__dict__["description"] = description
        __props__.__dict__["disk_format"] = disk_format
        __props__.__dict__["enterprise_project_id"] = enterprise_project_id
        __props__.__dict__["image_size"] = image_size
        __props__.__dict__["image_url"] = image_url
        __props__.__dict__["instance_id"] = instance_id
        __props__.__dict__["is_config"] = is_config
        __props__.__dict__["max_ram"] = max_ram
        __props__.__dict__["min_disk"] = min_disk
        __props__.__dict__["min_ram"] = min_ram
        __props__.__dict__["name"] = name
        __props__.__dict__["os_version"] = os_version
        __props__.__dict__["region"] = region
        __props__.__dict__["status"] = status
        __props__.__dict__["tags"] = tags
        __props__.__dict__["type"] = type
        __props__.__dict__["vault_id"] = vault_id
        __props__.__dict__["visibility"] = visibility
        __props__.__dict__["volume_id"] = volume_id
        return ImagesImage(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="backupId")
    def backup_id(self) -> pulumi.Output[builtins.str]:
        return pulumi.get(self, "backup_id")

    @property
    @pulumi.getter
    def checksum(self) -> pulumi.Output[builtins.str]:
        return pulumi.get(self, "checksum")

    @property
    @pulumi.getter(name="cmkId")
    def cmk_id(self) -> pulumi.Output[Optional[builtins.str]]:
        return pulumi.get(self, "cmk_id")

    @property
    @pulumi.getter(name="dataOrigin")
    def data_origin(self) -> pulumi.Output[builtins.str]:
        return pulumi.get(self, "data_origin")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[builtins.str]]:
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="diskFormat")
    def disk_format(self) -> pulumi.Output[builtins.str]:
        return pulumi.get(self, "disk_format")

    @property
    @pulumi.getter(name="enterpriseProjectId")
    def enterprise_project_id(self) -> pulumi.Output[builtins.str]:
        return pulumi.get(self, "enterprise_project_id")

    @property
    @pulumi.getter(name="imageSize")
    def image_size(self) -> pulumi.Output[builtins.str]:
        return pulumi.get(self, "image_size")

    @property
    @pulumi.getter(name="imageUrl")
    def image_url(self) -> pulumi.Output[Optional[builtins.str]]:
        return pulumi.get(self, "image_url")

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> pulumi.Output[builtins.str]:
        return pulumi.get(self, "instance_id")

    @property
    @pulumi.getter(name="isConfig")
    def is_config(self) -> pulumi.Output[Optional[builtins.bool]]:
        return pulumi.get(self, "is_config")

    @property
    @pulumi.getter(name="maxRam")
    def max_ram(self) -> pulumi.Output[builtins.int]:
        return pulumi.get(self, "max_ram")

    @property
    @pulumi.getter(name="minDisk")
    def min_disk(self) -> pulumi.Output[Optional[builtins.int]]:
        return pulumi.get(self, "min_disk")

    @property
    @pulumi.getter(name="minRam")
    def min_ram(self) -> pulumi.Output[builtins.int]:
        return pulumi.get(self, "min_ram")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[builtins.str]:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="osVersion")
    def os_version(self) -> pulumi.Output[builtins.str]:
        return pulumi.get(self, "os_version")

    @property
    @pulumi.getter
    def region(self) -> pulumi.Output[builtins.str]:
        return pulumi.get(self, "region")

    @property
    @pulumi.getter
    def status(self) -> pulumi.Output[builtins.str]:
        return pulumi.get(self, "status")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Mapping[str, builtins.str]]]:
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter
    def type(self) -> pulumi.Output[Optional[builtins.str]]:
        return pulumi.get(self, "type")

    @property
    @pulumi.getter(name="vaultId")
    def vault_id(self) -> pulumi.Output[Optional[builtins.str]]:
        return pulumi.get(self, "vault_id")

    @property
    @pulumi.getter
    def visibility(self) -> pulumi.Output[builtins.str]:
        return pulumi.get(self, "visibility")

    @property
    @pulumi.getter(name="volumeId")
    def volume_id(self) -> pulumi.Output[builtins.str]:
        return pulumi.get(self, "volume_id")

