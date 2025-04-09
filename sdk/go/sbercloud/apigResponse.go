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

type ApigResponse struct {
	pulumi.CustomResourceState

	// The creation time of the API custom response.
	CreatedAt pulumi.StringOutput `pulumi:"createdAt"`
	// The ID of the API group to which the API custom response belongs.
	GroupId pulumi.StringOutput `pulumi:"groupId"`
	// The ID of the dedicated instance to which the API group and the API custom response belongs.
	InstanceId pulumi.StringOutput `pulumi:"instanceId"`
	// The name of the API custom response.
	Name pulumi.StringOutput `pulumi:"name"`
	// The region where the API custom response is located.
	Region pulumi.StringOutput `pulumi:"region"`
	// The API custom response rules definition.
	Rules ApigResponseRuleArrayOutput `pulumi:"rules"`
	// The latest update time of the API custom response.
	UpdatedAt pulumi.StringOutput `pulumi:"updatedAt"`
}

// NewApigResponse registers a new resource with the given unique name, arguments, and options.
func NewApigResponse(ctx *pulumi.Context,
	name string, args *ApigResponseArgs, opts ...pulumi.ResourceOption) (*ApigResponse, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.GroupId == nil {
		return nil, errors.New("invalid value for required argument 'GroupId'")
	}
	if args.InstanceId == nil {
		return nil, errors.New("invalid value for required argument 'InstanceId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource ApigResponse
	err := ctx.RegisterResource("sbercloud:index/apigResponse:ApigResponse", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetApigResponse gets an existing ApigResponse resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetApigResponse(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ApigResponseState, opts ...pulumi.ResourceOption) (*ApigResponse, error) {
	var resource ApigResponse
	err := ctx.ReadResource("sbercloud:index/apigResponse:ApigResponse", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ApigResponse resources.
type apigResponseState struct {
	// The creation time of the API custom response.
	CreatedAt *string `pulumi:"createdAt"`
	// The ID of the API group to which the API custom response belongs.
	GroupId *string `pulumi:"groupId"`
	// The ID of the dedicated instance to which the API group and the API custom response belongs.
	InstanceId *string `pulumi:"instanceId"`
	// The name of the API custom response.
	Name *string `pulumi:"name"`
	// The region where the API custom response is located.
	Region *string `pulumi:"region"`
	// The API custom response rules definition.
	Rules []ApigResponseRule `pulumi:"rules"`
	// The latest update time of the API custom response.
	UpdatedAt *string `pulumi:"updatedAt"`
}

type ApigResponseState struct {
	// The creation time of the API custom response.
	CreatedAt pulumi.StringPtrInput
	// The ID of the API group to which the API custom response belongs.
	GroupId pulumi.StringPtrInput
	// The ID of the dedicated instance to which the API group and the API custom response belongs.
	InstanceId pulumi.StringPtrInput
	// The name of the API custom response.
	Name pulumi.StringPtrInput
	// The region where the API custom response is located.
	Region pulumi.StringPtrInput
	// The API custom response rules definition.
	Rules ApigResponseRuleArrayInput
	// The latest update time of the API custom response.
	UpdatedAt pulumi.StringPtrInput
}

func (ApigResponseState) ElementType() reflect.Type {
	return reflect.TypeOf((*apigResponseState)(nil)).Elem()
}

type apigResponseArgs struct {
	// The ID of the API group to which the API custom response belongs.
	GroupId string `pulumi:"groupId"`
	// The ID of the dedicated instance to which the API group and the API custom response belongs.
	InstanceId string `pulumi:"instanceId"`
	// The name of the API custom response.
	Name *string `pulumi:"name"`
	// The region where the API custom response is located.
	Region *string `pulumi:"region"`
	// The API custom response rules definition.
	Rules []ApigResponseRule `pulumi:"rules"`
}

// The set of arguments for constructing a ApigResponse resource.
type ApigResponseArgs struct {
	// The ID of the API group to which the API custom response belongs.
	GroupId pulumi.StringInput
	// The ID of the dedicated instance to which the API group and the API custom response belongs.
	InstanceId pulumi.StringInput
	// The name of the API custom response.
	Name pulumi.StringPtrInput
	// The region where the API custom response is located.
	Region pulumi.StringPtrInput
	// The API custom response rules definition.
	Rules ApigResponseRuleArrayInput
}

func (ApigResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*apigResponseArgs)(nil)).Elem()
}

type ApigResponseInput interface {
	pulumi.Input

	ToApigResponseOutput() ApigResponseOutput
	ToApigResponseOutputWithContext(ctx context.Context) ApigResponseOutput
}

func (*ApigResponse) ElementType() reflect.Type {
	return reflect.TypeOf((**ApigResponse)(nil)).Elem()
}

func (i *ApigResponse) ToApigResponseOutput() ApigResponseOutput {
	return i.ToApigResponseOutputWithContext(context.Background())
}

func (i *ApigResponse) ToApigResponseOutputWithContext(ctx context.Context) ApigResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApigResponseOutput)
}

// ApigResponseArrayInput is an input type that accepts ApigResponseArray and ApigResponseArrayOutput values.
// You can construct a concrete instance of `ApigResponseArrayInput` via:
//
//	ApigResponseArray{ ApigResponseArgs{...} }
type ApigResponseArrayInput interface {
	pulumi.Input

	ToApigResponseArrayOutput() ApigResponseArrayOutput
	ToApigResponseArrayOutputWithContext(context.Context) ApigResponseArrayOutput
}

type ApigResponseArray []ApigResponseInput

func (ApigResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ApigResponse)(nil)).Elem()
}

