// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as enums from "../../types/enums";
import * as utilities from "../../utilities";

export function getServiceAccountToken(args: GetServiceAccountTokenArgs, opts?: pulumi.InvokeOptions): Promise<outputs.workplace.v3.GetServiceAccountTokenProperties> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("doppler-native:workplace/v3:getServiceAccountToken", {
        "apiToken": args.apiToken,
        "serviceAccount": args.serviceAccount,
    }, opts);
}

export interface GetServiceAccountTokenArgs {
    /**
     * Slug of the API token
     */
    apiToken: string;
    /**
     * Slug of the service account
     */
    serviceAccount: string;
}
export function getServiceAccountTokenOutput(args: GetServiceAccountTokenOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<outputs.workplace.v3.GetServiceAccountTokenProperties> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("doppler-native:workplace/v3:getServiceAccountToken", {
        "apiToken": args.apiToken,
        "serviceAccount": args.serviceAccount,
    }, opts);
}

export interface GetServiceAccountTokenOutputArgs {
    /**
     * Slug of the API token
     */
    apiToken: pulumi.Input<string>;
    /**
     * Slug of the service account
     */
    serviceAccount: pulumi.Input<string>;
}
