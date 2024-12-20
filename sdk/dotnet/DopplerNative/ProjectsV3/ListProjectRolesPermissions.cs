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
    public static class ListProjectRolesPermissions
    {
        public static Task<Outputs.ListProjectRolesPermissionsProperties> InvokeAsync(ListProjectRolesPermissionsArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<Outputs.ListProjectRolesPermissionsProperties>("doppler-native:projects/v3:listProjectRolesPermissions", args ?? new ListProjectRolesPermissionsArgs(), options.WithDefaults());

        public static Output<Outputs.ListProjectRolesPermissionsProperties> Invoke(InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<Outputs.ListProjectRolesPermissionsProperties>("doppler-native:projects/v3:listProjectRolesPermissions", InvokeArgs.Empty, options.WithDefaults());

        public static Output<Outputs.ListProjectRolesPermissionsProperties> Invoke(InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<Outputs.ListProjectRolesPermissionsProperties>("doppler-native:projects/v3:listProjectRolesPermissions", InvokeArgs.Empty, options.WithDefaults());
    }


    public sealed class ListProjectRolesPermissionsArgs : global::Pulumi.InvokeArgs
    {
        public ListProjectRolesPermissionsArgs()
        {
        }
        public static new ListProjectRolesPermissionsArgs Empty => new ListProjectRolesPermissionsArgs();
    }
}
