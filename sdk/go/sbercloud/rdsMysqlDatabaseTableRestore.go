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

type RdsMysqlDatabaseTableRestore struct {
	pulumi.CustomResourceState

	// Specifies the databases that will be restored.
	Databases RdsMysqlDatabaseTableRestoreDatabaseArrayOutput `pulumi:"databases"`
	// Specifies the ID of RDS MySQL instance.
	InstanceId pulumi.StringOutput `pulumi:"instanceId"`
	// Specifies whether to use fast restoration.
	IsFastRestore pulumi.BoolPtrOutput `pulumi:"isFastRestore"`
	Region        pulumi.StringOutput  `pulumi:"region"`
	// Specifies the tables that will be restored.
	RestoreTables RdsMysqlDatabaseTableRestoreRestoreTableArrayOutput `pulumi:"restoreTables"`
	// Specifies the restoration time point.
	RestoreTime pulumi.IntOutput `pulumi:"restoreTime"`
}

// NewRdsMysqlDatabaseTableRestore registers a new resource with the given unique name, arguments, and options.
func NewRdsMysqlDatabaseTableRestore(ctx *pulumi.Context,
	name string, args *RdsMysqlDatabaseTableRestoreArgs, opts ...pulumi.ResourceOption) (*RdsMysqlDatabaseTableRestore, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.InstanceId == nil {
		return nil, errors.New("invalid value for required argument 'InstanceId'")
	}
	if args.RestoreTime == nil {
		return nil, errors.New("invalid value for required argument 'RestoreTime'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource RdsMysqlDatabaseTableRestore
	err := ctx.RegisterResource("sbercloud:index/rdsMysqlDatabaseTableRestore:RdsMysqlDatabaseTableRestore", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRdsMysqlDatabaseTableRestore gets an existing RdsMysqlDatabaseTableRestore resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRdsMysqlDatabaseTableRestore(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RdsMysqlDatabaseTableRestoreState, opts ...pulumi.ResourceOption) (*RdsMysqlDatabaseTableRestore, error) {
	var resource RdsMysqlDatabaseTableRestore
	err := ctx.ReadResource("sbercloud:index/rdsMysqlDatabaseTableRestore:RdsMysqlDatabaseTableRestore", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering RdsMysqlDatabaseTableRestore resources.
type rdsMysqlDatabaseTableRestoreState struct {
	// Specifies the databases that will be restored.
	Databases []RdsMysqlDatabaseTableRestoreDatabase `pulumi:"databases"`
	// Specifies the ID of RDS MySQL instance.
	InstanceId *string `pulumi:"instanceId"`
	// Specifies whether to use fast restoration.
	IsFastRestore *bool   `pulumi:"isFastRestore"`
	Region        *string `pulumi:"region"`
	// Specifies the tables that will be restored.
	RestoreTables []RdsMysqlDatabaseTableRestoreRestoreTable `pulumi:"restoreTables"`
	// Specifies the restoration time point.
	RestoreTime *int `pulumi:"restoreTime"`
}

type RdsMysqlDatabaseTableRestoreState struct {
	// Specifies the databases that will be restored.
	Databases RdsMysqlDatabaseTableRestoreDatabaseArrayInput
	// Specifies the ID of RDS MySQL instance.
	InstanceId pulumi.StringPtrInput
	// Specifies whether to use fast restoration.
	IsFastRestore pulumi.BoolPtrInput
	Region        pulumi.StringPtrInput
	// Specifies the tables that will be restored.
	RestoreTables RdsMysqlDatabaseTableRestoreRestoreTableArrayInput
	// Specifies the restoration time point.
	RestoreTime pulumi.IntPtrInput
}

func (RdsMysqlDatabaseTableRestoreState) ElementType() reflect.Type {
	return reflect.TypeOf((*rdsMysqlDatabaseTableRestoreState)(nil)).Elem()
}

type rdsMysqlDatabaseTableRestoreArgs struct {
	// Specifies the databases that will be restored.
	Databases []RdsMysqlDatabaseTableRestoreDatabase `pulumi:"databases"`
	// Specifies the ID of RDS MySQL instance.
	InstanceId string `pulumi:"instanceId"`
	// Specifies whether to use fast restoration.
	IsFastRestore *bool   `pulumi:"isFastRestore"`
	Region        *string `pulumi:"region"`
	// Specifies the tables that will be restored.
	RestoreTables []RdsMysqlDatabaseTableRestoreRestoreTable `pulumi:"restoreTables"`
	// Specifies the restoration time point.
	RestoreTime int `pulumi:"restoreTime"`
}

// The set of arguments for constructing a RdsMysqlDatabaseTableRestore resource.
type RdsMysqlDatabaseTableRestoreArgs struct {
	// Specifies the databases that will be restored.
	Databases RdsMysqlDatabaseTableRestoreDatabaseArrayInput
	// Specifies the ID of RDS MySQL instance.
	InstanceId pulumi.StringInput
	// Specifies whether to use fast restoration.
	IsFastRestore pulumi.BoolPtrInput
	Region        pulumi.StringPtrInput
	// Specifies the tables that will be restored.
	RestoreTables RdsMysqlDatabaseTableRestoreRestoreTableArrayInput
	// Specifies the restoration time point.
	RestoreTime pulumi.IntInput
}

func (RdsMysqlDatabaseTableRestoreArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*rdsMysqlDatabaseTableRestoreArgs)(nil)).Elem()
}

type RdsMysqlDatabaseTableRestoreInput interface {
	pulumi.Input

	ToRdsMysqlDatabaseTableRestoreOutput() RdsMysqlDatabaseTableRestoreOutput
	ToRdsMysqlDatabaseTableRestoreOutputWithContext(ctx context.Context) RdsMysqlDatabaseTableRestoreOutput
}

func (*RdsMysqlDatabaseTableRestore) ElementType() reflect.Type {
	return reflect.TypeOf((**RdsMysqlDatabaseTableRestore)(nil)).Elem()
}

func (i *RdsMysqlDatabaseTableRestore) ToRdsMysqlDatabaseTableRestoreOutput() RdsMysqlDatabaseTableRestoreOutput {
	return i.ToRdsMysqlDatabaseTableRestoreOutputWithContext(context.Background())
}

func (i *RdsMysqlDatabaseTableRestore) ToRdsMysqlDatabaseTableRestoreOutputWithContext(ctx context.Context) RdsMysqlDatabaseTableRestoreOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RdsMysqlDatabaseTableRestoreOutput)
}

// RdsMysqlDatabaseTableRestoreArrayInput is an input type that accepts RdsMysqlDatabaseTableRestoreArray and RdsMysqlDatabaseTableRestoreArrayOutput values.
// You can construct a concrete instance of `RdsMysqlDatabaseTableRestoreArrayInput` via:
//
//	RdsMysqlDatabaseTableRestoreArray{ RdsMysqlDatabaseTableRestoreArgs{...} }
type RdsMysqlDatabaseTableRestoreArrayInput interface {
	pulumi.Input

	ToRdsMysqlDatabaseTableRestoreArrayOutput() RdsMysqlDatabaseTableRestoreArrayOutput
	ToRdsMysqlDatabaseTableRestoreArrayOutputWithContext(context.Context) RdsMysqlDatabaseTableRestoreArrayOutput
}

type RdsMysqlDatabaseTableRestoreArray []RdsMysqlDatabaseTableRestoreInput

func (RdsMysqlDatabaseTableRestoreArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*RdsMysqlDatabaseTableRestore)(nil)).Elem()
}

