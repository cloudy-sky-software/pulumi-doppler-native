// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../../utilities";

// Export members:
export { GetProjectArgs } from "./getProject";
export const getProject: typeof import("./getProject").getProject = null as any;
export const getProjectOutput: typeof import("./getProject").getProjectOutput = null as any;
utilities.lazyLoad(exports, ["getProject","getProjectOutput"], () => require("./getProject"));

export { GetProjectMemberArgs, GetProjectMemberOutputArgs } from "./getProjectMember";
export const getProjectMember: typeof import("./getProjectMember").getProjectMember = null as any;
export const getProjectMemberOutput: typeof import("./getProjectMember").getProjectMemberOutput = null as any;
utilities.lazyLoad(exports, ["getProjectMember","getProjectMemberOutput"], () => require("./getProjectMember"));

export { GetProjectRoleArgs, GetProjectRoleOutputArgs } from "./getProjectRole";
export const getProjectRole: typeof import("./getProjectRole").getProjectRole = null as any;
export const getProjectRoleOutput: typeof import("./getProjectRole").getProjectRoleOutput = null as any;
utilities.lazyLoad(exports, ["getProjectRole","getProjectRoleOutput"], () => require("./getProjectRole"));

export { ListProjectMembersArgs } from "./listProjectMembers";
export const listProjectMembers: typeof import("./listProjectMembers").listProjectMembers = null as any;
export const listProjectMembersOutput: typeof import("./listProjectMembers").listProjectMembersOutput = null as any;
utilities.lazyLoad(exports, ["listProjectMembers","listProjectMembersOutput"], () => require("./listProjectMembers"));

export { ListProjectRolesArgs } from "./listProjectRoles";
export const listProjectRoles: typeof import("./listProjectRoles").listProjectRoles = null as any;
export const listProjectRolesOutput: typeof import("./listProjectRoles").listProjectRolesOutput = null as any;
utilities.lazyLoad(exports, ["listProjectRoles","listProjectRolesOutput"], () => require("./listProjectRoles"));

export { ListProjectRolesPermissionsArgs } from "./listProjectRolesPermissions";
export const listProjectRolesPermissions: typeof import("./listProjectRolesPermissions").listProjectRolesPermissions = null as any;
export const listProjectRolesPermissionsOutput: typeof import("./listProjectRolesPermissions").listProjectRolesPermissionsOutput = null as any;
utilities.lazyLoad(exports, ["listProjectRolesPermissions","listProjectRolesPermissionsOutput"], () => require("./listProjectRolesPermissions"));

export { ListProjectsArgs } from "./listProjects";
export const listProjects: typeof import("./listProjects").listProjects = null as any;
export const listProjectsOutput: typeof import("./listProjects").listProjectsOutput = null as any;
utilities.lazyLoad(exports, ["listProjects","listProjectsOutput"], () => require("./listProjects"));

export { ProjectArgs } from "./project";
export type Project = import("./project").Project;
export const Project: typeof import("./project").Project = null as any;
utilities.lazyLoad(exports, ["Project"], () => require("./project"));

export { ProjectMemberArgs } from "./projectMember";
export type ProjectMember = import("./projectMember").ProjectMember;
export const ProjectMember: typeof import("./projectMember").ProjectMember = null as any;
utilities.lazyLoad(exports, ["ProjectMember"], () => require("./projectMember"));

export { ProjectRoleArgs } from "./projectRole";
export type ProjectRole = import("./projectRole").ProjectRole;
export const ProjectRole: typeof import("./projectRole").ProjectRole = null as any;
utilities.lazyLoad(exports, ["ProjectRole"], () => require("./projectRole"));

export { ProjectsUpdateArgs } from "./projectsUpdate";
export type ProjectsUpdate = import("./projectsUpdate").ProjectsUpdate;
export const ProjectsUpdate: typeof import("./projectsUpdate").ProjectsUpdate = null as any;
utilities.lazyLoad(exports, ["ProjectsUpdate"], () => require("./projectsUpdate"));


// Export enums:
export * from "../../types/enums/projects/v3";

const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "doppler-native:projects/v3:Project":
                return new Project(name, <any>undefined, { urn })
            case "doppler-native:projects/v3:ProjectMember":
                return new ProjectMember(name, <any>undefined, { urn })
            case "doppler-native:projects/v3:ProjectRole":
                return new ProjectRole(name, <any>undefined, { urn })
            case "doppler-native:projects/v3:ProjectsUpdate":
                return new ProjectsUpdate(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("doppler-native", "projects/v3", _module)
