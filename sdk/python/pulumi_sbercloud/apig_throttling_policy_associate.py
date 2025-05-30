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

__all__ = ['ApigThrottlingPolicyAssociateArgs', 'ApigThrottlingPolicyAssociate']

@pulumi.input_type
class ApigThrottlingPolicyAssociateArgs:
    def __init__(__self__, *,
                 instance_id: pulumi.Input[builtins.str],
                 policy_id: pulumi.Input[builtins.str],
                 publish_ids: pulumi.Input[Sequence[pulumi.Input[builtins.str]]],
                 region: Optional[pulumi.Input[builtins.str]] = None):
        """
        The set of arguments for constructing a ApigThrottlingPolicyAssociate resource.
        :param pulumi.Input[builtins.str] instance_id: The ID of the dedicated instance to which the APIs and the throttling policy belongs.
        :param pulumi.Input[builtins.str] policy_id: The ID of the throttling policy.
        :param pulumi.Input[Sequence[pulumi.Input[builtins.str]]] publish_ids: The publish IDs corresponding to the APIs bound by the throttling policy.
        :param pulumi.Input[builtins.str] region: The region where the dedicated instance and the throttling policy are located.
        """
        pulumi.set(__self__, "instance_id", instance_id)
        pulumi.set(__self__, "policy_id", policy_id)
        pulumi.set(__self__, "publish_ids", publish_ids)
        if region is not None:
            pulumi.set(__self__, "region", region)

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> pulumi.Input[builtins.str]:
        """
        The ID of the dedicated instance to which the APIs and the throttling policy belongs.
        """
        return pulumi.get(self, "instance_id")

    @instance_id.setter
    def instance_id(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "instance_id", value)

    @property
    @pulumi.getter(name="policyId")
    def policy_id(self) -> pulumi.Input[builtins.str]:
        """
        The ID of the throttling policy.
        """
        return pulumi.get(self, "policy_id")

    @policy_id.setter
    def policy_id(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "policy_id", value)

    @property
    @pulumi.getter(name="publishIds")
    def publish_ids(self) -> pulumi.Input[Sequence[pulumi.Input[builtins.str]]]:
        """
        The publish IDs corresponding to the APIs bound by the throttling policy.
        """
        return pulumi.get(self, "publish_ids")

    @publish_ids.setter
    def publish_ids(self, value: pulumi.Input[Sequence[pulumi.Input[builtins.str]]]):
        pulumi.set(self, "publish_ids", value)

    @property
    @pulumi.getter
    def region(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The region where the dedicated instance and the throttling policy are located.
        """
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "region", value)


@pulumi.input_type
class _ApigThrottlingPolicyAssociateState:
    def __init__(__self__, *,
                 instance_id: Optional[pulumi.Input[builtins.str]] = None,
                 policy_id: Optional[pulumi.Input[builtins.str]] = None,
                 publish_ids: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]] = None,
                 region: Optional[pulumi.Input[builtins.str]] = None):
        """
        Input properties used for looking up and filtering ApigThrottlingPolicyAssociate resources.
        :param pulumi.Input[builtins.str] instance_id: The ID of the dedicated instance to which the APIs and the throttling policy belongs.
        :param pulumi.Input[builtins.str] policy_id: The ID of the throttling policy.
        :param pulumi.Input[Sequence[pulumi.Input[builtins.str]]] publish_ids: The publish IDs corresponding to the APIs bound by the throttling policy.
        :param pulumi.Input[builtins.str] region: The region where the dedicated instance and the throttling policy are located.
        """
        if instance_id is not None:
            pulumi.set(__self__, "instance_id", instance_id)
        if policy_id is not None:
            pulumi.set(__self__, "policy_id", policy_id)
        if publish_ids is not None:
            pulumi.set(__self__, "publish_ids", publish_ids)
        if region is not None:
            pulumi.set(__self__, "region", region)

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The ID of the dedicated instance to which the APIs and the throttling policy belongs.
        """
        return pulumi.get(self, "instance_id")

    @instance_id.setter
    def instance_id(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "instance_id", value)

    @property
    @pulumi.getter(name="policyId")
    def policy_id(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The ID of the throttling policy.
        """
        return pulumi.get(self, "policy_id")

    @policy_id.setter
    def policy_id(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "policy_id", value)

    @property
    @pulumi.getter(name="publishIds")
    def publish_ids(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]]:
        """
        The publish IDs corresponding to the APIs bound by the throttling policy.
        """
        return pulumi.get(self, "publish_ids")

    @publish_ids.setter
    def publish_ids(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]]):
        pulumi.set(self, "publish_ids", value)

    @property
    @pulumi.getter
    def region(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The region where the dedicated instance and the throttling policy are located.
        """
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "region", value)


class ApigThrottlingPolicyAssociate(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 instance_id: Optional[pulumi.Input[builtins.str]] = None,
                 policy_id: Optional[pulumi.Input[builtins.str]] = None,
                 publish_ids: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]] = None,
                 region: Optional[pulumi.Input[builtins.str]] = None,
                 __props__=None):
        """
        Create a ApigThrottlingPolicyAssociate resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[builtins.str] instance_id: The ID of the dedicated instance to which the APIs and the throttling policy belongs.
        :param pulumi.Input[builtins.str] policy_id: The ID of the throttling policy.
        :param pulumi.Input[Sequence[pulumi.Input[builtins.str]]] publish_ids: The publish IDs corresponding to the APIs bound by the throttling policy.
        :param pulumi.Input[builtins.str] region: The region where the dedicated instance and the throttling policy are located.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ApigThrottlingPolicyAssociateArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a ApigThrottlingPolicyAssociate resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param ApigThrottlingPolicyAssociateArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ApigThrottlingPolicyAssociateArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 instance_id: Optional[pulumi.Input[builtins.str]] = None,
                 policy_id: Optional[pulumi.Input[builtins.str]] = None,
                 publish_ids: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]] = None,
                 region: Optional[pulumi.Input[builtins.str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ApigThrottlingPolicyAssociateArgs.__new__(ApigThrottlingPolicyAssociateArgs)

            if instance_id is None and not opts.urn:
                raise TypeError("Missing required property 'instance_id'")
            __props__.__dict__["instance_id"] = instance_id
            if policy_id is None and not opts.urn:
                raise TypeError("Missing required property 'policy_id'")
            __props__.__dict__["policy_id"] = policy_id
            if publish_ids is None and not opts.urn:
                raise TypeError("Missing required property 'publish_ids'")
            __props__.__dict__["publish_ids"] = publish_ids
            __props__.__dict__["region"] = region
        super(ApigThrottlingPolicyAssociate, __self__).__init__(
            'sbercloud:index/apigThrottlingPolicyAssociate:ApigThrottlingPolicyAssociate',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            instance_id: Optional[pulumi.Input[builtins.str]] = None,
            policy_id: Optional[pulumi.Input[builtins.str]] = None,
            publish_ids: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]] = None,
            region: Optional[pulumi.Input[builtins.str]] = None) -> 'ApigThrottlingPolicyAssociate':
        """
        Get an existing ApigThrottlingPolicyAssociate resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[builtins.str] instance_id: The ID of the dedicated instance to which the APIs and the throttling policy belongs.
        :param pulumi.Input[builtins.str] policy_id: The ID of the throttling policy.
        :param pulumi.Input[Sequence[pulumi.Input[builtins.str]]] publish_ids: The publish IDs corresponding to the APIs bound by the throttling policy.
        :param pulumi.Input[builtins.str] region: The region where the dedicated instance and the throttling policy are located.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _ApigThrottlingPolicyAssociateState.__new__(_ApigThrottlingPolicyAssociateState)

        __props__.__dict__["instance_id"] = instance_id
        __props__.__dict__["policy_id"] = policy_id
        __props__.__dict__["publish_ids"] = publish_ids
        __props__.__dict__["region"] = region
        return ApigThrottlingPolicyAssociate(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> pulumi.Output[builtins.str]:
        """
        The ID of the dedicated instance to which the APIs and the throttling policy belongs.
        """
        return pulumi.get(self, "instance_id")

    @property
    @pulumi.getter(name="policyId")
    def policy_id(self) -> pulumi.Output[builtins.str]:
        """
        The ID of the throttling policy.
        """
        return pulumi.get(self, "policy_id")

    @property
    @pulumi.getter(name="publishIds")
    def publish_ids(self) -> pulumi.Output[Sequence[builtins.str]]:
        """
        The publish IDs corresponding to the APIs bound by the throttling policy.
        """
        return pulumi.get(self, "publish_ids")

    @property
    @pulumi.getter
    def region(self) -> pulumi.Output[builtins.str]:
        """
        The region where the dedicated instance and the throttling policy are located.
        """
        return pulumi.get(self, "region")

