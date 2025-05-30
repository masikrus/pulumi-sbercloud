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

type VpcSubnet struct {
	pulumi.CustomResourceState

	AvailabilityZone pulumi.StringOutput      `pulumi:"availabilityZone"`
	Cidr             pulumi.StringOutput      `pulumi:"cidr"`
	Description      pulumi.StringPtrOutput   `pulumi:"description"`
	DhcpEnable       pulumi.BoolPtrOutput     `pulumi:"dhcpEnable"`
	DhcpLeaseTime    pulumi.StringOutput      `pulumi:"dhcpLeaseTime"`
	DnsLists         pulumi.StringArrayOutput `pulumi:"dnsLists"`
	GatewayIp        pulumi.StringOutput      `pulumi:"gatewayIp"`
	Ipv4SubnetId     pulumi.StringOutput      `pulumi:"ipv4SubnetId"`
	Ipv6Cidr         pulumi.StringOutput      `pulumi:"ipv6Cidr"`
	Ipv6Enable       pulumi.BoolPtrOutput     `pulumi:"ipv6Enable"`
	Ipv6Gateway      pulumi.StringOutput      `pulumi:"ipv6Gateway"`
	Ipv6SubnetId     pulumi.StringOutput      `pulumi:"ipv6SubnetId"`
	Name             pulumi.StringOutput      `pulumi:"name"`
	NtpServerAddress pulumi.StringPtrOutput   `pulumi:"ntpServerAddress"`
	PrimaryDns       pulumi.StringOutput      `pulumi:"primaryDns"`
	Region           pulumi.StringOutput      `pulumi:"region"`
	SecondaryDns     pulumi.StringOutput      `pulumi:"secondaryDns"`
	// schema: Deprecated
	SubnetId pulumi.StringOutput    `pulumi:"subnetId"`
	Tags     pulumi.StringMapOutput `pulumi:"tags"`
	VpcId    pulumi.StringOutput    `pulumi:"vpcId"`
}

