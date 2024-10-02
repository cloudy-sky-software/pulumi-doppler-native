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
    public static class GetProject
    {
        public static Task<Outputs.GetProjectProperties> InvokeAsync(GetProjectArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<Outputs.GetProjectProperties>("doppler-native:projects/v3:getProject", args ?? new GetProjectArgs(), options.WithDefaults());

        public static Output<Outputs.GetProjectProperties> Invoke(InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<Outputs.GetProjectProperties>("doppler-native:projects/v3:getProject", InvokeArgs.Empty, options.WithDefaults());
    }


    public sealed class GetProjectArgs : global::Pulumi.InvokeArgs
    {
        public GetProjectArgs()
        {
        }
        public static new GetProjectArgs Empty => new GetProjectArgs();
    }
}
