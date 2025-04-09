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

__all__ = ['DliSparkJobArgs', 'DliSparkJob']

@pulumi.input_type
class DliSparkJobArgs:
    def __init__(__self__, *,
                 app_name: pulumi.Input[builtins.str],
                 queue_name: pulumi.Input[builtins.str],
                 app_parameters: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]] = None,
                 configurations: Optional[pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]]] = None,
                 dependent_packages: Optional[pulumi.Input[Sequence[pulumi.Input['DliSparkJobDependentPackageArgs']]]] = None,
                 driver_cores: Optional[pulumi.Input[builtins.int]] = None,
                 driver_memory: Optional[pulumi.Input[builtins.str]] = None,
                 executor_cores: Optional[pulumi.Input[builtins.int]] = None,
                 executor_memory: Optional[pulumi.Input[builtins.str]] = None,
                 executors: Optional[pulumi.Input[builtins.int]] = None,
                 feature: Optional[pulumi.Input[builtins.str]] = None,
                 files: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]] = None,
                 jars: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]] = None,
                 main_class: Optional[pulumi.Input[builtins.str]] = None,
                 max_retries: Optional[pulumi.Input[builtins.int]] = None,
                 modules: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]] = None,
                 name: Optional[pulumi.Input[builtins.str]] = None,
                 python_files: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]] = None,
                 region: Optional[pulumi.Input[builtins.str]] = None,
                 spark_version: Optional[pulumi.Input[builtins.str]] = None,
                 specification: Optional[pulumi.Input[builtins.str]] = None):
        """
        The set of arguments for constructing a DliSparkJob resource.
        """
        pulumi.set(__self__, "app_name", app_name)
        pulumi.set(__self__, "queue_name", queue_name)
        if app_parameters is not None:
            pulumi.set(__self__, "app_parameters", app_parameters)
        if configurations is not None:
            pulumi.set(__self__, "configurations", configurations)
        if dependent_packages is not None:
            pulumi.set(__self__, "dependent_packages", dependent_packages)
        if driver_cores is not None:
            pulumi.set(__self__, "driver_cores", driver_cores)
        if driver_memory is not None:
            pulumi.set(__self__, "driver_memory", driver_memory)
        if executor_cores is not None:
            pulumi.set(__self__, "executor_cores", executor_cores)
        if executor_memory is not None:
            pulumi.set(__self__, "executor_memory", executor_memory)
        if executors is not None:
            pulumi.set(__self__, "executors", executors)
        if feature is not None:
            pulumi.set(__self__, "feature", feature)
        if files is not None:
            pulumi.set(__self__, "files", files)
        if jars is not None:
            pulumi.set(__self__, "jars", jars)
        if main_class is not None:
            pulumi.set(__self__, "main_class", main_class)
        if max_retries is not None:
            pulumi.set(__self__, "max_retries", max_retries)
        if modules is not None:
            pulumi.set(__self__, "modules", modules)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if python_files is not None:
            pulumi.set(__self__, "python_files", python_files)
        if region is not None:
            pulumi.set(__self__, "region", region)
        if spark_version is not None:
            pulumi.set(__self__, "spark_version", spark_version)
        if specification is not None:
            pulumi.set(__self__, "specification", specification)

    @property
    @pulumi.getter(name="appName")
    def app_name(self) -> pulumi.Input[builtins.str]:
        return pulumi.get(self, "app_name")

    @app_name.setter
    def app_name(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "app_name", value)

    @property
    @pulumi.getter(name="queueName")
    def queue_name(self) -> pulumi.Input[builtins.str]:
        return pulumi.get(self, "queue_name")

    @queue_name.setter
    def queue_name(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "queue_name", value)

    @property
    @pulumi.getter(name="appParameters")
    def app_parameters(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]]:
        return pulumi.get(self, "app_parameters")

    @app_parameters.setter
    def app_parameters(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]]):
        pulumi.set(self, "app_parameters", value)

    @property
    @pulumi.getter
    def configurations(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]]]:
        return pulumi.get(self, "configurations")

    @configurations.setter
    def configurations(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]]]):
        pulumi.set(self, "configurations", value)

    @property
    @pulumi.getter(name="dependentPackages")
    def dependent_packages(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['DliSparkJobDependentPackageArgs']]]]:
        return pulumi.get(self, "dependent_packages")

    @dependent_packages.setter
    def dependent_packages(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['DliSparkJobDependentPackageArgs']]]]):
        pulumi.set(self, "dependent_packages", value)

    @property
    @pulumi.getter(name="driverCores")
    def driver_cores(self) -> Optional[pulumi.Input[builtins.int]]:
        return pulumi.get(self, "driver_cores")

    @driver_cores.setter
    def driver_cores(self, value: Optional[pulumi.Input[builtins.int]]):
        pulumi.set(self, "driver_cores", value)

    @property
    @pulumi.getter(name="driverMemory")
    def driver_memory(self) -> Optional[pulumi.Input[builtins.str]]:
        return pulumi.get(self, "driver_memory")

    @driver_memory.setter
    def driver_memory(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "driver_memory", value)

    @property
    @pulumi.getter(name="executorCores")
    def executor_cores(self) -> Optional[pulumi.Input[builtins.int]]:
        return pulumi.get(self, "executor_cores")

    @executor_cores.setter
    def executor_cores(self, value: Optional[pulumi.Input[builtins.int]]):
        pulumi.set(self, "executor_cores", value)

    @property
    @pulumi.getter(name="executorMemory")
    def executor_memory(self) -> Optional[pulumi.Input[builtins.str]]:
        return pulumi.get(self, "executor_memory")

    @executor_memory.setter
    def executor_memory(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "executor_memory", value)

    @property
    @pulumi.getter
    def executors(self) -> Optional[pulumi.Input[builtins.int]]:
        return pulumi.get(self, "executors")

    @executors.setter
    def executors(self, value: Optional[pulumi.Input[builtins.int]]):
        pulumi.set(self, "executors", value)

    @property
    @pulumi.getter
    def feature(self) -> Optional[pulumi.Input[builtins.str]]:
        return pulumi.get(self, "feature")

    @feature.setter
    def feature(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "feature", value)

    @property
    @pulumi.getter
    def files(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]]:
        return pulumi.get(self, "files")

    @files.setter
    def files(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]]):
        pulumi.set(self, "files", value)

    @property
    @pulumi.getter
    def jars(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]]:
        return pulumi.get(self, "jars")

    @jars.setter
    def jars(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]]):
        pulumi.set(self, "jars", value)

    @property
    @pulumi.getter(name="mainClass")
    def main_class(self) -> Optional[pulumi.Input[builtins.str]]:
        return pulumi.get(self, "main_class")

    @main_class.setter
    def main_class(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "main_class", value)

    @property
    @pulumi.getter(name="maxRetries")
    def max_retries(self) -> Optional[pulumi.Input[builtins.int]]:
        return pulumi.get(self, "max_retries")

    @max_retries.setter
    def max_retries(self, value: Optional[pulumi.Input[builtins.int]]):
        pulumi.set(self, "max_retries", value)

    @property
    @pulumi.getter
    def modules(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]]:
        return pulumi.get(self, "modules")

    @modules.setter
    def modules(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]]):
        pulumi.set(self, "modules", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[builtins.str]]:
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="pythonFiles")
    def python_files(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]]:
        return pulumi.get(self, "python_files")

    @python_files.setter
    def python_files(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]]):
        pulumi.set(self, "python_files", value)

    @property
    @pulumi.getter
    def region(self) -> Optional[pulumi.Input[builtins.str]]:
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "region", value)

    @property
    @pulumi.getter(name="sparkVersion")
    def spark_version(self) -> Optional[pulumi.Input[builtins.str]]:
        return pulumi.get(self, "spark_version")

    @spark_version.setter
    def spark_version(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "spark_version", value)

    @property
    @pulumi.getter
    def specification(self) -> Optional[pulumi.Input[builtins.str]]:
        return pulumi.get(self, "specification")

    @specification.setter
    def specification(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "specification", value)


