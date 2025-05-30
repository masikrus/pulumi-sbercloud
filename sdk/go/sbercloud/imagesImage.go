// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package sbercloud

import (
	"context"
	"reflect"

	"github.com/masikrus/pulumi-sbercloud/sdk/go/sbercloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ImagesImage struct {
	pulumi.CustomResourceState

	BackupId            pulumi.StringOutput    `pulumi:"backupId"`
	Checksum            pulumi.StringOutput    `pulumi:"checksum"`
	CmkId               pulumi.StringPtrOutput `pulumi:"cmkId"`
	DataOrigin          pulumi.StringOutput    `pulumi:"dataOrigin"`
	Description         pulumi.StringPtrOutput `pulumi:"description"`
	DiskFormat          pulumi.StringOutput    `pulumi:"diskFormat"`
	EnterpriseProjectId pulumi.StringOutput    `pulumi:"enterpriseProjectId"`
	ImageSize           pulumi.StringOutput    `pulumi:"imageSize"`
	ImageUrl            pulumi.StringPtrOutput `pulumi:"imageUrl"`
	InstanceId          pulumi.StringOutput    `pulumi:"instanceId"`
	IsConfig            pulumi.BoolPtrOutput   `pulumi:"isConfig"`
	MaxRam              pulumi.IntOutput       `pulumi:"maxRam"`
	MinDisk             pulumi.IntPtrOutput    `pulumi:"minDisk"`
	MinRam              pulumi.IntOutput       `pulumi:"minRam"`
	Name                pulumi.StringOutput    `pulumi:"name"`
	OsVersion           pulumi.StringOutput    `pulumi:"osVersion"`
	Region              pulumi.StringOutput    `pulumi:"region"`
	Status              pulumi.StringOutput    `pulumi:"status"`
	Tags                pulumi.StringMapOutput `pulumi:"tags"`
	Type                pulumi.StringPtrOutput `pulumi:"type"`
	VaultId             pulumi.StringPtrOutput `pulumi:"vaultId"`
	Visibility          pulumi.StringOutput    `pulumi:"visibility"`
	VolumeId            pulumi.StringOutput    `pulumi:"volumeId"`
}

