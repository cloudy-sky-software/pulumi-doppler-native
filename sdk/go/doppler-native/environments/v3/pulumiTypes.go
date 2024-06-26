// Code generated by pulumigen DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v3

import (
	"context"
	"reflect"

	"github.com/cloudy-sky-software/pulumi-doppler-native/sdk/go/doppler-native/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

var _ = internal.GetEnvOrDefault

type EnvironmentProperties struct {
	CreatedAt      *string `pulumi:"createdAt"`
	Id             *string `pulumi:"id"`
	InitialFetchAt *string `pulumi:"initialFetchAt"`
	Name           *string `pulumi:"name"`
	Project        *string `pulumi:"project"`
}

type EnvironmentPropertiesOutput struct{ *pulumi.OutputState }

func (EnvironmentPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EnvironmentProperties)(nil)).Elem()
}

func (o EnvironmentPropertiesOutput) ToEnvironmentPropertiesOutput() EnvironmentPropertiesOutput {
	return o
}

func (o EnvironmentPropertiesOutput) ToEnvironmentPropertiesOutputWithContext(ctx context.Context) EnvironmentPropertiesOutput {
	return o
}

func (o EnvironmentPropertiesOutput) CreatedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EnvironmentProperties) *string { return v.CreatedAt }).(pulumi.StringPtrOutput)
}

func (o EnvironmentPropertiesOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EnvironmentProperties) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o EnvironmentPropertiesOutput) InitialFetchAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EnvironmentProperties) *string { return v.InitialFetchAt }).(pulumi.StringPtrOutput)
}

func (o EnvironmentPropertiesOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EnvironmentProperties) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o EnvironmentPropertiesOutput) Project() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EnvironmentProperties) *string { return v.Project }).(pulumi.StringPtrOutput)
}

type EnvironmentPropertiesPtrOutput struct{ *pulumi.OutputState }

func (EnvironmentPropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EnvironmentProperties)(nil)).Elem()
}

func (o EnvironmentPropertiesPtrOutput) ToEnvironmentPropertiesPtrOutput() EnvironmentPropertiesPtrOutput {
	return o
}

func (o EnvironmentPropertiesPtrOutput) ToEnvironmentPropertiesPtrOutputWithContext(ctx context.Context) EnvironmentPropertiesPtrOutput {
	return o
}

func (o EnvironmentPropertiesPtrOutput) Elem() EnvironmentPropertiesOutput {
	return o.ApplyT(func(v *EnvironmentProperties) EnvironmentProperties {
		if v != nil {
			return *v
		}
		var ret EnvironmentProperties
		return ret
	}).(EnvironmentPropertiesOutput)
}

func (o EnvironmentPropertiesPtrOutput) CreatedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EnvironmentProperties) *string {
		if v == nil {
			return nil
		}
		return v.CreatedAt
	}).(pulumi.StringPtrOutput)
}

func (o EnvironmentPropertiesPtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EnvironmentProperties) *string {
		if v == nil {
			return nil
		}
		return v.Id
	}).(pulumi.StringPtrOutput)
}

func (o EnvironmentPropertiesPtrOutput) InitialFetchAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EnvironmentProperties) *string {
		if v == nil {
			return nil
		}
		return v.InitialFetchAt
	}).(pulumi.StringPtrOutput)
}

func (o EnvironmentPropertiesPtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EnvironmentProperties) *string {
		if v == nil {
			return nil
		}
		return v.Name
	}).(pulumi.StringPtrOutput)
}

func (o EnvironmentPropertiesPtrOutput) Project() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EnvironmentProperties) *string {
		if v == nil {
			return nil
		}
		return v.Project
	}).(pulumi.StringPtrOutput)
}

type GetEnvironmentProperties struct {
	Environment *GetEnvironmentPropertiesEnvironmentProperties `pulumi:"environment"`
}

type GetEnvironmentPropertiesOutput struct{ *pulumi.OutputState }

