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
    'GetMemberProperties',
    'AwaitableGetMemberProperties',
    'get_member',
    'get_member_output',
]

@pulumi.output_type
class GetMemberProperties:
    def __init__(__self__, group=None):
        if group and not isinstance(group, dict):
            raise TypeError("Expected argument 'group' to be a dict")
        pulumi.set(__self__, "group", group)

    @property
    @pulumi.getter
    def group(self) -> Optional['outputs.GetMemberPropertiesGroupProperties']:
        return pulumi.get(self, "group")


class AwaitableGetMemberProperties(GetMemberProperties):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetMemberProperties(
            group=self.group)


def get_member(group_slug: Optional[str] = None,
               member_slug: Optional[str] = None,
               member_type: Optional[str] = None,
               opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetMemberProperties:
    """
    Use this data source to access information about an existing resource.

    :param str group_slug: The group's slug
    :param str member_slug: The member's slug
    """
    __args__ = dict()
    __args__['groupSlug'] = group_slug
    __args__['memberSlug'] = member_slug
    __args__['memberType'] = member_type
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('doppler-native:workplace/v3:getMember', __args__, opts=opts, typ=GetMemberProperties).value

    return AwaitableGetMemberProperties(
        group=pulumi.get(__ret__, 'group'))
def get_member_output(group_slug: Optional[pulumi.Input[str]] = None,
                      member_slug: Optional[pulumi.Input[str]] = None,
                      member_type: Optional[pulumi.Input[str]] = None,
                      opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetMemberProperties]:
    """
    Use this data source to access information about an existing resource.

    :param str group_slug: The group's slug
    :param str member_slug: The member's slug
    """
    __args__ = dict()
    __args__['groupSlug'] = group_slug
    __args__['memberSlug'] = member_slug
    __args__['memberType'] = member_type
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('doppler-native:workplace/v3:getMember', __args__, opts=opts, typ=GetMemberProperties)
    return __ret__.apply(lambda __response__: GetMemberProperties(
        group=pulumi.get(__response__, 'group')))
