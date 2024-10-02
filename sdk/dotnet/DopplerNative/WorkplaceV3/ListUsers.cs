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
    public static class ListUsers
    {
        public static Task<Outputs.ListUsersProperties> InvokeAsync(ListUsersArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<Outputs.ListUsersProperties>("doppler-native:workplace/v3:listUsers", args ?? new ListUsersArgs(), options.WithDefaults());

        public static Output<Outputs.ListUsersProperties> Invoke(InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<Outputs.ListUsersProperties>("doppler-native:workplace/v3:listUsers", InvokeArgs.Empty, options.WithDefaults());
    }


    public sealed class ListUsersArgs : global::Pulumi.InvokeArgs
    {
        public ListUsersArgs()
        {
        }
        public static new ListUsersArgs Empty => new ListUsersArgs();
    }
}