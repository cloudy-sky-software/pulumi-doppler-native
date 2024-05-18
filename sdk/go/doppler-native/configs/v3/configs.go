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

type Configs struct {
	pulumi.CustomResourceState

	Config ConfigPropertiesPtrOutput `pulumi:"config"`
	// Identifier for the environment object.
	Environment pulumi.StringOutput `pulumi:"environment"`
	// Name of the new branch config.
	Name pulumi.StringOutput `pulumi:"name"`
	// Unique identifier for the project object.
	Project pulumi.StringOutput `pulumi:"project"`
}

// NewConfigs registers a new resource with the given unique name, arguments, and options.
func NewConfigs(ctx *pulumi.Context,
	name string, args *ConfigsArgs, opts ...pulumi.ResourceOption) (*Configs, error) {
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
	var resource Configs
	err := ctx.RegisterResource("doppler-native:configs/v3:Configs", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetConfigs gets an existing Configs resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetConfigs(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ConfigsState, opts ...pulumi.ResourceOption) (*Configs, error) {
	var resource Configs
	err := ctx.ReadResource("doppler-native:configs/v3:Configs", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Configs resources.
type configsState struct {
}

type ConfigsState struct {
}

func (ConfigsState) ElementType() reflect.Type {
	return reflect.TypeOf((*configsState)(nil)).Elem()
}

type configsArgs struct {
	// Identifier for the environment object.
	Environment string `pulumi:"environment"`
	// Name of the new branch config.
	Name *string `pulumi:"name"`
	// Unique identifier for the project object.
	Project string `pulumi:"project"`
}

// The set of arguments for constructing a Configs resource.
type ConfigsArgs struct {
	// Identifier for the environment object.
	Environment pulumi.StringInput
	// Name of the new branch config.
	Name pulumi.StringPtrInput
	// Unique identifier for the project object.
	Project pulumi.StringInput
}

func (ConfigsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*configsArgs)(nil)).Elem()
}

type ConfigsInput interface {
	pulumi.Input

	ToConfigsOutput() ConfigsOutput
	ToConfigsOutputWithContext(ctx context.Context) ConfigsOutput
}

func (*Configs) ElementType() reflect.Type {
	return reflect.TypeOf((**Configs)(nil)).Elem()
}

func (i *Configs) ToConfigsOutput() ConfigsOutput {
	return i.ToConfigsOutputWithContext(context.Background())
}

func (i *Configs) ToConfigsOutputWithContext(ctx context.Context) ConfigsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConfigsOutput)
}

type ConfigsOutput struct{ *pulumi.OutputState }

func (ConfigsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Configs)(nil)).Elem()
}

func (o ConfigsOutput) ToConfigsOutput() ConfigsOutput {
	return o
}

func (o ConfigsOutput) ToConfigsOutputWithContext(ctx context.Context) ConfigsOutput {
	return o
}

func (o ConfigsOutput) Config() ConfigPropertiesPtrOutput {
	return o.ApplyT(func(v *Configs) ConfigPropertiesPtrOutput { return v.Config }).(ConfigPropertiesPtrOutput)
}

// Identifier for the environment object.
func (o ConfigsOutput) Environment() pulumi.StringOutput {
	return o.ApplyT(func(v *Configs) pulumi.StringOutput { return v.Environment }).(pulumi.StringOutput)
}

// Name of the new branch config.
func (o ConfigsOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Configs) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Unique identifier for the project object.
func (o ConfigsOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *Configs) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ConfigsInput)(nil)).Elem(), &Configs{})
	pulumi.RegisterOutputType(ConfigsOutput{})
}