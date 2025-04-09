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

type DmsInstance struct {
	pulumi.CustomResourceState

	AccessUser        pulumi.StringPtrOutput   `pulumi:"accessUser"`
	AvailableZones    pulumi.StringArrayOutput `pulumi:"availableZones"`
	ConnectAddress    pulumi.StringOutput      `pulumi:"connectAddress"`
	CreatedAt         pulumi.StringOutput      `pulumi:"createdAt"`
	Description       pulumi.StringOutput      `pulumi:"description"`
	Engine            pulumi.StringOutput      `pulumi:"engine"`
	EngineVersion     pulumi.StringPtrOutput   `pulumi:"engineVersion"`
	MaintainBegin     pulumi.StringOutput      `pulumi:"maintainBegin"`
	MaintainEnd       pulumi.StringOutput      `pulumi:"maintainEnd"`
	Name              pulumi.StringOutput      `pulumi:"name"`
	OrderId           pulumi.StringOutput      `pulumi:"orderId"`
	PartitionNum      pulumi.IntPtrOutput      `pulumi:"partitionNum"`
	Password          pulumi.StringPtrOutput   `pulumi:"password"`
	Port              pulumi.StringOutput      `pulumi:"port"`
	ProductId         pulumi.StringOutput      `pulumi:"productId"`
	Region            pulumi.StringOutput      `pulumi:"region"`
	ResourceSpecCode  pulumi.StringOutput      `pulumi:"resourceSpecCode"`
	SecurityGroupId   pulumi.StringOutput      `pulumi:"securityGroupId"`
	SecurityGroupName pulumi.StringOutput      `pulumi:"securityGroupName"`
	Specification     pulumi.StringOutput      `pulumi:"specification"`
	Status            pulumi.StringOutput      `pulumi:"status"`
	StorageSpace      pulumi.IntOutput         `pulumi:"storageSpace"`
	StorageSpecCode   pulumi.StringOutput      `pulumi:"storageSpecCode"`
	SubnetId          pulumi.StringOutput      `pulumi:"subnetId"`
	SubnetName        pulumi.StringOutput      `pulumi:"subnetName"`
	Tags              pulumi.StringMapOutput   `pulumi:"tags"`
	Type              pulumi.StringOutput      `pulumi:"type"`
	UsedStorageSpace  pulumi.IntOutput         `pulumi:"usedStorageSpace"`
	UserId            pulumi.StringOutput      `pulumi:"userId"`
	VpcId             pulumi.StringOutput      `pulumi:"vpcId"`
	VpcName           pulumi.StringOutput      `pulumi:"vpcName"`
}

