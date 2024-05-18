// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";

import * as utilities from "../utilities";

export namespace configs {
    export namespace v3 {
        export interface ChangeRequestsItemPropertiesArgs {
            /**
             * The name of the secret.
             */
            name: pulumi.Input<string>;
            /**
             * The original name of the secret. Use `null` (an actual `null`, not the string `null`) or omit this parameter for new secrets. If it differs from `name` then a rename is inferred.
             */
            originalName: pulumi.Input<string>;
            /**
             * The value you expect the secret to have before `name` is applied. If specified, the request will only be processed if the provided value matches what's found in Doppler.
             */
            originalValue?: pulumi.Input<string>;
            /**
             * The valueType you expect the secret to have before `valueType` is applied. If specified, the request will only be processed if the provided valueType matches what's found in Doppler.
             */
            originalValueType?: pulumi.Input<inputs.configs.v3.ChangeRequestsItemPropertiesOriginalValueTypePropertiesArgs>;
            /**
             * Must be set to either `masked`, `unmasked`, or `restricted`. The visibility you expect the secret to have before `visibility` is applied. If specified, the request will only be processed if the provided visibility matches what's found in Doppler.
             */
            originalVisibility?: pulumi.Input<string>;
            /**
             * Defaults to `false`. Can only be set to `true` if the config being updated is a branch config and there is a secret with the same name in the root config. In this case, the branch secret will inherit the value and visibility type from the root secret.
             */
            shouldConverge?: pulumi.Input<boolean>;
            /**
             * Defaults to `false`. If set to `true`, will delete the secret matching the `name` field.
             */
            shouldDelete?: pulumi.Input<boolean>;
            /**
             * Defaults to `false`. Can only be set to `true` if the config being updated is a branch config. If set to `true`, the provided secret will be set in both the branch config as well as the root config in that environment.
             */
            shouldPromote?: pulumi.Input<boolean>;
            /**
             * The value the secret should have. Use `null` (an actual `null`, not the string `null`) to leave the existing secret value unchanged.
             */
            value: pulumi.Input<string>;
            /**
             * The default valueType (string) will result in no validations being applied.
             */
            valueType?: pulumi.Input<inputs.configs.v3.ChangeRequestsItemPropertiesValueTypePropertiesArgs>;
            /**
             * Must be set to either `masked`, `unmasked`, or `restricted`.
             */
            visibility?: pulumi.Input<string>;
        }

        /**
         * The valueType you expect the secret to have before `valueType` is applied. If specified, the request will only be processed if the provided valueType matches what's found in Doppler.
         */
        export interface ChangeRequestsItemPropertiesOriginalValueTypePropertiesArgs {
            type?: pulumi.Input<enums.configs.v3.SecretsUpdateChangeRequestsItemPropertiesOriginalValueTypePropertiesType>;
        }

        /**
         * The default valueType (string) will result in no validations being applied.
         */
        export interface ChangeRequestsItemPropertiesValueTypePropertiesArgs {
            type?: pulumi.Input<enums.configs.v3.SecretsUpdateChangeRequestsItemPropertiesValueTypePropertiesType>;
        }

        export interface SecretsPropertiesArgs {
            algolia?: pulumi.Input<inputs.configs.v3.SecretsPropertiesAlgoliaPropertiesArgs>;
            database?: pulumi.Input<inputs.configs.v3.SecretsPropertiesDatabasePropertiesArgs>;
            stripe?: pulumi.Input<inputs.configs.v3.SecretsPropertiesStripePropertiesArgs>;
        }

        export interface SecretsPropertiesAlgoliaPropertiesArgs {
            computed?: pulumi.Input<string>;
            note?: pulumi.Input<string>;
            raw?: pulumi.Input<string>;
        }

        export interface SecretsPropertiesDatabasePropertiesArgs {
            computed?: pulumi.Input<string>;
            note?: pulumi.Input<string>;
            raw?: pulumi.Input<string>;
        }

        export interface SecretsPropertiesStripePropertiesArgs {
            computed?: pulumi.Input<string>;
            note?: pulumi.Input<string>;
            raw?: pulumi.Input<string>;
        }

    }
}

export namespace environments {
    export namespace v3 {
    }
}

export namespace integrations {
    export namespace v3 {
    }
}

export namespace me {
    export namespace v3 {
    }
}

export namespace projects {
    export namespace v3 {
    }
}

export namespace webhooks {
    export namespace v3 {
        export interface AuthenticationPropertiesArgs {
            /**
             * Used when type = Basic
             */
            password?: pulumi.Input<string>;
            /**
             * Used when type = Bearer
             */
            token?: pulumi.Input<string>;
            type?: pulumi.Input<enums.webhooks.v3.WebhooksAuthenticationPropertiesType>;
            /**
             * Used when type = Basic
             */
            username?: pulumi.Input<string>;
        }
    }
}

export namespace workplace {
    export namespace v3 {
        /**
         * You may provide an identifier OR permissions, but not both
         */
        export interface WorkplaceRolePropertiesArgs {
            /**
             * Identifier of an existing workplace role
             */
            identifier?: pulumi.Input<string>;
            /**
             * Workplace permissions to grant
             */
            permissions?: pulumi.Input<pulumi.Input<string>[]>;
        }
    }
}