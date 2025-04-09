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

type ApigInstanceFeature struct {
	pulumi.CustomResourceState

	// Specified the detailed configuration of the feature.
	Config pulumi.StringPtrOutput `pulumi:"config"`
	// Specified whether to enable the feature.
	Enabled pulumi.BoolPtrOutput `pulumi:"enabled"`
	// Specified the ID of the dedicated instance to which the feature belongs.
	InstanceId pulumi.StringOutput `pulumi:"instanceId"`
	// Specified the name of the feature.
	Name   pulumi.StringOutput `pulumi:"name"`
	Region pulumi.StringOutput `pulumi:"region"`
}

// NewApigInstanceFeature registers a new resource with the given unique name, arguments, and options.
func NewApigInstanceFeature(ctx *pulumi.Context,
	name string, args *ApigInstanceFeatureArgs, opts ...pulumi.ResourceOption) (*ApigInstanceFeature, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.InstanceId == nil {
		return nil, errors.New("invalid value for required argument 'InstanceId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource ApigInstanceFeature
	err := ctx.RegisterResource("sbercloud:index/apigInstanceFeature:ApigInstanceFeature", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetApigInstanceFeature gets an existing ApigInstanceFeature resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetApigInstanceFeature(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ApigInstanceFeatureState, opts ...pulumi.ResourceOption) (*ApigInstanceFeature, error) {
	var resource ApigInstanceFeature
	err := ctx.ReadResource("sbercloud:index/apigInstanceFeature:ApigInstanceFeature", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ApigInstanceFeature resources.
type apigInstanceFeatureState struct {
	// Specified the detailed configuration of the feature.
	Config *string `pulumi:"config"`
	// Specified whether to enable the feature.
	Enabled *bool `pulumi:"enabled"`
	// Specified the ID of the dedicated instance to which the feature belongs.
	InstanceId *string `pulumi:"instanceId"`
	// Specified the name of the feature.
	Name   *string `pulumi:"name"`
	Region *string `pulumi:"region"`
}

type ApigInstanceFeatureState struct {
	// Specified the detailed configuration of the feature.
	Config pulumi.StringPtrInput
	// Specified whether to enable the feature.
	Enabled pulumi.BoolPtrInput
	// Specified the ID of the dedicated instance to which the feature belongs.
	InstanceId pulumi.StringPtrInput
	// Specified the name of the feature.
	Name   pulumi.StringPtrInput
	Region pulumi.StringPtrInput
}

func (ApigInstanceFeatureState) ElementType() reflect.Type {
	return reflect.TypeOf((*apigInstanceFeatureState)(nil)).Elem()
}

type apigInstanceFeatureArgs struct {
	// Specified the detailed configuration of the feature.
	Config *string `pulumi:"config"`
	// Specified whether to enable the feature.
	Enabled *bool `pulumi:"enabled"`
	// Specified the ID of the dedicated instance to which the feature belongs.
	InstanceId string `pulumi:"instanceId"`
	// Specified the name of the feature.
	Name   *string `pulumi:"name"`
	Region *string `pulumi:"region"`
}

// The set of arguments for constructing a ApigInstanceFeature resource.
type ApigInstanceFeatureArgs struct {
	// Specified the detailed configuration of the feature.
	Config pulumi.StringPtrInput
	// Specified whether to enable the feature.
	Enabled pulumi.BoolPtrInput
	// Specified the ID of the dedicated instance to which the feature belongs.
	InstanceId pulumi.StringInput
	// Specified the name of the feature.
	Name   pulumi.StringPtrInput
	Region pulumi.StringPtrInput
}

func (ApigInstanceFeatureArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*apigInstanceFeatureArgs)(nil)).Elem()
}

type ApigInstanceFeatureInput interface {
	pulumi.Input

	ToApigInstanceFeatureOutput() ApigInstanceFeatureOutput
	ToApigInstanceFeatureOutputWithContext(ctx context.Context) ApigInstanceFeatureOutput
}

func (*ApigInstanceFeature) ElementType() reflect.Type {
	return reflect.TypeOf((**ApigInstanceFeature)(nil)).Elem()
}

func (i *ApigInstanceFeature) ToApigInstanceFeatureOutput() ApigInstanceFeatureOutput {
	return i.ToApigInstanceFeatureOutputWithContext(context.Background())
}

func (i *ApigInstanceFeature) ToApigInstanceFeatureOutputWithContext(ctx context.Context) ApigInstanceFeatureOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApigInstanceFeatureOutput)
}

// ApigInstanceFeatureArrayInput is an input type that accepts ApigInstanceFeatureArray and ApigInstanceFeatureArrayOutput values.
// You can construct a concrete instance of `ApigInstanceFeatureArrayInput` via:
//
//	ApigInstanceFeatureArray{ ApigInstanceFeatureArgs{...} }
type ApigInstanceFeatureArrayInput interface {
	pulumi.Input

	ToApigInstanceFeatureArrayOutput() ApigInstanceFeatureArrayOutput
	ToApigInstanceFeatureArrayOutputWithContext(context.Context) ApigInstanceFeatureArrayOutput
}

type ApigInstanceFeatureArray []ApigInstanceFeatureInput

func (ApigInstanceFeatureArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ApigInstanceFeature)(nil)).Elem()
}

