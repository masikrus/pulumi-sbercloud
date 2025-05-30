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

type ApigApplicationQuotaAssociate struct {
	pulumi.CustomResourceState

	// The configuration of applications bound to the quota.
	Applications ApigApplicationQuotaAssociateApplicationArrayOutput `pulumi:"applications"`
	// The ID of the dedicated instance to which the application quota (policy) belongs.
	InstanceId pulumi.StringOutput `pulumi:"instanceId"`
	// The ID of the application quota (policy).
	QuotaId pulumi.StringOutput `pulumi:"quotaId"`
	// The region where the application quota (policy) is located.
	Region pulumi.StringOutput `pulumi:"region"`
}

// NewApigApplicationQuotaAssociate registers a new resource with the given unique name, arguments, and options.
func NewApigApplicationQuotaAssociate(ctx *pulumi.Context,
	name string, args *ApigApplicationQuotaAssociateArgs, opts ...pulumi.ResourceOption) (*ApigApplicationQuotaAssociate, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Applications == nil {
		return nil, errors.New("invalid value for required argument 'Applications'")
	}
	if args.InstanceId == nil {
		return nil, errors.New("invalid value for required argument 'InstanceId'")
	}
	if args.QuotaId == nil {
		return nil, errors.New("invalid value for required argument 'QuotaId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource ApigApplicationQuotaAssociate
	err := ctx.RegisterResource("sbercloud:index/apigApplicationQuotaAssociate:ApigApplicationQuotaAssociate", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetApigApplicationQuotaAssociate gets an existing ApigApplicationQuotaAssociate resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetApigApplicationQuotaAssociate(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ApigApplicationQuotaAssociateState, opts ...pulumi.ResourceOption) (*ApigApplicationQuotaAssociate, error) {
	var resource ApigApplicationQuotaAssociate
	err := ctx.ReadResource("sbercloud:index/apigApplicationQuotaAssociate:ApigApplicationQuotaAssociate", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ApigApplicationQuotaAssociate resources.
type apigApplicationQuotaAssociateState struct {
	// The configuration of applications bound to the quota.
	Applications []ApigApplicationQuotaAssociateApplication `pulumi:"applications"`
	// The ID of the dedicated instance to which the application quota (policy) belongs.
	InstanceId *string `pulumi:"instanceId"`
	// The ID of the application quota (policy).
	QuotaId *string `pulumi:"quotaId"`
	// The region where the application quota (policy) is located.
	Region *string `pulumi:"region"`
}

type ApigApplicationQuotaAssociateState struct {
	// The configuration of applications bound to the quota.
	Applications ApigApplicationQuotaAssociateApplicationArrayInput
	// The ID of the dedicated instance to which the application quota (policy) belongs.
	InstanceId pulumi.StringPtrInput
	// The ID of the application quota (policy).
	QuotaId pulumi.StringPtrInput
	// The region where the application quota (policy) is located.
	Region pulumi.StringPtrInput
}

func (ApigApplicationQuotaAssociateState) ElementType() reflect.Type {
	return reflect.TypeOf((*apigApplicationQuotaAssociateState)(nil)).Elem()
}

type apigApplicationQuotaAssociateArgs struct {
	// The configuration of applications bound to the quota.
	Applications []ApigApplicationQuotaAssociateApplication `pulumi:"applications"`
	// The ID of the dedicated instance to which the application quota (policy) belongs.
	InstanceId string `pulumi:"instanceId"`
	// The ID of the application quota (policy).
	QuotaId string `pulumi:"quotaId"`
	// The region where the application quota (policy) is located.
	Region *string `pulumi:"region"`
}

// The set of arguments for constructing a ApigApplicationQuotaAssociate resource.
type ApigApplicationQuotaAssociateArgs struct {
	// The configuration of applications bound to the quota.
	Applications ApigApplicationQuotaAssociateApplicationArrayInput
	// The ID of the dedicated instance to which the application quota (policy) belongs.
	InstanceId pulumi.StringInput
	// The ID of the application quota (policy).
	QuotaId pulumi.StringInput
	// The region where the application quota (policy) is located.
	Region pulumi.StringPtrInput
}

func (ApigApplicationQuotaAssociateArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*apigApplicationQuotaAssociateArgs)(nil)).Elem()
}

type ApigApplicationQuotaAssociateInput interface {
	pulumi.Input

	ToApigApplicationQuotaAssociateOutput() ApigApplicationQuotaAssociateOutput
	ToApigApplicationQuotaAssociateOutputWithContext(ctx context.Context) ApigApplicationQuotaAssociateOutput
}

func (*ApigApplicationQuotaAssociate) ElementType() reflect.Type {
	return reflect.TypeOf((**ApigApplicationQuotaAssociate)(nil)).Elem()
}

func (i *ApigApplicationQuotaAssociate) ToApigApplicationQuotaAssociateOutput() ApigApplicationQuotaAssociateOutput {
	return i.ToApigApplicationQuotaAssociateOutputWithContext(context.Background())
}

func (i *ApigApplicationQuotaAssociate) ToApigApplicationQuotaAssociateOutputWithContext(ctx context.Context) ApigApplicationQuotaAssociateOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApigApplicationQuotaAssociateOutput)
}

// ApigApplicationQuotaAssociateArrayInput is an input type that accepts ApigApplicationQuotaAssociateArray and ApigApplicationQuotaAssociateArrayOutput values.
// You can construct a concrete instance of `ApigApplicationQuotaAssociateArrayInput` via:
//
//	ApigApplicationQuotaAssociateArray{ ApigApplicationQuotaAssociateArgs{...} }
type ApigApplicationQuotaAssociateArrayInput interface {
	pulumi.Input

	ToApigApplicationQuotaAssociateArrayOutput() ApigApplicationQuotaAssociateArrayOutput
	ToApigApplicationQuotaAssociateArrayOutputWithContext(context.Context) ApigApplicationQuotaAssociateArrayOutput
}

type ApigApplicationQuotaAssociateArray []ApigApplicationQuotaAssociateInput

func (ApigApplicationQuotaAssociateArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ApigApplicationQuotaAssociate)(nil)).Elem()
}

