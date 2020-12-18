import firebase from '~/plugins/firebase'

export default function ({ redirect }) {
  let currentUser = firebase.auth().currentUser
  if (!currentUser) {
    return redirect('/auth/signin')
  }
}
