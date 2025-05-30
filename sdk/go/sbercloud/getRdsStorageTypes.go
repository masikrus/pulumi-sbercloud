// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package sbercloud

import (
	"context"
	"reflect"

	"github.com/masikrus/pulumi-sbercloud/sdk/go/sbercloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetRdsStorageTypes(ctx *pulumi.Context, args *GetRdsStorageTypesArgs, opts ...pulumi.InvokeOption) (*GetRdsStorageTypesResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetRdsStorageTypesResult
	err := ctx.Invoke("sbercloud:index/getRdsStorageTypes:getRdsStorageTypes", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getRdsStorageTypes.
type GetRdsStorageTypesArgs struct {
	DbType       string  `pulumi:"dbType"`
	DbVersion    string  `pulumi:"dbVersion"`
	InstanceMode *string `pulumi:"instanceMode"`
	Region       *string `pulumi:"region"`
}

// A collection of values returned by getRdsStorageTypes.
type GetRdsStorageTypesResult struct {
	DbType    string `pulumi:"dbType"`
	DbVersion string `pulumi:"dbVersion"`
	// The provider-assigned unique ID for this managed resource.
	Id           string                          `pulumi:"id"`
	InstanceMode *string                         `pulumi:"instanceMode"`
	Region       string                          `pulumi:"region"`
	StorageTypes []GetRdsStorageTypesStorageType `pulumi:"storageTypes"`
}

func GetRdsStorageTypesOutput(ctx *pulumi.Context, args GetRdsStorageTypesOutputArgs, opts ...pulumi.InvokeOption) GetRdsStorageTypesResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (GetRdsStorageTypesResultOutput, error) {
			args := v.(GetRdsStorageTypesArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("sbercloud:index/getRdsStorageTypes:getRdsStorageTypes", args, GetRdsStorageTypesResultOutput{}, options).(GetRdsStorageTypesResultOutput), nil
		}).(GetRdsStorageTypesResultOutput)
}

// A collection of arguments for invoking getRdsStorageTypes.
type GetRdsStorageTypesOutputArgs struct {
	DbType       pulumi.StringInput    `pulumi:"dbType"`
	DbVersion    pulumi.StringInput    `pulumi:"dbVersion"`
	InstanceMode pulumi.StringPtrInput `pulumi:"instanceMode"`
	Region       pulumi.StringPtrInput `pulumi:"region"`
}

func (GetRdsStorageTypesOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetRdsStorageTypesArgs)(nil)).Elem()
}

// A collection of values returned by getRdsStorageTypes.
type GetRdsStorageTypesResultOutput struct{ *pulumi.OutputState }

func (GetRdsStorageTypesResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetRdsStorageTypesResult)(nil)).Elem()
}

func (o GetRdsStorageTypesResultOutput) ToGetRdsStorageTypesResultOutput() GetRdsStorageTypesResultOutput {
	return o
}

func (o GetRdsStorageTypesResultOutput) ToGetRdsStorageTypesResultOutputWithContext(ctx context.Context) GetRdsStorageTypesResultOutput {
	return o
}

func (o GetRdsStorageTypesResultOutput) DbType() pulumi.StringOutput {
	return o.ApplyT(func(v GetRdsStorageTypesResult) string { return v.DbType }).(pulumi.StringOutput)
}

func (o GetRdsStorageTypesResultOutput) DbVersion() pulumi.StringOutput {
	return o.ApplyT(func(v GetRdsStorageTypesResult) string { return v.DbVersion }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetRdsStorageTypesResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetRdsStorageTypesResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetRdsStorageTypesResultOutput) InstanceMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetRdsStorageTypesResult) *string { return v.InstanceMode }).(pulumi.StringPtrOutput)
}

func (o GetRdsStorageTypesResultOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v GetRdsStorageTypesResult) string { return v.Region }).(pulumi.StringOutput)
}

func (o GetRdsStorageTypesResultOutput) StorageTypes() GetRdsStorageTypesStorageTypeArrayOutput {
	return o.ApplyT(func(v GetRdsStorageTypesResult) []GetRdsStorageTypesStorageType { return v.StorageTypes }).(GetRdsStorageTypesStorageTypeArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(GetRdsStorageTypesResultOutput{})
}
