// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as enums from "../../types/enums";
import * as utilities from "../../utilities";

export function listConfigs(args?: ListConfigsArgs, opts?: pulumi.InvokeOptions): Promise<outputs.configs.v3.ListConfigsProperties> {
    args = args || {};
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("doppler-native:configs/v3:listConfigs", {
    }, opts);
}

export interface ListConfigsArgs {
}
export function listConfigsOutput(opts?: pulumi.InvokeOptions): pulumi.Output<outputs.configs.v3.ListConfigsProperties> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("doppler-native:configs/v3:listConfigs", {
    }, opts);
}

