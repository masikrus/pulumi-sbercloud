// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package sbercloud

import (
	"context"
	"reflect"

	"github.com/masikrus/pulumi-sbercloud/sdk/go/sbercloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetVpcIds(ctx *pulumi.Context, args *GetVpcIdsArgs, opts ...pulumi.InvokeOption) (*GetVpcIdsResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetVpcIdsResult
	err := ctx.Invoke("sbercloud:index/getVpcIds:getVpcIds", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getVpcIds.
type GetVpcIdsArgs struct {
	Region *string `pulumi:"region"`
}

// A collection of values returned by getVpcIds.
type GetVpcIdsResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id     string   `pulumi:"id"`
	Ids    []string `pulumi:"ids"`
	Region string   `pulumi:"region"`
}

func GetVpcIdsOutput(ctx *pulumi.Context, args GetVpcIdsOutputArgs, opts ...pulumi.InvokeOption) GetVpcIdsResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (GetVpcIdsResultOutput, error) {
			args := v.(GetVpcIdsArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("sbercloud:index/getVpcIds:getVpcIds", args, GetVpcIdsResultOutput{}, options).(GetVpcIdsResultOutput), nil
		}).(GetVpcIdsResultOutput)
}

// A collection of arguments for invoking getVpcIds.
type GetVpcIdsOutputArgs struct {
	Region pulumi.StringPtrInput `pulumi:"region"`
}

func (GetVpcIdsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetVpcIdsArgs)(nil)).Elem()
}

// A collection of values returned by getVpcIds.
type GetVpcIdsResultOutput struct{ *pulumi.OutputState }

func (GetVpcIdsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetVpcIdsResult)(nil)).Elem()
}

func (o GetVpcIdsResultOutput) ToGetVpcIdsResultOutput() GetVpcIdsResultOutput {
	return o
}

func (o GetVpcIdsResultOutput) ToGetVpcIdsResultOutputWithContext(ctx context.Context) GetVpcIdsResultOutput {
	return o
}

// The provider-assigned unique ID for this managed resource.
func (o GetVpcIdsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetVpcIdsResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetVpcIdsResultOutput) Ids() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetVpcIdsResult) []string { return v.Ids }).(pulumi.StringArrayOutput)
}

func (o GetVpcIdsResultOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v GetVpcIdsResult) string { return v.Region }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetVpcIdsResultOutput{})
}
