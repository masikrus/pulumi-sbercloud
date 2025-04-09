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

__all__ = ['DcsRestoreArgs', 'DcsRestore']

@pulumi.input_type
class DcsRestoreArgs:
    def __init__(__self__, *,
                 backup_id: pulumi.Input[builtins.str],
                 instance_id: pulumi.Input[builtins.str],
                 project_id: pulumi.Input[builtins.str],
                 remark: Optional[pulumi.Input[builtins.str]] = None):
        """
        The set of arguments for constructing a DcsRestore resource.
        """
        pulumi.set(__self__, "backup_id", backup_id)
        pulumi.set(__self__, "instance_id", instance_id)
        pulumi.set(__self__, "project_id", project_id)
        if remark is not None:
            pulumi.set(__self__, "remark", remark)

    @property
    @pulumi.getter(name="backupId")
    def backup_id(self) -> pulumi.Input[builtins.str]:
        return pulumi.get(self, "backup_id")

    @backup_id.setter
    def backup_id(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "backup_id", value)

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> pulumi.Input[builtins.str]:
        return pulumi.get(self, "instance_id")

    @instance_id.setter
    def instance_id(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "instance_id", value)

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> pulumi.Input[builtins.str]:
        return pulumi.get(self, "project_id")

    @project_id.setter
    def project_id(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "project_id", value)

    @property
    @pulumi.getter
    def remark(self) -> Optional[pulumi.Input[builtins.str]]:
        return pulumi.get(self, "remark")

    @remark.setter
    def remark(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "remark", value)


@pulumi.input_type
class _DcsRestoreState:
    def __init__(__self__, *,
                 backup_id: Optional[pulumi.Input[builtins.str]] = None,
                 instance_id: Optional[pulumi.Input[builtins.str]] = None,
                 project_id: Optional[pulumi.Input[builtins.str]] = None,
                 remark: Optional[pulumi.Input[builtins.str]] = None,
                 restore_records: Optional[pulumi.Input[Sequence[pulumi.Input['DcsRestoreRestoreRecordArgs']]]] = None):
        """
        Input properties used for looking up and filtering DcsRestore resources.
        """
        if backup_id is not None:
            pulumi.set(__self__, "backup_id", backup_id)
        if instance_id is not None:
            pulumi.set(__self__, "instance_id", instance_id)
        if project_id is not None:
            pulumi.set(__self__, "project_id", project_id)
        if remark is not None:
            pulumi.set(__self__, "remark", remark)
        if restore_records is not None:
            pulumi.set(__self__, "restore_records", restore_records)

    @property
    @pulumi.getter(name="backupId")
    def backup_id(self) -> Optional[pulumi.Input[builtins.str]]:
        return pulumi.get(self, "backup_id")

    @backup_id.setter
    def backup_id(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "backup_id", value)

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> Optional[pulumi.Input[builtins.str]]:
        return pulumi.get(self, "instance_id")

    @instance_id.setter
    def instance_id(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "instance_id", value)

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> Optional[pulumi.Input[builtins.str]]:
        return pulumi.get(self, "project_id")

    @project_id.setter
    def project_id(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "project_id", value)

    @property
    @pulumi.getter
    def remark(self) -> Optional[pulumi.Input[builtins.str]]:
        return pulumi.get(self, "remark")

    @remark.setter
    def remark(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "remark", value)

    @property
    @pulumi.getter(name="restoreRecords")
    def restore_records(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['DcsRestoreRestoreRecordArgs']]]]:
        return pulumi.get(self, "restore_records")

    @restore_records.setter
    def restore_records(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['DcsRestoreRestoreRecordArgs']]]]):
        pulumi.set(self, "restore_records", value)


class DcsRestore(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 backup_id: Optional[pulumi.Input[builtins.str]] = None,
                 instance_id: Optional[pulumi.Input[builtins.str]] = None,
                 project_id: Optional[pulumi.Input[builtins.str]] = None,
                 remark: Optional[pulumi.Input[builtins.str]] = None,
                 __props__=None):
        """
        Create a DcsRestore resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: DcsRestoreArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a DcsRestore resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param DcsRestoreArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(DcsRestoreArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 backup_id: Optional[pulumi.Input[builtins.str]] = None,
                 instance_id: Optional[pulumi.Input[builtins.str]] = None,
                 project_id: Optional[pulumi.Input[builtins.str]] = None,
                 remark: Optional[pulumi.Input[builtins.str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = DcsRestoreArgs.__new__(DcsRestoreArgs)

            if backup_id is None and not opts.urn:
                raise TypeError("Missing required property 'backup_id'")
            __props__.__dict__["backup_id"] = backup_id
            if instance_id is None and not opts.urn:
                raise TypeError("Missing required property 'instance_id'")
            __props__.__dict__["instance_id"] = instance_id
            if project_id is None and not opts.urn:
                raise TypeError("Missing required property 'project_id'")
            __props__.__dict__["project_id"] = project_id
            __props__.__dict__["remark"] = remark
            __props__.__dict__["restore_records"] = None
        super(DcsRestore, __self__).__init__(
            'sbercloud:index/dcsRestore:DcsRestore',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            backup_id: Optional[pulumi.Input[builtins.str]] = None,
            instance_id: Optional[pulumi.Input[builtins.str]] = None,
            project_id: Optional[pulumi.Input[builtins.str]] = None,
            remark: Optional[pulumi.Input[builtins.str]] = None,
            restore_records: Optional[pulumi.Input[Sequence[pulumi.Input[Union['DcsRestoreRestoreRecordArgs', 'DcsRestoreRestoreRecordArgsDict']]]]] = None) -> 'DcsRestore':
        """
        Get an existing DcsRestore resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _DcsRestoreState.__new__(_DcsRestoreState)

        __props__.__dict__["backup_id"] = backup_id
        __props__.__dict__["instance_id"] = instance_id
        __props__.__dict__["project_id"] = project_id
        __props__.__dict__["remark"] = remark
        __props__.__dict__["restore_records"] = restore_records
        return DcsRestore(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="backupId")
    def backup_id(self) -> pulumi.Output[builtins.str]:
        return pulumi.get(self, "backup_id")

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> pulumi.Output[builtins.str]:
        return pulumi.get(self, "instance_id")

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> pulumi.Output[builtins.str]:
        return pulumi.get(self, "project_id")

    @property
    @pulumi.getter
    def remark(self) -> pulumi.Output[Optional[builtins.str]]:
        return pulumi.get(self, "remark")

    @property
    @pulumi.getter(name="restoreRecords")
    def restore_records(self) -> pulumi.Output[Sequence['outputs.DcsRestoreRestoreRecord']]:
        return pulumi.get(self, "restore_records")

