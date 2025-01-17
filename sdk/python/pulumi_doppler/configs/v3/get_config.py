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
    'GetConfigProperties',
    'AwaitableGetConfigProperties',
    'get_config',
    'get_config_output',
]

@pulumi.output_type
class GetConfigProperties:
    def __init__(__self__, config=None):
        if config and not isinstance(config, dict):
            raise TypeError("Expected argument 'config' to be a dict")
        pulumi.set(__self__, "config", config)

    @property
    @pulumi.getter
    def config(self) -> Optional['outputs.GetConfigPropertiesConfigProperties']:
        return pulumi.get(self, "config")


class AwaitableGetConfigProperties(GetConfigProperties):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetConfigProperties(
            config=self.config)


def get_config(opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetConfigProperties:
    """
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('doppler-native:configs/v3:getConfig', __args__, opts=opts, typ=GetConfigProperties).value

    return AwaitableGetConfigProperties(
        config=pulumi.get(__ret__, 'config'))
def get_config_output(opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetConfigProperties]:
    """
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('doppler-native:configs/v3:getConfig', __args__, opts=opts, typ=GetConfigProperties)
    return __ret__.apply(lambda __response__: GetConfigProperties(
        config=pulumi.get(__response__, 'config')))
