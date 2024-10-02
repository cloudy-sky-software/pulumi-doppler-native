// Code generated by pulumigen DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v3

import (
	"context"
	"reflect"

	"github.com/cloudy-sky-software/pulumi-doppler-native/sdk/go/doppler-native/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupEnvironment(ctx *pulumi.Context, args *LookupEnvironmentArgs, opts ...pulumi.InvokeOption) (*LookupEnvironmentResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupEnvironmentResult
	err := ctx.Invoke("doppler-native:environments/v3:getEnvironment", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupEnvironmentArgs struct {
}

type LookupEnvironmentResult struct {
	Environment *GetEnvironmentPropertiesEnvironmentProperties `pulumi:"environment"`
}

func LookupEnvironmentOutput(ctx *pulumi.Context, args LookupEnvironmentOutputArgs, opts ...pulumi.InvokeOption) LookupEnvironmentResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupEnvironmentResultOutput, error) {
			args := v.(LookupEnvironmentArgs)
			opts = internal.PkgInvokeDefaultOpts(opts)
			var rv LookupEnvironmentResult
			secret, err := ctx.InvokePackageRaw("doppler-native:environments/v3:getEnvironment", args, &rv, "", opts...)
			if err != nil {
				return LookupEnvironmentResultOutput{}, err
			}

			output := pulumi.ToOutput(rv).(LookupEnvironmentResultOutput)
			if secret {
				return pulumi.ToSecret(output).(LookupEnvironmentResultOutput), nil
			}
			return output, nil
		}).(LookupEnvironmentResultOutput)
}

type LookupEnvironmentOutputArgs struct {
}

func (LookupEnvironmentOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupEnvironmentArgs)(nil)).Elem()
}

type LookupEnvironmentResultOutput struct{ *pulumi.OutputState }

func (LookupEnvironmentResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupEnvironmentResult)(nil)).Elem()
}

func (o LookupEnvironmentResultOutput) ToLookupEnvironmentResultOutput() LookupEnvironmentResultOutput {
	return o
}

func (o LookupEnvironmentResultOutput) ToLookupEnvironmentResultOutputWithContext(ctx context.Context) LookupEnvironmentResultOutput {
	return o
}

func (o LookupEnvironmentResultOutput) Environment() GetEnvironmentPropertiesEnvironmentPropertiesPtrOutput {
	return o.ApplyT(func(v LookupEnvironmentResult) *GetEnvironmentPropertiesEnvironmentProperties { return v.Environment }).(GetEnvironmentPropertiesEnvironmentPropertiesPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupEnvironmentResultOutput{})
}
