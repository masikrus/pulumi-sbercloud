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

type RdsInstance struct {
	pulumi.CustomResourceState

	// Deprecated: Deprecated
	AutoPay                        pulumi.StringPtrOutput          `pulumi:"autoPay"`
	AutoRenew                      pulumi.StringPtrOutput          `pulumi:"autoRenew"`
	AvailabilityZones              pulumi.StringArrayOutput        `pulumi:"availabilityZones"`
	BackupStrategy                 RdsInstanceBackupStrategyOutput `pulumi:"backupStrategy"`
	BinlogRetentionHours           pulumi.IntPtrOutput             `pulumi:"binlogRetentionHours"`
	ChargingMode                   pulumi.StringOutput             `pulumi:"chargingMode"`
	Collation                      pulumi.StringOutput             `pulumi:"collation"`
	Created                        pulumi.StringOutput             `pulumi:"created"`
	Db                             RdsInstanceDbOutput             `pulumi:"db"`
	Description                    pulumi.StringPtrOutput          `pulumi:"description"`
	DssPoolId                      pulumi.StringPtrOutput          `pulumi:"dssPoolId"`
	EnterpriseProjectId            pulumi.StringOutput             `pulumi:"enterpriseProjectId"`
	FixedIp                        pulumi.StringOutput             `pulumi:"fixedIp"`
	Flavor                         pulumi.StringOutput             `pulumi:"flavor"`
	HaReplicationMode              pulumi.StringOutput             `pulumi:"haReplicationMode"`
	LowerCaseTableNames            pulumi.StringPtrOutput          `pulumi:"lowerCaseTableNames"`
	MaintainBegin                  pulumi.StringOutput             `pulumi:"maintainBegin"`
	MaintainEnd                    pulumi.StringOutput             `pulumi:"maintainEnd"`
	MsdtcHosts                     RdsInstanceMsdtcHostArrayOutput `pulumi:"msdtcHosts"`
	Name                           pulumi.StringOutput             `pulumi:"name"`
	Nodes                          RdsInstanceNodeArrayOutput      `pulumi:"nodes"`
	ParamGroupId                   pulumi.StringPtrOutput          `pulumi:"paramGroupId"`
	Parameters                     RdsInstanceParameterArrayOutput `pulumi:"parameters"`
	Period                         pulumi.IntPtrOutput             `pulumi:"period"`
	PeriodUnit                     pulumi.StringPtrOutput          `pulumi:"periodUnit"`
	PowerAction                    pulumi.StringOutput             `pulumi:"powerAction"`
	PrivateDnsNamePrefix           pulumi.StringOutput             `pulumi:"privateDnsNamePrefix"`
	PrivateDnsNames                pulumi.StringArrayOutput        `pulumi:"privateDnsNames"`
	PrivateIps                     pulumi.StringArrayOutput        `pulumi:"privateIps"`
	PublicIps                      pulumi.StringArrayOutput        `pulumi:"publicIps"`
	ReadWritePermissions           pulumi.StringPtrOutput          `pulumi:"readWritePermissions"`
	Region                         pulumi.StringOutput             `pulumi:"region"`
	Restore                        RdsInstanceRestorePtrOutput     `pulumi:"restore"`
	RotateDay                      pulumi.IntPtrOutput             `pulumi:"rotateDay"`
	SecondsLevelMonitoringEnabled  pulumi.BoolOutput               `pulumi:"secondsLevelMonitoringEnabled"`
	SecondsLevelMonitoringInterval pulumi.IntOutput                `pulumi:"secondsLevelMonitoringInterval"`
	SecretId                       pulumi.StringPtrOutput          `pulumi:"secretId"`
	SecretName                     pulumi.StringPtrOutput          `pulumi:"secretName"`
	SecretVersion                  pulumi.StringPtrOutput          `pulumi:"secretVersion"`
	SecurityGroupId                pulumi.StringOutput             `pulumi:"securityGroupId"`
	SlowLogShowOriginalStatus      pulumi.StringPtrOutput          `pulumi:"slowLogShowOriginalStatus"`
	SslEnable                      pulumi.BoolOutput               `pulumi:"sslEnable"`
	Status                         pulumi.StringOutput             `pulumi:"status"`
	SubnetId                       pulumi.StringOutput             `pulumi:"subnetId"`
	SwitchStrategy                 pulumi.StringOutput             `pulumi:"switchStrategy"`
	Tags                           pulumi.StringMapOutput          `pulumi:"tags"`
	TdeEnabled                     pulumi.BoolOutput               `pulumi:"tdeEnabled"`
	TimeZone                       pulumi.StringOutput             `pulumi:"timeZone"`
	Volume                         RdsInstanceVolumeOutput         `pulumi:"volume"`
	VpcId                          pulumi.StringOutput             `pulumi:"vpcId"`
}

