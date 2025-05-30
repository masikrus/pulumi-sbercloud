// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package sbercloud

import (
	"context"
	"reflect"

	"github.com/masikrus/pulumi-sbercloud/sdk/go/sbercloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type VpnCustomerGateway struct {
	pulumi.CustomResourceState

	// The BGP ASN number of the customer gateway, the default value is 65000.
	Asn                pulumi.IntPtrOutput    `pulumi:"asn"`
	CertificateContent pulumi.StringPtrOutput `pulumi:"certificateContent"`
	CertificateId      pulumi.StringOutput    `pulumi:"certificateId"`
	// The create time.
	CreatedAt  pulumi.StringOutput `pulumi:"createdAt"`
	ExpireTime pulumi.StringOutput `pulumi:"expireTime"`
	// The identifier type of a customer gateway.
	IdType pulumi.StringPtrOutput `pulumi:"idType"`
	// The identifier of a customer gateway.
	IdValue pulumi.StringOutput `pulumi:"idValue"`
	// The IP address of the customer gateway.
	Ip          pulumi.StringPtrOutput `pulumi:"ip"`
	IsUpdatable pulumi.BoolOutput      `pulumi:"isUpdatable"`
	Issuer      pulumi.StringOutput    `pulumi:"issuer"`
	// The customer gateway name.
	Name   pulumi.StringOutput `pulumi:"name"`
	Region pulumi.StringOutput `pulumi:"region"`
	// The route mode of the customer gateway.
	RouteMode          pulumi.StringPtrOutput `pulumi:"routeMode"`
	SerialNumber       pulumi.StringOutput    `pulumi:"serialNumber"`
	SignatureAlgorithm pulumi.StringOutput    `pulumi:"signatureAlgorithm"`
	Subject            pulumi.StringOutput    `pulumi:"subject"`
	Tags               pulumi.StringMapOutput `pulumi:"tags"`
	// The update time.
	UpdatedAt pulumi.StringOutput `pulumi:"updatedAt"`
}

