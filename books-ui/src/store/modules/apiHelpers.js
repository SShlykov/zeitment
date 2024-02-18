import axios from "axios";

const post = async (path, data, onError) => {
  return axios.post(`${import.meta.env.VITE_API_ADDR}${path}`, data)
  .then(r => r.data)
  .catch(error => {
    onError && onError(error)
    console.error('Error post request:', error)
    authViewer(error)
  })
}

const get = async (path) => {
  return axios.get(`${import.meta.env.VITE_API_ADDR}${path}`)
  .catch(error => {
      console.error('Error post request:', error)
      authViewer(error)
  })
}


export {get, post}
