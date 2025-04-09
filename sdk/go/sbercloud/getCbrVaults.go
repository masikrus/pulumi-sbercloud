// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package sbercloud

import (
	"context"
	"reflect"

	"github.com/masikrus/pulumi-sbercloud/sdk/go/sbercloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetCbrVaults(ctx *pulumi.Context, args *GetCbrVaultsArgs, opts ...pulumi.InvokeOption) (*GetCbrVaultsResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetCbrVaultsResult
	err := ctx.Invoke("sbercloud:index/getCbrVaults:getCbrVaults", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getCbrVaults.
type GetCbrVaultsArgs struct {
	AutoExpandEnabled   *bool   `pulumi:"autoExpandEnabled"`
	ConsistentLevel     *string `pulumi:"consistentLevel"`
	EnterpriseProjectId *string `pulumi:"enterpriseProjectId"`
	Name                *string `pulumi:"name"`
	PolicyId            *string `pulumi:"policyId"`
	ProtectionType      *string `pulumi:"protectionType"`
	Region              *string `pulumi:"region"`
	Size                *int    `pulumi:"size"`
	Status              *string `pulumi:"status"`
	Type                *string `pulumi:"type"`
}

// A collection of values returned by getCbrVaults.
type GetCbrVaultsResult struct {
	AutoExpandEnabled   *bool   `pulumi:"autoExpandEnabled"`
	ConsistentLevel     *string `pulumi:"consistentLevel"`
	EnterpriseProjectId *string `pulumi:"enterpriseProjectId"`
	// The provider-assigned unique ID for this managed resource.
	Id             string              `pulumi:"id"`
	Name           *string             `pulumi:"name"`
	PolicyId       *string             `pulumi:"policyId"`
	ProtectionType *string             `pulumi:"protectionType"`
	Region         *string             `pulumi:"region"`
	Size           *int                `pulumi:"size"`
	Status         *string             `pulumi:"status"`
	Type           *string             `pulumi:"type"`
	Vaults         []GetCbrVaultsVault `pulumi:"vaults"`
}

func GetCbrVaultsOutput(ctx *pulumi.Context, args GetCbrVaultsOutputArgs, opts ...pulumi.InvokeOption) GetCbrVaultsResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (GetCbrVaultsResultOutput, error) {
			args := v.(GetCbrVaultsArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("sbercloud:index/getCbrVaults:getCbrVaults", args, GetCbrVaultsResultOutput{}, options).(GetCbrVaultsResultOutput), nil
		}).(GetCbrVaultsResultOutput)
}

// A collection of arguments for invoking getCbrVaults.
type GetCbrVaultsOutputArgs struct {
	AutoExpandEnabled   pulumi.BoolPtrInput   `pulumi:"autoExpandEnabled"`
	ConsistentLevel     pulumi.StringPtrInput `pulumi:"consistentLevel"`
	EnterpriseProjectId pulumi.StringPtrInput `pulumi:"enterpriseProjectId"`
	Name                pulumi.StringPtrInput `pulumi:"name"`
	PolicyId            pulumi.StringPtrInput `pulumi:"policyId"`
	ProtectionType      pulumi.StringPtrInput `pulumi:"protectionType"`
	Region              pulumi.StringPtrInput `pulumi:"region"`
	Size                pulumi.IntPtrInput    `pulumi:"size"`
	Status              pulumi.StringPtrInput `pulumi:"status"`
	Type                pulumi.StringPtrInput `pulumi:"type"`
}

func (GetCbrVaultsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetCbrVaultsArgs)(nil)).Elem()
}

// A collection of values returned by getCbrVaults.
type GetCbrVaultsResultOutput struct{ *pulumi.OutputState }

func (GetCbrVaultsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetCbrVaultsResult)(nil)).Elem()
}

func (o GetCbrVaultsResultOutput) ToGetCbrVaultsResultOutput() GetCbrVaultsResultOutput {
	return o
}

func (o GetCbrVaultsResultOutput) ToGetCbrVaultsResultOutputWithContext(ctx context.Context) GetCbrVaultsResultOutput {
	return o
}

func (o GetCbrVaultsResultOutput) AutoExpandEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v GetCbrVaultsResult) *bool { return v.AutoExpandEnabled }).(pulumi.BoolPtrOutput)
}

func (o GetCbrVaultsResultOutput) ConsistentLevel() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetCbrVaultsResult) *string { return v.ConsistentLevel }).(pulumi.StringPtrOutput)
}

func (o GetCbrVaultsResultOutput) EnterpriseProjectId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetCbrVaultsResult) *string { return v.EnterpriseProjectId }).(pulumi.StringPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetCbrVaultsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetCbrVaultsResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetCbrVaultsResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetCbrVaultsResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o GetCbrVaultsResultOutput) PolicyId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetCbrVaultsResult) *string { return v.PolicyId }).(pulumi.StringPtrOutput)
}

func (o GetCbrVaultsResultOutput) ProtectionType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetCbrVaultsResult) *string { return v.ProtectionType }).(pulumi.StringPtrOutput)
}

func (o GetCbrVaultsResultOutput) Region() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetCbrVaultsResult) *string { return v.Region }).(pulumi.StringPtrOutput)
}

func (o GetCbrVaultsResultOutput) Size() pulumi.IntPtrOutput {
	return o.ApplyT(func(v GetCbrVaultsResult) *int { return v.Size }).(pulumi.IntPtrOutput)
}

func (o GetCbrVaultsResultOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetCbrVaultsResult) *string { return v.Status }).(pulumi.StringPtrOutput)
}

func (o GetCbrVaultsResultOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetCbrVaultsResult) *string { return v.Type }).(pulumi.StringPtrOutput)
}

func (o GetCbrVaultsResultOutput) Vaults() GetCbrVaultsVaultArrayOutput {
	return o.ApplyT(func(v GetCbrVaultsResult) []GetCbrVaultsVault { return v.Vaults }).(GetCbrVaultsVaultArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(GetCbrVaultsResultOutput{})
}
