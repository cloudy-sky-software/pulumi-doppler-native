// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.ComponentModel;
using Pulumi;

namespace Pulumi.DopplerNative.ProjectsV3
{
    [EnumType]
    public readonly struct ProjectMembersType : IEquatable<ProjectMembersType>
    {
        private readonly string _value;

        private ProjectMembersType(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static ProjectMembersType WorkplaceUser { get; } = new ProjectMembersType("workplace_user");
        public static ProjectMembersType Group { get; } = new ProjectMembersType("group");
        public static ProjectMembersType Invite { get; } = new ProjectMembersType("invite");
        public static ProjectMembersType ServiceAccount { get; } = new ProjectMembersType("service_account");

        public static bool operator ==(ProjectMembersType left, ProjectMembersType right) => left.Equals(right);
        public static bool operator !=(ProjectMembersType left, ProjectMembersType right) => !left.Equals(right);

        public static explicit operator string(ProjectMembersType value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is ProjectMembersType other && Equals(other);
        public bool Equals(ProjectMembersType other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }
}