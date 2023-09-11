<script setup lang="ts">
import { reactive, ref } from "vue";
import { useI18n } from "vue-i18n";
import { useLoading } from "../store/loading";
import { useColorTheme } from "../store/color-theme";
// import {
//   GetOptions,
//   ChooseLogFile,
//   ChooseDirectory,
//   SavePortHttp,
//   SavePortHttps,
//   SaveAutoStart,
//   GetSuperUserAccount,
//   UpdateSuperUserPassword,
// } from "../../wailsjs/go/local/service";
// import {
//   ChangeDisplayLanguage,
//   ChangeColorTheme,
//   ChangeWebServiceState,
// } from "../../wailsjs/go/tray/tray";
import { ResponseState } from "../../../components/types";

const { startLoading, endLoading } = useLoading();

const { t, locale, availableLocales } = useI18n();
const colorTheme = useColorTheme();

const options = reactive({
  LogFile: "",
  DirLanguages: "",
  DirAssets: "",
  DirUserData: "",
  DirDocs: "",
  WebAutoStart: false,
  WebPortHttp: 80,
  WebPortHttps: 443,
  WebDirCerts: "",
  SuperUserOldPassword: "",
  SuperUserNewPassword: "",
});
// GetOptions().then((config) => {
//   options.LogFile = config.LogFile;
//   options.DirLanguages = config.DirLanguages;
//   options.DirAssets = config.DirAssets;
//   options.DirUserData = config.DirUserData;
//   options.DirDocs = config.DirDocs;
//   options.WebAutoStart = config.WebAutoStart == "true";
//   options.WebPortHttp = parseInt(config.WebPortHttp.substring(1));
//   options.WebPortHttps = parseInt(config.WebPortHttps.substring(1));
//   options.WebDirCerts = config.WebDirCerts;
// });

const superUserAccount = ref("");
// GetSuperUserAccount().then((account: string) => {
//   superUserAccount.value = account;
// });

const changeDisplayLanguage = async (newLang: string) => {
  startLoading();
  // await ChangeDisplayLanguage(newLang);
};
const changeColorTheme = async (newTheme: string) => {
  startLoading();
  // await ChangeColorTheme(newTheme);
};

const changeLogFile = async (res: (state: ResponseState) => void) => {
  startLoading();
  // const newLogFile = await ChooseLogFile(
  //   options.LogFile,
  //   t("settings.ChooseLogFile")
  // );
  // if (newLogFile) {
  //   options.LogFile = newLogFile;
  //   res("success");
  // } else {
  //   res("warning");
  // }
  endLoading();
};
const changeDirCerts = async (res: (state: ResponseState) => void) => {
  startLoading();
  // const newDir = await ChooseDirectory(
  //   "certs",
  //   options.WebDirCerts,
  //   t("settings.chooseDirCerts")
  // );
  // if (newDir) {
  //   options.WebDirCerts = newDir;
  //   res("success");
  // } else {
  //   res("warning");
  // }
  endLoading();
};
const changeDirLanguages = async (res: (state: ResponseState) => void) => {
  startLoading();
  // const newDir = await ChooseDirectory(
  //   "languages",
  //   options.DirLanguages,
  //   t("settings.chooseDirLanguages")
  // );
  // if (newDir) {
  //   options.DirLanguages = newDir;
  //   res("success");
  // } else {
  //   res("warning");
  // }
  endLoading();
};
const changeDirAssets = async (res: (state: ResponseState) => void) => {
  startLoading();
  // const newDir = await ChooseDirectory(
  //   "assets",
  //   options.DirAssets,
  //   t("settings.chooseDirAssets")
  // );
  // if (newDir) {
  //   options.DirAssets = newDir;
  //   res("success");
  // } else {
  //   res("warning");
  // }
  endLoading();
};
const changeDirUserData = async (res: (state: ResponseState) => void) => {
  startLoading();
  // const newDir = await ChooseDirectory(
  //   "userdata",
  //   options.DirUserData,
  //   t("settings.chooseDirUserData")
  // );
  // if (newDir) {
  //   options.DirUserData = newDir;
  //   res("success");
  // } else {
  //   res("warning");
  // }
  endLoading();
};
const changeDirDocs = async (res: (state: ResponseState) => void) => {
  startLoading();
  // const newDir = await ChooseDirectory(
  //   "docs",
  //   options.DirDocs,
  //   t("settings.chooseDirDocs")
  // );
  // if (newDir) {
  //   options.DirDocs = newDir;
  //   res("success");
  // } else {
  //   res("warning");
  // }
  endLoading();
};
const changePortHttp = async (res: (state: ResponseState) => void) => {
  if (options.WebPortHttp) {
    startLoading();
    // const success = await SavePortHttp(options.WebPortHttp);
    // if (success) {
    //   res("success");
    // } else {
    //   res("error");
    // }
    endLoading();
  }
};
const changePortHttps = async (res: (state: ResponseState) => void) => {
  if (options.WebPortHttps) {
    startLoading();
    // const success = await SavePortHttps(options.WebPortHttps);
    // if (success) {
    //   res("success");
    // } else {
    //   res("error");
    // }
    endLoading();
  }
};
const changeAutoStart = async () => {
  startLoading();
  // const success = await SaveAutoStart(options.WebAutoStart);
  // if (success) {
  //   ChangeWebServiceState(true);
  // } else {
  //   options.WebAutoStart = !options.WebAutoStart;
  // }
  endLoading();
};
const clearSuperUserPassword = () => {
  options.SuperUserOldPassword = "";
  options.SuperUserNewPassword = "";
};
const changeSuperUserPassword = async (res: (state: ResponseState) => void) => {
  startLoading();
  // const success = await UpdateSuperUserPassword(
  //   options.SuperUserOldPassword,
  //   options.SuperUserNewPassword
  // );
  // if (success) {
  //   res("success");
  // } else {
  //   res("error");
  // }
  endLoading();
};
</script>

