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

type ApigEndpointConnectionManagement struct {
	pulumi.CustomResourceState

	// Specifies the operation type endpoint connection.
	Action pulumi.StringOutput `pulumi:"action"`
	// Specifies the ID of the endpoint connection.
	EndpointId pulumi.StringOutput `pulumi:"endpointId"`
	// Specifies the ID of the dedicated instance to which the endpoint connection belongs.
	InstanceId pulumi.StringOutput `pulumi:"instanceId"`
	Region     pulumi.StringOutput `pulumi:"region"`
	// The current ststus of the endpoint connection.
	Status pulumi.StringOutput `pulumi:"status"`
}

// NewApigEndpointConnectionManagement registers a new resource with the given unique name, arguments, and options.
func NewApigEndpointConnectionManagement(ctx *pulumi.Context,
	name string, args *ApigEndpointConnectionManagementArgs, opts ...pulumi.ResourceOption) (*ApigEndpointConnectionManagement, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Action == nil {
		return nil, errors.New("invalid value for required argument 'Action'")
	}
	if args.EndpointId == nil {
		return nil, errors.New("invalid value for required argument 'EndpointId'")
	}
	if args.InstanceId == nil {
		return nil, errors.New("invalid value for required argument 'InstanceId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource ApigEndpointConnectionManagement
	err := ctx.RegisterResource("sbercloud:index/apigEndpointConnectionManagement:ApigEndpointConnectionManagement", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetApigEndpointConnectionManagement gets an existing ApigEndpointConnectionManagement resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetApigEndpointConnectionManagement(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ApigEndpointConnectionManagementState, opts ...pulumi.ResourceOption) (*ApigEndpointConnectionManagement, error) {
	var resource ApigEndpointConnectionManagement
	err := ctx.ReadResource("sbercloud:index/apigEndpointConnectionManagement:ApigEndpointConnectionManagement", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ApigEndpointConnectionManagement resources.
type apigEndpointConnectionManagementState struct {
	// Specifies the operation type endpoint connection.
	Action *string `pulumi:"action"`
	// Specifies the ID of the endpoint connection.
	EndpointId *string `pulumi:"endpointId"`
	// Specifies the ID of the dedicated instance to which the endpoint connection belongs.
	InstanceId *string `pulumi:"instanceId"`
	Region     *string `pulumi:"region"`
	// The current ststus of the endpoint connection.
	Status *string `pulumi:"status"`
}

type ApigEndpointConnectionManagementState struct {
	// Specifies the operation type endpoint connection.
	Action pulumi.StringPtrInput
	// Specifies the ID of the endpoint connection.
	EndpointId pulumi.StringPtrInput
	// Specifies the ID of the dedicated instance to which the endpoint connection belongs.
	InstanceId pulumi.StringPtrInput
	Region     pulumi.StringPtrInput
	// The current ststus of the endpoint connection.
	Status pulumi.StringPtrInput
}

func (ApigEndpointConnectionManagementState) ElementType() reflect.Type {
	return reflect.TypeOf((*apigEndpointConnectionManagementState)(nil)).Elem()
}

type apigEndpointConnectionManagementArgs struct {
	// Specifies the operation type endpoint connection.
	Action string `pulumi:"action"`
	// Specifies the ID of the endpoint connection.
	EndpointId string `pulumi:"endpointId"`
	// Specifies the ID of the dedicated instance to which the endpoint connection belongs.
	InstanceId string  `pulumi:"instanceId"`
	Region     *string `pulumi:"region"`
}

// The set of arguments for constructing a ApigEndpointConnectionManagement resource.
type ApigEndpointConnectionManagementArgs struct {
	// Specifies the operation type endpoint connection.
	Action pulumi.StringInput
	// Specifies the ID of the endpoint connection.
	EndpointId pulumi.StringInput
	// Specifies the ID of the dedicated instance to which the endpoint connection belongs.
	InstanceId pulumi.StringInput
	Region     pulumi.StringPtrInput
}

func (ApigEndpointConnectionManagementArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*apigEndpointConnectionManagementArgs)(nil)).Elem()
}

type ApigEndpointConnectionManagementInput interface {
	pulumi.Input

	ToApigEndpointConnectionManagementOutput() ApigEndpointConnectionManagementOutput
	ToApigEndpointConnectionManagementOutputWithContext(ctx context.Context) ApigEndpointConnectionManagementOutput
}

func (*ApigEndpointConnectionManagement) ElementType() reflect.Type {
	return reflect.TypeOf((**ApigEndpointConnectionManagement)(nil)).Elem()
}

func (i *ApigEndpointConnectionManagement) ToApigEndpointConnectionManagementOutput() ApigEndpointConnectionManagementOutput {
	return i.ToApigEndpointConnectionManagementOutputWithContext(context.Background())
}

func (i *ApigEndpointConnectionManagement) ToApigEndpointConnectionManagementOutputWithContext(ctx context.Context) ApigEndpointConnectionManagementOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApigEndpointConnectionManagementOutput)
}

// ApigEndpointConnectionManagementArrayInput is an input type that accepts ApigEndpointConnectionManagementArray and ApigEndpointConnectionManagementArrayOutput values.
// You can construct a concrete instance of `ApigEndpointConnectionManagementArrayInput` via:
//
//	ApigEndpointConnectionManagementArray{ ApigEndpointConnectionManagementArgs{...} }
type ApigEndpointConnectionManagementArrayInput interface {
	pulumi.Input

	ToApigEndpointConnectionManagementArrayOutput() ApigEndpointConnectionManagementArrayOutput
	ToApigEndpointConnectionManagementArrayOutputWithContext(context.Context) ApigEndpointConnectionManagementArrayOutput
}

type ApigEndpointConnectionManagementArray []ApigEndpointConnectionManagementInput

func (ApigEndpointConnectionManagementArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ApigEndpointConnectionManagement)(nil)).Elem()
}

