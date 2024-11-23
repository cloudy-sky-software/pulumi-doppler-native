// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as enums from "../../types/enums";
import * as utilities from "../../utilities";

export function listProjectMembers(args?: ListProjectMembersArgs, opts?: pulumi.InvokeOptions): Promise<outputs.projects.v3.ListProjectMembersProperties> {
    args = args || {};
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("doppler-native:projects/v3:listProjectMembers", {
    }, opts);
}

export interface ListProjectMembersArgs {
}
export function listProjectMembersOutput(opts?: pulumi.InvokeOutputOptions): pulumi.Output<outputs.projects.v3.ListProjectMembersProperties> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("doppler-native:projects/v3:listProjectMembers", {
    }, opts);
}

