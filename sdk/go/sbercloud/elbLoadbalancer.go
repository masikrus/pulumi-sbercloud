// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package sbercloud

import (
	"context"
	"reflect"

	"errors"
	"github.com/masikrus/pulumi-sbercloud/sdk/go/sbercloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ElbLoadbalancer struct {
	pulumi.CustomResourceState

	// Deprecated: Deprecated
	AutoPay                  pulumi.StringPtrOutput   `pulumi:"autoPay"`
	AutoRenew                pulumi.StringPtrOutput   `pulumi:"autoRenew"`
	AutoscalingEnabled       pulumi.BoolOutput        `pulumi:"autoscalingEnabled"`
	AvailabilityZones        pulumi.StringArrayOutput `pulumi:"availabilityZones"`
	BackendSubnets           pulumi.StringArrayOutput `pulumi:"backendSubnets"`
	BandwidthChargeMode      pulumi.StringOutput      `pulumi:"bandwidthChargeMode"`
	BandwidthId              pulumi.StringOutput      `pulumi:"bandwidthId"`
	BandwidthSize            pulumi.IntOutput         `pulumi:"bandwidthSize"`
	ChargeMode               pulumi.StringOutput      `pulumi:"chargeMode"`
	ChargingMode             pulumi.StringOutput      `pulumi:"chargingMode"`
	CreatedAt                pulumi.StringOutput      `pulumi:"createdAt"`
	CrossVpcBackend          pulumi.BoolOutput        `pulumi:"crossVpcBackend"`
	DeletionProtectionEnable pulumi.BoolPtrOutput     `pulumi:"deletionProtectionEnable"`
	Description              pulumi.StringPtrOutput   `pulumi:"description"`
	EnterpriseProjectId      pulumi.StringOutput      `pulumi:"enterpriseProjectId"`
	ForceDelete              pulumi.BoolPtrOutput     `pulumi:"forceDelete"`
	Guaranteed               pulumi.BoolOutput        `pulumi:"guaranteed"`
	GwFlavorId               pulumi.StringOutput      `pulumi:"gwFlavorId"`
	Iptype                   pulumi.StringOutput      `pulumi:"iptype"`
	Ipv4Address              pulumi.StringOutput      `pulumi:"ipv4Address"`
	Ipv4Eip                  pulumi.StringOutput      `pulumi:"ipv4Eip"`
	Ipv4EipId                pulumi.StringOutput      `pulumi:"ipv4EipId"`
	Ipv4PortId               pulumi.StringOutput      `pulumi:"ipv4PortId"`
	// the IPv4 subnet ID of the subnet where the load balancer resides
	Ipv4SubnetId    pulumi.StringPtrOutput `pulumi:"ipv4SubnetId"`
	Ipv6Address     pulumi.StringOutput    `pulumi:"ipv6Address"`
	Ipv6BandwidthId pulumi.StringPtrOutput `pulumi:"ipv6BandwidthId"`
	Ipv6Eip         pulumi.StringOutput    `pulumi:"ipv6Eip"`
	Ipv6EipId       pulumi.StringOutput    `pulumi:"ipv6EipId"`
	// the ID of the subnet where the load balancer resides
	Ipv6NetworkId    pulumi.StringPtrOutput `pulumi:"ipv6NetworkId"`
	L4FlavorId       pulumi.StringOutput    `pulumi:"l4FlavorId"`
	L7FlavorId       pulumi.StringOutput    `pulumi:"l7FlavorId"`
	LoadbalancerType pulumi.StringOutput    `pulumi:"loadbalancerType"`
	MinL7FlavorId    pulumi.StringOutput    `pulumi:"minL7FlavorId"`
	Name             pulumi.StringOutput    `pulumi:"name"`
	Period           pulumi.IntPtrOutput    `pulumi:"period"`
	PeriodUnit       pulumi.StringPtrOutput `pulumi:"periodUnit"`
	ProtectionReason pulumi.StringPtrOutput `pulumi:"protectionReason"`
	ProtectionStatus pulumi.StringOutput    `pulumi:"protectionStatus"`
	Region           pulumi.StringOutput    `pulumi:"region"`
	Sharetype        pulumi.StringOutput    `pulumi:"sharetype"`
	Tags             pulumi.StringMapOutput `pulumi:"tags"`
	UpdatedAt        pulumi.StringOutput    `pulumi:"updatedAt"`
	VpcId            pulumi.StringOutput    `pulumi:"vpcId"`
	WafFailureAction pulumi.StringOutput    `pulumi:"wafFailureAction"`
}

