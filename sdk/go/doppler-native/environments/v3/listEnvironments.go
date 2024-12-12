// Code generated by pulumigen DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v3

import (
	"context"
	"reflect"

	"github.com/cloudy-sky-software/pulumi-doppler-native/sdk/go/doppler-native/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListEnvironments(ctx *pulumi.Context, args *ListEnvironmentsArgs, opts ...pulumi.InvokeOption) (*ListEnvironmentsResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv ListEnvironmentsResult
	err := ctx.Invoke("doppler-native:environments/v3:listEnvironments", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type ListEnvironmentsArgs struct {
}

type ListEnvironmentsResult struct {
	Environments []ListEnvironmentsPropertiesEnvironmentsItemProperties `pulumi:"environments"`
	Page         *int                                                   `pulumi:"page"`
}

// Defaults sets the appropriate defaults for ListEnvironmentsResult
func (val *ListEnvironmentsResult) Defaults() *ListEnvironmentsResult {
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
func ListEnvironmentsOutput(ctx *pulumi.Context, args ListEnvironmentsOutputArgs, opts ...pulumi.InvokeOption) ListEnvironmentsResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (ListEnvironmentsResultOutput, error) {
			args := v.(ListEnvironmentsArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("doppler-native:environments/v3:listEnvironments", args, ListEnvironmentsResultOutput{}, options).(ListEnvironmentsResultOutput), nil
		}).(ListEnvironmentsResultOutput)
}

type ListEnvironmentsOutputArgs struct {
}

func (ListEnvironmentsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListEnvironmentsArgs)(nil)).Elem()
}

type ListEnvironmentsResultOutput struct{ *pulumi.OutputState }

func (ListEnvironmentsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListEnvironmentsResult)(nil)).Elem()
}

func (o ListEnvironmentsResultOutput) ToListEnvironmentsResultOutput() ListEnvironmentsResultOutput {
	return o
}

func (o ListEnvironmentsResultOutput) ToListEnvironmentsResultOutputWithContext(ctx context.Context) ListEnvironmentsResultOutput {
	return o
}

func (o ListEnvironmentsResultOutput) Environments() ListEnvironmentsPropertiesEnvironmentsItemPropertiesArrayOutput {
	return o.ApplyT(func(v ListEnvironmentsResult) []ListEnvironmentsPropertiesEnvironmentsItemProperties {
		return v.Environments
	}).(ListEnvironmentsPropertiesEnvironmentsItemPropertiesArrayOutput)
}

func (o ListEnvironmentsResultOutput) Page() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ListEnvironmentsResult) *int { return v.Page }).(pulumi.IntPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(ListEnvironmentsResultOutput{})
}
