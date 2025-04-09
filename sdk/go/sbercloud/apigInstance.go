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

type ApigInstance struct {
	pulumi.CustomResourceState

	// schema: Required; The name list of availability zones for the dedicated instance.
	AvailabilityZones pulumi.StringArrayOutput `pulumi:"availabilityZones"`
	// schema: Deprecated; The name list of availability zones for the dedicated instance.
	AvailableZones pulumi.StringArrayOutput `pulumi:"availableZones"`
	// The egress bandwidth size of the dedicated instance.
	BandwidthSize pulumi.IntPtrOutput `pulumi:"bandwidthSize"`
	// schema: Deprecated; Time when the dedicated instance is created.
	//
	// Deprecated: Use 'created_at' instead
	CreateTime pulumi.StringOutput `pulumi:"createTime"`
	// Time when the dedicated instance is created, in RFC-3339 format.
	CreatedAt pulumi.StringOutput `pulumi:"createdAt"`
	// Specified the list of the instance custom ingress ports.
	CustomIngressPorts ApigInstanceCustomIngressPortArrayOutput `pulumi:"customIngressPorts"`
	// The description of the dedicated instance.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The edition of the dedicated instance.
	Edition pulumi.StringOutput `pulumi:"edition"`
	// The egress (NAT) public IP address.
	EgressAddress pulumi.StringOutput `pulumi:"egressAddress"`
	// The EIP ID associated with the dedicated instance.
	EipId pulumi.StringOutput `pulumi:"eipId"`
	// The enterprise project ID to which the dedicated instance belongs.
	EnterpriseProjectId pulumi.StringOutput `pulumi:"enterpriseProjectId"`
	// The ingress EIP address.
	IngressAddress               pulumi.StringOutput    `pulumi:"ingressAddress"`
	IngressBandwidthChargingMode pulumi.StringPtrOutput `pulumi:"ingressBandwidthChargingMode"`
	IngressBandwidthSize         pulumi.IntPtrOutput    `pulumi:"ingressBandwidthSize"`
	// Whether public access with an IPv6 address is supported.
	Ipv6Enable pulumi.BoolOutput `pulumi:"ipv6Enable"`
	// The type of loadbalancer provider used by the instance.
	LoadbalancerProvider pulumi.StringOutput `pulumi:"loadbalancerProvider"`
	// The start time of the maintenance time window.
	MaintainBegin pulumi.StringOutput `pulumi:"maintainBegin"`
	// End time of the maintenance time window, 4-hour difference between the start time and end time.
	MaintainEnd pulumi.StringOutput `pulumi:"maintainEnd"`
	// The name of the dedicated instance.
	Name pulumi.StringOutput `pulumi:"name"`
	// The region in which to create the dedicated instance resource.
	Region pulumi.StringOutput `pulumi:"region"`
	// The ID of the security group to which the dedicated instance belongs to.
	SecurityGroupId pulumi.StringOutput `pulumi:"securityGroupId"`
	// Status of the dedicated instance.
	Status pulumi.StringOutput `pulumi:"status"`
	// The ID of the VPC subnet used to create the dedicated instance.
	SubnetId pulumi.StringOutput `pulumi:"subnetId"`
	// The supported features of the dedicated instance.
	SupportedFeatures pulumi.StringArrayOutput `pulumi:"supportedFeatures"`
	Tags              pulumi.StringMapOutput   `pulumi:"tags"`
	// The ID of the VPC used to create the dedicated instance.
	VpcId pulumi.StringOutput `pulumi:"vpcId"`
	// The ingress private IP address of the VPC.
	VpcIngressAddress pulumi.StringOutput `pulumi:"vpcIngressAddress"`
	// The address (full name) of the VPC endpoint service.
	VpcepServiceAddress pulumi.StringOutput `pulumi:"vpcepServiceAddress"`
	// Name of the VPC endpoint service.
	VpcepServiceName pulumi.StringOutput `pulumi:"vpcepServiceName"`
}

