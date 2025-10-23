import { Button } from 'antd'
import React from 'react'
import { FormattedMessage, history } from '@umijs/max'

type QueryButtonEditProps = {
  name: string
  id: string
}

export const QueryButtonEdit: React.FC<QueryButtonEditProps> = ({ name, id }) => {
  return (
    <Button size="small" color="primary" variant="text" onClick={() => history.push(`/${name}/query/update/${id}`)}>
      <FormattedMessage id="component.query.edit" defaultMessage="编辑" />
    </Button>
  )
}
