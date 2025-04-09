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

type DcsBackup struct {
	pulumi.CustomResourceState

	// Specifies the format of the DCS instance backup.
	BackupFormat pulumi.StringOutput `pulumi:"backupFormat"`
	// Indicates the ID of the DCS instance backup.
	BackupId pulumi.StringOutput `pulumi:"backupId"`
	// Indicates the time when the backup task is created.
	BeginTime pulumi.StringOutput `pulumi:"beginTime"`
	// Specifies the description of DCS instance backup.
	Description pulumi.StringOutput `pulumi:"description"`
	// Indicates the time at which DCS instance backup is completed.
	EndTime pulumi.StringOutput `pulumi:"endTime"`
	// Specifies the ID of the DCS instance.
	InstanceId pulumi.StringOutput `pulumi:"instanceId"`
	// Indicates whether restoration is supported. Value Options: **TRUE**, **FALSE**.
	IsSupportRestore pulumi.StringOutput `pulumi:"isSupportRestore"`
	// Indicates the backup name.
	Name   pulumi.StringOutput `pulumi:"name"`
	Region pulumi.StringOutput `pulumi:"region"`
	// Indicates the size of the backup file (byte).
	Size pulumi.IntOutput `pulumi:"size"`
	// Indicates the backup status.
	Status pulumi.StringOutput `pulumi:"status"`
	// Indicates the backup type.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewDcsBackup registers a new resource with the given unique name, arguments, and options.
func NewDcsBackup(ctx *pulumi.Context,
	name string, args *DcsBackupArgs, opts ...pulumi.ResourceOption) (*DcsBackup, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.InstanceId == nil {
		return nil, errors.New("invalid value for required argument 'InstanceId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource DcsBackup
	err := ctx.RegisterResource("sbercloud:index/dcsBackup:DcsBackup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDcsBackup gets an existing DcsBackup resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDcsBackup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DcsBackupState, opts ...pulumi.ResourceOption) (*DcsBackup, error) {
	var resource DcsBackup
	err := ctx.ReadResource("sbercloud:index/dcsBackup:DcsBackup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DcsBackup resources.
type dcsBackupState struct {
	// Specifies the format of the DCS instance backup.
	BackupFormat *string `pulumi:"backupFormat"`
	// Indicates the ID of the DCS instance backup.
	BackupId *string `pulumi:"backupId"`
	// Indicates the time when the backup task is created.
	BeginTime *string `pulumi:"beginTime"`
	// Specifies the description of DCS instance backup.
	Description *string `pulumi:"description"`
	// Indicates the time at which DCS instance backup is completed.
	EndTime *string `pulumi:"endTime"`
	// Specifies the ID of the DCS instance.
	InstanceId *string `pulumi:"instanceId"`
	// Indicates whether restoration is supported. Value Options: **TRUE**, **FALSE**.
	IsSupportRestore *string `pulumi:"isSupportRestore"`
	// Indicates the backup name.
	Name   *string `pulumi:"name"`
	Region *string `pulumi:"region"`
	// Indicates the size of the backup file (byte).
	Size *int `pulumi:"size"`
	// Indicates the backup status.
	Status *string `pulumi:"status"`
	// Indicates the backup type.
	Type *string `pulumi:"type"`
}

type DcsBackupState struct {
	// Specifies the format of the DCS instance backup.
	BackupFormat pulumi.StringPtrInput
	// Indicates the ID of the DCS instance backup.
	BackupId pulumi.StringPtrInput
	// Indicates the time when the backup task is created.
	BeginTime pulumi.StringPtrInput
	// Specifies the description of DCS instance backup.
	Description pulumi.StringPtrInput
	// Indicates the time at which DCS instance backup is completed.
	EndTime pulumi.StringPtrInput
	// Specifies the ID of the DCS instance.
	InstanceId pulumi.StringPtrInput
	// Indicates whether restoration is supported. Value Options: **TRUE**, **FALSE**.
	IsSupportRestore pulumi.StringPtrInput
	// Indicates the backup name.
	Name   pulumi.StringPtrInput
	Region pulumi.StringPtrInput
	// Indicates the size of the backup file (byte).
	Size pulumi.IntPtrInput
	// Indicates the backup status.
	Status pulumi.StringPtrInput
	// Indicates the backup type.
	Type pulumi.StringPtrInput
}

func (DcsBackupState) ElementType() reflect.Type {
	return reflect.TypeOf((*dcsBackupState)(nil)).Elem()
}

type dcsBackupArgs struct {
	// Specifies the format of the DCS instance backup.
	BackupFormat *string `pulumi:"backupFormat"`
	// Specifies the description of DCS instance backup.
	Description *string `pulumi:"description"`
	// Specifies the ID of the DCS instance.
	InstanceId string  `pulumi:"instanceId"`
	Region     *string `pulumi:"region"`
}

// The set of arguments for constructing a DcsBackup resource.
type DcsBackupArgs struct {
	// Specifies the format of the DCS instance backup.
	BackupFormat pulumi.StringPtrInput
	// Specifies the description of DCS instance backup.
	Description pulumi.StringPtrInput
	// Specifies the ID of the DCS instance.
	InstanceId pulumi.StringInput
	Region     pulumi.StringPtrInput
}

func (DcsBackupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*dcsBackupArgs)(nil)).Elem()
}

type DcsBackupInput interface {
	pulumi.Input

	ToDcsBackupOutput() DcsBackupOutput
	ToDcsBackupOutputWithContext(ctx context.Context) DcsBackupOutput
}

func (*DcsBackup) ElementType() reflect.Type {
	return reflect.TypeOf((**DcsBackup)(nil)).Elem()
}

func (i *DcsBackup) ToDcsBackupOutput() DcsBackupOutput {
	return i.ToDcsBackupOutputWithContext(context.Background())
}

func (i *DcsBackup) ToDcsBackupOutputWithContext(ctx context.Context) DcsBackupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DcsBackupOutput)
}

// DcsBackupArrayInput is an input type that accepts DcsBackupArray and DcsBackupArrayOutput values.
// You can construct a concrete instance of `DcsBackupArrayInput` via:
//
//	DcsBackupArray{ DcsBackupArgs{...} }
type DcsBackupArrayInput interface {
	pulumi.Input

	ToDcsBackupArrayOutput() DcsBackupArrayOutput
	ToDcsBackupArrayOutputWithContext(context.Context) DcsBackupArrayOutput
}

type DcsBackupArray []DcsBackupInput

func (DcsBackupArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DcsBackup)(nil)).Elem()
}

