// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package sbercloud

import (
	"context"
	"reflect"

	"github.com/masikrus/pulumi-sbercloud/sdk/go/sbercloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type NetworkingSecgroup struct {
	pulumi.CustomResourceState

	CreatedAt           pulumi.StringOutput                   `pulumi:"createdAt"`
	DeleteDefaultRules  pulumi.BoolPtrOutput                  `pulumi:"deleteDefaultRules"`
	Description         pulumi.StringPtrOutput                `pulumi:"description"`
	EnterpriseProjectId pulumi.StringOutput                   `pulumi:"enterpriseProjectId"`
	Name                pulumi.StringOutput                   `pulumi:"name"`
	Region              pulumi.StringOutput                   `pulumi:"region"`
	Rules               NetworkingSecgroupRuleTypeArrayOutput `pulumi:"rules"`
	Tags                pulumi.StringMapOutput                `pulumi:"tags"`
	UpdatedAt           pulumi.StringOutput                   `pulumi:"updatedAt"`
}

// NewNetworkingSecgroup registers a new resource with the given unique name, arguments, and options.
func NewNetworkingSecgroup(ctx *pulumi.Context,
	name string, args *NetworkingSecgroupArgs, opts ...pulumi.ResourceOption) (*NetworkingSecgroup, error) {
	if args == nil {
		args = &NetworkingSecgroupArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource NetworkingSecgroup
	err := ctx.RegisterResource("sbercloud:index/networkingSecgroup:NetworkingSecgroup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetNetworkingSecgroup gets an existing NetworkingSecgroup resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetNetworkingSecgroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *NetworkingSecgroupState, opts ...pulumi.ResourceOption) (*NetworkingSecgroup, error) {
	var resource NetworkingSecgroup
	err := ctx.ReadResource("sbercloud:index/networkingSecgroup:NetworkingSecgroup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering NetworkingSecgroup resources.
type networkingSecgroupState struct {
	CreatedAt           *string                      `pulumi:"createdAt"`
	DeleteDefaultRules  *bool                        `pulumi:"deleteDefaultRules"`
	Description         *string                      `pulumi:"description"`
	EnterpriseProjectId *string                      `pulumi:"enterpriseProjectId"`
	Name                *string                      `pulumi:"name"`
	Region              *string                      `pulumi:"region"`
	Rules               []NetworkingSecgroupRuleType `pulumi:"rules"`
	Tags                map[string]string            `pulumi:"tags"`
	UpdatedAt           *string                      `pulumi:"updatedAt"`
}

type NetworkingSecgroupState struct {
	CreatedAt           pulumi.StringPtrInput
	DeleteDefaultRules  pulumi.BoolPtrInput
	Description         pulumi.StringPtrInput
	EnterpriseProjectId pulumi.StringPtrInput
	Name                pulumi.StringPtrInput
	Region              pulumi.StringPtrInput
	Rules               NetworkingSecgroupRuleTypeArrayInput
	Tags                pulumi.StringMapInput
	UpdatedAt           pulumi.StringPtrInput
}

func (NetworkingSecgroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*networkingSecgroupState)(nil)).Elem()
}

type networkingSecgroupArgs struct {
	DeleteDefaultRules  *bool             `pulumi:"deleteDefaultRules"`
	Description         *string           `pulumi:"description"`
	EnterpriseProjectId *string           `pulumi:"enterpriseProjectId"`
	Name                *string           `pulumi:"name"`
	Region              *string           `pulumi:"region"`
	Tags                map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a NetworkingSecgroup resource.
type NetworkingSecgroupArgs struct {
	DeleteDefaultRules  pulumi.BoolPtrInput
	Description         pulumi.StringPtrInput
	EnterpriseProjectId pulumi.StringPtrInput
	Name                pulumi.StringPtrInput
	Region              pulumi.StringPtrInput
	Tags                pulumi.StringMapInput
}

func (NetworkingSecgroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*networkingSecgroupArgs)(nil)).Elem()
}

type NetworkingSecgroupInput interface {
	pulumi.Input

	ToNetworkingSecgroupOutput() NetworkingSecgroupOutput
	ToNetworkingSecgroupOutputWithContext(ctx context.Context) NetworkingSecgroupOutput
}

func (*NetworkingSecgroup) ElementType() reflect.Type {
	return reflect.TypeOf((**NetworkingSecgroup)(nil)).Elem()
}

func (i *NetworkingSecgroup) ToNetworkingSecgroupOutput() NetworkingSecgroupOutput {
	return i.ToNetworkingSecgroupOutputWithContext(context.Background())
}

func (i *NetworkingSecgroup) ToNetworkingSecgroupOutputWithContext(ctx context.Context) NetworkingSecgroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkingSecgroupOutput)
}

// NetworkingSecgroupArrayInput is an input type that accepts NetworkingSecgroupArray and NetworkingSecgroupArrayOutput values.
// You can construct a concrete instance of `NetworkingSecgroupArrayInput` via:
//
//	NetworkingSecgroupArray{ NetworkingSecgroupArgs{...} }
type NetworkingSecgroupArrayInput interface {
	pulumi.Input

	ToNetworkingSecgroupArrayOutput() NetworkingSecgroupArrayOutput
	ToNetworkingSecgroupArrayOutputWithContext(context.Context) NetworkingSecgroupArrayOutput
}

type NetworkingSecgroupArray []NetworkingSecgroupInput

func (NetworkingSecgroupArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*NetworkingSecgroup)(nil)).Elem()
}

