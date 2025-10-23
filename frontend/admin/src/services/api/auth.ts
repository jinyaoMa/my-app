// @ts-ignore
/* eslint-disable */
import { request } from '@umijs/max'

/** Login Login by username and password, and return access token and refresh token. BTW, visitor id is required for distinguishing clients/devices. POST /api/auth/login */
export async function authPostLogin(body: API.AuthLoginRequest, options?: { [key: string]: any }) {
  return request<API.ResponseAuthLoginBody>('/api/auth/login', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json'
    },
    data: body,
    ...(options || {})
  })
}

/** Login Temporary Login by username and temporary password, and return access token and refresh token. BTW, visitor id is required for distinguishing clients/devices. POST /api/auth/login/temporary */
export async function authPostLoginTemporary(body: API.AuthLoginRequest, options?: { [key: string]: any }) {
  return request<API.ResponseAuthLoginBody>('/api/auth/login/temporary', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json'
    },
    data: body,
    ...(options || {})
  })
}

/** Refresh Refresh by refresh token, and return access token and refresh token. GET /api/auth/refresh/${param0} */
export async function authGetRefresh(
  // 叠加生成的Param类型 (非body参数swagger默认没有生成对象)
  params: API.authGetRefreshParams,
  options?: { [key: string]: any }
) {
  const { refresh: param0, ...queryParams } = params
  return request<API.ResponseAuthLoginBody>(`/api/auth/refresh/${param0}`, {
    method: 'GET',
    params: { ...queryParams },
    ...(options || {})
  })
}

/** Switch Role Switch current user role to specified role. GET /api/auth/role */
export async function authGetRole(
  // 叠加生成的Param类型 (非body参数swagger默认没有生成对象)
  params: API.authGetRoleParams,
  options?: { [key: string]: any }
) {
  return request<API.ResponseAuthLoginBody>('/api/auth/role', {
    method: 'GET',
    params: {
      ...params
    },
    ...(options || {})
  })
}
