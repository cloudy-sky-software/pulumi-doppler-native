// Code generated by pulumigen DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v3

import (
	"context"
	"reflect"

	"github.com/cloudy-sky-software/pulumi-doppler-native/sdk/go/doppler-native/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListConfigs(ctx *pulumi.Context, args *ListConfigsArgs, opts ...pulumi.InvokeOption) (*ListConfigsResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv ListConfigsResult
	err := ctx.Invoke("doppler-native:configs/v3:listConfigs", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type ListConfigsArgs struct {
}

type ListConfigsResult struct {
	Configs []ListConfigsPropertiesConfigsItemProperties `pulumi:"configs"`
	Page    *int                                         `pulumi:"page"`
}

// Defaults sets the appropriate defaults for ListConfigsResult
func (val *ListConfigsResult) Defaults() *ListConfigsResult {
	if val == nil {
		return nil
	}
	tmp := *val
	if tmp.Page == nil {
		page_ := 0
		tmp.Page = &page_
	}
	return &tmp
}
func ListConfigsOutput(ctx *pulumi.Context, args ListConfigsOutputArgs, opts ...pulumi.InvokeOption) ListConfigsResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (ListConfigsResultOutput, error) {
			args := v.(ListConfigsArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("doppler-native:configs/v3:listConfigs", args, ListConfigsResultOutput{}, options).(ListConfigsResultOutput), nil
		}).(ListConfigsResultOutput)
}

type ListConfigsOutputArgs struct {
}

func (ListConfigsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListConfigsArgs)(nil)).Elem()
}

type ListConfigsResultOutput struct{ *pulumi.OutputState }

func (ListConfigsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListConfigsResult)(nil)).Elem()
}

func (o ListConfigsResultOutput) ToListConfigsResultOutput() ListConfigsResultOutput {
	return o
}

func (o ListConfigsResultOutput) ToListConfigsResultOutputWithContext(ctx context.Context) ListConfigsResultOutput {
	return o
}

func (o ListConfigsResultOutput) Configs() ListConfigsPropertiesConfigsItemPropertiesArrayOutput {
	return o.ApplyT(func(v ListConfigsResult) []ListConfigsPropertiesConfigsItemProperties { return v.Configs }).(ListConfigsPropertiesConfigsItemPropertiesArrayOutput)
}

func (o ListConfigsResultOutput) Page() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ListConfigsResult) *int { return v.Page }).(pulumi.IntPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(ListConfigsResultOutput{})
}
