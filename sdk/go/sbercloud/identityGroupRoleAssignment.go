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

type IdentityGroupRoleAssignment struct {
	pulumi.CustomResourceState

	DomainId            pulumi.StringPtrOutput `pulumi:"domainId"`
	EnterpriseProjectId pulumi.StringPtrOutput `pulumi:"enterpriseProjectId"`
	GroupId             pulumi.StringOutput    `pulumi:"groupId"`
	ProjectId           pulumi.StringPtrOutput `pulumi:"projectId"`
	RoleId              pulumi.StringOutput    `pulumi:"roleId"`
}

// NewIdentityGroupRoleAssignment registers a new resource with the given unique name, arguments, and options.
func NewIdentityGroupRoleAssignment(ctx *pulumi.Context,
	name string, args *IdentityGroupRoleAssignmentArgs, opts ...pulumi.ResourceOption) (*IdentityGroupRoleAssignment, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.GroupId == nil {
		return nil, errors.New("invalid value for required argument 'GroupId'")
	}
	if args.RoleId == nil {
		return nil, errors.New("invalid value for required argument 'RoleId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource IdentityGroupRoleAssignment
	err := ctx.RegisterResource("sbercloud:index/identityGroupRoleAssignment:IdentityGroupRoleAssignment", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetIdentityGroupRoleAssignment gets an existing IdentityGroupRoleAssignment resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetIdentityGroupRoleAssignment(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *IdentityGroupRoleAssignmentState, opts ...pulumi.ResourceOption) (*IdentityGroupRoleAssignment, error) {
	var resource IdentityGroupRoleAssignment
	err := ctx.ReadResource("sbercloud:index/identityGroupRoleAssignment:IdentityGroupRoleAssignment", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering IdentityGroupRoleAssignment resources.
type identityGroupRoleAssignmentState struct {
	DomainId            *string `pulumi:"domainId"`
	EnterpriseProjectId *string `pulumi:"enterpriseProjectId"`
	GroupId             *string `pulumi:"groupId"`
	ProjectId           *string `pulumi:"projectId"`
	RoleId              *string `pulumi:"roleId"`
}

type IdentityGroupRoleAssignmentState struct {
	DomainId            pulumi.StringPtrInput
	EnterpriseProjectId pulumi.StringPtrInput
	GroupId             pulumi.StringPtrInput
	ProjectId           pulumi.StringPtrInput
	RoleId              pulumi.StringPtrInput
}

func (IdentityGroupRoleAssignmentState) ElementType() reflect.Type {
	return reflect.TypeOf((*identityGroupRoleAssignmentState)(nil)).Elem()
}

type identityGroupRoleAssignmentArgs struct {
	DomainId            *string `pulumi:"domainId"`
	EnterpriseProjectId *string `pulumi:"enterpriseProjectId"`
	GroupId             string  `pulumi:"groupId"`
	ProjectId           *string `pulumi:"projectId"`
	RoleId              string  `pulumi:"roleId"`
}

// The set of arguments for constructing a IdentityGroupRoleAssignment resource.
type IdentityGroupRoleAssignmentArgs struct {
	DomainId            pulumi.StringPtrInput
	EnterpriseProjectId pulumi.StringPtrInput
	GroupId             pulumi.StringInput
	ProjectId           pulumi.StringPtrInput
	RoleId              pulumi.StringInput
}

func (IdentityGroupRoleAssignmentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*identityGroupRoleAssignmentArgs)(nil)).Elem()
}

type IdentityGroupRoleAssignmentInput interface {
	pulumi.Input

	ToIdentityGroupRoleAssignmentOutput() IdentityGroupRoleAssignmentOutput
	ToIdentityGroupRoleAssignmentOutputWithContext(ctx context.Context) IdentityGroupRoleAssignmentOutput
}

func (*IdentityGroupRoleAssignment) ElementType() reflect.Type {
	return reflect.TypeOf((**IdentityGroupRoleAssignment)(nil)).Elem()
}

func (i *IdentityGroupRoleAssignment) ToIdentityGroupRoleAssignmentOutput() IdentityGroupRoleAssignmentOutput {
	return i.ToIdentityGroupRoleAssignmentOutputWithContext(context.Background())
}

func (i *IdentityGroupRoleAssignment) ToIdentityGroupRoleAssignmentOutputWithContext(ctx context.Context) IdentityGroupRoleAssignmentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IdentityGroupRoleAssignmentOutput)
}

// IdentityGroupRoleAssignmentArrayInput is an input type that accepts IdentityGroupRoleAssignmentArray and IdentityGroupRoleAssignmentArrayOutput values.
// You can construct a concrete instance of `IdentityGroupRoleAssignmentArrayInput` via:
//
//	IdentityGroupRoleAssignmentArray{ IdentityGroupRoleAssignmentArgs{...} }
type IdentityGroupRoleAssignmentArrayInput interface {
	pulumi.Input

	ToIdentityGroupRoleAssignmentArrayOutput() IdentityGroupRoleAssignmentArrayOutput
	ToIdentityGroupRoleAssignmentArrayOutputWithContext(context.Context) IdentityGroupRoleAssignmentArrayOutput
}

type IdentityGroupRoleAssignmentArray []IdentityGroupRoleAssignmentInput

func (IdentityGroupRoleAssignmentArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*IdentityGroupRoleAssignment)(nil)).Elem()
}

