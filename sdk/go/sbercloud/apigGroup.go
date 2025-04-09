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

type ApigGroup struct {
	pulumi.CustomResourceState

	// The creation time of the group, in RFC3339 format.
	CreatedAt pulumi.StringOutput `pulumi:"createdAt"`
	// The group description.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Specifies whether to use the debugging domain name to access the APIs within the group.
	DomainAccessEnabled pulumi.BoolOutput `pulumi:"domainAccessEnabled"`
	// The array of one or more environments of the associated group.
	Environments ApigGroupEnvironmentArrayOutput `pulumi:"environments"`
	// Whether to delete all sub-resources (for API) from this group.
	ForceDestroy pulumi.BoolOutput `pulumi:"forceDestroy"`
	// The ID of the dedicated instance to which the group belongs.
	InstanceId pulumi.StringOutput `pulumi:"instanceId"`
	// The group name.
	Name pulumi.StringOutput `pulumi:"name"`
	// The region where the dedicated instance is located.
	Region pulumi.StringOutput `pulumi:"region"`
	// The registration time.
	RegistrationTime pulumi.StringOutput `pulumi:"registrationTime"`
	// The latest update time of the group.
	UpdateTime pulumi.StringOutput `pulumi:"updateTime"`
	// The latest update time of the group, in RFC3339 format.
	UpdatedAt  pulumi.StringOutput           `pulumi:"updatedAt"`
	UrlDomains ApigGroupUrlDomainArrayOutput `pulumi:"urlDomains"`
}

// NewApigGroup registers a new resource with the given unique name, arguments, and options.
func NewApigGroup(ctx *pulumi.Context,
	name string, args *ApigGroupArgs, opts ...pulumi.ResourceOption) (*ApigGroup, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.InstanceId == nil {
		return nil, errors.New("invalid value for required argument 'InstanceId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource ApigGroup
	err := ctx.RegisterResource("sbercloud:index/apigGroup:ApigGroup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetApigGroup gets an existing ApigGroup resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetApigGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ApigGroupState, opts ...pulumi.ResourceOption) (*ApigGroup, error) {
	var resource ApigGroup
	err := ctx.ReadResource("sbercloud:index/apigGroup:ApigGroup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ApigGroup resources.
type apigGroupState struct {
	// The creation time of the group, in RFC3339 format.
	CreatedAt *string `pulumi:"createdAt"`
	// The group description.
	Description *string `pulumi:"description"`
	// Specifies whether to use the debugging domain name to access the APIs within the group.
	DomainAccessEnabled *bool `pulumi:"domainAccessEnabled"`
	// The array of one or more environments of the associated group.
	Environments []ApigGroupEnvironment `pulumi:"environments"`
	// Whether to delete all sub-resources (for API) from this group.
	ForceDestroy *bool `pulumi:"forceDestroy"`
	// The ID of the dedicated instance to which the group belongs.
	InstanceId *string `pulumi:"instanceId"`
	// The group name.
	Name *string `pulumi:"name"`
	// The region where the dedicated instance is located.
	Region *string `pulumi:"region"`
	// The registration time.
	RegistrationTime *string `pulumi:"registrationTime"`
	// The latest update time of the group.
	UpdateTime *string `pulumi:"updateTime"`
	// The latest update time of the group, in RFC3339 format.
	UpdatedAt  *string              `pulumi:"updatedAt"`
	UrlDomains []ApigGroupUrlDomain `pulumi:"urlDomains"`
}

type ApigGroupState struct {
	// The creation time of the group, in RFC3339 format.
	CreatedAt pulumi.StringPtrInput
	// The group description.
	Description pulumi.StringPtrInput
	// Specifies whether to use the debugging domain name to access the APIs within the group.
	DomainAccessEnabled pulumi.BoolPtrInput
	// The array of one or more environments of the associated group.
	Environments ApigGroupEnvironmentArrayInput
	// Whether to delete all sub-resources (for API) from this group.
	ForceDestroy pulumi.BoolPtrInput
	// The ID of the dedicated instance to which the group belongs.
	InstanceId pulumi.StringPtrInput
	// The group name.
	Name pulumi.StringPtrInput
	// The region where the dedicated instance is located.
	Region pulumi.StringPtrInput
	// The registration time.
	RegistrationTime pulumi.StringPtrInput
	// The latest update time of the group.
	UpdateTime pulumi.StringPtrInput
	// The latest update time of the group, in RFC3339 format.
	UpdatedAt  pulumi.StringPtrInput
	UrlDomains ApigGroupUrlDomainArrayInput
}

func (ApigGroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*apigGroupState)(nil)).Elem()
}

type apigGroupArgs struct {
	// The group description.
	Description *string `pulumi:"description"`
	// Specifies whether to use the debugging domain name to access the APIs within the group.
	DomainAccessEnabled *bool `pulumi:"domainAccessEnabled"`
	// The array of one or more environments of the associated group.
	Environments []ApigGroupEnvironment `pulumi:"environments"`
	// Whether to delete all sub-resources (for API) from this group.
	ForceDestroy *bool `pulumi:"forceDestroy"`
	// The ID of the dedicated instance to which the group belongs.
	InstanceId string `pulumi:"instanceId"`
	// The group name.
	Name *string `pulumi:"name"`
	// The region where the dedicated instance is located.
	Region     *string              `pulumi:"region"`
	UrlDomains []ApigGroupUrlDomain `pulumi:"urlDomains"`
}

// The set of arguments for constructing a ApigGroup resource.
type ApigGroupArgs struct {
	// The group description.
	Description pulumi.StringPtrInput
	// Specifies whether to use the debugging domain name to access the APIs within the group.
	DomainAccessEnabled pulumi.BoolPtrInput
	// The array of one or more environments of the associated group.
	Environments ApigGroupEnvironmentArrayInput
	// Whether to delete all sub-resources (for API) from this group.
	ForceDestroy pulumi.BoolPtrInput
	// The ID of the dedicated instance to which the group belongs.
	InstanceId pulumi.StringInput
	// The group name.
	Name pulumi.StringPtrInput
	// The region where the dedicated instance is located.
	Region     pulumi.StringPtrInput
	UrlDomains ApigGroupUrlDomainArrayInput
}

func (ApigGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*apigGroupArgs)(nil)).Elem()
}

type ApigGroupInput interface {
	pulumi.Input

	ToApigGroupOutput() ApigGroupOutput
	ToApigGroupOutputWithContext(ctx context.Context) ApigGroupOutput
}

func (*ApigGroup) ElementType() reflect.Type {
	return reflect.TypeOf((**ApigGroup)(nil)).Elem()
}

func (i *ApigGroup) ToApigGroupOutput() ApigGroupOutput {
	return i.ToApigGroupOutputWithContext(context.Background())
}

func (i *ApigGroup) ToApigGroupOutputWithContext(ctx context.Context) ApigGroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApigGroupOutput)
}

