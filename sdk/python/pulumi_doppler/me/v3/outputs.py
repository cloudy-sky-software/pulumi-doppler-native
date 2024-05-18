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
    'GetAuthMeProperties',
    'GetAuthMePropertiesWorkplaceProperties',
]

@pulumi.output_type
class GetAuthMeProperties(dict):
    def __init__(__self__, *,
                 created_at: Optional[str] = None,
                 last_seen_at: Optional[str] = None,
                 name: Optional[str] = None,
                 slug: Optional[str] = None,
                 token_preview: Optional[str] = None,
                 type: Optional[str] = None,
                 workplace: Optional['outputs.GetAuthMePropertiesWorkplaceProperties'] = None):
        if created_at is not None:
            pulumi.set(__self__, "created_at", created_at)
        if last_seen_at is not None:
            pulumi.set(__self__, "last_seen_at", last_seen_at)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if slug is not None:
            pulumi.set(__self__, "slug", slug)
        if token_preview is not None:
            pulumi.set(__self__, "token_preview", token_preview)
        if type is not None:
            pulumi.set(__self__, "type", type)
        if workplace is not None:
            pulumi.set(__self__, "workplace", workplace)

    @property
    @pulumi.getter(name="createdAt")
    def created_at(self) -> Optional[str]:
        return pulumi.get(self, "created_at")

    @property
    @pulumi.getter(name="lastSeenAt")
    def last_seen_at(self) -> Optional[str]:
        return pulumi.get(self, "last_seen_at")

    @property
    @pulumi.getter
    def name(self) -> Optional[str]:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def slug(self) -> Optional[str]:
        return pulumi.get(self, "slug")

    @property
    @pulumi.getter(name="tokenPreview")
    def token_preview(self) -> Optional[str]:
        return pulumi.get(self, "token_preview")

    @property
    @pulumi.getter
    def type(self) -> Optional[str]:
        return pulumi.get(self, "type")

    @property
    @pulumi.getter
    def workplace(self) -> Optional['outputs.GetAuthMePropertiesWorkplaceProperties']:
        return pulumi.get(self, "workplace")


@pulumi.output_type
class GetAuthMePropertiesWorkplaceProperties(dict):
    def __init__(__self__, *,
                 name: Optional[str] = None,
                 slug: Optional[str] = None):
        if name is not None:
            pulumi.set(__self__, "name", name)
        if slug is not None:
            pulumi.set(__self__, "slug", slug)

    @property
    @pulumi.getter
    def name(self) -> Optional[str]:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def slug(self) -> Optional[str]:
        return pulumi.get(self, "slug")