func (i ApigEndpointConnectionManagementArray) ToApigEndpointConnectionManagementArrayOutput() ApigEndpointConnectionManagementArrayOutput {
	return i.ToApigEndpointConnectionManagementArrayOutputWithContext(context.Background())
}

func (i ApigEndpointConnectionManagementArray) ToApigEndpointConnectionManagementArrayOutputWithContext(ctx context.Context) ApigEndpointConnectionManagementArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApigEndpointConnectionManagementArrayOutput)
}

// ApigEndpointConnectionManagementMapInput is an input type that accepts ApigEndpointConnectionManagementMap and ApigEndpointConnectionManagementMapOutput values.
// You can construct a concrete instance of `ApigEndpointConnectionManagementMapInput` via:
//
//	ApigEndpointConnectionManagementMap{ "key": ApigEndpointConnectionManagementArgs{...} }
type ApigEndpointConnectionManagementMapInput interface {
	pulumi.Input

	ToApigEndpointConnectionManagementMapOutput() ApigEndpointConnectionManagementMapOutput
	ToApigEndpointConnectionManagementMapOutputWithContext(context.Context) ApigEndpointConnectionManagementMapOutput
}

type ApigEndpointConnectionManagementMap map[string]ApigEndpointConnectionManagementInput

func (ApigEndpointConnectionManagementMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ApigEndpointConnectionManagement)(nil)).Elem()
}

func (i ApigEndpointConnectionManagementMap) ToApigEndpointConnectionManagementMapOutput() ApigEndpointConnectionManagementMapOutput {
	return i.ToApigEndpointConnectionManagementMapOutputWithContext(context.Background())
}

func (i ApigEndpointConnectionManagementMap) ToApigEndpointConnectionManagementMapOutputWithContext(ctx context.Context) ApigEndpointConnectionManagementMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApigEndpointConnectionManagementMapOutput)
}

type ApigEndpointConnectionManagementOutput struct{ *pulumi.OutputState }

func (ApigEndpointConnectionManagementOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ApigEndpointConnectionManagement)(nil)).Elem()
}

func (o ApigEndpointConnectionManagementOutput) ToApigEndpointConnectionManagementOutput() ApigEndpointConnectionManagementOutput {
	return o
}

