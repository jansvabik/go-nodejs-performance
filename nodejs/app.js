/**
 * Entry point of the Node.js app for testing Go vs. Node.js REST API performance
 * @author Jan Svabik
 * @see https://github.com/jansvabik/go-nodejs-performance
 */

// load express, config and db functions
const app = require('express')();
const bodyParser = require('body-parser');
const config = require('./config');
const db = require('./db');

// parse json from request body
app.use(bodyParser.json())

// connect to the database
db.connect((err) => {
    if (err) {
        console.error(err);
        process.exit(1);
    }

    // use one router for all paths
    const router = require('./router');
    app.use('/', router);

    // listen to requests
    app.listen(config.port, (err) => {
        if (err) {
            console.error(err);
            process.exit(1);
        }

        console.log(`Listening on port ${config.port}.`);
    })
});