// NewApigInstance registers a new resource with the given unique name, arguments, and options.
func NewApigInstance(ctx *pulumi.Context,
	name string, args *ApigInstanceArgs, opts ...pulumi.ResourceOption) (*ApigInstance, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Edition == nil {
		return nil, errors.New("invalid value for required argument 'Edition'")
	}
	if args.SecurityGroupId == nil {
		return nil, errors.New("invalid value for required argument 'SecurityGroupId'")
	}
	if args.SubnetId == nil {
		return nil, errors.New("invalid value for required argument 'SubnetId'")
	}
	if args.VpcId == nil {
		return nil, errors.New("invalid value for required argument 'VpcId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource ApigInstance
	err := ctx.RegisterResource("sbercloud:index/apigInstance:ApigInstance", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetApigInstance gets an existing ApigInstance resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetApigInstance(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ApigInstanceState, opts ...pulumi.ResourceOption) (*ApigInstance, error) {
	var resource ApigInstance
	err := ctx.ReadResource("sbercloud:index/apigInstance:ApigInstance", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ApigInstance resources.
type apigInstanceState struct {
	// schema: Required; The name list of availability zones for the dedicated instance.
	AvailabilityZones []string `pulumi:"availabilityZones"`
	// schema: Deprecated; The name list of availability zones for the dedicated instance.
	AvailableZones []string `pulumi:"availableZones"`
	// The egress bandwidth size of the dedicated instance.
	BandwidthSize *int `pulumi:"bandwidthSize"`
	// schema: Deprecated; Time when the dedicated instance is created.
	//
	// Deprecated: Use 'created_at' instead
	CreateTime *string `pulumi:"createTime"`
	// Time when the dedicated instance is created, in RFC-3339 format.
	CreatedAt *string `pulumi:"createdAt"`
	// Specified the list of the instance custom ingress ports.
	CustomIngressPorts []ApigInstanceCustomIngressPort `pulumi:"customIngressPorts"`
	// The description of the dedicated instance.
	Description *string `pulumi:"description"`
	// The edition of the dedicated instance.
	Edition *string `pulumi:"edition"`
	// The egress (NAT) public IP address.
	EgressAddress *string `pulumi:"egressAddress"`
	// The EIP ID associated with the dedicated instance.
	EipId *string `pulumi:"eipId"`
	// The enterprise project ID to which the dedicated instance belongs.
	EnterpriseProjectId *string `pulumi:"enterpriseProjectId"`
	// The ingress EIP address.
	IngressAddress               *string `pulumi:"ingressAddress"`
	IngressBandwidthChargingMode *string `pulumi:"ingressBandwidthChargingMode"`
	IngressBandwidthSize         *int    `pulumi:"ingressBandwidthSize"`
	// Whether public access with an IPv6 address is supported.
	Ipv6Enable *bool `pulumi:"ipv6Enable"`
	// The type of loadbalancer provider used by the instance.
	LoadbalancerProvider *string `pulumi:"loadbalancerProvider"`
	// The start time of the maintenance time window.
	MaintainBegin *string `pulumi:"maintainBegin"`
	// End time of the maintenance time window, 4-hour difference between the start time and end time.
	MaintainEnd *string `pulumi:"maintainEnd"`
	// The name of the dedicated instance.
	Name *string `pulumi:"name"`
	// The region in which to create the dedicated instance resource.
	Region *string `pulumi:"region"`
	// The ID of the security group to which the dedicated instance belongs to.
	SecurityGroupId *string `pulumi:"securityGroupId"`
	// Status of the dedicated instance.
	Status *string `pulumi:"status"`
	// The ID of the VPC subnet used to create the dedicated instance.
	SubnetId *string `pulumi:"subnetId"`
	// The supported features of the dedicated instance.
	SupportedFeatures []string          `pulumi:"supportedFeatures"`
	Tags              map[string]string `pulumi:"tags"`
	// The ID of the VPC used to create the dedicated instance.
	VpcId *string `pulumi:"vpcId"`
	// The ingress private IP address of the VPC.
	VpcIngressAddress *string `pulumi:"vpcIngressAddress"`
	// The address (full name) of the VPC endpoint service.
	VpcepServiceAddress *string `pulumi:"vpcepServiceAddress"`
	// Name of the VPC endpoint service.
	VpcepServiceName *string `pulumi:"vpcepServiceName"`
}

type ApigInstanceState struct {
	// schema: Required; The name list of availability zones for the dedicated instance.
	AvailabilityZones pulumi.StringArrayInput
	// schema: Deprecated; The name list of availability zones for the dedicated instance.
	AvailableZones pulumi.StringArrayInput
	// The egress bandwidth size of the dedicated instance.
	BandwidthSize pulumi.IntPtrInput
	// schema: Deprecated; Time when the dedicated instance is created.
	//
	// Deprecated: Use 'created_at' instead
	CreateTime pulumi.StringPtrInput
	// Time when the dedicated instance is created, in RFC-3339 format.
	CreatedAt pulumi.StringPtrInput
	// Specified the list of the instance custom ingress ports.
	CustomIngressPorts ApigInstanceCustomIngressPortArrayInput
	// The description of the dedicated instance.
	Description pulumi.StringPtrInput
	// The edition of the dedicated instance.
	Edition pulumi.StringPtrInput
	// The egress (NAT) public IP address.
	EgressAddress pulumi.StringPtrInput
	// The EIP ID associated with the dedicated instance.
	EipId pulumi.StringPtrInput
	// The enterprise project ID to which the dedicated instance belongs.
	EnterpriseProjectId pulumi.StringPtrInput
	// The ingress EIP address.
	IngressAddress               pulumi.StringPtrInput
	IngressBandwidthChargingMode pulumi.StringPtrInput
	IngressBandwidthSize         pulumi.IntPtrInput
	// Whether public access with an IPv6 address is supported.
	Ipv6Enable pulumi.BoolPtrInput
	// The type of loadbalancer provider used by the instance.
	LoadbalancerProvider pulumi.StringPtrInput
	// The start time of the maintenance time window.
	MaintainBegin pulumi.StringPtrInput
	// End time of the maintenance time window, 4-hour difference between the start time and end time.
	MaintainEnd pulumi.StringPtrInput
	// The name of the dedicated instance.
	Name pulumi.StringPtrInput
	// The region in which to create the dedicated instance resource.
	Region pulumi.StringPtrInput
	// The ID of the security group to which the dedicated instance belongs to.
	SecurityGroupId pulumi.StringPtrInput
	// Status of the dedicated instance.
	Status pulumi.StringPtrInput
	// The ID of the VPC subnet used to create the dedicated instance.
	SubnetId pulumi.StringPtrInput
	// The supported features of the dedicated instance.
	SupportedFeatures pulumi.StringArrayInput
	Tags              pulumi.StringMapInput
	// The ID of the VPC used to create the dedicated instance.
	VpcId pulumi.StringPtrInput
	// The ingress private IP address of the VPC.
	VpcIngressAddress pulumi.StringPtrInput
	// The address (full name) of the VPC endpoint service.
	VpcepServiceAddress pulumi.StringPtrInput
	// Name of the VPC endpoint service.
	VpcepServiceName pulumi.StringPtrInput
}

func (ApigInstanceState) ElementType() reflect.Type {
	return reflect.TypeOf((*apigInstanceState)(nil)).Elem()
}

type apigInstanceArgs struct {
	// schema: Required; The name list of availability zones for the dedicated instance.
	AvailabilityZones []string `pulumi:"availabilityZones"`
	// schema: Deprecated; The name list of availability zones for the dedicated instance.
	AvailableZones []string `pulumi:"availableZones"`
	// The egress bandwidth size of the dedicated instance.
	BandwidthSize *int `pulumi:"bandwidthSize"`
	// Specified the list of the instance custom ingress ports.
	CustomIngressPorts []ApigInstanceCustomIngressPort `pulumi:"customIngressPorts"`
	// The description of the dedicated instance.
	Description *string `pulumi:"description"`
	// The edition of the dedicated instance.
	Edition string `pulumi:"edition"`
	// The EIP ID associated with the dedicated instance.
	EipId *string `pulumi:"eipId"`
	// The enterprise project ID to which the dedicated instance belongs.
	EnterpriseProjectId          *string `pulumi:"enterpriseProjectId"`
	IngressBandwidthChargingMode *string `pulumi:"ingressBandwidthChargingMode"`
	IngressBandwidthSize         *int    `pulumi:"ingressBandwidthSize"`
	// Whether public access with an IPv6 address is supported.
	Ipv6Enable *bool `pulumi:"ipv6Enable"`
	// The type of loadbalancer provider used by the instance.
	LoadbalancerProvider *string `pulumi:"loadbalancerProvider"`
	// The start time of the maintenance time window.
	MaintainBegin *string `pulumi:"maintainBegin"`
	// The name of the dedicated instance.
	Name *string `pulumi:"name"`
	// The region in which to create the dedicated instance resource.
	Region *string `pulumi:"region"`
	// The ID of the security group to which the dedicated instance belongs to.
	SecurityGroupId string `pulumi:"securityGroupId"`
	// The ID of the VPC subnet used to create the dedicated instance.
	SubnetId string            `pulumi:"subnetId"`
	Tags     map[string]string `pulumi:"tags"`
	// The ID of the VPC used to create the dedicated instance.
	VpcId string `pulumi:"vpcId"`
	// Name of the VPC endpoint service.
	VpcepServiceName *string `pulumi:"vpcepServiceName"`
}

// The set of arguments for constructing a ApigInstance resource.
type ApigInstanceArgs struct {
	// schema: Required; The name list of availability zones for the dedicated instance.
	AvailabilityZones pulumi.StringArrayInput
	// schema: Deprecated; The name list of availability zones for the dedicated instance.
	AvailableZones pulumi.StringArrayInput
	// The egress bandwidth size of the dedicated instance.
	BandwidthSize pulumi.IntPtrInput
	// Specified the list of the instance custom ingress ports.
	CustomIngressPorts ApigInstanceCustomIngressPortArrayInput
	// The description of the dedicated instance.
	Description pulumi.StringPtrInput
	// The edition of the dedicated instance.
	Edition pulumi.StringInput
	// The EIP ID associated with the dedicated instance.
	EipId pulumi.StringPtrInput
	// The enterprise project ID to which the dedicated instance belongs.
	EnterpriseProjectId          pulumi.StringPtrInput
	IngressBandwidthChargingMode pulumi.StringPtrInput
	IngressBandwidthSize         pulumi.IntPtrInput
	// Whether public access with an IPv6 address is supported.
	Ipv6Enable pulumi.BoolPtrInput
	// The type of loadbalancer provider used by the instance.
	LoadbalancerProvider pulumi.StringPtrInput
	// The start time of the maintenance time window.
	MaintainBegin pulumi.StringPtrInput
	// The name of the dedicated instance.
	Name pulumi.StringPtrInput
	// The region in which to create the dedicated instance resource.
	Region pulumi.StringPtrInput
	// The ID of the security group to which the dedicated instance belongs to.
	SecurityGroupId pulumi.StringInput
	// The ID of the VPC subnet used to create the dedicated instance.
	SubnetId pulumi.StringInput
	Tags     pulumi.StringMapInput
	// The ID of the VPC used to create the dedicated instance.
	VpcId pulumi.StringInput
	// Name of the VPC endpoint service.
	VpcepServiceName pulumi.StringPtrInput
}

func (ApigInstanceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*apigInstanceArgs)(nil)).Elem()
}

type ApigInstanceInput interface {
	pulumi.Input

	ToApigInstanceOutput() ApigInstanceOutput
	ToApigInstanceOutputWithContext(ctx context.Context) ApigInstanceOutput
}

func (*ApigInstance) ElementType() reflect.Type {
	return reflect.TypeOf((**ApigInstance)(nil)).Elem()
}

func (i *ApigInstance) ToApigInstanceOutput() ApigInstanceOutput {
	return i.ToApigInstanceOutputWithContext(context.Background())
}

func (i *ApigInstance) ToApigInstanceOutputWithContext(ctx context.Context) ApigInstanceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApigInstanceOutput)
}

// ApigInstanceArrayInput is an input type that accepts ApigInstanceArray and ApigInstanceArrayOutput values.
// You can construct a concrete instance of `ApigInstanceArrayInput` via:
//
//	ApigInstanceArray{ ApigInstanceArgs{...} }
type ApigInstanceArrayInput interface {
	pulumi.Input

	ToApigInstanceArrayOutput() ApigInstanceArrayOutput
	ToApigInstanceArrayOutputWithContext(context.Context) ApigInstanceArrayOutput
}

type ApigInstanceArray []ApigInstanceInput

func (ApigInstanceArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ApigInstance)(nil)).Elem()
}

