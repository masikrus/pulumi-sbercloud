// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package sbercloud

import (
	"context"
	"reflect"

	"github.com/masikrus/pulumi-sbercloud/sdk/go/sbercloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetApigApiAssociatedAclPolicies(ctx *pulumi.Context, args *GetApigApiAssociatedAclPoliciesArgs, opts ...pulumi.InvokeOption) (*GetApigApiAssociatedAclPoliciesResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetApigApiAssociatedAclPoliciesResult
	err := ctx.Invoke("sbercloud:index/getApigApiAssociatedAclPolicies:getApigApiAssociatedAclPolicies", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getApigApiAssociatedAclPolicies.
type GetApigApiAssociatedAclPoliciesArgs struct {
	ApiId      string  `pulumi:"apiId"`
	EntityType *string `pulumi:"entityType"`
	EnvId      *string `pulumi:"envId"`
	EnvName    *string `pulumi:"envName"`
	InstanceId string  `pulumi:"instanceId"`
	Name       *string `pulumi:"name"`
	PolicyId   *string `pulumi:"policyId"`
	Region     *string `pulumi:"region"`
	Type       *string `pulumi:"type"`
}

// A collection of values returned by getApigApiAssociatedAclPolicies.
type GetApigApiAssociatedAclPoliciesResult struct {
	ApiId      string  `pulumi:"apiId"`
	EntityType *string `pulumi:"entityType"`
	EnvId      *string `pulumi:"envId"`
	EnvName    *string `pulumi:"envName"`
	// The provider-assigned unique ID for this managed resource.
	Id         string                                  `pulumi:"id"`
	InstanceId string                                  `pulumi:"instanceId"`
	Name       *string                                 `pulumi:"name"`
	Policies   []GetApigApiAssociatedAclPoliciesPolicy `pulumi:"policies"`
	PolicyId   *string                                 `pulumi:"policyId"`
	Region     string                                  `pulumi:"region"`
	Type       *string                                 `pulumi:"type"`
}

func GetApigApiAssociatedAclPoliciesOutput(ctx *pulumi.Context, args GetApigApiAssociatedAclPoliciesOutputArgs, opts ...pulumi.InvokeOption) GetApigApiAssociatedAclPoliciesResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (GetApigApiAssociatedAclPoliciesResultOutput, error) {
			args := v.(GetApigApiAssociatedAclPoliciesArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("sbercloud:index/getApigApiAssociatedAclPolicies:getApigApiAssociatedAclPolicies", args, GetApigApiAssociatedAclPoliciesResultOutput{}, options).(GetApigApiAssociatedAclPoliciesResultOutput), nil
		}).(GetApigApiAssociatedAclPoliciesResultOutput)
}

// A collection of arguments for invoking getApigApiAssociatedAclPolicies.
type GetApigApiAssociatedAclPoliciesOutputArgs struct {
	ApiId      pulumi.StringInput    `pulumi:"apiId"`
	EntityType pulumi.StringPtrInput `pulumi:"entityType"`
	EnvId      pulumi.StringPtrInput `pulumi:"envId"`
	EnvName    pulumi.StringPtrInput `pulumi:"envName"`
	InstanceId pulumi.StringInput    `pulumi:"instanceId"`
	Name       pulumi.StringPtrInput `pulumi:"name"`
	PolicyId   pulumi.StringPtrInput `pulumi:"policyId"`
	Region     pulumi.StringPtrInput `pulumi:"region"`
	Type       pulumi.StringPtrInput `pulumi:"type"`
}

func (GetApigApiAssociatedAclPoliciesOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetApigApiAssociatedAclPoliciesArgs)(nil)).Elem()
}

// A collection of values returned by getApigApiAssociatedAclPolicies.
type GetApigApiAssociatedAclPoliciesResultOutput struct{ *pulumi.OutputState }

func (GetApigApiAssociatedAclPoliciesResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetApigApiAssociatedAclPoliciesResult)(nil)).Elem()
}

func (o GetApigApiAssociatedAclPoliciesResultOutput) ToGetApigApiAssociatedAclPoliciesResultOutput() GetApigApiAssociatedAclPoliciesResultOutput {
	return o
}

func (o GetApigApiAssociatedAclPoliciesResultOutput) ToGetApigApiAssociatedAclPoliciesResultOutputWithContext(ctx context.Context) GetApigApiAssociatedAclPoliciesResultOutput {
	return o
}

func (o GetApigApiAssociatedAclPoliciesResultOutput) ApiId() pulumi.StringOutput {
	return o.ApplyT(func(v GetApigApiAssociatedAclPoliciesResult) string { return v.ApiId }).(pulumi.StringOutput)
}

func (o GetApigApiAssociatedAclPoliciesResultOutput) EntityType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetApigApiAssociatedAclPoliciesResult) *string { return v.EntityType }).(pulumi.StringPtrOutput)
}

func (o GetApigApiAssociatedAclPoliciesResultOutput) EnvId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetApigApiAssociatedAclPoliciesResult) *string { return v.EnvId }).(pulumi.StringPtrOutput)
}

func (o GetApigApiAssociatedAclPoliciesResultOutput) EnvName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetApigApiAssociatedAclPoliciesResult) *string { return v.EnvName }).(pulumi.StringPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetApigApiAssociatedAclPoliciesResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetApigApiAssociatedAclPoliciesResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetApigApiAssociatedAclPoliciesResultOutput) InstanceId() pulumi.StringOutput {
	return o.ApplyT(func(v GetApigApiAssociatedAclPoliciesResult) string { return v.InstanceId }).(pulumi.StringOutput)
}

func (o GetApigApiAssociatedAclPoliciesResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetApigApiAssociatedAclPoliciesResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o GetApigApiAssociatedAclPoliciesResultOutput) Policies() GetApigApiAssociatedAclPoliciesPolicyArrayOutput {
	return o.ApplyT(func(v GetApigApiAssociatedAclPoliciesResult) []GetApigApiAssociatedAclPoliciesPolicy {
		return v.Policies
	}).(GetApigApiAssociatedAclPoliciesPolicyArrayOutput)
}

func (o GetApigApiAssociatedAclPoliciesResultOutput) PolicyId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetApigApiAssociatedAclPoliciesResult) *string { return v.PolicyId }).(pulumi.StringPtrOutput)
}

func (o GetApigApiAssociatedAclPoliciesResultOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v GetApigApiAssociatedAclPoliciesResult) string { return v.Region }).(pulumi.StringOutput)
}

func (o GetApigApiAssociatedAclPoliciesResultOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetApigApiAssociatedAclPoliciesResult) *string { return v.Type }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetApigApiAssociatedAclPoliciesResultOutput{})
}
