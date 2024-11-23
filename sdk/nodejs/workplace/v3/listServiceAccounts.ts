// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as enums from "../../types/enums";
import * as utilities from "../../utilities";

export function listServiceAccounts(args?: ListServiceAccountsArgs, opts?: pulumi.InvokeOptions): Promise<outputs.workplace.v3.ListServiceAccountsProperties> {
    args = args || {};
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("doppler-native:workplace/v3:listServiceAccounts", {
    }, opts);
}

export interface ListServiceAccountsArgs {
}
export function listServiceAccountsOutput(opts?: pulumi.InvokeOutputOptions): pulumi.Output<outputs.workplace.v3.ListServiceAccountsProperties> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("doppler-native:workplace/v3:listServiceAccounts", {
    }, opts);
}

