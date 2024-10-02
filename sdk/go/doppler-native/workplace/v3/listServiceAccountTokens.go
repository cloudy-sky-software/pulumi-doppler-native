// Code generated by pulumigen DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v3

import (
	"context"
	"reflect"

	"github.com/cloudy-sky-software/pulumi-doppler-native/sdk/go/doppler-native/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListServiceAccountTokens(ctx *pulumi.Context, args *ListServiceAccountTokensArgs, opts ...pulumi.InvokeOption) (*ListServiceAccountTokensResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv ListServiceAccountTokensResult
	err := ctx.Invoke("doppler-native:workplace/v3:listServiceAccountTokens", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type ListServiceAccountTokensArgs struct {
	// Slug of the service account
	ServiceAccount string `pulumi:"serviceAccount"`
}

type ListServiceAccountTokensResult struct {
	ApiTokens []ListServiceAccountTokensPropertiesApiTokensItemProperties `pulumi:"apiTokens"`
	Success   *bool                                                       `pulumi:"success"`
}

// Defaults sets the appropriate defaults for ListServiceAccountTokensResult
func (val *ListServiceAccountTokensResult) Defaults() *ListServiceAccountTokensResult {
	if val == nil {
		return nil
	}
	tmp := *val
	if tmp.Success == nil {
		success_ := true
		tmp.Success = &success_
	}
	return &tmp
}

func ListServiceAccountTokensOutput(ctx *pulumi.Context, args ListServiceAccountTokensOutputArgs, opts ...pulumi.InvokeOption) ListServiceAccountTokensResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListServiceAccountTokensResultOutput, error) {
			args := v.(ListServiceAccountTokensArgs)
			opts = internal.PkgInvokeDefaultOpts(opts)
			var rv ListServiceAccountTokensResult
			secret, err := ctx.InvokePackageRaw("doppler-native:workplace/v3:listServiceAccountTokens", args, &rv, "", opts...)
			if err != nil {
				return ListServiceAccountTokensResultOutput{}, err
			}

			output := pulumi.ToOutput(rv).(ListServiceAccountTokensResultOutput)
			if secret {
				return pulumi.ToSecret(output).(ListServiceAccountTokensResultOutput), nil
			}
			return output, nil
		}).(ListServiceAccountTokensResultOutput)
}

type ListServiceAccountTokensOutputArgs struct {
	// Slug of the service account
	ServiceAccount pulumi.StringInput `pulumi:"serviceAccount"`
}

func (ListServiceAccountTokensOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListServiceAccountTokensArgs)(nil)).Elem()
}

type ListServiceAccountTokensResultOutput struct{ *pulumi.OutputState }

func (ListServiceAccountTokensResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListServiceAccountTokensResult)(nil)).Elem()
}

func (o ListServiceAccountTokensResultOutput) ToListServiceAccountTokensResultOutput() ListServiceAccountTokensResultOutput {
	return o
}

func (o ListServiceAccountTokensResultOutput) ToListServiceAccountTokensResultOutputWithContext(ctx context.Context) ListServiceAccountTokensResultOutput {
	return o
}

func (o ListServiceAccountTokensResultOutput) ApiTokens() ListServiceAccountTokensPropertiesApiTokensItemPropertiesArrayOutput {
	return o.ApplyT(func(v ListServiceAccountTokensResult) []ListServiceAccountTokensPropertiesApiTokensItemProperties {
		return v.ApiTokens
	}).(ListServiceAccountTokensPropertiesApiTokensItemPropertiesArrayOutput)
}

func (o ListServiceAccountTokensResultOutput) Success() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v ListServiceAccountTokensResult) *bool { return v.Success }).(pulumi.BoolPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(ListServiceAccountTokensResultOutput{})
}
