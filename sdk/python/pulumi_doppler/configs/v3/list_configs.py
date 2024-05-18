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
    'ListConfigsResult',
    'AwaitableListConfigsResult',
    'list_configs',
    'list_configs_output',
]

@pulumi.output_type
class ListConfigsResult:
    def __init__(__self__, items=None):
        if items and not isinstance(items, dict):
            raise TypeError("Expected argument 'items' to be a dict")
        pulumi.set(__self__, "items", items)

    @property
    @pulumi.getter
    def items(self) -> 'outputs.ListConfigsProperties':
        return pulumi.get(self, "items")


class AwaitableListConfigsResult(ListConfigsResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return ListConfigsResult(
            items=self.items)


def list_configs(opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableListConfigsResult:
    """
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('doppler-native:configs/v3:listConfigs', __args__, opts=opts, typ=ListConfigsResult).value

    return AwaitableListConfigsResult(
        items=pulumi.get(__ret__, 'items'))


@_utilities.lift_output_func(list_configs)
def list_configs_output(opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[ListConfigsResult]:
    """
    Use this data source to access information about an existing resource.
    """
    ...