func (i ApigInstanceArray) ToApigInstanceArrayOutput() ApigInstanceArrayOutput {
	return i.ToApigInstanceArrayOutputWithContext(context.Background())
}

func (i ApigInstanceArray) ToApigInstanceArrayOutputWithContext(ctx context.Context) ApigInstanceArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApigInstanceArrayOutput)
}

// ApigInstanceMapInput is an input type that accepts ApigInstanceMap and ApigInstanceMapOutput values.
// You can construct a concrete instance of `ApigInstanceMapInput` via:
//
//	ApigInstanceMap{ "key": ApigInstanceArgs{...} }
type ApigInstanceMapInput interface {
	pulumi.Input

	ToApigInstanceMapOutput() ApigInstanceMapOutput
	ToApigInstanceMapOutputWithContext(context.Context) ApigInstanceMapOutput
}

type ApigInstanceMap map[string]ApigInstanceInput

func (ApigInstanceMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ApigInstance)(nil)).Elem()
}

func (i ApigInstanceMap) ToApigInstanceMapOutput() ApigInstanceMapOutput {
	return i.ToApigInstanceMapOutputWithContext(context.Background())
}

func (i ApigInstanceMap) ToApigInstanceMapOutputWithContext(ctx context.Context) ApigInstanceMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApigInstanceMapOutput)
}

