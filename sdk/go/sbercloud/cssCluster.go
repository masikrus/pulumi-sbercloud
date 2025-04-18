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

type CssCluster struct {
	pulumi.CustomResourceState

	BackupStrategy      CssClusterBackupStrategyPtrOutput `pulumi:"backupStrategy"`
	Created             pulumi.StringOutput               `pulumi:"created"`
	Endpoint            pulumi.StringOutput               `pulumi:"endpoint"`
	EngineType          pulumi.StringPtrOutput            `pulumi:"engineType"`
	EngineVersion       pulumi.StringOutput               `pulumi:"engineVersion"`
	EnterpriseProjectId pulumi.StringOutput               `pulumi:"enterpriseProjectId"`
	ExpectNodeNum       pulumi.IntPtrOutput               `pulumi:"expectNodeNum"`
	Name                pulumi.StringOutput               `pulumi:"name"`
	NodeConfig          CssClusterNodeConfigOutput        `pulumi:"nodeConfig"`
	Nodes               CssClusterNodeArrayOutput         `pulumi:"nodes"`
	Password            pulumi.StringPtrOutput            `pulumi:"password"`
	Region              pulumi.StringOutput               `pulumi:"region"`
	SecurityMode        pulumi.BoolPtrOutput              `pulumi:"securityMode"`
	Status              pulumi.StringOutput               `pulumi:"status"`
	Tags                pulumi.StringMapOutput            `pulumi:"tags"`
}

// NewCssCluster registers a new resource with the given unique name, arguments, and options.
func NewCssCluster(ctx *pulumi.Context,
	name string, args *CssClusterArgs, opts ...pulumi.ResourceOption) (*CssCluster, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.EngineVersion == nil {
		return nil, errors.New("invalid value for required argument 'EngineVersion'")
	}
	if args.NodeConfig == nil {
		return nil, errors.New("invalid value for required argument 'NodeConfig'")
	}
	if args.Password != nil {
		args.Password = pulumi.ToSecret(args.Password).(pulumi.StringPtrInput)
	}
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"password",
	})
	opts = append(opts, secrets)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource CssCluster
	err := ctx.RegisterResource("sbercloud:index/cssCluster:CssCluster", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCssCluster gets an existing CssCluster resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCssCluster(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CssClusterState, opts ...pulumi.ResourceOption) (*CssCluster, error) {
	var resource CssCluster
	err := ctx.ReadResource("sbercloud:index/cssCluster:CssCluster", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering CssCluster resources.
type cssClusterState struct {
	BackupStrategy      *CssClusterBackupStrategy `pulumi:"backupStrategy"`
	Created             *string                   `pulumi:"created"`
	Endpoint            *string                   `pulumi:"endpoint"`
	EngineType          *string                   `pulumi:"engineType"`
	EngineVersion       *string                   `pulumi:"engineVersion"`
	EnterpriseProjectId *string                   `pulumi:"enterpriseProjectId"`
	ExpectNodeNum       *int                      `pulumi:"expectNodeNum"`
	Name                *string                   `pulumi:"name"`
	NodeConfig          *CssClusterNodeConfig     `pulumi:"nodeConfig"`
	Nodes               []CssClusterNode          `pulumi:"nodes"`
	Password            *string                   `pulumi:"password"`
	Region              *string                   `pulumi:"region"`
	SecurityMode        *bool                     `pulumi:"securityMode"`
	Status              *string                   `pulumi:"status"`
	Tags                map[string]string         `pulumi:"tags"`
}

type CssClusterState struct {
	BackupStrategy      CssClusterBackupStrategyPtrInput
	Created             pulumi.StringPtrInput
	Endpoint            pulumi.StringPtrInput
	EngineType          pulumi.StringPtrInput
	EngineVersion       pulumi.StringPtrInput
	EnterpriseProjectId pulumi.StringPtrInput
	ExpectNodeNum       pulumi.IntPtrInput
	Name                pulumi.StringPtrInput
	NodeConfig          CssClusterNodeConfigPtrInput
	Nodes               CssClusterNodeArrayInput
	Password            pulumi.StringPtrInput
	Region              pulumi.StringPtrInput
	SecurityMode        pulumi.BoolPtrInput
	Status              pulumi.StringPtrInput
	Tags                pulumi.StringMapInput
}

func (CssClusterState) ElementType() reflect.Type {
	return reflect.TypeOf((*cssClusterState)(nil)).Elem()
}

type cssClusterArgs struct {
	BackupStrategy      *CssClusterBackupStrategy `pulumi:"backupStrategy"`
	EngineType          *string                   `pulumi:"engineType"`
	EngineVersion       string                    `pulumi:"engineVersion"`
	EnterpriseProjectId *string                   `pulumi:"enterpriseProjectId"`
	ExpectNodeNum       *int                      `pulumi:"expectNodeNum"`
	Name                *string                   `pulumi:"name"`
	NodeConfig          CssClusterNodeConfig      `pulumi:"nodeConfig"`
	Password            *string                   `pulumi:"password"`
	Region              *string                   `pulumi:"region"`
	SecurityMode        *bool                     `pulumi:"securityMode"`
	Tags                map[string]string         `pulumi:"tags"`
}

// The set of arguments for constructing a CssCluster resource.
type CssClusterArgs struct {
	BackupStrategy      CssClusterBackupStrategyPtrInput
	EngineType          pulumi.StringPtrInput
	EngineVersion       pulumi.StringInput
	EnterpriseProjectId pulumi.StringPtrInput
	ExpectNodeNum       pulumi.IntPtrInput
	Name                pulumi.StringPtrInput
	NodeConfig          CssClusterNodeConfigInput
	Password            pulumi.StringPtrInput
	Region              pulumi.StringPtrInput
	SecurityMode        pulumi.BoolPtrInput
	Tags                pulumi.StringMapInput
}

func (CssClusterArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*cssClusterArgs)(nil)).Elem()
}

type CssClusterInput interface {
	pulumi.Input

	ToCssClusterOutput() CssClusterOutput
	ToCssClusterOutputWithContext(ctx context.Context) CssClusterOutput
}

func (*CssCluster) ElementType() reflect.Type {
	return reflect.TypeOf((**CssCluster)(nil)).Elem()
}

func (i *CssCluster) ToCssClusterOutput() CssClusterOutput {
	return i.ToCssClusterOutputWithContext(context.Background())
}

func (i *CssCluster) ToCssClusterOutputWithContext(ctx context.Context) CssClusterOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CssClusterOutput)
}

