class ServiceOfLayout {

  /**
   *
   * @param {Object} store
   */
  constructor(store) {
    this.store = store;
  }

  /**
   *
   * @param {Object} notification
   */
  addNotification(notification) {
    const id = Math.random().toString(36).slice(-4)
    const timer = notification.timer || 1000
    const newNotification = {
      ...notification,
      id,
    }
    this.store.dispatch('layout/addNotification', newNotification)
    setTimeout(() => {
      this.store.dispatch('layout/removeNotification', id)
    }, timer)
  }
}

export default ServiceOfLayout;