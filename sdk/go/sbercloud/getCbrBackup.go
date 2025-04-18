// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package sbercloud

import (
	"context"
	"reflect"

	"github.com/masikrus/pulumi-sbercloud/sdk/go/sbercloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetCbrBackup(ctx *pulumi.Context, args *GetCbrBackupArgs, opts ...pulumi.InvokeOption) (*GetCbrBackupResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetCbrBackupResult
	err := ctx.Invoke("sbercloud:index/getCbrBackup:getCbrBackup", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getCbrBackup.
type GetCbrBackupArgs struct {
	Id     string  `pulumi:"id"`
	Region *string `pulumi:"region"`
}

// A collection of values returned by getCbrBackup.
type GetCbrBackupResult struct {
	CheckpointId        string                          `pulumi:"checkpointId"`
	Childrens           []GetCbrBackupChildren          `pulumi:"childrens"`
	CreatedAt           string                          `pulumi:"createdAt"`
	Description         string                          `pulumi:"description"`
	EnterpriseProjectId string                          `pulumi:"enterpriseProjectId"`
	ExpiredAt           string                          `pulumi:"expiredAt"`
	ExtendInfos         []GetCbrBackupExtendInfo        `pulumi:"extendInfos"`
	Id                  string                          `pulumi:"id"`
	Name                string                          `pulumi:"name"`
	ParentId            string                          `pulumi:"parentId"`
	Region              *string                         `pulumi:"region"`
	ReplicationRecords  []GetCbrBackupReplicationRecord `pulumi:"replicationRecords"`
	ResourceAz          string                          `pulumi:"resourceAz"`
	ResourceId          string                          `pulumi:"resourceId"`
	ResourceName        string                          `pulumi:"resourceName"`
	ResourceSize        int                             `pulumi:"resourceSize"`
	ResourceType        string                          `pulumi:"resourceType"`
	Status              string                          `pulumi:"status"`
	Type                string                          `pulumi:"type"`
	UpdatedAt           string                          `pulumi:"updatedAt"`
	VaultId             string                          `pulumi:"vaultId"`
}

func GetCbrBackupOutput(ctx *pulumi.Context, args GetCbrBackupOutputArgs, opts ...pulumi.InvokeOption) GetCbrBackupResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (GetCbrBackupResultOutput, error) {
			args := v.(GetCbrBackupArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("sbercloud:index/getCbrBackup:getCbrBackup", args, GetCbrBackupResultOutput{}, options).(GetCbrBackupResultOutput), nil
		}).(GetCbrBackupResultOutput)
}

// A collection of arguments for invoking getCbrBackup.
type GetCbrBackupOutputArgs struct {
	Id     pulumi.StringInput    `pulumi:"id"`
	Region pulumi.StringPtrInput `pulumi:"region"`
}

func (GetCbrBackupOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetCbrBackupArgs)(nil)).Elem()
}

// A collection of values returned by getCbrBackup.
type GetCbrBackupResultOutput struct{ *pulumi.OutputState }

func (GetCbrBackupResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetCbrBackupResult)(nil)).Elem()
}

func (o GetCbrBackupResultOutput) ToGetCbrBackupResultOutput() GetCbrBackupResultOutput {
	return o
}

func (o GetCbrBackupResultOutput) ToGetCbrBackupResultOutputWithContext(ctx context.Context) GetCbrBackupResultOutput {
	return o
}

func (o GetCbrBackupResultOutput) CheckpointId() pulumi.StringOutput {
	return o.ApplyT(func(v GetCbrBackupResult) string { return v.CheckpointId }).(pulumi.StringOutput)
}

func (o GetCbrBackupResultOutput) Childrens() GetCbrBackupChildrenArrayOutput {
	return o.ApplyT(func(v GetCbrBackupResult) []GetCbrBackupChildren { return v.Childrens }).(GetCbrBackupChildrenArrayOutput)
}

func (o GetCbrBackupResultOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v GetCbrBackupResult) string { return v.CreatedAt }).(pulumi.StringOutput)
}

func (o GetCbrBackupResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v GetCbrBackupResult) string { return v.Description }).(pulumi.StringOutput)
}

func (o GetCbrBackupResultOutput) EnterpriseProjectId() pulumi.StringOutput {
	return o.ApplyT(func(v GetCbrBackupResult) string { return v.EnterpriseProjectId }).(pulumi.StringOutput)
}

func (o GetCbrBackupResultOutput) ExpiredAt() pulumi.StringOutput {
	return o.ApplyT(func(v GetCbrBackupResult) string { return v.ExpiredAt }).(pulumi.StringOutput)
}

func (o GetCbrBackupResultOutput) ExtendInfos() GetCbrBackupExtendInfoArrayOutput {
	return o.ApplyT(func(v GetCbrBackupResult) []GetCbrBackupExtendInfo { return v.ExtendInfos }).(GetCbrBackupExtendInfoArrayOutput)
}

func (o GetCbrBackupResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetCbrBackupResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetCbrBackupResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v GetCbrBackupResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o GetCbrBackupResultOutput) ParentId() pulumi.StringOutput {
	return o.ApplyT(func(v GetCbrBackupResult) string { return v.ParentId }).(pulumi.StringOutput)
}

func (o GetCbrBackupResultOutput) Region() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetCbrBackupResult) *string { return v.Region }).(pulumi.StringPtrOutput)
}

func (o GetCbrBackupResultOutput) ReplicationRecords() GetCbrBackupReplicationRecordArrayOutput {
	return o.ApplyT(func(v GetCbrBackupResult) []GetCbrBackupReplicationRecord { return v.ReplicationRecords }).(GetCbrBackupReplicationRecordArrayOutput)
}

func (o GetCbrBackupResultOutput) ResourceAz() pulumi.StringOutput {
	return o.ApplyT(func(v GetCbrBackupResult) string { return v.ResourceAz }).(pulumi.StringOutput)
}

func (o GetCbrBackupResultOutput) ResourceId() pulumi.StringOutput {
	return o.ApplyT(func(v GetCbrBackupResult) string { return v.ResourceId }).(pulumi.StringOutput)
}

func (o GetCbrBackupResultOutput) ResourceName() pulumi.StringOutput {
	return o.ApplyT(func(v GetCbrBackupResult) string { return v.ResourceName }).(pulumi.StringOutput)
}

func (o GetCbrBackupResultOutput) ResourceSize() pulumi.IntOutput {
	return o.ApplyT(func(v GetCbrBackupResult) int { return v.ResourceSize }).(pulumi.IntOutput)
}

func (o GetCbrBackupResultOutput) ResourceType() pulumi.StringOutput {
	return o.ApplyT(func(v GetCbrBackupResult) string { return v.ResourceType }).(pulumi.StringOutput)
}

func (o GetCbrBackupResultOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v GetCbrBackupResult) string { return v.Status }).(pulumi.StringOutput)
}

func (o GetCbrBackupResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v GetCbrBackupResult) string { return v.Type }).(pulumi.StringOutput)
}

func (o GetCbrBackupResultOutput) UpdatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v GetCbrBackupResult) string { return v.UpdatedAt }).(pulumi.StringOutput)
}

func (o GetCbrBackupResultOutput) VaultId() pulumi.StringOutput {
	return o.ApplyT(func(v GetCbrBackupResult) string { return v.VaultId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetCbrBackupResultOutput{})
}