@pulumi.input_type
class _DliSparkJobState:
    def __init__(__self__, *,
                 app_name: Optional[pulumi.Input[builtins.str]] = None,
                 app_parameters: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]] = None,
                 configurations: Optional[pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]]] = None,
                 created_at: Optional[pulumi.Input[builtins.str]] = None,
                 dependent_packages: Optional[pulumi.Input[Sequence[pulumi.Input['DliSparkJobDependentPackageArgs']]]] = None,
                 driver_cores: Optional[pulumi.Input[builtins.int]] = None,
                 driver_memory: Optional[pulumi.Input[builtins.str]] = None,
                 executor_cores: Optional[pulumi.Input[builtins.int]] = None,
                 executor_memory: Optional[pulumi.Input[builtins.str]] = None,
                 executors: Optional[pulumi.Input[builtins.int]] = None,
                 feature: Optional[pulumi.Input[builtins.str]] = None,
                 files: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]] = None,
                 jars: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]] = None,
                 main_class: Optional[pulumi.Input[builtins.str]] = None,
                 max_retries: Optional[pulumi.Input[builtins.int]] = None,
                 modules: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]] = None,
                 name: Optional[pulumi.Input[builtins.str]] = None,
                 owner: Optional[pulumi.Input[builtins.str]] = None,
                 python_files: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]] = None,
                 queue_name: Optional[pulumi.Input[builtins.str]] = None,
                 region: Optional[pulumi.Input[builtins.str]] = None,
                 spark_version: Optional[pulumi.Input[builtins.str]] = None,
                 specification: Optional[pulumi.Input[builtins.str]] = None):
        """
        Input properties used for looking up and filtering DliSparkJob resources.
        """
        if app_name is not None:
            pulumi.set(__self__, "app_name", app_name)
        if app_parameters is not None:
            pulumi.set(__self__, "app_parameters", app_parameters)
        if configurations is not None:
            pulumi.set(__self__, "configurations", configurations)
        if created_at is not None:
            pulumi.set(__self__, "created_at", created_at)
        if dependent_packages is not None:
            pulumi.set(__self__, "dependent_packages", dependent_packages)
        if driver_cores is not None:
            pulumi.set(__self__, "driver_cores", driver_cores)
        if driver_memory is not None:
            pulumi.set(__self__, "driver_memory", driver_memory)
        if executor_cores is not None:
            pulumi.set(__self__, "executor_cores", executor_cores)
        if executor_memory is not None:
            pulumi.set(__self__, "executor_memory", executor_memory)
        if executors is not None:
            pulumi.set(__self__, "executors", executors)
        if feature is not None:
            pulumi.set(__self__, "feature", feature)
        if files is not None:
            pulumi.set(__self__, "files", files)
        if jars is not None:
            pulumi.set(__self__, "jars", jars)
        if main_class is not None:
            pulumi.set(__self__, "main_class", main_class)
        if max_retries is not None:
            pulumi.set(__self__, "max_retries", max_retries)
        if modules is not None:
            pulumi.set(__self__, "modules", modules)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if owner is not None:
            pulumi.set(__self__, "owner", owner)
        if python_files is not None:
            pulumi.set(__self__, "python_files", python_files)
        if queue_name is not None:
            pulumi.set(__self__, "queue_name", queue_name)
        if region is not None:
            pulumi.set(__self__, "region", region)
        if spark_version is not None:
            pulumi.set(__self__, "spark_version", spark_version)
        if specification is not None:
            pulumi.set(__self__, "specification", specification)

    @property
    @pulumi.getter(name="appName")
    def app_name(self) -> Optional[pulumi.Input[builtins.str]]:
        return pulumi.get(self, "app_name")

    @app_name.setter
    def app_name(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "app_name", value)

    @property
    @pulumi.getter(name="appParameters")
    def app_parameters(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]]:
        return pulumi.get(self, "app_parameters")

    @app_parameters.setter
    def app_parameters(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]]):
        pulumi.set(self, "app_parameters", value)

    @property
    @pulumi.getter
    def configurations(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]]]:
        return pulumi.get(self, "configurations")

    @configurations.setter
    def configurations(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]]]):
        pulumi.set(self, "configurations", value)

    @property
    @pulumi.getter(name="createdAt")
    def created_at(self) -> Optional[pulumi.Input[builtins.str]]:
        return pulumi.get(self, "created_at")

    @created_at.setter
    def created_at(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "created_at", value)

    @property
    @pulumi.getter(name="dependentPackages")
    def dependent_packages(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['DliSparkJobDependentPackageArgs']]]]:
        return pulumi.get(self, "dependent_packages")

    @dependent_packages.setter
    def dependent_packages(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['DliSparkJobDependentPackageArgs']]]]):
        pulumi.set(self, "dependent_packages", value)

    @property
    @pulumi.getter(name="driverCores")
    def driver_cores(self) -> Optional[pulumi.Input[builtins.int]]:
        return pulumi.get(self, "driver_cores")

    @driver_cores.setter
    def driver_cores(self, value: Optional[pulumi.Input[builtins.int]]):
        pulumi.set(self, "driver_cores", value)

    @property
    @pulumi.getter(name="driverMemory")
    def driver_memory(self) -> Optional[pulumi.Input[builtins.str]]:
        return pulumi.get(self, "driver_memory")

    @driver_memory.setter
    def driver_memory(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "driver_memory", value)

    @property
    @pulumi.getter(name="executorCores")
    def executor_cores(self) -> Optional[pulumi.Input[builtins.int]]:
        return pulumi.get(self, "executor_cores")

    @executor_cores.setter
    def executor_cores(self, value: Optional[pulumi.Input[builtins.int]]):
        pulumi.set(self, "executor_cores", value)

    @property
    @pulumi.getter(name="executorMemory")
    def executor_memory(self) -> Optional[pulumi.Input[builtins.str]]:
        return pulumi.get(self, "executor_memory")

    @executor_memory.setter
    def executor_memory(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "executor_memory", value)

    @property
    @pulumi.getter
    def executors(self) -> Optional[pulumi.Input[builtins.int]]:
        return pulumi.get(self, "executors")

    @executors.setter
    def executors(self, value: Optional[pulumi.Input[builtins.int]]):
        pulumi.set(self, "executors", value)

    @property
    @pulumi.getter
    def feature(self) -> Optional[pulumi.Input[builtins.str]]:
        return pulumi.get(self, "feature")

    @feature.setter
    def feature(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "feature", value)

    @property
    @pulumi.getter
    def files(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]]:
        return pulumi.get(self, "files")

    @files.setter
    def files(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]]):
        pulumi.set(self, "files", value)

    @property
    @pulumi.getter
    def jars(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]]:
        return pulumi.get(self, "jars")

    @jars.setter
    def jars(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]]):
        pulumi.set(self, "jars", value)

    @property
    @pulumi.getter(name="mainClass")
    def main_class(self) -> Optional[pulumi.Input[builtins.str]]:
        return pulumi.get(self, "main_class")

    @main_class.setter
    def main_class(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "main_class", value)

    @property
    @pulumi.getter(name="maxRetries")
    def max_retries(self) -> Optional[pulumi.Input[builtins.int]]:
        return pulumi.get(self, "max_retries")

    @max_retries.setter
    def max_retries(self, value: Optional[pulumi.Input[builtins.int]]):
        pulumi.set(self, "max_retries", value)

    @property
    @pulumi.getter
    def modules(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]]:
        return pulumi.get(self, "modules")

    @modules.setter
    def modules(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]]):
        pulumi.set(self, "modules", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[builtins.str]]:
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def owner(self) -> Optional[pulumi.Input[builtins.str]]:
        return pulumi.get(self, "owner")

    @owner.setter
    def owner(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "owner", value)

    @property
    @pulumi.getter(name="pythonFiles")
    def python_files(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]]:
        return pulumi.get(self, "python_files")

    @python_files.setter
    def python_files(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]]):
        pulumi.set(self, "python_files", value)

    @property
    @pulumi.getter(name="queueName")
    def queue_name(self) -> Optional[pulumi.Input[builtins.str]]:
        return pulumi.get(self, "queue_name")

    @queue_name.setter
    def queue_name(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "queue_name", value)

    @property
    @pulumi.getter
    def region(self) -> Optional[pulumi.Input[builtins.str]]:
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "region", value)

    @property
    @pulumi.getter(name="sparkVersion")
    def spark_version(self) -> Optional[pulumi.Input[builtins.str]]:
        return pulumi.get(self, "spark_version")

    @spark_version.setter
    def spark_version(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "spark_version", value)

    @property
    @pulumi.getter
    def specification(self) -> Optional[pulumi.Input[builtins.str]]:
        return pulumi.get(self, "specification")

    @specification.setter
    def specification(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "specification", value)


