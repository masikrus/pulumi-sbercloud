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

type LtsGroup struct {
	pulumi.CustomResourceState

	CreatedAt pulumi.StringOutput `pulumi:"createdAt"`
	// The enterprise project ID to which the log group belongs.
	EnterpriseProjectId pulumi.StringOutput    `pulumi:"enterpriseProjectId"`
	GroupName           pulumi.StringOutput    `pulumi:"groupName"`
	Region              pulumi.StringOutput    `pulumi:"region"`
	Tags                pulumi.StringMapOutput `pulumi:"tags"`
	TtlInDays           pulumi.IntOutput       `pulumi:"ttlInDays"`
}

// NewLtsGroup registers a new resource with the given unique name, arguments, and options.
func NewLtsGroup(ctx *pulumi.Context,
	name string, args *LtsGroupArgs, opts ...pulumi.ResourceOption) (*LtsGroup, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.GroupName == nil {
		return nil, errors.New("invalid value for required argument 'GroupName'")
	}
	if args.TtlInDays == nil {
		return nil, errors.New("invalid value for required argument 'TtlInDays'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource LtsGroup
	err := ctx.RegisterResource("sbercloud:index/ltsGroup:LtsGroup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetLtsGroup gets an existing LtsGroup resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetLtsGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *LtsGroupState, opts ...pulumi.ResourceOption) (*LtsGroup, error) {
	var resource LtsGroup
	err := ctx.ReadResource("sbercloud:index/ltsGroup:LtsGroup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering LtsGroup resources.
type ltsGroupState struct {
	CreatedAt *string `pulumi:"createdAt"`
	// The enterprise project ID to which the log group belongs.
	EnterpriseProjectId *string           `pulumi:"enterpriseProjectId"`
	GroupName           *string           `pulumi:"groupName"`
	Region              *string           `pulumi:"region"`
	Tags                map[string]string `pulumi:"tags"`
	TtlInDays           *int              `pulumi:"ttlInDays"`
}

type LtsGroupState struct {
	CreatedAt pulumi.StringPtrInput
	// The enterprise project ID to which the log group belongs.
	EnterpriseProjectId pulumi.StringPtrInput
	GroupName           pulumi.StringPtrInput
	Region              pulumi.StringPtrInput
	Tags                pulumi.StringMapInput
	TtlInDays           pulumi.IntPtrInput
}

func (LtsGroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*ltsGroupState)(nil)).Elem()
}

type ltsGroupArgs struct {
	// The enterprise project ID to which the log group belongs.
	EnterpriseProjectId *string           `pulumi:"enterpriseProjectId"`
	GroupName           string            `pulumi:"groupName"`
	Region              *string           `pulumi:"region"`
	Tags                map[string]string `pulumi:"tags"`
	TtlInDays           int               `pulumi:"ttlInDays"`
}

// The set of arguments for constructing a LtsGroup resource.
type LtsGroupArgs struct {
	// The enterprise project ID to which the log group belongs.
	EnterpriseProjectId pulumi.StringPtrInput
	GroupName           pulumi.StringInput
	Region              pulumi.StringPtrInput
	Tags                pulumi.StringMapInput
	TtlInDays           pulumi.IntInput
}

func (LtsGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ltsGroupArgs)(nil)).Elem()
}

type LtsGroupInput interface {
	pulumi.Input

	ToLtsGroupOutput() LtsGroupOutput
	ToLtsGroupOutputWithContext(ctx context.Context) LtsGroupOutput
}

func (*LtsGroup) ElementType() reflect.Type {
	return reflect.TypeOf((**LtsGroup)(nil)).Elem()
}

func (i *LtsGroup) ToLtsGroupOutput() LtsGroupOutput {
	return i.ToLtsGroupOutputWithContext(context.Background())
}

func (i *LtsGroup) ToLtsGroupOutputWithContext(ctx context.Context) LtsGroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LtsGroupOutput)
}

// LtsGroupArrayInput is an input type that accepts LtsGroupArray and LtsGroupArrayOutput values.
// You can construct a concrete instance of `LtsGroupArrayInput` via:
//
//	LtsGroupArray{ LtsGroupArgs{...} }
type LtsGroupArrayInput interface {
	pulumi.Input

	ToLtsGroupArrayOutput() LtsGroupArrayOutput
	ToLtsGroupArrayOutputWithContext(context.Context) LtsGroupArrayOutput
}

type LtsGroupArray []LtsGroupInput

func (LtsGroupArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*LtsGroup)(nil)).Elem()
}

