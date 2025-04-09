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

type CcePvc struct {
	pulumi.CustomResourceState

	AccessModes       pulumi.StringArrayOutput `pulumi:"accessModes"`
	Annotations       pulumi.StringMapOutput   `pulumi:"annotations"`
	ClusterId         pulumi.StringOutput      `pulumi:"clusterId"`
	CreationTimestamp pulumi.StringOutput      `pulumi:"creationTimestamp"`
	Labels            pulumi.StringMapOutput   `pulumi:"labels"`
	Name              pulumi.StringOutput      `pulumi:"name"`
	Namespace         pulumi.StringOutput      `pulumi:"namespace"`
	Region            pulumi.StringOutput      `pulumi:"region"`
	Status            pulumi.StringOutput      `pulumi:"status"`
	Storage           pulumi.StringOutput      `pulumi:"storage"`
	StorageClassName  pulumi.StringOutput      `pulumi:"storageClassName"`
}

// NewCcePvc registers a new resource with the given unique name, arguments, and options.
func NewCcePvc(ctx *pulumi.Context,
	name string, args *CcePvcArgs, opts ...pulumi.ResourceOption) (*CcePvc, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AccessModes == nil {
		return nil, errors.New("invalid value for required argument 'AccessModes'")
	}
	if args.ClusterId == nil {
		return nil, errors.New("invalid value for required argument 'ClusterId'")
	}
	if args.Namespace == nil {
		return nil, errors.New("invalid value for required argument 'Namespace'")
	}
	if args.Storage == nil {
		return nil, errors.New("invalid value for required argument 'Storage'")
	}
	if args.StorageClassName == nil {
		return nil, errors.New("invalid value for required argument 'StorageClassName'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource CcePvc
	err := ctx.RegisterResource("sbercloud:index/ccePvc:CcePvc", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCcePvc gets an existing CcePvc resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCcePvc(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CcePvcState, opts ...pulumi.ResourceOption) (*CcePvc, error) {
	var resource CcePvc
	err := ctx.ReadResource("sbercloud:index/ccePvc:CcePvc", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering CcePvc resources.
type ccePvcState struct {
	AccessModes       []string          `pulumi:"accessModes"`
	Annotations       map[string]string `pulumi:"annotations"`
	ClusterId         *string           `pulumi:"clusterId"`
	CreationTimestamp *string           `pulumi:"creationTimestamp"`
	Labels            map[string]string `pulumi:"labels"`
	Name              *string           `pulumi:"name"`
	Namespace         *string           `pulumi:"namespace"`
	Region            *string           `pulumi:"region"`
	Status            *string           `pulumi:"status"`
	Storage           *string           `pulumi:"storage"`
	StorageClassName  *string           `pulumi:"storageClassName"`
}

type CcePvcState struct {
	AccessModes       pulumi.StringArrayInput
	Annotations       pulumi.StringMapInput
	ClusterId         pulumi.StringPtrInput
	CreationTimestamp pulumi.StringPtrInput
	Labels            pulumi.StringMapInput
	Name              pulumi.StringPtrInput
	Namespace         pulumi.StringPtrInput
	Region            pulumi.StringPtrInput
	Status            pulumi.StringPtrInput
	Storage           pulumi.StringPtrInput
	StorageClassName  pulumi.StringPtrInput
}

func (CcePvcState) ElementType() reflect.Type {
	return reflect.TypeOf((*ccePvcState)(nil)).Elem()
}

type ccePvcArgs struct {
	AccessModes      []string          `pulumi:"accessModes"`
	Annotations      map[string]string `pulumi:"annotations"`
	ClusterId        string            `pulumi:"clusterId"`
	Labels           map[string]string `pulumi:"labels"`
	Name             *string           `pulumi:"name"`
	Namespace        string            `pulumi:"namespace"`
	Region           *string           `pulumi:"region"`
	Storage          string            `pulumi:"storage"`
	StorageClassName string            `pulumi:"storageClassName"`
}

// The set of arguments for constructing a CcePvc resource.
type CcePvcArgs struct {
	AccessModes      pulumi.StringArrayInput
	Annotations      pulumi.StringMapInput
	ClusterId        pulumi.StringInput
	Labels           pulumi.StringMapInput
	Name             pulumi.StringPtrInput
	Namespace        pulumi.StringInput
	Region           pulumi.StringPtrInput
	Storage          pulumi.StringInput
	StorageClassName pulumi.StringInput
}

func (CcePvcArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ccePvcArgs)(nil)).Elem()
}

type CcePvcInput interface {
	pulumi.Input

	ToCcePvcOutput() CcePvcOutput
	ToCcePvcOutputWithContext(ctx context.Context) CcePvcOutput
}

func (*CcePvc) ElementType() reflect.Type {
	return reflect.TypeOf((**CcePvc)(nil)).Elem()
}

func (i *CcePvc) ToCcePvcOutput() CcePvcOutput {
	return i.ToCcePvcOutputWithContext(context.Background())
}

func (i *CcePvc) ToCcePvcOutputWithContext(ctx context.Context) CcePvcOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CcePvcOutput)
}

// CcePvcArrayInput is an input type that accepts CcePvcArray and CcePvcArrayOutput values.
// You can construct a concrete instance of `CcePvcArrayInput` via:
//
//	CcePvcArray{ CcePvcArgs{...} }
type CcePvcArrayInput interface {
	pulumi.Input

	ToCcePvcArrayOutput() CcePvcArrayOutput
	ToCcePvcArrayOutputWithContext(context.Context) CcePvcArrayOutput
}

type CcePvcArray []CcePvcInput

func (CcePvcArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*CcePvc)(nil)).Elem()
}

