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

type GesGraph struct {
	pulumi.CustomResourceState

	// AZ code
	AzCode pulumi.StringOutput `pulumi:"azCode"`
	// Graph instance's CPU architecture type.
	CpuArch pulumi.StringOutput `pulumi:"cpuArch"`
	// Graph instance cryptography algorithm.
	CryptAlgorithm pulumi.StringOutput `pulumi:"cryptAlgorithm"`
	// Whether to enable full-text index control for the created graph.
	EnableFullTextIndex pulumi.BoolOutput `pulumi:"enableFullTextIndex"`
	// Whether to enable the security mode. This mode may damage GES performance greatly.
	EnableHttps pulumi.BoolOutput `pulumi:"enableHttps"`
	// Whether to enable HyG for the graph. This parameter is available for database edition graphs only.
	EnableHyg pulumi.BoolOutput `pulumi:"enableHyg"`
	// Whether the created graph supports the cross-AZ mode. The default value is false.
	EnableMultiAz pulumi.BoolOutput `pulumi:"enableMultiAz"`
	// Whether to enable granular permission control for the created graph.
	EnableRbac pulumi.BoolOutput `pulumi:"enableRbac"`
	// Whether to enable data encryption The value can be true or false.
	Encryption GesGraphEncryptionOutput `pulumi:"encryption"`
	// The enterprise project ID.
	EnterpriseProjectId pulumi.StringOutput `pulumi:"enterpriseProjectId"`
	// Graph size type index.
	GraphSizeTypeIndex pulumi.StringOutput `pulumi:"graphSizeTypeIndex"`
	// Whether to retain the backups of a graph after it is deleted.
	KeepBackup        pulumi.BoolOutput               `pulumi:"keepBackup"`
	LtsOperationTrace GesGraphLtsOperationTraceOutput `pulumi:"ltsOperationTrace"`
	// The graph name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Floating IP address of a graph instance.
	PrivateIp pulumi.StringOutput `pulumi:"privateIp"`
	// Graph product type
	ProductType pulumi.StringOutput `pulumi:"productType"`
	// The information about public IP.
	PublicIp    GesGraphPublicIpOutput `pulumi:"publicIp"`
	Region      pulumi.StringOutput    `pulumi:"region"`
	Replication pulumi.IntOutput       `pulumi:"replication"`
	// The security group ID.
	SecurityGroupId pulumi.StringOutput `pulumi:"securityGroupId"`
	// Status of a graph.
	Status pulumi.StringOutput `pulumi:"status"`
	// The subnet ID.
	SubnetId pulumi.StringOutput `pulumi:"subnetId"`
	// The key/value pairs to associate with the graph.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Physical addresses of a graph instance for access from private networks.
	TrafficIpLists pulumi.StringArrayOutput `pulumi:"trafficIpLists"`
	// ID type of vertices. This parameter is mandatory only for database edition graphs.
	VertexIdType GesGraphVertexIdTypeOutput `pulumi:"vertexIdType"`
	// The VPC ID.
	VpcId pulumi.StringOutput `pulumi:"vpcId"`
}

