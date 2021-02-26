<template>
  <v-app class="main">
    <v-main class="flex justify-center align-center">
      <v-card width="600" height="480" class="center-posi">
        <v-form class="form">
          <v-text-field v-model="name" :disabled="!isEditable" outlined type="text" class="mb-1" label="通知する名前"
                        hint="Slackの通知名に使われます"></v-text-field>
          <v-text-field v-model="url" :disabled="!isEditable" outlined type="url" class="mb-1"
                        label="通知するSlackの Webhook URL"></v-text-field>
          <v-text-field v-model="confirmationInterval" :disabled="!isEditable" outlined type="number" class="mb-1"
                        label="バッテリーを確認する間隔(秒)" hint=""></v-text-field>
          <v-text-field v-model="notificationCondition" :disabled="!isEditable" outlined type="number" class="mb-1"
                        label="何％以下になったら通知するか"></v-text-field>

          <v-btn :disabled="isConfirming" class="float-left" color="success" @click="saveSettings">
            {{ isEditable ? "保存する" : "編集する" }}
          </v-btn>
          <v-btn :disabled="isEditable" class="float-right" color="success"
                 @click="isConfirming ? stopBatteryConfirmation() : startBatteryConfirmation()">{{
              isConfirming ? "監視ストップする！！" : "監視スタート！！"
            }}
          </v-btn>
        </v-form>
      </v-card>
      <p class="pt-2 text">バッテリー残量を定期的に確認して、指定した％以下になればSlackに通知してくれるアプリ。使い方は設定項目を入力して、監視スタート！！を押すだけ。</p>
      <p class="text-center pt-2">Slackの Webhook URL の取得は
        <a href="https://media-radar.jp/contents/meditsubu/slack_incoming_webhook/">こちらを参照</a>
      </p>
      <h2 v-if="errorMessage.length > 0" class="pt-2">{{ errorMessage }}</h2>
    </v-main>

    <v-footer app class="transparent">
      <v-col class="text-center">
        © バッテリー残量通知くん {{ new Date().getFullYear() }} All rights reserve. <br>
        <a href="http://www.onlinewebfonts.com">oNline Web Fonts</a>
      </v-col>
    </v-footer>
  </v-app>
</template>

<script>
export default {
  data() {
    return {
      isEditable: false,
      isConfirming: false,
      name: "",
      url: "",
      confirmationInterval: 60,
      notificationCondition: 30,
      settings: {},
      errorMessage: "",
    }
  },
  mounted() {
    window.backend.Settings.Load()
        .then(settings => {
          try {
            this.settings = JSON.parse(settings)
            this.name = this.settings.name
            this.url = this.settings.webhookUrl
            this.confirmationInterval = this.settings.confirmationInterval
            this.notificationCondition = this.settings.notificationCondition
          } catch (e) {
            this.errorMessage = "Unable to load todo settings"
            setTimeout(() => {
              this.errorMessage = ""
            }, 3000)
          }
        })
        .catch(error => {
          this.errorMessage = error
          setTimeout(() => {
            this.errorMessage = ""
          }, 3000)
        })
  },
  watch: {
    settings: {
      handler: function (settings) {
        window.backend.Settings.Save(JSON.stringify(settings))
      },
      deep: true
    }
  },
  methods: {
    saveSettings() {
      if (this.isEditable) {
        let name = this.name && this.name.trim()
        let url = this.url && this.url.trim()
        let confirmationInterval = this.confirmationInterval && this.confirmationInterval
        let notificationCondition = this.notificationCondition && this.notificationCondition
        if (name === "" || url === "") return
        this.settings = {
          name: name,
          webhookUrl: url,
          confirmationInterval: confirmationInterval,
          notificationCondition: notificationCondition
        }
      }
      this.isEditable = !this.isEditable
    },
    startBatteryConfirmation() {
      this.isConfirming = true
      window.backend.ConfirmBattery.Start(this.name, this.url,
          Number(this.confirmationInterval), Number(this.notificationCondition))
    },
    stopBatteryConfirmation() {
      this.isConfirming = false
      window.backend.ConfirmBattery.Stop()
    }
  }
}
</script>

<style scoped>
.main {
  background: rgb(0, 255, 199);
  background: linear-gradient(90deg, rgba(0, 255, 199, 0.68531162464986) 0%, rgba(7, 255, 0, 0.8169642857142857) 100%);
}

.center-posi {
  margin: 0 auto;
}

.form {
  max-width: 80%;
  margin: 0 auto;
  padding-top: 32px;
}

h2 {
  text-align: center;
  color: white;
  background-color: red;
  min-width: 230px;
  max-width: 550px;
  padding: 1rem;
  border-radius: 0.5rem;
  margin: 40px auto;

}

.text {
  max-width: 600px;
  margin: 0 auto;
}
</style>
