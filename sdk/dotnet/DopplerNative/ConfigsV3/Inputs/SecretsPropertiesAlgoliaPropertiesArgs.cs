// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace CloudySkySoftware.Pulumi.DopplerNative.ConfigsV3.Inputs
{

    public sealed class SecretsPropertiesAlgoliaPropertiesArgs : global::Pulumi.ResourceArgs
    {
        [Input("computed")]
        public Input<string>? Computed { get; set; }

        [Input("note")]
        public Input<string>? Note { get; set; }

        [Input("raw")]
        public Input<string>? Raw { get; set; }

        public SecretsPropertiesAlgoliaPropertiesArgs()
        {
        }
        public static new SecretsPropertiesAlgoliaPropertiesArgs Empty => new SecretsPropertiesAlgoliaPropertiesArgs();
    }
}
