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

type ElbIpgroup struct {
	pulumi.CustomResourceState

	CreatedAt           pulumi.StringOutput         `pulumi:"createdAt"`
	Description         pulumi.StringPtrOutput      `pulumi:"description"`
	EnterpriseProjectId pulumi.StringOutput         `pulumi:"enterpriseProjectId"`
	IpLists             ElbIpgroupIpListArrayOutput `pulumi:"ipLists"`
	ListenerIds         pulumi.StringArrayOutput    `pulumi:"listenerIds"`
	Name                pulumi.StringOutput         `pulumi:"name"`
	Region              pulumi.StringOutput         `pulumi:"region"`
	UpdatedAt           pulumi.StringOutput         `pulumi:"updatedAt"`
}

// NewElbIpgroup registers a new resource with the given unique name, arguments, and options.
func NewElbIpgroup(ctx *pulumi.Context,
	name string, args *ElbIpgroupArgs, opts ...pulumi.ResourceOption) (*ElbIpgroup, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.IpLists == nil {
		return nil, errors.New("invalid value for required argument 'IpLists'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource ElbIpgroup
	err := ctx.RegisterResource("sbercloud:index/elbIpgroup:ElbIpgroup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetElbIpgroup gets an existing ElbIpgroup resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetElbIpgroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ElbIpgroupState, opts ...pulumi.ResourceOption) (*ElbIpgroup, error) {
	var resource ElbIpgroup
	err := ctx.ReadResource("sbercloud:index/elbIpgroup:ElbIpgroup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ElbIpgroup resources.
type elbIpgroupState struct {
	CreatedAt           *string            `pulumi:"createdAt"`
	Description         *string            `pulumi:"description"`
	EnterpriseProjectId *string            `pulumi:"enterpriseProjectId"`
	IpLists             []ElbIpgroupIpList `pulumi:"ipLists"`
	ListenerIds         []string           `pulumi:"listenerIds"`
	Name                *string            `pulumi:"name"`
	Region              *string            `pulumi:"region"`
	UpdatedAt           *string            `pulumi:"updatedAt"`
}

type ElbIpgroupState struct {
	CreatedAt           pulumi.StringPtrInput
	Description         pulumi.StringPtrInput
	EnterpriseProjectId pulumi.StringPtrInput
	IpLists             ElbIpgroupIpListArrayInput
	ListenerIds         pulumi.StringArrayInput
	Name                pulumi.StringPtrInput
	Region              pulumi.StringPtrInput
	UpdatedAt           pulumi.StringPtrInput
}

func (ElbIpgroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*elbIpgroupState)(nil)).Elem()
}

type elbIpgroupArgs struct {
	Description         *string            `pulumi:"description"`
	EnterpriseProjectId *string            `pulumi:"enterpriseProjectId"`
	IpLists             []ElbIpgroupIpList `pulumi:"ipLists"`
	Name                *string            `pulumi:"name"`
	Region              *string            `pulumi:"region"`
}

// The set of arguments for constructing a ElbIpgroup resource.
type ElbIpgroupArgs struct {
	Description         pulumi.StringPtrInput
	EnterpriseProjectId pulumi.StringPtrInput
	IpLists             ElbIpgroupIpListArrayInput
	Name                pulumi.StringPtrInput
	Region              pulumi.StringPtrInput
}

func (ElbIpgroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*elbIpgroupArgs)(nil)).Elem()
}

type ElbIpgroupInput interface {
	pulumi.Input

	ToElbIpgroupOutput() ElbIpgroupOutput
	ToElbIpgroupOutputWithContext(ctx context.Context) ElbIpgroupOutput
}

func (*ElbIpgroup) ElementType() reflect.Type {
	return reflect.TypeOf((**ElbIpgroup)(nil)).Elem()
}

func (i *ElbIpgroup) ToElbIpgroupOutput() ElbIpgroupOutput {
	return i.ToElbIpgroupOutputWithContext(context.Background())
}

func (i *ElbIpgroup) ToElbIpgroupOutputWithContext(ctx context.Context) ElbIpgroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ElbIpgroupOutput)
}

// ElbIpgroupArrayInput is an input type that accepts ElbIpgroupArray and ElbIpgroupArrayOutput values.
// You can construct a concrete instance of `ElbIpgroupArrayInput` via:
//
//	ElbIpgroupArray{ ElbIpgroupArgs{...} }
type ElbIpgroupArrayInput interface {
	pulumi.Input

	ToElbIpgroupArrayOutput() ElbIpgroupArrayOutput
	ToElbIpgroupArrayOutputWithContext(context.Context) ElbIpgroupArrayOutput
}

type ElbIpgroupArray []ElbIpgroupInput

func (ElbIpgroupArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ElbIpgroup)(nil)).Elem()
}

func (i ElbIpgroupArray) ToElbIpgroupArrayOutput() ElbIpgroupArrayOutput {
	return i.ToElbIpgroupArrayOutputWithContext(context.Background())
}

