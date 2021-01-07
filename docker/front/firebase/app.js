import firebase from "firebase";

if (!firebase.apps.length) {
  firebase.initializeApp(
    {
      apiKey: process.env.apiKey,
      authDomain: `${process.env.projectId}.firebaseapp.com`,
      databaseURL: `https://${process.env.databaseURL}.firebaseio.com`,
      projectId: process.env.projectId,
      storageBucket: `${process.env.projectId}.appspot.com`,
      messagingSenderId: process.env.messagingSenderId,
      appId: process.env.appId,
      measurementId: process.env.measurementId
    }
  )
}

const googleProvider = new firebase.auth.GoogleAuthProvider()

export default firebase
export {googleProvider}
