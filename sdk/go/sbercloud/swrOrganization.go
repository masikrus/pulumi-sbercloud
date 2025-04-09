// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package sbercloud

import (
	"context"
	"reflect"

	"github.com/masikrus/pulumi-sbercloud/sdk/go/sbercloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type SwrOrganization struct {
	pulumi.CustomResourceState

	Creator     pulumi.StringOutput `pulumi:"creator"`
	LoginServer pulumi.StringOutput `pulumi:"loginServer"`
	Name        pulumi.StringOutput `pulumi:"name"`
	Permission  pulumi.StringOutput `pulumi:"permission"`
	Region      pulumi.StringOutput `pulumi:"region"`
}

// NewSwrOrganization registers a new resource with the given unique name, arguments, and options.
func NewSwrOrganization(ctx *pulumi.Context,
	name string, args *SwrOrganizationArgs, opts ...pulumi.ResourceOption) (*SwrOrganization, error) {
	if args == nil {
		args = &SwrOrganizationArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource SwrOrganization
	err := ctx.RegisterResource("sbercloud:index/swrOrganization:SwrOrganization", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSwrOrganization gets an existing SwrOrganization resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSwrOrganization(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SwrOrganizationState, opts ...pulumi.ResourceOption) (*SwrOrganization, error) {
	var resource SwrOrganization
	err := ctx.ReadResource("sbercloud:index/swrOrganization:SwrOrganization", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SwrOrganization resources.
type swrOrganizationState struct {
	Creator     *string `pulumi:"creator"`
	LoginServer *string `pulumi:"loginServer"`
	Name        *string `pulumi:"name"`
	Permission  *string `pulumi:"permission"`
	Region      *string `pulumi:"region"`
}

type SwrOrganizationState struct {
	Creator     pulumi.StringPtrInput
	LoginServer pulumi.StringPtrInput
	Name        pulumi.StringPtrInput
	Permission  pulumi.StringPtrInput
	Region      pulumi.StringPtrInput
}

func (SwrOrganizationState) ElementType() reflect.Type {
	return reflect.TypeOf((*swrOrganizationState)(nil)).Elem()
}

type swrOrganizationArgs struct {
	Name   *string `pulumi:"name"`
	Region *string `pulumi:"region"`
}

// The set of arguments for constructing a SwrOrganization resource.
type SwrOrganizationArgs struct {
	Name   pulumi.StringPtrInput
	Region pulumi.StringPtrInput
}

func (SwrOrganizationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*swrOrganizationArgs)(nil)).Elem()
}

type SwrOrganizationInput interface {
	pulumi.Input

	ToSwrOrganizationOutput() SwrOrganizationOutput
	ToSwrOrganizationOutputWithContext(ctx context.Context) SwrOrganizationOutput
}

func (*SwrOrganization) ElementType() reflect.Type {
	return reflect.TypeOf((**SwrOrganization)(nil)).Elem()
}

func (i *SwrOrganization) ToSwrOrganizationOutput() SwrOrganizationOutput {
	return i.ToSwrOrganizationOutputWithContext(context.Background())
}

func (i *SwrOrganization) ToSwrOrganizationOutputWithContext(ctx context.Context) SwrOrganizationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SwrOrganizationOutput)
}

// SwrOrganizationArrayInput is an input type that accepts SwrOrganizationArray and SwrOrganizationArrayOutput values.
// You can construct a concrete instance of `SwrOrganizationArrayInput` via:
//
//	SwrOrganizationArray{ SwrOrganizationArgs{...} }
type SwrOrganizationArrayInput interface {
	pulumi.Input

	ToSwrOrganizationArrayOutput() SwrOrganizationArrayOutput
	ToSwrOrganizationArrayOutputWithContext(context.Context) SwrOrganizationArrayOutput
}

type SwrOrganizationArray []SwrOrganizationInput

func (SwrOrganizationArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SwrOrganization)(nil)).Elem()
}

func (i SwrOrganizationArray) ToSwrOrganizationArrayOutput() SwrOrganizationArrayOutput {
	return i.ToSwrOrganizationArrayOutputWithContext(context.Background())
}

func (i SwrOrganizationArray) ToSwrOrganizationArrayOutputWithContext(ctx context.Context) SwrOrganizationArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SwrOrganizationArrayOutput)
}

