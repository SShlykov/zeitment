import axios from "axios";

const post = async (path, data, onError) => {
  return axios.post(`${path}`, data)
    .then(r => r.data)
    .catch(error => {
      onError && onError(error)
      console.error('Error post request:', error)
    })
}

const get = async (path) => {
  return axios.get(`${path}`)
    .then(r => r.data)
    .catch(error => {
      console.error('Error post request:', error)
    })
}

const put = async (path, data, onError) => {
  return axios.put(`${path}`, data)
    .then(r => r.data)
    .catch(error => {
      onError && onError(error)
      console.error('Error put request:', error)
    })
}

const remove = async (path, onError) => {
  return axios.delete(`${path}`)
    .then(r => r.data)
    .catch(error => {
      onError && onError(error)
      console.error('Error delete request:', error)
    })
}


export {get, post, put, remove}
