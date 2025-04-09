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

type CtsDataTracker struct {
	pulumi.CustomResourceState

	AgencyName         pulumi.StringOutput      `pulumi:"agencyName"`
	BucketName         pulumi.StringPtrOutput   `pulumi:"bucketName"`
	CompressType       pulumi.StringPtrOutput   `pulumi:"compressType"`
	CreateTime         pulumi.IntOutput         `pulumi:"createTime"`
	DataBucket         pulumi.StringOutput      `pulumi:"dataBucket"`
	DataOperations     pulumi.StringArrayOutput `pulumi:"dataOperations"`
	Detail             pulumi.StringOutput      `pulumi:"detail"`
	DomainId           pulumi.StringOutput      `pulumi:"domainId"`
	Enabled            pulumi.BoolPtrOutput     `pulumi:"enabled"`
	FilePrefix         pulumi.StringPtrOutput   `pulumi:"filePrefix"`
	GroupId            pulumi.StringOutput      `pulumi:"groupId"`
	IsAuthorizedBucket pulumi.BoolOutput        `pulumi:"isAuthorizedBucket"`
	IsSortByService    pulumi.BoolPtrOutput     `pulumi:"isSortByService"`
	LogGroupName       pulumi.StringOutput      `pulumi:"logGroupName"`
	LogTopicName       pulumi.StringOutput      `pulumi:"logTopicName"`
	LtsEnabled         pulumi.BoolPtrOutput     `pulumi:"ltsEnabled"`
	Name               pulumi.StringOutput      `pulumi:"name"`
	ObsRetentionPeriod pulumi.IntPtrOutput      `pulumi:"obsRetentionPeriod"`
	Region             pulumi.StringOutput      `pulumi:"region"`
	Status             pulumi.StringOutput      `pulumi:"status"`
	StreamId           pulumi.StringOutput      `pulumi:"streamId"`
	Tags               pulumi.StringMapOutput   `pulumi:"tags"`
	TransferEnabled    pulumi.BoolOutput        `pulumi:"transferEnabled"`
	Type               pulumi.StringOutput      `pulumi:"type"`
	ValidateFile       pulumi.BoolPtrOutput     `pulumi:"validateFile"`
}

