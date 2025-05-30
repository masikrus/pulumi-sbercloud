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

type SfsTurbo struct {
	pulumi.CustomResourceState

	AutoRenew           pulumi.StringPtrOutput `pulumi:"autoRenew"`
	AvailabilityZone    pulumi.StringOutput    `pulumi:"availabilityZone"`
	AvailableCapacity   pulumi.StringOutput    `pulumi:"availableCapacity"`
	BackupId            pulumi.StringOutput    `pulumi:"backupId"`
	ChargingMode        pulumi.StringOutput    `pulumi:"chargingMode"`
	CryptKeyId          pulumi.StringPtrOutput `pulumi:"cryptKeyId"`
	DedicatedFlavor     pulumi.StringPtrOutput `pulumi:"dedicatedFlavor"`
	DedicatedStorageId  pulumi.StringPtrOutput `pulumi:"dedicatedStorageId"`
	Enhanced            pulumi.BoolOutput      `pulumi:"enhanced"`
	EnterpriseProjectId pulumi.StringOutput    `pulumi:"enterpriseProjectId"`
	ExportLocation      pulumi.StringOutput    `pulumi:"exportLocation"`
	HpcBandwidth        pulumi.StringOutput    `pulumi:"hpcBandwidth"`
	HpcCacheBandwidth   pulumi.StringOutput    `pulumi:"hpcCacheBandwidth"`
	Name                pulumi.StringOutput    `pulumi:"name"`
	Period              pulumi.IntPtrOutput    `pulumi:"period"`
	PeriodUnit          pulumi.StringPtrOutput `pulumi:"periodUnit"`
	Region              pulumi.StringOutput    `pulumi:"region"`
	SecurityGroupId     pulumi.StringOutput    `pulumi:"securityGroupId"`
	ShareProto          pulumi.StringPtrOutput `pulumi:"shareProto"`
	ShareType           pulumi.StringPtrOutput `pulumi:"shareType"`
	Size                pulumi.IntOutput       `pulumi:"size"`
	Status              pulumi.StringOutput    `pulumi:"status"`
	SubnetId            pulumi.StringOutput    `pulumi:"subnetId"`
	Tags                pulumi.StringMapOutput `pulumi:"tags"`
	Version             pulumi.StringOutput    `pulumi:"version"`
	VpcId               pulumi.StringOutput    `pulumi:"vpcId"`
}