func (o ApigEndpointConnectionManagementOutput) ToApigEndpointConnectionManagementOutputWithContext(ctx context.Context) ApigEndpointConnectionManagementOutput {
	return o
}

// Specifies the operation type endpoint connection.
func (o ApigEndpointConnectionManagementOutput) Action() pulumi.StringOutput {
	return o.ApplyT(func(v *ApigEndpointConnectionManagement) pulumi.StringOutput { return v.Action }).(pulumi.StringOutput)
}

// Specifies the ID of the endpoint connection.
func (o ApigEndpointConnectionManagementOutput) EndpointId() pulumi.StringOutput {
	return o.ApplyT(func(v *ApigEndpointConnectionManagement) pulumi.StringOutput { return v.EndpointId }).(pulumi.StringOutput)
}

// Specifies the ID of the dedicated instance to which the endpoint connection belongs.
func (o ApigEndpointConnectionManagementOutput) InstanceId() pulumi.StringOutput {
	return o.ApplyT(func(v *ApigEndpointConnectionManagement) pulumi.StringOutput { return v.InstanceId }).(pulumi.StringOutput)
}

func (o ApigEndpointConnectionManagementOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v *ApigEndpointConnectionManagement) pulumi.StringOutput { return v.Region }).(pulumi.StringOutput)
}

// The current ststus of the endpoint connection.
func (o ApigEndpointConnectionManagementOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *ApigEndpointConnectionManagement) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

type ApigEndpointConnectionManagementArrayOutput struct{ *pulumi.OutputState }

func (ApigEndpointConnectionManagementArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ApigEndpointConnectionManagement)(nil)).Elem()
}

func (o ApigEndpointConnectionManagementArrayOutput) ToApigEndpointConnectionManagementArrayOutput() ApigEndpointConnectionManagementArrayOutput {
	return o
}

func (o ApigEndpointConnectionManagementArrayOutput) ToApigEndpointConnectionManagementArrayOutputWithContext(ctx context.Context) ApigEndpointConnectionManagementArrayOutput {
	return o
}

func (o ApigEndpointConnectionManagementArrayOutput) Index(i pulumi.IntInput) ApigEndpointConnectionManagementOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ApigEndpointConnectionManagement {
		return vs[0].([]*ApigEndpointConnectionManagement)[vs[1].(int)]
	}).(ApigEndpointConnectionManagementOutput)
}

type ApigEndpointConnectionManagementMapOutput struct{ *pulumi.OutputState }

func (ApigEndpointConnectionManagementMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ApigEndpointConnectionManagement)(nil)).Elem()
}

func (o ApigEndpointConnectionManagementMapOutput) ToApigEndpointConnectionManagementMapOutput() ApigEndpointConnectionManagementMapOutput {
	return o
}

func (o ApigEndpointConnectionManagementMapOutput) ToApigEndpointConnectionManagementMapOutputWithContext(ctx context.Context) ApigEndpointConnectionManagementMapOutput {
	return o
}

func (o ApigEndpointConnectionManagementMapOutput) MapIndex(k pulumi.StringInput) ApigEndpointConnectionManagementOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ApigEndpointConnectionManagement {
		return vs[0].(map[string]*ApigEndpointConnectionManagement)[vs[1].(string)]
	}).(ApigEndpointConnectionManagementOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ApigEndpointConnectionManagementInput)(nil)).Elem(), &ApigEndpointConnectionManagement{})
	pulumi.RegisterInputType(reflect.TypeOf((*ApigEndpointConnectionManagementArrayInput)(nil)).Elem(), ApigEndpointConnectionManagementArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ApigEndpointConnectionManagementMapInput)(nil)).Elem(), ApigEndpointConnectionManagementMap{})
	pulumi.RegisterOutputType(ApigEndpointConnectionManagementOutput{})
	pulumi.RegisterOutputType(ApigEndpointConnectionManagementArrayOutput{})
	pulumi.RegisterOutputType(ApigEndpointConnectionManagementMapOutput{})
}
