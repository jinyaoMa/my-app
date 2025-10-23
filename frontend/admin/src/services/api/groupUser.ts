// @ts-ignore
/* eslint-disable */
import { request } from '@umijs/max'

/** Delete Group User Delete group user by id. DELETE /api/group-user/delete/${param0} */
export async function groupUserDeleteDelete(
  // 叠加生成的Param类型 (非body参数swagger默认没有生成对象)
  params: API.groupUserDeleteDeleteParams,
  options?: { [key: string]: any }
) {
  const { id: param0, ...queryParams } = params
  return request<API.ResponseStringBody>(`/api/group-user/delete/${param0}`, {
    method: 'DELETE',
    params: { ...queryParams },
    ...(options || {})
  })
}

/** Get Group User Detail Get group user detail by id, and users can include associated entities by passing includes. POST /api/group-user/detail/${param0} */
export async function groupUserPostDetail(
  // 叠加生成的Param类型 (非body参数swagger默认没有生成对象)
  params: API.groupUserPostDetailParams,
  body: API.Include[],
  options?: { [key: string]: any }
) {
  const { id: param0, ...queryParams } = params
  return request<API.ResponseGroupUserDetailBody>(`/api/group-user/detail/${param0}`, {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json'
    },
    params: { ...queryParams },
    data: body,
    ...(options || {})
  })
}

/** Query Group User List Query group user and present results in list, and users can filter, sort and page data by passing criteria. POST /api/group-user/query */
export async function groupUserPostQuery(body: API.Criteria, options?: { [key: string]: any }) {
  return request<API.ResponseListGroupUserItemBody>('/api/group-user/query', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json'
    },
    data: body,
    ...(options || {})
  })
}

/** Save Group User Save group user, if it is a new entity which is transient or id=0, it will be created, otherwise it will be updated, and only the exposed fields will be updated. PUT /api/group-user/save */
export async function groupUserPutSave(
  // 叠加生成的Param类型 (非body参数swagger默认没有生成对象)
  params: API.groupUserPutSaveParams,
  body: API.GroupUserSave,
  options?: { [key: string]: any }
) {
  return request<API.ResponseGroupUserSaveBody>('/api/group-user/save', {
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

/** Search Group User Search group user by given keyword and fields, and limit the result count. POST /api/group-user/search */
export async function groupUserPostSearch(body: API.CrudSearchRequest, options?: { [key: string]: any }) {
  return request<API.ResponseListGroupUserSearchItemBody>('/api/group-user/search', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json'
    },
    data: body,
    ...(options || {})
  })
}
