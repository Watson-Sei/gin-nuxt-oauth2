<template>
  <div class="container">
    <div class="signin">
      <h2>Sign in</h2>
      <input type="text" placeholder="email" v-model="email">
      <input type="password" placeholder="password" v-model="password">
      <button @click="signIn">Signin</button>
      <p>You don't have an account?
        <nuxt-link to="/auth/signup">create account now!!</nuxt-link>
      </p>
    </div>
  </div>
</template>

<script>
import firebase from '~/plugins/firebase'
export default {
  name: "signin",
  data() {
    return {
      email: '',
      password: '',
    }
  },
  methods: {
    signIn: function () {
      firebase.auth().signInWithEmailAndPassword(this.email, this.password)
      .then(res => {
        res.user.getIdToken().then(idToken => {
          localStorage.setItem('jwt', idToken.toString())
        })
        this.$router.push('/')
      }).catch(error => {
        alert(error.message)
      })
    }
  }
}
</script>

<style scoped>

</style>
