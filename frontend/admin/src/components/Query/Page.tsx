import type { ActionType, ProColumns, RequestData } from '@ant-design/pro-components'
import { FooterToolbar, PageContainer, ProTable } from '@ant-design/pro-components'
import React, { useState } from 'react'
import { Space } from 'antd'
import { SortOrder } from 'antd/es/table/interface'
import type { QueryEntityIdReserved } from './types'
import { QueryButtonNew } from './ButtonNew'
import { QueryButtonDetail } from './ButtonDetail'
import { QueryButtonEdit } from './ButtonEdit'
import { QueryButtonDelete } from './ButtonDelete'
import { QueryButtonRecover } from './ButtonRecover'
import { QueryFooterCount } from './FooterCount'
import { QueryButtonBatchDelete } from './ButtonBatchDelete'
import { useIntl } from '@umijs/max'

type QueryPageRequestParams<U> = U & {
  current?: number
  pageSize?: number
}

type QueryPageProps<TItem extends QueryEntityIdReserved, U> = {
  name: string
  actionRef: React.RefObject<ActionType | null>
  tableColumns: ProColumns<TItem>[]
  selectedRows: TItem[]
  setSelectedRows: React.Dispatch<React.SetStateAction<TItem[]>>
  postQuery: (
    params: QueryPageRequestParams<U>,
    sort: Record<string, SortOrder>,
    filter: Record<string, (string | number)[] | null>
  ) => Promise<Partial<RequestData<TItem>>>
  deleteDelete: (id: string) => Promise<API.ResponseStringBody>
  deleteBatchDelete: (ids: string[]) => Promise<API.ResponseListStringBody>
}

export function QueryPage<TItem extends QueryEntityIdReserved, U>({
  name,
  actionRef,
  tableColumns,
  selectedRows,
  setSelectedRows,
  postQuery,
  deleteDelete,
  deleteBatchDelete
}: QueryPageProps<TItem, U>) {
  const intl = useIntl()
  const [tableLoading, setTableLoading] = useState<boolean>(false)

  const request = async (
    params: QueryPageRequestParams<U>,
    sort: Record<string, SortOrder>,
    filter: Record<string, (string | number)[] | null>
  ): Promise<Partial<RequestData<TItem>>> => {
    setTableLoading(true)
    const res = await postQuery(params, sort, filter)
    setTableLoading(false)
    return res
  }

  return (
    <PageContainer extra={[<QueryButtonRecover name={name} />]}>
      <ProTable<TItem, QueryPageRequestParams<U>>
        loading={tableLoading}
        pagination={{
          showSizeChanger: true,
          defaultPageSize: 10,
          pageSizeOptions: [5, 10, 20]
        }}
        actionRef={actionRef}
        rowKey="id"
        search={{
          labelWidth: 80
        }}
        headerTitle={intl.formatMessage({ id: `component.query.list.${name}`, defaultMessage: '列表' })}
        toolBarRender={() => [<QueryButtonNew name={name} />]}
        request={request}
        columns={[
          ...tableColumns,
          {
            title: intl.formatMessage({ id: `component.query.operation`, defaultMessage: '操作' }),
            dataIndex: 'option',
            valueType: 'option',
            fixed: 'right',
            render: (_, record) => (
              <Space.Compact block>
                {record.id && <QueryButtonDetail name={name} id={record.id} />}
                {record.reserved || (record.id && <QueryButtonEdit name={name} id={record.id} />)}
                {record.reserved ||
                  (record.id && (
                    <QueryButtonDelete<TItem>
                      id={record.id}
                      actionRef={actionRef}
                      selectedRows={selectedRows}
                      setSelectedRows={setSelectedRows}
                      deleteDelete={deleteDelete}
                    />
                  ))}
              </Space.Compact>
            )
          }
        ]}
        rowSelection={{
          getCheckboxProps: (record) => ({
            disabled: record.reserved
          }),
          onChange: (_, selectedRows) => {
            setSelectedRows(selectedRows)
          }
        }}
      />
      {selectedRows.length > 0 && (
        <FooterToolbar extra={<QueryFooterCount<TItem> selectedRows={selectedRows} />}>
          <QueryButtonBatchDelete<TItem>
            actionRef={actionRef}
            selectedRows={selectedRows}
            setSelectedRows={setSelectedRows}
            deleteBatchDelete={deleteBatchDelete}
          />
        </FooterToolbar>
      )}
    </PageContainer>
  )
}