// NewImagesImage registers a new resource with the given unique name, arguments, and options.
func NewImagesImage(ctx *pulumi.Context,
	name string, args *ImagesImageArgs, opts ...pulumi.ResourceOption) (*ImagesImage, error) {
	if args == nil {
		args = &ImagesImageArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource ImagesImage
	err := ctx.RegisterResource("sbercloud:index/imagesImage:ImagesImage", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetImagesImage gets an existing ImagesImage resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetImagesImage(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ImagesImageState, opts ...pulumi.ResourceOption) (*ImagesImage, error) {
	var resource ImagesImage
	err := ctx.ReadResource("sbercloud:index/imagesImage:ImagesImage", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ImagesImage resources.
type imagesImageState struct {
	BackupId            *string           `pulumi:"backupId"`
	Checksum            *string           `pulumi:"checksum"`
	CmkId               *string           `pulumi:"cmkId"`
	DataOrigin          *string           `pulumi:"dataOrigin"`
	Description         *string           `pulumi:"description"`
	DiskFormat          *string           `pulumi:"diskFormat"`
	EnterpriseProjectId *string           `pulumi:"enterpriseProjectId"`
	ImageSize           *string           `pulumi:"imageSize"`
	ImageUrl            *string           `pulumi:"imageUrl"`
	InstanceId          *string           `pulumi:"instanceId"`
	IsConfig            *bool             `pulumi:"isConfig"`
	MaxRam              *int              `pulumi:"maxRam"`
	MinDisk             *int              `pulumi:"minDisk"`
	MinRam              *int              `pulumi:"minRam"`
	Name                *string           `pulumi:"name"`
	OsVersion           *string           `pulumi:"osVersion"`
	Region              *string           `pulumi:"region"`
	Status              *string           `pulumi:"status"`
	Tags                map[string]string `pulumi:"tags"`
	Type                *string           `pulumi:"type"`
	VaultId             *string           `pulumi:"vaultId"`
	Visibility          *string           `pulumi:"visibility"`
	VolumeId            *string           `pulumi:"volumeId"`
}

type ImagesImageState struct {
	BackupId            pulumi.StringPtrInput
	Checksum            pulumi.StringPtrInput
	CmkId               pulumi.StringPtrInput
	DataOrigin          pulumi.StringPtrInput
	Description         pulumi.StringPtrInput
	DiskFormat          pulumi.StringPtrInput
	EnterpriseProjectId pulumi.StringPtrInput
	ImageSize           pulumi.StringPtrInput
	ImageUrl            pulumi.StringPtrInput
	InstanceId          pulumi.StringPtrInput
	IsConfig            pulumi.BoolPtrInput
	MaxRam              pulumi.IntPtrInput
	MinDisk             pulumi.IntPtrInput
	MinRam              pulumi.IntPtrInput
	Name                pulumi.StringPtrInput
	OsVersion           pulumi.StringPtrInput
	Region              pulumi.StringPtrInput
	Status              pulumi.StringPtrInput
	Tags                pulumi.StringMapInput
	Type                pulumi.StringPtrInput
	VaultId             pulumi.StringPtrInput
	Visibility          pulumi.StringPtrInput
	VolumeId            pulumi.StringPtrInput
}

func (ImagesImageState) ElementType() reflect.Type {
	return reflect.TypeOf((*imagesImageState)(nil)).Elem()
}

type imagesImageArgs struct {
	BackupId            *string           `pulumi:"backupId"`
	CmkId               *string           `pulumi:"cmkId"`
	Description         *string           `pulumi:"description"`
	EnterpriseProjectId *string           `pulumi:"enterpriseProjectId"`
	ImageUrl            *string           `pulumi:"imageUrl"`
	InstanceId          *string           `pulumi:"instanceId"`
	IsConfig            *bool             `pulumi:"isConfig"`
	MaxRam              *int              `pulumi:"maxRam"`
	MinDisk             *int              `pulumi:"minDisk"`
	MinRam              *int              `pulumi:"minRam"`
	Name                *string           `pulumi:"name"`
	OsVersion           *string           `pulumi:"osVersion"`
	Region              *string           `pulumi:"region"`
	Tags                map[string]string `pulumi:"tags"`
	Type                *string           `pulumi:"type"`
	VaultId             *string           `pulumi:"vaultId"`
	VolumeId            *string           `pulumi:"volumeId"`
}

// The set of arguments for constructing a ImagesImage resource.
type ImagesImageArgs struct {
	BackupId            pulumi.StringPtrInput
	CmkId               pulumi.StringPtrInput
	Description         pulumi.StringPtrInput
	EnterpriseProjectId pulumi.StringPtrInput
	ImageUrl            pulumi.StringPtrInput
	InstanceId          pulumi.StringPtrInput
	IsConfig            pulumi.BoolPtrInput
	MaxRam              pulumi.IntPtrInput
	MinDisk             pulumi.IntPtrInput
	MinRam              pulumi.IntPtrInput
	Name                pulumi.StringPtrInput
	OsVersion           pulumi.StringPtrInput
	Region              pulumi.StringPtrInput
	Tags                pulumi.StringMapInput
	Type                pulumi.StringPtrInput
	VaultId             pulumi.StringPtrInput
	VolumeId            pulumi.StringPtrInput
}

func (ImagesImageArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*imagesImageArgs)(nil)).Elem()
}

type ImagesImageInput interface {
	pulumi.Input

	ToImagesImageOutput() ImagesImageOutput
	ToImagesImageOutputWithContext(ctx context.Context) ImagesImageOutput
}

func (*ImagesImage) ElementType() reflect.Type {
	return reflect.TypeOf((**ImagesImage)(nil)).Elem()
}

func (i *ImagesImage) ToImagesImageOutput() ImagesImageOutput {
	return i.ToImagesImageOutputWithContext(context.Background())
}

func (i *ImagesImage) ToImagesImageOutputWithContext(ctx context.Context) ImagesImageOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ImagesImageOutput)
}

// ImagesImageArrayInput is an input type that accepts ImagesImageArray and ImagesImageArrayOutput values.
// You can construct a concrete instance of `ImagesImageArrayInput` via:
//
//	ImagesImageArray{ ImagesImageArgs{...} }
type ImagesImageArrayInput interface {
	pulumi.Input

	ToImagesImageArrayOutput() ImagesImageArrayOutput
	ToImagesImageArrayOutputWithContext(context.Context) ImagesImageArrayOutput
}

type ImagesImageArray []ImagesImageInput

func (ImagesImageArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ImagesImage)(nil)).Elem()
}

