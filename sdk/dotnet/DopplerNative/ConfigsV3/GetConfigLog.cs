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
    public static class GetConfigLog
    {
        public static Task<Outputs.GetConfigLogProperties> InvokeAsync(GetConfigLogArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<Outputs.GetConfigLogProperties>("doppler-native:configs/v3:getConfigLog", args ?? new GetConfigLogArgs(), options.WithDefaults());

        public static Output<Outputs.GetConfigLogProperties> Invoke(InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<Outputs.GetConfigLogProperties>("doppler-native:configs/v3:getConfigLog", InvokeArgs.Empty, options.WithDefaults());

        public static Output<Outputs.GetConfigLogProperties> Invoke(InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<Outputs.GetConfigLogProperties>("doppler-native:configs/v3:getConfigLog", InvokeArgs.Empty, options.WithDefaults());
    }


    public sealed class GetConfigLogArgs : global::Pulumi.InvokeArgs
    {
        public GetConfigLogArgs()
        {
        }
        public static new GetConfigLogArgs Empty => new GetConfigLogArgs();
    }
}
