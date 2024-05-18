// Code generated by pulumigen DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v3

import (
	"context"
	"reflect"

	"github.com/cloudy-sky-software/pulumi-doppler-native/sdk/go/doppler-native/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListWorkplaceRoles(ctx *pulumi.Context, args *ListWorkplaceRolesArgs, opts ...pulumi.InvokeOption) (*ListWorkplaceRolesResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv ListWorkplaceRolesResult
	err := ctx.Invoke("doppler-native:workplace/v3:listWorkplaceRoles", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListWorkplaceRolesArgs struct {
}

type ListWorkplaceRolesResult struct {
	Items ListWorkplaceRolesProperties `pulumi:"items"`
}

func ListWorkplaceRolesOutput(ctx *pulumi.Context, args ListWorkplaceRolesOutputArgs, opts ...pulumi.InvokeOption) ListWorkplaceRolesResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListWorkplaceRolesResult, error) {
			args := v.(ListWorkplaceRolesArgs)
			r, err := ListWorkplaceRoles(ctx, &args, opts...)
			var s ListWorkplaceRolesResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListWorkplaceRolesResultOutput)
}

type ListWorkplaceRolesOutputArgs struct {
}

func (ListWorkplaceRolesOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListWorkplaceRolesArgs)(nil)).Elem()
}

type ListWorkplaceRolesResultOutput struct{ *pulumi.OutputState }

func (ListWorkplaceRolesResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListWorkplaceRolesResult)(nil)).Elem()
}

func (o ListWorkplaceRolesResultOutput) ToListWorkplaceRolesResultOutput() ListWorkplaceRolesResultOutput {
	return o
}

func (o ListWorkplaceRolesResultOutput) ToListWorkplaceRolesResultOutputWithContext(ctx context.Context) ListWorkplaceRolesResultOutput {
	return o
}

func (o ListWorkplaceRolesResultOutput) Items() ListWorkplaceRolesPropertiesOutput {
	return o.ApplyT(func(v ListWorkplaceRolesResult) ListWorkplaceRolesProperties { return v.Items }).(ListWorkplaceRolesPropertiesOutput)
}

func init() {
	pulumi.RegisterOutputType(ListWorkplaceRolesResultOutput{})
}