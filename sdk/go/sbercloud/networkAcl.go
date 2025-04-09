// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package sbercloud

import (
	"context"
	"reflect"

	"github.com/masikrus/pulumi-sbercloud/sdk/go/sbercloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type NetworkAcl struct {
	pulumi.CustomResourceState

	Description      pulumi.StringPtrOutput   `pulumi:"description"`
	InboundPolicyId  pulumi.StringOutput      `pulumi:"inboundPolicyId"`
	InboundRules     pulumi.StringArrayOutput `pulumi:"inboundRules"`
	Name             pulumi.StringOutput      `pulumi:"name"`
	OutboundPolicyId pulumi.StringOutput      `pulumi:"outboundPolicyId"`
	OutboundRules    pulumi.StringArrayOutput `pulumi:"outboundRules"`
	Ports            pulumi.StringArrayOutput `pulumi:"ports"`
	Region           pulumi.StringOutput      `pulumi:"region"`
	Status           pulumi.StringOutput      `pulumi:"status"`
	Subnets          pulumi.StringArrayOutput `pulumi:"subnets"`
}

// NewNetworkAcl registers a new resource with the given unique name, arguments, and options.
func NewNetworkAcl(ctx *pulumi.Context,
	name string, args *NetworkAclArgs, opts ...pulumi.ResourceOption) (*NetworkAcl, error) {
	if args == nil {
		args = &NetworkAclArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource NetworkAcl
	err := ctx.RegisterResource("sbercloud:index/networkAcl:NetworkAcl", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetNetworkAcl gets an existing NetworkAcl resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetNetworkAcl(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *NetworkAclState, opts ...pulumi.ResourceOption) (*NetworkAcl, error) {
	var resource NetworkAcl
	err := ctx.ReadResource("sbercloud:index/networkAcl:NetworkAcl", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering NetworkAcl resources.
type networkAclState struct {
	Description      *string  `pulumi:"description"`
	InboundPolicyId  *string  `pulumi:"inboundPolicyId"`
	InboundRules     []string `pulumi:"inboundRules"`
	Name             *string  `pulumi:"name"`
	OutboundPolicyId *string  `pulumi:"outboundPolicyId"`
	OutboundRules    []string `pulumi:"outboundRules"`
	Ports            []string `pulumi:"ports"`
	Region           *string  `pulumi:"region"`
	Status           *string  `pulumi:"status"`
	Subnets          []string `pulumi:"subnets"`
}

type NetworkAclState struct {
	Description      pulumi.StringPtrInput
	InboundPolicyId  pulumi.StringPtrInput
	InboundRules     pulumi.StringArrayInput
	Name             pulumi.StringPtrInput
	OutboundPolicyId pulumi.StringPtrInput
	OutboundRules    pulumi.StringArrayInput
	Ports            pulumi.StringArrayInput
	Region           pulumi.StringPtrInput
	Status           pulumi.StringPtrInput
	Subnets          pulumi.StringArrayInput
}

func (NetworkAclState) ElementType() reflect.Type {
	return reflect.TypeOf((*networkAclState)(nil)).Elem()
}

type networkAclArgs struct {
	Description   *string  `pulumi:"description"`
	InboundRules  []string `pulumi:"inboundRules"`
	Name          *string  `pulumi:"name"`
	OutboundRules []string `pulumi:"outboundRules"`
	Region        *string  `pulumi:"region"`
	Subnets       []string `pulumi:"subnets"`
}

// The set of arguments for constructing a NetworkAcl resource.
type NetworkAclArgs struct {
	Description   pulumi.StringPtrInput
	InboundRules  pulumi.StringArrayInput
	Name          pulumi.StringPtrInput
	OutboundRules pulumi.StringArrayInput
	Region        pulumi.StringPtrInput
	Subnets       pulumi.StringArrayInput
}

func (NetworkAclArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*networkAclArgs)(nil)).Elem()
}

type NetworkAclInput interface {
	pulumi.Input

	ToNetworkAclOutput() NetworkAclOutput
	ToNetworkAclOutputWithContext(ctx context.Context) NetworkAclOutput
}

func (*NetworkAcl) ElementType() reflect.Type {
	return reflect.TypeOf((**NetworkAcl)(nil)).Elem()
}

func (i *NetworkAcl) ToNetworkAclOutput() NetworkAclOutput {
	return i.ToNetworkAclOutputWithContext(context.Background())
}

func (i *NetworkAcl) ToNetworkAclOutputWithContext(ctx context.Context) NetworkAclOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkAclOutput)
}

// NetworkAclArrayInput is an input type that accepts NetworkAclArray and NetworkAclArrayOutput values.
// You can construct a concrete instance of `NetworkAclArrayInput` via:
//
//	NetworkAclArray{ NetworkAclArgs{...} }
type NetworkAclArrayInput interface {
	pulumi.Input

	ToNetworkAclArrayOutput() NetworkAclArrayOutput
	ToNetworkAclArrayOutputWithContext(context.Context) NetworkAclArrayOutput
}

type NetworkAclArray []NetworkAclInput

func (NetworkAclArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*NetworkAcl)(nil)).Elem()
}