// NewGesGraph registers a new resource with the given unique name, arguments, and options.
func NewGesGraph(ctx *pulumi.Context,
	name string, args *GesGraphArgs, opts ...pulumi.ResourceOption) (*GesGraph, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.CryptAlgorithm == nil {
		return nil, errors.New("invalid value for required argument 'CryptAlgorithm'")
	}
	if args.EnableHttps == nil {
		return nil, errors.New("invalid value for required argument 'EnableHttps'")
	}
	if args.GraphSizeTypeIndex == nil {
		return nil, errors.New("invalid value for required argument 'GraphSizeTypeIndex'")
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
	var resource GesGraph
	err := ctx.RegisterResource("sbercloud:index/gesGraph:GesGraph", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetGesGraph gets an existing GesGraph resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetGesGraph(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *GesGraphState, opts ...pulumi.ResourceOption) (*GesGraph, error) {
	var resource GesGraph
	err := ctx.ReadResource("sbercloud:index/gesGraph:GesGraph", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering GesGraph resources.
type gesGraphState struct {
	// AZ code
	AzCode *string `pulumi:"azCode"`
	// Graph instance's CPU architecture type.
	CpuArch *string `pulumi:"cpuArch"`
	// Graph instance cryptography algorithm.
	CryptAlgorithm *string `pulumi:"cryptAlgorithm"`
	// Whether to enable full-text index control for the created graph.
	EnableFullTextIndex *bool `pulumi:"enableFullTextIndex"`
	// Whether to enable the security mode. This mode may damage GES performance greatly.
	EnableHttps *bool `pulumi:"enableHttps"`
	// Whether to enable HyG for the graph. This parameter is available for database edition graphs only.
	EnableHyg *bool `pulumi:"enableHyg"`
	// Whether the created graph supports the cross-AZ mode. The default value is false.
	EnableMultiAz *bool `pulumi:"enableMultiAz"`
	// Whether to enable granular permission control for the created graph.
	EnableRbac *bool `pulumi:"enableRbac"`
	// Whether to enable data encryption The value can be true or false.
	Encryption *GesGraphEncryption `pulumi:"encryption"`
	// The enterprise project ID.
	EnterpriseProjectId *string `pulumi:"enterpriseProjectId"`
	// Graph size type index.
	GraphSizeTypeIndex *string `pulumi:"graphSizeTypeIndex"`
	// Whether to retain the backups of a graph after it is deleted.
	KeepBackup        *bool                      `pulumi:"keepBackup"`
	LtsOperationTrace *GesGraphLtsOperationTrace `pulumi:"ltsOperationTrace"`
	// The graph name.
	Name *string `pulumi:"name"`
	// Floating IP address of a graph instance.
	PrivateIp *string `pulumi:"privateIp"`
	// Graph product type
	ProductType *string `pulumi:"productType"`
	// The information about public IP.
	PublicIp    *GesGraphPublicIp `pulumi:"publicIp"`
	Region      *string           `pulumi:"region"`
	Replication *int              `pulumi:"replication"`
	// The security group ID.
	SecurityGroupId *string `pulumi:"securityGroupId"`
	// Status of a graph.
	Status *string `pulumi:"status"`
	// The subnet ID.
	SubnetId *string `pulumi:"subnetId"`
	// The key/value pairs to associate with the graph.
	Tags map[string]string `pulumi:"tags"`
	// Physical addresses of a graph instance for access from private networks.
	TrafficIpLists []string `pulumi:"trafficIpLists"`
	// ID type of vertices. This parameter is mandatory only for database edition graphs.
	VertexIdType *GesGraphVertexIdType `pulumi:"vertexIdType"`
	// The VPC ID.
	VpcId *string `pulumi:"vpcId"`
}

type GesGraphState struct {
	// AZ code
	AzCode pulumi.StringPtrInput
	// Graph instance's CPU architecture type.
	CpuArch pulumi.StringPtrInput
	// Graph instance cryptography algorithm.
	CryptAlgorithm pulumi.StringPtrInput
	// Whether to enable full-text index control for the created graph.
	EnableFullTextIndex pulumi.BoolPtrInput
	// Whether to enable the security mode. This mode may damage GES performance greatly.
	EnableHttps pulumi.BoolPtrInput
	// Whether to enable HyG for the graph. This parameter is available for database edition graphs only.
	EnableHyg pulumi.BoolPtrInput
	// Whether the created graph supports the cross-AZ mode. The default value is false.
	EnableMultiAz pulumi.BoolPtrInput
	// Whether to enable granular permission control for the created graph.
	EnableRbac pulumi.BoolPtrInput
	// Whether to enable data encryption The value can be true or false.
	Encryption GesGraphEncryptionPtrInput
	// The enterprise project ID.
	EnterpriseProjectId pulumi.StringPtrInput
	// Graph size type index.
	GraphSizeTypeIndex pulumi.StringPtrInput
	// Whether to retain the backups of a graph after it is deleted.
	KeepBackup        pulumi.BoolPtrInput
	LtsOperationTrace GesGraphLtsOperationTracePtrInput
	// The graph name.
	Name pulumi.StringPtrInput
	// Floating IP address of a graph instance.
	PrivateIp pulumi.StringPtrInput
	// Graph product type
	ProductType pulumi.StringPtrInput
	// The information about public IP.
	PublicIp    GesGraphPublicIpPtrInput
	Region      pulumi.StringPtrInput
	Replication pulumi.IntPtrInput
	// The security group ID.
	SecurityGroupId pulumi.StringPtrInput
	// Status of a graph.
	Status pulumi.StringPtrInput
	// The subnet ID.
	SubnetId pulumi.StringPtrInput
	// The key/value pairs to associate with the graph.
	Tags pulumi.StringMapInput
	// Physical addresses of a graph instance for access from private networks.
	TrafficIpLists pulumi.StringArrayInput
	// ID type of vertices. This parameter is mandatory only for database edition graphs.
	VertexIdType GesGraphVertexIdTypePtrInput
	// The VPC ID.
	VpcId pulumi.StringPtrInput
}

func (GesGraphState) ElementType() reflect.Type {
	return reflect.TypeOf((*gesGraphState)(nil)).Elem()
}

type gesGraphArgs struct {
	// Graph instance's CPU architecture type.
	CpuArch *string `pulumi:"cpuArch"`
	// Graph instance cryptography algorithm.
	CryptAlgorithm string `pulumi:"cryptAlgorithm"`
	// Whether to enable full-text index control for the created graph.
	EnableFullTextIndex *bool `pulumi:"enableFullTextIndex"`
	// Whether to enable the security mode. This mode may damage GES performance greatly.
	EnableHttps bool `pulumi:"enableHttps"`
	// Whether to enable HyG for the graph. This parameter is available for database edition graphs only.
	EnableHyg *bool `pulumi:"enableHyg"`
	// Whether the created graph supports the cross-AZ mode. The default value is false.
	EnableMultiAz *bool `pulumi:"enableMultiAz"`
	// Whether to enable granular permission control for the created graph.
	EnableRbac *bool `pulumi:"enableRbac"`
	// Whether to enable data encryption The value can be true or false.
	Encryption *GesGraphEncryption `pulumi:"encryption"`
	// The enterprise project ID.
	EnterpriseProjectId *string `pulumi:"enterpriseProjectId"`
	// Graph size type index.
	GraphSizeTypeIndex string `pulumi:"graphSizeTypeIndex"`
	// Whether to retain the backups of a graph after it is deleted.
	KeepBackup        *bool                      `pulumi:"keepBackup"`
	LtsOperationTrace *GesGraphLtsOperationTrace `pulumi:"ltsOperationTrace"`
	// The graph name.
	Name *string `pulumi:"name"`
	// Graph product type
	ProductType *string `pulumi:"productType"`
	// The information about public IP.
	PublicIp    *GesGraphPublicIp `pulumi:"publicIp"`
	Region      *string           `pulumi:"region"`
	Replication *int              `pulumi:"replication"`
	// The security group ID.
	SecurityGroupId string `pulumi:"securityGroupId"`
	// The subnet ID.
	SubnetId string `pulumi:"subnetId"`
	// The key/value pairs to associate with the graph.
	Tags map[string]string `pulumi:"tags"`
	// ID type of vertices. This parameter is mandatory only for database edition graphs.
	VertexIdType *GesGraphVertexIdType `pulumi:"vertexIdType"`
	// The VPC ID.
	VpcId string `pulumi:"vpcId"`
}

// The set of arguments for constructing a GesGraph resource.
type GesGraphArgs struct {
	// Graph instance's CPU architecture type.
	CpuArch pulumi.StringPtrInput
	// Graph instance cryptography algorithm.
	CryptAlgorithm pulumi.StringInput
	// Whether to enable full-text index control for the created graph.
	EnableFullTextIndex pulumi.BoolPtrInput
	// Whether to enable the security mode. This mode may damage GES performance greatly.
	EnableHttps pulumi.BoolInput
	// Whether to enable HyG for the graph. This parameter is available for database edition graphs only.
	EnableHyg pulumi.BoolPtrInput
	// Whether the created graph supports the cross-AZ mode. The default value is false.
	EnableMultiAz pulumi.BoolPtrInput
	// Whether to enable granular permission control for the created graph.
	EnableRbac pulumi.BoolPtrInput
	// Whether to enable data encryption The value can be true or false.
	Encryption GesGraphEncryptionPtrInput
	// The enterprise project ID.
	EnterpriseProjectId pulumi.StringPtrInput
	// Graph size type index.
	GraphSizeTypeIndex pulumi.StringInput
	// Whether to retain the backups of a graph after it is deleted.
	KeepBackup        pulumi.BoolPtrInput
	LtsOperationTrace GesGraphLtsOperationTracePtrInput
	// The graph name.
	Name pulumi.StringPtrInput
	// Graph product type
	ProductType pulumi.StringPtrInput
	// The information about public IP.
	PublicIp    GesGraphPublicIpPtrInput
	Region      pulumi.StringPtrInput
	Replication pulumi.IntPtrInput
	// The security group ID.
	SecurityGroupId pulumi.StringInput
	// The subnet ID.
	SubnetId pulumi.StringInput
	// The key/value pairs to associate with the graph.
	Tags pulumi.StringMapInput
	// ID type of vertices. This parameter is mandatory only for database edition graphs.
	VertexIdType GesGraphVertexIdTypePtrInput
	// The VPC ID.
	VpcId pulumi.StringInput
}

func (GesGraphArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*gesGraphArgs)(nil)).Elem()
}

type GesGraphInput interface {
	pulumi.Input

	ToGesGraphOutput() GesGraphOutput
	ToGesGraphOutputWithContext(ctx context.Context) GesGraphOutput
}

func (*GesGraph) ElementType() reflect.Type {
	return reflect.TypeOf((**GesGraph)(nil)).Elem()
}

func (i *GesGraph) ToGesGraphOutput() GesGraphOutput {
	return i.ToGesGraphOutputWithContext(context.Background())
}

func (i *GesGraph) ToGesGraphOutputWithContext(ctx context.Context) GesGraphOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GesGraphOutput)
}

// GesGraphArrayInput is an input type that accepts GesGraphArray and GesGraphArrayOutput values.
// You can construct a concrete instance of `GesGraphArrayInput` via:
//
//	GesGraphArray{ GesGraphArgs{...} }
type GesGraphArrayInput interface {
	pulumi.Input

	ToGesGraphArrayOutput() GesGraphArrayOutput
	ToGesGraphArrayOutputWithContext(context.Context) GesGraphArrayOutput
}

type GesGraphArray []GesGraphInput

func (GesGraphArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*GesGraph)(nil)).Elem()
}

