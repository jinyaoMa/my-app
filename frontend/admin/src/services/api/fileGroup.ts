// @ts-ignore
/* eslint-disable */
import { request } from '@umijs/max'

/** Delete File Group Delete file group by id. DELETE /api/file-group/delete/${param0} */
export async function fileGroupDeleteDelete(
  // 叠加生成的Param类型 (非body参数swagger默认没有生成对象)
  params: API.fileGroupDeleteDeleteParams,
  options?: { [key: string]: any }
) {
  const { id: param0, ...queryParams } = params
  return request<API.ResponseStringBody>(`/api/file-group/delete/${param0}`, {
    method: 'DELETE',
    params: { ...queryParams },
    ...(options || {})
  })
}

/** Get File Group Detail Get file group detail by id, and users can include associated entities by passing includes. POST /api/file-group/detail/${param0} */
export async function fileGroupPostDetail(
  // 叠加生成的Param类型 (非body参数swagger默认没有生成对象)
  params: API.fileGroupPostDetailParams,
  body: API.Include[],
  options?: { [key: string]: any }
) {
  const { id: param0, ...queryParams } = params
  return request<API.ResponseFileGroupDetailBody>(`/api/file-group/detail/${param0}`, {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json'
    },
    params: { ...queryParams },
    data: body,
    ...(options || {})
  })
}

/** Query File Group List Query file group and present results in list, and users can filter, sort and page data by passing criteria. POST /api/file-group/query */
export async function fileGroupPostQuery(body: API.Criteria, options?: { [key: string]: any }) {
  return request<API.ResponseListFileGroupItemBody>('/api/file-group/query', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json'
    },
    data: body,
    ...(options || {})
  })
}

/** Save File Group Save file group, if it is a new entity which is transient or id=0, it will be created, otherwise it will be updated, and only the exposed fields will be updated. PUT /api/file-group/save */
export async function fileGroupPutSave(
  // 叠加生成的Param类型 (非body参数swagger默认没有生成对象)
  params: API.fileGroupPutSaveParams,
  body: API.FileGroupSave,
  options?: { [key: string]: any }
) {
  return request<API.ResponseFileGroupSaveBody>('/api/file-group/save', {
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

/** Search File Group Search file group by given keyword and fields, and limit the result count. POST /api/file-group/search */
export async function fileGroupPostSearch(body: API.CrudSearchRequest, options?: { [key: string]: any }) {
  return request<API.ResponseListFileGroupSearchItemBody>('/api/file-group/search', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json'
    },
    data: body,
    ...(options || {})
  })
}
