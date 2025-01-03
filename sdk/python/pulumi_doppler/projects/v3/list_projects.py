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
    'ListProjectsProperties',
    'AwaitableListProjectsProperties',
    'list_projects',
    'list_projects_output',
]

@pulumi.output_type
class ListProjectsProperties:
    def __init__(__self__, page=None, projects=None):
        if page and not isinstance(page, int):
            raise TypeError("Expected argument 'page' to be a int")
        pulumi.set(__self__, "page", page)
        if projects and not isinstance(projects, list):
            raise TypeError("Expected argument 'projects' to be a list")
        pulumi.set(__self__, "projects", projects)

    @property
    @pulumi.getter
    def page(self) -> Optional[int]:
        return pulumi.get(self, "page")

    @property
    @pulumi.getter
    def projects(self) -> Optional[Sequence['outputs.ListProjectsPropertiesProjectsItemProperties']]:
        return pulumi.get(self, "projects")


class AwaitableListProjectsProperties(ListProjectsProperties):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return ListProjectsProperties(
            page=self.page,
            projects=self.projects)


def list_projects(opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableListProjectsProperties:
    """
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('doppler-native:projects/v3:listProjects', __args__, opts=opts, typ=ListProjectsProperties).value

    return AwaitableListProjectsProperties(
        page=pulumi.get(__ret__, 'page'),
        projects=pulumi.get(__ret__, 'projects'))
def list_projects_output(opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[ListProjectsProperties]:
    """
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('doppler-native:projects/v3:listProjects', __args__, opts=opts, typ=ListProjectsProperties)
    return __ret__.apply(lambda __response__: ListProjectsProperties(
        page=pulumi.get(__response__, 'page'),
        projects=pulumi.get(__response__, 'projects')))