// CssClusterArrayInput is an input type that accepts CssClusterArray and CssClusterArrayOutput values.
// You can construct a concrete instance of `CssClusterArrayInput` via:
//
//	CssClusterArray{ CssClusterArgs{...} }
type CssClusterArrayInput interface {
	pulumi.Input

	ToCssClusterArrayOutput() CssClusterArrayOutput
	ToCssClusterArrayOutputWithContext(context.Context) CssClusterArrayOutput
}

type CssClusterArray []CssClusterInput

func (CssClusterArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*CssCluster)(nil)).Elem()
}

func (i CssClusterArray) ToCssClusterArrayOutput() CssClusterArrayOutput {
	return i.ToCssClusterArrayOutputWithContext(context.Background())
}

func (i CssClusterArray) ToCssClusterArrayOutputWithContext(ctx context.Context) CssClusterArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CssClusterArrayOutput)
}

// CssClusterMapInput is an input type that accepts CssClusterMap and CssClusterMapOutput values.
// You can construct a concrete instance of `CssClusterMapInput` via:
//
//	CssClusterMap{ "key": CssClusterArgs{...} }
type CssClusterMapInput interface {
	pulumi.Input

	ToCssClusterMapOutput() CssClusterMapOutput
	ToCssClusterMapOutputWithContext(context.Context) CssClusterMapOutput
}

type CssClusterMap map[string]CssClusterInput

func (CssClusterMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*CssCluster)(nil)).Elem()
}

func (i CssClusterMap) ToCssClusterMapOutput() CssClusterMapOutput {
	return i.ToCssClusterMapOutputWithContext(context.Background())
}

func (i CssClusterMap) ToCssClusterMapOutputWithContext(ctx context.Context) CssClusterMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CssClusterMapOutput)
}

type CssClusterOutput struct{ *pulumi.OutputState }

func (CssClusterOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CssCluster)(nil)).Elem()
}

func (o CssClusterOutput) ToCssClusterOutput() CssClusterOutput {
	return o
}

func (o CssClusterOutput) ToCssClusterOutputWithContext(ctx context.Context) CssClusterOutput {
	return o
}

