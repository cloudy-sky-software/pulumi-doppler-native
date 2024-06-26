// Code generated by pulumigen DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v3

import (
	"context"
	"reflect"

	"github.com/cloudy-sky-software/pulumi-doppler-native/sdk/go/doppler-native/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetWebhook(ctx *pulumi.Context, args *GetWebhookArgs, opts ...pulumi.InvokeOption) (*GetWebhookResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetWebhookResult
	err := ctx.Invoke("doppler-native:webhooks/v3:getWebhook", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type GetWebhookArgs struct {
	// Webhook's slug
	Slug string `pulumi:"slug"`
}

type GetWebhookResult struct {
	Items interface{} `pulumi:"items"`
}

func GetWebhookOutput(ctx *pulumi.Context, args GetWebhookOutputArgs, opts ...pulumi.InvokeOption) GetWebhookResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetWebhookResult, error) {
			args := v.(GetWebhookArgs)
			r, err := GetWebhook(ctx, &args, opts...)
			var s GetWebhookResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetWebhookResultOutput)
}

type GetWebhookOutputArgs struct {
	// Webhook's slug
	Slug pulumi.StringInput `pulumi:"slug"`
}

func (GetWebhookOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetWebhookArgs)(nil)).Elem()
}

type GetWebhookResultOutput struct{ *pulumi.OutputState }

func (GetWebhookResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetWebhookResult)(nil)).Elem()
}

func (o GetWebhookResultOutput) ToGetWebhookResultOutput() GetWebhookResultOutput {
	return o
}

func (o GetWebhookResultOutput) ToGetWebhookResultOutputWithContext(ctx context.Context) GetWebhookResultOutput {
	return o
}

func (o GetWebhookResultOutput) Items() pulumi.AnyOutput {
	return o.ApplyT(func(v GetWebhookResult) interface{} { return v.Items }).(pulumi.AnyOutput)
}

func init() {
	pulumi.RegisterOutputType(GetWebhookResultOutput{})
}
