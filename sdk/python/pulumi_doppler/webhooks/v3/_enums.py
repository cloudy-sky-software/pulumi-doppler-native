# coding=utf-8
# *** WARNING: this file was generated by pulumigen. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

from enum import Enum

__all__ = [
    'AuthenticationPropertiesType',
]


class AuthenticationPropertiesType(str, Enum):
    NONE = "None"
    BEARER = "Bearer"
    BASIC = "Basic"
