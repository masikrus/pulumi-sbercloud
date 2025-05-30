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

type IdentityGroupMembership struct {
	pulumi.CustomResourceState

	Group pulumi.StringOutput      `pulumi:"group"`
	Users pulumi.StringArrayOutput `pulumi:"users"`
}

// NewIdentityGroupMembership registers a new resource with the given unique name, arguments, and options.
func NewIdentityGroupMembership(ctx *pulumi.Context,
	name string, args *IdentityGroupMembershipArgs, opts ...pulumi.ResourceOption) (*IdentityGroupMembership, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Group == nil {
		return nil, errors.New("invalid value for required argument 'Group'")
	}
	if args.Users == nil {
		return nil, errors.New("invalid value for required argument 'Users'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource IdentityGroupMembership
	err := ctx.RegisterResource("sbercloud:index/identityGroupMembership:IdentityGroupMembership", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetIdentityGroupMembership gets an existing IdentityGroupMembership resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetIdentityGroupMembership(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *IdentityGroupMembershipState, opts ...pulumi.ResourceOption) (*IdentityGroupMembership, error) {
	var resource IdentityGroupMembership
	err := ctx.ReadResource("sbercloud:index/identityGroupMembership:IdentityGroupMembership", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering IdentityGroupMembership resources.
type identityGroupMembershipState struct {
	Group *string  `pulumi:"group"`
	Users []string `pulumi:"users"`
}

type IdentityGroupMembershipState struct {
	Group pulumi.StringPtrInput
	Users pulumi.StringArrayInput
}

func (IdentityGroupMembershipState) ElementType() reflect.Type {
	return reflect.TypeOf((*identityGroupMembershipState)(nil)).Elem()
}

type identityGroupMembershipArgs struct {
	Group string   `pulumi:"group"`
	Users []string `pulumi:"users"`
}

// The set of arguments for constructing a IdentityGroupMembership resource.
type IdentityGroupMembershipArgs struct {
	Group pulumi.StringInput
	Users pulumi.StringArrayInput
}

func (IdentityGroupMembershipArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*identityGroupMembershipArgs)(nil)).Elem()
}

type IdentityGroupMembershipInput interface {
	pulumi.Input

	ToIdentityGroupMembershipOutput() IdentityGroupMembershipOutput
	ToIdentityGroupMembershipOutputWithContext(ctx context.Context) IdentityGroupMembershipOutput
}

func (*IdentityGroupMembership) ElementType() reflect.Type {
	return reflect.TypeOf((**IdentityGroupMembership)(nil)).Elem()
}

func (i *IdentityGroupMembership) ToIdentityGroupMembershipOutput() IdentityGroupMembershipOutput {
	return i.ToIdentityGroupMembershipOutputWithContext(context.Background())
}

func (i *IdentityGroupMembership) ToIdentityGroupMembershipOutputWithContext(ctx context.Context) IdentityGroupMembershipOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IdentityGroupMembershipOutput)
}

// IdentityGroupMembershipArrayInput is an input type that accepts IdentityGroupMembershipArray and IdentityGroupMembershipArrayOutput values.
// You can construct a concrete instance of `IdentityGroupMembershipArrayInput` via:
//
//	IdentityGroupMembershipArray{ IdentityGroupMembershipArgs{...} }
type IdentityGroupMembershipArrayInput interface {
	pulumi.Input

	ToIdentityGroupMembershipArrayOutput() IdentityGroupMembershipArrayOutput
	ToIdentityGroupMembershipArrayOutputWithContext(context.Context) IdentityGroupMembershipArrayOutput
}

type IdentityGroupMembershipArray []IdentityGroupMembershipInput

func (IdentityGroupMembershipArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*IdentityGroupMembership)(nil)).Elem()
}

func (i IdentityGroupMembershipArray) ToIdentityGroupMembershipArrayOutput() IdentityGroupMembershipArrayOutput {
	return i.ToIdentityGroupMembershipArrayOutputWithContext(context.Background())
}

func (i IdentityGroupMembershipArray) ToIdentityGroupMembershipArrayOutputWithContext(ctx context.Context) IdentityGroupMembershipArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IdentityGroupMembershipArrayOutput)
}