// NewDmsInstance registers a new resource with the given unique name, arguments, and options.
func NewDmsInstance(ctx *pulumi.Context,
	name string, args *DmsInstanceArgs, opts ...pulumi.ResourceOption) (*DmsInstance, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AvailableZones == nil {
		return nil, errors.New("invalid value for required argument 'AvailableZones'")
	}
	if args.Engine == nil {
		return nil, errors.New("invalid value for required argument 'Engine'")
	}
	if args.ProductId == nil {
		return nil, errors.New("invalid value for required argument 'ProductId'")
	}
	if args.SecurityGroupId == nil {
		return nil, errors.New("invalid value for required argument 'SecurityGroupId'")
	}
	if args.StorageSpace == nil {
		return nil, errors.New("invalid value for required argument 'StorageSpace'")
	}
	if args.StorageSpecCode == nil {
		return nil, errors.New("invalid value for required argument 'StorageSpecCode'")
	}
	if args.SubnetId == nil {
		return nil, errors.New("invalid value for required argument 'SubnetId'")
	}
	if args.VpcId == nil {
		return nil, errors.New("invalid value for required argument 'VpcId'")
	}
	if args.Password != nil {
		args.Password = pulumi.ToSecret(args.Password).(pulumi.StringPtrInput)
	}
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"password",
	})
	opts = append(opts, secrets)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource DmsInstance
	err := ctx.RegisterResource("sbercloud:index/dmsInstance:DmsInstance", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDmsInstance gets an existing DmsInstance resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDmsInstance(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DmsInstanceState, opts ...pulumi.ResourceOption) (*DmsInstance, error) {
	var resource DmsInstance
	err := ctx.ReadResource("sbercloud:index/dmsInstance:DmsInstance", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DmsInstance resources.
type dmsInstanceState struct {
	AccessUser        *string           `pulumi:"accessUser"`
	AvailableZones    []string          `pulumi:"availableZones"`
	ConnectAddress    *string           `pulumi:"connectAddress"`
	CreatedAt         *string           `pulumi:"createdAt"`
	Description       *string           `pulumi:"description"`
	Engine            *string           `pulumi:"engine"`
	EngineVersion     *string           `pulumi:"engineVersion"`
	MaintainBegin     *string           `pulumi:"maintainBegin"`
	MaintainEnd       *string           `pulumi:"maintainEnd"`
	Name              *string           `pulumi:"name"`
	OrderId           *string           `pulumi:"orderId"`
	PartitionNum      *int              `pulumi:"partitionNum"`
	Password          *string           `pulumi:"password"`
	Port              *string           `pulumi:"port"`
	ProductId         *string           `pulumi:"productId"`
	Region            *string           `pulumi:"region"`
	ResourceSpecCode  *string           `pulumi:"resourceSpecCode"`
	SecurityGroupId   *string           `pulumi:"securityGroupId"`
	SecurityGroupName *string           `pulumi:"securityGroupName"`
	Specification     *string           `pulumi:"specification"`
	Status            *string           `pulumi:"status"`
	StorageSpace      *int              `pulumi:"storageSpace"`
	StorageSpecCode   *string           `pulumi:"storageSpecCode"`
	SubnetId          *string           `pulumi:"subnetId"`
	SubnetName        *string           `pulumi:"subnetName"`
	Tags              map[string]string `pulumi:"tags"`
	Type              *string           `pulumi:"type"`
	UsedStorageSpace  *int              `pulumi:"usedStorageSpace"`
	UserId            *string           `pulumi:"userId"`
	VpcId             *string           `pulumi:"vpcId"`
	VpcName           *string           `pulumi:"vpcName"`
}

type DmsInstanceState struct {
	AccessUser        pulumi.StringPtrInput
	AvailableZones    pulumi.StringArrayInput
	ConnectAddress    pulumi.StringPtrInput
	CreatedAt         pulumi.StringPtrInput
	Description       pulumi.StringPtrInput
	Engine            pulumi.StringPtrInput
	EngineVersion     pulumi.StringPtrInput
	MaintainBegin     pulumi.StringPtrInput
	MaintainEnd       pulumi.StringPtrInput
	Name              pulumi.StringPtrInput
	OrderId           pulumi.StringPtrInput
	PartitionNum      pulumi.IntPtrInput
	Password          pulumi.StringPtrInput
	Port              pulumi.StringPtrInput
	ProductId         pulumi.StringPtrInput
	Region            pulumi.StringPtrInput
	ResourceSpecCode  pulumi.StringPtrInput
	SecurityGroupId   pulumi.StringPtrInput
	SecurityGroupName pulumi.StringPtrInput
	Specification     pulumi.StringPtrInput
	Status            pulumi.StringPtrInput
	StorageSpace      pulumi.IntPtrInput
	StorageSpecCode   pulumi.StringPtrInput
	SubnetId          pulumi.StringPtrInput
	SubnetName        pulumi.StringPtrInput
	Tags              pulumi.StringMapInput
	Type              pulumi.StringPtrInput
	UsedStorageSpace  pulumi.IntPtrInput
	UserId            pulumi.StringPtrInput
	VpcId             pulumi.StringPtrInput
	VpcName           pulumi.StringPtrInput
}

func (DmsInstanceState) ElementType() reflect.Type {
	return reflect.TypeOf((*dmsInstanceState)(nil)).Elem()
}

type dmsInstanceArgs struct {
	AccessUser      *string           `pulumi:"accessUser"`
	AvailableZones  []string          `pulumi:"availableZones"`
	Description     *string           `pulumi:"description"`
	Engine          string            `pulumi:"engine"`
	EngineVersion   *string           `pulumi:"engineVersion"`
	MaintainBegin   *string           `pulumi:"maintainBegin"`
	MaintainEnd     *string           `pulumi:"maintainEnd"`
	Name            *string           `pulumi:"name"`
	PartitionNum    *int              `pulumi:"partitionNum"`
	Password        *string           `pulumi:"password"`
	ProductId       string            `pulumi:"productId"`
	Region          *string           `pulumi:"region"`
	SecurityGroupId string            `pulumi:"securityGroupId"`
	Specification   *string           `pulumi:"specification"`
	StorageSpace    int               `pulumi:"storageSpace"`
	StorageSpecCode string            `pulumi:"storageSpecCode"`
	SubnetId        string            `pulumi:"subnetId"`
	Tags            map[string]string `pulumi:"tags"`
	VpcId           string            `pulumi:"vpcId"`
}

// The set of arguments for constructing a DmsInstance resource.
type DmsInstanceArgs struct {
	AccessUser      pulumi.StringPtrInput
	AvailableZones  pulumi.StringArrayInput
	Description     pulumi.StringPtrInput
	Engine          pulumi.StringInput
	EngineVersion   pulumi.StringPtrInput
	MaintainBegin   pulumi.StringPtrInput
	MaintainEnd     pulumi.StringPtrInput
	Name            pulumi.StringPtrInput
	PartitionNum    pulumi.IntPtrInput
	Password        pulumi.StringPtrInput
	ProductId       pulumi.StringInput
	Region          pulumi.StringPtrInput
	SecurityGroupId pulumi.StringInput
	Specification   pulumi.StringPtrInput
	StorageSpace    pulumi.IntInput
	StorageSpecCode pulumi.StringInput
	SubnetId        pulumi.StringInput
	Tags            pulumi.StringMapInput
	VpcId           pulumi.StringInput
}

func (DmsInstanceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*dmsInstanceArgs)(nil)).Elem()
}

type DmsInstanceInput interface {
	pulumi.Input

	ToDmsInstanceOutput() DmsInstanceOutput
	ToDmsInstanceOutputWithContext(ctx context.Context) DmsInstanceOutput
}

func (*DmsInstance) ElementType() reflect.Type {
	return reflect.TypeOf((**DmsInstance)(nil)).Elem()
}

func (i *DmsInstance) ToDmsInstanceOutput() DmsInstanceOutput {
	return i.ToDmsInstanceOutputWithContext(context.Background())
}

func (i *DmsInstance) ToDmsInstanceOutputWithContext(ctx context.Context) DmsInstanceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DmsInstanceOutput)
}

