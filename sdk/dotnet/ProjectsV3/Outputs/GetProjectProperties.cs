// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.DopplerNative.ProjectsV3.Outputs
{

    [OutputType]
    public sealed class GetProjectProperties
    {
        public readonly Outputs.GetProjectPropertiesProjectProperties? Project;

        [OutputConstructor]
        private GetProjectProperties(Outputs.GetProjectPropertiesProjectProperties? project)
        {
            Project = project;
        }
    }
}