func (i ApigApplicationQuotaAssociateArray) ToApigApplicationQuotaAssociateArrayOutput() ApigApplicationQuotaAssociateArrayOutput {
	return i.ToApigApplicationQuotaAssociateArrayOutputWithContext(context.Background())
}

func (i ApigApplicationQuotaAssociateArray) ToApigApplicationQuotaAssociateArrayOutputWithContext(ctx context.Context) ApigApplicationQuotaAssociateArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApigApplicationQuotaAssociateArrayOutput)
}

// ApigApplicationQuotaAssociateMapInput is an input type that accepts ApigApplicationQuotaAssociateMap and ApigApplicationQuotaAssociateMapOutput values.
// You can construct a concrete instance of `ApigApplicationQuotaAssociateMapInput` via:
//
//	ApigApplicationQuotaAssociateMap{ "key": ApigApplicationQuotaAssociateArgs{...} }
type ApigApplicationQuotaAssociateMapInput interface {
	pulumi.Input

	ToApigApplicationQuotaAssociateMapOutput() ApigApplicationQuotaAssociateMapOutput
	ToApigApplicationQuotaAssociateMapOutputWithContext(context.Context) ApigApplicationQuotaAssociateMapOutput
}

type ApigApplicationQuotaAssociateMap map[string]ApigApplicationQuotaAssociateInput

func (ApigApplicationQuotaAssociateMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ApigApplicationQuotaAssociate)(nil)).Elem()
}

func (i ApigApplicationQuotaAssociateMap) ToApigApplicationQuotaAssociateMapOutput() ApigApplicationQuotaAssociateMapOutput {
	return i.ToApigApplicationQuotaAssociateMapOutputWithContext(context.Background())
}

func (i ApigApplicationQuotaAssociateMap) ToApigApplicationQuotaAssociateMapOutputWithContext(ctx context.Context) ApigApplicationQuotaAssociateMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApigApplicationQuotaAssociateMapOutput)
}

type ApigApplicationQuotaAssociateOutput struct{ *pulumi.OutputState }

func (ApigApplicationQuotaAssociateOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ApigApplicationQuotaAssociate)(nil)).Elem()
}

