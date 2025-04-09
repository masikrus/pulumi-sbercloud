// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package sbercloud

import (
	"context"
	"reflect"

	"github.com/masikrus/pulumi-sbercloud/sdk/go/sbercloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetCceAddonTemplate(ctx *pulumi.Context, args *GetCceAddonTemplateArgs, opts ...pulumi.InvokeOption) (*GetCceAddonTemplateResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetCceAddonTemplateResult
	err := ctx.Invoke("sbercloud:index/getCceAddonTemplate:getCceAddonTemplate", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getCceAddonTemplate.
type GetCceAddonTemplateArgs struct {
	ClusterId string  `pulumi:"clusterId"`
	Name      string  `pulumi:"name"`
	Region    *string `pulumi:"region"`
	Version   string  `pulumi:"version"`
}

// A collection of values returned by getCceAddonTemplate.
type GetCceAddonTemplateResult struct {
	ClusterId   string `pulumi:"clusterId"`
	Description string `pulumi:"description"`
	// The provider-assigned unique ID for this managed resource.
	Id              string                              `pulumi:"id"`
	Name            string                              `pulumi:"name"`
	Region          string                              `pulumi:"region"`
	Spec            string                              `pulumi:"spec"`
	Stable          bool                                `pulumi:"stable"`
	SupportVersions []GetCceAddonTemplateSupportVersion `pulumi:"supportVersions"`
	Version         string                              `pulumi:"version"`
}

func GetCceAddonTemplateOutput(ctx *pulumi.Context, args GetCceAddonTemplateOutputArgs, opts ...pulumi.InvokeOption) GetCceAddonTemplateResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (GetCceAddonTemplateResultOutput, error) {
			args := v.(GetCceAddonTemplateArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("sbercloud:index/getCceAddonTemplate:getCceAddonTemplate", args, GetCceAddonTemplateResultOutput{}, options).(GetCceAddonTemplateResultOutput), nil
		}).(GetCceAddonTemplateResultOutput)
}

// A collection of arguments for invoking getCceAddonTemplate.
type GetCceAddonTemplateOutputArgs struct {
	ClusterId pulumi.StringInput    `pulumi:"clusterId"`
	Name      pulumi.StringInput    `pulumi:"name"`
	Region    pulumi.StringPtrInput `pulumi:"region"`
	Version   pulumi.StringInput    `pulumi:"version"`
}

func (GetCceAddonTemplateOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetCceAddonTemplateArgs)(nil)).Elem()
}

// A collection of values returned by getCceAddonTemplate.
type GetCceAddonTemplateResultOutput struct{ *pulumi.OutputState }

func (GetCceAddonTemplateResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetCceAddonTemplateResult)(nil)).Elem()
}

func (o GetCceAddonTemplateResultOutput) ToGetCceAddonTemplateResultOutput() GetCceAddonTemplateResultOutput {
	return o
}

func (o GetCceAddonTemplateResultOutput) ToGetCceAddonTemplateResultOutputWithContext(ctx context.Context) GetCceAddonTemplateResultOutput {
	return o
}

func (o GetCceAddonTemplateResultOutput) ClusterId() pulumi.StringOutput {
	return o.ApplyT(func(v GetCceAddonTemplateResult) string { return v.ClusterId }).(pulumi.StringOutput)
}

func (o GetCceAddonTemplateResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v GetCceAddonTemplateResult) string { return v.Description }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetCceAddonTemplateResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetCceAddonTemplateResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetCceAddonTemplateResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v GetCceAddonTemplateResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o GetCceAddonTemplateResultOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v GetCceAddonTemplateResult) string { return v.Region }).(pulumi.StringOutput)
}

func (o GetCceAddonTemplateResultOutput) Spec() pulumi.StringOutput {
	return o.ApplyT(func(v GetCceAddonTemplateResult) string { return v.Spec }).(pulumi.StringOutput)
}

func (o GetCceAddonTemplateResultOutput) Stable() pulumi.BoolOutput {
	return o.ApplyT(func(v GetCceAddonTemplateResult) bool { return v.Stable }).(pulumi.BoolOutput)
}

func (o GetCceAddonTemplateResultOutput) SupportVersions() GetCceAddonTemplateSupportVersionArrayOutput {
	return o.ApplyT(func(v GetCceAddonTemplateResult) []GetCceAddonTemplateSupportVersion { return v.SupportVersions }).(GetCceAddonTemplateSupportVersionArrayOutput)
}

func (o GetCceAddonTemplateResultOutput) Version() pulumi.StringOutput {
	return o.ApplyT(func(v GetCceAddonTemplateResult) string { return v.Version }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetCceAddonTemplateResultOutput{})
}
