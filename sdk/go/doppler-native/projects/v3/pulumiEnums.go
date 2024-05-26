// Code generated by pulumigen DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v3

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Type string

const (
	TypeWorkplaceUser  = Type("workplace_user")
	TypeGroup          = Type("group")
	TypeInvite         = Type("invite")
	TypeServiceAccount = Type("service_account")
)

func (Type) ElementType() reflect.Type {
	return reflect.TypeOf((*Type)(nil)).Elem()
}

func (e Type) ToTypeOutput() TypeOutput {
	return pulumi.ToOutput(e).(TypeOutput)
}

func (e Type) ToTypeOutputWithContext(ctx context.Context) TypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(TypeOutput)
}

func (e Type) ToTypePtrOutput() TypePtrOutput {
	return e.ToTypePtrOutputWithContext(context.Background())
}

func (e Type) ToTypePtrOutputWithContext(ctx context.Context) TypePtrOutput {
	return Type(e).ToTypeOutputWithContext(ctx).ToTypePtrOutputWithContext(ctx)
}

func (e Type) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e Type) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e Type) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e Type) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type TypeOutput struct{ *pulumi.OutputState }

func (TypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Type)(nil)).Elem()
}

func (o TypeOutput) ToTypeOutput() TypeOutput {
	return o
}

func (o TypeOutput) ToTypeOutputWithContext(ctx context.Context) TypeOutput {
	return o
}

func (o TypeOutput) ToTypePtrOutput() TypePtrOutput {
	return o.ToTypePtrOutputWithContext(context.Background())
}

func (o TypeOutput) ToTypePtrOutputWithContext(ctx context.Context) TypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Type) *Type {
		return &v
	}).(TypePtrOutput)
}

func (o TypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o TypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e Type) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o TypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o TypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e Type) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type TypePtrOutput struct{ *pulumi.OutputState }

func (TypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Type)(nil)).Elem()
}

func (o TypePtrOutput) ToTypePtrOutput() TypePtrOutput {
	return o
}

func (o TypePtrOutput) ToTypePtrOutputWithContext(ctx context.Context) TypePtrOutput {
	return o
}

func (o TypePtrOutput) Elem() TypeOutput {
	return o.ApplyT(func(v *Type) Type {
		if v != nil {
			return *v
		}
		var ret Type
		return ret
	}).(TypeOutput)
}

func (o TypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o TypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *Type) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// TypeInput is an input type that accepts values of the Type enum
// A concrete instance of `TypeInput` can be one of the following:
//
//	TypeWorkplaceUser
//	TypeGroup
//	TypeInvite
//	TypeServiceAccount
type TypeInput interface {
	pulumi.Input

	ToTypeOutput() TypeOutput
	ToTypeOutputWithContext(context.Context) TypeOutput
}

var typePtrType = reflect.TypeOf((**Type)(nil)).Elem()

type TypePtrInput interface {
	pulumi.Input

	ToTypePtrOutput() TypePtrOutput
	ToTypePtrOutputWithContext(context.Context) TypePtrOutput
}

type typePtr string

func TypePtr(v string) TypePtrInput {
	return (*typePtr)(&v)
}

func (*typePtr) ElementType() reflect.Type {
	return typePtrType
}

func (in *typePtr) ToTypePtrOutput() TypePtrOutput {
	return pulumi.ToOutput(in).(TypePtrOutput)
}

func (in *typePtr) ToTypePtrOutputWithContext(ctx context.Context) TypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(TypePtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*TypeInput)(nil)).Elem(), Type("workplace_user"))
	pulumi.RegisterInputType(reflect.TypeOf((*TypePtrInput)(nil)).Elem(), Type("workplace_user"))
	pulumi.RegisterOutputType(TypeOutput{})
	pulumi.RegisterOutputType(TypePtrOutput{})
}