func (o ApigApplicationQuotaAssociateOutput) ToApigApplicationQuotaAssociateOutput() ApigApplicationQuotaAssociateOutput {
	return o
}

func (o ApigApplicationQuotaAssociateOutput) ToApigApplicationQuotaAssociateOutputWithContext(ctx context.Context) ApigApplicationQuotaAssociateOutput {
	return o
}

// The configuration of applications bound to the quota.
func (o ApigApplicationQuotaAssociateOutput) Applications() ApigApplicationQuotaAssociateApplicationArrayOutput {
	return o.ApplyT(func(v *ApigApplicationQuotaAssociate) ApigApplicationQuotaAssociateApplicationArrayOutput {
		return v.Applications
	}).(ApigApplicationQuotaAssociateApplicationArrayOutput)
}

// The ID of the dedicated instance to which the application quota (policy) belongs.
func (o ApigApplicationQuotaAssociateOutput) InstanceId() pulumi.StringOutput {
	return o.ApplyT(func(v *ApigApplicationQuotaAssociate) pulumi.StringOutput { return v.InstanceId }).(pulumi.StringOutput)
}

// The ID of the application quota (policy).
func (o ApigApplicationQuotaAssociateOutput) QuotaId() pulumi.StringOutput {
	return o.ApplyT(func(v *ApigApplicationQuotaAssociate) pulumi.StringOutput { return v.QuotaId }).(pulumi.StringOutput)
}

// The region where the application quota (policy) is located.
func (o ApigApplicationQuotaAssociateOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v *ApigApplicationQuotaAssociate) pulumi.StringOutput { return v.Region }).(pulumi.StringOutput)
}

type ApigApplicationQuotaAssociateArrayOutput struct{ *pulumi.OutputState }

func (ApigApplicationQuotaAssociateArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ApigApplicationQuotaAssociate)(nil)).Elem()
}

func (o ApigApplicationQuotaAssociateArrayOutput) ToApigApplicationQuotaAssociateArrayOutput() ApigApplicationQuotaAssociateArrayOutput {
	return o
}

func (o ApigApplicationQuotaAssociateArrayOutput) ToApigApplicationQuotaAssociateArrayOutputWithContext(ctx context.Context) ApigApplicationQuotaAssociateArrayOutput {
	return o
}

func (o ApigApplicationQuotaAssociateArrayOutput) Index(i pulumi.IntInput) ApigApplicationQuotaAssociateOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ApigApplicationQuotaAssociate {
		return vs[0].([]*ApigApplicationQuotaAssociate)[vs[1].(int)]
	}).(ApigApplicationQuotaAssociateOutput)
}

type ApigApplicationQuotaAssociateMapOutput struct{ *pulumi.OutputState }

func (ApigApplicationQuotaAssociateMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ApigApplicationQuotaAssociate)(nil)).Elem()
}

func (o ApigApplicationQuotaAssociateMapOutput) ToApigApplicationQuotaAssociateMapOutput() ApigApplicationQuotaAssociateMapOutput {
	return o
}

func (o ApigApplicationQuotaAssociateMapOutput) ToApigApplicationQuotaAssociateMapOutputWithContext(ctx context.Context) ApigApplicationQuotaAssociateMapOutput {
	return o
}

func (o ApigApplicationQuotaAssociateMapOutput) MapIndex(k pulumi.StringInput) ApigApplicationQuotaAssociateOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ApigApplicationQuotaAssociate {
		return vs[0].(map[string]*ApigApplicationQuotaAssociate)[vs[1].(string)]
	}).(ApigApplicationQuotaAssociateOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ApigApplicationQuotaAssociateInput)(nil)).Elem(), &ApigApplicationQuotaAssociate{})
	pulumi.RegisterInputType(reflect.TypeOf((*ApigApplicationQuotaAssociateArrayInput)(nil)).Elem(), ApigApplicationQuotaAssociateArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ApigApplicationQuotaAssociateMapInput)(nil)).Elem(), ApigApplicationQuotaAssociateMap{})
	pulumi.RegisterOutputType(ApigApplicationQuotaAssociateOutput{})
	pulumi.RegisterOutputType(ApigApplicationQuotaAssociateArrayOutput{})
	pulumi.RegisterOutputType(ApigApplicationQuotaAssociateMapOutput{})
}
