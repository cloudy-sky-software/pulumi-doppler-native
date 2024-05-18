# coding=utf-8
# *** WARNING: this file was generated by pulumigen. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from ... import _utilities
from . import outputs

__all__ = [
    'ListUsersResult',
    'AwaitableListUsersResult',
    'list_users',
    'list_users_output',
]

@pulumi.output_type
class ListUsersResult:
    def __init__(__self__, items=None):
        if items and not isinstance(items, dict):
            raise TypeError("Expected argument 'items' to be a dict")
        pulumi.set(__self__, "items", items)

    @property
    @pulumi.getter
    def items(self) -> 'outputs.ListUsersProperties':
        return pulumi.get(self, "items")


class AwaitableListUsersResult(ListUsersResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return ListUsersResult(
            items=self.items)


def list_users(opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableListUsersResult:
    """
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('doppler-native:workplace/v3:listUsers', __args__, opts=opts, typ=ListUsersResult).value

    return AwaitableListUsersResult(
        items=pulumi.get(__ret__, 'items'))


@_utilities.lift_output_func(list_users)
def list_users_output(opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[ListUsersResult]:
    """
    Use this data source to access information about an existing resource.
    """
    ...