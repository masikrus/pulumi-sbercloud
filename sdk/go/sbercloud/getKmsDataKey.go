// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package sbercloud

import (
	"context"
	"reflect"

	"github.com/masikrus/pulumi-sbercloud/sdk/go/sbercloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetKmsDataKey(ctx *pulumi.Context, args *GetKmsDataKeyArgs, opts ...pulumi.InvokeOption) (*GetKmsDataKeyResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetKmsDataKeyResult
	err := ctx.Invoke("sbercloud:index/getKmsDataKey:getKmsDataKey", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getKmsDataKey.
type GetKmsDataKeyArgs struct {
	DatakeyLength     string  `pulumi:"datakeyLength"`
	EncryptionContext *string `pulumi:"encryptionContext"`
	KeyId             string  `pulumi:"keyId"`
	Region            *string `pulumi:"region"`
}

// A collection of values returned by getKmsDataKey.
type GetKmsDataKeyResult struct {
	CipherText        string  `pulumi:"cipherText"`
	DatakeyLength     string  `pulumi:"datakeyLength"`
	EncryptionContext *string `pulumi:"encryptionContext"`
	// The provider-assigned unique ID for this managed resource.
	Id        string `pulumi:"id"`
	KeyId     string `pulumi:"keyId"`
	PlainText string `pulumi:"plainText"`
	Region    string `pulumi:"region"`
}

func GetKmsDataKeyOutput(ctx *pulumi.Context, args GetKmsDataKeyOutputArgs, opts ...pulumi.InvokeOption) GetKmsDataKeyResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (GetKmsDataKeyResultOutput, error) {
			args := v.(GetKmsDataKeyArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("sbercloud:index/getKmsDataKey:getKmsDataKey", args, GetKmsDataKeyResultOutput{}, options).(GetKmsDataKeyResultOutput), nil
		}).(GetKmsDataKeyResultOutput)
}

// A collection of arguments for invoking getKmsDataKey.
type GetKmsDataKeyOutputArgs struct {
	DatakeyLength     pulumi.StringInput    `pulumi:"datakeyLength"`
	EncryptionContext pulumi.StringPtrInput `pulumi:"encryptionContext"`
	KeyId             pulumi.StringInput    `pulumi:"keyId"`
	Region            pulumi.StringPtrInput `pulumi:"region"`
}

func (GetKmsDataKeyOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetKmsDataKeyArgs)(nil)).Elem()
}

// A collection of values returned by getKmsDataKey.
type GetKmsDataKeyResultOutput struct{ *pulumi.OutputState }

func (GetKmsDataKeyResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetKmsDataKeyResult)(nil)).Elem()
}

func (o GetKmsDataKeyResultOutput) ToGetKmsDataKeyResultOutput() GetKmsDataKeyResultOutput {
	return o
}

func (o GetKmsDataKeyResultOutput) ToGetKmsDataKeyResultOutputWithContext(ctx context.Context) GetKmsDataKeyResultOutput {
	return o
}

func (o GetKmsDataKeyResultOutput) CipherText() pulumi.StringOutput {
	return o.ApplyT(func(v GetKmsDataKeyResult) string { return v.CipherText }).(pulumi.StringOutput)
}

func (o GetKmsDataKeyResultOutput) DatakeyLength() pulumi.StringOutput {
	return o.ApplyT(func(v GetKmsDataKeyResult) string { return v.DatakeyLength }).(pulumi.StringOutput)
}

func (o GetKmsDataKeyResultOutput) EncryptionContext() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetKmsDataKeyResult) *string { return v.EncryptionContext }).(pulumi.StringPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetKmsDataKeyResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetKmsDataKeyResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetKmsDataKeyResultOutput) KeyId() pulumi.StringOutput {
	return o.ApplyT(func(v GetKmsDataKeyResult) string { return v.KeyId }).(pulumi.StringOutput)
}

func (o GetKmsDataKeyResultOutput) PlainText() pulumi.StringOutput {
	return o.ApplyT(func(v GetKmsDataKeyResult) string { return v.PlainText }).(pulumi.StringOutput)
}

func (o GetKmsDataKeyResultOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v GetKmsDataKeyResult) string { return v.Region }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetKmsDataKeyResultOutput{})
}
