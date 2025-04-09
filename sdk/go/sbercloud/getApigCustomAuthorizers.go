// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package sbercloud

import (
	"context"
	"reflect"

	"github.com/masikrus/pulumi-sbercloud/sdk/go/sbercloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetApigCustomAuthorizers(ctx *pulumi.Context, args *GetApigCustomAuthorizersArgs, opts ...pulumi.InvokeOption) (*GetApigCustomAuthorizersResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetApigCustomAuthorizersResult
	err := ctx.Invoke("sbercloud:index/getApigCustomAuthorizers:getApigCustomAuthorizers", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getApigCustomAuthorizers.
type GetApigCustomAuthorizersArgs struct {
	AuthorizerId *string `pulumi:"authorizerId"`
	InstanceId   string  `pulumi:"instanceId"`
	Name         *string `pulumi:"name"`
	Region       *string `pulumi:"region"`
	Type         *string `pulumi:"type"`
}

// A collection of values returned by getApigCustomAuthorizers.
type GetApigCustomAuthorizersResult struct {
	AuthorizerId *string                              `pulumi:"authorizerId"`
	Authorizers  []GetApigCustomAuthorizersAuthorizer `pulumi:"authorizers"`
	// The provider-assigned unique ID for this managed resource.
	Id         string  `pulumi:"id"`
	InstanceId string  `pulumi:"instanceId"`
	Name       *string `pulumi:"name"`
	Region     string  `pulumi:"region"`
	Type       *string `pulumi:"type"`
}

func GetApigCustomAuthorizersOutput(ctx *pulumi.Context, args GetApigCustomAuthorizersOutputArgs, opts ...pulumi.InvokeOption) GetApigCustomAuthorizersResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (GetApigCustomAuthorizersResultOutput, error) {
			args := v.(GetApigCustomAuthorizersArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("sbercloud:index/getApigCustomAuthorizers:getApigCustomAuthorizers", args, GetApigCustomAuthorizersResultOutput{}, options).(GetApigCustomAuthorizersResultOutput), nil
		}).(GetApigCustomAuthorizersResultOutput)
}

// A collection of arguments for invoking getApigCustomAuthorizers.
type GetApigCustomAuthorizersOutputArgs struct {
	AuthorizerId pulumi.StringPtrInput `pulumi:"authorizerId"`
	InstanceId   pulumi.StringInput    `pulumi:"instanceId"`
	Name         pulumi.StringPtrInput `pulumi:"name"`
	Region       pulumi.StringPtrInput `pulumi:"region"`
	Type         pulumi.StringPtrInput `pulumi:"type"`
}

func (GetApigCustomAuthorizersOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetApigCustomAuthorizersArgs)(nil)).Elem()
}

// A collection of values returned by getApigCustomAuthorizers.
type GetApigCustomAuthorizersResultOutput struct{ *pulumi.OutputState }

func (GetApigCustomAuthorizersResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetApigCustomAuthorizersResult)(nil)).Elem()
}

func (o GetApigCustomAuthorizersResultOutput) ToGetApigCustomAuthorizersResultOutput() GetApigCustomAuthorizersResultOutput {
	return o
}

func (o GetApigCustomAuthorizersResultOutput) ToGetApigCustomAuthorizersResultOutputWithContext(ctx context.Context) GetApigCustomAuthorizersResultOutput {
	return o
}

func (o GetApigCustomAuthorizersResultOutput) AuthorizerId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetApigCustomAuthorizersResult) *string { return v.AuthorizerId }).(pulumi.StringPtrOutput)
}

func (o GetApigCustomAuthorizersResultOutput) Authorizers() GetApigCustomAuthorizersAuthorizerArrayOutput {
	return o.ApplyT(func(v GetApigCustomAuthorizersResult) []GetApigCustomAuthorizersAuthorizer { return v.Authorizers }).(GetApigCustomAuthorizersAuthorizerArrayOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetApigCustomAuthorizersResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetApigCustomAuthorizersResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetApigCustomAuthorizersResultOutput) InstanceId() pulumi.StringOutput {
	return o.ApplyT(func(v GetApigCustomAuthorizersResult) string { return v.InstanceId }).(pulumi.StringOutput)
}

func (o GetApigCustomAuthorizersResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetApigCustomAuthorizersResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o GetApigCustomAuthorizersResultOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v GetApigCustomAuthorizersResult) string { return v.Region }).(pulumi.StringOutput)
}

func (o GetApigCustomAuthorizersResultOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetApigCustomAuthorizersResult) *string { return v.Type }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetApigCustomAuthorizersResultOutput{})
}
