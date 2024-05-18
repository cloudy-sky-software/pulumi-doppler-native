// Code generated by pulumigen DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v3

import (
	"context"
	"reflect"

	"github.com/cloudy-sky-software/pulumi-doppler-native/sdk/go/doppler-native/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetGroup(ctx *pulumi.Context, args *GetGroupArgs, opts ...pulumi.InvokeOption) (*GetGroupResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetGroupResult
	err := ctx.Invoke("doppler-native:workplace/v3:getGroup", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type GetGroupArgs struct {
	// The group's slug
	Slug string `pulumi:"slug"`
}

type GetGroupResult struct {
	Items GetGroupProperties `pulumi:"items"`
}

func GetGroupOutput(ctx *pulumi.Context, args GetGroupOutputArgs, opts ...pulumi.InvokeOption) GetGroupResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetGroupResult, error) {
			args := v.(GetGroupArgs)
			r, err := GetGroup(ctx, &args, opts...)
			var s GetGroupResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetGroupResultOutput)
}

type GetGroupOutputArgs struct {
	// The group's slug
	Slug pulumi.StringInput `pulumi:"slug"`
}

func (GetGroupOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetGroupArgs)(nil)).Elem()
}

type GetGroupResultOutput struct{ *pulumi.OutputState }

func (GetGroupResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetGroupResult)(nil)).Elem()
}

func (o GetGroupResultOutput) ToGetGroupResultOutput() GetGroupResultOutput {
	return o
}

func (o GetGroupResultOutput) ToGetGroupResultOutputWithContext(ctx context.Context) GetGroupResultOutput {
	return o
}

func (o GetGroupResultOutput) Items() GetGroupPropertiesOutput {
	return o.ApplyT(func(v GetGroupResult) GetGroupProperties { return v.Items }).(GetGroupPropertiesOutput)
}

func init() {
	pulumi.RegisterOutputType(GetGroupResultOutput{})
}