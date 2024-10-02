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
    'GetSyncProperties',
    'AwaitableGetSyncProperties',
    'get_sync',
    'get_sync_output',
]

@pulumi.output_type
class GetSyncProperties:
    def __init__(__self__, sync=None):
        if sync and not isinstance(sync, dict):
            raise TypeError("Expected argument 'sync' to be a dict")
        pulumi.set(__self__, "sync", sync)

    @property
    @pulumi.getter
    def sync(self) -> Optional['outputs.GetSyncPropertiesSyncProperties']:
        return pulumi.get(self, "sync")


class AwaitableGetSyncProperties(GetSyncProperties):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetSyncProperties(
            sync=self.sync)


def get_sync(opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetSyncProperties:
    """
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('doppler-native:configs/v3:getSync', __args__, opts=opts, typ=GetSyncProperties).value

    return AwaitableGetSyncProperties(
        sync=pulumi.get(__ret__, 'sync'))
def get_sync_output(opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetSyncProperties]:
    """
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('doppler-native:configs/v3:getSync', __args__, opts=opts, typ=GetSyncProperties)
    return __ret__.apply(lambda __response__: GetSyncProperties(
        sync=pulumi.get(__response__, 'sync')))
