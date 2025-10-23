import type { ActionType } from '@ant-design/pro-components'
import React, { useRef, useState } from 'react'
import { QueryPage } from '@/components/Query'
import api from '@/services/api'
import { Space, Tag } from 'antd'
import { useIntl } from '@umijs/max'

type PermissionQueryParams = {
  code?: string
  name?: string
  description?: string
}

const Query: React.FC = () => {
  const intl = useIntl()
  const actionRef = useRef<ActionType>(null)
  const [selectedRows, setSelectedRows] = useState<API.PermissionItem[]>([])

  return (
    <QueryPage<API.PermissionItem, PermissionQueryParams>
      name="permission"
      actionRef={actionRef}
      selectedRows={selectedRows}
      setSelectedRows={setSelectedRows}
      tableColumns={[
        {
          title: 'ID',
          dataIndex: 'id',
          hideInSearch: true,
          tooltip: intl.formatMessage({ id: `pages.permission.query.tooltip.id`, defaultMessage: 'ID 是权限的唯一标识符' }),
          render: (id, record) => (
            <Space>
              {id}
              {record.reserved && <Tag color="warning">{intl.formatMessage({ id: `pages.permission.query.reserved`, defaultMessage: '系统保留' })}</Tag>}
            </Space>
          )
        },
        {
          title: intl.formatMessage({ id: `pages.permission.query.code`, defaultMessage: '代码' }),
          dataIndex: 'code',
          tooltip: intl.formatMessage({ id: `pages.permission.query.tooltip.code`, defaultMessage: '权限代码是唯一的' }),
          render: (code, record) => (
            <Space>
              {code}
              {record.reserved && <Tag color="magenta">{intl.formatMessage({ id: `pages.permission.query.${record.code}`, defaultMessage: record.code })}</Tag>}
            </Space>
          )
        },
        {
          title: intl.formatMessage({ id: `pages.permission.query.name`, defaultMessage: '名称' }),
          dataIndex: 'name'
        },
        {
          title: intl.formatMessage({ id: `pages.permission.query.description`, defaultMessage: '描述' }),
          dataIndex: 'description',
          valueType: 'textarea'
        },
        {
          title: intl.formatMessage({ id: `pages.permission.query.createdAt`, defaultMessage: '创建时间' }),
          sorter: true,
          dataIndex: 'createdAt',
          valueType: 'dateTime',
          hideInSearch: true
        },
        {
          title: intl.formatMessage({ id: `pages.permission.query.updatedAt`, defaultMessage: '更新时间' }),
          sorter: true,
          dataIndex: 'updatedAt',
          valueType: 'dateTime',
          hideInSearch: true,
          defaultSortOrder: 'descend'
        }
      ]}
      postQuery={async (params, sort, filter) => {
        const filters: API.Filter[] = []
        if ('code' in params && params.code !== '') {
          filters.push({
            field: 'code',
            type: 'contains',
            params: [params.code]
          })
        }
        if ('name' in params && params.name !== '') {
          filters.push({
            field: 'name',
            type: 'contains',
            params: [params.name]
          })
        }
        if ('description' in params && params.description !== '') {
          filters.push({
            field: 'description',
            type: 'contains',
            params: [params.description]
          })
        }
        const sorts: API.Sort[] = []
        if ('createdAt' in sort) {
          sorts.push({
            field: 'created_at',
            desc: sort.createdAt === 'descend'
          })
        }
        if ('updatedAt' in sort) {
          sorts.push({
            field: 'updated_at',
            desc: sort.updatedAt === 'descend'
          })
        }
        return await api.permission.permissionPostQuery({
          page: params.current || 1,
          size: params.pageSize || 10,
          filters,
          sorts
        })
      }}
      deleteDelete={(id) => api.permission.permissionDeleteDelete({ id })}
      deleteBatchDelete={(ids) => api.permission.permissionDeleteBatchDelete(ids)}
    />
  )
}

export default Query
