// @ts-check
// 此文件由 wails3 generate bindings 自动生成
// 开发时请运行 wails3 dev 或 wails3 generate bindings 来更新

import {Call as $Call} from "/wails/runtime.js";

/**
 * @returns {Promise<{hostname: string, os: string, arch: string, numCPU: number}>}
 */
export function GetSystemInfo() {
    return $Call.ByName("cleanC/services/system.SystemService.GetSystemInfo");
}

/**
 * @returns {Promise<string>}
 */
export function GetHostname() {
    return $Call.ByName("cleanC/services/system.SystemService.GetHostname");
}

/**
 * @returns {Promise<string>}
 */
export function GetOS() {
    return $Call.ByName("cleanC/services/system.SystemService.GetOS");
}

/**
 * @returns {Promise<string>}
 */
export function GetArch() {
    return $Call.ByName("cleanC/services/system.SystemService.GetArch");
}

/**
 * @returns {Promise<number>}
 */
export function GetNumCPU() {
    return $Call.ByName("cleanC/services/system.SystemService.GetNumCPU");
}
