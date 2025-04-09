// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package sbercloud

import (
	"context"
	"reflect"

	"github.com/masikrus/pulumi-sbercloud/sdk/go/sbercloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetApigApiBasicConfigurations(ctx *pulumi.Context, args *GetApigApiBasicConfigurationsArgs, opts ...pulumi.InvokeOption) (*GetApigApiBasicConfigurationsResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetApigApiBasicConfigurationsResult
	err := ctx.Invoke("sbercloud:index/getApigApiBasicConfigurations:getApigApiBasicConfigurations", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getApigApiBasicConfigurations.
type GetApigApiBasicConfigurationsArgs struct {
	ApiId                  *string `pulumi:"apiId"`
	BackendType            *string `pulumi:"backendType"`
	EnvId                  *string `pulumi:"envId"`
	EnvName                *string `pulumi:"envName"`
	GroupId                *string `pulumi:"groupId"`
	InstanceId             string  `pulumi:"instanceId"`
	Name                   *string `pulumi:"name"`
	PreciseSearch          *string `pulumi:"preciseSearch"`
	Region                 *string `pulumi:"region"`
	RequestMethod          *string `pulumi:"requestMethod"`
	RequestPath            *string `pulumi:"requestPath"`
	RequestProtocol        *string `pulumi:"requestProtocol"`
	SecurityAuthentication *string `pulumi:"securityAuthentication"`
	Type                   *string `pulumi:"type"`
	VpcChannelName         *string `pulumi:"vpcChannelName"`
}

// A collection of values returned by getApigApiBasicConfigurations.
type GetApigApiBasicConfigurationsResult struct {
	ApiId          *string                                      `pulumi:"apiId"`
	BackendType    *string                                      `pulumi:"backendType"`
	Configurations []GetApigApiBasicConfigurationsConfiguration `pulumi:"configurations"`
	EnvId          *string                                      `pulumi:"envId"`
	EnvName        *string                                      `pulumi:"envName"`
	GroupId        *string                                      `pulumi:"groupId"`
	// The provider-assigned unique ID for this managed resource.
	Id                     string  `pulumi:"id"`
	InstanceId             string  `pulumi:"instanceId"`
	Name                   *string `pulumi:"name"`
	PreciseSearch          *string `pulumi:"preciseSearch"`
	Region                 string  `pulumi:"region"`
	RequestMethod          *string `pulumi:"requestMethod"`
	RequestPath            *string `pulumi:"requestPath"`
	RequestProtocol        *string `pulumi:"requestProtocol"`
	SecurityAuthentication *string `pulumi:"securityAuthentication"`
	Type                   *string `pulumi:"type"`
	VpcChannelName         *string `pulumi:"vpcChannelName"`
}

func GetApigApiBasicConfigurationsOutput(ctx *pulumi.Context, args GetApigApiBasicConfigurationsOutputArgs, opts ...pulumi.InvokeOption) GetApigApiBasicConfigurationsResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (GetApigApiBasicConfigurationsResultOutput, error) {
			args := v.(GetApigApiBasicConfigurationsArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("sbercloud:index/getApigApiBasicConfigurations:getApigApiBasicConfigurations", args, GetApigApiBasicConfigurationsResultOutput{}, options).(GetApigApiBasicConfigurationsResultOutput), nil
		}).(GetApigApiBasicConfigurationsResultOutput)
}

// A collection of arguments for invoking getApigApiBasicConfigurations.
type GetApigApiBasicConfigurationsOutputArgs struct {
	ApiId                  pulumi.StringPtrInput `pulumi:"apiId"`
	BackendType            pulumi.StringPtrInput `pulumi:"backendType"`
	EnvId                  pulumi.StringPtrInput `pulumi:"envId"`
	EnvName                pulumi.StringPtrInput `pulumi:"envName"`
	GroupId                pulumi.StringPtrInput `pulumi:"groupId"`
	InstanceId             pulumi.StringInput    `pulumi:"instanceId"`
	Name                   pulumi.StringPtrInput `pulumi:"name"`
	PreciseSearch          pulumi.StringPtrInput `pulumi:"preciseSearch"`
	Region                 pulumi.StringPtrInput `pulumi:"region"`
	RequestMethod          pulumi.StringPtrInput `pulumi:"requestMethod"`
	RequestPath            pulumi.StringPtrInput `pulumi:"requestPath"`
	RequestProtocol        pulumi.StringPtrInput `pulumi:"requestProtocol"`
	SecurityAuthentication pulumi.StringPtrInput `pulumi:"securityAuthentication"`
	Type                   pulumi.StringPtrInput `pulumi:"type"`
	VpcChannelName         pulumi.StringPtrInput `pulumi:"vpcChannelName"`
}

func (GetApigApiBasicConfigurationsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetApigApiBasicConfigurationsArgs)(nil)).Elem()
}