func (i RdsMysqlDatabaseTableRestoreArray) ToRdsMysqlDatabaseTableRestoreArrayOutput() RdsMysqlDatabaseTableRestoreArrayOutput {
	return i.ToRdsMysqlDatabaseTableRestoreArrayOutputWithContext(context.Background())
}

func (i RdsMysqlDatabaseTableRestoreArray) ToRdsMysqlDatabaseTableRestoreArrayOutputWithContext(ctx context.Context) RdsMysqlDatabaseTableRestoreArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RdsMysqlDatabaseTableRestoreArrayOutput)
}

// RdsMysqlDatabaseTableRestoreMapInput is an input type that accepts RdsMysqlDatabaseTableRestoreMap and RdsMysqlDatabaseTableRestoreMapOutput values.
// You can construct a concrete instance of `RdsMysqlDatabaseTableRestoreMapInput` via:
//
//	RdsMysqlDatabaseTableRestoreMap{ "key": RdsMysqlDatabaseTableRestoreArgs{...} }
type RdsMysqlDatabaseTableRestoreMapInput interface {
	pulumi.Input

	ToRdsMysqlDatabaseTableRestoreMapOutput() RdsMysqlDatabaseTableRestoreMapOutput
	ToRdsMysqlDatabaseTableRestoreMapOutputWithContext(context.Context) RdsMysqlDatabaseTableRestoreMapOutput
}

type RdsMysqlDatabaseTableRestoreMap map[string]RdsMysqlDatabaseTableRestoreInput

func (RdsMysqlDatabaseTableRestoreMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*RdsMysqlDatabaseTableRestore)(nil)).Elem()
}

func (i RdsMysqlDatabaseTableRestoreMap) ToRdsMysqlDatabaseTableRestoreMapOutput() RdsMysqlDatabaseTableRestoreMapOutput {
	return i.ToRdsMysqlDatabaseTableRestoreMapOutputWithContext(context.Background())
}

func (i RdsMysqlDatabaseTableRestoreMap) ToRdsMysqlDatabaseTableRestoreMapOutputWithContext(ctx context.Context) RdsMysqlDatabaseTableRestoreMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RdsMysqlDatabaseTableRestoreMapOutput)
}

type RdsMysqlDatabaseTableRestoreOutput struct{ *pulumi.OutputState }

func (RdsMysqlDatabaseTableRestoreOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RdsMysqlDatabaseTableRestore)(nil)).Elem()
}

func (o RdsMysqlDatabaseTableRestoreOutput) ToRdsMysqlDatabaseTableRestoreOutput() RdsMysqlDatabaseTableRestoreOutput {
	return o
}

