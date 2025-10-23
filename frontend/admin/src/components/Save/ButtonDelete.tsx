import { Button, message } from 'antd'
import React from 'react'
import { FormattedMessage, useIntl, history } from '@umijs/max'

type SaveButtonDeleteProps = {
  name: string
  id: string
  deleteDelete: (id: string) => Promise<API.ResponseStringBody>
}

export const SaveButtonDelete: React.FC<SaveButtonDeleteProps> = ({ name, id, deleteDelete }) => {
  const intl = useIntl()

  return (
    <Button
      danger={true}
      onClick={async () => {
        const hide = message.loading(intl.formatMessage({ id: 'component.save.deleting', defaultMessage: '正在删除' }))
        try {
          await deleteDelete(id)
          history.replace(`/${name}/query`)
          hide()
          message.success(intl.formatMessage({ id: 'component.save.deleteSuccess', defaultMessage: '删除成功' }))
        } catch (_error) {
          hide()
          message.error(intl.formatMessage({ id: 'component.save.deleteError', defaultMessage: '删除失败，请重试' }))
        }
      }}
    >
      <FormattedMessage id="component.save.delete" defaultMessage="删除" />
    </Button>
  )
}