// SwrOrganizationMapInput is an input type that accepts SwrOrganizationMap and SwrOrganizationMapOutput values.
// You can construct a concrete instance of `SwrOrganizationMapInput` via:
//
//	SwrOrganizationMap{ "key": SwrOrganizationArgs{...} }
type SwrOrganizationMapInput interface {
	pulumi.Input

	ToSwrOrganizationMapOutput() SwrOrganizationMapOutput
	ToSwrOrganizationMapOutputWithContext(context.Context) SwrOrganizationMapOutput
}

type SwrOrganizationMap map[string]SwrOrganizationInput

func (SwrOrganizationMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SwrOrganization)(nil)).Elem()
}

func (i SwrOrganizationMap) ToSwrOrganizationMapOutput() SwrOrganizationMapOutput {
	return i.ToSwrOrganizationMapOutputWithContext(context.Background())
}

func (i SwrOrganizationMap) ToSwrOrganizationMapOutputWithContext(ctx context.Context) SwrOrganizationMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SwrOrganizationMapOutput)
}

type SwrOrganizationOutput struct{ *pulumi.OutputState }

func (SwrOrganizationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SwrOrganization)(nil)).Elem()
}

func (o SwrOrganizationOutput) ToSwrOrganizationOutput() SwrOrganizationOutput {
	return o
}

func (o SwrOrganizationOutput) ToSwrOrganizationOutputWithContext(ctx context.Context) SwrOrganizationOutput {
	return o
}

func (o SwrOrganizationOutput) Creator() pulumi.StringOutput {
	return o.ApplyT(func(v *SwrOrganization) pulumi.StringOutput { return v.Creator }).(pulumi.StringOutput)
}

func (o SwrOrganizationOutput) LoginServer() pulumi.StringOutput {
	return o.ApplyT(func(v *SwrOrganization) pulumi.StringOutput { return v.LoginServer }).(pulumi.StringOutput)
}

func (o SwrOrganizationOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *SwrOrganization) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o SwrOrganizationOutput) Permission() pulumi.StringOutput {
	return o.ApplyT(func(v *SwrOrganization) pulumi.StringOutput { return v.Permission }).(pulumi.StringOutput)
}

func (o SwrOrganizationOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v *SwrOrganization) pulumi.StringOutput { return v.Region }).(pulumi.StringOutput)
}

type SwrOrganizationArrayOutput struct{ *pulumi.OutputState }

func (SwrOrganizationArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SwrOrganization)(nil)).Elem()
}

func (o SwrOrganizationArrayOutput) ToSwrOrganizationArrayOutput() SwrOrganizationArrayOutput {
	return o
}

func (o SwrOrganizationArrayOutput) ToSwrOrganizationArrayOutputWithContext(ctx context.Context) SwrOrganizationArrayOutput {
	return o
}

func (o SwrOrganizationArrayOutput) Index(i pulumi.IntInput) SwrOrganizationOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SwrOrganization {
		return vs[0].([]*SwrOrganization)[vs[1].(int)]
	}).(SwrOrganizationOutput)
}

type SwrOrganizationMapOutput struct{ *pulumi.OutputState }

func (SwrOrganizationMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SwrOrganization)(nil)).Elem()
}

func (o SwrOrganizationMapOutput) ToSwrOrganizationMapOutput() SwrOrganizationMapOutput {
	return o
}

func (o SwrOrganizationMapOutput) ToSwrOrganizationMapOutputWithContext(ctx context.Context) SwrOrganizationMapOutput {
	return o
}

func (o SwrOrganizationMapOutput) MapIndex(k pulumi.StringInput) SwrOrganizationOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SwrOrganization {
		return vs[0].(map[string]*SwrOrganization)[vs[1].(string)]
	}).(SwrOrganizationOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SwrOrganizationInput)(nil)).Elem(), &SwrOrganization{})
	pulumi.RegisterInputType(reflect.TypeOf((*SwrOrganizationArrayInput)(nil)).Elem(), SwrOrganizationArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SwrOrganizationMapInput)(nil)).Elem(), SwrOrganizationMap{})
	pulumi.RegisterOutputType(SwrOrganizationOutput{})
	pulumi.RegisterOutputType(SwrOrganizationArrayOutput{})
	pulumi.RegisterOutputType(SwrOrganizationMapOutput{})
}