// ApigGroupArrayInput is an input type that accepts ApigGroupArray and ApigGroupArrayOutput values.
// You can construct a concrete instance of `ApigGroupArrayInput` via:
//
//	ApigGroupArray{ ApigGroupArgs{...} }
type ApigGroupArrayInput interface {
	pulumi.Input

	ToApigGroupArrayOutput() ApigGroupArrayOutput
	ToApigGroupArrayOutputWithContext(context.Context) ApigGroupArrayOutput
}

type ApigGroupArray []ApigGroupInput

func (ApigGroupArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ApigGroup)(nil)).Elem()
}

func (i ApigGroupArray) ToApigGroupArrayOutput() ApigGroupArrayOutput {
	return i.ToApigGroupArrayOutputWithContext(context.Background())
}

func (i ApigGroupArray) ToApigGroupArrayOutputWithContext(ctx context.Context) ApigGroupArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApigGroupArrayOutput)
}

// ApigGroupMapInput is an input type that accepts ApigGroupMap and ApigGroupMapOutput values.
// You can construct a concrete instance of `ApigGroupMapInput` via:
//
//	ApigGroupMap{ "key": ApigGroupArgs{...} }
type ApigGroupMapInput interface {
	pulumi.Input

	ToApigGroupMapOutput() ApigGroupMapOutput
	ToApigGroupMapOutputWithContext(context.Context) ApigGroupMapOutput
}

type ApigGroupMap map[string]ApigGroupInput

func (ApigGroupMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ApigGroup)(nil)).Elem()
}

func (i ApigGroupMap) ToApigGroupMapOutput() ApigGroupMapOutput {
	return i.ToApigGroupMapOutputWithContext(context.Background())
}

func (i ApigGroupMap) ToApigGroupMapOutputWithContext(ctx context.Context) ApigGroupMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApigGroupMapOutput)
}

type ApigGroupOutput struct{ *pulumi.OutputState }

func (ApigGroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ApigGroup)(nil)).Elem()
}

func (o ApigGroupOutput) ToApigGroupOutput() ApigGroupOutput {
	return o
}

