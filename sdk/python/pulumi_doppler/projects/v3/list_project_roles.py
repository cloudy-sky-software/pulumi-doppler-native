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
    'ListProjectRolesProperties',
    'AwaitableListProjectRolesProperties',
    'list_project_roles',
    'list_project_roles_output',
]

@pulumi.output_type
class ListProjectRolesProperties:
    def __init__(__self__, roles=None):
        if roles and not isinstance(roles, list):
            raise TypeError("Expected argument 'roles' to be a list")
        pulumi.set(__self__, "roles", roles)

    @property
    @pulumi.getter
    def roles(self) -> Optional[Sequence['outputs.ListProjectRolesPropertiesRolesItemProperties']]:
        return pulumi.get(self, "roles")


class AwaitableListProjectRolesProperties(ListProjectRolesProperties):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return ListProjectRolesProperties(
            roles=self.roles)


def list_project_roles(opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableListProjectRolesProperties:
    """
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('doppler-native:projects/v3:listProjectRoles', __args__, opts=opts, typ=ListProjectRolesProperties).value

    return AwaitableListProjectRolesProperties(
        roles=pulumi.get(__ret__, 'roles'))
def list_project_roles_output(opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[ListProjectRolesProperties]:
    """
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('doppler-native:projects/v3:listProjectRoles', __args__, opts=opts, typ=ListProjectRolesProperties)
    return __ret__.apply(lambda __response__: ListProjectRolesProperties(
        roles=pulumi.get(__response__, 'roles')))