func (i GesGraphArray) ToGesGraphArrayOutput() GesGraphArrayOutput {
	return i.ToGesGraphArrayOutputWithContext(context.Background())
}

func (i GesGraphArray) ToGesGraphArrayOutputWithContext(ctx context.Context) GesGraphArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GesGraphArrayOutput)
}

// GesGraphMapInput is an input type that accepts GesGraphMap and GesGraphMapOutput values.
// You can construct a concrete instance of `GesGraphMapInput` via:
//
//	GesGraphMap{ "key": GesGraphArgs{...} }
type GesGraphMapInput interface {
	pulumi.Input

	ToGesGraphMapOutput() GesGraphMapOutput
	ToGesGraphMapOutputWithContext(context.Context) GesGraphMapOutput
}

type GesGraphMap map[string]GesGraphInput

func (GesGraphMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*GesGraph)(nil)).Elem()
}

func (i GesGraphMap) ToGesGraphMapOutput() GesGraphMapOutput {
	return i.ToGesGraphMapOutputWithContext(context.Background())
}

func (i GesGraphMap) ToGesGraphMapOutputWithContext(ctx context.Context) GesGraphMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GesGraphMapOutput)
}

type GesGraphOutput struct{ *pulumi.OutputState }

func (GesGraphOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**GesGraph)(nil)).Elem()
}