// NewRdsInstance registers a new resource with the given unique name, arguments, and options.
func NewRdsInstance(ctx *pulumi.Context,
	name string, args *RdsInstanceArgs, opts ...pulumi.ResourceOption) (*RdsInstance, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AvailabilityZones == nil {
		return nil, errors.New("invalid value for required argument 'AvailabilityZones'")
	}
	if args.Db == nil {
		return nil, errors.New("invalid value for required argument 'Db'")
	}
	if args.Flavor == nil {
		return nil, errors.New("invalid value for required argument 'Flavor'")
	}
	if args.SecurityGroupId == nil {
		return nil, errors.New("invalid value for required argument 'SecurityGroupId'")
	}
	if args.SubnetId == nil {
		return nil, errors.New("invalid value for required argument 'SubnetId'")
	}
	if args.Volume == nil {
		return nil, errors.New("invalid value for required argument 'Volume'")
	}
	if args.VpcId == nil {
		return nil, errors.New("invalid value for required argument 'VpcId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource RdsInstance
	err := ctx.RegisterResource("sbercloud:index/rdsInstance:RdsInstance", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRdsInstance gets an existing RdsInstance resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRdsInstance(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RdsInstanceState, opts ...pulumi.ResourceOption) (*RdsInstance, error) {
	var resource RdsInstance
	err := ctx.ReadResource("sbercloud:index/rdsInstance:RdsInstance", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering RdsInstance resources.
type rdsInstanceState struct {
	// Deprecated: Deprecated
	AutoPay                        *string                    `pulumi:"autoPay"`
	AutoRenew                      *string                    `pulumi:"autoRenew"`
	AvailabilityZones              []string                   `pulumi:"availabilityZones"`
	BackupStrategy                 *RdsInstanceBackupStrategy `pulumi:"backupStrategy"`
	BinlogRetentionHours           *int                       `pulumi:"binlogRetentionHours"`
	ChargingMode                   *string                    `pulumi:"chargingMode"`
	Collation                      *string                    `pulumi:"collation"`
	Created                        *string                    `pulumi:"created"`
	Db                             *RdsInstanceDb             `pulumi:"db"`
	Description                    *string                    `pulumi:"description"`
	DssPoolId                      *string                    `pulumi:"dssPoolId"`
	EnterpriseProjectId            *string                    `pulumi:"enterpriseProjectId"`
	FixedIp                        *string                    `pulumi:"fixedIp"`
	Flavor                         *string                    `pulumi:"flavor"`
	HaReplicationMode              *string                    `pulumi:"haReplicationMode"`
	LowerCaseTableNames            *string                    `pulumi:"lowerCaseTableNames"`
	MaintainBegin                  *string                    `pulumi:"maintainBegin"`
	MaintainEnd                    *string                    `pulumi:"maintainEnd"`
	MsdtcHosts                     []RdsInstanceMsdtcHost     `pulumi:"msdtcHosts"`
	Name                           *string                    `pulumi:"name"`
	Nodes                          []RdsInstanceNode          `pulumi:"nodes"`
	ParamGroupId                   *string                    `pulumi:"paramGroupId"`
	Parameters                     []RdsInstanceParameter     `pulumi:"parameters"`
	Period                         *int                       `pulumi:"period"`
	PeriodUnit                     *string                    `pulumi:"periodUnit"`
	PowerAction                    *string                    `pulumi:"powerAction"`
	PrivateDnsNamePrefix           *string                    `pulumi:"privateDnsNamePrefix"`
	PrivateDnsNames                []string                   `pulumi:"privateDnsNames"`
	PrivateIps                     []string                   `pulumi:"privateIps"`
	PublicIps                      []string                   `pulumi:"publicIps"`
	ReadWritePermissions           *string                    `pulumi:"readWritePermissions"`
	Region                         *string                    `pulumi:"region"`
	Restore                        *RdsInstanceRestore        `pulumi:"restore"`
	RotateDay                      *int                       `pulumi:"rotateDay"`
	SecondsLevelMonitoringEnabled  *bool                      `pulumi:"secondsLevelMonitoringEnabled"`
	SecondsLevelMonitoringInterval *int                       `pulumi:"secondsLevelMonitoringInterval"`
	SecretId                       *string                    `pulumi:"secretId"`
	SecretName                     *string                    `pulumi:"secretName"`
	SecretVersion                  *string                    `pulumi:"secretVersion"`
	SecurityGroupId                *string                    `pulumi:"securityGroupId"`
	SlowLogShowOriginalStatus      *string                    `pulumi:"slowLogShowOriginalStatus"`
	SslEnable                      *bool                      `pulumi:"sslEnable"`
	Status                         *string                    `pulumi:"status"`
	SubnetId                       *string                    `pulumi:"subnetId"`
	SwitchStrategy                 *string                    `pulumi:"switchStrategy"`
	Tags                           map[string]string          `pulumi:"tags"`
	TdeEnabled                     *bool                      `pulumi:"tdeEnabled"`
	TimeZone                       *string                    `pulumi:"timeZone"`
	Volume                         *RdsInstanceVolume         `pulumi:"volume"`
	VpcId                          *string                    `pulumi:"vpcId"`
}

type RdsInstanceState struct {
	// Deprecated: Deprecated
	AutoPay                        pulumi.StringPtrInput
	AutoRenew                      pulumi.StringPtrInput
	AvailabilityZones              pulumi.StringArrayInput
	BackupStrategy                 RdsInstanceBackupStrategyPtrInput
	BinlogRetentionHours           pulumi.IntPtrInput
	ChargingMode                   pulumi.StringPtrInput
	Collation                      pulumi.StringPtrInput
	Created                        pulumi.StringPtrInput
	Db                             RdsInstanceDbPtrInput
	Description                    pulumi.StringPtrInput
	DssPoolId                      pulumi.StringPtrInput
	EnterpriseProjectId            pulumi.StringPtrInput
	FixedIp                        pulumi.StringPtrInput
	Flavor                         pulumi.StringPtrInput
	HaReplicationMode              pulumi.StringPtrInput
	LowerCaseTableNames            pulumi.StringPtrInput
	MaintainBegin                  pulumi.StringPtrInput
	MaintainEnd                    pulumi.StringPtrInput
	MsdtcHosts                     RdsInstanceMsdtcHostArrayInput
	Name                           pulumi.StringPtrInput
	Nodes                          RdsInstanceNodeArrayInput
	ParamGroupId                   pulumi.StringPtrInput
	Parameters                     RdsInstanceParameterArrayInput
	Period                         pulumi.IntPtrInput
	PeriodUnit                     pulumi.StringPtrInput
	PowerAction                    pulumi.StringPtrInput
	PrivateDnsNamePrefix           pulumi.StringPtrInput
	PrivateDnsNames                pulumi.StringArrayInput
	PrivateIps                     pulumi.StringArrayInput
	PublicIps                      pulumi.StringArrayInput
	ReadWritePermissions           pulumi.StringPtrInput
	Region                         pulumi.StringPtrInput
	Restore                        RdsInstanceRestorePtrInput
	RotateDay                      pulumi.IntPtrInput
	SecondsLevelMonitoringEnabled  pulumi.BoolPtrInput
	SecondsLevelMonitoringInterval pulumi.IntPtrInput
	SecretId                       pulumi.StringPtrInput
	SecretName                     pulumi.StringPtrInput
	SecretVersion                  pulumi.StringPtrInput
	SecurityGroupId                pulumi.StringPtrInput
	SlowLogShowOriginalStatus      pulumi.StringPtrInput
	SslEnable                      pulumi.BoolPtrInput
	Status                         pulumi.StringPtrInput
	SubnetId                       pulumi.StringPtrInput
	SwitchStrategy                 pulumi.StringPtrInput
	Tags                           pulumi.StringMapInput
	TdeEnabled                     pulumi.BoolPtrInput
	TimeZone                       pulumi.StringPtrInput
	Volume                         RdsInstanceVolumePtrInput
	VpcId                          pulumi.StringPtrInput
}

func (RdsInstanceState) ElementType() reflect.Type {
	return reflect.TypeOf((*rdsInstanceState)(nil)).Elem()
}

type rdsInstanceArgs struct {
	// Deprecated: Deprecated
	AutoPay                        *string                    `pulumi:"autoPay"`
	AutoRenew                      *string                    `pulumi:"autoRenew"`
	AvailabilityZones              []string                   `pulumi:"availabilityZones"`
	BackupStrategy                 *RdsInstanceBackupStrategy `pulumi:"backupStrategy"`
	BinlogRetentionHours           *int                       `pulumi:"binlogRetentionHours"`
	ChargingMode                   *string                    `pulumi:"chargingMode"`
	Collation                      *string                    `pulumi:"collation"`
	Db                             RdsInstanceDb              `pulumi:"db"`
	Description                    *string                    `pulumi:"description"`
	DssPoolId                      *string                    `pulumi:"dssPoolId"`
	EnterpriseProjectId            *string                    `pulumi:"enterpriseProjectId"`
	FixedIp                        *string                    `pulumi:"fixedIp"`
	Flavor                         string                     `pulumi:"flavor"`
	HaReplicationMode              *string                    `pulumi:"haReplicationMode"`
	LowerCaseTableNames            *string                    `pulumi:"lowerCaseTableNames"`
	MaintainBegin                  *string                    `pulumi:"maintainBegin"`
	MaintainEnd                    *string                    `pulumi:"maintainEnd"`
	MsdtcHosts                     []RdsInstanceMsdtcHost     `pulumi:"msdtcHosts"`
	Name                           *string                    `pulumi:"name"`
	ParamGroupId                   *string                    `pulumi:"paramGroupId"`
	Parameters                     []RdsInstanceParameter     `pulumi:"parameters"`
	Period                         *int                       `pulumi:"period"`
	PeriodUnit                     *string                    `pulumi:"periodUnit"`
	PowerAction                    *string                    `pulumi:"powerAction"`
	PrivateDnsNamePrefix           *string                    `pulumi:"privateDnsNamePrefix"`
	ReadWritePermissions           *string                    `pulumi:"readWritePermissions"`
	Region                         *string                    `pulumi:"region"`
	Restore                        *RdsInstanceRestore        `pulumi:"restore"`
	RotateDay                      *int                       `pulumi:"rotateDay"`
	SecondsLevelMonitoringEnabled  *bool                      `pulumi:"secondsLevelMonitoringEnabled"`
	SecondsLevelMonitoringInterval *int                       `pulumi:"secondsLevelMonitoringInterval"`
	SecretId                       *string                    `pulumi:"secretId"`
	SecretName                     *string                    `pulumi:"secretName"`
	SecretVersion                  *string                    `pulumi:"secretVersion"`
	SecurityGroupId                string                     `pulumi:"securityGroupId"`
	SlowLogShowOriginalStatus      *string                    `pulumi:"slowLogShowOriginalStatus"`
	SslEnable                      *bool                      `pulumi:"sslEnable"`
	SubnetId                       string                     `pulumi:"subnetId"`
	SwitchStrategy                 *string                    `pulumi:"switchStrategy"`
	Tags                           map[string]string          `pulumi:"tags"`
	TdeEnabled                     *bool                      `pulumi:"tdeEnabled"`
	TimeZone                       *string                    `pulumi:"timeZone"`
	Volume                         RdsInstanceVolume          `pulumi:"volume"`
	VpcId                          string                     `pulumi:"vpcId"`
}

// The set of arguments for constructing a RdsInstance resource.
type RdsInstanceArgs struct {
	// Deprecated: Deprecated
	AutoPay                        pulumi.StringPtrInput
	AutoRenew                      pulumi.StringPtrInput
	AvailabilityZones              pulumi.StringArrayInput
	BackupStrategy                 RdsInstanceBackupStrategyPtrInput
	BinlogRetentionHours           pulumi.IntPtrInput
	ChargingMode                   pulumi.StringPtrInput
	Collation                      pulumi.StringPtrInput
	Db                             RdsInstanceDbInput
	Description                    pulumi.StringPtrInput
	DssPoolId                      pulumi.StringPtrInput
	EnterpriseProjectId            pulumi.StringPtrInput
	FixedIp                        pulumi.StringPtrInput
	Flavor                         pulumi.StringInput
	HaReplicationMode              pulumi.StringPtrInput
	LowerCaseTableNames            pulumi.StringPtrInput
	MaintainBegin                  pulumi.StringPtrInput
	MaintainEnd                    pulumi.StringPtrInput
	MsdtcHosts                     RdsInstanceMsdtcHostArrayInput
	Name                           pulumi.StringPtrInput
	ParamGroupId                   pulumi.StringPtrInput
	Parameters                     RdsInstanceParameterArrayInput
	Period                         pulumi.IntPtrInput
	PeriodUnit                     pulumi.StringPtrInput
	PowerAction                    pulumi.StringPtrInput
	PrivateDnsNamePrefix           pulumi.StringPtrInput
	ReadWritePermissions           pulumi.StringPtrInput
	Region                         pulumi.StringPtrInput
	Restore                        RdsInstanceRestorePtrInput
	RotateDay                      pulumi.IntPtrInput
	SecondsLevelMonitoringEnabled  pulumi.BoolPtrInput
	SecondsLevelMonitoringInterval pulumi.IntPtrInput
	SecretId                       pulumi.StringPtrInput
	SecretName                     pulumi.StringPtrInput
	SecretVersion                  pulumi.StringPtrInput
	SecurityGroupId                pulumi.StringInput
	SlowLogShowOriginalStatus      pulumi.StringPtrInput
	SslEnable                      pulumi.BoolPtrInput
	SubnetId                       pulumi.StringInput
	SwitchStrategy                 pulumi.StringPtrInput
	Tags                           pulumi.StringMapInput
	TdeEnabled                     pulumi.BoolPtrInput
	TimeZone                       pulumi.StringPtrInput
	Volume                         RdsInstanceVolumeInput
	VpcId                          pulumi.StringInput
}

func (RdsInstanceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*rdsInstanceArgs)(nil)).Elem()
}

type RdsInstanceInput interface {
	pulumi.Input

	ToRdsInstanceOutput() RdsInstanceOutput
	ToRdsInstanceOutputWithContext(ctx context.Context) RdsInstanceOutput
}

func (*RdsInstance) ElementType() reflect.Type {
	return reflect.TypeOf((**RdsInstance)(nil)).Elem()
}

func (i *RdsInstance) ToRdsInstanceOutput() RdsInstanceOutput {
	return i.ToRdsInstanceOutputWithContext(context.Background())
}

func (i *RdsInstance) ToRdsInstanceOutputWithContext(ctx context.Context) RdsInstanceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RdsInstanceOutput)
}

// RdsInstanceArrayInput is an input type that accepts RdsInstanceArray and RdsInstanceArrayOutput values.
// You can construct a concrete instance of `RdsInstanceArrayInput` via:
//
//	RdsInstanceArray{ RdsInstanceArgs{...} }
type RdsInstanceArrayInput interface {
	pulumi.Input

	ToRdsInstanceArrayOutput() RdsInstanceArrayOutput
	ToRdsInstanceArrayOutputWithContext(context.Context) RdsInstanceArrayOutput
}

type RdsInstanceArray []RdsInstanceInput

func (RdsInstanceArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*RdsInstance)(nil)).Elem()
}