// NewElbLoadbalancer registers a new resource with the given unique name, arguments, and options.
func NewElbLoadbalancer(ctx *pulumi.Context,
	name string, args *ElbLoadbalancerArgs, opts ...pulumi.ResourceOption) (*ElbLoadbalancer, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AvailabilityZones == nil {
		return nil, errors.New("invalid value for required argument 'AvailabilityZones'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource ElbLoadbalancer
	err := ctx.RegisterResource("sbercloud:index/elbLoadbalancer:ElbLoadbalancer", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetElbLoadbalancer gets an existing ElbLoadbalancer resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetElbLoadbalancer(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ElbLoadbalancerState, opts ...pulumi.ResourceOption) (*ElbLoadbalancer, error) {
	var resource ElbLoadbalancer
	err := ctx.ReadResource("sbercloud:index/elbLoadbalancer:ElbLoadbalancer", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ElbLoadbalancer resources.
type elbLoadbalancerState struct {
	// Deprecated: Deprecated
	AutoPay                  *string  `pulumi:"autoPay"`
	AutoRenew                *string  `pulumi:"autoRenew"`
	AutoscalingEnabled       *bool    `pulumi:"autoscalingEnabled"`
	AvailabilityZones        []string `pulumi:"availabilityZones"`
	BackendSubnets           []string `pulumi:"backendSubnets"`
	BandwidthChargeMode      *string  `pulumi:"bandwidthChargeMode"`
	BandwidthId              *string  `pulumi:"bandwidthId"`
	BandwidthSize            *int     `pulumi:"bandwidthSize"`
	ChargeMode               *string  `pulumi:"chargeMode"`
	ChargingMode             *string  `pulumi:"chargingMode"`
	CreatedAt                *string  `pulumi:"createdAt"`
	CrossVpcBackend          *bool    `pulumi:"crossVpcBackend"`
	DeletionProtectionEnable *bool    `pulumi:"deletionProtectionEnable"`
	Description              *string  `pulumi:"description"`
	EnterpriseProjectId      *string  `pulumi:"enterpriseProjectId"`
	ForceDelete              *bool    `pulumi:"forceDelete"`
	Guaranteed               *bool    `pulumi:"guaranteed"`
	GwFlavorId               *string  `pulumi:"gwFlavorId"`
	Iptype                   *string  `pulumi:"iptype"`
	Ipv4Address              *string  `pulumi:"ipv4Address"`
	Ipv4Eip                  *string  `pulumi:"ipv4Eip"`
	Ipv4EipId                *string  `pulumi:"ipv4EipId"`
	Ipv4PortId               *string  `pulumi:"ipv4PortId"`
	// the IPv4 subnet ID of the subnet where the load balancer resides
	Ipv4SubnetId    *string `pulumi:"ipv4SubnetId"`
	Ipv6Address     *string `pulumi:"ipv6Address"`
	Ipv6BandwidthId *string `pulumi:"ipv6BandwidthId"`
	Ipv6Eip         *string `pulumi:"ipv6Eip"`
	Ipv6EipId       *string `pulumi:"ipv6EipId"`
	// the ID of the subnet where the load balancer resides
	Ipv6NetworkId    *string           `pulumi:"ipv6NetworkId"`
	L4FlavorId       *string           `pulumi:"l4FlavorId"`
	L7FlavorId       *string           `pulumi:"l7FlavorId"`
	LoadbalancerType *string           `pulumi:"loadbalancerType"`
	MinL7FlavorId    *string           `pulumi:"minL7FlavorId"`
	Name             *string           `pulumi:"name"`
	Period           *int              `pulumi:"period"`
	PeriodUnit       *string           `pulumi:"periodUnit"`
	ProtectionReason *string           `pulumi:"protectionReason"`
	ProtectionStatus *string           `pulumi:"protectionStatus"`
	Region           *string           `pulumi:"region"`
	Sharetype        *string           `pulumi:"sharetype"`
	Tags             map[string]string `pulumi:"tags"`
	UpdatedAt        *string           `pulumi:"updatedAt"`
	VpcId            *string           `pulumi:"vpcId"`
	WafFailureAction *string           `pulumi:"wafFailureAction"`
}

type ElbLoadbalancerState struct {
	// Deprecated: Deprecated
	AutoPay                  pulumi.StringPtrInput
	AutoRenew                pulumi.StringPtrInput
	AutoscalingEnabled       pulumi.BoolPtrInput
	AvailabilityZones        pulumi.StringArrayInput
	BackendSubnets           pulumi.StringArrayInput
	BandwidthChargeMode      pulumi.StringPtrInput
	BandwidthId              pulumi.StringPtrInput
	BandwidthSize            pulumi.IntPtrInput
	ChargeMode               pulumi.StringPtrInput
	ChargingMode             pulumi.StringPtrInput
	CreatedAt                pulumi.StringPtrInput
	CrossVpcBackend          pulumi.BoolPtrInput
	DeletionProtectionEnable pulumi.BoolPtrInput
	Description              pulumi.StringPtrInput
	EnterpriseProjectId      pulumi.StringPtrInput
	ForceDelete              pulumi.BoolPtrInput
	Guaranteed               pulumi.BoolPtrInput
	GwFlavorId               pulumi.StringPtrInput
	Iptype                   pulumi.StringPtrInput
	Ipv4Address              pulumi.StringPtrInput
	Ipv4Eip                  pulumi.StringPtrInput
	Ipv4EipId                pulumi.StringPtrInput
	Ipv4PortId               pulumi.StringPtrInput
	// the IPv4 subnet ID of the subnet where the load balancer resides
	Ipv4SubnetId    pulumi.StringPtrInput
	Ipv6Address     pulumi.StringPtrInput
	Ipv6BandwidthId pulumi.StringPtrInput
	Ipv6Eip         pulumi.StringPtrInput
	Ipv6EipId       pulumi.StringPtrInput
	// the ID of the subnet where the load balancer resides
	Ipv6NetworkId    pulumi.StringPtrInput
	L4FlavorId       pulumi.StringPtrInput
	L7FlavorId       pulumi.StringPtrInput
	LoadbalancerType pulumi.StringPtrInput
	MinL7FlavorId    pulumi.StringPtrInput
	Name             pulumi.StringPtrInput
	Period           pulumi.IntPtrInput
	PeriodUnit       pulumi.StringPtrInput
	ProtectionReason pulumi.StringPtrInput
	ProtectionStatus pulumi.StringPtrInput
	Region           pulumi.StringPtrInput
	Sharetype        pulumi.StringPtrInput
	Tags             pulumi.StringMapInput
	UpdatedAt        pulumi.StringPtrInput
	VpcId            pulumi.StringPtrInput
	WafFailureAction pulumi.StringPtrInput
}

func (ElbLoadbalancerState) ElementType() reflect.Type {
	return reflect.TypeOf((*elbLoadbalancerState)(nil)).Elem()
}

type elbLoadbalancerArgs struct {
	// Deprecated: Deprecated
	AutoPay                  *string  `pulumi:"autoPay"`
	AutoRenew                *string  `pulumi:"autoRenew"`
	AutoscalingEnabled       *bool    `pulumi:"autoscalingEnabled"`
	AvailabilityZones        []string `pulumi:"availabilityZones"`
	BackendSubnets           []string `pulumi:"backendSubnets"`
	BandwidthChargeMode      *string  `pulumi:"bandwidthChargeMode"`
	BandwidthId              *string  `pulumi:"bandwidthId"`
	BandwidthSize            *int     `pulumi:"bandwidthSize"`
	ChargingMode             *string  `pulumi:"chargingMode"`
	CrossVpcBackend          *bool    `pulumi:"crossVpcBackend"`
	DeletionProtectionEnable *bool    `pulumi:"deletionProtectionEnable"`
	Description              *string  `pulumi:"description"`
	EnterpriseProjectId      *string  `pulumi:"enterpriseProjectId"`
	ForceDelete              *bool    `pulumi:"forceDelete"`
	Iptype                   *string  `pulumi:"iptype"`
	Ipv4Address              *string  `pulumi:"ipv4Address"`
	Ipv4EipId                *string  `pulumi:"ipv4EipId"`
	// the IPv4 subnet ID of the subnet where the load balancer resides
	Ipv4SubnetId    *string `pulumi:"ipv4SubnetId"`
	Ipv6Address     *string `pulumi:"ipv6Address"`
	Ipv6BandwidthId *string `pulumi:"ipv6BandwidthId"`
	// the ID of the subnet where the load balancer resides
	Ipv6NetworkId    *string           `pulumi:"ipv6NetworkId"`
	L4FlavorId       *string           `pulumi:"l4FlavorId"`
	L7FlavorId       *string           `pulumi:"l7FlavorId"`
	LoadbalancerType *string           `pulumi:"loadbalancerType"`
	MinL7FlavorId    *string           `pulumi:"minL7FlavorId"`
	Name             *string           `pulumi:"name"`
	Period           *int              `pulumi:"period"`
	PeriodUnit       *string           `pulumi:"periodUnit"`
	ProtectionReason *string           `pulumi:"protectionReason"`
	ProtectionStatus *string           `pulumi:"protectionStatus"`
	Region           *string           `pulumi:"region"`
	Sharetype        *string           `pulumi:"sharetype"`
	Tags             map[string]string `pulumi:"tags"`
	VpcId            *string           `pulumi:"vpcId"`
	WafFailureAction *string           `pulumi:"wafFailureAction"`
}

// The set of arguments for constructing a ElbLoadbalancer resource.
type ElbLoadbalancerArgs struct {
	// Deprecated: Deprecated
	AutoPay                  pulumi.StringPtrInput
	AutoRenew                pulumi.StringPtrInput
	AutoscalingEnabled       pulumi.BoolPtrInput
	AvailabilityZones        pulumi.StringArrayInput
	BackendSubnets           pulumi.StringArrayInput
	BandwidthChargeMode      pulumi.StringPtrInput
	BandwidthId              pulumi.StringPtrInput
	BandwidthSize            pulumi.IntPtrInput
	ChargingMode             pulumi.StringPtrInput
	CrossVpcBackend          pulumi.BoolPtrInput
	DeletionProtectionEnable pulumi.BoolPtrInput
	Description              pulumi.StringPtrInput
	EnterpriseProjectId      pulumi.StringPtrInput
	ForceDelete              pulumi.BoolPtrInput
	Iptype                   pulumi.StringPtrInput
	Ipv4Address              pulumi.StringPtrInput
	Ipv4EipId                pulumi.StringPtrInput
	// the IPv4 subnet ID of the subnet where the load balancer resides
	Ipv4SubnetId    pulumi.StringPtrInput
	Ipv6Address     pulumi.StringPtrInput
	Ipv6BandwidthId pulumi.StringPtrInput
	// the ID of the subnet where the load balancer resides
	Ipv6NetworkId    pulumi.StringPtrInput
	L4FlavorId       pulumi.StringPtrInput
	L7FlavorId       pulumi.StringPtrInput
	LoadbalancerType pulumi.StringPtrInput
	MinL7FlavorId    pulumi.StringPtrInput
	Name             pulumi.StringPtrInput
	Period           pulumi.IntPtrInput
	PeriodUnit       pulumi.StringPtrInput
	ProtectionReason pulumi.StringPtrInput
	ProtectionStatus pulumi.StringPtrInput
	Region           pulumi.StringPtrInput
	Sharetype        pulumi.StringPtrInput
	Tags             pulumi.StringMapInput
	VpcId            pulumi.StringPtrInput
	WafFailureAction pulumi.StringPtrInput
}

func (ElbLoadbalancerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*elbLoadbalancerArgs)(nil)).Elem()
}

type ElbLoadbalancerInput interface {
	pulumi.Input

	ToElbLoadbalancerOutput() ElbLoadbalancerOutput
	ToElbLoadbalancerOutputWithContext(ctx context.Context) ElbLoadbalancerOutput
}

func (*ElbLoadbalancer) ElementType() reflect.Type {
	return reflect.TypeOf((**ElbLoadbalancer)(nil)).Elem()
}

func (i *ElbLoadbalancer) ToElbLoadbalancerOutput() ElbLoadbalancerOutput {
	return i.ToElbLoadbalancerOutputWithContext(context.Background())
}

func (i *ElbLoadbalancer) ToElbLoadbalancerOutputWithContext(ctx context.Context) ElbLoadbalancerOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ElbLoadbalancerOutput)
}

// ElbLoadbalancerArrayInput is an input type that accepts ElbLoadbalancerArray and ElbLoadbalancerArrayOutput values.
// You can construct a concrete instance of `ElbLoadbalancerArrayInput` via:
//
//	ElbLoadbalancerArray{ ElbLoadbalancerArgs{...} }
type ElbLoadbalancerArrayInput interface {
	pulumi.Input

	ToElbLoadbalancerArrayOutput() ElbLoadbalancerArrayOutput
	ToElbLoadbalancerArrayOutputWithContext(context.Context) ElbLoadbalancerArrayOutput
}

type ElbLoadbalancerArray []ElbLoadbalancerInput

func (ElbLoadbalancerArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ElbLoadbalancer)(nil)).Elem()
}

