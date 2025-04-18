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

type ApigEnvironment struct {
	pulumi.CustomResourceState

	// schema: Deprecated; The time when the environment was created.
	//
	// Deprecated: Use 'created_at' instead
	CreateTime pulumi.StringOutput `pulumi:"createTime"`
	// The time when the environment was created.
	CreatedAt pulumi.StringOutput `pulumi:"createdAt"`
	// The environment description.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The ID of the dedicated instance to which the environment belongs.
	InstanceId pulumi.StringOutput `pulumi:"instanceId"`
	// The environment name.
	Name pulumi.StringOutput `pulumi:"name"`
	// The region where the dedicated instance is located.
	Region pulumi.StringOutput `pulumi:"region"`
}

// NewApigEnvironment registers a new resource with the given unique name, arguments, and options.
func NewApigEnvironment(ctx *pulumi.Context,
	name string, args *ApigEnvironmentArgs, opts ...pulumi.ResourceOption) (*ApigEnvironment, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.InstanceId == nil {
		return nil, errors.New("invalid value for required argument 'InstanceId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource ApigEnvironment
	err := ctx.RegisterResource("sbercloud:index/apigEnvironment:ApigEnvironment", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetApigEnvironment gets an existing ApigEnvironment resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetApigEnvironment(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ApigEnvironmentState, opts ...pulumi.ResourceOption) (*ApigEnvironment, error) {
	var resource ApigEnvironment
	err := ctx.ReadResource("sbercloud:index/apigEnvironment:ApigEnvironment", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ApigEnvironment resources.
type apigEnvironmentState struct {
	// schema: Deprecated; The time when the environment was created.
	//
	// Deprecated: Use 'created_at' instead
	CreateTime *string `pulumi:"createTime"`
	// The time when the environment was created.
	CreatedAt *string `pulumi:"createdAt"`
	// The environment description.
	Description *string `pulumi:"description"`
	// The ID of the dedicated instance to which the environment belongs.
	InstanceId *string `pulumi:"instanceId"`
	// The environment name.
	Name *string `pulumi:"name"`
	// The region where the dedicated instance is located.
	Region *string `pulumi:"region"`
}

type ApigEnvironmentState struct {
	// schema: Deprecated; The time when the environment was created.
	//
	// Deprecated: Use 'created_at' instead
	CreateTime pulumi.StringPtrInput
	// The time when the environment was created.
	CreatedAt pulumi.StringPtrInput
	// The environment description.
	Description pulumi.StringPtrInput
	// The ID of the dedicated instance to which the environment belongs.
	InstanceId pulumi.StringPtrInput
	// The environment name.
	Name pulumi.StringPtrInput
	// The region where the dedicated instance is located.
	Region pulumi.StringPtrInput
}

func (ApigEnvironmentState) ElementType() reflect.Type {
	return reflect.TypeOf((*apigEnvironmentState)(nil)).Elem()
}

type apigEnvironmentArgs struct {
	// The environment description.
	Description *string `pulumi:"description"`
	// The ID of the dedicated instance to which the environment belongs.
	InstanceId string `pulumi:"instanceId"`
	// The environment name.
	Name *string `pulumi:"name"`
	// The region where the dedicated instance is located.
	Region *string `pulumi:"region"`
}

// The set of arguments for constructing a ApigEnvironment resource.
type ApigEnvironmentArgs struct {
	// The environment description.
	Description pulumi.StringPtrInput
	// The ID of the dedicated instance to which the environment belongs.
	InstanceId pulumi.StringInput
	// The environment name.
	Name pulumi.StringPtrInput
	// The region where the dedicated instance is located.
	Region pulumi.StringPtrInput
}

func (ApigEnvironmentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*apigEnvironmentArgs)(nil)).Elem()
}

type ApigEnvironmentInput interface {
	pulumi.Input

	ToApigEnvironmentOutput() ApigEnvironmentOutput
	ToApigEnvironmentOutputWithContext(ctx context.Context) ApigEnvironmentOutput
}

func (*ApigEnvironment) ElementType() reflect.Type {
	return reflect.TypeOf((**ApigEnvironment)(nil)).Elem()
}

func (i *ApigEnvironment) ToApigEnvironmentOutput() ApigEnvironmentOutput {
	return i.ToApigEnvironmentOutputWithContext(context.Background())
}

func (i *ApigEnvironment) ToApigEnvironmentOutputWithContext(ctx context.Context) ApigEnvironmentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApigEnvironmentOutput)
}

// ApigEnvironmentArrayInput is an input type that accepts ApigEnvironmentArray and ApigEnvironmentArrayOutput values.
// You can construct a concrete instance of `ApigEnvironmentArrayInput` via:
//
//	ApigEnvironmentArray{ ApigEnvironmentArgs{...} }
type ApigEnvironmentArrayInput interface {
	pulumi.Input

	ToApigEnvironmentArrayOutput() ApigEnvironmentArrayOutput
	ToApigEnvironmentArrayOutputWithContext(context.Context) ApigEnvironmentArrayOutput
}

type ApigEnvironmentArray []ApigEnvironmentInput

func (ApigEnvironmentArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ApigEnvironment)(nil)).Elem()
}