func (i CcePvcArray) ToCcePvcArrayOutput() CcePvcArrayOutput {
	return i.ToCcePvcArrayOutputWithContext(context.Background())
}

func (i CcePvcArray) ToCcePvcArrayOutputWithContext(ctx context.Context) CcePvcArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CcePvcArrayOutput)
}

// CcePvcMapInput is an input type that accepts CcePvcMap and CcePvcMapOutput values.
// You can construct a concrete instance of `CcePvcMapInput` via:
//
//	CcePvcMap{ "key": CcePvcArgs{...} }
type CcePvcMapInput interface {
	pulumi.Input

	ToCcePvcMapOutput() CcePvcMapOutput
	ToCcePvcMapOutputWithContext(context.Context) CcePvcMapOutput
}

type CcePvcMap map[string]CcePvcInput

func (CcePvcMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*CcePvc)(nil)).Elem()
}

func (i CcePvcMap) ToCcePvcMapOutput() CcePvcMapOutput {
	return i.ToCcePvcMapOutputWithContext(context.Background())
}

func (i CcePvcMap) ToCcePvcMapOutputWithContext(ctx context.Context) CcePvcMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CcePvcMapOutput)
}

type CcePvcOutput struct{ *pulumi.OutputState }

func (CcePvcOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CcePvc)(nil)).Elem()
}

func (o CcePvcOutput) ToCcePvcOutput() CcePvcOutput {
	return o
}

func (o CcePvcOutput) ToCcePvcOutputWithContext(ctx context.Context) CcePvcOutput {
	return o
}

func (o CcePvcOutput) AccessModes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *CcePvc) pulumi.StringArrayOutput { return v.AccessModes }).(pulumi.StringArrayOutput)
}

func (o CcePvcOutput) Annotations() pulumi.StringMapOutput {
	return o.ApplyT(func(v *CcePvc) pulumi.StringMapOutput { return v.Annotations }).(pulumi.StringMapOutput)
}

func (o CcePvcOutput) ClusterId() pulumi.StringOutput {
	return o.ApplyT(func(v *CcePvc) pulumi.StringOutput { return v.ClusterId }).(pulumi.StringOutput)
}

func (o CcePvcOutput) CreationTimestamp() pulumi.StringOutput {
	return o.ApplyT(func(v *CcePvc) pulumi.StringOutput { return v.CreationTimestamp }).(pulumi.StringOutput)
}

func (o CcePvcOutput) Labels() pulumi.StringMapOutput {
	return o.ApplyT(func(v *CcePvc) pulumi.StringMapOutput { return v.Labels }).(pulumi.StringMapOutput)
}

func (o CcePvcOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *CcePvc) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o CcePvcOutput) Namespace() pulumi.StringOutput {
	return o.ApplyT(func(v *CcePvc) pulumi.StringOutput { return v.Namespace }).(pulumi.StringOutput)
}

func (o CcePvcOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v *CcePvc) pulumi.StringOutput { return v.Region }).(pulumi.StringOutput)
}

func (o CcePvcOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *CcePvc) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

func (o CcePvcOutput) Storage() pulumi.StringOutput {
	return o.ApplyT(func(v *CcePvc) pulumi.StringOutput { return v.Storage }).(pulumi.StringOutput)
}

func (o CcePvcOutput) StorageClassName() pulumi.StringOutput {
	return o.ApplyT(func(v *CcePvc) pulumi.StringOutput { return v.StorageClassName }).(pulumi.StringOutput)
}

type CcePvcArrayOutput struct{ *pulumi.OutputState }

func (CcePvcArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*CcePvc)(nil)).Elem()
}

func (o CcePvcArrayOutput) ToCcePvcArrayOutput() CcePvcArrayOutput {
	return o
}

func (o CcePvcArrayOutput) ToCcePvcArrayOutputWithContext(ctx context.Context) CcePvcArrayOutput {
	return o
}

func (o CcePvcArrayOutput) Index(i pulumi.IntInput) CcePvcOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *CcePvc {
		return vs[0].([]*CcePvc)[vs[1].(int)]
	}).(CcePvcOutput)
}

type CcePvcMapOutput struct{ *pulumi.OutputState }

func (CcePvcMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*CcePvc)(nil)).Elem()
}

func (o CcePvcMapOutput) ToCcePvcMapOutput() CcePvcMapOutput {
	return o
}

func (o CcePvcMapOutput) ToCcePvcMapOutputWithContext(ctx context.Context) CcePvcMapOutput {
	return o
}

func (o CcePvcMapOutput) MapIndex(k pulumi.StringInput) CcePvcOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *CcePvc {
		return vs[0].(map[string]*CcePvc)[vs[1].(string)]
	}).(CcePvcOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*CcePvcInput)(nil)).Elem(), &CcePvc{})
	pulumi.RegisterInputType(reflect.TypeOf((*CcePvcArrayInput)(nil)).Elem(), CcePvcArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*CcePvcMapInput)(nil)).Elem(), CcePvcMap{})
	pulumi.RegisterOutputType(CcePvcOutput{})
	pulumi.RegisterOutputType(CcePvcArrayOutput{})
	pulumi.RegisterOutputType(CcePvcMapOutput{})
}
