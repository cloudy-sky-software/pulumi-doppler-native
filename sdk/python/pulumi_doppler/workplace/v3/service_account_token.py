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

__all__ = ['ServiceAccountTokenArgs', 'ServiceAccountToken']

@pulumi.input_type
class ServiceAccountTokenArgs:
    def __init__(__self__, *,
                 expires_at: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 service_account: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a ServiceAccountToken resource.
        :param pulumi.Input[str] expires_at: The datetime at which the API token should expire. If not provided, the API token will remain vaild indefinitely unless manually revoked
        :param pulumi.Input[str] name: The display name of the API token
        :param pulumi.Input[str] service_account: Slug of the service account
        """
        if expires_at is not None:
            pulumi.set(__self__, "expires_at", expires_at)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if service_account is not None:
            pulumi.set(__self__, "service_account", service_account)

    @property
    @pulumi.getter(name="expiresAt")
    def expires_at(self) -> Optional[pulumi.Input[str]]:
        """
        The datetime at which the API token should expire. If not provided, the API token will remain vaild indefinitely unless manually revoked
        """
        return pulumi.get(self, "expires_at")

    @expires_at.setter
    def expires_at(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "expires_at", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The display name of the API token
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="serviceAccount")
    def service_account(self) -> Optional[pulumi.Input[str]]:
        """
        Slug of the service account
        """
        return pulumi.get(self, "service_account")

    @service_account.setter
    def service_account(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "service_account", value)


class ServiceAccountToken(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 expires_at: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 service_account: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Create a ServiceAccountToken resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] expires_at: The datetime at which the API token should expire. If not provided, the API token will remain vaild indefinitely unless manually revoked
        :param pulumi.Input[str] name: The display name of the API token
        :param pulumi.Input[str] service_account: Slug of the service account
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[ServiceAccountTokenArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a ServiceAccountToken resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param ServiceAccountTokenArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ServiceAccountTokenArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 expires_at: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 service_account: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ServiceAccountTokenArgs.__new__(ServiceAccountTokenArgs)

            __props__.__dict__["expires_at"] = expires_at
            __props__.__dict__["name"] = name
            __props__.__dict__["service_account"] = service_account
            __props__.__dict__["api_key"] = None
            __props__.__dict__["api_token"] = None
            __props__.__dict__["success"] = None
        super(ServiceAccountToken, __self__).__init__(
            'doppler-native:workplace/v3:ServiceAccountToken',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'ServiceAccountToken':
        """
        Get an existing ServiceAccountToken resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = ServiceAccountTokenArgs.__new__(ServiceAccountTokenArgs)

        __props__.__dict__["api_key"] = None
        __props__.__dict__["api_token"] = None
        __props__.__dict__["expires_at"] = None
        __props__.__dict__["name"] = None
        __props__.__dict__["success"] = None
        return ServiceAccountToken(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="apiKey")
    def api_key(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "api_key")

    @property
    @pulumi.getter(name="apiToken")
    def api_token(self) -> pulumi.Output[Optional['outputs.ApiTokenProperties']]:
        return pulumi.get(self, "api_token")

    @property
    @pulumi.getter(name="expiresAt")
    def expires_at(self) -> pulumi.Output[Optional[str]]:
        """
        The datetime at which the API token should expire. If not provided, the API token will remain vaild indefinitely unless manually revoked
        """
        return pulumi.get(self, "expires_at")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[Optional[str]]:
        """
        The display name of the API token
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def success(self) -> pulumi.Output[Optional[bool]]:
        return pulumi.get(self, "success")
