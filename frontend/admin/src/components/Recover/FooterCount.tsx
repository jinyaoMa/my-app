import { FormattedMessage } from '@umijs/max'
import { Space } from 'antd'
import { RecoverEntityId } from './types'

type RecoverFooterCountProps<T extends RecoverEntityId> = {
  selectedRows: T[]
}

export function RecoverFooterCount<T extends RecoverEntityId>({ selectedRows }: RecoverFooterCountProps<T>) {
  return (
    <Space>
      <FormattedMessage
        id={selectedRows.length === 1 ? 'component.recover.footerCount' : 'component.recover.footerCounts'}
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
