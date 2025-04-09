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

type DdsInstance struct {
	pulumi.CustomResourceState

	// Deprecated: Deprecated
	AutoPay             pulumi.StringPtrOutput              `pulumi:"autoPay"`
	AutoRenew           pulumi.StringPtrOutput              `pulumi:"autoRenew"`
	AvailabilityZone    pulumi.StringOutput                 `pulumi:"availabilityZone"`
	BackupStrategy      DdsInstanceBackupStrategyOutput     `pulumi:"backupStrategy"`
	BalancerActiveBegin pulumi.StringPtrOutput              `pulumi:"balancerActiveBegin"`
	BalancerActiveEnd   pulumi.StringPtrOutput              `pulumi:"balancerActiveEnd"`
	BalancerStatus      pulumi.StringOutput                 `pulumi:"balancerStatus"`
	ChargingMode        pulumi.StringOutput                 `pulumi:"chargingMode"`
	Configurations      DdsInstanceConfigurationArrayOutput `pulumi:"configurations"`
	CreatedAt           pulumi.StringOutput                 `pulumi:"createdAt"`
	Datastore           DdsInstanceDatastoreOutput          `pulumi:"datastore"`
	DbUsername          pulumi.StringOutput                 `pulumi:"dbUsername"`
	Description         pulumi.StringPtrOutput              `pulumi:"description"`
	DiskEncryptionId    pulumi.StringPtrOutput              `pulumi:"diskEncryptionId"`
	EnterpriseProjectId pulumi.StringOutput                 `pulumi:"enterpriseProjectId"`
	Flavors             DdsInstanceFlavorArrayOutput        `pulumi:"flavors"`
	Groups              DdsInstanceGroupArrayOutput         `pulumi:"groups"`
	Mode                pulumi.StringOutput                 `pulumi:"mode"`
	Name                pulumi.StringOutput                 `pulumi:"name"`
	// This field is deprecated.
	Nodes           DdsInstanceNodeArrayOutput `pulumi:"nodes"`
	Password        pulumi.StringPtrOutput     `pulumi:"password"`
	Period          pulumi.IntPtrOutput        `pulumi:"period"`
	PeriodUnit      pulumi.StringPtrOutput     `pulumi:"periodUnit"`
	Port            pulumi.IntOutput           `pulumi:"port"`
	Region          pulumi.StringOutput        `pulumi:"region"`
	SecurityGroupId pulumi.StringOutput        `pulumi:"securityGroupId"`
	Ssl             pulumi.BoolPtrOutput       `pulumi:"ssl"`
	Status          pulumi.StringOutput        `pulumi:"status"`
	SubnetId        pulumi.StringOutput        `pulumi:"subnetId"`
	Tags            pulumi.StringMapOutput     `pulumi:"tags"`
	TimeZone        pulumi.StringOutput        `pulumi:"timeZone"`
	UpdatedAt       pulumi.StringOutput        `pulumi:"updatedAt"`
	VpcId           pulumi.StringOutput        `pulumi:"vpcId"`
}

