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
    public static class GetServiceAccount
    {
        public static Task<Outputs.GetServiceAccountProperties> InvokeAsync(GetServiceAccountArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<Outputs.GetServiceAccountProperties>("doppler-native:workplace/v3:getServiceAccount", args ?? new GetServiceAccountArgs(), options.WithDefaults());

        public static Output<Outputs.GetServiceAccountProperties> Invoke(GetServiceAccountInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<Outputs.GetServiceAccountProperties>("doppler-native:workplace/v3:getServiceAccount", args ?? new GetServiceAccountInvokeArgs(), options.WithDefaults());

        public static Output<Outputs.GetServiceAccountProperties> Invoke(GetServiceAccountInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<Outputs.GetServiceAccountProperties>("doppler-native:workplace/v3:getServiceAccount", args ?? new GetServiceAccountInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetServiceAccountArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Slug of the service account
        /// </summary>
        [Input("slug", required: true)]
        public string Slug { get; set; } = null!;

        public GetServiceAccountArgs()
        {
        }
        public static new GetServiceAccountArgs Empty => new GetServiceAccountArgs();
    }

    public sealed class GetServiceAccountInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Slug of the service account
        /// </summary>
        [Input("slug", required: true)]
        public Input<string> Slug { get; set; } = null!;

        public GetServiceAccountInvokeArgs()
        {
        }
        public static new GetServiceAccountInvokeArgs Empty => new GetServiceAccountInvokeArgs();
    }
}
