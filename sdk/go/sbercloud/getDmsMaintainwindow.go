// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package sbercloud

import (
	"context"
	"reflect"

	"github.com/masikrus/pulumi-sbercloud/sdk/go/sbercloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetDmsMaintainwindow(ctx *pulumi.Context, args *GetDmsMaintainwindowArgs, opts ...pulumi.InvokeOption) (*GetDmsMaintainwindowResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetDmsMaintainwindowResult
	err := ctx.Invoke("sbercloud:index/getDmsMaintainwindow:getDmsMaintainwindow", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getDmsMaintainwindow.
type GetDmsMaintainwindowArgs struct {
	Begin   *string `pulumi:"begin"`
	Default *bool   `pulumi:"default"`
	End     *string `pulumi:"end"`
	Region  *string `pulumi:"region"`
	Seq     *int    `pulumi:"seq"`
}

// A collection of values returned by getDmsMaintainwindow.
type GetDmsMaintainwindowResult struct {
	Begin   string `pulumi:"begin"`
	Default bool   `pulumi:"default"`
	End     string `pulumi:"end"`
	// The provider-assigned unique ID for this managed resource.
	Id     string `pulumi:"id"`
	Region string `pulumi:"region"`
	Seq    int    `pulumi:"seq"`
}

func GetDmsMaintainwindowOutput(ctx *pulumi.Context, args GetDmsMaintainwindowOutputArgs, opts ...pulumi.InvokeOption) GetDmsMaintainwindowResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (GetDmsMaintainwindowResultOutput, error) {
			args := v.(GetDmsMaintainwindowArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("sbercloud:index/getDmsMaintainwindow:getDmsMaintainwindow", args, GetDmsMaintainwindowResultOutput{}, options).(GetDmsMaintainwindowResultOutput), nil
		}).(GetDmsMaintainwindowResultOutput)
}

// A collection of arguments for invoking getDmsMaintainwindow.
type GetDmsMaintainwindowOutputArgs struct {
	Begin   pulumi.StringPtrInput `pulumi:"begin"`
	Default pulumi.BoolPtrInput   `pulumi:"default"`
	End     pulumi.StringPtrInput `pulumi:"end"`
	Region  pulumi.StringPtrInput `pulumi:"region"`
	Seq     pulumi.IntPtrInput    `pulumi:"seq"`
}

func (GetDmsMaintainwindowOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetDmsMaintainwindowArgs)(nil)).Elem()
}

// A collection of values returned by getDmsMaintainwindow.
type GetDmsMaintainwindowResultOutput struct{ *pulumi.OutputState }

func (GetDmsMaintainwindowResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetDmsMaintainwindowResult)(nil)).Elem()
}

func (o GetDmsMaintainwindowResultOutput) ToGetDmsMaintainwindowResultOutput() GetDmsMaintainwindowResultOutput {
	return o
}

func (o GetDmsMaintainwindowResultOutput) ToGetDmsMaintainwindowResultOutputWithContext(ctx context.Context) GetDmsMaintainwindowResultOutput {
	return o
}

func (o GetDmsMaintainwindowResultOutput) Begin() pulumi.StringOutput {
	return o.ApplyT(func(v GetDmsMaintainwindowResult) string { return v.Begin }).(pulumi.StringOutput)
}

func (o GetDmsMaintainwindowResultOutput) Default() pulumi.BoolOutput {
	return o.ApplyT(func(v GetDmsMaintainwindowResult) bool { return v.Default }).(pulumi.BoolOutput)
}

func (o GetDmsMaintainwindowResultOutput) End() pulumi.StringOutput {
	return o.ApplyT(func(v GetDmsMaintainwindowResult) string { return v.End }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetDmsMaintainwindowResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetDmsMaintainwindowResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetDmsMaintainwindowResultOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v GetDmsMaintainwindowResult) string { return v.Region }).(pulumi.StringOutput)
}

func (o GetDmsMaintainwindowResultOutput) Seq() pulumi.IntOutput {
	return o.ApplyT(func(v GetDmsMaintainwindowResult) int { return v.Seq }).(pulumi.IntOutput)
}

func init() {
	pulumi.RegisterOutputType(GetDmsMaintainwindowResultOutput{})
}
