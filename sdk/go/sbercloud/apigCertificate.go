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

type ApigCertificate struct {
	pulumi.CustomResourceState

	// The certificate content.
	Content pulumi.StringOutput `pulumi:"content"`
	// The effective time of the certificate.
	EffectedAt pulumi.StringOutput `pulumi:"effectedAt"`
	// The expiration time of the certificate.
	ExpiresAt pulumi.StringOutput `pulumi:"expiresAt"`
	// The dedicated instance ID to which the certificate belongs.
	InstanceId pulumi.StringOutput `pulumi:"instanceId"`
	// The certificate name.
	Name pulumi.StringOutput `pulumi:"name"`
	// The private key of the certificate.
	PrivateKey pulumi.StringOutput `pulumi:"privateKey"`
	// The region where the certificate is located.
	Region pulumi.StringOutput `pulumi:"region"`
	// The SAN (Subject Alternative Names) of the certificate.
	Sans pulumi.StringArrayOutput `pulumi:"sans"`
	// What signature algorithm the certificate uses.
	SignatureAlgorithm pulumi.StringOutput `pulumi:"signatureAlgorithm"`
	// The trusted root CA certificate.
	TrustedRootCa pulumi.StringPtrOutput `pulumi:"trustedRootCa"`
	// The certificate type.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewApigCertificate registers a new resource with the given unique name, arguments, and options.
func NewApigCertificate(ctx *pulumi.Context,
	name string, args *ApigCertificateArgs, opts ...pulumi.ResourceOption) (*ApigCertificate, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Content == nil {
		return nil, errors.New("invalid value for required argument 'Content'")
	}
	if args.PrivateKey == nil {
		return nil, errors.New("invalid value for required argument 'PrivateKey'")
	}
	if args.Content != nil {
		args.Content = pulumi.ToSecret(args.Content).(pulumi.StringInput)
	}
	if args.PrivateKey != nil {
		args.PrivateKey = pulumi.ToSecret(args.PrivateKey).(pulumi.StringInput)
	}
	if args.TrustedRootCa != nil {
		args.TrustedRootCa = pulumi.ToSecret(args.TrustedRootCa).(pulumi.StringPtrInput)
	}
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"content",
		"privateKey",
		"trustedRootCa",
	})
	opts = append(opts, secrets)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource ApigCertificate
	err := ctx.RegisterResource("sbercloud:index/apigCertificate:ApigCertificate", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetApigCertificate gets an existing ApigCertificate resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetApigCertificate(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ApigCertificateState, opts ...pulumi.ResourceOption) (*ApigCertificate, error) {
	var resource ApigCertificate
	err := ctx.ReadResource("sbercloud:index/apigCertificate:ApigCertificate", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ApigCertificate resources.
type apigCertificateState struct {
	// The certificate content.
	Content *string `pulumi:"content"`
	// The effective time of the certificate.
	EffectedAt *string `pulumi:"effectedAt"`
	// The expiration time of the certificate.
	ExpiresAt *string `pulumi:"expiresAt"`
	// The dedicated instance ID to which the certificate belongs.
	InstanceId *string `pulumi:"instanceId"`
	// The certificate name.
	Name *string `pulumi:"name"`
	// The private key of the certificate.
	PrivateKey *string `pulumi:"privateKey"`
	// The region where the certificate is located.
	Region *string `pulumi:"region"`
	// The SAN (Subject Alternative Names) of the certificate.
	Sans []string `pulumi:"sans"`
	// What signature algorithm the certificate uses.
	SignatureAlgorithm *string `pulumi:"signatureAlgorithm"`
	// The trusted root CA certificate.
	TrustedRootCa *string `pulumi:"trustedRootCa"`
	// The certificate type.
	Type *string `pulumi:"type"`
}

type ApigCertificateState struct {
	// The certificate content.
	Content pulumi.StringPtrInput
	// The effective time of the certificate.
	EffectedAt pulumi.StringPtrInput
	// The expiration time of the certificate.
	ExpiresAt pulumi.StringPtrInput
	// The dedicated instance ID to which the certificate belongs.
	InstanceId pulumi.StringPtrInput
	// The certificate name.
	Name pulumi.StringPtrInput
	// The private key of the certificate.
	PrivateKey pulumi.StringPtrInput
	// The region where the certificate is located.
	Region pulumi.StringPtrInput
	// The SAN (Subject Alternative Names) of the certificate.
	Sans pulumi.StringArrayInput
	// What signature algorithm the certificate uses.
	SignatureAlgorithm pulumi.StringPtrInput
	// The trusted root CA certificate.
	TrustedRootCa pulumi.StringPtrInput
	// The certificate type.
	Type pulumi.StringPtrInput
}

func (ApigCertificateState) ElementType() reflect.Type {
	return reflect.TypeOf((*apigCertificateState)(nil)).Elem()
}

type apigCertificateArgs struct {
	// The certificate content.
	Content string `pulumi:"content"`
	// The dedicated instance ID to which the certificate belongs.
	InstanceId *string `pulumi:"instanceId"`
	// The certificate name.
	Name *string `pulumi:"name"`
	// The private key of the certificate.
	PrivateKey string `pulumi:"privateKey"`
	// The region where the certificate is located.
	Region *string `pulumi:"region"`
	// The trusted root CA certificate.
	TrustedRootCa *string `pulumi:"trustedRootCa"`
	// The certificate type.
	Type *string `pulumi:"type"`
}

// The set of arguments for constructing a ApigCertificate resource.
type ApigCertificateArgs struct {
	// The certificate content.
	Content pulumi.StringInput
	// The dedicated instance ID to which the certificate belongs.
	InstanceId pulumi.StringPtrInput
	// The certificate name.
	Name pulumi.StringPtrInput
	// The private key of the certificate.
	PrivateKey pulumi.StringInput
	// The region where the certificate is located.
	Region pulumi.StringPtrInput
	// The trusted root CA certificate.
	TrustedRootCa pulumi.StringPtrInput
	// The certificate type.
	Type pulumi.StringPtrInput
}

func (ApigCertificateArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*apigCertificateArgs)(nil)).Elem()
}

type ApigCertificateInput interface {
	pulumi.Input

	ToApigCertificateOutput() ApigCertificateOutput
	ToApigCertificateOutputWithContext(ctx context.Context) ApigCertificateOutput
}

func (*ApigCertificate) ElementType() reflect.Type {
	return reflect.TypeOf((**ApigCertificate)(nil)).Elem()
}

func (i *ApigCertificate) ToApigCertificateOutput() ApigCertificateOutput {
	return i.ToApigCertificateOutputWithContext(context.Background())
}

func (i *ApigCertificate) ToApigCertificateOutputWithContext(ctx context.Context) ApigCertificateOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApigCertificateOutput)
}