// NewDdsInstance registers a new resource with the given unique name, arguments, and options.
func NewDdsInstance(ctx *pulumi.Context,
	name string, args *DdsInstanceArgs, opts ...pulumi.ResourceOption) (*DdsInstance, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AvailabilityZone == nil {
		return nil, errors.New("invalid value for required argument 'AvailabilityZone'")
	}
	if args.Datastore == nil {
		return nil, errors.New("invalid value for required argument 'Datastore'")
	}
	if args.Flavors == nil {
		return nil, errors.New("invalid value for required argument 'Flavors'")
	}
	if args.Mode == nil {
		return nil, errors.New("invalid value for required argument 'Mode'")
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
	if args.Password != nil {
		args.Password = pulumi.ToSecret(args.Password).(pulumi.StringPtrInput)
	}
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"password",
	})
	opts = append(opts, secrets)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource DdsInstance
	err := ctx.RegisterResource("sbercloud:index/ddsInstance:DdsInstance", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDdsInstance gets an existing DdsInstance resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDdsInstance(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DdsInstanceState, opts ...pulumi.ResourceOption) (*DdsInstance, error) {
	var resource DdsInstance
	err := ctx.ReadResource("sbercloud:index/ddsInstance:DdsInstance", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DdsInstance resources.
type ddsInstanceState struct {
	// Deprecated: Deprecated
	AutoPay             *string                    `pulumi:"autoPay"`
	AutoRenew           *string                    `pulumi:"autoRenew"`
	AvailabilityZone    *string                    `pulumi:"availabilityZone"`
	BackupStrategy      *DdsInstanceBackupStrategy `pulumi:"backupStrategy"`
	BalancerActiveBegin *string                    `pulumi:"balancerActiveBegin"`
	BalancerActiveEnd   *string                    `pulumi:"balancerActiveEnd"`
	BalancerStatus      *string                    `pulumi:"balancerStatus"`
	ChargingMode        *string                    `pulumi:"chargingMode"`
	Configurations      []DdsInstanceConfiguration `pulumi:"configurations"`
	CreatedAt           *string                    `pulumi:"createdAt"`
	Datastore           *DdsInstanceDatastore      `pulumi:"datastore"`
	DbUsername          *string                    `pulumi:"dbUsername"`
	Description         *string                    `pulumi:"description"`
	DiskEncryptionId    *string                    `pulumi:"diskEncryptionId"`
	EnterpriseProjectId *string                    `pulumi:"enterpriseProjectId"`
	Flavors             []DdsInstanceFlavor        `pulumi:"flavors"`
	Groups              []DdsInstanceGroup         `pulumi:"groups"`
	Mode                *string                    `pulumi:"mode"`
	Name                *string                    `pulumi:"name"`
	// This field is deprecated.
	Nodes           []DdsInstanceNode `pulumi:"nodes"`
	Password        *string           `pulumi:"password"`
	Period          *int              `pulumi:"period"`
	PeriodUnit      *string           `pulumi:"periodUnit"`
	Port            *int              `pulumi:"port"`
	Region          *string           `pulumi:"region"`
	SecurityGroupId *string           `pulumi:"securityGroupId"`
	Ssl             *bool             `pulumi:"ssl"`
	Status          *string           `pulumi:"status"`
	SubnetId        *string           `pulumi:"subnetId"`
	Tags            map[string]string `pulumi:"tags"`
	TimeZone        *string           `pulumi:"timeZone"`
	UpdatedAt       *string           `pulumi:"updatedAt"`
	VpcId           *string           `pulumi:"vpcId"`
}

type DdsInstanceState struct {
	// Deprecated: Deprecated
	AutoPay             pulumi.StringPtrInput
	AutoRenew           pulumi.StringPtrInput
	AvailabilityZone    pulumi.StringPtrInput
	BackupStrategy      DdsInstanceBackupStrategyPtrInput
	BalancerActiveBegin pulumi.StringPtrInput
	BalancerActiveEnd   pulumi.StringPtrInput
	BalancerStatus      pulumi.StringPtrInput
	ChargingMode        pulumi.StringPtrInput
	Configurations      DdsInstanceConfigurationArrayInput
	CreatedAt           pulumi.StringPtrInput
	Datastore           DdsInstanceDatastorePtrInput
	DbUsername          pulumi.StringPtrInput
	Description         pulumi.StringPtrInput
	DiskEncryptionId    pulumi.StringPtrInput
	EnterpriseProjectId pulumi.StringPtrInput
	Flavors             DdsInstanceFlavorArrayInput
	Groups              DdsInstanceGroupArrayInput
	Mode                pulumi.StringPtrInput
	Name                pulumi.StringPtrInput
	// This field is deprecated.
	Nodes           DdsInstanceNodeArrayInput
	Password        pulumi.StringPtrInput
	Period          pulumi.IntPtrInput
	PeriodUnit      pulumi.StringPtrInput
	Port            pulumi.IntPtrInput
	Region          pulumi.StringPtrInput
	SecurityGroupId pulumi.StringPtrInput
	Ssl             pulumi.BoolPtrInput
	Status          pulumi.StringPtrInput
	SubnetId        pulumi.StringPtrInput
	Tags            pulumi.StringMapInput
	TimeZone        pulumi.StringPtrInput
	UpdatedAt       pulumi.StringPtrInput
	VpcId           pulumi.StringPtrInput
}

func (DdsInstanceState) ElementType() reflect.Type {
	return reflect.TypeOf((*ddsInstanceState)(nil)).Elem()
}

type ddsInstanceArgs struct {
	// Deprecated: Deprecated
	AutoPay             *string                    `pulumi:"autoPay"`
	AutoRenew           *string                    `pulumi:"autoRenew"`
	AvailabilityZone    string                     `pulumi:"availabilityZone"`
	BackupStrategy      *DdsInstanceBackupStrategy `pulumi:"backupStrategy"`
	BalancerActiveBegin *string                    `pulumi:"balancerActiveBegin"`
	BalancerActiveEnd   *string                    `pulumi:"balancerActiveEnd"`
	BalancerStatus      *string                    `pulumi:"balancerStatus"`
	ChargingMode        *string                    `pulumi:"chargingMode"`
	Configurations      []DdsInstanceConfiguration `pulumi:"configurations"`
	Datastore           DdsInstanceDatastore       `pulumi:"datastore"`
	Description         *string                    `pulumi:"description"`
	DiskEncryptionId    *string                    `pulumi:"diskEncryptionId"`
	EnterpriseProjectId *string                    `pulumi:"enterpriseProjectId"`
	Flavors             []DdsInstanceFlavor        `pulumi:"flavors"`
	Mode                string                     `pulumi:"mode"`
	Name                *string                    `pulumi:"name"`
	Password            *string                    `pulumi:"password"`
	Period              *int                       `pulumi:"period"`
	PeriodUnit          *string                    `pulumi:"periodUnit"`
	Port                *int                       `pulumi:"port"`
	Region              *string                    `pulumi:"region"`
	SecurityGroupId     string                     `pulumi:"securityGroupId"`
	Ssl                 *bool                      `pulumi:"ssl"`
	SubnetId            string                     `pulumi:"subnetId"`
	Tags                map[string]string          `pulumi:"tags"`
	VpcId               string                     `pulumi:"vpcId"`
}

// The set of arguments for constructing a DdsInstance resource.
type DdsInstanceArgs struct {
	// Deprecated: Deprecated
	AutoPay             pulumi.StringPtrInput
	AutoRenew           pulumi.StringPtrInput
	AvailabilityZone    pulumi.StringInput
	BackupStrategy      DdsInstanceBackupStrategyPtrInput
	BalancerActiveBegin pulumi.StringPtrInput
	BalancerActiveEnd   pulumi.StringPtrInput
	BalancerStatus      pulumi.StringPtrInput
	ChargingMode        pulumi.StringPtrInput
	Configurations      DdsInstanceConfigurationArrayInput
	Datastore           DdsInstanceDatastoreInput
	Description         pulumi.StringPtrInput
	DiskEncryptionId    pulumi.StringPtrInput
	EnterpriseProjectId pulumi.StringPtrInput
	Flavors             DdsInstanceFlavorArrayInput
	Mode                pulumi.StringInput
	Name                pulumi.StringPtrInput
	Password            pulumi.StringPtrInput
	Period              pulumi.IntPtrInput
	PeriodUnit          pulumi.StringPtrInput
	Port                pulumi.IntPtrInput
	Region              pulumi.StringPtrInput
	SecurityGroupId     pulumi.StringInput
	Ssl                 pulumi.BoolPtrInput
	SubnetId            pulumi.StringInput
	Tags                pulumi.StringMapInput
	VpcId               pulumi.StringInput
}

func (DdsInstanceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ddsInstanceArgs)(nil)).Elem()
}

type DdsInstanceInput interface {
	pulumi.Input

	ToDdsInstanceOutput() DdsInstanceOutput
	ToDdsInstanceOutputWithContext(ctx context.Context) DdsInstanceOutput
}

func (*DdsInstance) ElementType() reflect.Type {
	return reflect.TypeOf((**DdsInstance)(nil)).Elem()
}

func (i *DdsInstance) ToDdsInstanceOutput() DdsInstanceOutput {
	return i.ToDdsInstanceOutputWithContext(context.Background())
}

func (i *DdsInstance) ToDdsInstanceOutputWithContext(ctx context.Context) DdsInstanceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DdsInstanceOutput)
}

