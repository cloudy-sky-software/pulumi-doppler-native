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
    'GetProjectRoleProperties',
    'AwaitableGetProjectRoleProperties',
    'get_project_role',
    'get_project_role_output',
]

@pulumi.output_type
class GetProjectRoleProperties:
    def __init__(__self__, role=None):
        if role and not isinstance(role, dict):
            raise TypeError("Expected argument 'role' to be a dict")
        pulumi.set(__self__, "role", role)

    @property
    @pulumi.getter
    def role(self) -> Optional['outputs.GetProjectRolePropertiesRoleProperties']:
        return pulumi.get(self, "role")


class AwaitableGetProjectRoleProperties(GetProjectRoleProperties):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetProjectRoleProperties(
            role=self.role)


def get_project_role(role: Optional[str] = None,
                     opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetProjectRoleProperties:
    """
    Use this data source to access information about an existing resource.

    :param str role: The role's unique identifier
    """
    __args__ = dict()
    __args__['role'] = role
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('doppler-native:projects/v3:getProjectRole', __args__, opts=opts, typ=GetProjectRoleProperties).value

    return AwaitableGetProjectRoleProperties(
        role=pulumi.get(__ret__, 'role'))
def get_project_role_output(role: Optional[pulumi.Input[str]] = None,
                            opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetProjectRoleProperties]:
    """
    Use this data source to access information about an existing resource.

    :param str role: The role's unique identifier
    """
    __args__ = dict()
    __args__['role'] = role
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('doppler-native:projects/v3:getProjectRole', __args__, opts=opts, typ=GetProjectRoleProperties)
    return __ret__.apply(lambda __response__: GetProjectRoleProperties(
        role=pulumi.get(__response__, 'role')))