// NewSfsTurbo registers a new resource with the given unique name, arguments, and options.
func NewSfsTurbo(ctx *pulumi.Context,
	name string, args *SfsTurboArgs, opts ...pulumi.ResourceOption) (*SfsTurbo, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AvailabilityZone == nil {
		return nil, errors.New("invalid value for required argument 'AvailabilityZone'")
	}
	if args.SecurityGroupId == nil {
		return nil, errors.New("invalid value for required argument 'SecurityGroupId'")
	}
	if args.Size == nil {
		return nil, errors.New("invalid value for required argument 'Size'")
	}
	if args.SubnetId == nil {
		return nil, errors.New("invalid value for required argument 'SubnetId'")
	}
	if args.VpcId == nil {
		return nil, errors.New("invalid value for required argument 'VpcId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource SfsTurbo
	err := ctx.RegisterResource("sbercloud:index/sfsTurbo:SfsTurbo", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSfsTurbo gets an existing SfsTurbo resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSfsTurbo(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SfsTurboState, opts ...pulumi.ResourceOption) (*SfsTurbo, error) {
	var resource SfsTurbo
	err := ctx.ReadResource("sbercloud:index/sfsTurbo:SfsTurbo", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SfsTurbo resources.
type sfsTurboState struct {
	AutoRenew           *string           `pulumi:"autoRenew"`
	AvailabilityZone    *string           `pulumi:"availabilityZone"`
	AvailableCapacity   *string           `pulumi:"availableCapacity"`
	BackupId            *string           `pulumi:"backupId"`
	ChargingMode        *string           `pulumi:"chargingMode"`
	CryptKeyId          *string           `pulumi:"cryptKeyId"`
	DedicatedFlavor     *string           `pulumi:"dedicatedFlavor"`
	DedicatedStorageId  *string           `pulumi:"dedicatedStorageId"`
	Enhanced            *bool             `pulumi:"enhanced"`
	EnterpriseProjectId *string           `pulumi:"enterpriseProjectId"`
	ExportLocation      *string           `pulumi:"exportLocation"`
	HpcBandwidth        *string           `pulumi:"hpcBandwidth"`
	HpcCacheBandwidth   *string           `pulumi:"hpcCacheBandwidth"`
	Name                *string           `pulumi:"name"`
	Period              *int              `pulumi:"period"`
	PeriodUnit          *string           `pulumi:"periodUnit"`
	Region              *string           `pulumi:"region"`
	SecurityGroupId     *string           `pulumi:"securityGroupId"`
	ShareProto          *string           `pulumi:"shareProto"`
	ShareType           *string           `pulumi:"shareType"`
	Size                *int              `pulumi:"size"`
	Status              *string           `pulumi:"status"`
	SubnetId            *string           `pulumi:"subnetId"`
	Tags                map[string]string `pulumi:"tags"`
	Version             *string           `pulumi:"version"`
	VpcId               *string           `pulumi:"vpcId"`
}

type SfsTurboState struct {
	AutoRenew           pulumi.StringPtrInput
	AvailabilityZone    pulumi.StringPtrInput
	AvailableCapacity   pulumi.StringPtrInput
	BackupId            pulumi.StringPtrInput
	ChargingMode        pulumi.StringPtrInput
	CryptKeyId          pulumi.StringPtrInput
	DedicatedFlavor     pulumi.StringPtrInput
	DedicatedStorageId  pulumi.StringPtrInput
	Enhanced            pulumi.BoolPtrInput
	EnterpriseProjectId pulumi.StringPtrInput
	ExportLocation      pulumi.StringPtrInput
	HpcBandwidth        pulumi.StringPtrInput
	HpcCacheBandwidth   pulumi.StringPtrInput
	Name                pulumi.StringPtrInput
	Period              pulumi.IntPtrInput
	PeriodUnit          pulumi.StringPtrInput
	Region              pulumi.StringPtrInput
	SecurityGroupId     pulumi.StringPtrInput
	ShareProto          pulumi.StringPtrInput
	ShareType           pulumi.StringPtrInput
	Size                pulumi.IntPtrInput
	Status              pulumi.StringPtrInput
	SubnetId            pulumi.StringPtrInput
	Tags                pulumi.StringMapInput
	Version             pulumi.StringPtrInput
	VpcId               pulumi.StringPtrInput
}

func (SfsTurboState) ElementType() reflect.Type {
	return reflect.TypeOf((*sfsTurboState)(nil)).Elem()
}

type sfsTurboArgs struct {
	AutoRenew           *string           `pulumi:"autoRenew"`
	AvailabilityZone    string            `pulumi:"availabilityZone"`
	BackupId            *string           `pulumi:"backupId"`
	ChargingMode        *string           `pulumi:"chargingMode"`
	CryptKeyId          *string           `pulumi:"cryptKeyId"`
	DedicatedFlavor     *string           `pulumi:"dedicatedFlavor"`
	DedicatedStorageId  *string           `pulumi:"dedicatedStorageId"`
	Enhanced            *bool             `pulumi:"enhanced"`
	EnterpriseProjectId *string           `pulumi:"enterpriseProjectId"`
	HpcBandwidth        *string           `pulumi:"hpcBandwidth"`
	HpcCacheBandwidth   *string           `pulumi:"hpcCacheBandwidth"`
	Name                *string           `pulumi:"name"`
	Period              *int              `pulumi:"period"`
	PeriodUnit          *string           `pulumi:"periodUnit"`
	Region              *string           `pulumi:"region"`
	SecurityGroupId     string            `pulumi:"securityGroupId"`
	ShareProto          *string           `pulumi:"shareProto"`
	ShareType           *string           `pulumi:"shareType"`
	Size                int               `pulumi:"size"`
	SubnetId            string            `pulumi:"subnetId"`
	Tags                map[string]string `pulumi:"tags"`
	VpcId               string            `pulumi:"vpcId"`
}

// The set of arguments for constructing a SfsTurbo resource.
type SfsTurboArgs struct {
	AutoRenew           pulumi.StringPtrInput
	AvailabilityZone    pulumi.StringInput
	BackupId            pulumi.StringPtrInput
	ChargingMode        pulumi.StringPtrInput
	CryptKeyId          pulumi.StringPtrInput
	DedicatedFlavor     pulumi.StringPtrInput
	DedicatedStorageId  pulumi.StringPtrInput
	Enhanced            pulumi.BoolPtrInput
	EnterpriseProjectId pulumi.StringPtrInput
	HpcBandwidth        pulumi.StringPtrInput
	HpcCacheBandwidth   pulumi.StringPtrInput
	Name                pulumi.StringPtrInput
	Period              pulumi.IntPtrInput
	PeriodUnit          pulumi.StringPtrInput
	Region              pulumi.StringPtrInput
	SecurityGroupId     pulumi.StringInput
	ShareProto          pulumi.StringPtrInput
	ShareType           pulumi.StringPtrInput
	Size                pulumi.IntInput
	SubnetId            pulumi.StringInput
	Tags                pulumi.StringMapInput
	VpcId               pulumi.StringInput
}

func (SfsTurboArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*sfsTurboArgs)(nil)).Elem()
}

type SfsTurboInput interface {
	pulumi.Input

	ToSfsTurboOutput() SfsTurboOutput
	ToSfsTurboOutputWithContext(ctx context.Context) SfsTurboOutput
}

func (*SfsTurbo) ElementType() reflect.Type {
	return reflect.TypeOf((**SfsTurbo)(nil)).Elem()
}

func (i *SfsTurbo) ToSfsTurboOutput() SfsTurboOutput {
	return i.ToSfsTurboOutputWithContext(context.Background())
}

func (i *SfsTurbo) ToSfsTurboOutputWithContext(ctx context.Context) SfsTurboOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SfsTurboOutput)
}

// SfsTurboArrayInput is an input type that accepts SfsTurboArray and SfsTurboArrayOutput values.
// You can construct a concrete instance of `SfsTurboArrayInput` via:
//
//	SfsTurboArray{ SfsTurboArgs{...} }
type SfsTurboArrayInput interface {
	pulumi.Input

	ToSfsTurboArrayOutput() SfsTurboArrayOutput
	ToSfsTurboArrayOutputWithContext(context.Context) SfsTurboArrayOutput
}

type SfsTurboArray []SfsTurboInput

func (SfsTurboArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SfsTurbo)(nil)).Elem()
}

