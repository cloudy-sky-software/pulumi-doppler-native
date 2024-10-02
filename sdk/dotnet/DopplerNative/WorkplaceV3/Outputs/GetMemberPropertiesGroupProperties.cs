// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace CloudySkySoftware.Pulumi.DopplerNative.WorkplaceV3.Outputs
{

    [OutputType]
    public sealed class GetMemberPropertiesGroupProperties
    {
        public readonly string? CreatedAt;
        public readonly Outputs.GetMemberPropertiesGroupPropertiesDefaultProjectRoleProperties? DefaultProjectRole;
        public readonly ImmutableArray<Outputs.GetMemberPropertiesGroupPropertiesMembersItemProperties> Members;
        public readonly string? Name;
        public readonly ImmutableArray<Outputs.GetMemberPropertiesGroupPropertiesProjectsItemProperties> Projects;
        public readonly string? Slug;

        [OutputConstructor]
        private GetMemberPropertiesGroupProperties(
            string? createdAt,

            Outputs.GetMemberPropertiesGroupPropertiesDefaultProjectRoleProperties? defaultProjectRole,

            ImmutableArray<Outputs.GetMemberPropertiesGroupPropertiesMembersItemProperties> members,

            string? name,

            ImmutableArray<Outputs.GetMemberPropertiesGroupPropertiesProjectsItemProperties> projects,

            string? slug)
        {
            CreatedAt = createdAt;
            DefaultProjectRole = defaultProjectRole;
            Members = members;
            Name = name;
            Projects = projects;
            Slug = slug;
        }
    }
}