func (GetEnvironmentPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetEnvironmentProperties)(nil)).Elem()
}

func (o GetEnvironmentPropertiesOutput) ToGetEnvironmentPropertiesOutput() GetEnvironmentPropertiesOutput {
	return o
}

func (o GetEnvironmentPropertiesOutput) ToGetEnvironmentPropertiesOutputWithContext(ctx context.Context) GetEnvironmentPropertiesOutput {
	return o
}

func (o GetEnvironmentPropertiesOutput) Environment() GetEnvironmentPropertiesEnvironmentPropertiesPtrOutput {
	return o.ApplyT(func(v GetEnvironmentProperties) *GetEnvironmentPropertiesEnvironmentProperties { return v.Environment }).(GetEnvironmentPropertiesEnvironmentPropertiesPtrOutput)
}

type GetEnvironmentPropertiesEnvironmentProperties struct {
	CreatedAt      *string `pulumi:"createdAt"`
	Id             *string `pulumi:"id"`
	InitialFetchAt *string `pulumi:"initialFetchAt"`
	Name           *string `pulumi:"name"`
	Project        *string `pulumi:"project"`
}

type GetEnvironmentPropertiesEnvironmentPropertiesOutput struct{ *pulumi.OutputState }

func (GetEnvironmentPropertiesEnvironmentPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetEnvironmentPropertiesEnvironmentProperties)(nil)).Elem()
}

func (o GetEnvironmentPropertiesEnvironmentPropertiesOutput) ToGetEnvironmentPropertiesEnvironmentPropertiesOutput() GetEnvironmentPropertiesEnvironmentPropertiesOutput {
	return o
}

func (o GetEnvironmentPropertiesEnvironmentPropertiesOutput) ToGetEnvironmentPropertiesEnvironmentPropertiesOutputWithContext(ctx context.Context) GetEnvironmentPropertiesEnvironmentPropertiesOutput {
	return o
}

func (o GetEnvironmentPropertiesEnvironmentPropertiesOutput) CreatedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetEnvironmentPropertiesEnvironmentProperties) *string { return v.CreatedAt }).(pulumi.StringPtrOutput)
}

func (o GetEnvironmentPropertiesEnvironmentPropertiesOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetEnvironmentPropertiesEnvironmentProperties) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o GetEnvironmentPropertiesEnvironmentPropertiesOutput) InitialFetchAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetEnvironmentPropertiesEnvironmentProperties) *string { return v.InitialFetchAt }).(pulumi.StringPtrOutput)
}

func (o GetEnvironmentPropertiesEnvironmentPropertiesOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetEnvironmentPropertiesEnvironmentProperties) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o GetEnvironmentPropertiesEnvironmentPropertiesOutput) Project() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetEnvironmentPropertiesEnvironmentProperties) *string { return v.Project }).(pulumi.StringPtrOutput)
}

type GetEnvironmentPropertiesEnvironmentPropertiesPtrOutput struct{ *pulumi.OutputState }

func (GetEnvironmentPropertiesEnvironmentPropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**GetEnvironmentPropertiesEnvironmentProperties)(nil)).Elem()
}

func (o GetEnvironmentPropertiesEnvironmentPropertiesPtrOutput) ToGetEnvironmentPropertiesEnvironmentPropertiesPtrOutput() GetEnvironmentPropertiesEnvironmentPropertiesPtrOutput {
	return o
}

func (o GetEnvironmentPropertiesEnvironmentPropertiesPtrOutput) ToGetEnvironmentPropertiesEnvironmentPropertiesPtrOutputWithContext(ctx context.Context) GetEnvironmentPropertiesEnvironmentPropertiesPtrOutput {
	return o
}