func (i IdentityGroupRoleAssignmentArray) ToIdentityGroupRoleAssignmentArrayOutput() IdentityGroupRoleAssignmentArrayOutput {
	return i.ToIdentityGroupRoleAssignmentArrayOutputWithContext(context.Background())
}

func (i IdentityGroupRoleAssignmentArray) ToIdentityGroupRoleAssignmentArrayOutputWithContext(ctx context.Context) IdentityGroupRoleAssignmentArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IdentityGroupRoleAssignmentArrayOutput)
}

// IdentityGroupRoleAssignmentMapInput is an input type that accepts IdentityGroupRoleAssignmentMap and IdentityGroupRoleAssignmentMapOutput values.
// You can construct a concrete instance of `IdentityGroupRoleAssignmentMapInput` via:
//
//	IdentityGroupRoleAssignmentMap{ "key": IdentityGroupRoleAssignmentArgs{...} }
type IdentityGroupRoleAssignmentMapInput interface {
	pulumi.Input

	ToIdentityGroupRoleAssignmentMapOutput() IdentityGroupRoleAssignmentMapOutput
	ToIdentityGroupRoleAssignmentMapOutputWithContext(context.Context) IdentityGroupRoleAssignmentMapOutput
}

type IdentityGroupRoleAssignmentMap map[string]IdentityGroupRoleAssignmentInput

func (IdentityGroupRoleAssignmentMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*IdentityGroupRoleAssignment)(nil)).Elem()
}

func (i IdentityGroupRoleAssignmentMap) ToIdentityGroupRoleAssignmentMapOutput() IdentityGroupRoleAssignmentMapOutput {
	return i.ToIdentityGroupRoleAssignmentMapOutputWithContext(context.Background())
}

func (i IdentityGroupRoleAssignmentMap) ToIdentityGroupRoleAssignmentMapOutputWithContext(ctx context.Context) IdentityGroupRoleAssignmentMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IdentityGroupRoleAssignmentMapOutput)
}

type IdentityGroupRoleAssignmentOutput struct{ *pulumi.OutputState }

func (IdentityGroupRoleAssignmentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IdentityGroupRoleAssignment)(nil)).Elem()
}

func (o IdentityGroupRoleAssignmentOutput) ToIdentityGroupRoleAssignmentOutput() IdentityGroupRoleAssignmentOutput {
	return o
}

func (o IdentityGroupRoleAssignmentOutput) ToIdentityGroupRoleAssignmentOutputWithContext(ctx context.Context) IdentityGroupRoleAssignmentOutput {
	return o
}