// DdsInstanceArrayInput is an input type that accepts DdsInstanceArray and DdsInstanceArrayOutput values.
// You can construct a concrete instance of `DdsInstanceArrayInput` via:
//
//	DdsInstanceArray{ DdsInstanceArgs{...} }
type DdsInstanceArrayInput interface {
	pulumi.Input

	ToDdsInstanceArrayOutput() DdsInstanceArrayOutput
	ToDdsInstanceArrayOutputWithContext(context.Context) DdsInstanceArrayOutput
}

type DdsInstanceArray []DdsInstanceInput

func (DdsInstanceArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DdsInstance)(nil)).Elem()
}

func (i DdsInstanceArray) ToDdsInstanceArrayOutput() DdsInstanceArrayOutput {
	return i.ToDdsInstanceArrayOutputWithContext(context.Background())
}

func (i DdsInstanceArray) ToDdsInstanceArrayOutputWithContext(ctx context.Context) DdsInstanceArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DdsInstanceArrayOutput)
}

// DdsInstanceMapInput is an input type that accepts DdsInstanceMap and DdsInstanceMapOutput values.
// You can construct a concrete instance of `DdsInstanceMapInput` via:
//
//	DdsInstanceMap{ "key": DdsInstanceArgs{...} }
type DdsInstanceMapInput interface {
	pulumi.Input

	ToDdsInstanceMapOutput() DdsInstanceMapOutput
	ToDdsInstanceMapOutputWithContext(context.Context) DdsInstanceMapOutput
}

