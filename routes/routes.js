var routes = require('./controller.js');

module.exports = function(app) {
    app.get('/', function(req, res) {
        routes.index(req, res);
    });

}
