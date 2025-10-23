import { ExclamationCircleOutlined } from '@ant-design/icons'
import type { ProDescriptionsItemProps } from '@ant-design/pro-components'
import { PageContainer, ProDescriptions } from '@ant-design/pro-components'
import { useIntl, useParams, useRequest } from '@umijs/max'
import { BreadcrumbProps, Space, TabPaneProps, Tag } from 'antd'
import React, { useEffect } from 'react'
import type { DetailEntityIdReserved, DetailRequestData } from './types'
import { DetailButtonDelete, DetailButtonEdit } from '@/components/Detail'

type DetailPageProps<TDetail extends DetailEntityIdReserved> = {
  name: string
  breadcrumb: BreadcrumbProps
  columns: ProDescriptionsItemProps<TDetail>[]
  tabList?: (TabPaneProps & {
    key?: React.Key
  })[]
  tabActiveKey?: string
  setTabActiveKey?: React.Dispatch<React.SetStateAction<string>>
  postDetail: (id: string) => Promise<Partial<DetailRequestData<TDetail>>>
  deleteDelete: (id: string) => Promise<API.ResponseStringBody>
  children?: React.ReactNode
}

export const DetailPageContext = React.createContext<any>({})

export function DetailPage<TDetail extends DetailEntityIdReserved>({
  name,
  breadcrumb,
  columns,
  tabList,
  tabActiveKey,
  setTabActiveKey,
  postDetail,
  deleteDelete,
  children
}: DetailPageProps<TDetail>) {
  const intl = useIntl()
  const { id = '' } = useParams()

  const { data, loading, run } = useRequest(() => postDetail(id), { manual: true })

  useEffect(() => {
    if (id) {
      run()
    }
  }, [run, id])

  return (
    <PageContainer
      loading={loading}
      title={
        <Space>
          {`ID: ${id}`}
          {data?.reserved && (
            <Tag icon={<ExclamationCircleOutlined />} color="warning">
              {intl.formatMessage({
                id: 'component.detail.reserved',
                defaultMessage: '系统保留'
              })}
            </Tag>
          )}
        </Space>
      }
      extra={[
        data?.reserved || <DetailButtonEdit name={name} id={id} />,
        data?.reserved || <DetailButtonDelete name={name} id={id} deleteDelete={deleteDelete} />
      ]}
      breadcrumb={breadcrumb}
      content={
        <ProDescriptions<TDetail>
          style={{
            marginBottom: '12px'
          }}
          dataSource={data}
          columns={columns}
        />
      }
      tabActiveKey={tabActiveKey}
      tabList={tabList}
      onTabChange={setTabActiveKey}
    >
      <DetailPageContext.Provider value={data}>{children}</DetailPageContext.Provider>
    </PageContainer>
  )
}