// DmsInstanceArrayInput is an input type that accepts DmsInstanceArray and DmsInstanceArrayOutput values.
// You can construct a concrete instance of `DmsInstanceArrayInput` via:
//
//	DmsInstanceArray{ DmsInstanceArgs{...} }
type DmsInstanceArrayInput interface {
	pulumi.Input

	ToDmsInstanceArrayOutput() DmsInstanceArrayOutput
	ToDmsInstanceArrayOutputWithContext(context.Context) DmsInstanceArrayOutput
}

type DmsInstanceArray []DmsInstanceInput

func (DmsInstanceArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DmsInstance)(nil)).Elem()
}

func (i DmsInstanceArray) ToDmsInstanceArrayOutput() DmsInstanceArrayOutput {
	return i.ToDmsInstanceArrayOutputWithContext(context.Background())
}

func (i DmsInstanceArray) ToDmsInstanceArrayOutputWithContext(ctx context.Context) DmsInstanceArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DmsInstanceArrayOutput)
}

// DmsInstanceMapInput is an input type that accepts DmsInstanceMap and DmsInstanceMapOutput values.
// You can construct a concrete instance of `DmsInstanceMapInput` via:
//
//	DmsInstanceMap{ "key": DmsInstanceArgs{...} }
type DmsInstanceMapInput interface {
	pulumi.Input

	ToDmsInstanceMapOutput() DmsInstanceMapOutput
	ToDmsInstanceMapOutputWithContext(context.Context) DmsInstanceMapOutput
}

type DmsInstanceMap map[string]DmsInstanceInput

func (DmsInstanceMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DmsInstance)(nil)).Elem()
}

