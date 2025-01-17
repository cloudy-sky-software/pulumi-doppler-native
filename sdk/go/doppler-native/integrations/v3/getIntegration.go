// Code generated by pulumigen DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v3

import (
	"context"
	"reflect"

	"github.com/cloudy-sky-software/pulumi-doppler-native/sdk/go/doppler-native/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupIntegration(ctx *pulumi.Context, args *LookupIntegrationArgs, opts ...pulumi.InvokeOption) (*LookupIntegrationResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupIntegrationResult
	err := ctx.Invoke("doppler-native:integrations/v3:getIntegration", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupIntegrationArgs struct {
}

type LookupIntegrationResult struct {
	Integration *GetIntegrationPropertiesIntegrationProperties `pulumi:"integration"`
}

func LookupIntegrationOutput(ctx *pulumi.Context, args LookupIntegrationOutputArgs, opts ...pulumi.InvokeOption) LookupIntegrationResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupIntegrationResultOutput, error) {
			args := v.(LookupIntegrationArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("doppler-native:integrations/v3:getIntegration", args, LookupIntegrationResultOutput{}, options).(LookupIntegrationResultOutput), nil
		}).(LookupIntegrationResultOutput)
}

type LookupIntegrationOutputArgs struct {
}

func (LookupIntegrationOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupIntegrationArgs)(nil)).Elem()
}

type LookupIntegrationResultOutput struct{ *pulumi.OutputState }

func (LookupIntegrationResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupIntegrationResult)(nil)).Elem()
}

func (o LookupIntegrationResultOutput) ToLookupIntegrationResultOutput() LookupIntegrationResultOutput {
	return o
}

func (o LookupIntegrationResultOutput) ToLookupIntegrationResultOutputWithContext(ctx context.Context) LookupIntegrationResultOutput {
	return o
}

func (o LookupIntegrationResultOutput) Integration() GetIntegrationPropertiesIntegrationPropertiesPtrOutput {
	return o.ApplyT(func(v LookupIntegrationResult) *GetIntegrationPropertiesIntegrationProperties { return v.Integration }).(GetIntegrationPropertiesIntegrationPropertiesPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupIntegrationResultOutput{})
}
