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
    [DopplerNativeResourceType("doppler-native:configs/v3:SecretsUpdate")]
    public partial class SecretsUpdate : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Either `secrets` or `change_requests` is required (can't use both). Object of secrets you would like to save to the config. Try it with the sample secrets below.
        /// </summary>
        [Output("changeRequests")]
        public Output<ImmutableArray<Outputs.ChangeRequestsItemProperties>> ChangeRequests { get; private set; } = null!;

        /// <summary>
        /// Name of the config object.
        /// </summary>
        [Output("config")]
        public Output<string> Config { get; private set; } = null!;

        /// <summary>
        /// Unique identifier for the project object.
        /// </summary>
        [Output("project")]
        public Output<string> Project { get; private set; } = null!;

        [Output("secrets")]
        public Output<Outputs.SecretsProperties?> Secrets { get; private set; } = null!;


        /// <summary>
        /// Create a SecretsUpdate resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public SecretsUpdate(string name, SecretsUpdateArgs args, CustomResourceOptions? options = null)
            : base("doppler-native:configs/v3:SecretsUpdate", name, args ?? new SecretsUpdateArgs(), MakeResourceOptions(options, ""))
        {
        }

        private SecretsUpdate(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("doppler-native:configs/v3:SecretsUpdate", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                PluginDownloadURL = "github://api.github.com/cloudy-sky-software/pulumi-doppler-native",
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing SecretsUpdate resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static SecretsUpdate Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new SecretsUpdate(name, id, options);
        }
    }

    public sealed class SecretsUpdateArgs : global::Pulumi.ResourceArgs
    {
        [Input("changeRequests")]
        private InputList<Inputs.ChangeRequestsItemPropertiesArgs>? _changeRequests;

        /// <summary>
        /// Either `secrets` or `change_requests` is required (can't use both). Object of secrets you would like to save to the config. Try it with the sample secrets below.
        /// </summary>
        public InputList<Inputs.ChangeRequestsItemPropertiesArgs> ChangeRequests
        {
            get => _changeRequests ?? (_changeRequests = new InputList<Inputs.ChangeRequestsItemPropertiesArgs>());
            set => _changeRequests = value;
        }

        /// <summary>
        /// Name of the config object.
        /// </summary>
        [Input("config", required: true)]
        public Input<string> Config { get; set; } = null!;

        /// <summary>
        /// Unique identifier for the project object.
        /// </summary>
        [Input("project", required: true)]
        public Input<string> Project { get; set; } = null!;

        /// <summary>
        /// Either `secrets` or `change_requests` is required (can't use both). Object of secrets you would like to save to the config. Try it with the sample secrets below.
        /// </summary>
        [Input("secrets")]
        public Input<Inputs.SecretsPropertiesArgs>? Secrets { get; set; }

        public SecretsUpdateArgs()
        {
            Config = "CONFIG_NAME";
            Project = "PROJECT_NAME";
        }
        public static new SecretsUpdateArgs Empty => new SecretsUpdateArgs();
    }
}
