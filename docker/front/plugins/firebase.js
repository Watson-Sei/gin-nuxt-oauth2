import firebase from "firebase";

if (!firebase.apps.length) {
  firebase.initializeApp(
    {
      apiKey: process.env.apiKey,
      authDomain: process.env.authDomain,
      databaseURL: process.env.databaseURL,
      storageBucket: process.env.storageBucket,
      messagingSenderId: process.env.messagingSenderId,
      appId: process.env.appId,
    }
  )
}

export default firebase