type DdsInstanceMap map[string]DdsInstanceInput

func (DdsInstanceMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DdsInstance)(nil)).Elem()
}

func (i DdsInstanceMap) ToDdsInstanceMapOutput() DdsInstanceMapOutput {
	return i.ToDdsInstanceMapOutputWithContext(context.Background())
}

func (i DdsInstanceMap) ToDdsInstanceMapOutputWithContext(ctx context.Context) DdsInstanceMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DdsInstanceMapOutput)
}

type DdsInstanceOutput struct{ *pulumi.OutputState }

func (DdsInstanceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DdsInstance)(nil)).Elem()
}

func (o DdsInstanceOutput) ToDdsInstanceOutput() DdsInstanceOutput {
	return o
}

func (o DdsInstanceOutput) ToDdsInstanceOutputWithContext(ctx context.Context) DdsInstanceOutput {
	return o
}

// Deprecated: Deprecated
func (o DdsInstanceOutput) AutoPay() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DdsInstance) pulumi.StringPtrOutput { return v.AutoPay }).(pulumi.StringPtrOutput)
}

func (o DdsInstanceOutput) AutoRenew() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DdsInstance) pulumi.StringPtrOutput { return v.AutoRenew }).(pulumi.StringPtrOutput)
}

func (o DdsInstanceOutput) AvailabilityZone() pulumi.StringOutput {
	return o.ApplyT(func(v *DdsInstance) pulumi.StringOutput { return v.AvailabilityZone }).(pulumi.StringOutput)
}

