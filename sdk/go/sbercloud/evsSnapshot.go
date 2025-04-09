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

type EvsSnapshot struct {
	pulumi.CustomResourceState

	Description pulumi.StringPtrOutput `pulumi:"description"`
	Force       pulumi.BoolPtrOutput   `pulumi:"force"`
	Metadata    pulumi.StringMapOutput `pulumi:"metadata"`
	Name        pulumi.StringOutput    `pulumi:"name"`
	Region      pulumi.StringOutput    `pulumi:"region"`
	Size        pulumi.IntOutput       `pulumi:"size"`
	Status      pulumi.StringOutput    `pulumi:"status"`
	VolumeId    pulumi.StringOutput    `pulumi:"volumeId"`
}

// NewEvsSnapshot registers a new resource with the given unique name, arguments, and options.
func NewEvsSnapshot(ctx *pulumi.Context,
	name string, args *EvsSnapshotArgs, opts ...pulumi.ResourceOption) (*EvsSnapshot, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.VolumeId == nil {
		return nil, errors.New("invalid value for required argument 'VolumeId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource EvsSnapshot
	err := ctx.RegisterResource("sbercloud:index/evsSnapshot:EvsSnapshot", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetEvsSnapshot gets an existing EvsSnapshot resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetEvsSnapshot(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *EvsSnapshotState, opts ...pulumi.ResourceOption) (*EvsSnapshot, error) {
	var resource EvsSnapshot
	err := ctx.ReadResource("sbercloud:index/evsSnapshot:EvsSnapshot", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering EvsSnapshot resources.
type evsSnapshotState struct {
	Description *string           `pulumi:"description"`
	Force       *bool             `pulumi:"force"`
	Metadata    map[string]string `pulumi:"metadata"`
	Name        *string           `pulumi:"name"`
	Region      *string           `pulumi:"region"`
	Size        *int              `pulumi:"size"`
	Status      *string           `pulumi:"status"`
	VolumeId    *string           `pulumi:"volumeId"`
}

type EvsSnapshotState struct {
	Description pulumi.StringPtrInput
	Force       pulumi.BoolPtrInput
	Metadata    pulumi.StringMapInput
	Name        pulumi.StringPtrInput
	Region      pulumi.StringPtrInput
	Size        pulumi.IntPtrInput
	Status      pulumi.StringPtrInput
	VolumeId    pulumi.StringPtrInput
}

func (EvsSnapshotState) ElementType() reflect.Type {
	return reflect.TypeOf((*evsSnapshotState)(nil)).Elem()
}

type evsSnapshotArgs struct {
	Description *string           `pulumi:"description"`
	Force       *bool             `pulumi:"force"`
	Metadata    map[string]string `pulumi:"metadata"`
	Name        *string           `pulumi:"name"`
	Region      *string           `pulumi:"region"`
	VolumeId    string            `pulumi:"volumeId"`
}

// The set of arguments for constructing a EvsSnapshot resource.
type EvsSnapshotArgs struct {
	Description pulumi.StringPtrInput
	Force       pulumi.BoolPtrInput
	Metadata    pulumi.StringMapInput
	Name        pulumi.StringPtrInput
	Region      pulumi.StringPtrInput
	VolumeId    pulumi.StringInput
}

func (EvsSnapshotArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*evsSnapshotArgs)(nil)).Elem()
}

type EvsSnapshotInput interface {
	pulumi.Input

	ToEvsSnapshotOutput() EvsSnapshotOutput
	ToEvsSnapshotOutputWithContext(ctx context.Context) EvsSnapshotOutput
}

func (*EvsSnapshot) ElementType() reflect.Type {
	return reflect.TypeOf((**EvsSnapshot)(nil)).Elem()
}

func (i *EvsSnapshot) ToEvsSnapshotOutput() EvsSnapshotOutput {
	return i.ToEvsSnapshotOutputWithContext(context.Background())
}

func (i *EvsSnapshot) ToEvsSnapshotOutputWithContext(ctx context.Context) EvsSnapshotOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EvsSnapshotOutput)
}

// EvsSnapshotArrayInput is an input type that accepts EvsSnapshotArray and EvsSnapshotArrayOutput values.
// You can construct a concrete instance of `EvsSnapshotArrayInput` via:
//
//	EvsSnapshotArray{ EvsSnapshotArgs{...} }
type EvsSnapshotArrayInput interface {
	pulumi.Input

	ToEvsSnapshotArrayOutput() EvsSnapshotArrayOutput
	ToEvsSnapshotArrayOutputWithContext(context.Context) EvsSnapshotArrayOutput
}

type EvsSnapshotArray []EvsSnapshotInput

func (EvsSnapshotArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*EvsSnapshot)(nil)).Elem()
}

func (i EvsSnapshotArray) ToEvsSnapshotArrayOutput() EvsSnapshotArrayOutput {
	return i.ToEvsSnapshotArrayOutputWithContext(context.Background())
}