func (i RdsInstanceArray) ToRdsInstanceArrayOutput() RdsInstanceArrayOutput {
	return i.ToRdsInstanceArrayOutputWithContext(context.Background())
}

func (i RdsInstanceArray) ToRdsInstanceArrayOutputWithContext(ctx context.Context) RdsInstanceArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RdsInstanceArrayOutput)
}

// RdsInstanceMapInput is an input type that accepts RdsInstanceMap and RdsInstanceMapOutput values.
// You can construct a concrete instance of `RdsInstanceMapInput` via:
//
//	RdsInstanceMap{ "key": RdsInstanceArgs{...} }
type RdsInstanceMapInput interface {
	pulumi.Input

	ToRdsInstanceMapOutput() RdsInstanceMapOutput
	ToRdsInstanceMapOutputWithContext(context.Context) RdsInstanceMapOutput
}

type RdsInstanceMap map[string]RdsInstanceInput

func (RdsInstanceMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*RdsInstance)(nil)).Elem()
}

func (i RdsInstanceMap) ToRdsInstanceMapOutput() RdsInstanceMapOutput {
	return i.ToRdsInstanceMapOutputWithContext(context.Background())
}

func (i RdsInstanceMap) ToRdsInstanceMapOutputWithContext(ctx context.Context) RdsInstanceMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RdsInstanceMapOutput)
}

