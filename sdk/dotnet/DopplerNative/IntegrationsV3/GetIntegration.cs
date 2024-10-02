// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace CloudySkySoftware.Pulumi.DopplerNative.IntegrationsV3
{
    public static class GetIntegration
    {
        public static Task<Outputs.GetIntegrationProperties> InvokeAsync(GetIntegrationArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<Outputs.GetIntegrationProperties>("doppler-native:integrations/v3:getIntegration", args ?? new GetIntegrationArgs(), options.WithDefaults());

        public static Output<Outputs.GetIntegrationProperties> Invoke(InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<Outputs.GetIntegrationProperties>("doppler-native:integrations/v3:getIntegration", InvokeArgs.Empty, options.WithDefaults());
    }


    public sealed class GetIntegrationArgs : global::Pulumi.InvokeArgs
    {
        public GetIntegrationArgs()
        {
        }
        public static new GetIntegrationArgs Empty => new GetIntegrationArgs();
    }
}