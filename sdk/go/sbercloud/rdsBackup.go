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

type RdsBackup struct {
	pulumi.CustomResourceState

	// Whether a DDM instance has been associated.
	AssociatedWithDdm pulumi.BoolOutput `pulumi:"associatedWithDdm"`
	// Backup start time in the "yyyy-mm-ddThh:mm:ssZ" format.
	BeginTime pulumi.StringOutput `pulumi:"beginTime"`
	// List of self-built Microsoft SQL Server databases that are partially backed up.
	Databases RdsBackupDatabaseArrayOutput `pulumi:"databases"`
	// The description about the backup.
	Description pulumi.StringOutput `pulumi:"description"`
	// Backup end time in the "yyyy-mm-ddThh:mm:ssZ" format.
	EndTime pulumi.StringOutput `pulumi:"endTime"`
	// Instance ID.
	InstanceId pulumi.StringOutput `pulumi:"instanceId"`
	// Backup name.
	Name   pulumi.StringOutput `pulumi:"name"`
	Region pulumi.StringOutput `pulumi:"region"`
	// Backup size in KB.
	Size pulumi.IntOutput `pulumi:"size"`
	// Backup status.
	Status pulumi.StringOutput `pulumi:"status"`
}

// NewRdsBackup registers a new resource with the given unique name, arguments, and options.
func NewRdsBackup(ctx *pulumi.Context,
	name string, args *RdsBackupArgs, opts ...pulumi.ResourceOption) (*RdsBackup, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.InstanceId == nil {
		return nil, errors.New("invalid value for required argument 'InstanceId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource RdsBackup
	err := ctx.RegisterResource("sbercloud:index/rdsBackup:RdsBackup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRdsBackup gets an existing RdsBackup resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRdsBackup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RdsBackupState, opts ...pulumi.ResourceOption) (*RdsBackup, error) {
	var resource RdsBackup
	err := ctx.ReadResource("sbercloud:index/rdsBackup:RdsBackup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering RdsBackup resources.
type rdsBackupState struct {
	// Whether a DDM instance has been associated.
	AssociatedWithDdm *bool `pulumi:"associatedWithDdm"`
	// Backup start time in the "yyyy-mm-ddThh:mm:ssZ" format.
	BeginTime *string `pulumi:"beginTime"`
	// List of self-built Microsoft SQL Server databases that are partially backed up.
	Databases []RdsBackupDatabase `pulumi:"databases"`
	// The description about the backup.
	Description *string `pulumi:"description"`
	// Backup end time in the "yyyy-mm-ddThh:mm:ssZ" format.
	EndTime *string `pulumi:"endTime"`
	// Instance ID.
	InstanceId *string `pulumi:"instanceId"`
	// Backup name.
	Name   *string `pulumi:"name"`
	Region *string `pulumi:"region"`
	// Backup size in KB.
	Size *int `pulumi:"size"`
	// Backup status.
	Status *string `pulumi:"status"`
}

type RdsBackupState struct {
	// Whether a DDM instance has been associated.
	AssociatedWithDdm pulumi.BoolPtrInput
	// Backup start time in the "yyyy-mm-ddThh:mm:ssZ" format.
	BeginTime pulumi.StringPtrInput
	// List of self-built Microsoft SQL Server databases that are partially backed up.
	Databases RdsBackupDatabaseArrayInput
	// The description about the backup.
	Description pulumi.StringPtrInput
	// Backup end time in the "yyyy-mm-ddThh:mm:ssZ" format.
	EndTime pulumi.StringPtrInput
	// Instance ID.
	InstanceId pulumi.StringPtrInput
	// Backup name.
	Name   pulumi.StringPtrInput
	Region pulumi.StringPtrInput
	// Backup size in KB.
	Size pulumi.IntPtrInput
	// Backup status.
	Status pulumi.StringPtrInput
}

func (RdsBackupState) ElementType() reflect.Type {
	return reflect.TypeOf((*rdsBackupState)(nil)).Elem()
}

type rdsBackupArgs struct {
	// List of self-built Microsoft SQL Server databases that are partially backed up.
	Databases []RdsBackupDatabase `pulumi:"databases"`
	// The description about the backup.
	Description *string `pulumi:"description"`
	// Instance ID.
	InstanceId string `pulumi:"instanceId"`
	// Backup name.
	Name   *string `pulumi:"name"`
	Region *string `pulumi:"region"`
}

// The set of arguments for constructing a RdsBackup resource.
type RdsBackupArgs struct {
	// List of self-built Microsoft SQL Server databases that are partially backed up.
	Databases RdsBackupDatabaseArrayInput
	// The description about the backup.
	Description pulumi.StringPtrInput
	// Instance ID.
	InstanceId pulumi.StringInput
	// Backup name.
	Name   pulumi.StringPtrInput
	Region pulumi.StringPtrInput
}

func (RdsBackupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*rdsBackupArgs)(nil)).Elem()
}

type RdsBackupInput interface {
	pulumi.Input

	ToRdsBackupOutput() RdsBackupOutput
	ToRdsBackupOutputWithContext(ctx context.Context) RdsBackupOutput
}

func (*RdsBackup) ElementType() reflect.Type {
	return reflect.TypeOf((**RdsBackup)(nil)).Elem()
}

func (i *RdsBackup) ToRdsBackupOutput() RdsBackupOutput {
	return i.ToRdsBackupOutputWithContext(context.Background())
}

func (i *RdsBackup) ToRdsBackupOutputWithContext(ctx context.Context) RdsBackupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RdsBackupOutput)
}

// RdsBackupArrayInput is an input type that accepts RdsBackupArray and RdsBackupArrayOutput values.
// You can construct a concrete instance of `RdsBackupArrayInput` via:
//
//	RdsBackupArray{ RdsBackupArgs{...} }
type RdsBackupArrayInput interface {
	pulumi.Input

	ToRdsBackupArrayOutput() RdsBackupArrayOutput
	ToRdsBackupArrayOutputWithContext(context.Context) RdsBackupArrayOutput
}

type RdsBackupArray []RdsBackupInput

func (RdsBackupArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*RdsBackup)(nil)).Elem()
}