type ApigInstanceOutput struct{ *pulumi.OutputState }

func (ApigInstanceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ApigInstance)(nil)).Elem()
}

func (o ApigInstanceOutput) ToApigInstanceOutput() ApigInstanceOutput {
	return o
}

func (o ApigInstanceOutput) ToApigInstanceOutputWithContext(ctx context.Context) ApigInstanceOutput {
	return o
}

// schema: Required; The name list of availability zones for the dedicated instance.
func (o ApigInstanceOutput) AvailabilityZones() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ApigInstance) pulumi.StringArrayOutput { return v.AvailabilityZones }).(pulumi.StringArrayOutput)
}

// schema: Deprecated; The name list of availability zones for the dedicated instance.
func (o ApigInstanceOutput) AvailableZones() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ApigInstance) pulumi.StringArrayOutput { return v.AvailableZones }).(pulumi.StringArrayOutput)
}

// The egress bandwidth size of the dedicated instance.
func (o ApigInstanceOutput) BandwidthSize() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ApigInstance) pulumi.IntPtrOutput { return v.BandwidthSize }).(pulumi.IntPtrOutput)
}

// schema: Deprecated; Time when the dedicated instance is created.
//
// Deprecated: Use 'created_at' instead
func (o ApigInstanceOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *ApigInstance) pulumi.StringOutput { return v.CreateTime }).(pulumi.StringOutput)
}

