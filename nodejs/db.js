/**
 * Libraries
 */
const mongodb = require('mongodb');

/**
 * Configuration
 */
const config = require('./config');

// database connection
let _db;

/**
 * Connect to database
 * @description This function creates the database connection and stores it in the _db variable
 * @param {Function} callback Callback function
 */
module.exports.connect = (callback) => {
    if (_db) {
        console.warn('Database connection already exists.');
        return callback(null);
    }
    
    // try to connect to the database
    mongodb.MongoClient.connect(config.databaseConnection.string, config.databaseConnection.options, (err, db) => {
        if (err) {
            return callback(err);
        }
    
        console.log('Connected to database');
        _db = db;
    
        return callback(null);
    });
};

/**
 * Get database connection
 * @description This function returns connection created in connect() and stored in _db
 */
module.exports.get = (name) => {
    if (!_db) {
        return console.warn('Database connection was not created.');
    }
    return _db.db(name);
}
