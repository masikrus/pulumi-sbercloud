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

type CceNamespace struct {
	pulumi.CustomResourceState

	Annotations       pulumi.StringMapOutput `pulumi:"annotations"`
	ClusterId         pulumi.StringOutput    `pulumi:"clusterId"`
	CreationTimestamp pulumi.StringOutput    `pulumi:"creationTimestamp"`
	Labels            pulumi.StringMapOutput `pulumi:"labels"`
	Name              pulumi.StringOutput    `pulumi:"name"`
	Prefix            pulumi.StringPtrOutput `pulumi:"prefix"`
	Region            pulumi.StringOutput    `pulumi:"region"`
	Status            pulumi.StringOutput    `pulumi:"status"`
}

// NewCceNamespace registers a new resource with the given unique name, arguments, and options.
func NewCceNamespace(ctx *pulumi.Context,
	name string, args *CceNamespaceArgs, opts ...pulumi.ResourceOption) (*CceNamespace, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ClusterId == nil {
		return nil, errors.New("invalid value for required argument 'ClusterId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource CceNamespace
	err := ctx.RegisterResource("sbercloud:index/cceNamespace:CceNamespace", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCceNamespace gets an existing CceNamespace resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCceNamespace(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CceNamespaceState, opts ...pulumi.ResourceOption) (*CceNamespace, error) {
	var resource CceNamespace
	err := ctx.ReadResource("sbercloud:index/cceNamespace:CceNamespace", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering CceNamespace resources.
type cceNamespaceState struct {
	Annotations       map[string]string `pulumi:"annotations"`
	ClusterId         *string           `pulumi:"clusterId"`
	CreationTimestamp *string           `pulumi:"creationTimestamp"`
	Labels            map[string]string `pulumi:"labels"`
	Name              *string           `pulumi:"name"`
	Prefix            *string           `pulumi:"prefix"`
	Region            *string           `pulumi:"region"`
	Status            *string           `pulumi:"status"`
}

type CceNamespaceState struct {
	Annotations       pulumi.StringMapInput
	ClusterId         pulumi.StringPtrInput
	CreationTimestamp pulumi.StringPtrInput
	Labels            pulumi.StringMapInput
	Name              pulumi.StringPtrInput
	Prefix            pulumi.StringPtrInput
	Region            pulumi.StringPtrInput
	Status            pulumi.StringPtrInput
}

func (CceNamespaceState) ElementType() reflect.Type {
	return reflect.TypeOf((*cceNamespaceState)(nil)).Elem()
}

type cceNamespaceArgs struct {
	Annotations map[string]string `pulumi:"annotations"`
	ClusterId   string            `pulumi:"clusterId"`
	Labels      map[string]string `pulumi:"labels"`
	Name        *string           `pulumi:"name"`
	Prefix      *string           `pulumi:"prefix"`
	Region      *string           `pulumi:"region"`
}

// The set of arguments for constructing a CceNamespace resource.
type CceNamespaceArgs struct {
	Annotations pulumi.StringMapInput
	ClusterId   pulumi.StringInput
	Labels      pulumi.StringMapInput
	Name        pulumi.StringPtrInput
	Prefix      pulumi.StringPtrInput
	Region      pulumi.StringPtrInput
}

func (CceNamespaceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*cceNamespaceArgs)(nil)).Elem()
}

type CceNamespaceInput interface {
	pulumi.Input

	ToCceNamespaceOutput() CceNamespaceOutput
	ToCceNamespaceOutputWithContext(ctx context.Context) CceNamespaceOutput
}

func (*CceNamespace) ElementType() reflect.Type {
	return reflect.TypeOf((**CceNamespace)(nil)).Elem()
}

func (i *CceNamespace) ToCceNamespaceOutput() CceNamespaceOutput {
	return i.ToCceNamespaceOutputWithContext(context.Background())
}

func (i *CceNamespace) ToCceNamespaceOutputWithContext(ctx context.Context) CceNamespaceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CceNamespaceOutput)
}

// CceNamespaceArrayInput is an input type that accepts CceNamespaceArray and CceNamespaceArrayOutput values.
// You can construct a concrete instance of `CceNamespaceArrayInput` via:
//
//	CceNamespaceArray{ CceNamespaceArgs{...} }
type CceNamespaceArrayInput interface {
	pulumi.Input

	ToCceNamespaceArrayOutput() CceNamespaceArrayOutput
	ToCceNamespaceArrayOutputWithContext(context.Context) CceNamespaceArrayOutput
}

type CceNamespaceArray []CceNamespaceInput

func (CceNamespaceArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*CceNamespace)(nil)).Elem()
}

func (i CceNamespaceArray) ToCceNamespaceArrayOutput() CceNamespaceArrayOutput {
	return i.ToCceNamespaceArrayOutputWithContext(context.Background())
}

