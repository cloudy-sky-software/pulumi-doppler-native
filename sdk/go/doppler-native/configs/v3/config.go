// Code generated by pulumigen DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v3

import (
	"context"
	"reflect"

	"errors"
	"github.com/cloudy-sky-software/pulumi-doppler-native/sdk/go/doppler-native/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Config struct {
	pulumi.CustomResourceState

	Config ConfigPropertiesPtrOutput `pulumi:"config"`
	// Identifier for the environment object.
	Environment pulumi.StringOutput `pulumi:"environment"`
	// Name of the new branch config.
	Name pulumi.StringOutput `pulumi:"name"`
	// Unique identifier for the project object.
	Project pulumi.StringOutput `pulumi:"project"`
}

// NewConfig registers a new resource with the given unique name, arguments, and options.
func NewConfig(ctx *pulumi.Context,
	name string, args *ConfigArgs, opts ...pulumi.ResourceOption) (*Config, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Environment == nil {
		args.Environment = pulumi.String("ENVIRONMENT_ID")
	}
	if args.Name == nil {
		args.Name = pulumi.StringPtr("CONFIG_NAME")
	}
	if args.Project == nil {
		args.Project = pulumi.String("PROJECT_NAME")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Config
	err := ctx.RegisterResource("doppler-native:configs/v3:Config", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetConfig gets an existing Config resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetConfig(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ConfigState, opts ...pulumi.ResourceOption) (*Config, error) {
	var resource Config
	err := ctx.ReadResource("doppler-native:configs/v3:Config", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Config resources.
type configState struct {
}

type ConfigState struct {
}

func (ConfigState) ElementType() reflect.Type {
	return reflect.TypeOf((*configState)(nil)).Elem()
}

type configArgs struct {
	// Identifier for the environment object.
	Environment string `pulumi:"environment"`
	// Name of the new branch config.
	Name *string `pulumi:"name"`
	// Unique identifier for the project object.
	Project string `pulumi:"project"`
}

// The set of arguments for constructing a Config resource.
type ConfigArgs struct {
	// Identifier for the environment object.
	Environment pulumi.StringInput
	// Name of the new branch config.
	Name pulumi.StringPtrInput
	// Unique identifier for the project object.
	Project pulumi.StringInput
}

func (ConfigArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*configArgs)(nil)).Elem()
}

type ConfigInput interface {
	pulumi.Input

	ToConfigOutput() ConfigOutput
	ToConfigOutputWithContext(ctx context.Context) ConfigOutput
}

func (*Config) ElementType() reflect.Type {
	return reflect.TypeOf((**Config)(nil)).Elem()
}

func (i *Config) ToConfigOutput() ConfigOutput {
	return i.ToConfigOutputWithContext(context.Background())
}

func (i *Config) ToConfigOutputWithContext(ctx context.Context) ConfigOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConfigOutput)
}

type ConfigOutput struct{ *pulumi.OutputState }

func (ConfigOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Config)(nil)).Elem()
}

func (o ConfigOutput) ToConfigOutput() ConfigOutput {
	return o
}

func (o ConfigOutput) ToConfigOutputWithContext(ctx context.Context) ConfigOutput {
	return o
}

func (o ConfigOutput) Config() ConfigPropertiesPtrOutput {
	return o.ApplyT(func(v *Config) ConfigPropertiesPtrOutput { return v.Config }).(ConfigPropertiesPtrOutput)
}

// Identifier for the environment object.
func (o ConfigOutput) Environment() pulumi.StringOutput {
	return o.ApplyT(func(v *Config) pulumi.StringOutput { return v.Environment }).(pulumi.StringOutput)
}

// Name of the new branch config.
func (o ConfigOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Config) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Unique identifier for the project object.
func (o ConfigOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *Config) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ConfigInput)(nil)).Elem(), &Config{})
	pulumi.RegisterOutputType(ConfigOutput{})
}