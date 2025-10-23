// @ts-ignore
/* eslint-disable */
import request from '@/utils/request'

/** Delete File Delete file by id. 返回值: Error DELETE /api/file/delete/${param0} */
export async function fileDeleteDelete(
  // 叠加生成的Param类型 (非body参数swagger默认没有生成对象)
  params: API.fileDeleteDeleteParams,
  options?: { [key: string]: any }
) {
  const { id: param0, ...queryParams } = params
  return request<API.ErrorModel>(`/api/file/delete/${param0}`, {
    method: 'DELETE',
    params: { ...queryParams },
    ...(options || {})
  })
}

/** Get File Detail Get file detail by id, and users can include associated entities by passing includes. 返回值: Error POST /api/file/detail/${param0} */
export async function filePostDetail(
  // 叠加生成的Param类型 (非body参数swagger默认没有生成对象)
  params: API.filePostDetailParams,
  body: any,
  options?: { [key: string]: any }
) {
  const { id: param0, ...queryParams } = params
  return request<API.ErrorModel>(`/api/file/detail/${param0}`, {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json'
    },
    params: { ...queryParams },
    data: body,
    ...(options || {})
  })
}

/** Query File List Query file and present results in list, and users can filter, sort and page data by passing criteria. 返回值: Error POST /api/file/query */
export async function filePostQuery(body: API.Criteria, options?: { [key: string]: any }) {
  return request<API.ErrorModel>('/api/file/query', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json'
    },
    data: body,
    ...(options || {})
  })
}

/** Recover File Recover file by id. 返回值: Error PUT /api/file/recover/${param0} */
export async function filePutRecover(
  // 叠加生成的Param类型 (非body参数swagger默认没有生成对象)
  params: API.filePutRecoverParams,
  options?: { [key: string]: any }
) {
  const { id: param0, ...queryParams } = params
  return request<API.ErrorModel>(`/api/file/recover/${param0}`, {
    method: 'PUT',
    params: { ...queryParams },
    ...(options || {})
  })
}

/** Get File Recycle List Get deleted file and present results in list, and users can filter, sort and page data by passing criteria. 返回值: Error POST /api/file/recycle */
export async function filePostRecycle(body: API.Criteria, options?: { [key: string]: any }) {
  return request<API.ErrorModel>('/api/file/recycle', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json'
    },
    data: body,
    ...(options || {})
  })
}

/** Save File Save file, if it is a new entity which is transient or id=0, it will be created, otherwise it will be updated, and only the exposed fields will be updated. 返回值: Error PUT /api/file/save */
export async function filePutSave(body: API.FileSave, options?: { [key: string]: any }) {
  return request<API.ErrorModel>('/api/file/save', {
    method: 'PUT',
    headers: {
      'Content-Type': 'application/json'
    },
    data: body,
    ...(options || {})
  })
}

/** Search File Search file by given keyword and fields, and limit the result count. 返回值: Error POST /api/file/search */
export async function filePostSearch(body: API.CrudSearchRequest, options?: { [key: string]: any }) {
  return request<API.ErrorModel>('/api/file/search', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json'
    },
    data: body,
    ...(options || {})
  })
}