func (i SfsTurboArray) ToSfsTurboArrayOutput() SfsTurboArrayOutput {
	return i.ToSfsTurboArrayOutputWithContext(context.Background())
}

func (i SfsTurboArray) ToSfsTurboArrayOutputWithContext(ctx context.Context) SfsTurboArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SfsTurboArrayOutput)
}

// SfsTurboMapInput is an input type that accepts SfsTurboMap and SfsTurboMapOutput values.
// You can construct a concrete instance of `SfsTurboMapInput` via:
//
//	SfsTurboMap{ "key": SfsTurboArgs{...} }
type SfsTurboMapInput interface {
	pulumi.Input

	ToSfsTurboMapOutput() SfsTurboMapOutput
	ToSfsTurboMapOutputWithContext(context.Context) SfsTurboMapOutput
}

type SfsTurboMap map[string]SfsTurboInput

func (SfsTurboMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SfsTurbo)(nil)).Elem()
}

func (i SfsTurboMap) ToSfsTurboMapOutput() SfsTurboMapOutput {
	return i.ToSfsTurboMapOutputWithContext(context.Background())
}

func (i SfsTurboMap) ToSfsTurboMapOutputWithContext(ctx context.Context) SfsTurboMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SfsTurboMapOutput)
}

type SfsTurboOutput struct{ *pulumi.OutputState }

func (SfsTurboOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SfsTurbo)(nil)).Elem()
}

func (o SfsTurboOutput) ToSfsTurboOutput() SfsTurboOutput {
	return o
}

func (o SfsTurboOutput) ToSfsTurboOutputWithContext(ctx context.Context) SfsTurboOutput {
	return o
}

func (o SfsTurboOutput) AutoRenew() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SfsTurbo) pulumi.StringPtrOutput { return v.AutoRenew }).(pulumi.StringPtrOutput)
}

func (o SfsTurboOutput) AvailabilityZone() pulumi.StringOutput {
	return o.ApplyT(func(v *SfsTurbo) pulumi.StringOutput { return v.AvailabilityZone }).(pulumi.StringOutput)
}

func (o SfsTurboOutput) AvailableCapacity() pulumi.StringOutput {
	return o.ApplyT(func(v *SfsTurbo) pulumi.StringOutput { return v.AvailableCapacity }).(pulumi.StringOutput)
}

func (o SfsTurboOutput) BackupId() pulumi.StringOutput {
	return o.ApplyT(func(v *SfsTurbo) pulumi.StringOutput { return v.BackupId }).(pulumi.StringOutput)
}

func (o SfsTurboOutput) ChargingMode() pulumi.StringOutput {
	return o.ApplyT(func(v *SfsTurbo) pulumi.StringOutput { return v.ChargingMode }).(pulumi.StringOutput)
}

func (o SfsTurboOutput) CryptKeyId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SfsTurbo) pulumi.StringPtrOutput { return v.CryptKeyId }).(pulumi.StringPtrOutput)
}

func (o SfsTurboOutput) DedicatedFlavor() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SfsTurbo) pulumi.StringPtrOutput { return v.DedicatedFlavor }).(pulumi.StringPtrOutput)
}

func (o SfsTurboOutput) DedicatedStorageId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SfsTurbo) pulumi.StringPtrOutput { return v.DedicatedStorageId }).(pulumi.StringPtrOutput)
}

func (o SfsTurboOutput) Enhanced() pulumi.BoolOutput {
	return o.ApplyT(func(v *SfsTurbo) pulumi.BoolOutput { return v.Enhanced }).(pulumi.BoolOutput)
}

