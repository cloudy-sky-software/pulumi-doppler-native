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
    public sealed class GroupProperties
    {
        public readonly string? CreatedAt;
        public readonly Outputs.GroupPropertiesDefaultProjectRoleProperties? DefaultProjectRole;
        public readonly ImmutableArray<Outputs.GroupPropertiesMembersItemProperties> Members;
        public readonly string? Name;
        public readonly ImmutableArray<Outputs.GroupPropertiesProjectsItemProperties> Projects;
        public readonly string? Slug;

        [OutputConstructor]
        private GroupProperties(
            string? createdAt,

            Outputs.GroupPropertiesDefaultProjectRoleProperties? defaultProjectRole,

            ImmutableArray<Outputs.GroupPropertiesMembersItemProperties> members,

            string? name,

            ImmutableArray<Outputs.GroupPropertiesProjectsItemProperties> projects,

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