func (i LtsGroupArray) ToLtsGroupArrayOutput() LtsGroupArrayOutput {
	return i.ToLtsGroupArrayOutputWithContext(context.Background())
}

func (i LtsGroupArray) ToLtsGroupArrayOutputWithContext(ctx context.Context) LtsGroupArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LtsGroupArrayOutput)
}

// LtsGroupMapInput is an input type that accepts LtsGroupMap and LtsGroupMapOutput values.
// You can construct a concrete instance of `LtsGroupMapInput` via:
//
//	LtsGroupMap{ "key": LtsGroupArgs{...} }
type LtsGroupMapInput interface {
	pulumi.Input

	ToLtsGroupMapOutput() LtsGroupMapOutput
	ToLtsGroupMapOutputWithContext(context.Context) LtsGroupMapOutput
}

type LtsGroupMap map[string]LtsGroupInput

func (LtsGroupMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*LtsGroup)(nil)).Elem()
}

func (i LtsGroupMap) ToLtsGroupMapOutput() LtsGroupMapOutput {
	return i.ToLtsGroupMapOutputWithContext(context.Background())
}

func (i LtsGroupMap) ToLtsGroupMapOutputWithContext(ctx context.Context) LtsGroupMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LtsGroupMapOutput)
}

type LtsGroupOutput struct{ *pulumi.OutputState }

func (LtsGroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LtsGroup)(nil)).Elem()
}

func (o LtsGroupOutput) ToLtsGroupOutput() LtsGroupOutput {
	return o
}

func (o LtsGroupOutput) ToLtsGroupOutputWithContext(ctx context.Context) LtsGroupOutput {
	return o
}

func (o LtsGroupOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *LtsGroup) pulumi.StringOutput { return v.CreatedAt }).(pulumi.StringOutput)
}

// The enterprise project ID to which the log group belongs.
func (o LtsGroupOutput) EnterpriseProjectId() pulumi.StringOutput {
	return o.ApplyT(func(v *LtsGroup) pulumi.StringOutput { return v.EnterpriseProjectId }).(pulumi.StringOutput)
}

func (o LtsGroupOutput) GroupName() pulumi.StringOutput {
	return o.ApplyT(func(v *LtsGroup) pulumi.StringOutput { return v.GroupName }).(pulumi.StringOutput)
}

func (o LtsGroupOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v *LtsGroup) pulumi.StringOutput { return v.Region }).(pulumi.StringOutput)
}

func (o LtsGroupOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *LtsGroup) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LtsGroupOutput) TtlInDays() pulumi.IntOutput {
	return o.ApplyT(func(v *LtsGroup) pulumi.IntOutput { return v.TtlInDays }).(pulumi.IntOutput)
}

type LtsGroupArrayOutput struct{ *pulumi.OutputState }

func (LtsGroupArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*LtsGroup)(nil)).Elem()
}

func (o LtsGroupArrayOutput) ToLtsGroupArrayOutput() LtsGroupArrayOutput {
	return o
}

func (o LtsGroupArrayOutput) ToLtsGroupArrayOutputWithContext(ctx context.Context) LtsGroupArrayOutput {
	return o
}

func (o LtsGroupArrayOutput) Index(i pulumi.IntInput) LtsGroupOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *LtsGroup {
		return vs[0].([]*LtsGroup)[vs[1].(int)]
	}).(LtsGroupOutput)
}

type LtsGroupMapOutput struct{ *pulumi.OutputState }

func (LtsGroupMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*LtsGroup)(nil)).Elem()
}

func (o LtsGroupMapOutput) ToLtsGroupMapOutput() LtsGroupMapOutput {
	return o
}

func (o LtsGroupMapOutput) ToLtsGroupMapOutputWithContext(ctx context.Context) LtsGroupMapOutput {
	return o
}

func (o LtsGroupMapOutput) MapIndex(k pulumi.StringInput) LtsGroupOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *LtsGroup {
		return vs[0].(map[string]*LtsGroup)[vs[1].(string)]
	}).(LtsGroupOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*LtsGroupInput)(nil)).Elem(), &LtsGroup{})
	pulumi.RegisterInputType(reflect.TypeOf((*LtsGroupArrayInput)(nil)).Elem(), LtsGroupArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*LtsGroupMapInput)(nil)).Elem(), LtsGroupMap{})
	pulumi.RegisterOutputType(LtsGroupOutput{})
	pulumi.RegisterOutputType(LtsGroupArrayOutput{})
	pulumi.RegisterOutputType(LtsGroupMapOutput{})
}