func (o IdentityGroupRoleAssignmentOutput) DomainId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IdentityGroupRoleAssignment) pulumi.StringPtrOutput { return v.DomainId }).(pulumi.StringPtrOutput)
}

func (o IdentityGroupRoleAssignmentOutput) EnterpriseProjectId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IdentityGroupRoleAssignment) pulumi.StringPtrOutput { return v.EnterpriseProjectId }).(pulumi.StringPtrOutput)
}

func (o IdentityGroupRoleAssignmentOutput) GroupId() pulumi.StringOutput {
	return o.ApplyT(func(v *IdentityGroupRoleAssignment) pulumi.StringOutput { return v.GroupId }).(pulumi.StringOutput)
}

func (o IdentityGroupRoleAssignmentOutput) ProjectId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IdentityGroupRoleAssignment) pulumi.StringPtrOutput { return v.ProjectId }).(pulumi.StringPtrOutput)
}

func (o IdentityGroupRoleAssignmentOutput) RoleId() pulumi.StringOutput {
	return o.ApplyT(func(v *IdentityGroupRoleAssignment) pulumi.StringOutput { return v.RoleId }).(pulumi.StringOutput)
}

type IdentityGroupRoleAssignmentArrayOutput struct{ *pulumi.OutputState }

func (IdentityGroupRoleAssignmentArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*IdentityGroupRoleAssignment)(nil)).Elem()
}

func (o IdentityGroupRoleAssignmentArrayOutput) ToIdentityGroupRoleAssignmentArrayOutput() IdentityGroupRoleAssignmentArrayOutput {
	return o
}

func (o IdentityGroupRoleAssignmentArrayOutput) ToIdentityGroupRoleAssignmentArrayOutputWithContext(ctx context.Context) IdentityGroupRoleAssignmentArrayOutput {
	return o
}

func (o IdentityGroupRoleAssignmentArrayOutput) Index(i pulumi.IntInput) IdentityGroupRoleAssignmentOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *IdentityGroupRoleAssignment {
		return vs[0].([]*IdentityGroupRoleAssignment)[vs[1].(int)]
	}).(IdentityGroupRoleAssignmentOutput)
}

type IdentityGroupRoleAssignmentMapOutput struct{ *pulumi.OutputState }

func (IdentityGroupRoleAssignmentMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*IdentityGroupRoleAssignment)(nil)).Elem()
}

func (o IdentityGroupRoleAssignmentMapOutput) ToIdentityGroupRoleAssignmentMapOutput() IdentityGroupRoleAssignmentMapOutput {
	return o
}

func (o IdentityGroupRoleAssignmentMapOutput) ToIdentityGroupRoleAssignmentMapOutputWithContext(ctx context.Context) IdentityGroupRoleAssignmentMapOutput {
	return o
}

func (o IdentityGroupRoleAssignmentMapOutput) MapIndex(k pulumi.StringInput) IdentityGroupRoleAssignmentOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *IdentityGroupRoleAssignment {
		return vs[0].(map[string]*IdentityGroupRoleAssignment)[vs[1].(string)]
	}).(IdentityGroupRoleAssignmentOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*IdentityGroupRoleAssignmentInput)(nil)).Elem(), &IdentityGroupRoleAssignment{})
	pulumi.RegisterInputType(reflect.TypeOf((*IdentityGroupRoleAssignmentArrayInput)(nil)).Elem(), IdentityGroupRoleAssignmentArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*IdentityGroupRoleAssignmentMapInput)(nil)).Elem(), IdentityGroupRoleAssignmentMap{})
	pulumi.RegisterOutputType(IdentityGroupRoleAssignmentOutput{})
	pulumi.RegisterOutputType(IdentityGroupRoleAssignmentArrayOutput{})
	pulumi.RegisterOutputType(IdentityGroupRoleAssignmentMapOutput{})
}