type RdsInstanceOutput struct{ *pulumi.OutputState }

func (RdsInstanceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RdsInstance)(nil)).Elem()
}

func (o RdsInstanceOutput) ToRdsInstanceOutput() RdsInstanceOutput {
	return o
}

func (o RdsInstanceOutput) ToRdsInstanceOutputWithContext(ctx context.Context) RdsInstanceOutput {
	return o
}

// Deprecated: Deprecated
func (o RdsInstanceOutput) AutoPay() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RdsInstance) pulumi.StringPtrOutput { return v.AutoPay }).(pulumi.StringPtrOutput)
}

func (o RdsInstanceOutput) AutoRenew() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RdsInstance) pulumi.StringPtrOutput { return v.AutoRenew }).(pulumi.StringPtrOutput)
}

func (o RdsInstanceOutput) AvailabilityZones() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *RdsInstance) pulumi.StringArrayOutput { return v.AvailabilityZones }).(pulumi.StringArrayOutput)
}

func (o RdsInstanceOutput) BackupStrategy() RdsInstanceBackupStrategyOutput {
	return o.ApplyT(func(v *RdsInstance) RdsInstanceBackupStrategyOutput { return v.BackupStrategy }).(RdsInstanceBackupStrategyOutput)
}

func (o RdsInstanceOutput) BinlogRetentionHours() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *RdsInstance) pulumi.IntPtrOutput { return v.BinlogRetentionHours }).(pulumi.IntPtrOutput)
}

