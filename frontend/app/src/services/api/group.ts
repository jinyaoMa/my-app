// @ts-ignore
/* eslint-disable */
import request from '@/utils/request'

/** Delete Group Delete group by id. 返回值: Error DELETE /api/group/delete/${param0} */
export async function groupDeleteDelete(
  // 叠加生成的Param类型 (非body参数swagger默认没有生成对象)
  params: API.groupDeleteDeleteParams,
  options?: { [key: string]: any }
) {
  const { id: param0, ...queryParams } = params
  return request<API.ErrorModel>(`/api/group/delete/${param0}`, {
    method: 'DELETE',
    params: { ...queryParams },
    ...(options || {})
  })
}

/** Get Group Detail Get group detail by id, and users can include associated entities by passing includes. 返回值: Error POST /api/group/detail/${param0} */
export async function groupPostDetail(
  // 叠加生成的Param类型 (非body参数swagger默认没有生成对象)
  params: API.groupPostDetailParams,
  body: any,
  options?: { [key: string]: any }
) {
  const { id: param0, ...queryParams } = params
  return request<API.ErrorModel>(`/api/group/detail/${param0}`, {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json'
    },
    params: { ...queryParams },
    data: body,
    ...(options || {})
  })
}

/** Query Group List Query group and present results in list, and users can filter, sort and page data by passing criteria. 返回值: Error POST /api/group/query */
export async function groupPostQuery(body: API.Criteria, options?: { [key: string]: any }) {
  return request<API.ErrorModel>('/api/group/query', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json'
    },
    data: body,
    ...(options || {})
  })
}

/** Recover Group Recover group by id. 返回值: Error PUT /api/group/recover/${param0} */
export async function groupPutRecover(
  // 叠加生成的Param类型 (非body参数swagger默认没有生成对象)
  params: API.groupPutRecoverParams,
  options?: { [key: string]: any }
) {
  const { id: param0, ...queryParams } = params
  return request<API.ErrorModel>(`/api/group/recover/${param0}`, {
    method: 'PUT',
    params: { ...queryParams },
    ...(options || {})
  })
}

/** Get Group Recycle List Get deleted group and present results in list, and users can filter, sort and page data by passing criteria. 返回值: Error POST /api/group/recycle */
export async function groupPostRecycle(body: API.Criteria, options?: { [key: string]: any }) {
  return request<API.ErrorModel>('/api/group/recycle', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json'
    },
    data: body,
    ...(options || {})
  })
}

/** Save Group Save group, if it is a new entity which is transient or id=0, it will be created, otherwise it will be updated, and only the exposed fields will be updated. 返回值: Error PUT /api/group/save */
export async function groupPutSave(body: API.GroupSave, options?: { [key: string]: any }) {
  return request<API.ErrorModel>('/api/group/save', {
    method: 'PUT',
    headers: {
      'Content-Type': 'application/json'
    },
    data: body,
    ...(options || {})
  })
}

/** Search Group Search group by given keyword and fields, and limit the result count. 返回值: Error POST /api/group/search */
export async function groupPostSearch(body: API.CrudSearchRequest, options?: { [key: string]: any }) {
  return request<API.ErrorModel>('/api/group/search', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json'
    },
    data: body,
    ...(options || {})
  })
}