// A collection of values returned by getApigApiBasicConfigurations.
type GetApigApiBasicConfigurationsResultOutput struct{ *pulumi.OutputState }

func (GetApigApiBasicConfigurationsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetApigApiBasicConfigurationsResult)(nil)).Elem()
}

func (o GetApigApiBasicConfigurationsResultOutput) ToGetApigApiBasicConfigurationsResultOutput() GetApigApiBasicConfigurationsResultOutput {
	return o
}

func (o GetApigApiBasicConfigurationsResultOutput) ToGetApigApiBasicConfigurationsResultOutputWithContext(ctx context.Context) GetApigApiBasicConfigurationsResultOutput {
	return o
}

func (o GetApigApiBasicConfigurationsResultOutput) ApiId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetApigApiBasicConfigurationsResult) *string { return v.ApiId }).(pulumi.StringPtrOutput)
}

func (o GetApigApiBasicConfigurationsResultOutput) BackendType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetApigApiBasicConfigurationsResult) *string { return v.BackendType }).(pulumi.StringPtrOutput)
}

func (o GetApigApiBasicConfigurationsResultOutput) Configurations() GetApigApiBasicConfigurationsConfigurationArrayOutput {
	return o.ApplyT(func(v GetApigApiBasicConfigurationsResult) []GetApigApiBasicConfigurationsConfiguration {
		return v.Configurations
	}).(GetApigApiBasicConfigurationsConfigurationArrayOutput)
}

func (o GetApigApiBasicConfigurationsResultOutput) EnvId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetApigApiBasicConfigurationsResult) *string { return v.EnvId }).(pulumi.StringPtrOutput)
}

func (o GetApigApiBasicConfigurationsResultOutput) EnvName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetApigApiBasicConfigurationsResult) *string { return v.EnvName }).(pulumi.StringPtrOutput)
}

func (o GetApigApiBasicConfigurationsResultOutput) GroupId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetApigApiBasicConfigurationsResult) *string { return v.GroupId }).(pulumi.StringPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetApigApiBasicConfigurationsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetApigApiBasicConfigurationsResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetApigApiBasicConfigurationsResultOutput) InstanceId() pulumi.StringOutput {
	return o.ApplyT(func(v GetApigApiBasicConfigurationsResult) string { return v.InstanceId }).(pulumi.StringOutput)
}

func (o GetApigApiBasicConfigurationsResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetApigApiBasicConfigurationsResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o GetApigApiBasicConfigurationsResultOutput) PreciseSearch() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetApigApiBasicConfigurationsResult) *string { return v.PreciseSearch }).(pulumi.StringPtrOutput)
}

func (o GetApigApiBasicConfigurationsResultOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v GetApigApiBasicConfigurationsResult) string { return v.Region }).(pulumi.StringOutput)
}

func (o GetApigApiBasicConfigurationsResultOutput) RequestMethod() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetApigApiBasicConfigurationsResult) *string { return v.RequestMethod }).(pulumi.StringPtrOutput)
}

func (o GetApigApiBasicConfigurationsResultOutput) RequestPath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetApigApiBasicConfigurationsResult) *string { return v.RequestPath }).(pulumi.StringPtrOutput)
}

func (o GetApigApiBasicConfigurationsResultOutput) RequestProtocol() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetApigApiBasicConfigurationsResult) *string { return v.RequestProtocol }).(pulumi.StringPtrOutput)
}

func (o GetApigApiBasicConfigurationsResultOutput) SecurityAuthentication() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetApigApiBasicConfigurationsResult) *string { return v.SecurityAuthentication }).(pulumi.StringPtrOutput)
}

func (o GetApigApiBasicConfigurationsResultOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetApigApiBasicConfigurationsResult) *string { return v.Type }).(pulumi.StringPtrOutput)
}

func (o GetApigApiBasicConfigurationsResultOutput) VpcChannelName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetApigApiBasicConfigurationsResult) *string { return v.VpcChannelName }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetApigApiBasicConfigurationsResultOutput{})
}