func (i DcsBackupArray) ToDcsBackupArrayOutput() DcsBackupArrayOutput {
	return i.ToDcsBackupArrayOutputWithContext(context.Background())
}

func (i DcsBackupArray) ToDcsBackupArrayOutputWithContext(ctx context.Context) DcsBackupArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DcsBackupArrayOutput)
}

// DcsBackupMapInput is an input type that accepts DcsBackupMap and DcsBackupMapOutput values.
// You can construct a concrete instance of `DcsBackupMapInput` via:
//
//	DcsBackupMap{ "key": DcsBackupArgs{...} }
type DcsBackupMapInput interface {
	pulumi.Input

	ToDcsBackupMapOutput() DcsBackupMapOutput
	ToDcsBackupMapOutputWithContext(context.Context) DcsBackupMapOutput
}

type DcsBackupMap map[string]DcsBackupInput

func (DcsBackupMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DcsBackup)(nil)).Elem()
}

func (i DcsBackupMap) ToDcsBackupMapOutput() DcsBackupMapOutput {
	return i.ToDcsBackupMapOutputWithContext(context.Background())
}

func (i DcsBackupMap) ToDcsBackupMapOutputWithContext(ctx context.Context) DcsBackupMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DcsBackupMapOutput)
}

type DcsBackupOutput struct{ *pulumi.OutputState }

func (DcsBackupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DcsBackup)(nil)).Elem()
}

func (o DcsBackupOutput) ToDcsBackupOutput() DcsBackupOutput {
	return o
}

func (o DcsBackupOutput) ToDcsBackupOutputWithContext(ctx context.Context) DcsBackupOutput {
	return o
}

