// @ts-ignore
/* eslint-disable */
import { request } from '../utils'

/** auth login auth login by username and password, and return access token and refresh token. BTW, visitor id is required for distinguishing clients/devices. 返回值: Error POST /api/auth/login */
export async function authLogin(body: API.AuthLoginRequest, options?: { [key: string]: any }) {
  return request<API.ErrorModel>('/api/auth/login', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json'
    },
    data: body,
    ...(options || {})
  })
}

/** auth refresh auth refresh by refresh token, and return access token and refresh token. 返回值: Error GET /api/auth/refresh */
export async function authRefresh(
  // 叠加生成的Param类型 (非body参数swagger默认没有生成对象)
  params: API.authRefreshParams,
  options?: { [key: string]: any }
) {
  return request<API.ErrorModel>('/api/auth/refresh', {
    method: 'GET',
    params: {
      ...params
    },
    ...(options || {})
  })
}
