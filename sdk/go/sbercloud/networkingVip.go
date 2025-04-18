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

type NetworkingVip struct {
	pulumi.CustomResourceState

	DeviceOwner pulumi.StringOutput `pulumi:"deviceOwner"`
	IpAddress   pulumi.StringOutput `pulumi:"ipAddress"`
	IpVersion   pulumi.IntOutput    `pulumi:"ipVersion"`
	MacAddress  pulumi.StringOutput `pulumi:"macAddress"`
	Name        pulumi.StringOutput `pulumi:"name"`
	NetworkId   pulumi.StringOutput `pulumi:"networkId"`
	Region      pulumi.StringOutput `pulumi:"region"`
	Status      pulumi.StringOutput `pulumi:"status"`
	// Deprecated: use ipVersion instead
	SubnetId pulumi.StringOutput `pulumi:"subnetId"`
}

// NewNetworkingVip registers a new resource with the given unique name, arguments, and options.
func NewNetworkingVip(ctx *pulumi.Context,
	name string, args *NetworkingVipArgs, opts ...pulumi.ResourceOption) (*NetworkingVip, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.NetworkId == nil {
		return nil, errors.New("invalid value for required argument 'NetworkId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource NetworkingVip
	err := ctx.RegisterResource("sbercloud:index/networkingVip:NetworkingVip", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetNetworkingVip gets an existing NetworkingVip resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetNetworkingVip(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *NetworkingVipState, opts ...pulumi.ResourceOption) (*NetworkingVip, error) {
	var resource NetworkingVip
	err := ctx.ReadResource("sbercloud:index/networkingVip:NetworkingVip", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering NetworkingVip resources.
type networkingVipState struct {
	DeviceOwner *string `pulumi:"deviceOwner"`
	IpAddress   *string `pulumi:"ipAddress"`
	IpVersion   *int    `pulumi:"ipVersion"`
	MacAddress  *string `pulumi:"macAddress"`
	Name        *string `pulumi:"name"`
	NetworkId   *string `pulumi:"networkId"`
	Region      *string `pulumi:"region"`
	Status      *string `pulumi:"status"`
	// Deprecated: use ipVersion instead
	SubnetId *string `pulumi:"subnetId"`
}

type NetworkingVipState struct {
	DeviceOwner pulumi.StringPtrInput
	IpAddress   pulumi.StringPtrInput
	IpVersion   pulumi.IntPtrInput
	MacAddress  pulumi.StringPtrInput
	Name        pulumi.StringPtrInput
	NetworkId   pulumi.StringPtrInput
	Region      pulumi.StringPtrInput
	Status      pulumi.StringPtrInput
	// Deprecated: use ipVersion instead
	SubnetId pulumi.StringPtrInput
}

func (NetworkingVipState) ElementType() reflect.Type {
	return reflect.TypeOf((*networkingVipState)(nil)).Elem()
}

type networkingVipArgs struct {
	IpAddress *string `pulumi:"ipAddress"`
	IpVersion *int    `pulumi:"ipVersion"`
	Name      *string `pulumi:"name"`
	NetworkId string  `pulumi:"networkId"`
	Region    *string `pulumi:"region"`
	// Deprecated: use ipVersion instead
	SubnetId *string `pulumi:"subnetId"`
}

// The set of arguments for constructing a NetworkingVip resource.
type NetworkingVipArgs struct {
	IpAddress pulumi.StringPtrInput
	IpVersion pulumi.IntPtrInput
	Name      pulumi.StringPtrInput
	NetworkId pulumi.StringInput
	Region    pulumi.StringPtrInput
	// Deprecated: use ipVersion instead
	SubnetId pulumi.StringPtrInput
}

func (NetworkingVipArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*networkingVipArgs)(nil)).Elem()
}

type NetworkingVipInput interface {
	pulumi.Input

	ToNetworkingVipOutput() NetworkingVipOutput
	ToNetworkingVipOutputWithContext(ctx context.Context) NetworkingVipOutput
}

func (*NetworkingVip) ElementType() reflect.Type {
	return reflect.TypeOf((**NetworkingVip)(nil)).Elem()
}

func (i *NetworkingVip) ToNetworkingVipOutput() NetworkingVipOutput {
	return i.ToNetworkingVipOutputWithContext(context.Background())
}

func (i *NetworkingVip) ToNetworkingVipOutputWithContext(ctx context.Context) NetworkingVipOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkingVipOutput)
}

// NetworkingVipArrayInput is an input type that accepts NetworkingVipArray and NetworkingVipArrayOutput values.
// You can construct a concrete instance of `NetworkingVipArrayInput` via:
//
//	NetworkingVipArray{ NetworkingVipArgs{...} }
type NetworkingVipArrayInput interface {
	pulumi.Input

	ToNetworkingVipArrayOutput() NetworkingVipArrayOutput
	ToNetworkingVipArrayOutputWithContext(context.Context) NetworkingVipArrayOutput
}

type NetworkingVipArray []NetworkingVipInput

func (NetworkingVipArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*NetworkingVip)(nil)).Elem()
}

