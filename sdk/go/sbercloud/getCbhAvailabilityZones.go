// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package sbercloud

import (
	"context"
	"reflect"

	"github.com/masikrus/pulumi-sbercloud/sdk/go/sbercloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetCbhAvailabilityZones(ctx *pulumi.Context, args *GetCbhAvailabilityZonesArgs, opts ...pulumi.InvokeOption) (*GetCbhAvailabilityZonesResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetCbhAvailabilityZonesResult
	err := ctx.Invoke("sbercloud:index/getCbhAvailabilityZones:getCbhAvailabilityZones", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getCbhAvailabilityZones.
type GetCbhAvailabilityZonesArgs struct {
	DisplayName *string `pulumi:"displayName"`
	Name        *string `pulumi:"name"`
	Region      *string `pulumi:"region"`
}

// A collection of values returned by getCbhAvailabilityZones.
type GetCbhAvailabilityZonesResult struct {
	AvailabilityZones []GetCbhAvailabilityZonesAvailabilityZone `pulumi:"availabilityZones"`
	DisplayName       *string                                   `pulumi:"displayName"`
	// The provider-assigned unique ID for this managed resource.
	Id     string  `pulumi:"id"`
	Name   *string `pulumi:"name"`
	Region string  `pulumi:"region"`
}

func GetCbhAvailabilityZonesOutput(ctx *pulumi.Context, args GetCbhAvailabilityZonesOutputArgs, opts ...pulumi.InvokeOption) GetCbhAvailabilityZonesResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (GetCbhAvailabilityZonesResultOutput, error) {
			args := v.(GetCbhAvailabilityZonesArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("sbercloud:index/getCbhAvailabilityZones:getCbhAvailabilityZones", args, GetCbhAvailabilityZonesResultOutput{}, options).(GetCbhAvailabilityZonesResultOutput), nil
		}).(GetCbhAvailabilityZonesResultOutput)
}

// A collection of arguments for invoking getCbhAvailabilityZones.
type GetCbhAvailabilityZonesOutputArgs struct {
	DisplayName pulumi.StringPtrInput `pulumi:"displayName"`
	Name        pulumi.StringPtrInput `pulumi:"name"`
	Region      pulumi.StringPtrInput `pulumi:"region"`
}

func (GetCbhAvailabilityZonesOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetCbhAvailabilityZonesArgs)(nil)).Elem()
}

// A collection of values returned by getCbhAvailabilityZones.
type GetCbhAvailabilityZonesResultOutput struct{ *pulumi.OutputState }

func (GetCbhAvailabilityZonesResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetCbhAvailabilityZonesResult)(nil)).Elem()
}

func (o GetCbhAvailabilityZonesResultOutput) ToGetCbhAvailabilityZonesResultOutput() GetCbhAvailabilityZonesResultOutput {
	return o
}

func (o GetCbhAvailabilityZonesResultOutput) ToGetCbhAvailabilityZonesResultOutputWithContext(ctx context.Context) GetCbhAvailabilityZonesResultOutput {
	return o
}

func (o GetCbhAvailabilityZonesResultOutput) AvailabilityZones() GetCbhAvailabilityZonesAvailabilityZoneArrayOutput {
	return o.ApplyT(func(v GetCbhAvailabilityZonesResult) []GetCbhAvailabilityZonesAvailabilityZone {
		return v.AvailabilityZones
	}).(GetCbhAvailabilityZonesAvailabilityZoneArrayOutput)
}

func (o GetCbhAvailabilityZonesResultOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetCbhAvailabilityZonesResult) *string { return v.DisplayName }).(pulumi.StringPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetCbhAvailabilityZonesResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetCbhAvailabilityZonesResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetCbhAvailabilityZonesResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetCbhAvailabilityZonesResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o GetCbhAvailabilityZonesResultOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v GetCbhAvailabilityZonesResult) string { return v.Region }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetCbhAvailabilityZonesResultOutput{})
}
