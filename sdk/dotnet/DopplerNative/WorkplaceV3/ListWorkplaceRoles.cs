// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace CloudySkySoftware.Pulumi.DopplerNative.WorkplaceV3
{
    public static class ListWorkplaceRoles
    {
        public static Task<Outputs.ListWorkplaceRolesProperties> InvokeAsync(ListWorkplaceRolesArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<Outputs.ListWorkplaceRolesProperties>("doppler-native:workplace/v3:listWorkplaceRoles", args ?? new ListWorkplaceRolesArgs(), options.WithDefaults());

        public static Output<Outputs.ListWorkplaceRolesProperties> Invoke(InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<Outputs.ListWorkplaceRolesProperties>("doppler-native:workplace/v3:listWorkplaceRoles", InvokeArgs.Empty, options.WithDefaults());
    }


    public sealed class ListWorkplaceRolesArgs : global::Pulumi.InvokeArgs
    {
        public ListWorkplaceRolesArgs()
        {
        }
        public static new ListWorkplaceRolesArgs Empty => new ListWorkplaceRolesArgs();
    }
}
