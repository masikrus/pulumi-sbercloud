// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package sbercloud

import (
	"context"
	"reflect"

	"github.com/masikrus/pulumi-sbercloud/sdk/go/sbercloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetIdentityCustomRole(ctx *pulumi.Context, args *GetIdentityCustomRoleArgs, opts ...pulumi.InvokeOption) (*GetIdentityCustomRoleResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetIdentityCustomRoleResult
	err := ctx.Invoke("sbercloud:index/getIdentityCustomRole:getIdentityCustomRole", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getIdentityCustomRole.
type GetIdentityCustomRoleArgs struct {
	Description *string `pulumi:"description"`
	DomainId    *string `pulumi:"domainId"`
	Id          *string `pulumi:"id"`
	Name        *string `pulumi:"name"`
	References  *int    `pulumi:"references"`
	Type        *string `pulumi:"type"`
}

// A collection of values returned by getIdentityCustomRole.
type GetIdentityCustomRoleResult struct {
	Catalog     string `pulumi:"catalog"`
	Description string `pulumi:"description"`
	DomainId    string `pulumi:"domainId"`
	Id          string `pulumi:"id"`
	Name        string `pulumi:"name"`
	Policy      string `pulumi:"policy"`
	References  int    `pulumi:"references"`
	Type        string `pulumi:"type"`
}

func GetIdentityCustomRoleOutput(ctx *pulumi.Context, args GetIdentityCustomRoleOutputArgs, opts ...pulumi.InvokeOption) GetIdentityCustomRoleResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (GetIdentityCustomRoleResultOutput, error) {
			args := v.(GetIdentityCustomRoleArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("sbercloud:index/getIdentityCustomRole:getIdentityCustomRole", args, GetIdentityCustomRoleResultOutput{}, options).(GetIdentityCustomRoleResultOutput), nil
		}).(GetIdentityCustomRoleResultOutput)
}

// A collection of arguments for invoking getIdentityCustomRole.
type GetIdentityCustomRoleOutputArgs struct {
	Description pulumi.StringPtrInput `pulumi:"description"`
	DomainId    pulumi.StringPtrInput `pulumi:"domainId"`
	Id          pulumi.StringPtrInput `pulumi:"id"`
	Name        pulumi.StringPtrInput `pulumi:"name"`
	References  pulumi.IntPtrInput    `pulumi:"references"`
	Type        pulumi.StringPtrInput `pulumi:"type"`
}

func (GetIdentityCustomRoleOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetIdentityCustomRoleArgs)(nil)).Elem()
}

// A collection of values returned by getIdentityCustomRole.
type GetIdentityCustomRoleResultOutput struct{ *pulumi.OutputState }

func (GetIdentityCustomRoleResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetIdentityCustomRoleResult)(nil)).Elem()
}

func (o GetIdentityCustomRoleResultOutput) ToGetIdentityCustomRoleResultOutput() GetIdentityCustomRoleResultOutput {
	return o
}

func (o GetIdentityCustomRoleResultOutput) ToGetIdentityCustomRoleResultOutputWithContext(ctx context.Context) GetIdentityCustomRoleResultOutput {
	return o
}

func (o GetIdentityCustomRoleResultOutput) Catalog() pulumi.StringOutput {
	return o.ApplyT(func(v GetIdentityCustomRoleResult) string { return v.Catalog }).(pulumi.StringOutput)
}

func (o GetIdentityCustomRoleResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v GetIdentityCustomRoleResult) string { return v.Description }).(pulumi.StringOutput)
}

func (o GetIdentityCustomRoleResultOutput) DomainId() pulumi.StringOutput {
	return o.ApplyT(func(v GetIdentityCustomRoleResult) string { return v.DomainId }).(pulumi.StringOutput)
}

func (o GetIdentityCustomRoleResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetIdentityCustomRoleResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetIdentityCustomRoleResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v GetIdentityCustomRoleResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o GetIdentityCustomRoleResultOutput) Policy() pulumi.StringOutput {
	return o.ApplyT(func(v GetIdentityCustomRoleResult) string { return v.Policy }).(pulumi.StringOutput)
}

func (o GetIdentityCustomRoleResultOutput) References() pulumi.IntOutput {
	return o.ApplyT(func(v GetIdentityCustomRoleResult) int { return v.References }).(pulumi.IntOutput)
}

func (o GetIdentityCustomRoleResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v GetIdentityCustomRoleResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetIdentityCustomRoleResultOutput{})
}
