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

type SwrRepository struct {
	pulumi.CustomResourceState

	Category     pulumi.StringPtrOutput `pulumi:"category"`
	Description  pulumi.StringPtrOutput `pulumi:"description"`
	InternalPath pulumi.StringOutput    `pulumi:"internalPath"`
	IsPublic     pulumi.BoolPtrOutput   `pulumi:"isPublic"`
	Name         pulumi.StringOutput    `pulumi:"name"`
	NumImages    pulumi.IntOutput       `pulumi:"numImages"`
	Organization pulumi.StringOutput    `pulumi:"organization"`
	Path         pulumi.StringOutput    `pulumi:"path"`
	Region       pulumi.StringOutput    `pulumi:"region"`
	RepositoryId pulumi.IntOutput       `pulumi:"repositoryId"`
	Size         pulumi.IntOutput       `pulumi:"size"`
}

// NewSwrRepository registers a new resource with the given unique name, arguments, and options.
func NewSwrRepository(ctx *pulumi.Context,
	name string, args *SwrRepositoryArgs, opts ...pulumi.ResourceOption) (*SwrRepository, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Organization == nil {
		return nil, errors.New("invalid value for required argument 'Organization'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource SwrRepository
	err := ctx.RegisterResource("sbercloud:index/swrRepository:SwrRepository", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSwrRepository gets an existing SwrRepository resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSwrRepository(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SwrRepositoryState, opts ...pulumi.ResourceOption) (*SwrRepository, error) {
	var resource SwrRepository
	err := ctx.ReadResource("sbercloud:index/swrRepository:SwrRepository", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SwrRepository resources.
type swrRepositoryState struct {
	Category     *string `pulumi:"category"`
	Description  *string `pulumi:"description"`
	InternalPath *string `pulumi:"internalPath"`
	IsPublic     *bool   `pulumi:"isPublic"`
	Name         *string `pulumi:"name"`
	NumImages    *int    `pulumi:"numImages"`
	Organization *string `pulumi:"organization"`
	Path         *string `pulumi:"path"`
	Region       *string `pulumi:"region"`
	RepositoryId *int    `pulumi:"repositoryId"`
	Size         *int    `pulumi:"size"`
}

type SwrRepositoryState struct {
	Category     pulumi.StringPtrInput
	Description  pulumi.StringPtrInput
	InternalPath pulumi.StringPtrInput
	IsPublic     pulumi.BoolPtrInput
	Name         pulumi.StringPtrInput
	NumImages    pulumi.IntPtrInput
	Organization pulumi.StringPtrInput
	Path         pulumi.StringPtrInput
	Region       pulumi.StringPtrInput
	RepositoryId pulumi.IntPtrInput
	Size         pulumi.IntPtrInput
}

func (SwrRepositoryState) ElementType() reflect.Type {
	return reflect.TypeOf((*swrRepositoryState)(nil)).Elem()
}

type swrRepositoryArgs struct {
	Category     *string `pulumi:"category"`
	Description  *string `pulumi:"description"`
	IsPublic     *bool   `pulumi:"isPublic"`
	Name         *string `pulumi:"name"`
	Organization string  `pulumi:"organization"`
	Region       *string `pulumi:"region"`
}

// The set of arguments for constructing a SwrRepository resource.
type SwrRepositoryArgs struct {
	Category     pulumi.StringPtrInput
	Description  pulumi.StringPtrInput
	IsPublic     pulumi.BoolPtrInput
	Name         pulumi.StringPtrInput
	Organization pulumi.StringInput
	Region       pulumi.StringPtrInput
}

func (SwrRepositoryArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*swrRepositoryArgs)(nil)).Elem()
}

type SwrRepositoryInput interface {
	pulumi.Input

	ToSwrRepositoryOutput() SwrRepositoryOutput
	ToSwrRepositoryOutputWithContext(ctx context.Context) SwrRepositoryOutput
}

func (*SwrRepository) ElementType() reflect.Type {
	return reflect.TypeOf((**SwrRepository)(nil)).Elem()
}

func (i *SwrRepository) ToSwrRepositoryOutput() SwrRepositoryOutput {
	return i.ToSwrRepositoryOutputWithContext(context.Background())
}

func (i *SwrRepository) ToSwrRepositoryOutputWithContext(ctx context.Context) SwrRepositoryOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SwrRepositoryOutput)
}

// SwrRepositoryArrayInput is an input type that accepts SwrRepositoryArray and SwrRepositoryArrayOutput values.
// You can construct a concrete instance of `SwrRepositoryArrayInput` via:
//
//	SwrRepositoryArray{ SwrRepositoryArgs{...} }
type SwrRepositoryArrayInput interface {
	pulumi.Input

	ToSwrRepositoryArrayOutput() SwrRepositoryArrayOutput
	ToSwrRepositoryArrayOutputWithContext(context.Context) SwrRepositoryArrayOutput
}

type SwrRepositoryArray []SwrRepositoryInput

func (SwrRepositoryArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SwrRepository)(nil)).Elem()
}

func (i SwrRepositoryArray) ToSwrRepositoryArrayOutput() SwrRepositoryArrayOutput {
	return i.ToSwrRepositoryArrayOutputWithContext(context.Background())
}

func (i SwrRepositoryArray) ToSwrRepositoryArrayOutputWithContext(ctx context.Context) SwrRepositoryArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SwrRepositoryArrayOutput)
}

