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

type ObsBucket struct {
	pulumi.CustomResourceState

	Acl                 pulumi.StringPtrOutput            `pulumi:"acl"`
	Bucket              pulumi.StringOutput               `pulumi:"bucket"`
	BucketDomainName    pulumi.StringOutput               `pulumi:"bucketDomainName"`
	BucketVersion       pulumi.StringOutput               `pulumi:"bucketVersion"`
	CorsRules           ObsBucketCorsRuleArrayOutput      `pulumi:"corsRules"`
	Encryption          pulumi.BoolPtrOutput              `pulumi:"encryption"`
	EnterpriseProjectId pulumi.StringOutput               `pulumi:"enterpriseProjectId"`
	ForceDestroy        pulumi.BoolPtrOutput              `pulumi:"forceDestroy"`
	KmsKeyId            pulumi.StringPtrOutput            `pulumi:"kmsKeyId"`
	KmsKeyProjectId     pulumi.StringOutput               `pulumi:"kmsKeyProjectId"`
	LifecycleRules      ObsBucketLifecycleRuleArrayOutput `pulumi:"lifecycleRules"`
	Loggings            ObsBucketLoggingArrayOutput       `pulumi:"loggings"`
	MultiAz             pulumi.BoolOutput                 `pulumi:"multiAz"`
	ParallelFs          pulumi.BoolPtrOutput              `pulumi:"parallelFs"`
	Policy              pulumi.StringOutput               `pulumi:"policy"`
	PolicyFormat        pulumi.StringPtrOutput            `pulumi:"policyFormat"`
	Quota               pulumi.IntPtrOutput               `pulumi:"quota"`
	Region              pulumi.StringOutput               `pulumi:"region"`
	SseAlgorithm        pulumi.StringOutput               `pulumi:"sseAlgorithm"`
	StorageClass        pulumi.StringPtrOutput            `pulumi:"storageClass"`
	StorageInfos        ObsBucketStorageInfoArrayOutput   `pulumi:"storageInfos"`
	Tags                pulumi.StringMapOutput            `pulumi:"tags"`
	UserDomainNames     pulumi.StringArrayOutput          `pulumi:"userDomainNames"`
	Versioning          pulumi.BoolPtrOutput              `pulumi:"versioning"`
	Website             ObsBucketWebsitePtrOutput         `pulumi:"website"`
}