// NewVpnCustomerGateway registers a new resource with the given unique name, arguments, and options.
func NewVpnCustomerGateway(ctx *pulumi.Context,
	name string, args *VpnCustomerGatewayArgs, opts ...pulumi.ResourceOption) (*VpnCustomerGateway, error) {
	if args == nil {
		args = &VpnCustomerGatewayArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource VpnCustomerGateway
	err := ctx.RegisterResource("sbercloud:index/vpnCustomerGateway:VpnCustomerGateway", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetVpnCustomerGateway gets an existing VpnCustomerGateway resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVpnCustomerGateway(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VpnCustomerGatewayState, opts ...pulumi.ResourceOption) (*VpnCustomerGateway, error) {
	var resource VpnCustomerGateway
	err := ctx.ReadResource("sbercloud:index/vpnCustomerGateway:VpnCustomerGateway", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering VpnCustomerGateway resources.
type vpnCustomerGatewayState struct {
	// The BGP ASN number of the customer gateway, the default value is 65000.
	Asn                *int    `pulumi:"asn"`
	CertificateContent *string `pulumi:"certificateContent"`
	CertificateId      *string `pulumi:"certificateId"`
	// The create time.
	CreatedAt  *string `pulumi:"createdAt"`
	ExpireTime *string `pulumi:"expireTime"`
	// The identifier type of a customer gateway.
	IdType *string `pulumi:"idType"`
	// The identifier of a customer gateway.
	IdValue *string `pulumi:"idValue"`
	// The IP address of the customer gateway.
	Ip          *string `pulumi:"ip"`
	IsUpdatable *bool   `pulumi:"isUpdatable"`
	Issuer      *string `pulumi:"issuer"`
	// The customer gateway name.
	Name   *string `pulumi:"name"`
	Region *string `pulumi:"region"`
	// The route mode of the customer gateway.
	RouteMode          *string           `pulumi:"routeMode"`
	SerialNumber       *string           `pulumi:"serialNumber"`
	SignatureAlgorithm *string           `pulumi:"signatureAlgorithm"`
	Subject            *string           `pulumi:"subject"`
	Tags               map[string]string `pulumi:"tags"`
	// The update time.
	UpdatedAt *string `pulumi:"updatedAt"`
}

type VpnCustomerGatewayState struct {
	// The BGP ASN number of the customer gateway, the default value is 65000.
	Asn                pulumi.IntPtrInput
	CertificateContent pulumi.StringPtrInput
	CertificateId      pulumi.StringPtrInput
	// The create time.
	CreatedAt  pulumi.StringPtrInput
	ExpireTime pulumi.StringPtrInput
	// The identifier type of a customer gateway.
	IdType pulumi.StringPtrInput
	// The identifier of a customer gateway.
	IdValue pulumi.StringPtrInput
	// The IP address of the customer gateway.
	Ip          pulumi.StringPtrInput
	IsUpdatable pulumi.BoolPtrInput
	Issuer      pulumi.StringPtrInput
	// The customer gateway name.
	Name   pulumi.StringPtrInput
	Region pulumi.StringPtrInput
	// The route mode of the customer gateway.
	RouteMode          pulumi.StringPtrInput
	SerialNumber       pulumi.StringPtrInput
	SignatureAlgorithm pulumi.StringPtrInput
	Subject            pulumi.StringPtrInput
	Tags               pulumi.StringMapInput
	// The update time.
	UpdatedAt pulumi.StringPtrInput
}

func (VpnCustomerGatewayState) ElementType() reflect.Type {
	return reflect.TypeOf((*vpnCustomerGatewayState)(nil)).Elem()
}

type vpnCustomerGatewayArgs struct {
	// The BGP ASN number of the customer gateway, the default value is 65000.
	Asn                *int    `pulumi:"asn"`
	CertificateContent *string `pulumi:"certificateContent"`
	// The identifier type of a customer gateway.
	IdType *string `pulumi:"idType"`
	// The identifier of a customer gateway.
	IdValue *string `pulumi:"idValue"`
	// The IP address of the customer gateway.
	Ip *string `pulumi:"ip"`
	// The customer gateway name.
	Name   *string `pulumi:"name"`
	Region *string `pulumi:"region"`
	// The route mode of the customer gateway.
	RouteMode *string           `pulumi:"routeMode"`
	Tags      map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a VpnCustomerGateway resource.
type VpnCustomerGatewayArgs struct {
	// The BGP ASN number of the customer gateway, the default value is 65000.
	Asn                pulumi.IntPtrInput
	CertificateContent pulumi.StringPtrInput
	// The identifier type of a customer gateway.
	IdType pulumi.StringPtrInput
	// The identifier of a customer gateway.
	IdValue pulumi.StringPtrInput
	// The IP address of the customer gateway.
	Ip pulumi.StringPtrInput
	// The customer gateway name.
	Name   pulumi.StringPtrInput
	Region pulumi.StringPtrInput
	// The route mode of the customer gateway.
	RouteMode pulumi.StringPtrInput
	Tags      pulumi.StringMapInput
}

func (VpnCustomerGatewayArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*vpnCustomerGatewayArgs)(nil)).Elem()
}

type VpnCustomerGatewayInput interface {
	pulumi.Input

	ToVpnCustomerGatewayOutput() VpnCustomerGatewayOutput
	ToVpnCustomerGatewayOutputWithContext(ctx context.Context) VpnCustomerGatewayOutput
}

func (*VpnCustomerGateway) ElementType() reflect.Type {
	return reflect.TypeOf((**VpnCustomerGateway)(nil)).Elem()
}

func (i *VpnCustomerGateway) ToVpnCustomerGatewayOutput() VpnCustomerGatewayOutput {
	return i.ToVpnCustomerGatewayOutputWithContext(context.Background())
}

func (i *VpnCustomerGateway) ToVpnCustomerGatewayOutputWithContext(ctx context.Context) VpnCustomerGatewayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpnCustomerGatewayOutput)
}

// VpnCustomerGatewayArrayInput is an input type that accepts VpnCustomerGatewayArray and VpnCustomerGatewayArrayOutput values.
// You can construct a concrete instance of `VpnCustomerGatewayArrayInput` via:
//
//	VpnCustomerGatewayArray{ VpnCustomerGatewayArgs{...} }
type VpnCustomerGatewayArrayInput interface {
	pulumi.Input

	ToVpnCustomerGatewayArrayOutput() VpnCustomerGatewayArrayOutput
	ToVpnCustomerGatewayArrayOutputWithContext(context.Context) VpnCustomerGatewayArrayOutput
}

type VpnCustomerGatewayArray []VpnCustomerGatewayInput

func (VpnCustomerGatewayArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*VpnCustomerGateway)(nil)).Elem()
}

