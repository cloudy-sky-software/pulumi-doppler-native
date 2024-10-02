// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace CloudySkySoftware.Pulumi.DopplerNative.ProjectsV3
{
    public static class ListProjects
    {
        public static Task<Outputs.ListProjectsProperties> InvokeAsync(ListProjectsArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<Outputs.ListProjectsProperties>("doppler-native:projects/v3:listProjects", args ?? new ListProjectsArgs(), options.WithDefaults());

        public static Output<Outputs.ListProjectsProperties> Invoke(InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<Outputs.ListProjectsProperties>("doppler-native:projects/v3:listProjects", InvokeArgs.Empty, options.WithDefaults());
    }


    public sealed class ListProjectsArgs : global::Pulumi.InvokeArgs
    {
        public ListProjectsArgs()
        {
        }
        public static new ListProjectsArgs Empty => new ListProjectsArgs();
    }
}