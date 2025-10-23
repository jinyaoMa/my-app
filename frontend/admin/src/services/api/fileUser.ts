// @ts-ignore
/* eslint-disable */
import { request } from '@umijs/max'

/** Delete File User Delete file user by id. DELETE /api/file-user/delete/${param0} */
export async function fileUserDeleteDelete(
  // 叠加生成的Param类型 (非body参数swagger默认没有生成对象)
  params: API.fileUserDeleteDeleteParams,
  options?: { [key: string]: any }
) {
  const { id: param0, ...queryParams } = params
  return request<API.ResponseStringBody>(`/api/file-user/delete/${param0}`, {
    method: 'DELETE',
    params: { ...queryParams },
    ...(options || {})
  })
}

/** Get File User Detail Get file user detail by id, and users can include associated entities by passing includes. POST /api/file-user/detail/${param0} */
export async function fileUserPostDetail(
  // 叠加生成的Param类型 (非body参数swagger默认没有生成对象)
  params: API.fileUserPostDetailParams,
  body: API.Include[],
  options?: { [key: string]: any }
) {
  const { id: param0, ...queryParams } = params
  return request<API.ResponseFileUserDetailBody>(`/api/file-user/detail/${param0}`, {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json'
    },
    params: { ...queryParams },
    data: body,
    ...(options || {})
  })
}

/** Query File User List Query file user and present results in list, and users can filter, sort and page data by passing criteria. POST /api/file-user/query */
export async function fileUserPostQuery(body: API.Criteria, options?: { [key: string]: any }) {
  return request<API.ResponseListFileUserItemBody>('/api/file-user/query', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json'
    },
    data: body,
    ...(options || {})
  })
}

/** Save File User Save file user, if it is a new entity which is transient or id=0, it will be created, otherwise it will be updated, and only the exposed fields will be updated. PUT /api/file-user/save */
export async function fileUserPutSave(
  // 叠加生成的Param类型 (非body参数swagger默认没有生成对象)
  params: API.fileUserPutSaveParams,
  body: API.FileUserSave,
  options?: { [key: string]: any }
) {
  return request<API.ResponseFileUserSaveBody>('/api/file-user/save', {
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

/** Search File User Search file user by given keyword and fields, and limit the result count. POST /api/file-user/search */
export async function fileUserPostSearch(body: API.CrudSearchRequest, options?: { [key: string]: any }) {
  return request<API.ResponseListFileUserSearchItemBody>('/api/file-user/search', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json'
    },
    data: body,
    ...(options || {})
  })
}