func (i NetworkingSecgroupArray) ToNetworkingSecgroupArrayOutput() NetworkingSecgroupArrayOutput {
	return i.ToNetworkingSecgroupArrayOutputWithContext(context.Background())
}

func (i NetworkingSecgroupArray) ToNetworkingSecgroupArrayOutputWithContext(ctx context.Context) NetworkingSecgroupArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkingSecgroupArrayOutput)
}

// NetworkingSecgroupMapInput is an input type that accepts NetworkingSecgroupMap and NetworkingSecgroupMapOutput values.
// You can construct a concrete instance of `NetworkingSecgroupMapInput` via:
//
//	NetworkingSecgroupMap{ "key": NetworkingSecgroupArgs{...} }
type NetworkingSecgroupMapInput interface {
	pulumi.Input

	ToNetworkingSecgroupMapOutput() NetworkingSecgroupMapOutput
	ToNetworkingSecgroupMapOutputWithContext(context.Context) NetworkingSecgroupMapOutput
}

type NetworkingSecgroupMap map[string]NetworkingSecgroupInput

func (NetworkingSecgroupMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*NetworkingSecgroup)(nil)).Elem()
}

func (i NetworkingSecgroupMap) ToNetworkingSecgroupMapOutput() NetworkingSecgroupMapOutput {
	return i.ToNetworkingSecgroupMapOutputWithContext(context.Background())
}

func (i NetworkingSecgroupMap) ToNetworkingSecgroupMapOutputWithContext(ctx context.Context) NetworkingSecgroupMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkingSecgroupMapOutput)
}

type NetworkingSecgroupOutput struct{ *pulumi.OutputState }

func (NetworkingSecgroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NetworkingSecgroup)(nil)).Elem()
}

func (o NetworkingSecgroupOutput) ToNetworkingSecgroupOutput() NetworkingSecgroupOutput {
	return o
}

func (o NetworkingSecgroupOutput) ToNetworkingSecgroupOutputWithContext(ctx context.Context) NetworkingSecgroupOutput {
	return o
}

func (o NetworkingSecgroupOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkingSecgroup) pulumi.StringOutput { return v.CreatedAt }).(pulumi.StringOutput)
}

func (o NetworkingSecgroupOutput) DeleteDefaultRules() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *NetworkingSecgroup) pulumi.BoolPtrOutput { return v.DeleteDefaultRules }).(pulumi.BoolPtrOutput)
}

func (o NetworkingSecgroupOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NetworkingSecgroup) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

func (o NetworkingSecgroupOutput) EnterpriseProjectId() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkingSecgroup) pulumi.StringOutput { return v.EnterpriseProjectId }).(pulumi.StringOutput)
}

func (o NetworkingSecgroupOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkingSecgroup) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o NetworkingSecgroupOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkingSecgroup) pulumi.StringOutput { return v.Region }).(pulumi.StringOutput)
}

func (o NetworkingSecgroupOutput) Rules() NetworkingSecgroupRuleTypeArrayOutput {
	return o.ApplyT(func(v *NetworkingSecgroup) NetworkingSecgroupRuleTypeArrayOutput { return v.Rules }).(NetworkingSecgroupRuleTypeArrayOutput)
}

func (o NetworkingSecgroupOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *NetworkingSecgroup) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o NetworkingSecgroupOutput) UpdatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkingSecgroup) pulumi.StringOutput { return v.UpdatedAt }).(pulumi.StringOutput)
}

type NetworkingSecgroupArrayOutput struct{ *pulumi.OutputState }

func (NetworkingSecgroupArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*NetworkingSecgroup)(nil)).Elem()
}

func (o NetworkingSecgroupArrayOutput) ToNetworkingSecgroupArrayOutput() NetworkingSecgroupArrayOutput {
	return o
}

func (o NetworkingSecgroupArrayOutput) ToNetworkingSecgroupArrayOutputWithContext(ctx context.Context) NetworkingSecgroupArrayOutput {
	return o
}

func (o NetworkingSecgroupArrayOutput) Index(i pulumi.IntInput) NetworkingSecgroupOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *NetworkingSecgroup {
		return vs[0].([]*NetworkingSecgroup)[vs[1].(int)]
	}).(NetworkingSecgroupOutput)
}

type NetworkingSecgroupMapOutput struct{ *pulumi.OutputState }

func (NetworkingSecgroupMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*NetworkingSecgroup)(nil)).Elem()
}

func (o NetworkingSecgroupMapOutput) ToNetworkingSecgroupMapOutput() NetworkingSecgroupMapOutput {
	return o
}

func (o NetworkingSecgroupMapOutput) ToNetworkingSecgroupMapOutputWithContext(ctx context.Context) NetworkingSecgroupMapOutput {
	return o
}

func (o NetworkingSecgroupMapOutput) MapIndex(k pulumi.StringInput) NetworkingSecgroupOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *NetworkingSecgroup {
		return vs[0].(map[string]*NetworkingSecgroup)[vs[1].(string)]
	}).(NetworkingSecgroupOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*NetworkingSecgroupInput)(nil)).Elem(), &NetworkingSecgroup{})
	pulumi.RegisterInputType(reflect.TypeOf((*NetworkingSecgroupArrayInput)(nil)).Elem(), NetworkingSecgroupArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*NetworkingSecgroupMapInput)(nil)).Elem(), NetworkingSecgroupMap{})
	pulumi.RegisterOutputType(NetworkingSecgroupOutput{})
	pulumi.RegisterOutputType(NetworkingSecgroupArrayOutput{})
	pulumi.RegisterOutputType(NetworkingSecgroupMapOutput{})
}