func (o RdsInstanceOutput) ChargingMode() pulumi.StringOutput {
	return o.ApplyT(func(v *RdsInstance) pulumi.StringOutput { return v.ChargingMode }).(pulumi.StringOutput)
}

func (o RdsInstanceOutput) Collation() pulumi.StringOutput {
	return o.ApplyT(func(v *RdsInstance) pulumi.StringOutput { return v.Collation }).(pulumi.StringOutput)
}

func (o RdsInstanceOutput) Created() pulumi.StringOutput {
	return o.ApplyT(func(v *RdsInstance) pulumi.StringOutput { return v.Created }).(pulumi.StringOutput)
}

func (o RdsInstanceOutput) Db() RdsInstanceDbOutput {
	return o.ApplyT(func(v *RdsInstance) RdsInstanceDbOutput { return v.Db }).(RdsInstanceDbOutput)
}

func (o RdsInstanceOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RdsInstance) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

func (o RdsInstanceOutput) DssPoolId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RdsInstance) pulumi.StringPtrOutput { return v.DssPoolId }).(pulumi.StringPtrOutput)
}

func (o RdsInstanceOutput) EnterpriseProjectId() pulumi.StringOutput {
	return o.ApplyT(func(v *RdsInstance) pulumi.StringOutput { return v.EnterpriseProjectId }).(pulumi.StringOutput)
}

