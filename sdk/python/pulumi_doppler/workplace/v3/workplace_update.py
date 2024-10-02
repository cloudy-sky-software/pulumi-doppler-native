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

__all__ = ['WorkplaceUpdateArgs', 'WorkplaceUpdate']

@pulumi.input_type
class WorkplaceUpdateArgs:
    def __init__(__self__, *,
                 billing_email: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 security_email: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a WorkplaceUpdate resource.
        :param pulumi.Input[str] name: Workplace name
        """
        if billing_email is not None:
            pulumi.set(__self__, "billing_email", billing_email)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if security_email is not None:
            pulumi.set(__self__, "security_email", security_email)

    @property
    @pulumi.getter(name="billingEmail")
    def billing_email(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "billing_email")

    @billing_email.setter
    def billing_email(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "billing_email", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Workplace name
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="securityEmail")
    def security_email(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "security_email")

    @security_email.setter
    def security_email(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "security_email", value)


class WorkplaceUpdate(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 billing_email: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 security_email: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Create a WorkplaceUpdate resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] name: Workplace name
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[WorkplaceUpdateArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a WorkplaceUpdate resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param WorkplaceUpdateArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(WorkplaceUpdateArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 billing_email: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 security_email: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = WorkplaceUpdateArgs.__new__(WorkplaceUpdateArgs)

            __props__.__dict__["billing_email"] = billing_email
            __props__.__dict__["name"] = name
            __props__.__dict__["security_email"] = security_email
            __props__.__dict__["workplace"] = None
        super(WorkplaceUpdate, __self__).__init__(
            'doppler-native:workplace/v3:WorkplaceUpdate',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'WorkplaceUpdate':
        """
        Get an existing WorkplaceUpdate resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = WorkplaceUpdateArgs.__new__(WorkplaceUpdateArgs)

        __props__.__dict__["billing_email"] = None
        __props__.__dict__["name"] = None
        __props__.__dict__["security_email"] = None
        __props__.__dict__["workplace"] = None
        return WorkplaceUpdate(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="billingEmail")
    def billing_email(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "billing_email")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[Optional[str]]:
        """
        Workplace name
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="securityEmail")
    def security_email(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "security_email")

    @property
    @pulumi.getter
    def workplace(self) -> pulumi.Output[Optional['outputs.WorkplaceProperties']]:
        return pulumi.get(self, "workplace")

