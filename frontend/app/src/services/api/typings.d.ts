declare namespace API {
  type authGetRefreshParams = {
    /** Refresh Token */
    refresh: string
  }

  type authGetRoleParams = {
    /** Role Id (Base36) */
    roleId: string
    /** Refresh Token */
    refresh: string
  }

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

  type CrudSearchRequest = {
    /** A URL to the JSON Schema for this object. */
    $schema?: string
    /** Is Exact Search */
    exact?: boolean
    /** Search Fields */
    fields: any
    /** Included Associations */
    includes?: any
    /** Search Keyword */
    keyword: string
    /** Search Limit */
    limit: number
  }

  type CurrentUser = {
    /** User Id (Base36) */
    id: string
    /** Name */
    name: string
    /** Available Roles */
    roles: any
  }

  type CurrentUserRole = {
    /** Code */
    code: string
    /** Role Id (Base36) */
    id: string
    /** Is Current Role */
    isCurrent: boolean
    /** Name */
    name: string
  }

  type endpointsGetListParams = {
    /** Filter by Tag */
    tag?: string
    /** Filter by Operation ID */
    operationId?: string
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

  type fileCategoryDeleteDeleteParams = {
    /** Model Id (Base36) */
    id: string
  }

  type FileCategoryDetail = {
    /** Code */
    code: string
    /** Created At */
    createAt: string
    /** Deleted At */
    deleteAt: string
    /** File Extensions */
    fileExtensions: any
    /** Model Id (Base36) */
    id?: string
    /** Name */
    name: string
    /** Updated At */
    updateAt: string
  }

  type FileCategoryItem = {
    /** Code */
    code: string
    /** Created At */
    createAt: string
    /** Deleted At */
    deleteAt: string
    /** Model Id (Base36) */
    id?: string
    /** Name */
    name: string
    /** Updated At */
    updateAt: string
  }

  type fileCategoryPostDetailParams = {
    /** Model Id (Base36) */
    id: string
  }

  type fileCategoryPutRecoverParams = {
    /** Model Id (Base36) */
    id: string
  }

  type FileCategorySave = {
    /** A URL to the JSON Schema for this object. */
    $schema?: string
    /** Code */
    code: string
    /** Model Id (Base36) */
    id?: string
    /** Name */
    name: string
  }

  type FileCategorySearchItem = {
    /** Code */
    code: string
    /** Model Id (Base36) */
    id?: string
    /** Name */
    name: string
  }

  type fileDeleteDeleteParams = {
    /** Model Id (Base36) */
    id: string
  }

  type FileDetail = {
    /** Checksums */
    checksums: any
    /** Created At */
    createAt: string
    /** Deleted At */
    deleteAt: string
    /** File Extension */
    fileExtension: FileExtensionItem
    /** File Extension Id */
    fileExtensionId: any
    /** File Groups */
    fileGroups: any
    /** File Users */
    fileUsers: any
    /** Groups */
    groups: any
    /** Is Hidden or Visible */
    hidden: boolean
    /** Model Id (Base36) */
    id?: string
    /** Is Directory or Not */
    isDir: boolean
    /** Name */
    name: string
    /** Oid */
    oid: string
    /** Is Read-only or Not */
    readonly: boolean
    /** Size */
    size: number
    /** Status: uploading, persisted, notfound, forbidden */
    status: string
    /** Updated At */
    updateAt: string
    /** Users */
    users: any
  }

  type fileExtensionDeleteDeleteParams = {
    /** Model Id (Base36) */
    id: string
  }

  type FileExtensionDetail = {
    /** Created At */
    createAt: string
    /** Deleted At */
    deleteAt: string
    /** Ext */
    ext: string
    /** File Category */
    fileCategory: FileCategoryItem
    /** File Category Id */
    fileCategoryId: any
    /** Files */
    files: any
    /** Model Id (Base36) */
    id?: string
    /** Mime */
    mime: string
    /** Name */
    name: string
    /** Updated At */
    updateAt: string
  }

  type FileExtensionItem = {
    /** Created At */
    createAt: string
    /** Deleted At */
    deleteAt: string
    /** Ext */
    ext: string
    /** File Category */
    fileCategory: FileCategoryItem
    /** File Category Id */
    fileCategoryId: any
    /** Model Id (Base36) */
    id?: string
    /** Mime */
    mime: string
    /** Name */
    name: string
    /** Updated At */
    updateAt: string
  }

  type fileExtensionPostDetailParams = {
    /** Model Id (Base36) */
    id: string
  }

  type fileExtensionPutRecoverParams = {
    /** Model Id (Base36) */
    id: string
  }

  type FileExtensionSave = {
    /** A URL to the JSON Schema for this object. */
    $schema?: string
    /** Ext */
    ext: string
    /** File Category Id */
    fileCategoryId: any
    /** Model Id (Base36) */
    id?: string
    /** Mime */
    mime: string
    /** Name */
    name: string
  }

  type FileExtensionSearchItem = {
    /** Ext */
    ext: string
    /** File Category */
    fileCategory: FileCategorySearchItem
    /** File Category Id */
    fileCategoryId: any
    /** Model Id (Base36) */
    id?: string
    /** Name */
    name: string
  }

  type fileGroupDeleteDeleteParams = {
    /** Model Id (Base36) */
    id: string
  }

  type FileGroupDetail = {
    /** Created At */
    createAt: string
    /** Deleted At */
    deleteAt: string
    /** File */
    file: FileItem
    /** File Id */
    fileId: string
    /** Group */
    group: GroupItem
    /** Group Id */
    groupId: string
    /** Model Id (Base36) */
    id?: string
    /** No Create */
    noCreate: boolean
    /** No Delete */
    noDelete: boolean
    /** No Read */
    noRead: boolean
    /** No Update */
    noUpdate: boolean
    /** Updated At */
    updateAt: string
  }

  type FileGroupItem = {
    /** Created At */
    createAt: string
    /** Deleted At */
    deleteAt: string
    /** File */
    file: FileItem
    /** File Id */
    fileId: string
    /** Group */
    group: GroupItem
    /** Group Id */
    groupId: string
    /** Model Id (Base36) */
    id?: string
    /** No Create */
    noCreate: boolean
    /** No Delete */
    noDelete: boolean
    /** No Read */
    noRead: boolean
    /** No Update */
    noUpdate: boolean
    /** Updated At */
    updateAt: string
  }

  type fileGroupPostDetailParams = {
    /** Model Id (Base36) */
    id: string
  }

  type FileGroupSave = {
    /** A URL to the JSON Schema for this object. */
    $schema?: string
    /** File Id */
    fileId: string
    /** Group Id */
    groupId: string
    /** Model Id (Base36) */
    id?: string
    /** No Create */
    noCreate: boolean
    /** No Delete */
    noDelete: boolean
    /** No Read */
    noRead: boolean
    /** No Update */
    noUpdate: boolean
  }

  type FileGroupSearchItem = {
    /** File */
    file: FileSearchItem
    /** File Id */
    fileId: string
    /** Group */
    group: GroupSearchItem
    /** Group Id */
    groupId: string
    /** Model Id (Base36) */
    id?: string
    /** No Create */
    noCreate: boolean
    /** No Delete */
    noDelete: boolean
    /** No Read */
    noRead: boolean
    /** No Update */
    noUpdate: boolean
  }

  type FileItem = {
    /** Created At */
    createAt: string
    /** Deleted At */
    deleteAt: string
    /** File Extension */
    fileExtension: FileExtensionItem
    /** File Extension Id */
    fileExtensionId: any
    /** Is Hidden or Visible */
    hidden: boolean
    /** Model Id (Base36) */
    id?: string
    /** Is Directory or Not */
    isDir: boolean
    /** Name */
    name: string
    /** Oid */
    oid: string
    /** Is Read-only or Not */
    readonly: boolean
    /** Size */
    size: number
    /** Status: uploading, persisted, notfound, forbidden */
    status: string
    /** Updated At */
    updateAt: string
  }

  type filePostDetailParams = {
    /** Model Id (Base36) */
    id: string
  }

  type filePutRecoverParams = {
    /** Model Id (Base36) */
    id: string
  }

  type FileSave = {
    /** A URL to the JSON Schema for this object. */
    $schema?: string
    /** Checksums */
    checksums: any
    /** File Extension Id */
    fileExtensionId: any
    /** Is Hidden or Visible */
    hidden: boolean
    /** Model Id (Base36) */
    id?: string
    /** Is Directory or Not */
    isDir: boolean
    /** Name */
    name: string
    /** Oid */
    oid: string
    /** Is Read-only or Not */
    readonly: boolean
    /** Size */
    size: number
    /** Status: uploading, persisted, notfound, forbidden */
    status: string
  }

  type FileSearchItem = {
    /** File Extension */
    fileExtension: FileExtensionSearchItem
    /** File Extension Id */
    fileExtensionId: any
    /** Model Id (Base36) */
    id?: string
    /** Is Directory or Not */
    isDir: boolean
    /** Name */
    name: string
    /** Is Read-only or Not */
    readonly: boolean
    /** Size */
    size: number
    /** Status: uploading, persisted, notfound, forbidden */
    status: string
  }

  type fileUserDeleteDeleteParams = {
    /** Model Id (Base36) */
    id: string
  }

  type FileUserDetail = {
    /** Created At */
    createAt: string
    /** Deleted At */
    deleteAt: string
    /** File */
    file: FileItem
    /** File Id */
    fileId: string
    /** Model Id (Base36) */
    id?: string
    /** Is Used as Avatar or Not */
    isAvatar: boolean
    /** No Create */
    noCreate: boolean
    /** No Delete */
    noDelete: boolean
    /** No Read */
    noRead: boolean
    /** No Update */
    noUpdate: boolean
    /** Updated At */
    updateAt: string
    /** User */
    user: UserItem
    /** User Id */
    userId: string
  }

  type FileUserItem = {
    /** Created At */
    createAt: string
    /** Deleted At */
    deleteAt: string
    /** File */
    file: FileItem
    /** File Id */
    fileId: string
    /** Model Id (Base36) */
    id?: string
    /** Is Used as Avatar or Not */
    isAvatar: boolean
    /** No Create */
    noCreate: boolean
    /** No Delete */
    noDelete: boolean
    /** No Read */
    noRead: boolean
    /** No Update */
    noUpdate: boolean
    /** Updated At */
    updateAt: string
    /** User */
    user: UserItem
    /** User Id */
    userId: string
  }

  type fileUserPostDetailParams = {
    /** Model Id (Base36) */
    id: string
  }

  type FileUserSave = {
    /** A URL to the JSON Schema for this object. */
    $schema?: string
    /** File Id */
    fileId: string
    /** Model Id (Base36) */
    id?: string
    /** Is Used as Avatar or Not */
    isAvatar: boolean
    /** No Create */
    noCreate: boolean
    /** No Delete */
    noDelete: boolean
    /** No Read */
    noRead: boolean
    /** No Update */
    noUpdate: boolean
    /** User Id */
    userId: string
  }

  type FileUserSearchItem = {
    /** File */
    file: FileSearchItem
    /** File Id */
    fileId: string
    /** Model Id (Base36) */
    id?: string
    /** Is Used as Avatar or Not */
    isAvatar: boolean
    /** No Create */
    noCreate: boolean
    /** No Delete */
    noDelete: boolean
    /** No Read */
    noRead: boolean
    /** No Update */
    noUpdate: boolean
    /** User */
    user: UserSearchItem
    /** User Id */
    userId: string
  }

  type Filter = {
    /** Filter Field Name */
    field: string
    /** Or or And */
    or?: boolean
    /** Filter Condition Parameters */
    params?: any
    /** Filter Special */
    special?: number
    /** Filter Type */
    type: number
  }

  type groupDeleteDeleteParams = {
    /** Model Id (Base36) */
    id: string
  }

  type GroupDetail = {
    /** Code */
    code: string
    /** Created At */
    createAt: string
    /** Deleted At */
    deleteAt: string
    /** Description */
    description: string
    /** File Groups */
    fileGroups: any
    /** Files */
    files: any
    /** Group Roles */
    groupRoles: any
    /** Group Users */
    groupUsers: any
    /** Model Id (Base36) */
    id?: string
    /** Name */
    name: string
    /** Roles */
    roles: any
    /** Updated At */
    updateAt: string
    /** Users */
    users: any
  }

  type GroupItem = {
    /** Code */
    code: string
    /** Created At */
    createAt: string
    /** Deleted At */
    deleteAt: string
    /** Description */
    description: string
    /** Model Id (Base36) */
    id?: string
    /** Name */
    name: string
    /** Updated At */
    updateAt: string
  }

  type groupPostDetailParams = {
    /** Model Id (Base36) */
    id: string
  }

  type groupPutRecoverParams = {
    /** Model Id (Base36) */
    id: string
  }

  type groupRoleDeleteDeleteParams = {
    /** Model Id (Base36) */
    id: string
  }

  type GroupRoleDetail = {
    /** Created At */
    createAt: string
    /** Deleted At */
    deleteAt: string
    /** Group */
    group: GroupItem
    /** Group Id */
    groupId: string
    /** Model Id (Base36) */
    id?: string
    /** Role */
    role: RoleItem
    /** Role Id */
    roleId: string
    /** Updated At */
    updateAt: string
  }

  type GroupRoleItem = {
    /** Created At */
    createAt: string
    /** Deleted At */
    deleteAt: string
    /** Group */
    group: GroupItem
    /** Group Id */
    groupId: string
    /** Model Id (Base36) */
    id?: string
    /** Role */
    role: RoleItem
    /** Role Id */
    roleId: string
    /** Updated At */
    updateAt: string
  }

  type groupRolePostDetailParams = {
    /** Model Id (Base36) */
    id: string
  }

  type GroupRoleSave = {
    /** A URL to the JSON Schema for this object. */
    $schema?: string
    /** Group Id */
    groupId: string
    /** Model Id (Base36) */
    id?: string
    /** Role Id */
    roleId: string
  }

  type GroupRoleSearchItem = {
    /** Group */
    group: GroupSearchItem
    /** Group Id */
    groupId: string
    /** Model Id (Base36) */
    id?: string
    /** Role */
    role: RoleSearchItem
    /** Role Id */
    roleId: string
  }

  type GroupSave = {
    /** A URL to the JSON Schema for this object. */
    $schema?: string
    /** Code */
    code: string
    /** Description */
    description: string
    /** Model Id (Base36) */
    id?: string
    /** Name */
    name: string
  }

  type GroupSearchItem = {
    /** Code */
    code: string
    /** Model Id (Base36) */
    id?: string
    /** Name */
    name: string
  }

  type groupUserDeleteDeleteParams = {
    /** Model Id (Base36) */
    id: string
  }

  type GroupUserDetail = {
    /** Created At */
    createAt: string
    /** Deleted At */
    deleteAt: string
    /** Group */
    group: GroupItem
    /** Group Id */
    groupId: string
    /** Model Id (Base36) */
    id?: string
    /** Updated At */
    updateAt: string
    /** User */
    user: UserItem
    /** User Id */
    userId: string
  }

  type GroupUserItem = {
    /** Created At */
    createAt: string
    /** Deleted At */
    deleteAt: string
    /** Group */
    group: GroupItem
    /** Group Id */
    groupId: string
    /** Model Id (Base36) */
    id?: string
    /** Updated At */
    updateAt: string
    /** User */
    user: UserItem
    /** User Id */
    userId: string
  }

  type groupUserPostDetailParams = {
    /** Model Id (Base36) */
    id: string
  }

  type GroupUserSave = {
    /** A URL to the JSON Schema for this object. */
    $schema?: string
    /** Group Id */
    groupId: string
    /** Model Id (Base36) */
    id?: string
    /** User Id */
    userId: string
  }

  type GroupUserSearchItem = {
    /** Group */
    group: GroupSearchItem
    /** Group Id */
    groupId: string
    /** Model Id (Base36) */
    id?: string
    /** User */
    user: UserSearchItem
    /** User Id */
    userId: string
  }

  type Include = {
    /** Included Association Name */
    association: string
    /** Included Associations */
    includes?: any
    /** Omitted Fields */
    omits?: any
    /** Selected Fields */
    selects?: any
    /** Size Limit */
    size?: number
    /** List Sort Conditions */
    sorts?: any
  }

  type operationIdEnumPairGetListParams = {
    /** Filter by Operation ID */
    operationId?: string
  }

  type OperationIdEnumPairItem = {
    /** Enum Value */
    enum: number
    /** Operation ID */
    operationId: string
  }

  type operationIdEnumPairPostFromFlagParams = {
    /** Flag (Base64) */
    flag: string
  }

  type optionDeleteDeleteParams = {
    /** Model Id (Base36) */
    id: string
  }

  type OptionDetail = {
    /** Created At */
    createAt: string
    /** Deleted At */
    deleteAt: string
    /** Model Id (Base36) */
    id?: string
    /** Key */
    key: string
    /** Updated At */
    updateAt: string
    /** Value */
    value: string
  }

  type OptionItem = {
    /** Created At */
    createAt: string
    /** Deleted At */
    deleteAt: string
    /** Model Id (Base36) */
    id?: string
    /** Key */
    key: string
    /** Updated At */
    updateAt: string
    /** Value */
    value: string
  }

  type optionPostDetailParams = {
    /** Model Id (Base36) */
    id: string
  }

  type optionPutRecoverParams = {
    /** Model Id (Base36) */
    id: string
  }

  type OptionSave = {
    /** A URL to the JSON Schema for this object. */
    $schema?: string
    /** Model Id (Base36) */
    id?: string
    /** Key */
    key: string
    /** Value */
    value: string
  }

  type OptionSearchItem = {
    /** Model Id (Base36) */
    id?: string
    /** Key */
    key: string
    /** Value */
    value: string
  }

  type permissionDeleteDeleteParams = {
    /** Model Id (Base36) */
    id: string
  }

  type PermissionDetail = {
    /** Code */
    code: string
    /** Created At */
    createAt: string
    /** Deleted At */
    deleteAt: string
    /** Description */
    description: string
    /** Flag (Base64) */
    flagBase64: string
    /** Model Id (Base36) */
    id?: string
    /** Name */
    name: string
    /** Role Permissions */
    rolePermissions: any
    /** Roles */
    roles: any
    /** Updated At */
    updateAt: string
  }

  type PermissionItem = {
    /** Code */
    code: string
    /** Created At */
    createAt: string
    /** Deleted At */
    deleteAt: string
    /** Description */
    description: string
    /** Model Id (Base36) */
    id?: string
    /** Name */
    name: string
    /** Updated At */
    updateAt: string
  }

  type permissionPostDetailParams = {
    /** Model Id (Base36) */
    id: string
  }

  type permissionPutRecoverParams = {
    /** Model Id (Base36) */
    id: string
  }

  type PermissionSave = {
    /** A URL to the JSON Schema for this object. */
    $schema?: string
    /** Code */
    code: string
    /** Description */
    description: string
    /** Flag (Base64) */
    flagBase64: string
    /** Model Id (Base36) */
    id?: string
    /** Name */
    name: string
  }

  type PermissionSearchItem = {
    /** Code */
    code: string
    /** Model Id (Base36) */
    id?: string
    /** Name */
    name: string
  }

  type ResponseAuthLoginBody = {
    /** A URL to the JSON Schema for this object. */
    $schema?: string
    /** 0 => success; Otherwise, fail */
    code: number
    /** Response Payload */
    data: AuthLogin
    /** Hint Message */
    message: string
    /** Has the action done successfully? */
    success: boolean
    /** If request is list query, total represents the count of the list w/o any user query conditions but conditions built-in; Or represents affected rows for other requests */
    total: number
  }

  type ResponseCurrentUserBody = {
    /** A URL to the JSON Schema for this object. */
    $schema?: string
    /** 0 => success; Otherwise, fail */
    code: number
    /** Response Payload */
    data: CurrentUser
    /** Hint Message */
    message: string
    /** Has the action done successfully? */
    success: boolean
    /** If request is list query, total represents the count of the list w/o any user query conditions but conditions built-in; Or represents affected rows for other requests */
    total: number
  }

  type ResponseFileCategoryDetailBody = {
    /** A URL to the JSON Schema for this object. */
    $schema?: string
    /** 0 => success; Otherwise, fail */
    code: number
    /** Response Payload */
    data: FileCategoryDetail
    /** Hint Message */
    message: string
    /** Has the action done successfully? */
    success: boolean
    /** If request is list query, total represents the count of the list w/o any user query conditions but conditions built-in; Or represents affected rows for other requests */
    total: number
  }

  type ResponseFileCategorySaveBody = {
    /** A URL to the JSON Schema for this object. */
    $schema?: string
    /** 0 => success; Otherwise, fail */
    code: number
    /** Response Payload */
    data: FileCategorySave
    /** Hint Message */
    message: string
    /** Has the action done successfully? */
    success: boolean
    /** If request is list query, total represents the count of the list w/o any user query conditions but conditions built-in; Or represents affected rows for other requests */
    total: number
  }

  type ResponseFileDetailBody = {
    /** A URL to the JSON Schema for this object. */
    $schema?: string
    /** 0 => success; Otherwise, fail */
    code: number
    /** Response Payload */
    data: FileDetail
    /** Hint Message */
    message: string
    /** Has the action done successfully? */
    success: boolean
    /** If request is list query, total represents the count of the list w/o any user query conditions but conditions built-in; Or represents affected rows for other requests */
    total: number
  }

  type ResponseFileExtensionDetailBody = {
    /** A URL to the JSON Schema for this object. */
    $schema?: string
    /** 0 => success; Otherwise, fail */
    code: number
    /** Response Payload */
    data: FileExtensionDetail
    /** Hint Message */
    message: string
    /** Has the action done successfully? */
    success: boolean
    /** If request is list query, total represents the count of the list w/o any user query conditions but conditions built-in; Or represents affected rows for other requests */
    total: number
  }

  type ResponseFileExtensionSaveBody = {
    /** A URL to the JSON Schema for this object. */
    $schema?: string
    /** 0 => success; Otherwise, fail */
    code: number
    /** Response Payload */
    data: FileExtensionSave
    /** Hint Message */
    message: string
    /** Has the action done successfully? */
    success: boolean
    /** If request is list query, total represents the count of the list w/o any user query conditions but conditions built-in; Or represents affected rows for other requests */
    total: number
  }

  type ResponseFileGroupDetailBody = {
    /** A URL to the JSON Schema for this object. */
    $schema?: string
    /** 0 => success; Otherwise, fail */
    code: number
    /** Response Payload */
    data: FileGroupDetail
    /** Hint Message */
    message: string
    /** Has the action done successfully? */
    success: boolean
    /** If request is list query, total represents the count of the list w/o any user query conditions but conditions built-in; Or represents affected rows for other requests */
    total: number
  }

  type ResponseFileGroupSaveBody = {
    /** A URL to the JSON Schema for this object. */
    $schema?: string
    /** 0 => success; Otherwise, fail */
    code: number
    /** Response Payload */
    data: FileGroupSave
    /** Hint Message */
    message: string
    /** Has the action done successfully? */
    success: boolean
    /** If request is list query, total represents the count of the list w/o any user query conditions but conditions built-in; Or represents affected rows for other requests */
    total: number
  }

  type ResponseFileSaveBody = {
    /** A URL to the JSON Schema for this object. */
    $schema?: string
    /** 0 => success; Otherwise, fail */
    code: number
    /** Response Payload */
    data: FileSave
    /** Hint Message */
    message: string
    /** Has the action done successfully? */
    success: boolean
    /** If request is list query, total represents the count of the list w/o any user query conditions but conditions built-in; Or represents affected rows for other requests */
    total: number
  }

  type ResponseFileUserDetailBody = {
    /** A URL to the JSON Schema for this object. */
    $schema?: string
    /** 0 => success; Otherwise, fail */
    code: number
    /** Response Payload */
    data: FileUserDetail
    /** Hint Message */
    message: string
    /** Has the action done successfully? */
    success: boolean
    /** If request is list query, total represents the count of the list w/o any user query conditions but conditions built-in; Or represents affected rows for other requests */
    total: number
  }

  type ResponseFileUserSaveBody = {
    /** A URL to the JSON Schema for this object. */
    $schema?: string
    /** 0 => success; Otherwise, fail */
    code: number
    /** Response Payload */
    data: FileUserSave
    /** Hint Message */
    message: string
    /** Has the action done successfully? */
    success: boolean
    /** If request is list query, total represents the count of the list w/o any user query conditions but conditions built-in; Or represents affected rows for other requests */
    total: number
  }

  type ResponseGroupDetailBody = {
    /** A URL to the JSON Schema for this object. */
    $schema?: string
    /** 0 => success; Otherwise, fail */
    code: number
    /** Response Payload */
    data: GroupDetail
    /** Hint Message */
    message: string
    /** Has the action done successfully? */
    success: boolean
    /** If request is list query, total represents the count of the list w/o any user query conditions but conditions built-in; Or represents affected rows for other requests */
    total: number
  }

  type ResponseGroupRoleDetailBody = {
    /** A URL to the JSON Schema for this object. */
    $schema?: string
    /** 0 => success; Otherwise, fail */
    code: number
    /** Response Payload */
    data: GroupRoleDetail
    /** Hint Message */
    message: string
    /** Has the action done successfully? */
    success: boolean
    /** If request is list query, total represents the count of the list w/o any user query conditions but conditions built-in; Or represents affected rows for other requests */
    total: number
  }

  type ResponseGroupRoleSaveBody = {
    /** A URL to the JSON Schema for this object. */
    $schema?: string
    /** 0 => success; Otherwise, fail */
    code: number
    /** Response Payload */
    data: GroupRoleSave
    /** Hint Message */
    message: string
    /** Has the action done successfully? */
    success: boolean
    /** If request is list query, total represents the count of the list w/o any user query conditions but conditions built-in; Or represents affected rows for other requests */
    total: number
  }

  type ResponseGroupSaveBody = {
    /** A URL to the JSON Schema for this object. */
    $schema?: string
    /** 0 => success; Otherwise, fail */
    code: number
    /** Response Payload */
    data: GroupSave
    /** Hint Message */
    message: string
    /** Has the action done successfully? */
    success: boolean
    /** If request is list query, total represents the count of the list w/o any user query conditions but conditions built-in; Or represents affected rows for other requests */
    total: number
  }

  type ResponseGroupUserDetailBody = {
    /** A URL to the JSON Schema for this object. */
    $schema?: string
    /** 0 => success; Otherwise, fail */
    code: number
    /** Response Payload */
    data: GroupUserDetail
    /** Hint Message */
    message: string
    /** Has the action done successfully? */
    success: boolean
    /** If request is list query, total represents the count of the list w/o any user query conditions but conditions built-in; Or represents affected rows for other requests */
    total: number
  }

  type ResponseGroupUserSaveBody = {
    /** A URL to the JSON Schema for this object. */
    $schema?: string
    /** 0 => success; Otherwise, fail */
    code: number
    /** Response Payload */
    data: GroupUserSave
    /** Hint Message */
    message: string
    /** Has the action done successfully? */
    success: boolean
    /** If request is list query, total represents the count of the list w/o any user query conditions but conditions built-in; Or represents affected rows for other requests */
    total: number
  }

  type ResponseListEndpointsItemBody = {
    /** A URL to the JSON Schema for this object. */
    $schema?: string
    /** 0 => success; Otherwise, fail */
    code: number
    /** Response Payload */
    data: any
    /** Hint Message */
    message: string
    /** Has the action done successfully? */
    success: boolean
    /** If request is list query, total represents the count of the list w/o any user query conditions but conditions built-in; Or represents affected rows for other requests */
    total: number
  }

  type ResponseListFileCategoryItemBody = {
    /** A URL to the JSON Schema for this object. */
    $schema?: string
    /** 0 => success; Otherwise, fail */
    code: number
    /** Response Payload */
    data: any
    /** Hint Message */
    message: string
    /** Has the action done successfully? */
    success: boolean
    /** If request is list query, total represents the count of the list w/o any user query conditions but conditions built-in; Or represents affected rows for other requests */
    total: number
  }

  type ResponseListFileCategorySearchItemBody = {
    /** A URL to the JSON Schema for this object. */
    $schema?: string
    /** 0 => success; Otherwise, fail */
    code: number
    /** Response Payload */
    data: any
    /** Hint Message */
    message: string
    /** Has the action done successfully? */
    success: boolean
    /** If request is list query, total represents the count of the list w/o any user query conditions but conditions built-in; Or represents affected rows for other requests */
    total: number
  }

  type ResponseListFileExtensionItemBody = {
    /** A URL to the JSON Schema for this object. */
    $schema?: string
    /** 0 => success; Otherwise, fail */
    code: number
    /** Response Payload */
    data: any
    /** Hint Message */
    message: string
    /** Has the action done successfully? */
    success: boolean
    /** If request is list query, total represents the count of the list w/o any user query conditions but conditions built-in; Or represents affected rows for other requests */
    total: number
  }

  type ResponseListFileExtensionSearchItemBody = {
    /** A URL to the JSON Schema for this object. */
    $schema?: string
    /** 0 => success; Otherwise, fail */
    code: number
    /** Response Payload */
    data: any
    /** Hint Message */
    message: string
    /** Has the action done successfully? */
    success: boolean
    /** If request is list query, total represents the count of the list w/o any user query conditions but conditions built-in; Or represents affected rows for other requests */
    total: number
  }

  type ResponseListFileGroupItemBody = {
    /** A URL to the JSON Schema for this object. */
    $schema?: string
    /** 0 => success; Otherwise, fail */
    code: number
    /** Response Payload */
    data: any
    /** Hint Message */
    message: string
    /** Has the action done successfully? */
    success: boolean
    /** If request is list query, total represents the count of the list w/o any user query conditions but conditions built-in; Or represents affected rows for other requests */
    total: number
  }

  type ResponseListFileGroupSearchItemBody = {
    /** A URL to the JSON Schema for this object. */
    $schema?: string
    /** 0 => success; Otherwise, fail */
    code: number
    /** Response Payload */
    data: any
    /** Hint Message */
    message: string
    /** Has the action done successfully? */
    success: boolean
    /** If request is list query, total represents the count of the list w/o any user query conditions but conditions built-in; Or represents affected rows for other requests */
    total: number
  }

  type ResponseListFileItemBody = {
    /** A URL to the JSON Schema for this object. */
    $schema?: string
    /** 0 => success; Otherwise, fail */
    code: number
    /** Response Payload */
    data: any
    /** Hint Message */
    message: string
    /** Has the action done successfully? */
    success: boolean
    /** If request is list query, total represents the count of the list w/o any user query conditions but conditions built-in; Or represents affected rows for other requests */
    total: number
  }

  type ResponseListFileSearchItemBody = {
    /** A URL to the JSON Schema for this object. */
    $schema?: string
    /** 0 => success; Otherwise, fail */
    code: number
    /** Response Payload */
    data: any
    /** Hint Message */
    message: string
    /** Has the action done successfully? */
    success: boolean
    /** If request is list query, total represents the count of the list w/o any user query conditions but conditions built-in; Or represents affected rows for other requests */
    total: number
  }

  type ResponseListFileUserItemBody = {
    /** A URL to the JSON Schema for this object. */
    $schema?: string
    /** 0 => success; Otherwise, fail */
    code: number
    /** Response Payload */
    data: any
    /** Hint Message */
    message: string
    /** Has the action done successfully? */
    success: boolean
    /** If request is list query, total represents the count of the list w/o any user query conditions but conditions built-in; Or represents affected rows for other requests */
    total: number
  }

  type ResponseListFileUserSearchItemBody = {
    /** A URL to the JSON Schema for this object. */
    $schema?: string
    /** 0 => success; Otherwise, fail */
    code: number
    /** Response Payload */
    data: any
    /** Hint Message */
    message: string
    /** Has the action done successfully? */
    success: boolean
    /** If request is list query, total represents the count of the list w/o any user query conditions but conditions built-in; Or represents affected rows for other requests */
    total: number
  }

  type ResponseListGroupItemBody = {
    /** A URL to the JSON Schema for this object. */
    $schema?: string
    /** 0 => success; Otherwise, fail */
    code: number
    /** Response Payload */
    data: any
    /** Hint Message */
    message: string
    /** Has the action done successfully? */
    success: boolean
    /** If request is list query, total represents the count of the list w/o any user query conditions but conditions built-in; Or represents affected rows for other requests */
    total: number
  }

  type ResponseListGroupRoleItemBody = {
    /** A URL to the JSON Schema for this object. */
    $schema?: string
    /** 0 => success; Otherwise, fail */
    code: number
    /** Response Payload */
    data: any
    /** Hint Message */
    message: string
    /** Has the action done successfully? */
    success: boolean
    /** If request is list query, total represents the count of the list w/o any user query conditions but conditions built-in; Or represents affected rows for other requests */
    total: number
  }

  type ResponseListGroupRoleSearchItemBody = {
    /** A URL to the JSON Schema for this object. */
    $schema?: string
    /** 0 => success; Otherwise, fail */
    code: number
    /** Response Payload */
    data: any
    /** Hint Message */
    message: string
    /** Has the action done successfully? */
    success: boolean
    /** If request is list query, total represents the count of the list w/o any user query conditions but conditions built-in; Or represents affected rows for other requests */
    total: number
  }

  type ResponseListGroupSearchItemBody = {
    /** A URL to the JSON Schema for this object. */
    $schema?: string
    /** 0 => success; Otherwise, fail */
    code: number
    /** Response Payload */
    data: any
    /** Hint Message */
    message: string
    /** Has the action done successfully? */
    success: boolean
    /** If request is list query, total represents the count of the list w/o any user query conditions but conditions built-in; Or represents affected rows for other requests */
    total: number
  }

  type ResponseListGroupUserItemBody = {
    /** A URL to the JSON Schema for this object. */
    $schema?: string
    /** 0 => success; Otherwise, fail */
    code: number
    /** Response Payload */
    data: any
    /** Hint Message */
    message: string
    /** Has the action done successfully? */
    success: boolean
    /** If request is list query, total represents the count of the list w/o any user query conditions but conditions built-in; Or represents affected rows for other requests */
    total: number
  }

  type ResponseListGroupUserSearchItemBody = {
    /** A URL to the JSON Schema for this object. */
    $schema?: string
    /** 0 => success; Otherwise, fail */
    code: number
    /** Response Payload */
    data: any
    /** Hint Message */
    message: string
    /** Has the action done successfully? */
    success: boolean
    /** If request is list query, total represents the count of the list w/o any user query conditions but conditions built-in; Or represents affected rows for other requests */
    total: number
  }

  type ResponseListOperationIdEnumPairItemBody = {
    /** A URL to the JSON Schema for this object. */
    $schema?: string
    /** 0 => success; Otherwise, fail */
    code: number
    /** Response Payload */
    data: any
    /** Hint Message */
    message: string
    /** Has the action done successfully? */
    success: boolean
    /** If request is list query, total represents the count of the list w/o any user query conditions but conditions built-in; Or represents affected rows for other requests */
    total: number
  }

  type ResponseListOptionItemBody = {
    /** A URL to the JSON Schema for this object. */
    $schema?: string
    /** 0 => success; Otherwise, fail */
    code: number
    /** Response Payload */
    data: any
    /** Hint Message */
    message: string
    /** Has the action done successfully? */
    success: boolean
    /** If request is list query, total represents the count of the list w/o any user query conditions but conditions built-in; Or represents affected rows for other requests */
    total: number
  }

  type ResponseListOptionSearchItemBody = {
    /** A URL to the JSON Schema for this object. */
    $schema?: string
    /** 0 => success; Otherwise, fail */
    code: number
    /** Response Payload */
    data: any
    /** Hint Message */
    message: string
    /** Has the action done successfully? */
    success: boolean
    /** If request is list query, total represents the count of the list w/o any user query conditions but conditions built-in; Or represents affected rows for other requests */
    total: number
  }

  type ResponseListPermissionItemBody = {
    /** A URL to the JSON Schema for this object. */
    $schema?: string
    /** 0 => success; Otherwise, fail */
    code: number
    /** Response Payload */
    data: any
    /** Hint Message */
    message: string
    /** Has the action done successfully? */
    success: boolean
    /** If request is list query, total represents the count of the list w/o any user query conditions but conditions built-in; Or represents affected rows for other requests */
    total: number
  }

  type ResponseListPermissionSearchItemBody = {
    /** A URL to the JSON Schema for this object. */
    $schema?: string
    /** 0 => success; Otherwise, fail */
    code: number
    /** Response Payload */
    data: any
    /** Hint Message */
    message: string
    /** Has the action done successfully? */
    success: boolean
    /** If request is list query, total represents the count of the list w/o any user query conditions but conditions built-in; Or represents affected rows for other requests */
    total: number
  }

  type ResponseListRoleItemBody = {
    /** A URL to the JSON Schema for this object. */
    $schema?: string
    /** 0 => success; Otherwise, fail */
    code: number
    /** Response Payload */
    data: any
    /** Hint Message */
    message: string
    /** Has the action done successfully? */
    success: boolean
    /** If request is list query, total represents the count of the list w/o any user query conditions but conditions built-in; Or represents affected rows for other requests */
    total: number
  }

  type ResponseListRolePermissionItemBody = {
    /** A URL to the JSON Schema for this object. */
    $schema?: string
    /** 0 => success; Otherwise, fail */
    code: number
    /** Response Payload */
    data: any
    /** Hint Message */
    message: string
    /** Has the action done successfully? */
    success: boolean
    /** If request is list query, total represents the count of the list w/o any user query conditions but conditions built-in; Or represents affected rows for other requests */
    total: number
  }

  type ResponseListRolePermissionSearchItemBody = {
    /** A URL to the JSON Schema for this object. */
    $schema?: string
    /** 0 => success; Otherwise, fail */
    code: number
    /** Response Payload */
    data: any
    /** Hint Message */
    message: string
    /** Has the action done successfully? */
    success: boolean
    /** If request is list query, total represents the count of the list w/o any user query conditions but conditions built-in; Or represents affected rows for other requests */
    total: number
  }

  type ResponseListRoleSearchItemBody = {
    /** A URL to the JSON Schema for this object. */
    $schema?: string
    /** 0 => success; Otherwise, fail */
    code: number
    /** Response Payload */
    data: any
    /** Hint Message */
    message: string
    /** Has the action done successfully? */
    success: boolean
    /** If request is list query, total represents the count of the list w/o any user query conditions but conditions built-in; Or represents affected rows for other requests */
    total: number
  }

  type ResponseListStringBody = {
    /** A URL to the JSON Schema for this object. */
    $schema?: string
    /** 0 => success; Otherwise, fail */
    code: number
    /** Response Payload */
    data: any
    /** Hint Message */
    message: string
    /** Has the action done successfully? */
    success: boolean
    /** If request is list query, total represents the count of the list w/o any user query conditions but conditions built-in; Or represents affected rows for other requests */
    total: number
  }

  type ResponseListUserItemBody = {
    /** A URL to the JSON Schema for this object. */
    $schema?: string
    /** 0 => success; Otherwise, fail */
    code: number
    /** Response Payload */
    data: any
    /** Hint Message */
    message: string
    /** Has the action done successfully? */
    success: boolean
    /** If request is list query, total represents the count of the list w/o any user query conditions but conditions built-in; Or represents affected rows for other requests */
    total: number
  }

  type ResponseListUserPasswordItemBody = {
    /** A URL to the JSON Schema for this object. */
    $schema?: string
    /** 0 => success; Otherwise, fail */
    code: number
    /** Response Payload */
    data: any
    /** Hint Message */
    message: string
    /** Has the action done successfully? */
    success: boolean
    /** If request is list query, total represents the count of the list w/o any user query conditions but conditions built-in; Or represents affected rows for other requests */
    total: number
  }

  type ResponseListUserPasswordSearchItemBody = {
    /** A URL to the JSON Schema for this object. */
    $schema?: string
    /** 0 => success; Otherwise, fail */
    code: number
    /** Response Payload */
    data: any
    /** Hint Message */
    message: string
    /** Has the action done successfully? */
    success: boolean
    /** If request is list query, total represents the count of the list w/o any user query conditions but conditions built-in; Or represents affected rows for other requests */
    total: number
  }

  type ResponseListUserRoleItemBody = {
    /** A URL to the JSON Schema for this object. */
    $schema?: string
    /** 0 => success; Otherwise, fail */
    code: number
    /** Response Payload */
    data: any
    /** Hint Message */
    message: string
    /** Has the action done successfully? */
    success: boolean
    /** If request is list query, total represents the count of the list w/o any user query conditions but conditions built-in; Or represents affected rows for other requests */
    total: number
  }

  type ResponseListUserRoleSearchItemBody = {
    /** A URL to the JSON Schema for this object. */
    $schema?: string
    /** 0 => success; Otherwise, fail */
    code: number
    /** Response Payload */
    data: any
    /** Hint Message */
    message: string
    /** Has the action done successfully? */
    success: boolean
    /** If request is list query, total represents the count of the list w/o any user query conditions but conditions built-in; Or represents affected rows for other requests */
    total: number
  }

  type ResponseListUserSearchItemBody = {
    /** A URL to the JSON Schema for this object. */
    $schema?: string
    /** 0 => success; Otherwise, fail */
    code: number
    /** Response Payload */
    data: any
    /** Hint Message */
    message: string
    /** Has the action done successfully? */
    success: boolean
    /** If request is list query, total represents the count of the list w/o any user query conditions but conditions built-in; Or represents affected rows for other requests */
    total: number
  }

  type ResponseOptionDetailBody = {
    /** A URL to the JSON Schema for this object. */
    $schema?: string
    /** 0 => success; Otherwise, fail */
    code: number
    /** Response Payload */
    data: OptionDetail
    /** Hint Message */
    message: string
    /** Has the action done successfully? */
    success: boolean
    /** If request is list query, total represents the count of the list w/o any user query conditions but conditions built-in; Or represents affected rows for other requests */
    total: number
  }

  type ResponseOptionSaveBody = {
    /** A URL to the JSON Schema for this object. */
    $schema?: string
    /** 0 => success; Otherwise, fail */
    code: number
    /** Response Payload */
    data: OptionSave
    /** Hint Message */
    message: string
    /** Has the action done successfully? */
    success: boolean
    /** If request is list query, total represents the count of the list w/o any user query conditions but conditions built-in; Or represents affected rows for other requests */
    total: number
  }

  type ResponsePermissionDetailBody = {
    /** A URL to the JSON Schema for this object. */
    $schema?: string
    /** 0 => success; Otherwise, fail */
    code: number
    /** Response Payload */
    data: PermissionDetail
    /** Hint Message */
    message: string
    /** Has the action done successfully? */
    success: boolean
    /** If request is list query, total represents the count of the list w/o any user query conditions but conditions built-in; Or represents affected rows for other requests */
    total: number
  }

  type ResponsePermissionSaveBody = {
    /** A URL to the JSON Schema for this object. */
    $schema?: string
    /** 0 => success; Otherwise, fail */
    code: number
    /** Response Payload */
    data: PermissionSave
    /** Hint Message */
    message: string
    /** Has the action done successfully? */
    success: boolean
    /** If request is list query, total represents the count of the list w/o any user query conditions but conditions built-in; Or represents affected rows for other requests */
    total: number
  }

  type ResponseRoleDetailBody = {
    /** A URL to the JSON Schema for this object. */
    $schema?: string
    /** 0 => success; Otherwise, fail */
    code: number
    /** Response Payload */
    data: RoleDetail
    /** Hint Message */
    message: string
    /** Has the action done successfully? */
    success: boolean
    /** If request is list query, total represents the count of the list w/o any user query conditions but conditions built-in; Or represents affected rows for other requests */
    total: number
  }

  type ResponseRolePermissionDetailBody = {
    /** A URL to the JSON Schema for this object. */
    $schema?: string
    /** 0 => success; Otherwise, fail */
    code: number
    /** Response Payload */
    data: RolePermissionDetail
    /** Hint Message */
    message: string
    /** Has the action done successfully? */
    success: boolean
    /** If request is list query, total represents the count of the list w/o any user query conditions but conditions built-in; Or represents affected rows for other requests */
    total: number
  }

  type ResponseRolePermissionSaveBody = {
    /** A URL to the JSON Schema for this object. */
    $schema?: string
    /** 0 => success; Otherwise, fail */
    code: number
    /** Response Payload */
    data: RolePermissionSave
    /** Hint Message */
    message: string
    /** Has the action done successfully? */
    success: boolean
    /** If request is list query, total represents the count of the list w/o any user query conditions but conditions built-in; Or represents affected rows for other requests */
    total: number
  }

  type ResponseRoleSaveBody = {
    /** A URL to the JSON Schema for this object. */
    $schema?: string
    /** 0 => success; Otherwise, fail */
    code: number
    /** Response Payload */
    data: RoleSave
    /** Hint Message */
    message: string
    /** Has the action done successfully? */
    success: boolean
    /** If request is list query, total represents the count of the list w/o any user query conditions but conditions built-in; Or represents affected rows for other requests */
    total: number
  }

  type ResponseStringBody = {
    /** A URL to the JSON Schema for this object. */
    $schema?: string
    /** 0 => success; Otherwise, fail */
    code: number
    /** Response Payload */
    data: string
    /** Hint Message */
    message: string
    /** Has the action done successfully? */
    success: boolean
    /** If request is list query, total represents the count of the list w/o any user query conditions but conditions built-in; Or represents affected rows for other requests */
    total: number
  }

  type ResponseUserDetailBody = {
    /** A URL to the JSON Schema for this object. */
    $schema?: string
    /** 0 => success; Otherwise, fail */
    code: number
    /** Response Payload */
    data: UserDetail
    /** Hint Message */
    message: string
    /** Has the action done successfully? */
    success: boolean
    /** If request is list query, total represents the count of the list w/o any user query conditions but conditions built-in; Or represents affected rows for other requests */
    total: number
  }

  type ResponseUserPasswordDetailBody = {
    /** A URL to the JSON Schema for this object. */
    $schema?: string
    /** 0 => success; Otherwise, fail */
    code: number
    /** Response Payload */
    data: UserPasswordDetail
    /** Hint Message */
    message: string
    /** Has the action done successfully? */
    success: boolean
    /** If request is list query, total represents the count of the list w/o any user query conditions but conditions built-in; Or represents affected rows for other requests */
    total: number
  }

  type ResponseUserPasswordSaveBody = {
    /** A URL to the JSON Schema for this object. */
    $schema?: string
    /** 0 => success; Otherwise, fail */
    code: number
    /** Response Payload */
    data: UserPasswordSave
    /** Hint Message */
    message: string
    /** Has the action done successfully? */
    success: boolean
    /** If request is list query, total represents the count of the list w/o any user query conditions but conditions built-in; Or represents affected rows for other requests */
    total: number
  }

  type ResponseUserRoleDetailBody = {
    /** A URL to the JSON Schema for this object. */
    $schema?: string
    /** 0 => success; Otherwise, fail */
    code: number
    /** Response Payload */
    data: UserRoleDetail
    /** Hint Message */
    message: string
    /** Has the action done successfully? */
    success: boolean
    /** If request is list query, total represents the count of the list w/o any user query conditions but conditions built-in; Or represents affected rows for other requests */
    total: number
  }

  type ResponseUserRoleSaveBody = {
    /** A URL to the JSON Schema for this object. */
    $schema?: string
    /** 0 => success; Otherwise, fail */
    code: number
    /** Response Payload */
    data: UserRoleSave
    /** Hint Message */
    message: string
    /** Has the action done successfully? */
    success: boolean
    /** If request is list query, total represents the count of the list w/o any user query conditions but conditions built-in; Or represents affected rows for other requests */
    total: number
  }

  type ResponseUserSaveBody = {
    /** A URL to the JSON Schema for this object. */
    $schema?: string
    /** 0 => success; Otherwise, fail */
    code: number
    /** Response Payload */
    data: UserSave
    /** Hint Message */
    message: string
    /** Has the action done successfully? */
    success: boolean
    /** If request is list query, total represents the count of the list w/o any user query conditions but conditions built-in; Or represents affected rows for other requests */
    total: number
  }

  type roleDeleteDeleteParams = {
    /** Model Id (Base36) */
    id: string
  }

  type RoleDetail = {
    /** Code */
    code: string
    /** Created At */
    createAt: string
    /** Deleted At */
    deleteAt: string
    /** Description */
    description: string
    /** Group Roles */
    groupRoles: any
    /** Groups */
    groups: any
    /** Model Id (Base36) */
    id?: string
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

  type RoleItem = {
    /** Code */
    code: string
    /** Created At */
    createAt: string
    /** Deleted At */
    deleteAt: string
    /** Description */
    description: string
    /** Model Id (Base36) */
    id?: string
    /** Name */
    name: string
    /** Updated At */
    updateAt: string
  }

  type rolePermissionDeleteDeleteParams = {
    /** Model Id (Base36) */
    id: string
  }

  type RolePermissionDetail = {
    /** Created At */
    createAt: string
    /** Deleted At */
    deleteAt: string
    /** Model Id (Base36) */
    id?: string
    /** Permission */
    permission: PermissionItem
    /** Permission Id */
    permissionId: string
    /** Role */
    role: RoleItem
    /** Role Id */
    roleId: string
    /** Updated At */
    updateAt: string
  }

  type RolePermissionItem = {
    /** Created At */
    createAt: string
    /** Deleted At */
    deleteAt: string
    /** Model Id (Base36) */
    id?: string
    /** Permission */
    permission: PermissionItem
    /** Permission Id */
    permissionId: string
    /** Role */
    role: RoleItem
    /** Role Id */
    roleId: string
    /** Updated At */
    updateAt: string
  }

  type rolePermissionPostDetailParams = {
    /** Model Id (Base36) */
    id: string
  }

  type RolePermissionSave = {
    /** A URL to the JSON Schema for this object. */
    $schema?: string
    /** Model Id (Base36) */
    id?: string
    /** Permission Id */
    permissionId: string
    /** Role Id */
    roleId: string
  }

  type RolePermissionSearchItem = {
    /** Model Id (Base36) */
    id?: string
    /** Permission */
    permission: PermissionSearchItem
    /** Permission Id */
    permissionId: string
    /** Role */
    role: RoleSearchItem
    /** Role Id */
    roleId: string
  }

  type rolePostDetailParams = {
    /** Model Id (Base36) */
    id: string
  }

  type rolePutRecoverParams = {
    /** Model Id (Base36) */
    id: string
  }

  type RoleSave = {
    /** A URL to the JSON Schema for this object. */
    $schema?: string
    /** Code */
    code: string
    /** Description */
    description: string
    /** Model Id (Base36) */
    id?: string
    /** Name */
    name: string
  }

  type RoleSearchItem = {
    /** Code */
    code: string
    /** Model Id (Base36) */
    id?: string
    /** Name */
    name: string
  }

  type Sort = {
    /** Desc or Asc */
    desc?: boolean
    /** Sorted Field Name */
    field: string
  }

  type userDeleteDeleteParams = {
    /** Model Id (Base36) */
    id: string
  }

  type UserDetail = {
    /** Account */
    account: string
    /** Created At */
    createAt: string
    /** Deleted At */
    deleteAt: string
    /** Description */
    description: string
    /** File Users */
    fileUsers: any
    /** Files */
    files: any
    /** Group Users */
    groupUsers: any
    /** Groups */
    groups: any
    /** Model Id (Base36) */
    id?: string
    /** Name */
    name: string
    /** Password */
    password: string
    /** Roles */
    roles: any
    /** Updated At */
    updateAt: string
    /** User Passwords */
    userPasswords: any
    /** User Roles */
    userRoles: any
  }

  type UserItem = {
    /** Account */
    account: string
    /** Created At */
    createAt: string
    /** Deleted At */
    deleteAt: string
    /** Description */
    description: string
    /** Model Id (Base36) */
    id?: string
    /** Name */
    name: string
    /** Password */
    password: string
    /** Updated At */
    updateAt: string
  }

  type userPasswordDeleteDeleteParams = {
    /** Model Id (Base36) */
    id: string
  }

  type UserPasswordDetail = {
    /** Created At */
    createAt: string
    /** Deleted At */
    deleteAt: string
    /** Model Id (Base36) */
    id?: string
    /** Password */
    password: string
    /** Updated At */
    updateAt: string
    /** User */
    user: UserItem
    /** User Id */
    userId: string
  }

  type UserPasswordItem = {
    /** Created At */
    createAt: string
    /** Deleted At */
    deleteAt: string
    /** Model Id (Base36) */
    id?: string
    /** Password */
    password: string
    /** Updated At */
    updateAt: string
    /** User */
    user: UserItem
    /** User Id */
    userId: string
  }

  type userPasswordPostDetailParams = {
    /** Model Id (Base36) */
    id: string
  }

  type UserPasswordSave = {
    /** A URL to the JSON Schema for this object. */
    $schema?: string
    /** Model Id (Base36) */
    id?: string
    /** Password */
    password: string
    /** User Id */
    userId: string
  }

  type UserPasswordSearchItem = {
    /** Model Id (Base36) */
    id?: string
    /** Password */
    password: string
    /** User */
    user: UserSearchItem
    /** User Id */
    userId: string
  }

  type userPostDetailParams = {
    /** Model Id (Base36) */
    id: string
  }

  type userPutRecoverParams = {
    /** Model Id (Base36) */
    id: string
  }

  type userRoleDeleteDeleteParams = {
    /** Model Id (Base36) */
    id: string
  }

  type UserRoleDetail = {
    /** Created At */
    createAt: string
    /** Deleted At */
    deleteAt: string
    /** Model Id (Base36) */
    id?: string
    /** Is Default Role */
    isDefault: boolean
    /** Role */
    role: RoleItem
    /** Role Id */
    roleId: string
    /** Updated At */
    updateAt: string
    /** User */
    user: UserItem
    /** User Id */
    userId: string
  }

  type UserRoleItem = {
    /** Created At */
    createAt: string
    /** Deleted At */
    deleteAt: string
    /** Model Id (Base36) */
    id?: string
    /** Is Default Role */
    isDefault: boolean
    /** Role */
    role: RoleItem
    /** Role Id */
    roleId: string
    /** Updated At */
    updateAt: string
    /** User */
    user: UserItem
    /** User Id */
    userId: string
  }

  type userRolePostDetailParams = {
    /** Model Id (Base36) */
    id: string
  }

  type UserRoleSave = {
    /** A URL to the JSON Schema for this object. */
    $schema?: string
    /** Model Id (Base36) */
    id?: string
    /** Is Default Role */
    isDefault: boolean
    /** Role Id */
    roleId: string
    /** User Id */
    userId: string
  }

  type UserRoleSearchItem = {
    /** Model Id (Base36) */
    id?: string
    /** Is Default Role */
    isDefault: boolean
    /** Role */
    role: RoleSearchItem
    /** Role Id */
    roleId: string
    /** User */
    user: UserSearchItem
    /** User Id */
    userId: string
  }

  type UserSave = {
    /** A URL to the JSON Schema for this object. */
    $schema?: string
    /** Account */
    account: string
    /** Description */
    description: string
    /** Model Id (Base36) */
    id?: string
    /** Name */
    name: string
    /** Password */
    password: string
  }

  type UserSearchItem = {
    /** Account */
    account: string
    /** Model Id (Base36) */
    id?: string
    /** Name */
    name: string
  }
}
