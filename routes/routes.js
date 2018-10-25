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
        routes.addItem(req, res, function_name);
    });

    app.get('/queryFood', function (req, res) {
        routes.queryFood(req, res);
    });

    app.get('/food/:id', function (req, res) {
        var function_name = 'getFoodProInfo';
        routes.queryItem(req, res, function_name);
    });
}
