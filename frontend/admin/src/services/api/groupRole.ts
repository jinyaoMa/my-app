// @ts-ignore
/* eslint-disable */
import { request } from '@umijs/max'

/** Delete Group Role Delete group role by id. DELETE /api/group-role/delete/${param0} */
export async function groupRoleDeleteDelete(
  // 叠加生成的Param类型 (非body参数swagger默认没有生成对象)
  params: API.groupRoleDeleteDeleteParams,
  options?: { [key: string]: any }
) {
  const { id: param0, ...queryParams } = params
  return request<API.ResponseStringBody>(`/api/group-role/delete/${param0}`, {
    method: 'DELETE',
    params: { ...queryParams },
    ...(options || {})
  })
}

/** Get Group Role Detail Get group role detail by id, and users can include associated entities by passing includes. POST /api/group-role/detail/${param0} */
export async function groupRolePostDetail(
  // 叠加生成的Param类型 (非body参数swagger默认没有生成对象)
  params: API.groupRolePostDetailParams,
  body: API.Include[],
  options?: { [key: string]: any }
) {
  const { id: param0, ...queryParams } = params
  return request<API.ResponseGroupRoleDetailBody>(`/api/group-role/detail/${param0}`, {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json'
    },
    params: { ...queryParams },
    data: body,
    ...(options || {})
  })
}

/** Query Group Role List Query group role and present results in list, and users can filter, sort and page data by passing criteria. POST /api/group-role/query */
export async function groupRolePostQuery(body: API.Criteria, options?: { [key: string]: any }) {
  return request<API.ResponseListGroupRoleItemBody>('/api/group-role/query', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json'
    },
    data: body,
    ...(options || {})
  })
}

/** Save Group Role Save group role, if it is a new entity which is transient or id=0, it will be created, otherwise it will be updated, and only the exposed fields will be updated. PUT /api/group-role/save */
export async function groupRolePutSave(
  // 叠加生成的Param类型 (非body参数swagger默认没有生成对象)
  params: API.groupRolePutSaveParams,
  body: API.GroupRoleSave,
  options?: { [key: string]: any }
) {
  return request<API.ResponseGroupRoleSaveBody>('/api/group-role/save', {
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

/** Search Group Role Search group role by given keyword and fields, and limit the result count. POST /api/group-role/search */
export async function groupRolePostSearch(body: API.CrudSearchRequest, options?: { [key: string]: any }) {
  return request<API.ResponseListGroupRoleSearchItemBody>('/api/group-role/search', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json'
    },
    data: body,
    ...(options || {})
  })
}
