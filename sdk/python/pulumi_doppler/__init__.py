# coding=utf-8
# *** WARNING: this file was generated by pulumigen. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

from . import _utilities
import typing
# Export this package's modules as members:
from .provider import *

# Make subpackages available:
if typing.TYPE_CHECKING:
    import pulumi_doppler.config as __config
    config = __config
    import pulumi_doppler.configs as __configs
    configs = __configs
    import pulumi_doppler.environments as __environments
    environments = __environments
    import pulumi_doppler.integrations as __integrations
    integrations = __integrations
    import pulumi_doppler.me as __me
    me = __me
    import pulumi_doppler.projects as __projects
    projects = __projects
    import pulumi_doppler.webhooks as __webhooks
    webhooks = __webhooks
    import pulumi_doppler.workplace as __workplace
    workplace = __workplace
else:
    config = _utilities.lazy_import('pulumi_doppler.config')
    configs = _utilities.lazy_import('pulumi_doppler.configs')
    environments = _utilities.lazy_import('pulumi_doppler.environments')
    integrations = _utilities.lazy_import('pulumi_doppler.integrations')
    me = _utilities.lazy_import('pulumi_doppler.me')
    projects = _utilities.lazy_import('pulumi_doppler.projects')
    webhooks = _utilities.lazy_import('pulumi_doppler.webhooks')
    workplace = _utilities.lazy_import('pulumi_doppler.workplace')

_utilities.register(
    resource_modules="""
[
 {
  "pkg": "doppler-native",
  "mod": "configs/v3",
  "fqn": "pulumi_doppler.configs.v3",
  "classes": {
   "doppler-native:configs/v3:ConfigLogsRollback": "ConfigLogsRollback",
   "doppler-native:configs/v3:Configs": "Configs",
   "doppler-native:configs/v3:ConfigsClone": "ConfigsClone",
   "doppler-native:configs/v3:ConfigsLock": "ConfigsLock",
   "doppler-native:configs/v3:ConfigsTrustedIp": "ConfigsTrustedIp",
   "doppler-native:configs/v3:ConfigsUnlock": "ConfigsUnlock",
   "doppler-native:configs/v3:ConfigsUpdate": "ConfigsUpdate",
   "doppler-native:configs/v3:DynamicSecretsIssueLease": "DynamicSecretsIssueLease",
   "doppler-native:configs/v3:SecretsUpdate": "SecretsUpdate",
   "doppler-native:configs/v3:SecretsUpdateNote": "SecretsUpdateNote",
   "doppler-native:configs/v3:ServiceTokens": "ServiceTokens",
   "doppler-native:configs/v3:Syncs": "Syncs"
  }
 },
 {
  "pkg": "doppler-native",
  "mod": "environments/v3",
  "fqn": "pulumi_doppler.environments.v3",
  "classes": {
   "doppler-native:environments/v3:Environments": "Environments",
   "doppler-native:environments/v3:EnvironmentsRename": "EnvironmentsRename"
  }
 },
 {
  "pkg": "doppler-native",
  "mod": "integrations/v3",
  "fqn": "pulumi_doppler.integrations.v3",
  "classes": {
   "doppler-native:integrations/v3:Integrations": "Integrations"
  }
 },
 {
  "pkg": "doppler-native",
  "mod": "projects/v3",
  "fqn": "pulumi_doppler.projects.v3",
  "classes": {
   "doppler-native:projects/v3:ProjectMembers": "ProjectMembers",
   "doppler-native:projects/v3:ProjectRoles": "ProjectRoles",
   "doppler-native:projects/v3:Projects": "Projects",
   "doppler-native:projects/v3:ProjectsUpdate": "ProjectsUpdate"
  }
 },
 {
  "pkg": "doppler-native",
  "mod": "webhooks/v3",
  "fqn": "pulumi_doppler.webhooks.v3",
  "classes": {
   "doppler-native:webhooks/v3:Webhooks": "Webhooks",
   "doppler-native:webhooks/v3:WebhooksDisable": "WebhooksDisable",
   "doppler-native:webhooks/v3:WebhooksEnable": "WebhooksEnable"
  }
 },
 {
  "pkg": "doppler-native",
  "mod": "workplace/v3",
  "fqn": "pulumi_doppler.workplace.v3",
  "classes": {
   "doppler-native:workplace/v3:Groups": "Groups",
   "doppler-native:workplace/v3:GroupsMember": "GroupsMember",
   "doppler-native:workplace/v3:ServiceAccountTokens": "ServiceAccountTokens",
   "doppler-native:workplace/v3:ServiceAccounts": "ServiceAccounts",
   "doppler-native:workplace/v3:WorkplaceRoles": "WorkplaceRoles",
   "doppler-native:workplace/v3:WorkplaceUpdate": "WorkplaceUpdate"
  }
 }
]
""",
    resource_packages="""
[
 {
  "pkg": "doppler-native",
  "token": "pulumi:providers:doppler-native",
  "fqn": "pulumi_doppler",
  "class": "Provider"
 }
]
"""
)