// Time when the dedicated instance is created, in RFC-3339 format.
func (o ApigInstanceOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *ApigInstance) pulumi.StringOutput { return v.CreatedAt }).(pulumi.StringOutput)
}

// Specified the list of the instance custom ingress ports.
func (o ApigInstanceOutput) CustomIngressPorts() ApigInstanceCustomIngressPortArrayOutput {
	return o.ApplyT(func(v *ApigInstance) ApigInstanceCustomIngressPortArrayOutput { return v.CustomIngressPorts }).(ApigInstanceCustomIngressPortArrayOutput)
}

// The description of the dedicated instance.
func (o ApigInstanceOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApigInstance) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// The edition of the dedicated instance.
func (o ApigInstanceOutput) Edition() pulumi.StringOutput {
	return o.ApplyT(func(v *ApigInstance) pulumi.StringOutput { return v.Edition }).(pulumi.StringOutput)
}

// The egress (NAT) public IP address.
func (o ApigInstanceOutput) EgressAddress() pulumi.StringOutput {
	return o.ApplyT(func(v *ApigInstance) pulumi.StringOutput { return v.EgressAddress }).(pulumi.StringOutput)
}

// The EIP ID associated with the dedicated instance.
func (o ApigInstanceOutput) EipId() pulumi.StringOutput {
	return o.ApplyT(func(v *ApigInstance) pulumi.StringOutput { return v.EipId }).(pulumi.StringOutput)
}