func (o GesGraphOutput) ToGesGraphOutput() GesGraphOutput {
	return o
}

func (o GesGraphOutput) ToGesGraphOutputWithContext(ctx context.Context) GesGraphOutput {
	return o
}

// AZ code
func (o GesGraphOutput) AzCode() pulumi.StringOutput {
	return o.ApplyT(func(v *GesGraph) pulumi.StringOutput { return v.AzCode }).(pulumi.StringOutput)
}

// Graph instance's CPU architecture type.
func (o GesGraphOutput) CpuArch() pulumi.StringOutput {
	return o.ApplyT(func(v *GesGraph) pulumi.StringOutput { return v.CpuArch }).(pulumi.StringOutput)
}

// Graph instance cryptography algorithm.
func (o GesGraphOutput) CryptAlgorithm() pulumi.StringOutput {
	return o.ApplyT(func(v *GesGraph) pulumi.StringOutput { return v.CryptAlgorithm }).(pulumi.StringOutput)
}

// Whether to enable full-text index control for the created graph.
func (o GesGraphOutput) EnableFullTextIndex() pulumi.BoolOutput {
	return o.ApplyT(func(v *GesGraph) pulumi.BoolOutput { return v.EnableFullTextIndex }).(pulumi.BoolOutput)
}

// Whether to enable the security mode. This mode may damage GES performance greatly.
func (o GesGraphOutput) EnableHttps() pulumi.BoolOutput {
	return o.ApplyT(func(v *GesGraph) pulumi.BoolOutput { return v.EnableHttps }).(pulumi.BoolOutput)
}

