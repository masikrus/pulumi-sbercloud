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

type LbPool struct {
	pulumi.CustomResourceState

	// Deprecated: this field is deprecated
	AdminStateUp     pulumi.BoolPtrOutput         `pulumi:"adminStateUp"`
	Description      pulumi.StringPtrOutput       `pulumi:"description"`
	LbMethod         pulumi.StringOutput          `pulumi:"lbMethod"`
	ListenerId       pulumi.StringOutput          `pulumi:"listenerId"`
	LoadbalancerId   pulumi.StringOutput          `pulumi:"loadbalancerId"`
	MonitorId        pulumi.StringOutput          `pulumi:"monitorId"`
	Name             pulumi.StringOutput          `pulumi:"name"`
	Persistences     LbPoolPersistenceArrayOutput `pulumi:"persistences"`
	ProtectionReason pulumi.StringPtrOutput       `pulumi:"protectionReason"`
	ProtectionStatus pulumi.StringOutput          `pulumi:"protectionStatus"`
	Protocol         pulumi.StringOutput          `pulumi:"protocol"`
	Region           pulumi.StringOutput          `pulumi:"region"`
	// Deprecated: tenant_id is deprecated
	TenantId pulumi.StringOutput `pulumi:"tenantId"`
}

// NewLbPool registers a new resource with the given unique name, arguments, and options.
func NewLbPool(ctx *pulumi.Context,
	name string, args *LbPoolArgs, opts ...pulumi.ResourceOption) (*LbPool, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.LbMethod == nil {
		return nil, errors.New("invalid value for required argument 'LbMethod'")
	}
	if args.Protocol == nil {
		return nil, errors.New("invalid value for required argument 'Protocol'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource LbPool
	err := ctx.RegisterResource("sbercloud:index/lbPool:LbPool", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetLbPool gets an existing LbPool resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetLbPool(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *LbPoolState, opts ...pulumi.ResourceOption) (*LbPool, error) {
	var resource LbPool
	err := ctx.ReadResource("sbercloud:index/lbPool:LbPool", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering LbPool resources.
type lbPoolState struct {
	// Deprecated: this field is deprecated
	AdminStateUp     *bool               `pulumi:"adminStateUp"`
	Description      *string             `pulumi:"description"`
	LbMethod         *string             `pulumi:"lbMethod"`
	ListenerId       *string             `pulumi:"listenerId"`
	LoadbalancerId   *string             `pulumi:"loadbalancerId"`
	MonitorId        *string             `pulumi:"monitorId"`
	Name             *string             `pulumi:"name"`
	Persistences     []LbPoolPersistence `pulumi:"persistences"`
	ProtectionReason *string             `pulumi:"protectionReason"`
	ProtectionStatus *string             `pulumi:"protectionStatus"`
	Protocol         *string             `pulumi:"protocol"`
	Region           *string             `pulumi:"region"`
	// Deprecated: tenant_id is deprecated
	TenantId *string `pulumi:"tenantId"`
}

type LbPoolState struct {
	// Deprecated: this field is deprecated
	AdminStateUp     pulumi.BoolPtrInput
	Description      pulumi.StringPtrInput
	LbMethod         pulumi.StringPtrInput
	ListenerId       pulumi.StringPtrInput
	LoadbalancerId   pulumi.StringPtrInput
	MonitorId        pulumi.StringPtrInput
	Name             pulumi.StringPtrInput
	Persistences     LbPoolPersistenceArrayInput
	ProtectionReason pulumi.StringPtrInput
	ProtectionStatus pulumi.StringPtrInput
	Protocol         pulumi.StringPtrInput
	Region           pulumi.StringPtrInput
	// Deprecated: tenant_id is deprecated
	TenantId pulumi.StringPtrInput
}

func (LbPoolState) ElementType() reflect.Type {
	return reflect.TypeOf((*lbPoolState)(nil)).Elem()
}

type lbPoolArgs struct {
	// Deprecated: this field is deprecated
	AdminStateUp     *bool               `pulumi:"adminStateUp"`
	Description      *string             `pulumi:"description"`
	LbMethod         string              `pulumi:"lbMethod"`
	ListenerId       *string             `pulumi:"listenerId"`
	LoadbalancerId   *string             `pulumi:"loadbalancerId"`
	Name             *string             `pulumi:"name"`
	Persistences     []LbPoolPersistence `pulumi:"persistences"`
	ProtectionReason *string             `pulumi:"protectionReason"`
	ProtectionStatus *string             `pulumi:"protectionStatus"`
	Protocol         string              `pulumi:"protocol"`
	Region           *string             `pulumi:"region"`
	// Deprecated: tenant_id is deprecated
	TenantId *string `pulumi:"tenantId"`
}

// The set of arguments for constructing a LbPool resource.
type LbPoolArgs struct {
	// Deprecated: this field is deprecated
	AdminStateUp     pulumi.BoolPtrInput
	Description      pulumi.StringPtrInput
	LbMethod         pulumi.StringInput
	ListenerId       pulumi.StringPtrInput
	LoadbalancerId   pulumi.StringPtrInput
	Name             pulumi.StringPtrInput
	Persistences     LbPoolPersistenceArrayInput
	ProtectionReason pulumi.StringPtrInput
	ProtectionStatus pulumi.StringPtrInput
	Protocol         pulumi.StringInput
	Region           pulumi.StringPtrInput
	// Deprecated: tenant_id is deprecated
	TenantId pulumi.StringPtrInput
}

func (LbPoolArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*lbPoolArgs)(nil)).Elem()
}

type LbPoolInput interface {
	pulumi.Input

	ToLbPoolOutput() LbPoolOutput
	ToLbPoolOutputWithContext(ctx context.Context) LbPoolOutput
}

func (*LbPool) ElementType() reflect.Type {
	return reflect.TypeOf((**LbPool)(nil)).Elem()
}

func (i *LbPool) ToLbPoolOutput() LbPoolOutput {
	return i.ToLbPoolOutputWithContext(context.Background())
}

func (i *LbPool) ToLbPoolOutputWithContext(ctx context.Context) LbPoolOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LbPoolOutput)
}

// LbPoolArrayInput is an input type that accepts LbPoolArray and LbPoolArrayOutput values.
// You can construct a concrete instance of `LbPoolArrayInput` via:
//
//	LbPoolArray{ LbPoolArgs{...} }
type LbPoolArrayInput interface {
	pulumi.Input

	ToLbPoolArrayOutput() LbPoolArrayOutput
	ToLbPoolArrayOutputWithContext(context.Context) LbPoolArrayOutput
}

type LbPoolArray []LbPoolInput

func (LbPoolArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*LbPool)(nil)).Elem()
}

