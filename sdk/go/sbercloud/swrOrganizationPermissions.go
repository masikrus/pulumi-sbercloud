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

type SwrOrganizationPermissions struct {
	pulumi.CustomResourceState

	Creator         pulumi.StringOutput                                 `pulumi:"creator"`
	Organization    pulumi.StringOutput                                 `pulumi:"organization"`
	Region          pulumi.StringOutput                                 `pulumi:"region"`
	SelfPermissions SwrOrganizationPermissionsSelfPermissionArrayOutput `pulumi:"selfPermissions"`
	Users           SwrOrganizationPermissionsUserArrayOutput           `pulumi:"users"`
}

// NewSwrOrganizationPermissions registers a new resource with the given unique name, arguments, and options.
func NewSwrOrganizationPermissions(ctx *pulumi.Context,
	name string, args *SwrOrganizationPermissionsArgs, opts ...pulumi.ResourceOption) (*SwrOrganizationPermissions, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Organization == nil {
		return nil, errors.New("invalid value for required argument 'Organization'")
	}
	if args.Users == nil {
		return nil, errors.New("invalid value for required argument 'Users'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource SwrOrganizationPermissions
	err := ctx.RegisterResource("sbercloud:index/swrOrganizationPermissions:SwrOrganizationPermissions", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSwrOrganizationPermissions gets an existing SwrOrganizationPermissions resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSwrOrganizationPermissions(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SwrOrganizationPermissionsState, opts ...pulumi.ResourceOption) (*SwrOrganizationPermissions, error) {
	var resource SwrOrganizationPermissions
	err := ctx.ReadResource("sbercloud:index/swrOrganizationPermissions:SwrOrganizationPermissions", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SwrOrganizationPermissions resources.
type swrOrganizationPermissionsState struct {
	Creator         *string                                    `pulumi:"creator"`
	Organization    *string                                    `pulumi:"organization"`
	Region          *string                                    `pulumi:"region"`
	SelfPermissions []SwrOrganizationPermissionsSelfPermission `pulumi:"selfPermissions"`
	Users           []SwrOrganizationPermissionsUser           `pulumi:"users"`
}

type SwrOrganizationPermissionsState struct {
	Creator         pulumi.StringPtrInput
	Organization    pulumi.StringPtrInput
	Region          pulumi.StringPtrInput
	SelfPermissions SwrOrganizationPermissionsSelfPermissionArrayInput
	Users           SwrOrganizationPermissionsUserArrayInput
}

func (SwrOrganizationPermissionsState) ElementType() reflect.Type {
	return reflect.TypeOf((*swrOrganizationPermissionsState)(nil)).Elem()
}

type swrOrganizationPermissionsArgs struct {
	Organization string                           `pulumi:"organization"`
	Region       *string                          `pulumi:"region"`
	Users        []SwrOrganizationPermissionsUser `pulumi:"users"`
}

// The set of arguments for constructing a SwrOrganizationPermissions resource.
type SwrOrganizationPermissionsArgs struct {
	Organization pulumi.StringInput
	Region       pulumi.StringPtrInput
	Users        SwrOrganizationPermissionsUserArrayInput
}

func (SwrOrganizationPermissionsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*swrOrganizationPermissionsArgs)(nil)).Elem()
}

type SwrOrganizationPermissionsInput interface {
	pulumi.Input

	ToSwrOrganizationPermissionsOutput() SwrOrganizationPermissionsOutput
	ToSwrOrganizationPermissionsOutputWithContext(ctx context.Context) SwrOrganizationPermissionsOutput
}

func (*SwrOrganizationPermissions) ElementType() reflect.Type {
	return reflect.TypeOf((**SwrOrganizationPermissions)(nil)).Elem()
}

func (i *SwrOrganizationPermissions) ToSwrOrganizationPermissionsOutput() SwrOrganizationPermissionsOutput {
	return i.ToSwrOrganizationPermissionsOutputWithContext(context.Background())
}

func (i *SwrOrganizationPermissions) ToSwrOrganizationPermissionsOutputWithContext(ctx context.Context) SwrOrganizationPermissionsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SwrOrganizationPermissionsOutput)
}

// SwrOrganizationPermissionsArrayInput is an input type that accepts SwrOrganizationPermissionsArray and SwrOrganizationPermissionsArrayOutput values.
// You can construct a concrete instance of `SwrOrganizationPermissionsArrayInput` via:
//
//	SwrOrganizationPermissionsArray{ SwrOrganizationPermissionsArgs{...} }
type SwrOrganizationPermissionsArrayInput interface {
	pulumi.Input

	ToSwrOrganizationPermissionsArrayOutput() SwrOrganizationPermissionsArrayOutput
	ToSwrOrganizationPermissionsArrayOutputWithContext(context.Context) SwrOrganizationPermissionsArrayOutput
}

type SwrOrganizationPermissionsArray []SwrOrganizationPermissionsInput

func (SwrOrganizationPermissionsArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SwrOrganizationPermissions)(nil)).Elem()
}

func (i SwrOrganizationPermissionsArray) ToSwrOrganizationPermissionsArrayOutput() SwrOrganizationPermissionsArrayOutput {
	return i.ToSwrOrganizationPermissionsArrayOutputWithContext(context.Background())
}

func (i SwrOrganizationPermissionsArray) ToSwrOrganizationPermissionsArrayOutputWithContext(ctx context.Context) SwrOrganizationPermissionsArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SwrOrganizationPermissionsArrayOutput)
}

