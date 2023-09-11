<script setup lang="ts">
import { reactive, ref } from "vue";
import { useI18n } from "vue-i18n";
import { useLoading } from "../store/loading";
// import {
//   GetSuperUserAccount,
//   CheckSuperUserPassword,
// } from "../../wailsjs/go/local/service";
import { ResponseState } from "../../../components/types";

const { startLoading, endLoading } = useLoading();

const { t } = useI18n();

const superUserAccount = ref("");
// GetSuperUserAccount().then((account: string) => {
//   superUserAccount.value = account;
// });

const auth = reactive({
  isAuth: false,
  enteredPassword: "",
  usedPassword: "",
});
const checkPassword = async (res: (state: ResponseState) => void) => {
  startLoading();
  // const success = await CheckSuperUserPassword(auth.enteredPassword);
  // if (success) {
  //   auth.isAuth = true;
  //   auth.usedPassword = auth.enteredPassword;
  //   res("success");
  // } else {
  //   res("error");
  // }
  endLoading();
};

const search = ref("");
const cards = ref<
  {
    ID: number;
  }[]
>([]);
</script>

<template>
  <my-container class="keyring">
    <my-main v-if="auth.isAuth">
      <my-form>
        <my-input
          class="input--centered"
          width="20em"
          v-model="search"
          :placeholder="t('keyring.search')"
        ></my-input>
      </my-form>
      <my-container>
        <div class="card" v-for="card in cards" :key="card.ID"></div>
      </my-container>
    </my-main>
    <my-main v-else>
      <my-form>
        <my-input
          class="input--centered"
          type="password"
          width="20em"
          v-model="auth.enteredPassword"
          :placeholder="
            t('keyring.auth.password', { account: superUserAccount })
          "
        >
          <template #append>
            <my-button type="primary" @click="checkPassword">
              {{ t("keyring.auth.confirm") }}
            </my-button>
          </template>
        </my-input>
      </my-form>
    </my-main>
  </my-container>
</template>

<style lang="scss" scoped>
.keyring {
  width: 100%;
  max-width: 1024px;
}

.input--centered {
  justify-content: center;
}
</style>
