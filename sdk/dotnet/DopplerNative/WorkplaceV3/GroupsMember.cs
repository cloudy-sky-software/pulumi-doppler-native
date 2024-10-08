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
    [DopplerNativeResourceType("doppler-native:workplace/v3:GroupsMember")]
    public partial class GroupsMember : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The member's slug
        /// </summary>
        [Output("slug")]
        public Output<string> Slug { get; private set; } = null!;

        [Output("type")]
        public Output<CloudySkySoftware.Pulumi.DopplerNative.WorkplaceV3.Type> Type { get; private set; } = null!;


        /// <summary>
        /// Create a GroupsMember resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public GroupsMember(string name, GroupsMemberArgs args, CustomResourceOptions? options = null)
            : base("doppler-native:workplace/v3:GroupsMember", name, args ?? new GroupsMemberArgs(), MakeResourceOptions(options, ""))
        {
        }

        private GroupsMember(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("doppler-native:workplace/v3:GroupsMember", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing GroupsMember resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static GroupsMember Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new GroupsMember(name, id, options);
        }
    }

    public sealed class GroupsMemberArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The group's slug
        /// </summary>
        [Input("slug", required: true)]
        public Input<string> Slug { get; set; } = null!;

        [Input("type", required: true)]
        public Input<CloudySkySoftware.Pulumi.DopplerNative.WorkplaceV3.Type> Type { get; set; } = null!;

        public GroupsMemberArgs()
        {
        }
        public static new GroupsMemberArgs Empty => new GroupsMemberArgs();
    }
}
