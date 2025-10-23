import Detail from './Detail'
import Query from './Query'
import Recover from './Recover'
import Save from './Save'

export default {
  ...Detail,
  ...Query,
  ...Recover,
  ...Save,
  'component.tagSelect.expand': 'Expand',
  'component.tagSelect.collapse': 'Collapse',
  'component.tagSelect.all': 'All'
}