class DliSparkJob(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 app_name: Optional[pulumi.Input[builtins.str]] = None,
                 app_parameters: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]] = None,
                 configurations: Optional[pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]]] = None,
                 dependent_packages: Optional[pulumi.Input[Sequence[pulumi.Input[Union['DliSparkJobDependentPackageArgs', 'DliSparkJobDependentPackageArgsDict']]]]] = None,
                 driver_cores: Optional[pulumi.Input[builtins.int]] = None,
                 driver_memory: Optional[pulumi.Input[builtins.str]] = None,
                 executor_cores: Optional[pulumi.Input[builtins.int]] = None,
                 executor_memory: Optional[pulumi.Input[builtins.str]] = None,
                 executors: Optional[pulumi.Input[builtins.int]] = None,
                 feature: Optional[pulumi.Input[builtins.str]] = None,
                 files: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]] = None,
                 jars: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]] = None,
                 main_class: Optional[pulumi.Input[builtins.str]] = None,
                 max_retries: Optional[pulumi.Input[builtins.int]] = None,
                 modules: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]] = None,
                 name: Optional[pulumi.Input[builtins.str]] = None,
                 python_files: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]] = None,
                 queue_name: Optional[pulumi.Input[builtins.str]] = None,
                 region: Optional[pulumi.Input[builtins.str]] = None,
                 spark_version: Optional[pulumi.Input[builtins.str]] = None,
                 specification: Optional[pulumi.Input[builtins.str]] = None,
                 __props__=None):
        """
        Create a DliSparkJob resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: DliSparkJobArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a DliSparkJob resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param DliSparkJobArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(DliSparkJobArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 app_name: Optional[pulumi.Input[builtins.str]] = None,
                 app_parameters: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]] = None,
                 configurations: Optional[pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]]] = None,
                 dependent_packages: Optional[pulumi.Input[Sequence[pulumi.Input[Union['DliSparkJobDependentPackageArgs', 'DliSparkJobDependentPackageArgsDict']]]]] = None,
                 driver_cores: Optional[pulumi.Input[builtins.int]] = None,
                 driver_memory: Optional[pulumi.Input[builtins.str]] = None,
                 executor_cores: Optional[pulumi.Input[builtins.int]] = None,
                 executor_memory: Optional[pulumi.Input[builtins.str]] = None,
                 executors: Optional[pulumi.Input[builtins.int]] = None,
                 feature: Optional[pulumi.Input[builtins.str]] = None,
                 files: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]] = None,
                 jars: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]] = None,
                 main_class: Optional[pulumi.Input[builtins.str]] = None,
                 max_retries: Optional[pulumi.Input[builtins.int]] = None,
                 modules: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]] = None,
                 name: Optional[pulumi.Input[builtins.str]] = None,
                 python_files: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]] = None,
                 queue_name: Optional[pulumi.Input[builtins.str]] = None,
                 region: Optional[pulumi.Input[builtins.str]] = None,
                 spark_version: Optional[pulumi.Input[builtins.str]] = None,
                 specification: Optional[pulumi.Input[builtins.str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = DliSparkJobArgs.__new__(DliSparkJobArgs)

            if app_name is None and not opts.urn:
                raise TypeError("Missing required property 'app_name'")
            __props__.__dict__["app_name"] = app_name
            __props__.__dict__["app_parameters"] = app_parameters
            __props__.__dict__["configurations"] = configurations
            __props__.__dict__["dependent_packages"] = dependent_packages
            __props__.__dict__["driver_cores"] = driver_cores
            __props__.__dict__["driver_memory"] = driver_memory
            __props__.__dict__["executor_cores"] = executor_cores
            __props__.__dict__["executor_memory"] = executor_memory
            __props__.__dict__["executors"] = executors
            __props__.__dict__["feature"] = feature
            __props__.__dict__["files"] = files
            __props__.__dict__["jars"] = jars
            __props__.__dict__["main_class"] = main_class
            __props__.__dict__["max_retries"] = max_retries
            __props__.__dict__["modules"] = modules
            __props__.__dict__["name"] = name
            __props__.__dict__["python_files"] = python_files
            if queue_name is None and not opts.urn:
                raise TypeError("Missing required property 'queue_name'")
            __props__.__dict__["queue_name"] = queue_name
            __props__.__dict__["region"] = region
            __props__.__dict__["spark_version"] = spark_version
            __props__.__dict__["specification"] = specification
            __props__.__dict__["created_at"] = None
            __props__.__dict__["owner"] = None
        super(DliSparkJob, __self__).__init__(
            'sbercloud:index/dliSparkJob:DliSparkJob',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            app_name: Optional[pulumi.Input[builtins.str]] = None,
            app_parameters: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]] = None,
            configurations: Optional[pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]]] = None,
            created_at: Optional[pulumi.Input[builtins.str]] = None,
            dependent_packages: Optional[pulumi.Input[Sequence[pulumi.Input[Union['DliSparkJobDependentPackageArgs', 'DliSparkJobDependentPackageArgsDict']]]]] = None,
            driver_cores: Optional[pulumi.Input[builtins.int]] = None,
            driver_memory: Optional[pulumi.Input[builtins.str]] = None,
            executor_cores: Optional[pulumi.Input[builtins.int]] = None,
            executor_memory: Optional[pulumi.Input[builtins.str]] = None,
            executors: Optional[pulumi.Input[builtins.int]] = None,
            feature: Optional[pulumi.Input[builtins.str]] = None,
            files: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]] = None,
            jars: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]] = None,
            main_class: Optional[pulumi.Input[builtins.str]] = None,
            max_retries: Optional[pulumi.Input[builtins.int]] = None,
            modules: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]] = None,
            name: Optional[pulumi.Input[builtins.str]] = None,
            owner: Optional[pulumi.Input[builtins.str]] = None,
            python_files: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]] = None,
            queue_name: Optional[pulumi.Input[builtins.str]] = None,
            region: Optional[pulumi.Input[builtins.str]] = None,
            spark_version: Optional[pulumi.Input[builtins.str]] = None,
            specification: Optional[pulumi.Input[builtins.str]] = None) -> 'DliSparkJob':
        """
        Get an existing DliSparkJob resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _DliSparkJobState.__new__(_DliSparkJobState)

        __props__.__dict__["app_name"] = app_name
        __props__.__dict__["app_parameters"] = app_parameters
        __props__.__dict__["configurations"] = configurations
        __props__.__dict__["created_at"] = created_at
        __props__.__dict__["dependent_packages"] = dependent_packages
        __props__.__dict__["driver_cores"] = driver_cores
        __props__.__dict__["driver_memory"] = driver_memory
        __props__.__dict__["executor_cores"] = executor_cores
        __props__.__dict__["executor_memory"] = executor_memory
        __props__.__dict__["executors"] = executors
        __props__.__dict__["feature"] = feature
        __props__.__dict__["files"] = files
        __props__.__dict__["jars"] = jars
        __props__.__dict__["main_class"] = main_class
        __props__.__dict__["max_retries"] = max_retries
        __props__.__dict__["modules"] = modules
        __props__.__dict__["name"] = name
        __props__.__dict__["owner"] = owner
        __props__.__dict__["python_files"] = python_files
        __props__.__dict__["queue_name"] = queue_name
        __props__.__dict__["region"] = region
        __props__.__dict__["spark_version"] = spark_version
        __props__.__dict__["specification"] = specification
        return DliSparkJob(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="appName")
    def app_name(self) -> pulumi.Output[builtins.str]:
        return pulumi.get(self, "app_name")

    @property
    @pulumi.getter(name="appParameters")
    def app_parameters(self) -> pulumi.Output[Optional[Sequence[builtins.str]]]:
        return pulumi.get(self, "app_parameters")

    @property
    @pulumi.getter
    def configurations(self) -> pulumi.Output[Optional[Mapping[str, builtins.str]]]:
        return pulumi.get(self, "configurations")

    @property
    @pulumi.getter(name="createdAt")
    def created_at(self) -> pulumi.Output[builtins.str]:
        return pulumi.get(self, "created_at")

    @property
    @pulumi.getter(name="dependentPackages")
    def dependent_packages(self) -> pulumi.Output[Optional[Sequence['outputs.DliSparkJobDependentPackage']]]:
        return pulumi.get(self, "dependent_packages")

    @property
    @pulumi.getter(name="driverCores")
    def driver_cores(self) -> pulumi.Output[Optional[builtins.int]]:
        return pulumi.get(self, "driver_cores")

    @property
    @pulumi.getter(name="driverMemory")
    def driver_memory(self) -> pulumi.Output[Optional[builtins.str]]:
        return pulumi.get(self, "driver_memory")

    @property
    @pulumi.getter(name="executorCores")
    def executor_cores(self) -> pulumi.Output[Optional[builtins.int]]:
        return pulumi.get(self, "executor_cores")

    @property
    @pulumi.getter(name="executorMemory")
    def executor_memory(self) -> pulumi.Output[Optional[builtins.str]]:
        return pulumi.get(self, "executor_memory")

    @property
    @pulumi.getter
    def executors(self) -> pulumi.Output[Optional[builtins.int]]:
        return pulumi.get(self, "executors")

    @property
    @pulumi.getter
    def feature(self) -> pulumi.Output[Optional[builtins.str]]:
        return pulumi.get(self, "feature")

    @property
    @pulumi.getter
    def files(self) -> pulumi.Output[Optional[Sequence[builtins.str]]]:
        return pulumi.get(self, "files")

    @property
    @pulumi.getter
    def jars(self) -> pulumi.Output[Optional[Sequence[builtins.str]]]:
        return pulumi.get(self, "jars")

    @property
    @pulumi.getter(name="mainClass")
    def main_class(self) -> pulumi.Output[Optional[builtins.str]]:
        return pulumi.get(self, "main_class")

    @property
    @pulumi.getter(name="maxRetries")
    def max_retries(self) -> pulumi.Output[Optional[builtins.int]]:
        return pulumi.get(self, "max_retries")

    @property
    @pulumi.getter
    def modules(self) -> pulumi.Output[Optional[Sequence[builtins.str]]]:
        return pulumi.get(self, "modules")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[builtins.str]:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def owner(self) -> pulumi.Output[builtins.str]:
        return pulumi.get(self, "owner")

    @property
    @pulumi.getter(name="pythonFiles")
    def python_files(self) -> pulumi.Output[Optional[Sequence[builtins.str]]]:
        return pulumi.get(self, "python_files")

    @property
    @pulumi.getter(name="queueName")
    def queue_name(self) -> pulumi.Output[builtins.str]:
        return pulumi.get(self, "queue_name")

    @property
    @pulumi.getter
    def region(self) -> pulumi.Output[builtins.str]:
        return pulumi.get(self, "region")

    @property
    @pulumi.getter(name="sparkVersion")
    def spark_version(self) -> pulumi.Output[Optional[builtins.str]]:
        return pulumi.get(self, "spark_version")

    @property
    @pulumi.getter
    def specification(self) -> pulumi.Output[Optional[builtins.str]]:
        return pulumi.get(self, "specification")

