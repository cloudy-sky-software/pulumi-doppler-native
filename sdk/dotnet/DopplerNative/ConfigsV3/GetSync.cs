// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace CloudySkySoftware.Pulumi.DopplerNative.ConfigsV3
{
    public static class GetSync
    {
        public static Task<Outputs.GetSyncProperties> InvokeAsync(GetSyncArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<Outputs.GetSyncProperties>("doppler-native:configs/v3:getSync", args ?? new GetSyncArgs(), options.WithDefaults());

        public static Output<Outputs.GetSyncProperties> Invoke(InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<Outputs.GetSyncProperties>("doppler-native:configs/v3:getSync", InvokeArgs.Empty, options.WithDefaults());
    }


    public sealed class GetSyncArgs : global::Pulumi.InvokeArgs
    {
        public GetSyncArgs()
        {
        }
        public static new GetSyncArgs Empty => new GetSyncArgs();
    }
}