// NewVpcSubnet registers a new resource with the given unique name, arguments, and options.
func NewVpcSubnet(ctx *pulumi.Context,
	name string, args *VpcSubnetArgs, opts ...pulumi.ResourceOption) (*VpcSubnet, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Cidr == nil {
		return nil, errors.New("invalid value for required argument 'Cidr'")
	}
	if args.GatewayIp == nil {
		return nil, errors.New("invalid value for required argument 'GatewayIp'")
	}
	if args.VpcId == nil {
		return nil, errors.New("invalid value for required argument 'VpcId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource VpcSubnet
	err := ctx.RegisterResource("sbercloud:index/vpcSubnet:VpcSubnet", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetVpcSubnet gets an existing VpcSubnet resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVpcSubnet(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VpcSubnetState, opts ...pulumi.ResourceOption) (*VpcSubnet, error) {
	var resource VpcSubnet
	err := ctx.ReadResource("sbercloud:index/vpcSubnet:VpcSubnet", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering VpcSubnet resources.
type vpcSubnetState struct {
	AvailabilityZone *string  `pulumi:"availabilityZone"`
	Cidr             *string  `pulumi:"cidr"`
	Description      *string  `pulumi:"description"`
	DhcpEnable       *bool    `pulumi:"dhcpEnable"`
	DhcpLeaseTime    *string  `pulumi:"dhcpLeaseTime"`
	DnsLists         []string `pulumi:"dnsLists"`
	GatewayIp        *string  `pulumi:"gatewayIp"`
	Ipv4SubnetId     *string  `pulumi:"ipv4SubnetId"`
	Ipv6Cidr         *string  `pulumi:"ipv6Cidr"`
	Ipv6Enable       *bool    `pulumi:"ipv6Enable"`
	Ipv6Gateway      *string  `pulumi:"ipv6Gateway"`
	Ipv6SubnetId     *string  `pulumi:"ipv6SubnetId"`
	Name             *string  `pulumi:"name"`
	NtpServerAddress *string  `pulumi:"ntpServerAddress"`
	PrimaryDns       *string  `pulumi:"primaryDns"`
	Region           *string  `pulumi:"region"`
	SecondaryDns     *string  `pulumi:"secondaryDns"`
	// schema: Deprecated
	SubnetId *string           `pulumi:"subnetId"`
	Tags     map[string]string `pulumi:"tags"`
	VpcId    *string           `pulumi:"vpcId"`
}

type VpcSubnetState struct {
	AvailabilityZone pulumi.StringPtrInput
	Cidr             pulumi.StringPtrInput
	Description      pulumi.StringPtrInput
	DhcpEnable       pulumi.BoolPtrInput
	DhcpLeaseTime    pulumi.StringPtrInput
	DnsLists         pulumi.StringArrayInput
	GatewayIp        pulumi.StringPtrInput
	Ipv4SubnetId     pulumi.StringPtrInput
	Ipv6Cidr         pulumi.StringPtrInput
	Ipv6Enable       pulumi.BoolPtrInput
	Ipv6Gateway      pulumi.StringPtrInput
	Ipv6SubnetId     pulumi.StringPtrInput
	Name             pulumi.StringPtrInput
	NtpServerAddress pulumi.StringPtrInput
	PrimaryDns       pulumi.StringPtrInput
	Region           pulumi.StringPtrInput
	SecondaryDns     pulumi.StringPtrInput
	// schema: Deprecated
	SubnetId pulumi.StringPtrInput
	Tags     pulumi.StringMapInput
	VpcId    pulumi.StringPtrInput
}

func (VpcSubnetState) ElementType() reflect.Type {
	return reflect.TypeOf((*vpcSubnetState)(nil)).Elem()
}

type vpcSubnetArgs struct {
	AvailabilityZone *string           `pulumi:"availabilityZone"`
	Cidr             string            `pulumi:"cidr"`
	Description      *string           `pulumi:"description"`
	DhcpEnable       *bool             `pulumi:"dhcpEnable"`
	DhcpLeaseTime    *string           `pulumi:"dhcpLeaseTime"`
	DnsLists         []string          `pulumi:"dnsLists"`
	GatewayIp        string            `pulumi:"gatewayIp"`
	Ipv6Enable       *bool             `pulumi:"ipv6Enable"`
	Name             *string           `pulumi:"name"`
	NtpServerAddress *string           `pulumi:"ntpServerAddress"`
	PrimaryDns       *string           `pulumi:"primaryDns"`
	Region           *string           `pulumi:"region"`
	SecondaryDns     *string           `pulumi:"secondaryDns"`
	Tags             map[string]string `pulumi:"tags"`
	VpcId            string            `pulumi:"vpcId"`
}

// The set of arguments for constructing a VpcSubnet resource.
type VpcSubnetArgs struct {
	AvailabilityZone pulumi.StringPtrInput
	Cidr             pulumi.StringInput
	Description      pulumi.StringPtrInput
	DhcpEnable       pulumi.BoolPtrInput
	DhcpLeaseTime    pulumi.StringPtrInput
	DnsLists         pulumi.StringArrayInput
	GatewayIp        pulumi.StringInput
	Ipv6Enable       pulumi.BoolPtrInput
	Name             pulumi.StringPtrInput
	NtpServerAddress pulumi.StringPtrInput
	PrimaryDns       pulumi.StringPtrInput
	Region           pulumi.StringPtrInput
	SecondaryDns     pulumi.StringPtrInput
	Tags             pulumi.StringMapInput
	VpcId            pulumi.StringInput
}

func (VpcSubnetArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*vpcSubnetArgs)(nil)).Elem()
}

type VpcSubnetInput interface {
	pulumi.Input

	ToVpcSubnetOutput() VpcSubnetOutput
	ToVpcSubnetOutputWithContext(ctx context.Context) VpcSubnetOutput
}

func (*VpcSubnet) ElementType() reflect.Type {
	return reflect.TypeOf((**VpcSubnet)(nil)).Elem()
}

func (i *VpcSubnet) ToVpcSubnetOutput() VpcSubnetOutput {
	return i.ToVpcSubnetOutputWithContext(context.Background())
}

func (i *VpcSubnet) ToVpcSubnetOutputWithContext(ctx context.Context) VpcSubnetOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpcSubnetOutput)
}

// VpcSubnetArrayInput is an input type that accepts VpcSubnetArray and VpcSubnetArrayOutput values.
// You can construct a concrete instance of `VpcSubnetArrayInput` via:
//
//	VpcSubnetArray{ VpcSubnetArgs{...} }
type VpcSubnetArrayInput interface {
	pulumi.Input

	ToVpcSubnetArrayOutput() VpcSubnetArrayOutput
	ToVpcSubnetArrayOutputWithContext(context.Context) VpcSubnetArrayOutput
}

type VpcSubnetArray []VpcSubnetInput

func (VpcSubnetArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*VpcSubnet)(nil)).Elem()
}

