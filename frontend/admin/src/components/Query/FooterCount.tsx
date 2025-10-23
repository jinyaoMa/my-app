import { FormattedMessage } from '@umijs/max'
import { Space } from 'antd'
import { QueryEntityId } from './types'

type QueryFooterCountProps<T extends QueryEntityId> = {
  selectedRows: T[]
}

export function QueryFooterCount<T extends QueryEntityId>({ selectedRows }: QueryFooterCountProps<T>) {
  return (
    <Space>
      <FormattedMessage
        id={selectedRows.length === 1 ? 'component.query.footerCount' : 'component.query.footerCounts'}
        values={{
          count: (
            <a>
              <b>{selectedRows.length}</b>
            </a>
          )
        }}
        defaultMessage="已选择 {count} 项"
      />
    </Space>
  )
}