func (i VpnCustomerGatewayArray) ToVpnCustomerGatewayArrayOutput() VpnCustomerGatewayArrayOutput {
	return i.ToVpnCustomerGatewayArrayOutputWithContext(context.Background())
}

func (i VpnCustomerGatewayArray) ToVpnCustomerGatewayArrayOutputWithContext(ctx context.Context) VpnCustomerGatewayArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpnCustomerGatewayArrayOutput)
}

// VpnCustomerGatewayMapInput is an input type that accepts VpnCustomerGatewayMap and VpnCustomerGatewayMapOutput values.
// You can construct a concrete instance of `VpnCustomerGatewayMapInput` via:
//
//	VpnCustomerGatewayMap{ "key": VpnCustomerGatewayArgs{...} }
type VpnCustomerGatewayMapInput interface {
	pulumi.Input

	ToVpnCustomerGatewayMapOutput() VpnCustomerGatewayMapOutput
	ToVpnCustomerGatewayMapOutputWithContext(context.Context) VpnCustomerGatewayMapOutput
}

type VpnCustomerGatewayMap map[string]VpnCustomerGatewayInput

func (VpnCustomerGatewayMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*VpnCustomerGateway)(nil)).Elem()
}

func (i VpnCustomerGatewayMap) ToVpnCustomerGatewayMapOutput() VpnCustomerGatewayMapOutput {
	return i.ToVpnCustomerGatewayMapOutputWithContext(context.Background())
}

func (i VpnCustomerGatewayMap) ToVpnCustomerGatewayMapOutputWithContext(ctx context.Context) VpnCustomerGatewayMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpnCustomerGatewayMapOutput)
}

type VpnCustomerGatewayOutput struct{ *pulumi.OutputState }

func (VpnCustomerGatewayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VpnCustomerGateway)(nil)).Elem()
}

func (o VpnCustomerGatewayOutput) ToVpnCustomerGatewayOutput() VpnCustomerGatewayOutput {
	return o
}

func (o VpnCustomerGatewayOutput) ToVpnCustomerGatewayOutputWithContext(ctx context.Context) VpnCustomerGatewayOutput {
	return o
}

// The BGP ASN number of the customer gateway, the default value is 65000.
func (o VpnCustomerGatewayOutput) Asn() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *VpnCustomerGateway) pulumi.IntPtrOutput { return v.Asn }).(pulumi.IntPtrOutput)
}

func (o VpnCustomerGatewayOutput) CertificateContent() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VpnCustomerGateway) pulumi.StringPtrOutput { return v.CertificateContent }).(pulumi.StringPtrOutput)
}

func (o VpnCustomerGatewayOutput) CertificateId() pulumi.StringOutput {
	return o.ApplyT(func(v *VpnCustomerGateway) pulumi.StringOutput { return v.CertificateId }).(pulumi.StringOutput)
}

// The create time.
func (o VpnCustomerGatewayOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *VpnCustomerGateway) pulumi.StringOutput { return v.CreatedAt }).(pulumi.StringOutput)
}

func (o VpnCustomerGatewayOutput) ExpireTime() pulumi.StringOutput {
	return o.ApplyT(func(v *VpnCustomerGateway) pulumi.StringOutput { return v.ExpireTime }).(pulumi.StringOutput)
}

// The identifier type of a customer gateway.
func (o VpnCustomerGatewayOutput) IdType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VpnCustomerGateway) pulumi.StringPtrOutput { return v.IdType }).(pulumi.StringPtrOutput)
}

// The identifier of a customer gateway.
func (o VpnCustomerGatewayOutput) IdValue() pulumi.StringOutput {
	return o.ApplyT(func(v *VpnCustomerGateway) pulumi.StringOutput { return v.IdValue }).(pulumi.StringOutput)
}

// The IP address of the customer gateway.
func (o VpnCustomerGatewayOutput) Ip() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VpnCustomerGateway) pulumi.StringPtrOutput { return v.Ip }).(pulumi.StringPtrOutput)
}

