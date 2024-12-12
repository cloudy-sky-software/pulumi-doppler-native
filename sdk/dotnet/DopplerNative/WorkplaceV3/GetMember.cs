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
    public static class GetMember
    {
        public static Task<Outputs.GetMemberProperties> InvokeAsync(GetMemberArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<Outputs.GetMemberProperties>("doppler-native:workplace/v3:getMember", args ?? new GetMemberArgs(), options.WithDefaults());

        public static Output<Outputs.GetMemberProperties> Invoke(GetMemberInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<Outputs.GetMemberProperties>("doppler-native:workplace/v3:getMember", args ?? new GetMemberInvokeArgs(), options.WithDefaults());

        public static Output<Outputs.GetMemberProperties> Invoke(GetMemberInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<Outputs.GetMemberProperties>("doppler-native:workplace/v3:getMember", args ?? new GetMemberInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetMemberArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The group's slug
        /// </summary>
        [Input("groupSlug", required: true)]
        public string GroupSlug { get; set; } = null!;

        /// <summary>
        /// The member's slug
        /// </summary>
        [Input("memberSlug", required: true)]
        public string MemberSlug { get; set; } = null!;

        [Input("memberType", required: true)]
        public string MemberType { get; set; } = null!;

        public GetMemberArgs()
        {
        }
        public static new GetMemberArgs Empty => new GetMemberArgs();
    }

    public sealed class GetMemberInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The group's slug
        /// </summary>
        [Input("groupSlug", required: true)]
        public Input<string> GroupSlug { get; set; } = null!;

        /// <summary>
        /// The member's slug
        /// </summary>
        [Input("memberSlug", required: true)]
        public Input<string> MemberSlug { get; set; } = null!;

        [Input("memberType", required: true)]
        public Input<string> MemberType { get; set; } = null!;

        public GetMemberInvokeArgs()
        {
        }
        public static new GetMemberInvokeArgs Empty => new GetMemberInvokeArgs();
    }
}