// ApigCertificateArrayInput is an input type that accepts ApigCertificateArray and ApigCertificateArrayOutput values.
// You can construct a concrete instance of `ApigCertificateArrayInput` via:
//
//	ApigCertificateArray{ ApigCertificateArgs{...} }
type ApigCertificateArrayInput interface {
	pulumi.Input

	ToApigCertificateArrayOutput() ApigCertificateArrayOutput
	ToApigCertificateArrayOutputWithContext(context.Context) ApigCertificateArrayOutput
}

type ApigCertificateArray []ApigCertificateInput

func (ApigCertificateArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ApigCertificate)(nil)).Elem()
}

func (i ApigCertificateArray) ToApigCertificateArrayOutput() ApigCertificateArrayOutput {
	return i.ToApigCertificateArrayOutputWithContext(context.Background())
}

func (i ApigCertificateArray) ToApigCertificateArrayOutputWithContext(ctx context.Context) ApigCertificateArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApigCertificateArrayOutput)
}

// ApigCertificateMapInput is an input type that accepts ApigCertificateMap and ApigCertificateMapOutput values.
// You can construct a concrete instance of `ApigCertificateMapInput` via:
//
//	ApigCertificateMap{ "key": ApigCertificateArgs{...} }
type ApigCertificateMapInput interface {
	pulumi.Input

	ToApigCertificateMapOutput() ApigCertificateMapOutput
	ToApigCertificateMapOutputWithContext(context.Context) ApigCertificateMapOutput
}

type ApigCertificateMap map[string]ApigCertificateInput

func (ApigCertificateMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ApigCertificate)(nil)).Elem()
}

func (i ApigCertificateMap) ToApigCertificateMapOutput() ApigCertificateMapOutput {
	return i.ToApigCertificateMapOutputWithContext(context.Background())
}

func (i ApigCertificateMap) ToApigCertificateMapOutputWithContext(ctx context.Context) ApigCertificateMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApigCertificateMapOutput)
}

type ApigCertificateOutput struct{ *pulumi.OutputState }

func (ApigCertificateOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ApigCertificate)(nil)).Elem()
}

