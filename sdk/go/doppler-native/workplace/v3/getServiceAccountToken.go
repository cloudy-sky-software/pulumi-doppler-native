// Code generated by pulumigen DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v3

import (
	"context"
	"reflect"

	"github.com/cloudy-sky-software/pulumi-doppler-native/sdk/go/doppler-native/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupServiceAccountToken(ctx *pulumi.Context, args *LookupServiceAccountTokenArgs, opts ...pulumi.InvokeOption) (*LookupServiceAccountTokenResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupServiceAccountTokenResult
	err := ctx.Invoke("doppler-native:workplace/v3:getServiceAccountToken", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupServiceAccountTokenArgs struct {
	// Slug of the API token
	ApiToken string `pulumi:"apiToken"`
	// Slug of the service account
	ServiceAccount string `pulumi:"serviceAccount"`
}

type LookupServiceAccountTokenResult struct {
	ApiToken *GetServiceAccountTokenPropertiesApiTokenProperties `pulumi:"apiToken"`
	Success  *bool                                               `pulumi:"success"`
}

// Defaults sets the appropriate defaults for LookupServiceAccountTokenResult
func (val *LookupServiceAccountTokenResult) Defaults() *LookupServiceAccountTokenResult {
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
func LookupServiceAccountTokenOutput(ctx *pulumi.Context, args LookupServiceAccountTokenOutputArgs, opts ...pulumi.InvokeOption) LookupServiceAccountTokenResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupServiceAccountTokenResultOutput, error) {
			args := v.(LookupServiceAccountTokenArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("doppler-native:workplace/v3:getServiceAccountToken", args, LookupServiceAccountTokenResultOutput{}, options).(LookupServiceAccountTokenResultOutput), nil
		}).(LookupServiceAccountTokenResultOutput)
}

type LookupServiceAccountTokenOutputArgs struct {
	// Slug of the API token
	ApiToken pulumi.StringInput `pulumi:"apiToken"`
	// Slug of the service account
	ServiceAccount pulumi.StringInput `pulumi:"serviceAccount"`
}

func (LookupServiceAccountTokenOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupServiceAccountTokenArgs)(nil)).Elem()
}

type LookupServiceAccountTokenResultOutput struct{ *pulumi.OutputState }

func (LookupServiceAccountTokenResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupServiceAccountTokenResult)(nil)).Elem()
}

func (o LookupServiceAccountTokenResultOutput) ToLookupServiceAccountTokenResultOutput() LookupServiceAccountTokenResultOutput {
	return o
}

func (o LookupServiceAccountTokenResultOutput) ToLookupServiceAccountTokenResultOutputWithContext(ctx context.Context) LookupServiceAccountTokenResultOutput {
	return o
}

func (o LookupServiceAccountTokenResultOutput) ApiToken() GetServiceAccountTokenPropertiesApiTokenPropertiesPtrOutput {
	return o.ApplyT(func(v LookupServiceAccountTokenResult) *GetServiceAccountTokenPropertiesApiTokenProperties {
		return v.ApiToken
	}).(GetServiceAccountTokenPropertiesApiTokenPropertiesPtrOutput)
}

func (o LookupServiceAccountTokenResultOutput) Success() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupServiceAccountTokenResult) *bool { return v.Success }).(pulumi.BoolPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupServiceAccountTokenResultOutput{})
}