func (i LbPoolArray) ToLbPoolArrayOutput() LbPoolArrayOutput {
	return i.ToLbPoolArrayOutputWithContext(context.Background())
}

func (i LbPoolArray) ToLbPoolArrayOutputWithContext(ctx context.Context) LbPoolArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LbPoolArrayOutput)
}

// LbPoolMapInput is an input type that accepts LbPoolMap and LbPoolMapOutput values.
// You can construct a concrete instance of `LbPoolMapInput` via:
//
//	LbPoolMap{ "key": LbPoolArgs{...} }
type LbPoolMapInput interface {
	pulumi.Input

	ToLbPoolMapOutput() LbPoolMapOutput
	ToLbPoolMapOutputWithContext(context.Context) LbPoolMapOutput
}

type LbPoolMap map[string]LbPoolInput

func (LbPoolMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*LbPool)(nil)).Elem()
}

func (i LbPoolMap) ToLbPoolMapOutput() LbPoolMapOutput {
	return i.ToLbPoolMapOutputWithContext(context.Background())
}

func (i LbPoolMap) ToLbPoolMapOutputWithContext(ctx context.Context) LbPoolMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LbPoolMapOutput)
}

type LbPoolOutput struct{ *pulumi.OutputState }

func (LbPoolOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LbPool)(nil)).Elem()
}

func (o LbPoolOutput) ToLbPoolOutput() LbPoolOutput {
	return o
}

func (o LbPoolOutput) ToLbPoolOutputWithContext(ctx context.Context) LbPoolOutput {
	return o
}

