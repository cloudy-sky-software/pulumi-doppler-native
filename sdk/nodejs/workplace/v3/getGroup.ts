// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as enums from "../../types/enums";
import * as utilities from "../../utilities";

export function getGroup(args: GetGroupArgs, opts?: pulumi.InvokeOptions): Promise<GetGroupResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("doppler-native:workplace/v3:getGroup", {
        "slug": args.slug,
    }, opts);
}

export interface GetGroupArgs {
    /**
     * The group's slug
     */
    slug: string;
}

export interface GetGroupResult {
    readonly items: outputs.workplace.v3.GetGroupProperties;
}
export function getGroupOutput(args: GetGroupOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetGroupResult> {
    return pulumi.output(args).apply((a: any) => getGroup(a, opts))
}

export interface GetGroupOutputArgs {
    /**
     * The group's slug
     */
    slug: pulumi.Input<string>;
}