func (o RdsInstanceOutput) FixedIp() pulumi.StringOutput {
	return o.ApplyT(func(v *RdsInstance) pulumi.StringOutput { return v.FixedIp }).(pulumi.StringOutput)
}

func (o RdsInstanceOutput) Flavor() pulumi.StringOutput {
	return o.ApplyT(func(v *RdsInstance) pulumi.StringOutput { return v.Flavor }).(pulumi.StringOutput)
}

func (o RdsInstanceOutput) HaReplicationMode() pulumi.StringOutput {
	return o.ApplyT(func(v *RdsInstance) pulumi.StringOutput { return v.HaReplicationMode }).(pulumi.StringOutput)
}

func (o RdsInstanceOutput) LowerCaseTableNames() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RdsInstance) pulumi.StringPtrOutput { return v.LowerCaseTableNames }).(pulumi.StringPtrOutput)
}

func (o RdsInstanceOutput) MaintainBegin() pulumi.StringOutput {
	return o.ApplyT(func(v *RdsInstance) pulumi.StringOutput { return v.MaintainBegin }).(pulumi.StringOutput)
}

func (o RdsInstanceOutput) MaintainEnd() pulumi.StringOutput {
	return o.ApplyT(func(v *RdsInstance) pulumi.StringOutput { return v.MaintainEnd }).(pulumi.StringOutput)
}

