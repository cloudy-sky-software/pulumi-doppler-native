// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace CloudySkySoftware.Pulumi.DopplerNative.ProjectsV3.Outputs
{

    [OutputType]
    public sealed class MemberProperties
    {
        public readonly bool? AccessAllEnvironments;
        public readonly ImmutableArray<string> Environments;
        public readonly Outputs.MemberPropertiesRoleProperties? Role;
        public readonly string? Slug;
        public readonly string? Type;

        [OutputConstructor]
        private MemberProperties(
            bool? accessAllEnvironments,

            ImmutableArray<string> environments,

            Outputs.MemberPropertiesRoleProperties? role,

            string? slug,

            string? type)
        {
            AccessAllEnvironments = accessAllEnvironments;
            Environments = environments;
            Role = role;
            Slug = slug;
            Type = type;
        }
    }
}