func (i ElbLoadbalancerArray) ToElbLoadbalancerArrayOutput() ElbLoadbalancerArrayOutput {
	return i.ToElbLoadbalancerArrayOutputWithContext(context.Background())
}

func (i ElbLoadbalancerArray) ToElbLoadbalancerArrayOutputWithContext(ctx context.Context) ElbLoadbalancerArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ElbLoadbalancerArrayOutput)
}

// ElbLoadbalancerMapInput is an input type that accepts ElbLoadbalancerMap and ElbLoadbalancerMapOutput values.
// You can construct a concrete instance of `ElbLoadbalancerMapInput` via:
//
//	ElbLoadbalancerMap{ "key": ElbLoadbalancerArgs{...} }
type ElbLoadbalancerMapInput interface {
	pulumi.Input

	ToElbLoadbalancerMapOutput() ElbLoadbalancerMapOutput
	ToElbLoadbalancerMapOutputWithContext(context.Context) ElbLoadbalancerMapOutput
}

type ElbLoadbalancerMap map[string]ElbLoadbalancerInput

func (ElbLoadbalancerMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ElbLoadbalancer)(nil)).Elem()
}

func (i ElbLoadbalancerMap) ToElbLoadbalancerMapOutput() ElbLoadbalancerMapOutput {
	return i.ToElbLoadbalancerMapOutputWithContext(context.Background())
}

