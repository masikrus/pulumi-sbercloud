// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package sbercloud

import (
	"context"
	"reflect"

	"github.com/masikrus/pulumi-sbercloud/sdk/go/sbercloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetApigApiAssociatedPlugins(ctx *pulumi.Context, args *GetApigApiAssociatedPluginsArgs, opts ...pulumi.InvokeOption) (*GetApigApiAssociatedPluginsResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetApigApiAssociatedPluginsResult
	err := ctx.Invoke("sbercloud:index/getApigApiAssociatedPlugins:getApigApiAssociatedPlugins", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getApigApiAssociatedPlugins.
type GetApigApiAssociatedPluginsArgs struct {
	ApiId      string  `pulumi:"apiId"`
	EnvId      *string `pulumi:"envId"`
	EnvName    *string `pulumi:"envName"`
	InstanceId string  `pulumi:"instanceId"`
	Name       *string `pulumi:"name"`
	PluginId   *string `pulumi:"pluginId"`
	Region     *string `pulumi:"region"`
	Type       *string `pulumi:"type"`
}

// A collection of values returned by getApigApiAssociatedPlugins.
type GetApigApiAssociatedPluginsResult struct {
	ApiId   string  `pulumi:"apiId"`
	EnvId   *string `pulumi:"envId"`
	EnvName *string `pulumi:"envName"`
	// The provider-assigned unique ID for this managed resource.
	Id         string                              `pulumi:"id"`
	InstanceId string                              `pulumi:"instanceId"`
	Name       *string                             `pulumi:"name"`
	PluginId   *string                             `pulumi:"pluginId"`
	Plugins    []GetApigApiAssociatedPluginsPlugin `pulumi:"plugins"`
	Region     string                              `pulumi:"region"`
	Type       *string                             `pulumi:"type"`
}

func GetApigApiAssociatedPluginsOutput(ctx *pulumi.Context, args GetApigApiAssociatedPluginsOutputArgs, opts ...pulumi.InvokeOption) GetApigApiAssociatedPluginsResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (GetApigApiAssociatedPluginsResultOutput, error) {
			args := v.(GetApigApiAssociatedPluginsArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("sbercloud:index/getApigApiAssociatedPlugins:getApigApiAssociatedPlugins", args, GetApigApiAssociatedPluginsResultOutput{}, options).(GetApigApiAssociatedPluginsResultOutput), nil
		}).(GetApigApiAssociatedPluginsResultOutput)
}

// A collection of arguments for invoking getApigApiAssociatedPlugins.
type GetApigApiAssociatedPluginsOutputArgs struct {
	ApiId      pulumi.StringInput    `pulumi:"apiId"`
	EnvId      pulumi.StringPtrInput `pulumi:"envId"`
	EnvName    pulumi.StringPtrInput `pulumi:"envName"`
	InstanceId pulumi.StringInput    `pulumi:"instanceId"`
	Name       pulumi.StringPtrInput `pulumi:"name"`
	PluginId   pulumi.StringPtrInput `pulumi:"pluginId"`
	Region     pulumi.StringPtrInput `pulumi:"region"`
	Type       pulumi.StringPtrInput `pulumi:"type"`
}

func (GetApigApiAssociatedPluginsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetApigApiAssociatedPluginsArgs)(nil)).Elem()
}

// A collection of values returned by getApigApiAssociatedPlugins.
type GetApigApiAssociatedPluginsResultOutput struct{ *pulumi.OutputState }

func (GetApigApiAssociatedPluginsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetApigApiAssociatedPluginsResult)(nil)).Elem()
}

func (o GetApigApiAssociatedPluginsResultOutput) ToGetApigApiAssociatedPluginsResultOutput() GetApigApiAssociatedPluginsResultOutput {
	return o
}

func (o GetApigApiAssociatedPluginsResultOutput) ToGetApigApiAssociatedPluginsResultOutputWithContext(ctx context.Context) GetApigApiAssociatedPluginsResultOutput {
	return o
}

func (o GetApigApiAssociatedPluginsResultOutput) ApiId() pulumi.StringOutput {
	return o.ApplyT(func(v GetApigApiAssociatedPluginsResult) string { return v.ApiId }).(pulumi.StringOutput)
}

func (o GetApigApiAssociatedPluginsResultOutput) EnvId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetApigApiAssociatedPluginsResult) *string { return v.EnvId }).(pulumi.StringPtrOutput)
}

func (o GetApigApiAssociatedPluginsResultOutput) EnvName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetApigApiAssociatedPluginsResult) *string { return v.EnvName }).(pulumi.StringPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetApigApiAssociatedPluginsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetApigApiAssociatedPluginsResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetApigApiAssociatedPluginsResultOutput) InstanceId() pulumi.StringOutput {
	return o.ApplyT(func(v GetApigApiAssociatedPluginsResult) string { return v.InstanceId }).(pulumi.StringOutput)
}

func (o GetApigApiAssociatedPluginsResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetApigApiAssociatedPluginsResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o GetApigApiAssociatedPluginsResultOutput) PluginId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetApigApiAssociatedPluginsResult) *string { return v.PluginId }).(pulumi.StringPtrOutput)
}

func (o GetApigApiAssociatedPluginsResultOutput) Plugins() GetApigApiAssociatedPluginsPluginArrayOutput {
	return o.ApplyT(func(v GetApigApiAssociatedPluginsResult) []GetApigApiAssociatedPluginsPlugin { return v.Plugins }).(GetApigApiAssociatedPluginsPluginArrayOutput)
}

func (o GetApigApiAssociatedPluginsResultOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v GetApigApiAssociatedPluginsResult) string { return v.Region }).(pulumi.StringOutput)
}

func (o GetApigApiAssociatedPluginsResultOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetApigApiAssociatedPluginsResult) *string { return v.Type }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetApigApiAssociatedPluginsResultOutput{})
}
