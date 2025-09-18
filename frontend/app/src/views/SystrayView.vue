<script setup lang="ts">
import { ref } from 'vue'
import { UserService } from '@jinyaoma/my-sdk/bindings/majinyao.cn/my-app/backend/cmd/wails/services'

const searchParams = new URLSearchParams(window.location.search)
const hasReservedUsersOnStartup = ref(searchParams.get('hasReservedUsersOnStartup') === 'true')

const adminName = ref('')
const adminAccount = ref('')
const adminPassword = ref('')

if (hasReservedUsersOnStartup.value) {
  UserService.GetReservedUserInfo().then((user) => {
    if (user) {
      adminName.value = user.Name
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
      hasReservedUsersOnStartup.value = true
    }
  }
}
</script>

<template>
  <div class="systray">
    <div v-if="hasReservedUsersOnStartup">
      <h1>Load Admin</h1>
      <div>{{ adminName }}</div>
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
