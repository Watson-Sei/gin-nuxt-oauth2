<template>
  <div class="container">
    <div class="hello">
      <h1>{{ msg }}</h1>
      <h2>Essential Links</h2>
      <button @click="signOut">Sign out</button>
      <button @click="apiPublic">public</button>
      <button @click="apiPrivate">private</button>
    </div>
  </div>
</template>
<script>
import firebase from '~/plugins/firebase'
export default {
  middleware: 'authenticated',
  data() {
    return {
      msg: 'Welcome to Your Vue.js App',
      name: firebase.auth().currentUser.email
    }
  },
  methods: {
    signOut() {
      firebase.auth().signOut().then(() => {
        localStorage.removeItem('jwt')
        this.$router.push('/auth/signin')
      })
    },
    apiPublic() {
      const response = this.$axios.$get('http://localhost/api/public')
      .then(function (response) {
        console.log(response.msg);
      });
    },
    apiPrivate() {
      const response = this.$axios.$get('http://localhost/api/private', {
        headers: {'Authorization': `Bearer ${localStorage.getItem('jwt')}`}
      })
      .then(function (response) {
        console.log(response.msg)
      })
    }
  }
}
</script>

<style scoped>
h1, h2 {
  font-weight: normal;
}
ul {
  list-style-type: none;
  padding: 0;
}
li {
  display: inline-block;
  margin: 0 10px;
}
a {
  color: #42b983;
}
button {
  margin: 10px 0;
  padding: 10px;
}
</style>
