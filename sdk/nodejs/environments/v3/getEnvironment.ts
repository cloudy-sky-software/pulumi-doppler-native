// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as enums from "../../types/enums";
import * as utilities from "../../utilities";

export function getEnvironment(args?: GetEnvironmentArgs, opts?: pulumi.InvokeOptions): Promise<outputs.environments.v3.GetEnvironmentProperties> {
    args = args || {};
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("doppler-native:environments/v3:getEnvironment", {
    }, opts);
}

export interface GetEnvironmentArgs {
}
export function getEnvironmentOutput(opts?: pulumi.InvokeOutputOptions): pulumi.Output<outputs.environments.v3.GetEnvironmentProperties> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("doppler-native:environments/v3:getEnvironment", {
    }, opts);
}

