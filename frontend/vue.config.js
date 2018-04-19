module.exports = {
  lintOnSave: false,
  devServer: {
    proxy: {
      '/projects': {
        target: 'http://localhost:3000/'
      },
      '/login': {
        target: 'http://localhost:3000/'
      }
    }
  }
};
