const PostHeader = {
  headers: {
    'Content-Type': 'application/x-www-form-urlencoded; charset=UTF-8'
  },
  transformRequest: [
    data => {
      const pairs = []

      Object.keys(data).forEach(key => {
        pairs.push(`${key}=${data[key]}`)
      })

      return pairs.join('&')
    }
  ]
}

export default PostHeader
