<template>
  <div class="container">
    <div class="signup">
      <h2>Sign up</h2>
      <input type="text" placeholder="Username" v-model="email">
      <input type="password" placeholder="Password" v-model="password">
      <button @click="signUp">Register</button>
      <p>Do you have an account?
        <nuxt-link to="/auth/signin">sign in now!!</nuxt-link>
      </p>
    </div>
  </div>
</template>

<script>
import firebase from '~/plugins/firebase'
export default {
  name: "signup",
  data() {
    return {
      email: '',
      password: '',
    }
  },
  methods: {
    signUp: function () {
      firebase.auth().createUserWithEmailAndPassword(this.email, this.password)
      .then(res => {
        console.log('Create account: ', res.user.email)
        firebase.auth().signInWithEmailAndPassword(this.email, this.password)
        .then(r => {
          r.user.getIdToken().then(idToken => {
            localStorage.setItem('jwt', idToken.toString())
          })
          this.$router.push('/')
        })
      }).catch(error => {
        alert(error.message)
      })
    }
  }
}
</script>

<style scoped>

</style>
