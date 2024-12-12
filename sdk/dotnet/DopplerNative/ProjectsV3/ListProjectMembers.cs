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
    public static class ListProjectMembers
    {
        public static Task<Outputs.ListProjectMembersProperties> InvokeAsync(ListProjectMembersArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<Outputs.ListProjectMembersProperties>("doppler-native:projects/v3:listProjectMembers", args ?? new ListProjectMembersArgs(), options.WithDefaults());

        public static Output<Outputs.ListProjectMembersProperties> Invoke(InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<Outputs.ListProjectMembersProperties>("doppler-native:projects/v3:listProjectMembers", InvokeArgs.Empty, options.WithDefaults());

        public static Output<Outputs.ListProjectMembersProperties> Invoke(InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<Outputs.ListProjectMembersProperties>("doppler-native:projects/v3:listProjectMembers", InvokeArgs.Empty, options.WithDefaults());
    }


    public sealed class ListProjectMembersArgs : global::Pulumi.InvokeArgs
    {
        public ListProjectMembersArgs()
        {
        }
        public static new ListProjectMembersArgs Empty => new ListProjectMembersArgs();
    }
}