func (i DmsInstanceMap) ToDmsInstanceMapOutput() DmsInstanceMapOutput {
	return i.ToDmsInstanceMapOutputWithContext(context.Background())
}

func (i DmsInstanceMap) ToDmsInstanceMapOutputWithContext(ctx context.Context) DmsInstanceMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DmsInstanceMapOutput)
}

type DmsInstanceOutput struct{ *pulumi.OutputState }

func (DmsInstanceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DmsInstance)(nil)).Elem()
}

func (o DmsInstanceOutput) ToDmsInstanceOutput() DmsInstanceOutput {
	return o
}

func (o DmsInstanceOutput) ToDmsInstanceOutputWithContext(ctx context.Context) DmsInstanceOutput {
	return o
}

func (o DmsInstanceOutput) AccessUser() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DmsInstance) pulumi.StringPtrOutput { return v.AccessUser }).(pulumi.StringPtrOutput)
}

func (o DmsInstanceOutput) AvailableZones() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *DmsInstance) pulumi.StringArrayOutput { return v.AvailableZones }).(pulumi.StringArrayOutput)
}

func (o DmsInstanceOutput) ConnectAddress() pulumi.StringOutput {
	return o.ApplyT(func(v *DmsInstance) pulumi.StringOutput { return v.ConnectAddress }).(pulumi.StringOutput)
}

func (o DmsInstanceOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *DmsInstance) pulumi.StringOutput { return v.CreatedAt }).(pulumi.StringOutput)
}

func (o DmsInstanceOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v *DmsInstance) pulumi.StringOutput { return v.Description }).(pulumi.StringOutput)
}

func (o DmsInstanceOutput) Engine() pulumi.StringOutput {
	return o.ApplyT(func(v *DmsInstance) pulumi.StringOutput { return v.Engine }).(pulumi.StringOutput)
}

func (o DmsInstanceOutput) EngineVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DmsInstance) pulumi.StringPtrOutput { return v.EngineVersion }).(pulumi.StringPtrOutput)
}

func (o DmsInstanceOutput) MaintainBegin() pulumi.StringOutput {
	return o.ApplyT(func(v *DmsInstance) pulumi.StringOutput { return v.MaintainBegin }).(pulumi.StringOutput)
}

func (o DmsInstanceOutput) MaintainEnd() pulumi.StringOutput {
	return o.ApplyT(func(v *DmsInstance) pulumi.StringOutput { return v.MaintainEnd }).(pulumi.StringOutput)
}

func (o DmsInstanceOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *DmsInstance) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o DmsInstanceOutput) OrderId() pulumi.StringOutput {
	return o.ApplyT(func(v *DmsInstance) pulumi.StringOutput { return v.OrderId }).(pulumi.StringOutput)
}

func (o DmsInstanceOutput) PartitionNum() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *DmsInstance) pulumi.IntPtrOutput { return v.PartitionNum }).(pulumi.IntPtrOutput)
}

func (o DmsInstanceOutput) Password() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DmsInstance) pulumi.StringPtrOutput { return v.Password }).(pulumi.StringPtrOutput)
}

func (o DmsInstanceOutput) Port() pulumi.StringOutput {
	return o.ApplyT(func(v *DmsInstance) pulumi.StringOutput { return v.Port }).(pulumi.StringOutput)
}

func (o DmsInstanceOutput) ProductId() pulumi.StringOutput {
	return o.ApplyT(func(v *DmsInstance) pulumi.StringOutput { return v.ProductId }).(pulumi.StringOutput)
}

func (o DmsInstanceOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v *DmsInstance) pulumi.StringOutput { return v.Region }).(pulumi.StringOutput)
}

func (o DmsInstanceOutput) ResourceSpecCode() pulumi.StringOutput {
	return o.ApplyT(func(v *DmsInstance) pulumi.StringOutput { return v.ResourceSpecCode }).(pulumi.StringOutput)
}

func (o DmsInstanceOutput) SecurityGroupId() pulumi.StringOutput {
	return o.ApplyT(func(v *DmsInstance) pulumi.StringOutput { return v.SecurityGroupId }).(pulumi.StringOutput)
}

func (o DmsInstanceOutput) SecurityGroupName() pulumi.StringOutput {
	return o.ApplyT(func(v *DmsInstance) pulumi.StringOutput { return v.SecurityGroupName }).(pulumi.StringOutput)
}