func (o RdsMysqlDatabaseTableRestoreOutput) ToRdsMysqlDatabaseTableRestoreOutputWithContext(ctx context.Context) RdsMysqlDatabaseTableRestoreOutput {
	return o
}

// Specifies the databases that will be restored.
func (o RdsMysqlDatabaseTableRestoreOutput) Databases() RdsMysqlDatabaseTableRestoreDatabaseArrayOutput {
	return o.ApplyT(func(v *RdsMysqlDatabaseTableRestore) RdsMysqlDatabaseTableRestoreDatabaseArrayOutput {
		return v.Databases
	}).(RdsMysqlDatabaseTableRestoreDatabaseArrayOutput)
}

// Specifies the ID of RDS MySQL instance.
func (o RdsMysqlDatabaseTableRestoreOutput) InstanceId() pulumi.StringOutput {
	return o.ApplyT(func(v *RdsMysqlDatabaseTableRestore) pulumi.StringOutput { return v.InstanceId }).(pulumi.StringOutput)
}

// Specifies whether to use fast restoration.
func (o RdsMysqlDatabaseTableRestoreOutput) IsFastRestore() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *RdsMysqlDatabaseTableRestore) pulumi.BoolPtrOutput { return v.IsFastRestore }).(pulumi.BoolPtrOutput)
}

func (o RdsMysqlDatabaseTableRestoreOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v *RdsMysqlDatabaseTableRestore) pulumi.StringOutput { return v.Region }).(pulumi.StringOutput)
}

// Specifies the tables that will be restored.
func (o RdsMysqlDatabaseTableRestoreOutput) RestoreTables() RdsMysqlDatabaseTableRestoreRestoreTableArrayOutput {
	return o.ApplyT(func(v *RdsMysqlDatabaseTableRestore) RdsMysqlDatabaseTableRestoreRestoreTableArrayOutput {
		return v.RestoreTables
	}).(RdsMysqlDatabaseTableRestoreRestoreTableArrayOutput)
}

// Specifies the restoration time point.
func (o RdsMysqlDatabaseTableRestoreOutput) RestoreTime() pulumi.IntOutput {
	return o.ApplyT(func(v *RdsMysqlDatabaseTableRestore) pulumi.IntOutput { return v.RestoreTime }).(pulumi.IntOutput)
}

type RdsMysqlDatabaseTableRestoreArrayOutput struct{ *pulumi.OutputState }

func (RdsMysqlDatabaseTableRestoreArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*RdsMysqlDatabaseTableRestore)(nil)).Elem()
}

func (o RdsMysqlDatabaseTableRestoreArrayOutput) ToRdsMysqlDatabaseTableRestoreArrayOutput() RdsMysqlDatabaseTableRestoreArrayOutput {
	return o
}

func (o RdsMysqlDatabaseTableRestoreArrayOutput) ToRdsMysqlDatabaseTableRestoreArrayOutputWithContext(ctx context.Context) RdsMysqlDatabaseTableRestoreArrayOutput {
	return o
}

func (o RdsMysqlDatabaseTableRestoreArrayOutput) Index(i pulumi.IntInput) RdsMysqlDatabaseTableRestoreOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *RdsMysqlDatabaseTableRestore {
		return vs[0].([]*RdsMysqlDatabaseTableRestore)[vs[1].(int)]
	}).(RdsMysqlDatabaseTableRestoreOutput)
}

type RdsMysqlDatabaseTableRestoreMapOutput struct{ *pulumi.OutputState }

func (RdsMysqlDatabaseTableRestoreMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*RdsMysqlDatabaseTableRestore)(nil)).Elem()
}

func (o RdsMysqlDatabaseTableRestoreMapOutput) ToRdsMysqlDatabaseTableRestoreMapOutput() RdsMysqlDatabaseTableRestoreMapOutput {
	return o
}

func (o RdsMysqlDatabaseTableRestoreMapOutput) ToRdsMysqlDatabaseTableRestoreMapOutputWithContext(ctx context.Context) RdsMysqlDatabaseTableRestoreMapOutput {
	return o
}

func (o RdsMysqlDatabaseTableRestoreMapOutput) MapIndex(k pulumi.StringInput) RdsMysqlDatabaseTableRestoreOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *RdsMysqlDatabaseTableRestore {
		return vs[0].(map[string]*RdsMysqlDatabaseTableRestore)[vs[1].(string)]
	}).(RdsMysqlDatabaseTableRestoreOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*RdsMysqlDatabaseTableRestoreInput)(nil)).Elem(), &RdsMysqlDatabaseTableRestore{})
	pulumi.RegisterInputType(reflect.TypeOf((*RdsMysqlDatabaseTableRestoreArrayInput)(nil)).Elem(), RdsMysqlDatabaseTableRestoreArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*RdsMysqlDatabaseTableRestoreMapInput)(nil)).Elem(), RdsMysqlDatabaseTableRestoreMap{})
	pulumi.RegisterOutputType(RdsMysqlDatabaseTableRestoreOutput{})
	pulumi.RegisterOutputType(RdsMysqlDatabaseTableRestoreArrayOutput{})
	pulumi.RegisterOutputType(RdsMysqlDatabaseTableRestoreMapOutput{})
}
