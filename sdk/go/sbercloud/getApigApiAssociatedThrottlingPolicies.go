// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package sbercloud

import (
	"context"
	"reflect"

	"github.com/masikrus/pulumi-sbercloud/sdk/go/sbercloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetApigApiAssociatedThrottlingPolicies(ctx *pulumi.Context, args *GetApigApiAssociatedThrottlingPoliciesArgs, opts ...pulumi.InvokeOption) (*GetApigApiAssociatedThrottlingPoliciesResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetApigApiAssociatedThrottlingPoliciesResult
	err := ctx.Invoke("sbercloud:index/getApigApiAssociatedThrottlingPolicies:getApigApiAssociatedThrottlingPolicies", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getApigApiAssociatedThrottlingPolicies.
type GetApigApiAssociatedThrottlingPoliciesArgs struct {
	ApiId      string  `pulumi:"apiId"`
	EnvName    *string `pulumi:"envName"`
	InstanceId string  `pulumi:"instanceId"`
	Name       *string `pulumi:"name"`
	PeriodUnit *string `pulumi:"periodUnit"`
	PolicyId   *string `pulumi:"policyId"`
	Region     *string `pulumi:"region"`
	Type       *string `pulumi:"type"`
}

// A collection of values returned by getApigApiAssociatedThrottlingPolicies.
type GetApigApiAssociatedThrottlingPoliciesResult struct {
	ApiId   string  `pulumi:"apiId"`
	EnvName *string `pulumi:"envName"`
	// The provider-assigned unique ID for this managed resource.
	Id         string                                         `pulumi:"id"`
	InstanceId string                                         `pulumi:"instanceId"`
	Name       *string                                        `pulumi:"name"`
	PeriodUnit *string                                        `pulumi:"periodUnit"`
	Policies   []GetApigApiAssociatedThrottlingPoliciesPolicy `pulumi:"policies"`
	PolicyId   *string                                        `pulumi:"policyId"`
	Region     string                                         `pulumi:"region"`
	Type       *string                                        `pulumi:"type"`
}

func GetApigApiAssociatedThrottlingPoliciesOutput(ctx *pulumi.Context, args GetApigApiAssociatedThrottlingPoliciesOutputArgs, opts ...pulumi.InvokeOption) GetApigApiAssociatedThrottlingPoliciesResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (GetApigApiAssociatedThrottlingPoliciesResultOutput, error) {
			args := v.(GetApigApiAssociatedThrottlingPoliciesArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("sbercloud:index/getApigApiAssociatedThrottlingPolicies:getApigApiAssociatedThrottlingPolicies", args, GetApigApiAssociatedThrottlingPoliciesResultOutput{}, options).(GetApigApiAssociatedThrottlingPoliciesResultOutput), nil
		}).(GetApigApiAssociatedThrottlingPoliciesResultOutput)
}

// A collection of arguments for invoking getApigApiAssociatedThrottlingPolicies.
type GetApigApiAssociatedThrottlingPoliciesOutputArgs struct {
	ApiId      pulumi.StringInput    `pulumi:"apiId"`
	EnvName    pulumi.StringPtrInput `pulumi:"envName"`
	InstanceId pulumi.StringInput    `pulumi:"instanceId"`
	Name       pulumi.StringPtrInput `pulumi:"name"`
	PeriodUnit pulumi.StringPtrInput `pulumi:"periodUnit"`
	PolicyId   pulumi.StringPtrInput `pulumi:"policyId"`
	Region     pulumi.StringPtrInput `pulumi:"region"`
	Type       pulumi.StringPtrInput `pulumi:"type"`
}

func (GetApigApiAssociatedThrottlingPoliciesOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetApigApiAssociatedThrottlingPoliciesArgs)(nil)).Elem()
}

// A collection of values returned by getApigApiAssociatedThrottlingPolicies.
type GetApigApiAssociatedThrottlingPoliciesResultOutput struct{ *pulumi.OutputState }

func (GetApigApiAssociatedThrottlingPoliciesResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetApigApiAssociatedThrottlingPoliciesResult)(nil)).Elem()
}

func (o GetApigApiAssociatedThrottlingPoliciesResultOutput) ToGetApigApiAssociatedThrottlingPoliciesResultOutput() GetApigApiAssociatedThrottlingPoliciesResultOutput {
	return o
}

func (o GetApigApiAssociatedThrottlingPoliciesResultOutput) ToGetApigApiAssociatedThrottlingPoliciesResultOutputWithContext(ctx context.Context) GetApigApiAssociatedThrottlingPoliciesResultOutput {
	return o
}

func (o GetApigApiAssociatedThrottlingPoliciesResultOutput) ApiId() pulumi.StringOutput {
	return o.ApplyT(func(v GetApigApiAssociatedThrottlingPoliciesResult) string { return v.ApiId }).(pulumi.StringOutput)
}

func (o GetApigApiAssociatedThrottlingPoliciesResultOutput) EnvName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetApigApiAssociatedThrottlingPoliciesResult) *string { return v.EnvName }).(pulumi.StringPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetApigApiAssociatedThrottlingPoliciesResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetApigApiAssociatedThrottlingPoliciesResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetApigApiAssociatedThrottlingPoliciesResultOutput) InstanceId() pulumi.StringOutput {
	return o.ApplyT(func(v GetApigApiAssociatedThrottlingPoliciesResult) string { return v.InstanceId }).(pulumi.StringOutput)
}

func (o GetApigApiAssociatedThrottlingPoliciesResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetApigApiAssociatedThrottlingPoliciesResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o GetApigApiAssociatedThrottlingPoliciesResultOutput) PeriodUnit() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetApigApiAssociatedThrottlingPoliciesResult) *string { return v.PeriodUnit }).(pulumi.StringPtrOutput)
}

func (o GetApigApiAssociatedThrottlingPoliciesResultOutput) Policies() GetApigApiAssociatedThrottlingPoliciesPolicyArrayOutput {
	return o.ApplyT(func(v GetApigApiAssociatedThrottlingPoliciesResult) []GetApigApiAssociatedThrottlingPoliciesPolicy {
		return v.Policies
	}).(GetApigApiAssociatedThrottlingPoliciesPolicyArrayOutput)
}

func (o GetApigApiAssociatedThrottlingPoliciesResultOutput) PolicyId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetApigApiAssociatedThrottlingPoliciesResult) *string { return v.PolicyId }).(pulumi.StringPtrOutput)
}

func (o GetApigApiAssociatedThrottlingPoliciesResultOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v GetApigApiAssociatedThrottlingPoliciesResult) string { return v.Region }).(pulumi.StringOutput)
}

func (o GetApigApiAssociatedThrottlingPoliciesResultOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetApigApiAssociatedThrottlingPoliciesResult) *string { return v.Type }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetApigApiAssociatedThrottlingPoliciesResultOutput{})
}