func (i CceNamespaceArray) ToCceNamespaceArrayOutputWithContext(ctx context.Context) CceNamespaceArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CceNamespaceArrayOutput)
}

// CceNamespaceMapInput is an input type that accepts CceNamespaceMap and CceNamespaceMapOutput values.
// You can construct a concrete instance of `CceNamespaceMapInput` via:
//
//	CceNamespaceMap{ "key": CceNamespaceArgs{...} }
type CceNamespaceMapInput interface {
	pulumi.Input

	ToCceNamespaceMapOutput() CceNamespaceMapOutput
	ToCceNamespaceMapOutputWithContext(context.Context) CceNamespaceMapOutput
}

type CceNamespaceMap map[string]CceNamespaceInput

func (CceNamespaceMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*CceNamespace)(nil)).Elem()
}

func (i CceNamespaceMap) ToCceNamespaceMapOutput() CceNamespaceMapOutput {
	return i.ToCceNamespaceMapOutputWithContext(context.Background())
}

func (i CceNamespaceMap) ToCceNamespaceMapOutputWithContext(ctx context.Context) CceNamespaceMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CceNamespaceMapOutput)
}

type CceNamespaceOutput struct{ *pulumi.OutputState }

func (CceNamespaceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CceNamespace)(nil)).Elem()
}

func (o CceNamespaceOutput) ToCceNamespaceOutput() CceNamespaceOutput {
	return o
}

func (o CceNamespaceOutput) ToCceNamespaceOutputWithContext(ctx context.Context) CceNamespaceOutput {
	return o
}

func (o CceNamespaceOutput) Annotations() pulumi.StringMapOutput {
	return o.ApplyT(func(v *CceNamespace) pulumi.StringMapOutput { return v.Annotations }).(pulumi.StringMapOutput)
}

func (o CceNamespaceOutput) ClusterId() pulumi.StringOutput {
	return o.ApplyT(func(v *CceNamespace) pulumi.StringOutput { return v.ClusterId }).(pulumi.StringOutput)
}

func (o CceNamespaceOutput) CreationTimestamp() pulumi.StringOutput {
	return o.ApplyT(func(v *CceNamespace) pulumi.StringOutput { return v.CreationTimestamp }).(pulumi.StringOutput)
}

func (o CceNamespaceOutput) Labels() pulumi.StringMapOutput {
	return o.ApplyT(func(v *CceNamespace) pulumi.StringMapOutput { return v.Labels }).(pulumi.StringMapOutput)
}

func (o CceNamespaceOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *CceNamespace) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o CceNamespaceOutput) Prefix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CceNamespace) pulumi.StringPtrOutput { return v.Prefix }).(pulumi.StringPtrOutput)
}

func (o CceNamespaceOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v *CceNamespace) pulumi.StringOutput { return v.Region }).(pulumi.StringOutput)
}

func (o CceNamespaceOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *CceNamespace) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

type CceNamespaceArrayOutput struct{ *pulumi.OutputState }

func (CceNamespaceArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*CceNamespace)(nil)).Elem()
}

func (o CceNamespaceArrayOutput) ToCceNamespaceArrayOutput() CceNamespaceArrayOutput {
	return o
}

func (o CceNamespaceArrayOutput) ToCceNamespaceArrayOutputWithContext(ctx context.Context) CceNamespaceArrayOutput {
	return o
}

func (o CceNamespaceArrayOutput) Index(i pulumi.IntInput) CceNamespaceOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *CceNamespace {
		return vs[0].([]*CceNamespace)[vs[1].(int)]
	}).(CceNamespaceOutput)
}

type CceNamespaceMapOutput struct{ *pulumi.OutputState }

func (CceNamespaceMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*CceNamespace)(nil)).Elem()
}

func (o CceNamespaceMapOutput) ToCceNamespaceMapOutput() CceNamespaceMapOutput {
	return o
}

func (o CceNamespaceMapOutput) ToCceNamespaceMapOutputWithContext(ctx context.Context) CceNamespaceMapOutput {
	return o
}

func (o CceNamespaceMapOutput) MapIndex(k pulumi.StringInput) CceNamespaceOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *CceNamespace {
		return vs[0].(map[string]*CceNamespace)[vs[1].(string)]
	}).(CceNamespaceOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*CceNamespaceInput)(nil)).Elem(), &CceNamespace{})
	pulumi.RegisterInputType(reflect.TypeOf((*CceNamespaceArrayInput)(nil)).Elem(), CceNamespaceArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*CceNamespaceMapInput)(nil)).Elem(), CceNamespaceMap{})
	pulumi.RegisterOutputType(CceNamespaceOutput{})
	pulumi.RegisterOutputType(CceNamespaceArrayOutput{})
	pulumi.RegisterOutputType(CceNamespaceMapOutput{})
}
