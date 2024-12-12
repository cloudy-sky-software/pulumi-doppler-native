// Code generated by pulumigen DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v3

import (
	"context"
	"reflect"

	"github.com/cloudy-sky-software/pulumi-doppler-native/sdk/go/doppler-native/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListConfigsTrustedIps(ctx *pulumi.Context, args *ListConfigsTrustedIpsArgs, opts ...pulumi.InvokeOption) (*ListConfigsTrustedIpsResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv ListConfigsTrustedIpsResult
	err := ctx.Invoke("doppler-native:configs/v3:listConfigsTrustedIps", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListConfigsTrustedIpsArgs struct {
}

type ListConfigsTrustedIpsResult struct {
	Ips []string `pulumi:"ips"`
}

func ListConfigsTrustedIpsOutput(ctx *pulumi.Context, args ListConfigsTrustedIpsOutputArgs, opts ...pulumi.InvokeOption) ListConfigsTrustedIpsResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (ListConfigsTrustedIpsResultOutput, error) {
			args := v.(ListConfigsTrustedIpsArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("doppler-native:configs/v3:listConfigsTrustedIps", args, ListConfigsTrustedIpsResultOutput{}, options).(ListConfigsTrustedIpsResultOutput), nil
		}).(ListConfigsTrustedIpsResultOutput)
}

type ListConfigsTrustedIpsOutputArgs struct {
}

func (ListConfigsTrustedIpsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListConfigsTrustedIpsArgs)(nil)).Elem()
}

type ListConfigsTrustedIpsResultOutput struct{ *pulumi.OutputState }

func (ListConfigsTrustedIpsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListConfigsTrustedIpsResult)(nil)).Elem()
}

func (o ListConfigsTrustedIpsResultOutput) ToListConfigsTrustedIpsResultOutput() ListConfigsTrustedIpsResultOutput {
	return o
}

func (o ListConfigsTrustedIpsResultOutput) ToListConfigsTrustedIpsResultOutputWithContext(ctx context.Context) ListConfigsTrustedIpsResultOutput {
	return o
}

func (o ListConfigsTrustedIpsResultOutput) Ips() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ListConfigsTrustedIpsResult) []string { return v.Ips }).(pulumi.StringArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(ListConfigsTrustedIpsResultOutput{})
}