// IdentityGroupMembershipMapInput is an input type that accepts IdentityGroupMembershipMap and IdentityGroupMembershipMapOutput values.
// You can construct a concrete instance of `IdentityGroupMembershipMapInput` via:
//
//	IdentityGroupMembershipMap{ "key": IdentityGroupMembershipArgs{...} }
type IdentityGroupMembershipMapInput interface {
	pulumi.Input

	ToIdentityGroupMembershipMapOutput() IdentityGroupMembershipMapOutput
	ToIdentityGroupMembershipMapOutputWithContext(context.Context) IdentityGroupMembershipMapOutput
}

type IdentityGroupMembershipMap map[string]IdentityGroupMembershipInput

func (IdentityGroupMembershipMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*IdentityGroupMembership)(nil)).Elem()
}

func (i IdentityGroupMembershipMap) ToIdentityGroupMembershipMapOutput() IdentityGroupMembershipMapOutput {
	return i.ToIdentityGroupMembershipMapOutputWithContext(context.Background())
}

func (i IdentityGroupMembershipMap) ToIdentityGroupMembershipMapOutputWithContext(ctx context.Context) IdentityGroupMembershipMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IdentityGroupMembershipMapOutput)
}

type IdentityGroupMembershipOutput struct{ *pulumi.OutputState }

func (IdentityGroupMembershipOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IdentityGroupMembership)(nil)).Elem()
}

func (o IdentityGroupMembershipOutput) ToIdentityGroupMembershipOutput() IdentityGroupMembershipOutput {
	return o
}

func (o IdentityGroupMembershipOutput) ToIdentityGroupMembershipOutputWithContext(ctx context.Context) IdentityGroupMembershipOutput {
	return o
}

func (o IdentityGroupMembershipOutput) Group() pulumi.StringOutput {
	return o.ApplyT(func(v *IdentityGroupMembership) pulumi.StringOutput { return v.Group }).(pulumi.StringOutput)
}

func (o IdentityGroupMembershipOutput) Users() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *IdentityGroupMembership) pulumi.StringArrayOutput { return v.Users }).(pulumi.StringArrayOutput)
}

type IdentityGroupMembershipArrayOutput struct{ *pulumi.OutputState }

func (IdentityGroupMembershipArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*IdentityGroupMembership)(nil)).Elem()
}

func (o IdentityGroupMembershipArrayOutput) ToIdentityGroupMembershipArrayOutput() IdentityGroupMembershipArrayOutput {
	return o
}

func (o IdentityGroupMembershipArrayOutput) ToIdentityGroupMembershipArrayOutputWithContext(ctx context.Context) IdentityGroupMembershipArrayOutput {
	return o
}

func (o IdentityGroupMembershipArrayOutput) Index(i pulumi.IntInput) IdentityGroupMembershipOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *IdentityGroupMembership {
		return vs[0].([]*IdentityGroupMembership)[vs[1].(int)]
	}).(IdentityGroupMembershipOutput)
}

type IdentityGroupMembershipMapOutput struct{ *pulumi.OutputState }

func (IdentityGroupMembershipMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*IdentityGroupMembership)(nil)).Elem()
}

func (o IdentityGroupMembershipMapOutput) ToIdentityGroupMembershipMapOutput() IdentityGroupMembershipMapOutput {
	return o
}

func (o IdentityGroupMembershipMapOutput) ToIdentityGroupMembershipMapOutputWithContext(ctx context.Context) IdentityGroupMembershipMapOutput {
	return o
}

func (o IdentityGroupMembershipMapOutput) MapIndex(k pulumi.StringInput) IdentityGroupMembershipOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *IdentityGroupMembership {
		return vs[0].(map[string]*IdentityGroupMembership)[vs[1].(string)]
	}).(IdentityGroupMembershipOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*IdentityGroupMembershipInput)(nil)).Elem(), &IdentityGroupMembership{})
	pulumi.RegisterInputType(reflect.TypeOf((*IdentityGroupMembershipArrayInput)(nil)).Elem(), IdentityGroupMembershipArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*IdentityGroupMembershipMapInput)(nil)).Elem(), IdentityGroupMembershipMap{})
	pulumi.RegisterOutputType(IdentityGroupMembershipOutput{})
	pulumi.RegisterOutputType(IdentityGroupMembershipArrayOutput{})
	pulumi.RegisterOutputType(IdentityGroupMembershipMapOutput{})
}
