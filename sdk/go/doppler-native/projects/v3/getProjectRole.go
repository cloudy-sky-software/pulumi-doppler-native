// Code generated by pulumigen DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v3

import (
	"context"
	"reflect"

	"github.com/cloudy-sky-software/pulumi-doppler-native/sdk/go/doppler-native/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupProjectRole(ctx *pulumi.Context, args *LookupProjectRoleArgs, opts ...pulumi.InvokeOption) (*LookupProjectRoleResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupProjectRoleResult
	err := ctx.Invoke("doppler-native:projects/v3:getProjectRole", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupProjectRoleArgs struct {
	// The role's unique identifier
	Role string `pulumi:"role"`
}

type LookupProjectRoleResult struct {
	Role *GetProjectRolePropertiesRoleProperties `pulumi:"role"`
}

// Defaults sets the appropriate defaults for LookupProjectRoleResult
func (val *LookupProjectRoleResult) Defaults() *LookupProjectRoleResult {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.Role = tmp.Role.Defaults()

	return &tmp
}

func LookupProjectRoleOutput(ctx *pulumi.Context, args LookupProjectRoleOutputArgs, opts ...pulumi.InvokeOption) LookupProjectRoleResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupProjectRoleResultOutput, error) {
			args := v.(LookupProjectRoleArgs)
			opts = internal.PkgInvokeDefaultOpts(opts)
			var rv LookupProjectRoleResult
			secret, err := ctx.InvokePackageRaw("doppler-native:projects/v3:getProjectRole", args, &rv, "", opts...)
			if err != nil {
				return LookupProjectRoleResultOutput{}, err
			}

			output := pulumi.ToOutput(rv).(LookupProjectRoleResultOutput)
			if secret {
				return pulumi.ToSecret(output).(LookupProjectRoleResultOutput), nil
			}
			return output, nil
		}).(LookupProjectRoleResultOutput)
}

type LookupProjectRoleOutputArgs struct {
	// The role's unique identifier
	Role pulumi.StringInput `pulumi:"role"`
}

func (LookupProjectRoleOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupProjectRoleArgs)(nil)).Elem()
}

type LookupProjectRoleResultOutput struct{ *pulumi.OutputState }

func (LookupProjectRoleResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupProjectRoleResult)(nil)).Elem()
}

func (o LookupProjectRoleResultOutput) ToLookupProjectRoleResultOutput() LookupProjectRoleResultOutput {
	return o
}

func (o LookupProjectRoleResultOutput) ToLookupProjectRoleResultOutputWithContext(ctx context.Context) LookupProjectRoleResultOutput {
	return o
}

func (o LookupProjectRoleResultOutput) Role() GetProjectRolePropertiesRolePropertiesPtrOutput {
	return o.ApplyT(func(v LookupProjectRoleResult) *GetProjectRolePropertiesRoleProperties { return v.Role }).(GetProjectRolePropertiesRolePropertiesPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupProjectRoleResultOutput{})
}
