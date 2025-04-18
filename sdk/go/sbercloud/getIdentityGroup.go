// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package sbercloud

import (
	"context"
	"reflect"

	"github.com/masikrus/pulumi-sbercloud/sdk/go/sbercloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupIdentityGroup(ctx *pulumi.Context, args *LookupIdentityGroupArgs, opts ...pulumi.InvokeOption) (*LookupIdentityGroupResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupIdentityGroupResult
	err := ctx.Invoke("sbercloud:index/getIdentityGroup:getIdentityGroup", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getIdentityGroup.
type LookupIdentityGroupArgs struct {
	Description *string `pulumi:"description"`
	Id          *string `pulumi:"id"`
	Name        *string `pulumi:"name"`
}

// A collection of values returned by getIdentityGroup.
type LookupIdentityGroupResult struct {
	Description string                 `pulumi:"description"`
	DomainId    string                 `pulumi:"domainId"`
	Id          string                 `pulumi:"id"`
	Name        string                 `pulumi:"name"`
	Users       []GetIdentityGroupUser `pulumi:"users"`
}

func LookupIdentityGroupOutput(ctx *pulumi.Context, args LookupIdentityGroupOutputArgs, opts ...pulumi.InvokeOption) LookupIdentityGroupResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupIdentityGroupResultOutput, error) {
			args := v.(LookupIdentityGroupArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("sbercloud:index/getIdentityGroup:getIdentityGroup", args, LookupIdentityGroupResultOutput{}, options).(LookupIdentityGroupResultOutput), nil
		}).(LookupIdentityGroupResultOutput)
}

// A collection of arguments for invoking getIdentityGroup.
type LookupIdentityGroupOutputArgs struct {
	Description pulumi.StringPtrInput `pulumi:"description"`
	Id          pulumi.StringPtrInput `pulumi:"id"`
	Name        pulumi.StringPtrInput `pulumi:"name"`
}

func (LookupIdentityGroupOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupIdentityGroupArgs)(nil)).Elem()
}

// A collection of values returned by getIdentityGroup.
type LookupIdentityGroupResultOutput struct{ *pulumi.OutputState }

func (LookupIdentityGroupResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupIdentityGroupResult)(nil)).Elem()
}

func (o LookupIdentityGroupResultOutput) ToLookupIdentityGroupResultOutput() LookupIdentityGroupResultOutput {
	return o
}

func (o LookupIdentityGroupResultOutput) ToLookupIdentityGroupResultOutputWithContext(ctx context.Context) LookupIdentityGroupResultOutput {
	return o
}

func (o LookupIdentityGroupResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIdentityGroupResult) string { return v.Description }).(pulumi.StringOutput)
}

func (o LookupIdentityGroupResultOutput) DomainId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIdentityGroupResult) string { return v.DomainId }).(pulumi.StringOutput)
}

func (o LookupIdentityGroupResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIdentityGroupResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupIdentityGroupResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIdentityGroupResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupIdentityGroupResultOutput) Users() GetIdentityGroupUserArrayOutput {
	return o.ApplyT(func(v LookupIdentityGroupResult) []GetIdentityGroupUser { return v.Users }).(GetIdentityGroupUserArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupIdentityGroupResultOutput{})
}
