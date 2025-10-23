import { Button, message } from 'antd'
import React from 'react'
import { FormattedMessage, useIntl, history } from '@umijs/max'

type DetailButtonDeleteProps = {
  name: string
  id: string
  deleteDelete: (id: string) => Promise<API.ResponseStringBody>
}

export const DetailButtonDelete: React.FC<DetailButtonDeleteProps> = ({ name, id, deleteDelete }) => {
  const intl = useIntl()

  return (
    <Button
      danger={true}
      onClick={async () => {
        const hide = message.loading(intl.formatMessage({ id: 'component.detail.deleting', defaultMessage: '正在删除' }))
        try {
          await deleteDelete(id)
          history.replace(`/${name}/query`)
          hide()
          message.success(intl.formatMessage({ id: 'component.detail.deleteSuccess', defaultMessage: '删除成功' }))
        } catch (_error) {
          hide()
          message.error(intl.formatMessage({ id: 'component.detail.deleteError', defaultMessage: '删除失败，请重试' }))
        }
      }}
    >
      <FormattedMessage id="component.detail.delete" defaultMessage="删除" />
    </Button>
  )
}