func (i NetworkingVipArray) ToNetworkingVipArrayOutput() NetworkingVipArrayOutput {
	return i.ToNetworkingVipArrayOutputWithContext(context.Background())
}

func (i NetworkingVipArray) ToNetworkingVipArrayOutputWithContext(ctx context.Context) NetworkingVipArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkingVipArrayOutput)
}

// NetworkingVipMapInput is an input type that accepts NetworkingVipMap and NetworkingVipMapOutput values.
// You can construct a concrete instance of `NetworkingVipMapInput` via:
//
//	NetworkingVipMap{ "key": NetworkingVipArgs{...} }
type NetworkingVipMapInput interface {
	pulumi.Input

	ToNetworkingVipMapOutput() NetworkingVipMapOutput
	ToNetworkingVipMapOutputWithContext(context.Context) NetworkingVipMapOutput
}

type NetworkingVipMap map[string]NetworkingVipInput

func (NetworkingVipMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*NetworkingVip)(nil)).Elem()
}

func (i NetworkingVipMap) ToNetworkingVipMapOutput() NetworkingVipMapOutput {
	return i.ToNetworkingVipMapOutputWithContext(context.Background())
}

func (i NetworkingVipMap) ToNetworkingVipMapOutputWithContext(ctx context.Context) NetworkingVipMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkingVipMapOutput)
}

type NetworkingVipOutput struct{ *pulumi.OutputState }

func (NetworkingVipOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NetworkingVip)(nil)).Elem()
}

func (o NetworkingVipOutput) ToNetworkingVipOutput() NetworkingVipOutput {
	return o
}

func (o NetworkingVipOutput) ToNetworkingVipOutputWithContext(ctx context.Context) NetworkingVipOutput {
	return o
}

func (o NetworkingVipOutput) DeviceOwner() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkingVip) pulumi.StringOutput { return v.DeviceOwner }).(pulumi.StringOutput)
}

func (o NetworkingVipOutput) IpAddress() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkingVip) pulumi.StringOutput { return v.IpAddress }).(pulumi.StringOutput)
}

func (o NetworkingVipOutput) IpVersion() pulumi.IntOutput {
	return o.ApplyT(func(v *NetworkingVip) pulumi.IntOutput { return v.IpVersion }).(pulumi.IntOutput)
}

func (o NetworkingVipOutput) MacAddress() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkingVip) pulumi.StringOutput { return v.MacAddress }).(pulumi.StringOutput)
}

func (o NetworkingVipOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkingVip) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o NetworkingVipOutput) NetworkId() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkingVip) pulumi.StringOutput { return v.NetworkId }).(pulumi.StringOutput)
}

func (o NetworkingVipOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkingVip) pulumi.StringOutput { return v.Region }).(pulumi.StringOutput)
}

func (o NetworkingVipOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkingVip) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

// Deprecated: use ipVersion instead
func (o NetworkingVipOutput) SubnetId() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkingVip) pulumi.StringOutput { return v.SubnetId }).(pulumi.StringOutput)
}

type NetworkingVipArrayOutput struct{ *pulumi.OutputState }

func (NetworkingVipArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*NetworkingVip)(nil)).Elem()
}

func (o NetworkingVipArrayOutput) ToNetworkingVipArrayOutput() NetworkingVipArrayOutput {
	return o
}

func (o NetworkingVipArrayOutput) ToNetworkingVipArrayOutputWithContext(ctx context.Context) NetworkingVipArrayOutput {
	return o
}

func (o NetworkingVipArrayOutput) Index(i pulumi.IntInput) NetworkingVipOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *NetworkingVip {
		return vs[0].([]*NetworkingVip)[vs[1].(int)]
	}).(NetworkingVipOutput)
}

type NetworkingVipMapOutput struct{ *pulumi.OutputState }

func (NetworkingVipMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*NetworkingVip)(nil)).Elem()
}

func (o NetworkingVipMapOutput) ToNetworkingVipMapOutput() NetworkingVipMapOutput {
	return o
}

func (o NetworkingVipMapOutput) ToNetworkingVipMapOutputWithContext(ctx context.Context) NetworkingVipMapOutput {
	return o
}

func (o NetworkingVipMapOutput) MapIndex(k pulumi.StringInput) NetworkingVipOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *NetworkingVip {
		return vs[0].(map[string]*NetworkingVip)[vs[1].(string)]
	}).(NetworkingVipOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*NetworkingVipInput)(nil)).Elem(), &NetworkingVip{})
	pulumi.RegisterInputType(reflect.TypeOf((*NetworkingVipArrayInput)(nil)).Elem(), NetworkingVipArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*NetworkingVipMapInput)(nil)).Elem(), NetworkingVipMap{})
	pulumi.RegisterOutputType(NetworkingVipOutput{})
	pulumi.RegisterOutputType(NetworkingVipArrayOutput{})
	pulumi.RegisterOutputType(NetworkingVipMapOutput{})
}
