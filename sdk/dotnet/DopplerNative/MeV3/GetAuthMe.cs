// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace CloudySkySoftware.Pulumi.DopplerNative.MeV3
{
    public static class GetAuthMe
    {
        public static Task<Outputs.GetAuthMeProperties> InvokeAsync(GetAuthMeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<Outputs.GetAuthMeProperties>("doppler-native:me/v3:getAuthMe", args ?? new GetAuthMeArgs(), options.WithDefaults());

        public static Output<Outputs.GetAuthMeProperties> Invoke(InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<Outputs.GetAuthMeProperties>("doppler-native:me/v3:getAuthMe", InvokeArgs.Empty, options.WithDefaults());

        public static Output<Outputs.GetAuthMeProperties> Invoke(InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<Outputs.GetAuthMeProperties>("doppler-native:me/v3:getAuthMe", InvokeArgs.Empty, options.WithDefaults());
    }


    public sealed class GetAuthMeArgs : global::Pulumi.InvokeArgs
    {
        public GetAuthMeArgs()
        {
        }
        public static new GetAuthMeArgs Empty => new GetAuthMeArgs();
    }
}