// Whether to enable HyG for the graph. This parameter is available for database edition graphs only.
func (o GesGraphOutput) EnableHyg() pulumi.BoolOutput {
	return o.ApplyT(func(v *GesGraph) pulumi.BoolOutput { return v.EnableHyg }).(pulumi.BoolOutput)
}

// Whether the created graph supports the cross-AZ mode. The default value is false.
func (o GesGraphOutput) EnableMultiAz() pulumi.BoolOutput {
	return o.ApplyT(func(v *GesGraph) pulumi.BoolOutput { return v.EnableMultiAz }).(pulumi.BoolOutput)
}

// Whether to enable granular permission control for the created graph.
func (o GesGraphOutput) EnableRbac() pulumi.BoolOutput {
	return o.ApplyT(func(v *GesGraph) pulumi.BoolOutput { return v.EnableRbac }).(pulumi.BoolOutput)
}

// Whether to enable data encryption The value can be true or false.
func (o GesGraphOutput) Encryption() GesGraphEncryptionOutput {
	return o.ApplyT(func(v *GesGraph) GesGraphEncryptionOutput { return v.Encryption }).(GesGraphEncryptionOutput)
}

// The enterprise project ID.
func (o GesGraphOutput) EnterpriseProjectId() pulumi.StringOutput {
	return o.ApplyT(func(v *GesGraph) pulumi.StringOutput { return v.EnterpriseProjectId }).(pulumi.StringOutput)
}

// Graph size type index.
func (o GesGraphOutput) GraphSizeTypeIndex() pulumi.StringOutput {
	return o.ApplyT(func(v *GesGraph) pulumi.StringOutput { return v.GraphSizeTypeIndex }).(pulumi.StringOutput)
}

// Whether to retain the backups of a graph after it is deleted.
func (o GesGraphOutput) KeepBackup() pulumi.BoolOutput {
	return o.ApplyT(func(v *GesGraph) pulumi.BoolOutput { return v.KeepBackup }).(pulumi.BoolOutput)
}

func (o GesGraphOutput) LtsOperationTrace() GesGraphLtsOperationTraceOutput {
	return o.ApplyT(func(v *GesGraph) GesGraphLtsOperationTraceOutput { return v.LtsOperationTrace }).(GesGraphLtsOperationTraceOutput)
}

// The graph name.
func (o GesGraphOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *GesGraph) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Floating IP address of a graph instance.
func (o GesGraphOutput) PrivateIp() pulumi.StringOutput {
	return o.ApplyT(func(v *GesGraph) pulumi.StringOutput { return v.PrivateIp }).(pulumi.StringOutput)
}

