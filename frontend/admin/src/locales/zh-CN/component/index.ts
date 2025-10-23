import Detail from './Detail'
import Query from './Query'
import Recover from './Recover'
import Save from './Save'

export default {
  ...Detail,
  ...Query,
  ...Recover,
  ...Save,
  'component.tagSelect.expand': '展开',
  'component.tagSelect.collapse': '收起',
  'component.tagSelect.all': '全部'
}