// Specifies the format of the DCS instance backup.
func (o DcsBackupOutput) BackupFormat() pulumi.StringOutput {
	return o.ApplyT(func(v *DcsBackup) pulumi.StringOutput { return v.BackupFormat }).(pulumi.StringOutput)
}

// Indicates the ID of the DCS instance backup.
func (o DcsBackupOutput) BackupId() pulumi.StringOutput {
	return o.ApplyT(func(v *DcsBackup) pulumi.StringOutput { return v.BackupId }).(pulumi.StringOutput)
}

// Indicates the time when the backup task is created.
func (o DcsBackupOutput) BeginTime() pulumi.StringOutput {
	return o.ApplyT(func(v *DcsBackup) pulumi.StringOutput { return v.BeginTime }).(pulumi.StringOutput)
}

// Specifies the description of DCS instance backup.
func (o DcsBackupOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v *DcsBackup) pulumi.StringOutput { return v.Description }).(pulumi.StringOutput)
}

// Indicates the time at which DCS instance backup is completed.
func (o DcsBackupOutput) EndTime() pulumi.StringOutput {
	return o.ApplyT(func(v *DcsBackup) pulumi.StringOutput { return v.EndTime }).(pulumi.StringOutput)
}

// Specifies the ID of the DCS instance.
func (o DcsBackupOutput) InstanceId() pulumi.StringOutput {
	return o.ApplyT(func(v *DcsBackup) pulumi.StringOutput { return v.InstanceId }).(pulumi.StringOutput)
}

// Indicates whether restoration is supported. Value Options: **TRUE**, **FALSE**.
func (o DcsBackupOutput) IsSupportRestore() pulumi.StringOutput {
	return o.ApplyT(func(v *DcsBackup) pulumi.StringOutput { return v.IsSupportRestore }).(pulumi.StringOutput)
}

// Indicates the backup name.
func (o DcsBackupOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *DcsBackup) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o DcsBackupOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v *DcsBackup) pulumi.StringOutput { return v.Region }).(pulumi.StringOutput)
}

// Indicates the size of the backup file (byte).
func (o DcsBackupOutput) Size() pulumi.IntOutput {
	return o.ApplyT(func(v *DcsBackup) pulumi.IntOutput { return v.Size }).(pulumi.IntOutput)
}

// Indicates the backup status.
func (o DcsBackupOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *DcsBackup) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

// Indicates the backup type.
func (o DcsBackupOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *DcsBackup) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

type DcsBackupArrayOutput struct{ *pulumi.OutputState }

func (DcsBackupArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DcsBackup)(nil)).Elem()
}

func (o DcsBackupArrayOutput) ToDcsBackupArrayOutput() DcsBackupArrayOutput {
	return o
}

func (o DcsBackupArrayOutput) ToDcsBackupArrayOutputWithContext(ctx context.Context) DcsBackupArrayOutput {
	return o
}

func (o DcsBackupArrayOutput) Index(i pulumi.IntInput) DcsBackupOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *DcsBackup {
		return vs[0].([]*DcsBackup)[vs[1].(int)]
	}).(DcsBackupOutput)
}

type DcsBackupMapOutput struct{ *pulumi.OutputState }

func (DcsBackupMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DcsBackup)(nil)).Elem()
}

func (o DcsBackupMapOutput) ToDcsBackupMapOutput() DcsBackupMapOutput {
	return o
}

func (o DcsBackupMapOutput) ToDcsBackupMapOutputWithContext(ctx context.Context) DcsBackupMapOutput {
	return o
}

func (o DcsBackupMapOutput) MapIndex(k pulumi.StringInput) DcsBackupOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *DcsBackup {
		return vs[0].(map[string]*DcsBackup)[vs[1].(string)]
	}).(DcsBackupOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DcsBackupInput)(nil)).Elem(), &DcsBackup{})
	pulumi.RegisterInputType(reflect.TypeOf((*DcsBackupArrayInput)(nil)).Elem(), DcsBackupArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*DcsBackupMapInput)(nil)).Elem(), DcsBackupMap{})
	pulumi.RegisterOutputType(DcsBackupOutput{})
	pulumi.RegisterOutputType(DcsBackupArrayOutput{})
	pulumi.RegisterOutputType(DcsBackupMapOutput{})
}
