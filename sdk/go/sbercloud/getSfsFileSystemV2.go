// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package sbercloud

import (
	"context"
	"reflect"

	"github.com/masikrus/pulumi-sbercloud/sdk/go/sbercloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetSfsFileSystemV2(ctx *pulumi.Context, args *GetSfsFileSystemV2Args, opts ...pulumi.InvokeOption) (*GetSfsFileSystemV2Result, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetSfsFileSystemV2Result
	err := ctx.Invoke("sbercloud:index/getSfsFileSystemV2:getSfsFileSystemV2", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getSfsFileSystemV2.
type GetSfsFileSystemV2Args struct {
	Id     *string `pulumi:"id"`
	Name   *string `pulumi:"name"`
	Region *string `pulumi:"region"`
	Status *string `pulumi:"status"`
}

// A collection of values returned by getSfsFileSystemV2.
type GetSfsFileSystemV2Result struct {
	AccessLevel      string            `pulumi:"accessLevel"`
	AccessTo         string            `pulumi:"accessTo"`
	AccessType       string            `pulumi:"accessType"`
	AvailabilityZone string            `pulumi:"availabilityZone"`
	Description      string            `pulumi:"description"`
	ExportLocation   string            `pulumi:"exportLocation"`
	Id               string            `pulumi:"id"`
	IsPublic         bool              `pulumi:"isPublic"`
	Metadata         map[string]string `pulumi:"metadata"`
	MountId          string            `pulumi:"mountId"`
	Name             string            `pulumi:"name"`
	Preferred        bool              `pulumi:"preferred"`
	Region           string            `pulumi:"region"`
	ShareAccessId    string            `pulumi:"shareAccessId"`
	ShareInstanceId  string            `pulumi:"shareInstanceId"`
	ShareProto       string            `pulumi:"shareProto"`
	Size             int               `pulumi:"size"`
	State            string            `pulumi:"state"`
	Status           string            `pulumi:"status"`
}

func GetSfsFileSystemV2Output(ctx *pulumi.Context, args GetSfsFileSystemV2OutputArgs, opts ...pulumi.InvokeOption) GetSfsFileSystemV2ResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (GetSfsFileSystemV2ResultOutput, error) {
			args := v.(GetSfsFileSystemV2Args)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("sbercloud:index/getSfsFileSystemV2:getSfsFileSystemV2", args, GetSfsFileSystemV2ResultOutput{}, options).(GetSfsFileSystemV2ResultOutput), nil
		}).(GetSfsFileSystemV2ResultOutput)
}

// A collection of arguments for invoking getSfsFileSystemV2.
type GetSfsFileSystemV2OutputArgs struct {
	Id     pulumi.StringPtrInput `pulumi:"id"`
	Name   pulumi.StringPtrInput `pulumi:"name"`
	Region pulumi.StringPtrInput `pulumi:"region"`
	Status pulumi.StringPtrInput `pulumi:"status"`
}

func (GetSfsFileSystemV2OutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetSfsFileSystemV2Args)(nil)).Elem()
}

// A collection of values returned by getSfsFileSystemV2.
type GetSfsFileSystemV2ResultOutput struct{ *pulumi.OutputState }

func (GetSfsFileSystemV2ResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetSfsFileSystemV2Result)(nil)).Elem()
}

func (o GetSfsFileSystemV2ResultOutput) ToGetSfsFileSystemV2ResultOutput() GetSfsFileSystemV2ResultOutput {
	return o
}

func (o GetSfsFileSystemV2ResultOutput) ToGetSfsFileSystemV2ResultOutputWithContext(ctx context.Context) GetSfsFileSystemV2ResultOutput {
	return o
}

func (o GetSfsFileSystemV2ResultOutput) AccessLevel() pulumi.StringOutput {
	return o.ApplyT(func(v GetSfsFileSystemV2Result) string { return v.AccessLevel }).(pulumi.StringOutput)
}

func (o GetSfsFileSystemV2ResultOutput) AccessTo() pulumi.StringOutput {
	return o.ApplyT(func(v GetSfsFileSystemV2Result) string { return v.AccessTo }).(pulumi.StringOutput)
}

func (o GetSfsFileSystemV2ResultOutput) AccessType() pulumi.StringOutput {
	return o.ApplyT(func(v GetSfsFileSystemV2Result) string { return v.AccessType }).(pulumi.StringOutput)
}

func (o GetSfsFileSystemV2ResultOutput) AvailabilityZone() pulumi.StringOutput {
	return o.ApplyT(func(v GetSfsFileSystemV2Result) string { return v.AvailabilityZone }).(pulumi.StringOutput)
}

func (o GetSfsFileSystemV2ResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v GetSfsFileSystemV2Result) string { return v.Description }).(pulumi.StringOutput)
}

func (o GetSfsFileSystemV2ResultOutput) ExportLocation() pulumi.StringOutput {
	return o.ApplyT(func(v GetSfsFileSystemV2Result) string { return v.ExportLocation }).(pulumi.StringOutput)
}

func (o GetSfsFileSystemV2ResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetSfsFileSystemV2Result) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetSfsFileSystemV2ResultOutput) IsPublic() pulumi.BoolOutput {
	return o.ApplyT(func(v GetSfsFileSystemV2Result) bool { return v.IsPublic }).(pulumi.BoolOutput)
}

func (o GetSfsFileSystemV2ResultOutput) Metadata() pulumi.StringMapOutput {
	return o.ApplyT(func(v GetSfsFileSystemV2Result) map[string]string { return v.Metadata }).(pulumi.StringMapOutput)
}

func (o GetSfsFileSystemV2ResultOutput) MountId() pulumi.StringOutput {
	return o.ApplyT(func(v GetSfsFileSystemV2Result) string { return v.MountId }).(pulumi.StringOutput)
}

func (o GetSfsFileSystemV2ResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v GetSfsFileSystemV2Result) string { return v.Name }).(pulumi.StringOutput)
}

func (o GetSfsFileSystemV2ResultOutput) Preferred() pulumi.BoolOutput {
	return o.ApplyT(func(v GetSfsFileSystemV2Result) bool { return v.Preferred }).(pulumi.BoolOutput)
}

func (o GetSfsFileSystemV2ResultOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v GetSfsFileSystemV2Result) string { return v.Region }).(pulumi.StringOutput)
}

func (o GetSfsFileSystemV2ResultOutput) ShareAccessId() pulumi.StringOutput {
	return o.ApplyT(func(v GetSfsFileSystemV2Result) string { return v.ShareAccessId }).(pulumi.StringOutput)
}

func (o GetSfsFileSystemV2ResultOutput) ShareInstanceId() pulumi.StringOutput {
	return o.ApplyT(func(v GetSfsFileSystemV2Result) string { return v.ShareInstanceId }).(pulumi.StringOutput)
}

func (o GetSfsFileSystemV2ResultOutput) ShareProto() pulumi.StringOutput {
	return o.ApplyT(func(v GetSfsFileSystemV2Result) string { return v.ShareProto }).(pulumi.StringOutput)
}

func (o GetSfsFileSystemV2ResultOutput) Size() pulumi.IntOutput {
	return o.ApplyT(func(v GetSfsFileSystemV2Result) int { return v.Size }).(pulumi.IntOutput)
}

func (o GetSfsFileSystemV2ResultOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v GetSfsFileSystemV2Result) string { return v.State }).(pulumi.StringOutput)
}

func (o GetSfsFileSystemV2ResultOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v GetSfsFileSystemV2Result) string { return v.Status }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetSfsFileSystemV2ResultOutput{})
}