func (i ImagesImageArray) ToImagesImageArrayOutput() ImagesImageArrayOutput {
	return i.ToImagesImageArrayOutputWithContext(context.Background())
}

func (i ImagesImageArray) ToImagesImageArrayOutputWithContext(ctx context.Context) ImagesImageArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ImagesImageArrayOutput)
}

// ImagesImageMapInput is an input type that accepts ImagesImageMap and ImagesImageMapOutput values.
// You can construct a concrete instance of `ImagesImageMapInput` via:
//
//	ImagesImageMap{ "key": ImagesImageArgs{...} }
type ImagesImageMapInput interface {
	pulumi.Input

	ToImagesImageMapOutput() ImagesImageMapOutput
	ToImagesImageMapOutputWithContext(context.Context) ImagesImageMapOutput
}

type ImagesImageMap map[string]ImagesImageInput

func (ImagesImageMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ImagesImage)(nil)).Elem()
}

func (i ImagesImageMap) ToImagesImageMapOutput() ImagesImageMapOutput {
	return i.ToImagesImageMapOutputWithContext(context.Background())
}

func (i ImagesImageMap) ToImagesImageMapOutputWithContext(ctx context.Context) ImagesImageMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ImagesImageMapOutput)
}

type ImagesImageOutput struct{ *pulumi.OutputState }

func (ImagesImageOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ImagesImage)(nil)).Elem()
}

func (o ImagesImageOutput) ToImagesImageOutput() ImagesImageOutput {
	return o
}

func (o ImagesImageOutput) ToImagesImageOutputWithContext(ctx context.Context) ImagesImageOutput {
	return o
}

func (o ImagesImageOutput) BackupId() pulumi.StringOutput {
	return o.ApplyT(func(v *ImagesImage) pulumi.StringOutput { return v.BackupId }).(pulumi.StringOutput)
}

func (o ImagesImageOutput) Checksum() pulumi.StringOutput {
	return o.ApplyT(func(v *ImagesImage) pulumi.StringOutput { return v.Checksum }).(pulumi.StringOutput)
}

func (o ImagesImageOutput) CmkId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ImagesImage) pulumi.StringPtrOutput { return v.CmkId }).(pulumi.StringPtrOutput)
}

func (o ImagesImageOutput) DataOrigin() pulumi.StringOutput {
	return o.ApplyT(func(v *ImagesImage) pulumi.StringOutput { return v.DataOrigin }).(pulumi.StringOutput)
}

func (o ImagesImageOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ImagesImage) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

func (o ImagesImageOutput) DiskFormat() pulumi.StringOutput {
	return o.ApplyT(func(v *ImagesImage) pulumi.StringOutput { return v.DiskFormat }).(pulumi.StringOutput)
}

func (o ImagesImageOutput) EnterpriseProjectId() pulumi.StringOutput {
	return o.ApplyT(func(v *ImagesImage) pulumi.StringOutput { return v.EnterpriseProjectId }).(pulumi.StringOutput)
}

func (o ImagesImageOutput) ImageSize() pulumi.StringOutput {
	return o.ApplyT(func(v *ImagesImage) pulumi.StringOutput { return v.ImageSize }).(pulumi.StringOutput)
}

func (o ImagesImageOutput) ImageUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ImagesImage) pulumi.StringPtrOutput { return v.ImageUrl }).(pulumi.StringPtrOutput)
}

