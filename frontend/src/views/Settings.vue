<script setup lang="ts">
import { useI18n } from "vue-i18n";
import { useColorTheme } from "../store/color-theme";
import { EventsOn } from "../../wailsjs/runtime";
import {
  GetOptions,
  ChooseLogPath,
  ChooseDirCerts,
  SavePortHttp,
  SavePortHttps,
  SaveAutoStart,
} from "../../wailsjs/go/service/settings";
import { ChangeLanguage, ChangeColorTheme } from "../../wailsjs/go/tray/tray";
import { ResponseState } from "../../packages/components/types";
import { reactive, ref } from "vue";

const { t, locale, availableLocales } = useI18n();
const colorTheme = useColorTheme();

const options = reactive({
  LogPath: "",
  Web: {
    AutoStart: false,
    PortHttp: "",
    PortHttps: "",
    DirCerts: "",
  },
});
GetOptions().then((config: app.Config) => {
  options.LogPath = config.LogPath;
  options.Web.AutoStart = config.Web.AutoStart == "true";
  options.Web.PortHttp = config.Web.PortHttp.substring(1);
  options.Web.PortHttps = config.Web.PortHttps.substring(1);
  options.Web.DirCerts = config.Web.DirCerts;
});

const changeDisplayLanguage = (newLang: string) => {
  ChangeLanguage(newLang);
};
const changeColorTheme = (newTheme: string) => {
  ChangeColorTheme(newTheme);
};

const changeLogPath = async (res: (state: ResponseState) => void) => {
  const newLogPath = await ChooseLogPath(
    options.LogPath,
    t("settings.chooseLogPath")
  );
  if (newLogPath === "") {
    res("warning");
    return;
  }
  options.LogPath = newLogPath;
  res("success");
};
const changeDirCerts = async (res: (state: ResponseState) => void) => {
  const newDirCerts = await ChooseDirCerts(
    options.Web.DirCerts,
    t("settings.chooseDirCerts")
  );
  if (newDirCerts === "") {
    res("warning");
    return;
  }
  options.Web.DirCerts = newDirCerts;
  res("success");
};
const changePortHttp = async (res: (state: ResponseState) => void) => {
  if (options.Web.PortHttp) {
    const newPort = ":" + options.Web.PortHttp;
    const ok = await SavePortHttp(newPort);
    if (!ok) {
      res("error");
      return;
    }
    res("success");
  }
};
const changePortHttps = async (res: (state: ResponseState) => void) => {
  if (options.Web.PortHttps) {
    const newPort = ":" + options.Web.PortHttps;
    const ok = await SavePortHttps(newPort);
    if (!ok) {
      res("error");
      return;
    }
    res("success");
  }
};
const changeAutoStart = async (res: (state: ResponseState) => void) => {
  SaveAutoStart(`${options.Web.AutoStart}`);
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
              v-model="options.Web.AutoStart"
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
              v-model="options.Web.PortHttp"
            >
              <template #prepend>:</template>
              <template #append>
                <my-button type="primary" @click="changePortHttp">
                  {{ t("settings.save") }}
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
              v-model="options.Web.PortHttps"
            >
              <template #prepend>:</template>
              <template #append>
                <my-button type="primary" @click="changePortHttps">
                  {{ t("settings.save") }}
                </my-button>
              </template>
            </my-input>
          </my-form-item>
          <my-form-item :label="t('settings.webService.dirCerts')">
            <my-input
              name="Web.DirCerts"
              width="100%"
              v-model="options.Web.DirCerts"
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
      </my-form>
      <p>{{ t("settings.footnote") }}</p>
    </my-main>
  </my-container>
</template>

<style lang="scss" scoped>
.settings {
  width: 100%;
  max-width: 560px;
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
