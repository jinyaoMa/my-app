import type { ProDescriptionsItemProps } from '@ant-design/pro-components'
import { useIntl } from '@umijs/max'
import { Badge, BreadcrumbProps, Space } from 'antd'
import React, { useState } from 'react'
import api from '@/services/api'
import { DetailCardEndpointsConfig, DetailPage } from '@/components/Detail'

const Detail: React.FC = () => {
  const intl = useIntl()
  const [tabActiveKey, setTabActiveKey] = useState<string>('endpoints')
  const [availEndpointsCount, setAvailEndpointsCount] = useState<number>(0)

  const breadcrumb: BreadcrumbProps = {
    items: [
      {
        title: intl.formatMessage({ id: 'menu.permission', defaultMessage: '权限管理' }),
        path: '/permission'
      },
      {
        title: intl.formatMessage({ id: 'menu.permission.query', defaultMessage: '查询权限' }),
        path: '/query'
      },
      {
        title: intl.formatMessage({ id: 'menu.permission.detail', defaultMessage: '权限详情' })
      }
    ]
  }

  const columns: ProDescriptionsItemProps<API.PermissionDetail>[] = [
    {
      title: intl.formatMessage({ id: 'pages.permission.detail.createdAt', defaultMessage: '创建时间' }),
      dataIndex: 'createdAt',
      valueType: 'dateTime'
    },
    {
      title: intl.formatMessage({ id: 'pages.permission.detail.updatedAt', defaultMessage: '更新时间' }),
      dataIndex: 'updatedAt',
      valueType: 'dateTime'
    },
    {
      title: intl.formatMessage({ id: 'pages.permission.detail.code', defaultMessage: '代码' }),
      dataIndex: 'code',
      tooltip: intl.formatMessage({ id: 'pages.permission.detail.tooltip.code', defaultMessage: '权限代码是唯一的' })
    },
    {
      title: intl.formatMessage({ id: 'pages.permission.detail.name', defaultMessage: '名称' }),
      dataIndex: 'name'
    },
    {
      title: intl.formatMessage({ id: 'pages.permission.detail.description', defaultMessage: '描述' }),
      dataIndex: 'description',
      valueType: 'textarea'
    }
  ]

  return (
    <DetailPage<API.PermissionDetail>
      name="permission"
      breadcrumb={breadcrumb}
      columns={columns}
      tabList={[
        {
          key: 'endpoints',
          tab: (
            <Space>
              {intl.formatMessage({ id: 'pages.permission.detail.endpoints', defaultMessage: '可访问接口' })}
              <Badge color="blue" count={availEndpointsCount} overflowCount={availEndpointsCount} />
            </Space>
          )
        }
      ]}
      tabActiveKey={tabActiveKey}
      setTabActiveKey={setTabActiveKey}
      postDetail={(id: string) => api.permission.permissionPostDetail({ id }, [])}
      deleteDelete={(id: string) => api.permission.permissionDeleteDelete({ id })}
    >
      {tabActiveKey === 'endpoints' && (
        <DetailCardEndpointsConfig<API.PermissionDetail> setAvailEndpointsCount={setAvailEndpointsCount} />
      )}
    </DetailPage>
  )
}
export default Detail