func (o ImagesImageOutput) InstanceId() pulumi.StringOutput {
	return o.ApplyT(func(v *ImagesImage) pulumi.StringOutput { return v.InstanceId }).(pulumi.StringOutput)
}

func (o ImagesImageOutput) IsConfig() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ImagesImage) pulumi.BoolPtrOutput { return v.IsConfig }).(pulumi.BoolPtrOutput)
}

func (o ImagesImageOutput) MaxRam() pulumi.IntOutput {
	return o.ApplyT(func(v *ImagesImage) pulumi.IntOutput { return v.MaxRam }).(pulumi.IntOutput)
}

func (o ImagesImageOutput) MinDisk() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ImagesImage) pulumi.IntPtrOutput { return v.MinDisk }).(pulumi.IntPtrOutput)
}

func (o ImagesImageOutput) MinRam() pulumi.IntOutput {
	return o.ApplyT(func(v *ImagesImage) pulumi.IntOutput { return v.MinRam }).(pulumi.IntOutput)
}

func (o ImagesImageOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ImagesImage) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o ImagesImageOutput) OsVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *ImagesImage) pulumi.StringOutput { return v.OsVersion }).(pulumi.StringOutput)
}

func (o ImagesImageOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v *ImagesImage) pulumi.StringOutput { return v.Region }).(pulumi.StringOutput)
}

func (o ImagesImageOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *ImagesImage) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

func (o ImagesImageOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *ImagesImage) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o ImagesImageOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ImagesImage) pulumi.StringPtrOutput { return v.Type }).(pulumi.StringPtrOutput)
}

func (o ImagesImageOutput) VaultId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ImagesImage) pulumi.StringPtrOutput { return v.VaultId }).(pulumi.StringPtrOutput)
}

func (o ImagesImageOutput) Visibility() pulumi.StringOutput {
	return o.ApplyT(func(v *ImagesImage) pulumi.StringOutput { return v.Visibility }).(pulumi.StringOutput)
}

func (o ImagesImageOutput) VolumeId() pulumi.StringOutput {
	return o.ApplyT(func(v *ImagesImage) pulumi.StringOutput { return v.VolumeId }).(pulumi.StringOutput)
}

type ImagesImageArrayOutput struct{ *pulumi.OutputState }

func (ImagesImageArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ImagesImage)(nil)).Elem()
}

func (o ImagesImageArrayOutput) ToImagesImageArrayOutput() ImagesImageArrayOutput {
	return o
}

func (o ImagesImageArrayOutput) ToImagesImageArrayOutputWithContext(ctx context.Context) ImagesImageArrayOutput {
	return o
}

func (o ImagesImageArrayOutput) Index(i pulumi.IntInput) ImagesImageOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ImagesImage {
		return vs[0].([]*ImagesImage)[vs[1].(int)]
	}).(ImagesImageOutput)
}

type ImagesImageMapOutput struct{ *pulumi.OutputState }

func (ImagesImageMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ImagesImage)(nil)).Elem()
}

func (o ImagesImageMapOutput) ToImagesImageMapOutput() ImagesImageMapOutput {
	return o
}

func (o ImagesImageMapOutput) ToImagesImageMapOutputWithContext(ctx context.Context) ImagesImageMapOutput {
	return o
}

func (o ImagesImageMapOutput) MapIndex(k pulumi.StringInput) ImagesImageOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ImagesImage {
		return vs[0].(map[string]*ImagesImage)[vs[1].(string)]
	}).(ImagesImageOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ImagesImageInput)(nil)).Elem(), &ImagesImage{})
	pulumi.RegisterInputType(reflect.TypeOf((*ImagesImageArrayInput)(nil)).Elem(), ImagesImageArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ImagesImageMapInput)(nil)).Elem(), ImagesImageMap{})
	pulumi.RegisterOutputType(ImagesImageOutput{})
	pulumi.RegisterOutputType(ImagesImageArrayOutput{})
	pulumi.RegisterOutputType(ImagesImageMapOutput{})
}