func (o SfsTurboOutput) EnterpriseProjectId() pulumi.StringOutput {
	return o.ApplyT(func(v *SfsTurbo) pulumi.StringOutput { return v.EnterpriseProjectId }).(pulumi.StringOutput)
}

func (o SfsTurboOutput) ExportLocation() pulumi.StringOutput {
	return o.ApplyT(func(v *SfsTurbo) pulumi.StringOutput { return v.ExportLocation }).(pulumi.StringOutput)
}

func (o SfsTurboOutput) HpcBandwidth() pulumi.StringOutput {
	return o.ApplyT(func(v *SfsTurbo) pulumi.StringOutput { return v.HpcBandwidth }).(pulumi.StringOutput)
}

func (o SfsTurboOutput) HpcCacheBandwidth() pulumi.StringOutput {
	return o.ApplyT(func(v *SfsTurbo) pulumi.StringOutput { return v.HpcCacheBandwidth }).(pulumi.StringOutput)
}

func (o SfsTurboOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *SfsTurbo) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o SfsTurboOutput) Period() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *SfsTurbo) pulumi.IntPtrOutput { return v.Period }).(pulumi.IntPtrOutput)
}

func (o SfsTurboOutput) PeriodUnit() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SfsTurbo) pulumi.StringPtrOutput { return v.PeriodUnit }).(pulumi.StringPtrOutput)
}

func (o SfsTurboOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v *SfsTurbo) pulumi.StringOutput { return v.Region }).(pulumi.StringOutput)
}

func (o SfsTurboOutput) SecurityGroupId() pulumi.StringOutput {
	return o.ApplyT(func(v *SfsTurbo) pulumi.StringOutput { return v.SecurityGroupId }).(pulumi.StringOutput)
}

func (o SfsTurboOutput) ShareProto() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SfsTurbo) pulumi.StringPtrOutput { return v.ShareProto }).(pulumi.StringPtrOutput)
}

func (o SfsTurboOutput) ShareType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SfsTurbo) pulumi.StringPtrOutput { return v.ShareType }).(pulumi.StringPtrOutput)
}

func (o SfsTurboOutput) Size() pulumi.IntOutput {
	return o.ApplyT(func(v *SfsTurbo) pulumi.IntOutput { return v.Size }).(pulumi.IntOutput)
}

func (o SfsTurboOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *SfsTurbo) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

func (o SfsTurboOutput) SubnetId() pulumi.StringOutput {
	return o.ApplyT(func(v *SfsTurbo) pulumi.StringOutput { return v.SubnetId }).(pulumi.StringOutput)
}

func (o SfsTurboOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *SfsTurbo) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o SfsTurboOutput) Version() pulumi.StringOutput {
	return o.ApplyT(func(v *SfsTurbo) pulumi.StringOutput { return v.Version }).(pulumi.StringOutput)
}

func (o SfsTurboOutput) VpcId() pulumi.StringOutput {
	return o.ApplyT(func(v *SfsTurbo) pulumi.StringOutput { return v.VpcId }).(pulumi.StringOutput)
}

type SfsTurboArrayOutput struct{ *pulumi.OutputState }

func (SfsTurboArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SfsTurbo)(nil)).Elem()
}

func (o SfsTurboArrayOutput) ToSfsTurboArrayOutput() SfsTurboArrayOutput {
	return o
}

func (o SfsTurboArrayOutput) ToSfsTurboArrayOutputWithContext(ctx context.Context) SfsTurboArrayOutput {
	return o
}

func (o SfsTurboArrayOutput) Index(i pulumi.IntInput) SfsTurboOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SfsTurbo {
		return vs[0].([]*SfsTurbo)[vs[1].(int)]
	}).(SfsTurboOutput)
}

type SfsTurboMapOutput struct{ *pulumi.OutputState }

func (SfsTurboMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SfsTurbo)(nil)).Elem()
}

func (o SfsTurboMapOutput) ToSfsTurboMapOutput() SfsTurboMapOutput {
	return o
}

func (o SfsTurboMapOutput) ToSfsTurboMapOutputWithContext(ctx context.Context) SfsTurboMapOutput {
	return o
}

func (o SfsTurboMapOutput) MapIndex(k pulumi.StringInput) SfsTurboOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SfsTurbo {
		return vs[0].(map[string]*SfsTurbo)[vs[1].(string)]
	}).(SfsTurboOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SfsTurboInput)(nil)).Elem(), &SfsTurbo{})
	pulumi.RegisterInputType(reflect.TypeOf((*SfsTurboArrayInput)(nil)).Elem(), SfsTurboArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SfsTurboMapInput)(nil)).Elem(), SfsTurboMap{})
	pulumi.RegisterOutputType(SfsTurboOutput{})
	pulumi.RegisterOutputType(SfsTurboArrayOutput{})
	pulumi.RegisterOutputType(SfsTurboMapOutput{})
}
