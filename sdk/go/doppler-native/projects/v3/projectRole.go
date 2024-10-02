// Code generated by pulumigen DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v3

import (
	"context"
	"reflect"

	"errors"
	"github.com/cloudy-sky-software/pulumi-doppler-native/sdk/go/doppler-native/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ProjectRole struct {
	pulumi.CustomResourceState

	// The name of the role
	Name pulumi.StringOutput `pulumi:"name"`
	// An array containing the permissions the role has. Valid permissions are: `enclave_config_logs`, `enclave_project_config_secrets_read`, `enclave_project_config_dynamic_secrets_read`, `enclave_project_config_dynamic_secrets_leases_write`, `enclave_project_config_rotated_secrets_read`, `enclave_config_syncs_manage`, `enclave_project_secrets_notes_manage`, `enclave_project_config_create`, `enclave_project_config_duplicate`, `enclave_project_config_secrets_write`, `enclave_project_config_service_tokens`, `enclave_project_config_trusted_ips`, `enclave_project_config_logs_rollback`, `enclave_project_config_list_all`, `enclave_project_pull_request_create`, `enclave_project_pull_request_respond`, `enclave_project_pull_request_view`, `enclave_secret_reminders`, `enclave_config_access_logs`, `enclave_project_members`, `enclave_project_rename`, `enclave_project_delete`, `enclave_project_webhooks`, `enclave_project_config_dynamic_secrets_manage`, `enclave_project_config_rotated_secrets_manage`, `enclave_project_config_rename`, `enclave_project_config_lock`, `enclave_project_config_delete`, `enclave_project_environment_list_all`, `enclave_project_environment_all`, `enclave_project_environment_order`, `enclave_project_environment_create`, `enclave_project_environment_delete`, `enclave_project_environment_rename`, `enclave_project_environment_settings_manage`, `enclave_project_secrets_referencing`, `enclave_config_secrets_referencing`
	Permissions pulumi.StringArrayOutput `pulumi:"permissions"`
	Role        RolePropertiesPtrOutput  `pulumi:"role"`
}

// NewProjectRole registers a new resource with the given unique name, arguments, and options.
func NewProjectRole(ctx *pulumi.Context,
	name string, args *ProjectRoleArgs, opts ...pulumi.ResourceOption) (*ProjectRole, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Permissions == nil {
		return nil, errors.New("invalid value for required argument 'Permissions'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource ProjectRole
	err := ctx.RegisterResource("doppler-native:projects/v3:ProjectRole", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetProjectRole gets an existing ProjectRole resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetProjectRole(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ProjectRoleState, opts ...pulumi.ResourceOption) (*ProjectRole, error) {
	var resource ProjectRole
	err := ctx.ReadResource("doppler-native:projects/v3:ProjectRole", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ProjectRole resources.
type projectRoleState struct {
}

type ProjectRoleState struct {
}

func (ProjectRoleState) ElementType() reflect.Type {
	return reflect.TypeOf((*projectRoleState)(nil)).Elem()
}

type projectRoleArgs struct {
	// The name of the role
	Name *string `pulumi:"name"`
	// An array containing the permissions the role has. Valid permissions are: `enclave_config_logs`, `enclave_project_config_secrets_read`, `enclave_project_config_dynamic_secrets_read`, `enclave_project_config_dynamic_secrets_leases_write`, `enclave_project_config_rotated_secrets_read`, `enclave_config_syncs_manage`, `enclave_project_secrets_notes_manage`, `enclave_project_config_create`, `enclave_project_config_duplicate`, `enclave_project_config_secrets_write`, `enclave_project_config_service_tokens`, `enclave_project_config_trusted_ips`, `enclave_project_config_logs_rollback`, `enclave_project_config_list_all`, `enclave_project_pull_request_create`, `enclave_project_pull_request_respond`, `enclave_project_pull_request_view`, `enclave_secret_reminders`, `enclave_config_access_logs`, `enclave_project_members`, `enclave_project_rename`, `enclave_project_delete`, `enclave_project_webhooks`, `enclave_project_config_dynamic_secrets_manage`, `enclave_project_config_rotated_secrets_manage`, `enclave_project_config_rename`, `enclave_project_config_lock`, `enclave_project_config_delete`, `enclave_project_environment_list_all`, `enclave_project_environment_all`, `enclave_project_environment_order`, `enclave_project_environment_create`, `enclave_project_environment_delete`, `enclave_project_environment_rename`, `enclave_project_environment_settings_manage`, `enclave_project_secrets_referencing`, `enclave_config_secrets_referencing`
	Permissions []string `pulumi:"permissions"`
}

// The set of arguments for constructing a ProjectRole resource.
type ProjectRoleArgs struct {
	// The name of the role
	Name pulumi.StringPtrInput
	// An array containing the permissions the role has. Valid permissions are: `enclave_config_logs`, `enclave_project_config_secrets_read`, `enclave_project_config_dynamic_secrets_read`, `enclave_project_config_dynamic_secrets_leases_write`, `enclave_project_config_rotated_secrets_read`, `enclave_config_syncs_manage`, `enclave_project_secrets_notes_manage`, `enclave_project_config_create`, `enclave_project_config_duplicate`, `enclave_project_config_secrets_write`, `enclave_project_config_service_tokens`, `enclave_project_config_trusted_ips`, `enclave_project_config_logs_rollback`, `enclave_project_config_list_all`, `enclave_project_pull_request_create`, `enclave_project_pull_request_respond`, `enclave_project_pull_request_view`, `enclave_secret_reminders`, `enclave_config_access_logs`, `enclave_project_members`, `enclave_project_rename`, `enclave_project_delete`, `enclave_project_webhooks`, `enclave_project_config_dynamic_secrets_manage`, `enclave_project_config_rotated_secrets_manage`, `enclave_project_config_rename`, `enclave_project_config_lock`, `enclave_project_config_delete`, `enclave_project_environment_list_all`, `enclave_project_environment_all`, `enclave_project_environment_order`, `enclave_project_environment_create`, `enclave_project_environment_delete`, `enclave_project_environment_rename`, `enclave_project_environment_settings_manage`, `enclave_project_secrets_referencing`, `enclave_config_secrets_referencing`
	Permissions pulumi.StringArrayInput
}

func (ProjectRoleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*projectRoleArgs)(nil)).Elem()
}

type ProjectRoleInput interface {
	pulumi.Input

	ToProjectRoleOutput() ProjectRoleOutput
	ToProjectRoleOutputWithContext(ctx context.Context) ProjectRoleOutput
}

func (*ProjectRole) ElementType() reflect.Type {
	return reflect.TypeOf((**ProjectRole)(nil)).Elem()
}

func (i *ProjectRole) ToProjectRoleOutput() ProjectRoleOutput {
	return i.ToProjectRoleOutputWithContext(context.Background())
}

func (i *ProjectRole) ToProjectRoleOutputWithContext(ctx context.Context) ProjectRoleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProjectRoleOutput)
}

type ProjectRoleOutput struct{ *pulumi.OutputState }

func (ProjectRoleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ProjectRole)(nil)).Elem()
}

