// @ts-ignore
/* eslint-disable */
import request from '@/utils/request'

/** Delete Role Delete role by id. 返回值: Error DELETE /api/role/delete/${param0} */
export async function roleDeleteDelete(
  // 叠加生成的Param类型 (非body参数swagger默认没有生成对象)
  params: API.roleDeleteDeleteParams,
  options?: { [key: string]: any }
) {
  const { id: param0, ...queryParams } = params
  return request<API.ErrorModel>(`/api/role/delete/${param0}`, {
    method: 'DELETE',
    params: { ...queryParams },
    ...(options || {})
  })
}

/** Get Role Detail Get role detail by id, and users can include associated entities by passing includes. 返回值: Error POST /api/role/detail/${param0} */
export async function rolePostDetail(
  // 叠加生成的Param类型 (非body参数swagger默认没有生成对象)
  params: API.rolePostDetailParams,
  body: any,
  options?: { [key: string]: any }
) {
  const { id: param0, ...queryParams } = params
  return request<API.ErrorModel>(`/api/role/detail/${param0}`, {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json'
    },
    params: { ...queryParams },
    data: body,
    ...(options || {})
  })
}

/** Query Role List Query role and present results in list, and users can filter, sort and page data by passing criteria. 返回值: Error POST /api/role/query */
export async function rolePostQuery(body: API.Criteria, options?: { [key: string]: any }) {
  return request<API.ErrorModel>('/api/role/query', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json'
    },
    data: body,
    ...(options || {})
  })
}

/** Recover Role Recover role by id. 返回值: Error PUT /api/role/recover/${param0} */
export async function rolePutRecover(
  // 叠加生成的Param类型 (非body参数swagger默认没有生成对象)
  params: API.rolePutRecoverParams,
  options?: { [key: string]: any }
) {
  const { id: param0, ...queryParams } = params
  return request<API.ErrorModel>(`/api/role/recover/${param0}`, {
    method: 'PUT',
    params: { ...queryParams },
    ...(options || {})
  })
}

/** Get Role Recycle List Get deleted role and present results in list, and users can filter, sort and page data by passing criteria. 返回值: Error POST /api/role/recycle */
export async function rolePostRecycle(body: API.Criteria, options?: { [key: string]: any }) {
  return request<API.ErrorModel>('/api/role/recycle', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json'
    },
    data: body,
    ...(options || {})
  })
}

/** Save Role Save role, if it is a new entity which is transient or id=0, it will be created, otherwise it will be updated, and only the exposed fields will be updated. 返回值: Error PUT /api/role/save */
export async function rolePutSave(body: API.RoleSave, options?: { [key: string]: any }) {
  return request<API.ErrorModel>('/api/role/save', {
    method: 'PUT',
    headers: {
      'Content-Type': 'application/json'
    },
    data: body,
    ...(options || {})
  })
}

/** Search Role Search role by given keyword and fields, and limit the result count. 返回值: Error POST /api/role/search */
export async function rolePostSearch(body: API.CrudSearchRequest, options?: { [key: string]: any }) {
  return request<API.ErrorModel>('/api/role/search', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json'
    },
    data: body,
    ...(options || {})
  })
}