func (i ElbLoadbalancerMap) ToElbLoadbalancerMapOutputWithContext(ctx context.Context) ElbLoadbalancerMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ElbLoadbalancerMapOutput)
}

type ElbLoadbalancerOutput struct{ *pulumi.OutputState }

func (ElbLoadbalancerOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ElbLoadbalancer)(nil)).Elem()
}

func (o ElbLoadbalancerOutput) ToElbLoadbalancerOutput() ElbLoadbalancerOutput {
	return o
}

func (o ElbLoadbalancerOutput) ToElbLoadbalancerOutputWithContext(ctx context.Context) ElbLoadbalancerOutput {
	return o
}

// Deprecated: Deprecated
func (o ElbLoadbalancerOutput) AutoPay() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ElbLoadbalancer) pulumi.StringPtrOutput { return v.AutoPay }).(pulumi.StringPtrOutput)
}

func (o ElbLoadbalancerOutput) AutoRenew() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ElbLoadbalancer) pulumi.StringPtrOutput { return v.AutoRenew }).(pulumi.StringPtrOutput)
}

func (o ElbLoadbalancerOutput) AutoscalingEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v *ElbLoadbalancer) pulumi.BoolOutput { return v.AutoscalingEnabled }).(pulumi.BoolOutput)
}

func (o ElbLoadbalancerOutput) AvailabilityZones() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ElbLoadbalancer) pulumi.StringArrayOutput { return v.AvailabilityZones }).(pulumi.StringArrayOutput)
}