func (o RdsInstanceOutput) MsdtcHosts() RdsInstanceMsdtcHostArrayOutput {
	return o.ApplyT(func(v *RdsInstance) RdsInstanceMsdtcHostArrayOutput { return v.MsdtcHosts }).(RdsInstanceMsdtcHostArrayOutput)
}

func (o RdsInstanceOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *RdsInstance) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o RdsInstanceOutput) Nodes() RdsInstanceNodeArrayOutput {
	return o.ApplyT(func(v *RdsInstance) RdsInstanceNodeArrayOutput { return v.Nodes }).(RdsInstanceNodeArrayOutput)
}

func (o RdsInstanceOutput) ParamGroupId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RdsInstance) pulumi.StringPtrOutput { return v.ParamGroupId }).(pulumi.StringPtrOutput)
}

func (o RdsInstanceOutput) Parameters() RdsInstanceParameterArrayOutput {
	return o.ApplyT(func(v *RdsInstance) RdsInstanceParameterArrayOutput { return v.Parameters }).(RdsInstanceParameterArrayOutput)
}

func (o RdsInstanceOutput) Period() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *RdsInstance) pulumi.IntPtrOutput { return v.Period }).(pulumi.IntPtrOutput)
}

func (o RdsInstanceOutput) PeriodUnit() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RdsInstance) pulumi.StringPtrOutput { return v.PeriodUnit }).(pulumi.StringPtrOutput)
}

func (o RdsInstanceOutput) PowerAction() pulumi.StringOutput {
	return o.ApplyT(func(v *RdsInstance) pulumi.StringOutput { return v.PowerAction }).(pulumi.StringOutput)
}

func (o RdsInstanceOutput) PrivateDnsNamePrefix() pulumi.StringOutput {
	return o.ApplyT(func(v *RdsInstance) pulumi.StringOutput { return v.PrivateDnsNamePrefix }).(pulumi.StringOutput)
}

func (o RdsInstanceOutput) PrivateDnsNames() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *RdsInstance) pulumi.StringArrayOutput { return v.PrivateDnsNames }).(pulumi.StringArrayOutput)
}

func (o RdsInstanceOutput) PrivateIps() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *RdsInstance) pulumi.StringArrayOutput { return v.PrivateIps }).(pulumi.StringArrayOutput)
}

func (o RdsInstanceOutput) PublicIps() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *RdsInstance) pulumi.StringArrayOutput { return v.PublicIps }).(pulumi.StringArrayOutput)
}

func (o RdsInstanceOutput) ReadWritePermissions() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RdsInstance) pulumi.StringPtrOutput { return v.ReadWritePermissions }).(pulumi.StringPtrOutput)
}

func (o RdsInstanceOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v *RdsInstance) pulumi.StringOutput { return v.Region }).(pulumi.StringOutput)
}

func (o RdsInstanceOutput) Restore() RdsInstanceRestorePtrOutput {
	return o.ApplyT(func(v *RdsInstance) RdsInstanceRestorePtrOutput { return v.Restore }).(RdsInstanceRestorePtrOutput)
}

func (o RdsInstanceOutput) RotateDay() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *RdsInstance) pulumi.IntPtrOutput { return v.RotateDay }).(pulumi.IntPtrOutput)
}

func (o RdsInstanceOutput) SecondsLevelMonitoringEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v *RdsInstance) pulumi.BoolOutput { return v.SecondsLevelMonitoringEnabled }).(pulumi.BoolOutput)
}

func (o RdsInstanceOutput) SecondsLevelMonitoringInterval() pulumi.IntOutput {
	return o.ApplyT(func(v *RdsInstance) pulumi.IntOutput { return v.SecondsLevelMonitoringInterval }).(pulumi.IntOutput)
}

func (o RdsInstanceOutput) SecretId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RdsInstance) pulumi.StringPtrOutput { return v.SecretId }).(pulumi.StringPtrOutput)
}

func (o RdsInstanceOutput) SecretName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RdsInstance) pulumi.StringPtrOutput { return v.SecretName }).(pulumi.StringPtrOutput)
}

