// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package sbercloud

import (
	"context"
	"reflect"

	"github.com/masikrus/pulumi-sbercloud/sdk/go/sbercloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetDcsMaintainwindow(ctx *pulumi.Context, args *GetDcsMaintainwindowArgs, opts ...pulumi.InvokeOption) (*GetDcsMaintainwindowResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetDcsMaintainwindowResult
	err := ctx.Invoke("sbercloud:index/getDcsMaintainwindow:getDcsMaintainwindow", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getDcsMaintainwindow.
type GetDcsMaintainwindowArgs struct {
	Begin   *string `pulumi:"begin"`
	Default *bool   `pulumi:"default"`
	End     *string `pulumi:"end"`
	Region  *string `pulumi:"region"`
	Seq     *int    `pulumi:"seq"`
}

// A collection of values returned by getDcsMaintainwindow.
type GetDcsMaintainwindowResult struct {
	Begin   string `pulumi:"begin"`
	Default bool   `pulumi:"default"`
	End     string `pulumi:"end"`
	// The provider-assigned unique ID for this managed resource.
	Id     string `pulumi:"id"`
	Region string `pulumi:"region"`
	Seq    int    `pulumi:"seq"`
}

func GetDcsMaintainwindowOutput(ctx *pulumi.Context, args GetDcsMaintainwindowOutputArgs, opts ...pulumi.InvokeOption) GetDcsMaintainwindowResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (GetDcsMaintainwindowResultOutput, error) {
			args := v.(GetDcsMaintainwindowArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("sbercloud:index/getDcsMaintainwindow:getDcsMaintainwindow", args, GetDcsMaintainwindowResultOutput{}, options).(GetDcsMaintainwindowResultOutput), nil
		}).(GetDcsMaintainwindowResultOutput)
}

// A collection of arguments for invoking getDcsMaintainwindow.
type GetDcsMaintainwindowOutputArgs struct {
	Begin   pulumi.StringPtrInput `pulumi:"begin"`
	Default pulumi.BoolPtrInput   `pulumi:"default"`
	End     pulumi.StringPtrInput `pulumi:"end"`
	Region  pulumi.StringPtrInput `pulumi:"region"`
	Seq     pulumi.IntPtrInput    `pulumi:"seq"`
}

func (GetDcsMaintainwindowOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetDcsMaintainwindowArgs)(nil)).Elem()
}

// A collection of values returned by getDcsMaintainwindow.
type GetDcsMaintainwindowResultOutput struct{ *pulumi.OutputState }

func (GetDcsMaintainwindowResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetDcsMaintainwindowResult)(nil)).Elem()
}

func (o GetDcsMaintainwindowResultOutput) ToGetDcsMaintainwindowResultOutput() GetDcsMaintainwindowResultOutput {
	return o
}

func (o GetDcsMaintainwindowResultOutput) ToGetDcsMaintainwindowResultOutputWithContext(ctx context.Context) GetDcsMaintainwindowResultOutput {
	return o
}

func (o GetDcsMaintainwindowResultOutput) Begin() pulumi.StringOutput {
	return o.ApplyT(func(v GetDcsMaintainwindowResult) string { return v.Begin }).(pulumi.StringOutput)
}

func (o GetDcsMaintainwindowResultOutput) Default() pulumi.BoolOutput {
	return o.ApplyT(func(v GetDcsMaintainwindowResult) bool { return v.Default }).(pulumi.BoolOutput)
}

func (o GetDcsMaintainwindowResultOutput) End() pulumi.StringOutput {
	return o.ApplyT(func(v GetDcsMaintainwindowResult) string { return v.End }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetDcsMaintainwindowResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetDcsMaintainwindowResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetDcsMaintainwindowResultOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v GetDcsMaintainwindowResult) string { return v.Region }).(pulumi.StringOutput)
}

func (o GetDcsMaintainwindowResultOutput) Seq() pulumi.IntOutput {
	return o.ApplyT(func(v GetDcsMaintainwindowResult) int { return v.Seq }).(pulumi.IntOutput)
}

func init() {
	pulumi.RegisterOutputType(GetDcsMaintainwindowResultOutput{})
}
