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

type RdsMysqlBinlog struct {
	pulumi.CustomResourceState

	BinlogRetentionHours pulumi.IntOutput    `pulumi:"binlogRetentionHours"`
	InstanceId           pulumi.StringOutput `pulumi:"instanceId"`
	Region               pulumi.StringOutput `pulumi:"region"`
}

// NewRdsMysqlBinlog registers a new resource with the given unique name, arguments, and options.
func NewRdsMysqlBinlog(ctx *pulumi.Context,
	name string, args *RdsMysqlBinlogArgs, opts ...pulumi.ResourceOption) (*RdsMysqlBinlog, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.BinlogRetentionHours == nil {
		return nil, errors.New("invalid value for required argument 'BinlogRetentionHours'")
	}
	if args.InstanceId == nil {
		return nil, errors.New("invalid value for required argument 'InstanceId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource RdsMysqlBinlog
	err := ctx.RegisterResource("sbercloud:index/rdsMysqlBinlog:RdsMysqlBinlog", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRdsMysqlBinlog gets an existing RdsMysqlBinlog resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRdsMysqlBinlog(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RdsMysqlBinlogState, opts ...pulumi.ResourceOption) (*RdsMysqlBinlog, error) {
	var resource RdsMysqlBinlog
	err := ctx.ReadResource("sbercloud:index/rdsMysqlBinlog:RdsMysqlBinlog", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering RdsMysqlBinlog resources.
type rdsMysqlBinlogState struct {
	BinlogRetentionHours *int    `pulumi:"binlogRetentionHours"`
	InstanceId           *string `pulumi:"instanceId"`
	Region               *string `pulumi:"region"`
}

type RdsMysqlBinlogState struct {
	BinlogRetentionHours pulumi.IntPtrInput
	InstanceId           pulumi.StringPtrInput
	Region               pulumi.StringPtrInput
}

func (RdsMysqlBinlogState) ElementType() reflect.Type {
	return reflect.TypeOf((*rdsMysqlBinlogState)(nil)).Elem()
}

type rdsMysqlBinlogArgs struct {
	BinlogRetentionHours int     `pulumi:"binlogRetentionHours"`
	InstanceId           string  `pulumi:"instanceId"`
	Region               *string `pulumi:"region"`
}

// The set of arguments for constructing a RdsMysqlBinlog resource.
type RdsMysqlBinlogArgs struct {
	BinlogRetentionHours pulumi.IntInput
	InstanceId           pulumi.StringInput
	Region               pulumi.StringPtrInput
}

func (RdsMysqlBinlogArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*rdsMysqlBinlogArgs)(nil)).Elem()
}

type RdsMysqlBinlogInput interface {
	pulumi.Input

	ToRdsMysqlBinlogOutput() RdsMysqlBinlogOutput
	ToRdsMysqlBinlogOutputWithContext(ctx context.Context) RdsMysqlBinlogOutput
}

func (*RdsMysqlBinlog) ElementType() reflect.Type {
	return reflect.TypeOf((**RdsMysqlBinlog)(nil)).Elem()
}

func (i *RdsMysqlBinlog) ToRdsMysqlBinlogOutput() RdsMysqlBinlogOutput {
	return i.ToRdsMysqlBinlogOutputWithContext(context.Background())
}

func (i *RdsMysqlBinlog) ToRdsMysqlBinlogOutputWithContext(ctx context.Context) RdsMysqlBinlogOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RdsMysqlBinlogOutput)
}

// RdsMysqlBinlogArrayInput is an input type that accepts RdsMysqlBinlogArray and RdsMysqlBinlogArrayOutput values.
// You can construct a concrete instance of `RdsMysqlBinlogArrayInput` via:
//
//	RdsMysqlBinlogArray{ RdsMysqlBinlogArgs{...} }
type RdsMysqlBinlogArrayInput interface {
	pulumi.Input

	ToRdsMysqlBinlogArrayOutput() RdsMysqlBinlogArrayOutput
	ToRdsMysqlBinlogArrayOutputWithContext(context.Context) RdsMysqlBinlogArrayOutput
}

type RdsMysqlBinlogArray []RdsMysqlBinlogInput

func (RdsMysqlBinlogArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*RdsMysqlBinlog)(nil)).Elem()
}

func (i RdsMysqlBinlogArray) ToRdsMysqlBinlogArrayOutput() RdsMysqlBinlogArrayOutput {
	return i.ToRdsMysqlBinlogArrayOutputWithContext(context.Background())
}

func (i RdsMysqlBinlogArray) ToRdsMysqlBinlogArrayOutputWithContext(ctx context.Context) RdsMysqlBinlogArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RdsMysqlBinlogArrayOutput)
}

