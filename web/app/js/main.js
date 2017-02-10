ng = require("angular");
router = require("angular-ui-router");

app = ng.module("StockAnalysis",["ui-router"])

app.component("sectorsView", require("./components/sectorsView.component.js"));

app = require('./routes.js')(app)

ng.bootstrap('body',['stockAnalysisView']);
