// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package sbercloud

import (
	"context"
	"reflect"

	"github.com/masikrus/pulumi-sbercloud/sdk/go/sbercloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetApigInstanceFeatures(ctx *pulumi.Context, args *GetApigInstanceFeaturesArgs, opts ...pulumi.InvokeOption) (*GetApigInstanceFeaturesResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetApigInstanceFeaturesResult
	err := ctx.Invoke("sbercloud:index/getApigInstanceFeatures:getApigInstanceFeatures", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getApigInstanceFeatures.
type GetApigInstanceFeaturesArgs struct {
	InstanceId string  `pulumi:"instanceId"`
	Name       *string `pulumi:"name"`
	Region     *string `pulumi:"region"`
}

// A collection of values returned by getApigInstanceFeatures.
type GetApigInstanceFeaturesResult struct {
	Features []GetApigInstanceFeaturesFeature `pulumi:"features"`
	// The provider-assigned unique ID for this managed resource.
	Id         string  `pulumi:"id"`
	InstanceId string  `pulumi:"instanceId"`
	Name       *string `pulumi:"name"`
	Region     string  `pulumi:"region"`
}

func GetApigInstanceFeaturesOutput(ctx *pulumi.Context, args GetApigInstanceFeaturesOutputArgs, opts ...pulumi.InvokeOption) GetApigInstanceFeaturesResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (GetApigInstanceFeaturesResultOutput, error) {
			args := v.(GetApigInstanceFeaturesArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("sbercloud:index/getApigInstanceFeatures:getApigInstanceFeatures", args, GetApigInstanceFeaturesResultOutput{}, options).(GetApigInstanceFeaturesResultOutput), nil
		}).(GetApigInstanceFeaturesResultOutput)
}

// A collection of arguments for invoking getApigInstanceFeatures.
type GetApigInstanceFeaturesOutputArgs struct {
	InstanceId pulumi.StringInput    `pulumi:"instanceId"`
	Name       pulumi.StringPtrInput `pulumi:"name"`
	Region     pulumi.StringPtrInput `pulumi:"region"`
}

func (GetApigInstanceFeaturesOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetApigInstanceFeaturesArgs)(nil)).Elem()
}

// A collection of values returned by getApigInstanceFeatures.
type GetApigInstanceFeaturesResultOutput struct{ *pulumi.OutputState }

func (GetApigInstanceFeaturesResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetApigInstanceFeaturesResult)(nil)).Elem()
}

func (o GetApigInstanceFeaturesResultOutput) ToGetApigInstanceFeaturesResultOutput() GetApigInstanceFeaturesResultOutput {
	return o
}

func (o GetApigInstanceFeaturesResultOutput) ToGetApigInstanceFeaturesResultOutputWithContext(ctx context.Context) GetApigInstanceFeaturesResultOutput {
	return o
}

func (o GetApigInstanceFeaturesResultOutput) Features() GetApigInstanceFeaturesFeatureArrayOutput {
	return o.ApplyT(func(v GetApigInstanceFeaturesResult) []GetApigInstanceFeaturesFeature { return v.Features }).(GetApigInstanceFeaturesFeatureArrayOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetApigInstanceFeaturesResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetApigInstanceFeaturesResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetApigInstanceFeaturesResultOutput) InstanceId() pulumi.StringOutput {
	return o.ApplyT(func(v GetApigInstanceFeaturesResult) string { return v.InstanceId }).(pulumi.StringOutput)
}

func (o GetApigInstanceFeaturesResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetApigInstanceFeaturesResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o GetApigInstanceFeaturesResultOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v GetApigInstanceFeaturesResult) string { return v.Region }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetApigInstanceFeaturesResultOutput{})
}
