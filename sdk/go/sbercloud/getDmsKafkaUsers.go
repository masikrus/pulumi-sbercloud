// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package sbercloud

import (
	"context"
	"reflect"

	"github.com/masikrus/pulumi-sbercloud/sdk/go/sbercloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetDmsKafkaUsers(ctx *pulumi.Context, args *GetDmsKafkaUsersArgs, opts ...pulumi.InvokeOption) (*GetDmsKafkaUsersResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetDmsKafkaUsersResult
	err := ctx.Invoke("sbercloud:index/getDmsKafkaUsers:getDmsKafkaUsers", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getDmsKafkaUsers.
type GetDmsKafkaUsersArgs struct {
	Description *string `pulumi:"description"`
	InstanceId  string  `pulumi:"instanceId"`
	Name        *string `pulumi:"name"`
	Region      *string `pulumi:"region"`
}

// A collection of values returned by getDmsKafkaUsers.
type GetDmsKafkaUsersResult struct {
	Description *string `pulumi:"description"`
	// The provider-assigned unique ID for this managed resource.
	Id         string                 `pulumi:"id"`
	InstanceId string                 `pulumi:"instanceId"`
	Name       *string                `pulumi:"name"`
	Region     string                 `pulumi:"region"`
	Users      []GetDmsKafkaUsersUser `pulumi:"users"`
}

func GetDmsKafkaUsersOutput(ctx *pulumi.Context, args GetDmsKafkaUsersOutputArgs, opts ...pulumi.InvokeOption) GetDmsKafkaUsersResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (GetDmsKafkaUsersResultOutput, error) {
			args := v.(GetDmsKafkaUsersArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("sbercloud:index/getDmsKafkaUsers:getDmsKafkaUsers", args, GetDmsKafkaUsersResultOutput{}, options).(GetDmsKafkaUsersResultOutput), nil
		}).(GetDmsKafkaUsersResultOutput)
}

// A collection of arguments for invoking getDmsKafkaUsers.
type GetDmsKafkaUsersOutputArgs struct {
	Description pulumi.StringPtrInput `pulumi:"description"`
	InstanceId  pulumi.StringInput    `pulumi:"instanceId"`
	Name        pulumi.StringPtrInput `pulumi:"name"`
	Region      pulumi.StringPtrInput `pulumi:"region"`
}

func (GetDmsKafkaUsersOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetDmsKafkaUsersArgs)(nil)).Elem()
}

// A collection of values returned by getDmsKafkaUsers.
type GetDmsKafkaUsersResultOutput struct{ *pulumi.OutputState }

func (GetDmsKafkaUsersResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetDmsKafkaUsersResult)(nil)).Elem()
}

func (o GetDmsKafkaUsersResultOutput) ToGetDmsKafkaUsersResultOutput() GetDmsKafkaUsersResultOutput {
	return o
}

func (o GetDmsKafkaUsersResultOutput) ToGetDmsKafkaUsersResultOutputWithContext(ctx context.Context) GetDmsKafkaUsersResultOutput {
	return o
}

func (o GetDmsKafkaUsersResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetDmsKafkaUsersResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetDmsKafkaUsersResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetDmsKafkaUsersResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetDmsKafkaUsersResultOutput) InstanceId() pulumi.StringOutput {
	return o.ApplyT(func(v GetDmsKafkaUsersResult) string { return v.InstanceId }).(pulumi.StringOutput)
}

func (o GetDmsKafkaUsersResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetDmsKafkaUsersResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o GetDmsKafkaUsersResultOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v GetDmsKafkaUsersResult) string { return v.Region }).(pulumi.StringOutput)
}

func (o GetDmsKafkaUsersResultOutput) Users() GetDmsKafkaUsersUserArrayOutput {
	return o.ApplyT(func(v GetDmsKafkaUsersResult) []GetDmsKafkaUsersUser { return v.Users }).(GetDmsKafkaUsersUserArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(GetDmsKafkaUsersResultOutput{})
}
