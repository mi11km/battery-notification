<template>
  <v-app class="main">
    <v-main class="flex justify-center align-center">
            <v-col class="text-center text-h3 green--text">バッテリー残量通知くん</v-col>
      <v-card width="600" height="480" class="center-posi">
        <v-form class="form">
          <v-text-field outlined type="text" class="mb-1" label="通知する名前" hint="Slackの通知名に使われます"></v-text-field>
          <v-text-field outlined type="url" class="mb-1" label="通知するSlackの Webhook URL"></v-text-field>
          <v-text-field outlined type="number" class="mb-1" label="バッテリーを確認する間隔(秒）"></v-text-field>
          <v-text-field outlined type="number" class="mb-1" label="何％以下になったら通知するか"></v-text-field>

          <v-btn class="float-left" color="success">編集する</v-btn>
          <v-btn class="float-right" color="success">オン</v-btn>
        </v-form>
      </v-card>
    </v-main>

    <v-footer app class="transparent">
      <v-col class="text-center">
        © バッテリー残量通知くん {{ new Date().getFullYear() }} All rights reserve.
      </v-col>
    </v-footer>
  </v-app>
</template>

<script>
export default {
  data() {
    return {
      isEditable: true,
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
</style>