func (i ApigInstanceFeatureArray) ToApigInstanceFeatureArrayOutput() ApigInstanceFeatureArrayOutput {
	return i.ToApigInstanceFeatureArrayOutputWithContext(context.Background())
}

func (i ApigInstanceFeatureArray) ToApigInstanceFeatureArrayOutputWithContext(ctx context.Context) ApigInstanceFeatureArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApigInstanceFeatureArrayOutput)
}

// ApigInstanceFeatureMapInput is an input type that accepts ApigInstanceFeatureMap and ApigInstanceFeatureMapOutput values.
// You can construct a concrete instance of `ApigInstanceFeatureMapInput` via:
//
//	ApigInstanceFeatureMap{ "key": ApigInstanceFeatureArgs{...} }
type ApigInstanceFeatureMapInput interface {
	pulumi.Input

	ToApigInstanceFeatureMapOutput() ApigInstanceFeatureMapOutput
	ToApigInstanceFeatureMapOutputWithContext(context.Context) ApigInstanceFeatureMapOutput
}

type ApigInstanceFeatureMap map[string]ApigInstanceFeatureInput

func (ApigInstanceFeatureMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ApigInstanceFeature)(nil)).Elem()
}

func (i ApigInstanceFeatureMap) ToApigInstanceFeatureMapOutput() ApigInstanceFeatureMapOutput {
	return i.ToApigInstanceFeatureMapOutputWithContext(context.Background())
}

func (i ApigInstanceFeatureMap) ToApigInstanceFeatureMapOutputWithContext(ctx context.Context) ApigInstanceFeatureMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApigInstanceFeatureMapOutput)
}

type ApigInstanceFeatureOutput struct{ *pulumi.OutputState }

func (ApigInstanceFeatureOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ApigInstanceFeature)(nil)).Elem()
}

func (o ApigInstanceFeatureOutput) ToApigInstanceFeatureOutput() ApigInstanceFeatureOutput {
	return o
}

func (o ApigInstanceFeatureOutput) ToApigInstanceFeatureOutputWithContext(ctx context.Context) ApigInstanceFeatureOutput {
	return o
}

// Specified the detailed configuration of the feature.
func (o ApigInstanceFeatureOutput) Config() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApigInstanceFeature) pulumi.StringPtrOutput { return v.Config }).(pulumi.StringPtrOutput)
}

// Specified whether to enable the feature.
func (o ApigInstanceFeatureOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ApigInstanceFeature) pulumi.BoolPtrOutput { return v.Enabled }).(pulumi.BoolPtrOutput)
}

// Specified the ID of the dedicated instance to which the feature belongs.
func (o ApigInstanceFeatureOutput) InstanceId() pulumi.StringOutput {
	return o.ApplyT(func(v *ApigInstanceFeature) pulumi.StringOutput { return v.InstanceId }).(pulumi.StringOutput)
}

// Specified the name of the feature.
func (o ApigInstanceFeatureOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ApigInstanceFeature) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o ApigInstanceFeatureOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v *ApigInstanceFeature) pulumi.StringOutput { return v.Region }).(pulumi.StringOutput)
}

type ApigInstanceFeatureArrayOutput struct{ *pulumi.OutputState }

func (ApigInstanceFeatureArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ApigInstanceFeature)(nil)).Elem()
}

func (o ApigInstanceFeatureArrayOutput) ToApigInstanceFeatureArrayOutput() ApigInstanceFeatureArrayOutput {
	return o
}

func (o ApigInstanceFeatureArrayOutput) ToApigInstanceFeatureArrayOutputWithContext(ctx context.Context) ApigInstanceFeatureArrayOutput {
	return o
}

func (o ApigInstanceFeatureArrayOutput) Index(i pulumi.IntInput) ApigInstanceFeatureOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ApigInstanceFeature {
		return vs[0].([]*ApigInstanceFeature)[vs[1].(int)]
	}).(ApigInstanceFeatureOutput)
}

type ApigInstanceFeatureMapOutput struct{ *pulumi.OutputState }

func (ApigInstanceFeatureMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ApigInstanceFeature)(nil)).Elem()
}

func (o ApigInstanceFeatureMapOutput) ToApigInstanceFeatureMapOutput() ApigInstanceFeatureMapOutput {
	return o
}

func (o ApigInstanceFeatureMapOutput) ToApigInstanceFeatureMapOutputWithContext(ctx context.Context) ApigInstanceFeatureMapOutput {
	return o
}

func (o ApigInstanceFeatureMapOutput) MapIndex(k pulumi.StringInput) ApigInstanceFeatureOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ApigInstanceFeature {
		return vs[0].(map[string]*ApigInstanceFeature)[vs[1].(string)]
	}).(ApigInstanceFeatureOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ApigInstanceFeatureInput)(nil)).Elem(), &ApigInstanceFeature{})
	pulumi.RegisterInputType(reflect.TypeOf((*ApigInstanceFeatureArrayInput)(nil)).Elem(), ApigInstanceFeatureArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ApigInstanceFeatureMapInput)(nil)).Elem(), ApigInstanceFeatureMap{})
	pulumi.RegisterOutputType(ApigInstanceFeatureOutput{})
	pulumi.RegisterOutputType(ApigInstanceFeatureArrayOutput{})
	pulumi.RegisterOutputType(ApigInstanceFeatureMapOutput{})
}
