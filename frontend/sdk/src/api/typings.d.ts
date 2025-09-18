declare namespace API {
  type AuthLogin = {
    /** Access Token */
    accessToken: string
    /** Expired At */
    expiredAt: string
    /** Refresh Token */
    refreshToken: string
  }

  type AuthLoginRequest = {
    /** A URL to the JSON Schema for this object. */
    $schema?: string
    /** User Account */
    account: string
    /** User Password */
    password: string
  }

  type authRefreshParams = {
    /** Refresh Token */
    token?: string
  }

  type Criteria = {
    /** A URL to the JSON Schema for this object. */
    $schema?: string
    /** List Filter Conditions */
    filters?: any
    /** Included Associations */
    includes?: any
    /** Join 121/belong2 Associations */
    joins?: any
    /** Omitted Fields */
    omits?: any
    /** Page Number */
    page: number
    /** Selected Fields */
    selects?: any
    /** Page Size */
    size: number
    /** List Sort Conditions */
    sorts?: any
  }

  type EndpointsItem = {
    /** Method */
    method: string
    /** Operation ID */
    operationId: string
    /** Path */
    path: string
    /** Summary */
    summary: string
    /** Tags */
    tags: any
  }

  type endpointsListParams = {
    /** Filter by Tag */
    tag?: string
    /** Filter by Operation ID */
    operationId?: string
  }

  type ErrorDetail = {
    /** Where the error occurred, e.g. 'body.items[3].tags' or 'path.thing-id' */
    location?: string
    /** Error message text */
    message?: string
    /** The value at the given location */
    value?: any
  }

  type ErrorModel = {
    /** A URL to the JSON Schema for this object. */
    $schema?: string
    /** A human-readable explanation specific to this occurrence of the problem. */
    detail?: string
    /** Optional list of individual error details */
    errors?: any
    /** A URI reference that identifies the specific occurrence of the problem. */
    instance?: string
    /** HTTP status code */
    status?: number
    /** A short, human-readable summary of the problem type. This value should not change between occurrences of the error. */
    title?: string
    /** A URI reference to human-readable documentation for the error. */
    type?: string
  }

  type Filter = {
    /** Filter Field Name */
    field: string
    /** Or or And */
    or?: boolean
    /** Filter Condition Parameters */
    params: any
    /** Filter Special */
    special?: number
    /** Filter Type */
    type: number
  }

  type groupDeleteParams = {
    id?: string
  }

  type GroupDetail = {
    /** Created At */
    createAt: string
    /** Deleted At */
    deleteAt?: string
    /** Entity ID */
    id: string
    /** Updated At */
    updateAt: string
  }

  type groupDetailParams = {
    id?: string
    /** Included Associations */
    includes?: any
  }

  type GroupItem = {
    /** Code */
    code: string
    /** Created At */
    createAt: string
    /** Deleted At */
    deleteAt?: string
    /** Description */
    description: string
    /** Group Roles */
    groupRoles: any
    /** Group Users */
    groupUsers: any
    /** Entity ID */
    id: string
    /** Name */
    name: string
    /** Roles */
    roles: any
    /** Updated At */
    updateAt: string
    /** Users */
    users: any
  }

  type groupRoleDeleteParams = {
    id?: string
  }

  type GroupRoleDetail = {
    /** Created At */
    createAt: string
    /** Deleted At */
    deleteAt?: string
    /** Entity ID */
    id: string
    /** Updated At */
    updateAt: string
  }

  type groupRoleDetailParams = {
    id?: string
    /** Included Associations */
    includes?: any
  }

  type GroupRoleItem = {
    /** Created At */
    createAt: string
    /** Deleted At */
    deleteAt?: string
    /** Group */
    group: GroupItem
    /** Group ID */
    groupId: string
    /** Entity ID */
    id: string
    /** Role */
    role: RoleItem
    /** Role ID */
    roleId: string
    /** Updated At */
    updateAt: string
  }

  type GroupRoleSave = {
    /** A URL to the JSON Schema for this object. */
    $schema?: string
    /** Created At */
    createAt: string
    /** Deleted At */
    deleteAt?: string
    /** Entity ID */
    id: string
    /** Updated At */
    updateAt: string
  }

  type GroupSave = {
    /** A URL to the JSON Schema for this object. */
    $schema?: string
    /** Created At */
    createAt: string
    /** Deleted At */
    deleteAt?: string
    /** Entity ID */
    id: string
    /** Updated At */
    updateAt: string
  }

  type groupUserDeleteParams = {
    id?: string
  }

  type GroupUserDetail = {
    /** Created At */
    createAt: string
    /** Deleted At */
    deleteAt?: string
    /** Entity ID */
    id: string
    /** Updated At */
    updateAt: string
  }

  type groupUserDetailParams = {
    id?: string
    /** Included Associations */
    includes?: any
  }

  type GroupUserItem = {
    /** Created At */
    createAt: string
    /** Deleted At */
    deleteAt?: string
    /** Group */
    group: GroupItem
    /** Group ID */
    groupId: string
    /** Entity ID */
    id: string
    /** Updated At */
    updateAt: string
    /** User */
    user: UserItem
    /** User ID */
    userId: string
  }

  type GroupUserSave = {
    /** A URL to the JSON Schema for this object. */
    $schema?: string
    /** Created At */
    createAt: string
    /** Deleted At */
    deleteAt?: string
    /** Entity ID */
    id: string
    /** Updated At */
    updateAt: string
  }

  type OperationIdEnumPairItem = {
    /** Enum Value */
    enum: number
    /** Operation ID */
    operationId: string
  }

  type operationIdEnumPairListParams = {
    /** Filter by Operation ID */
    operationId?: string
  }

  type optionDeleteParams = {
    id?: string
  }

  type OptionDetail = {
    /** Created At */
    createAt: string
    /** Deleted At */
    deleteAt?: string
    /** Entity ID */
    id: string
    /** Key */
    key: string
    /** Updated At */
    updateAt: string
    /** Value */
    value: string
  }

  type optionDetailParams = {
    id?: string
    /** Included Associations */
    includes?: any
  }

  type OptionItem = {
    /** Created At */
    createAt: string
    /** Deleted At */
    deleteAt?: string
    /** Entity ID */
    id: string
    /** Key */
    key: string
    /** Updated At */
    updateAt: string
    /** Value */
    value: string
  }

  type OptionSave = {
    /** A URL to the JSON Schema for this object. */
    $schema?: string
    /** Created At */
    createAt: string
    /** Deleted At */
    deleteAt?: string
    /** Entity ID */
    id: string
    /** Key */
    key: string
    /** Updated At */
    updateAt: string
    /** Value */
    value: string
  }

  type permissionDeleteParams = {
    id?: string
  }

  type PermissionDetail = {
    /** Created At */
    createAt: string
    /** Deleted At */
    deleteAt?: string
    /** Entity ID */
    id: string
    /** Updated At */
    updateAt: string
  }

  type permissionDetailParams = {
    id?: string
    /** Included Associations */
    includes?: any
  }

  type PermissionItem = {
    /** Code */
    code: string
    /** Created At */
    createAt: string
    /** Deleted At */
    deleteAt?: string
    /** Description */
    description: string
    /** Flag */
    flag: string
    /** Entity ID */
    id: string
    /** Name */
    name: string
    /** Role Permissions */
    rolePermissions: any
    /** Roles */
    roles: any
    /** Updated At */
    updateAt: string
  }

  type PermissionSave = {
    /** A URL to the JSON Schema for this object. */
    $schema?: string
    /** Created At */
    createAt: string
    /** Deleted At */
    deleteAt?: string
    /** Entity ID */
    id: string
    /** Updated At */
    updateAt: string
  }

  type ResponseAuthLoginBody = {
    /** A URL to the JSON Schema for this object. */
    $schema?: string
    /** 0 => success; otherwise, fail */
    code: number
    /** response payload */
    data: AuthLogin
    /** hint message */
    message: string
    /** Has the action done successfully? */
    success: boolean
    /** if request is list query, total represents the count of the list w/o any user query conditions but conditions built-in; or represents affected rows for other requests */
    total: number
  }

  type ResponseGroupDetailBody = {
    /** A URL to the JSON Schema for this object. */
    $schema?: string
    /** 0 => success; otherwise, fail */
    code: number
    /** response payload */
    data: GroupDetail
    /** hint message */
    message: string
    /** Has the action done successfully? */
    success: boolean
    /** if request is list query, total represents the count of the list w/o any user query conditions but conditions built-in; or represents affected rows for other requests */
    total: number
  }

  type ResponseGroupRoleDetailBody = {
    /** A URL to the JSON Schema for this object. */
    $schema?: string
    /** 0 => success; otherwise, fail */
    code: number
    /** response payload */
    data: GroupRoleDetail
    /** hint message */
    message: string
    /** Has the action done successfully? */
    success: boolean
    /** if request is list query, total represents the count of the list w/o any user query conditions but conditions built-in; or represents affected rows for other requests */
    total: number
  }

  type ResponseGroupUserDetailBody = {
    /** A URL to the JSON Schema for this object. */
    $schema?: string
    /** 0 => success; otherwise, fail */
    code: number
    /** response payload */
    data: GroupUserDetail
    /** hint message */
    message: string
    /** Has the action done successfully? */
    success: boolean
    /** if request is list query, total represents the count of the list w/o any user query conditions but conditions built-in; or represents affected rows for other requests */
    total: number
  }

  type ResponseListEndpointsItemBody = {
    /** A URL to the JSON Schema for this object. */
    $schema?: string
    /** 0 => success; otherwise, fail */
    code: number
    /** response payload */
    data: any
    /** hint message */
    message: string
    /** Has the action done successfully? */
    success: boolean
    /** if request is list query, total represents the count of the list w/o any user query conditions but conditions built-in; or represents affected rows for other requests */
    total: number
  }

  type ResponseListGroupItemBody = {
    /** A URL to the JSON Schema for this object. */
    $schema?: string
    /** 0 => success; otherwise, fail */
    code: number
    /** response payload */
    data: any
    /** hint message */
    message: string
    /** Has the action done successfully? */
    success: boolean
    /** if request is list query, total represents the count of the list w/o any user query conditions but conditions built-in; or represents affected rows for other requests */
    total: number
  }

  type ResponseListGroupRoleItemBody = {
    /** A URL to the JSON Schema for this object. */
    $schema?: string
    /** 0 => success; otherwise, fail */
    code: number
    /** response payload */
    data: any
    /** hint message */
    message: string
    /** Has the action done successfully? */
    success: boolean
    /** if request is list query, total represents the count of the list w/o any user query conditions but conditions built-in; or represents affected rows for other requests */
    total: number
  }

  type ResponseListGroupUserItemBody = {
    /** A URL to the JSON Schema for this object. */
    $schema?: string
    /** 0 => success; otherwise, fail */
    code: number
    /** response payload */
    data: any
    /** hint message */
    message: string
    /** Has the action done successfully? */
    success: boolean
    /** if request is list query, total represents the count of the list w/o any user query conditions but conditions built-in; or represents affected rows for other requests */
    total: number
  }

  type ResponseListOperationIdEnumPairItemBody = {
    /** A URL to the JSON Schema for this object. */
    $schema?: string
    /** 0 => success; otherwise, fail */
    code: number
    /** response payload */
    data: any
    /** hint message */
    message: string
    /** Has the action done successfully? */
    success: boolean
    /** if request is list query, total represents the count of the list w/o any user query conditions but conditions built-in; or represents affected rows for other requests */
    total: number
  }

  type ResponseListOptionItemBody = {
    /** A URL to the JSON Schema for this object. */
    $schema?: string
    /** 0 => success; otherwise, fail */
    code: number
    /** response payload */
    data: any
    /** hint message */
    message: string
    /** Has the action done successfully? */
    success: boolean
    /** if request is list query, total represents the count of the list w/o any user query conditions but conditions built-in; or represents affected rows for other requests */
    total: number
  }

  type ResponseListPermissionItemBody = {
    /** A URL to the JSON Schema for this object. */
    $schema?: string
    /** 0 => success; otherwise, fail */
    code: number
    /** response payload */
    data: any
    /** hint message */
    message: string
    /** Has the action done successfully? */
    success: boolean
    /** if request is list query, total represents the count of the list w/o any user query conditions but conditions built-in; or represents affected rows for other requests */
    total: number
  }

  type ResponseListRoleItemBody = {
    /** A URL to the JSON Schema for this object. */
    $schema?: string
    /** 0 => success; otherwise, fail */
    code: number
    /** response payload */
    data: any
    /** hint message */
    message: string
    /** Has the action done successfully? */
    success: boolean
    /** if request is list query, total represents the count of the list w/o any user query conditions but conditions built-in; or represents affected rows for other requests */
    total: number
  }

  type ResponseListRolePermissionItemBody = {
    /** A URL to the JSON Schema for this object. */
    $schema?: string
    /** 0 => success; otherwise, fail */
    code: number
    /** response payload */
    data: any
    /** hint message */
    message: string
    /** Has the action done successfully? */
    success: boolean
    /** if request is list query, total represents the count of the list w/o any user query conditions but conditions built-in; or represents affected rows for other requests */
    total: number
  }

  type ResponseListStringBody = {
    /** A URL to the JSON Schema for this object. */
    $schema?: string
    /** 0 => success; otherwise, fail */
    code: number
    /** response payload */
    data: any
    /** hint message */
    message: string
    /** Has the action done successfully? */
    success: boolean
    /** if request is list query, total represents the count of the list w/o any user query conditions but conditions built-in; or represents affected rows for other requests */
    total: number
  }

  type ResponseListUserItemBody = {
    /** A URL to the JSON Schema for this object. */
    $schema?: string
    /** 0 => success; otherwise, fail */
    code: number
    /** response payload */
    data: any
    /** hint message */
    message: string
    /** Has the action done successfully? */
    success: boolean
    /** if request is list query, total represents the count of the list w/o any user query conditions but conditions built-in; or represents affected rows for other requests */
    total: number
  }

  type ResponseListUserRoleItemBody = {
    /** A URL to the JSON Schema for this object. */
    $schema?: string
    /** 0 => success; otherwise, fail */
    code: number
    /** response payload */
    data: any
    /** hint message */
    message: string
    /** Has the action done successfully? */
    success: boolean
    /** if request is list query, total represents the count of the list w/o any user query conditions but conditions built-in; or represents affected rows for other requests */
    total: number
  }

  type ResponseOptionDetailBody = {
    /** A URL to the JSON Schema for this object. */
    $schema?: string
    /** 0 => success; otherwise, fail */
    code: number
    /** response payload */
    data: OptionDetail
    /** hint message */
    message: string
    /** Has the action done successfully? */
    success: boolean
    /** if request is list query, total represents the count of the list w/o any user query conditions but conditions built-in; or represents affected rows for other requests */
    total: number
  }

  type ResponsePermissionDetailBody = {
    /** A URL to the JSON Schema for this object. */
    $schema?: string
    /** 0 => success; otherwise, fail */
    code: number
    /** response payload */
    data: PermissionDetail
    /** hint message */
    message: string
    /** Has the action done successfully? */
    success: boolean
    /** if request is list query, total represents the count of the list w/o any user query conditions but conditions built-in; or represents affected rows for other requests */
    total: number
  }

  type ResponseRoleDetailBody = {
    /** A URL to the JSON Schema for this object. */
    $schema?: string
    /** 0 => success; otherwise, fail */
    code: number
    /** response payload */
    data: RoleDetail
    /** hint message */
    message: string
    /** Has the action done successfully? */
    success: boolean
    /** if request is list query, total represents the count of the list w/o any user query conditions but conditions built-in; or represents affected rows for other requests */
    total: number
  }

  type ResponseRolePermissionDetailBody = {
    /** A URL to the JSON Schema for this object. */
    $schema?: string
    /** 0 => success; otherwise, fail */
    code: number
    /** response payload */
    data: RolePermissionDetail
    /** hint message */
    message: string
    /** Has the action done successfully? */
    success: boolean
    /** if request is list query, total represents the count of the list w/o any user query conditions but conditions built-in; or represents affected rows for other requests */
    total: number
  }

  type ResponseStringBody = {
    /** A URL to the JSON Schema for this object. */
    $schema?: string
    /** 0 => success; otherwise, fail */
    code: number
    /** response payload */
    data: string
    /** hint message */
    message: string
    /** Has the action done successfully? */
    success: boolean
    /** if request is list query, total represents the count of the list w/o any user query conditions but conditions built-in; or represents affected rows for other requests */
    total: number
  }

  type ResponseUserDetailBody = {
    /** A URL to the JSON Schema for this object. */
    $schema?: string
    /** 0 => success; otherwise, fail */
    code: number
    /** response payload */
    data: UserDetail
    /** hint message */
    message: string
    /** Has the action done successfully? */
    success: boolean
    /** if request is list query, total represents the count of the list w/o any user query conditions but conditions built-in; or represents affected rows for other requests */
    total: number
  }

  type ResponseUserRoleDetailBody = {
    /** A URL to the JSON Schema for this object. */
    $schema?: string
    /** 0 => success; otherwise, fail */
    code: number
    /** response payload */
    data: UserRoleDetail
    /** hint message */
    message: string
    /** Has the action done successfully? */
    success: boolean
    /** if request is list query, total represents the count of the list w/o any user query conditions but conditions built-in; or represents affected rows for other requests */
    total: number
  }

  type roleDeleteParams = {
    id?: string
  }

  type RoleDetail = {
    /** Created At */
    createAt: string
    /** Deleted At */
    deleteAt?: string
    /** Entity ID */
    id: string
    /** Updated At */
    updateAt: string
  }

  type roleDetailParams = {
    id?: string
    /** Included Associations */
    includes?: any
  }

  type RoleItem = {
    /** Code */
    code: string
    /** Created At */
    createAt: string
    /** Deleted At */
    deleteAt?: string
    /** Description */
    description: string
    /** Group Roles */
    groupRoles: any
    /** Groups */
    groups: any
    /** Entity ID */
    id: string
    /** Name */
    name: string
    /** Permissions */
    permissions: any
    /** Role Permissions */
    rolePermissions: any
    /** Updated At */
    updateAt: string
    /** User Roles */
    userRoles: any
    /** Users */
    users: any
  }

  type rolePermissionDeleteParams = {
    id?: string
  }

  type RolePermissionDetail = {
    /** Created At */
    createAt: string
    /** Deleted At */
    deleteAt?: string
    /** Entity ID */
    id: string
    /** Updated At */
    updateAt: string
  }

  type rolePermissionDetailParams = {
    id?: string
    /** Included Associations */
    includes?: any
  }

  type RolePermissionItem = {
    /** Created At */
    createAt: string
    /** Deleted At */
    deleteAt?: string
    /** Entity ID */
    id: string
    /** Permission */
    permission: PermissionItem
    /** Permission ID */
    permissionId: string
    /** Role */
    role: RoleItem
    /** Role ID */
    roleId: string
    /** Updated At */
    updateAt: string
  }

  type RolePermissionSave = {
    /** A URL to the JSON Schema for this object. */
    $schema?: string
    /** Created At */
    createAt: string
    /** Deleted At */
    deleteAt?: string
    /** Entity ID */
    id: string
    /** Updated At */
    updateAt: string
  }

  type RoleSave = {
    /** A URL to the JSON Schema for this object. */
    $schema?: string
    /** Created At */
    createAt: string
    /** Deleted At */
    deleteAt?: string
    /** Entity ID */
    id: string
    /** Updated At */
    updateAt: string
  }

  type Sort = {
    /** Desc or Asc */
    desc?: boolean
    /** Sorted Field Name */
    field: string
  }

  type userDeleteParams = {
    id?: string
  }

  type UserDetail = {
    /** Created At */
    createAt: string
    /** Deleted At */
    deleteAt?: string
    /** Entity ID */
    id: string
    /** Updated At */
    updateAt: string
  }

  type userDetailParams = {
    id?: string
    /** Included Associations */
    includes?: any
  }

  type UserItem = {
    /** Account */
    account: string
    /** Created At */
    createAt: string
    /** Deleted At */
    deleteAt?: string
    /** Description */
    description: string
    /** Group Users */
    groupUsers: any
    /** Groups */
    groups: any
    /** Entity ID */
    id: string
    /** Name */
    name: string
    /** Password */
    password: string
    /** Roles */
    roles: any
    /** Updated At */
    updateAt: string
    /** User Roles */
    userRoles: any
  }

  type userRoleDeleteParams = {
    id?: string
  }

  type UserRoleDetail = {
    /** Created At */
    createAt: string
    /** Deleted At */
    deleteAt?: string
    /** Entity ID */
    id: string
    /** Updated At */
    updateAt: string
  }

  type userRoleDetailParams = {
    id?: string
    /** Included Associations */
    includes?: any
  }

  type UserRoleItem = {
    /** Created At */
    createAt: string
    /** Deleted At */
    deleteAt?: string
    /** Entity ID */
    id: string
    /** Role */
    role: RoleItem
    /** Role ID */
    roleId: string
    /** Updated At */
    updateAt: string
    /** User */
    user: UserItem
    /** User ID */
    userId: string
  }

  type UserRoleSave = {
    /** A URL to the JSON Schema for this object. */
    $schema?: string
    /** Created At */
    createAt: string
    /** Deleted At */
    deleteAt?: string
    /** Entity ID */
    id: string
    /** Updated At */
    updateAt: string
  }

  type UserSave = {
    /** A URL to the JSON Schema for this object. */
    $schema?: string
    /** Created At */
    createAt: string
    /** Deleted At */
    deleteAt?: string
    /** Entity ID */
    id: string
    /** Updated At */
    updateAt: string
  }
}
