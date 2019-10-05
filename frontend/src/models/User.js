import JwtDecode from 'jwt-decode'

export default class User {
  static from (token) {
    try {
      console.log('### received JWT token: ', token)
      let obj = JwtDecode(token)
      console.log('### decoded object', obj)
      return new User(obj)
    } catch (_) {
      return null
    }
  }

  constructor ({ sub, name, email }) {
    this.id = sub // eslint-disable-line camelcase
    this.name = name
    this.email = email
  }

}