func (o ElbLoadbalancerOutput) BackendSubnets() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ElbLoadbalancer) pulumi.StringArrayOutput { return v.BackendSubnets }).(pulumi.StringArrayOutput)
}

func (o ElbLoadbalancerOutput) BandwidthChargeMode() pulumi.StringOutput {
	return o.ApplyT(func(v *ElbLoadbalancer) pulumi.StringOutput { return v.BandwidthChargeMode }).(pulumi.StringOutput)
}

func (o ElbLoadbalancerOutput) BandwidthId() pulumi.StringOutput {
	return o.ApplyT(func(v *ElbLoadbalancer) pulumi.StringOutput { return v.BandwidthId }).(pulumi.StringOutput)
}

func (o ElbLoadbalancerOutput) BandwidthSize() pulumi.IntOutput {
	return o.ApplyT(func(v *ElbLoadbalancer) pulumi.IntOutput { return v.BandwidthSize }).(pulumi.IntOutput)
}

func (o ElbLoadbalancerOutput) ChargeMode() pulumi.StringOutput {
	return o.ApplyT(func(v *ElbLoadbalancer) pulumi.StringOutput { return v.ChargeMode }).(pulumi.StringOutput)
}

func (o ElbLoadbalancerOutput) ChargingMode() pulumi.StringOutput {
	return o.ApplyT(func(v *ElbLoadbalancer) pulumi.StringOutput { return v.ChargingMode }).(pulumi.StringOutput)
}

