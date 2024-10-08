// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace CloudySkySoftware.Pulumi.DopplerNative.WorkplaceV3.Outputs
{

    [OutputType]
    public sealed class ListServiceAccountTokensProperties
    {
        public readonly ImmutableArray<Outputs.ListServiceAccountTokensPropertiesApiTokensItemProperties> ApiTokens;
        public readonly bool? Success;

        [OutputConstructor]
        private ListServiceAccountTokensProperties(
            ImmutableArray<Outputs.ListServiceAccountTokensPropertiesApiTokensItemProperties> apiTokens,

            bool? success)
        {
            ApiTokens = apiTokens;
            Success = success;
        }
    }
}
