// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package sbercloud

import (
	"context"
	"reflect"

	"errors"
	"github.com/masikrus/pulumi-sbercloud/sdk/go/sbercloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type IdentityProvider struct {
	pulumi.CustomResourceState

	AccessConfig    IdentityProviderAccessConfigPtrOutput     `pulumi:"accessConfig"`
	ConversionRules IdentityProviderConversionRuleArrayOutput `pulumi:"conversionRules"`
	Description     pulumi.StringPtrOutput                    `pulumi:"description"`
	LoginLink       pulumi.StringOutput                       `pulumi:"loginLink"`
	Metadata        pulumi.StringPtrOutput                    `pulumi:"metadata"`
	Name            pulumi.StringOutput                       `pulumi:"name"`
	Protocol        pulumi.StringOutput                       `pulumi:"protocol"`
	SsoType         pulumi.StringOutput                       `pulumi:"ssoType"`
	Status          pulumi.BoolPtrOutput                      `pulumi:"status"`
}

// NewIdentityProvider registers a new resource with the given unique name, arguments, and options.
func NewIdentityProvider(ctx *pulumi.Context,
	name string, args *IdentityProviderArgs, opts ...pulumi.ResourceOption) (*IdentityProvider, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Protocol == nil {
		return nil, errors.New("invalid value for required argument 'Protocol'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource IdentityProvider
	err := ctx.RegisterResource("sbercloud:index/identityProvider:IdentityProvider", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetIdentityProvider gets an existing IdentityProvider resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetIdentityProvider(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *IdentityProviderState, opts ...pulumi.ResourceOption) (*IdentityProvider, error) {
	var resource IdentityProvider
	err := ctx.ReadResource("sbercloud:index/identityProvider:IdentityProvider", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering IdentityProvider resources.
type identityProviderState struct {
	AccessConfig    *IdentityProviderAccessConfig    `pulumi:"accessConfig"`
	ConversionRules []IdentityProviderConversionRule `pulumi:"conversionRules"`
	Description     *string                          `pulumi:"description"`
	LoginLink       *string                          `pulumi:"loginLink"`
	Metadata        *string                          `pulumi:"metadata"`
	Name            *string                          `pulumi:"name"`
	Protocol        *string                          `pulumi:"protocol"`
	SsoType         *string                          `pulumi:"ssoType"`
	Status          *bool                            `pulumi:"status"`
}

type IdentityProviderState struct {
	AccessConfig    IdentityProviderAccessConfigPtrInput
	ConversionRules IdentityProviderConversionRuleArrayInput
	Description     pulumi.StringPtrInput
	LoginLink       pulumi.StringPtrInput
	Metadata        pulumi.StringPtrInput
	Name            pulumi.StringPtrInput
	Protocol        pulumi.StringPtrInput
	SsoType         pulumi.StringPtrInput
	Status          pulumi.BoolPtrInput
}

func (IdentityProviderState) ElementType() reflect.Type {
	return reflect.TypeOf((*identityProviderState)(nil)).Elem()
}

type identityProviderArgs struct {
	AccessConfig *IdentityProviderAccessConfig `pulumi:"accessConfig"`
	Description  *string                       `pulumi:"description"`
	Metadata     *string                       `pulumi:"metadata"`
	Name         *string                       `pulumi:"name"`
	Protocol     string                        `pulumi:"protocol"`
	SsoType      *string                       `pulumi:"ssoType"`
	Status       *bool                         `pulumi:"status"`
}

// The set of arguments for constructing a IdentityProvider resource.
type IdentityProviderArgs struct {
	AccessConfig IdentityProviderAccessConfigPtrInput
	Description  pulumi.StringPtrInput
	Metadata     pulumi.StringPtrInput
	Name         pulumi.StringPtrInput
	Protocol     pulumi.StringInput
	SsoType      pulumi.StringPtrInput
	Status       pulumi.BoolPtrInput
}

func (IdentityProviderArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*identityProviderArgs)(nil)).Elem()
}

type IdentityProviderInput interface {
	pulumi.Input

	ToIdentityProviderOutput() IdentityProviderOutput
	ToIdentityProviderOutputWithContext(ctx context.Context) IdentityProviderOutput
}

func (*IdentityProvider) ElementType() reflect.Type {
	return reflect.TypeOf((**IdentityProvider)(nil)).Elem()
}

func (i *IdentityProvider) ToIdentityProviderOutput() IdentityProviderOutput {
	return i.ToIdentityProviderOutputWithContext(context.Background())
}

func (i *IdentityProvider) ToIdentityProviderOutputWithContext(ctx context.Context) IdentityProviderOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IdentityProviderOutput)
}

// IdentityProviderArrayInput is an input type that accepts IdentityProviderArray and IdentityProviderArrayOutput values.
// You can construct a concrete instance of `IdentityProviderArrayInput` via:
//
//	IdentityProviderArray{ IdentityProviderArgs{...} }
type IdentityProviderArrayInput interface {
	pulumi.Input

	ToIdentityProviderArrayOutput() IdentityProviderArrayOutput
	ToIdentityProviderArrayOutputWithContext(context.Context) IdentityProviderArrayOutput
}

type IdentityProviderArray []IdentityProviderInput

func (IdentityProviderArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*IdentityProvider)(nil)).Elem()
}