func (o ElbLoadbalancerOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *ElbLoadbalancer) pulumi.StringOutput { return v.CreatedAt }).(pulumi.StringOutput)
}

func (o ElbLoadbalancerOutput) CrossVpcBackend() pulumi.BoolOutput {
	return o.ApplyT(func(v *ElbLoadbalancer) pulumi.BoolOutput { return v.CrossVpcBackend }).(pulumi.BoolOutput)
}

func (o ElbLoadbalancerOutput) DeletionProtectionEnable() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ElbLoadbalancer) pulumi.BoolPtrOutput { return v.DeletionProtectionEnable }).(pulumi.BoolPtrOutput)
}

func (o ElbLoadbalancerOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ElbLoadbalancer) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

func (o ElbLoadbalancerOutput) EnterpriseProjectId() pulumi.StringOutput {
	return o.ApplyT(func(v *ElbLoadbalancer) pulumi.StringOutput { return v.EnterpriseProjectId }).(pulumi.StringOutput)
}

func (o ElbLoadbalancerOutput) ForceDelete() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ElbLoadbalancer) pulumi.BoolPtrOutput { return v.ForceDelete }).(pulumi.BoolPtrOutput)
}

func (o ElbLoadbalancerOutput) Guaranteed() pulumi.BoolOutput {
	return o.ApplyT(func(v *ElbLoadbalancer) pulumi.BoolOutput { return v.Guaranteed }).(pulumi.BoolOutput)
}

func (o ElbLoadbalancerOutput) GwFlavorId() pulumi.StringOutput {
	return o.ApplyT(func(v *ElbLoadbalancer) pulumi.StringOutput { return v.GwFlavorId }).(pulumi.StringOutput)
}