// NewObsBucket registers a new resource with the given unique name, arguments, and options.
func NewObsBucket(ctx *pulumi.Context,
	name string, args *ObsBucketArgs, opts ...pulumi.ResourceOption) (*ObsBucket, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Bucket == nil {
		return nil, errors.New("invalid value for required argument 'Bucket'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource ObsBucket
	err := ctx.RegisterResource("sbercloud:index/obsBucket:ObsBucket", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetObsBucket gets an existing ObsBucket resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetObsBucket(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ObsBucketState, opts ...pulumi.ResourceOption) (*ObsBucket, error) {
	var resource ObsBucket
	err := ctx.ReadResource("sbercloud:index/obsBucket:ObsBucket", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ObsBucket resources.
type obsBucketState struct {
	Acl                 *string                  `pulumi:"acl"`
	Bucket              *string                  `pulumi:"bucket"`
	BucketDomainName    *string                  `pulumi:"bucketDomainName"`
	BucketVersion       *string                  `pulumi:"bucketVersion"`
	CorsRules           []ObsBucketCorsRule      `pulumi:"corsRules"`
	Encryption          *bool                    `pulumi:"encryption"`
	EnterpriseProjectId *string                  `pulumi:"enterpriseProjectId"`
	ForceDestroy        *bool                    `pulumi:"forceDestroy"`
	KmsKeyId            *string                  `pulumi:"kmsKeyId"`
	KmsKeyProjectId     *string                  `pulumi:"kmsKeyProjectId"`
	LifecycleRules      []ObsBucketLifecycleRule `pulumi:"lifecycleRules"`
	Loggings            []ObsBucketLogging       `pulumi:"loggings"`
	MultiAz             *bool                    `pulumi:"multiAz"`
	ParallelFs          *bool                    `pulumi:"parallelFs"`
	Policy              *string                  `pulumi:"policy"`
	PolicyFormat        *string                  `pulumi:"policyFormat"`
	Quota               *int                     `pulumi:"quota"`
	Region              *string                  `pulumi:"region"`
	SseAlgorithm        *string                  `pulumi:"sseAlgorithm"`
	StorageClass        *string                  `pulumi:"storageClass"`
	StorageInfos        []ObsBucketStorageInfo   `pulumi:"storageInfos"`
	Tags                map[string]string        `pulumi:"tags"`
	UserDomainNames     []string                 `pulumi:"userDomainNames"`
	Versioning          *bool                    `pulumi:"versioning"`
	Website             *ObsBucketWebsite        `pulumi:"website"`
}

type ObsBucketState struct {
	Acl                 pulumi.StringPtrInput
	Bucket              pulumi.StringPtrInput
	BucketDomainName    pulumi.StringPtrInput
	BucketVersion       pulumi.StringPtrInput
	CorsRules           ObsBucketCorsRuleArrayInput
	Encryption          pulumi.BoolPtrInput
	EnterpriseProjectId pulumi.StringPtrInput
	ForceDestroy        pulumi.BoolPtrInput
	KmsKeyId            pulumi.StringPtrInput
	KmsKeyProjectId     pulumi.StringPtrInput
	LifecycleRules      ObsBucketLifecycleRuleArrayInput
	Loggings            ObsBucketLoggingArrayInput
	MultiAz             pulumi.BoolPtrInput
	ParallelFs          pulumi.BoolPtrInput
	Policy              pulumi.StringPtrInput
	PolicyFormat        pulumi.StringPtrInput
	Quota               pulumi.IntPtrInput
	Region              pulumi.StringPtrInput
	SseAlgorithm        pulumi.StringPtrInput
	StorageClass        pulumi.StringPtrInput
	StorageInfos        ObsBucketStorageInfoArrayInput
	Tags                pulumi.StringMapInput
	UserDomainNames     pulumi.StringArrayInput
	Versioning          pulumi.BoolPtrInput
	Website             ObsBucketWebsitePtrInput
}

func (ObsBucketState) ElementType() reflect.Type {
	return reflect.TypeOf((*obsBucketState)(nil)).Elem()
}

type obsBucketArgs struct {
	Acl                 *string                  `pulumi:"acl"`
	Bucket              string                   `pulumi:"bucket"`
	CorsRules           []ObsBucketCorsRule      `pulumi:"corsRules"`
	Encryption          *bool                    `pulumi:"encryption"`
	EnterpriseProjectId *string                  `pulumi:"enterpriseProjectId"`
	ForceDestroy        *bool                    `pulumi:"forceDestroy"`
	KmsKeyId            *string                  `pulumi:"kmsKeyId"`
	KmsKeyProjectId     *string                  `pulumi:"kmsKeyProjectId"`
	LifecycleRules      []ObsBucketLifecycleRule `pulumi:"lifecycleRules"`
	Loggings            []ObsBucketLogging       `pulumi:"loggings"`
	MultiAz             *bool                    `pulumi:"multiAz"`
	ParallelFs          *bool                    `pulumi:"parallelFs"`
	Policy              *string                  `pulumi:"policy"`
	PolicyFormat        *string                  `pulumi:"policyFormat"`
	Quota               *int                     `pulumi:"quota"`
	Region              *string                  `pulumi:"region"`
	SseAlgorithm        *string                  `pulumi:"sseAlgorithm"`
	StorageClass        *string                  `pulumi:"storageClass"`
	Tags                map[string]string        `pulumi:"tags"`
	UserDomainNames     []string                 `pulumi:"userDomainNames"`
	Versioning          *bool                    `pulumi:"versioning"`
	Website             *ObsBucketWebsite        `pulumi:"website"`
}

// The set of arguments for constructing a ObsBucket resource.
type ObsBucketArgs struct {
	Acl                 pulumi.StringPtrInput
	Bucket              pulumi.StringInput
	CorsRules           ObsBucketCorsRuleArrayInput
	Encryption          pulumi.BoolPtrInput
	EnterpriseProjectId pulumi.StringPtrInput
	ForceDestroy        pulumi.BoolPtrInput
	KmsKeyId            pulumi.StringPtrInput
	KmsKeyProjectId     pulumi.StringPtrInput
	LifecycleRules      ObsBucketLifecycleRuleArrayInput
	Loggings            ObsBucketLoggingArrayInput
	MultiAz             pulumi.BoolPtrInput
	ParallelFs          pulumi.BoolPtrInput
	Policy              pulumi.StringPtrInput
	PolicyFormat        pulumi.StringPtrInput
	Quota               pulumi.IntPtrInput
	Region              pulumi.StringPtrInput
	SseAlgorithm        pulumi.StringPtrInput
	StorageClass        pulumi.StringPtrInput
	Tags                pulumi.StringMapInput
	UserDomainNames     pulumi.StringArrayInput
	Versioning          pulumi.BoolPtrInput
	Website             ObsBucketWebsitePtrInput
}

func (ObsBucketArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*obsBucketArgs)(nil)).Elem()
}

type ObsBucketInput interface {
	pulumi.Input

	ToObsBucketOutput() ObsBucketOutput
	ToObsBucketOutputWithContext(ctx context.Context) ObsBucketOutput
}

func (*ObsBucket) ElementType() reflect.Type {
	return reflect.TypeOf((**ObsBucket)(nil)).Elem()
}

func (i *ObsBucket) ToObsBucketOutput() ObsBucketOutput {
	return i.ToObsBucketOutputWithContext(context.Background())
}

func (i *ObsBucket) ToObsBucketOutputWithContext(ctx context.Context) ObsBucketOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ObsBucketOutput)
}

// ObsBucketArrayInput is an input type that accepts ObsBucketArray and ObsBucketArrayOutput values.
// You can construct a concrete instance of `ObsBucketArrayInput` via:
//
//	ObsBucketArray{ ObsBucketArgs{...} }
type ObsBucketArrayInput interface {
	pulumi.Input

	ToObsBucketArrayOutput() ObsBucketArrayOutput
	ToObsBucketArrayOutputWithContext(context.Context) ObsBucketArrayOutput
}

type ObsBucketArray []ObsBucketInput

func (ObsBucketArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ObsBucket)(nil)).Elem()
}

