<script setup lang="ts">
import { ref } from 'vue'
import { Events } from '@wailsio/runtime'
import { FileService, UserService } from '@/bindings/majinyao.cn/my-app/backend/cmd/wails/services'

const searchParams = new URLSearchParams(window.location.search)
const hasReservedUsersOnStartup = ref(searchParams.get('hasReservedUsersOnStartup') === 'true')

const adminId = ref('')
const adminName = ref('')
const adminAccount = ref('')
const adminPassword = ref('')
const adminAvatar = ref('')

if (hasReservedUsersOnStartup.value) {
  UserService.GetReservedUserInfo().then((user) => {
    if (user) {
      adminId.value = user.Id
      adminName.value = user.Name
      adminAccount.value = user.Account
      adminAvatar.value = user.Avatar
    }
  })
}

const handleSetup = async () => {
  if (adminName.value && adminAccount.value && adminPassword.value) {
    console.log('Setup Admin', adminName.value, adminAccount.value, adminPassword.value)
    const user = await UserService.CreateReservedUser({
      Name: adminName.value,
      Account: adminAccount.value,
      Password: adminPassword.value
    })
    if (user) {
      console.log('Setup Admin Success', user)
      adminId.value = user.Id
      adminAvatar.value = user.Avatar
      hasReservedUsersOnStartup.value = true
    }
  }
}

const handleUploadAvatar = async () => {
  const avatar = await FileService.UploadAvatar(adminId.value)
  if (avatar) {
    adminAvatar.value = avatar
  }
}
</script>

<template>
  <div class="systray">
    <button @click="Events.Emit(new Events.WailsEvent('open-main-window'))">Open Main Window</button>
    <div v-if="hasReservedUsersOnStartup">
      <h1>Load Admin</h1>
      <div>{{ adminName }}</div>
      <br />
      <img v-if="adminAvatar" :src="adminAvatar" alt="Admin Avatar" width="64" height="64" />
      <button v-else @click="handleUploadAvatar">Upload Avatar</button>
    </div>
    <div v-else>
      <h1>Setup Admin</h1>
      <input type="text" placeholder="Admin Name" v-model="adminName" />
      <input type="text" placeholder="Admin Account" v-model="adminAccount" />
      <input type="text" placeholder="Admin Password" v-model="adminPassword" />
      <button @click="handleSetup">Setup</button>
    </div>
  </div>
</template>

<style scoped></style>
