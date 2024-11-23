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
    'ListIntegrationsProperties',
    'AwaitableListIntegrationsProperties',
    'list_integrations',
    'list_integrations_output',
]

@pulumi.output_type
class ListIntegrationsProperties:
    def __init__(__self__, integrations=None, success=None):
        if integrations and not isinstance(integrations, list):
            raise TypeError("Expected argument 'integrations' to be a list")
        pulumi.set(__self__, "integrations", integrations)
        if success and not isinstance(success, bool):
            raise TypeError("Expected argument 'success' to be a bool")
        pulumi.set(__self__, "success", success)

    @property
    @pulumi.getter
    def integrations(self) -> Optional[Sequence['outputs.ListIntegrationsPropertiesIntegrationsItemProperties']]:
        return pulumi.get(self, "integrations")

    @property
    @pulumi.getter
    def success(self) -> Optional[bool]:
        return pulumi.get(self, "success")


class AwaitableListIntegrationsProperties(ListIntegrationsProperties):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return ListIntegrationsProperties(
            integrations=self.integrations,
            success=self.success)


def list_integrations(opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableListIntegrationsProperties:
    """
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('doppler-native:integrations/v3:listIntegrations', __args__, opts=opts, typ=ListIntegrationsProperties).value

    return AwaitableListIntegrationsProperties(
        integrations=pulumi.get(__ret__, 'integrations'),
        success=pulumi.get(__ret__, 'success'))
def list_integrations_output(opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[ListIntegrationsProperties]:
    """
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('doppler-native:integrations/v3:listIntegrations', __args__, opts=opts, typ=ListIntegrationsProperties)
    return __ret__.apply(lambda __response__: ListIntegrationsProperties(
        integrations=pulumi.get(__response__, 'integrations'),
        success=pulumi.get(__response__, 'success')))
