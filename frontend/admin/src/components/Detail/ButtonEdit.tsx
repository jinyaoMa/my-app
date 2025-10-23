import { Button } from 'antd'
import React from 'react'
import { history, FormattedMessage } from '@umijs/max'

type DetailButtonEditProps = {
  name: string
  id: string
}

export const DetailButtonEdit: React.FC<DetailButtonEditProps> = ({ name, id }) => {
  return (
    <Button onClick={() => history.push(`/${name}/query/update/${id}`)}>
      <FormattedMessage id="component.detail.edit" defaultMessage="编辑" />
    </Button>
  )
}