func (o ProjectRoleOutput) ToProjectRoleOutput() ProjectRoleOutput {
	return o
}

func (o ProjectRoleOutput) ToProjectRoleOutputWithContext(ctx context.Context) ProjectRoleOutput {
	return o
}

// The name of the role
func (o ProjectRoleOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ProjectRole) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// An array containing the permissions the role has. Valid permissions are: `enclave_config_logs`, `enclave_project_config_secrets_read`, `enclave_project_config_dynamic_secrets_read`, `enclave_project_config_dynamic_secrets_leases_write`, `enclave_project_config_rotated_secrets_read`, `enclave_config_syncs_manage`, `enclave_project_secrets_notes_manage`, `enclave_project_config_create`, `enclave_project_config_duplicate`, `enclave_project_config_secrets_write`, `enclave_project_config_service_tokens`, `enclave_project_config_trusted_ips`, `enclave_project_config_logs_rollback`, `enclave_project_config_list_all`, `enclave_project_pull_request_create`, `enclave_project_pull_request_respond`, `enclave_project_pull_request_view`, `enclave_secret_reminders`, `enclave_config_access_logs`, `enclave_project_members`, `enclave_project_rename`, `enclave_project_delete`, `enclave_project_webhooks`, `enclave_project_config_dynamic_secrets_manage`, `enclave_project_config_rotated_secrets_manage`, `enclave_project_config_rename`, `enclave_project_config_lock`, `enclave_project_config_delete`, `enclave_project_environment_list_all`, `enclave_project_environment_all`, `enclave_project_environment_order`, `enclave_project_environment_create`, `enclave_project_environment_delete`, `enclave_project_environment_rename`, `enclave_project_environment_settings_manage`, `enclave_project_secrets_referencing`, `enclave_config_secrets_referencing`
func (o ProjectRoleOutput) Permissions() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ProjectRole) pulumi.StringArrayOutput { return v.Permissions }).(pulumi.StringArrayOutput)
}

func (o ProjectRoleOutput) Role() RolePropertiesPtrOutput {
	return o.ApplyT(func(v *ProjectRole) RolePropertiesPtrOutput { return v.Role }).(RolePropertiesPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ProjectRoleInput)(nil)).Elem(), &ProjectRole{})
	pulumi.RegisterOutputType(ProjectRoleOutput{})
}