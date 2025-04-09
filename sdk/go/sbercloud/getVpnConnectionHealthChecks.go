// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package sbercloud

import (
	"context"
	"reflect"

	"github.com/masikrus/pulumi-sbercloud/sdk/go/sbercloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetVpnConnectionHealthChecks(ctx *pulumi.Context, args *GetVpnConnectionHealthChecksArgs, opts ...pulumi.InvokeOption) (*GetVpnConnectionHealthChecksResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetVpnConnectionHealthChecksResult
	err := ctx.Invoke("sbercloud:index/getVpnConnectionHealthChecks:getVpnConnectionHealthChecks", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getVpnConnectionHealthChecks.
type GetVpnConnectionHealthChecksArgs struct {
	ConnectionId  *string `pulumi:"connectionId"`
	DestinationIp *string `pulumi:"destinationIp"`
	Region        *string `pulumi:"region"`
	SourceIp      *string `pulumi:"sourceIp"`
	Status        *string `pulumi:"status"`
}

// A collection of values returned by getVpnConnectionHealthChecks.
type GetVpnConnectionHealthChecksResult struct {
	ConnectionHealthChecks []GetVpnConnectionHealthChecksConnectionHealthCheck `pulumi:"connectionHealthChecks"`
	ConnectionId           *string                                             `pulumi:"connectionId"`
	DestinationIp          *string                                             `pulumi:"destinationIp"`
	// The provider-assigned unique ID for this managed resource.
	Id       string  `pulumi:"id"`
	Region   string  `pulumi:"region"`
	SourceIp *string `pulumi:"sourceIp"`
	Status   *string `pulumi:"status"`
}

func GetVpnConnectionHealthChecksOutput(ctx *pulumi.Context, args GetVpnConnectionHealthChecksOutputArgs, opts ...pulumi.InvokeOption) GetVpnConnectionHealthChecksResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (GetVpnConnectionHealthChecksResultOutput, error) {
			args := v.(GetVpnConnectionHealthChecksArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("sbercloud:index/getVpnConnectionHealthChecks:getVpnConnectionHealthChecks", args, GetVpnConnectionHealthChecksResultOutput{}, options).(GetVpnConnectionHealthChecksResultOutput), nil
		}).(GetVpnConnectionHealthChecksResultOutput)
}

// A collection of arguments for invoking getVpnConnectionHealthChecks.
type GetVpnConnectionHealthChecksOutputArgs struct {
	ConnectionId  pulumi.StringPtrInput `pulumi:"connectionId"`
	DestinationIp pulumi.StringPtrInput `pulumi:"destinationIp"`
	Region        pulumi.StringPtrInput `pulumi:"region"`
	SourceIp      pulumi.StringPtrInput `pulumi:"sourceIp"`
	Status        pulumi.StringPtrInput `pulumi:"status"`
}

func (GetVpnConnectionHealthChecksOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetVpnConnectionHealthChecksArgs)(nil)).Elem()
}

// A collection of values returned by getVpnConnectionHealthChecks.
type GetVpnConnectionHealthChecksResultOutput struct{ *pulumi.OutputState }

func (GetVpnConnectionHealthChecksResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetVpnConnectionHealthChecksResult)(nil)).Elem()
}

func (o GetVpnConnectionHealthChecksResultOutput) ToGetVpnConnectionHealthChecksResultOutput() GetVpnConnectionHealthChecksResultOutput {
	return o
}

func (o GetVpnConnectionHealthChecksResultOutput) ToGetVpnConnectionHealthChecksResultOutputWithContext(ctx context.Context) GetVpnConnectionHealthChecksResultOutput {
	return o
}

func (o GetVpnConnectionHealthChecksResultOutput) ConnectionHealthChecks() GetVpnConnectionHealthChecksConnectionHealthCheckArrayOutput {
	return o.ApplyT(func(v GetVpnConnectionHealthChecksResult) []GetVpnConnectionHealthChecksConnectionHealthCheck {
		return v.ConnectionHealthChecks
	}).(GetVpnConnectionHealthChecksConnectionHealthCheckArrayOutput)
}

func (o GetVpnConnectionHealthChecksResultOutput) ConnectionId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetVpnConnectionHealthChecksResult) *string { return v.ConnectionId }).(pulumi.StringPtrOutput)
}

func (o GetVpnConnectionHealthChecksResultOutput) DestinationIp() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetVpnConnectionHealthChecksResult) *string { return v.DestinationIp }).(pulumi.StringPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetVpnConnectionHealthChecksResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetVpnConnectionHealthChecksResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetVpnConnectionHealthChecksResultOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v GetVpnConnectionHealthChecksResult) string { return v.Region }).(pulumi.StringOutput)
}

func (o GetVpnConnectionHealthChecksResultOutput) SourceIp() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetVpnConnectionHealthChecksResult) *string { return v.SourceIp }).(pulumi.StringPtrOutput)
}

func (o GetVpnConnectionHealthChecksResultOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetVpnConnectionHealthChecksResult) *string { return v.Status }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetVpnConnectionHealthChecksResultOutput{})
}
