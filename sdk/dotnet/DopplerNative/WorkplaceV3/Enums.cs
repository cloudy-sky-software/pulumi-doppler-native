// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.ComponentModel;
using Pulumi;

namespace CloudySkySoftware.Pulumi.DopplerNative.WorkplaceV3
{
    [EnumType]
    public readonly struct Type : IEquatable<Type>
    {
        private readonly string _value;

        private Type(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static Type WorkplaceUser { get; } = new Type("workplace_user");

        public static bool operator ==(Type left, Type right) => left.Equals(right);
        public static bool operator !=(Type left, Type right) => !left.Equals(right);

        public static explicit operator string(Type value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is Type other && Equals(other);
        public bool Equals(Type other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }
}