// The enterprise project ID to which the dedicated instance belongs.
func (o ApigInstanceOutput) EnterpriseProjectId() pulumi.StringOutput {
	return o.ApplyT(func(v *ApigInstance) pulumi.StringOutput { return v.EnterpriseProjectId }).(pulumi.StringOutput)
}

// The ingress EIP address.
func (o ApigInstanceOutput) IngressAddress() pulumi.StringOutput {
	return o.ApplyT(func(v *ApigInstance) pulumi.StringOutput { return v.IngressAddress }).(pulumi.StringOutput)
}

func (o ApigInstanceOutput) IngressBandwidthChargingMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApigInstance) pulumi.StringPtrOutput { return v.IngressBandwidthChargingMode }).(pulumi.StringPtrOutput)
}

func (o ApigInstanceOutput) IngressBandwidthSize() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ApigInstance) pulumi.IntPtrOutput { return v.IngressBandwidthSize }).(pulumi.IntPtrOutput)
}

// Whether public access with an IPv6 address is supported.
func (o ApigInstanceOutput) Ipv6Enable() pulumi.BoolOutput {
	return o.ApplyT(func(v *ApigInstance) pulumi.BoolOutput { return v.Ipv6Enable }).(pulumi.BoolOutput)
}

// The type of loadbalancer provider used by the instance.
func (o ApigInstanceOutput) LoadbalancerProvider() pulumi.StringOutput {
	return o.ApplyT(func(v *ApigInstance) pulumi.StringOutput { return v.LoadbalancerProvider }).(pulumi.StringOutput)
}

// The start time of the maintenance time window.
func (o ApigInstanceOutput) MaintainBegin() pulumi.StringOutput {
	return o.ApplyT(func(v *ApigInstance) pulumi.StringOutput { return v.MaintainBegin }).(pulumi.StringOutput)
}

// End time of the maintenance time window, 4-hour difference between the start time and end time.
func (o ApigInstanceOutput) MaintainEnd() pulumi.StringOutput {
	return o.ApplyT(func(v *ApigInstance) pulumi.StringOutput { return v.MaintainEnd }).(pulumi.StringOutput)
}

// The name of the dedicated instance.
func (o ApigInstanceOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ApigInstance) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The region in which to create the dedicated instance resource.
func (o ApigInstanceOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v *ApigInstance) pulumi.StringOutput { return v.Region }).(pulumi.StringOutput)
}

