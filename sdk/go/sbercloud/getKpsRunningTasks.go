// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package sbercloud

import (
	"context"
	"reflect"

	"github.com/masikrus/pulumi-sbercloud/sdk/go/sbercloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetKpsRunningTasks(ctx *pulumi.Context, args *GetKpsRunningTasksArgs, opts ...pulumi.InvokeOption) (*GetKpsRunningTasksResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetKpsRunningTasksResult
	err := ctx.Invoke("sbercloud:index/getKpsRunningTasks:getKpsRunningTasks", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getKpsRunningTasks.
type GetKpsRunningTasksArgs struct {
	Region *string `pulumi:"region"`
}

// A collection of values returned by getKpsRunningTasks.
type GetKpsRunningTasksResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id     string                   `pulumi:"id"`
	Region string                   `pulumi:"region"`
	Tasks  []GetKpsRunningTasksTask `pulumi:"tasks"`
}

func GetKpsRunningTasksOutput(ctx *pulumi.Context, args GetKpsRunningTasksOutputArgs, opts ...pulumi.InvokeOption) GetKpsRunningTasksResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (GetKpsRunningTasksResultOutput, error) {
			args := v.(GetKpsRunningTasksArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("sbercloud:index/getKpsRunningTasks:getKpsRunningTasks", args, GetKpsRunningTasksResultOutput{}, options).(GetKpsRunningTasksResultOutput), nil
		}).(GetKpsRunningTasksResultOutput)
}

// A collection of arguments for invoking getKpsRunningTasks.
type GetKpsRunningTasksOutputArgs struct {
	Region pulumi.StringPtrInput `pulumi:"region"`
}

func (GetKpsRunningTasksOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetKpsRunningTasksArgs)(nil)).Elem()
}

// A collection of values returned by getKpsRunningTasks.
type GetKpsRunningTasksResultOutput struct{ *pulumi.OutputState }

func (GetKpsRunningTasksResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetKpsRunningTasksResult)(nil)).Elem()
}

func (o GetKpsRunningTasksResultOutput) ToGetKpsRunningTasksResultOutput() GetKpsRunningTasksResultOutput {
	return o
}

func (o GetKpsRunningTasksResultOutput) ToGetKpsRunningTasksResultOutputWithContext(ctx context.Context) GetKpsRunningTasksResultOutput {
	return o
}

// The provider-assigned unique ID for this managed resource.
func (o GetKpsRunningTasksResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetKpsRunningTasksResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetKpsRunningTasksResultOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v GetKpsRunningTasksResult) string { return v.Region }).(pulumi.StringOutput)
}

func (o GetKpsRunningTasksResultOutput) Tasks() GetKpsRunningTasksTaskArrayOutput {
	return o.ApplyT(func(v GetKpsRunningTasksResult) []GetKpsRunningTasksTask { return v.Tasks }).(GetKpsRunningTasksTaskArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(GetKpsRunningTasksResultOutput{})
}
