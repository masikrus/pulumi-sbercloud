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

__all__ = ['NatGatewayArgs', 'NatGateway']

@pulumi.input_type
class NatGatewayArgs:
    def __init__(__self__, *,
                 spec: pulumi.Input[builtins.str],
                 subnet_id: pulumi.Input[builtins.str],
                 vpc_id: pulumi.Input[builtins.str],
                 auto_renew: Optional[pulumi.Input[builtins.str]] = None,
                 charging_mode: Optional[pulumi.Input[builtins.str]] = None,
                 description: Optional[pulumi.Input[builtins.str]] = None,
                 enterprise_project_id: Optional[pulumi.Input[builtins.str]] = None,
                 name: Optional[pulumi.Input[builtins.str]] = None,
                 ngport_ip_address: Optional[pulumi.Input[builtins.str]] = None,
                 period: Optional[pulumi.Input[builtins.int]] = None,
                 period_unit: Optional[pulumi.Input[builtins.str]] = None,
                 region: Optional[pulumi.Input[builtins.str]] = None,
                 session_conf: Optional[pulumi.Input['NatGatewaySessionConfArgs']] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]]] = None):
        """
        The set of arguments for constructing a NatGateway resource.
        :param pulumi.Input[builtins.str] spec: The specification of the NAT gateway.
        :param pulumi.Input[builtins.str] subnet_id: The network ID of the downstream interface (the next hop of the DVR) of the NAT gateway.
        :param pulumi.Input[builtins.str] vpc_id: The ID of the VPC to which the NAT gateway belongs.
        :param pulumi.Input[builtins.str] description: The description of the NAT gateway.
        :param pulumi.Input[builtins.str] enterprise_project_id: The enterprise project ID of the NAT gateway.
        :param pulumi.Input[builtins.str] name: The NAT gateway name.
        :param pulumi.Input[builtins.str] ngport_ip_address: The IP address used for the NG port of the NAT gateway.
        :param pulumi.Input[builtins.str] region: The region where the NAT gateway is located.
        :param pulumi.Input['NatGatewaySessionConfArgs'] session_conf: The session configuration of the NAT gateway.
        """
        pulumi.set(__self__, "spec", spec)
        pulumi.set(__self__, "subnet_id", subnet_id)
        pulumi.set(__self__, "vpc_id", vpc_id)
        if auto_renew is not None:
            pulumi.set(__self__, "auto_renew", auto_renew)
        if charging_mode is not None:
            pulumi.set(__self__, "charging_mode", charging_mode)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if enterprise_project_id is not None:
            pulumi.set(__self__, "enterprise_project_id", enterprise_project_id)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if ngport_ip_address is not None:
            pulumi.set(__self__, "ngport_ip_address", ngport_ip_address)
        if period is not None:
            pulumi.set(__self__, "period", period)
        if period_unit is not None:
            pulumi.set(__self__, "period_unit", period_unit)
        if region is not None:
            pulumi.set(__self__, "region", region)
        if session_conf is not None:
            pulumi.set(__self__, "session_conf", session_conf)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter
    def spec(self) -> pulumi.Input[builtins.str]:
        """
        The specification of the NAT gateway.
        """
        return pulumi.get(self, "spec")

    @spec.setter
    def spec(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "spec", value)

    @property
    @pulumi.getter(name="subnetId")
    def subnet_id(self) -> pulumi.Input[builtins.str]:
        """
        The network ID of the downstream interface (the next hop of the DVR) of the NAT gateway.
        """
        return pulumi.get(self, "subnet_id")

    @subnet_id.setter
    def subnet_id(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "subnet_id", value)

    @property
    @pulumi.getter(name="vpcId")
    def vpc_id(self) -> pulumi.Input[builtins.str]:
        """
        The ID of the VPC to which the NAT gateway belongs.
        """
        return pulumi.get(self, "vpc_id")

    @vpc_id.setter
    def vpc_id(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "vpc_id", value)

    @property
    @pulumi.getter(name="autoRenew")
    def auto_renew(self) -> Optional[pulumi.Input[builtins.str]]:
        return pulumi.get(self, "auto_renew")

    @auto_renew.setter
    def auto_renew(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "auto_renew", value)

    @property
    @pulumi.getter(name="chargingMode")
    def charging_mode(self) -> Optional[pulumi.Input[builtins.str]]:
        return pulumi.get(self, "charging_mode")

    @charging_mode.setter
    def charging_mode(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "charging_mode", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The description of the NAT gateway.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="enterpriseProjectId")
    def enterprise_project_id(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The enterprise project ID of the NAT gateway.
        """
        return pulumi.get(self, "enterprise_project_id")

    @enterprise_project_id.setter
    def enterprise_project_id(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "enterprise_project_id", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The NAT gateway name.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="ngportIpAddress")
    def ngport_ip_address(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The IP address used for the NG port of the NAT gateway.
        """
        return pulumi.get(self, "ngport_ip_address")

    @ngport_ip_address.setter
    def ngport_ip_address(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "ngport_ip_address", value)

    @property
    @pulumi.getter
    def period(self) -> Optional[pulumi.Input[builtins.int]]:
        return pulumi.get(self, "period")

    @period.setter
    def period(self, value: Optional[pulumi.Input[builtins.int]]):
        pulumi.set(self, "period", value)

    @property
    @pulumi.getter(name="periodUnit")
    def period_unit(self) -> Optional[pulumi.Input[builtins.str]]:
        return pulumi.get(self, "period_unit")

    @period_unit.setter
    def period_unit(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "period_unit", value)

    @property
    @pulumi.getter
    def region(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The region where the NAT gateway is located.
        """
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "region", value)

    @property
    @pulumi.getter(name="sessionConf")
    def session_conf(self) -> Optional[pulumi.Input['NatGatewaySessionConfArgs']]:
        """
        The session configuration of the NAT gateway.
        """
        return pulumi.get(self, "session_conf")

    @session_conf.setter
    def session_conf(self, value: Optional[pulumi.Input['NatGatewaySessionConfArgs']]):
        pulumi.set(self, "session_conf", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]]]:
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]]]):
        pulumi.set(self, "tags", value)


@pulumi.input_type
class _NatGatewayState:
    def __init__(__self__, *,
                 auto_renew: Optional[pulumi.Input[builtins.str]] = None,
                 billing_info: Optional[pulumi.Input[builtins.str]] = None,
                 bps_max: Optional[pulumi.Input[builtins.int]] = None,
                 charging_mode: Optional[pulumi.Input[builtins.str]] = None,
                 created_at: Optional[pulumi.Input[builtins.str]] = None,
                 description: Optional[pulumi.Input[builtins.str]] = None,
                 dnat_rules_limit: Optional[pulumi.Input[builtins.int]] = None,
                 enterprise_project_id: Optional[pulumi.Input[builtins.str]] = None,
                 name: Optional[pulumi.Input[builtins.str]] = None,
                 ngport_ip_address: Optional[pulumi.Input[builtins.str]] = None,
                 period: Optional[pulumi.Input[builtins.int]] = None,
                 period_unit: Optional[pulumi.Input[builtins.str]] = None,
                 pps_max: Optional[pulumi.Input[builtins.int]] = None,
                 region: Optional[pulumi.Input[builtins.str]] = None,
                 session_conf: Optional[pulumi.Input['NatGatewaySessionConfArgs']] = None,
                 snat_rule_public_ip_limit: Optional[pulumi.Input[builtins.int]] = None,
                 spec: Optional[pulumi.Input[builtins.str]] = None,
                 status: Optional[pulumi.Input[builtins.str]] = None,
                 subnet_id: Optional[pulumi.Input[builtins.str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]]] = None,
                 vpc_id: Optional[pulumi.Input[builtins.str]] = None):
        """
        Input properties used for looking up and filtering NatGateway resources.
        :param pulumi.Input[builtins.str] billing_info: The order information of the NAT gateway.
        :param pulumi.Input[builtins.int] bps_max: The bandwidth that the NAT gateway can receive or send per second.
        :param pulumi.Input[builtins.str] created_at: The creation time of the NAT gateway.
        :param pulumi.Input[builtins.str] description: The description of the NAT gateway.
        :param pulumi.Input[builtins.int] dnat_rules_limit: The maximum number of DNAT rules on the NAT gateway.
        :param pulumi.Input[builtins.str] enterprise_project_id: The enterprise project ID of the NAT gateway.
        :param pulumi.Input[builtins.str] name: The NAT gateway name.
        :param pulumi.Input[builtins.str] ngport_ip_address: The IP address used for the NG port of the NAT gateway.
        :param pulumi.Input[builtins.int] pps_max: The number of packets that the NAT gateway can receive or send per second.
        :param pulumi.Input[builtins.str] region: The region where the NAT gateway is located.
        :param pulumi.Input['NatGatewaySessionConfArgs'] session_conf: The session configuration of the NAT gateway.
        :param pulumi.Input[builtins.int] snat_rule_public_ip_limit: The maximum number of SNAT rules on the NAT gateway.
        :param pulumi.Input[builtins.str] spec: The specification of the NAT gateway.
        :param pulumi.Input[builtins.str] status: The current status of the NAT gateway.
        :param pulumi.Input[builtins.str] subnet_id: The network ID of the downstream interface (the next hop of the DVR) of the NAT gateway.
        :param pulumi.Input[builtins.str] vpc_id: The ID of the VPC to which the NAT gateway belongs.
        """
        if auto_renew is not None:
            pulumi.set(__self__, "auto_renew", auto_renew)
        if billing_info is not None:
            pulumi.set(__self__, "billing_info", billing_info)
        if bps_max is not None:
            pulumi.set(__self__, "bps_max", bps_max)
        if charging_mode is not None:
            pulumi.set(__self__, "charging_mode", charging_mode)
        if created_at is not None:
            pulumi.set(__self__, "created_at", created_at)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if dnat_rules_limit is not None:
            pulumi.set(__self__, "dnat_rules_limit", dnat_rules_limit)
        if enterprise_project_id is not None:
            pulumi.set(__self__, "enterprise_project_id", enterprise_project_id)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if ngport_ip_address is not None:
            pulumi.set(__self__, "ngport_ip_address", ngport_ip_address)
        if period is not None:
            pulumi.set(__self__, "period", period)
        if period_unit is not None:
            pulumi.set(__self__, "period_unit", period_unit)
        if pps_max is not None:
            pulumi.set(__self__, "pps_max", pps_max)
        if region is not None:
            pulumi.set(__self__, "region", region)
        if session_conf is not None:
            pulumi.set(__self__, "session_conf", session_conf)
        if snat_rule_public_ip_limit is not None:
            pulumi.set(__self__, "snat_rule_public_ip_limit", snat_rule_public_ip_limit)
        if spec is not None:
            pulumi.set(__self__, "spec", spec)
        if status is not None:
            pulumi.set(__self__, "status", status)
        if subnet_id is not None:
            pulumi.set(__self__, "subnet_id", subnet_id)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)
        if vpc_id is not None:
            pulumi.set(__self__, "vpc_id", vpc_id)

    @property
    @pulumi.getter(name="autoRenew")
    def auto_renew(self) -> Optional[pulumi.Input[builtins.str]]:
        return pulumi.get(self, "auto_renew")

    @auto_renew.setter
    def auto_renew(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "auto_renew", value)

    @property
    @pulumi.getter(name="billingInfo")
    def billing_info(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The order information of the NAT gateway.
        """
        return pulumi.get(self, "billing_info")

    @billing_info.setter
    def billing_info(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "billing_info", value)

    @property
    @pulumi.getter(name="bpsMax")
    def bps_max(self) -> Optional[pulumi.Input[builtins.int]]:
        """
        The bandwidth that the NAT gateway can receive or send per second.
        """
        return pulumi.get(self, "bps_max")

    @bps_max.setter
    def bps_max(self, value: Optional[pulumi.Input[builtins.int]]):
        pulumi.set(self, "bps_max", value)

    @property
    @pulumi.getter(name="chargingMode")
    def charging_mode(self) -> Optional[pulumi.Input[builtins.str]]:
        return pulumi.get(self, "charging_mode")

    @charging_mode.setter
    def charging_mode(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "charging_mode", value)

    @property
    @pulumi.getter(name="createdAt")
    def created_at(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The creation time of the NAT gateway.
        """
        return pulumi.get(self, "created_at")

    @created_at.setter
    def created_at(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "created_at", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The description of the NAT gateway.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="dnatRulesLimit")
    def dnat_rules_limit(self) -> Optional[pulumi.Input[builtins.int]]:
        """
        The maximum number of DNAT rules on the NAT gateway.
        """
        return pulumi.get(self, "dnat_rules_limit")

    @dnat_rules_limit.setter
    def dnat_rules_limit(self, value: Optional[pulumi.Input[builtins.int]]):
        pulumi.set(self, "dnat_rules_limit", value)

    @property
    @pulumi.getter(name="enterpriseProjectId")
    def enterprise_project_id(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The enterprise project ID of the NAT gateway.
        """
        return pulumi.get(self, "enterprise_project_id")

    @enterprise_project_id.setter
    def enterprise_project_id(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "enterprise_project_id", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The NAT gateway name.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="ngportIpAddress")
    def ngport_ip_address(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The IP address used for the NG port of the NAT gateway.
        """
        return pulumi.get(self, "ngport_ip_address")

    @ngport_ip_address.setter
    def ngport_ip_address(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "ngport_ip_address", value)

    @property
    @pulumi.getter
    def period(self) -> Optional[pulumi.Input[builtins.int]]:
        return pulumi.get(self, "period")

    @period.setter
    def period(self, value: Optional[pulumi.Input[builtins.int]]):
        pulumi.set(self, "period", value)

    @property
    @pulumi.getter(name="periodUnit")
    def period_unit(self) -> Optional[pulumi.Input[builtins.str]]:
        return pulumi.get(self, "period_unit")

    @period_unit.setter
    def period_unit(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "period_unit", value)

    @property
    @pulumi.getter(name="ppsMax")
    def pps_max(self) -> Optional[pulumi.Input[builtins.int]]:
        """
        The number of packets that the NAT gateway can receive or send per second.
        """
        return pulumi.get(self, "pps_max")

    @pps_max.setter
    def pps_max(self, value: Optional[pulumi.Input[builtins.int]]):
        pulumi.set(self, "pps_max", value)

    @property
    @pulumi.getter
    def region(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The region where the NAT gateway is located.
        """
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "region", value)

    @property
    @pulumi.getter(name="sessionConf")
    def session_conf(self) -> Optional[pulumi.Input['NatGatewaySessionConfArgs']]:
        """
        The session configuration of the NAT gateway.
        """
        return pulumi.get(self, "session_conf")

    @session_conf.setter
    def session_conf(self, value: Optional[pulumi.Input['NatGatewaySessionConfArgs']]):
        pulumi.set(self, "session_conf", value)

    @property
    @pulumi.getter(name="snatRulePublicIpLimit")
    def snat_rule_public_ip_limit(self) -> Optional[pulumi.Input[builtins.int]]:
        """
        The maximum number of SNAT rules on the NAT gateway.
        """
        return pulumi.get(self, "snat_rule_public_ip_limit")

    @snat_rule_public_ip_limit.setter
    def snat_rule_public_ip_limit(self, value: Optional[pulumi.Input[builtins.int]]):
        pulumi.set(self, "snat_rule_public_ip_limit", value)

    @property
    @pulumi.getter
    def spec(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The specification of the NAT gateway.
        """
        return pulumi.get(self, "spec")

    @spec.setter
    def spec(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "spec", value)

    @property
    @pulumi.getter
    def status(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The current status of the NAT gateway.
        """
        return pulumi.get(self, "status")

    @status.setter
    def status(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "status", value)

    @property
    @pulumi.getter(name="subnetId")
    def subnet_id(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The network ID of the downstream interface (the next hop of the DVR) of the NAT gateway.
        """
        return pulumi.get(self, "subnet_id")

    @subnet_id.setter
    def subnet_id(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "subnet_id", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]]]:
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]]]):
        pulumi.set(self, "tags", value)

    @property
    @pulumi.getter(name="vpcId")
    def vpc_id(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The ID of the VPC to which the NAT gateway belongs.
        """
        return pulumi.get(self, "vpc_id")

    @vpc_id.setter
    def vpc_id(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "vpc_id", value)


class NatGateway(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 auto_renew: Optional[pulumi.Input[builtins.str]] = None,
                 charging_mode: Optional[pulumi.Input[builtins.str]] = None,
                 description: Optional[pulumi.Input[builtins.str]] = None,
                 enterprise_project_id: Optional[pulumi.Input[builtins.str]] = None,
                 name: Optional[pulumi.Input[builtins.str]] = None,
                 ngport_ip_address: Optional[pulumi.Input[builtins.str]] = None,
                 period: Optional[pulumi.Input[builtins.int]] = None,
                 period_unit: Optional[pulumi.Input[builtins.str]] = None,
                 region: Optional[pulumi.Input[builtins.str]] = None,
                 session_conf: Optional[pulumi.Input[Union['NatGatewaySessionConfArgs', 'NatGatewaySessionConfArgsDict']]] = None,
                 spec: Optional[pulumi.Input[builtins.str]] = None,
                 subnet_id: Optional[pulumi.Input[builtins.str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]]] = None,
                 vpc_id: Optional[pulumi.Input[builtins.str]] = None,
                 __props__=None):
        """
        Create a NatGateway resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[builtins.str] description: The description of the NAT gateway.
        :param pulumi.Input[builtins.str] enterprise_project_id: The enterprise project ID of the NAT gateway.
        :param pulumi.Input[builtins.str] name: The NAT gateway name.
        :param pulumi.Input[builtins.str] ngport_ip_address: The IP address used for the NG port of the NAT gateway.
        :param pulumi.Input[builtins.str] region: The region where the NAT gateway is located.
        :param pulumi.Input[Union['NatGatewaySessionConfArgs', 'NatGatewaySessionConfArgsDict']] session_conf: The session configuration of the NAT gateway.
        :param pulumi.Input[builtins.str] spec: The specification of the NAT gateway.
        :param pulumi.Input[builtins.str] subnet_id: The network ID of the downstream interface (the next hop of the DVR) of the NAT gateway.
        :param pulumi.Input[builtins.str] vpc_id: The ID of the VPC to which the NAT gateway belongs.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: NatGatewayArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a NatGateway resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param NatGatewayArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(NatGatewayArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 auto_renew: Optional[pulumi.Input[builtins.str]] = None,
                 charging_mode: Optional[pulumi.Input[builtins.str]] = None,
                 description: Optional[pulumi.Input[builtins.str]] = None,
                 enterprise_project_id: Optional[pulumi.Input[builtins.str]] = None,
                 name: Optional[pulumi.Input[builtins.str]] = None,
                 ngport_ip_address: Optional[pulumi.Input[builtins.str]] = None,
                 period: Optional[pulumi.Input[builtins.int]] = None,
                 period_unit: Optional[pulumi.Input[builtins.str]] = None,
                 region: Optional[pulumi.Input[builtins.str]] = None,
                 session_conf: Optional[pulumi.Input[Union['NatGatewaySessionConfArgs', 'NatGatewaySessionConfArgsDict']]] = None,
                 spec: Optional[pulumi.Input[builtins.str]] = None,
                 subnet_id: Optional[pulumi.Input[builtins.str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]]] = None,
                 vpc_id: Optional[pulumi.Input[builtins.str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = NatGatewayArgs.__new__(NatGatewayArgs)

            __props__.__dict__["auto_renew"] = auto_renew
            __props__.__dict__["charging_mode"] = charging_mode
            __props__.__dict__["description"] = description
            __props__.__dict__["enterprise_project_id"] = enterprise_project_id
            __props__.__dict__["name"] = name
            __props__.__dict__["ngport_ip_address"] = ngport_ip_address
            __props__.__dict__["period"] = period
            __props__.__dict__["period_unit"] = period_unit
            __props__.__dict__["region"] = region
            __props__.__dict__["session_conf"] = session_conf
            if spec is None and not opts.urn:
                raise TypeError("Missing required property 'spec'")
            __props__.__dict__["spec"] = spec
            if subnet_id is None and not opts.urn:
                raise TypeError("Missing required property 'subnet_id'")
            __props__.__dict__["subnet_id"] = subnet_id
            __props__.__dict__["tags"] = tags
            if vpc_id is None and not opts.urn:
                raise TypeError("Missing required property 'vpc_id'")
            __props__.__dict__["vpc_id"] = vpc_id
            __props__.__dict__["billing_info"] = None
            __props__.__dict__["bps_max"] = None
            __props__.__dict__["created_at"] = None
            __props__.__dict__["dnat_rules_limit"] = None
            __props__.__dict__["pps_max"] = None
            __props__.__dict__["snat_rule_public_ip_limit"] = None
            __props__.__dict__["status"] = None
        super(NatGateway, __self__).__init__(
            'sbercloud:index/natGateway:NatGateway',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            auto_renew: Optional[pulumi.Input[builtins.str]] = None,
            billing_info: Optional[pulumi.Input[builtins.str]] = None,
            bps_max: Optional[pulumi.Input[builtins.int]] = None,
            charging_mode: Optional[pulumi.Input[builtins.str]] = None,
            created_at: Optional[pulumi.Input[builtins.str]] = None,
            description: Optional[pulumi.Input[builtins.str]] = None,
            dnat_rules_limit: Optional[pulumi.Input[builtins.int]] = None,
            enterprise_project_id: Optional[pulumi.Input[builtins.str]] = None,
            name: Optional[pulumi.Input[builtins.str]] = None,
            ngport_ip_address: Optional[pulumi.Input[builtins.str]] = None,
            period: Optional[pulumi.Input[builtins.int]] = None,
            period_unit: Optional[pulumi.Input[builtins.str]] = None,
            pps_max: Optional[pulumi.Input[builtins.int]] = None,
            region: Optional[pulumi.Input[builtins.str]] = None,
            session_conf: Optional[pulumi.Input[Union['NatGatewaySessionConfArgs', 'NatGatewaySessionConfArgsDict']]] = None,
            snat_rule_public_ip_limit: Optional[pulumi.Input[builtins.int]] = None,
            spec: Optional[pulumi.Input[builtins.str]] = None,
            status: Optional[pulumi.Input[builtins.str]] = None,
            subnet_id: Optional[pulumi.Input[builtins.str]] = None,
            tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]]] = None,
            vpc_id: Optional[pulumi.Input[builtins.str]] = None) -> 'NatGateway':
        """
        Get an existing NatGateway resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[builtins.str] billing_info: The order information of the NAT gateway.
        :param pulumi.Input[builtins.int] bps_max: The bandwidth that the NAT gateway can receive or send per second.
        :param pulumi.Input[builtins.str] created_at: The creation time of the NAT gateway.
        :param pulumi.Input[builtins.str] description: The description of the NAT gateway.
        :param pulumi.Input[builtins.int] dnat_rules_limit: The maximum number of DNAT rules on the NAT gateway.
        :param pulumi.Input[builtins.str] enterprise_project_id: The enterprise project ID of the NAT gateway.
        :param pulumi.Input[builtins.str] name: The NAT gateway name.
        :param pulumi.Input[builtins.str] ngport_ip_address: The IP address used for the NG port of the NAT gateway.
        :param pulumi.Input[builtins.int] pps_max: The number of packets that the NAT gateway can receive or send per second.
        :param pulumi.Input[builtins.str] region: The region where the NAT gateway is located.
        :param pulumi.Input[Union['NatGatewaySessionConfArgs', 'NatGatewaySessionConfArgsDict']] session_conf: The session configuration of the NAT gateway.
        :param pulumi.Input[builtins.int] snat_rule_public_ip_limit: The maximum number of SNAT rules on the NAT gateway.
        :param pulumi.Input[builtins.str] spec: The specification of the NAT gateway.
        :param pulumi.Input[builtins.str] status: The current status of the NAT gateway.
        :param pulumi.Input[builtins.str] subnet_id: The network ID of the downstream interface (the next hop of the DVR) of the NAT gateway.
        :param pulumi.Input[builtins.str] vpc_id: The ID of the VPC to which the NAT gateway belongs.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _NatGatewayState.__new__(_NatGatewayState)

        __props__.__dict__["auto_renew"] = auto_renew
        __props__.__dict__["billing_info"] = billing_info
        __props__.__dict__["bps_max"] = bps_max
        __props__.__dict__["charging_mode"] = charging_mode
        __props__.__dict__["created_at"] = created_at
        __props__.__dict__["description"] = description
        __props__.__dict__["dnat_rules_limit"] = dnat_rules_limit
        __props__.__dict__["enterprise_project_id"] = enterprise_project_id
        __props__.__dict__["name"] = name
        __props__.__dict__["ngport_ip_address"] = ngport_ip_address
        __props__.__dict__["period"] = period
        __props__.__dict__["period_unit"] = period_unit
        __props__.__dict__["pps_max"] = pps_max
        __props__.__dict__["region"] = region
        __props__.__dict__["session_conf"] = session_conf
        __props__.__dict__["snat_rule_public_ip_limit"] = snat_rule_public_ip_limit
        __props__.__dict__["spec"] = spec
        __props__.__dict__["status"] = status
        __props__.__dict__["subnet_id"] = subnet_id
        __props__.__dict__["tags"] = tags
        __props__.__dict__["vpc_id"] = vpc_id
        return NatGateway(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="autoRenew")
    def auto_renew(self) -> pulumi.Output[Optional[builtins.str]]:
        return pulumi.get(self, "auto_renew")

    @property
    @pulumi.getter(name="billingInfo")
    def billing_info(self) -> pulumi.Output[builtins.str]:
        """
        The order information of the NAT gateway.
        """
        return pulumi.get(self, "billing_info")

    @property
    @pulumi.getter(name="bpsMax")
    def bps_max(self) -> pulumi.Output[builtins.int]:
        """
        The bandwidth that the NAT gateway can receive or send per second.
        """
        return pulumi.get(self, "bps_max")

    @property
    @pulumi.getter(name="chargingMode")
    def charging_mode(self) -> pulumi.Output[builtins.str]:
        return pulumi.get(self, "charging_mode")

    @property
    @pulumi.getter(name="createdAt")
    def created_at(self) -> pulumi.Output[builtins.str]:
        """
        The creation time of the NAT gateway.
        """
        return pulumi.get(self, "created_at")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[builtins.str]]:
        """
        The description of the NAT gateway.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="dnatRulesLimit")
    def dnat_rules_limit(self) -> pulumi.Output[builtins.int]:
        """
        The maximum number of DNAT rules on the NAT gateway.
        """
        return pulumi.get(self, "dnat_rules_limit")

    @property
    @pulumi.getter(name="enterpriseProjectId")
    def enterprise_project_id(self) -> pulumi.Output[builtins.str]:
        """
        The enterprise project ID of the NAT gateway.
        """
        return pulumi.get(self, "enterprise_project_id")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[builtins.str]:
        """
        The NAT gateway name.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="ngportIpAddress")
    def ngport_ip_address(self) -> pulumi.Output[builtins.str]:
        """
        The IP address used for the NG port of the NAT gateway.
        """
        return pulumi.get(self, "ngport_ip_address")

    @property
    @pulumi.getter
    def period(self) -> pulumi.Output[Optional[builtins.int]]:
        return pulumi.get(self, "period")

    @property
    @pulumi.getter(name="periodUnit")
    def period_unit(self) -> pulumi.Output[Optional[builtins.str]]:
        return pulumi.get(self, "period_unit")

    @property
    @pulumi.getter(name="ppsMax")
    def pps_max(self) -> pulumi.Output[builtins.int]:
        """
        The number of packets that the NAT gateway can receive or send per second.
        """
        return pulumi.get(self, "pps_max")

    @property
    @pulumi.getter
    def region(self) -> pulumi.Output[builtins.str]:
        """
        The region where the NAT gateway is located.
        """
        return pulumi.get(self, "region")

    @property
    @pulumi.getter(name="sessionConf")
    def session_conf(self) -> pulumi.Output['outputs.NatGatewaySessionConf']:
        """
        The session configuration of the NAT gateway.
        """
        return pulumi.get(self, "session_conf")

    @property
    @pulumi.getter(name="snatRulePublicIpLimit")
    def snat_rule_public_ip_limit(self) -> pulumi.Output[builtins.int]:
        """
        The maximum number of SNAT rules on the NAT gateway.
        """
        return pulumi.get(self, "snat_rule_public_ip_limit")

    @property
    @pulumi.getter
    def spec(self) -> pulumi.Output[builtins.str]:
        """
        The specification of the NAT gateway.
        """
        return pulumi.get(self, "spec")

    @property
    @pulumi.getter
    def status(self) -> pulumi.Output[builtins.str]:
        """
        The current status of the NAT gateway.
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter(name="subnetId")
    def subnet_id(self) -> pulumi.Output[builtins.str]:
        """
        The network ID of the downstream interface (the next hop of the DVR) of the NAT gateway.
        """
        return pulumi.get(self, "subnet_id")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Mapping[str, builtins.str]]]:
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="vpcId")
    def vpc_id(self) -> pulumi.Output[builtins.str]:
        """
        The ID of the VPC to which the NAT gateway belongs.
        """
        return pulumi.get(self, "vpc_id")