func (o RdsInstanceOutput) SecretVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RdsInstance) pulumi.StringPtrOutput { return v.SecretVersion }).(pulumi.StringPtrOutput)
}

func (o RdsInstanceOutput) SecurityGroupId() pulumi.StringOutput {
	return o.ApplyT(func(v *RdsInstance) pulumi.StringOutput { return v.SecurityGroupId }).(pulumi.StringOutput)
}

func (o RdsInstanceOutput) SlowLogShowOriginalStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RdsInstance) pulumi.StringPtrOutput { return v.SlowLogShowOriginalStatus }).(pulumi.StringPtrOutput)
}

func (o RdsInstanceOutput) SslEnable() pulumi.BoolOutput {
	return o.ApplyT(func(v *RdsInstance) pulumi.BoolOutput { return v.SslEnable }).(pulumi.BoolOutput)
}

func (o RdsInstanceOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *RdsInstance) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

func (o RdsInstanceOutput) SubnetId() pulumi.StringOutput {
	return o.ApplyT(func(v *RdsInstance) pulumi.StringOutput { return v.SubnetId }).(pulumi.StringOutput)
}

func (o RdsInstanceOutput) SwitchStrategy() pulumi.StringOutput {
	return o.ApplyT(func(v *RdsInstance) pulumi.StringOutput { return v.SwitchStrategy }).(pulumi.StringOutput)
}

func (o RdsInstanceOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *RdsInstance) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o RdsInstanceOutput) TdeEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v *RdsInstance) pulumi.BoolOutput { return v.TdeEnabled }).(pulumi.BoolOutput)
}

func (o RdsInstanceOutput) TimeZone() pulumi.StringOutput {
	return o.ApplyT(func(v *RdsInstance) pulumi.StringOutput { return v.TimeZone }).(pulumi.StringOutput)
}

func (o RdsInstanceOutput) Volume() RdsInstanceVolumeOutput {
	return o.ApplyT(func(v *RdsInstance) RdsInstanceVolumeOutput { return v.Volume }).(RdsInstanceVolumeOutput)
}

func (o RdsInstanceOutput) VpcId() pulumi.StringOutput {
	return o.ApplyT(func(v *RdsInstance) pulumi.StringOutput { return v.VpcId }).(pulumi.StringOutput)
}

type RdsInstanceArrayOutput struct{ *pulumi.OutputState }

func (RdsInstanceArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*RdsInstance)(nil)).Elem()
}

func (o RdsInstanceArrayOutput) ToRdsInstanceArrayOutput() RdsInstanceArrayOutput {
	return o
}

func (o RdsInstanceArrayOutput) ToRdsInstanceArrayOutputWithContext(ctx context.Context) RdsInstanceArrayOutput {
	return o
}

func (o RdsInstanceArrayOutput) Index(i pulumi.IntInput) RdsInstanceOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *RdsInstance {
		return vs[0].([]*RdsInstance)[vs[1].(int)]
	}).(RdsInstanceOutput)
}

type RdsInstanceMapOutput struct{ *pulumi.OutputState }

func (RdsInstanceMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*RdsInstance)(nil)).Elem()
}

func (o RdsInstanceMapOutput) ToRdsInstanceMapOutput() RdsInstanceMapOutput {
	return o
}

func (o RdsInstanceMapOutput) ToRdsInstanceMapOutputWithContext(ctx context.Context) RdsInstanceMapOutput {
	return o
}

func (o RdsInstanceMapOutput) MapIndex(k pulumi.StringInput) RdsInstanceOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *RdsInstance {
		return vs[0].(map[string]*RdsInstance)[vs[1].(string)]
	}).(RdsInstanceOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*RdsInstanceInput)(nil)).Elem(), &RdsInstance{})
	pulumi.RegisterInputType(reflect.TypeOf((*RdsInstanceArrayInput)(nil)).Elem(), RdsInstanceArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*RdsInstanceMapInput)(nil)).Elem(), RdsInstanceMap{})
	pulumi.RegisterOutputType(RdsInstanceOutput{})
	pulumi.RegisterOutputType(RdsInstanceArrayOutput{})
	pulumi.RegisterOutputType(RdsInstanceMapOutput{})
}
