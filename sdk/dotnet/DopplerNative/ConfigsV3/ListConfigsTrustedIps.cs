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
    public static class ListConfigsTrustedIps
    {
        public static Task<Outputs.ListConfigsTrustedIpsProperties> InvokeAsync(ListConfigsTrustedIpsArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<Outputs.ListConfigsTrustedIpsProperties>("doppler-native:configs/v3:listConfigsTrustedIps", args ?? new ListConfigsTrustedIpsArgs(), options.WithDefaults());

        public static Output<Outputs.ListConfigsTrustedIpsProperties> Invoke(InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<Outputs.ListConfigsTrustedIpsProperties>("doppler-native:configs/v3:listConfigsTrustedIps", InvokeArgs.Empty, options.WithDefaults());
    }


    public sealed class ListConfigsTrustedIpsArgs : global::Pulumi.InvokeArgs
    {
        public ListConfigsTrustedIpsArgs()
        {
        }
        public static new ListConfigsTrustedIpsArgs Empty => new ListConfigsTrustedIpsArgs();
    }
}