func (o DdsInstanceOutput) BackupStrategy() DdsInstanceBackupStrategyOutput {
	return o.ApplyT(func(v *DdsInstance) DdsInstanceBackupStrategyOutput { return v.BackupStrategy }).(DdsInstanceBackupStrategyOutput)
}

func (o DdsInstanceOutput) BalancerActiveBegin() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DdsInstance) pulumi.StringPtrOutput { return v.BalancerActiveBegin }).(pulumi.StringPtrOutput)
}

func (o DdsInstanceOutput) BalancerActiveEnd() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DdsInstance) pulumi.StringPtrOutput { return v.BalancerActiveEnd }).(pulumi.StringPtrOutput)
}

func (o DdsInstanceOutput) BalancerStatus() pulumi.StringOutput {
	return o.ApplyT(func(v *DdsInstance) pulumi.StringOutput { return v.BalancerStatus }).(pulumi.StringOutput)
}

func (o DdsInstanceOutput) ChargingMode() pulumi.StringOutput {
	return o.ApplyT(func(v *DdsInstance) pulumi.StringOutput { return v.ChargingMode }).(pulumi.StringOutput)
}

func (o DdsInstanceOutput) Configurations() DdsInstanceConfigurationArrayOutput {
	return o.ApplyT(func(v *DdsInstance) DdsInstanceConfigurationArrayOutput { return v.Configurations }).(DdsInstanceConfigurationArrayOutput)
}

func (o DdsInstanceOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *DdsInstance) pulumi.StringOutput { return v.CreatedAt }).(pulumi.StringOutput)
}

func (o DdsInstanceOutput) Datastore() DdsInstanceDatastoreOutput {
	return o.ApplyT(func(v *DdsInstance) DdsInstanceDatastoreOutput { return v.Datastore }).(DdsInstanceDatastoreOutput)
}

func (o DdsInstanceOutput) DbUsername() pulumi.StringOutput {
	return o.ApplyT(func(v *DdsInstance) pulumi.StringOutput { return v.DbUsername }).(pulumi.StringOutput)
}

func (o DdsInstanceOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DdsInstance) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

func (o DdsInstanceOutput) DiskEncryptionId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DdsInstance) pulumi.StringPtrOutput { return v.DiskEncryptionId }).(pulumi.StringPtrOutput)
}

func (o DdsInstanceOutput) EnterpriseProjectId() pulumi.StringOutput {
	return o.ApplyT(func(v *DdsInstance) pulumi.StringOutput { return v.EnterpriseProjectId }).(pulumi.StringOutput)
}

func (o DdsInstanceOutput) Flavors() DdsInstanceFlavorArrayOutput {
	return o.ApplyT(func(v *DdsInstance) DdsInstanceFlavorArrayOutput { return v.Flavors }).(DdsInstanceFlavorArrayOutput)
}

func (o DdsInstanceOutput) Groups() DdsInstanceGroupArrayOutput {
	return o.ApplyT(func(v *DdsInstance) DdsInstanceGroupArrayOutput { return v.Groups }).(DdsInstanceGroupArrayOutput)
}

func (o DdsInstanceOutput) Mode() pulumi.StringOutput {
	return o.ApplyT(func(v *DdsInstance) pulumi.StringOutput { return v.Mode }).(pulumi.StringOutput)
}

func (o DdsInstanceOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *DdsInstance) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// This field is deprecated.
func (o DdsInstanceOutput) Nodes() DdsInstanceNodeArrayOutput {
	return o.ApplyT(func(v *DdsInstance) DdsInstanceNodeArrayOutput { return v.Nodes }).(DdsInstanceNodeArrayOutput)
}

func (o DdsInstanceOutput) Password() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DdsInstance) pulumi.StringPtrOutput { return v.Password }).(pulumi.StringPtrOutput)
}

