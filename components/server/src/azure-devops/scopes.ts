/**
 * Copyright (c) 2020 Gitpod GmbH. All rights reserved.
 * Licensed under the GNU Affero General Public License (AGPL).
 * See License-AGPL.txt in the project root for license information.
 */


export namespace AzureDevOpsScope {
    export const PROFILE = "vso.profile";
    export const CODE = "vso.code";

    export const All = [PROFILE, CODE];
    export const Requirements = {
        /**
         * Minimal required permission.
         * GitHub's API is not restricted any further.
         */
        DEFAULT: [PROFILE, CODE],
    }
}
