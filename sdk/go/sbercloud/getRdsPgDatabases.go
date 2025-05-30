// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package sbercloud

import (
	"context"
	"reflect"

	"github.com/masikrus/pulumi-sbercloud/sdk/go/sbercloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetRdsPgDatabases(ctx *pulumi.Context, args *GetRdsPgDatabasesArgs, opts ...pulumi.InvokeOption) (*GetRdsPgDatabasesResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetRdsPgDatabasesResult
	err := ctx.Invoke("sbercloud:index/getRdsPgDatabases:getRdsPgDatabases", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getRdsPgDatabases.
type GetRdsPgDatabasesArgs struct {
	CharacterSet *string `pulumi:"characterSet"`
	InstanceId   string  `pulumi:"instanceId"`
	LcCollate    *string `pulumi:"lcCollate"`
	Name         *string `pulumi:"name"`
	Owner        *string `pulumi:"owner"`
	Region       *string `pulumi:"region"`
	Size         *int    `pulumi:"size"`
}

// A collection of values returned by getRdsPgDatabases.
type GetRdsPgDatabasesResult struct {
	CharacterSet *string                     `pulumi:"characterSet"`
	Databases    []GetRdsPgDatabasesDatabase `pulumi:"databases"`
	// The provider-assigned unique ID for this managed resource.
	Id         string  `pulumi:"id"`
	InstanceId string  `pulumi:"instanceId"`
	LcCollate  *string `pulumi:"lcCollate"`
	Name       *string `pulumi:"name"`
	Owner      *string `pulumi:"owner"`
	Region     string  `pulumi:"region"`
	Size       *int    `pulumi:"size"`
}

func GetRdsPgDatabasesOutput(ctx *pulumi.Context, args GetRdsPgDatabasesOutputArgs, opts ...pulumi.InvokeOption) GetRdsPgDatabasesResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (GetRdsPgDatabasesResultOutput, error) {
			args := v.(GetRdsPgDatabasesArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("sbercloud:index/getRdsPgDatabases:getRdsPgDatabases", args, GetRdsPgDatabasesResultOutput{}, options).(GetRdsPgDatabasesResultOutput), nil
		}).(GetRdsPgDatabasesResultOutput)
}

// A collection of arguments for invoking getRdsPgDatabases.
type GetRdsPgDatabasesOutputArgs struct {
	CharacterSet pulumi.StringPtrInput `pulumi:"characterSet"`
	InstanceId   pulumi.StringInput    `pulumi:"instanceId"`
	LcCollate    pulumi.StringPtrInput `pulumi:"lcCollate"`
	Name         pulumi.StringPtrInput `pulumi:"name"`
	Owner        pulumi.StringPtrInput `pulumi:"owner"`
	Region       pulumi.StringPtrInput `pulumi:"region"`
	Size         pulumi.IntPtrInput    `pulumi:"size"`
}

func (GetRdsPgDatabasesOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetRdsPgDatabasesArgs)(nil)).Elem()
}

// A collection of values returned by getRdsPgDatabases.
type GetRdsPgDatabasesResultOutput struct{ *pulumi.OutputState }

func (GetRdsPgDatabasesResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetRdsPgDatabasesResult)(nil)).Elem()
}

func (o GetRdsPgDatabasesResultOutput) ToGetRdsPgDatabasesResultOutput() GetRdsPgDatabasesResultOutput {
	return o
}

func (o GetRdsPgDatabasesResultOutput) ToGetRdsPgDatabasesResultOutputWithContext(ctx context.Context) GetRdsPgDatabasesResultOutput {
	return o
}

func (o GetRdsPgDatabasesResultOutput) CharacterSet() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetRdsPgDatabasesResult) *string { return v.CharacterSet }).(pulumi.StringPtrOutput)
}

func (o GetRdsPgDatabasesResultOutput) Databases() GetRdsPgDatabasesDatabaseArrayOutput {
	return o.ApplyT(func(v GetRdsPgDatabasesResult) []GetRdsPgDatabasesDatabase { return v.Databases }).(GetRdsPgDatabasesDatabaseArrayOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetRdsPgDatabasesResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetRdsPgDatabasesResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetRdsPgDatabasesResultOutput) InstanceId() pulumi.StringOutput {
	return o.ApplyT(func(v GetRdsPgDatabasesResult) string { return v.InstanceId }).(pulumi.StringOutput)
}

func (o GetRdsPgDatabasesResultOutput) LcCollate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetRdsPgDatabasesResult) *string { return v.LcCollate }).(pulumi.StringPtrOutput)
}

func (o GetRdsPgDatabasesResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetRdsPgDatabasesResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o GetRdsPgDatabasesResultOutput) Owner() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetRdsPgDatabasesResult) *string { return v.Owner }).(pulumi.StringPtrOutput)
}

func (o GetRdsPgDatabasesResultOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v GetRdsPgDatabasesResult) string { return v.Region }).(pulumi.StringOutput)
}

func (o GetRdsPgDatabasesResultOutput) Size() pulumi.IntPtrOutput {
	return o.ApplyT(func(v GetRdsPgDatabasesResult) *int { return v.Size }).(pulumi.IntPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetRdsPgDatabasesResultOutput{})
}
