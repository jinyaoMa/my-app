// @ts-ignore
/* eslint-disable */
import { request } from '../utils'

/** detail group Detail group by id, and users can include associated entities by passing includes. 返回值: Error GET /api/group */
export async function groupDetail(
  // 叠加生成的Param类型 (非body参数swagger默认没有生成对象)
  params: API.groupDetailParams,
  options?: { [key: string]: any }
) {
  return request<API.ErrorModel>('/api/group', {
    method: 'GET',
    params: {
      ...params
    },
    ...(options || {})
  })
}

/** save group Save group, if it is a new entity which is transient or id=0, it will be created, otherwise it will be updated, and only the exposed fields will be updated. 返回值: Error POST /api/group */
export async function groupSave(body: API.GroupSave, options?: { [key: string]: any }) {
  return request<API.ErrorModel>('/api/group', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json'
    },
    data: body,
    ...(options || {})
  })
}

/** delete group Delete group by id. 返回值: Error DELETE /api/group */
export async function groupDelete(
  // 叠加生成的Param类型 (非body参数swagger默认没有生成对象)
  params: API.groupDeleteParams,
  options?: { [key: string]: any }
) {
  return request<API.ErrorModel>('/api/group', {
    method: 'DELETE',
    params: {
      ...params
    },
    ...(options || {})
  })
}

/** query group list Query group and present results in list, and users can filter, sort and page data by passing criteria. 返回值: Error POST /api/group/query */
export async function groupQuery(body: API.Criteria, options?: { [key: string]: any }) {
  return request<API.ErrorModel>('/api/group/query', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json'
    },
    data: body,
    ...(options || {})
  })
}