func (o ApigCertificateOutput) ToApigCertificateOutput() ApigCertificateOutput {
	return o
}

func (o ApigCertificateOutput) ToApigCertificateOutputWithContext(ctx context.Context) ApigCertificateOutput {
	return o
}

// The certificate content.
func (o ApigCertificateOutput) Content() pulumi.StringOutput {
	return o.ApplyT(func(v *ApigCertificate) pulumi.StringOutput { return v.Content }).(pulumi.StringOutput)
}

// The effective time of the certificate.
func (o ApigCertificateOutput) EffectedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *ApigCertificate) pulumi.StringOutput { return v.EffectedAt }).(pulumi.StringOutput)
}

// The expiration time of the certificate.
func (o ApigCertificateOutput) ExpiresAt() pulumi.StringOutput {
	return o.ApplyT(func(v *ApigCertificate) pulumi.StringOutput { return v.ExpiresAt }).(pulumi.StringOutput)
}

// The dedicated instance ID to which the certificate belongs.
func (o ApigCertificateOutput) InstanceId() pulumi.StringOutput {
	return o.ApplyT(func(v *ApigCertificate) pulumi.StringOutput { return v.InstanceId }).(pulumi.StringOutput)
}

// The certificate name.
func (o ApigCertificateOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ApigCertificate) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The private key of the certificate.
func (o ApigCertificateOutput) PrivateKey() pulumi.StringOutput {
	return o.ApplyT(func(v *ApigCertificate) pulumi.StringOutput { return v.PrivateKey }).(pulumi.StringOutput)
}

// The region where the certificate is located.
func (o ApigCertificateOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v *ApigCertificate) pulumi.StringOutput { return v.Region }).(pulumi.StringOutput)
}

// The SAN (Subject Alternative Names) of the certificate.
func (o ApigCertificateOutput) Sans() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ApigCertificate) pulumi.StringArrayOutput { return v.Sans }).(pulumi.StringArrayOutput)
}

// What signature algorithm the certificate uses.
func (o ApigCertificateOutput) SignatureAlgorithm() pulumi.StringOutput {
	return o.ApplyT(func(v *ApigCertificate) pulumi.StringOutput { return v.SignatureAlgorithm }).(pulumi.StringOutput)
}

// The trusted root CA certificate.
func (o ApigCertificateOutput) TrustedRootCa() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApigCertificate) pulumi.StringPtrOutput { return v.TrustedRootCa }).(pulumi.StringPtrOutput)
}

// The certificate type.
func (o ApigCertificateOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *ApigCertificate) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

type ApigCertificateArrayOutput struct{ *pulumi.OutputState }

func (ApigCertificateArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ApigCertificate)(nil)).Elem()
}

func (o ApigCertificateArrayOutput) ToApigCertificateArrayOutput() ApigCertificateArrayOutput {
	return o
}

func (o ApigCertificateArrayOutput) ToApigCertificateArrayOutputWithContext(ctx context.Context) ApigCertificateArrayOutput {
	return o
}

func (o ApigCertificateArrayOutput) Index(i pulumi.IntInput) ApigCertificateOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ApigCertificate {
		return vs[0].([]*ApigCertificate)[vs[1].(int)]
	}).(ApigCertificateOutput)
}

type ApigCertificateMapOutput struct{ *pulumi.OutputState }

func (ApigCertificateMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ApigCertificate)(nil)).Elem()
}

func (o ApigCertificateMapOutput) ToApigCertificateMapOutput() ApigCertificateMapOutput {
	return o
}

func (o ApigCertificateMapOutput) ToApigCertificateMapOutputWithContext(ctx context.Context) ApigCertificateMapOutput {
	return o
}

func (o ApigCertificateMapOutput) MapIndex(k pulumi.StringInput) ApigCertificateOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ApigCertificate {
		return vs[0].(map[string]*ApigCertificate)[vs[1].(string)]
	}).(ApigCertificateOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ApigCertificateInput)(nil)).Elem(), &ApigCertificate{})
	pulumi.RegisterInputType(reflect.TypeOf((*ApigCertificateArrayInput)(nil)).Elem(), ApigCertificateArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ApigCertificateMapInput)(nil)).Elem(), ApigCertificateMap{})
	pulumi.RegisterOutputType(ApigCertificateOutput{})
	pulumi.RegisterOutputType(ApigCertificateArrayOutput{})
	pulumi.RegisterOutputType(ApigCertificateMapOutput{})
}
