// Code generated by pulumigen DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v3

import (
	"context"
	"reflect"

	"github.com/cloudy-sky-software/pulumi-doppler-native/sdk/go/doppler-native/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ServiceAccounts struct {
	pulumi.CustomResourceState

	Name           pulumi.StringPtrOutput            `pulumi:"name"`
	ServiceAccount ServiceAccountPropertiesPtrOutput `pulumi:"serviceAccount"`
	// You may provide an identifier OR permissions, but not both
	WorkplaceRole WorkplaceRolePropertiesPtrOutput `pulumi:"workplaceRole"`
}

// NewServiceAccounts registers a new resource with the given unique name, arguments, and options.
func NewServiceAccounts(ctx *pulumi.Context,
	name string, args *ServiceAccountsArgs, opts ...pulumi.ResourceOption) (*ServiceAccounts, error) {
	if args == nil {
		args = &ServiceAccountsArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource ServiceAccounts
	err := ctx.RegisterResource("doppler-native:workplace/v3:ServiceAccounts", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetServiceAccounts gets an existing ServiceAccounts resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetServiceAccounts(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ServiceAccountsState, opts ...pulumi.ResourceOption) (*ServiceAccounts, error) {
	var resource ServiceAccounts
	err := ctx.ReadResource("doppler-native:workplace/v3:ServiceAccounts", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ServiceAccounts resources.
type serviceAccountsState struct {
}

type ServiceAccountsState struct {
}

func (ServiceAccountsState) ElementType() reflect.Type {
	return reflect.TypeOf((*serviceAccountsState)(nil)).Elem()
}

type serviceAccountsArgs struct {
	Name *string `pulumi:"name"`
	// You may provide an identifier OR permissions, but not both
	WorkplaceRole *WorkplaceRoleProperties `pulumi:"workplaceRole"`
}

// The set of arguments for constructing a ServiceAccounts resource.
type ServiceAccountsArgs struct {
	Name pulumi.StringPtrInput
	// You may provide an identifier OR permissions, but not both
	WorkplaceRole WorkplaceRolePropertiesPtrInput
}

func (ServiceAccountsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*serviceAccountsArgs)(nil)).Elem()
}

type ServiceAccountsInput interface {
	pulumi.Input

	ToServiceAccountsOutput() ServiceAccountsOutput
	ToServiceAccountsOutputWithContext(ctx context.Context) ServiceAccountsOutput
}

func (*ServiceAccounts) ElementType() reflect.Type {
	return reflect.TypeOf((**ServiceAccounts)(nil)).Elem()
}

func (i *ServiceAccounts) ToServiceAccountsOutput() ServiceAccountsOutput {
	return i.ToServiceAccountsOutputWithContext(context.Background())
}

func (i *ServiceAccounts) ToServiceAccountsOutputWithContext(ctx context.Context) ServiceAccountsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceAccountsOutput)
}

type ServiceAccountsOutput struct{ *pulumi.OutputState }

func (ServiceAccountsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ServiceAccounts)(nil)).Elem()
}

func (o ServiceAccountsOutput) ToServiceAccountsOutput() ServiceAccountsOutput {
	return o
}

func (o ServiceAccountsOutput) ToServiceAccountsOutputWithContext(ctx context.Context) ServiceAccountsOutput {
	return o
}

func (o ServiceAccountsOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServiceAccounts) pulumi.StringPtrOutput { return v.Name }).(pulumi.StringPtrOutput)
}

func (o ServiceAccountsOutput) ServiceAccount() ServiceAccountPropertiesPtrOutput {
	return o.ApplyT(func(v *ServiceAccounts) ServiceAccountPropertiesPtrOutput { return v.ServiceAccount }).(ServiceAccountPropertiesPtrOutput)
}

// You may provide an identifier OR permissions, but not both
func (o ServiceAccountsOutput) WorkplaceRole() WorkplaceRolePropertiesPtrOutput {
	return o.ApplyT(func(v *ServiceAccounts) WorkplaceRolePropertiesPtrOutput { return v.WorkplaceRole }).(WorkplaceRolePropertiesPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ServiceAccountsInput)(nil)).Elem(), &ServiceAccounts{})
	pulumi.RegisterOutputType(ServiceAccountsOutput{})
}
