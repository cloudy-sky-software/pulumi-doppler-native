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
    [DopplerNativeResourceType("doppler-native:configs/v3:ConfigsUnlock")]
    public partial class ConfigsUnlock : global::Pulumi.CustomResource
    {
        [Output("config")]
        public Output<Outputs.ConfigProperties> Config { get; private set; } = null!;

        /// <summary>
        /// Unique identifier for the project object.
        /// </summary>
        [Output("project")]
        public Output<string> Project { get; private set; } = null!;


        /// <summary>
        /// Create a ConfigsUnlock resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ConfigsUnlock(string name, ConfigsUnlockArgs args, CustomResourceOptions? options = null)
            : base("doppler-native:configs/v3:ConfigsUnlock", name, args ?? new ConfigsUnlockArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ConfigsUnlock(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("doppler-native:configs/v3:ConfigsUnlock", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing ConfigsUnlock resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ConfigsUnlock Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new ConfigsUnlock(name, id, options);
        }
    }

    public sealed class ConfigsUnlockArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Name of the config.
        /// </summary>
        [Input("config", required: true)]
        public Input<string> Config { get; set; } = null!;

        /// <summary>
        /// Unique identifier for the project object.
        /// </summary>
        [Input("project", required: true)]
        public Input<string> Project { get; set; } = null!;

        public ConfigsUnlockArgs()
        {
            Config = "CONFIG_NAME";
            Project = "PROJECT_NAME";
        }
        public static new ConfigsUnlockArgs Empty => new ConfigsUnlockArgs();
    }
}
