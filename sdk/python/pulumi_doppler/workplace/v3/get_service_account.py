# coding=utf-8
# *** WARNING: this file was generated by pulumigen. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import sys
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload, Awaitable
if sys.version_info >= (3, 11):
    from typing import NotRequired, TypedDict, TypeAlias
else:
    from typing_extensions import NotRequired, TypedDict, TypeAlias
from ... import _utilities
from . import outputs

__all__ = [
    'GetServiceAccountProperties',
    'AwaitableGetServiceAccountProperties',
    'get_service_account',
    'get_service_account_output',
]

@pulumi.output_type
class GetServiceAccountProperties:
    def __init__(__self__, service_account=None):
        if service_account and not isinstance(service_account, dict):
            raise TypeError("Expected argument 'service_account' to be a dict")
        pulumi.set(__self__, "service_account", service_account)

    @property
    @pulumi.getter(name="serviceAccount")
    def service_account(self) -> Optional['outputs.GetServiceAccountPropertiesServiceAccountProperties']:
        return pulumi.get(self, "service_account")


class AwaitableGetServiceAccountProperties(GetServiceAccountProperties):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetServiceAccountProperties(
            service_account=self.service_account)


def get_service_account(slug: Optional[str] = None,
                        opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetServiceAccountProperties:
    """
    Use this data source to access information about an existing resource.

    :param str slug: Slug of the service account
    """
    __args__ = dict()
    __args__['slug'] = slug
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('doppler-native:workplace/v3:getServiceAccount', __args__, opts=opts, typ=GetServiceAccountProperties).value

    return AwaitableGetServiceAccountProperties(
        service_account=pulumi.get(__ret__, 'service_account'))
def get_service_account_output(slug: Optional[pulumi.Input[str]] = None,
                               opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetServiceAccountProperties]:
    """
    Use this data source to access information about an existing resource.

    :param str slug: Slug of the service account
    """
    __args__ = dict()
    __args__['slug'] = slug
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('doppler-native:workplace/v3:getServiceAccount', __args__, opts=opts, typ=GetServiceAccountProperties)
    return __ret__.apply(lambda __response__: GetServiceAccountProperties(
        service_account=pulumi.get(__response__, 'service_account')))
