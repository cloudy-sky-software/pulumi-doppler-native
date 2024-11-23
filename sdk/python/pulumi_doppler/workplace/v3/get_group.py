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
    'GetGroupProperties',
    'AwaitableGetGroupProperties',
    'get_group',
    'get_group_output',
]

@pulumi.output_type
class GetGroupProperties:
    def __init__(__self__, group=None):
        if group and not isinstance(group, dict):
            raise TypeError("Expected argument 'group' to be a dict")
        pulumi.set(__self__, "group", group)

    @property
    @pulumi.getter
    def group(self) -> Optional['outputs.GetGroupPropertiesGroupProperties']:
        return pulumi.get(self, "group")


class AwaitableGetGroupProperties(GetGroupProperties):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetGroupProperties(
            group=self.group)


def get_group(slug: Optional[str] = None,
              opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetGroupProperties:
    """
    Use this data source to access information about an existing resource.

    :param str slug: The group's slug
    """
    __args__ = dict()
    __args__['slug'] = slug
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('doppler-native:workplace/v3:getGroup', __args__, opts=opts, typ=GetGroupProperties).value

    return AwaitableGetGroupProperties(
        group=pulumi.get(__ret__, 'group'))
def get_group_output(slug: Optional[pulumi.Input[str]] = None,
                     opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetGroupProperties]:
    """
    Use this data source to access information about an existing resource.

    :param str slug: The group's slug
    """
    __args__ = dict()
    __args__['slug'] = slug
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('doppler-native:workplace/v3:getGroup', __args__, opts=opts, typ=GetGroupProperties)
    return __ret__.apply(lambda __response__: GetGroupProperties(
        group=pulumi.get(__response__, 'group')))