// SwrOrganizationPermissionsMapInput is an input type that accepts SwrOrganizationPermissionsMap and SwrOrganizationPermissionsMapOutput values.
// You can construct a concrete instance of `SwrOrganizationPermissionsMapInput` via:
//
//	SwrOrganizationPermissionsMap{ "key": SwrOrganizationPermissionsArgs{...} }
type SwrOrganizationPermissionsMapInput interface {
	pulumi.Input

	ToSwrOrganizationPermissionsMapOutput() SwrOrganizationPermissionsMapOutput
	ToSwrOrganizationPermissionsMapOutputWithContext(context.Context) SwrOrganizationPermissionsMapOutput
}

type SwrOrganizationPermissionsMap map[string]SwrOrganizationPermissionsInput

func (SwrOrganizationPermissionsMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SwrOrganizationPermissions)(nil)).Elem()
}

func (i SwrOrganizationPermissionsMap) ToSwrOrganizationPermissionsMapOutput() SwrOrganizationPermissionsMapOutput {
	return i.ToSwrOrganizationPermissionsMapOutputWithContext(context.Background())
}

func (i SwrOrganizationPermissionsMap) ToSwrOrganizationPermissionsMapOutputWithContext(ctx context.Context) SwrOrganizationPermissionsMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SwrOrganizationPermissionsMapOutput)
}

type SwrOrganizationPermissionsOutput struct{ *pulumi.OutputState }

func (SwrOrganizationPermissionsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SwrOrganizationPermissions)(nil)).Elem()
}

func (o SwrOrganizationPermissionsOutput) ToSwrOrganizationPermissionsOutput() SwrOrganizationPermissionsOutput {
	return o
}

func (o SwrOrganizationPermissionsOutput) ToSwrOrganizationPermissionsOutputWithContext(ctx context.Context) SwrOrganizationPermissionsOutput {
	return o
}

func (o SwrOrganizationPermissionsOutput) Creator() pulumi.StringOutput {
	return o.ApplyT(func(v *SwrOrganizationPermissions) pulumi.StringOutput { return v.Creator }).(pulumi.StringOutput)
}

func (o SwrOrganizationPermissionsOutput) Organization() pulumi.StringOutput {
	return o.ApplyT(func(v *SwrOrganizationPermissions) pulumi.StringOutput { return v.Organization }).(pulumi.StringOutput)
}

func (o SwrOrganizationPermissionsOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v *SwrOrganizationPermissions) pulumi.StringOutput { return v.Region }).(pulumi.StringOutput)
}

func (o SwrOrganizationPermissionsOutput) SelfPermissions() SwrOrganizationPermissionsSelfPermissionArrayOutput {
	return o.ApplyT(func(v *SwrOrganizationPermissions) SwrOrganizationPermissionsSelfPermissionArrayOutput {
		return v.SelfPermissions
	}).(SwrOrganizationPermissionsSelfPermissionArrayOutput)
}

func (o SwrOrganizationPermissionsOutput) Users() SwrOrganizationPermissionsUserArrayOutput {
	return o.ApplyT(func(v *SwrOrganizationPermissions) SwrOrganizationPermissionsUserArrayOutput { return v.Users }).(SwrOrganizationPermissionsUserArrayOutput)
}

type SwrOrganizationPermissionsArrayOutput struct{ *pulumi.OutputState }

func (SwrOrganizationPermissionsArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SwrOrganizationPermissions)(nil)).Elem()
}

func (o SwrOrganizationPermissionsArrayOutput) ToSwrOrganizationPermissionsArrayOutput() SwrOrganizationPermissionsArrayOutput {
	return o
}

func (o SwrOrganizationPermissionsArrayOutput) ToSwrOrganizationPermissionsArrayOutputWithContext(ctx context.Context) SwrOrganizationPermissionsArrayOutput {
	return o
}

func (o SwrOrganizationPermissionsArrayOutput) Index(i pulumi.IntInput) SwrOrganizationPermissionsOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SwrOrganizationPermissions {
		return vs[0].([]*SwrOrganizationPermissions)[vs[1].(int)]
	}).(SwrOrganizationPermissionsOutput)
}

type SwrOrganizationPermissionsMapOutput struct{ *pulumi.OutputState }

func (SwrOrganizationPermissionsMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SwrOrganizationPermissions)(nil)).Elem()
}

func (o SwrOrganizationPermissionsMapOutput) ToSwrOrganizationPermissionsMapOutput() SwrOrganizationPermissionsMapOutput {
	return o
}

func (o SwrOrganizationPermissionsMapOutput) ToSwrOrganizationPermissionsMapOutputWithContext(ctx context.Context) SwrOrganizationPermissionsMapOutput {
	return o
}

func (o SwrOrganizationPermissionsMapOutput) MapIndex(k pulumi.StringInput) SwrOrganizationPermissionsOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SwrOrganizationPermissions {
		return vs[0].(map[string]*SwrOrganizationPermissions)[vs[1].(string)]
	}).(SwrOrganizationPermissionsOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SwrOrganizationPermissionsInput)(nil)).Elem(), &SwrOrganizationPermissions{})
	pulumi.RegisterInputType(reflect.TypeOf((*SwrOrganizationPermissionsArrayInput)(nil)).Elem(), SwrOrganizationPermissionsArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SwrOrganizationPermissionsMapInput)(nil)).Elem(), SwrOrganizationPermissionsMap{})
	pulumi.RegisterOutputType(SwrOrganizationPermissionsOutput{})
	pulumi.RegisterOutputType(SwrOrganizationPermissionsArrayOutput{})
	pulumi.RegisterOutputType(SwrOrganizationPermissionsMapOutput{})
}
