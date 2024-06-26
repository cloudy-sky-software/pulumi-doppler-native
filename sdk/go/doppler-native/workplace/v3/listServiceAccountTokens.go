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
	Items ListServiceAccountTokensProperties `pulumi:"items"`
}

// Defaults sets the appropriate defaults for ListServiceAccountTokensResult
func (val *ListServiceAccountTokensResult) Defaults() *ListServiceAccountTokensResult {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.Items = *tmp.Items.Defaults()

	return &tmp
}

func ListServiceAccountTokensOutput(ctx *pulumi.Context, args ListServiceAccountTokensOutputArgs, opts ...pulumi.InvokeOption) ListServiceAccountTokensResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListServiceAccountTokensResult, error) {
			args := v.(ListServiceAccountTokensArgs)
			r, err := ListServiceAccountTokens(ctx, &args, opts...)
			var s ListServiceAccountTokensResult
			if r != nil {
				s = *r
			}
			return s, err
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

func (o ListServiceAccountTokensResultOutput) Items() ListServiceAccountTokensPropertiesOutput {
	return o.ApplyT(func(v ListServiceAccountTokensResult) ListServiceAccountTokensProperties { return v.Items }).(ListServiceAccountTokensPropertiesOutput)
}

func init() {
	pulumi.RegisterOutputType(ListServiceAccountTokensResultOutput{})
}
