import { Button } from 'antd'
import React from 'react'
import { history, FormattedMessage } from '@umijs/max'

type QueryButtonDetailProps = {
  name: string
  id: string
}

export const QueryButtonDetail: React.FC<QueryButtonDetailProps> = ({ name, id }) => {
  return (
    <Button size="small" color="primary" variant="text" onClick={() => history.push(`/${name}/query/detail/${id}`)}>
      <FormattedMessage id="component.query.detail" defaultMessage="详情" />
    </Button>
  )
}
