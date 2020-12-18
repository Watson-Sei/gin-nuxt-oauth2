<template>
  <div class="container">
    <div class="signin">
      <h2>Sign in</h2>
      <div>
        <form @submit.prevent="submit">
          <label for="usernameTxt">Username:</label>
          <input id="usernameTxt" type="email" v-model="email">
          <label for="passwordTxt">Password:</label>
          <input id="passwordTxt" type="password" v-model="password">
          <button type="submit">Sign In</button>
        </form>
        <button class="button" @click.prevent="fbGoogleLogin">Google Login</button>
        <button class="button" @click.prevent="fbGoogleLogout">Google Logout</button>
      </div>
      <p>You don't have an account?
        <nuxt-link to="/auth/signup">create account now!!</nuxt-link>
      </p>
    </div>
  </div>
</template>

<script>
import { mapActions, mapGetters } from 'vuex'
import firebase, {googleProvider} from '~/firebase/app'
export default {
  name: "signin",
  data() {
    return {
      email: '',
      password: '',
    }
  },
  computed: {
    ...mapGetters('modules/user', [
      'uid'
    ])
  },
  methods: {
    ...mapActions("modules/user", [ 'login' ]),
    async submit () {
      firebase.auth().signInWithEmailAndPassword(this.email, this.password).then((firebaseUser) => {
        return this.login(firebaseUser.user)
      }).then(() => {
        this.$router.push('/')
      }).catch((error) => {
        console.log(error.message)
      })
    },
    // Googleアカウントログイン
    async fbGoogleLogin() {
      const { user } = await firebase.auth().signInWithPopup(googleProvider);
      await this.login(user);
      await this.$router.push('/')
    },
    async fbGoogleLogout() {
      await this.logout();
      await this.$router.push('/')
    }
  }
}
</script>

<style scoped>

</style>
