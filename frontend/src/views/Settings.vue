<script setup lang="ts">
import { reactive, ref } from "vue";
import { useI18n } from "vue-i18n";
import { useLoading } from "../store/loading";
import { useColorTheme } from "../store/color-theme";
import {
  GetOptions,
  ChooseLogPath,
  ChooseDirCerts,
  SavePortHttp,
  SavePortHttps,
  SaveAutoStart,
  GetSuperUserAccount,
  UpdateSuperUserPassword,
} from "../../wailsjs/go/service/service";
import {
  ChangeDisplayLanguage,
  ChangeColorTheme,
  IsWebServiceRunning,
  StartWebService,
} from "../../wailsjs/go/tray/tray";
import { ResponseState } from "../../packages/components/types";

const { startLoading, endLoading } = useLoading();

const { t, locale, availableLocales } = useI18n();
const colorTheme = useColorTheme();

const options = reactive({
  LogPath: "",
  AutoStart: false,
  PortHttp: 80,
  PortHttps: 443,
  DirCerts: "",
  SuperUserOldPassword: "",
  SuperUserNewPassword: "",
});
GetOptions().then((config: config.Config) => {
  options.LogPath = config.LogPath;
  options.AutoStart = config.AutoStart == "true";
  options.PortHttp = parseInt(config.PortHttp.substring(1));
  options.PortHttps = parseInt(config.PortHttps.substring(1));
  options.DirCerts = config.DirCerts;
});

const superUserAccount = ref("");
GetSuperUserAccount().then((account: string) => {
  superUserAccount.value = account;
});

const changeDisplayLanguage = async (newLang: string) => {
  startLoading();
  await ChangeDisplayLanguage(newLang);
};
const changeColorTheme = async (newTheme: string) => {
  startLoading();
  await ChangeColorTheme(newTheme);
};

const changeLogPath = async (res: (state: ResponseState) => void) => {
  startLoading();
  const newLogPath = await ChooseLogPath(
    options.LogPath,
    t("settings.chooseLogPath")
  );
  if (newLogPath) {
    options.LogPath = newLogPath;
    res("success");
  } else {
    res("warning");
  }
  endLoading();
};
const changeDirCerts = async (res: (state: ResponseState) => void) => {
  startLoading();
  const newDirCerts = await ChooseDirCerts(
    options.DirCerts,
    t("settings.chooseDirCerts")
  );
  if (newDirCerts) {
    options.DirCerts = newDirCerts;
    res("success");
  } else {
    res("warning");
  }
  endLoading();
};
const changePortHttp = async (res: (state: ResponseState) => void) => {
  if (options.PortHttp) {
    startLoading();
    const success = await SavePortHttp(options.PortHttp);
    if (success) {
      res("success");
    } else {
      res("error");
    }
    endLoading();
  }
};
const changePortHttps = async (res: (state: ResponseState) => void) => {
  if (options.PortHttps) {
    startLoading();
    const success = await SavePortHttps(options.PortHttps);
    if (success) {
      res("success");
    } else {
      res("error");
    }
    endLoading();
  }
};
const changeAutoStart = async () => {
  startLoading();
  const success = await SaveAutoStart(options.AutoStart);
  if (success) {
    !(await IsWebServiceRunning()) && StartWebService();
  } else {
    options.AutoStart = !options.AutoStart;
  }
  endLoading();
};
const clearSuperUserPassword = () => {
  options.SuperUserOldPassword = "";
  options.SuperUserNewPassword = "";
};
const changeSuperUserPassword = async (res: (state: ResponseState) => void) => {
  startLoading();
  const success = await UpdateSuperUserPassword(
    options.SuperUserOldPassword,
    options.SuperUserNewPassword
  );
  if (success) {
    res("success");
  } else {
    res("error");
  }
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
          <my-form-item :label="t('settings.general.logPath')">
            <my-input
              name="LogPath"
              width="100%"
              v-model="options.LogPath"
              disabled
            >
              <template #append>
                <my-button type="primary" @click="changeLogPath">
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
              name="AutoStart"
              v-model="options.AutoStart"
              @change="changeAutoStart"
            >
            </my-input>
          </my-form-item>
          <my-form-item :label="t('settings.webService.portHttp')">
            <my-input
              type="number"
              name="Web.PortHttp"
              width="6.5em"
              :min="0"
              :max="65536"
              v-model="options.PortHttp"
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
              name="Web.PortHttps"
              width="6.5em"
              :min="0"
              :max="65536"
              v-model="options.PortHttps"
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
              name="Web.DirCerts"
              width="100%"
              v-model="options.DirCerts"
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
