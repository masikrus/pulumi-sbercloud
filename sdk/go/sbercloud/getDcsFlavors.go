// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package sbercloud

import (
	"context"
	"reflect"

	"github.com/masikrus/pulumi-sbercloud/sdk/go/sbercloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetDcsFlavors(ctx *pulumi.Context, args *GetDcsFlavorsArgs, opts ...pulumi.InvokeOption) (*GetDcsFlavorsResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetDcsFlavorsResult
	err := ctx.Invoke("sbercloud:index/getDcsFlavors:getDcsFlavors", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getDcsFlavors.
type GetDcsFlavorsArgs struct {
	CacheMode       *string  `pulumi:"cacheMode"`
	Capacity        *float64 `pulumi:"capacity"`
	CpuArchitecture *string  `pulumi:"cpuArchitecture"`
	Engine          *string  `pulumi:"engine"`
	EngineVersion   *string  `pulumi:"engineVersion"`
	Name            *string  `pulumi:"name"`
	Region          *string  `pulumi:"region"`
}

// A collection of values returned by getDcsFlavors.
type GetDcsFlavorsResult struct {
	CacheMode       *string               `pulumi:"cacheMode"`
	Capacity        *float64              `pulumi:"capacity"`
	CpuArchitecture *string               `pulumi:"cpuArchitecture"`
	Engine          *string               `pulumi:"engine"`
	EngineVersion   *string               `pulumi:"engineVersion"`
	Flavors         []GetDcsFlavorsFlavor `pulumi:"flavors"`
	// The provider-assigned unique ID for this managed resource.
	Id     string  `pulumi:"id"`
	Name   *string `pulumi:"name"`
	Region string  `pulumi:"region"`
}

func GetDcsFlavorsOutput(ctx *pulumi.Context, args GetDcsFlavorsOutputArgs, opts ...pulumi.InvokeOption) GetDcsFlavorsResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (GetDcsFlavorsResultOutput, error) {
			args := v.(GetDcsFlavorsArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("sbercloud:index/getDcsFlavors:getDcsFlavors", args, GetDcsFlavorsResultOutput{}, options).(GetDcsFlavorsResultOutput), nil
		}).(GetDcsFlavorsResultOutput)
}

// A collection of arguments for invoking getDcsFlavors.
type GetDcsFlavorsOutputArgs struct {
	CacheMode       pulumi.StringPtrInput  `pulumi:"cacheMode"`
	Capacity        pulumi.Float64PtrInput `pulumi:"capacity"`
	CpuArchitecture pulumi.StringPtrInput  `pulumi:"cpuArchitecture"`
	Engine          pulumi.StringPtrInput  `pulumi:"engine"`
	EngineVersion   pulumi.StringPtrInput  `pulumi:"engineVersion"`
	Name            pulumi.StringPtrInput  `pulumi:"name"`
	Region          pulumi.StringPtrInput  `pulumi:"region"`
}

func (GetDcsFlavorsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetDcsFlavorsArgs)(nil)).Elem()
}

// A collection of values returned by getDcsFlavors.
type GetDcsFlavorsResultOutput struct{ *pulumi.OutputState }

func (GetDcsFlavorsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetDcsFlavorsResult)(nil)).Elem()
}

func (o GetDcsFlavorsResultOutput) ToGetDcsFlavorsResultOutput() GetDcsFlavorsResultOutput {
	return o
}

func (o GetDcsFlavorsResultOutput) ToGetDcsFlavorsResultOutputWithContext(ctx context.Context) GetDcsFlavorsResultOutput {
	return o
}

func (o GetDcsFlavorsResultOutput) CacheMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetDcsFlavorsResult) *string { return v.CacheMode }).(pulumi.StringPtrOutput)
}

func (o GetDcsFlavorsResultOutput) Capacity() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v GetDcsFlavorsResult) *float64 { return v.Capacity }).(pulumi.Float64PtrOutput)
}

func (o GetDcsFlavorsResultOutput) CpuArchitecture() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetDcsFlavorsResult) *string { return v.CpuArchitecture }).(pulumi.StringPtrOutput)
}

func (o GetDcsFlavorsResultOutput) Engine() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetDcsFlavorsResult) *string { return v.Engine }).(pulumi.StringPtrOutput)
}

func (o GetDcsFlavorsResultOutput) EngineVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetDcsFlavorsResult) *string { return v.EngineVersion }).(pulumi.StringPtrOutput)
}

func (o GetDcsFlavorsResultOutput) Flavors() GetDcsFlavorsFlavorArrayOutput {
	return o.ApplyT(func(v GetDcsFlavorsResult) []GetDcsFlavorsFlavor { return v.Flavors }).(GetDcsFlavorsFlavorArrayOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetDcsFlavorsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetDcsFlavorsResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetDcsFlavorsResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetDcsFlavorsResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o GetDcsFlavorsResultOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v GetDcsFlavorsResult) string { return v.Region }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetDcsFlavorsResultOutput{})
}