// SwrRepositoryMapInput is an input type that accepts SwrRepositoryMap and SwrRepositoryMapOutput values.
// You can construct a concrete instance of `SwrRepositoryMapInput` via:
//
//	SwrRepositoryMap{ "key": SwrRepositoryArgs{...} }
type SwrRepositoryMapInput interface {
	pulumi.Input

	ToSwrRepositoryMapOutput() SwrRepositoryMapOutput
	ToSwrRepositoryMapOutputWithContext(context.Context) SwrRepositoryMapOutput
}

type SwrRepositoryMap map[string]SwrRepositoryInput

func (SwrRepositoryMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SwrRepository)(nil)).Elem()
}

func (i SwrRepositoryMap) ToSwrRepositoryMapOutput() SwrRepositoryMapOutput {
	return i.ToSwrRepositoryMapOutputWithContext(context.Background())
}

func (i SwrRepositoryMap) ToSwrRepositoryMapOutputWithContext(ctx context.Context) SwrRepositoryMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SwrRepositoryMapOutput)
}

type SwrRepositoryOutput struct{ *pulumi.OutputState }

func (SwrRepositoryOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SwrRepository)(nil)).Elem()
}

func (o SwrRepositoryOutput) ToSwrRepositoryOutput() SwrRepositoryOutput {
	return o
}

func (o SwrRepositoryOutput) ToSwrRepositoryOutputWithContext(ctx context.Context) SwrRepositoryOutput {
	return o
}

func (o SwrRepositoryOutput) Category() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SwrRepository) pulumi.StringPtrOutput { return v.Category }).(pulumi.StringPtrOutput)
}

func (o SwrRepositoryOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SwrRepository) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

func (o SwrRepositoryOutput) InternalPath() pulumi.StringOutput {
	return o.ApplyT(func(v *SwrRepository) pulumi.StringOutput { return v.InternalPath }).(pulumi.StringOutput)
}

func (o SwrRepositoryOutput) IsPublic() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *SwrRepository) pulumi.BoolPtrOutput { return v.IsPublic }).(pulumi.BoolPtrOutput)
}

func (o SwrRepositoryOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *SwrRepository) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o SwrRepositoryOutput) NumImages() pulumi.IntOutput {
	return o.ApplyT(func(v *SwrRepository) pulumi.IntOutput { return v.NumImages }).(pulumi.IntOutput)
}

func (o SwrRepositoryOutput) Organization() pulumi.StringOutput {
	return o.ApplyT(func(v *SwrRepository) pulumi.StringOutput { return v.Organization }).(pulumi.StringOutput)
}

func (o SwrRepositoryOutput) Path() pulumi.StringOutput {
	return o.ApplyT(func(v *SwrRepository) pulumi.StringOutput { return v.Path }).(pulumi.StringOutput)
}

func (o SwrRepositoryOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v *SwrRepository) pulumi.StringOutput { return v.Region }).(pulumi.StringOutput)
}

func (o SwrRepositoryOutput) RepositoryId() pulumi.IntOutput {
	return o.ApplyT(func(v *SwrRepository) pulumi.IntOutput { return v.RepositoryId }).(pulumi.IntOutput)
}

func (o SwrRepositoryOutput) Size() pulumi.IntOutput {
	return o.ApplyT(func(v *SwrRepository) pulumi.IntOutput { return v.Size }).(pulumi.IntOutput)
}

type SwrRepositoryArrayOutput struct{ *pulumi.OutputState }

func (SwrRepositoryArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SwrRepository)(nil)).Elem()
}

func (o SwrRepositoryArrayOutput) ToSwrRepositoryArrayOutput() SwrRepositoryArrayOutput {
	return o
}

func (o SwrRepositoryArrayOutput) ToSwrRepositoryArrayOutputWithContext(ctx context.Context) SwrRepositoryArrayOutput {
	return o
}

func (o SwrRepositoryArrayOutput) Index(i pulumi.IntInput) SwrRepositoryOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SwrRepository {
		return vs[0].([]*SwrRepository)[vs[1].(int)]
	}).(SwrRepositoryOutput)
}

type SwrRepositoryMapOutput struct{ *pulumi.OutputState }

func (SwrRepositoryMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SwrRepository)(nil)).Elem()
}

func (o SwrRepositoryMapOutput) ToSwrRepositoryMapOutput() SwrRepositoryMapOutput {
	return o
}

func (o SwrRepositoryMapOutput) ToSwrRepositoryMapOutputWithContext(ctx context.Context) SwrRepositoryMapOutput {
	return o
}

func (o SwrRepositoryMapOutput) MapIndex(k pulumi.StringInput) SwrRepositoryOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SwrRepository {
		return vs[0].(map[string]*SwrRepository)[vs[1].(string)]
	}).(SwrRepositoryOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SwrRepositoryInput)(nil)).Elem(), &SwrRepository{})
	pulumi.RegisterInputType(reflect.TypeOf((*SwrRepositoryArrayInput)(nil)).Elem(), SwrRepositoryArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SwrRepositoryMapInput)(nil)).Elem(), SwrRepositoryMap{})
	pulumi.RegisterOutputType(SwrRepositoryOutput{})
	pulumi.RegisterOutputType(SwrRepositoryArrayOutput{})
	pulumi.RegisterOutputType(SwrRepositoryMapOutput{})
}