// NewCtsDataTracker registers a new resource with the given unique name, arguments, and options.
func NewCtsDataTracker(ctx *pulumi.Context,
	name string, args *CtsDataTrackerArgs, opts ...pulumi.ResourceOption) (*CtsDataTracker, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DataBucket == nil {
		return nil, errors.New("invalid value for required argument 'DataBucket'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource CtsDataTracker
	err := ctx.RegisterResource("sbercloud:index/ctsDataTracker:CtsDataTracker", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCtsDataTracker gets an existing CtsDataTracker resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCtsDataTracker(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CtsDataTrackerState, opts ...pulumi.ResourceOption) (*CtsDataTracker, error) {
	var resource CtsDataTracker
	err := ctx.ReadResource("sbercloud:index/ctsDataTracker:CtsDataTracker", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering CtsDataTracker resources.
type ctsDataTrackerState struct {
	AgencyName         *string           `pulumi:"agencyName"`
	BucketName         *string           `pulumi:"bucketName"`
	CompressType       *string           `pulumi:"compressType"`
	CreateTime         *int              `pulumi:"createTime"`
	DataBucket         *string           `pulumi:"dataBucket"`
	DataOperations     []string          `pulumi:"dataOperations"`
	Detail             *string           `pulumi:"detail"`
	DomainId           *string           `pulumi:"domainId"`
	Enabled            *bool             `pulumi:"enabled"`
	FilePrefix         *string           `pulumi:"filePrefix"`
	GroupId            *string           `pulumi:"groupId"`
	IsAuthorizedBucket *bool             `pulumi:"isAuthorizedBucket"`
	IsSortByService    *bool             `pulumi:"isSortByService"`
	LogGroupName       *string           `pulumi:"logGroupName"`
	LogTopicName       *string           `pulumi:"logTopicName"`
	LtsEnabled         *bool             `pulumi:"ltsEnabled"`
	Name               *string           `pulumi:"name"`
	ObsRetentionPeriod *int              `pulumi:"obsRetentionPeriod"`
	Region             *string           `pulumi:"region"`
	Status             *string           `pulumi:"status"`
	StreamId           *string           `pulumi:"streamId"`
	Tags               map[string]string `pulumi:"tags"`
	TransferEnabled    *bool             `pulumi:"transferEnabled"`
	Type               *string           `pulumi:"type"`
	ValidateFile       *bool             `pulumi:"validateFile"`
}

type CtsDataTrackerState struct {
	AgencyName         pulumi.StringPtrInput
	BucketName         pulumi.StringPtrInput
	CompressType       pulumi.StringPtrInput
	CreateTime         pulumi.IntPtrInput
	DataBucket         pulumi.StringPtrInput
	DataOperations     pulumi.StringArrayInput
	Detail             pulumi.StringPtrInput
	DomainId           pulumi.StringPtrInput
	Enabled            pulumi.BoolPtrInput
	FilePrefix         pulumi.StringPtrInput
	GroupId            pulumi.StringPtrInput
	IsAuthorizedBucket pulumi.BoolPtrInput
	IsSortByService    pulumi.BoolPtrInput
	LogGroupName       pulumi.StringPtrInput
	LogTopicName       pulumi.StringPtrInput
	LtsEnabled         pulumi.BoolPtrInput
	Name               pulumi.StringPtrInput
	ObsRetentionPeriod pulumi.IntPtrInput
	Region             pulumi.StringPtrInput
	Status             pulumi.StringPtrInput
	StreamId           pulumi.StringPtrInput
	Tags               pulumi.StringMapInput
	TransferEnabled    pulumi.BoolPtrInput
	Type               pulumi.StringPtrInput
	ValidateFile       pulumi.BoolPtrInput
}

func (CtsDataTrackerState) ElementType() reflect.Type {
	return reflect.TypeOf((*ctsDataTrackerState)(nil)).Elem()
}

type ctsDataTrackerArgs struct {
	BucketName         *string           `pulumi:"bucketName"`
	CompressType       *string           `pulumi:"compressType"`
	DataBucket         string            `pulumi:"dataBucket"`
	DataOperations     []string          `pulumi:"dataOperations"`
	Enabled            *bool             `pulumi:"enabled"`
	FilePrefix         *string           `pulumi:"filePrefix"`
	IsSortByService    *bool             `pulumi:"isSortByService"`
	LtsEnabled         *bool             `pulumi:"ltsEnabled"`
	Name               *string           `pulumi:"name"`
	ObsRetentionPeriod *int              `pulumi:"obsRetentionPeriod"`
	Region             *string           `pulumi:"region"`
	Tags               map[string]string `pulumi:"tags"`
	ValidateFile       *bool             `pulumi:"validateFile"`
}

// The set of arguments for constructing a CtsDataTracker resource.
type CtsDataTrackerArgs struct {
	BucketName         pulumi.StringPtrInput
	CompressType       pulumi.StringPtrInput
	DataBucket         pulumi.StringInput
	DataOperations     pulumi.StringArrayInput
	Enabled            pulumi.BoolPtrInput
	FilePrefix         pulumi.StringPtrInput
	IsSortByService    pulumi.BoolPtrInput
	LtsEnabled         pulumi.BoolPtrInput
	Name               pulumi.StringPtrInput
	ObsRetentionPeriod pulumi.IntPtrInput
	Region             pulumi.StringPtrInput
	Tags               pulumi.StringMapInput
	ValidateFile       pulumi.BoolPtrInput
}

func (CtsDataTrackerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ctsDataTrackerArgs)(nil)).Elem()
}

type CtsDataTrackerInput interface {
	pulumi.Input

	ToCtsDataTrackerOutput() CtsDataTrackerOutput
	ToCtsDataTrackerOutputWithContext(ctx context.Context) CtsDataTrackerOutput
}

func (*CtsDataTracker) ElementType() reflect.Type {
	return reflect.TypeOf((**CtsDataTracker)(nil)).Elem()
}

func (i *CtsDataTracker) ToCtsDataTrackerOutput() CtsDataTrackerOutput {
	return i.ToCtsDataTrackerOutputWithContext(context.Background())
}

func (i *CtsDataTracker) ToCtsDataTrackerOutputWithContext(ctx context.Context) CtsDataTrackerOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CtsDataTrackerOutput)
}

// CtsDataTrackerArrayInput is an input type that accepts CtsDataTrackerArray and CtsDataTrackerArrayOutput values.
// You can construct a concrete instance of `CtsDataTrackerArrayInput` via:
//
//	CtsDataTrackerArray{ CtsDataTrackerArgs{...} }
type CtsDataTrackerArrayInput interface {
	pulumi.Input

	ToCtsDataTrackerArrayOutput() CtsDataTrackerArrayOutput
	ToCtsDataTrackerArrayOutputWithContext(context.Context) CtsDataTrackerArrayOutput
}

type CtsDataTrackerArray []CtsDataTrackerInput

func (CtsDataTrackerArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*CtsDataTracker)(nil)).Elem()
}

