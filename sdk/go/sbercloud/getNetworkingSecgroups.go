// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package sbercloud

import (
	"context"
	"reflect"

	"github.com/masikrus/pulumi-sbercloud/sdk/go/sbercloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetNetworkingSecgroups(ctx *pulumi.Context, args *GetNetworkingSecgroupsArgs, opts ...pulumi.InvokeOption) (*GetNetworkingSecgroupsResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetNetworkingSecgroupsResult
	err := ctx.Invoke("sbercloud:index/getNetworkingSecgroups:getNetworkingSecgroups", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getNetworkingSecgroups.
type GetNetworkingSecgroupsArgs struct {
	Description         *string `pulumi:"description"`
	EnterpriseProjectId *string `pulumi:"enterpriseProjectId"`
	Id                  *string `pulumi:"id"`
	Name                *string `pulumi:"name"`
	Region              *string `pulumi:"region"`
}

// A collection of values returned by getNetworkingSecgroups.
type GetNetworkingSecgroupsResult struct {
	Description         *string                               `pulumi:"description"`
	EnterpriseProjectId *string                               `pulumi:"enterpriseProjectId"`
	Id                  string                                `pulumi:"id"`
	Name                *string                               `pulumi:"name"`
	Region              string                                `pulumi:"region"`
	SecurityGroups      []GetNetworkingSecgroupsSecurityGroup `pulumi:"securityGroups"`
}

func GetNetworkingSecgroupsOutput(ctx *pulumi.Context, args GetNetworkingSecgroupsOutputArgs, opts ...pulumi.InvokeOption) GetNetworkingSecgroupsResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (GetNetworkingSecgroupsResultOutput, error) {
			args := v.(GetNetworkingSecgroupsArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("sbercloud:index/getNetworkingSecgroups:getNetworkingSecgroups", args, GetNetworkingSecgroupsResultOutput{}, options).(GetNetworkingSecgroupsResultOutput), nil
		}).(GetNetworkingSecgroupsResultOutput)
}

// A collection of arguments for invoking getNetworkingSecgroups.
type GetNetworkingSecgroupsOutputArgs struct {
	Description         pulumi.StringPtrInput `pulumi:"description"`
	EnterpriseProjectId pulumi.StringPtrInput `pulumi:"enterpriseProjectId"`
	Id                  pulumi.StringPtrInput `pulumi:"id"`
	Name                pulumi.StringPtrInput `pulumi:"name"`
	Region              pulumi.StringPtrInput `pulumi:"region"`
}

func (GetNetworkingSecgroupsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetNetworkingSecgroupsArgs)(nil)).Elem()
}

// A collection of values returned by getNetworkingSecgroups.
type GetNetworkingSecgroupsResultOutput struct{ *pulumi.OutputState }

func (GetNetworkingSecgroupsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetNetworkingSecgroupsResult)(nil)).Elem()
}

func (o GetNetworkingSecgroupsResultOutput) ToGetNetworkingSecgroupsResultOutput() GetNetworkingSecgroupsResultOutput {
	return o
}

func (o GetNetworkingSecgroupsResultOutput) ToGetNetworkingSecgroupsResultOutputWithContext(ctx context.Context) GetNetworkingSecgroupsResultOutput {
	return o
}

func (o GetNetworkingSecgroupsResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetNetworkingSecgroupsResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o GetNetworkingSecgroupsResultOutput) EnterpriseProjectId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetNetworkingSecgroupsResult) *string { return v.EnterpriseProjectId }).(pulumi.StringPtrOutput)
}

func (o GetNetworkingSecgroupsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetNetworkingSecgroupsResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetNetworkingSecgroupsResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetNetworkingSecgroupsResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o GetNetworkingSecgroupsResultOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v GetNetworkingSecgroupsResult) string { return v.Region }).(pulumi.StringOutput)
}

func (o GetNetworkingSecgroupsResultOutput) SecurityGroups() GetNetworkingSecgroupsSecurityGroupArrayOutput {
	return o.ApplyT(func(v GetNetworkingSecgroupsResult) []GetNetworkingSecgroupsSecurityGroup { return v.SecurityGroups }).(GetNetworkingSecgroupsSecurityGroupArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(GetNetworkingSecgroupsResultOutput{})
}