func (i VpcSubnetArray) ToVpcSubnetArrayOutput() VpcSubnetArrayOutput {
	return i.ToVpcSubnetArrayOutputWithContext(context.Background())
}

func (i VpcSubnetArray) ToVpcSubnetArrayOutputWithContext(ctx context.Context) VpcSubnetArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpcSubnetArrayOutput)
}

// VpcSubnetMapInput is an input type that accepts VpcSubnetMap and VpcSubnetMapOutput values.
// You can construct a concrete instance of `VpcSubnetMapInput` via:
//
//	VpcSubnetMap{ "key": VpcSubnetArgs{...} }
type VpcSubnetMapInput interface {
	pulumi.Input

	ToVpcSubnetMapOutput() VpcSubnetMapOutput
	ToVpcSubnetMapOutputWithContext(context.Context) VpcSubnetMapOutput
}

type VpcSubnetMap map[string]VpcSubnetInput

func (VpcSubnetMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*VpcSubnet)(nil)).Elem()
}

func (i VpcSubnetMap) ToVpcSubnetMapOutput() VpcSubnetMapOutput {
	return i.ToVpcSubnetMapOutputWithContext(context.Background())
}

func (i VpcSubnetMap) ToVpcSubnetMapOutputWithContext(ctx context.Context) VpcSubnetMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpcSubnetMapOutput)
}

type VpcSubnetOutput struct{ *pulumi.OutputState }

func (VpcSubnetOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VpcSubnet)(nil)).Elem()
}

func (o VpcSubnetOutput) ToVpcSubnetOutput() VpcSubnetOutput {
	return o
}

func (o VpcSubnetOutput) ToVpcSubnetOutputWithContext(ctx context.Context) VpcSubnetOutput {
	return o
}

func (o VpcSubnetOutput) AvailabilityZone() pulumi.StringOutput {
	return o.ApplyT(func(v *VpcSubnet) pulumi.StringOutput { return v.AvailabilityZone }).(pulumi.StringOutput)
}

func (o VpcSubnetOutput) Cidr() pulumi.StringOutput {
	return o.ApplyT(func(v *VpcSubnet) pulumi.StringOutput { return v.Cidr }).(pulumi.StringOutput)
}

func (o VpcSubnetOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VpcSubnet) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

func (o VpcSubnetOutput) DhcpEnable() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *VpcSubnet) pulumi.BoolPtrOutput { return v.DhcpEnable }).(pulumi.BoolPtrOutput)
}

func (o VpcSubnetOutput) DhcpLeaseTime() pulumi.StringOutput {
	return o.ApplyT(func(v *VpcSubnet) pulumi.StringOutput { return v.DhcpLeaseTime }).(pulumi.StringOutput)
}

func (o VpcSubnetOutput) DnsLists() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *VpcSubnet) pulumi.StringArrayOutput { return v.DnsLists }).(pulumi.StringArrayOutput)
}

func (o VpcSubnetOutput) GatewayIp() pulumi.StringOutput {
	return o.ApplyT(func(v *VpcSubnet) pulumi.StringOutput { return v.GatewayIp }).(pulumi.StringOutput)
}

func (o VpcSubnetOutput) Ipv4SubnetId() pulumi.StringOutput {
	return o.ApplyT(func(v *VpcSubnet) pulumi.StringOutput { return v.Ipv4SubnetId }).(pulumi.StringOutput)
}