func (o GetEnvironmentPropertiesEnvironmentPropertiesPtrOutput) Elem() GetEnvironmentPropertiesEnvironmentPropertiesOutput {
	return o.ApplyT(func(v *GetEnvironmentPropertiesEnvironmentProperties) GetEnvironmentPropertiesEnvironmentProperties {
		if v != nil {
			return *v
		}
		var ret GetEnvironmentPropertiesEnvironmentProperties
		return ret
	}).(GetEnvironmentPropertiesEnvironmentPropertiesOutput)
}

func (o GetEnvironmentPropertiesEnvironmentPropertiesPtrOutput) CreatedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GetEnvironmentPropertiesEnvironmentProperties) *string {
		if v == nil {
			return nil
		}
		return v.CreatedAt
	}).(pulumi.StringPtrOutput)
}

func (o GetEnvironmentPropertiesEnvironmentPropertiesPtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GetEnvironmentPropertiesEnvironmentProperties) *string {
		if v == nil {
			return nil
		}
		return v.Id
	}).(pulumi.StringPtrOutput)
}

func (o GetEnvironmentPropertiesEnvironmentPropertiesPtrOutput) InitialFetchAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GetEnvironmentPropertiesEnvironmentProperties) *string {
		if v == nil {
			return nil
		}
		return v.InitialFetchAt
	}).(pulumi.StringPtrOutput)
}

func (o GetEnvironmentPropertiesEnvironmentPropertiesPtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GetEnvironmentPropertiesEnvironmentProperties) *string {
		if v == nil {
			return nil
		}
		return v.Name
	}).(pulumi.StringPtrOutput)
}

func (o GetEnvironmentPropertiesEnvironmentPropertiesPtrOutput) Project() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GetEnvironmentPropertiesEnvironmentProperties) *string {
		if v == nil {
			return nil
		}
		return v.Project
	}).(pulumi.StringPtrOutput)
}

type ListEnvironmentsProperties struct {
	Environments []ListEnvironmentsPropertiesEnvironmentsItemProperties `pulumi:"environments"`
	Page         *int                                                   `pulumi:"page"`
}

// Defaults sets the appropriate defaults for ListEnvironmentsProperties
func (val *ListEnvironmentsProperties) Defaults() *ListEnvironmentsProperties {
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

type ListEnvironmentsPropertiesOutput struct{ *pulumi.OutputState }

func (ListEnvironmentsPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListEnvironmentsProperties)(nil)).Elem()
}

func (o ListEnvironmentsPropertiesOutput) ToListEnvironmentsPropertiesOutput() ListEnvironmentsPropertiesOutput {
	return o
}

func (o ListEnvironmentsPropertiesOutput) ToListEnvironmentsPropertiesOutputWithContext(ctx context.Context) ListEnvironmentsPropertiesOutput {
	return o
}

func (o ListEnvironmentsPropertiesOutput) Environments() ListEnvironmentsPropertiesEnvironmentsItemPropertiesArrayOutput {
	return o.ApplyT(func(v ListEnvironmentsProperties) []ListEnvironmentsPropertiesEnvironmentsItemProperties {
		return v.Environments
	}).(ListEnvironmentsPropertiesEnvironmentsItemPropertiesArrayOutput)
}

func (o ListEnvironmentsPropertiesOutput) Page() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ListEnvironmentsProperties) *int { return v.Page }).(pulumi.IntPtrOutput)
}

type ListEnvironmentsPropertiesEnvironmentsItemProperties struct {
	CreatedAt      *string `pulumi:"createdAt"`
	Id             *string `pulumi:"id"`
	InitialFetchAt *string `pulumi:"initialFetchAt"`
	Name           *string `pulumi:"name"`
	Project        *string `pulumi:"project"`
}

type ListEnvironmentsPropertiesEnvironmentsItemPropertiesOutput struct{ *pulumi.OutputState }

func (ListEnvironmentsPropertiesEnvironmentsItemPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListEnvironmentsPropertiesEnvironmentsItemProperties)(nil)).Elem()
}

