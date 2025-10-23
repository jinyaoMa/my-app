import type { ActionType, ProColumns, RequestData } from '@ant-design/pro-components'
import { FooterToolbar, PageContainer, ProTable } from '@ant-design/pro-components'
import React, { useState } from 'react'
import { Space } from 'antd'
import { SortOrder } from 'antd/es/table/interface'
import { useIntl } from '@umijs/max'
import type { RecoverEntityIdReserved } from './types'
import { RecoverButtonRecover } from './ButtonRecover'
import { RecoverButtonQuery } from './ButtonQuery'
import { RecoverFooterCount } from './FooterCount'
import { RecoverButtonBatchRecover } from './ButtonBatchRecover'

type RecoverPageRequestParams<U> = U & {
  current?: number
  pageSize?: number
}

type RecoverPageProps<TItem extends RecoverEntityIdReserved, U> = {
  name: string
  actionRef: React.RefObject<ActionType | null>
  tableColumns: ProColumns<TItem>[]
  selectedRows: TItem[]
  setSelectedRows: React.Dispatch<React.SetStateAction<TItem[]>>
  postRecycle: (
    params: RecoverPageRequestParams<U>,
    sort: Record<string, SortOrder>,
    filter: Record<string, (string | number)[] | null>
  ) => Promise<Partial<RequestData<TItem>>>
  putRecover: (id: string) => Promise<API.ResponseStringBody>
  putBatchRecover: (ids: string[]) => Promise<API.ResponseListStringBody>
}

export function RecoverPage<TItem extends RecoverEntityIdReserved, U>({
  name,
  actionRef,
  tableColumns,
  selectedRows,
  setSelectedRows,
  postRecycle,
  putRecover,
  putBatchRecover
}: RecoverPageProps<TItem, U>) {
  const intl = useIntl()
  const [tableLoading, setTableLoading] = useState<boolean>(false)

  const request = async (
    params: RecoverPageRequestParams<U>,
    sort: Record<string, SortOrder>,
    filter: Record<string, (string | number)[] | null>
  ): Promise<Partial<RequestData<TItem>>> => {
    setTableLoading(true)
    const res = await postRecycle(params, sort, filter)
    setTableLoading(false)
    return res
  }

  return (
    <PageContainer extra={[<RecoverButtonQuery name={name} />]}>
      <ProTable<TItem, RecoverPageRequestParams<U>>
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
        headerTitle={intl.formatMessage({ id: `component.recover.list.${name}`, defaultMessage: '列表' })}
        request={request}
        columns={[
          ...tableColumns,
          {
            title: intl.formatMessage({ id: `component.recover.operation`, defaultMessage: '操作' }),
            dataIndex: 'option',
            valueType: 'option',
            fixed: 'right',
            render: (_, record) => (
              <Space.Compact block>
                <RecoverButtonRecover<TItem>
                  id={record.id}
                  actionRef={actionRef}
                  selectedRows={selectedRows}
                  setSelectedRows={setSelectedRows}
                  putRecover={putRecover}
                />
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
        <FooterToolbar extra={<RecoverFooterCount<TItem> selectedRows={selectedRows} />}>
          <RecoverButtonBatchRecover<TItem>
            actionRef={actionRef}
            selectedRows={selectedRows}
            setSelectedRows={setSelectedRows}
            putBatchRecover={putBatchRecover}
          />
        </FooterToolbar>
      )}
    </PageContainer>
  )
}
