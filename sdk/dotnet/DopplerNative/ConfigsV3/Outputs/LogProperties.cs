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
    public sealed class LogProperties
    {
        public readonly string? Config;
        public readonly string? CreatedAt;
        public readonly ImmutableArray<Outputs.LogPropertiesDiffItemProperties> Diff;
        public readonly string? Environment;
        public readonly string? Html;
        public readonly string? Id;
        public readonly string? Project;
        public readonly bool? Rollback;
        public readonly string? Text;
        public readonly Outputs.LogPropertiesUserProperties? User;

        [OutputConstructor]
        private LogProperties(
            string? config,

            string? createdAt,

            ImmutableArray<Outputs.LogPropertiesDiffItemProperties> diff,

            string? environment,

            string? html,

            string? id,

            string? project,

            bool? rollback,

            string? text,

            Outputs.LogPropertiesUserProperties? user)
        {
            Config = config;
            CreatedAt = createdAt;
            Diff = diff;
            Environment = environment;
            Html = html;
            Id = id;
            Project = project;
            Rollback = rollback;
            Text = text;
            User = user;
        }
    }
}
