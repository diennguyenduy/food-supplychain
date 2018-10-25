var routes = require('./controller.js');

module.exports = function(app) {
    app.get('/', function(req, res) {
        routes.index(req, res);
    });

    app.get('/addFood', function (req, res) {
        routes.addFood(req, res);
    });

    app.post('/addFood', function (req, res) {
        var function_name = 'addFoodProInfo';
        routes.additem(req, res, function_name);
    });

    app.get('/')
}