<template>
  <my-container class="settings">
    <my-main>
      <h2>{{ t("menu.main.settings") }}</h2>
      <my-form>
        <my-form-group
          :legend="t('settings.general.legend')"
          :label-width="locale === 'zh' ? '7em' : '9em'"
        >
          <my-form-item :label="t('settings.general.displayLanguage')">
            <my-select
              name="DisplayLanguage"
              :display-value="locale"
              @change="changeDisplayLanguage"
            >
              <my-option
                v-for="lang in availableLocales"
                :key="lang"
                :value="lang"
              >
                {{ t(`lang.${lang}`) }}
              </my-option>
            </my-select>
          </my-form-item>
          <my-form-item :label="t('settings.general.colorTheme')">
            <my-select
              name="ColorTheme"
              :display-value="colorTheme.theme"
              @change="changeColorTheme"
            >
              <my-option value="system">
                {{ t(`theme.system`) }}
              </my-option>
              <my-option value="light">
                {{ t(`theme.light`) }}
              </my-option>
              <my-option value="dark">
                {{ t(`theme.dark`) }}
              </my-option>
            </my-select>
          </my-form-item>
          <my-form-item :label="t('settings.general.LogFile')">
            <my-input
              name="LogFile"
              width="100%"
              v-model="options.LogFile"
              disabled
            >
              <template #append>
                <my-button type="primary" @click="changeLogFile">
                  {{ t("settings.choosePath") }}
                </my-button>
              </template>
            </my-input>
          </my-form-item>
        </my-form-group>
        <my-form-group
          :legend="t('settings.webService.legend')"
          :label-width="locale === 'zh' ? '7em' : '10em'"
        >
          <my-form-item :label="t('settings.webService.autoStart')">
            <my-input
              type="checkbox"
              name="WebAutoStart"
              v-model="options.WebAutoStart"
              @change="changeAutoStart"
            >
            </my-input>
          </my-form-item>
          <my-form-item :label="t('settings.webService.portHttp')">
            <my-input
              type="number"
              name="WebPortHttp"
              width="6.5em"
              :min="0"
              :max="65536"
              v-model="options.WebPortHttp"
            >
              <template #prepend>:</template>
              <template #append>
                <my-button type="primary" @click="changePortHttp">
                  {{ t("settings.save") }}
                  {{ t("settings.webService.portHttp") }}
                </my-button>
              </template>
            </my-input>
          </my-form-item>
          <my-form-item :label="t('settings.webService.portHttps')">
            <my-input
              type="number"
              name="WebPortHttps"
              width="6.5em"
              :min="0"
              :max="65536"
              v-model="options.WebPortHttps"
            >
              <template #prepend>:</template>
              <template #append>
                <my-button type="primary" @click="changePortHttps">
                  {{ t("settings.save") }}
                  {{ t("settings.webService.portHttps") }}
                </my-button>
              </template>
            </my-input>
          </my-form-item>
          <my-form-item :label="t('settings.webService.dirCerts')">
            <my-input
              name="WebDirCerts"
              width="100%"
              v-model="options.WebDirCerts"
              :placeholder="t('settings.webService.dirCertsPlaceholder')"
              disabled
            >
              <template #append>
                <my-button type="primary" @click="changeDirCerts">
                  {{ t("settings.choosePath") }}
                </my-button>
              </template>
            </my-input>
          </my-form-item>
        </my-form-group>
        <my-form-group
          :legend="t('settings.otherPaths.legend')"
          label-width="8em"
        >
          <my-form-item :label="t('settings.otherPaths.dirLanguages')">
            <my-input
              name="DirLanguages"
              width="100%"
              v-model="options.DirLanguages"
              disabled
            >
              <template #append>
                <my-button type="primary" @click="changeDirLanguages">
                  {{ t("settings.choosePath") }}
                </my-button>
              </template>
            </my-input>
          </my-form-item>
          <my-form-item :label="t('settings.otherPaths.dirAssets')">
            <my-input
              name="DirAssets"
              width="100%"
              v-model="options.DirAssets"
              disabled
            >
              <template #append>
                <my-button type="primary" @click="changeDirAssets">
                  {{ t("settings.choosePath") }}
                </my-button>
              </template>
            </my-input>
          </my-form-item>
          <my-form-item :label="t('settings.otherPaths.dirUserData')">
            <my-input
              name="DirUserData"
              width="100%"
              v-model="options.DirUserData"
              disabled
            >
              <template #append>
                <my-button type="primary" @click="changeDirUserData">
                  {{ t("settings.choosePath") }}
                </my-button>
              </template>
            </my-input>
          </my-form-item>
          <my-form-item :label="t('settings.otherPaths.dirDocs')">
            <my-input
              name="DirDocs"
              width="100%"
              v-model="options.DirDocs"
              disabled
            >
              <template #append>
                <my-button type="primary" @click="changeDirDocs">
                  {{ t("settings.choosePath") }}
                </my-button>
              </template>
            </my-input>
          </my-form-item>
        </my-form-group>
        <my-form-group
          :legend="
            t('settings.changeSuperUserPassword.legend', {
              account: superUserAccount,
            })
          "
          :label-width="locale === 'zh' ? '5em' : '8em'"
        >
          <my-form-item
            :label="t('settings.changeSuperUserPassword.oldPassword')"
          >
            <my-input
              type="password"
              width="20em"
              v-model="options.SuperUserOldPassword"
              :placeholder="t('settings.changeSuperUserPassword.oldPassword')"
            >
            </my-input>
          </my-form-item>
          <my-form-item
            :label="t('settings.changeSuperUserPassword.newPassword')"
          >
            <my-input
              type="password"
              width="20em"
              v-model="options.SuperUserNewPassword"
              :placeholder="t('settings.changeSuperUserPassword.newPassword')"
            >
            </my-input>
          </my-form-item>
          <my-form-item>
            <my-button type="primary" @click="changeSuperUserPassword">
              {{ t("settings.changeSuperUserPassword.change") }}
            </my-button>
            <my-button @click="clearSuperUserPassword">
              {{ t("settings.changeSuperUserPassword.clear") }}
            </my-button>
          </my-form-item>
        </my-form-group>
      </my-form>
      <p>{{ t("settings.footnote") }}</p>
    </my-main>
  </my-container>
</template>

<style lang="scss" scoped>
.settings {
  width: 100%;
  max-width: 580px;
}
h2 {
  font-size: 2em;
  margin-top: 0;
}
p {
  font-size: 0.85em;
  color: var(--my-color-danger-1);
}
</style>
