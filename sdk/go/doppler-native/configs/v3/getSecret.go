// Code generated by pulumigen DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v3

import (
	"context"
	"reflect"

	"github.com/cloudy-sky-software/pulumi-doppler-native/sdk/go/doppler-native/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetSecret(ctx *pulumi.Context, args *GetSecretArgs, opts ...pulumi.InvokeOption) (*GetSecretResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetSecretResult
	err := ctx.Invoke("doppler-native:configs/v3:getSecret", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type GetSecretArgs struct {
}

type GetSecretResult struct {
	Items GetSecretProperties `pulumi:"items"`
}

func GetSecretOutput(ctx *pulumi.Context, args GetSecretOutputArgs, opts ...pulumi.InvokeOption) GetSecretResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetSecretResult, error) {
			args := v.(GetSecretArgs)
			r, err := GetSecret(ctx, &args, opts...)
			var s GetSecretResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetSecretResultOutput)
}

type GetSecretOutputArgs struct {
}

func (GetSecretOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetSecretArgs)(nil)).Elem()
}

type GetSecretResultOutput struct{ *pulumi.OutputState }

func (GetSecretResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetSecretResult)(nil)).Elem()
}

func (o GetSecretResultOutput) ToGetSecretResultOutput() GetSecretResultOutput {
	return o
}

func (o GetSecretResultOutput) ToGetSecretResultOutputWithContext(ctx context.Context) GetSecretResultOutput {
	return o
}

func (o GetSecretResultOutput) Items() GetSecretPropertiesOutput {
	return o.ApplyT(func(v GetSecretResult) GetSecretProperties { return v.Items }).(GetSecretPropertiesOutput)
}

func init() {
	pulumi.RegisterOutputType(GetSecretResultOutput{})
}
