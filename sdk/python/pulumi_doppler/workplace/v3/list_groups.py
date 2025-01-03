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
    'ListGroupsProperties',
    'AwaitableListGroupsProperties',
    'list_groups',
    'list_groups_output',
]

@pulumi.output_type
class ListGroupsProperties:
    def __init__(__self__, groups=None):
        if groups and not isinstance(groups, list):
            raise TypeError("Expected argument 'groups' to be a list")
        pulumi.set(__self__, "groups", groups)

    @property
    @pulumi.getter
    def groups(self) -> Optional[Sequence['outputs.ListGroupsPropertiesGroupsItemProperties']]:
        return pulumi.get(self, "groups")


class AwaitableListGroupsProperties(ListGroupsProperties):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return ListGroupsProperties(
            groups=self.groups)


def list_groups(opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableListGroupsProperties:
    """
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('doppler-native:workplace/v3:listGroups', __args__, opts=opts, typ=ListGroupsProperties).value

    return AwaitableListGroupsProperties(
        groups=pulumi.get(__ret__, 'groups'))
def list_groups_output(opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[ListGroupsProperties]:
    """
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('doppler-native:workplace/v3:listGroups', __args__, opts=opts, typ=ListGroupsProperties)
    return __ret__.apply(lambda __response__: ListGroupsProperties(
        groups=pulumi.get(__response__, 'groups')))