func (i ApigEnvironmentArray) ToApigEnvironmentArrayOutput() ApigEnvironmentArrayOutput {
	return i.ToApigEnvironmentArrayOutputWithContext(context.Background())
}

func (i ApigEnvironmentArray) ToApigEnvironmentArrayOutputWithContext(ctx context.Context) ApigEnvironmentArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApigEnvironmentArrayOutput)
}

// ApigEnvironmentMapInput is an input type that accepts ApigEnvironmentMap and ApigEnvironmentMapOutput values.
// You can construct a concrete instance of `ApigEnvironmentMapInput` via:
//
//	ApigEnvironmentMap{ "key": ApigEnvironmentArgs{...} }
type ApigEnvironmentMapInput interface {
	pulumi.Input

	ToApigEnvironmentMapOutput() ApigEnvironmentMapOutput
	ToApigEnvironmentMapOutputWithContext(context.Context) ApigEnvironmentMapOutput
}

type ApigEnvironmentMap map[string]ApigEnvironmentInput

func (ApigEnvironmentMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ApigEnvironment)(nil)).Elem()
}

func (i ApigEnvironmentMap) ToApigEnvironmentMapOutput() ApigEnvironmentMapOutput {
	return i.ToApigEnvironmentMapOutputWithContext(context.Background())
}

func (i ApigEnvironmentMap) ToApigEnvironmentMapOutputWithContext(ctx context.Context) ApigEnvironmentMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApigEnvironmentMapOutput)
}

type ApigEnvironmentOutput struct{ *pulumi.OutputState }

func (ApigEnvironmentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ApigEnvironment)(nil)).Elem()
}

func (o ApigEnvironmentOutput) ToApigEnvironmentOutput() ApigEnvironmentOutput {
	return o
}

func (o ApigEnvironmentOutput) ToApigEnvironmentOutputWithContext(ctx context.Context) ApigEnvironmentOutput {
	return o
}

// schema: Deprecated; The time when the environment was created.
//
// Deprecated: Use 'created_at' instead
func (o ApigEnvironmentOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *ApigEnvironment) pulumi.StringOutput { return v.CreateTime }).(pulumi.StringOutput)
}

// The time when the environment was created.
func (o ApigEnvironmentOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *ApigEnvironment) pulumi.StringOutput { return v.CreatedAt }).(pulumi.StringOutput)
}

// The environment description.
func (o ApigEnvironmentOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApigEnvironment) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// The ID of the dedicated instance to which the environment belongs.
func (o ApigEnvironmentOutput) InstanceId() pulumi.StringOutput {
	return o.ApplyT(func(v *ApigEnvironment) pulumi.StringOutput { return v.InstanceId }).(pulumi.StringOutput)
}

// The environment name.
func (o ApigEnvironmentOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ApigEnvironment) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The region where the dedicated instance is located.
func (o ApigEnvironmentOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v *ApigEnvironment) pulumi.StringOutput { return v.Region }).(pulumi.StringOutput)
}

type ApigEnvironmentArrayOutput struct{ *pulumi.OutputState }

func (ApigEnvironmentArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ApigEnvironment)(nil)).Elem()
}

func (o ApigEnvironmentArrayOutput) ToApigEnvironmentArrayOutput() ApigEnvironmentArrayOutput {
	return o
}

func (o ApigEnvironmentArrayOutput) ToApigEnvironmentArrayOutputWithContext(ctx context.Context) ApigEnvironmentArrayOutput {
	return o
}

func (o ApigEnvironmentArrayOutput) Index(i pulumi.IntInput) ApigEnvironmentOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ApigEnvironment {
		return vs[0].([]*ApigEnvironment)[vs[1].(int)]
	}).(ApigEnvironmentOutput)
}

type ApigEnvironmentMapOutput struct{ *pulumi.OutputState }

func (ApigEnvironmentMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ApigEnvironment)(nil)).Elem()
}

func (o ApigEnvironmentMapOutput) ToApigEnvironmentMapOutput() ApigEnvironmentMapOutput {
	return o
}

func (o ApigEnvironmentMapOutput) ToApigEnvironmentMapOutputWithContext(ctx context.Context) ApigEnvironmentMapOutput {
	return o
}

func (o ApigEnvironmentMapOutput) MapIndex(k pulumi.StringInput) ApigEnvironmentOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ApigEnvironment {
		return vs[0].(map[string]*ApigEnvironment)[vs[1].(string)]
	}).(ApigEnvironmentOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ApigEnvironmentInput)(nil)).Elem(), &ApigEnvironment{})
	pulumi.RegisterInputType(reflect.TypeOf((*ApigEnvironmentArrayInput)(nil)).Elem(), ApigEnvironmentArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ApigEnvironmentMapInput)(nil)).Elem(), ApigEnvironmentMap{})
	pulumi.RegisterOutputType(ApigEnvironmentOutput{})
	pulumi.RegisterOutputType(ApigEnvironmentArrayOutput{})
	pulumi.RegisterOutputType(ApigEnvironmentMapOutput{})
}