func (i RdsBackupArray) ToRdsBackupArrayOutput() RdsBackupArrayOutput {
	return i.ToRdsBackupArrayOutputWithContext(context.Background())
}

func (i RdsBackupArray) ToRdsBackupArrayOutputWithContext(ctx context.Context) RdsBackupArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RdsBackupArrayOutput)
}

// RdsBackupMapInput is an input type that accepts RdsBackupMap and RdsBackupMapOutput values.
// You can construct a concrete instance of `RdsBackupMapInput` via:
//
//	RdsBackupMap{ "key": RdsBackupArgs{...} }
type RdsBackupMapInput interface {
	pulumi.Input

	ToRdsBackupMapOutput() RdsBackupMapOutput
	ToRdsBackupMapOutputWithContext(context.Context) RdsBackupMapOutput
}

type RdsBackupMap map[string]RdsBackupInput

func (RdsBackupMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*RdsBackup)(nil)).Elem()
}

func (i RdsBackupMap) ToRdsBackupMapOutput() RdsBackupMapOutput {
	return i.ToRdsBackupMapOutputWithContext(context.Background())
}

func (i RdsBackupMap) ToRdsBackupMapOutputWithContext(ctx context.Context) RdsBackupMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RdsBackupMapOutput)
}

type RdsBackupOutput struct{ *pulumi.OutputState }

func (RdsBackupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RdsBackup)(nil)).Elem()
}

func (o RdsBackupOutput) ToRdsBackupOutput() RdsBackupOutput {
	return o
}

func (o RdsBackupOutput) ToRdsBackupOutputWithContext(ctx context.Context) RdsBackupOutput {
	return o
}

