// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as enums from "../../types/enums";
import * as utilities from "../../utilities";

export function getSecretsName(args?: GetSecretsNameArgs, opts?: pulumi.InvokeOptions): Promise<outputs.configs.v3.GetSecretsNameProperties> {
    args = args || {};
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("doppler-native:configs/v3:getSecretsName", {
    }, opts);
}

export interface GetSecretsNameArgs {
}
export function getSecretsNameOutput(opts?: pulumi.InvokeOutputOptions): pulumi.Output<outputs.configs.v3.GetSecretsNameProperties> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("doppler-native:configs/v3:getSecretsName", {
    }, opts);
}

