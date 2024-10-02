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
    'ListWorkplaceRolesProperties',
    'AwaitableListWorkplaceRolesProperties',
    'list_workplace_roles',
    'list_workplace_roles_output',
]

@pulumi.output_type
class ListWorkplaceRolesProperties:
    def __init__(__self__, roles=None):
        if roles and not isinstance(roles, list):
            raise TypeError("Expected argument 'roles' to be a list")
        pulumi.set(__self__, "roles", roles)

    @property
    @pulumi.getter
    def roles(self) -> Optional[Sequence['outputs.ListWorkplaceRolesPropertiesRolesItemProperties']]:
        return pulumi.get(self, "roles")


class AwaitableListWorkplaceRolesProperties(ListWorkplaceRolesProperties):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return ListWorkplaceRolesProperties(
            roles=self.roles)


def list_workplace_roles(opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableListWorkplaceRolesProperties:
    """
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('doppler-native:workplace/v3:listWorkplaceRoles', __args__, opts=opts, typ=ListWorkplaceRolesProperties).value

    return AwaitableListWorkplaceRolesProperties(
        roles=pulumi.get(__ret__, 'roles'))
def list_workplace_roles_output(opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[ListWorkplaceRolesProperties]:
    """
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('doppler-native:workplace/v3:listWorkplaceRoles', __args__, opts=opts, typ=ListWorkplaceRolesProperties)
    return __ret__.apply(lambda __response__: ListWorkplaceRolesProperties(
        roles=pulumi.get(__response__, 'roles')))