// The ID of the security group to which the dedicated instance belongs to.
func (o ApigInstanceOutput) SecurityGroupId() pulumi.StringOutput {
	return o.ApplyT(func(v *ApigInstance) pulumi.StringOutput { return v.SecurityGroupId }).(pulumi.StringOutput)
}

// Status of the dedicated instance.
func (o ApigInstanceOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *ApigInstance) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

// The ID of the VPC subnet used to create the dedicated instance.
func (o ApigInstanceOutput) SubnetId() pulumi.StringOutput {
	return o.ApplyT(func(v *ApigInstance) pulumi.StringOutput { return v.SubnetId }).(pulumi.StringOutput)
}

// The supported features of the dedicated instance.
func (o ApigInstanceOutput) SupportedFeatures() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ApigInstance) pulumi.StringArrayOutput { return v.SupportedFeatures }).(pulumi.StringArrayOutput)
}

func (o ApigInstanceOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *ApigInstance) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// The ID of the VPC used to create the dedicated instance.
func (o ApigInstanceOutput) VpcId() pulumi.StringOutput {
	return o.ApplyT(func(v *ApigInstance) pulumi.StringOutput { return v.VpcId }).(pulumi.StringOutput)
}

// The ingress private IP address of the VPC.
func (o ApigInstanceOutput) VpcIngressAddress() pulumi.StringOutput {
	return o.ApplyT(func(v *ApigInstance) pulumi.StringOutput { return v.VpcIngressAddress }).(pulumi.StringOutput)
}

// The address (full name) of the VPC endpoint service.
func (o ApigInstanceOutput) VpcepServiceAddress() pulumi.StringOutput {
	return o.ApplyT(func(v *ApigInstance) pulumi.StringOutput { return v.VpcepServiceAddress }).(pulumi.StringOutput)
}

// Name of the VPC endpoint service.
func (o ApigInstanceOutput) VpcepServiceName() pulumi.StringOutput {
	return o.ApplyT(func(v *ApigInstance) pulumi.StringOutput { return v.VpcepServiceName }).(pulumi.StringOutput)
}

type ApigInstanceArrayOutput struct{ *pulumi.OutputState }

func (ApigInstanceArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ApigInstance)(nil)).Elem()
}

func (o ApigInstanceArrayOutput) ToApigInstanceArrayOutput() ApigInstanceArrayOutput {
	return o
}

func (o ApigInstanceArrayOutput) ToApigInstanceArrayOutputWithContext(ctx context.Context) ApigInstanceArrayOutput {
	return o
}

func (o ApigInstanceArrayOutput) Index(i pulumi.IntInput) ApigInstanceOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ApigInstance {
		return vs[0].([]*ApigInstance)[vs[1].(int)]
	}).(ApigInstanceOutput)
}

type ApigInstanceMapOutput struct{ *pulumi.OutputState }

func (ApigInstanceMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ApigInstance)(nil)).Elem()
}

func (o ApigInstanceMapOutput) ToApigInstanceMapOutput() ApigInstanceMapOutput {
	return o
}

func (o ApigInstanceMapOutput) ToApigInstanceMapOutputWithContext(ctx context.Context) ApigInstanceMapOutput {
	return o
}

func (o ApigInstanceMapOutput) MapIndex(k pulumi.StringInput) ApigInstanceOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ApigInstance {
		return vs[0].(map[string]*ApigInstance)[vs[1].(string)]
	}).(ApigInstanceOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ApigInstanceInput)(nil)).Elem(), &ApigInstance{})
	pulumi.RegisterInputType(reflect.TypeOf((*ApigInstanceArrayInput)(nil)).Elem(), ApigInstanceArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ApigInstanceMapInput)(nil)).Elem(), ApigInstanceMap{})
	pulumi.RegisterOutputType(ApigInstanceOutput{})
	pulumi.RegisterOutputType(ApigInstanceArrayOutput{})
	pulumi.RegisterOutputType(ApigInstanceMapOutput{})
}
