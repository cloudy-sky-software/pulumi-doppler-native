// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.DopplerNative.WorkplaceV3.Outputs
{

    [OutputType]
    public sealed class ListServiceAccountsPropertiesServiceAccountsItemProperties
    {
        public readonly string? CreatedAt;
        public readonly string? Name;
        public readonly string? Slug;
        public readonly Outputs.ListServiceAccountsPropertiesServiceAccountsItemPropertiesWorkplaceRoleProperties? WorkplaceRole;

        [OutputConstructor]
        private ListServiceAccountsPropertiesServiceAccountsItemProperties(
            string? createdAt,

            string? name,

            string? slug,

            Outputs.ListServiceAccountsPropertiesServiceAccountsItemPropertiesWorkplaceRoleProperties? workplaceRole)
        {
            CreatedAt = createdAt;
            Name = name;
            Slug = slug;
            WorkplaceRole = workplaceRole;
        }
    }
}