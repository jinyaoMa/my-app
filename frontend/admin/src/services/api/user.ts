// @ts-ignore
/* eslint-disable */
import { request } from '@umijs/max'

/** Batch Delete User Batch delete user by ids. DELETE /api/user/batch/delete */
export async function userDeleteBatchDelete(body: string[], options?: { [key: string]: any }) {
  return request<API.ResponseListStringBody>('/api/user/batch/delete', {
    method: 'DELETE',
    headers: {
      'Content-Type': 'application/json'
    },
    data: body,
    ...(options || {})
  })
}

/** Batch Recover User Batch recover user by ids. PUT /api/user/batch/recover */
export async function userPutBatchRecover(body: string[], options?: { [key: string]: any }) {
  return request<API.ResponseListStringBody>('/api/user/batch/recover', {
    method: 'PUT',
    headers: {
      'Content-Type': 'application/json'
    },
    data: body,
    ...(options || {})
  })
}

/** Get Current User Get current user. GET /api/user/current */
export async function userGetCurrent(options?: { [key: string]: any }) {
  return request<API.ResponseCurrentUserBody>('/api/user/current', {
    method: 'GET',
    ...(options || {})
  })
}

/** Delete User Delete user by id. DELETE /api/user/delete/${param0} */
export async function userDeleteDelete(
  // 叠加生成的Param类型 (非body参数swagger默认没有生成对象)
  params: API.userDeleteDeleteParams,
  options?: { [key: string]: any }
) {
  const { id: param0, ...queryParams } = params
  return request<API.ResponseStringBody>(`/api/user/delete/${param0}`, {
    method: 'DELETE',
    params: { ...queryParams },
    ...(options || {})
  })
}

/** Get User Detail Get user detail by id, and users can include associated entities by passing includes. POST /api/user/detail/${param0} */
export async function userPostDetail(
  // 叠加生成的Param类型 (非body参数swagger默认没有生成对象)
  params: API.userPostDetailParams,
  body: API.Include[],
  options?: { [key: string]: any }
) {
  const { id: param0, ...queryParams } = params
  return request<API.ResponseUserDetailBody>(`/api/user/detail/${param0}`, {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json'
    },
    params: { ...queryParams },
    data: body,
    ...(options || {})
  })
}

/** Query User List Query user and present results in list, and users can filter, sort and page data by passing criteria. POST /api/user/query */
export async function userPostQuery(body: API.Criteria, options?: { [key: string]: any }) {
  return request<API.ResponseListUserItemBody>('/api/user/query', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json'
    },
    data: body,
    ...(options || {})
  })
}

/** Recover User Recover user by id. PUT /api/user/recover/${param0} */
export async function userPutRecover(
  // 叠加生成的Param类型 (非body参数swagger默认没有生成对象)
  params: API.userPutRecoverParams,
  options?: { [key: string]: any }
) {
  const { id: param0, ...queryParams } = params
  return request<API.ResponseStringBody>(`/api/user/recover/${param0}`, {
    method: 'PUT',
    params: { ...queryParams },
    ...(options || {})
  })
}

/** Get User Recycle List Get deleted user and present results in list, and users can filter, sort and page data by passing criteria. POST /api/user/recycle */
export async function userPostRecycle(body: API.Criteria, options?: { [key: string]: any }) {
  return request<API.ResponseListUserItemBody>('/api/user/recycle', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json'
    },
    data: body,
    ...(options || {})
  })
}

/** Save User Save user, if it is a new entity which is transient or id=0, it will be created, otherwise it will be updated, and only the exposed fields will be updated. PUT /api/user/save */
export async function userPutSave(
  // 叠加生成的Param类型 (非body参数swagger默认没有生成对象)
  params: API.userPutSaveParams,
  body: API.UserSave,
  options?: { [key: string]: any }
) {
  return request<API.ResponseUserSaveBody>('/api/user/save', {
    method: 'PUT',
    headers: {
      'Content-Type': 'application/json'
    },
    params: {
      ...params
    },
    data: body,
    ...(options || {})
  })
}

/** Search User Search user by given keyword and fields, and limit the result count. POST /api/user/search */
export async function userPostSearch(body: API.CrudSearchRequest, options?: { [key: string]: any }) {
  return request<API.ResponseListUserSearchItemBody>('/api/user/search', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json'
    },
    data: body,
    ...(options || {})
  })
}
