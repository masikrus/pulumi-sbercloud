// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package sbercloud

import (
	"context"
	"reflect"

	"github.com/masikrus/pulumi-sbercloud/sdk/go/sbercloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetVpcEips(ctx *pulumi.Context, args *GetVpcEipsArgs, opts ...pulumi.InvokeOption) (*GetVpcEipsResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetVpcEipsResult
	err := ctx.Invoke("sbercloud:index/getVpcEips:getVpcEips", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getVpcEips.
type GetVpcEipsArgs struct {
	EnterpriseProjectId *string           `pulumi:"enterpriseProjectId"`
	Ids                 []string          `pulumi:"ids"`
	IpVersion           *int              `pulumi:"ipVersion"`
	PortIds             []string          `pulumi:"portIds"`
	PrivateIps          []string          `pulumi:"privateIps"`
	PublicIps           []string          `pulumi:"publicIps"`
	Region              *string           `pulumi:"region"`
	Tags                map[string]string `pulumi:"tags"`
}

// A collection of values returned by getVpcEips.
type GetVpcEipsResult struct {
	Eips                []GetVpcEipsEip `pulumi:"eips"`
	EnterpriseProjectId *string         `pulumi:"enterpriseProjectId"`
	// The provider-assigned unique ID for this managed resource.
	Id         string            `pulumi:"id"`
	Ids        []string          `pulumi:"ids"`
	IpVersion  *int              `pulumi:"ipVersion"`
	PortIds    []string          `pulumi:"portIds"`
	PrivateIps []string          `pulumi:"privateIps"`
	PublicIps  []string          `pulumi:"publicIps"`
	Region     string            `pulumi:"region"`
	Tags       map[string]string `pulumi:"tags"`
}

func GetVpcEipsOutput(ctx *pulumi.Context, args GetVpcEipsOutputArgs, opts ...pulumi.InvokeOption) GetVpcEipsResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (GetVpcEipsResultOutput, error) {
			args := v.(GetVpcEipsArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("sbercloud:index/getVpcEips:getVpcEips", args, GetVpcEipsResultOutput{}, options).(GetVpcEipsResultOutput), nil
		}).(GetVpcEipsResultOutput)
}

// A collection of arguments for invoking getVpcEips.
type GetVpcEipsOutputArgs struct {
	EnterpriseProjectId pulumi.StringPtrInput   `pulumi:"enterpriseProjectId"`
	Ids                 pulumi.StringArrayInput `pulumi:"ids"`
	IpVersion           pulumi.IntPtrInput      `pulumi:"ipVersion"`
	PortIds             pulumi.StringArrayInput `pulumi:"portIds"`
	PrivateIps          pulumi.StringArrayInput `pulumi:"privateIps"`
	PublicIps           pulumi.StringArrayInput `pulumi:"publicIps"`
	Region              pulumi.StringPtrInput   `pulumi:"region"`
	Tags                pulumi.StringMapInput   `pulumi:"tags"`
}

func (GetVpcEipsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetVpcEipsArgs)(nil)).Elem()
}

// A collection of values returned by getVpcEips.
type GetVpcEipsResultOutput struct{ *pulumi.OutputState }

func (GetVpcEipsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetVpcEipsResult)(nil)).Elem()
}

func (o GetVpcEipsResultOutput) ToGetVpcEipsResultOutput() GetVpcEipsResultOutput {
	return o
}

func (o GetVpcEipsResultOutput) ToGetVpcEipsResultOutputWithContext(ctx context.Context) GetVpcEipsResultOutput {
	return o
}

func (o GetVpcEipsResultOutput) Eips() GetVpcEipsEipArrayOutput {
	return o.ApplyT(func(v GetVpcEipsResult) []GetVpcEipsEip { return v.Eips }).(GetVpcEipsEipArrayOutput)
}

func (o GetVpcEipsResultOutput) EnterpriseProjectId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetVpcEipsResult) *string { return v.EnterpriseProjectId }).(pulumi.StringPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetVpcEipsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetVpcEipsResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetVpcEipsResultOutput) Ids() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetVpcEipsResult) []string { return v.Ids }).(pulumi.StringArrayOutput)
}

func (o GetVpcEipsResultOutput) IpVersion() pulumi.IntPtrOutput {
	return o.ApplyT(func(v GetVpcEipsResult) *int { return v.IpVersion }).(pulumi.IntPtrOutput)
}

func (o GetVpcEipsResultOutput) PortIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetVpcEipsResult) []string { return v.PortIds }).(pulumi.StringArrayOutput)
}

func (o GetVpcEipsResultOutput) PrivateIps() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetVpcEipsResult) []string { return v.PrivateIps }).(pulumi.StringArrayOutput)
}

func (o GetVpcEipsResultOutput) PublicIps() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetVpcEipsResult) []string { return v.PublicIps }).(pulumi.StringArrayOutput)
}

func (o GetVpcEipsResultOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v GetVpcEipsResult) string { return v.Region }).(pulumi.StringOutput)
}

func (o GetVpcEipsResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v GetVpcEipsResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func init() {
	pulumi.RegisterOutputType(GetVpcEipsResultOutput{})
}