func (o VpnCustomerGatewayOutput) IsUpdatable() pulumi.BoolOutput {
	return o.ApplyT(func(v *VpnCustomerGateway) pulumi.BoolOutput { return v.IsUpdatable }).(pulumi.BoolOutput)
}

func (o VpnCustomerGatewayOutput) Issuer() pulumi.StringOutput {
	return o.ApplyT(func(v *VpnCustomerGateway) pulumi.StringOutput { return v.Issuer }).(pulumi.StringOutput)
}

// The customer gateway name.
func (o VpnCustomerGatewayOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *VpnCustomerGateway) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o VpnCustomerGatewayOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v *VpnCustomerGateway) pulumi.StringOutput { return v.Region }).(pulumi.StringOutput)
}

// The route mode of the customer gateway.
func (o VpnCustomerGatewayOutput) RouteMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VpnCustomerGateway) pulumi.StringPtrOutput { return v.RouteMode }).(pulumi.StringPtrOutput)
}

func (o VpnCustomerGatewayOutput) SerialNumber() pulumi.StringOutput {
	return o.ApplyT(func(v *VpnCustomerGateway) pulumi.StringOutput { return v.SerialNumber }).(pulumi.StringOutput)
}

func (o VpnCustomerGatewayOutput) SignatureAlgorithm() pulumi.StringOutput {
	return o.ApplyT(func(v *VpnCustomerGateway) pulumi.StringOutput { return v.SignatureAlgorithm }).(pulumi.StringOutput)
}

func (o VpnCustomerGatewayOutput) Subject() pulumi.StringOutput {
	return o.ApplyT(func(v *VpnCustomerGateway) pulumi.StringOutput { return v.Subject }).(pulumi.StringOutput)
}

func (o VpnCustomerGatewayOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *VpnCustomerGateway) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// The update time.
func (o VpnCustomerGatewayOutput) UpdatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *VpnCustomerGateway) pulumi.StringOutput { return v.UpdatedAt }).(pulumi.StringOutput)
}

type VpnCustomerGatewayArrayOutput struct{ *pulumi.OutputState }

func (VpnCustomerGatewayArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*VpnCustomerGateway)(nil)).Elem()
}

func (o VpnCustomerGatewayArrayOutput) ToVpnCustomerGatewayArrayOutput() VpnCustomerGatewayArrayOutput {
	return o
}

func (o VpnCustomerGatewayArrayOutput) ToVpnCustomerGatewayArrayOutputWithContext(ctx context.Context) VpnCustomerGatewayArrayOutput {
	return o
}

func (o VpnCustomerGatewayArrayOutput) Index(i pulumi.IntInput) VpnCustomerGatewayOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *VpnCustomerGateway {
		return vs[0].([]*VpnCustomerGateway)[vs[1].(int)]
	}).(VpnCustomerGatewayOutput)
}

type VpnCustomerGatewayMapOutput struct{ *pulumi.OutputState }

func (VpnCustomerGatewayMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*VpnCustomerGateway)(nil)).Elem()
}

func (o VpnCustomerGatewayMapOutput) ToVpnCustomerGatewayMapOutput() VpnCustomerGatewayMapOutput {
	return o
}

func (o VpnCustomerGatewayMapOutput) ToVpnCustomerGatewayMapOutputWithContext(ctx context.Context) VpnCustomerGatewayMapOutput {
	return o
}

func (o VpnCustomerGatewayMapOutput) MapIndex(k pulumi.StringInput) VpnCustomerGatewayOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *VpnCustomerGateway {
		return vs[0].(map[string]*VpnCustomerGateway)[vs[1].(string)]
	}).(VpnCustomerGatewayOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*VpnCustomerGatewayInput)(nil)).Elem(), &VpnCustomerGateway{})
	pulumi.RegisterInputType(reflect.TypeOf((*VpnCustomerGatewayArrayInput)(nil)).Elem(), VpnCustomerGatewayArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*VpnCustomerGatewayMapInput)(nil)).Elem(), VpnCustomerGatewayMap{})
	pulumi.RegisterOutputType(VpnCustomerGatewayOutput{})
	pulumi.RegisterOutputType(VpnCustomerGatewayArrayOutput{})
	pulumi.RegisterOutputType(VpnCustomerGatewayMapOutput{})
}