// Whether a DDM instance has been associated.
func (o RdsBackupOutput) AssociatedWithDdm() pulumi.BoolOutput {
	return o.ApplyT(func(v *RdsBackup) pulumi.BoolOutput { return v.AssociatedWithDdm }).(pulumi.BoolOutput)
}

// Backup start time in the "yyyy-mm-ddThh:mm:ssZ" format.
func (o RdsBackupOutput) BeginTime() pulumi.StringOutput {
	return o.ApplyT(func(v *RdsBackup) pulumi.StringOutput { return v.BeginTime }).(pulumi.StringOutput)
}

// List of self-built Microsoft SQL Server databases that are partially backed up.
func (o RdsBackupOutput) Databases() RdsBackupDatabaseArrayOutput {
	return o.ApplyT(func(v *RdsBackup) RdsBackupDatabaseArrayOutput { return v.Databases }).(RdsBackupDatabaseArrayOutput)
}

// The description about the backup.
func (o RdsBackupOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v *RdsBackup) pulumi.StringOutput { return v.Description }).(pulumi.StringOutput)
}

// Backup end time in the "yyyy-mm-ddThh:mm:ssZ" format.
func (o RdsBackupOutput) EndTime() pulumi.StringOutput {
	return o.ApplyT(func(v *RdsBackup) pulumi.StringOutput { return v.EndTime }).(pulumi.StringOutput)
}

// Instance ID.
func (o RdsBackupOutput) InstanceId() pulumi.StringOutput {
	return o.ApplyT(func(v *RdsBackup) pulumi.StringOutput { return v.InstanceId }).(pulumi.StringOutput)
}

// Backup name.
func (o RdsBackupOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *RdsBackup) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o RdsBackupOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v *RdsBackup) pulumi.StringOutput { return v.Region }).(pulumi.StringOutput)
}

// Backup size in KB.
func (o RdsBackupOutput) Size() pulumi.IntOutput {
	return o.ApplyT(func(v *RdsBackup) pulumi.IntOutput { return v.Size }).(pulumi.IntOutput)
}

// Backup status.
func (o RdsBackupOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *RdsBackup) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

type RdsBackupArrayOutput struct{ *pulumi.OutputState }

func (RdsBackupArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*RdsBackup)(nil)).Elem()
}

func (o RdsBackupArrayOutput) ToRdsBackupArrayOutput() RdsBackupArrayOutput {
	return o
}

func (o RdsBackupArrayOutput) ToRdsBackupArrayOutputWithContext(ctx context.Context) RdsBackupArrayOutput {
	return o
}

func (o RdsBackupArrayOutput) Index(i pulumi.IntInput) RdsBackupOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *RdsBackup {
		return vs[0].([]*RdsBackup)[vs[1].(int)]
	}).(RdsBackupOutput)
}

type RdsBackupMapOutput struct{ *pulumi.OutputState }

func (RdsBackupMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*RdsBackup)(nil)).Elem()
}

func (o RdsBackupMapOutput) ToRdsBackupMapOutput() RdsBackupMapOutput {
	return o
}

func (o RdsBackupMapOutput) ToRdsBackupMapOutputWithContext(ctx context.Context) RdsBackupMapOutput {
	return o
}

func (o RdsBackupMapOutput) MapIndex(k pulumi.StringInput) RdsBackupOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *RdsBackup {
		return vs[0].(map[string]*RdsBackup)[vs[1].(string)]
	}).(RdsBackupOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*RdsBackupInput)(nil)).Elem(), &RdsBackup{})
	pulumi.RegisterInputType(reflect.TypeOf((*RdsBackupArrayInput)(nil)).Elem(), RdsBackupArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*RdsBackupMapInput)(nil)).Elem(), RdsBackupMap{})
	pulumi.RegisterOutputType(RdsBackupOutput{})
	pulumi.RegisterOutputType(RdsBackupArrayOutput{})
	pulumi.RegisterOutputType(RdsBackupMapOutput{})
}