// Deprecated: this field is deprecated
func (o LbPoolOutput) AdminStateUp() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *LbPool) pulumi.BoolPtrOutput { return v.AdminStateUp }).(pulumi.BoolPtrOutput)
}

func (o LbPoolOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LbPool) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

func (o LbPoolOutput) LbMethod() pulumi.StringOutput {
	return o.ApplyT(func(v *LbPool) pulumi.StringOutput { return v.LbMethod }).(pulumi.StringOutput)
}

func (o LbPoolOutput) ListenerId() pulumi.StringOutput {
	return o.ApplyT(func(v *LbPool) pulumi.StringOutput { return v.ListenerId }).(pulumi.StringOutput)
}

func (o LbPoolOutput) LoadbalancerId() pulumi.StringOutput {
	return o.ApplyT(func(v *LbPool) pulumi.StringOutput { return v.LoadbalancerId }).(pulumi.StringOutput)
}

func (o LbPoolOutput) MonitorId() pulumi.StringOutput {
	return o.ApplyT(func(v *LbPool) pulumi.StringOutput { return v.MonitorId }).(pulumi.StringOutput)
}

func (o LbPoolOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *LbPool) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o LbPoolOutput) Persistences() LbPoolPersistenceArrayOutput {
	return o.ApplyT(func(v *LbPool) LbPoolPersistenceArrayOutput { return v.Persistences }).(LbPoolPersistenceArrayOutput)
}

func (o LbPoolOutput) ProtectionReason() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LbPool) pulumi.StringPtrOutput { return v.ProtectionReason }).(pulumi.StringPtrOutput)
}

func (o LbPoolOutput) ProtectionStatus() pulumi.StringOutput {
	return o.ApplyT(func(v *LbPool) pulumi.StringOutput { return v.ProtectionStatus }).(pulumi.StringOutput)
}

func (o LbPoolOutput) Protocol() pulumi.StringOutput {
	return o.ApplyT(func(v *LbPool) pulumi.StringOutput { return v.Protocol }).(pulumi.StringOutput)
}

func (o LbPoolOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v *LbPool) pulumi.StringOutput { return v.Region }).(pulumi.StringOutput)
}

// Deprecated: tenant_id is deprecated
func (o LbPoolOutput) TenantId() pulumi.StringOutput {
	return o.ApplyT(func(v *LbPool) pulumi.StringOutput { return v.TenantId }).(pulumi.StringOutput)
}

type LbPoolArrayOutput struct{ *pulumi.OutputState }

func (LbPoolArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*LbPool)(nil)).Elem()
}

func (o LbPoolArrayOutput) ToLbPoolArrayOutput() LbPoolArrayOutput {
	return o
}

func (o LbPoolArrayOutput) ToLbPoolArrayOutputWithContext(ctx context.Context) LbPoolArrayOutput {
	return o
}

func (o LbPoolArrayOutput) Index(i pulumi.IntInput) LbPoolOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *LbPool {
		return vs[0].([]*LbPool)[vs[1].(int)]
	}).(LbPoolOutput)
}

type LbPoolMapOutput struct{ *pulumi.OutputState }

func (LbPoolMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*LbPool)(nil)).Elem()
}

func (o LbPoolMapOutput) ToLbPoolMapOutput() LbPoolMapOutput {
	return o
}

func (o LbPoolMapOutput) ToLbPoolMapOutputWithContext(ctx context.Context) LbPoolMapOutput {
	return o
}

func (o LbPoolMapOutput) MapIndex(k pulumi.StringInput) LbPoolOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *LbPool {
		return vs[0].(map[string]*LbPool)[vs[1].(string)]
	}).(LbPoolOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*LbPoolInput)(nil)).Elem(), &LbPool{})
	pulumi.RegisterInputType(reflect.TypeOf((*LbPoolArrayInput)(nil)).Elem(), LbPoolArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*LbPoolMapInput)(nil)).Elem(), LbPoolMap{})
	pulumi.RegisterOutputType(LbPoolOutput{})
	pulumi.RegisterOutputType(LbPoolArrayOutput{})
	pulumi.RegisterOutputType(LbPoolMapOutput{})
}