func (o DmsInstanceOutput) Specification() pulumi.StringOutput {
	return o.ApplyT(func(v *DmsInstance) pulumi.StringOutput { return v.Specification }).(pulumi.StringOutput)
}

func (o DmsInstanceOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *DmsInstance) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

func (o DmsInstanceOutput) StorageSpace() pulumi.IntOutput {
	return o.ApplyT(func(v *DmsInstance) pulumi.IntOutput { return v.StorageSpace }).(pulumi.IntOutput)
}

func (o DmsInstanceOutput) StorageSpecCode() pulumi.StringOutput {
	return o.ApplyT(func(v *DmsInstance) pulumi.StringOutput { return v.StorageSpecCode }).(pulumi.StringOutput)
}

func (o DmsInstanceOutput) SubnetId() pulumi.StringOutput {
	return o.ApplyT(func(v *DmsInstance) pulumi.StringOutput { return v.SubnetId }).(pulumi.StringOutput)
}

func (o DmsInstanceOutput) SubnetName() pulumi.StringOutput {
	return o.ApplyT(func(v *DmsInstance) pulumi.StringOutput { return v.SubnetName }).(pulumi.StringOutput)
}

func (o DmsInstanceOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *DmsInstance) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o DmsInstanceOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *DmsInstance) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func (o DmsInstanceOutput) UsedStorageSpace() pulumi.IntOutput {
	return o.ApplyT(func(v *DmsInstance) pulumi.IntOutput { return v.UsedStorageSpace }).(pulumi.IntOutput)
}

func (o DmsInstanceOutput) UserId() pulumi.StringOutput {
	return o.ApplyT(func(v *DmsInstance) pulumi.StringOutput { return v.UserId }).(pulumi.StringOutput)
}

func (o DmsInstanceOutput) VpcId() pulumi.StringOutput {
	return o.ApplyT(func(v *DmsInstance) pulumi.StringOutput { return v.VpcId }).(pulumi.StringOutput)
}

func (o DmsInstanceOutput) VpcName() pulumi.StringOutput {
	return o.ApplyT(func(v *DmsInstance) pulumi.StringOutput { return v.VpcName }).(pulumi.StringOutput)
}

type DmsInstanceArrayOutput struct{ *pulumi.OutputState }

func (DmsInstanceArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DmsInstance)(nil)).Elem()
}

func (o DmsInstanceArrayOutput) ToDmsInstanceArrayOutput() DmsInstanceArrayOutput {
	return o
}

func (o DmsInstanceArrayOutput) ToDmsInstanceArrayOutputWithContext(ctx context.Context) DmsInstanceArrayOutput {
	return o
}

func (o DmsInstanceArrayOutput) Index(i pulumi.IntInput) DmsInstanceOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *DmsInstance {
		return vs[0].([]*DmsInstance)[vs[1].(int)]
	}).(DmsInstanceOutput)
}

type DmsInstanceMapOutput struct{ *pulumi.OutputState }

func (DmsInstanceMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DmsInstance)(nil)).Elem()
}

func (o DmsInstanceMapOutput) ToDmsInstanceMapOutput() DmsInstanceMapOutput {
	return o
}

func (o DmsInstanceMapOutput) ToDmsInstanceMapOutputWithContext(ctx context.Context) DmsInstanceMapOutput {
	return o
}

func (o DmsInstanceMapOutput) MapIndex(k pulumi.StringInput) DmsInstanceOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *DmsInstance {
		return vs[0].(map[string]*DmsInstance)[vs[1].(string)]
	}).(DmsInstanceOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DmsInstanceInput)(nil)).Elem(), &DmsInstance{})
	pulumi.RegisterInputType(reflect.TypeOf((*DmsInstanceArrayInput)(nil)).Elem(), DmsInstanceArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*DmsInstanceMapInput)(nil)).Elem(), DmsInstanceMap{})
	pulumi.RegisterOutputType(DmsInstanceOutput{})
	pulumi.RegisterOutputType(DmsInstanceArrayOutput{})
	pulumi.RegisterOutputType(DmsInstanceMapOutput{})
}