func (o DdsInstanceOutput) Period() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *DdsInstance) pulumi.IntPtrOutput { return v.Period }).(pulumi.IntPtrOutput)
}

func (o DdsInstanceOutput) PeriodUnit() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DdsInstance) pulumi.StringPtrOutput { return v.PeriodUnit }).(pulumi.StringPtrOutput)
}

func (o DdsInstanceOutput) Port() pulumi.IntOutput {
	return o.ApplyT(func(v *DdsInstance) pulumi.IntOutput { return v.Port }).(pulumi.IntOutput)
}

func (o DdsInstanceOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v *DdsInstance) pulumi.StringOutput { return v.Region }).(pulumi.StringOutput)
}

func (o DdsInstanceOutput) SecurityGroupId() pulumi.StringOutput {
	return o.ApplyT(func(v *DdsInstance) pulumi.StringOutput { return v.SecurityGroupId }).(pulumi.StringOutput)
}

func (o DdsInstanceOutput) Ssl() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *DdsInstance) pulumi.BoolPtrOutput { return v.Ssl }).(pulumi.BoolPtrOutput)
}

func (o DdsInstanceOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *DdsInstance) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

func (o DdsInstanceOutput) SubnetId() pulumi.StringOutput {
	return o.ApplyT(func(v *DdsInstance) pulumi.StringOutput { return v.SubnetId }).(pulumi.StringOutput)
}

func (o DdsInstanceOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *DdsInstance) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o DdsInstanceOutput) TimeZone() pulumi.StringOutput {
	return o.ApplyT(func(v *DdsInstance) pulumi.StringOutput { return v.TimeZone }).(pulumi.StringOutput)
}

func (o DdsInstanceOutput) UpdatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *DdsInstance) pulumi.StringOutput { return v.UpdatedAt }).(pulumi.StringOutput)
}

func (o DdsInstanceOutput) VpcId() pulumi.StringOutput {
	return o.ApplyT(func(v *DdsInstance) pulumi.StringOutput { return v.VpcId }).(pulumi.StringOutput)
}

type DdsInstanceArrayOutput struct{ *pulumi.OutputState }

func (DdsInstanceArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DdsInstance)(nil)).Elem()
}

func (o DdsInstanceArrayOutput) ToDdsInstanceArrayOutput() DdsInstanceArrayOutput {
	return o
}

func (o DdsInstanceArrayOutput) ToDdsInstanceArrayOutputWithContext(ctx context.Context) DdsInstanceArrayOutput {
	return o
}

func (o DdsInstanceArrayOutput) Index(i pulumi.IntInput) DdsInstanceOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *DdsInstance {
		return vs[0].([]*DdsInstance)[vs[1].(int)]
	}).(DdsInstanceOutput)
}

type DdsInstanceMapOutput struct{ *pulumi.OutputState }

func (DdsInstanceMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DdsInstance)(nil)).Elem()
}

func (o DdsInstanceMapOutput) ToDdsInstanceMapOutput() DdsInstanceMapOutput {
	return o
}

func (o DdsInstanceMapOutput) ToDdsInstanceMapOutputWithContext(ctx context.Context) DdsInstanceMapOutput {
	return o
}

func (o DdsInstanceMapOutput) MapIndex(k pulumi.StringInput) DdsInstanceOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *DdsInstance {
		return vs[0].(map[string]*DdsInstance)[vs[1].(string)]
	}).(DdsInstanceOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DdsInstanceInput)(nil)).Elem(), &DdsInstance{})
	pulumi.RegisterInputType(reflect.TypeOf((*DdsInstanceArrayInput)(nil)).Elem(), DdsInstanceArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*DdsInstanceMapInput)(nil)).Elem(), DdsInstanceMap{})
	pulumi.RegisterOutputType(DdsInstanceOutput{})
	pulumi.RegisterOutputType(DdsInstanceArrayOutput{})
	pulumi.RegisterOutputType(DdsInstanceMapOutput{})
}