func (i IdentityProviderArray) ToIdentityProviderArrayOutput() IdentityProviderArrayOutput {
	return i.ToIdentityProviderArrayOutputWithContext(context.Background())
}

func (i IdentityProviderArray) ToIdentityProviderArrayOutputWithContext(ctx context.Context) IdentityProviderArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IdentityProviderArrayOutput)
}

// IdentityProviderMapInput is an input type that accepts IdentityProviderMap and IdentityProviderMapOutput values.
// You can construct a concrete instance of `IdentityProviderMapInput` via:
//
//	IdentityProviderMap{ "key": IdentityProviderArgs{...} }
type IdentityProviderMapInput interface {
	pulumi.Input

	ToIdentityProviderMapOutput() IdentityProviderMapOutput
	ToIdentityProviderMapOutputWithContext(context.Context) IdentityProviderMapOutput
}

type IdentityProviderMap map[string]IdentityProviderInput

func (IdentityProviderMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*IdentityProvider)(nil)).Elem()
}

func (i IdentityProviderMap) ToIdentityProviderMapOutput() IdentityProviderMapOutput {
	return i.ToIdentityProviderMapOutputWithContext(context.Background())
}

func (i IdentityProviderMap) ToIdentityProviderMapOutputWithContext(ctx context.Context) IdentityProviderMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IdentityProviderMapOutput)
}

type IdentityProviderOutput struct{ *pulumi.OutputState }

func (IdentityProviderOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IdentityProvider)(nil)).Elem()
}

func (o IdentityProviderOutput) ToIdentityProviderOutput() IdentityProviderOutput {
	return o
}

func (o IdentityProviderOutput) ToIdentityProviderOutputWithContext(ctx context.Context) IdentityProviderOutput {
	return o
}

func (o IdentityProviderOutput) AccessConfig() IdentityProviderAccessConfigPtrOutput {
	return o.ApplyT(func(v *IdentityProvider) IdentityProviderAccessConfigPtrOutput { return v.AccessConfig }).(IdentityProviderAccessConfigPtrOutput)
}

func (o IdentityProviderOutput) ConversionRules() IdentityProviderConversionRuleArrayOutput {
	return o.ApplyT(func(v *IdentityProvider) IdentityProviderConversionRuleArrayOutput { return v.ConversionRules }).(IdentityProviderConversionRuleArrayOutput)
}

func (o IdentityProviderOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IdentityProvider) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

func (o IdentityProviderOutput) LoginLink() pulumi.StringOutput {
	return o.ApplyT(func(v *IdentityProvider) pulumi.StringOutput { return v.LoginLink }).(pulumi.StringOutput)
}

func (o IdentityProviderOutput) Metadata() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IdentityProvider) pulumi.StringPtrOutput { return v.Metadata }).(pulumi.StringPtrOutput)
}

func (o IdentityProviderOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *IdentityProvider) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o IdentityProviderOutput) Protocol() pulumi.StringOutput {
	return o.ApplyT(func(v *IdentityProvider) pulumi.StringOutput { return v.Protocol }).(pulumi.StringOutput)
}

func (o IdentityProviderOutput) SsoType() pulumi.StringOutput {
	return o.ApplyT(func(v *IdentityProvider) pulumi.StringOutput { return v.SsoType }).(pulumi.StringOutput)
}

func (o IdentityProviderOutput) Status() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *IdentityProvider) pulumi.BoolPtrOutput { return v.Status }).(pulumi.BoolPtrOutput)
}

type IdentityProviderArrayOutput struct{ *pulumi.OutputState }

func (IdentityProviderArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*IdentityProvider)(nil)).Elem()
}

func (o IdentityProviderArrayOutput) ToIdentityProviderArrayOutput() IdentityProviderArrayOutput {
	return o
}

func (o IdentityProviderArrayOutput) ToIdentityProviderArrayOutputWithContext(ctx context.Context) IdentityProviderArrayOutput {
	return o
}

func (o IdentityProviderArrayOutput) Index(i pulumi.IntInput) IdentityProviderOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *IdentityProvider {
		return vs[0].([]*IdentityProvider)[vs[1].(int)]
	}).(IdentityProviderOutput)
}

type IdentityProviderMapOutput struct{ *pulumi.OutputState }

func (IdentityProviderMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*IdentityProvider)(nil)).Elem()
}

func (o IdentityProviderMapOutput) ToIdentityProviderMapOutput() IdentityProviderMapOutput {
	return o
}

func (o IdentityProviderMapOutput) ToIdentityProviderMapOutputWithContext(ctx context.Context) IdentityProviderMapOutput {
	return o
}

func (o IdentityProviderMapOutput) MapIndex(k pulumi.StringInput) IdentityProviderOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *IdentityProvider {
		return vs[0].(map[string]*IdentityProvider)[vs[1].(string)]
	}).(IdentityProviderOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*IdentityProviderInput)(nil)).Elem(), &IdentityProvider{})
	pulumi.RegisterInputType(reflect.TypeOf((*IdentityProviderArrayInput)(nil)).Elem(), IdentityProviderArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*IdentityProviderMapInput)(nil)).Elem(), IdentityProviderMap{})
	pulumi.RegisterOutputType(IdentityProviderOutput{})
	pulumi.RegisterOutputType(IdentityProviderArrayOutput{})
	pulumi.RegisterOutputType(IdentityProviderMapOutput{})
}
