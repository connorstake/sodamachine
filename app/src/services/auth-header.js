export function authHeader() {
    const token = localStorage.getItem('token');
    console.log(token)
    if (token) {
      return { Authorization: token };
    } else {
      return {};
    }
  }
  