func (o ElbLoadbalancerOutput) Iptype() pulumi.StringOutput {
	return o.ApplyT(func(v *ElbLoadbalancer) pulumi.StringOutput { return v.Iptype }).(pulumi.StringOutput)
}

func (o ElbLoadbalancerOutput) Ipv4Address() pulumi.StringOutput {
	return o.ApplyT(func(v *ElbLoadbalancer) pulumi.StringOutput { return v.Ipv4Address }).(pulumi.StringOutput)
}

func (o ElbLoadbalancerOutput) Ipv4Eip() pulumi.StringOutput {
	return o.ApplyT(func(v *ElbLoadbalancer) pulumi.StringOutput { return v.Ipv4Eip }).(pulumi.StringOutput)
}

func (o ElbLoadbalancerOutput) Ipv4EipId() pulumi.StringOutput {
	return o.ApplyT(func(v *ElbLoadbalancer) pulumi.StringOutput { return v.Ipv4EipId }).(pulumi.StringOutput)
}

func (o ElbLoadbalancerOutput) Ipv4PortId() pulumi.StringOutput {
	return o.ApplyT(func(v *ElbLoadbalancer) pulumi.StringOutput { return v.Ipv4PortId }).(pulumi.StringOutput)
}

// the IPv4 subnet ID of the subnet where the load balancer resides
func (o ElbLoadbalancerOutput) Ipv4SubnetId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ElbLoadbalancer) pulumi.StringPtrOutput { return v.Ipv4SubnetId }).(pulumi.StringPtrOutput)
}

func (o ElbLoadbalancerOutput) Ipv6Address() pulumi.StringOutput {
	return o.ApplyT(func(v *ElbLoadbalancer) pulumi.StringOutput { return v.Ipv6Address }).(pulumi.StringOutput)
}

func (o ElbLoadbalancerOutput) Ipv6BandwidthId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ElbLoadbalancer) pulumi.StringPtrOutput { return v.Ipv6BandwidthId }).(pulumi.StringPtrOutput)
}

func (o ElbLoadbalancerOutput) Ipv6Eip() pulumi.StringOutput {
	return o.ApplyT(func(v *ElbLoadbalancer) pulumi.StringOutput { return v.Ipv6Eip }).(pulumi.StringOutput)
}

func (o ElbLoadbalancerOutput) Ipv6EipId() pulumi.StringOutput {
	return o.ApplyT(func(v *ElbLoadbalancer) pulumi.StringOutput { return v.Ipv6EipId }).(pulumi.StringOutput)
}

// the ID of the subnet where the load balancer resides
func (o ElbLoadbalancerOutput) Ipv6NetworkId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ElbLoadbalancer) pulumi.StringPtrOutput { return v.Ipv6NetworkId }).(pulumi.StringPtrOutput)
}

func (o ElbLoadbalancerOutput) L4FlavorId() pulumi.StringOutput {
	return o.ApplyT(func(v *ElbLoadbalancer) pulumi.StringOutput { return v.L4FlavorId }).(pulumi.StringOutput)
}

func (o ElbLoadbalancerOutput) L7FlavorId() pulumi.StringOutput {
	return o.ApplyT(func(v *ElbLoadbalancer) pulumi.StringOutput { return v.L7FlavorId }).(pulumi.StringOutput)
}

func (o ElbLoadbalancerOutput) LoadbalancerType() pulumi.StringOutput {
	return o.ApplyT(func(v *ElbLoadbalancer) pulumi.StringOutput { return v.LoadbalancerType }).(pulumi.StringOutput)
}

func (o ElbLoadbalancerOutput) MinL7FlavorId() pulumi.StringOutput {
	return o.ApplyT(func(v *ElbLoadbalancer) pulumi.StringOutput { return v.MinL7FlavorId }).(pulumi.StringOutput)
}

func (o ElbLoadbalancerOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ElbLoadbalancer) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o ElbLoadbalancerOutput) Period() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ElbLoadbalancer) pulumi.IntPtrOutput { return v.Period }).(pulumi.IntPtrOutput)
}

