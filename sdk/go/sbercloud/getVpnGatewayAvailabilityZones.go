// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package sbercloud

import (
	"context"
	"reflect"

	"github.com/masikrus/pulumi-sbercloud/sdk/go/sbercloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetVpnGatewayAvailabilityZones(ctx *pulumi.Context, args *GetVpnGatewayAvailabilityZonesArgs, opts ...pulumi.InvokeOption) (*GetVpnGatewayAvailabilityZonesResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetVpnGatewayAvailabilityZonesResult
	err := ctx.Invoke("sbercloud:index/getVpnGatewayAvailabilityZones:getVpnGatewayAvailabilityZones", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getVpnGatewayAvailabilityZones.
type GetVpnGatewayAvailabilityZonesArgs struct {
	AttachmentType *string `pulumi:"attachmentType"`
	Flavor         string  `pulumi:"flavor"`
	Region         *string `pulumi:"region"`
}

// A collection of values returned by getVpnGatewayAvailabilityZones.
type GetVpnGatewayAvailabilityZonesResult struct {
	AttachmentType *string `pulumi:"attachmentType"`
	Flavor         string  `pulumi:"flavor"`
	// The provider-assigned unique ID for this managed resource.
	Id     string   `pulumi:"id"`
	Names  []string `pulumi:"names"`
	Region string   `pulumi:"region"`
}

func GetVpnGatewayAvailabilityZonesOutput(ctx *pulumi.Context, args GetVpnGatewayAvailabilityZonesOutputArgs, opts ...pulumi.InvokeOption) GetVpnGatewayAvailabilityZonesResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (GetVpnGatewayAvailabilityZonesResultOutput, error) {
			args := v.(GetVpnGatewayAvailabilityZonesArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("sbercloud:index/getVpnGatewayAvailabilityZones:getVpnGatewayAvailabilityZones", args, GetVpnGatewayAvailabilityZonesResultOutput{}, options).(GetVpnGatewayAvailabilityZonesResultOutput), nil
		}).(GetVpnGatewayAvailabilityZonesResultOutput)
}

// A collection of arguments for invoking getVpnGatewayAvailabilityZones.
type GetVpnGatewayAvailabilityZonesOutputArgs struct {
	AttachmentType pulumi.StringPtrInput `pulumi:"attachmentType"`
	Flavor         pulumi.StringInput    `pulumi:"flavor"`
	Region         pulumi.StringPtrInput `pulumi:"region"`
}

func (GetVpnGatewayAvailabilityZonesOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetVpnGatewayAvailabilityZonesArgs)(nil)).Elem()
}

// A collection of values returned by getVpnGatewayAvailabilityZones.
type GetVpnGatewayAvailabilityZonesResultOutput struct{ *pulumi.OutputState }

func (GetVpnGatewayAvailabilityZonesResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetVpnGatewayAvailabilityZonesResult)(nil)).Elem()
}

func (o GetVpnGatewayAvailabilityZonesResultOutput) ToGetVpnGatewayAvailabilityZonesResultOutput() GetVpnGatewayAvailabilityZonesResultOutput {
	return o
}

func (o GetVpnGatewayAvailabilityZonesResultOutput) ToGetVpnGatewayAvailabilityZonesResultOutputWithContext(ctx context.Context) GetVpnGatewayAvailabilityZonesResultOutput {
	return o
}

func (o GetVpnGatewayAvailabilityZonesResultOutput) AttachmentType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetVpnGatewayAvailabilityZonesResult) *string { return v.AttachmentType }).(pulumi.StringPtrOutput)
}

func (o GetVpnGatewayAvailabilityZonesResultOutput) Flavor() pulumi.StringOutput {
	return o.ApplyT(func(v GetVpnGatewayAvailabilityZonesResult) string { return v.Flavor }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetVpnGatewayAvailabilityZonesResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetVpnGatewayAvailabilityZonesResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetVpnGatewayAvailabilityZonesResultOutput) Names() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetVpnGatewayAvailabilityZonesResult) []string { return v.Names }).(pulumi.StringArrayOutput)
}

func (o GetVpnGatewayAvailabilityZonesResultOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v GetVpnGatewayAvailabilityZonesResult) string { return v.Region }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetVpnGatewayAvailabilityZonesResultOutput{})
}