func (o CssClusterOutput) BackupStrategy() CssClusterBackupStrategyPtrOutput {
	return o.ApplyT(func(v *CssCluster) CssClusterBackupStrategyPtrOutput { return v.BackupStrategy }).(CssClusterBackupStrategyPtrOutput)
}

func (o CssClusterOutput) Created() pulumi.StringOutput {
	return o.ApplyT(func(v *CssCluster) pulumi.StringOutput { return v.Created }).(pulumi.StringOutput)
}

func (o CssClusterOutput) Endpoint() pulumi.StringOutput {
	return o.ApplyT(func(v *CssCluster) pulumi.StringOutput { return v.Endpoint }).(pulumi.StringOutput)
}

func (o CssClusterOutput) EngineType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CssCluster) pulumi.StringPtrOutput { return v.EngineType }).(pulumi.StringPtrOutput)
}

func (o CssClusterOutput) EngineVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *CssCluster) pulumi.StringOutput { return v.EngineVersion }).(pulumi.StringOutput)
}

func (o CssClusterOutput) EnterpriseProjectId() pulumi.StringOutput {
	return o.ApplyT(func(v *CssCluster) pulumi.StringOutput { return v.EnterpriseProjectId }).(pulumi.StringOutput)
}

func (o CssClusterOutput) ExpectNodeNum() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *CssCluster) pulumi.IntPtrOutput { return v.ExpectNodeNum }).(pulumi.IntPtrOutput)
}

func (o CssClusterOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *CssCluster) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o CssClusterOutput) NodeConfig() CssClusterNodeConfigOutput {
	return o.ApplyT(func(v *CssCluster) CssClusterNodeConfigOutput { return v.NodeConfig }).(CssClusterNodeConfigOutput)
}

func (o CssClusterOutput) Nodes() CssClusterNodeArrayOutput {
	return o.ApplyT(func(v *CssCluster) CssClusterNodeArrayOutput { return v.Nodes }).(CssClusterNodeArrayOutput)
}

func (o CssClusterOutput) Password() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CssCluster) pulumi.StringPtrOutput { return v.Password }).(pulumi.StringPtrOutput)
}

func (o CssClusterOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v *CssCluster) pulumi.StringOutput { return v.Region }).(pulumi.StringOutput)
}

func (o CssClusterOutput) SecurityMode() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *CssCluster) pulumi.BoolPtrOutput { return v.SecurityMode }).(pulumi.BoolPtrOutput)
}

func (o CssClusterOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *CssCluster) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

func (o CssClusterOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *CssCluster) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

type CssClusterArrayOutput struct{ *pulumi.OutputState }

func (CssClusterArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*CssCluster)(nil)).Elem()
}

func (o CssClusterArrayOutput) ToCssClusterArrayOutput() CssClusterArrayOutput {
	return o
}

func (o CssClusterArrayOutput) ToCssClusterArrayOutputWithContext(ctx context.Context) CssClusterArrayOutput {
	return o
}

func (o CssClusterArrayOutput) Index(i pulumi.IntInput) CssClusterOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *CssCluster {
		return vs[0].([]*CssCluster)[vs[1].(int)]
	}).(CssClusterOutput)
}

type CssClusterMapOutput struct{ *pulumi.OutputState }

func (CssClusterMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*CssCluster)(nil)).Elem()
}

func (o CssClusterMapOutput) ToCssClusterMapOutput() CssClusterMapOutput {
	return o
}

func (o CssClusterMapOutput) ToCssClusterMapOutputWithContext(ctx context.Context) CssClusterMapOutput {
	return o
}

func (o CssClusterMapOutput) MapIndex(k pulumi.StringInput) CssClusterOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *CssCluster {
		return vs[0].(map[string]*CssCluster)[vs[1].(string)]
	}).(CssClusterOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*CssClusterInput)(nil)).Elem(), &CssCluster{})
	pulumi.RegisterInputType(reflect.TypeOf((*CssClusterArrayInput)(nil)).Elem(), CssClusterArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*CssClusterMapInput)(nil)).Elem(), CssClusterMap{})
	pulumi.RegisterOutputType(CssClusterOutput{})
	pulumi.RegisterOutputType(CssClusterArrayOutput{})
	pulumi.RegisterOutputType(CssClusterMapOutput{})
}
