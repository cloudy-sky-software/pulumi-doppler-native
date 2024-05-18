// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.DopplerNative.ConfigsV3
{
    public static class ListSecrets
    {
        public static Task<ListSecretsResult> InvokeAsync(ListSecretsArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<ListSecretsResult>("doppler-native:configs/v3:listSecrets", args ?? new ListSecretsArgs(), options.WithDefaults());

        public static Output<ListSecretsResult> Invoke(InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<ListSecretsResult>("doppler-native:configs/v3:listSecrets", InvokeArgs.Empty, options.WithDefaults());
    }


    public sealed class ListSecretsArgs : global::Pulumi.InvokeArgs
    {
        public ListSecretsArgs()
        {
        }
        public static new ListSecretsArgs Empty => new ListSecretsArgs();
    }


    [OutputType]
    public sealed class ListSecretsResult
    {
        public readonly Outputs.ListSecretsProperties Items;

        [OutputConstructor]
        private ListSecretsResult(Outputs.ListSecretsProperties items)
        {
            Items = items;
        }
    }
}