func (i EvsSnapshotArray) ToEvsSnapshotArrayOutputWithContext(ctx context.Context) EvsSnapshotArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EvsSnapshotArrayOutput)
}

// EvsSnapshotMapInput is an input type that accepts EvsSnapshotMap and EvsSnapshotMapOutput values.
// You can construct a concrete instance of `EvsSnapshotMapInput` via:
//
//	EvsSnapshotMap{ "key": EvsSnapshotArgs{...} }
type EvsSnapshotMapInput interface {
	pulumi.Input

	ToEvsSnapshotMapOutput() EvsSnapshotMapOutput
	ToEvsSnapshotMapOutputWithContext(context.Context) EvsSnapshotMapOutput
}

type EvsSnapshotMap map[string]EvsSnapshotInput

func (EvsSnapshotMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*EvsSnapshot)(nil)).Elem()
}

func (i EvsSnapshotMap) ToEvsSnapshotMapOutput() EvsSnapshotMapOutput {
	return i.ToEvsSnapshotMapOutputWithContext(context.Background())
}

func (i EvsSnapshotMap) ToEvsSnapshotMapOutputWithContext(ctx context.Context) EvsSnapshotMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EvsSnapshotMapOutput)
}

type EvsSnapshotOutput struct{ *pulumi.OutputState }

func (EvsSnapshotOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EvsSnapshot)(nil)).Elem()
}

func (o EvsSnapshotOutput) ToEvsSnapshotOutput() EvsSnapshotOutput {
	return o
}

func (o EvsSnapshotOutput) ToEvsSnapshotOutputWithContext(ctx context.Context) EvsSnapshotOutput {
	return o
}

func (o EvsSnapshotOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EvsSnapshot) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

func (o EvsSnapshotOutput) Force() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *EvsSnapshot) pulumi.BoolPtrOutput { return v.Force }).(pulumi.BoolPtrOutput)
}

func (o EvsSnapshotOutput) Metadata() pulumi.StringMapOutput {
	return o.ApplyT(func(v *EvsSnapshot) pulumi.StringMapOutput { return v.Metadata }).(pulumi.StringMapOutput)
}

func (o EvsSnapshotOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *EvsSnapshot) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o EvsSnapshotOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v *EvsSnapshot) pulumi.StringOutput { return v.Region }).(pulumi.StringOutput)
}

func (o EvsSnapshotOutput) Size() pulumi.IntOutput {
	return o.ApplyT(func(v *EvsSnapshot) pulumi.IntOutput { return v.Size }).(pulumi.IntOutput)
}

func (o EvsSnapshotOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *EvsSnapshot) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

func (o EvsSnapshotOutput) VolumeId() pulumi.StringOutput {
	return o.ApplyT(func(v *EvsSnapshot) pulumi.StringOutput { return v.VolumeId }).(pulumi.StringOutput)
}

type EvsSnapshotArrayOutput struct{ *pulumi.OutputState }

func (EvsSnapshotArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*EvsSnapshot)(nil)).Elem()
}

func (o EvsSnapshotArrayOutput) ToEvsSnapshotArrayOutput() EvsSnapshotArrayOutput {
	return o
}

func (o EvsSnapshotArrayOutput) ToEvsSnapshotArrayOutputWithContext(ctx context.Context) EvsSnapshotArrayOutput {
	return o
}

func (o EvsSnapshotArrayOutput) Index(i pulumi.IntInput) EvsSnapshotOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *EvsSnapshot {
		return vs[0].([]*EvsSnapshot)[vs[1].(int)]
	}).(EvsSnapshotOutput)
}

type EvsSnapshotMapOutput struct{ *pulumi.OutputState }

func (EvsSnapshotMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*EvsSnapshot)(nil)).Elem()
}

func (o EvsSnapshotMapOutput) ToEvsSnapshotMapOutput() EvsSnapshotMapOutput {
	return o
}

func (o EvsSnapshotMapOutput) ToEvsSnapshotMapOutputWithContext(ctx context.Context) EvsSnapshotMapOutput {
	return o
}

func (o EvsSnapshotMapOutput) MapIndex(k pulumi.StringInput) EvsSnapshotOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *EvsSnapshot {
		return vs[0].(map[string]*EvsSnapshot)[vs[1].(string)]
	}).(EvsSnapshotOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*EvsSnapshotInput)(nil)).Elem(), &EvsSnapshot{})
	pulumi.RegisterInputType(reflect.TypeOf((*EvsSnapshotArrayInput)(nil)).Elem(), EvsSnapshotArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*EvsSnapshotMapInput)(nil)).Elem(), EvsSnapshotMap{})
	pulumi.RegisterOutputType(EvsSnapshotOutput{})
	pulumi.RegisterOutputType(EvsSnapshotArrayOutput{})
	pulumi.RegisterOutputType(EvsSnapshotMapOutput{})
}
