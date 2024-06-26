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
    'ListIntegrationsResult',
    'AwaitableListIntegrationsResult',
    'list_integrations',
    'list_integrations_output',
]

@pulumi.output_type
class ListIntegrationsResult:
    def __init__(__self__, items=None):
        if items and not isinstance(items, dict):
            raise TypeError("Expected argument 'items' to be a dict")
        pulumi.set(__self__, "items", items)

    @property
    @pulumi.getter
    def items(self) -> 'outputs.ListIntegrationsProperties':
        return pulumi.get(self, "items")


class AwaitableListIntegrationsResult(ListIntegrationsResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return ListIntegrationsResult(
            items=self.items)


def list_integrations(opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableListIntegrationsResult:
    """
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('doppler-native:integrations/v3:listIntegrations', __args__, opts=opts, typ=ListIntegrationsResult).value

    return AwaitableListIntegrationsResult(
        items=pulumi.get(__ret__, 'items'))


@_utilities.lift_output_func(list_integrations)
def list_integrations_output(opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[ListIntegrationsResult]:
    """
    Use this data source to access information about an existing resource.
    """
    ...
