// Code generated by pulumigen DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v3

import (
	"context"
	"reflect"

	"github.com/cloudy-sky-software/pulumi-doppler-native/sdk/go/doppler-native/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetEnvironment(ctx *pulumi.Context, args *GetEnvironmentArgs, opts ...pulumi.InvokeOption) (*GetEnvironmentResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetEnvironmentResult
	err := ctx.Invoke("doppler-native:environments/v3:getEnvironment", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type GetEnvironmentArgs struct {
}

type GetEnvironmentResult struct {
	Items GetEnvironmentProperties `pulumi:"items"`
}

func GetEnvironmentOutput(ctx *pulumi.Context, args GetEnvironmentOutputArgs, opts ...pulumi.InvokeOption) GetEnvironmentResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetEnvironmentResult, error) {
			args := v.(GetEnvironmentArgs)
			r, err := GetEnvironment(ctx, &args, opts...)
			var s GetEnvironmentResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetEnvironmentResultOutput)
}

type GetEnvironmentOutputArgs struct {
}

func (GetEnvironmentOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetEnvironmentArgs)(nil)).Elem()
}

type GetEnvironmentResultOutput struct{ *pulumi.OutputState }

func (GetEnvironmentResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetEnvironmentResult)(nil)).Elem()
}

func (o GetEnvironmentResultOutput) ToGetEnvironmentResultOutput() GetEnvironmentResultOutput {
	return o
}

func (o GetEnvironmentResultOutput) ToGetEnvironmentResultOutputWithContext(ctx context.Context) GetEnvironmentResultOutput {
	return o
}

func (o GetEnvironmentResultOutput) Items() GetEnvironmentPropertiesOutput {
	return o.ApplyT(func(v GetEnvironmentResult) GetEnvironmentProperties { return v.Items }).(GetEnvironmentPropertiesOutput)
}

func init() {
	pulumi.RegisterOutputType(GetEnvironmentResultOutput{})
}