func (i ObsBucketArray) ToObsBucketArrayOutput() ObsBucketArrayOutput {
	return i.ToObsBucketArrayOutputWithContext(context.Background())
}

func (i ObsBucketArray) ToObsBucketArrayOutputWithContext(ctx context.Context) ObsBucketArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ObsBucketArrayOutput)
}

// ObsBucketMapInput is an input type that accepts ObsBucketMap and ObsBucketMapOutput values.
// You can construct a concrete instance of `ObsBucketMapInput` via:
//
//	ObsBucketMap{ "key": ObsBucketArgs{...} }
type ObsBucketMapInput interface {
	pulumi.Input

	ToObsBucketMapOutput() ObsBucketMapOutput
	ToObsBucketMapOutputWithContext(context.Context) ObsBucketMapOutput
}

type ObsBucketMap map[string]ObsBucketInput

func (ObsBucketMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ObsBucket)(nil)).Elem()
}

func (i ObsBucketMap) ToObsBucketMapOutput() ObsBucketMapOutput {
	return i.ToObsBucketMapOutputWithContext(context.Background())
}

func (i ObsBucketMap) ToObsBucketMapOutputWithContext(ctx context.Context) ObsBucketMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ObsBucketMapOutput)
}

type ObsBucketOutput struct{ *pulumi.OutputState }

func (ObsBucketOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ObsBucket)(nil)).Elem()
}

func (o ObsBucketOutput) ToObsBucketOutput() ObsBucketOutput {
	return o
}

func (o ObsBucketOutput) ToObsBucketOutputWithContext(ctx context.Context) ObsBucketOutput {
	return o
}

func (o ObsBucketOutput) Acl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ObsBucket) pulumi.StringPtrOutput { return v.Acl }).(pulumi.StringPtrOutput)
}

func (o ObsBucketOutput) Bucket() pulumi.StringOutput {
	return o.ApplyT(func(v *ObsBucket) pulumi.StringOutput { return v.Bucket }).(pulumi.StringOutput)
}

func (o ObsBucketOutput) BucketDomainName() pulumi.StringOutput {
	return o.ApplyT(func(v *ObsBucket) pulumi.StringOutput { return v.BucketDomainName }).(pulumi.StringOutput)
}

func (o ObsBucketOutput) BucketVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *ObsBucket) pulumi.StringOutput { return v.BucketVersion }).(pulumi.StringOutput)
}

func (o ObsBucketOutput) CorsRules() ObsBucketCorsRuleArrayOutput {
	return o.ApplyT(func(v *ObsBucket) ObsBucketCorsRuleArrayOutput { return v.CorsRules }).(ObsBucketCorsRuleArrayOutput)
}

func (o ObsBucketOutput) Encryption() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ObsBucket) pulumi.BoolPtrOutput { return v.Encryption }).(pulumi.BoolPtrOutput)
}

func (o ObsBucketOutput) EnterpriseProjectId() pulumi.StringOutput {
	return o.ApplyT(func(v *ObsBucket) pulumi.StringOutput { return v.EnterpriseProjectId }).(pulumi.StringOutput)
}

func (o ObsBucketOutput) ForceDestroy() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ObsBucket) pulumi.BoolPtrOutput { return v.ForceDestroy }).(pulumi.BoolPtrOutput)
}

func (o ObsBucketOutput) KmsKeyId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ObsBucket) pulumi.StringPtrOutput { return v.KmsKeyId }).(pulumi.StringPtrOutput)
}

func (o ObsBucketOutput) KmsKeyProjectId() pulumi.StringOutput {
	return o.ApplyT(func(v *ObsBucket) pulumi.StringOutput { return v.KmsKeyProjectId }).(pulumi.StringOutput)
}

