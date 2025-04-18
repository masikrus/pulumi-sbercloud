// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package sbercloud

import (
	"context"
	"reflect"

	"github.com/masikrus/pulumi-sbercloud/sdk/go/sbercloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetCdmFlavors(ctx *pulumi.Context, args *GetCdmFlavorsArgs, opts ...pulumi.InvokeOption) (*GetCdmFlavorsResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetCdmFlavorsResult
	err := ctx.Invoke("sbercloud:index/getCdmFlavors:getCdmFlavors", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getCdmFlavors.
type GetCdmFlavorsArgs struct {
	Region *string `pulumi:"region"`
}

// A collection of values returned by getCdmFlavors.
type GetCdmFlavorsResult struct {
	Flavors []GetCdmFlavorsFlavor `pulumi:"flavors"`
	// The provider-assigned unique ID for this managed resource.
	Id      string `pulumi:"id"`
	Region  string `pulumi:"region"`
	Version string `pulumi:"version"`
}

func GetCdmFlavorsOutput(ctx *pulumi.Context, args GetCdmFlavorsOutputArgs, opts ...pulumi.InvokeOption) GetCdmFlavorsResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (GetCdmFlavorsResultOutput, error) {
			args := v.(GetCdmFlavorsArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("sbercloud:index/getCdmFlavors:getCdmFlavors", args, GetCdmFlavorsResultOutput{}, options).(GetCdmFlavorsResultOutput), nil
		}).(GetCdmFlavorsResultOutput)
}

// A collection of arguments for invoking getCdmFlavors.
type GetCdmFlavorsOutputArgs struct {
	Region pulumi.StringPtrInput `pulumi:"region"`
}

func (GetCdmFlavorsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetCdmFlavorsArgs)(nil)).Elem()
}

// A collection of values returned by getCdmFlavors.
type GetCdmFlavorsResultOutput struct{ *pulumi.OutputState }

func (GetCdmFlavorsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetCdmFlavorsResult)(nil)).Elem()
}

func (o GetCdmFlavorsResultOutput) ToGetCdmFlavorsResultOutput() GetCdmFlavorsResultOutput {
	return o
}

func (o GetCdmFlavorsResultOutput) ToGetCdmFlavorsResultOutputWithContext(ctx context.Context) GetCdmFlavorsResultOutput {
	return o
}

func (o GetCdmFlavorsResultOutput) Flavors() GetCdmFlavorsFlavorArrayOutput {
	return o.ApplyT(func(v GetCdmFlavorsResult) []GetCdmFlavorsFlavor { return v.Flavors }).(GetCdmFlavorsFlavorArrayOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetCdmFlavorsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetCdmFlavorsResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetCdmFlavorsResultOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v GetCdmFlavorsResult) string { return v.Region }).(pulumi.StringOutput)
}

func (o GetCdmFlavorsResultOutput) Version() pulumi.StringOutput {
	return o.ApplyT(func(v GetCdmFlavorsResult) string { return v.Version }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetCdmFlavorsResultOutput{})
}
