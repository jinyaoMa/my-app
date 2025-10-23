import type { ActionType } from '@ant-design/pro-components'
import { Space, Tag } from 'antd'
import React, { useRef, useState } from 'react'
import api from '@/services/api'
import { RecoverPage } from '@/components/Recover'
import { useIntl } from '@umijs/max'

type PermissionRecoverParams = {
  code?: string
  name?: string
  description?: string
}

const Recover: React.FC = () => {
  const intl = useIntl()
  const actionRef = useRef<ActionType>(null)
  const [selectedRows, setSelectedRows] = useState<API.PermissionItem[]>([])

  return (
    <RecoverPage<API.PermissionItem, PermissionRecoverParams>
      name="permission"
      actionRef={actionRef}
      selectedRows={selectedRows}
      setSelectedRows={setSelectedRows}
      tableColumns={[
        {
          title: 'ID',
          dataIndex: 'id',
          hideInSearch: true,
          tooltip: intl.formatMessage({ id: `pages.permission.recover.tooltip.id`, defaultMessage: 'ID 是权限的唯一标识符' }),
          render: (id, record) => (
            <Space>
              {id}
              {record.reserved && <Tag color="warning">{intl.formatMessage({ id: `pages.permission.recover.reserved`, defaultMessage: '系统保留' })}</Tag>}
            </Space>
          )
        },
        {
          title: intl.formatMessage({ id: `pages.permission.recover.code`, defaultMessage: '代码' }),
          dataIndex: 'code',
          tooltip: intl.formatMessage({ id: `pages.permission.recover.tooltip.code`, defaultMessage: '权限代码是唯一的' }),
          render: (code, record) => (
            <Space>
              {code}
              {record.reserved && <Tag color="magenta">{intl.formatMessage({ id: `pages.permission.recover.${record.code}`, defaultMessage: record.code })}</Tag>}
            </Space>
          )
        },
        {
          title: intl.formatMessage({ id: `pages.permission.recover.name`, defaultMessage: '名称' }),
          dataIndex: 'name'
        },
        {
          title: intl.formatMessage({ id: `pages.permission.recover.description`, defaultMessage: '描述' }),
          dataIndex: 'description',
          valueType: 'textarea'
        },
        {
          title: intl.formatMessage({ id: `pages.permission.recover.deletedAt`, defaultMessage: '删除时间' }),
          sorter: true,
          dataIndex: 'deletedAt',
          valueType: 'dateTime',
          hideInSearch: true,
          defaultSortOrder: 'descend'
        }
      ]}
      postRecycle={async (params, sort, filter) => {
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
        if ('deletedAt' in sort) {
          sorts.push({
            field: 'deleted_at',
            desc: sort.deletedAt === 'descend'
          })
        }
        return await api.permission.permissionPostRecycle({
          page: params.current || 1,
          size: params.pageSize || 10,
          filters,
          sorts
        })
      }}
      putRecover={(id) => api.permission.permissionPutRecover({ id })}
      putBatchRecover={(ids) => api.permission.permissionPutBatchRecover(ids)}
    />
  )
}

export default Recover
