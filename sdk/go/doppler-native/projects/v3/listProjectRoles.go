// Code generated by pulumigen DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v3

import (
	"context"
	"reflect"

	"github.com/cloudy-sky-software/pulumi-doppler-native/sdk/go/doppler-native/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListProjectRoles(ctx *pulumi.Context, args *ListProjectRolesArgs, opts ...pulumi.InvokeOption) (*ListProjectRolesResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv ListProjectRolesResult
	err := ctx.Invoke("doppler-native:projects/v3:listProjectRoles", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListProjectRolesArgs struct {
}

type ListProjectRolesResult struct {
	Roles []ListProjectRolesPropertiesRolesItemProperties `pulumi:"roles"`
}

func ListProjectRolesOutput(ctx *pulumi.Context, args ListProjectRolesOutputArgs, opts ...pulumi.InvokeOption) ListProjectRolesResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListProjectRolesResultOutput, error) {
			args := v.(ListProjectRolesArgs)
			opts = internal.PkgInvokeDefaultOpts(opts)
			var rv ListProjectRolesResult
			secret, err := ctx.InvokePackageRaw("doppler-native:projects/v3:listProjectRoles", args, &rv, "", opts...)
			if err != nil {
				return ListProjectRolesResultOutput{}, err
			}

			output := pulumi.ToOutput(rv).(ListProjectRolesResultOutput)
			if secret {
				return pulumi.ToSecret(output).(ListProjectRolesResultOutput), nil
			}
			return output, nil
		}).(ListProjectRolesResultOutput)
}

type ListProjectRolesOutputArgs struct {
}

func (ListProjectRolesOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListProjectRolesArgs)(nil)).Elem()
}

type ListProjectRolesResultOutput struct{ *pulumi.OutputState }

func (ListProjectRolesResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListProjectRolesResult)(nil)).Elem()
}

func (o ListProjectRolesResultOutput) ToListProjectRolesResultOutput() ListProjectRolesResultOutput {
	return o
}

func (o ListProjectRolesResultOutput) ToListProjectRolesResultOutputWithContext(ctx context.Context) ListProjectRolesResultOutput {
	return o
}

func (o ListProjectRolesResultOutput) Roles() ListProjectRolesPropertiesRolesItemPropertiesArrayOutput {
	return o.ApplyT(func(v ListProjectRolesResult) []ListProjectRolesPropertiesRolesItemProperties { return v.Roles }).(ListProjectRolesPropertiesRolesItemPropertiesArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(ListProjectRolesResultOutput{})
}