func (i NetworkAclArray) ToNetworkAclArrayOutput() NetworkAclArrayOutput {
	return i.ToNetworkAclArrayOutputWithContext(context.Background())
}

func (i NetworkAclArray) ToNetworkAclArrayOutputWithContext(ctx context.Context) NetworkAclArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkAclArrayOutput)
}

// NetworkAclMapInput is an input type that accepts NetworkAclMap and NetworkAclMapOutput values.
// You can construct a concrete instance of `NetworkAclMapInput` via:
//
//	NetworkAclMap{ "key": NetworkAclArgs{...} }
type NetworkAclMapInput interface {
	pulumi.Input

	ToNetworkAclMapOutput() NetworkAclMapOutput
	ToNetworkAclMapOutputWithContext(context.Context) NetworkAclMapOutput
}

type NetworkAclMap map[string]NetworkAclInput

func (NetworkAclMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*NetworkAcl)(nil)).Elem()
}

func (i NetworkAclMap) ToNetworkAclMapOutput() NetworkAclMapOutput {
	return i.ToNetworkAclMapOutputWithContext(context.Background())
}

func (i NetworkAclMap) ToNetworkAclMapOutputWithContext(ctx context.Context) NetworkAclMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkAclMapOutput)
}

type NetworkAclOutput struct{ *pulumi.OutputState }

func (NetworkAclOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NetworkAcl)(nil)).Elem()
}

func (o NetworkAclOutput) ToNetworkAclOutput() NetworkAclOutput {
	return o
}

func (o NetworkAclOutput) ToNetworkAclOutputWithContext(ctx context.Context) NetworkAclOutput {
	return o
}

func (o NetworkAclOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NetworkAcl) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

func (o NetworkAclOutput) InboundPolicyId() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkAcl) pulumi.StringOutput { return v.InboundPolicyId }).(pulumi.StringOutput)
}

func (o NetworkAclOutput) InboundRules() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *NetworkAcl) pulumi.StringArrayOutput { return v.InboundRules }).(pulumi.StringArrayOutput)
}

func (o NetworkAclOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkAcl) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o NetworkAclOutput) OutboundPolicyId() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkAcl) pulumi.StringOutput { return v.OutboundPolicyId }).(pulumi.StringOutput)
}

func (o NetworkAclOutput) OutboundRules() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *NetworkAcl) pulumi.StringArrayOutput { return v.OutboundRules }).(pulumi.StringArrayOutput)
}

func (o NetworkAclOutput) Ports() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *NetworkAcl) pulumi.StringArrayOutput { return v.Ports }).(pulumi.StringArrayOutput)
}

func (o NetworkAclOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkAcl) pulumi.StringOutput { return v.Region }).(pulumi.StringOutput)
}

func (o NetworkAclOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkAcl) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

func (o NetworkAclOutput) Subnets() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *NetworkAcl) pulumi.StringArrayOutput { return v.Subnets }).(pulumi.StringArrayOutput)
}

type NetworkAclArrayOutput struct{ *pulumi.OutputState }

func (NetworkAclArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*NetworkAcl)(nil)).Elem()
}

func (o NetworkAclArrayOutput) ToNetworkAclArrayOutput() NetworkAclArrayOutput {
	return o
}

func (o NetworkAclArrayOutput) ToNetworkAclArrayOutputWithContext(ctx context.Context) NetworkAclArrayOutput {
	return o
}

func (o NetworkAclArrayOutput) Index(i pulumi.IntInput) NetworkAclOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *NetworkAcl {
		return vs[0].([]*NetworkAcl)[vs[1].(int)]
	}).(NetworkAclOutput)
}

type NetworkAclMapOutput struct{ *pulumi.OutputState }

func (NetworkAclMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*NetworkAcl)(nil)).Elem()
}

func (o NetworkAclMapOutput) ToNetworkAclMapOutput() NetworkAclMapOutput {
	return o
}

func (o NetworkAclMapOutput) ToNetworkAclMapOutputWithContext(ctx context.Context) NetworkAclMapOutput {
	return o
}

func (o NetworkAclMapOutput) MapIndex(k pulumi.StringInput) NetworkAclOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *NetworkAcl {
		return vs[0].(map[string]*NetworkAcl)[vs[1].(string)]
	}).(NetworkAclOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*NetworkAclInput)(nil)).Elem(), &NetworkAcl{})
	pulumi.RegisterInputType(reflect.TypeOf((*NetworkAclArrayInput)(nil)).Elem(), NetworkAclArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*NetworkAclMapInput)(nil)).Elem(), NetworkAclMap{})
	pulumi.RegisterOutputType(NetworkAclOutput{})
	pulumi.RegisterOutputType(NetworkAclArrayOutput{})
	pulumi.RegisterOutputType(NetworkAclMapOutput{})
}
