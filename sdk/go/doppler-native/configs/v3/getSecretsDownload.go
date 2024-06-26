// Code generated by pulumigen DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v3

import (
	"context"
	"reflect"

	"github.com/cloudy-sky-software/pulumi-doppler-native/sdk/go/doppler-native/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetSecretsDownload(ctx *pulumi.Context, args *GetSecretsDownloadArgs, opts ...pulumi.InvokeOption) (*GetSecretsDownloadResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetSecretsDownloadResult
	err := ctx.Invoke("doppler-native:configs/v3:getSecretsDownload", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type GetSecretsDownloadArgs struct {
}

type GetSecretsDownloadResult struct {
	Items GetSecretsDownloadProperties `pulumi:"items"`
}

func GetSecretsDownloadOutput(ctx *pulumi.Context, args GetSecretsDownloadOutputArgs, opts ...pulumi.InvokeOption) GetSecretsDownloadResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetSecretsDownloadResult, error) {
			args := v.(GetSecretsDownloadArgs)
			r, err := GetSecretsDownload(ctx, &args, opts...)
			var s GetSecretsDownloadResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetSecretsDownloadResultOutput)
}

type GetSecretsDownloadOutputArgs struct {
}

func (GetSecretsDownloadOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetSecretsDownloadArgs)(nil)).Elem()
}

type GetSecretsDownloadResultOutput struct{ *pulumi.OutputState }

func (GetSecretsDownloadResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetSecretsDownloadResult)(nil)).Elem()
}

func (o GetSecretsDownloadResultOutput) ToGetSecretsDownloadResultOutput() GetSecretsDownloadResultOutput {
	return o
}

func (o GetSecretsDownloadResultOutput) ToGetSecretsDownloadResultOutputWithContext(ctx context.Context) GetSecretsDownloadResultOutput {
	return o
}

func (o GetSecretsDownloadResultOutput) Items() GetSecretsDownloadPropertiesOutput {
	return o.ApplyT(func(v GetSecretsDownloadResult) GetSecretsDownloadProperties { return v.Items }).(GetSecretsDownloadPropertiesOutput)
}

func init() {
	pulumi.RegisterOutputType(GetSecretsDownloadResultOutput{})
}
