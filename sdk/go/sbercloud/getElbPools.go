// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package sbercloud

import (
	"context"
	"reflect"

	"github.com/masikrus/pulumi-sbercloud/sdk/go/sbercloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetElbPools(ctx *pulumi.Context, args *GetElbPoolsArgs, opts ...pulumi.InvokeOption) (*GetElbPoolsResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetElbPoolsResult
	err := ctx.Invoke("sbercloud:index/getElbPools:getElbPools", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getElbPools.
type GetElbPoolsArgs struct {
	Description      *string `pulumi:"description"`
	HealthmonitorId  *string `pulumi:"healthmonitorId"`
	LbMethod         *string `pulumi:"lbMethod"`
	ListenerId       *string `pulumi:"listenerId"`
	LoadbalancerId   *string `pulumi:"loadbalancerId"`
	Name             *string `pulumi:"name"`
	PoolId           *string `pulumi:"poolId"`
	ProtectionStatus *string `pulumi:"protectionStatus"`
	Protocol         *string `pulumi:"protocol"`
	Region           *string `pulumi:"region"`
	Type             *string `pulumi:"type"`
	VpcId            *string `pulumi:"vpcId"`
}

// A collection of values returned by getElbPools.
type GetElbPoolsResult struct {
	Description     *string `pulumi:"description"`
	HealthmonitorId *string `pulumi:"healthmonitorId"`
	// The provider-assigned unique ID for this managed resource.
	Id               string            `pulumi:"id"`
	LbMethod         *string           `pulumi:"lbMethod"`
	ListenerId       *string           `pulumi:"listenerId"`
	LoadbalancerId   *string           `pulumi:"loadbalancerId"`
	Name             *string           `pulumi:"name"`
	PoolId           *string           `pulumi:"poolId"`
	Pools            []GetElbPoolsPool `pulumi:"pools"`
	ProtectionStatus *string           `pulumi:"protectionStatus"`
	Protocol         *string           `pulumi:"protocol"`
	Region           string            `pulumi:"region"`
	Type             *string           `pulumi:"type"`
	VpcId            *string           `pulumi:"vpcId"`
}

func GetElbPoolsOutput(ctx *pulumi.Context, args GetElbPoolsOutputArgs, opts ...pulumi.InvokeOption) GetElbPoolsResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (GetElbPoolsResultOutput, error) {
			args := v.(GetElbPoolsArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("sbercloud:index/getElbPools:getElbPools", args, GetElbPoolsResultOutput{}, options).(GetElbPoolsResultOutput), nil
		}).(GetElbPoolsResultOutput)
}

// A collection of arguments for invoking getElbPools.
type GetElbPoolsOutputArgs struct {
	Description      pulumi.StringPtrInput `pulumi:"description"`
	HealthmonitorId  pulumi.StringPtrInput `pulumi:"healthmonitorId"`
	LbMethod         pulumi.StringPtrInput `pulumi:"lbMethod"`
	ListenerId       pulumi.StringPtrInput `pulumi:"listenerId"`
	LoadbalancerId   pulumi.StringPtrInput `pulumi:"loadbalancerId"`
	Name             pulumi.StringPtrInput `pulumi:"name"`
	PoolId           pulumi.StringPtrInput `pulumi:"poolId"`
	ProtectionStatus pulumi.StringPtrInput `pulumi:"protectionStatus"`
	Protocol         pulumi.StringPtrInput `pulumi:"protocol"`
	Region           pulumi.StringPtrInput `pulumi:"region"`
	Type             pulumi.StringPtrInput `pulumi:"type"`
	VpcId            pulumi.StringPtrInput `pulumi:"vpcId"`
}

func (GetElbPoolsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetElbPoolsArgs)(nil)).Elem()
}

// A collection of values returned by getElbPools.
type GetElbPoolsResultOutput struct{ *pulumi.OutputState }

func (GetElbPoolsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetElbPoolsResult)(nil)).Elem()
}

func (o GetElbPoolsResultOutput) ToGetElbPoolsResultOutput() GetElbPoolsResultOutput {
	return o
}

func (o GetElbPoolsResultOutput) ToGetElbPoolsResultOutputWithContext(ctx context.Context) GetElbPoolsResultOutput {
	return o
}

func (o GetElbPoolsResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetElbPoolsResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o GetElbPoolsResultOutput) HealthmonitorId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetElbPoolsResult) *string { return v.HealthmonitorId }).(pulumi.StringPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetElbPoolsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetElbPoolsResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetElbPoolsResultOutput) LbMethod() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetElbPoolsResult) *string { return v.LbMethod }).(pulumi.StringPtrOutput)
}

func (o GetElbPoolsResultOutput) ListenerId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetElbPoolsResult) *string { return v.ListenerId }).(pulumi.StringPtrOutput)
}

func (o GetElbPoolsResultOutput) LoadbalancerId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetElbPoolsResult) *string { return v.LoadbalancerId }).(pulumi.StringPtrOutput)
}

func (o GetElbPoolsResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetElbPoolsResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o GetElbPoolsResultOutput) PoolId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetElbPoolsResult) *string { return v.PoolId }).(pulumi.StringPtrOutput)
}

func (o GetElbPoolsResultOutput) Pools() GetElbPoolsPoolArrayOutput {
	return o.ApplyT(func(v GetElbPoolsResult) []GetElbPoolsPool { return v.Pools }).(GetElbPoolsPoolArrayOutput)
}

func (o GetElbPoolsResultOutput) ProtectionStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetElbPoolsResult) *string { return v.ProtectionStatus }).(pulumi.StringPtrOutput)
}

func (o GetElbPoolsResultOutput) Protocol() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetElbPoolsResult) *string { return v.Protocol }).(pulumi.StringPtrOutput)
}

func (o GetElbPoolsResultOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v GetElbPoolsResult) string { return v.Region }).(pulumi.StringOutput)
}

func (o GetElbPoolsResultOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetElbPoolsResult) *string { return v.Type }).(pulumi.StringPtrOutput)
}

func (o GetElbPoolsResultOutput) VpcId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetElbPoolsResult) *string { return v.VpcId }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetElbPoolsResultOutput{})
}
