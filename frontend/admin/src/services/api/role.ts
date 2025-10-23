// @ts-ignore
/* eslint-disable */
import { request } from '@umijs/max'

/** Batch Delete Role Batch delete role by ids. DELETE /api/role/batch/delete */
export async function roleDeleteBatchDelete(body: string[], options?: { [key: string]: any }) {
  return request<API.ResponseListStringBody>('/api/role/batch/delete', {
    method: 'DELETE',
    headers: {
      'Content-Type': 'application/json'
    },
    data: body,
    ...(options || {})
  })
}

/** Batch Recover Role Batch recover role by ids. PUT /api/role/batch/recover */
export async function rolePutBatchRecover(body: string[], options?: { [key: string]: any }) {
  return request<API.ResponseListStringBody>('/api/role/batch/recover', {
    method: 'PUT',
    headers: {
      'Content-Type': 'application/json'
    },
    data: body,
    ...(options || {})
  })
}

/** Delete Role Delete role by id. DELETE /api/role/delete/${param0} */
export async function roleDeleteDelete(
  // 叠加生成的Param类型 (非body参数swagger默认没有生成对象)
  params: API.roleDeleteDeleteParams,
  options?: { [key: string]: any }
) {
  const { id: param0, ...queryParams } = params
  return request<API.ResponseStringBody>(`/api/role/delete/${param0}`, {
    method: 'DELETE',
    params: { ...queryParams },
    ...(options || {})
  })
}

/** Get Role Detail Get role detail by id, and users can include associated entities by passing includes. POST /api/role/detail/${param0} */
export async function rolePostDetail(
  // 叠加生成的Param类型 (非body参数swagger默认没有生成对象)
  params: API.rolePostDetailParams,
  body: API.Include[],
  options?: { [key: string]: any }
) {
  const { id: param0, ...queryParams } = params
  return request<API.ResponseRoleDetailBody>(`/api/role/detail/${param0}`, {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json'
    },
    params: { ...queryParams },
    data: body,
    ...(options || {})
  })
}

/** Query Role List Query role and present results in list, and users can filter, sort and page data by passing criteria. POST /api/role/query */
export async function rolePostQuery(body: API.Criteria, options?: { [key: string]: any }) {
  return request<API.ResponseListRoleItemBody>('/api/role/query', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json'
    },
    data: body,
    ...(options || {})
  })
}

/** Recover Role Recover role by id. PUT /api/role/recover/${param0} */
export async function rolePutRecover(
  // 叠加生成的Param类型 (非body参数swagger默认没有生成对象)
  params: API.rolePutRecoverParams,
  options?: { [key: string]: any }
) {
  const { id: param0, ...queryParams } = params
  return request<API.ResponseStringBody>(`/api/role/recover/${param0}`, {
    method: 'PUT',
    params: { ...queryParams },
    ...(options || {})
  })
}

/** Get Role Recycle List Get deleted role and present results in list, and users can filter, sort and page data by passing criteria. POST /api/role/recycle */
export async function rolePostRecycle(body: API.Criteria, options?: { [key: string]: any }) {
  return request<API.ResponseListRoleItemBody>('/api/role/recycle', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json'
    },
    data: body,
    ...(options || {})
  })
}

/** Save Role Save role, if it is a new entity which is transient or id=0, it will be created, otherwise it will be updated, and only the exposed fields will be updated. PUT /api/role/save */
export async function rolePutSave(
  // 叠加生成的Param类型 (非body参数swagger默认没有生成对象)
  params: API.rolePutSaveParams,
  body: API.RoleSave,
  options?: { [key: string]: any }
) {
  return request<API.ResponseRoleSaveBody>('/api/role/save', {
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

/** Search Role Search role by given keyword and fields, and limit the result count. POST /api/role/search */
export async function rolePostSearch(body: API.CrudSearchRequest, options?: { [key: string]: any }) {
  return request<API.ResponseListRoleSearchItemBody>('/api/role/search', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json'
    },
    data: body,
    ...(options || {})
  })
}
