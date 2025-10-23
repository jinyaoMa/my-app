import { Button } from 'antd'
import React from 'react'
import { history, FormattedMessage } from '@umijs/max'

type SaveButtonDetailProps = {
  name: string
  id: string
}

export const SaveButtonDetail: React.FC<SaveButtonDetailProps> = ({ name, id }) => {
  return (
    <Button onClick={() => history.push(`/${name}/query/detail/${id}`)}>
      <FormattedMessage id="component.save.detail" defaultMessage="详情" />
    </Button>
  )
}
