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
    public sealed class GetAuditUserPropertiesWorkplaceUserPropertiesUserProperties
    {
        public readonly string? Email;
        public readonly bool? MfaEnabled;
        public readonly string? Name;
        public readonly string? ProfileImageUrl;
        public readonly bool? SamlSsoEnabled;
        public readonly bool? ThirdpartySsoEnabled;
        public readonly string? Username;

        [OutputConstructor]
        private GetAuditUserPropertiesWorkplaceUserPropertiesUserProperties(
            string? email,

            bool? mfaEnabled,

            string? name,

            string? profileImageUrl,

            bool? samlSsoEnabled,

            bool? thirdpartySsoEnabled,

            string? username)
        {
            Email = email;
            MfaEnabled = mfaEnabled;
            Name = name;
            ProfileImageUrl = profileImageUrl;
            SamlSsoEnabled = samlSsoEnabled;
            ThirdpartySsoEnabled = thirdpartySsoEnabled;
            Username = username;
        }
    }
}
