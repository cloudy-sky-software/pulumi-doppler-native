# coding=utf-8
# *** WARNING: this file was generated by pulumigen. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import sys
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
if sys.version_info >= (3, 11):
    from typing import NotRequired, TypedDict, TypeAlias
else:
    from typing_extensions import NotRequired, TypedDict, TypeAlias
from ... import _utilities
from . import outputs

__all__ = ['ProjectRoleArgs', 'ProjectRole']

@pulumi.input_type
class ProjectRoleArgs:
    def __init__(__self__, *,
                 permissions: pulumi.Input[Sequence[pulumi.Input[str]]],
                 name: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a ProjectRole resource.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] permissions: An array containing the permissions the role has. Valid permissions are: `enclave_config_logs`, `enclave_project_config_secrets_read`, `enclave_project_config_dynamic_secrets_read`, `enclave_project_config_dynamic_secrets_leases_write`, `enclave_project_config_rotated_secrets_read`, `enclave_config_syncs_manage`, `enclave_project_secrets_notes_manage`, `enclave_project_config_create`, `enclave_project_config_duplicate`, `enclave_project_config_secrets_write`, `enclave_project_config_service_tokens`, `enclave_project_config_trusted_ips`, `enclave_project_config_logs_rollback`, `enclave_project_config_list_all`, `enclave_project_pull_request_create`, `enclave_project_pull_request_respond`, `enclave_project_pull_request_view`, `enclave_secret_reminders`, `enclave_config_access_logs`, `enclave_project_members`, `enclave_project_rename`, `enclave_project_delete`, `enclave_project_webhooks`, `enclave_project_config_dynamic_secrets_manage`, `enclave_project_config_rotated_secrets_manage`, `enclave_project_config_rename`, `enclave_project_config_lock`, `enclave_project_config_delete`, `enclave_project_environment_list_all`, `enclave_project_environment_all`, `enclave_project_environment_order`, `enclave_project_environment_create`, `enclave_project_environment_delete`, `enclave_project_environment_rename`, `enclave_project_environment_settings_manage`, `enclave_project_secrets_referencing`, `enclave_config_secrets_referencing`
        :param pulumi.Input[str] name: The name of the role
        """
        pulumi.set(__self__, "permissions", permissions)
        if name is not None:
            pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter
    def permissions(self) -> pulumi.Input[Sequence[pulumi.Input[str]]]:
        """
        An array containing the permissions the role has. Valid permissions are: `enclave_config_logs`, `enclave_project_config_secrets_read`, `enclave_project_config_dynamic_secrets_read`, `enclave_project_config_dynamic_secrets_leases_write`, `enclave_project_config_rotated_secrets_read`, `enclave_config_syncs_manage`, `enclave_project_secrets_notes_manage`, `enclave_project_config_create`, `enclave_project_config_duplicate`, `enclave_project_config_secrets_write`, `enclave_project_config_service_tokens`, `enclave_project_config_trusted_ips`, `enclave_project_config_logs_rollback`, `enclave_project_config_list_all`, `enclave_project_pull_request_create`, `enclave_project_pull_request_respond`, `enclave_project_pull_request_view`, `enclave_secret_reminders`, `enclave_config_access_logs`, `enclave_project_members`, `enclave_project_rename`, `enclave_project_delete`, `enclave_project_webhooks`, `enclave_project_config_dynamic_secrets_manage`, `enclave_project_config_rotated_secrets_manage`, `enclave_project_config_rename`, `enclave_project_config_lock`, `enclave_project_config_delete`, `enclave_project_environment_list_all`, `enclave_project_environment_all`, `enclave_project_environment_order`, `enclave_project_environment_create`, `enclave_project_environment_delete`, `enclave_project_environment_rename`, `enclave_project_environment_settings_manage`, `enclave_project_secrets_referencing`, `enclave_config_secrets_referencing`
        """
        return pulumi.get(self, "permissions")

    @permissions.setter
    def permissions(self, value: pulumi.Input[Sequence[pulumi.Input[str]]]):
        pulumi.set(self, "permissions", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the role
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)


class ProjectRole(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 permissions: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 __props__=None):
        """
        Create a ProjectRole resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] name: The name of the role
        :param pulumi.Input[Sequence[pulumi.Input[str]]] permissions: An array containing the permissions the role has. Valid permissions are: `enclave_config_logs`, `enclave_project_config_secrets_read`, `enclave_project_config_dynamic_secrets_read`, `enclave_project_config_dynamic_secrets_leases_write`, `enclave_project_config_rotated_secrets_read`, `enclave_config_syncs_manage`, `enclave_project_secrets_notes_manage`, `enclave_project_config_create`, `enclave_project_config_duplicate`, `enclave_project_config_secrets_write`, `enclave_project_config_service_tokens`, `enclave_project_config_trusted_ips`, `enclave_project_config_logs_rollback`, `enclave_project_config_list_all`, `enclave_project_pull_request_create`, `enclave_project_pull_request_respond`, `enclave_project_pull_request_view`, `enclave_secret_reminders`, `enclave_config_access_logs`, `enclave_project_members`, `enclave_project_rename`, `enclave_project_delete`, `enclave_project_webhooks`, `enclave_project_config_dynamic_secrets_manage`, `enclave_project_config_rotated_secrets_manage`, `enclave_project_config_rename`, `enclave_project_config_lock`, `enclave_project_config_delete`, `enclave_project_environment_list_all`, `enclave_project_environment_all`, `enclave_project_environment_order`, `enclave_project_environment_create`, `enclave_project_environment_delete`, `enclave_project_environment_rename`, `enclave_project_environment_settings_manage`, `enclave_project_secrets_referencing`, `enclave_config_secrets_referencing`
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ProjectRoleArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a ProjectRole resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param ProjectRoleArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ProjectRoleArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 permissions: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ProjectRoleArgs.__new__(ProjectRoleArgs)

            __props__.__dict__["name"] = name
            if permissions is None and not opts.urn:
                raise TypeError("Missing required property 'permissions'")
            __props__.__dict__["permissions"] = permissions
            __props__.__dict__["role"] = None
        super(ProjectRole, __self__).__init__(
            'doppler-native:projects/v3:ProjectRole',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'ProjectRole':
        """
        Get an existing ProjectRole resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = ProjectRoleArgs.__new__(ProjectRoleArgs)

        __props__.__dict__["name"] = None
        __props__.__dict__["permissions"] = None
        __props__.__dict__["role"] = None
        return ProjectRole(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The name of the role
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def permissions(self) -> pulumi.Output[Sequence[str]]:
        """
        An array containing the permissions the role has. Valid permissions are: `enclave_config_logs`, `enclave_project_config_secrets_read`, `enclave_project_config_dynamic_secrets_read`, `enclave_project_config_dynamic_secrets_leases_write`, `enclave_project_config_rotated_secrets_read`, `enclave_config_syncs_manage`, `enclave_project_secrets_notes_manage`, `enclave_project_config_create`, `enclave_project_config_duplicate`, `enclave_project_config_secrets_write`, `enclave_project_config_service_tokens`, `enclave_project_config_trusted_ips`, `enclave_project_config_logs_rollback`, `enclave_project_config_list_all`, `enclave_project_pull_request_create`, `enclave_project_pull_request_respond`, `enclave_project_pull_request_view`, `enclave_secret_reminders`, `enclave_config_access_logs`, `enclave_project_members`, `enclave_project_rename`, `enclave_project_delete`, `enclave_project_webhooks`, `enclave_project_config_dynamic_secrets_manage`, `enclave_project_config_rotated_secrets_manage`, `enclave_project_config_rename`, `enclave_project_config_lock`, `enclave_project_config_delete`, `enclave_project_environment_list_all`, `enclave_project_environment_all`, `enclave_project_environment_order`, `enclave_project_environment_create`, `enclave_project_environment_delete`, `enclave_project_environment_rename`, `enclave_project_environment_settings_manage`, `enclave_project_secrets_referencing`, `enclave_config_secrets_referencing`
        """
        return pulumi.get(self, "permissions")

    @property
    @pulumi.getter
    def role(self) -> pulumi.Output[Optional['outputs.RoleProperties']]:
        return pulumi.get(self, "role")