func (o ElbLoadbalancerOutput) PeriodUnit() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ElbLoadbalancer) pulumi.StringPtrOutput { return v.PeriodUnit }).(pulumi.StringPtrOutput)
}

func (o ElbLoadbalancerOutput) ProtectionReason() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ElbLoadbalancer) pulumi.StringPtrOutput { return v.ProtectionReason }).(pulumi.StringPtrOutput)
}

func (o ElbLoadbalancerOutput) ProtectionStatus() pulumi.StringOutput {
	return o.ApplyT(func(v *ElbLoadbalancer) pulumi.StringOutput { return v.ProtectionStatus }).(pulumi.StringOutput)
}

func (o ElbLoadbalancerOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v *ElbLoadbalancer) pulumi.StringOutput { return v.Region }).(pulumi.StringOutput)
}

func (o ElbLoadbalancerOutput) Sharetype() pulumi.StringOutput {
	return o.ApplyT(func(v *ElbLoadbalancer) pulumi.StringOutput { return v.Sharetype }).(pulumi.StringOutput)
}

func (o ElbLoadbalancerOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *ElbLoadbalancer) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o ElbLoadbalancerOutput) UpdatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *ElbLoadbalancer) pulumi.StringOutput { return v.UpdatedAt }).(pulumi.StringOutput)
}

func (o ElbLoadbalancerOutput) VpcId() pulumi.StringOutput {
	return o.ApplyT(func(v *ElbLoadbalancer) pulumi.StringOutput { return v.VpcId }).(pulumi.StringOutput)
}

func (o ElbLoadbalancerOutput) WafFailureAction() pulumi.StringOutput {
	return o.ApplyT(func(v *ElbLoadbalancer) pulumi.StringOutput { return v.WafFailureAction }).(pulumi.StringOutput)
}

type ElbLoadbalancerArrayOutput struct{ *pulumi.OutputState }

func (ElbLoadbalancerArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ElbLoadbalancer)(nil)).Elem()
}

func (o ElbLoadbalancerArrayOutput) ToElbLoadbalancerArrayOutput() ElbLoadbalancerArrayOutput {
	return o
}

func (o ElbLoadbalancerArrayOutput) ToElbLoadbalancerArrayOutputWithContext(ctx context.Context) ElbLoadbalancerArrayOutput {
	return o
}

func (o ElbLoadbalancerArrayOutput) Index(i pulumi.IntInput) ElbLoadbalancerOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ElbLoadbalancer {
		return vs[0].([]*ElbLoadbalancer)[vs[1].(int)]
	}).(ElbLoadbalancerOutput)
}

type ElbLoadbalancerMapOutput struct{ *pulumi.OutputState }

func (ElbLoadbalancerMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ElbLoadbalancer)(nil)).Elem()
}

func (o ElbLoadbalancerMapOutput) ToElbLoadbalancerMapOutput() ElbLoadbalancerMapOutput {
	return o
}

func (o ElbLoadbalancerMapOutput) ToElbLoadbalancerMapOutputWithContext(ctx context.Context) ElbLoadbalancerMapOutput {
	return o
}

func (o ElbLoadbalancerMapOutput) MapIndex(k pulumi.StringInput) ElbLoadbalancerOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ElbLoadbalancer {
		return vs[0].(map[string]*ElbLoadbalancer)[vs[1].(string)]
	}).(ElbLoadbalancerOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ElbLoadbalancerInput)(nil)).Elem(), &ElbLoadbalancer{})
	pulumi.RegisterInputType(reflect.TypeOf((*ElbLoadbalancerArrayInput)(nil)).Elem(), ElbLoadbalancerArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ElbLoadbalancerMapInput)(nil)).Elem(), ElbLoadbalancerMap{})
	pulumi.RegisterOutputType(ElbLoadbalancerOutput{})
	pulumi.RegisterOutputType(ElbLoadbalancerArrayOutput{})
	pulumi.RegisterOutputType(ElbLoadbalancerMapOutput{})
}
