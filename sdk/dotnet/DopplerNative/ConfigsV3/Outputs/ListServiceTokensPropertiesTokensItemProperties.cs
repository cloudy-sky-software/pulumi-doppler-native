// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace CloudySkySoftware.Pulumi.DopplerNative.ConfigsV3.Outputs
{

    [OutputType]
    public sealed class ListServiceTokensPropertiesTokensItemProperties
    {
        public readonly string? Config;
        public readonly string? CreatedAt;
        public readonly string? Environment;
        public readonly string? ExpiresAt;
        public readonly string? Name;
        public readonly string? Project;
        public readonly string? Slug;

        [OutputConstructor]
        private ListServiceTokensPropertiesTokensItemProperties(
            string? config,

            string? createdAt,

            string? environment,

            string? expiresAt,

            string? name,

            string? project,

            string? slug)
        {
            Config = config;
            CreatedAt = createdAt;
            Environment = environment;
            ExpiresAt = expiresAt;
            Name = name;
            Project = project;
            Slug = slug;
        }
    }
}