func (i CtsDataTrackerArray) ToCtsDataTrackerArrayOutput() CtsDataTrackerArrayOutput {
	return i.ToCtsDataTrackerArrayOutputWithContext(context.Background())
}

func (i CtsDataTrackerArray) ToCtsDataTrackerArrayOutputWithContext(ctx context.Context) CtsDataTrackerArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CtsDataTrackerArrayOutput)
}

// CtsDataTrackerMapInput is an input type that accepts CtsDataTrackerMap and CtsDataTrackerMapOutput values.
// You can construct a concrete instance of `CtsDataTrackerMapInput` via:
//
//	CtsDataTrackerMap{ "key": CtsDataTrackerArgs{...} }
type CtsDataTrackerMapInput interface {
	pulumi.Input

	ToCtsDataTrackerMapOutput() CtsDataTrackerMapOutput
	ToCtsDataTrackerMapOutputWithContext(context.Context) CtsDataTrackerMapOutput
}

type CtsDataTrackerMap map[string]CtsDataTrackerInput

func (CtsDataTrackerMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*CtsDataTracker)(nil)).Elem()
}

func (i CtsDataTrackerMap) ToCtsDataTrackerMapOutput() CtsDataTrackerMapOutput {
	return i.ToCtsDataTrackerMapOutputWithContext(context.Background())
}

func (i CtsDataTrackerMap) ToCtsDataTrackerMapOutputWithContext(ctx context.Context) CtsDataTrackerMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CtsDataTrackerMapOutput)
}

type CtsDataTrackerOutput struct{ *pulumi.OutputState }

func (CtsDataTrackerOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CtsDataTracker)(nil)).Elem()
}

func (o CtsDataTrackerOutput) ToCtsDataTrackerOutput() CtsDataTrackerOutput {
	return o
}

func (o CtsDataTrackerOutput) ToCtsDataTrackerOutputWithContext(ctx context.Context) CtsDataTrackerOutput {
	return o
}

func (o CtsDataTrackerOutput) AgencyName() pulumi.StringOutput {
	return o.ApplyT(func(v *CtsDataTracker) pulumi.StringOutput { return v.AgencyName }).(pulumi.StringOutput)
}

func (o CtsDataTrackerOutput) BucketName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CtsDataTracker) pulumi.StringPtrOutput { return v.BucketName }).(pulumi.StringPtrOutput)
}

func (o CtsDataTrackerOutput) CompressType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CtsDataTracker) pulumi.StringPtrOutput { return v.CompressType }).(pulumi.StringPtrOutput)
}

func (o CtsDataTrackerOutput) CreateTime() pulumi.IntOutput {
	return o.ApplyT(func(v *CtsDataTracker) pulumi.IntOutput { return v.CreateTime }).(pulumi.IntOutput)
}

func (o CtsDataTrackerOutput) DataBucket() pulumi.StringOutput {
	return o.ApplyT(func(v *CtsDataTracker) pulumi.StringOutput { return v.DataBucket }).(pulumi.StringOutput)
}

func (o CtsDataTrackerOutput) DataOperations() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *CtsDataTracker) pulumi.StringArrayOutput { return v.DataOperations }).(pulumi.StringArrayOutput)
}

func (o CtsDataTrackerOutput) Detail() pulumi.StringOutput {
	return o.ApplyT(func(v *CtsDataTracker) pulumi.StringOutput { return v.Detail }).(pulumi.StringOutput)
}

func (o CtsDataTrackerOutput) DomainId() pulumi.StringOutput {
	return o.ApplyT(func(v *CtsDataTracker) pulumi.StringOutput { return v.DomainId }).(pulumi.StringOutput)
}

func (o CtsDataTrackerOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *CtsDataTracker) pulumi.BoolPtrOutput { return v.Enabled }).(pulumi.BoolPtrOutput)
}

func (o CtsDataTrackerOutput) FilePrefix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CtsDataTracker) pulumi.StringPtrOutput { return v.FilePrefix }).(pulumi.StringPtrOutput)
}

func (o CtsDataTrackerOutput) GroupId() pulumi.StringOutput {
	return o.ApplyT(func(v *CtsDataTracker) pulumi.StringOutput { return v.GroupId }).(pulumi.StringOutput)
}

