// @ts-ignore
/* eslint-disable */
import request from '@/utils/request'

/** Delete File Group Delete file group by id. 返回值: Error DELETE /api/file-group/delete/${param0} */
export async function fileGroupDeleteDelete(
  // 叠加生成的Param类型 (非body参数swagger默认没有生成对象)
  params: API.fileGroupDeleteDeleteParams,
  options?: { [key: string]: any }
) {
  const { id: param0, ...queryParams } = params
  return request<API.ErrorModel>(`/api/file-group/delete/${param0}`, {
    method: 'DELETE',
    params: { ...queryParams },
    ...(options || {})
  })
}

/** Get File Group Detail Get file group detail by id, and users can include associated entities by passing includes. 返回值: Error POST /api/file-group/detail/${param0} */
export async function fileGroupPostDetail(
  // 叠加生成的Param类型 (非body参数swagger默认没有生成对象)
  params: API.fileGroupPostDetailParams,
  body: any,
  options?: { [key: string]: any }
) {
  const { id: param0, ...queryParams } = params
  return request<API.ErrorModel>(`/api/file-group/detail/${param0}`, {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json'
    },
    params: { ...queryParams },
    data: body,
    ...(options || {})
  })
}

/** Query File Group List Query file group and present results in list, and users can filter, sort and page data by passing criteria. 返回值: Error POST /api/file-group/query */
export async function fileGroupPostQuery(body: API.Criteria, options?: { [key: string]: any }) {
  return request<API.ErrorModel>('/api/file-group/query', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json'
    },
    data: body,
    ...(options || {})
  })
}

/** Save File Group Save file group, if it is a new entity which is transient or id=0, it will be created, otherwise it will be updated, and only the exposed fields will be updated. 返回值: Error PUT /api/file-group/save */
export async function fileGroupPutSave(body: API.FileGroupSave, options?: { [key: string]: any }) {
  return request<API.ErrorModel>('/api/file-group/save', {
    method: 'PUT',
    headers: {
      'Content-Type': 'application/json'
    },
    data: body,
    ...(options || {})
  })
}

/** Search File Group Search file group by given keyword and fields, and limit the result count. 返回值: Error POST /api/file-group/search */
export async function fileGroupPostSearch(body: API.CrudSearchRequest, options?: { [key: string]: any }) {
  return request<API.ErrorModel>('/api/file-group/search', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json'
    },
    data: body,
    ...(options || {})
  })
}
