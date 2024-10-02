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
    'GetUserProperties',
    'AwaitableGetUserProperties',
    'get_user',
    'get_user_output',
]

@pulumi.output_type
class GetUserProperties:
    def __init__(__self__, success=None, workplace_user=None):
        if success and not isinstance(success, bool):
            raise TypeError("Expected argument 'success' to be a bool")
        pulumi.set(__self__, "success", success)
        if workplace_user and not isinstance(workplace_user, dict):
            raise TypeError("Expected argument 'workplace_user' to be a dict")
        pulumi.set(__self__, "workplace_user", workplace_user)

    @property
    @pulumi.getter
    def success(self) -> Optional[bool]:
        return pulumi.get(self, "success")

    @property
    @pulumi.getter(name="workplaceUser")
    def workplace_user(self) -> Optional['outputs.GetUserPropertiesWorkplaceUserProperties']:
        return pulumi.get(self, "workplace_user")


class AwaitableGetUserProperties(GetUserProperties):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetUserProperties(
            success=self.success,
            workplace_user=self.workplace_user)


def get_user(slug: Optional[str] = None,
             opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetUserProperties:
    """
    Use this data source to access information about an existing resource.

    :param str slug: The slug of the workplace user
    """
    __args__ = dict()
    __args__['slug'] = slug
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('doppler-native:workplace/v3:getUser', __args__, opts=opts, typ=GetUserProperties).value

    return AwaitableGetUserProperties(
        success=pulumi.get(__ret__, 'success'),
        workplace_user=pulumi.get(__ret__, 'workplace_user'))
def get_user_output(slug: Optional[pulumi.Input[str]] = None,
                    opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetUserProperties]:
    """
    Use this data source to access information about an existing resource.

    :param str slug: The slug of the workplace user
    """
    __args__ = dict()
    __args__['slug'] = slug
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('doppler-native:workplace/v3:getUser', __args__, opts=opts, typ=GetUserProperties)
    return __ret__.apply(lambda __response__: GetUserProperties(
        success=pulumi.get(__response__, 'success'),
        workplace_user=pulumi.get(__response__, 'workplace_user')))
