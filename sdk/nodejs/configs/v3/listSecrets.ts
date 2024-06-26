// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as enums from "../../types/enums";
import * as utilities from "../../utilities";

export function listSecrets(args?: ListSecretsArgs, opts?: pulumi.InvokeOptions): Promise<ListSecretsResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("doppler-native:configs/v3:listSecrets", {
    }, opts);
}

export interface ListSecretsArgs {
}

export interface ListSecretsResult {
    readonly items: outputs.configs.v3.ListSecretsProperties;
}
export function listSecretsOutput(opts?: pulumi.InvokeOptions): pulumi.Output<ListSecretsResult> {
    return pulumi.output(listSecrets(opts))
}