func (i ApigResponseArray) ToApigResponseArrayOutput() ApigResponseArrayOutput {
	return i.ToApigResponseArrayOutputWithContext(context.Background())
}

func (i ApigResponseArray) ToApigResponseArrayOutputWithContext(ctx context.Context) ApigResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApigResponseArrayOutput)
}

// ApigResponseMapInput is an input type that accepts ApigResponseMap and ApigResponseMapOutput values.
// You can construct a concrete instance of `ApigResponseMapInput` via:
//
//	ApigResponseMap{ "key": ApigResponseArgs{...} }
type ApigResponseMapInput interface {
	pulumi.Input

	ToApigResponseMapOutput() ApigResponseMapOutput
	ToApigResponseMapOutputWithContext(context.Context) ApigResponseMapOutput
}

type ApigResponseMap map[string]ApigResponseInput

func (ApigResponseMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ApigResponse)(nil)).Elem()
}

func (i ApigResponseMap) ToApigResponseMapOutput() ApigResponseMapOutput {
	return i.ToApigResponseMapOutputWithContext(context.Background())
}

func (i ApigResponseMap) ToApigResponseMapOutputWithContext(ctx context.Context) ApigResponseMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApigResponseMapOutput)
}

type ApigResponseOutput struct{ *pulumi.OutputState }

func (ApigResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ApigResponse)(nil)).Elem()
}

func (o ApigResponseOutput) ToApigResponseOutput() ApigResponseOutput {
	return o
}

func (o ApigResponseOutput) ToApigResponseOutputWithContext(ctx context.Context) ApigResponseOutput {
	return o
}

// The creation time of the API custom response.
func (o ApigResponseOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *ApigResponse) pulumi.StringOutput { return v.CreatedAt }).(pulumi.StringOutput)
}

// The ID of the API group to which the API custom response belongs.
func (o ApigResponseOutput) GroupId() pulumi.StringOutput {
	return o.ApplyT(func(v *ApigResponse) pulumi.StringOutput { return v.GroupId }).(pulumi.StringOutput)
}

// The ID of the dedicated instance to which the API group and the API custom response belongs.
func (o ApigResponseOutput) InstanceId() pulumi.StringOutput {
	return o.ApplyT(func(v *ApigResponse) pulumi.StringOutput { return v.InstanceId }).(pulumi.StringOutput)
}

// The name of the API custom response.
func (o ApigResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ApigResponse) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The region where the API custom response is located.
func (o ApigResponseOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v *ApigResponse) pulumi.StringOutput { return v.Region }).(pulumi.StringOutput)
}

// The API custom response rules definition.
func (o ApigResponseOutput) Rules() ApigResponseRuleArrayOutput {
	return o.ApplyT(func(v *ApigResponse) ApigResponseRuleArrayOutput { return v.Rules }).(ApigResponseRuleArrayOutput)
}

// The latest update time of the API custom response.
func (o ApigResponseOutput) UpdatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *ApigResponse) pulumi.StringOutput { return v.UpdatedAt }).(pulumi.StringOutput)
}

type ApigResponseArrayOutput struct{ *pulumi.OutputState }

func (ApigResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ApigResponse)(nil)).Elem()
}

func (o ApigResponseArrayOutput) ToApigResponseArrayOutput() ApigResponseArrayOutput {
	return o
}

func (o ApigResponseArrayOutput) ToApigResponseArrayOutputWithContext(ctx context.Context) ApigResponseArrayOutput {
	return o
}

func (o ApigResponseArrayOutput) Index(i pulumi.IntInput) ApigResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ApigResponse {
		return vs[0].([]*ApigResponse)[vs[1].(int)]
	}).(ApigResponseOutput)
}

type ApigResponseMapOutput struct{ *pulumi.OutputState }

func (ApigResponseMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ApigResponse)(nil)).Elem()
}

func (o ApigResponseMapOutput) ToApigResponseMapOutput() ApigResponseMapOutput {
	return o
}

func (o ApigResponseMapOutput) ToApigResponseMapOutputWithContext(ctx context.Context) ApigResponseMapOutput {
	return o
}

func (o ApigResponseMapOutput) MapIndex(k pulumi.StringInput) ApigResponseOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ApigResponse {
		return vs[0].(map[string]*ApigResponse)[vs[1].(string)]
	}).(ApigResponseOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ApigResponseInput)(nil)).Elem(), &ApigResponse{})
	pulumi.RegisterInputType(reflect.TypeOf((*ApigResponseArrayInput)(nil)).Elem(), ApigResponseArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ApigResponseMapInput)(nil)).Elem(), ApigResponseMap{})
	pulumi.RegisterOutputType(ApigResponseOutput{})
	pulumi.RegisterOutputType(ApigResponseArrayOutput{})
	pulumi.RegisterOutputType(ApigResponseMapOutput{})
}
