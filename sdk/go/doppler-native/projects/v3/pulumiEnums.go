// Code generated by pulumigen DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v3

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ProjectMembersType string

const (
	ProjectMembersTypeWorkplaceUser  = ProjectMembersType("workplace_user")
	ProjectMembersTypeGroup          = ProjectMembersType("group")
	ProjectMembersTypeInvite         = ProjectMembersType("invite")
	ProjectMembersTypeServiceAccount = ProjectMembersType("service_account")
)

func (ProjectMembersType) ElementType() reflect.Type {
	return reflect.TypeOf((*ProjectMembersType)(nil)).Elem()
}

func (e ProjectMembersType) ToProjectMembersTypeOutput() ProjectMembersTypeOutput {
	return pulumi.ToOutput(e).(ProjectMembersTypeOutput)
}

func (e ProjectMembersType) ToProjectMembersTypeOutputWithContext(ctx context.Context) ProjectMembersTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(ProjectMembersTypeOutput)
}

func (e ProjectMembersType) ToProjectMembersTypePtrOutput() ProjectMembersTypePtrOutput {
	return e.ToProjectMembersTypePtrOutputWithContext(context.Background())
}

func (e ProjectMembersType) ToProjectMembersTypePtrOutputWithContext(ctx context.Context) ProjectMembersTypePtrOutput {
	return ProjectMembersType(e).ToProjectMembersTypeOutputWithContext(ctx).ToProjectMembersTypePtrOutputWithContext(ctx)
}

func (e ProjectMembersType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e ProjectMembersType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e ProjectMembersType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e ProjectMembersType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type ProjectMembersTypeOutput struct{ *pulumi.OutputState }

func (ProjectMembersTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ProjectMembersType)(nil)).Elem()
}

func (o ProjectMembersTypeOutput) ToProjectMembersTypeOutput() ProjectMembersTypeOutput {
	return o
}

func (o ProjectMembersTypeOutput) ToProjectMembersTypeOutputWithContext(ctx context.Context) ProjectMembersTypeOutput {
	return o
}

func (o ProjectMembersTypeOutput) ToProjectMembersTypePtrOutput() ProjectMembersTypePtrOutput {
	return o.ToProjectMembersTypePtrOutputWithContext(context.Background())
}

func (o ProjectMembersTypeOutput) ToProjectMembersTypePtrOutputWithContext(ctx context.Context) ProjectMembersTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ProjectMembersType) *ProjectMembersType {
		return &v
	}).(ProjectMembersTypePtrOutput)
}

func (o ProjectMembersTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o ProjectMembersTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ProjectMembersType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o ProjectMembersTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ProjectMembersTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ProjectMembersType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type ProjectMembersTypePtrOutput struct{ *pulumi.OutputState }

func (ProjectMembersTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ProjectMembersType)(nil)).Elem()
}

func (o ProjectMembersTypePtrOutput) ToProjectMembersTypePtrOutput() ProjectMembersTypePtrOutput {
	return o
}

func (o ProjectMembersTypePtrOutput) ToProjectMembersTypePtrOutputWithContext(ctx context.Context) ProjectMembersTypePtrOutput {
	return o
}

func (o ProjectMembersTypePtrOutput) Elem() ProjectMembersTypeOutput {
	return o.ApplyT(func(v *ProjectMembersType) ProjectMembersType {
		if v != nil {
			return *v
		}
		var ret ProjectMembersType
		return ret
	}).(ProjectMembersTypeOutput)
}

func (o ProjectMembersTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ProjectMembersTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *ProjectMembersType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// ProjectMembersTypeInput is an input type that accepts values of the ProjectMembersType enum
// A concrete instance of `ProjectMembersTypeInput` can be one of the following:
//
//	ProjectMembersTypeWorkplaceUser
//	ProjectMembersTypeGroup
//	ProjectMembersTypeInvite
//	ProjectMembersTypeServiceAccount
type ProjectMembersTypeInput interface {
	pulumi.Input

	ToProjectMembersTypeOutput() ProjectMembersTypeOutput
	ToProjectMembersTypeOutputWithContext(context.Context) ProjectMembersTypeOutput
}

var projectMembersTypePtrType = reflect.TypeOf((**ProjectMembersType)(nil)).Elem()

type ProjectMembersTypePtrInput interface {
	pulumi.Input

	ToProjectMembersTypePtrOutput() ProjectMembersTypePtrOutput
	ToProjectMembersTypePtrOutputWithContext(context.Context) ProjectMembersTypePtrOutput
}

type projectMembersTypePtr string

func ProjectMembersTypePtr(v string) ProjectMembersTypePtrInput {
	return (*projectMembersTypePtr)(&v)
}

func (*projectMembersTypePtr) ElementType() reflect.Type {
	return projectMembersTypePtrType
}

func (in *projectMembersTypePtr) ToProjectMembersTypePtrOutput() ProjectMembersTypePtrOutput {
	return pulumi.ToOutput(in).(ProjectMembersTypePtrOutput)
}

func (in *projectMembersTypePtr) ToProjectMembersTypePtrOutputWithContext(ctx context.Context) ProjectMembersTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(ProjectMembersTypePtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ProjectMembersTypeInput)(nil)).Elem(), ProjectMembersType("workplace_user"))
	pulumi.RegisterInputType(reflect.TypeOf((*ProjectMembersTypePtrInput)(nil)).Elem(), ProjectMembersType("workplace_user"))
	pulumi.RegisterOutputType(ProjectMembersTypeOutput{})
	pulumi.RegisterOutputType(ProjectMembersTypePtrOutput{})
}