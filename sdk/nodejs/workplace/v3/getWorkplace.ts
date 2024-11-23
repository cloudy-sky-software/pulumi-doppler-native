// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as enums from "../../types/enums";
import * as utilities from "../../utilities";

export function getWorkplace(args?: GetWorkplaceArgs, opts?: pulumi.InvokeOptions): Promise<outputs.workplace.v3.GetWorkplaceProperties> {
    args = args || {};
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("doppler-native:workplace/v3:getWorkplace", {
    }, opts);
}

export interface GetWorkplaceArgs {
}
export function getWorkplaceOutput(opts?: pulumi.InvokeOutputOptions): pulumi.Output<outputs.workplace.v3.GetWorkplaceProperties> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("doppler-native:workplace/v3:getWorkplace", {
    }, opts);
}