func (o ListEnvironmentsPropertiesEnvironmentsItemPropertiesOutput) ToListEnvironmentsPropertiesEnvironmentsItemPropertiesOutput() ListEnvironmentsPropertiesEnvironmentsItemPropertiesOutput {
	return o
}

func (o ListEnvironmentsPropertiesEnvironmentsItemPropertiesOutput) ToListEnvironmentsPropertiesEnvironmentsItemPropertiesOutputWithContext(ctx context.Context) ListEnvironmentsPropertiesEnvironmentsItemPropertiesOutput {
	return o
}

func (o ListEnvironmentsPropertiesEnvironmentsItemPropertiesOutput) CreatedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListEnvironmentsPropertiesEnvironmentsItemProperties) *string { return v.CreatedAt }).(pulumi.StringPtrOutput)
}

func (o ListEnvironmentsPropertiesEnvironmentsItemPropertiesOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListEnvironmentsPropertiesEnvironmentsItemProperties) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o ListEnvironmentsPropertiesEnvironmentsItemPropertiesOutput) InitialFetchAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListEnvironmentsPropertiesEnvironmentsItemProperties) *string { return v.InitialFetchAt }).(pulumi.StringPtrOutput)
}

func (o ListEnvironmentsPropertiesEnvironmentsItemPropertiesOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListEnvironmentsPropertiesEnvironmentsItemProperties) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o ListEnvironmentsPropertiesEnvironmentsItemPropertiesOutput) Project() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListEnvironmentsPropertiesEnvironmentsItemProperties) *string { return v.Project }).(pulumi.StringPtrOutput)
}

type ListEnvironmentsPropertiesEnvironmentsItemPropertiesArrayOutput struct{ *pulumi.OutputState }

func (ListEnvironmentsPropertiesEnvironmentsItemPropertiesArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ListEnvironmentsPropertiesEnvironmentsItemProperties)(nil)).Elem()
}

func (o ListEnvironmentsPropertiesEnvironmentsItemPropertiesArrayOutput) ToListEnvironmentsPropertiesEnvironmentsItemPropertiesArrayOutput() ListEnvironmentsPropertiesEnvironmentsItemPropertiesArrayOutput {
	return o
}

func (o ListEnvironmentsPropertiesEnvironmentsItemPropertiesArrayOutput) ToListEnvironmentsPropertiesEnvironmentsItemPropertiesArrayOutputWithContext(ctx context.Context) ListEnvironmentsPropertiesEnvironmentsItemPropertiesArrayOutput {
	return o
}

func (o ListEnvironmentsPropertiesEnvironmentsItemPropertiesArrayOutput) Index(i pulumi.IntInput) ListEnvironmentsPropertiesEnvironmentsItemPropertiesOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ListEnvironmentsPropertiesEnvironmentsItemProperties {
		return vs[0].([]ListEnvironmentsPropertiesEnvironmentsItemProperties)[vs[1].(int)]
	}).(ListEnvironmentsPropertiesEnvironmentsItemPropertiesOutput)
}

func init() {
	pulumi.RegisterOutputType(EnvironmentPropertiesOutput{})
	pulumi.RegisterOutputType(EnvironmentPropertiesPtrOutput{})
	pulumi.RegisterOutputType(GetEnvironmentPropertiesOutput{})
	pulumi.RegisterOutputType(GetEnvironmentPropertiesEnvironmentPropertiesOutput{})
	pulumi.RegisterOutputType(GetEnvironmentPropertiesEnvironmentPropertiesPtrOutput{})
	pulumi.RegisterOutputType(ListEnvironmentsPropertiesOutput{})
	pulumi.RegisterOutputType(ListEnvironmentsPropertiesEnvironmentsItemPropertiesOutput{})
	pulumi.RegisterOutputType(ListEnvironmentsPropertiesEnvironmentsItemPropertiesArrayOutput{})
}
