// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.DopplerNative.ConfigsV3.Inputs
{

    /// <summary>
    /// The valueType you expect the secret to have before `valueType` is applied. If specified, the request will only be processed if the provided valueType matches what's found in Doppler.
    /// </summary>
    public sealed class ChangeRequestsItemPropertiesOriginalValueTypePropertiesArgs : global::Pulumi.ResourceArgs
    {
        [Input("type")]
        public Input<Pulumi.DopplerNative.ConfigsV3.ChangeRequestsItemPropertiesOriginalValueTypePropertiesType>? Type { get; set; }

        public ChangeRequestsItemPropertiesOriginalValueTypePropertiesArgs()
        {
        }
        public static new ChangeRequestsItemPropertiesOriginalValueTypePropertiesArgs Empty => new ChangeRequestsItemPropertiesOriginalValueTypePropertiesArgs();
    }
}