func (o ObsBucketOutput) LifecycleRules() ObsBucketLifecycleRuleArrayOutput {
	return o.ApplyT(func(v *ObsBucket) ObsBucketLifecycleRuleArrayOutput { return v.LifecycleRules }).(ObsBucketLifecycleRuleArrayOutput)
}

func (o ObsBucketOutput) Loggings() ObsBucketLoggingArrayOutput {
	return o.ApplyT(func(v *ObsBucket) ObsBucketLoggingArrayOutput { return v.Loggings }).(ObsBucketLoggingArrayOutput)
}

func (o ObsBucketOutput) MultiAz() pulumi.BoolOutput {
	return o.ApplyT(func(v *ObsBucket) pulumi.BoolOutput { return v.MultiAz }).(pulumi.BoolOutput)
}

func (o ObsBucketOutput) ParallelFs() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ObsBucket) pulumi.BoolPtrOutput { return v.ParallelFs }).(pulumi.BoolPtrOutput)
}

func (o ObsBucketOutput) Policy() pulumi.StringOutput {
	return o.ApplyT(func(v *ObsBucket) pulumi.StringOutput { return v.Policy }).(pulumi.StringOutput)
}

func (o ObsBucketOutput) PolicyFormat() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ObsBucket) pulumi.StringPtrOutput { return v.PolicyFormat }).(pulumi.StringPtrOutput)
}

func (o ObsBucketOutput) Quota() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ObsBucket) pulumi.IntPtrOutput { return v.Quota }).(pulumi.IntPtrOutput)
}

func (o ObsBucketOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v *ObsBucket) pulumi.StringOutput { return v.Region }).(pulumi.StringOutput)
}

func (o ObsBucketOutput) SseAlgorithm() pulumi.StringOutput {
	return o.ApplyT(func(v *ObsBucket) pulumi.StringOutput { return v.SseAlgorithm }).(pulumi.StringOutput)
}

func (o ObsBucketOutput) StorageClass() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ObsBucket) pulumi.StringPtrOutput { return v.StorageClass }).(pulumi.StringPtrOutput)
}

func (o ObsBucketOutput) StorageInfos() ObsBucketStorageInfoArrayOutput {
	return o.ApplyT(func(v *ObsBucket) ObsBucketStorageInfoArrayOutput { return v.StorageInfos }).(ObsBucketStorageInfoArrayOutput)
}

func (o ObsBucketOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *ObsBucket) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o ObsBucketOutput) UserDomainNames() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ObsBucket) pulumi.StringArrayOutput { return v.UserDomainNames }).(pulumi.StringArrayOutput)
}

func (o ObsBucketOutput) Versioning() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ObsBucket) pulumi.BoolPtrOutput { return v.Versioning }).(pulumi.BoolPtrOutput)
}

func (o ObsBucketOutput) Website() ObsBucketWebsitePtrOutput {
	return o.ApplyT(func(v *ObsBucket) ObsBucketWebsitePtrOutput { return v.Website }).(ObsBucketWebsitePtrOutput)
}

type ObsBucketArrayOutput struct{ *pulumi.OutputState }

func (ObsBucketArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ObsBucket)(nil)).Elem()
}

func (o ObsBucketArrayOutput) ToObsBucketArrayOutput() ObsBucketArrayOutput {
	return o
}

func (o ObsBucketArrayOutput) ToObsBucketArrayOutputWithContext(ctx context.Context) ObsBucketArrayOutput {
	return o
}

func (o ObsBucketArrayOutput) Index(i pulumi.IntInput) ObsBucketOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ObsBucket {
		return vs[0].([]*ObsBucket)[vs[1].(int)]
	}).(ObsBucketOutput)
}

type ObsBucketMapOutput struct{ *pulumi.OutputState }

func (ObsBucketMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ObsBucket)(nil)).Elem()
}

func (o ObsBucketMapOutput) ToObsBucketMapOutput() ObsBucketMapOutput {
	return o
}

func (o ObsBucketMapOutput) ToObsBucketMapOutputWithContext(ctx context.Context) ObsBucketMapOutput {
	return o
}

func (o ObsBucketMapOutput) MapIndex(k pulumi.StringInput) ObsBucketOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ObsBucket {
		return vs[0].(map[string]*ObsBucket)[vs[1].(string)]
	}).(ObsBucketOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ObsBucketInput)(nil)).Elem(), &ObsBucket{})
	pulumi.RegisterInputType(reflect.TypeOf((*ObsBucketArrayInput)(nil)).Elem(), ObsBucketArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ObsBucketMapInput)(nil)).Elem(), ObsBucketMap{})
	pulumi.RegisterOutputType(ObsBucketOutput{})
	pulumi.RegisterOutputType(ObsBucketArrayOutput{})
	pulumi.RegisterOutputType(ObsBucketMapOutput{})
}