func (o VpcSubnetOutput) Ipv6Cidr() pulumi.StringOutput {
	return o.ApplyT(func(v *VpcSubnet) pulumi.StringOutput { return v.Ipv6Cidr }).(pulumi.StringOutput)
}

func (o VpcSubnetOutput) Ipv6Enable() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *VpcSubnet) pulumi.BoolPtrOutput { return v.Ipv6Enable }).(pulumi.BoolPtrOutput)
}

func (o VpcSubnetOutput) Ipv6Gateway() pulumi.StringOutput {
	return o.ApplyT(func(v *VpcSubnet) pulumi.StringOutput { return v.Ipv6Gateway }).(pulumi.StringOutput)
}

func (o VpcSubnetOutput) Ipv6SubnetId() pulumi.StringOutput {
	return o.ApplyT(func(v *VpcSubnet) pulumi.StringOutput { return v.Ipv6SubnetId }).(pulumi.StringOutput)
}

func (o VpcSubnetOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *VpcSubnet) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o VpcSubnetOutput) NtpServerAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VpcSubnet) pulumi.StringPtrOutput { return v.NtpServerAddress }).(pulumi.StringPtrOutput)
}

func (o VpcSubnetOutput) PrimaryDns() pulumi.StringOutput {
	return o.ApplyT(func(v *VpcSubnet) pulumi.StringOutput { return v.PrimaryDns }).(pulumi.StringOutput)
}

func (o VpcSubnetOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v *VpcSubnet) pulumi.StringOutput { return v.Region }).(pulumi.StringOutput)
}

func (o VpcSubnetOutput) SecondaryDns() pulumi.StringOutput {
	return o.ApplyT(func(v *VpcSubnet) pulumi.StringOutput { return v.SecondaryDns }).(pulumi.StringOutput)
}

// schema: Deprecated
func (o VpcSubnetOutput) SubnetId() pulumi.StringOutput {
	return o.ApplyT(func(v *VpcSubnet) pulumi.StringOutput { return v.SubnetId }).(pulumi.StringOutput)
}

func (o VpcSubnetOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *VpcSubnet) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o VpcSubnetOutput) VpcId() pulumi.StringOutput {
	return o.ApplyT(func(v *VpcSubnet) pulumi.StringOutput { return v.VpcId }).(pulumi.StringOutput)
}

type VpcSubnetArrayOutput struct{ *pulumi.OutputState }

func (VpcSubnetArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*VpcSubnet)(nil)).Elem()
}

func (o VpcSubnetArrayOutput) ToVpcSubnetArrayOutput() VpcSubnetArrayOutput {
	return o
}

func (o VpcSubnetArrayOutput) ToVpcSubnetArrayOutputWithContext(ctx context.Context) VpcSubnetArrayOutput {
	return o
}

func (o VpcSubnetArrayOutput) Index(i pulumi.IntInput) VpcSubnetOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *VpcSubnet {
		return vs[0].([]*VpcSubnet)[vs[1].(int)]
	}).(VpcSubnetOutput)
}

type VpcSubnetMapOutput struct{ *pulumi.OutputState }

func (VpcSubnetMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*VpcSubnet)(nil)).Elem()
}

func (o VpcSubnetMapOutput) ToVpcSubnetMapOutput() VpcSubnetMapOutput {
	return o
}

func (o VpcSubnetMapOutput) ToVpcSubnetMapOutputWithContext(ctx context.Context) VpcSubnetMapOutput {
	return o
}

func (o VpcSubnetMapOutput) MapIndex(k pulumi.StringInput) VpcSubnetOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *VpcSubnet {
		return vs[0].(map[string]*VpcSubnet)[vs[1].(string)]
	}).(VpcSubnetOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*VpcSubnetInput)(nil)).Elem(), &VpcSubnet{})
	pulumi.RegisterInputType(reflect.TypeOf((*VpcSubnetArrayInput)(nil)).Elem(), VpcSubnetArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*VpcSubnetMapInput)(nil)).Elem(), VpcSubnetMap{})
	pulumi.RegisterOutputType(VpcSubnetOutput{})
	pulumi.RegisterOutputType(VpcSubnetArrayOutput{})
	pulumi.RegisterOutputType(VpcSubnetMapOutput{})
}
