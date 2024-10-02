// Code generated by pulumigen DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v3

import (
	"context"
	"reflect"

	"github.com/cloudy-sky-software/pulumi-doppler-native/sdk/go/doppler-native/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetConfigLog(ctx *pulumi.Context, args *GetConfigLogArgs, opts ...pulumi.InvokeOption) (*GetConfigLogResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetConfigLogResult
	err := ctx.Invoke("doppler-native:configs/v3:getConfigLog", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type GetConfigLogArgs struct {
}

type GetConfigLogResult struct {
	Log *GetConfigLogPropertiesLogProperties `pulumi:"log"`
}

// Defaults sets the appropriate defaults for GetConfigLogResult
func (val *GetConfigLogResult) Defaults() *GetConfigLogResult {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.Log = tmp.Log.Defaults()

	return &tmp
}

func GetConfigLogOutput(ctx *pulumi.Context, args GetConfigLogOutputArgs, opts ...pulumi.InvokeOption) GetConfigLogResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetConfigLogResultOutput, error) {
			args := v.(GetConfigLogArgs)
			opts = internal.PkgInvokeDefaultOpts(opts)
			var rv GetConfigLogResult
			secret, err := ctx.InvokePackageRaw("doppler-native:configs/v3:getConfigLog", args, &rv, "", opts...)
			if err != nil {
				return GetConfigLogResultOutput{}, err
			}

			output := pulumi.ToOutput(rv).(GetConfigLogResultOutput)
			if secret {
				return pulumi.ToSecret(output).(GetConfigLogResultOutput), nil
			}
			return output, nil
		}).(GetConfigLogResultOutput)
}

type GetConfigLogOutputArgs struct {
}

func (GetConfigLogOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetConfigLogArgs)(nil)).Elem()
}

type GetConfigLogResultOutput struct{ *pulumi.OutputState }

func (GetConfigLogResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetConfigLogResult)(nil)).Elem()
}

func (o GetConfigLogResultOutput) ToGetConfigLogResultOutput() GetConfigLogResultOutput {
	return o
}

func (o GetConfigLogResultOutput) ToGetConfigLogResultOutputWithContext(ctx context.Context) GetConfigLogResultOutput {
	return o
}

func (o GetConfigLogResultOutput) Log() GetConfigLogPropertiesLogPropertiesPtrOutput {
	return o.ApplyT(func(v GetConfigLogResult) *GetConfigLogPropertiesLogProperties { return v.Log }).(GetConfigLogPropertiesLogPropertiesPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetConfigLogResultOutput{})
}