// Graph product type
func (o GesGraphOutput) ProductType() pulumi.StringOutput {
	return o.ApplyT(func(v *GesGraph) pulumi.StringOutput { return v.ProductType }).(pulumi.StringOutput)
}

// The information about public IP.
func (o GesGraphOutput) PublicIp() GesGraphPublicIpOutput {
	return o.ApplyT(func(v *GesGraph) GesGraphPublicIpOutput { return v.PublicIp }).(GesGraphPublicIpOutput)
}

func (o GesGraphOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v *GesGraph) pulumi.StringOutput { return v.Region }).(pulumi.StringOutput)
}

func (o GesGraphOutput) Replication() pulumi.IntOutput {
	return o.ApplyT(func(v *GesGraph) pulumi.IntOutput { return v.Replication }).(pulumi.IntOutput)
}

// The security group ID.
func (o GesGraphOutput) SecurityGroupId() pulumi.StringOutput {
	return o.ApplyT(func(v *GesGraph) pulumi.StringOutput { return v.SecurityGroupId }).(pulumi.StringOutput)
}

// Status of a graph.
func (o GesGraphOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *GesGraph) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

// The subnet ID.
func (o GesGraphOutput) SubnetId() pulumi.StringOutput {
	return o.ApplyT(func(v *GesGraph) pulumi.StringOutput { return v.SubnetId }).(pulumi.StringOutput)
}

// The key/value pairs to associate with the graph.
func (o GesGraphOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *GesGraph) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// Physical addresses of a graph instance for access from private networks.
func (o GesGraphOutput) TrafficIpLists() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *GesGraph) pulumi.StringArrayOutput { return v.TrafficIpLists }).(pulumi.StringArrayOutput)
}

// ID type of vertices. This parameter is mandatory only for database edition graphs.
func (o GesGraphOutput) VertexIdType() GesGraphVertexIdTypeOutput {
	return o.ApplyT(func(v *GesGraph) GesGraphVertexIdTypeOutput { return v.VertexIdType }).(GesGraphVertexIdTypeOutput)
}

// The VPC ID.
func (o GesGraphOutput) VpcId() pulumi.StringOutput {
	return o.ApplyT(func(v *GesGraph) pulumi.StringOutput { return v.VpcId }).(pulumi.StringOutput)
}

type GesGraphArrayOutput struct{ *pulumi.OutputState }

func (GesGraphArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*GesGraph)(nil)).Elem()
}

func (o GesGraphArrayOutput) ToGesGraphArrayOutput() GesGraphArrayOutput {
	return o
}

func (o GesGraphArrayOutput) ToGesGraphArrayOutputWithContext(ctx context.Context) GesGraphArrayOutput {
	return o
}

func (o GesGraphArrayOutput) Index(i pulumi.IntInput) GesGraphOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *GesGraph {
		return vs[0].([]*GesGraph)[vs[1].(int)]
	}).(GesGraphOutput)
}

type GesGraphMapOutput struct{ *pulumi.OutputState }

func (GesGraphMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*GesGraph)(nil)).Elem()
}

func (o GesGraphMapOutput) ToGesGraphMapOutput() GesGraphMapOutput {
	return o
}

func (o GesGraphMapOutput) ToGesGraphMapOutputWithContext(ctx context.Context) GesGraphMapOutput {
	return o
}

func (o GesGraphMapOutput) MapIndex(k pulumi.StringInput) GesGraphOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *GesGraph {
		return vs[0].(map[string]*GesGraph)[vs[1].(string)]
	}).(GesGraphOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*GesGraphInput)(nil)).Elem(), &GesGraph{})
	pulumi.RegisterInputType(reflect.TypeOf((*GesGraphArrayInput)(nil)).Elem(), GesGraphArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*GesGraphMapInput)(nil)).Elem(), GesGraphMap{})
	pulumi.RegisterOutputType(GesGraphOutput{})
	pulumi.RegisterOutputType(GesGraphArrayOutput{})
	pulumi.RegisterOutputType(GesGraphMapOutput{})
}
