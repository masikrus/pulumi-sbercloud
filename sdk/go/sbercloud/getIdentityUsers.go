// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package sbercloud

import (
	"context"
	"reflect"

	"github.com/masikrus/pulumi-sbercloud/sdk/go/sbercloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetIdentityUsers(ctx *pulumi.Context, args *GetIdentityUsersArgs, opts ...pulumi.InvokeOption) (*GetIdentityUsersResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetIdentityUsersResult
	err := ctx.Invoke("sbercloud:index/getIdentityUsers:getIdentityUsers", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getIdentityUsers.
type GetIdentityUsersArgs struct {
	Enabled *bool   `pulumi:"enabled"`
	Name    *string `pulumi:"name"`
}

// A collection of values returned by getIdentityUsers.
type GetIdentityUsersResult struct {
	Enabled *bool `pulumi:"enabled"`
	// The provider-assigned unique ID for this managed resource.
	Id    string                 `pulumi:"id"`
	Name  *string                `pulumi:"name"`
	Users []GetIdentityUsersUser `pulumi:"users"`
}

func GetIdentityUsersOutput(ctx *pulumi.Context, args GetIdentityUsersOutputArgs, opts ...pulumi.InvokeOption) GetIdentityUsersResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (GetIdentityUsersResultOutput, error) {
			args := v.(GetIdentityUsersArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("sbercloud:index/getIdentityUsers:getIdentityUsers", args, GetIdentityUsersResultOutput{}, options).(GetIdentityUsersResultOutput), nil
		}).(GetIdentityUsersResultOutput)
}

// A collection of arguments for invoking getIdentityUsers.
type GetIdentityUsersOutputArgs struct {
	Enabled pulumi.BoolPtrInput   `pulumi:"enabled"`
	Name    pulumi.StringPtrInput `pulumi:"name"`
}

func (GetIdentityUsersOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetIdentityUsersArgs)(nil)).Elem()
}

// A collection of values returned by getIdentityUsers.
type GetIdentityUsersResultOutput struct{ *pulumi.OutputState }

func (GetIdentityUsersResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetIdentityUsersResult)(nil)).Elem()
}

func (o GetIdentityUsersResultOutput) ToGetIdentityUsersResultOutput() GetIdentityUsersResultOutput {
	return o
}

func (o GetIdentityUsersResultOutput) ToGetIdentityUsersResultOutputWithContext(ctx context.Context) GetIdentityUsersResultOutput {
	return o
}

func (o GetIdentityUsersResultOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v GetIdentityUsersResult) *bool { return v.Enabled }).(pulumi.BoolPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetIdentityUsersResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetIdentityUsersResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetIdentityUsersResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetIdentityUsersResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o GetIdentityUsersResultOutput) Users() GetIdentityUsersUserArrayOutput {
	return o.ApplyT(func(v GetIdentityUsersResult) []GetIdentityUsersUser { return v.Users }).(GetIdentityUsersUserArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(GetIdentityUsersResultOutput{})
}
