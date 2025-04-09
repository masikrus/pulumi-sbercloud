// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package sbercloud

import (
	"context"
	"reflect"

	"github.com/masikrus/pulumi-sbercloud/sdk/go/sbercloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type IdentityGroup struct {
	pulumi.CustomResourceState

	Description pulumi.StringPtrOutput `pulumi:"description"`
	Name        pulumi.StringOutput    `pulumi:"name"`
}

// NewIdentityGroup registers a new resource with the given unique name, arguments, and options.
func NewIdentityGroup(ctx *pulumi.Context,
	name string, args *IdentityGroupArgs, opts ...pulumi.ResourceOption) (*IdentityGroup, error) {
	if args == nil {
		args = &IdentityGroupArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource IdentityGroup
	err := ctx.RegisterResource("sbercloud:index/identityGroup:IdentityGroup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetIdentityGroup gets an existing IdentityGroup resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetIdentityGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *IdentityGroupState, opts ...pulumi.ResourceOption) (*IdentityGroup, error) {
	var resource IdentityGroup
	err := ctx.ReadResource("sbercloud:index/identityGroup:IdentityGroup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering IdentityGroup resources.
type identityGroupState struct {
	Description *string `pulumi:"description"`
	Name        *string `pulumi:"name"`
}

type IdentityGroupState struct {
	Description pulumi.StringPtrInput
	Name        pulumi.StringPtrInput
}

func (IdentityGroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*identityGroupState)(nil)).Elem()
}

type identityGroupArgs struct {
	Description *string `pulumi:"description"`
	Name        *string `pulumi:"name"`
}

// The set of arguments for constructing a IdentityGroup resource.
type IdentityGroupArgs struct {
	Description pulumi.StringPtrInput
	Name        pulumi.StringPtrInput
}

func (IdentityGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*identityGroupArgs)(nil)).Elem()
}

type IdentityGroupInput interface {
	pulumi.Input

	ToIdentityGroupOutput() IdentityGroupOutput
	ToIdentityGroupOutputWithContext(ctx context.Context) IdentityGroupOutput
}

func (*IdentityGroup) ElementType() reflect.Type {
	return reflect.TypeOf((**IdentityGroup)(nil)).Elem()
}

func (i *IdentityGroup) ToIdentityGroupOutput() IdentityGroupOutput {
	return i.ToIdentityGroupOutputWithContext(context.Background())
}

func (i *IdentityGroup) ToIdentityGroupOutputWithContext(ctx context.Context) IdentityGroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IdentityGroupOutput)
}

// IdentityGroupArrayInput is an input type that accepts IdentityGroupArray and IdentityGroupArrayOutput values.
// You can construct a concrete instance of `IdentityGroupArrayInput` via:
//
//	IdentityGroupArray{ IdentityGroupArgs{...} }
type IdentityGroupArrayInput interface {
	pulumi.Input

	ToIdentityGroupArrayOutput() IdentityGroupArrayOutput
	ToIdentityGroupArrayOutputWithContext(context.Context) IdentityGroupArrayOutput
}

type IdentityGroupArray []IdentityGroupInput

func (IdentityGroupArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*IdentityGroup)(nil)).Elem()
}

func (i IdentityGroupArray) ToIdentityGroupArrayOutput() IdentityGroupArrayOutput {
	return i.ToIdentityGroupArrayOutputWithContext(context.Background())
}

func (i IdentityGroupArray) ToIdentityGroupArrayOutputWithContext(ctx context.Context) IdentityGroupArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IdentityGroupArrayOutput)
}

// IdentityGroupMapInput is an input type that accepts IdentityGroupMap and IdentityGroupMapOutput values.
// You can construct a concrete instance of `IdentityGroupMapInput` via:
//
//	IdentityGroupMap{ "key": IdentityGroupArgs{...} }
type IdentityGroupMapInput interface {
	pulumi.Input

	ToIdentityGroupMapOutput() IdentityGroupMapOutput
	ToIdentityGroupMapOutputWithContext(context.Context) IdentityGroupMapOutput
}

type IdentityGroupMap map[string]IdentityGroupInput

func (IdentityGroupMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*IdentityGroup)(nil)).Elem()
}

func (i IdentityGroupMap) ToIdentityGroupMapOutput() IdentityGroupMapOutput {
	return i.ToIdentityGroupMapOutputWithContext(context.Background())
}

func (i IdentityGroupMap) ToIdentityGroupMapOutputWithContext(ctx context.Context) IdentityGroupMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IdentityGroupMapOutput)
}

type IdentityGroupOutput struct{ *pulumi.OutputState }

func (IdentityGroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IdentityGroup)(nil)).Elem()
}

func (o IdentityGroupOutput) ToIdentityGroupOutput() IdentityGroupOutput {
	return o
}

func (o IdentityGroupOutput) ToIdentityGroupOutputWithContext(ctx context.Context) IdentityGroupOutput {
	return o
}

func (o IdentityGroupOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IdentityGroup) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

func (o IdentityGroupOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *IdentityGroup) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

type IdentityGroupArrayOutput struct{ *pulumi.OutputState }

func (IdentityGroupArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*IdentityGroup)(nil)).Elem()
}

func (o IdentityGroupArrayOutput) ToIdentityGroupArrayOutput() IdentityGroupArrayOutput {
	return o
}

func (o IdentityGroupArrayOutput) ToIdentityGroupArrayOutputWithContext(ctx context.Context) IdentityGroupArrayOutput {
	return o
}

func (o IdentityGroupArrayOutput) Index(i pulumi.IntInput) IdentityGroupOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *IdentityGroup {
		return vs[0].([]*IdentityGroup)[vs[1].(int)]
	}).(IdentityGroupOutput)
}

type IdentityGroupMapOutput struct{ *pulumi.OutputState }

func (IdentityGroupMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*IdentityGroup)(nil)).Elem()
}

func (o IdentityGroupMapOutput) ToIdentityGroupMapOutput() IdentityGroupMapOutput {
	return o
}

func (o IdentityGroupMapOutput) ToIdentityGroupMapOutputWithContext(ctx context.Context) IdentityGroupMapOutput {
	return o
}

func (o IdentityGroupMapOutput) MapIndex(k pulumi.StringInput) IdentityGroupOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *IdentityGroup {
		return vs[0].(map[string]*IdentityGroup)[vs[1].(string)]
	}).(IdentityGroupOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*IdentityGroupInput)(nil)).Elem(), &IdentityGroup{})
	pulumi.RegisterInputType(reflect.TypeOf((*IdentityGroupArrayInput)(nil)).Elem(), IdentityGroupArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*IdentityGroupMapInput)(nil)).Elem(), IdentityGroupMap{})
	pulumi.RegisterOutputType(IdentityGroupOutput{})
	pulumi.RegisterOutputType(IdentityGroupArrayOutput{})
	pulumi.RegisterOutputType(IdentityGroupMapOutput{})
}