func (i ElbIpgroupArray) ToElbIpgroupArrayOutputWithContext(ctx context.Context) ElbIpgroupArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ElbIpgroupArrayOutput)
}

// ElbIpgroupMapInput is an input type that accepts ElbIpgroupMap and ElbIpgroupMapOutput values.
// You can construct a concrete instance of `ElbIpgroupMapInput` via:
//
//	ElbIpgroupMap{ "key": ElbIpgroupArgs{...} }
type ElbIpgroupMapInput interface {
	pulumi.Input

	ToElbIpgroupMapOutput() ElbIpgroupMapOutput
	ToElbIpgroupMapOutputWithContext(context.Context) ElbIpgroupMapOutput
}

type ElbIpgroupMap map[string]ElbIpgroupInput

func (ElbIpgroupMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ElbIpgroup)(nil)).Elem()
}

func (i ElbIpgroupMap) ToElbIpgroupMapOutput() ElbIpgroupMapOutput {
	return i.ToElbIpgroupMapOutputWithContext(context.Background())
}

func (i ElbIpgroupMap) ToElbIpgroupMapOutputWithContext(ctx context.Context) ElbIpgroupMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ElbIpgroupMapOutput)
}

type ElbIpgroupOutput struct{ *pulumi.OutputState }

func (ElbIpgroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ElbIpgroup)(nil)).Elem()
}

func (o ElbIpgroupOutput) ToElbIpgroupOutput() ElbIpgroupOutput {
	return o
}

func (o ElbIpgroupOutput) ToElbIpgroupOutputWithContext(ctx context.Context) ElbIpgroupOutput {
	return o
}

func (o ElbIpgroupOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *ElbIpgroup) pulumi.StringOutput { return v.CreatedAt }).(pulumi.StringOutput)
}

func (o ElbIpgroupOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ElbIpgroup) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

func (o ElbIpgroupOutput) EnterpriseProjectId() pulumi.StringOutput {
	return o.ApplyT(func(v *ElbIpgroup) pulumi.StringOutput { return v.EnterpriseProjectId }).(pulumi.StringOutput)
}

func (o ElbIpgroupOutput) IpLists() ElbIpgroupIpListArrayOutput {
	return o.ApplyT(func(v *ElbIpgroup) ElbIpgroupIpListArrayOutput { return v.IpLists }).(ElbIpgroupIpListArrayOutput)
}

func (o ElbIpgroupOutput) ListenerIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ElbIpgroup) pulumi.StringArrayOutput { return v.ListenerIds }).(pulumi.StringArrayOutput)
}

func (o ElbIpgroupOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ElbIpgroup) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o ElbIpgroupOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v *ElbIpgroup) pulumi.StringOutput { return v.Region }).(pulumi.StringOutput)
}

func (o ElbIpgroupOutput) UpdatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *ElbIpgroup) pulumi.StringOutput { return v.UpdatedAt }).(pulumi.StringOutput)
}

type ElbIpgroupArrayOutput struct{ *pulumi.OutputState }

func (ElbIpgroupArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ElbIpgroup)(nil)).Elem()
}

func (o ElbIpgroupArrayOutput) ToElbIpgroupArrayOutput() ElbIpgroupArrayOutput {
	return o
}

func (o ElbIpgroupArrayOutput) ToElbIpgroupArrayOutputWithContext(ctx context.Context) ElbIpgroupArrayOutput {
	return o
}

func (o ElbIpgroupArrayOutput) Index(i pulumi.IntInput) ElbIpgroupOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ElbIpgroup {
		return vs[0].([]*ElbIpgroup)[vs[1].(int)]
	}).(ElbIpgroupOutput)
}

type ElbIpgroupMapOutput struct{ *pulumi.OutputState }

func (ElbIpgroupMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ElbIpgroup)(nil)).Elem()
}

func (o ElbIpgroupMapOutput) ToElbIpgroupMapOutput() ElbIpgroupMapOutput {
	return o
}

func (o ElbIpgroupMapOutput) ToElbIpgroupMapOutputWithContext(ctx context.Context) ElbIpgroupMapOutput {
	return o
}

func (o ElbIpgroupMapOutput) MapIndex(k pulumi.StringInput) ElbIpgroupOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ElbIpgroup {
		return vs[0].(map[string]*ElbIpgroup)[vs[1].(string)]
	}).(ElbIpgroupOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ElbIpgroupInput)(nil)).Elem(), &ElbIpgroup{})
	pulumi.RegisterInputType(reflect.TypeOf((*ElbIpgroupArrayInput)(nil)).Elem(), ElbIpgroupArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ElbIpgroupMapInput)(nil)).Elem(), ElbIpgroupMap{})
	pulumi.RegisterOutputType(ElbIpgroupOutput{})
	pulumi.RegisterOutputType(ElbIpgroupArrayOutput{})
	pulumi.RegisterOutputType(ElbIpgroupMapOutput{})
}