func (o ApigGroupOutput) ToApigGroupOutputWithContext(ctx context.Context) ApigGroupOutput {
	return o
}

// The creation time of the group, in RFC3339 format.
func (o ApigGroupOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *ApigGroup) pulumi.StringOutput { return v.CreatedAt }).(pulumi.StringOutput)
}

// The group description.
func (o ApigGroupOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApigGroup) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// Specifies whether to use the debugging domain name to access the APIs within the group.
func (o ApigGroupOutput) DomainAccessEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v *ApigGroup) pulumi.BoolOutput { return v.DomainAccessEnabled }).(pulumi.BoolOutput)
}

// The array of one or more environments of the associated group.
func (o ApigGroupOutput) Environments() ApigGroupEnvironmentArrayOutput {
	return o.ApplyT(func(v *ApigGroup) ApigGroupEnvironmentArrayOutput { return v.Environments }).(ApigGroupEnvironmentArrayOutput)
}

// Whether to delete all sub-resources (for API) from this group.
func (o ApigGroupOutput) ForceDestroy() pulumi.BoolOutput {
	return o.ApplyT(func(v *ApigGroup) pulumi.BoolOutput { return v.ForceDestroy }).(pulumi.BoolOutput)
}

// The ID of the dedicated instance to which the group belongs.
func (o ApigGroupOutput) InstanceId() pulumi.StringOutput {
	return o.ApplyT(func(v *ApigGroup) pulumi.StringOutput { return v.InstanceId }).(pulumi.StringOutput)
}

// The group name.
func (o ApigGroupOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ApigGroup) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The region where the dedicated instance is located.
func (o ApigGroupOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v *ApigGroup) pulumi.StringOutput { return v.Region }).(pulumi.StringOutput)
}

// The registration time.
func (o ApigGroupOutput) RegistrationTime() pulumi.StringOutput {
	return o.ApplyT(func(v *ApigGroup) pulumi.StringOutput { return v.RegistrationTime }).(pulumi.StringOutput)
}

// The latest update time of the group.
func (o ApigGroupOutput) UpdateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *ApigGroup) pulumi.StringOutput { return v.UpdateTime }).(pulumi.StringOutput)
}

// The latest update time of the group, in RFC3339 format.
func (o ApigGroupOutput) UpdatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *ApigGroup) pulumi.StringOutput { return v.UpdatedAt }).(pulumi.StringOutput)
}

func (o ApigGroupOutput) UrlDomains() ApigGroupUrlDomainArrayOutput {
	return o.ApplyT(func(v *ApigGroup) ApigGroupUrlDomainArrayOutput { return v.UrlDomains }).(ApigGroupUrlDomainArrayOutput)
}

type ApigGroupArrayOutput struct{ *pulumi.OutputState }

func (ApigGroupArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ApigGroup)(nil)).Elem()
}

func (o ApigGroupArrayOutput) ToApigGroupArrayOutput() ApigGroupArrayOutput {
	return o
}

func (o ApigGroupArrayOutput) ToApigGroupArrayOutputWithContext(ctx context.Context) ApigGroupArrayOutput {
	return o
}

func (o ApigGroupArrayOutput) Index(i pulumi.IntInput) ApigGroupOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ApigGroup {
		return vs[0].([]*ApigGroup)[vs[1].(int)]
	}).(ApigGroupOutput)
}

type ApigGroupMapOutput struct{ *pulumi.OutputState }

func (ApigGroupMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ApigGroup)(nil)).Elem()
}

func (o ApigGroupMapOutput) ToApigGroupMapOutput() ApigGroupMapOutput {
	return o
}

func (o ApigGroupMapOutput) ToApigGroupMapOutputWithContext(ctx context.Context) ApigGroupMapOutput {
	return o
}

func (o ApigGroupMapOutput) MapIndex(k pulumi.StringInput) ApigGroupOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ApigGroup {
		return vs[0].(map[string]*ApigGroup)[vs[1].(string)]
	}).(ApigGroupOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ApigGroupInput)(nil)).Elem(), &ApigGroup{})
	pulumi.RegisterInputType(reflect.TypeOf((*ApigGroupArrayInput)(nil)).Elem(), ApigGroupArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ApigGroupMapInput)(nil)).Elem(), ApigGroupMap{})
	pulumi.RegisterOutputType(ApigGroupOutput{})
	pulumi.RegisterOutputType(ApigGroupArrayOutput{})
	pulumi.RegisterOutputType(ApigGroupMapOutput{})
}