// RdsMysqlBinlogMapInput is an input type that accepts RdsMysqlBinlogMap and RdsMysqlBinlogMapOutput values.
// You can construct a concrete instance of `RdsMysqlBinlogMapInput` via:
//
//	RdsMysqlBinlogMap{ "key": RdsMysqlBinlogArgs{...} }
type RdsMysqlBinlogMapInput interface {
	pulumi.Input

	ToRdsMysqlBinlogMapOutput() RdsMysqlBinlogMapOutput
	ToRdsMysqlBinlogMapOutputWithContext(context.Context) RdsMysqlBinlogMapOutput
}

type RdsMysqlBinlogMap map[string]RdsMysqlBinlogInput

func (RdsMysqlBinlogMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*RdsMysqlBinlog)(nil)).Elem()
}

func (i RdsMysqlBinlogMap) ToRdsMysqlBinlogMapOutput() RdsMysqlBinlogMapOutput {
	return i.ToRdsMysqlBinlogMapOutputWithContext(context.Background())
}

func (i RdsMysqlBinlogMap) ToRdsMysqlBinlogMapOutputWithContext(ctx context.Context) RdsMysqlBinlogMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RdsMysqlBinlogMapOutput)
}

type RdsMysqlBinlogOutput struct{ *pulumi.OutputState }

func (RdsMysqlBinlogOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RdsMysqlBinlog)(nil)).Elem()
}

func (o RdsMysqlBinlogOutput) ToRdsMysqlBinlogOutput() RdsMysqlBinlogOutput {
	return o
}

func (o RdsMysqlBinlogOutput) ToRdsMysqlBinlogOutputWithContext(ctx context.Context) RdsMysqlBinlogOutput {
	return o
}

func (o RdsMysqlBinlogOutput) BinlogRetentionHours() pulumi.IntOutput {
	return o.ApplyT(func(v *RdsMysqlBinlog) pulumi.IntOutput { return v.BinlogRetentionHours }).(pulumi.IntOutput)
}

func (o RdsMysqlBinlogOutput) InstanceId() pulumi.StringOutput {
	return o.ApplyT(func(v *RdsMysqlBinlog) pulumi.StringOutput { return v.InstanceId }).(pulumi.StringOutput)
}

func (o RdsMysqlBinlogOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v *RdsMysqlBinlog) pulumi.StringOutput { return v.Region }).(pulumi.StringOutput)
}

type RdsMysqlBinlogArrayOutput struct{ *pulumi.OutputState }

func (RdsMysqlBinlogArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*RdsMysqlBinlog)(nil)).Elem()
}

func (o RdsMysqlBinlogArrayOutput) ToRdsMysqlBinlogArrayOutput() RdsMysqlBinlogArrayOutput {
	return o
}

func (o RdsMysqlBinlogArrayOutput) ToRdsMysqlBinlogArrayOutputWithContext(ctx context.Context) RdsMysqlBinlogArrayOutput {
	return o
}

func (o RdsMysqlBinlogArrayOutput) Index(i pulumi.IntInput) RdsMysqlBinlogOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *RdsMysqlBinlog {
		return vs[0].([]*RdsMysqlBinlog)[vs[1].(int)]
	}).(RdsMysqlBinlogOutput)
}

type RdsMysqlBinlogMapOutput struct{ *pulumi.OutputState }

func (RdsMysqlBinlogMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*RdsMysqlBinlog)(nil)).Elem()
}

func (o RdsMysqlBinlogMapOutput) ToRdsMysqlBinlogMapOutput() RdsMysqlBinlogMapOutput {
	return o
}

func (o RdsMysqlBinlogMapOutput) ToRdsMysqlBinlogMapOutputWithContext(ctx context.Context) RdsMysqlBinlogMapOutput {
	return o
}

func (o RdsMysqlBinlogMapOutput) MapIndex(k pulumi.StringInput) RdsMysqlBinlogOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *RdsMysqlBinlog {
		return vs[0].(map[string]*RdsMysqlBinlog)[vs[1].(string)]
	}).(RdsMysqlBinlogOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*RdsMysqlBinlogInput)(nil)).Elem(), &RdsMysqlBinlog{})
	pulumi.RegisterInputType(reflect.TypeOf((*RdsMysqlBinlogArrayInput)(nil)).Elem(), RdsMysqlBinlogArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*RdsMysqlBinlogMapInput)(nil)).Elem(), RdsMysqlBinlogMap{})
	pulumi.RegisterOutputType(RdsMysqlBinlogOutput{})
	pulumi.RegisterOutputType(RdsMysqlBinlogArrayOutput{})
	pulumi.RegisterOutputType(RdsMysqlBinlogMapOutput{})
}
