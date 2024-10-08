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

__all__ = ['DynamicSecretsIssueLeaseArgs', 'DynamicSecretsIssueLease']

@pulumi.input_type
class DynamicSecretsIssueLeaseArgs:
    def __init__(__self__, *,
                 config: pulumi.Input[str],
                 dynamic_secret: pulumi.Input[str],
                 project: pulumi.Input[str],
                 ttl_sec: pulumi.Input[int]):
        """
        The set of arguments for constructing a DynamicSecretsIssueLease resource.
        :param pulumi.Input[str] config: The config where the dynamic secret is located
        :param pulumi.Input[str] dynamic_secret: The name of the dynamic secret for which to issue this lease
        :param pulumi.Input[str] project: The project where the dynamic secret is located
        :param pulumi.Input[int] ttl_sec: The number of seconds until this lease is automatically revoked
        """
        pulumi.set(__self__, "config", config)
        pulumi.set(__self__, "dynamic_secret", dynamic_secret)
        pulumi.set(__self__, "project", project)
        pulumi.set(__self__, "ttl_sec", ttl_sec)

    @property
    @pulumi.getter
    def config(self) -> pulumi.Input[str]:
        """
        The config where the dynamic secret is located
        """
        return pulumi.get(self, "config")

    @config.setter
    def config(self, value: pulumi.Input[str]):
        pulumi.set(self, "config", value)

    @property
    @pulumi.getter(name="dynamicSecret")
    def dynamic_secret(self) -> pulumi.Input[str]:
        """
        The name of the dynamic secret for which to issue this lease
        """
        return pulumi.get(self, "dynamic_secret")

    @dynamic_secret.setter
    def dynamic_secret(self, value: pulumi.Input[str]):
        pulumi.set(self, "dynamic_secret", value)

    @property
    @pulumi.getter
    def project(self) -> pulumi.Input[str]:
        """
        The project where the dynamic secret is located
        """
        return pulumi.get(self, "project")

    @project.setter
    def project(self, value: pulumi.Input[str]):
        pulumi.set(self, "project", value)

    @property
    @pulumi.getter(name="ttlSec")
    def ttl_sec(self) -> pulumi.Input[int]:
        """
        The number of seconds until this lease is automatically revoked
        """
        return pulumi.get(self, "ttl_sec")

    @ttl_sec.setter
    def ttl_sec(self, value: pulumi.Input[int]):
        pulumi.set(self, "ttl_sec", value)


class DynamicSecretsIssueLease(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 config: Optional[pulumi.Input[str]] = None,
                 dynamic_secret: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 ttl_sec: Optional[pulumi.Input[int]] = None,
                 __props__=None):
        """
        Create a DynamicSecretsIssueLease resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] config: The config where the dynamic secret is located
        :param pulumi.Input[str] dynamic_secret: The name of the dynamic secret for which to issue this lease
        :param pulumi.Input[str] project: The project where the dynamic secret is located
        :param pulumi.Input[int] ttl_sec: The number of seconds until this lease is automatically revoked
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: DynamicSecretsIssueLeaseArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a DynamicSecretsIssueLease resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param DynamicSecretsIssueLeaseArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(DynamicSecretsIssueLeaseArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 config: Optional[pulumi.Input[str]] = None,
                 dynamic_secret: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 ttl_sec: Optional[pulumi.Input[int]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = DynamicSecretsIssueLeaseArgs.__new__(DynamicSecretsIssueLeaseArgs)

            if config is None and not opts.urn:
                raise TypeError("Missing required property 'config'")
            __props__.__dict__["config"] = config
            if dynamic_secret is None and not opts.urn:
                raise TypeError("Missing required property 'dynamic_secret'")
            __props__.__dict__["dynamic_secret"] = dynamic_secret
            if project is None and not opts.urn:
                raise TypeError("Missing required property 'project'")
            __props__.__dict__["project"] = project
            if ttl_sec is None and not opts.urn:
                raise TypeError("Missing required property 'ttl_sec'")
            __props__.__dict__["ttl_sec"] = ttl_sec
            __props__.__dict__["expires_at"] = None
            __props__.__dict__["success"] = None
            __props__.__dict__["value"] = None
        super(DynamicSecretsIssueLease, __self__).__init__(
            'doppler-native:configs/v3:DynamicSecretsIssueLease',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'DynamicSecretsIssueLease':
        """
        Get an existing DynamicSecretsIssueLease resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = DynamicSecretsIssueLeaseArgs.__new__(DynamicSecretsIssueLeaseArgs)

        __props__.__dict__["config"] = None
        __props__.__dict__["dynamic_secret"] = None
        __props__.__dict__["expires_at"] = None
        __props__.__dict__["project"] = None
        __props__.__dict__["success"] = None
        __props__.__dict__["ttl_sec"] = None
        __props__.__dict__["value"] = None
        return DynamicSecretsIssueLease(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def config(self) -> pulumi.Output[str]:
        """
        The config where the dynamic secret is located
        """
        return pulumi.get(self, "config")

    @property
    @pulumi.getter(name="dynamicSecret")
    def dynamic_secret(self) -> pulumi.Output[str]:
        """
        The name of the dynamic secret for which to issue this lease
        """
        return pulumi.get(self, "dynamic_secret")

    @property
    @pulumi.getter(name="expiresAt")
    def expires_at(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "expires_at")

    @property
    @pulumi.getter
    def project(self) -> pulumi.Output[str]:
        """
        The project where the dynamic secret is located
        """
        return pulumi.get(self, "project")

    @property
    @pulumi.getter
    def success(self) -> pulumi.Output[Optional[bool]]:
        return pulumi.get(self, "success")

    @property
    @pulumi.getter(name="ttlSec")
    def ttl_sec(self) -> pulumi.Output[int]:
        """
        The number of seconds until this lease is automatically revoked
        """
        return pulumi.get(self, "ttl_sec")

    @property
    @pulumi.getter
    def value(self) -> pulumi.Output[Optional[Any]]:
        return pulumi.get(self, "value")