func (o CtsDataTrackerOutput) IsAuthorizedBucket() pulumi.BoolOutput {
	return o.ApplyT(func(v *CtsDataTracker) pulumi.BoolOutput { return v.IsAuthorizedBucket }).(pulumi.BoolOutput)
}

func (o CtsDataTrackerOutput) IsSortByService() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *CtsDataTracker) pulumi.BoolPtrOutput { return v.IsSortByService }).(pulumi.BoolPtrOutput)
}

func (o CtsDataTrackerOutput) LogGroupName() pulumi.StringOutput {
	return o.ApplyT(func(v *CtsDataTracker) pulumi.StringOutput { return v.LogGroupName }).(pulumi.StringOutput)
}

func (o CtsDataTrackerOutput) LogTopicName() pulumi.StringOutput {
	return o.ApplyT(func(v *CtsDataTracker) pulumi.StringOutput { return v.LogTopicName }).(pulumi.StringOutput)
}

func (o CtsDataTrackerOutput) LtsEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *CtsDataTracker) pulumi.BoolPtrOutput { return v.LtsEnabled }).(pulumi.BoolPtrOutput)
}

func (o CtsDataTrackerOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *CtsDataTracker) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o CtsDataTrackerOutput) ObsRetentionPeriod() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *CtsDataTracker) pulumi.IntPtrOutput { return v.ObsRetentionPeriod }).(pulumi.IntPtrOutput)
}

func (o CtsDataTrackerOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v *CtsDataTracker) pulumi.StringOutput { return v.Region }).(pulumi.StringOutput)
}

func (o CtsDataTrackerOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *CtsDataTracker) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

func (o CtsDataTrackerOutput) StreamId() pulumi.StringOutput {
	return o.ApplyT(func(v *CtsDataTracker) pulumi.StringOutput { return v.StreamId }).(pulumi.StringOutput)
}

func (o CtsDataTrackerOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *CtsDataTracker) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o CtsDataTrackerOutput) TransferEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v *CtsDataTracker) pulumi.BoolOutput { return v.TransferEnabled }).(pulumi.BoolOutput)
}

func (o CtsDataTrackerOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *CtsDataTracker) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func (o CtsDataTrackerOutput) ValidateFile() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *CtsDataTracker) pulumi.BoolPtrOutput { return v.ValidateFile }).(pulumi.BoolPtrOutput)
}

type CtsDataTrackerArrayOutput struct{ *pulumi.OutputState }

func (CtsDataTrackerArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*CtsDataTracker)(nil)).Elem()
}

func (o CtsDataTrackerArrayOutput) ToCtsDataTrackerArrayOutput() CtsDataTrackerArrayOutput {
	return o
}

func (o CtsDataTrackerArrayOutput) ToCtsDataTrackerArrayOutputWithContext(ctx context.Context) CtsDataTrackerArrayOutput {
	return o
}

func (o CtsDataTrackerArrayOutput) Index(i pulumi.IntInput) CtsDataTrackerOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *CtsDataTracker {
		return vs[0].([]*CtsDataTracker)[vs[1].(int)]
	}).(CtsDataTrackerOutput)
}

type CtsDataTrackerMapOutput struct{ *pulumi.OutputState }

func (CtsDataTrackerMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*CtsDataTracker)(nil)).Elem()
}

func (o CtsDataTrackerMapOutput) ToCtsDataTrackerMapOutput() CtsDataTrackerMapOutput {
	return o
}

func (o CtsDataTrackerMapOutput) ToCtsDataTrackerMapOutputWithContext(ctx context.Context) CtsDataTrackerMapOutput {
	return o
}

func (o CtsDataTrackerMapOutput) MapIndex(k pulumi.StringInput) CtsDataTrackerOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *CtsDataTracker {
		return vs[0].(map[string]*CtsDataTracker)[vs[1].(string)]
	}).(CtsDataTrackerOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*CtsDataTrackerInput)(nil)).Elem(), &CtsDataTracker{})
	pulumi.RegisterInputType(reflect.TypeOf((*CtsDataTrackerArrayInput)(nil)).Elem(), CtsDataTrackerArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*CtsDataTrackerMapInput)(nil)).Elem(), CtsDataTrackerMap{})
	pulumi.RegisterOutputType(CtsDataTrackerOutput{})
	pulumi.RegisterOutputType(CtsDataTrackerArrayOutput{})
	pulumi.RegisterOutputType(CtsDataTrackerMapOutput{})
}
